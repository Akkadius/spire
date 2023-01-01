<template>
  <div>

    <div class="card">
      <div class="card-body">
        <div class="row justify-content-between align-items-center">
          <div class="col-12 col-md-9 col-xl-7">
            <h2 class="mb-2">
              Discord Settings
            </h2>
            <p class="text-muted mb-md-0">
              Configure Discord Webhooks
            </p>
          </div>

          <div class="col-12 col-md-auto">
            <button type="submit" class="btn btn-primary ml-auto" @click="submitServerConfig()">
              <i class="fe fe-save"></i>
              Save
            </button>
          </div>
        </div>
      </div>
    </div>


    <div class="card">
      <div class="card-body">
        <div class="row">
          <div class="col-lg-12" v-if="Object.keys(serverConfig).length > 0">
            <div class="form-row">
              <div class="form-group col-md-12">
                <label class="form-label">Crash Logs Webhook</label>
                <input
                  type="text" class="form-control" placeholder="https://discord.com/api/webhooks/xxx/xxx"
                  v-model="serverConfig['web-admin'].discord.crash_log_webhook"
                />
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>

  </div>
</template>

<script>
import {EqemuAdminClient} from "@/app/api/eqemu-admin-client-occulus";

export default {
  data() {
    return {
      serverConfig: {},
      loaded: false
    }
  },
  async created() {
    this.serverConfig = await EqemuAdminClient.getServerConfig()

    this.loaded = true
  },
  methods: {
    submitServerConfig: async function () {
      const result = await EqemuAdminClient.postServerConfig(this.serverConfig)

      if (result.success) {
        this.$bvToast.toast(
          result.success,
          {
            title: "Configuration saved!",
            toaster: 'b-toaster-bottom-center',
            autoHideDelay: 3000,
            solid: true,
            appendToast: false
          }
        )
      }
    }
  }
}
</script>
