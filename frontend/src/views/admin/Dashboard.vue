<template>
  <div class="row" style="max-height: 85vh; overflow-y: scroll; top: -15px; position: inherit">
    <div class="col-lg-6">

      <div class="row">
        <div class="col-sm-6 col-lg-6">
          <dashboard-process-counts/>

          <dashboard-networking-info/>
          <dashboard-system-info-v2/>
        </div>
        <div class="col-sm-6 col-lg-6">
          <dashboard-cpu-info/>
        </div>
      </div>

    </div>

    <!-- Right side -->
    <div class="col-lg-6">
      <players-online-component/>
    </div>
  </div>
</template>

<script>
import ServerProcessButtonComponent from "@/views/admin/components/ServerProcessButtonComponent";
import DashboardProcessCounts       from "@/views/admin/components/DashboardProcessCounts";
import DashboardCpuInfo             from "@/views/admin/components/DashboardCpuInfo";
import DashboardSystemInfo          from "@/views/admin/components/DashboardSystemInfo";
import DashboardCounter             from "@/views/admin/components/DashboardCounter.vue";
import PlayersOnlineComponent       from "@/views/admin/components/PlayersOnlineComponent.vue";
import {SpireApi}                   from "@/app/api/spire-api";
import DashboardSystemInfoV2        from "@/views/admin/components/DashboardSystemInfoV2.vue";
import DashboardNetworkingInfo      from "@/views/admin/components/DashboardNetworkingInfo.vue";

export default {
  components: {
    DashboardNetworkingInfo,
    DashboardSystemInfoV2,
    PlayersOnlineComponent,
    DashboardCounter,
    DashboardSystemInfo,
    DashboardCpuInfo,
    DashboardProcessCounts,
    ServerProcessButtonComponent
  },
  data() {
    return {
      loaded: false,
      stats: {},
      sysinfo: {},
      statLoop: null,
      circleProgressInitialized: null,

      timer: null,
      statsTimer: null,
    }
  },
  beforeDestroy() {
    clearInterval(this.timer)
    clearInterval(this.statsTimer)
  },
  created: async function () {
    this.fetchDashboardStats()

    if (this.timer) {
      clearInterval(this.timer)
    }

    this.statsTimer = setInterval(() => {
      if (!document.hidden) {
        this.fetchDashboardStats()
      }
    }, 60 * 1000)

  },
  methods: {
    fetchDashboardStats() {
      SpireApi.v1().get("eqemuserver/dashboard-stats").then((r) => {
        if (r.status === 200) {
          this.stats = r.data
          this.checkLoaded()
        }
      })
    },

    checkLoaded() {
      this.loaded = (
        Object.keys(this.stats).length >= 0 &&
        Object.keys(this.sysinfo).length >= 0
      )
    },

  }
}
</script>
