<template>
  <div style="height: 90vh; overflow-y: scroll;">

    <app-loader :is-loading="!loaded" padding="6"/>

    <div v-if="filteredRaces && filteredRaces.length === 0" class="mt-6">
      No model preview found...
    </div>

    <div v-if="loaded" style="zoom: 115%" class="text-center">
      <div class="mb-3" v-if="filteredRaces && filteredRaces.length > 0">
        Use Arrow Keys <b>Up / Down</b> on material field
      </div>

      <div
        v-for="race in filteredRaces"
        :key="race"
        style="border-radius: 7px; border: 1px solid rgba(255, 255, 255, .5);"
        class="text-center d-inline-block m-2 p-3"
      >
        <div
          v-for="img in raceImages[race]"
          :key="img"
          class="d-inline-block"
          v-if="raceImages[race]"
        >
          <span :class="'race-models-ctn-' + img"></span>
        </div>

        <h6 class="eq-header" style="font-size: 24px">
          {{ (raceConstants[race] ? raceConstants[race] : "") }}
        </h6>

      </div>
    </div>
  </div>
</template>

<script>
import NpcModels         from "@/app/eq-assets/npc-models-map";
import util              from "util";
import {RACES}           from "@/app/constants/eq-race-constants"
import PageHeader        from "@/components/layout/PageHeader";
import {App}             from "@/constants/app";
import EqWindow          from "@/components/eq-ui/EQWindow";
import EqWindowSimple    from "@/components/eq-ui/EQWindowSimple";
import {debounce}        from "@/app/utility/debounce.js";
import {DB_PLAYER_RACES} from "@/app/constants/eq-races-constants";

const baseUrl     = App.ASSET_CDN_BASE_URL + "assets/npc_models/";
const MAX_RACE_ID = 700;
let modelFiles    = {};
let raceExists    = {};
let races         = [];

export default {
  name: "ItemMaterialPreview",
  components: { EqWindowSimple, EqWindow, PageHeader },
  data() {
    return {
      filteredRaces: null,
      raceImages: null,
      loaded: false,
      raceConstants: null,
      raceSearch: "",
    }
  },
  watch: {
    'selectedMaterial': function (newVal, oldVal) {
      this.initModels()
      this.loadModels()
    },
  },
  props: {
    selectedMaterial: {
      required: true,
      type: Number,
    },
  },
  methods: {
    triggerState() {
      this.loadModels()
    },

    isPlayerRace(raceId) {
      if (raceId > 522) {
        return false;
      }

      for (let key in DB_PLAYER_RACES) {
        if (parseInt(raceId) === parseInt(key)) {
          return true;
        }
      }

      return false;
    },

    loadModels() {
      this.loaded = false;

      let filteredRaceIds = [];
      for (let raceId = 0; raceId <= MAX_RACE_ID; raceId++) {

        if (!RACES[raceId]) {
          continue;
        }

        if (!raceExists[raceId]) {
          continue;
        }

        if (!this.isPlayerRace(raceId)) {
          continue;
        }

        if (this.getRaceImages(raceId).length === 0) {
          continue;
        }

        const raceName = RACES[raceId];

        if (this.raceSearch !== "" && !raceName.toLowerCase().includes(this.raceSearch)) {
          continue;
        }

        filteredRaceIds.push(raceId);
      }

      this.filteredRaces = filteredRaceIds
      this.loaded        = true
    },

    triggerStateDebounce: debounce(function () {
      this.triggerState()
    }, 300),
    getRaceImages: function (raceId) {

      let raceImages = [];
      for (let genderId = 0; genderId <= 2; genderId++) {
        for (let textureId = 0; textureId <= 20; textureId++) {

          // we don't care about non-selected material
          if (textureId !== parseInt(this.selectedMaterial)) {
            continue;
          }

          // only show one helm texture for this viewer
          for (let helmTextureId = 0; helmTextureId <= 0; helmTextureId++) {
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
    initModels() {
      var start = new Date().getTime();
      NpcModels[0].contents.forEach((row) => {
        const pieces   = row.name.split(/\//);
        const fileName = pieces[pieces.length - 1];

        modelFiles[fileName] = 1
      })

      this.raceImages = {};
      let raceImages  = {};
      races           = [];

      for (let raceId = 0; raceId <= MAX_RACE_ID; raceId++) {
        let modelKey = "";

        // don't bother crunching non-player race data
        if (!this.isPlayerRace(raceId)) {
          continue;
        }

        for (let genderId = 0; genderId <= 2; genderId++) {
          for (let textureId = 0; textureId <= 20; textureId++) {

            // only show one helm texture for this viewer
            for (let helmTextureId = 0; helmTextureId <= 0; helmTextureId++) {
              modelKey          = util.format("CTN_%s_%s_%s_%s.png", raceId, genderId, textureId, helmTextureId);
              const modelExists = modelFiles[modelKey]

              if (modelExists) {
                if (!raceExists[raceId] && this.getRaceImages(raceId).length > 0) {
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
    }
  },
  async created() {
    this.raceConstants = RACES
    this.initModels()

    setTimeout(() => {
      this.loadModels()
    }, 50);
  }
}
</script>

