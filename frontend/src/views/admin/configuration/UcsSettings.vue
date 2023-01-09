<template>
  <div>

    <div class="card">
      <div class="card-body">
        <div class="row justify-content-between align-items-center">
          <div class="col-12 col-md-9 col-xl-7">
            <h2 class="mb-2">
              Universal Chat Service
            </h2>
            <p class="text-muted mb-md-0">
              Mail and Chat server configuration properties
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
              <div class="form-group col-md-6">
                <label class="form-label">Chatserver Host</label>
                <input
                  type="text" class="form-control" placeholder="0.0.0.0"
                  v-model="serverConfig.server.chatserver.host"
                />
              </div>

              <div class="form-group col-md-6">
                <label class="form-label">Chatserver Port</label>
                <input type="text" class="form-control" v-model="serverConfig.server.chatserver.port"/>
              </div>
            </div>

            <div class="form-row">
              <div class="form-group col-md-6">
                <label class="form-label">Mailserver Host</label>
                <input
                  type="text" class="form-control" placeholder="0.0.0.0"
                  v-model="serverConfig.server.mailserver.host"
                />
              </div>

              <div class="form-group col-md-6">
                <label class="form-label">Mailserver Port</label>
                <input type="text" class="form-control" v-model="serverConfig.server.mailserver.port"/>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>

  </div>
</template>

<script>
import {OcculusClient} from "@/app/api/eqemu-admin-client-occulus";

export default {
  data() {
    return {
      serverConfig: {},
      passwordFieldType: 'password',
      loaded: false
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
