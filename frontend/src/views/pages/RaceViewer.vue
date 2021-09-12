<template>
  <div>
    <!--    <page-header title="Race Viewer" pre-title="Search and view races..."/>-->

    <!-- CONTENT -->
    <div>
      <div class="container-fluid">


        <eq-window title="Race Viewer" class="mt-5 text-center">

          <div class="row mb-4">
            <div class="col">

              <!-- Form -->

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

          <span v-if="filteredRaces.length === 0">
            No races found...
          </span>

          <div v-if="loaded">
            <div v-for="race in filteredRaces"
                 :key="race"
                 v-lazy-container="{ selector: 'img' }"
                 style="padding-bottom: 15px"
            >
              <h6 class="eq-header"> {{ (raceConstants[race] ? raceConstants[race] : "") }} ({{ race }}) </h6>
              <div class="inner-window mt-3">

            <span v-for="img in raceImages[race]" :key="img">
            <img :src="initialLoad === false ? '' : img" :data-src="img" :id="slug(img)" class="fade-in p-1">

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
              </div>

            </div>
          </div>

          <div class="mt-5">Images courtesy of Maudigan <3</div>
        </eq-window>

        <!--        <div class="card" v-for="race in filteredRaces" :key="race" v-if="loaded">-->
        <!--          <div class="card-header">-->

        <!--            <div v-if="raceConstants[race]">{{ raceConstants[race] }}</div>-->
        <!--            Race ({{ race }})-->
        <!--          </div>-->
        <!--          <div class="card-body" v-lazy-container="{ selector: 'img' }">-->
        <!--          <span v-for="img in raceImages[race]" :key="img">-->
        <!--            <img :data-src="img" :id="slug(img)" class="fade-in">-->

        <!--            &lt;!&ndash; Popover &ndash;&gt;-->
        <!--            <b-popover-->
        <!--              :target="slug(img)"-->
        <!--              placement="bottom"-->
        <!--              variant="light"-->
        <!--              triggers="hover focus"-->
        <!--            >-->
        <!--              <template v-slot:title>Info</template>-->

        <!--              <table>-->
        <!--                <tr>-->
        <!--                  <td><b>Race</b></td>-->
        <!--                  <td>{{ getRaceFromImage(img) }}</td>-->
        <!--                </tr>-->
        <!--                <tr>-->
        <!--                  <td><b>Gender</b></td>-->
        <!--                  <td>{{ getGenderFromImage(img) }}</td>-->
        <!--                </tr>-->
        <!--                <tr>-->
        <!--                  <td><b>Texture</b></td>-->
        <!--                  <td>{{ getTextureFromImage(img) }}</td>-->
        <!--                </tr>-->
        <!--                <tr>-->
        <!--                  <td><b>Helm Texture</b></td>-->
        <!--                  <td>{{ getHelmTextureFromImage(img) }}</td>-->
        <!--                </tr>-->
        <!--              </table>-->
        <!--            </b-popover>-->

        <!--          </span>-->
        <!--          </div>-->

        <!--        </div>-->

      </div>
    </div>

  </div>
</template>

<script>
import NpcModels from "@/app/asset-maps/npc-models-map";
import util from "util";
import slugify from "slugify"
import raceConstants from "@/app/constants/eq-race-constants"
import {RaceViewerStore} from "@/app/store/raceViewerStore";
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
      races: null,
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

        if (!RaceViewerStore.raceExists[raceId]) {
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
            const raceModelImage = util.format("CTN_%s_%s_%s_%s.png", raceId, genderId, textureId, helmTextureId);
            const modelExists    = modelFiles[raceModelImage]

            if (modelExists) {
              raceImages.push(baseUrl + raceModelImage);
            }
          }
        }
      }

      return raceImages
    },
    getMetaDataFromImage: function (img) {
      const pieces   = img.split(/\//);
      const fileName = pieces[pieces.length - 1];

      return fileName.replace("CTN_", "").replace(".png", "").split("_");
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
      if (typeof RaceViewerStore.raceImages !== "undefined" && Object.keys(RaceViewerStore.raceImages).length > 0) {
        this.raceImages    = RaceViewerStore.raceImages
        this.races         = RaceViewerStore.races
        this.filteredRaces = RaceViewerStore.races
        this.loaded        = true
        return
      }

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

        raceImages[raceId] = this.getRaceImages(raceId)
      }

      this.raceImages    = raceImages
      this.races         = races;
      this.filteredRaces = races;
      this.loaded        = true

      // Store
      RaceViewerStore.raceImages = raceImages
      RaceViewerStore.races      = races
      RaceViewerStore.raceExists = raceExists

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

