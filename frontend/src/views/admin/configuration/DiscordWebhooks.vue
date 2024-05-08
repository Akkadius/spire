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

    <eq-window title="Logging Webhooks" class="p-3 pt-4">
      <div class="form-group col-md-12 mb-0 p-0">

        <div class="row">
          <div class="col-12">
            <b-button
              size="sm"
              variant="outline-warning btn-dark"
              @click="add()"
            >
              <i class="fa fa-plus mr-1"></i>
              Add Webhook
            </b-button>
          </div>
        </div>

        <table class="eq-table eq-highlight-rows bordered log-settings mt-3" style="table-layout: fixed">
          <thead class="eq-table-floating-header">
          <tr>
            <th style="width: 20px"></th>
            <th class="text-center" style="width: 20px">ID</th>
            <th style="width: 100px; white-space: normal; ">
              Webhook Name
              <small class="text-muted d-block">
                This is what you name it and what you use to call out the webhook using Quest methods
              </small>
            </th>
            <th style="width: 20vw">
              Webhook URL

              <small class="text-muted d-block">
                The webhook URL. In Discord this is mapped to a single channel.
                <a
                  class=""
                  href="https://support.discord.com/hc/en-us/articles/228383668-Intro-to-Webhooks"
                  target="discord"
                >
                  Look here for an introduction to Discord webhooks.
                </a>

              </small>

            </th>
          </tr>

          </thead>
          <tbody>
          <tr v-for="w in discordWebhooks">
            <td class="text-center p-0 m-0">
              <b-button
                variant="primary"
                class="btn-dark btn-sm btn-outline-danger ml-1"
                style="padding: 0px 6px;"
                title="Delete"
                @click="deleteWebhook(w)"
              >
                <i class="fa fa-trash"></i>
              </b-button>

            </td>
            <td class="text-center">{{ w.id }}</td>
            <td>
              <input
                @change="updateWebhook(w)"
                type="text"
                class="form-control"
                v-model="w.webhook_name"
              >
            </td>
            <td>
              <input
                @change="updateWebhook(w)"
                type="text"
                class="form-control"
                v-model="w.webhook_url"
              >
            </td>
          </tr>
          </tbody>
        </table>
      </div>
    </eq-window>

  </div>
</template>

<script>
import EqWindow            from "@/components/eq-ui/EQWindow.vue";
import {DiscordWebhookApi} from "@/app/api/api/discord-webhook-api";
import {SpireApi}          from "@/app/api/spire-api";
import InfoErrorBanner     from "@/components/InfoErrorBanner.vue";

export default {
  components: { InfoErrorBanner, EqWindow },
  data() {
    return {
      loaded: false,

      discordWebhooks: [],

      // notification / errors
      notification: "",
      error: "",
    }
  },
  async created() {
    this.loaded = true

    this.loadWebhooks()
  },
  methods: {
    async deleteWebhook(e) {
      if (confirm(`Are you sure you want to delete this webhook?\n\n${e.webhook_name} (${e.id})`)) {
        try {
          const r = await (new DiscordWebhookApi(...SpireApi.cfg())).deleteDiscordWebhook(
            {
              id: e.id
            }
          )
          if (r.status === 200) {
            this.notification = `Deleted Discord Webhook [${e.webhook_name}] (${e.id})!`
            this.loadWebhooks()

            const r = await SpireApi.v1().post("eqemuserver/reload/logs")
            if (r.status === 200) {
              setTimeout(() => {
                this.notification = "Server log settings reloaded in-game!"
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
    },

    async updateWebhook(e) {
      try {
        const r = await (new DiscordWebhookApi(...SpireApi.cfg())).updateDiscordWebhook(
          {
            id: e.id,
            discordWebhook: e
          }
        )
        if (r.status === 200) {
          this.notification = `Updated Discord Webhook [${e.webhook_name}] (${e.id})!`
          this.loadWebhooks()

          const r = await SpireApi.v1().post("eqemuserver/reload/logs")
          if (r.status === 200) {
            setTimeout(() => {
              this.notification = "Server log settings reloaded in-game!"
            }, 1000)
          }
        }
      } catch (e) {
        // error notify
        if (e.response && e.response.data && e.response.data.error) {
          this.error = e.response.data.error
        }
      }
    },

    async loadWebhooks() {
      let r = await (new DiscordWebhookApi(...SpireApi.cfg())).listDiscordWebhooks()
      if (r.status === 200) {
        this.discordWebhooks = r.data
      }
    },

    async add() {
      try {
        const r = await (new DiscordWebhookApi(...SpireApi.cfg())).createDiscordWebhook(
          {
            discordWebhook: {
              webhook_name: "New Webhook",
              webhook_url: "",
            }
          }
        )
        if (r.status === 200) {
          this.notification = "New webhook added!"
          this.loadWebhooks()
        }
      } catch (e) {
        // error notify
        if (e.response && e.response.data && e.response.data.error) {
          this.error = e.response.data.error
        }
      }
    },
  }
}
</script>
