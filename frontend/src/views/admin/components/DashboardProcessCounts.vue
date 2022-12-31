<template>
  <div class="col-sm-6 col-lg-6" v-if="processCounts">
    <div class="card">
      <div class="card-header">
        <h4 class="card-header-title">Server Processes</h4>
        <b-spinner variant="primary" label="Spinning" small class="ml-3" v-if="!loaded"></b-spinner>
      </div>
      <table class="table card-table">
        <tbody>

        <!-- List Counts -->
        <tr v-for="(processCount, processName) in processCounts" :key="processName">
          <td> {{ processName.charAt(0).toUpperCase() + processName.slice(1) }}</td>
          <td class="text-right">
            <span class="badge badge-danger"
                  style="font-size: 12px"
                  v-if="processCount === 0">Offline</span>
            <span class="badge badge-success"
                  style="font-size: 12px"
                  v-if="processCount > 0">Online ({{ processCount }})</span>
          </td>
        </tr>
        </tbody>
      </table>
    </div>
  </div>
</template>

<script>

  import Timer              from "@/app/timer/timer";
  import {EqemuAdminClient} from "@/app/api/eqemu-admin-client-occulus";
  import {EventBus}         from "@/app/event-bus/event-bus";

  export default {
    name: 'DashboardProcessCounts',
    data () {
      return {
        processCounts: [],
        loaded: false
      }
    },

    beforeDestroy () {
      clearInterval(Timer.timer['process-counts'])
      EventBus.$off('process-change')
    },

    /**
     * Mounted
     */
    mounted () {
      let self = this
      EventBus.$on('process-change', async function (event) {
        self.processCounts = await EqemuAdminClient.getProcessCounts()
      })
    },

    /**
     * Create
     */
    async created () {
      this.processCounts = await EqemuAdminClient.getProcessCounts()
      this.loaded        = true

      /**
       * Timer update
       * @type {default}
       */
      let self = this;

      if (Timer.timer['process-counts']) {
        clearInterval(Timer.timer['process-counts'])
      }

      Timer.timer['process-counts'] = setInterval(async function () {
        self.loaded = false
        if (!document.hidden) {
          self.processCounts = await EqemuAdminClient.getProcessCounts()
        }
        self.loaded = true
      }, 5000)
    },
  }
</script>

<style scoped>

</style>
