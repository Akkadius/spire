<template>
  <div>
    <div class="container-fluid">
      <eq-window title="Race Viewer" class="mt-5 text-center">
        <div class="row mb-4">
          <div class="col">

            <!-- Input -->
            <input
              type="text"
              class="form-control form-control-prepended list-search"
              v-model.lazy="raceSearch"
              @keyup.enter="doRaceSearch()"
              placeholder="Filter by Race name">
          </div>
          <div class="col-auto">

          </div>
        </div>

        <app-loader :is-loading="!loaded" padding="6"/>

        <span v-if="filteredRaces && filteredRaces.length === 0">
            No races found...
          </span>

        <div v-if="loaded">
          <div v-for="race in filteredRaces"
               :key="race"
               style="padding-bottom: 15px; display: inline-block; border: 1px solid rgb(218, 218, 218); border-radius: 7px;"
               class="p-1 m-3"
          >

            <div class="mt-3" style="vertical-align: middle;">

            <span v-for="img in raceImages[race]" :key="img">
              <span :class="'race-models-ctn-' + img" :id="slug(img)"></span>

              <!-- Popover -->
              <b-popover
                :target="slug(img)"
                placement="bottom"
                variant="light"
                triggers="hover focus"
              >
                <template v-slot:title>Info</template>

                <table>
                  <tr>
                    <td><b>Race</b></td>
                    <td>{{ getRaceFromImage(img) }}</td>
                  </tr>
                  <tr>
                    <td><b>Gender</b></td>
                    <td>{{ getGenderFromImage(img) }}</td>
                  </tr>
                  <tr>
                    <td><b>Texture</b></td>
                    <td>{{ getTextureFromImage(img) }}</td>
                  </tr>
                  <tr>
                    <td><b>Helm Texture</b></td>
                    <td>{{ getHelmTextureFromImage(img) }}</td>
                  </tr>
                </table>
              </b-popover>

            </span>

              <h6 class="eq-header"> {{ (raceConstants[race] ? raceConstants[race] : "") }} ({{ race }}) </h6>
            </div>

          </div>
        </div>

        <div class="mt-5">Images courtesy of Maudigan <3</div>
      </eq-window>
    </div>

  </div>
</template>

<script>
import NpcModels from "@/app/asset-maps/npc-models-map";
import util from "util";
import slugify from "slugify"
import raceConstants from "@/app/constants/eq-race-constants"
import PageHeader from "@/views/layout/PageHeader";
import {App} from "@/constants/app";
import EqWindow from "@/components/eq-ui/EQWindow";
import EqWindowSimple from "@/components/eq-ui/EQWindowSimple";

const baseUrl = App.ASSET_CDN_BASE_URL + "assets/npc_models/";

const MAX_RACE_ID = 700;
let modelFiles    = {};
let raceExists    = {};

export default {
  components: { EqWindowSimple, EqWindow, PageHeader },
  data() {
    return {
      filteredRaces: null,
      raceImages: null,
      loaded: false,
      raceConstants: null,
      raceSearch: "",
      initialLoad: false
    }
  },
  methods: {
    doRaceSearch: function () {
      this.loaded = false

      let filteredRaceIds = [];
      for (let raceId = 0; raceId <= MAX_RACE_ID; raceId++) {

        if (!raceConstants[raceId]) {
          continue;
        }

        if (!raceExists[raceId]) {
          continue;
        }

        const raceName = raceConstants[raceId];
        if (!raceName.toLowerCase().includes(this.raceSearch)) {
          continue;
        }

        filteredRaceIds.push(raceId);
      }

      this.filteredRaces = filteredRaceIds
      this.loaded        = true
    },
    getRaceImages: function (raceId) {

      let raceImages = [];
      for (let genderId = 0; genderId <= 2; genderId++) {
        for (let textureId = 0; textureId <= 20; textureId++) {
          for (let helmTextureId = 0; helmTextureId <= 10; helmTextureId++) {
            const raceModel   = util.format("CTN_%s_%s_%s_%s.png", raceId, genderId, textureId, helmTextureId)
            const modelExists = modelFiles[raceModel]

            if (modelExists) {
              raceImages.push(util.format("%s-%s-%s-%s", raceId, genderId, textureId, helmTextureId));
            }
          }
        }
      }

      return raceImages
    },
    getMetaDataFromImage: function (img) {
      const pieces   = img.split(/\//);
      const fileName = pieces[pieces.length - 1];

      return fileName.split("-");
    },
    getRaceFromImage: function (img) {
      return this.getMetaDataFromImage(img)[0];
    },
    getGenderFromImage: function (img) {
      return this.getMetaDataFromImage(img)[1];
    },
    getTextureFromImage: function (img) {
      return this.getMetaDataFromImage(img)[2];
    },
    getHelmTextureFromImage: function (img) {
      return this.getMetaDataFromImage(img)[3];
    },
    slug: function (toSlug) {
      return slugify(toSlug.replace(/[&\/\\#, +()$~%.'":*?<>{}]/g, "-"))
    },
    loadModels: function () {
      var start = new Date().getTime();
      NpcModels[0].contents.forEach((row) => {
        const pieces   = row.name.split(/\//);
        const fileName = pieces[pieces.length - 1];

        modelFiles[fileName] = 1
      })

      this.raceImages = {};
      let races       = [];
      let raceImages  = {};

      for (let raceId = 0; raceId <= MAX_RACE_ID; raceId++) {
        let modelKey = "";

        for (let genderId = 0; genderId <= 2; genderId++) {
          for (let textureId = 0; textureId <= 20; textureId++) {
            for (let helmTextureId = 0; helmTextureId <= 10; helmTextureId++) {
              modelKey          = util.format("CTN_%s_%s_%s_%s.png", raceId, genderId, textureId, helmTextureId);
              const modelExists = modelFiles[modelKey]

              if (modelExists) {
                if (!raceExists[raceId]) {
                  raceExists[raceId] = 1
                  races.push(raceId);
                }
              }
            }
          }
        }

        if (raceExists[raceId]) {
          raceImages[raceId] = this.getRaceImages(raceId)
        }
      }

      this.raceImages    = raceImages
      this.filteredRaces = races;
      this.loaded        = true
    }
  },
  async mounted() {
    this.raceConstants = raceConstants

    setTimeout(() => {
      this.loadModels()
      this.initialLoad = true
    }, 100);
  }
}
</script>

