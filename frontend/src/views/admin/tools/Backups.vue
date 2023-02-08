<template>
  <div class="card">

    <div class="card-header">
      <h4 class="card-header-title">Backups</h4>
    </div>

    <div class="card-body">
      <!-- Row -->

      <div class="row">
        <div class="col-lg-12">
          <table class="table card-table table-vcenter">
            <tbody>
            <tr v-for="(backupType, index) in simpleDownloadTypes" :key="index">
              <td style="width: 300px">
                <button
                  :disabled="awaitingDownload"
                  type="button" class="btn btn-outline-primary mb-2"
                  style="width:100%"
                  @click="downloadBackup(backupType)"
                  v-bind:class="{ 'btn-loading': downloadStatus[backupType] }"
                >
                  <i class="fe fe-download"></i>
                  Download {{ ucFirst(backupType) }}
                </button>

                <div v-show="downloadStatus[backupType]" style="text-align:center">
                  <div class="progress progress-sm mt-3">
                    <div
                      class="progress-bar bg-green"
                      v-bind:style="{ width: fileDownloadProgress + '%' }"
                    ></div>
                  </div>
                  <span v-show="awaitingDownload">Download is preparing...</span>
                  {{ Math.round(downloadedBytes / 1024 / 1024) }}MB /
                  {{ Math.round(totalBytes / 1024 / 1024) }}MB
                </div>
              </td>
              <td>
                Downloads a ZIP archive of your server {{ ucFirst(backupType) }}
              </td>
            </tr>
            <tr>
              <td>
                <button
                  :disabled="awaitingDownload"
                  type="button" class="btn btn-outline-primary mb-2"
                  style="width:100%"
                  @click="downloadBackup('full')"
                  v-bind:class="{ 'btn-loading': downloadStatus['full'] }"
                >
                  <i class="fe fe-download"></i>
                  Full Server Download
                </button>

                <div v-show="downloadStatus['full']" style="text-align:center">
                  <div class="progress progress-sm mt-3">
                    <div
                      class="progress-bar bg-green"
                      v-bind:style="{ width: fileDownloadProgress + '%' }"
                    ></div>
                  </div>
                  <span v-show="awaitingDownload">Download is preparing...</span>
                  {{ Math.round(downloadedBytes / 1024 / 1024) }}MB /
                  {{ Math.round(totalBytes / 1024 / 1024) }}MB
                </div>
              </td>
              <td>
                Downloads a ZIP archive of your entire server, including the database
              </td>
            </tr>
            </tbody>
          </table>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import * as util       from 'util'
import {OcculusClient} from "@/app/api/eqemu-admin-client-occulus";

export default {
  data() {
    return {
      downloadStatus: {},
      backupQueued: false,
      awaitingDownload: false,
      fileDownloadProgress: 0,
      lastPercent: 0,
      downloadedBytes: 0,
      downloadedMbytes: 0,
      totalBytes: 0,
      simpleDownloadTypes: ['quests', 'maps', 'database']
    }
  },
  created() {
    window.addEventListener('file-download-progress', this.handleFileDownloadProgress)
  },
  destroyed() {
    window.removeEventListener('file-download-progress', this.handleFileDownloadProgress, false)
  },
  methods: {
    downloadBackup: async function (backupType) {
      if (this.backupQueued) {
        this.$bvToast.toast(
          'There is a backup currently in progress, please wait...',
          {
            title: "Backup Download",
            toaster: 'b-toaster-bottom-right',
            autoHideDelay: 3000,
            solid: true,
            appendToast: false
          }
        )

        return false
      }

      this.downloadStatus[backupType] = true
      this.lastPercent                = 0
      this.backupQueued               = true
      this.awaitingDownload           = true

      this.$bvToast.toast(
        util.format("Queueing download for '%s'", backupType),
        {
          title: "Backup Download",
          toaster: 'b-toaster-bottom-right',
          autoHideDelay: 3000,
          solid: true,
          appendToast: false
        }
      )

      const response           = await OcculusClient.downloadBackup(backupType)
      const contentDisposition = response.request.getResponseHeader('Content-Disposition')
      const fileName           = contentDisposition.match(/"(.*?)"/)[1]
      const url                = window.URL.createObjectURL(response.data)
      const link               = document.createElement('a')
      link.href                = url
      link.setAttribute('download', fileName)
      document.body.appendChild(link)
      link.click()

      this.downloadStatus[backupType] = false
      this.fileDownloadProgress       = 0
      this.downloadedMbytes           = 0
      this.downloadedBytes            = 0
      this.totalBytes                 = 0
      this.backupQueued               = false

      console.log("Downloading '%s'", fileName)
    },
    handleFileDownloadProgress: function (e) {
      if (this.awaitingDownload) {
        this.awaitingDownload = false
      }

      if (e.detail.downloadedMbytes !== this.downloadedMbytes) {
        this.fileDownloadProgress = e.detail.percent
        this.lastPercent          = e.detail.percent
        this.downloadedMbytes     = Math.round(e.detail.loaded / 1024 / 1024)
        this.downloadedBytes      = e.detail.loaded
        this.totalBytes           = e.detail.total
      }
    },
    ucFirst: function (string) {
      return string.charAt(0).toUpperCase() + string.slice(1)
    }
  }
}
</script>

<style scoped>
.td-center {
  text-align: center
}
</style>
