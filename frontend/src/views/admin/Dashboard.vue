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
        <dashboard-counter name="Server Uptime" :counter="stats.uptime"/>
      </div>

      <div class="row row-cards" v-if="loaded">
        <div class="col-lg-6">

          <div class="row">
            <div class="col-sm-6 col-lg-6">
              <dashboard-process-counts/>
              <dashboard-system-info :sysinfo="sysinfo"/>
            </div>
            <div class="col-sm-6 col-lg-6">
              <dashboard-cpu-info :sysinfo="sysinfo"/>
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
import {EqemuAdminClient}           from "@/app/api/eqemu-admin-client-occulus";
import DashboardCounter             from "@/views/admin/components/DashboardCounter.vue";
import {OS}                         from "@/app/os/os";
import PlayersOnlineComponent       from "@/views/admin/components/PlayersOnlineComponent.vue";

export default {
  components: {
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
      circleProgressInitialized: null
    }
  },
  beforeDestroy() {
    clearInterval(Timer.timer['sys-info'])
  },
  created: async function () {
    EqemuAdminClient.getDashboardStats().then(response => {
      if (response) {
        this.stats = response
        this.checkLoaded()
      }
    })

    this.loadSysInfo()

    if (Timer.timer['sys-info']) {
      clearInterval(Timer.timer['sys-info'])
    }

    const sysInfoTimer = (OS.get() === "Linux" ? 1000 : 5000);
    this.loadSysInfo();
    Timer.timer['sys-info'] = setInterval(() => {
      if (!document.hidden) {
        this.loadSysInfo()
      }
    }, sysInfoTimer)

  },
  methods: {
    checkLoaded() {
      this.loaded = (
        Object.keys(this.stats).length >= 0 &&
        Object.keys(this.sysinfo).length >= 0
      )
    },

    loadSysInfo: function () {
      EqemuAdminClient.getSysInfo().then(response => {
        if (response) {
          this.sysinfo = response
        }

        this.checkLoaded()
      })
    },
    kFormatter: function (number) {
      return number > 999 ? (number / 1000).toFixed(1) + 'k' : number
    }
  }
}
</script>
