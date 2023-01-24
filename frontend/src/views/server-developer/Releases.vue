<template>
  <div>
    <app-loader :is-loading="loading"/>

    <eq-window
      title="Release Analytics"
      class="p-3"
      v-if="!loading"
    >
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

      <div style="max-height:95vh; overflow-y: scroll">
        <table
          class="eq-table bordered eq-highlight-rows"
        >
          <thead class="eq-table-floating-header">
          <tr>
            <th style="width: 50px"></th>
            <th>Release</th>
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
          </tr>
          </thead>
          <tbody>
          <tr
            class="fade-in"
            v-for="(r, index) in releases"
            :key="r.tag_name"
          >
            <td class="text-center">
              <b-button
                variant="primary"
                class="btn-dark btn-sm btn-outline-warning"
                style="padding: 0px 6px;"
                title="View Release Notes"
                @click="viewReleaseNotes(r)"
              >
                <i class="fa fa-sticky-note-o"></i>
              </b-button>
            </td>
            <td>{{ r.name }}</td>
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

              <b-button
                variant="primary"
                class="btn-dark btn-sm btn-outline-warning ml-1"
                style="padding: 0px 6px;"
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
    </eq-window>
  </div>
</template>

<script>
import EqWindow   from "@/components/eq-ui/EQWindow.vue";
import {SpireApi} from "@/app/api/spire-api";
import util       from "util";
import {ROUTE}    from "@/routes";
import moment     from "moment";

export default {
  name: "Releases",
  components: {
    "v-runtime-template": () => import("v-runtime-template"),
    EqWindow
  },
  data() {
    return {
      loading: false,

      releaseNotes: "",

      releases: [],

      counts: []
    }
  },
  methods: {
    countChanges(r) {
      const count = r.body.split("\n* ").length
      return this.commify(count > 0 ? count : 0)
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
      this.$router.push(
        {
          path: util.format(ROUTE.RELEASE, r),
        }
      ).catch(() => {
      })
    },
    async loadCounts() {
      const r = await SpireApi.v1().get(`analytics/server-crash-report/counts`)
      if (r.status === 200) {
        this.counts = r.data
      }
    },
    async loadReleases() {
      const r = await SpireApi.v1().get(`analytics/releases`)
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
