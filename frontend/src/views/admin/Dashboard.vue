<template>
  <div class="col-12">

    <!-- Header -->
    <div class="card">
      <div class="card-body">
        <div class="row align-items-center">
          <div class="col">

            <h6 class="header-pretitle">
              Dashboard
            </h6>

            <h1 class="header-title">
              <span v-if="stats.long_name">{{ stats.long_name }}</span>
            </h1>



          </div>

          <div class="col-auto">
            <server-process-button-component/>
          </div>
        </div>

      </div>
    </div>


    <app-loader :is-loading="!loaded"></app-loader>

    <div class="row row-cards" v-if="loaded">
      <dashboard-counter name="Accounts" icon="user" :counter="this.kFormatter(stats.accounts)"></dashboard-counter>
      <dashboard-counter name="Characters" icon="user" :counter="this.kFormatter(stats.characters)"></dashboard-counter>
      <dashboard-counter name="Guilds" icon="shield" :counter="this.kFormatter(stats.guilds)"></dashboard-counter>
    </div>

    <div class="row row-cards" v-if="loaded">
      <dashboard-counter name="Items" icon="award" :counter="this.kFormatter(stats.items)"></dashboard-counter>
      <dashboard-counter name="NPCs" icon="gitlab" :counter="this.kFormatter(stats.npcs)"></dashboard-counter>
      <dashboard-counter name="Server Uptime" :counter="stats.uptime"></dashboard-counter>
    </div>

    <div class="row row-cards" v-if="loaded">
      <div class="col-lg-6">

        <div class="row">
          <dashboard-process-counts></dashboard-process-counts>
          <dashboard-cpu-info :sysinfo="sysinfo"></dashboard-cpu-info>
          <dashboard-system-info :sysinfo="sysinfo"></dashboard-system-info>
        </div>

        <div class="row">

        </div>

      </div>

      <!-- Right side -->
      <div class="col-lg-6">
        <players-online></players-online>
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
import PlayersOnline                from "@/views/admin/components/PlayersOnline";
import DashboardCounter             from "@/views/admin/components/DashboardCounter.vue";

export default {
  components: {
    DashboardCounter,
    PlayersOnline,
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

    var self = this

    if (Timer.timer['sys-info']) {
      clearInterval(Timer.timer['sys-info'])
    }

    const sysInfoTimer = navigator.appVersion.indexOf("Win") ? 2500 : 1000;
    this.loadSysInfo();
    Timer.timer['sys-info'] = setInterval(function () {
      if (!document.hidden) {
        self.loadSysInfo()
      }
    }, sysInfoTimer)

  },
  computed: {
    cpuLoadDisplay: function () {
      return (Object.keys(this.sysinfo).length > 0 ? (Math.round(this.sysinfo.cpu.load.currentload * 100) / 100) : 0)
    }
  },
  methods: {
    checkLoaded() {
      this.loaded = (
        Object.keys(this.stats).length >= 0 &&
        Object.keys(this.sysinfo).length >= 0
      )
    },

    /**
     * @param number
     * @returns {string}
     */
    commify: function (number) {
      return number.toString().replace(/\B(?=(\d{3})+(?!\d))/g, ',')
    },

    loadSysInfo: function () {
      EqemuAdminClient.getSysInfo().then(response => {
        if (response) {
          this.sysinfo = response
        }

        this.checkLoaded()
      })
    },

    /**
     * @param number
     * @returns {string}
     */
    kFormatter: function (number) {
      return number > 999 ? (number / 1000).toFixed(1) + 'k' : number
    }

  }
}
</script>
