<template>
  <div>
    <div class="card">
      <div class="card-body">
        <div class="row justify-content-between align-items-center">
          <div class="col-12 col-md-9 col-xl-7">
            <h2 class="mb-2">
              Zone Server Settings
            </h2>
            <p class="text-muted mb-md-0">
              Configure zone server properties
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
              <div class="form-group col-md-4">
                <label class="form-label">Default Player Account Status</label>
                <input type="text" class="form-control" v-model="serverConfig.server.zones.defaultstatus"/>
                <small class="form-text text-muted mt-3">This is the default status that new accounts are created
                  with, this most likely should be 0
                </small>
              </div>

              <div class="form-group col-md-4">
                <label class="form-label">Zone Port Range Start</label>
                <input
                  type="number" class="form-control" v-model="serverConfig.server.zones.ports.low"
                  min="7000" max="7500"
                />
                <small class="form-text text-muted mt-3">Port range start for zone assignment (7000-7500)</small>
              </div>

              <div class="form-group col-md-4">
                <label class="form-label">Zone Port Range End</label>
                <input
                  type="number" class="form-control" v-model="serverConfig.server.zones.ports.high"
                  min="7000" max="7500"
                />
                <small class="form-text text-muted mt-3">Port range start for zone assignment (7000-7500)</small>
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
      passwordFieldType: 'password',
      loaded: false
    }
  },
  async created() {
    this.serverConfig = await EqemuAdminClient.getServerConfig()
    this.loaded       = true
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
