<template>
  <eq-window
    title="Server Configuration"
  >
    <eq-tabs
      :selected="tabSelected"
      @on-selected="tabSelected = $event; updateQueryState()"
      v-if="Object.keys(config).length > 0"
    >
      <eq-tab
        class="fade-in"
        :name="`World Server`"
      >
        <eq-tabs content-class="mt-5" fill>
          <eq-tab name="Server Naming" :selected="true">

            <div class="form-row">
              <div class="form-group col-md-6">
                <label class="form-label">Server Long Name</label>
                <input type="text" class="form-control" v-model="config.server.world.longname"/>
                <small class="form-text text-muted mt-3">Displays on the Loginserver</small>
              </div>

              <div class="form-group col-md-6">
                <label class="form-label">Server Short Name</label>
                <input type="text" class="form-control" v-model="config.server.world.shortname"/>
                <small class="form-text text-muted mt-3">Used in the client .ini configuration files</small>
              </div>
            </div>

          </eq-tab>

          <eq-tab name="Networking">
            <div class="form-row">
              <div class="form-group col-md-6">
                <label class="form-label">Public Address</label>
                <input type="text" class="form-control" v-model="config.server.world.address"/>
                <small class="form-text mt-3 eq-alert">If your server is on the internet, you will need to specify your
                  public address. When on a LAN behind a NAT you will need both this and <b>localaddress</b> set
                </small>
              </div>

              <div class="form-group col-md-6">
                <label class="form-label">Local Address</label>
                <input type="text" class="form-control" v-model="config.server.world.localaddress"/>
                <small class="form-text mt-3 eq-alert">If you are on a LAN you will need this address set to the local
                  address of your host so that others on your network can properly be routed to your gameserver. Do not
                  use 127.0.0.1</small>
              </div>
            </div>
          </eq-tab>

          <eq-tab name="Key">
            <div class="form-group">
              <label class="form-label">Server Key</label>

              <div class="input-group">
                <input type="text" class="form-control" v-model="config.server.world.key"/>
                <span class="input-group-append">
              <button class="btn btn-light btn-sm" type="button" @click="generateRandomKey()"><i
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
          </eq-tab>

          <eq-tab name="Telnet / Websockets">

            <label class="pb-2 pt-2">

              <div class="custom-control custom-checkbox checklist-control">
                <input
                  class="custom-control-input" id="checklistOne" type="checkbox"
                  v-model="config.server.world.telnet.enabled"
                />
                <label class="custom-control-label" for="checklistOne"></label>
                Telnet / Websockets Enabled
                <small class="ml-2 text-muted">
                  (Keep this on or admin panel functionality will break)
                </small>
              </div>

            </label>

            <div class="form-row">
              <div class="form-group col-md-6" v-show="config.server.world.telnet.enabled">
                <label class="form-label">Telnet IP</label>
                <input
                  type="text" class="form-control"
                  v-model="config.server.world.telnet.ip"
                />
                <small class="form-text text-muted mt-3">
                  IP address to listen to telnet on. (0.0.0.0) works in most scenarios
                </small>
              </div>

              <div class="form-group col-md-6" v-show="config.server.world.telnet.enabled">
                <label class="form-label">Telnet IP</label>
                <input
                  type="text" class="form-control"
                  v-model="config.server.world.telnet.port"
                />
                <small class="form-text text-muted mt-3">
                  Port for telnet to listen on. Keep this at 9000
                </small>
              </div>
            </div>

          </eq-tab>

          <eq-tab name="TCP Connections">
            <div class="form-row">
              <div class="form-group col-md-6">
                <label class="form-label">IP</label>
                <input
                  type="text" class="form-control"
                  v-model="config.server.world.tcp.ip"
                />
                <small class="form-text text-muted mt-3">IP address to world to listen on. (0.0.0.0) works in most
                  scenarios</small>
              </div>

              <div class="form-group col-md-6">
                <label class="form-label">Port</label>
                <input
                  type="text" class="form-control"
                  v-model="config.server.world.tcp.port"
                />
                <small class="form-text text-muted mt-3">Port for world to listen on. Keep this at 9001</small>
              </div>
            </div>
          </eq-tab>

          <eq-tab
            v-for="i in maxLoginServers"
            :key="i"
            :name="`Loginserver #${i}`"
          >
            <div class="form-row">
              <div class="form-group col-md-4">
                <label class="form-label">Loginserver Host</label>
                <input
                  type="text" class="form-control"
                  v-model="config.server.world['loginserver' + i].host"
                />
                <small class="form-text text-muted mt-3">Loginserver host your server is connecting to</small>
              </div>

              <div class="form-group col-md-4">
                <label class="form-label">Loginserver Port</label>
                <input
                  type="text" class="form-control"
                  v-model="config.server.world['loginserver' + i].port"
                />
                <small class="form-text text-muted mt-3">Loginserver port your server is connecting to (usually 5998)</small>
              </div>

              <div class="form-group col-md-4">
                <label class="form-label">Legacy Network Connection</label>
                <input
                  type="text" class="form-control"
                  v-model="config.server.world['loginserver' + i].legacy"
                />
                <small class="form-text text-muted mt-3">Used to determine if this is a legacy network connection</small>
              </div>

              <div class="form-group col-md-6">
                <label class="form-label">Loginserver Account</label>
                <input
                  type="text" class="form-control"
                  v-model="config.server.world['loginserver' + i].account"
                />
                <small class="form-text text-muted mt-3">Used to authenticate your server as a registered server</small>
              </div>

              <div class="form-group col-md-6">
                <label class="form-label">Loginserver Password</label>
                <div class="input-group">
                  <input
                    :type="passwordFieldType" class="form-control"
                    v-model="config.server.world['loginserver' + i].password"
                  />
                  <span class="input-group-append">
                    <button class="btn btn-light btn-sm" type="button" @click="switchPasswordVisibility()">
                      <i class="fa fa-eye pr-1"></i>
                      Show / Hide
                    </button>
                  </span>
                </div>
              </div>
            </div>
          </eq-tab>


        </eq-tabs>
      </eq-tab>

      <eq-tab class="fade-in" name="Zone Server">
        <div class="form-row">
          <div class="form-group col-md-4">
            <label class="form-label">Default Player Account Status</label>
            <input type="text" class="form-control" v-model="config.server.zones.defaultstatus"/>
            <small class="form-text text-muted mt-3">This is the default status that new accounts are created
              with, this most likely should be 0
            </small>
          </div>

          <div class="form-group col-md-4">
            <label class="form-label">Zone Port Range Start</label>
            <input
              type="number" class="form-control" v-model="config.server.zones.ports.low"
              min="7000" max="7500"
            />
            <small class="form-text text-muted mt-3">Port range start for zone assignment (7000-7500)</small>
          </div>

          <div class="form-group col-md-4">
            <label class="form-label">Zone Port Range End</label>
            <input
              type="number" class="form-control" v-model="config.server.zones.ports.high"
              min="7000" max="7500"
            />
            <small class="form-text text-muted mt-3">Port range start for zone assignment (7000-7500)</small>
          </div>
        </div>
      </eq-tab>

      <eq-tab class="fade-in" name="UCS">

        <div class="mb-3">
          Universal Chat Service (UCS) is a service that allows you to connect to the in game chat server. It also runs
          mail services.
        </div>

        <div class="form-row">
          <div class="form-group col-md-6">
            <label class="form-label">Host</label>
            <input
              type="text" class="form-control" placeholder="0.0.0.0"
              v-model="config.server.ucs.host"
            />
          </div>

          <div class="form-group col-md-6">
            <label class="form-label">Port</label>
            <input type="text" class="form-control" v-model="config.server.ucs.port"/>
          </div>
        </div>
      </eq-tab>

      <eq-tab
        class="fade-in"
        :name="`Database`"
      >
        <eq-tabs>
          <eq-tab name="Game" :selected="true">
            <div class="form-row">
              <div class="form-group col-md-4">
                <label class="form-label">Database Name</label>
                <input type="text" class="form-control" v-model="config.server.database.db"/>
              </div>

              <div class="form-group col-md-4">
                <label class="form-label">Database Host</label>
                <input type="text" class="form-control" v-model="config.server.database.host"/>
              </div>

              <div class="form-group col-md-4">
                <label class="form-label">Database Port</label>
                <input type="text" class="form-control" v-model="config.server.database.port"/>
              </div>
            </div>

            <div class="form-row">
              <div class="form-group col-md-6">
                <label class="form-label">Database Username</label>
                <input type="text" class="form-control" v-model="config.server.database.username"/>
              </div>

              <div class="form-group col-md-6">
                <label class="form-label">Database Password</label>

                <div class="input-group">
                  <input
                    :type="passwordFieldType" class="form-control"
                    v-model="config.server.database.password"
                  />
                  <span class="input-group-append">
                  <button class="btn btn-light btn-sm" type="button" @click="switchPasswordVisibility()"><i
                    class="fa fa-eye"
                  ></i>
                      Show / Hide
                  </button>
                </span>
                </div>
              </div>
            </div>

          </eq-tab>
          <eq-tab name="Game (Content Database)">
            <div class="eq-alert">
              <i class="fe fe-info"></i>
              This connection is optional and is described in more detail
              <a
                href="https://docs.eqemu.io/server/database/multi-tenancy/"
                target="multi-tenancy-doc"
              >(here)</a> leave blank to use default connection
            </div>

            <div class="form-row mt-3">
              <div class="form-group col-md-4">
                <label class="form-label">Database Name</label>
                <input type="text" class="form-control" v-model="config.server.content_database.db"/>
              </div>

              <div class="form-group col-md-4">
                <label class="form-label">Database Host</label>
                <input type="text" class="form-control" v-model="config.server.content_database.host"/>
              </div>

              <div class="form-group col-md-4">
                <label class="form-label">Database Port</label>
                <input type="text" class="form-control" v-model="config.server.content_database.port"/>
              </div>
            </div>

            <div class="form-row">
              <div class="form-group col-md-6">
                <label class="form-label">Database Username</label>
                <input type="text" class="form-control" v-model="config.server.content_database.username"/>
              </div>

              <div class="form-group col-md-6">
                <label class="form-label">Database Password</label>

                <div class="input-group">
                  <input
                    :type="passwordFieldType" class="form-control"
                    v-model="config.server.content_database.password"
                  />
                  <span class="input-group-append">
                  <button class="btn btn-light btn-sm" type="button" @click="switchPasswordVisibility()"><i
                    class="fa fa-eye"
                  ></i>
                      Show / Hide
                  </button>
                </span>
                </div>
              </div>
            </div>

          </eq-tab>
          <eq-tab name="Logs">

            <div class="eq-alert">
              <i class="fe fe-info"></i>
              This service and connection is optional and is currently known as queryserver
            </div>

            <div class="form-row mt-3">
              <div class="form-group col-md-4">
                <label class="form-label">Database Name</label>
                <input type="text" class="form-control" v-model="config.server.qsdatabase.db"/>
              </div>

              <div class="form-group col-md-4">
                <label class="form-label">Database Host</label>
                <input type="text" class="form-control" v-model="config.server.qsdatabase.host"/>
              </div>

              <div class="form-group col-md-4">
                <label class="form-label">Database Port</label>
                <input type="text" class="form-control" v-model="config.server.qsdatabase.port"/>
              </div>
            </div>

            <div class="form-row">
              <div class="form-group col-md-6">
                <label class="form-label">Database Username</label>
                <input type="text" class="form-control" v-model="config.server.qsdatabase.username"/>
              </div>

              <div class="form-group col-md-6">
                <label class="form-label">Database Password</label>

                <div class="input-group">
                  <input
                    :type="passwordFieldType" class="form-control"
                    v-model="config.server.qsdatabase.password"
                  />
                  <span class="input-group-append">
                  <button class="btn btn-light btn-sm" type="button" @click="switchPasswordVisibility()"><i
                    class="fa fa-eye"
                  ></i>
                      Show / Hide
                  </button>
                </span>
                </div>
              </div>
            </div>
          </eq-tab>
        </eq-tabs>


      </eq-tab>

      <div class="row">
        <div class="justify-content-center col-12">
          <button type="submit" class="btn btn-dark btn-sm ml-auto" @click="submitServerConfig()">
            <i class="fe fe-save"></i>
            Save
          </button>
        </div>
      </div>

    </eq-tabs>


    <div
      class="row justify-content-center"
      style="position: absolute; bottom: 5%; z-index: 9999999; width: 100%"
    >
      <div class="col-6">
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
  </eq-window>
