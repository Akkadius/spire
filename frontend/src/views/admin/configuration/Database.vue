<template>
  <div class="row">
    <div class="col-lg-12" v-if="Object.keys(serverConfig).length > 0">

      <div class="card">
        <div class="card-body">
          <div class="row justify-content-between align-items-center">
            <div class="col-12 col-md-9 col-xl-7">
              <h2 class="mb-2">
                Database Connection Properties
              </h2>
              <p class="text-muted mb-md-0">
                Configure the server's different types of MySQL database connections
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
          <b-tabs content-class="mt-5" fill>
            <b-tab title="Game" active>
              <div class="form-row">
                <div class="form-group col-md-4">
                  <label class="form-label">Database Name</label>
                  <input type="text" class="form-control" v-model="serverConfig.server.database.db"/>
                </div>

                <div class="form-group col-md-4">
                  <label class="form-label">Database Host</label>
                  <input type="text" class="form-control" v-model="serverConfig.server.database.host"/>
                </div>

                <div class="form-group col-md-4">
                  <label class="form-label">Database Port</label>
                  <input type="text" class="form-control" v-model="serverConfig.server.database.port"/>
                </div>
              </div>

              <div class="form-row">
                <div class="form-group col-md-6">
                  <label class="form-label">Database Username</label>
                  <input type="text" class="form-control" v-model="serverConfig.server.database.username"/>
                </div>

                <div class="form-group col-md-6">
                  <label class="form-label">Database Password</label>

                  <div class="input-group">
                    <input
                      :type="passwordFieldType" class="form-control"
                      v-model="serverConfig.server.database.password"
                    />
                    <span class="input-group-append">
                  <button class="btn btn-primary" type="button" @click="switchPasswordVisibility()"><i
                    class="fa fa-eye"
                  ></i>
                      Show / Hide
                  </button>
                </span>
                  </div>
                </div>
              </div>

            </b-tab>
            <b-tab title="Game (Content Database)">
              <b-alert show variant="primary">
                <i class="fe fe-info"></i>
                This connection is optional and is described in more detail
                <a
                  href="https://eqemu.gitbook.io/server/categories/database/multi-tenancy"
                  target="multi-tenancy-doc"
                  style="color:white"
                >(here)</a>; leave blank to use default connction
              </b-alert>
              <div class="form-row">
                <div class="form-group col-md-4">
                  <label class="form-label">Database Name</label>
                  <input type="text" class="form-control" v-model="serverConfig.server.content_database.db"/>
                </div>

                <div class="form-group col-md-4">
                  <label class="form-label">Database Host</label>
                  <input type="text" class="form-control" v-model="serverConfig.server.content_database.host"/>
                </div>

                <div class="form-group col-md-4">
                  <label class="form-label">Database Port</label>
                  <input type="text" class="form-control" v-model="serverConfig.server.content_database.port"/>
                </div>
              </div>

              <div class="form-row">
                <div class="form-group col-md-6">
                  <label class="form-label">Database Username</label>
                  <input type="text" class="form-control" v-model="serverConfig.server.content_database.username"/>
                </div>

                <div class="form-group col-md-6">
                  <label class="form-label">Database Password</label>

                  <div class="input-group">
                    <input
                      :type="passwordFieldType" class="form-control"
                      v-model="serverConfig.server.content_database.password"
                    />
                    <span class="input-group-append">
                  <button class="btn btn-primary" type="button" @click="switchPasswordVisibility()"><i
                    class="fa fa-eye"
                  ></i>
                      Show / Hide
                  </button>
                </span>
                  </div>
                </div>
              </div>

            </b-tab>
            <b-tab title="QueryServer">

              <div class="form-row">
                <div class="form-group col-md-4">
                  <label class="form-label">Database Name</label>
                  <input type="text" class="form-control" v-model="serverConfig.server.qsdatabase.db"/>
                </div>

                <div class="form-group col-md-4">
                  <label class="form-label">Database Host</label>
                  <input type="text" class="form-control" v-model="serverConfig.server.qsdatabase.host"/>
                </div>

                <div class="form-group col-md-4">
                  <label class="form-label">Database Port</label>
                  <input type="text" class="form-control" v-model="serverConfig.server.qsdatabase.port"/>
                </div>
              </div>

              <div class="form-row">
                <div class="form-group col-md-6">
                  <label class="form-label">Database Username</label>
                  <input type="text" class="form-control" v-model="serverConfig.server.qsdatabase.username"/>
                </div>

                <div class="form-group col-md-6">
                  <label class="form-label">Database Password</label>

                  <div class="input-group">
                    <input
                      :type="passwordFieldType" class="form-control"
                      v-model="serverConfig.server.qsdatabase.password"
                    />
                    <span class="input-group-append">
                  <button class="btn btn-primary" type="button" @click="switchPasswordVisibility()"><i
                    class="fa fa-eye"
                  ></i>
                      Show / Hide
                  </button>
                </span>
                  </div>
                </div>
              </div>
            </b-tab>
          </b-tabs>
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

    if (!this.serverConfig.server.content_database) {
      this.serverConfig.server.content_database = {
        db: "",
        host: "",
        port: "",
        username: "",
        password: "",
      }
    }

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
    },
    switchPasswordVisibility() {
      this.passwordFieldType = this.passwordFieldType === 'password' ? 'text' : 'password'
    }
  }
}
</script>
