<template>
  <b-card
    header="Launcher Options"
  >
    <b-card-text>
      <div class="mb-3">
        <b-form-checkbox
          v-model="launcher.runSharedMemory"
          name="check-button"
          switch
          @change="saveLauncherOptions()"
        >
          Run Shared Memory
        </b-form-checkbox>
      </div>
      <div class="mb-3">
        <b-form-checkbox
          v-model="launcher.runLoginserver"
          name="check-button"
          switch
          class="custom-control custom-switch"
          @change="saveLauncherOptions()"
        >
          Run Loginserver
        </b-form-checkbox>
      </div>
      <div class="mb-3">
        <b-form-checkbox
          v-model="launcher.runQueryServ"
          name="check-button"
          switch
          @change="saveLauncherOptions()"
        >
          Run QueryServ
        </b-form-checkbox>
      </div>
    </b-card-text>
  </b-card>
</template>

<script>

  import {EqemuAdminClient} from "@/app/api/eqemu-admin-client-occulus";

  export default {
    name: 'LauncherOptions',
    props: ['launcherConfig'],
    data () {
      return {
        launcher: {
          runSharedMemory: false,
          runLoginserver: false,
          runQueryServ: false
        }
      }
    },
    created() {
      this.launcher = this.launcherConfig
    },
    watch: {
      launcherConfig: function (newValue) {
        this.launcher = newValue
      }
    },
    methods: {
      saveLauncherOptions () {
        EqemuAdminClient.postLauncherConfig(this.launcher)
      }
    }
  }
</script>

<style scoped>

</style>
