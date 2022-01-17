<template>
  <div>
    <div class="container-fluid">
      <eq-window title="Race Viewer" class="mt-5 text-center">
        <div class="row mb-4">
          <div class="col-8">

            <!-- Input -->
            <input
              type="text"
              class="form-control form-control-prepended list-search"
              v-model="raceSearch"
              @keyup="triggerStateDebounce()"
              @keyup.enter="triggerState()"
              placeholder="Filter by Race name"
            >

          </div>

          <div class="col-4">
            <select
              v-model.number="zoneSearch" class="form-control"
            >
              <option
                v-for="(z, index) in zoneList"
                :key="z.zoneId"
                :value="parseInt(z.zoneId)"
              >
                {{ z.shortName }} {{ z.zoneId }}) {{ z.longName }})
              </option>
            </select>
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
            style="padding-bottom: 15px; display: inline-block; border: 3px solid rgba(218, 218, 218, .1); border-radius: 7px;"
            class="p-3 m-3"
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

const baseUrl     = App.ASSET_CDN_BASE_URL + "assets/npc_models/";
const MAX_RACE_ID = 700;
let modelFiles    = {};
let raceExists    = {};
let races         = [];

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
    },

    triggerState() {
      this.updateQueryState();
      this.loadModels()
    },

    loadModels() {
      this.loaded = false;

      if (this.raceSearch !== "") {
        let filteredRaceIds = [];
        for (let raceId = 0; raceId <= MAX_RACE_ID; raceId++) {

          if (!RACES[raceId]) {
            continue;
          }

          if (!raceExists[raceId]) {
            continue;
          }

          const raceName = RACES[raceId];
          if (!raceName.toLowerCase().includes(this.raceSearch)) {
            continue;
          }

          filteredRaceIds.push(raceId);
        }

        this.filteredRaces = filteredRaceIds
        this.loaded        = true
        return;
      }

      this.filteredRaces = races
      this.loaded        = true;
    },

    triggerStateDebounce: debounce(function () {
      this.triggerState()
    }, 1000),
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
    },
    loadRaceInventory() {
      SpireApiClient.v1().get('/static-map/race-inventory-map.json').then((result) => {

        let zoneToRaceIdMapping = {};

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

        // console.log(zoneToRaceIdMapping)

        // console.log(result)
      });

      let zoneList = [];

      (new ZoneApi(SpireApiClient.getOpenApiConfig())).listZones({
        where: "version__0",
        orderBy: "zoneidnumber",
        groupBy: "zoneidnumber"
      }).then((result) => {
        if (result.status === 200) {
          result.data.forEach((row) => {
            // console.log(row)
            zoneList.push(
              {
                zoneId: row.zoneidnumber,
                shortName: row.short_name,
                longName: row.long_name,
              }
            )
          })

          this.zoneList = zoneList
        }
      })


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

