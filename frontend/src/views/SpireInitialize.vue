<template>
  <content-area class="text-center fade-in">
    <div class="row justify-content-center">
      <div class="col-sm-12 col-lg-6 justify-content-center">
        <h1 style="font-size: 100px" class="text-center eq-header small-mobile">Spire Setup</h1>

        <info-error-banner
          :slim="true"
          :notification="notification"
          :error="error"
          @dismiss-error="error = ''"
          @dismiss-notification="notification = ''"
          class="mt-0"
        />

        <div style="height: 85vh; overflow-y: scroll; overflow-x: hidden">
          <eq-window>
            <p style="font-size: 14px" class="m-0">
              It appears that you are using Spire for the first time. Let's walk you through a brief setup...
            </p>
          </eq-window>

          <eq-window title="Settings" class="mt-5">

            <!-- Auth enabled -->
            <div class="row mt-3 unselectable">
              <div class="col-6 text-right">
                <div class="font-weight-bold mb-3">Authentication</div>
                Do you want Spire to run with authentication? <br>
                <small class="text-muted">(This will require a user login in order to use this instance of
                  Spire.)</small>
              </div>
              <div class="col-6 text-left">
                <eq-checkbox
                  class="mt-4"
                  v-model.number="form.auth_enabled"
                />
              </div>
            </div>

            <!-- Username -->
            <div class="row mt-3 unselectable" v-if="form.auth_enabled">
              <div class="col-6 text-right">
                <div class="font-weight-bold">Username</div>
                <small class="text-muted">(This will be the administrator user for your Spire)</small>
                <small class="d-block">If you are using this only for yourself locally, you don't need
                  authentication.</small>
              </div>
              <div class="col-6 text-left">
                <input
                  class="form-control"
                  v-model="form.username"
                  placeholder="Your username..."
                />
              </div>
            </div>

            <!-- Password -->
            <div class="row mt-3 unselectable" v-if="form.auth_enabled">
              <div class="col-6 text-right">
                <div class="font-weight-bold mb-3">Password</div>
              </div>
              <div class="col-6 text-left">
                <input
                  class="form-control"
                  type="password"
                  v-model="form.password"
                  placeholder="Your password..."
                />
              </div>
            </div>

          </eq-window>

          <eq-window title="Database" class="mt-5">
            <small>
              Spire will use your database to store settings and configuration as it relates to Spire and will use your
              EverQuest Emulator database instance for communication with tooling. This user must have access to create
              tables within the following database.
            </small>

            <div
              class="mt-3"
              v-if="connection && Object.keys(connection).length > 0"
            >
              <div class="row" v-for="f in connfields">
                <div class="col-6 text-right">
                  <div class="font-weight-bold">{{ f.field }}</div>
                </div>
                <div class="col-6 text-left" v-if="f.value && connection[f.value]">
                  {{ formatConnectionValue(f.field, connection[f.value]) }}
                </div>
                <div class="col-6 text-left" v-if="f.value && !connection[f.value]">
                  {{ f.value }}
                </div>
              </div>
            </div>
          </eq-window>

          <eq-window title="Installation Actions" class="mt-5">
            <div
              class="mt-3"
              v-if="form.auth_enabled"
            >
              <small class="font-weight-bold">
                Spire will enable authentication
              </small>
            </div>

            <div
              class="mt-3"
              v-if="form.auth_enabled && form.username && form.username.length > 0"
            >
              <small class="font-weight-bold">
                Spire will create the following admin user ({{ form.username }})
              </small>
            </div>

            <div v-if="tables" class="mt-3">
              <small class="font-weight-bold">
                Spire install the following tables ({{ tables.length }})
              </small>
            </div>

            <div class="row justify-content-center">
              <div class="col-lg-4 col-sm-12">
                <div
                  class="mt-3"
                  v-if="tables && Object.keys(tables).length > 0"
                >
                  <div class="text-left" v-for="t in tables">
                    <div class="">• {{ t }}</div>
                  </div>
                </div>
              </div>
            </div>

            <button
              :disabled="result && Object.keys(result).length > 0"
              :style="'opacity: ' + (result && Object.keys(result).length > 0 ? '.5' : '1')"
              class='eq-button'
              @click="finish"
            >Finish
            </button>

            <!-- Debug -->

            <eq-window v-if="debug">
              <eq-debug :data="result" class="text-left"/>
              <eq-debug :data="form" class="text-left"/>
              <eq-debug :data="connection" class="text-left"/>
              <eq-debug :data="tables" class="text-left"/>
            </eq-window>

          </eq-window>
        </div>


      </div>
    </div>
  </content-area>
