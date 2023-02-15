<template>
  <div v-if="processCounts">
    <div class="card">
      <div class="card-header">
        <h4 class="card-header-title">Server Processes</h4>
      </div>
      <table class="table card-table">
        <tbody>

        <!-- List Counts -->
        <tr v-for="p in processCounts" :key="p.name">
          <td>
            {{ p.name }}
            <span class="text-muted" v-if="p.optional">(Optional Service)</span>
          </td>
          <td class="text-right">
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
      let p = []
      p.push({ name: "Occulus Launcher", count: e.launcher_online ? 1 : 0 })
      p.push({ name: "World (world)", count: e.world_online ? 1 : 0 })
      p.push({ name: "Zones (zone)", count: e.zone_list.data.length })
      p.push({ name: "Universal Chat Service (ucs)", count: e.ucs_online ? 1 : 0, optional: true })
      p.push({ name: "Loginserver (loginserver)", count: e.login_online ? 1 : 0, optional: true })
      p.push({ name: "Queryserv (queryserv)", count: e.query_serv_online ? 1 : 0, optional: true })

      this.processCounts = p
    })
  },
}
</script>
