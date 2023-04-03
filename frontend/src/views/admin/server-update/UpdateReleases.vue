<template>
  <div>
    <app-loader :is-loading="loading"/>

    <b-modal
      id="release-notes"
      centered
      :title="`Release Notes`"
      size="lg"
    >
      <v-runtime-template class="changelog" :template="releaseNotes"/>

      <template #modal-footer>
        <div class="">

        </div>
      </template>
    </b-modal>

    <eq-window
      title="Official Releases"
      class="p-3"
      v-if="!loading"
    >
      <div style="max-height:95vh; overflow-y: scroll">
        <table
          class="eq-table bordered eq-highlight-rows fade-in"
        >
          <thead class="eq-table-floating-header">
          <tr>
            <th style="width: 140px"></th>
            <th>Release</th>
            <th class="text-center">Install</th>
            <th>Change Count</th>

            <th>Released</th>
            <th>
              <i class="fa fa-windows"></i>
              Windows Downloads
            </th>
            <th>
              <i class="fa fa-linux"></i>
              Linux Downloads
            </th>
            <th>
              <i class="fa fa-download"></i>
              Download Links
            </th>
            <th class="text-center">Crash Count</th>
            <th></th>
          </tr>
          </thead>
          <tbody>
          <tr
            class="fade-in"
            v-for="(r, index) in releases"
            :key="r.tag_name"
            :style="(versionGreater(r.tag_name, version) ? 'opacity: 1' : 'opacity: .4')"
          >
            <td class="text-center">
              <b-button
                variant="primary"
                class="btn-dark btn-sm btn-outline-warning"
                style="padding: 0px 6px;"
                title="View Release Notes"
                @click="viewReleaseNotes(r)"
              >
                <i class="fa fa-sticky-note-o"></i> Changelog
              </b-button>
            </td>
            <td>{{ r.name }}</td>
            <td class="text-center">
              <b-button
                variant="primary"
                class="btn-dark btn-sm btn-outline-warning"
                style="padding: 0px 6px;"
                title="Install"
                @click="installRelease(r)"
              >
                <i class="fa fa-cloud-download"></i> Install
              </b-button>
            </td>
            <td>{{ countChanges(r) }}</td>
            <td>{{ formatTime(r.published_at) }} ({{ formatDate(r.published_at) }})</td>
            <td>{{ getWindowsDownloads(r) }}</td>
            <td>{{ getLinuxDownloads(r) }}</td>
            <td>
              <a
                :href="getWindowsDownloadLink(r)"
                v-if="getWindowsDownloadLink(r)"
                class="text-muted"
              >
                <i class="fa fa-windows"></i> Windows (x64)
              </a>
              <a
                :href="getLinuxDownloadLink(r)"
                v-if="getLinuxDownloadLink(r)"
                class="text-muted"
              >
                <i class="fa fa-linux"></i> Linux (x64)
              </a>
            </td>
            <td class="text-center">
              {{ getCrashCount(r) }}
            </td>
            <td class="text-center">
              <b-button
                variant="primary"
                class="btn-dark btn-sm btn-outline-warning ml-1"
                style="padding: 0px 6px; width: 50px"
                title="View Release Crashes"
                @click="goToRelease(r.name.replaceAll('v', ''))"
              >
                <i class="fa fa-arrow-right"></i>
              </b-button>
            </td>
          </tr>
          </tbody>
        </table>
      </div>

      <div
        class="row justify-content-center"
        style="position: absolute; bottom: -5%; z-index: 9999999; width: 100%"
      >
        <div class="col-6">
          <info-error-banner
            style="width: 100%"
            :slim="true"
            :notification="notification"
            :error="error"
            @dismiss-error="error = ''"
            @dismiss-notification="notification = ''"
            class="mt-3"
          />
        </div>
      </div>

    </eq-window>
  </div>

</template>

<script>
import EqWindow        from "@/components/eq-ui/EQWindow.vue";
import util            from "util";
import {ROUTE}         from "@/routes";
import moment          from "moment";
import EqTabs          from "@/components/eq-ui/EQTabs.vue";
import EqTab           from "@/components/eq-ui/EQTab.vue";
import axios           from "axios";
import semver          from "semver";
import {AppEnv}        from "@/app/env/app-env";
import {SpireApi}      from "@/app/api/spire-api";
import InfoErrorBanner from "@/components/InfoErrorBanner.vue";

