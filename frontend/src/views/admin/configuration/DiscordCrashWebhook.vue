<template>
  <div>

    <div
      class="row justify-content-center"
      style="position: absolute; top: 0%; z-index: 9999999; width: 100%"
    >
      <div class="col-4">
        <info-error-banner
          style="width: 100%"
          :slim="true"
          :notification="notification"
          :error="error"
          @dismiss-error="error = ''"
          @dismiss-notification="notification = ''"
          class="mt-3"
        />
      </div>
    </div>

    <eq-window title="Crash Logs" v-if="Object.keys(serverConfig).length > 0">
      <div class="form-group col-md-12 mb-0 p-0">
        <label class="form-label">Crash Logs Webhook</label>
        <small class="d-block text-muted">
          When a process crash occurs, Occulus will send crash logs to this webhook
        </small>

        <input
          type="text"
          class="form-control mt-3"
          placeholder="https://discord.com/api/webhooks/xxx/xxx"
          v-model="serverConfig['web-admin'].discord.crash_log_webhook"
        />

        <button type="submit" class="btn btn-sm mt-3 btn-outline-warning ml-auto" @click="submitServerConfig()">
          <i class="fe fe-save"></i>
          Save
        </button>
      </div>
    </eq-window>

  </div>
</template>

<script>
import {OcculusClient}     from "@/app/api/eqemu-admin-client-occulus";
import EqWindow            from "@/components/eq-ui/EQWindow.vue";
import {DiscordWebhookApi} from "@/app/api/api/discord-webhook-api";
import {SpireApi}          from "@/app/api/spire-api";
import InfoErrorBanner     from "@/components/InfoErrorBanner.vue";

export default {
  components: { InfoErrorBanner, EqWindow },
  data() {
    return {
      serverConfig: {},
      loaded: false,

      // notification / errors
      notification: "",
      error: "",
    }
  },
  async created() {
    this.serverConfig = await OcculusClient.getServerConfig()

    this.loaded = true
  },
  methods: {

    submitServerConfig: async function () {
      const result = await OcculusClient.postServerConfig(this.serverConfig)

      if (result.success) {
        this.notification = "Saved crash webhook!"
      }
    }
  }
}
</script>
