<template>
  <div>
    <eq-window title="Crash Logs" v-if="Object.keys(serverConfig).length > 0">

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
import EqWindow        from "@/components/eq-ui/EQWindow.vue";
import InfoErrorBanner from "@/components/InfoErrorBanner.vue";
import {SpireApi}      from "@/app/api/spire-api";

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
    try {
      const r = await SpireApi.v1().get('admin/serverconfig')
      if (r.status === 200) {
        this.serverConfig = r.data
      }
    } catch (e) {
      // error notify
      if (e.response && e.response.data && e.response.data.error) {
        this.error = e.response.data.error
      }
    }

    if (!this.serverConfig['web-admin'].discord) {
      this.serverConfig['web-admin'].discord = {};
      if (!this.serverConfig['web-admin'].discord.crash_log_webhook) {
        this.serverConfig['web-admin'].discord.crash_log_webhook = "";
      }
    }

    this.loaded = true
  },
  methods: {

    submitServerConfig: async function () {

      try {
        const r = await SpireApi.v1().post('admin/serverconfig', this.serverConfig)
        if (r.status === 200) {
          this.notification = "Saved crash webhook!"
        }
      } catch (e) {
        // error notify
        if (e.response && e.response.data && e.response.data.error) {
          this.error = e.response.data.error
        }
      }

    }
  }
}
</script>
