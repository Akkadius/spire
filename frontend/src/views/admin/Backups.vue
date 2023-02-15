<template>
  <div>
    <eq-window title="Manual Backups">

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

      <table
        class="eq-table eq-highlight-rows bordered log-settings minified-inputs mt-3"
      >
        <thead class="eq-table-floating-header">
        <tr>
          <th class="text-center" style="width: 120px">Asset</th>
          <th class="">Description</th>
        </tr>

        </thead>
        <tbody>
        <tr
          v-for="r in backups"
          :key="r.path"
        >
          <td class="text-center">
            <a class="eq-button" @click="download(r)" :target="r.path">{{ r.asset }}</a>
          </td>
          <td>
            <span class="font-weight-bold">{{ r.description }}</span>
          </td>
        </tr>
        </tbody>
      </table>

      <div v-if="downloading" class="mt-3">
        <div class="row">
          <div class="col-1 text-left">
            <span>Downloading...</span>
            <loader-fake-progress class="mt-3"/>
          </div>
        </div>

      </div>

    </eq-window>
  </div>
</template>

<script>
import EqWindow           from "@/components/eq-ui/EQWindow.vue";
import axios              from "axios";
import {SpireApi}         from "@/app/api/spire-api";
import LoaderFakeProgress from "@/components/LoaderFakeProgress.vue";
import InfoErrorBanner    from "@/components/InfoErrorBanner.vue";

export default {
  name: "Backups",
  components: { InfoErrorBanner, LoaderFakeProgress, EqWindow },
  data() {
    return {
      backups: [
        {
          asset: "Quests",
          description: "Downloads a compressed .zip of your Quests directory",
          path: SpireApi.getBaseV1Path() + '/eqemuserver/manual-backup/quests'
        },
        {
          asset: "Maps",
          description: "Downloads a compressed .zip of your Maps directory",
          path: SpireApi.getBaseV1Path() + '/eqemuserver/manual-backup/maps'
        },
      ],

      downloading: false,

      // notification / errors
      notification: "",
      error: "",
    }
  },
  methods: {
    async download(r) {
      try {
        await this.getFile(r.path)
      } catch (e) {
        // error notify
        if (e.response && e.response.data && e.response.data.error) {
          this.error = e.response.data.error
        }
      }
    },
    async getFile(url) {
      this.downloading    = true
      let config          = SpireApi.getAxiosConfig()
      config.responseType = 'blob';
      axios.create(config).get(url,
        {
          onDownloadProgress: (e) => {
            console.log(e)
            this.downloadPercent = Math.round((e.loaded * 100) / e.total)
          }
        }
      ).then((response) => {
        // create file link in browser's memory
        const href             = URL.createObjectURL(response.data);
        const fileDownloadName = response.headers["content-disposition"].split("filename=")[1].replaceAll("\"", "")
        // create "a" HTML element with href to file & click
        const link             = document.createElement('a');
        link.href              = href;
        link.setAttribute('download', fileDownloadName); //or any other extension
        document.body.appendChild(link);
        link.click();

        // clean up "a" element & remove ObjectURL
        document.body.removeChild(link);
        URL.revokeObjectURL(href);

        this.downloading = false
        this.notification = "File downloaded!"
      });

    }
  }
}
</script>

<style scoped>

</style>
