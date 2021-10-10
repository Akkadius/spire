<template>
  <div class="container-fluid">
    <div class="panel-body">
      <div class="panel panel-default">


        <div class="row">
          <div class="col-12">
            <eq-window class="mt-5" title="Zones">

              <app-loader :is-loading="!loaded" padding="8"/>

              <div class="row" style="justify-content: center" v-if="zones">
                <div class='' style="width: 100%" v-if="loaded">
                  <div class="ml-2 mt-1 pb-1">Showing ({{ resultCount }}) results</div>

                  <input type="text"
                         class="form-control"
                         placeholder="Zone filter"
                         v-model="zoneSearchText"
                         v-on:keyup="setStateDebounce">

                  <table id="zonetable" class="eq-table eq-highlight-rows" style="display: table; font-size: 14px; ">
                    <thead>
                    <tr>

                      <th style="width: 60px; white-space: nowrap;"></th>
                      <th style="width: 200px; white-space: nowrap;">Expansion</th>
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
                      <td style="text-align: left">{{ getExpansionName(zone.expansion) }}</td>
                      <td style="text-align: right">{{ zone.short_name }}</td>

                      <td>{{ zone.long_name }}</td>

                      <td style="text-align: center">{{ zone.zoneidnumber }}</td>
                      <td style="text-align: center">{{ zone.version }}</td>
                      <td style="text-align: center">
                        <eq-checkbox :is-checked="(zone.canbind > 0)"/>
                      </td>
                      <td style="text-align: center">
                        <eq-checkbox :is-checked="(zone.cancombat > 0)"/>
                      </td>
                      <td style="text-align: center">
                        <eq-checkbox :is-checked="(zone.canlevitate > 0)"/>
                      </td>
                      <td style="text-align: center">
                        <eq-checkbox :is-checked="(zone.castoutdoor > 0)"/>
                      </td>
                    </tr>
                    </tbody>
                  </table>
                </div>
              </div>

            </eq-window>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script type="ts">
import {ZoneApi} from "@/app/api/api";
import EqWindow from "@/components/eq-ui/EQWindow.vue";
import {SpireApiClient} from "@/app/api/spire-api-client";
import * as util from "util";
import Expansions from "@/app/utility/expansions";
import EqCheckbox from "@/components/eq-ui/EQCheckbox.vue";
import ZoneForm from "@/components/forms/ZoneForm.vue";
import {debounce} from "@/app/utility/debounce.js";
import EqTabs from "@/components/eq-ui/EQTabs.vue";
import EqTab from "@/components/eq-ui/EQTab.vue";

export default {
  components: {
    EqTab,
    EqTabs,
    ZoneForm,
    EqCheckbox,
    EqWindow,
  },
  data() {
    return {
      zones: null,
      filteredZones: null,
      resultCount: 0,
      zoneSearchText: "",
      limit: 10000,
      beginRange: 10000,
      endRange: 100000,

      // route watcher
      routeWatcher: null,

      // loaded state
      loaded: false,
    }
  },
  async activated() {
    await this.init()

    // route watcher
    this.routeWatcher = this.$watch('$route.query', () => {
      this.loadState()
    });

  },
  deactivated() {
    // remove route watcher
    this.routeWatcher()
  },
  methods: {
    async init() {
      this.loaded = false

      await this.listZones()
      this.loadState()
    },

    setStateDebounce: debounce(function () {
      this.setState()
    }, 300),

    // let querystring updates drive state change
    setState() {
      console.log("triggering setState()")

      let query = {}
      if (this.zoneSearchText) {
        query.q = this.zoneSearchText
      }

      this.$router.push(
        {
          path: '/zones',
          query: query
        }
      ).catch(() => {
      })
    },

    // load state from query params
    loadState() {
      console.log(this.$route.query.q)

      if (this.$route.query.q) {
        this.zoneSearchText = this.$route.query.q
      }

      this.triggerSearch()
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
      return Expansions.getExpansionIconUrlSmall(expansion - 1) // zone table is offset by 1
    },
    getExpansionName(expansion) {
      return Expansions.getExpansionName(expansion - 1) // zone table is offset by 1
    },
    clickZoneRow(zone) {
      this.$router.push(
        {
          path: '/zone/' + zone.id
        }
      ).catch(() => {
      })
    },
    listZones: async function () {
      const api = (new ZoneApi(SpireApiClient.getOpenApiConfig()))

      let filters = [
        // ["id", "_lte_", this.endRange],
      ]

      let wheres = [];
      filters.forEach((filter) => {
        const where = util.format("%s%s%s", filter[0], filter[1], filter[2])
        wheres.push(where)
      })

      const result = await api.listZones({limit: this.limit, where: wheres.join("."), orderBy: "expansion.short_name"})

      if (result.status === 200) {
        this.zones = result.data
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
.eq-table tr {
  border-bottom: .4px solid #ffffff1c;
}

.eq-table td {
  padding-top:    5px;
  padding-bottom: 5px;
  border-right: .1px solid #ffffff1c;
  border-left: .1px solid #ffffff1c;
}
</style>
