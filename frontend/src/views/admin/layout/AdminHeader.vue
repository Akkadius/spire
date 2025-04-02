<template>
  <eq-window-simple
    v-if="stats"
    class="admin-header-window mb-4 pb-0"
    style="font-family: 'Cerebri Sans', sans-serif; color: white; padding: 10px; z-index: 1;"
    :title="stats.server_name"
  >

    <div class="row align-items-center">

      <!-- Server Metrics -->
      <div class="col-lg-12 col-sm-12 pl-0 pr-0 ml-0 text-center">
        <div
          class="row align-items-center admin-header"
        >

          <!-- Server Metrics -->
          <div class="col-lg-2 col-sm-12 mt-3-mobile pl-0 pr-0">
            <div class="row" v-for="(metric, index) in serverStats.filter((e) => e.value !== '')" :key="index" style="height: 15px;">
              <!-- Left Label -->
              <div class="col-lg-6 col-3 p-0 m-0 text-right">
                  <span class="small font-weight-bold" style="font-size: 12px; ">
                    {{ metric.label }}
                  </span>
              </div>

              <!-- Right Value -->
              <div class="col-lg-6 col-3 p-0 m-0 pl-3 text-left">

                <span
                  v-if="metric.label !== 'Locked' && metric.label !== 'World'"
                  class="small" :style="'font-size: 12px;' + (metric.color ? 'color: ' + metric.color : '')"
                >
                    {{ metric.value.toLocaleString() }}
                  </span>

                <!-- Server Lock -->
                <a
                  href="javascript:void(0)"
                  v-if="metric.label === 'Locked'"
                  class="small font-weight-bold text-muted"
                  style="font-size: 12px;"
                  @click="toggleServerLock"
                >
                  <span v-if="serverLocked" class="text-danger">
                    <i class="fa fa-lock mr-1"></i> Locked
                  </span>
                  <span v-else class="text-success">
                    <i class="fa fa-unlock mr-1"></i> Unlocked
                  </span>
                </a>

                <!-- Server Process Button -->
                <div v-if="metric.label === 'World'">
                  <server-process-button-component
                    style="z-index:1000"
                    :server-status="metric.value"
                    class="d-inline-block mr-3"
                  />
                </div>

              </div>
            </div>
          </div>

          <!-- Dash Stats -->
          <div class="col-lg-1 col-sm-12 mt-3-mobile pl-0 pr-0">
            <div class="row" v-for="(metric, index) in dashStats" :key="index" style="height: 15px;">
              <!-- Left Label -->
              <div class="col-lg-6 col-3 p-0 m-0 text-right">
                  <span class="small font-weight-bold" style="font-size: 12px; ">
                    {{ metric.label }}
                  </span>
              </div>

              <!-- Right Value -->
              <div class="col-lg-6 col-3 p-0 m-0 pl-3 text-left">

                <span
                  v-if="metric.label !== 'Locked' && metric.label !== 'World'"
                  class="small" style="font-size: 12px;"
                >
                    {{ metric.value.toLocaleString() }}
                  </span>

              </div>
            </div>
          </div>

          <!-- Resource Metrics -->
          <div
            class="col-lg-2 col-sm-12 pl-0 pr-0 mt-3-mobile"
            v-for="(host, index) in hostMetrics"
            style="display: inline-block"
          >
            <!-- Render Metrics Dynamically -->
            <div class="row" v-for="(metric, index) in host" :key="index">
              <!-- Left Label -->
              <div class="col-3 p-0 m-0 text-right" style="line-height: .8 !important">
                  <span class="small font-weight-bold" style="font-size: 10px; opacity: .9">
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
  </eq-window-simple>
</template>

<script>
import ServerProcessButtonComponent from "@/views/admin/components/ServerProcessButtonComponent.vue";
import {EventBus}                   from "@/app/event-bus/event-bus";
import {SpireApi}                   from "@/app/api/spire-api";
import {SpireWebsocket}             from "@/app/api/spire-websocket";
import EqProgressBar                from "@/components/eq-ui/EQProgressBar.vue";
import Time                         from "@/app/time/time";
import EqWindow                     from "@/components/eq-ui/EQWindow.vue";
import EqWindowSimple               from "@/components/eq-ui/EQWindowSimple.vue";
import {Notify}                     from "@/app/Notify";

