<template>
  <eq-window>

    <div
      class="row justify-content-center"
      style="position: absolute; top: 0%; z-index: 9999999; width: 100%"
    >
      <div class="col-6">
        <info-error-banner
          style="width: 100%"
          :slim="true"
          :notification="notification"
          :error="error"
          @dismiss-error="error = ''"
          @dismiss-notification="notification = ''"
          class="mt-3"
        />
      </div>
    </div>

    <div style="max-height: 80vh; overflow-y: scroll; overflow-x: hidden">

      <div v-if="reloadTypes && reloadTypes.length === 0" class="text-center font-weight-bold">
        This menu requires a connection to the World Server
      </div>

      <table
        v-if="reloadTypes && reloadTypes.length > 0"
        class="eq-table eq-highlight-rows bordered log-settings minified-inputs"
      >
        <thead class="eq-table-floating-header">
        <tr>
          <th class="text-center" style="width: 120px">Reload</th>
          <th class="">Description</th>
        </tr>

        </thead>
        <tbody>
        <tr
          v-for="r in reloadTypes"
          :key="r.command"
        >
          <td class="text-right">
            <button class="eq-button" @click="reload(r)">Reload</button>
          </td>
          <td>
            Reload <span class="font-weight-bold">{{ r.description }}</span> globally
          </td>
        </tr>
        </tbody>
      </table>
    </div>

  </eq-window>
</template>

<script>
import EqWindow        from "@/components/eq-ui/EQWindow.vue";
import {SpireApi}      from "@/app/api/spire-api";
import InfoErrorBanner from "@/components/InfoErrorBanner.vue";

export default {
  name: "ServerReload",
  components: { InfoErrorBanner, EqWindow },
  data() {
    return {
      reloadTypes: [],

      // notification / errors
      notification: "",
      error: "",
    }
  },
  mounted() {
    this.getReloadTypes()
  },
  methods: {
    async getReloadTypes() {
      try {
        const r = await SpireApi.v1().get("eqemuserver/reload-types")
        if (r.status === 200) {
          this.reloadTypes = r.data.data
        }
      } catch (e) {
        if (e.response && e.response.data && e.response.data.error) {
          if (e.response.data.error.includes("Failed to connect to gameserver")) {
            this.error = ""
          } else {
            this.error = e.response.data.error
          }
        }
      }
    },
    async reload(e) {
      const r = await SpireApi.v1().post("eqemuserver/reload/" + e.command)
      if (r.status === 200) {
        this.notification = r.data.data.message.replaceAll("Reloading", "Reloaded")
      }
    }
  }
}
</script>

<style scoped>

</style>
