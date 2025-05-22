<template>
  <div>

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

    <info-error-banner
      :slim="true"
      :notification="notification"
      :error="error"
      @dismiss-error="error = ''"
      @dismiss-notification="notification = ''"
      class="mb-3"
    />

    <div class="card">
      <div class="card-header">

        <!-- Title -->
        <h4 class="card-header-title">
          <i class="fe fe-database"></i>
          User Database Connections ({{ connections.length }})
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
            :key="c.id"
            v-for="c in connections"
          >
            <div class="row align-items-center">
              <div class="col-auto">

                <img
                  :style="'height: 120px; width: auto; border-radius: 50px; opacity: ' + (c.active === 1 ? '1' : '.5')"
                  src="@/assets/img/eq-logo.jpg"
                >

              </div>
              <div class="col ml-n2">

                <!-- Title -->
                <h4 class="mb-1">
                  {{ c.database_connection.name }}
                </h4>

                <p class="card-text small text-muted mb-2">
                  <b>Host</b> {{ c.database_connection.db_host }}
                  <b>Port</b> {{ c.database_connection.db_port }}
                  <b>Database Name</b> {{ c.database_connection.db_name }}
                  <b>User</b> {{ c.database_connection.db_username }}
                </p>

                <h4 class="mb-1 text-muted" v-if="c.database_connection.content_db_username !== ''">
                  Content Connection
                </h4>

                <p
                  class="card-text small text-muted mb-2"
                  v-if="c.database_connection.content_db_username !== ''"
                >
                  <b>Host</b> {{ c.database_connection.content_db_host }}
                  <b>Port</b> {{ c.database_connection.content_db_port }}
                  <b>Database Name</b> {{ c.database_connection.content_db_name }}
                  <b>User</b> {{ c.database_connection.content_db_username }}
                </p>

                <!-- Users -->

                <h4
                  v-if="c.database_connection.user_server_database_connections.length > 0"
                  class="mb-1 mt-1"
                >
                  Developers
                  <a
                    class="btn btn-sm btn-white btn-danger ml-2 pt-0 pb-0"
                    v-if="isCurrentUserOwnerOfConnection(c)"
                    @click="addDeveloperToServerModal(c)"
                    v-b-modal.add-developer-server-connection-modal
                  ><i class="fa fa-plus"/> Add</a>
                </h4>

                <b-avatar-group rounded="lg" overlap="0.05" class="d-inline-block mt-1">
                  <div
                    v-for="user in c.database_connection.user_server_database_connections.
                      filter(u => u.user.deleted_at === null && u.user.id > 0)"
                    :key="user.user.id"
                    @click="manageUser(user.user, c)"
                  >
                    <b-avatar
                      :src="user.user.avatar"
                      class="hover-highlight mr-1"
                      variant="info"
                      v-b-tooltip.hover.v-dark.top
                      :title="user.user.user_name + (isUserOwnerOfConnection(user.user, c) ? ' (Owner)' : '')"
                    />
                  </div>
                </b-avatar-group>

              </div>
              <div class="col-auto">

                <!-- Loader -->
                <i
                  class="fa fa-spinner fa-pulse fa-3x fa-fw"
                  v-if="c.active === 0 && (!connectionStatuses[c.id])"
                />

                <!-- Button -->
                <a
                  class="btn btn-sm btn-outline-primary d-none d-md-inline-block"
                  @click="setActiveConnection(c.id)"
                  v-if="c.active === 0 && (connectionStatuses[c.id] && !connectionStatuses[c.id].includes('error'))"
                >
                  <i class="fe fe-activity"></i> Set Active
                </a>

                <a class="btn btn-sm btn-white" v-if="c.active === 1">
                  <span class="text-success">●</span> Active Connection
                </a>

                <a
                  class="btn btn-sm btn-danger ml-2"
                  v-if="connectionStatuses[c.id] && connectionStatuses[c.id].includes('error')"
                  v-b-tooltip.hover
                  v-b-tooltip.v-danger
                  :title="connectionStatuses[c.id]"
                >
                  <span class="text-white">●</span> Offline
                </a>

                <a
                  class="btn btn-sm btn-white btn-danger ml-2"
                  v-if="connectionStatuses[c.id] && !connectionStatuses[c.id].includes('error')"
                  v-b-tooltip.hover
                  v-b-tooltip.v-outline-success
                  :title="connectionStatuses[c.id]"
                >
                  <span class="text-success">●</span> Online
                </a>

                <!-- Audit log -->
                <a
                  class="btn btn-sm btn-white ml-1"
                  @click="viewAuditLog(c.server_database_connection_id)"
                  title="View audit log"
                  v-b-tooltip.hover.v-dark.top
                  v-if="isCurrentUserOwnerOfConnection(c)"
                >
                  <i class="ra ra-bleeding-eye"></i>
                </a>

                <!-- Discord -->
                <a
                  class="btn btn-sm btn-white ml-1"
                  @click="setDiscordWebhookLogs(c)"
                  title="Set Discord webhook logs"
                  v-b-tooltip.hover.v-dark.top
                  v-if="isCurrentUserOwnerOfConnection(c)"
                >
                  <img
                    src="~@/assets/img/discord-logo-small.8530e062.png"
                    style="width: 25px; height: 25px"
                  >
                </a>

              </div>

              <div class="col-auto p-0" v-if="isCurrentUserOwnerOfConnection(c)">

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
                    <a @click="deleteConnection(c.server_database_connection_id)" class="dropdown-item">
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
  </div>
