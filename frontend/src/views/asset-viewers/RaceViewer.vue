<template>
  <div>
    <eq-window title="Race Viewer">
      <div class="row">
        <div class="col-6">

          Filter by Race Name

          <!-- Input -->
          <input
            type="text"
            class="form-control form-control-prepended list-search mt-1"
            v-model="raceSearch"
            @keyup="zoneSearch = 0; triggerStateDebounce()"
            placeholder="Filter by Race name"
          >

        </div>

        <div class="col-5">
          Find Models Available by Zone

          <select
            @change="raceSearch = ''; triggerState()"
            v-model.number="zoneSearch"
            class="form-control mt-1"
          >
            <option value="0">--- Select Zone ---</option>
            <option
              v-for="(z, index) in zoneList"
              :key="z.zoneId"
              :value="parseInt(z.zoneId)"
            >
              {{ z.shortName }} {{ z.zoneId }}) Races ({{ z.modelCount }}) ({{ z.longName }})
            </option>
          </select>
        </div>

        <div class="col-1">

          <button
            class='btn btn-dark btn-sm mb-1 mr-2 mt-4'
            @click="reset"
          >
            <i class="fa fa-refresh"></i> Reset
          </button>

        </div>

      </div>

    </eq-window>

    <eq-window class="mt-3 p-0 mb-0 pl-1 pr-3">

      <!-- loader -->
      <div v-if="!loaded" class="text-center justify-content-center mt-5 mb-5">
        <div class="mb-3">
          {{ renderingImages ? 'Rendering images...' : 'Loading images...' }}
        </div>
        <loader-fake-progress v-if="!loaded && !renderingImages"/>
        <eq-progress-bar :percent="100" v-if="renderingImages"/>
      </div>

      <div v-if="filteredRaces && filteredRaces.length === 0" class="mt-3 text-center">
        No races found...
      </div>

      <div
        v-if="loaded"
        style="overflow-y: scroll; "
        id="race-viewer-viewport"
        class="row justify-content-center align-items-center text-center fill-screen"
      >
        <div
          v-for="race in filteredRaces"
          :key="race"
          style="padding-bottom: 15px; display: inline-block; border: 2px solid rgba(218, 218, 218, .1); border-radius: 5px; min-height: 200px"
          class="p-3 m-3 fade-in"
        >

          <div class="mt-3" style="vertical-align: middle;">

            <span
              v-for="img in raceImages[race]"
              :title="getImageTitleDescription(img)"
              :key="img"
            >
              <span :class="'race-models-ctn-' + img" style="filter: drop-shadow(10px 5px 7px #000);"></span>
            </span>

            <h6 class="eq-header mt-5"> {{ (raceConstants[race] ? raceConstants[race] : "") }} ({{ race }}) </h6>
          </div>

        </div>

        <div class="col-12 mt-3 text-center">Image Credits @Maudigan</div>
      </div>

    </eq-window>

  </div>
</template>

<script>
import util               from "util";
import {RACES}            from "@/app/constants/eq-race-constants"
import PageHeader         from "@/components/layout/PageHeader";
import {App}              from "@/constants/app";
import EqWindow           from "@/components/eq-ui/EQWindow";
import EqWindowSimple     from "@/components/eq-ui/EQWindowSimple";
import {debounce}         from "@/app/utility/debounce.js";
import {ROUTE}    from "../../routes";
import {SpireApi} from "../../app/api/spire-api";
import {ZoneApi}  from "../../app/api";
import LoaderFakeProgress from "../../components/LoaderFakeProgress";
import EqProgressBar      from "../../components/eq-ui/EQProgressBar";
import EqAssets           from "../../app/eq-assets/eq-assets";
import ContentArea        from "../../components/layout/ContentArea";
import {WindowManager} from "@/app/window";

const baseUrl           = App.ASSET_CDN_BASE_URL + "assets/npc_models/";
const MAX_RACE_ID       = 700;
let modelFiles          = {}
let races               = [];
let zoneToRaceIdMapping = {};

