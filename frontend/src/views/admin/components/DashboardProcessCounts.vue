<template>
  <div v-if="processCounts">
    <div class="card">
      <div class="card-header">
        <h4 class="card-header-title">Server Processes</h4>
      </div>

      <div class="p-3 text-center" v-if="processCounts && processCounts.length === 0">
        Server offline
      </div>

      <table class="table card-table process-counts" v-if="processCounts && processCounts.length > 0">
        <tbody>

        <!-- List Counts -->
        <tr
          v-b-tooltip.hover.v-dark.left
          :title="p.optional ? 'This is an optional service and not required to be running' : ''"
          v-for="p in processCounts"
          :key="p.name"
        >
          <td>
            {{ p.name }}
          </td>
          <td

            class="text-right"
          >
            <span
              class="badge badge-danger"
              style="font-size: 12px"
              v-if="p.count === 0"
            >Offline</span>
            <span
              class="badge badge-success"
              style="font-size: 12px"
              v-if="p.count > 0"
            >Online ({{ p.count }})</span>
          </td>
        </tr>
        </tbody>
      </table>
    </div>
  </div>
</template>

<script>

import {EventBus} from "@/app/event-bus/event-bus";

export default {
  name: 'DashboardProcessCounts',
  data() {
    return {
      processCounts: [],
      loaded: false
    }
  },

  beforeDestroy() {
    EventBus.$off('server-stats')
  },

  mounted() {
    EventBus.$on('server-stats', async (e) => {
      this.processCounts = []

      let p = []
      p.push({ name: "Occulus Launcher", count: e.launcher_online ? 1 : 0 })
      p.push({ name: "World (world)", count: e.world_online ? 1 : 0 })
      p.push({ name: "Zones (zone)", count: e && e.zone_list && e.zone_list.data ? e.zone_list.data.length : 0 })
      p.push({ name: "Universal Chat Service (ucs)", count: e.ucs_online ? 1 : 0, optional: true })
      p.push({ name: "Loginserver (loginserver)", count: e.login_online ? 1 : 0, optional: true })
      p.push({ name: "Queryserv (queryserv)", count: e.query_serv_online ? 1 : 0, optional: true })

      this.processCounts = p
    })
  },
}
</script>

<style>
.process-counts td {
  padding: 0.7rem;
  vertical-align: top;
  border-top: 1px solid rgba(0, 40, 100, 0.12);
}
</style>
