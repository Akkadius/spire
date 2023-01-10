<template>
  <div class="card">
    <div class="card-body">
      <div class="row align-items-center">
        <div class="col">
          <h6 class="header-pretitle">
            {{ pageName }}
          </h6>

          <h1 class="header-title">
            <span v-if="stats.server_name">{{ stats.server_name }}</span>
          </h1>
        </div>

        <div class="col-auto">
          <small class="text-muted text-uppercase">Launcher</small>
          <span
            :class="`badge badge-${stats.launcher_online ? 'success' : 'danger'} ml-3`"
            style="font-size: 12px"
          >{{ stats.launcher_online ? 'Online' : 'Offline' }}</span>
        </div>

        <div class="col-auto">
          <small class="text-muted text-uppercase">World</small>
          <span
            :class="`badge badge-${stats.world_online ? 'success' : 'danger'} ml-3`"
            style="font-size: 12px"
          >{{ stats.world_online ? 'Online' : 'Offline' }}</span>
        </div>

        <div class="col-auto">
          <small class="text-muted text-uppercase">UCS</small>
          <span
            :class="`badge badge-${stats.ucs_online ? 'success' : 'danger'} ml-3`"
            style="font-size: 12px"
          >{{ stats.ucs_online ? 'Online' : 'Offline' }}</span>
        </div>

        <div class="col-auto">
          <small class="text-muted text-uppercase">Zoneservers</small>
          <span class="h2 mb-0 ml-3">
            {{ stats && stats.zone_list && stats.zone_list.data ? stats.zone_list.data.length : 0 }}
          </span>
        </div>

        <div class="col-auto">
          <small class="text-muted text-uppercase">Players Online</small>
          <span class="h2 mb-0 ml-3">
            {{ stats && stats.client_list && stats.client_list.data ? stats.client_list.data.length : 0 }}
          </span>
        </div>

        <div class="col-auto">
          <server-process-button-component/>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import ServerProcessButtonComponent from "@/views/admin/components/ServerProcessButtonComponent.vue";
import {EventBus}                   from "@/app/event-bus/event-bus";
import Timer                        from "@/app/timer/timer";
import {SpireApi}                   from "@/app/api/spire-api";
import {OcculusClient}              from "@/app/api/eqemu-admin-client-occulus";

export default {
  name: "AdminHeader",
  components: { ServerProcessButtonComponent },
  data() {
    return {
      pageName: "",

      stats: {},
    }
  },
  beforeDestroy() {
    clearInterval(Timer.timer['server-stat-refresh'])

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

    Timer.timer['server-stat-refresh'] = setInterval(() => {
      if (!document.hidden) {
        this.loadServerStats()
      }
    }, 5000)

    // initial page name set
    if (this.$route.meta && this.$route.meta.title) {
      this.pageName = this.$route.meta.title
    }
  },
  methods: {

    async loadServerStats() {
      const r = await SpireApi.v1().get("eqemuserver/server-stats")
      if (r.status === 200) {
        this.stats = r.data
        this.$forceUpdate()

      }

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
        // case '1':
        //   this.$router.push(ROUTES.ROOT)
        //   break
        // case '2':
        //   this.$router.push(ROUTES.PLAYERS_ONLINE)
        //   break
        // case '3':
        //   this.$router.push(ROUTES.ZONESERVERS)
        //   break
        // case '4':
        //   this.$router.push(ROUTES.CONFIGURATION)
        //   break
        // case '5':
        //   this.$router.push(ROUTES.TOOLS_LOGS)
        //   break
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

<style scoped>

</style>