export default {
  name: "AdminHeader",
  components: {
    EqWindowSimple,
    EqWindow,
    EqProgressBar,
    ServerProcessButtonComponent,
  },
  data() {
    return {
      pageName: "",

      dashstats: {},
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
      shutdownTime: "",

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
          label: "Players Online",
          value: this.stats && this.stats.client_list && this.stats.client_list.data ? this.stats.client_list.data.length : 0,
        },
        {
          label: "Uptime",
          value: this.stats && this.stats.uptime ? this.formatUptime(this.stats.uptime) : "N/A",
        },
        {
          label: "Shutdown Timer",
          value: this.shutdownTime,
          color: 'red',
        },
        {
          label: "Locked",
          value: this.serverLocked ? "Yes" : "No",
        },
      ]
    },

    dashStats() {
      return [
        {
          label: "Accounts",
          value: this.dashstats && this.dashstats.accounts ? this.dashstats.accounts.toLocaleString() : 0,
        },
        {
          label: "Characters",
          value: this.dashstats && this.dashstats.characters ? this.dashstats.characters.toLocaleString() : 0,
        },
        {
          label: "Guilds",
          value: this.dashstats && this.dashstats.guilds ? this.dashstats.guilds.toLocaleString() : 0,
        },
        {
          label: "Items",
          value: this.dashstats && this.dashstats.items ? this.dashstats.items.toLocaleString() : 0,
        },
        {
          label: "NPCs",
          value: this.dashstats && this.dashstats.npcs ? this.dashstats.npcs.toLocaleString() : 0,
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
            value: `${e.diskStats.readBytesPerSec.toFixed(2) > 1000 ? 0 : e.diskStats.readBytesPerSec.toFixed(2)} MB/s`,
            percent: Math.min(e.diskStats.readBytesPerSec / 100, 100),
            color: "deepskyblue"
          },
          {
            label: "DISK W",
            value: `${e.diskStats.writeBytesPerSec.toFixed(2) > 1000 ? 0 : e.diskStats.writeBytesPerSec.toFixed(2)} MB/s`,
            percent: Math.min(e.diskStats.writeBytesPerSec / 100, 100),
            color: "tomato"
          },
          {
            label: "NET DL",
            value: `${this.bytesToMbytes(e.net['all'].bytes_recv_ps) > 10000 ? 0 : this.bytesToMbytes(e.net['all'].bytes_recv_ps)} Mbps`,
            percent: Math.min(this.bytesToMbytes(e.net['all'].bytes_recv_ps) / 10, 100),
            color: "limegreen"
          },
          {
            label: "NET UL",
            value: `${this.bytesToMbytes(e.net['all'].bytes_sent_ps) > 10000 ? 0 : this.bytesToMbytes(e.net['all'].bytes_sent_ps)} Mbps`,
            percent: Math.min(this.bytesToMbytes(e.net['all'].bytes_sent_ps) / 10, 100),
            color: "limegreen"
          }
        ]

        metrics.push(metric)
      }

      // sort by hostname value
      metrics.sort((a, b) => {
        if (a[0].label === "Host" && a[0].value < b[0].value) {
          return -1
        }
        if (a[0].label === "Host" && a[0].value > b[0].value) {
          return 1
        }
        return 0
      })

      return metrics
    },

  },

  beforeDestroy() {
    clearInterval(this.timer)

    EventBus.$off("ROUTE_CHANGE", this.handleRouteChange);
    EventBus.$off('process-change')

    SpireWebsocket.removeEventListener('message', this.handleWebsocketMessage);
  },
  created() {
    EventBus.$on("ROUTE_CHANGE", this.handleRouteChange);

    SpireWebsocket.addEventListener('message', this.handleWebsocketMessage);

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

    SpireApi.v1().get("eqemuserver/dashboard-stats").then((r) => {
      if (r.status === 200) {
        this.dashstats = r.data
      }
    })
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
      if (confirm("Are you sure you want to toggle the server lock?")) {
        SpireApi.v1().post("eqemuserver/toggle-server-lock").then((r) => {
          if (r.status === 200) {
            this.serverLocked = r.data.locked

            Notify.toast(r.data.message);
          }
        })
      }
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
            const disk = e && e.disk && e.disk.length > 0 ? e.disk[0] : null;
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
            this.shutdownTime = ""
            return
          }
          this.stopMessage = `${type.charAt(0).toUpperCase() + type.slice(1)} in ${remaining}`
          this.shutdownTime = remaining
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

<style scoped>
.admin-header-window .text-muted {

}
</style>
