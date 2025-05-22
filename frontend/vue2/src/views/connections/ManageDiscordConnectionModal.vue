<template>
  <b-modal
    v-if="connection && Object.keys(connection).length > 0"
    id="manage-discord-connection-modal"
    centered
    :title="`Manage Discord Webhook [${connection.database_connection.name}]`"
    size="xl"
    @show="init()"
  >
    <b-form-input
      v-model="webhookUrl"
      id="user-search"
      @keyup="update()"
      placeholder="https://discord.com/api/webhooks/xxx/xxx"
    />

    <div class="text-center">
      <a
        class="btn btn-sm btn-white mt-3 outline-primary"
        href="https://support.discord.com/hc/en-us/articles/228383668-Intro-to-Webhooks"
        target="discord"
      >
        Introduction to Discord Webhooks
      </a>
    </div>

    <info-error-banner
      :slim="true"
      :notification="notification"
      :error="error"
      @dismiss-error="error = ''"
      @dismiss-notification="notification = ''"
      class="mt-3"
    />




    <template #modal-footer>
      <div class="">

      </div>
    </template>
  </b-modal>
</template>

<script>
import InfoErrorBanner from "@/components/InfoErrorBanner";
import {SpireApi}      from "@/app/api/spire-api";

export default {
  name: "ManageDiscordConnectionModal",
  components: { InfoErrorBanner },
  props: {
    connection: {
      type: Object,
    },
  },
  data() {
    return {

      webhookUrl: "",

      // notification / errors
      notification: "",
      error: "",
    }
  },
  watch: {
    connection: {
      handler() {
        this.init()
      },
      deep: true
    },
  },

  mounted() {
    this.init()
  },
  methods: {
    async update() {
      try {
        const r = await SpireApi.v1().post(`connection/${this.connection.server_database_connection_id}/discord-webhook`, {
          webhook_url: this.webhookUrl
        })
        if (r.status === 200) {
          this.notification = r.data
        }
      } catch (err) {
        // error notify
        if (err.response && err.response.data && err.response.data.error) {
          this.error = err.response.data.error
        }
      }
    },

    async init() {
      this.webhookUrl = ""

      if (this.connection.server_database_connection_id > 0) {
        try {
          const r = await SpireApi.v1().get(`connection/${this.connection.server_database_connection_id}/discord-webhook`)
          if (r.status === 200) {
            this.webhookUrl = r.data
          }
        } catch (err) {
          // error notify
          if (err.response && err.response.data && err.response.data.error) {
            this.error = err.response.data.error
          }
        }
      }
    },

  }
}
</script>

<style scoped>

</style>
