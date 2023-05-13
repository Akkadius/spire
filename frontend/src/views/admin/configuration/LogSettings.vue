<template>
  <div>

    <eq-window title="Log Settings">

      <div
        class="row justify-content-center"
        style="position: absolute; top: 0; z-index: 9999999; width: 100%"
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
        <div class="row mb-3">
          <div class="col-11">
            <b-form-input
              type="text"
              class="form-control list-search"
              @keyup="updateQueryState()"
              v-model="search"
              placeholder="Search log settings..."
              autofocus
            />
          </div>

          <div class="col-1">
            <button
              title="Reset"
              class="eq-button m-0"
              @click="search = ''; updateQueryState()"
            ><i class="fa fa-refresh"></i> Reset
            </button>
          </div>
        </div>


        <table class="eq-table eq-highlight-rows bordered log-settings minified-inputs">
          <thead class="eq-table-floating-header">
          <tr>
            <th class="text-right" style="width: 200px">(ID) Category</th>
            <th style="width: 230px">Console Log Level</th>
            <th style="width: 230px">File Log Level</th>
            <th style="width: 230px">GM (In-Game) Log Level</th>
            <th style="width: 230px">Discord Log Level</th>
            <th>
              <router-link
                style="color: #8aa3ff"
                :to="ROUTE.ADMIN_DISCORD_WEBHOOK_SETTINGS"
              >
                Discord Webhook
                <i class="fa fa-arrow-right"></i>
              </router-link>
            </th>
          </tr>

          </thead>
          <tbody>
          <tr
            v-for="s in filteredSettings(settings)"
            :key="s.log_category_id"
          >
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
                >({{ w.id }}) {{ w.webhook_name }}
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
import {ROUTE}             from "@/routes";

export default {
  name: "LogSettings",
  components: { InfoErrorBanner, EqDebug, EqCheckbox, EqWindow },
  data() {
    return {
      search: "",

      ROUTE: ROUTE,

      settings: [],
      discordWebhooks: [],

      // notification / errors
      notification: "",
      error: "",
    }
  },
  async mounted() {
    this.loadQueryState()

    try {
      let r = await (new LogsysCategoryApi(...SpireApi.cfg())).listLogsysCategories()
      if (r.status === 200) {
        this.settings = r.data
      }

      r = await (new DiscordWebhookApi(...SpireApi.cfg())).listDiscordWebhooks()
      if (r.status === 200) {
        this.discordWebhooks = r.data
      }
    } catch (e) {
      // error notify
      if (e.response && e.response.data && e.response.data.error) {
        this.error = e.response.data.error
      }
    }
  },
  methods: {
    filteredSettings(s) {
      return s.filter((e) => {
        if (this.search && this.search.length > 0) {
          return e.log_category_description.toLowerCase().includes(this.search.toLowerCase())
        }

        return e
      })
    },

    updateQueryState() {
      let q = {};

      if (this.search !== "") {
        q.search = this.search
      }

      this.$router.push(
        {
          path: ROUTE.ADMIN_LOG_SETTINGS,
          query: q
        }
      ).catch(() => {
      })
    },

    loadQueryState() {
      if (this.$route.query.search && this.$route.query.search.length > 0) {
        this.search = this.$route.query.search
      }
    },

    async save(e) {
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
