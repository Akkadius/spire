<template>
  <div class="row justify-content-center" v-if="connections">
    <div class="col-12 col-lg-10 col-xl-8 content-pop card">
      <!-- CONTENT -->
      <div class="container-fluid">

        <div class="row justify-content-between align-items-center mt-5">
          <div class="col-12">
            <h2 class="mb-2">
              Database Connection Properties
            </h2>
            <p class="text-muted mb-md-0">
              Manage your database connections, you can only have one active primary connection at a time
            </p>
          </div>
        </div>

        <b-tabs class="mt-4" content-class="mt-5" fill v-if="connections">
          <b-tab
            title="Connections"
            :active="(connections && connections.length > 0)"
          >
            <!-- Default Local -->
            <div class="card" v-if="defaultConnection && isAppProduction()">
              <div class="card-header">

                <!-- Title -->
                <h4 class="card-header-title">
                  <i class="fe fe-database"></i>

                  {{ defaultConnection.database_connection.name }}
                </h4>
                <small class="text-muted">Latest of ProjectEQ data, active when not logged in, read-only</small>

              </div>
              <div class="card-body">

                <!-- List group -->
                <div class="list-group list-group-flush my-n3">

                  <div class="list-group-item">
                    <div class="row align-items-center">
                      <div class="col-auto">

                        <!-- Avatar -->
                        <!--                        <div class="avatar">-->
                        <!--                          <img src="~@/assets/img/mysql-logo.png" class="avatar-img rounded">-->
                        <!--                        </div>-->

                        <!--                        <i class="fe fe-database" style="font-size: 50px; color: rgb(189 195 204)"></i>-->

                        <img
                          :style="'height: 120px; width: auto; border-radius: 40px; opacity: ' + (isDefaultActive ? '1' : '.5')"
                          src="@/assets/img/eq-logo.jpg"
                        >

                      </div>
                      <div class="col ml-n2">

                        <!-- Title -->
                        <h4 class="mb-1">
                          {{ defaultConnection.database_connection.name }}
                        </h4>

                        <!-- Time -->
                        <p class="card-text small text-muted">
                          <b>Host</b> {{ defaultConnection.database_connection.db_host }}
                          <b>Port</b> {{ defaultConnection.database_connection.db_port }}
                          <b>Database Name</b> {{ defaultConnection.database_connection.db_name }}
                          <b>User</b> {{ defaultConnection.database_connection.db_username }}
                        </p>

                      </div>
                      <div class="col-auto">

                        <!-- Button -->
                        <a
                          class="btn btn-sm btn-outline-primary d-none d-md-inline-block"
                          @click="setDefaultActive()"
                          v-if="!isDefaultActive"
                        >
                          <i class="fe fe-activity"></i> Set Active
                        </a>

                        <a class="btn btn-sm btn-white" v-if="isDefaultActive">
                          <span class="text-success">●</span> Active Connection
                        </a>

                        <!-- The application shouldn't even boot if this isn't available -->
                        <a
                          class="btn btn-sm btn-white btn-danger ml-2"
                          v-b-tooltip.hover
                          v-b-tooltip.v-outline-success
                        >
                          <span class="text-success">●</span> Online
                        </a>


                      </div>

                    </div> <!-- / .row -->
                  </div>
                </div>

              </div> <!-- / .card-body -->
            </div>

            <!-- User Defined -->
            <user-connections
              @reload-connections="reloadConnections()"
              :connections="connections"
              :connection-statuses="connectionStatuses"
            />

          </b-tab>

          <b-tab
            title="Create New"
            v-if="canCreateNewConnection()"
            :active="(connections.length === 0)"
          >

            <div class="form-row">
              <div class="form-group col-md-12">
                <label class="form-label">Database Connection Name</label>
                <input
                  type="text" placeholder="My EverQuest Emulator Server" class="form-control"
                  v-model="database.connection_name"
                />
              </div>

              <div class="form-group col-md-4">
                <label class="form-label">Database Name</label>
                <input type="text" class="form-control" placeholder="peq" v-model="database.db_name"/>
              </div>

              <div class="form-group col-md-4">
                <label class="form-label">Database Host</label>
                <input type="text" class="form-control" placeholder="eqemu-server.com" v-model="database.db_host"/>
              </div>

              <div class="form-group col-md-4">
                <label class="form-label">Database Port</label>
                <input type="text" class="form-control" placeholder="3306" v-model="database.db_port"/>
              </div>
            </div>
            <div class="form-row">
              <div class="form-group col-md-6">
                <label class="form-label">Database Username</label>
                <input type="text" class="form-control" placeholder="username" v-model="database.db_username"/>
              </div>

              <div class="form-group col-md-6">
                <label class="form-label">Database Password</label>

                <div class="input-group">
                  <input
                    :type="passwordFieldType" class="form-control" placeholder="password"
                    v-model="database.db_password"
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
            <div class="form-row ml-0 mt-3">
              <div class="col">
                <button
                  class="btn btn-primary" type="button"
                  @click="addContentDb ? addContentDb = false : addContentDb = true"
                >
                  <i class="fe fe-database"></i>
                  Add Content Database
                </button>
              </div>
              <div class="col-auto">
                <button
                  class="btn btn-white ml-3" type="button"
                  @click="createConnection()"
                >
                  <i class="fe fe-plus"></i>
                  Create Connection
                </button>
              </div>
            </div>

            <b-alert variant="primary" show class="p-3 mt-4 mb-5" v-if="addContentDb">
              <i class="fe fe-info"></i>
              Note: Content database connections are optional and not required. See <a
              href="https://eqemu.gitbook.io/server/categories/database/multi-tenancy" style="color:white"
            >here</a> for
              more info
            </b-alert>

            <div class="form-row" v-if="addContentDb">
              <div class="form-group col-md-4">
                <label class="form-label">Content Database Name</label>
                <input type="text" class="form-control" v-model="database.content_db_name"/>
              </div>

              <div class="form-group col-md-4">
                <label class="form-label">Content Database Host</label>
                <input type="text" class="form-control" v-model="database.content_db_host"/>
              </div>

              <div class="form-group col-md-4">
                <label class="form-label">Content Database Port</label>
                <input type="text" class="form-control" v-model="database.content_db_port"/>
              </div>
            </div>
            <div class="form-row" v-if="addContentDb">
              <div class="form-group col-md-6">
                <label class="form-label">Content Database Username</label>
                <input type="text" class="form-control" v-model="database.content_db_username"/>
              </div>

              <div class="form-group col-md-6">
                <label class="form-label">Content Database Password</label>

                <div class="input-group">
                  <input :type="passwordFieldType" class="form-control" v-model="database.content_db_password"/>
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

        <b-alert
          :show="(errorMessage !== '')" v-if="errorMessage"
          class="mt-5 mb-5"
          variant="danger"
        >
          <i class="fe fe-alert-triangle"></i>
          {{ errorMessage }}
        </b-alert>

        <b-alert
          :show="(successMessage !== '')" v-if="successMessage"
          class="mt-5 mb-5"
          variant="success"
        >
          <i class="fe fe-database"></i>
          {{ successMessage }}
        </b-alert>

        <debug-display-component :data="database"/>

        <!--        <pre v-if="debug" class="mt-4 highlight html bg-dark hljs xml">{{ database }}</pre>-->
      </div>
    </div>

  </div>

