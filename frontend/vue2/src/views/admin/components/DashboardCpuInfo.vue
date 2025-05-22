<template>
  <eq-window
    :title="'CPU (s) (' + (cpu && cpu.cpu_percents ? cpu.cpu_percents.length : 0) + ') ' + (cpu && cpu.info ? cpu.info[0].modelName : '')"
    class="p-0 mb-4"
  >
    <table
      class="eq-table bordered eq-highlight-rows mb-0 eq-window-cpu"
    >
      <thead class="eq-table-floating-header">
      <tr>
        <th class="text-right">
          CPU
        </th>
        <th class="text-left">
          Percent
        </th>
        <th class="text-left">

        </th>
      </tr>
      </thead>

      <tbody>

      <!-- List Counts -->
      <tr
        v-for="(l, index) in cpu.cpu_percents"
        :key="index"
      >
        <td style="width: 100px" class="text-right font-weight-bold">
          #{{ index + 1 }}
        </td>
        <td style="width: 50px" class="text-right font-weight-bold text-muted">
          {{ (Math.round(l)) }}%
        </td>
        <td
          class="text-left"
        >
          <eq-progress-bar
            style="opacity: .9"
            :percent="(Math.round(l * 100) / 100)"
            :show-percent="false"
            :color="getCpuLoadColor(parseInt(l))"
          />
        </td>
      </tr>

      </tbody>
    </table>

  </eq-window>
</template>

<script>
import {SpireApi}    from "@/app/api/spire-api";
import EqWindow      from "@/components/eq-ui/EQWindow.vue";
import EqProgressBar from "@/components/eq-ui/EQProgressBar.vue";

export default {
  name: 'DashboardCpuInfo',
  components: { EqProgressBar, EqWindow },
  data() {
    return {
      cpu: {},

      timer: null,
    }
  },
  beforeDestroy() {
    clearInterval(this.timer)
  },
  mounted() {
    this.fetchStats()
    this.timer = setInterval(this.fetchStats, 1000)
  },
  methods: {
    fetchStats() {
      if (!document.hidden) {
        SpireApi.v1().get("admin/system/cpu").then((r) => {
          if (r.status === 200) {
            this.cpu = r.data
          }
        })
      }
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

    /**
     * @param input
     * @param length
     */
    truncate: function (input, length) {
      if (input.length > length) {
        return input.substring(0, length) + '...'
      } else {
        return input
      }
    }
  },
}
</script>

<style>
.eq-window-cpu td {
  padding-top: 0 !important;
  padding-bottom: 0 !important;
}

.eq-window-cpu .eq-progress-bar {
  margin: 0 !important;
}
</style>
