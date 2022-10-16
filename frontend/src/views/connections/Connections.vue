<template>
  <div class="row justify-content-center" v-if="connections">
    <div class="col-12 col-lg-10 col-xl-8 content-pop">
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

        <!-- Modal -->
        <add-developer-modal
          @reload-connections="reloadConnections()"
          :connection="addDeveloperToConnection"
        />

        <manage-developer-modal
          @reload-connections="reloadConnections()"
          :user="managedUser"
          :connection="managedUserConnection"
        />

        <manage-discord-connection-modal
          :connection="discordConnection"
        />

        <b-tabs class="mt-4" content-class="mt-5" fill>
          <b-tab title="Connections" :active="connections && connections.length > 0">

            <!-- Default Local -->
            <div class="card" v-if="defaultConnection">
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

                        <i class="fe fe-database" style="font-size: 50px; color: rgb(189 195 204)"></i>

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
            <div class="card">
              <div class="card-header">

                <!-- Title -->
                <h4 class="card-header-title">
                  <i class="fe fe-database"></i>
                  User Remote Database Connections ({{ connections.length }})
                </h4>

              </div>
              <div class="card-body">

                <!-- List group -->
                <div class="list-group list-group-flush my-n3">

                  <div v-if="connections.length === 0 && isLoggedIn()">
                    No database connections, you can create a new one
                  </div>

                  <div v-if="!isLoggedIn()">
                    You must log in to create connections
                  </div>

                  <div
                    class="list-group-item"
                    :key="connection.id"
                    v-for="connection in connections"
                  >
                    <div class="row align-items-center">
                      <div class="col-auto">

                        <!-- Avatar -->
                        <!--                        <a href="#!" class="avatar">-->
                        <!--                                                    <img src="~@/assets/img/eqemu-avatar.png" alt="..." class="avatar-img rounded">-->
                        <!--                          <img src="~@/assets/img/mysql-logo.png" alt="..." class="avatar-img rounded">-->
                        <!--                        </a>-->

                        <i class="fe fe-database" style="font-size: 50px; color: rgb(189 195 204)"></i>

                      </div>
                      <div class="col ml-n2">

                        <!-- Title -->
                        <h4 class="mb-1">
                          {{ connection.database_connection.name }}
                        </h4>

                        <p class="card-text small text-muted mb-2">
                          <b>Host</b> {{ connection.database_connection.db_host }}
                          <b>Port</b> {{ connection.database_connection.db_port }}
                          <b>Database Name</b> {{ connection.database_connection.db_name }}
                          <b>User</b> {{ connection.database_connection.db_username }}
                        </p>

                        <h4 class="mb-1 text-muted" v-if="connection.database_connection.content_db_username !== ''">
                          Content Connection
                        </h4>

                        <p
                          class="card-text small text-muted mb-2"
                          v-if="connection.database_connection.content_db_username !== ''"
                        >
                          <b>Host</b> {{ connection.database_connection.content_db_host }}
                          <b>Port</b> {{ connection.database_connection.content_db_port }}
                          <b>Database Name</b> {{ connection.database_connection.content_db_name }}
                          <b>User</b> {{ connection.database_connection.content_db_username }}
                        </p>

                        <!-- Users -->

                        <h4
                          v-if="connection.database_connection.user_server_database_connections.length > 0"
                          class="mb-1 mt-1"
                        >
                          Developers
                          <a
                            class="btn btn-sm btn-white btn-danger ml-2 pt-0 pb-0"
                            v-if="isCurrentUserOwnerOfConnection(connection)"
                            @click="addDeveloperToServerModal(connection)"
                            v-b-modal.add-developer-server-connection-modal
                          ><i class="fa fa-plus"/> Add</a>
                        </h4>

                        <b-avatar-group rounded="lg" overlap="0.05" class="d-inline-block mt-1">
                          <div
                            v-for="user in connection.database_connection.user_server_database_connections"
                            :key="user.user.id"
                            @click="manageUser(user.user, connection)"
                          >
                            <b-avatar
                              :src="user.user.avatar"
                              class="hover-highlight mr-1"
                              variant="info"
                              v-b-tooltip.hover.v-dark.top
                              :title="user.user.user_name + (isUserOwnerOfConnection(user.user, connection) ? ' (Owner)' : '')"
                            />
                          </div>
                        </b-avatar-group>

                      </div>
                      <div class="col-auto">

                        <!-- Loader -->
                        <i
                          class="fa fa-spinner fa-pulse fa-3x fa-fw"
                          v-if="connection.active === 0 && (!connectionStatuses[connection.id])"
                        />

                        <!-- Button -->
                        <a
                          class="btn btn-sm btn-outline-primary d-none d-md-inline-block"
                          @click="setActiveConnection(connection.id)"
                          v-if="connection.active === 0 && (connectionStatuses[connection.id] && !connectionStatuses[connection.id].includes('error'))"
                        >
                          <i class="fe fe-activity"></i> Set Active
                        </a>

                        <a class="btn btn-sm btn-white" v-if="connection.active === 1">
                          <span class="text-success">●</span> Active Connection
                        </a>

                        <a
                          class="btn btn-sm btn-danger ml-2"
                          v-if="connectionStatuses[connection.id] && connectionStatuses[connection.id].includes('error')"
                          v-b-tooltip.hover
                          v-b-tooltip.v-danger
                          :title="connectionStatuses[connection.id]"
                        >
                          <span class="text-white">●</span> Offline
                        </a>

                        <a
                          class="btn btn-sm btn-white btn-danger ml-2"
                          v-if="connectionStatuses[connection.id] && !connectionStatuses[connection.id].includes('error')"
                          v-b-tooltip.hover
                          v-b-tooltip.v-outline-success
                          :title="connectionStatuses[connection.id]"
                        >
                          <span class="text-success">●</span> Online
                        </a>

                        <!-- Audit log -->
                        <a
                          class="btn btn-sm btn-white ml-1"
                          @click="viewAuditLog(connection.id)"
                          title="View audit log"
                          v-b-tooltip.hover.v-dark.top
                          v-if="isCurrentUserOwnerOfConnection(connection)"
                        >
                          <i class="ra ra-bleeding-eye"></i>
                        </a>

                        <!-- Discord -->
                        <a
                          class="btn btn-sm btn-white ml-1"
                          @click="setDiscordWebhookLogs(connection)"
                          title="Set Discord webhook logs"
                          v-b-tooltip.hover.v-dark.top
                          v-if="isCurrentUserOwnerOfConnection(connection)"
                        >
                          <img
                            src="~@/assets/img/discord-logo-small.8530e062.png"
                            style="width: 25px; height: 25px"
                          >
                        </a>

                      </div>

                      <div class="col-auto p-0" v-if="isCurrentUserOwnerOfConnection(connection)">

                        <!-- Dropdown -->
                        <div class="dropdown">

                          <!-- Toggle -->
                          <a
                            href="#" class="dropdown-ellipses dropdown-toggle" role="button" data-toggle="dropdown"
                            aria-haspopup="true" aria-expanded="false"
                          >
                            <i class="fe fe-more-vertical"></i>
                          </a>

                          <!-- Menu -->
                          <div class="dropdown-menu dropdown-menu-right" style="">
                            <a @click="deleteConnection(connection.id)" class="dropdown-item">
                              Delete Connection
                            </a>
                          </div>

                        </div>

                      </div>

                    </div> <!-- / .row -->
                  </div>
                </div>

              </div> <!-- / .card-body -->
            </div>

          </b-tab>

          <b-tab
            title="Create New"
            v-if="isLoggedIn()"
            :active="connections.length === 0 && isLoggedIn()"
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
import {SpireApi} from "../../app/api/spire-api";
import DebugDisplayComponent from "@/components/DebugDisplayComponent.vue";
import {EventBus} from "@/app/event-bus/event-bus";
import UserContext from "@/app/user/UserContext";
import AddDeveloperModal from "@/views/connections/AddDeveloperModal.vue";
import ManageDeveloperModal from "@/views/connections/ManageDeveloperModal.vue";
import util from "util";
import {ROUTE} from "@/routes";
import ManageDiscordConnectionModal from "@/views/connections/ManageDiscordConnectionModal.vue";

