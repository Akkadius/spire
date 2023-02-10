<template>
  <div>
    <div class="row">
      <div class="col-3">
        <eq-window title="Current Version" style="height: 100%" class="p-0">


          <table class="eq-table eq-highlight-rows bordered m-0 mt-3" v-if="version">
            <tbody>
            <tr>
              <td class="text-right font-weight-bold">Operating System</td>
              <td>{{ version.os }}</td>
            </tr>
            <tr>
              <td class="text-right font-weight-bold">Server Version</td>
              <td>{{ version.server_version }}</td>
            </tr>
            <tr>
              <td class="text-right font-weight-bold">Compile Date</td>
              <td>{{ version.compile_date }}</td>
            </tr>
            <tr>
              <td class="text-right font-weight-bold">Compile Time</td>
              <td>{{ version.compile_time }}</td>
            </tr>
            <tr>
              <td class="text-right font-weight-bold">Database Version</td>
              <td>{{ version.database_version }}</td>
            </tr>
            <tr v-if="version.bots_database_version">
              <td class="text-right font-weight-bold">Bots Database Version</td>
              <td>{{ version.bots_database_version }}</td>
            </tr>
            </tbody>
          </table>

        </eq-window>
      </div>
      <div class="col-9">
        <eq-window title="Update Settings" style="height: 100%">
          <div class="row">
            <div class="col-3 text-right mt-3">
              <eq-checkbox
                label="Use Release Binaries"
                class="d-inline-block"
                true-value="release"
                false-value="self-compiled"
                v-model="updateType"
                @input="setUpdateOption('release')"
              />
            </div>
            <div class="col-9">
              <small class="text-muted">
                Uses official release binaries from
                <a href="https://github.com/EQEmu/Server/releases" target="releases">EverQuest Emulator Server</a>
              </small>
              <div class="mt-3">
                Use this when you are not developing server or making code changes to server code
              </div>
            </div>
          </div>
          <div class="row mt-4">
            <div class="col-3 text-right mt-3">
              <eq-checkbox
                label="Self-Compiled Binaries"
                class="d-inline-block"
                false-value="release"
                true-value="self-compiled"
                v-model="updateType"
                @input="setUpdateOption('self-compiled')"
              />
            </div>
            <div class="col-9">
              <small class="text-muted">
                Compiles binaries locally
              </small>
              <div class="mt-3">
                Use this if you are a developer or intend on making modifications to your server
              </div>
            </div>
          </div>

          <div
            class="row justify-content-center"
            style="position: absolute; bottom: -5%; z-index: 9999999; width: 100%"
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
        </eq-window>
      </div>
    </div>

    <update-releases
      class="mt-5"
      :version="version.server_version"
      v-if="updateType === 'release' && version"
    />

  </div>
</template>

<script>
import EqWindow        from "@/components/eq-ui/EQWindow.vue";
import EqCheckbox      from "@/components/eq-ui/EQCheckbox.vue";
import {SpireApi}      from "@/app/api/spire-api";
import InfoErrorBanner from "@/components/InfoErrorBanner.vue";
import UpdateReleases  from "@/views/admin/server-update/UpdateReleases.vue";
import {AppEnv}        from "@/app/env/app-env";

export default {
  name: "ServerUpdate",
  components: { UpdateReleases, InfoErrorBanner, EqCheckbox, EqWindow },
  data() {
    return {
      updateType: "",

      version: {},

      releases: [],

      // notification / errors
      notification: "",
      error: "",
    }
  },
  created() {
    this.init()
  },
  methods: {
    async init() {
      const v = await SpireApi.v1().get(`eqemuserver/version`)
      if (v.status === 200) {
        this.version = v.data
        this.version.os = AppEnv.getOS()
      }

      const r = await SpireApi.v1().get(`eqemuserver/update-type`)
      if (r.status === 200) {
        this.updateType = r.data.updateType
      }



      // if (this.updateType === "release") {
      //   this.loadReleaseData()
      // }
    },

    async setUpdateOption(option) {
      try {
        const r = await SpireApi.v1().post(`eqemuserver/update-type/${option}`)
        if (r.status === 200) {
          this.notification = r.data.message
        }
      } catch (err) {
        // error notify
        if (err.response && err.response.data && err.response.data.error) {
          this.error = err.response.data.error
        }
      }
    }
  }
}
</script>

<style scoped>

</style>
