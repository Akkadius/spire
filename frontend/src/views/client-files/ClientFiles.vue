<template>
  <div class="container-fluid">
    <div class="panel-body">
      <div class="panel panel-default">
        <div class="row">
          <div class="col-6">
            <eq-window-simple title="File Upload Dropzone">
              <div class="mb-3 text-center" style="font-size: 16px">
                <h3 class="mb-3 eq-header" style="font-size: 36px">
                  Files that can be uploaded
                </h3>
                <div v-for="file in ['spells_us.txt', 'dbstr_us.txt']"><b>{{ file }}</b></div>

                <div class="mt-3">
                  <b>Warning</b> Files will immediately overwrite all database values
                </div>

                <div class="mt-3" v-if="message" style="color: limegreen; font-weight: bold">
                  {{ message }}
                </div>
              </div>

              <vue-dropzone
                v-on:vdropzone-success="success"
                v-on:vdropzone-queue-complete="queueComplete"
                v-on:vdropzone-processing="processing"
                ref="myVueDropzone" id="dropzone" :options="dropzoneOptions"
              ></vue-dropzone>


            </eq-window-simple>

          </div>

          <div class="col-6">
            <eq-window-simple title="File Downloads">

              <div class="row">
                <div class="col-12 text-center">
                  <div class="row">
                    <div class="col-12 mt-3">
                      <b-button
                        @click="downloadSpells()"
                        class="ml-3"
                        variant="warning"
                      ><i class="fa fa-download"></i> Download Spells (spells_us.txt)
                      </b-button>
                    </div>
                  </div>
                  <div class="row mt-3">
                    <div class="col-12">
                      <b-button
                        @click="downloadDbStr()"
                        class="ml-3"
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
    </div>
  </div>
</template>

<script>
import vue2Dropzone     from 'vue2-dropzone'
import 'vue2-dropzone/dist/vue2Dropzone.min.css'
import {SpireApiClient} from "../../app/api/spire-api-client";
import EqWindowSimple   from "../../components/eq-ui/EQWindowSimple";

export default {
  name: "ClientFiles.vue",
  components: {
    EqWindowSimple,
    vueDropzone: vue2Dropzone
  },
  data: function () {
    return {
      dropzoneOptions: {
        url: SpireApiClient.getBasePath() + "/api/v1/client-files/import/file",
        resizeWidth: 50,
        resizeHeight: 50,
        thumbnailWidth: 150,
        thumbnailHeight: 150,
        thumbnailMethod: "crop",
        // maxFilesize: 0.5,
        addRemoveLinks: true,

        message: "",
      }
    }
  },
  methods: {

    downloadSpells() {
      window.open(SpireApiClient.getBasePath() + "/api/v1/client-files/export/spells", '_self');
    },
    downloadDbStr() {
      window.open(SpireApiClient.getBasePath() + "/api/v1/client-files/export/dbstr", '_self');
    },

    success(file, response) {
      console.log("file", file)
      console.log("response", response)

      this.message = response

      setTimeout(() => {
        this.$refs.myVueDropzone.removeAllFiles()
      }, 1000)

      setTimeout(() => {
        this.message = ""
      }, 6000)
    },
    queueComplete() {
      console.log("complete")
    },
    processing(event) {
      console.log("dropped")
      console.log(event)
    }
  }
}
</script>spells_us

<style>
.vue-dropzone:hover, .vue-dropzone {
  background-color: rgba(0, 0, 0, 0.1);
}

.dz-message {
  background-color: #161a25;
  color: yellow;
}

.dz-message:hover {
  color: yellow;
}

</style>
