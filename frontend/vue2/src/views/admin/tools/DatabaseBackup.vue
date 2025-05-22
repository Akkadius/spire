<template>
  <div>
    <eq-window title="Database Backup">
      <div
        class="row"
        v-for="o in options"
        :key="o.key"
      >
        <div class="col-3 text-right mt-3">
          <eq-checkbox
            :label="o.label"
            class="d-inline-block"
            :fade-when-not-true="true"
            :true-value="true"
            :false-value="false"
            v-model="request[o.key]"
            :style="disabled[o.key] ? 'opacity: .3' : 'opacity: 1'"
            @input="checked(o)"
            :disabled="loading"
          />
        </div>
        <div class="col-9">
          <div class="text-muted mt-3">
            {{ o.desc }}
          </div>
        </div>
      </div>

      <div class="row">
        <div class="col-3 text-right mt-3"></div>
        <div class="col-9 text-left mt-3">
          <button
            :style="loading ? 'opacity: .3' : 'opacity: 1'"
            class="d-inline-block"
            :disabled="loading"
            @click="backup"
          >
            Backup
          </button>
        </div>
      </div>

      <div class="row" v-if="downloading">
        <div class="col-3 text-right mt-3"></div>
        <div class="col-9 text-left mt-3 font-weight-bold">
          Downloading file...
        </div>
      </div>

      <div class="row">
        <div class="col-12 text-left mt-3">
          <info-error-banner
            :slim="true"
            :notification="notification"
            :error="error"
            @dismiss-error="error = ''"
            @dismiss-notification="notification = ''"
            class="mb-3"
          />
        </div>
      </div>

      <eq-debug :data="request"/>
    </eq-window>

    <eq-window
      title="Backup Result"
      v-if="loading || (backupResult && Object.keys(backupResult).length > 0)">
      <app-loader :is-loading="loading"/>

      <div v-if="backupResult && Object.keys(backupResult).length > 0">
        <span class="font-weight-bold">Command</span>
        <pre v-if="backupResult.command" style="width: 100%">{{ backupResult.command }}</pre>
        <span class="font-weight-bold">Output</span>
        <pre v-if="backupResult.command" style="width: 100%">{{ backupResult.stdout }}</pre>
      </div>
    </eq-window>
  </div>
</template>

<script>
import EqWindow        from "@/components/eq-ui/EQWindow.vue";
import EqCheckbox      from "@/components/eq-ui/EQCheckbox.vue";
import EqDebug         from "@/components/eq-ui/EQDebug.vue";
import {SpireApi}      from "@/app/api/spire-api";
import InfoErrorBanner from "@/components/InfoErrorBanner.vue";
import * as path       from "path";

export default {
  name: "DatabaseBackup",
  components: { InfoErrorBanner, EqDebug, EqCheckbox, EqWindow },
  data() {
    return {
      loading: false,
      downloading: false,

      backupResult: {},

      options: [
        {
          label: "All Tables",
          key: 'dump_all_tables',
          desc: "This will backup everything in your database with no exclusions"
        },
        { label: "Content Tables", key: 'content_tables', desc: "Dump content tables" },
        { label: "Player Tables", key: 'player_tables', desc: "Dump player tables" },
        { label: "Bot Tables", key: 'bot_tables', desc: "Dump bot tables" },
        { label: "State Tables", key: 'state_tables', desc: "Dump state tables" },
        { label: "System Tables", key: 'system_tables', desc: "Dump system tables" },
        { label: "Login Tables", key: 'login_tables', desc: "Dump login tables" },
        {
          label: "Compress",
          key: 'compress',
          desc: "Compresses the dump with the available compression found in the operating system"
        },
      ],

      disabled: {},

      request: {
        dump_all_tables: true,
        content_tables: false,
        player_tables: false,
        bot_tables: false,
        state_tables: false,
        system_tables: false,
        query_serv_tables: false,
        login_tables: false,
        compress: true,
      },

      // notification / errors
      notification: "",
      error: "",
    }
  },
  destroyed() {
    window.removeEventListener('file-download-progress', this.handleFileDownloadProgress, false)
  },
  created() {
    window.addEventListener('file-download-progress', this.handleFileDownloadProgress)

    this.reset()
  },
  methods: {
    handleFileDownloadProgress: function (e) {
      if (this.awaitingDownload) {
        this.awaitingDownload = false
      }

      if (e.detail.downloadedMbytes !== this.downloadedMbytes) {
        this.fileDownloadProgress = e.detail.percent
        this.lastPercent          = e.detail.percent
        this.downloadedMbytes     = Math.round(e.detail.loaded / 1024 / 1024)
        this.downloadedBytes      = e.detail.loaded
        this.totalBytes           = e.detail.total
      }
    },

    reset() {
      // disable all but "all"
      for (const [key, value] of Object.entries(this.request)) {
        if (!["dump_all_tables", "compress"].includes(key)) {
          this.disabled[key] = 1
          this.request[key]  = false
        }
      }
    },

    checked(e) {
      if (e.key === "compress") {
        return
      }

      this.disabled = {}

      // if we selected dump all again, reset everything else
      if (e.key === "dump_all_tables" && this.request[e.key]) {
        for (const [key, value] of Object.entries(this.request)) {
          if (!["dump_all_tables", "compress"].includes(key)) {
            this.disabled[key] = 1
            this.request[key]  = false
          }
          // console.log(`${key}: ${value}`);
        }
        return
      }

      let otherOptionsSet = false
      for (const [key, value] of Object.entries(this.request)) {
        if (this.request[key] && key !== "compress") {
          otherOptionsSet = true
        }
      }

      // we set any other option
      if (otherOptionsSet) {
        this.disabled["dump_all_tables"] = 1
        this.request["dump_all_tables"]  = false
      } else {
        this.disabled["dump_all_tables"] = 0
        this.request["dump_all_tables"]  = true
      }

      this.$forceUpdate()
    },

    async backup() {
      this.backupResult = {}
      this.loading      = true
      setTimeout(async () => {
        try {
          const r = await SpireApi.v1().post("/backup/mysql", this.request)
          if (r.status === 200) {
            this.backupResult = r.data
            this.loading      = false

            if (r.data.file_path) {
              this.downloadBackup(path.basename(r.data.file_path))
            }

            for (const l of r.data.stdout.split("\n")) {
              if (l.includes("Error")) {
                this.error = l
                break
              }
            }
          }
        } catch (e) {
          // error notify
          if (e.response && e.response.data && e.response.data.error) {
            this.error   = e.response.data.error
            this.loading = false
          }
        }
      }, 1)

    },

    async downloadBackup(filename) {
      this.downloading = true

      try {
        const r = await SpireApi.v1().request("/backup/mysql-dump-download/" + filename,
          {
            responseType: 'blob',
          }
        )

        console.log(r.data)

        // create file link in browser's memory
        const href = URL.createObjectURL(r.data);

        // create "a" HTML element with href to file & click
        const link = document.createElement('a');
        link.href  = href;
        link.setAttribute('download', filename); //or any other extension
        document.body.appendChild(link);
        link.click();

        // clean up "a" element & remove ObjectURL
        document.body.removeChild(link);
        URL.revokeObjectURL(href);

        this.downloading = false

        return r
      } catch (e) {
        // error notify
        if (e.response && e.response.data && e.response.data.error) {
          this.error   = e.response.data.error
          this.loading = false
        }
      }
    }
  }
}
</script>

<style scoped>

</style>
