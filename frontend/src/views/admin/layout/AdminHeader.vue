<template>
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
</template>

<script>
import ServerProcessButtonComponent from "@/views/admin/components/ServerProcessButtonComponent.vue";
import {EqemuAdminClient}           from "@/app/api/eqemu-admin-client-occulus";

export default {
  name: "AdminHeader",
  components: { ServerProcessButtonComponent },
  data() {
    return {
      stats: {},
    }
  },
  methods: {
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
  beforeDestroy() {
    window.removeEventListener('keypress', this.keypressHandler)
  },
  created() {
    EqemuAdminClient.getDashboardStats().then(r => {
      if (r) {
        this.stats = r
      }
    })

    window.addEventListener('keypress', this.keypressHandler)
  }
}
</script>

<style scoped>

</style>
