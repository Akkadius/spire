<template>
  <div class="dropdown">

    <a
      href="#"
      style="color: white"
      class="dropdown-toggle btn btn-primary lift"
      role="button"
      data-toggle="dropdown"
      aria-haspopup="true"
      aria-expanded="false"
    >
      <i class="fe fe-power"></i> Power
    </a>

    <div class="dropdown-menu dropdown-menu-right">
      <a href="#" @click="startServerModal" class="dropdown-item">Power On [p]</a>
      <a href="#" @click="stopServerModal" class="dropdown-item">Power Off [s]</a>
      <a href="#" @click="restartServerModal" class="dropdown-item">Restart [r]</a>
      <a href="#" @click="cancelServerRestartModal" class="dropdown-item">Cancel Restart [c]</a>
    </div>

    <!-- Start Server -->
    <b-modal
      centered
      no-fade
      id="start-server-modal"
      title="Start Server"
      header-text-variant="dark"
      body-text-variant="dark"
      ok-title="Start Server"
      @ok="startServer"
    >

      <p class="pt-3 pb-3">Are you sure you want to start server?</p>

      <LauncherOptions :launcherConfig="launcher"></LauncherOptions>

    </b-modal>

    <!-- Stop Server -->
    <b-modal
      centered
      no-fade
      id="stop-server-modal"
      title="Stop Server"
      header-text-variant="dark"
      body-text-variant="dark"
      ok-title="Stop Server"
      @ok="stopServer"
    >
      Are you sure you want to stop your server?
    </b-modal>

    <!-- Restart Server -->
    <b-modal
      centered
      no-fade
      id="restart-server-modal"
      title="Restart Server"
      header-text-variant="dark"
      body-text-variant="dark"
      ok-title="Restart Server"
      @ok="restartServer"
    >

      <p class="pt-3 pb-3">Are you sure you want to restart your server?</p>

      <LauncherOptions :launcherConfig="launcher"></LauncherOptions>

      <b-card
        header="Restart Announcement Warning"
      >
        <b-card-text>
          <b-form-radio v-model="delayedRestart" name="some-radios" value="0">None</b-form-radio>
          <b-form-radio v-model="delayedRestart" name="some-radios" :value="5 * 60">5 Minute(s)</b-form-radio>
          <b-form-radio v-model="delayedRestart" name="some-radios" :value="10 * 60">10 Minute(s)</b-form-radio>
          <b-form-radio v-model="delayedRestart" name="some-radios" :value="15 * 60">15 Minute(s)</b-form-radio>
          <b-form-radio v-model="delayedRestart" name="some-radios" :value="30 * 60">30 Minute(s)</b-form-radio>
        </b-card-text>
      </b-card>

    </b-modal>

    <!-- Cancel Restart Server -->
    <b-modal
      centered
      no-fade
      id="cancel-restart-server-modal"
      title="Cancel Server Restart"
      header-text-variant="dark"
      body-text-variant="dark"
      ok-title="Cancel Restart"
      @ok="cancelRestartServer"
    >
      Are you sure you want to cancel your timed restart?

    </b-modal>

  </div>
</template>

<script>

import {EqemuAdminClient} from "@/app/api/eqemu-admin-client-occulus";
import {EventBus}         from "@/app/event-bus/event-bus";
import LauncherOptions    from "@/views/admin/components/LauncherOptions.vue";

export default {
  name: 'ServerProcessButtonComponent',
  components: {
    LauncherOptions
  },
  data() {
    return {
      delayedRestart: 0,
      launcher: null,
    }
  },
  async mounted() {
    const result  = await EqemuAdminClient.getLauncherConfig();
    this.launcher = result.data;
  },
  methods: {
    /**
     * Start
     */
    startServerModal() {
      this.$root.$emit('bv::show::modal', 'start-server-modal')
    },
    startServer() {
      EqemuAdminClient.startServer();
      this.notify("Server Start", "Server is starting!");
      this.notifyProcessChange()
    },

    /**
     * Stop
     */
    stopServerModal() {
      this.$root.$emit('bv::show::modal', 'stop-server-modal')
    },
    stopServer() {
      EqemuAdminClient.stopServer();
      this.notify("Server Stopped", "Server has been stopped!");
      this.notifyProcessChange()
    },

    /**
     * Restart
     */
    restartServerModal() {
      this.$root.$emit('bv::show::modal', 'restart-server-modal')
      this.delayedRestart = 0;
    },
    restartServer() {
      EqemuAdminClient.restartServer({ timer: this.delayedRestart });
      if (this.delayedRestart > 0) {
        this.notify("Server Restarted", "Server restart warning timer has been started!");
      } else {
        this.notify("Server Restarted", "Server has been restarted!");
      }
      this.notifyProcessChange()
    },

    /**
     * Cancel
     */
    cancelServerRestartModal() {
      this.$root.$emit('bv::show::modal', 'cancel-restart-server-modal')
      this.delayedRestart = 0;
    },
    cancelRestartServer() {
      EqemuAdminClient.cancelRestartServer({ cancel: 1 });
      this.notify("Server Restart Cancelled", "Server restart has been cancelled");
    },

    notify(title, message) {
      this.$bvToast.toast(message, {
        title: title,
        solid: true
      })
    },

    notifyProcessChange() {
      setTimeout(function () {
        EventBus.$emit('process-change');
      }, 1000);

      setTimeout(function () {
        EventBus.$emit('process-change');
      }, 3000);
    },

  },
}
</script>

<style scoped>

</style>
