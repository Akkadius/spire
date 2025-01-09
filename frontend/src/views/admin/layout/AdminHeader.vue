<template>
  <div class="card" v-if="stats" style="min-height: 90px">
    <div class="card-body pl-4 pr-4 pt-3 pb-3">

      <div class="row align-items-center">

        <!-- Server Name -->
        <div class="col-lg-auto">
          <h3 class="d-inline-block mr-3">
            <a
              href="javascript:void(0)"
              @click="toggleServerLock()"
              class="truncate-server-name"
            >
              <i
                :class="'fe ' + (serverLocked ? 'fe-lock' : 'fe-unlock')  + ' pr-2'"
                :title="serverLocked ? 'Server is locked' : 'Server is unlocked'"
                :style="`color: ${serverLocked ? 'red' : 'gray'}`"
              />

              <span v-if="stats.server_name">{{ stats.server_name }}</span>
            </a>
          </h3>

          <div>
            <server-process-button-component class="d-inline-block mr-3"/>
            <small style="color: red" v-if="stopMessage !== ''">{{ stopMessage }}</small>
          </div>

        </div>

        <!-- Server Metrics -->
        <div class="col-lg-8 col-sm-12 pr-0 ml-0 text-center">
          <div class="row align-items-center">

            <!-- Server Metrics -->
            <div class="col-lg-2 col-sm-12 mt-3-mobile">
              <div class="row" v-for="(metric, index) in serverStats" :key="index">
                <!-- Left Label -->
                <div class="col-lg-6 col-3 p-0 m-0 text-right" style="line-height: 1 !important">
                  <span class="small font-weight-bold text-muted" style="font-size: 12px;">
                    {{ metric.label }}
                  </span>
                </div>

                <!-- Right Value -->
                <div class="col-lg-6 col-3 p-0 m-0 pl-3 text-left" style="line-height: 1 !important">
                  <span class="small font-weight-bold" style="font-size: 12px;">
                    {{ metric.value.toLocaleString() }}
                  </span>
                </div>
              </div>
            </div>

            <div
              class="d-none d-lg-block mr-3 ml-3"
              style="color: #95aac9; border-left: 1px solid #95aac9; height: 50px; opacity: .3"
            />

            <!-- Resource Metrics -->
            <div
              class="col-lg-3 col-sm-12 pl-0 pr-0 mt-3-mobile"
              v-for="(host, index) in hostMetrics"
            >
              <!-- Render Metrics Dynamically -->
              <div class="row" v-for="(metric, index) in host" :key="index">
                <!-- Left Label -->
                <div class="col-3 p-0 m-0 text-right" style="line-height: .8 !important">
                  <span class="small font-weight-bold text-muted" style="font-size: 10px;">
                    {{ metric.label }}
                  </span>
                </div>

                <!-- Progress Bar -->
                <div class="col-6 p-0 m-0 mt-1" style="max-width: 120px" v-if="typeof metric.percent !== 'undefined'">
                  <eq-progress-bar
                    style="opacity: .95"
                    :percent="metric.percent"
                    :show-percent="false"
                    :color="metric.color"
                  />
                </div>

                <!-- Right Value -->
                <div
                  :class="'p-0 m-0 text-left' + (typeof metric.percent === 'undefined' ? 'col-9 ml-4 mb-1' : 'col-3')"
                  style="line-height: .8 !important"
                >
                  <span
                    :class="' font-weight-bold' + (typeof metric.percent === 'undefined' ? '' : 'small text-muted')"
                    style="font-size: 10px;"
                  >
                    {{ metric.value }}
                  </span>
                </div>
              </div>
            </div>

          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import ServerProcessButtonComponent from "@/views/admin/components/ServerProcessButtonComponent.vue";
import {EventBus}                   from "@/app/event-bus/event-bus";
import {SpireApi}                   from "@/app/api/spire-api";
import {SpireWebsocket}             from "@/app/api/spire-websocket";
import EqProgressBar                from "@/components/eq-ui/EQProgressBar.vue";
import Time                         from "@/app/time/time";

