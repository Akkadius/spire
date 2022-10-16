<template>
  <div>
    <eq-window-simple title="Horse Selector">
      <b-input
        v-model="horseSearch"
        class="form-control"
        v-on:keyup="searchHorse"
        placeholder="Search by horse name..."
      />
    </eq-window-simple>

    <eq-window-simple
      id="horse-view-container"
      style="height: 85vh; overflow-y: scroll;" class="p-0"
    >
      <table
        id="horsetable"
        class="eq-table eq-highlight-rows"
        style="display: table; font-size: 14px; overflow-x: scroll; width: 1200px"
      >
        <thead
          class="eq-table-floating-header"
        >
        <tr>
          <th style="width: 30px"></th>
          <th style="width: 30px">Race</th>
          <th style="width: 30px">Gender</th>
          <th style="width: 30px">Texture</th>
          <th style="width: 30px">Speed</th>
          <th style="width: 100px">Name</th>
          <th>Notes</th>
        </tr>
        </thead>
        <tbody>
        <tr
          :id="'horse-' + stripNonAlpha(horse.filename)"
          :class="(isHorseSelected(horse) ? 'pulsate-highlight-white' : '')"
          v-for="(horse, index) in filteredHorses"
          :key="horse.id"
        >
          <td>
            <b-button
              class="btn-dark btn-sm btn-outline-warning"
              @click="selectHorse(horse)"
            >
              <i class="fa fa-arrow-left"></i>
            </b-button>
          </td>
          <td style="text-align: center" class="p-0">{{ horse.race }}</td>
          <td style="text-align: center" class="p-0">{{ horse.gender }}</td>
          <td style="text-align: center" class="p-0">{{ horse.texture }}</td>
          <td style="text-align: center" class="p-0">{{ horse.mountspeed }}</td>
          <td style="text-align: left">{{ horse.filename }}</td>
          <td style="text-align: left">{{ horse.notes }}</td>
        </tr>
        </tbody>
      </table>
    </eq-window-simple>
  </div>
</template>

<script>
import {TELEPORT_ZONE_SELECTOR_TYPE} from "@/app/constants/eq-spell-constants";
import EqWindowSimple                from "@/components/eq-ui/EQWindowSimple";
import {HorseApi}                    from "@/app/api";
import {SpireApi}              from "@/app/api/spire-api";
import util                          from "util";
import Expansions                    from "@/app/utility/expansions";
import EqCheckbox                    from "@/components/eq-ui/EQCheckbox";
import {SpireQueryBuilder}           from "@/app/api/spire-query-builder";

let horses = {}

export default {
  name: "SpellTeleportZoneSelectorHorse",
  components: { EqCheckbox, EqWindowSimple },
  data() {
    return {
      TELEPORT_ZONE_SELECTOR_TYPE: TELEPORT_ZONE_SELECTOR_TYPE,

      // filtered content
      filteredHorses: {},

      // search
      horseSearch: "",

      // model we work with after the prop is passed so we can manipulate it ourselves
      // props should not be mutated
      selectedHorse: "",
    }
  },
  props: {
    selectedHorseName: {
      type: String,
      required: true,
    },
  },
  methods: {

    stripNonAlpha(string) {
      return string.replace(/[\W_]+/g, " ");
    },

    isHorseSelected(horse) {
      return horse.filename.trim() === this.selectedHorse
    },

    selectHorse(horse) {
      this.$emit('input', {
        horse: horse,
      });

      this.selectedHorse = horse.filename
    },

    searchHorse() {
      const searchString = this.horseSearch.toLowerCase().trim()
      let filteredHorses = []
      horses.forEach((horse) => {
        if (this.horseSearch.trim() !== '' && horse.filename.toLowerCase().includes(searchString)) {
          filteredHorses.push(horse)
        }
      });
      this.filteredHorses = filteredHorses

      if (filteredHorses.length === 0) {
        this.filteredHorses = horses;
      }
    },
    
    async loadHorses() {
      const api    = (new HorseApi(...SpireApi.cfg()))
      const result = await api.listHorses(
        (new SpireQueryBuilder())
          .groupBy(["filename"])
          .get()
      )

      if (result.status === 200) {
        horses              = result.data
        this.filteredHorses = horses
      }
    },

    init() {
      this.loadHorses()
    }
  },
  mounted() {
    // model we work with after the prop is passed - we can manipulate it ourselves
    this.selectedHorse = this.selectedHorseName
    this.init()

    setTimeout(() => {
      const container = document.getElementById("horse-view-container");
      const target    = document.getElementById(util.format("horse-%s", this.selectedHorse))
      if (container && target) {
        const top           = target.getBoundingClientRect().top
        container.scrollTop = container.scrollTop + top - 300;
      }
    }, 1000)
  }
}
</script>

<style scoped>
#horsetable td {
  vertical-align: middle !important;
}
</style>
