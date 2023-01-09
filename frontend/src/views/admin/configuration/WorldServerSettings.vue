<template>
  <div>

    <div class="card">
      <div class="card-body">
        <div class="row justify-content-between align-items-center">
          <div class="col-12 col-md-9 col-xl-7">
            <h2 class="mb-2">
              World Server Settings
            </h2>
            <p class="text-muted mb-md-0">
              Configure world server properties such as server name, connection settings...
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

            <b-tabs content-class="mt-5" fill>
              <b-tab title="Server Naming" active>

                <div class="form-row">
                  <div class="form-group col-md-6">
                    <label class="form-label">Server Long Name</label>
                    <input type="text" class="form-control" v-model="serverConfig.server.world.longname"/>
                    <small class="form-text text-muted mt-3">Displays on the Loginserver</small>
                  </div>

                  <div class="form-group col-md-6">
                    <label class="form-label">Server Short Name</label>
                    <input type="text" class="form-control" v-model="serverConfig.server.world.shortname"/>
                    <small class="form-text text-muted mt-3">Used in the client .ini configuration files</small>
                  </div>
                </div>

              </b-tab>

              <b-tab title="Key">
                <div class="form-group">
                  <label class="form-label">Server Key</label>

                  <div class="input-group">
                    <input type="text" class="form-control" v-model="serverConfig.server.world.key"/>
                    <span class="input-group-append">
              <button class="btn btn-primary" type="button" @click="generateRandomKey()"><i
                class="fa fa-history pr-1"
              ></i>
                  Generate Random Key
              </button>
            </span>
                  </div>

                  <small class="form-text text-muted mt-3">Used in inter-server communication and allows remote
                    zoneservers
                    to connect back
                  </small>
                </div>
              </b-tab>

              <b-tab title="Telnet / Websockets">

                <label class="pb-2 pt-2">

                  <div class="custom-control custom-checkbox checklist-control">
                    <input
                      class="custom-control-input" id="checklistOne" type="checkbox"
                      v-model="serverConfig.server.world.telnet.enabled"
                    />
                    <label class="custom-control-label" for="checklistOne"></label>
                    Telnet / Websockets Enabled <small class="ml-2 text-muted">(Keep this on or admin panel
                    functionality
                    will break)</small>
                  </div>

                </label>

                <hr class="my-4">

                <div class="form-row">
                  <div class="form-group col-md-6" v-show="serverConfig.server.world.telnet.enabled">
                    <label class="form-label">Telnet IP</label>
                    <input
                      type="text" class="form-control"
                      v-model="serverConfig.server.world.telnet.ip"
                    />
                    <small class="form-text text-muted mt-3">IP address to listen to telnet on. (0.0.0.0) works in
                      most scenarios
                    </small>
                  </div>

                  <div class="form-group col-md-6" v-show="serverConfig.server.world.telnet.enabled">
                    <label class="form-label">Telnet IP</label>
                    <input
                      type="text" class="form-control"
                      v-model="serverConfig.server.world.telnet.port"
                    />
                    <small class="form-text text-muted mt-3">Port for telnet to listen on. Keep this at 9000
                    </small>
                  </div>
                </div>

              </b-tab>

              <b-tab title="TCP Connections">
                <div class="form-row">
                  <div class="form-group col-md-6">
                    <label class="form-label">IP</label>
                    <input
                      type="text" class="form-control"
                      v-model="serverConfig.server.world.tcp.ip"
                    />
                    <small class="form-text text-muted mt-3">IP address to world to listen on. (0.0.0.0) works in
                      most
                      scenarios
                    </small>
                  </div>

                  <div class="form-group col-md-6">
                    <label class="form-label">Port</label>
                    <input
                      type="text" class="form-control"
                      v-model="serverConfig.server.world.tcp.port"
                    />
                    <small class="form-text text-muted mt-3">Port for world to listen on. Keep this at 9001
                    </small>
                  </div>
                </div>
              </b-tab>

              <b-tab title="Loginserver #1">
                <div class="form-row" v-if="serverConfig.server.world.loginserver1">
                  <div class="form-group col-md-4">
                    <label class="form-label">Loginserver Host</label>
                    <input
                      type="text" class="form-control"
                      v-model="serverConfig.server.world.loginserver1.host"
                    />
                    <small class="form-text text-muted mt-3">Loginserver host your server is connecting to</small>
                  </div>

                  <div class="form-group col-md-4">
                    <label class="form-label">Loginserver Port</label>
                    <input
                      type="text" class="form-control"
                      v-model="serverConfig.server.world.loginserver1.port"
                    />
                    <small class="form-text text-muted mt-3">Loginserver port your server is connecting to (usually
                      5998)
                    </small>
                  </div>

                  <div class="form-group col-md-4">
                    <label class="form-label">Legacy Network Connection</label>
                    <input
                      type="text" class="form-control"
                      v-model="serverConfig.server.world.loginserver1.legacy"
                    />
                    <small class="form-text text-muted mt-3">Used to determine if this is a legacy network
                      connection, needed for eqemulator.net
                    </small>
                  </div>

                  <div class="form-group col-md-6">
                    <label class="form-label">Loginserver Account</label>
                    <input
                      type="text" class="form-control"
                      v-model="serverConfig.server.world.loginserver1.account"
                    />
                    <small class="form-text text-muted mt-3">Used to authenticate your server as a registered server
                      with the connecting loginserver <br> (Not Required)
                    </small>
                  </div>

                  <div class="form-group col-md-6">
                    <label class="form-label">Loginserver Password</label>
                    <div class="input-group">
                      <input
                        :type="passwordFieldType" class="form-control"
                        data-lpignore="true"
                        v-model="serverConfig.server.world.loginserver1.password"
                      />
                      <span class="input-group-append">
                <button class="btn btn-primary" type="button" @click="switchPasswordVisibility()"><i
                  class="fa fa-history pr-1"
                ></i>
                    Show / Hide
                </button>
              </span>
                    </div>
                  </div>

                </div>
              </b-tab>

              <b-tab title="Loginserver #2">
                <div class="form-row" v-if="serverConfig.server.world.loginserver2">
                  <div class="form-group col-md-4">
                    <label class="form-label">Loginserver Host</label>
                    <input
                      type="text" class="form-control"
                      v-model="serverConfig.server.world.loginserver2.host"
                    />
                    <small class="form-text text-muted mt-3">Loginserver host your server is connecting to</small>
                  </div>

                  <div class="form-group col-md-4">
                    <label class="form-label">Loginserver Port</label>
                    <input
                      type="text" class="form-control"
                      v-model="serverConfig.server.world.loginserver2.port"
                    />
                    <small class="form-text text-muted mt-3">Loginserver port your server is connecting to (usually
                      5998)
                    </small>
                  </div>

                  <div class="form-group col-md-4">
                    <label class="form-label">Legacy Network Connection</label>
                    <input
                      type="text" class="form-control"
                      v-model="serverConfig.server.world.loginserver2.legacy"
                    />
                    <small class="form-text text-muted mt-3">Used to determine if this is a legacy network
                      connection, needed for eqemulator.net
                    </small>
                  </div>

                  <div class="form-group col-md-6">
                    <label class="form-label">Loginserver Account</label>
                    <input
                      type="text" class="form-control"
                      v-model="serverConfig.server.world.loginserver2.account"
                    />
                    <small class="form-text text-muted mt-3">Used to authenticate your server as a registered server
                      with the connecting loginserver <br> (Not Required)
                    </small>
                  </div>

                  <div class="form-group col-md-6">
                    <label class="form-label">Loginserver Password</label>
                    <div class="input-group">
                      <input
                        :type="passwordFieldType" class="form-control"
                        data-lpignore="true"
                        v-model="serverConfig.server.world.loginserver2.password"
                      />
                      <span class="input-group-append">
                    <button class="btn btn-primary" type="button" @click="switchPasswordVisibility()"><i
                      class="fa fa-eye pr-1"
                    ></i>
                        Show / Hide
                    </button>
                  </span>
                    </div>
                  </div>

                </div>
              </b-tab>

              <b-tab title="Loginserver #3">
                <div class="form-row" v-if="serverConfig.server.world.loginserver3">
                  <div class="form-group col-md-4">
                    <label class="form-label">Loginserver Host</label>
                    <input
                      type="text" class="form-control"
                      v-model="serverConfig.server.world.loginserver3.host"
                    />
                    <small class="form-text text-muted mt-3">Loginserver host your server is connecting to</small>
                  </div>

                  <div class="form-group col-md-4">
                    <label class="form-label">Loginserver Port</label>
                    <input
                      type="text" class="form-control"
                      v-model="serverConfig.server.world.loginserver3.port"
                    />
                    <small class="form-text text-muted mt-3">Loginserver port your server is connecting to (usually
                      5998)
                    </small>
                  </div>

                  <div class="form-group col-md-4">
                    <label class="form-label">Legacy Network Connection</label>
                    <input
                      type="text" class="form-control"
                      v-model="serverConfig.server.world.loginserver3.legacy"
                    />
                    <small class="form-text text-muted mt-3">Used to determine if this is a legacy network
                      connection, needed for eqemulator.net
                    </small>
                  </div>

                  <div class="form-group col-md-6">
                    <label class="form-label">Loginserver Account</label>
                    <input
                      type="text" class="form-control"
                      v-model="serverConfig.server.world.loginserver3.account"
                    />
                    <small class="form-text text-muted mt-3">Used to authenticate your server as a registered server
                      with the connecting loginserver <br> (Not Required)
                    </small>
                  </div>

                  <div class="form-group col-md-6">
                    <label class="form-label">Loginserver Password</label>
                    <div class="input-group">
                      <input
                        :type="passwordFieldType" class="form-control"
                        v-model="serverConfig.server.world.loginserver3.password"
                      />
                      <span class="input-group-append">
                      <button class="btn btn-primary" type="button" @click="switchPasswordVisibility()"><i
                        class="fa fa-eye pr-1"
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

    const loginServerConfigModel = {
      'account': '',
      'password': '',
      'legacy': '0',
      'host': '',
      'port': '5998'
    }

    if (typeof this.serverConfig.server.world.loginserver1 === 'undefined') {
      this.serverConfig.server.world.loginserver1 = loginServerConfigModel
    }

    if (typeof this.serverConfig.server.world.loginserver2 === 'undefined') {
      this.serverConfig.server.world.loginserver2 = loginServerConfigModel
    }

    if (typeof this.serverConfig.server.world.loginserver3 === 'undefined') {
      this.serverConfig.server.world.loginserver3 = loginServerConfigModel
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
    generateRandomKey: function () {
      let result           = ''
      let characters       = 'ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789'
      let charactersLength = characters.length
      for (var i = 0; i < 40; i++) {
        result += characters.charAt(Math.floor(Math.random() * charactersLength))
      }

      this.serverConfig.server.world.key = result
    },
    switchPasswordVisibility() {
      this.passwordFieldType = this.passwordFieldType === 'password' ? 'text' : 'password'
    }
  }
}
</script>