</template>

<script>
import util                         from "util";
import {ROUTE}                      from "@/routes";
import {SpireApi}                   from "@/app/api/spire-api";
import {EventBus}                   from "@/app/event-bus/event-bus";
import AddDeveloperModal            from "@/views/connections/AddDeveloperModal.vue";
import ManageDeveloperModal         from "@/views/connections/ManageDeveloperModal.vue";
import ManageDiscordConnectionModal from "@/views/connections/ManageDiscordConnectionModal.vue";
import UserContext                  from "@/app/user/UserContext";
import InfoErrorBanner              from "@/components/InfoErrorBanner.vue";

export default {
  name: "UserConnections",
  components: { InfoErrorBanner, ManageDiscordConnectionModal, ManageDeveloperModal, AddDeveloperModal },
  props: {
    connections: {
      type: Array,
      required: true
    },
    connectionStatuses: {
      type: [Array, Object],
      required: true
    },
  },
  watch: {
    connectionStatuses: {
      deep: true,
      handler() {
        console.log("statuses")
      }
    },
    connections: {
      deep: true,
      handler() {
        console.log("connections")
      }
    },
  },
  data() {
    return {
      user: {},

      // notification / errors
      notification: "",
      error: "",

      // modals
      addDeveloperToConnection: {}, // AddDeveloperToModal

      discordConnection: {}, // ManageDiscordConnectionModal

      managedUser: {}, // ManageDeveloperModal
      managedUserConnection: {} // ManageDeveloperModal
    }
  },
  async mounted() {
    this.user = await (UserContext.getUser())

    setTimeout(() => {
      this.$forceUpdate()
    }, 1000)
  },
  methods: {

    isCurrentUserOwnerOfConnection(connection) {
      return this.user.id === connection.database_connection.created_by
    },

    reloadConnections() {
      this.$emit("reload-connections", true);
    },

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

    manageUser(user, connection) {
      // if (!this.canViewPermissions(user, connection)) {
      //   return;
      // }

      this.managedUser           = user
      this.managedUserConnection = connection

      // I don't remember why I needed to queue this
      setTimeout(() => {
        this.$bvModal.show('manage-developer-modal')
      }, 10)
    },

    isUserOwnerOfConnection(user, connection) {
      return user.id === connection.database_connection.created_by
    },

    isLoggedIn() {
      return this.user && Object.keys(this.user).length > 0
    },

    addDeveloperToServerModal(connection) {
      this.addDeveloperToConnection = connection
    },

    async deleteConnection(connectionId) {
      if (confirm("Are you sure you want to delete this connection and everything associated to it? This is permanent")) {

        try {
          const r = await SpireApi.v1().delete(`/connection/${connectionId}`)
          if (r.data && r.data.data) {
            this.reloadConnections()
            this.notification = "Connection deleted successfully"
          }
        } catch (err) {
          if (err.response.data.error) {
            if (err.response && err.response.data && err.response.data.error) {
              this.error = err.response.data.error
            }
          }
        }
      }
    },
    setActiveConnection(connectionId) {
      SpireApi.v1().post(`/connection/${connectionId}/set-active`).then((response) => {
        if (response.data && response.data.data) {
          EventBus.$emit('DB_CONNECTION_CHANGE', true);
          this.reloadConnections()
        }
      })
    },
  }
}
</script>

<style scoped>

</style>