export default {
  name: "AdminHeader",
  components: {
    EqProgressBar,
    ServerProcessButtonComponent,
  },
  data() {
    return {
      pageName: "",

      stats: {},

      updateIntervalSeconds: 1,

      sys: [],

      diskStats: {
        previousReadBytes: 0,
        previousWriteBytes: 0,
        readBytesPerSec: 0,
        writeBytesPerSec: 0,
      },

      serverLocked: false,

      stopMessage: "",

      timer: null,
    }
  },

  computed: {

    serverStats() {
      return [
        {
          label: "World",
          value: this.stats && this.stats.world_online ? "Online" : "Offline",
        },
        {
          label: "Zoneservers",
          value: this.stats && this.stats.zone_list && this.stats.zone_list.data ? this.stats.zone_list.data.length : 0,
        },
        {
          label: "Players",
          value: this.stats && this.stats.client_list && this.stats.client_list.data ? this.stats.client_list.data.length : 0,
        },
        {
          label: "Uptime",
          value: this.stats && this.stats.uptime ? this.formatUptime(this.stats.uptime) : "N/A",
        },
        {
          label: "Locked",
          value: this.serverLocked ? "Yes" : "No",
        },
      ]
    },

    hostMetrics() {
      let metrics = []

      for (const e of this.sys) {
        let metric = [
          {
            label: "Host",
            value: e.hostname,
          },
          {
            label: "CPU",
            value: `${e.cpu || "N/A"} %`,
            percent: parseFloat(e.cpu || 0),
            color: this.getCpuLoadColor(e.cpu)
          },
          {
            label: "MEM",
            value: `${e.mem || "N/A"} %`,
            percent: parseFloat(e.mem || 0),
            color: "lightgreen"
          },
          {
            label: "DISK R",
            value: `${e.diskStats.readBytesPerSec.toFixed(2)} MB/s`,
            percent: Math.min(e.diskStats.readBytesPerSec / 100, 100),
            color: "deepskyblue"
          },
          {
            label: "DISK W",
            value: `${e.diskStats.writeBytesPerSec.toFixed(2)} MB/s`,
            percent: Math.min(e.diskStats.writeBytesPerSec / 100, 100),
            color: "tomato"
          },
          {
            label: "NET DL",
            value: `${this.bytesToMbytes(e.net['all'].bytes_recv_ps)} Mbps`,
            percent: Math.min(this.bytesToMbytes(e.net['all'].bytes_recv_ps) / 10, 100),
            color: "limegreen"
          },
          {
            label: "NET UL",
            value: `${this.bytesToMbytes(e.net['all'].bytes_sent_ps)} Mbps`,
            percent: Math.min(this.bytesToMbytes(e.net['all'].bytes_sent_ps) / 10, 100),
            color: "limegreen"
          }
        ]

        metrics.push(metric)
      }

      return metrics
    },

  },

  beforeDestroy() {
    clearInterval(this.timer)

    window.removeEventListener('keypress', this.keypressHandler)

    EventBus.$off("ROUTE_CHANGE", this.handleRouteChange);
    EventBus.$off('process-change')

    SpireWebsocket.removeEventListener('message', this.handleWebsocketMessage);
  },
  created() {
    EventBus.$on("ROUTE_CHANGE", this.handleRouteChange);

    SpireWebsocket.addEventListener('message', this.handleWebsocketMessage);

    window.addEventListener('keypress', this.keypressHandler)

    this.loadServerStats()

    EventBus.$on('process-change', async (event) => {
      this.loadServerStats()
    })

    this.getServerLockedStatus()

    this.timer = setInterval(() => {
      if (!document.hidden) {
        this.loadServerStats()
      }
    }, 1000)

    // initial page name set
    if (this.$route.meta && this.$route.meta.title) {
      this.pageName = this.$route.meta.title
    }
  },
  methods: {

    formatUptime(t) {
      // reformat
      t = t.replace("Worldserver Uptime |", "")
      t = t.replace(new RegExp(',', 'g'), '');
      t = t.replace("and", "")
      t = t.replace(" Days", "d")
      t = t.replace(" Day", "d")
      t = t.replace(" Weeks", "w")
      t = t.replace(" Week", "w")
      t = t.replace(" Months", "m")
      t = t.replace(" Month", "m")
      t = t.replace(" Hours", "h")
      t = t.replace(" Hour", "h")
      t = t.replace(" Minutes", "m")
      t = t.replace(" Minute", "m")
      t = t.replace(" Seconds", "s")
      t = t.replace(" Second", "s")
      t = t.replace(/^\s+|\s+$/g, "")

      return t.trim();
    },

    getCpuLoadColor(load) {
      if (load > 80) {
        return 'red'
      }
      if (load > 50) {
        return 'orange'
      }

      return '#2c7be5'
    },

    getServerLockedStatus() {
      SpireApi.v1().get("eqemuserver/get-lock-status").then((r) => {
        if (r.status === 200) {
          this.serverLocked = r.data.locked
        }
      })
    },

    toggleServerLock() {
      SpireApi.v1().post("eqemuserver/toggle-server-lock").then((r) => {
        if (r.status === 200) {
          this.serverLocked = r.data.locked
          this.$bvToast.toast(r.data.message, {
            title: "Server Lock",
            autoHideDelay: 2000,
            solid: true,
            toaster: 'b-toaster-bottom-right',
          })
        }
      })
    },

    async loadServerStats() {
      SpireApi.v1().get("eqemuserver/server-stats").then((r) => {
        if (r.status === 200) {
          this.stats = r.data
          this.$forceUpdate()
          EventBus.$emit("server-stats", r.data)
        }
      })

      SpireApi.v1().get("eqemuserver/system-all").then((r) => {
        let lastSys = JSON.parse(JSON.stringify(this.sys))

        let systems = []
        if (r.status === 200) {
          for (const e of r.data) {
            let s = {
              cpu: 0,
              mem: 0,
              diskStats: {
                readBytes: 0,
                writeBytes: 0,
                previousReadBytes: null, // Initialize to null
                previousWriteBytes: null, // Initialize to null
                readBytesPerSec: 0,
                writeBytesPerSec: 0,
              },
              net: {
                all: {
                  bytes_recv: 0,
                  bytes_sent: 0,
                  bytes_recv_ps: 0,
                  bytes_sent_ps: 0
                }
              }
            };

            // Find last system stats
            // we need this for calculating per-second changes
            // in disk stats, network stats
            if (lastSys.length > 0) {
              const last = lastSys.find((x) => {
                return x.hostname === e.hostname
              })
              if (last) {
                s = JSON.parse(JSON.stringify(last))
              }
            }

            s.cpu      = Math.round(e.cpu)
            s.mem      = Math.round(e.mem_percent)
            s.hostname = e.hostname

            // Disk stats
            const disk = e.disk[0]; // Assuming one disk for simplicity
            if (disk) {
              const nowReadBytes  = disk.readBytes;
              const nowWriteBytes = disk.writeBytes;

              // If this is the first pass, initialize previous values
              if (s.diskStats.previousReadBytes === null || s.diskStats.previousWriteBytes === null) {
                s.diskStats.previousReadBytes  = nowReadBytes;
                s.diskStats.previousWriteBytes = nowWriteBytes;
                s.diskStats.readBytesPerSec    = 0;
                s.diskStats.writeBytesPerSec   = 0;
              } else {
                // Calculate per-second change
                s.diskStats.readBytesPerSec  = (nowReadBytes - s.diskStats.previousReadBytes) / 1024 / 1024;
                s.diskStats.writeBytesPerSec = (nowWriteBytes - s.diskStats.previousWriteBytes) / 1024 / 1024;

                // Update previous values
                s.diskStats.previousReadBytes  = nowReadBytes;
                s.diskStats.previousWriteBytes = nowWriteBytes;
              }
            }

            // network stats
            for (let n of e.net) {
              if (typeof s.net[n.name] === 'undefined') {
                s.net[n.name] = {
                  bytes_recv: 0,
                  bytes_sent: 0,
                  bytes_recv_ps: 0,
                  bytes_sent_ps: 0,
                };
              }

              if (s.net[n.name].bytes_recv === 0 && s.net[n.name].bytes_sent === 0) {
                // First iteration: initialize without calculation
                s.net[n.name].bytes_recv    = n.bytesRecv;
                s.net[n.name].bytes_sent    = n.bytesSent;
                s.net[n.name].bytes_recv_ps = 0;
                s.net[n.name].bytes_sent_ps = 0;
              } else {
                // Calculate per-second changes
                s.net[n.name].bytes_recv_ps = n.bytesRecv - s.net[n.name].bytes_recv;
                s.net[n.name].bytes_sent_ps = n.bytesSent - s.net[n.name].bytes_sent;

                // Update previous values
                s.net[n.name].bytes_recv = n.bytesRecv;
                s.net[n.name].bytes_sent = n.bytesSent;
              }
            }

            systems.push(s)
          }

          this.sys = systems
        }
      })
    },

    handleRouteChange(e) {
      if (e && e.meta && e.meta.title) {
        this.pageName = e.meta.title
      }
    },

    keypressHandler(e) {
      if (e.srcElement.tagName !== 'BODY' && e.srcElement.tagName !== 'A') {
        return
      }

      if (window.location.pathname === '/login') {
        return
      }

      switch (String.fromCharCode(e.keyCode)) {
        case 'p':
          this.$root.$emit('bv::show::modal', 'start-server-modal')
          break
        case 'r':
          this.$root.$emit('bv::show::modal', 'restart-server-modal')
          break
        case 'c':
          this.$root.$emit('bv::show::modal', 'cancel-restart-server-modal')
          break
        case 's':
          this.$root.$emit('bv::show::modal', 'stop-server-modal')
          break
      }
    },
    handleWebsocketMessage(e) {
      if (e && e.data) {
        const data = JSON.parse(e.data)
        if (data.type === "stopTimer") {
          const timerData = JSON.parse(data.message)
          const time      = timerData.time
          const type      = timerData.type
          const remaining = this.calcRemainingTime(time)

          if (remaining === "") {
            this.stopMessage = ""
            return
          }
          this.stopMessage = `${type.charAt(0).toUpperCase() + type.slice(1)} in ${remaining}`
        }
      }
    },
    calcRemainingTime(unixTimestamp) {
      return Time.calculateRemainingTimeServerReboot(unixTimestamp)
    },

    bytesToMbytes: function (bytes) {
      return parseFloat((bytes * 8) / 1024 / 1024).toFixed(2); // Convert to Mbps
    },
  },
}
</script>

<style>
.truncate-server-name {
  margin-top: 10px;
  display: inline-flex;
  -webkit-line-clamp: 8; /* Number of lines to show */
  -webkit-box-orient: vertical;
  overflow: hidden;
  text-overflow: ellipsis;
  max-width: 450px; /* Set the max-width to the desired length */
}
</style>
