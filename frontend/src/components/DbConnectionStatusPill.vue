<template>
  <div
    class="card mt-3"
    style="margin-bottom: 5px; background-color: rgba(0,0,0, .5); color: rgba(255,255,255,.7); border: 1px solid #000000;"
    v-if="connection && connection.database_connection && connection.database_connection.name"
    @click="navigateConnections"
  >
    <div
      class="card-body connection-status-box"
      style="padding: 5px; padding-left: 15px; text-align: left;"
      v-b-tooltip.v-dark.hover
      :title="getConnectionDescription()"
    >
      <div
        class="avatar avatar-sm mr-3"
        style="height: 10px; width: 10px"
      >
        <img
          :style="'background-color: ' + getConnectionStatusColor() + '; margin-bottom: 5px; transition: background-color 300ms;'"
          src="data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAYAAAAfFcSJAAAADUlEQVR42mNk+M9QDwADhgGAWjR9awAAAABJRU5ErkJggg=="
          class="avatar-img rounded-circle"
        >
      </div>
      <i class="fe fe-database"></i> {{ connection.database_connection.name }}
    </div>
  </div>
</template>

<script>
import {SpireApi} from "../app/api/spire-api";
import {EventBus} from "@/app/event-bus/event-bus";
import {ROUTE}          from "@/routes";
import util             from "util";

export default {
  name: "DbConnectionStatusPill",
  data() {
    return {
      connection: {},
      connectionStatus: "",
    }
  },
  mounted() {
    this.fetchConnection()
  },

  created() {
    EventBus.$on("DB_CONNECTION_CHANGE", this.fetchConnection);
  },
  destroyed() {
    EventBus.$off("DB_CONNECTION_CHANGE", this.fetchConnection);
  },

  methods: {
    getConnectionDescription() {
      return util.format(
        "Host: %s Status: %s",
        this.connection.database_connection.db_host,
        this.connectionStatus
      )
    },

    navigateConnections() {
      this.$router.push(
        {
          path: ROUTE.DATABASE_CONNECTIONS
        }
      ).catch(() => {
      })
    },

    getConnectionStatusColor() {
      if (this.connectionStatus.includes("connecting")) {
        return 'white';
      }
      if (!this.connectionStatus.includes("online")) {
        return 'red';
      }

      // connectionStatus
      return '#00d97e'
    },

    fetchConnection() {
      // connection status
      SpireApi.v1().get('/connections').then((r) => {

        this.connection       = {}
        this.connectionStatus = "connecting"

        if (r.data && r.data.data) {
          r.data.data.forEach(connection => {
            const connectionId = connection.id

            if (connection.active) {
              this.connection = connection

              SpireApi.v1().get(`/connection-check/${connectionId}`).then((cr) => {
                this.connectionStatus = cr.data.data.message
              })
            }
          })

          if (Object.keys(this.connection).length === 0) {
            // console.log("There is no non-default connection active")
            this.connection.database_connection      = {}
            this.connection.database_connection.name = "Local (Default)"
            this.connectionStatus                    = "online"
          }

        }
      })
    }
  }
}
</script>

<style scoped>

</style>
