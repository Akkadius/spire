<template>
  <div>
    <div class="container-fluid">
      <eq-window title="Race Viewer" class="mt-5 text-center">
        <div class="row mb-4">
          <div class="col-6">

            <!-- Input -->
            <input
              type="text"
              class="form-control form-control-prepended list-search"
              v-model="raceSearch"
              @keyup="zoneSearch = 0; triggerStateDebounce()"
              @keyup.enter="zoneSearch = 0; triggerState()"
              placeholder="Filter by Race name"
            >

          </div>

          <div class="col-5">
            <select
              @change="raceSearch = ''; triggerState()"
              v-model.number="zoneSearch" class="form-control"
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
              class='btn btn-outline-warning btn-sm mb-1 mr-2'
              @click="reset"
            >
              <i class="fa fa-refresh"></i> Reset
            </button>

          </div>

        </div>

        <app-loader :is-loading="!loaded" padding="6"/>

        <span v-if="filteredRaces && filteredRaces.length === 0">
            No races found...
          </span>

        <div v-if="loaded" class="row justify-content-center align-items-center">
          <div
            v-for="race in filteredRaces"
            :key="race"
            style="padding-bottom: 15px; display: inline-block; border: 3px solid rgba(218, 218, 218, .1); border-radius: 7px; min-height: 200px"
            class="p-3 m-3 fade-in"
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

              <h6 class="eq-header mt-5"> {{ (raceConstants[race] ? raceConstants[race] : "") }} ({{ race }}) </h6>
            </div>

          </div>
        </div>

        <div class="mt-5">Images courtesy of Maudigan <3</div>
      </eq-window>
    </div>

  </div>
</template>

<script>
import NpcModels        from "@/app/eq-assets/npc-models-map";
import util             from "util";
import slugify          from "slugify"
import {RACES}          from "@/app/constants/eq-race-constants"
import PageHeader       from "@/components/layout/PageHeader";
import {App}            from "@/constants/app";
import EqWindow         from "@/components/eq-ui/EQWindow";
import EqWindowSimple   from "@/components/eq-ui/EQWindowSimple";
import {debounce}       from "@/app/utility/debounce.js";
import {ROUTE}          from "../../routes";
import {SpireApiClient} from "../../app/api/spire-api-client";
import {ZoneApi}        from "../../app/api";

const baseUrl           = App.ASSET_CDN_BASE_URL + "assets/npc_models/";
const MAX_RACE_ID       = 700;
let modelFileExists     = {};
let modelFiles          = {}
let races               = [];
let zoneToRaceIdMapping = {};


export default {
  components: { EqWindowSimple, EqWindow, PageHeader },
  data() {
    return {
      filteredRaces: null,
      raceImages: null,
      loaded: false,
      raceConstants: null,

      raceSearch: "",
      zoneSearch: 0,

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
    },

    loadModels() {
      this.loaded = false;

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
    },

    triggerStateDebounce: debounce(function () {
      this.triggerState()
    }, 1000),
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
    initModels() {
      console.time('initModels');
      console.time('modelFiles');

      modelFiles = {}
      NpcModels[0].contents.forEach((row) => {
        const pieces     = row.name.split(/\//);
        const fileName   = pieces[pieces.length - 1];
        const paramSplit = fileName.split("_")
        const raceId     = paramSplit[1].trim();

        modelFileExists[fileName] = 1

        if (typeof modelFiles[raceId] === "undefined") {
          modelFiles[raceId] = []
        }
        modelFiles[raceId].push(fileName)
      })

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
    loadRaceInventory() {
      SpireApiClient.v1().get('/static-map/race-inventory-map.json').then((result) => {

        zoneToRaceIdMapping = {};

        result.data.races.forEach((race) => {
          // console.log(race)

          if (race.sources) {
            race.sources.forEach((source) => {
              if (source.zones) {
                source.zones.forEach((zone) => {
                  // console.log(zone)

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
        let zoneList = [];
        (new ZoneApi(SpireApiClient.getOpenApiConfig())).listZones({
          where: "version__0",
          orderBy: "short_name",
          groupBy: "zoneidnumber"
        }).then((result) => {
          if (result.status === 200) {
            result.data.forEach((row) => {

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
        })

        // console.log(zoneToRaceIdMapping)
        // console.log(result)
      });

    }
  },
  async mounted() {
    this.loadQueryState()
    this.raceConstants = RACES
    this.initModels()
    this.loadRaceInventory()

    setTimeout(() => {
      this.loadModels()
    }, 50);
  }
}
</script>

