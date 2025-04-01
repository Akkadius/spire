<template>
  <div>
    <div class="row">
      <div
        :class="(fileToWatch && fileToWatch.length > 0) || (searchResults && searchResults.length > 0) ? 'col-5' : 'col-12'"
      >
        <eq-window
          :title="`Files (${files ? files.length : 0})`"
        >
          <div style="max-height: 79vh; overflow-y: scroll;">

            <div
              class="row justify-content-center"
              style="position: absolute; top: 20px; z-index: 9999999; width: 100%"
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

            <div class="row">
              <div class="col-12">
                <b-input-group>
                  <input
                    type="text"
                    class="form-control"
                    v-model="searchAllString"
                    placeholder="Search all files by search string"
                    v-on:keyup.enter="searchAll()"
                    autofocus
                  >
                  <b-input-group-append>
                    <b-button size="sm" variant="outline-warning" @click="searchAll()">
                      <i class="fa fa-search"></i>
                    </b-button>
                  </b-input-group-append>

                </b-input-group>
              </div>
            </div>

            <div class="mt-3 ml-1 row">
              <div class="col-12 p-0">
                <b-button
                  v-for="f in filterTypes"
                  :key="`${f.label}-${countLogs(f.search)}`"
                  class="btn-sm mr-2 btn-dark btn-dark fade-in"
                  @click="filterOn(f)"
                  :style="`opacity: ${(filterType.length > 0 && filterType !== f.search ? '.2' : '1')}`"
                >
                  <i class="fa fa-filter"></i>
                  {{ f.label }}
                  ({{ countLogs(f.search) }})
                </b-button>

                <b-button
                  class="btn-sm mr-2 btn-outline-danger btn-dark"
                  @click="deleteAllLogFiles()"
                  title="Delete all log files in view"
                >
                  <i class="fa fa-trash"></i>
                </b-button>

                <b-button
                  class="btn-sm mr-2 btn-dark"
                  @click="resetAll()"
                  title="Reset"
                >
                  <i class="fa fa-dot-circle-o"></i>
                </b-button>
              </div>
            </div>

            <div
              class="font-weight-bold text-center p-5"
              v-if="getFilteredFiles().length === 0"
            >
              There are no files to be shown in the current filter
            </div>

            <div style="overflow-x: scroll">
              <table
                class="eq-table eq-highlight-rows bordered player-events mt-3"
                style="overflow: auto; width: 100%; border-collapse: collapse; white-space: nowrap; "
                v-if="files && files.length > 0 && getFilteredFiles().length > 0"
                id="file-logs"
              >
                <thead class="eq-table-floating-header">
                <tr>
                  <th style="width: 130px"></th>
                  <th style="width: 260px" class="text-center">Last Modified</th>
                  <th style="width: 100px" class="text-center">Size</th>
                  <th>File</th>
                </tr>
                </thead>
                <tbody>
                <tr
                  :class="'fade-in ' + (fileToWatch && f.path === fileToWatch ? 'pulsate-highlight-white' : '')"
                  v-for="f in getFilteredFiles()"
                  :key="`${f.path}-${f.modified_time}`"
                >
                  <td class="text-center p-1">
                    <b-button
                      class="btn-danger btn-sm"
                      @click="deleteFile(f)"
                      title="Delete File"
                    >
                      <i class="fa fa-trash"></i>
                    </b-button>

                    <b-button
                      class="btn-white btn-sm ml-1"
                      @click="viewLog(f)"
                      title="View and watch log file"
                    >
                      <i class="fa fa-eye"></i>
                    </b-button>
                  </td>
                  <td class="text-center">{{ formatTime(f.modified_time) }} {{
                      `(${formatTimeFromNow(f.modified_time)})`
                    }}
                  </td>
                  <td class="text-center">{{ formatBytes(f.size) }}</td>
                  <td>{{ f.path }}</td>
                </tr>
                </tbody>
              </table>
            </div>
          </div>

        </eq-window>
      </div>

      <div class="col-7" v-if="fileToWatch && fileToWatch.length > 0">
        <div
          class="row justify-content-center"
          style="position: absolute; top: 50px; z-index: 9999999; width: 100%"
        >
          <div class="col-4">
            <info-error-banner
              style="width: 100%"
              :slim="true"
              :notification="fileNotification"
              :error="fileError"
              @dismiss-error="fileError = ''"
              @dismiss-notification="fileNotification = ''"
              class="mt-3"
            />
          </div>
        </div>

        <eq-window :title="fileToWatch">
          <div
            style="position: absolute; top: 25px; z-index: 999999; left: 21px;"
          >
            <b-spinner small v-if="watchTimer"/>
          </div>

          <div
            class="minified-inputs row p- mb-3"
          >
            <div class="col-2 pr-0 text-center">
              <b-button
                class="btn-danger btn-sm mb-1"
                title="Stop log stream"
                @click="stopLogStream()"
                :disabled="!watchTimer"
              >
                <i class="fa fa-stop"></i>
              </b-button>

              <b-button
                class="btn-white btn-sm ml-1 mb-1"
                title="Play log stream"
                :disabled="watchTimer && parseInt(watchTimer) > 0"
                @click="startLogStream()"
              >
                <i class="fa fa-play"></i>
              </b-button>

              <b-button
                class="btn-white btn-sm ml-1 mb-1"
                title="Copy to clipboard"
                @click="copyFileContentsToClipboard()"
              >
                <i class="fa fa-copy"></i>
              </b-button>
            </div>

            <div class="col-10 pl-0">
              <input
                type="text"
                class="form-control mt-0 mb-0"
                placeholder="Filter log contents..."
                v-on:keyup="updateFilter()"
                v-model="fileLineFilter"
                style="width: 100%"
              >
            </div>

          </div>

          <div style="max-height: 75vh; overflow-y: scroll; " id="output">
            <pre
              class="mt-0 fade-in mb-1"
              style="width: 100%; word-wrap: break-word; overflow-x: scroll; padding-bottom: 0 !important"
              v-if="file && file.length > 0"
              id="file-contents"
              v-html="filterFileResults(file)"
            ></pre>

            <div>
              <span class="font-weight-bold">Line Buffer</span> {{ commify(currentLineBufferLength) }} /
              {{ commify(lineBufferLimit) }} (Max)
              <span class="font-weight-bold">Cursor</span> {{ commify(fileCursor) }}
            </div>
          </div>
        </eq-window>
      </div>

      <div
        class="col-7"
        v-if="searchResults && searchResults.length > 0"
      >

        <eq-window
          class="fade-in"
          :title="`Search All Results`"
        >

          <div style="max-height: 80vh; overflow-y: scroll">
            <div
              v-for="f in searchResults"
              :key="f.file"
            >
              <div class="font-weight-bold mt-3">{{ f.file }}</div>
              <pre
                class="mt-3 fade-in mb-0"
                style="width: 100%; word-wrap: break-word;  overflow-x: scroll; padding-bottom: 0 !important"
                v-if="searchResults && searchResults.length > 0"
                id="search-results"
                v-html="formatSearchResult(f.lines)"
              ></pre>
            </div>

          </div>
        </eq-window>
      </div>
    </div>

  </div>