export default {
  components: { ContentArea, EqProgressBar, LoaderFakeProgress, EqWindowSimple, EqWindow, PageHeader },
  data() {
    return {
      filteredRaces: null,
      raceImages: null,
      loaded: false,
      renderingImages: false,
      raceConstants: null,

      // search
      raceSearch: "",
      zoneSearch: 0,

      // data
      zoneList: [],
    }
  },
  methods: {

    // when inputs are triggered and state is updated
    updateQueryState: function () {
      let queryState = {};

      if (this.raceSearch !== "") {
        queryState.raceSearch = this.raceSearch
      }
      if (this.zoneSearch !== 0) {
        queryState.zoneSearch = this.zoneSearch
      }

      this.$router.push(
        {
          path: ROUTE.RACE_VIEWER,
          query: queryState
        }
      ).catch(() => {
      })
    },

    // usually from loading initial state
    loadQueryState: function () {
      if (this.$route.query.raceSearch) {
        this.raceSearch = this.$route.query.raceSearch;
      }
      if (this.$route.query.zoneSearch) {
        this.zoneSearch = this.$route.query.zoneSearch;
      }
    },

    triggerState() {
      this.updateQueryState();
      this.loadModels()

      const target = document.getElementById("race-viewer-viewport")
      if (target) {
        target.scrollTop = 0
      }
    },

    loadModels() {
      this.loaded = false;

      let curImg    = new Image();
      curImg.src    = '/eq-asset-preview-master/assets/sprites/race-models.png';
      curImg.onload = () => {
        this.renderingImages = true

        setTimeout(() => {
          this.renderingImages = false

          // filtering
          // race filter
          if (this.raceSearch !== "") {
            let filteredRaceIds = [];
            for (let raceId = 0; raceId <= MAX_RACE_ID; raceId++) {

              if (!RACES[raceId] || !modelFiles[raceId]) {
                continue;
              }

              const raceName = RACES[raceId];
              if (!raceName.toLowerCase().includes(this.raceSearch.toLowerCase())) {
                continue;
              }

              filteredRaceIds.push(raceId);
            }

            this.filteredRaces = filteredRaceIds
            this.loaded        = true
            return;
          }

          // zone filter
          if (this.zoneSearch !== 0) {
            let filteredRaceIds = [];
            for (let raceId = 0; raceId <= MAX_RACE_ID; raceId++) {
              if (!RACES[raceId] || !modelFiles[raceId]) {
                continue;
              }

              if (zoneToRaceIdMapping[this.zoneSearch]) {
                if (!zoneToRaceIdMapping[this.zoneSearch].includes(raceId)) {
                  continue;
                }
              }

              filteredRaceIds.push(raceId);
            }

            this.filteredRaces = filteredRaceIds
            this.loaded        = true
            return;
          }

          // set filtered races
          this.filteredRaces = races
          this.loaded        = true;

          setTimeout(() => {
            WindowManager.resizeFillScreenElements();
          }, 100);

        }, 1);
      }
    },

    triggerStateDebounce: debounce(function () {
      this.triggerState()
    }, 500),
    getRaceImages: function (raceId) {
      let raceImages = []
      modelFiles[raceId].forEach((file) => {
        if (file.includes(util.format("CTN_%s", raceId))) {

          // replace for css formatting
          file = file.replace(".png", "")
            .replace("CTN_", "")
            .replaceAll("_", "-")

          // images
          raceImages.push(file)
        }
      })

      return raceImages
    },
    getMetaDataFromImage: function (img) {
      const pieces   = img.split(/\//);
      const fileName = pieces[pieces.length - 1];

      return fileName.split("-");
    },
    getImageTitleDescription(img) {
      const meta = this.getMetaDataFromImage(img);

      return util.format(
        "Race: %s Gender: %s Texture: %s Helm: %s",
        meta[0],
        meta[1],
        meta[2],
        meta[3]
      )
    },
    async initModels() {
      console.time('initModels');
      console.time('modelFiles');

      modelFiles = {}

      const r = await EqAssets.getNpcModels()
      r.forEach((n) => {
        if (typeof modelFiles[n.raceId] === "undefined") {
          modelFiles[n.raceId] = []
        }
        modelFiles[n.raceId].push(n.fileName)
      })

      // console.log(modelFiles)

      console.timeEnd('modelFiles');

      this.raceImages = {};
      let raceImages  = {};
      races           = [];

      console.time('defineRaces');
      for (let raceId = 0; raceId <= MAX_RACE_ID; raceId++) {
        if (modelFiles[raceId] && modelFiles[raceId].length > 0) {
          races.push(raceId)
          raceImages[raceId] = this.getRaceImages(raceId)
        }
      }
      console.timeEnd('defineRaces');

      console.timeEnd('initModels');

      this.raceImages    = raceImages
      this.filteredRaces = races;
    },
    reset() {
      this.raceSearch = ""
      this.zoneSearch = 0
      this.updateQueryState()
      this.loadModels()
    },
    async loadRaceInventory() {
      const result = await SpireApi.v1().get('/static-map/race-inventory-map.json')

      // zero out
      zoneToRaceIdMapping = {};
      result.data.races.forEach((race) => {
        if (race.sources) {
          race.sources.forEach((source) => {
            if (source.zones) {
              source.zones.forEach((zone) => {
                if (typeof zoneToRaceIdMapping[zone.id] === "undefined") {
                  zoneToRaceIdMapping[zone.id] = []
                }

                zoneToRaceIdMapping[zone.id].push(race.race_id)
              })
            }
          })
        }
      })

      // after we load race inventory data
      let zoneList   = [];
      let zoneResult = await (new ZoneApi(...SpireApi.cfg())).listZones({
        where: "version__0",
        orderBy: "short_name",
        groupBy: "zoneidnumber"
      })

      // zone data
      if (zoneResult.status === 200) {
        zoneResult.data.forEach((row) => {
          zoneList.push(
            {
              zoneId: row.zoneidnumber,
              shortName: row.short_name,
              longName: row.long_name,
              modelCount: zoneToRaceIdMapping[row.zoneidnumber] ? zoneToRaceIdMapping[row.zoneidnumber].length : 0,
            }
          )
        })

        this.zoneList = zoneList
      }
    }
  },
  async mounted() {
    this.loadQueryState()
    this.raceConstants = RACES
    await this.initModels()
    await this.loadRaceInventory()
    this.loadModels()
  }
}
</script>

