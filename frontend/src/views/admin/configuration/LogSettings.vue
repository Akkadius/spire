<template>
  <div>
    <div
      class="row justify-content-center"
      style="position: absolute; top: 0%; z-index: 9999999; width: 100%"
    >
      <div class="col-4">
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

    <eq-window title="Log Settings" class="p-1 minified-inputs">

      <div style="height: 84vh; overflow-y: scroll">
        <table class="eq-table eq-highlight-rows bordered log-settings">
          <thead class="eq-table-floating-header">
          <tr>
            <th class="text-right" style="width: 200px">(ID) Category</th>
            <th style="width: 230px">Console Log Level</th>
            <th style="width: 230px">File Log Level</th>
            <th style="width: 230px">GM (In-Game) Log Level</th>
            <th style="width: 230px">Discord Log Level</th>
            <th>Discord Webhook</th>
          </tr>

          </thead>
          <tbody>
          <tr v-for="s in settings">
            <td class="text-right">({{ s.log_category_id }}) {{ s.log_category_description }}</td>
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
            <td>
              <div class="d-inline-block mr-3">
                <eq-checkbox
                  label="Off"
                  :fade-when-not-true="true"
                  class="d-inline-block"
                  :true-value="0"
                  :false-value="0"
                  v-model="s.log_to_discord"
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
                  v-model="s.log_to_discord"
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
                  v-model="s.log_to_discord"
                  @change="save(s)"
                />
              </div>
            </td>
            <td
              :style="(s.log_to_discord > 0 && s.discord_webhook_id === 0 ? 'color: red' : '')"
              :title="(s.log_to_discord > 0 && s.discord_webhook_id === 0 ? 'Webhook needs to be assigned' : '')"
            >
              <select
                :style="(s.log_to_discord > 0 && s.discord_webhook_id === 0 ? 'border-color: red' : '')"
                :title="(s.log_to_discord > 0 && s.discord_webhook_id === 0 ? 'Webhook needs to be assigned' : '')"
                @change="save(s)"
                class="form-control m-0" v-model="s.discord_webhook_id"
              >
                <option :value="0">--- None ---</option>
                <option
                  v-for="w in discordWebhooks"
                  :key="w.id"
                  :value="w.id"
                >Name [{{ w.webhook_name }}] ID ({{w.id}})
                </option>
              </select>

            </td>
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
import util                from "util";
import {DiscordWebhookApi} from "@/app/api/api/discord-webhook-api";

export default {
  name: "LogSettings",
  components: { InfoErrorBanner, EqDebug, EqCheckbox, EqWindow },
  data() {
    return {
      settings: [],
      discordWebhooks: [],

      // notification / errors
      notification: "",
      error: "",
    }
  },
  async mounted() {
    let r = await (new LogsysCategoryApi(...SpireApi.cfg())).listLogsysCategories()
    if (r.status === 200) {
      this.settings = r.data
    }

    r = await (new DiscordWebhookApi(...SpireApi.cfg())).listDiscordWebhooks()
    if (r.status === 200) {
      this.discordWebhooks = r.data
      console.log(this.discordWebhooks)
    }
  },
  methods: {
    async save(e) {
      console.log(e)

      if (e.log_to_discord > 0 && e.discord_webhook_id === 0) {
        this.error = `Discord Webhook needs to be assigned for this log category [${e.log_category_description}] to work!`
        return;
      }

      try {
        const r = await (new LogsysCategoryApi(...SpireApi.cfg()))
          .updateLogsysCategory(
            {
              id: e.log_category_id,
              logsysCategory: e
            }
          )
        if (r.status === 200) {
          // reset
          this.notification = ""
          // we have to queue timeout to reset the notification dismiss timer
          setTimeout(() => {
            this.notification =
              util.format(
                "Settings updated for [%s] (%s)!",
                e.log_category_description,
                e.log_category_id
              )
          }, 1)

          const r = await SpireApi.v1().post("eqemuserver/reload/logs")
          if (r.status === 200) {
            setTimeout(() => {
              this.notification = "Server logs reloaded in-game!"
            }, 1000)
          }
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
