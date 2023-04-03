<template>
  <div>
    <app-loader :is-loading="!loaded"></app-loader>

    <div style="max-height: 85vh; overflow-y: scroll; overflow-x: hidden">
      <div class="row row-cards" v-if="loaded">
        <dashboard-counter name="Accounts" icon="user" :counter="kFormatter(stats.accounts)"/>
        <dashboard-counter name="Characters" icon="user" :counter="kFormatter(stats.characters)"/>
        <dashboard-counter name="Guilds" icon="shield" :counter="kFormatter(stats.guilds)"/>
      </div>

      <div class="row row-cards" v-if="loaded">
        <dashboard-counter name="Items" icon="award" :counter="kFormatter(stats.items)"/>
        <dashboard-counter name="NPCs" icon="gitlab" :counter="kFormatter(stats.npcs)"/>
        <dashboard-counter name="Server Uptime" :counter="formatUptime(stats.uptime)"/>
      </div>

      <div class="row row-cards">
        <div class="col-lg-6">

          <div class="row">
            <div class="col-sm-6 col-lg-6">
              <dashboard-process-counts/>

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
    </div>


  </div>
</template>

<script>
import ServerProcessButtonComponent from "@/views/admin/components/ServerProcessButtonComponent";
import DashboardProcessCounts       from "@/views/admin/components/DashboardProcessCounts";
import DashboardCpuInfo             from "@/views/admin/components/DashboardCpuInfo";
import DashboardSystemInfo          from "@/views/admin/components/DashboardSystemInfo";
import Timer                        from "@/app/timer/timer";
import {OcculusClient}              from "@/app/api/eqemu-admin-client-occulus";
import DashboardCounter             from "@/views/admin/components/DashboardCounter.vue";
import {OS}                         from "@/app/os/os";
import PlayersOnlineComponent       from "@/views/admin/components/PlayersOnlineComponent.vue";
import {SpireApi}                   from "@/app/api/spire-api";
import DashboardSystemInfoV2        from "@/views/admin/components/DashboardSystemInfoV2.vue";

export default {
  components: {
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

    checkLoaded() {
      this.loaded = (
        Object.keys(this.stats).length >= 0 &&
        Object.keys(this.sysinfo).length >= 0
      )
    },

    kFormatter: function (number) {
      return number > 999 ? (number / 1000).toFixed(1) + 'k' : number
    }
  }
}
</script>
