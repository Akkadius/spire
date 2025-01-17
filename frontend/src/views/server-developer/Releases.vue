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
      title="Release Version Analytics"
      class="p-3"
      v-if="!loading"

    >
      <eq-tabs
        :selected="tabSelected"
        @on-selected="tabSelected = $event; updateQueryState()"
      >
        <eq-tab
          class="fade-in"
          :name="`Official Releases`"
        >
          <div style="max-height:85vh; overflow-y: scroll">
            <table
              class="eq-table bordered eq-highlight-rows"
            >
              <thead class="eq-table-floating-header">
              <tr>
                <th style="width: 50px"></th>
                <th class="text-center pl-0 pr-0">
                  <div title="Change Count">
                    <i class="fa fa-pencil"></i>
                  </div>
                </th>
                <th class="text-center">Release</th>

                <th style="width: 220px">Released</th>
                <th style="width: 220px">
                  <i class="fa fa-download"></i>
                  Download Links
                </th>
                <th class="text-center pl-0 pr-0" style="width: 75px">
                  <div title="Windows downloads">
                    <i class="fa fa-windows"></i>
                    <i class="fa fa-download ml-2"></i>
                  </div>
                </th>
                <th class="text-center pl-0 pr-0" style="width: 75px">
                  <div title="Linux downloads">
                    <i class="fa fa-linux"></i>
                    <i class="fa fa-download ml-2"></i>
                  </div>
                </th>
                <th class="text-center pl-0 pr-0 pt-2 pb-0" style="width: 100px">Unique <br>Crashes<br></th>
                <th class="text-center pl-0 pr-0 pt-2 pb-0" style="width: 100px">Unique <br>Resolved<br></th>
                <th class="text-center pl-0 pr-0 pt-2 pb-0" style="width: 100px">Total <br>Crashes</th>
                <th></th>
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
                    class="btn-dark btn-sm btn-dark"
                    style="padding: 0px 6px;"
                    title="View Release Notes"
                    @click="viewReleaseNotes(r)"
                  >
                    <i class="fa fa-sticky-note-o"></i>
                  </b-button>
                </td>
                <td class="text-center pl-1 pr-1" style="width: 75px">{{ countChanges(r) }}</td>
                <td class="text-center" style="width: 100px">{{ r.name }}</td>
                <td>{{ formatTime(r.published_at) }} ({{ formatDate(r.published_at) }})</td>
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
                <td class="text-center">{{ getWindowsDownloads(r) }}</td>
                <td class="text-center">{{ getLinuxDownloads(r) }}</td>
                <td class="text-center p-1">
                  {{ getUniqueCrashCount(r) }}
                </td>
                <td class="text-center p-1">
                  <span
                    title="Unique crashes that have been resolved and released in a newer version"
                    style="color: limegreen"
                  >{{ getUniqueCrashResolvedCount(r) }}

                    <span v-if="getUniqueCrashResolvedCount(r) > 0">
                      ({{ (getUniqueCrashResolvedCount(r) / getUniqueCrashCount(r) * 100).toFixed(2) }}%)
                    </span>
                  </span>
                </td>
                <td class="text-center pl-0 pr-0">
                  {{ getCrashCount(r) }}
                </td>
                <td class="">
                  <b-button
                    variant="primary"
                    class="btn-dark btn-sm btn-dark ml-1"
                    style="padding: 0px 6px;"
                    title="View Release Crashes"
                    @click="goToRelease(r.name.replaceAll('v', ''))"
                  >
                    View Crashes <i class="fa fa-arrow-right"></i>
                  </b-button>
                </td>
              </tr>
              </tbody>
            </table>
          </div>

        </eq-tab>
        <eq-tab
          class="fade-in"
          :name="`Self Built`"
        >
          <div style="max-height:85vh; overflow-y: scroll">
            <table
              class="eq-table bordered eq-highlight-rows"
            >
              <thead class="eq-table-floating-header">
              <tr>
                <th class="text-center">Release</th>
                <th class="text-center">Crash Count</th>
                <th></th>
              </tr>
              </thead>
              <tbody>
              <tr
                class="fade-in"
                v-for="(r, index) in selfBuilt"
                :key="r.server_version"
              >
                <td class="text-center">{{ r.server_version }}</td>
                <td class="text-center">
                  {{ r.crash_count }}
                </td>
                <td class="text-center">
                  <b-button
                    variant="primary"
                    class="btn-dark btn-sm btn-dark ml-1"
                    style="padding: 0px 6px; width: 50px"
                    title="View Release Crashes"
                    @click="goToRelease(r.server_version)"
                  >
                    <i class="fa fa-arrow-right"></i>
                  </b-button>
                </td>
              </tr>
              </tbody>
            </table>
          </div>

        </eq-tab>
      </eq-tabs>


    </eq-window>
  </div>

</template>

<script>
import EqWindow   from "@/components/eq-ui/EQWindow.vue";
import {SpireApi} from "@/app/api/spire-api";
import util       from "util";
import {ROUTE}    from "@/routes";
import EqTabs     from "@/components/eq-ui/EQTabs.vue";
import EqTab      from "@/components/eq-ui/EQTab.vue";
import Time       from "@/app/time/time";

export default {
  name: "Releases",
  components: {
    EqTab,
    EqTabs,
    "v-runtime-template": () => import("v-runtime-template"),
    EqWindow
  },
  data() {
    return {
      tabSelected: "Official Releases",

      loading: false,

      releaseNotes: "",

      releases: [],

      counts: [],
      uniqueCounts: [],
      selfBuilt: []
    }
  },
  watch: {
    $route(to, from) {
      this.tabSelected = "Official Releases"
      this.loadQueryState()
    }
  },
  methods: {
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
      return Time.fromNow(time)
    },
    formatDate(time) {
      return Time.format(time, "MMM D YYYY")
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
        this.counts       = r.data.crash_report_counts
        this.uniqueCounts = r.data.unique_crash_counts
        this.selfBuilt    = r.data.crash_report_counts.filter((e) => {
          return e.server_version.includes("-dev")
        }).sort((a, b) => {
          return b.server_version.localeCompare(a.server_version);
        });
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
    },
    getCrashResolvedCount(r) {
      const version = r.name.replaceAll("v", "")

      for (let v of this.counts) {
        if (v.server_version === version) {
          return v.resolved_count;
        }
      }

      return 0
    },
    getUniqueCrashCount(r) {
      const version = r.name.replaceAll("v", "")

      let total = 0;
      for (let v of this.uniqueCounts) {
        if (v.server_version === version) {
          total++;
        }
      }

      return total
    },
    getUniqueCrashResolvedCount(r) {
      const version = r.name.replaceAll("v", "")

      console.log(this.uniqueCounts)

      let total = 0;
      for (let v of this.uniqueCounts) {
        if (v.server_version === version && v.resolved_count > 0) {
          total++;
        }
      }

      return total
    },
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