</template>

<script>
import EqWindow           from "@/components/eq-ui/EQWindow.vue";
import {SpireApi}         from "@/app/api/spire-api";
import {Navbar}           from "@/app/navbar";
import InfoErrorBanner    from "@/components/InfoErrorBanner.vue";
import LoaderFakeProgress from "@/components/LoaderFakeProgress.vue";
import ClipBoard          from "@/app/clipboard/clipboard";
import {debounce}         from "@/app/utility/debounce";
import Time               from "@/app/time/time";

export default {
  name: "FileLogs",
  components: { LoaderFakeProgress, InfoErrorBanner, EqWindow },
  data() {
    return {
      searchResults: [],

      filterType: "",
      files: [], // file listing
      fileListingTimer: null,
      notification: "", // info/error
      error: "", // info/error

      currentLineBufferLength: 0,
      lineBufferLimit: 1000,
      outputContainer: null, // output container
      watchTimer: null,
      fileToWatch: "", // file name to watch
      fileCursor: 0,
      fileNotification: "", // info/error
      fileError: "", // info/error

      watching: false, // watching log file

      filterTypes: [
        { label: "World", search: "world" },
        { label: "Zone", search: "zone" },
        { label: "UCS", search: "ucs" },
        { label: "Login", search: "login" },
        { label: "QS", search: "query_server" },
        { label: "Crashes", search: "crashes" },
      ]
    }
  },
  watch: {
    '$route'() {
      this.fileToWatch = ""
      this.fileCursor  = 0
      this.stopLogStream()

      this.loadQueryState()
      this.init()
    },
  },

  beforeDestroy() {
    Navbar.expand()

    this.stopLogStream()
    if (this.fileListingTimer) {
      clearInterval(this.fileListingTimer);
    }
  },
  created() {
    this.file            = ""; // file contents to display
    this.fileLineFilter  = "";
    this.searchAllString = "";
  },
  async mounted() {
    Navbar.collapse()
    // eqemuserver/logs

    this.loadQueryState()
    await this.init()
  },
  methods: {

    resetAll() {
      this.filterType = ""
      this.stopLogStream()
      this.fileToWatch     = ""
      this.fileCursor      = 0
      this.searchAllString = ""
      this.searchResults   = []
      this.updateQueryState()
    },

    getFilteredFiles() {
      return this.files && this.files.length > 0 ? this.files.filter((e) => {
        return (this.filterType.length > 0 && e.path.includes(this.filterType)) || this.filterType.length === 0
      }) : []
    },

    async deleteAllLogFiles() {
      if (confirm(`Are you sure? This will delete all log files shown in the current filter`)) {
        for (const f of this.getFilteredFiles()) {
          try {
            const r = await SpireApi.v1().delete(`eqemuserver/log/${f.path}`)
          } catch (e) {
            if (e.response && e.response.data && e.response.data.error) {
              this.error = e.response.data.error
            }
          }
        }

        this.resetAll()

        this.notification = "Files deleted successfully!";

        await this.loadLogListing();
        this.updateQueryState()
      }
    },

    formatSearchResult(lines) {
      let newLines = '';
      lines.forEach((l) => {
        const info = `Line ${l.line_number}`

        newLines += `${info.padStart(8)} | ${this.formatContents(l.line)}\n`
      })

      return newLines
    },

    async searchAll() {
      // reset some other state
      this.filterType = ""
      this.stopLogStream()
      this.fileToWatch = ""

      this.updateQueryState()
    },

    async doSearchAll() {
      try {
        const r = await SpireApi.v1().get(`eqemuserver/log-search/${this.searchAllString}`)
        if (r.status === 200) {
          this.searchResults = r.data
        }
      } catch (e) {
        if (e.response && e.response.data && e.response.data.error) {
          this.error = e.response.data.error
        }
      }
    },

    filterOn(f) {
      this.filterType = f.search
      this.updateQueryState()
    },

    countLogs(search) {
      return this.commify(
        this.files && this.files.length > 0 ? this.files.filter((e) => {
          return (e.path.includes(search) && !e.path.includes("crash")) || (search.includes("crash") && e.path.includes(search))
        }).length : 0
      )
    },

    async deleteFile(f) {
      if (confirm(`Are you sure you want to delete this file? \n\n${f.path}`)) {
        try {
          const r = await SpireApi.v1().delete(`eqemuserver/log/${f.path}`)
          if (r.status === 200) {
            this.notification = "File deleted successfully!";
            this.loadLogListing();
            this.resetAll()
          }
        } catch (e) {
          if (e.response && e.response.data && e.response.data.error) {
            this.error = e.response.data.error
          }
        }
      }
    },

    commify(x) {
      return x.toString().replace(/\B(?=(\d{3})+(?!\d))/g, ",");
    },

    updateFilter: debounce(function () {
      this.updateQueryState()
    }, 600),

    filterFileResults(file) {
      const f = file.split("\n").filter((e) => {
        return e.toLowerCase().includes(this.fileLineFilter.toLowerCase())
      }).slice(-1000)

      this.currentLineBufferLength = f.length

      return f.join("\n")
    },

    copyFileContentsToClipboard() {
      ClipBoard.copyFromElement("file-contents");
      this.fileNotification = "Copied to clipboard!"
    },

    stopLogStream() {
      if (this.watchTimer) {
        clearInterval(this.watchTimer);
        this.watchTimer = null
      }
    },

    startLogStream() {
      if (this.watchTimer) {
        clearInterval(this.watchTimer);
      }

      this.watchTimer = setInterval(() => {
        if (document.hidden) {
          return
        }

        if (this.fileToWatch && this.fileToWatch.length > 0) {
          this.loadLogFile(this.fileToWatch)
        }
      }, 1000)
    },

    // state
    updateQueryState() {
      let q = {};
      if (this.fileToWatch !== "") {
        q.f = this.fileToWatch
      }
      if (this.fileLineFilter && this.fileLineFilter.length > 0) {
        q.s = this.fileLineFilter
      }
      if (this.filterType && this.filterType.length > 0) {
        q.ft = this.filterType
      }
      if (this.searchAllString && this.searchAllString.length > 0) {
        q.sa = this.searchAllString
      }

      this.$router.push(
        {
          path: this.$route.path,
          query: q
        }
      ).catch(() => {
      })
    },
    loadQueryState() {
      if (this.$route.query.f && this.$route.query.f.length > 0) {
        this.fileToWatch = this.$route.query.f
      }
      if (this.$route.query.s && this.$route.query.s.length > 0) {
        this.fileLineFilter = this.$route.query.s
      }
      if (this.$route.query.ft && this.$route.query.ft.length > 0) {
        this.filterType = this.$route.query.ft
      }
      if (this.$route.query.sa && this.$route.query.sa.length > 0) {
        this.searchAllString = this.$route.query.sa
      }
    },

    async loadLogListing() {
      try {
        const r = await SpireApi.v1().get('eqemuserver/logs')
        if (r.status === 200) {
          this.files = r.data
        }
      } catch (e) {
        if (e.response && e.response.data && e.response.data.error) {
          this.error = e.response.data.error
        }
      }
    },

    async init() {
      if (this.fileToWatch && this.fileToWatch.length > 0) {
        this.loadLogFile(this.fileToWatch)
        this.startLogStream()
      }

      this.loadLogListing()

      if (!this.fileListingTimer) {
        this.fileListingTimer = setInterval(() => {
          if (document.hidden) {
            return
          }

          this.loadLogListing()
        }, 15000)
      }

      if (this.searchAllString && this.searchAllString.length > 0) {
        this.doSearchAll()
      }
    },

    viewLog(f) {
      // reset
      this.fileError               = ""
      this.fileNotification        = ""
      this.error                   = ""
      this.notification            = ""
      this.currentLineBufferLength = 0;
      this.fileLineFilter          = ""
      this.fileCursor              = 0
      this.file                    = ""
      this.searchAllString         = ""
      this.searchResults           = []
      // watch
      this.fileToWatch             = f.path
      this.updateQueryState()
    },

    async loadLogFile(f) {
      try {
        let q = {};
        if (this.fileCursor > 0) {
          q.params        = {}
          q.params.cursor = this.fileCursor
        }

        const r = await SpireApi.v1().get(`eqemuserver/log/${f}`, q)
        if (r.status === 200) {
          if (this.fileCursor && this.fileCursor > 0) {
            this.fileCursor = parseInt(this.fileCursor) + parseInt(r.data.cursor)
            this.file += this.formatContents(r.data.contents)
          } else {
            this.fileCursor = parseInt(r.data.cursor)
            this.file       = this.formatContents(r.data.contents)
          }

          this.$forceUpdate()

          setTimeout(() => {
            if (!this.outputContainer) {
              const t = document.getElementById("output")
              if (t) {
                this.outputContainer = t
              }
            }

            if (this.outputContainer) {
              this.outputContainer.scrollTop = this.outputContainer.scrollHeight + 100;
            }
          }, 1)

        }
      } catch (e) {
        if (e.response && e.response.data && e.response.data.error) {
          this.fileError = e.response.data.error

          // if error is "makeslice: len out of range" then restart the log stream
          if (this.fileError.includes("makeslice: len out of range")) {
            console.log("len out of range, restart log stream")
            this.stopLogStream()
            this.startLogStream()
          }
        }
      }
    },

    formatBytes(bytes, decimals = 2) {
      if (!+bytes) return '0 Bytes'

      const k     = 1024
      const dm    = decimals < 0 ? 0 : decimals
      const sizes = ['Bytes', 'KB', 'MB', 'GB', 'TB', 'PB', 'EB', 'ZB', 'YB']

      const i = Math.floor(Math.log(bytes) / Math.log(k))

      return `${parseFloat((bytes / Math.pow(k, i)).toFixed(dm))} ${sizes[i]}`
    },
    formatTime(unix) {
      return Time.formatUnix(unix, 'M-D-YYYY h:mm:ss a')
    },
    formatTimeFromNow(unix) {
      return Time.fromNowUnix(unix)
    },

    formatContents(c) {
      let lines = []
      c.split("\n").forEach((line) => {
        let newLine = line

        let lbs = line.split("[")
        if (lbs.length > 0) {
          lbs.forEach((lb, lbi) => {
            let rbs = lb.split("]")
            if (rbs.length > 0) {
              let bc = rbs[0] // bracket contents
              if (bc && bc.length > 0) {
                if (lbi === 1) {
                  newLine = newLine.replaceAll(`[${bc}]`, `<span class="font-weight-bold" style="color: lightblue">${bc}</span> |`)
                } else if (lbi === 2) {
                  newLine = newLine.replaceAll(`[${bc}]`, `<span class="font-weight-bold" style="color: #575555">${bc}</span> |`)
                } else if (lbi === 3) {
                  newLine = newLine.replaceAll(`[${bc}]`, `<span class="font-weight-bold" style="color: #ffffff">${bc}</span> |`)
                } else {
                  newLine = newLine.replaceAll(`[${bc}]`, `[<span class="font-weight-bold" style="color: #ffffb1">${bc}</span>]`)
                }
              }
            }
          })
        }

        lines.push(newLine)
      })

      return lines.join("\n")
    },
  }
}
</script>

<style scoped>

</style>
