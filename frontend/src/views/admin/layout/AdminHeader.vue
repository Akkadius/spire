<template>
  <div class="card">
    <div class="card-body">
      <div class="row align-items-center">
        <div class="col">
          <h6 class="header-pretitle">
            {{pageName}}
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
</template>

<script>
import ServerProcessButtonComponent from "@/views/admin/components/ServerProcessButtonComponent.vue";
import {OcculusClient}              from "@/app/api/eqemu-admin-client-occulus";
import {EventBus}                   from "@/app/event-bus/event-bus";

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
    window.removeEventListener('keypress', this.keypressHandler)
    EventBus.$off("ROUTE_CHANGE", this.handleRouteChange);
  },
  created() {
    EventBus.$on("ROUTE_CHANGE", this.handleRouteChange);

    OcculusClient.getDashboardStats().then(r => {
      if (r) {
        this.stats = r
      }
    })

    window.addEventListener('keypress', this.keypressHandler)

    // initial page name set
    if (this.$route.meta && this.$route.meta.title) {
      this.pageName = this.$route.meta.title
    }
  },
  methods: {
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
