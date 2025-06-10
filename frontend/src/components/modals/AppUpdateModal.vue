<template>
  <EqModal
    title="Update Spire"
    @close="ignoreUpdate"
    size="xl"
  >
    <template #body>
      <div v-if="!reloading">
        <p>There is a new Spire update available</p>

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
          <p>
            Once Spire is done updating, it will exit and you will need to restart it manually.
            If you have Spire being run by a process manager or under akk-stack, it will restart automatically.
          </p>
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

      <div v-else>
        Spire has been updated. Waiting for restart to reload the page.<br><br>
        If you are not running Spire under a process manager or akk-stack, you will need to restart it manually.
      </div>

      <info-error-banner
        :slim="true"
        :notification="notification"
        :error="error"
        @dismiss-error="error = ''"
        @dismiss-notification="notification = ''"
        class="mt-3"
      />
    </template>

    <template #footer>
      <div class="mt-3">
        <button
          @click="ignoreUpdate"
          class="btn btn-sm mr-3 btn-default"
          v-if="!reloading"
        >
          <i class="fe fe-x"></i> Skip Update
        </button>

        <button
          @click="updateSpire"
          class="btn btn-sm mr-3 btn-success"
          v-if="!reloading"
        >
          <i class="fe fe-download"></i> Update
        </button>
      </div>
    </template>
  </EqModal>
</template>

<script>
import {LocalSettings} from "@/app/local-settings/localsettings";
import {SpireApi} from "@/app/api/spire-api";
import InfoErrorBanner from "@/components/InfoErrorBanner.vue";
import LoaderFakeProgress from "@/components/LoaderFakeProgress.vue";
import EqModal from "@/components/eq-ui/EQModal.vue";

export default {
  name: "AppUpdateModal",
  components: {
    EqModal,
    LoaderFakeProgress,
    InfoErrorBanner
  },
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
      notification: "",
      error: ""
    };
  },
  watch: {
    release() {
      const md = require("markdown-it")({
        html: true,
        breaks: true,
        linkify: true
      });
      this.releaseNotes = `<div>${md.render(this.release.body)}</div>`;
    }
  },
  methods: {
    async updateSpire() {
      this.updating = true;

      try {
        const r = await SpireApi.v1().post("app/update");
        if (r.status === 200) {
          LocalSettings.clearUpdateVariables();
          this.updating = false;
          this.reloading = true;

          setInterval(async () => {
            const r = await SpireApi.v1().get("app/env");
            if (r.status === 200) {
              location.reload();
            }
          }, 1000);
        }
      } catch (e) {
        if (e.response?.data?.error) {
          this.error = e.response.data.error;
        }
      }
    },
    ignoreUpdate() {
      LocalSettings.setIgnoredUpdateVersion(this.release.tag_name.replace("v", ""));
      this.$emit("close");
    }
  }
};
</script>

<style scoped>
/* Keep styles minimal or reuse existing app theme */
</style>
