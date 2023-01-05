<template>
  <div>
    <eq-window title="Log Settings" class="p-1">

      <info-error-banner
        style="position: absolute; top: -60px; z-index: 9999999; width: 100%"
        :slim="true"
        :notification="notification"
        :error="error"
        @dismiss-error="error = ''"
        @dismiss-notification="notification = ''"
        class="mt-3"
      />

      <div style="height: 84vh; overflow-y: scroll">
      <table class="eq-table eq-highlight-rows bordered log-settings">
        <thead>
        <tr>
          <th>ID</th>
          <th>Category</th>
          <th>Console</th>
          <th>File</th>
          <th>GM (In-Game)</th>
<!--          <th>Discord</th>-->
        </tr>

        </thead>
        <tbody>
        <tr v-for="s in settings">
          <td>{{ s.log_category_id }}</td>
          <td>{{ s.log_category_description }}</td>
          <td>
            <div class="d-inline-block mr-3">
              <eq-checkbox
                label="Off"
                :fade-when-not-true="true"
                class="d-inline-block"
                :true-value="0"
                :false-value="0"
                v-model="s.log_to_console"
                @change="save(s)"
              />
            </div>

            <div class="d-inline-block mr-3">
              <eq-checkbox
                label="Normal"
                :fade-when-not-true="true"
                class="d-inline-block"
                :true-value="1"
                :false-value="0"
                v-model="s.log_to_console"
                @change="save(s)"
              />
            </div>

            <div class="d-inline-block">
              <eq-checkbox
                label="Detail"
                :fade-when-not-true="true"
                class="d-inline-block"
                :true-value="3"
                :false-value="0"
                v-model="s.log_to_console"
                @change="save(s)"
              />
            </div>

          </td>
          <td>
            <div class="d-inline-block mr-3">
              <eq-checkbox
                label="Off"
                :fade-when-not-true="true"
                class="d-inline-block"
                :true-value="0"
                :false-value="0"
                v-model="s.log_to_file"
                @change="save(s)"
              />
            </div>

            <div class="d-inline-block mr-3">
              <eq-checkbox
                label="Normal"
                :fade-when-not-true="true"
                class="d-inline-block"
                :true-value="1"
                :false-value="0"
                v-model="s.log_to_file"
                @change="save(s)"
              />
            </div>

            <div class="d-inline-block">
              <eq-checkbox
                label="Detail"
                :fade-when-not-true="true"
                class="d-inline-block"
                :true-value="3"
                :false-value="0"
                v-model="s.log_to_file"
                @change="save(s)"
              />
            </div>

          </td>
          <td>
            <div class="d-inline-block mr-3">
              <eq-checkbox
                label="Off"
                :fade-when-not-true="true"
                class="d-inline-block"
                :true-value="0"
                :false-value="0"
                v-model="s.log_to_gmsay"
                @change="save(s)"
              />
            </div>

            <div class="d-inline-block mr-3">
              <eq-checkbox
                label="Normal"
                :fade-when-not-true="true"
                class="d-inline-block"
                :true-value="1"
                :false-value="0"
                v-model="s.log_to_gmsay"
                @change="save(s)"
              />
            </div>

            <div class="d-inline-block">
              <eq-checkbox
                label="Detail"
                :fade-when-not-true="true"
                class="d-inline-block"
                :true-value="3"
                :false-value="0"
                v-model="s.log_to_gmsay"
                @change="save(s)"
              />
            </div>
          </td>
<!--          <td>{{ s.log_to_discord }}</td>-->
        </tr>
        </tbody>
      </table>
      </div>

      <eq-debug :data="settings"/>

    </eq-window>
  </div>
</template>

<script>
import EqWindow            from "@/components/eq-ui/EQWindow.vue";
import {SpireApi}          from "@/app/api/spire-api";
import {LogsysCategoryApi} from "@/app/api/api/logsys-category-api";
import EqCheckbox          from "@/components/eq-ui/EQCheckbox.vue";
import EqDebug             from "@/components/eq-ui/EQDebug.vue";
import InfoErrorBanner     from "@/components/InfoErrorBanner.vue";

export default {
  name: "LogSettings",
  components: { InfoErrorBanner, EqDebug, EqCheckbox, EqWindow },
  data() {
    return {
      settings: [],

      // notification / errors
      notification: "",
      error: "",
    }
  },
  async mounted() {
    const r = await (new LogsysCategoryApi(...SpireApi.cfg())).listLogsysCategories()
    if (r.status === 200) {
      this.settings = r.data
    }
  },
  methods: {
    async save(e) {
      console.log(e)
      try {
        const r = await (new LogsysCategoryApi(...SpireApi.cfg()))
          .updateLogsysCategory(
            {
              id: e.log_category_id,
              logsysCategory: e
            }
          )
        if (r.status === 200) {
          this.notification = "Settings updated!"
        }
      } catch (e) {
        // error notify
        if (e.response && e.response.data && e.response.data.error) {
          this.error = e.response.data.error
        }
      }
    }
  }
}
</script>

<style scoped>
.log-settings td, .log-settings th {
  text-align: center;
}
</style>
