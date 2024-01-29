<template>
  <b-modal
    centered
    no-fade
    :ok-disabled="true"
    id="app-update-modal"
    title="Update Spire"
    header-text-variant="dark"
    body-text-variant="dark"
    cancel-title="Ignore Update"
    size="xl"
    @cancel="ignoreUpdate"
  >

    <div v-if="!reloading">
      There is a new Spire update available

      <div v-if="release && Object.keys(release).length > 0" class="mt-3">
        <div>
          <span class="font-weight-bold">Current Version</span> {{ currentVersion }}
        </div>
        <div>
          <span class="font-weight-bold">Latest Version</span> {{ release.tag_name.replace("v", "") }}
        </div>
      </div>

      <div class="mt-3">
        <h4>Updating Spire</h4>
        Once Spire is done updating, it will exit and you will need to restart it manually. If you have Spire being ran
        by a process manager or running it under akk-stack, it will automatically restart.
      </div>

      <div v-if="releaseNotes" class="mt-3">
        <div>
          <span class="font-weight-bold">Release Notes</span>
        </div>
        <div class="row" id="changelog">
          <div class="col-12">
            <div v-html="releaseNotes" class="mt-3 changelog markdown-body"></div>
          </div>
        </div>
      </div>

      <div class="mt-3" v-if="updating">
        <div class="text-center">
          <h4 class="text-muted">Updating Spire</h4>
        </div>

        <loader-fake-progress/>
      </div>
    </div>

    <div v-if="reloading">
      Spire has been updated, waiting for Spire to restart to reload the page.
    </div>

    <div>
      <info-error-banner
        :slim="true"
        :notification="notification"
        :error="error"
        @dismiss-error="error = ''"
        @dismiss-notification="notification = ''"
        class="mt-0"
      />
    </div>

    <template #modal-footer="{ cancel, ok }">
      <b-button
        @click="ignoreUpdate"
        variant="outline-secondary"
      >
        <i class="fe fe-x"></i> Skip Update
      </b-button>

      <b-button
        @click="updateSpire"
        variant="primary"
      >
        <i class="fe fe-download"></i> Update
      </b-button>

    </template>

  </b-modal>
</template>

<script>

import {LocalSettings}    from "@/app/local-settings/localsettings";
import {SpireApi}         from "@/app/api/spire-api";
import InfoErrorBanner    from "@/components/InfoErrorBanner.vue";
import LoaderFakeProgress from "@/components/LoaderFakeProgress.vue";

export default {
  name: "AppUpdateModal",
  components: { LoaderFakeProgress, InfoErrorBanner },
  props: {
    release: {
      type: Object,
      required: true
    },
    currentVersion: {
      type: String,
      required: true
    }
  },
  data() {
    return {
      releaseNotes: "",
      updating: false,
      reloading: false,

      // notification / errors
      notification: "",
      error: "",
    }
  },
  watch: {
    'release'() {
      const md = require("markdown-it")({
        html: true,
        xhtmlOut: false,
        breaks: true,
        typographer: false,
        linkify: true
      });

      let markdownRaw   = md.render(this.release.body);
      this.releaseNotes = "<div>" + markdownRaw + "</div>"
    }
  },

  methods: {
    async updateSpire() {
      this.updating = true;

      try {
        const r = await SpireApi.v1().post('app/update')
        if (r.status === 200) {
          LocalSettings.clearUpdateVariables()
          this.updating  = false
          this.reloading = true

          setInterval(async () => {
            // poll app/env to see if spire is running
            // if it is, then we need to reload our page
            const r = await SpireApi.v1().get('app/env')
            if (r.status === 200) {
              location.reload()
            }
          }, 1000)
        }
      } catch (e) {
        // error notify
        if (e.response && e.response.data && e.response.data.error) {
          this.error = e.response.data.error
        }
      }

      //
    },
    ignoreUpdate() {
      // set ignored version in local browser storage
      LocalSettings.setIgnoredUpdateVersion(this.release.tag_name.replace("v", ""))
      this.$bvModal.hide('app-update-modal')
    }
  },

}
</script>

<style scoped>

</style>