</template>

<script>
import ContentArea     from "../components/layout/ContentArea";
import {AppEnv}        from "@/app/env/app-env";
import {ROUTE}         from "@/routes";
import EqWindow        from "@/components/eq-ui/EQWindow.vue";
import EqCheckbox      from "@/components/eq-ui/EQCheckbox.vue";
import {SpireApi}      from "@/app/api/spire-api";
import EqDebug         from "@/components/eq-ui/EQDebug.vue";
import {LocalSettings} from "@/app/local-settings/localsettings";
import InfoErrorBanner from "@/components/InfoErrorBanner.vue";

export default {
  name: 'Login.vue',
  components: { InfoErrorBanner, EqDebug, EqCheckbox, EqWindow, ContentArea },
  data() {
    return {
      debug: false,

      githubAuthEnabled: AppEnv.isGithubAuthEnabled(),
      ROUTE: ROUTE,

      // api responses
      error: "",
      notification: "",

      // form data
      form: {
        auth_enabled: 0,
        username: "",
        password: "",
      },

      connection: {},
      tables: {},

      result: {},

      connfields: [
        { field: "Source", value: "Emulator Server Configuration" },
        { field: "Host", value: "Addr" },
        { field: "Database", value: "DBName" },
        { field: "Collation", value: "Collation" },
        { field: "Max Allowed Packet", value: "MaxAllowedPacket" },
      ]
    }
  },
  async mounted() {
    this.debug = LocalSettings.get("debug-mode") === "true"

    const r = await SpireApi.v1().get("/app/onboarding-info")
    if (r.data && r.data.data) {
      this.connection = r.data.data.connection_info
      this.tables     = r.data.data.tables
    }
  },
  methods: {
    async finish() {
      if (confirm(`Are you sure this is how you want to install your Spire installation?`)) {
        try {
          const r = await SpireApi.v1().post("/app/onboard-initialize", this.form)
          if (r.data && r.status === 200) {
            this.result       = r.data
            this.notification = "Installation succeeded! Redirecting momentarily..."

            const init = await AppEnv.init()
            if (init) {
              setTimeout(() => {
                this.$router.push(ROUTE.HOME).catch((e) => {
                })
              }, 3000)
            }
          }
        } catch (err) {
          // error notify
          if (err.response && err.response.data && err.response.data.error) {
            this.error = err.response.data.error
          }
        }
      }
    },

    hasAuthOptions() {
      return this.githubAuthEnabled
    },
    formatConnectionValue(field, value) {
      if (field === "Max Allowed Packet") {
        return this.formatBytes(value)
      }

      return value;
    },
    formatBytes(bytes, decimals = 2) {
      if (!+bytes) return '0 Bytes'

      const k     = 1024
      const dm    = decimals < 0 ? 0 : decimals
      const sizes = ['Bytes', 'KB', 'MB', 'GB', 'TB', 'PB', 'EB', 'ZB', 'YB']

      const i = Math.floor(Math.log(bytes) / Math.log(k))

      return `${parseFloat((bytes / Math.pow(k, i)).toFixed(dm))} ${sizes[i]}`
    }
  }
}
</script>
