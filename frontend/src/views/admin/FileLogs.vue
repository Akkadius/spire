<template>
  <div>
    <div class="row">
      <div :class="fileToWatch && fileToWatch.length > 0 ? 'col-5' : 'col-12'">
        <eq-window
          :title="`Files (${files ? files.length : 0})`"
        >
          <div style="max-height: 80vh; overflow-y: scroll">
            <table
              class="eq-table eq-highlight-rows bordered player-events"
              v-if="files.length > 0"
              id="file-logs"
            >
              <thead class="eq-table-floating-header">
              <tr>
                <th style="width: 50px"></th>
                <th style="width: 150px">File</th>
                <th>Size</th>
                <th>Last Modified</th>
              </tr>
              </thead>
              <tbody>
              <tr
                class="fade-in"
                v-for="f in files"
                :key="`${f.path}-${f.modified_time}`"
              >
                <td class="text-center">
                  <b-button
                    class="btn-white btn-sm"
                    @click="viewLog(f)"
                    title="View and watch log file"
                  >
                    <i class="fa fa-eye"></i>
                  </b-button>
                </td>
                <td>{{ f.path }}</td>
                <td>{{ formatBytes(f.size) }}</td>
                <td>{{ formatTime(f.modified_time) }} {{ `(${formatTimeFromNow(f.modified_time)})` }}</td>
              </tr>
              </tbody>
            </table>
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
            style="position: absolute; top: -10px; z-index: 999999; left: 21px;"
          >
            <b-spinner small v-if="watchTimer"/>
          </div>

          <div
            class="minified-inputs row p-1 mb-3"
          >
            <div class="col-1 text-center p-0">
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

            <div class="col-11 pl-0">
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

          <div style="max-height: 80vh; overflow-y: scroll; " id="output">
            <pre
              class="mt-0 fade-in mb-1"
              style="width: 100%; word-wrap: break-word;  overflow-x: scroll"
              v-if="file && file.length > 0"
              id="file-contents"
              v-html="filterFileResults(file)"
            ></pre>
          </div>
        </eq-window>
      </div>
    </div>

  </div>
</template>

<script>
import EqWindow           from "@/components/eq-ui/EQWindow.vue";
import {SpireApi}         from "@/app/api/spire-api";
import moment             from "moment";
import {Navbar}           from "@/app/navbar";
import InfoErrorBanner    from "@/components/InfoErrorBanner.vue";
import LoaderFakeProgress from "@/components/LoaderFakeProgress.vue";
import ClipBoard          from "@/app/clipboard/clipboard";
import {debounce}         from "@/app/utility/debounce";

export default {
  name: "FileLogs",
  components: { LoaderFakeProgress, InfoErrorBanner, EqWindow },
  data() {
    return {
      files: [], // file listing
      fileListingTimer: null,

      outputContainer: null, // output container
      watchTimer: null,
      fileLineFilter: "",
      fileToWatch: "", // file name to watch
      fileCursor: 0,
      fileNotification: "", // info/error
      fileError: "", // info/error

      watching: false, // watching log file
    }
  },
  watch: {
    '$route'() {
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
    this.file = ""; // file contents to display
  },
  async mounted() {
    Navbar.collapse()
    // eqemuserver/logs

    this.loadQueryState()
    await this.init()
  },
  methods: {

    updateFilter: debounce(function () {
      this.updateQueryState()
    }, 600),

    filterFileResults(file) {
      return file.split("\n").filter((e) => {
        return e.toLowerCase().includes(this.fileLineFilter.toLowerCase())
      }).join("\n")
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
      if (this.search !== "") {
        q.f = this.fileToWatch
      }
      if (this.fileLineFilter && this.fileLineFilter.length > 0) {
        q.s = this.fileLineFilter
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
    },

    viewLog(f) {
      // reset
      this.fileLineFilter = ""
      this.fileCursor     = 0
      this.file           = ""
      // watch
      this.fileToWatch    = f.path
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
      return moment.unix(unix).format('M-D-YYYY h:mm:ss a')
    },
    formatTimeFromNow(unix) {
      return moment.unix(unix).fromNow()
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
                }
                else if (lbi === 2) {
                  newLine = newLine.replaceAll(`[${bc}]`, `<span class="font-weight-bold" style="color: #575555">${bc}</span> |`)
                }
                else if (lbi === 3) {
                  newLine = newLine.replaceAll(`[${bc}]`, `<span class="font-weight-bold" style="color: #ffffff">${bc}</span> |`)
                }
                else {
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