</template>

<script type="ts">
import EqWindow from "@/components/eq-ui/EQWindow.vue";
import {App} from "@/constants/app";
import axios from "axios"
import {SpireApi} from "@/app/api/spire-api";
import DebugDisplayComponent from "@/components/DebugDisplayComponent.vue";
import {EventBus} from "@/app/event-bus/event-bus";
import AddDeveloperModal from "@/views/connections/AddDeveloperModal.vue";
import ManageDeveloperModal from "@/views/connections/ManageDeveloperModal.vue";
import ManageDiscordConnectionModal from "@/views/connections/ManageDiscordConnectionModal.vue";
import {AppEnv} from "@/app/env/app-env";
import UserConnections from "@/views/connections/UserConnections.vue";

export default {
  components: {
    UserConnections,
    ManageDiscordConnectionModal,
    ManageDeveloperModal,
    AddDeveloperModal,
    DebugDisplayComponent,
    EqWindow,
    "page-header": () => import("@/components/layout/PageHeader.vue")
  },
  data() {
    return {
      passwordFieldType: 'password',
      database: {},
      debug: App.DEBUG,
      addContentDb: false,
      errorMessage: null,
      successMessage: null,
      defaultConnection: null,
      isDefaultActive: false,
      connections: [],

      connectionStatuses: {},
    }
  },
  async mounted() {
    await this.listConnections()
  },
  methods: {
    isAppProduction() {
      return AppEnv.isAppProduction()
    },
    canCreateNewConnection() {
      if (AppEnv.isAppProduction()) {
        return true
      }

      return false;
    },
    reloadConnections() {
      this.listConnections()
    },
    switchPasswordVisibility() {
      this.passwordFieldType = this.passwordFieldType === 'password' ? 'text' : 'password'
    },
    setDefaultActive() {
      SpireApi.v1().post(`/connection-default/set-active`).then((response) => {
        if (response.data && response.data.data) {
          EventBus.$emit('DB_CONNECTION_CHANGE', true);
          this.listConnections()
        }
      })
    },
    async listConnections() {
      // default local
      SpireApi.v1().get('/connection-default').then((rd) => {
        if (rd.data && rd.data.data) {
          this.defaultConnection = rd.data.data
        }
      })

      // user defined
      const r = await SpireApi.v1().get('/connections')
      if (r.data && r.data.data) {
        this.connections = r.data.data.filter(c => c.database_connection.id > 0)

        let isDefaultActive = true
        for (const c of r.data.data) {
          const connectionId = c.id

          if (c.active) {
            isDefaultActive = false
          }

          const s = await SpireApi.v1().get(`/connection-check/${connectionId}`)
          if (s.status === 200) {
            this.connectionStatuses[connectionId] = s.data.data.message
          }
        }

        this.isDefaultActive = isDefaultActive
      }


    },
    createConnection() {
      this.errorMessage   = null
      this.successMessage = null

      SpireApi.v1().post('/connection', this.database).then((response) => {
        this.successMessage = response.data.data
        this.listConnections()
      }, (error) => {
        this.errorMessage = "Unknown error trying to contact the database"
        if (error.response && error.response.data) {
          this.errorMessage = error.response.data.error
        }
      }).catch((error) => {
        if (!axios.isCancel(error)) {
          console.log(error)
          this.errorMessage = "Unknown error trying to contact the database"
        }
      });

    }
  }
}

</script>