export default {
  components: {
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
      connections: null,
      connectionStatuses: {},

      user: {},

      // modals
      addDeveloperToConnection: {}, // AddDeveloperToModal

      discordConnection: {}, // ManageDiscordConnectionModal

      managedUser: {}, // ManageDeveloperModal
      managedUserConnection: {} // ManageDeveloperModal
    }
  },
  async mounted() {
    this.listConnections()

    this.user = await (UserContext.getUser())
  },
  methods: {

    setDiscordWebhookLogs(connection) {
      this.discordConnection = connection

      // I don't remember why I needed to queue this
      setTimeout(() => {
        this.$bvModal.show('manage-discord-connection-modal')
      }, 10)
    },

    viewAuditLog(connectionId) {
      this.$router.push(
        {
          path: util.format(ROUTE.DATABASE_CONNECTION_AUDIT_LOG, connectionId)
        }
      ).catch(() => {
      })
    },

    canViewPermissions(user, connection) {
      return this.isCurrentUserOwnerOfConnection(connection) || this.user.id === user.id
    },

    manageUser(user, connection) {
      // if (!this.canViewPermissions(user, connection)) {
      //   return;
      // }

      this.managedUser = user
      this.managedUserConnection = connection

      // I don't remember why I needed to queue this
      setTimeout(() => {
        this.$bvModal.show('manage-developer-modal')
      }, 10)
    },

    reloadConnections() {
      this.listConnections()
    },

    addDeveloperToServerModal(connection) {
      this.addDeveloperToConnection = connection
    },

    isCurrentUserOwnerOfConnection(connection) {
      return this.user.id === connection.database_connection.created_by
    },

    isUserOwnerOfConnection(user, connection) {
      return user.id === connection.database_connection.created_by
    },

    isLoggedIn() {
      return this.user && Object.keys(this.user).length
    },

    switchPasswordVisibility() {
      this.passwordFieldType = this.passwordFieldType === 'password' ? 'text' : 'password'
    },
    deleteConnection(connectionId) {
      if (confirm("Are you sure you want to delete this connection and everything associated to it? This is permanent")) {
        SpireApi.v1().delete(`/connection/${connectionId}`).then((response) => {
          if (response.data && response.data.data) {
            this.listConnections()
          }
        })
      }
    },
    setActiveConnection(connectionId) {
      SpireApi.v1().post(`/connection/${connectionId}/set-active`).then((response) => {
        if (response.data && response.data.data) {
          EventBus.$emit('DB_CONNECTION_CHANGE', true);
          this.listConnections()
        }
      })
    },
    setDefaultActive() {
      SpireApi.v1().post(`/connection-default/set-active`).then((response) => {
        if (response.data && response.data.data) {
          EventBus.$emit('DB_CONNECTION_CHANGE', true);
          this.listConnections()
        }
      })
    },
    listConnections() {

      // user defined
      SpireApi.v1().get('/connections').then((response) => {
        if (response.data && response.data.data) {
          this.connections = response.data.data

          let isDefaultActive = true
          response.data.data.forEach(connection => {
            const connectionId = connection.id

            if (connection.active) {
              isDefaultActive = false
            }

            SpireApi.v1().get(`/connection-check/${connectionId}`).then((response) => {
              this.connectionStatuses[connectionId] = response.data.data.message
              this.$forceUpdate()
            })
          })

          this.isDefaultActive = isDefaultActive
        }
      })

      // default local
      SpireApi.v1().get('/connection-default').then((response) => {
        if (response.data && response.data.data) {
          this.defaultConnection = response.data.data
        }
      })

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
