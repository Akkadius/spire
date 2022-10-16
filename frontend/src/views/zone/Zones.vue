<template>
  <div>
    <div class="row">
      <div class="col-12">
        <eq-window-simple title="Zones">

          <div class="row" style="justify-content: center">
            <div class='col-4'>
              <div>Showing ({{ resultCount }}) results</div>
              <input
                type="text"
                class="form-control"
                placeholder="Search by zone long or short name, expansion etc."
                v-model="zoneSearchText"
                v-on:keyup="selectedExpansion = -1; setStateDebounce();"
                v-on:keyup.enter="updateQueryState"
              >
            </div>
            <div class="col-8">

              <img
                v-for="(expansion, expansionId) in EXPANSIONS_FULL"
                v-if="!getExpansionIcon(expansionId).includes('base64')"
                :title="getExpansionName(expansionId) + ' (' + expansionId + ')'"
                :src="getExpansionIcon(expansionId)"
                @click="zoneSearchText = ''; selectedExpansion = expansionId; updateQueryState()"
                :style="'width: 56px; opacity: .7; ' + (isExpansionSelected(expansionId) ? 'border: 2px solid #dadada; border-radius: 7px;' : 'border: 2px solid rgb(218 218 218 / 30%); border-radius: 7px;')"
                class="mr-2 p-1 hover-highlight-inner"
              >

            </div>
          </div>

        </eq-window-simple>

        <eq-window
          v-if="zones"
          class="p-3 pt-0"
          style="height: 80vh; overflow-y: scroll; overflow-x: hidden"
        >
          <table
            id="zonetable"
            class="eq-table eq-highlight-rows"
            style="display: table; font-size: 14px;"
          >
            <thead>
            <tr>

              <th style="width: 60px; white-space: nowrap;"></th>
              <th style="width: 200px; text-align: center; white-space: nowrap;">Expansion</th>
              <th style="width: 100px; white-space: nowrap;">Short Name</th>

              <th style="width: 350px">Long Name</th>

              <th style="width: 50px">Zone ID</th>
              <th style="width: 50px">Version</th>
              <th style="text-align: center">Bind</th>
              <th style="text-align: center">Combat</th>
              <th style="text-align: center">Levitate</th>
              <th style="text-align: center">Outdoor</th>
            </tr>
            </thead>
            <tbody>
            <tr v-for="(zone, index) in filteredZones" :key="zone.id" @click="clickZoneRow(zone)">

              <td style="text-align: center"><img :src="getExpansionIcon(zone.expansion)"></td>
              <td style="text-align: center">{{ getExpansionName(zone.expansion) }}</td>
              <td style="text-align: right">{{ zone.short_name }}</td>

              <td>{{ zone.long_name }}</td>

              <td style="text-align: center">{{ zone.zoneidnumber }}</td>
              <td style="text-align: center">{{ zone.version }}</td>
              <td style="text-align: center">
                <eq-checkbox
                  :disabled="true"
                  :value="zone.canbind"
                />
              </td>
              <td style="text-align: center">
                <eq-checkbox
                  :disabled="true"
                  :value="zone.cancombat"
                />
              </td>
              <td style="text-align: center">
                <eq-checkbox
                  :disabled="true"
                  :value="zone.canlevitate"
                />
              </td>
              <td style="text-align: center">
                <eq-checkbox
                  :disabled="true"
                  :value="zone.castoutdoor"
                />
              </td>
            </tr>
            </tbody>
          </table>

        </eq-window>
      </div>
    </div>
  </div>
</template>

<script type="ts">
import {ZoneApi} from "@/app/api/api";
import EqWindow from "@/components/eq-ui/EQWindow.vue";
import {SpireApi} from "../../app/api/spire-api";
import Expansions from "@/app/utility/expansions";
import EqCheckbox from "@/components/eq-ui/EQCheckbox.vue";
import {debounce} from "@/app/utility/debounce.js";
import EqTabs from "@/components/eq-ui/EQTabs.vue";
import EqTab from "@/components/eq-ui/EQTab.vue";
import ContentArea from "@/components/layout/ContentArea.vue";
import {SpireQueryBuilder} from "@/app/api/spire-query-builder";
import EqWindowSimple from "@/components/eq-ui/EQWindowSimple.vue";
import {EXPANSIONS_FULL} from "@/app/constants/eq-expansions";
import {ROUTE} from "@/routes";

