<template>
  <eq-window
    title="Server Processes"
    v-if="processCounts"
    class="p-0 mb-4"
  >
    <div class="p-3 text-center" v-if="processCounts && processCounts.length === 0">
      Server offline
    </div>

    <table
      class="eq-table bordered eq-highlight-rows mb-0"
      v-if="processCounts && processCounts.length > 0"
    >
      <thead class="eq-table-floating-header">
      <tr>
        <th class="text-right">
          Process
        </th>
        <th class="text-left" style="width: 130px">
          Status
        </th>
      </tr>
      </thead>

      <tbody>

      <!-- List Counts -->
      <tr
        v-b-tooltip.hover.v-dark.left
        :title="p.optional ? 'This is an optional service and not required to be running' : ''"
        v-for="p in processCounts"
        :key="p.name"
      >
        <td class="text-right font-weight-bold">
          {{ p.name }}
        </td>
        <td
          class="text-left"
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
  </eq-window>
</template>

<script>

import {EventBus} from "@/app/event-bus/event-bus";
import {SpireApi} from "@/app/api/spire-api";
import EqWindow   from "@/components/eq-ui/EQWindow.vue";

export default {
  name: 'DashboardProcessCounts',
  components: { EqWindow },
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
    SpireApi.v1().get("eqemuserver/server-stats").then((r) => {
      if (r.status === 200) {
        EventBus.$emit("server-stats", r.data)
      }
    })

    EventBus.$on('server-stats', async (e) => {
      this.processCounts = []

      let p = []
      p.push({ name: "Spire Launcher", count: e.launcher_online ? 1 : 0 })
      p.push({ name: "World (world)", count: e.world_online ? 1 : 0 })
      p.push({ name: "Zones (zone)", count: e && e.zone_list && e.zone_list.data ? e.zone_list.data.length : 0 })
      p.push({ name: "Universal Chat Service (ucs)", count: e.ucs_online ? 1 : 0, optional: true })
      p.push({ name: "Loginserver (loginserver)", count: e.login_online ? 1 : 0, optional: true })
      p.push({ name: "Queryserv (queryserv)", count: e.query_serv_online ? 1 : 0, optional: true })

      this.processCounts = p

      this.loaded = true
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
