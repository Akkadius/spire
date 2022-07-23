<template>
  <div>
    <div class="row">
      <div class="col-6">
        <eq-window-simple title="Client File Dropzone">
          <div class="mb-3 text-center" style="font-size: 16px">
            <!--                <h3 class="mb-3 eq-header" style="font-size: 42px">-->
            <!--                  Client File DropZone Upload-->
            <!--                </h3>-->
            <div v-for="file in ['dbstr_us.txt', 'spells_us.txt']"><b>{{ file }}</b></div>

            <div class="mt-3">
              <b>Warning</b> Files will immediately overwrite all database values
            </div>

            <!-- Success -->
            <div class="mt-3 eq-header fade-in" v-if="successMessage" style="font-size: 36px">
              {{ successMessage }}
            </div>
            <div class="mt-3 fade-in" v-if="loading">
              <loader-fake-progress/>
            </div>
          </div>

          <vue-dropzone
            class="mt-4"
            v-on:vdropzone-success="success"
            v-on:vdropzone-queue-complete="queueComplete"
            v-on:vdropzone-processing="processing"
            ref="myVueDropzone" id="dropzone" :options="dropzoneOptions"
          ></vue-dropzone>


        </eq-window-simple>

      </div>

      <div class="col-3">
        <eq-window-simple title="File Downloads" class="p-3">

          <div class="row">
            <div class="col-12">
              <div class="row">
                <div class="col-12 mt-3">
                  <b-button
                    @click="downloadSpells()"
                    class="form-control"
                    size="sm"
                    variant="warning"
                  ><i class="fa fa-download"></i> Download Spells (spells_us.txt)
                  </b-button>
                </div>
              </div>
              <div class="row mt-3">
                <div class="col-12">
                  <b-button
                    @click="downloadDbStr()"
                    class="form-control"
                    size="sm"
                    variant="warning"
                  ><i class="fa fa-download"></i> Download Database Strings (dbstr_us.txt)
                  </b-button>
                </div>
              </div>
            </div>
          </div>

        </eq-window-simple>
      </div>

    </div>

  </div>
</template>

<script>
import querystring        from 'querystring'
import vue2Dropzone       from 'vue2-dropzone'
import 'vue2-dropzone/dist/vue2Dropzone.min.css'
import {SpireApiClient}   from "../../app/api/spire-api-client";
import EqWindowSimple     from "../../components/eq-ui/EQWindowSimple";
import EqWindow           from "../../components/eq-ui/EQWindow";
import LoaderFakeProgress from "../../components/LoaderFakeProgress";
import util               from "util";
import ContentArea        from "../../components/layout/ContentArea";

export default {
  name: "ClientFiles.vue",
  components: {
    ContentArea,
    LoaderFakeProgress,
    EqWindow,
    EqWindowSimple,
    vueDropzone: vue2Dropzone
  },
  data: function () {
    return {
      successMessage: "",
      loading: false,

      dropzoneOptions: {
        url: util.format(
          "%s/api/v1/client-files/import/file?%s",
          SpireApiClient.getBasePath(),
          querystring.stringify(SpireApiClient.getAccessTokenQueryString())
        ),
        resizeWidth: 50,
        resizeHeight: 50,
        thumbnailWidth: 150,
        thumbnailHeight: 150,
        thumbnailMethod: "crop",
        // maxFilesize: 0.5,
        addRemoveLinks: true,
      }
    }
  },
  methods: {

    downloadSpells() {
      window.open(
        util.format(
          "%s/api/v1/client-files/export/spells?%s",
          SpireApiClient.getBasePath(),
          querystring.stringify(SpireApiClient.getAccessTokenQueryString())
        ),
        '_blank'
      );
    },
    downloadDbStr() {
      window.open(
        util.format(
          "%s/api/v1/client-files/export/dbstr?%s",
          SpireApiClient.getBasePath(),
          querystring.stringify(SpireApiClient.getAccessTokenQueryString())
        ),
        '_blank'
      );
    },

    success(file, response) {
      this.loading = false
      console.log("file", file)
      console.log("response", response)

      this.successMessage = response

      setTimeout(() => {
        this.$refs.myVueDropzone.removeAllFiles()
      }, 1000)

      setTimeout(() => {
        this.successMessage = ""
      }, 6000)
    },
    queueComplete() {
      console.log("complete")
    },
    processing(event) {
      this.loading = true
      console.log("dropped")
      console.log(event)
    }
  }
}
</script>

<style>
.vue-dropzone:hover, .vue-dropzone {
  background-color: rgba(0, 0, 0, 0.6);
}

.dz-message {
  font-size: 14px;
  background-color: #161a25;
  color: yellow;
}

.dz-message:hover {
  color: yellow;
}

</style>