</template>

<script>
import EqWindow        from "@/components/eq-ui/EQWindow.vue";
import EqTabs          from "@/components/eq-ui/EQTabs.vue";
import EqTab           from "@/components/eq-ui/EQTab.vue";
import {SpireApi}      from "@/app/api/spire-api";
import InfoErrorBanner from "@/components/InfoErrorBanner.vue";
import EqWindowComplex from "@/components/eq-ui/EQWindowComplex.vue";
import {AppEnv}        from "@/app/env/app-env";

export default {
  name: "ServerConfig",
  components: { EqWindowComplex, InfoErrorBanner, EqTab, EqTabs, EqWindow },
  data() {
    return {
      config: {
        server: {
          world: {
            loginservers: [], // Array to hold loginserver configurations
          },
        },
      },
      tabSelected: "World Server",

      maxLoginServers: 5,

      passwordFieldType: 'password',
      loaded: false,

      // notification / errors
      notification: "",
      error: "",
    }
  },
  watch: {
    '$route'() {
      this.loadQueryState()
    },
  },
  async created() {

    try {
      const r = await SpireApi.v1().get('admin/serverconfig')
      if (r.status === 200) {
        this.config = r.data
      }
    } catch (e) {
      // error notify
      if (e.response && e.response.data && e.response.data.error) {
        this.error = e.response.data.error
      }
    }

    if (!this.config.server.qsdatabase) {
      this.config.server.qsdatabase = {
        db: "",
        host: "",
        port: "",
        username: "",
        password: "",
      }
    }

    if (!this.config.server.content_database) {
      this.config.server.content_database = {
        db: "",
        host: "",
        port: "",
        username: "",
        password: "",
      }
    }

    if (!this.config.server.ucs) {
      this.config.server.ucs = {
        host: "",
        port: "",
      }
    }

    const loginServerConfigModel = {
      'account': '',
      'password': '',
      'legacy': '0',
      'host': '',
      'port': '5998'
    }

    for (let i = 1; i <= this.maxLoginServers; i++) {
      const key = `loginserver${i}`;
      if (!this.config.server.world[key]) {
        this.config.server.world[key] = { ...loginServerConfigModel };
      }
    }

  },
  mounted() {
    this.loadQueryState()
  },
  methods: {
    updateQueryState() {
      console.log("trigger")
      let q = {};
      if (this.tabSelected !== "") {
        q.s = this.tabSelected
      }

      this.$router.push(
        {
          path: this.$route.path,
          query: q
        }
      ).catch(() => {
      })
    },
    loadQueryState() {
      if (this.$route.query.s && this.$route.query.s.length > 0) {
        this.tabSelected = this.$route.query.s
      }
    },

    generateRandomKey: function () {
      let result           = ''
      let characters       = 'ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789'
      let charactersLength = characters.length
      for (let i = 0; i < 40; i++) {
        result += characters.charAt(Math.floor(Math.random() * charactersLength))
      }

      this.config.server.world.key = result
    },

    submitServerConfig: async function () {

      // remove empty content_database configs
      if (!this.config.server.content_database) {
        delete this.config.server.content_database
      }
      if (this.config.server.content_database && Object.keys(this.config.server.content_database).length === 0) {
        delete this.config.server.content_database
      }

      for (let i = 1; i <= this.maxLoginServers; i++) {
        const key = `loginserver${i}`;
        const loginserver = this.config.server.world[key];

        if (!loginserver || Object.keys(loginserver).every((k) => loginserver[k] === '')) {
          delete this.config.server.world[key];
        }
      }

      let loginServers = [];
      for (let i = 1; i <= this.maxLoginServers; i++) {
        const key = `loginserver${i}`;
        const loginserver = this.config.server.world[key];

        if (loginserver && Object.keys(loginserver).length > 0 && loginserver.host) {
          loginServers.push(loginserver); // Only keep valid loginservers
        }

        // Clean up the original object to remove gaps
        delete this.config.server.world[key];
      }

      // Reassign valid loginservers back without gaps
      loginServers.forEach((loginserver, index) => {
        const key = `loginserver${index + 1}`;
        this.config.server.world[key] = loginserver;
      });

      try {
        const r = await SpireApi.v1().post('admin/serverconfig', this.config)
        if (r.status === 200) {
          this.notification = "Server configuration updated!"
        }
      } catch (e) {
        // error notify
        if (e.response && e.response.data && e.response.data.error) {
          this.error = e.response.data.error
        }
      }

      // hack to sync the db name for now in local setups
      if (AppEnv.isAppLocal()) {
        await SpireApi.v1().post('app/sync')
      }
    },
    switchPasswordVisibility() {
      this.passwordFieldType = this.passwordFieldType === 'password' ? 'text' : 'password'
    }
  }
}
</script>

<style scoped>

</style>
