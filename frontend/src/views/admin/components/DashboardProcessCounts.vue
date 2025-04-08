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
          <div class="mt-1">Process</div>
        </th>
        <th class="text-left">
          Status
          <button
            class="ml-3 btn btn-sm btn-dark"
            @click="showStats = !showStats"
          >
            <i class="fe fe-activity"></i>
            {{ showStats ? 'Hide' : 'Show' }} Stats
          </button>
        </th>
      </tr>
      </thead>

      <tbody>

      <!-- List Counts -->
      <tr
        v-b-tooltip.hover.v-dark.left
        :title="p.optional ? 'This is an optional service and not required to be running' : ''"
        v-for="p in processCounts"
        :key="p.description"
      >
        <td class="text-right font-weight-bold">
          {{ p.description }}
        </td>
        <td
          class="text-center"
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

          <div v-if="showStats && (p.cpu || p.memory)" class="mt-2">
            <div
              class="row"
              v-for="(metric, index) in getMetrics(p)"
              :key="metric.percent"
            >
              <div class="col-3 p-0 pr-3 m-0 text-right">
                <div class="small font-weight-bold" style="font-size: 10px;">
                  {{ metric.label }}
                </div>
              </div>
              <div class="col-6 p-0 m-0">
                <eq-progress-bar
                  style="opacity: .95; margin-top: 4px"
                  :percent="metric.percent"
                  :show-percent="false"
                  :color="metric.color"
                />

              </div>

              <div class="col-3 p-0 pl-2 m-0">
                <div class="small font-weight-bold" style="font-size: 10px; opacity: .8">
                  {{ metric.percent }} %
                </div>
              </div>
            </div>

            <span class="text-muted small">
              Uptime ({{formatProcessUptime(p.elapsed)}})
            </span>
          </div>
        </td>
      </tr>
      </tbody>
    </table>
  </eq-window>
</template>

<script>

import {EventBus}    from "@/app/event-bus/event-bus";
import {SpireApi}    from "@/app/api/spire-api";
import EqWindow      from "@/components/eq-ui/EQWindow.vue";
import EqProgressBar from "@/components/eq-ui/EQProgressBar.vue";
import Time          from "@/app/time/time";

export default {
  name: 'DashboardProcessCounts',
  components: { EqProgressBar, EqWindow },
  data() {
    return {
      processCounts: [],
      loaded: false,
      showStats: false,
    }
  },

  methods: {
    formatProcessUptime(time) {
      return Time.humanizeUnix(time)
    },

    getCpuLoadColor(load) {
      if (load > 80) {
        return 'red'
      }
      if (load > 50) {
        return 'orange'
      }

      return '#2c7be5'
    },
    getMetrics(p) {
      let metrics = []

      if (p.cpu) {
        metrics.push({
          label: 'CPU',
          percent: p.cpu,
          color: this.getCpuLoadColor(p.cpu),
        })
      }

      if (p.memory) {
        metrics.push({
          label: 'MEM',
          percent: p.memory,
          color: 'lightgreen',
        })
      }

      return metrics
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
      const newStats = e.main_process_stats || []

      newStats.forEach((newItem) => {
        const existing = this.processCounts.find(p => p.description === newItem.description)

        if (existing) {
          Object.assign(existing, newItem)
        } else {
          this.processCounts.push(newItem)
        }
      })

      // Remove any old items no longer in the new stats
      this.processCounts = this.processCounts.filter(p =>
        newStats.find(n => n.description === p.description)
      )

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
