<template>
  <content-area class="text-center" style="padding: 0 !important">
    <eq-window-simple
      title="Client File Drop Zone"
      style="height: 95vh"
      class="mt-3 p-0"
    >
      <div
        class="mb-3"
        style="font-size: 16px"
      >

        <!-- Success -->
        <div
          class="mt-3 eq-header fade-in"
          v-if="successMessage"
          style="font-size: 36px"
        >
          {{ successMessage }}
        </div>
        <div class="mt-3 fade-in" v-if="loading">
          <loader-fake-progress/>
        </div>
      </div>

      <!-- Buttons -->
      <div class="row">
        <div class="col-12">
          <b-button @click="downloadSpells" size="sm" variant="warning">
            <i class="fa fa-cloud-download"></i>
            Spells (spells_us.txt)
          </b-button>
          <b-button @click="downloadDbStr" size="sm" variant="warning" class="ml-3">
            <i class="fa fa-cloud-download"></i> DB Strings (dbstr_us.txt)
          </b-button>
        </div>
      </div>

      <!-- Dropzone -->
      <vue-dropzone
        class="mt-4"
        style="height: 70vh"
        v-on:vdropzone-success="success"
        v-on:vdropzone-queue-complete="queueComplete"
        v-on:vdropzone-processing="processing"
        ref="myVueDropzone"
        id="dropzone"
        :options="dropzoneOptions"
      />

      <div class="mt-4" style="color: red">
        <b>WARNING</b> Files will immediately overwrite all database values
      </div>


    </eq-window-simple>

  </content-area>
</template>

<script>
import querystring        from 'querystring'
import vue2Dropzone       from 'vue2-dropzone'
import 'vue2-dropzone/dist/vue2Dropzone.min.css'
import {SpireApi}         from "../../app/api/spire-api";
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
          SpireApi.getBasePath(),
          querystring.stringify(SpireApi.getAccessTokenQueryString())
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
          SpireApi.getBasePath(),
          querystring.stringify(SpireApi.getAccessTokenQueryString())
        ),
        '_blank'
      );
    },
    downloadDbStr() {
      window.open(
        util.format(
          "%s/api/v1/client-files/export/dbstr?%s",
          SpireApi.getBasePath(),
          querystring.stringify(SpireApi.getAccessTokenQueryString())
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

.dz-message {
  height: 80vh;
}

</style>
