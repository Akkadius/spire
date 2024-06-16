<template>
  <div class="card" v-if="stats" style="min-height: 90px">
    <div class="card-body pl-4 pr-4 pt-3 pb-3">

      <div class="row align-items-center">
        <div class="mr-3 pr-5">
          <h6 class="header-pretitle">
            {{ pageName }}
          </h6>

          <h1
            class="header-title"
            :title="stats.server_name"
            style="font-size: 1.1rem;">

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

          <small style="color: red" v-if="stopMessage !== ''">{{ stopMessage }}</small>



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
              <small class="text-muted text-uppercase mr-1">Players Online</small>
              <span class="h2 mb-0 ml-3">
              {{ stats && stats.client_list && stats.client_list.data ? stats.client_list.data.length : 0 }}
          </span>
            </div>

            <div
              class="d-none d-lg-block mr-3 ml-3"
              style="color: #95aac9; border-left: 1px solid #95aac9; height: 50px; opacity: .3"
            />

            <!-- Resource Utilization -->
            <div class="col-lg-auto col-sm-12 pl-3 pr-3 mt-3-mobile">
              <vue-ellipse-progress
                :progress="cpuPercent"
                animation="default 300 0"
                thickness="4"
                :legend-formatter="({ currentValue }) => `${currentValue}%`"
                :size="60"
                :color="getCpuLoadColor(cpuPercent)"
                empty-color="#95aac9"
                empty-thickness="1"
                font-size=".8rem"
                font-color="#95aac9"
              >
                <span
                  slot="legend-caption"
                  class="text-muted font-weight-bold"
                  style="font-size: 10px"
                > CPU </span>
              </vue-ellipse-progress>

              <vue-ellipse-progress
                class="ml-3"
                :progress="memoryPercent"
                animation="loop 600 0"
                thickness="4"
                :legend-formatter="({ currentValue }) => `${currentValue}%`"
                :size="60"
                color="#2c7be5"
                empty-color="#95aac9"
                empty-thickness="1"
                font-size=".8rem"
                font-color="#95aac9"
              >
              <span
                slot="legend-caption"
                class="text-muted font-weight-bold"
                style="font-size: 10px"
              > MEM</span>
              </vue-ellipse-progress>
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
import {VueEllipseProgress}         from "vue-ellipse-progress";
import {SpireWebsocket}             from "@/app/api/spire-websocket";
import moment                       from "moment";

export default {
  name: "AdminHeader",
  components: {
    ServerProcessButtonComponent,
    VueEllipseProgress,
  },
  data() {
    return {
      pageName: "",

      stats: {},

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

      SpireApi.v1().get("admin/system/resource-usage-summary").then((r) => {
        if (r.status === 200) {
          this.cpuPercent    = Math.round(r.data.cpu)
          this.memoryPercent = Math.round(r.data.memory.usedPercent)
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
    }
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
