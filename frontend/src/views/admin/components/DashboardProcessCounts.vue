<template>
  <div v-if="processCounts">
    <div class="card">
      <div class="card-header">
        <h4 class="card-header-title">Server Processes</h4>
        <!--        <b-spinner variant="primary" label="Spinning" small class="ml-3" v-if="!loaded"></b-spinner>-->
      </div>
      <table class="table card-table">
        <tbody>

        <!-- List Counts -->
        <tr v-for="(processCount, processName) in processCounts" :key="processName">
          <td> {{ processName.charAt(0).toUpperCase() + processName.slice(1) }}</td>
          <td class="text-right">
            <span
              class="badge badge-danger"
              style="font-size: 12px"
              v-if="processCount === 0"
            >Offline</span>
            <span
              class="badge badge-success"
              style="font-size: 12px"
              v-if="processCount > 0"
            >Online ({{ processCount }})</span>
          </td>
        </tr>
        </tbody>
      </table>
    </div>
  </div>
</template>

<script>

import Timer           from "@/app/timer/timer";
import {OcculusClient} from "@/app/api/eqemu-admin-client-occulus";
import {EventBus}      from "@/app/event-bus/event-bus";
import {OS}            from "@/app/os/os";

export default {
  name: 'DashboardProcessCounts',
  data() {
    return {
      processCounts: [],
      loaded: false
    }
  },

  beforeDestroy() {
    clearInterval(Timer.timer['process-counts'])
    EventBus.$off('process-change')
  },

  /**
   * Mounted
   */
  mounted() {
    EventBus.$on('process-change', async (event) => {
      this.processCounts = await OcculusClient.getProcessCounts()
    })
  },

  /**
   * Create
   */
  async created() {
    this.processCounts = await OcculusClient.getProcessCounts()
    this.loaded        = true

    /**
     * Timer update
     * @type {default}
     */
    if (Timer.timer['process-counts']) {
      clearInterval(Timer.timer['process-counts'])
    }

    Timer.timer['process-counts'] = setInterval(async () => {
      this.loaded = false
      if (!document.hidden) {
        this.processCounts = await OcculusClient.getProcessCounts()
      }
      this.loaded = true
    }, (OS.get() === "Linux" ? 1000 : 5000))
  },
}
</script>

<style scoped>

</style>
