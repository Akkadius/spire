<template>
  <div>
    <eq-window-simple title="Zone Selector">
      <b-input
        v-model="zoneSearch"
        id="search-selector"
        class="form-control"
        v-on:keyup="searchZone"
        placeholder="Search by zone name, id..."
      />
    </eq-window-simple>

    <eq-window-simple
      id="zone-view-container"
      style="height: 85vh; overflow-y: scroll;" class="p-0"
    >
      <table
        id="zonetable"
        class="eq-table eq-highlight-rows bordered"
        style="display: table; font-size: 14px; overflow-x: scroll"
      >
        <thead
          class="eq-table-floating-header"
        >
        <tr>
          <th></th>
          <th style="width: 60px; white-space: nowrap;"></th>
          <th style="width: 30px">ID</th>
          <th style="width: 30px">X</th>
          <th style="width: 30px">Y</th>
          <th style="width: 30px">Z</th>
          <th style="width: 100%">Name</th>
        </tr>
        </thead>
        <tbody>
        <tr
          :id="'zone-' + zone.zoneidnumber"
          :class="(isZoneSelected(zone) ? 'pulsate-highlight-white' : '')"
          v-for="(zone, index) in filteredZones"
          :key="zone.id"
        >
          <td>
            <b-button
              class="btn-dark btn-sm btn-outline-warning"
              @click="selectZone(zone)"
            >
              <i class="fa fa-arrow-left"></i>
            </b-button>
          </td>
          <td style="text-align: center"><img :src="getExpansionIcon(zone.expansion)"></td>
          <td style="text-align: center" class="p-0">{{ zone.zoneidnumber }}</td>
          <td style="text-align: center" class="p-0">{{ Math.round(zone.safe_x) }}</td>
          <td style="text-align: center" class="p-0">{{ Math.round(zone.safe_y) }}</td>
          <td style="text-align: center" class="p-0">{{ Math.round(zone.safe_z) }}</td>
          <td>{{ zone.long_name }} ({{ zone.short_name }})</td>
        </tr>
        </tbody>
      </table>
    </eq-window-simple>
  </div>
</template>

<script>
import {TELEPORT_ZONE_SELECTOR_TYPE} from "@/app/constants/eq-spell-constants";
import EqWindowSimple                from "@/components/eq-ui/EQWindowSimple";
import {ZoneApi}  from "@/app/api";
import {SpireApi} from "../../../app/api/spire-api";
import util       from "util";
import Expansions                    from "@/app/utility/expansions";
import EqCheckbox                    from "@/components/eq-ui/EQCheckbox";
import {SpireQueryBuilder}           from "@/app/api/spire-query-builder";

let zones = {}

export default {
  name: "TaskZoneSelector",
  components: { EqCheckbox, EqWindowSimple },
  data() {
    return {
      TELEPORT_ZONE_SELECTOR_TYPE: TELEPORT_ZONE_SELECTOR_TYPE,

      // filtered content
      filteredZones: {},

      // search
      zoneSearch: "",

      // model we work with after the prop is passed so we can manipulate it ourselves
      // props should not be mutated
      selectedZone: "",
    }
  },
  props: {
    selectedZoneId: {
      type: Number,
      required: true,
    },
  },
  methods: {
    isZoneSelected(zone) {
      return zone.zoneidnumber === parseInt(this.selectedZone)
    },

    selectZone(zone) {
      this.$emit('input', {
        zoneId: zone.zoneidnumber,
      });

      // window.open("https://www.youtube.com/watch?v=dQw4w9WgXcQ")

      this.selectedZone = zone.zoneidnumber
    },

    searchZone() {
      const searchString = this.zoneSearch.toLowerCase().trim()
      let filteredZones  = []
      zones.forEach((zone) => {
        if (
          this.zoneSearch.trim() !== '' &&
          (
            zone.short_name.toLowerCase().includes(searchString) ||
            zone.long_name.toLowerCase().includes(searchString) ||
            zone.zoneidnumber.toString().includes(searchString)
          )) {
          filteredZones.push(zone)
        }
      });
      this.filteredZones = filteredZones
      if (filteredZones.length === 0) {
        this.filteredZones = zones;
      }
    },

    getExpansionIcon(expansion) {
      return Expansions.getExpansionIconUrlSmall(expansion)
    },

    async loadZones() {
      const api    = (new ZoneApi(...SpireApi.cfg()))
      const result = await api.listZones(
        (new SpireQueryBuilder())
          .where("version", "=", "0")
          .orderBy(["expansion", "short_name"])
          .get()
      )

      if (result.status === 200) {
        zones              = result.data
        this.filteredZones = zones
      }
    },

    init() {
      const t = document.getElementById("search-selector")
      if (t) {
        t.focus()
      }

      this.loadZones().then(() => {
        this.scrollToSelected()
      })
    },

    scrollToSelected() {
      setTimeout(() => {
        const container = document.getElementById("zone-view-container");
        const target    = document.getElementById(util.format("zone-%s", this.selectedZone))
        if (container && target) {
          const top           = target.getBoundingClientRect().top
          container.scrollTop = container.scrollTop + top - 300;
        }
      }, 100)
    }
  },
  mounted() {
    // model we work with after the prop is passed - we can manipulate it ourselves
    if (this.selectedZoneId > 0) {
      this.selectedZone = this.selectedZoneId
      this.zoneSearch   = this.selectedZoneId
    }

    this.init()
  }
}
</script>

<style scoped>
#zonetable td {
  vertical-align: middle !important;
}
</style>