export default {
  name: "UpdateReleases",
  components: {
    InfoErrorBanner,
    EqTab,
    EqTabs,
    "v-runtime-template": () => import("v-runtime-template"),
    EqWindow
  },
  props: {
    version: {
      required: true
    },
  },

  data() {
    return {
      tabSelected: "Official Releases",

      loading: false,

      releaseNotes: "",

      releases: [],

      counts: [],
      selfBuilt: [],

      // notification / errors
      notification: "",
      error: "",
    }
  },
  watch: {
    $route(to, from) {
      this.tabSelected = "Official Releases"
      this.loadQueryState()
    }
  },
  methods: {
    async installRelease(r) {
      let release = this.getWindowsDownloadLink(r)
      if (AppEnv.getOS().includes("linux")) {
        release = this.getLinuxDownloadLink(r)
      }

      if (confirm(`Install this release?\n\nThis will replace your current binaries \n\n[${r.name}] ${release}`)) {
        try {
          const r = await SpireApi.v1().post(`eqemuserver/install-release/` + release)
          if (r.status === 200) {
            this.notification = r.data.message

            this.$emit("refresh-version", true)
          }
        } catch (err) {
          // error notify
          if (err.response && err.response.data && err.response.data.error) {
            this.error = err.response.data.error
          }
        }
      }

      console.log(release)

    },

    versionGreater(v1, v2) {
      if (!v1.includes(".") || !v2.includes(".")) {
        return false;
      }

      if (v1.split(".").length !== 3 || v2.split(".").length !== 3) {
        return false;
      }

      return semver.gt(v1, v2)
    },

    // state
    updateQueryState() {
      console.log("trigger")
      let q = {};
      if (this.tabSelected !== "") {
        q.s = this.tabSelected
      }

      this.$router.push(
        {
          path: ROUTE.RELEASES,
          query: q
        }
      ).catch(() => {
      })
    },
    loadQueryState() {
      if (this.$route.query.s && this.$route.query.s.length > 0) {
        this.tabSelected = this.$route.query.s
      }
    },

    countChanges(r) {
      const count = r.body.split("\n* ").length
      return this.commify(count > 0 ? (count - 1) : 0)
    },
    commify(x) {
      return x.toString().replace(/\B(?=(\d{3})+(?!\d))/g, ",");
    },

    viewReleaseNotes(r) {
      setTimeout(() => {
        this.$bvModal.show('release-notes')
        this.releaseNotes = r.body

        const md = require("markdown-it")({
          html: true,
          xhtmlOut: false,
          breaks: true,
          typographer: false,
          linkify: true
        });

        let markdownRaw   = md.render(r.body);
        this.releaseNotes = "<div>" + markdownRaw + "</div>"

      }, 10)
    },
    formatTime(time) {
      return moment(time).fromNow()
    },
    formatDate(time) {
      return moment(time).format("MMM Do YYYY")
    },
    goToRelease(r) {
      window.open(util.format("http://spire.akkadius.com/dev/release/%s", r), 'release_' + r);
    },
    async loadCounts() {
      const r = await axios.get(`http://spire.akkadius.com/api/v1/analytics/server-crash-report/counts`)
      if (r.status === 200) {
        this.counts    = r.data
        this.selfBuilt = r.data.filter((e) => {
          return e.server_version.includes("-dev")
        }).sort((a, b) => {
          return b.server_version.localeCompare(a.server_version);
        });
      }
    },
    async loadReleases() {
      const r = await axios.get(`http://spire.akkadius.com/api/v1/analytics/releases`)
      if (r.status === 200) {
        this.releases = r.data.data.filter((e) => {
          return e.name.split(".").length === 3
        })
      }
    },
    getWindowsDownloads(r) {
      const f = r.assets.find((e) => {
        return e.name.includes("windows")
      })

      if (f) {
        return f.download_count
      }

      return 0
    },
    getWindowsDownloadLink(r) {
      const f = r.assets.find((e) => {
        return e.name.includes("windows")
      })

      if (f) {
        return f.browser_download_url
      }

      return 0
    },
    getLinuxDownloadLink(r) {
      const f = r.assets.find((e) => {
        return e.name.includes("linux")
      })

      if (f) {
        return f.browser_download_url
      }

      return 0
    },
    getLinuxDownloads(r) {
      const f = r.assets.find((e) => {
        return e.name.includes("linux")
      })

      if (f) {
        return f.download_count
      }

      return 0
    },
    getCrashCount(r) {
      const version = r.name.replaceAll("v", "")

      for (let v of this.counts) {
        if (v.server_version === version) {
          return v.crash_count;
        }
      }

      return 0
    }
  },
  async mounted() {
    this.loadQueryState();
    this.loading = true;
    this.loadCounts().then((r) => {
      this.loadReleases()
      this.loading = false
    })
  },

}
</script>

<style scoped>

</style>
