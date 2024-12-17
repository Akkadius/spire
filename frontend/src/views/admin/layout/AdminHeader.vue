<template>
  <div class="card" v-if="stats" style="min-height: 90px">
    <div class="card-body pl-4 pr-4 pt-3 pb-3">

      <div class="row align-items-center">
        <div class="mr-3 pr-5">
          <h6 class="header-pretitle d-inline-block mr-3">
            {{ pageName }}
          </h6>

          <small style="color: red" v-if="stopMessage !== ''">{{ stopMessage }}</small>

          <h1
            class="header-title"
            :title="stats.server_name"
            style="font-size: 1.1rem;"
          >

            <server-process-button-component class="d-inline-block mr-3"/>

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

          </h1>


        </div>

        <div
          class="d-none d-lg-block mr-3 ml-3"
          style="color: #95aac9; border-left: 1px solid #95aac9; height: 50px; opacity: .3"
        />

        <div class="col-lg-8 col-sm-12 pl-3 pr-0 ml-0 text-center">

          <div class="row align-items-center">

            <div class="col-lg-auto col-sm-12 mt-3-mobile">
              <small class="text-muted text-uppercase mr-1">World</small>
              <span
                :class="`badge badge-${stats.world_online ? 'success' : 'danger'} ml-3`"
                style="font-size: 12px"
              >{{ stats.world_online ? 'Online' : 'Offline' }}</span>
            </div>

            <div
              class="d-none d-lg-block mr-3 ml-3"
              style="color: #95aac9; border-left: 1px solid #95aac9; height: 50px; opacity: .3"
            />

            <div class="col-lg-auto col-sm-12 mt-3-mobile">
              <small class="text-muted text-uppercase">Zoneservers</small>
              <span class="h2 mb-0 ml-3">
                {{ stats && stats.zone_list && stats.zone_list.data ? stats.zone_list.data.length : 0 }}
              </span>
            </div>

            <div
              class="d-none d-lg-block mr-3 ml-3"
              style="color: #95aac9; border-left: 1px solid #95aac9; height: 50px; opacity: .3"
            />

            <div class="col-lg-auto col-sm-12 mt-3-mobile">
              <small class="text-muted text-uppercase mr-1">Players</small>
              <span class="h2 mb-0 ml-3">
              {{ stats && stats.client_list && stats.client_list.data ? stats.client_list.data.length : 0 }}
          </span>
            </div>

            <div
              class="d-none d-lg-block mr-3 ml-3"
              style="color: #95aac9; border-left: 1px solid #95aac9; height: 50px; opacity: .3"
            />

            <!-- Resource Utilization -->
            <div class="col-lg-4 col-sm-12 pl-3 pr-3 mt-3-mobile mb-2">

              <!-- CPU -->
              <div class="row">

                <div class="col-6 p-0 m-0 text-right" style="line-height: .8 !important">
                  <span class="small font-weight-bold text-muted" style="font-size: 10px;">
                    CPU - {{ cpuPercent ? cpuPercent : "N/A" }} %
                  </span>
                </div>
                <div class="col-6 p-0 m-0 mt-1">
                  <eq-progress-bar
                    style="opacity: .95"
                    :percent="parseFloat(cpuPercent)"
                    :show-percent="false"
                    :color="getCpuLoadColor(cpuPercent)"
                  />
                </div>
              </div>

              <!-- Memory -->
              <div class="row">
                <div class="col-6 p-0 m-0 text-right" style="line-height: .8 !important">
                  <span class="small font-weight-bold text-muted" style="font-size: 10px;">
                    MEM - {{ memoryPercent ? memoryPercent : "N/A" }} %
                  </span>
                </div>
                <div class="col-6 p-0 m-0 mt-1">
                  <eq-progress-bar
                    style="opacity: .95"
                    :percent="parseFloat(memoryPercent)"
                    :show-percent="false"
                    color="lightgreen"
                  />
                </div>
              </div>

              <!-- Disk Read -->
              <div class="row">
                <div class="col-6 p-0 m-0 text-right" style="line-height: .8 !important">
                  <span class="small font-weight-bold text-muted" style="font-size: 10px;">
                    DISK READ - {{ diskStats.readBytesPerSec.toFixed(2) }} MB/s
                  </span>
                </div>
                <div class="col-6 p-0 m-0 mt-1">
                  <eq-progress-bar
                    style="opacity: .95"
                    :percent="Math.min(diskStats.readBytesPerSec / 100, 100)"
                    :show-percent="false"
                    color="deepskyblue"
                  />
                </div>
              </div>

              <!-- Disk Write -->
              <div class="row">
                <div class="col-6 p-0 m-0 text-right" style="line-height: .8 !important">
                  <span class="small font-weight-bold text-muted" style="font-size: 10px;">
                    DISK WRITE - {{ diskStats.writeBytesPerSec.toFixed(2) }} MB/s
                  </span>
                </div>
                <div class="col-6 p-0 m-0 mt-1">
                  <eq-progress-bar
                    style="opacity: .95"
                    :percent="Math.min(diskStats.writeBytesPerSec / 100, 100)"
                    :show-percent="false"
                    color="tomato"
                  />
                </div>
              </div>


              <!-- Network Download -->
              <div class="row" v-if="net && net['all'] && bytesToMbytes(net['all'].bytes_recv_ps)">
                <div class="col-6 p-0 m-0 text-right" style="line-height: .8 !important">
                  <span class="small font-weight-bold text-muted" style="font-size: 10px;">
                    NET DL - {{ bytesToMbytes(net['all'].bytes_recv_ps) < 10000 ? bytesToMbytes(net['all'].bytes_recv_ps) : 0 }} Mbps
                  </span>
                </div>
                <div class="col-6 p-0 m-0 mt-1">
                  <eq-progress-bar
                    style="opacity: .95"
                    :percent="bytesToMbytes(net['all'].bytes_recv_ps) < 10000 ? bytesToMbytes(net['all'].bytes_recv_ps) / 10 : 0"
                    :show-percent="false"
                    color="limegreen"
                  />
                </div>
              </div>

              <!-- Network Upload -->
              <div class="row" v-if="net && net['all'] && bytesToMbytes(net['all'].bytes_sent_ps)">
                <div class="col-6 p-0 m-0 text-right" style="line-height: .8 !important">
                  <span class="small font-weight-bold text-muted" style="font-size: 10px;">
                    NET UL - {{ bytesToMbytes(net['all'].bytes_sent_ps) < 10000 ? bytesToMbytes(net['all'].bytes_sent_ps) : 0 }} Mbps
                  </span>
                </div>
                <div class="col-6 p-0 m-0 mt-1">
                  <eq-progress-bar
                    style="opacity: .95"
                    :percent="bytesToMbytes(net['all'].bytes_sent_ps) < 10000 ? bytesToMbytes(net['all'].bytes_sent_ps) / 10 : 0"
                    :show-percent="false"
                    color="limegreen"
                  />
                </div>
              </div>

            </div>

            <!--            <div-->
            <!--              class="d-none d-lg-block ml-3 mr-3"-->
            <!--              style="color: #95aac9; border-left: 1px solid #95aac9; height: 50px; opacity: .3"-->
            <!--            />-->
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
import moment                       from "moment";
import EqProgressBar                from "@/components/eq-ui/EQProgressBar.vue";

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

      lastUpdateTime: 0,
      updateIntervalSeconds: 1,

      net: {
        all: {
          bytes_recv: 0,
          bytes_sent: 0,
          bytes_recv_ps: 0,
          bytes_sent_ps: 0
        }
      },

      diskStats: {
        previousReadBytes: 0,
        previousWriteBytes: 0,
        readBytesPerSec: 0,
        writeBytesPerSec: 0,
      },

      serverLocked: false,

      cpuPercent: 0,
      memoryPercent: 0,

      stopMessage: "",

      timer: null,
    }
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
      if (!document.hidden && this.readyToPoll()) {
        this.loadServerStats()
      }
    }, 1000)

    // initial page name set
    if (this.$route.meta && this.$route.meta.title) {
      this.pageName = this.$route.meta.title
    }
  },
  methods: {

    readyToPoll() {
      return Date.now() - this.lastUpdateTime > this.updateIntervalSeconds * 1000
    },

    getClientListCount() {
      return this.stats && this.stats.client_list && this.stats.client_list.data ? this.stats.client_list.data.length : 0
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

          const clientCount = this.getClientListCount()

          if (clientCount > 1000) {
            this.updateIntervalSeconds = 30
          }
          if (clientCount > 500) {
            this.updateIntervalSeconds = 10
          } else if (clientCount > 100) {
            this.updateIntervalSeconds = 5
          }

          this.lastUpdateTime = Date.now()
        }
      })

      SpireApi.v1().get("admin/system/all").then((r) => {
        if (r.status === 200) {
          this.cpuPercent    = Math.round(r.data.cpu)
          this.memoryPercent = Math.round(r.data.mem_percent)

          // Disk stats
          const disk = r.data.disk[0]; // Assuming one disk for simplicity
          if (disk) {
            const nowReadBytes = disk.readBytes;
            const nowWriteBytes = disk.writeBytes;

            // Calculate per-second change
            this.diskStats.readBytesPerSec = (nowReadBytes - this.diskStats.previousReadBytes) / 1024 / 1024;
            this.diskStats.writeBytesPerSec = (nowWriteBytes - this.diskStats.previousWriteBytes) / 1024 / 1024;

            // Update previous values
            this.diskStats.previousReadBytes = nowReadBytes;
            this.diskStats.previousWriteBytes = nowWriteBytes;
          }

          // network stats
          for (let n of r.data.net) {
            if (typeof this.net[n.name] === 'undefined') {
              this.net[n.name] = {}
            }

            if (typeof this.net[n.name]['bytes_recv'] === 'undefined') {
              this.net[n.name]['bytes_recv']    = 0
              this.net[n.name]['bytes_recv_ps'] = 0
            } else {
              this.net[n.name]['bytes_recv_ps'] = n['bytesRecv'] - this.net[n.name]['bytes_recv']
            }

            if (typeof this.net[n.name]['bytes_sent'] === 'undefined') {
              this.net[n.name]['bytes_sent']    = 0
              this.net[n.name]['bytes_sent_ps'] = 0
            } else {
              this.net[n.name]['bytes_sent_ps'] = n['bytesSent'] - this.net[n.name]['bytes_sent']
            }

            // track per second
            this.net[n.name]['bytes_recv'] = n['bytesRecv']
            this.net[n.name]['bytes_sent'] = n['bytesSent']
          }

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
          const remaining = this.calculateTimeRemainingWithMoment(time)

          if (remaining === "") {
            this.stopMessage = ""
            return
          }
          this.stopMessage = `${type.charAt(0).toUpperCase() + type.slice(1)} in ${remaining}`
        }
      }
    },
    calculateTimeRemainingWithMoment(unixTimestamp) {
      const now      = moment();
      const endTime  = moment.unix(unixTimestamp);
      const duration = moment.duration(endTime.diff(now));

      if (duration.asMilliseconds() < 0) {
        return "";
      }

      const minutes = duration.minutes();
      const seconds = duration.seconds();

      return `${minutes} minutes, ${seconds} seconds`;
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
  max-width: 240px; /* Set the max-width to the desired length */
}
</style>
