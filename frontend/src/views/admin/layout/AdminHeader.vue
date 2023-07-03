<template>
  <div class="card">
    <div class="card-body pl-4 pr-4 pt-3 pb-3">

      <div class="row align-items-center">
        <div class="col-lg-6 col-sm-12">
          <h6 class="header-pretitle">
            {{ pageName }}
          </h6>

          <h1 class="header-title">
            <span v-if="stats.server_name">{{ stats.server_name }}</span>
          </h1>
        </div>

        <div class="col-lg-6 col-sm-12 pl-0 pr-0">

          <div class="row align-items-center text-center mt-3-mobile">

          <!-- Resource Utilization -->
          <div class="col-lg-2 col-sm-12 pl-0 pr-0">
<!--            <small-->
<!--              style="position: absolute; top: 25px; left: 0px;"-->
<!--              class="text-muted text-uppercase"-->
<!--            >Resources</small>-->

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

          <div class="col-lg-auto col-sm-12">
            <small class="text-muted text-uppercase mr-1">Launcher</small>
            <span
              :class="`badge badge-${stats.launcher_online ? 'success' : 'danger'} ml-3`"
              style="font-size: 12px"
            >{{ stats.launcher_online ? 'Online' : 'Offline' }}</span>
          </div>

          <div class="col-lg-auto col-sm-12">
            <small class="text-muted text-uppercase mr-1">World</small>
            <span
              :class="`badge badge-${stats.world_online ? 'success' : 'danger'} ml-3`"
              style="font-size: 12px"
            >{{ stats.world_online ? 'Online' : 'Offline' }}</span>
          </div>

          <div class="col-lg-auto col-sm-12">
            <small class="text-muted text-uppercase mr-1">UCS</small>
            <span
              :class="`badge badge-${stats.ucs_online ? 'success' : 'danger'} ml-3`"
              style="font-size: 12px"
            >{{ stats.ucs_online ? 'Online' : 'Offline' }}</span>
          </div>

          <div class="col-lg-auto col-sm-12">
            <small class="text-muted text-uppercase">Zoneservers</small>
            <span class="h2 mb-0 ml-3">
            {{ stats && stats.zone_list && stats.zone_list.data ? stats.zone_list.data.length : 0 }}
          </span>
          </div>

          <div class="col-lg-auto col-sm-12">
            <small class="text-muted text-uppercase mr-1">Players Online</small>
            <span class="h2 mb-0 ml-3">
            {{ stats && stats.client_list && stats.client_list.data ? stats.client_list.data.length : 0 }}
          </span>
          </div>

          <div class="col-lg-2 col-sm-12 text-right-no-mobile">
            <server-process-button-component/>
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
import {VueEllipseProgress}         from "vue-ellipse-progress";

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


      cpuPercent: 0,
      memoryPercent: 0,

      timer: null,
    }
  },
  beforeDestroy() {
    clearInterval(this.timer)

    window.removeEventListener('keypress', this.keypressHandler)

    EventBus.$off("ROUTE_CHANGE", this.handleRouteChange);
    EventBus.$off('process-change')
  },
  created() {
    EventBus.$on("ROUTE_CHANGE", this.handleRouteChange);

    window.addEventListener('keypress', this.keypressHandler)

    this.loadServerStats()

    EventBus.$on('process-change', async (event) => {
      this.loadServerStats()
    })

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
    }
  },
}
</script>

<style>
</style>