export default {
  components: {
    EqWindowSimple,
    ContentArea,
    EqTab,
    EqTabs,
    EqCheckbox,
    EqWindow,
  },
  data() {
    return {
      zones: null,
      filteredZones: null,
      resultCount: 0,

      selectedExpansion: -1,
      zoneSearchText: "",

      // route watcher
      routeWatcher: null,

      // loaded state
      loaded: false,

      EXPANSIONS_FULL: EXPANSIONS_FULL,
    }
  },
  watch: {
    $route(to, from) {
      this.loadQueryState()
      this.init()
    }
  },
  async mounted() {
    this.loadQueryState()
    await this.init()
  },
  methods: {
    isExpansionSelected(expansion) {
      return expansion === this.selectedExpansion
    },

    async init() {
      this.loaded = false
      this.loadQueryState()
      await this.listZones()
      this.triggerSearch()
    },

    setStateDebounce: debounce(function () {
      this.updateQueryState()
    }, 300),


    // when inputs are triggered and state is updated
    updateQueryState: function () {
      let queryState = {};

      if (this.zoneSearchText !== "") {
        queryState.q = this.zoneSearchText
      }
      if (parseInt(this.selectedExpansion) !== -1) {
        queryState.expansion = this.selectedExpansion
      }

      this.$router.push(
        {
          path: ROUTE.ZONES,
          query: queryState
        }
      ).catch(() => {
      })
    },

    // usually from loading initial state
    loadQueryState: function () {
      if (this.$route.query.expansion) {
        this.selectedExpansion = this.$route.query.expansion;
      }
      if (this.$route.query.q) {
        this.zoneSearchText = this.$route.query.q
      }
    },

    triggerSearch: function () {

      this.filteredZones = this.zones.filter((e) => {
        const searchString = this.zoneSearchText.toLowerCase()

        // console.log(searchString)
        // console.log(e.short_name.toLowerCase())
        const expansion = e.expansion - 1 // zone table is offset by 1

        return e.short_name.toLowerCase().includes(searchString)
          || Expansions.getExpansionName(expansion).toLowerCase().includes(searchString)
          || e.long_name.toLowerCase().includes(searchString)
      });

      if (this.filteredZones.length === 0) {
        this.filteredZones = this.zones
      }

      this.resultCount = this.filteredZones.length
    },
    getExpansionIcon(expansion) {
      return Expansions.getExpansionIconUrlSmall(expansion)
    },
    getExpansionName(expansion) {
      return Expansions.getExpansionName(expansion)
    },
    clickZoneRow(zone) {
      this.$router.push(
        {
          path: '/zone/' + zone.short_name + '?v=' + zone.version
        }
      ).catch(() => {
      })
    },
    listZones: async function () {

      const builder   = (new SpireQueryBuilder())
      const expansion = parseInt(this.selectedExpansion)
      console.log(expansion)
      if (expansion > -1) {
        builder.where("expansion", "=", expansion)
      }

      builder.orderBy(["expansion", "long_name"])
        .limit(10000)

      const result = await (new ZoneApi(...SpireApi.cfg())).listZones(builder.get())
      if (result.status === 200) {
        this.zones         = result.data
        this.filteredZones = result.data
      }

      this.loaded = true
    }
  }
}

</script>

<style>
#zonetable TBODY TR TD {
  padding: 2px 4px;
}

#zonetable tr {
  border-bottom: .4px solid #ffffff1c;
}

#zonetable td {
  padding-top: 5px;
  padding-bottom: 5px;
  border-right: .1px solid #ffffff1c;
  border-left: .1px solid #ffffff1c;
}
</style>
