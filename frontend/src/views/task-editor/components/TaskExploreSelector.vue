<template>
  <div>
    <eq-window-simple title="Explore Proximity Selector">
      <b-input
        v-model="zoneSearch"
        id="search-selector"
        class="form-control"
        v-on:keyup="searchZone"
        placeholder="Search by zone name, id..."
      />
    </eq-window-simple>

    <eq-window-simple
      id="explore-container"
      style="height: 85vh; overflow-y: scroll;" class="p-0"
    >
      <table
        id="explore-table"
        class="eq-table eq-highlight-rows"
        style="display: table; font-size: 14px; overflow-x: scroll"
      >
        <thead
          class="eq-table-floating-header"
        >
        <tr>
          <th></th>
          <th style="width: 30px">ID</th>
          <th>Zone</th>
          <th>Min X</th>
          <th>Max X</th>
          <th>Min Y</th>
          <th>Max Y</th>
          <th>Min Z</th>
          <th>Max Z</th>
        </tr>
        </thead>
        <tbody>
        <tr
          :id="'proximity-' + e.exploreid"
          :class="(isZoneSelected(e) ? 'pulsate-highlight-white' : '')"
          v-for="(e, index) in filteredExplore"
          :key="e.id"
        >
          <td>
            <b-button
              class="btn-dark btn-sm btn-outline-warning"
              @click="selectZone(e)"
            >
              Select
            </b-button>
          </td>
          <td style="text-align: center" class="p-0">{{ e.exploreid }}</td>
          <td><span v-if="zones[e.zoneid]">{{ zones[e.zoneid].long_name }} ({{e.zoneid}})</span></td>
          <td>{{e.minx}}</td>
          <td>{{e.maxx}}</td>
          <td>{{e.miny}}</td>
          <td>{{e.maxy}}</td>
          <td>{{e.minz}}</td>
          <td>{{e.maxz}}</td>
        </tr>
        </tbody>
      </table>
    </eq-window-simple>
  </div>
</template>

<script>
import EqWindowSimple      from "@/components/eq-ui/EQWindowSimple";
import {ProximityApi}      from "@/app/api";
import {SpireApiClient}    from "@/app/api/spire-api-client";
import util                from "util";
import Expansions          from "@/app/utility/expansions";
import EqCheckbox          from "@/components/eq-ui/EQCheckbox";
import {SpireQueryBuilder} from "@/app/api/spire-query-builder";
import {Zones}             from "@/app/zones";

let zones = {}

export default {
  name: "TaskExploreSelector",
  components: { EqCheckbox, EqWindowSimple },
  data() {
    return {
      // filtered content
      filteredExplore: {},

      // search
      zoneSearch: "",

      // model we work with after the prop is passed so we can manipulate it ourselves
      // props should not be mutated
      selectedExplore: "",

      zones: {}
    }
  },
  props: {
    selectedExploreId: {
      type: Number,
      required: true,
    },
  },
  methods: {

    isZoneSelected(proximity) {
      return proximity.exploreid === parseInt(this.selectedExplore)
    },

    selectZone(zone) {
      this.$emit('input', {
        zoneId: zone.zoneidnumber,
      });

      this.selectedExplore = zone.short_name
    },

    searchZone() {
      const searchString  = this.zoneSearch.toLowerCase().trim()
      let filteredExplore = []
      zones.forEach((zone) => {
        if (
          this.zoneSearch.trim() !== '' &&
          (
            zone.short_name.toLowerCase().includes(searchString) ||
            zone.long_name.toLowerCase().includes(searchString) ||
            zone.zoneidnumber.toString().includes(searchString)
          )) {
          filteredExplore.push(zone)
        }
      });
      this.filteredExplore = filteredExplore
      if (filteredExplore.length === 0) {
        this.filteredExplore = zones;
      }
    },

    getExpansionIcon(expansion) {
      return Expansions.getExpansionIconUrlSmall(expansion - 1) // zone table is offset by 1
    },
    getExpansionName(expansion) {
      return Expansions.getExpansionName(expansion - 1) // zone table is offset by 1
    },

    async loadProximities() {
      const api    = (new ProximityApi(SpireApiClient.getOpenApiConfig()))
      const result = await api.listProximities(
        (new SpireQueryBuilder())
          .orderBy(["exploreid"])
          .get()
      )

      if (result.status === 200) {
        zones                = result.data
        this.filteredExplore = zones
      }
    },

    init() {
      const t = document.getElementById("search-selector")
      if (t) {
        t.focus()
      }

      this.loadProximities().then(() => {
        this.scrollToSelected()
      })
    },

    scrollToSelected() {
      setTimeout(() => {
        const container = document.getElementById("explore-container");
        const target    = document.getElementById(util.format("proximity-%s", this.selectedExplore))
        if (container && target) {
          const top           = target.getBoundingClientRect().top
          container.scrollTop = container.scrollTop + top - 300;
        }
      }, 100)
    }
  },
  mounted() {
    // model we work with after the prop is passed - we can manipulate it ourselves
    if (this.selectedExploreId > 0) {
      this.selectedExplore = this.selectedExploreId
      this.zoneSearch      = this.selectedExploreId
    }

    Zones.getZones().then((r) => {
      let zones = {}
      r.forEach((zone) => {
        zones[zone.zoneidnumber] = zone
      })
      this.zones = zones

      this.init()
    })

  }
}
</script>

<style scoped>
#explore-table td {
  vertical-align: middle !important;

}
#explore-table td, #explore-table th {
  padding: 3px;
}
</style>
