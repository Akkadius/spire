<template>
  <div class="row">
    <div class="mr-3 d-inline-block text-center">
      <div v-for="(race, index) in races" class="mb-1 text-center d-inline-block">
        <div class="text-center p-1 col-lg-12 col-sm-12">
          {{ race.short }}
          <div class="text-center">
            <img
              @click="selectRace(index)"
              :src="itemCdnUrl + 'item_' + race.icon + '.png'"
              :style="'width:auto;' + (isRaceSelected(index) ? 'border: 2px solid #dadada; border-radius: 7px;' : 'border: 2px solid rgb(218 218 218 / 0%); border-radius: 7px;')"
              class="mt-1 p-1">
          </div>
        </div>
      </div>
    </div>
    <div :class="'mt-4 d-inline-block ' + (centeredButtons ? 'text-center w-100' : '')" v-if="displayAllNone">
      <button class='eq-button mr-3' @click="selectAll()" style="display: inline-block; width: 80px">All</button>
      <button class='eq-button' @click="selectNone()" style="display: inline-block; width: 80px">None</button>
    </div>
  </div>
</template>

<script>
import {DB_PLAYER_RACES} from "@/app/constants/eq-races-constants";
import {App} from "@/constants/app";

export default {
  name: "RaceBitmaskCalculator",
  props: {
    debug: {
      type: Boolean,
      required: false
    },
    mask: {
      type: String,
      required: false
    },
    displayAllNone: {
      type: Boolean,
      required: false,
      default: true
    },
    centeredButtons: {
      type: Boolean,
      required: false,
      default: true
    }
  },
  watch: {
    mask: {
      // the callback will be called immediately after the start of the observation
      immediate: true,
      handler(val, oldVal) {
        this.currentMask = parseInt(this.mask)
        this.calculateFromBitmask();
      }
    }
  },
  data() {
    return {
      races: DB_PLAYER_RACES,
      itemCdnUrl: App.ASSET_ITEM_ICON_BASE_URL,
      selectedRaces: {},
      currentMask: 0
    }
  },
  mounted() {
    this.currentMask = parseInt(this.mask)
    this.calculateFromBitmask();
  },
  methods: {
    selectAll() {
      Object.keys(this.races).reverse().forEach((raceId) => {
        this.selectedRaces[raceId] = true;
      });
      this.$forceUpdate();
      this.calculateToBitmask();
    },
    selectNone() {
      Object.keys(this.races).reverse().forEach((raceId) => {
        this.selectedRaces[raceId] = false;
      });
      this.$forceUpdate();
      this.calculateToBitmask();
    },
    calculateFromBitmask() {
      Object.keys(this.races).reverse().forEach((raceId) => {
        const race                 = this.races[raceId];
        this.selectedRaces[raceId] = false
        if (this.currentMask >= race.mask) {
          this.currentMask -= race.mask;
          this.selectedRaces[raceId] = true;
        }
      });
      this.$forceUpdate()
    },
    calculateToBitmask() {
      let bitmask = 0;

      Object.keys(this.races).reverse().forEach((raceId) => {
        const race = this.races[raceId];
        if (this.selectedRaces[raceId]) {
          bitmask += parseInt(race.mask);
        }
      });

      this.$emit("update:inputData", bitmask.toString());
      this.$emit("fired", "true");
    },
    selectRace: function (raceId) {
      this.selectedRaces[raceId] = !this.selectedRaces[raceId];

      this.$forceUpdate()
      this.calculateToBitmask();
    },
    isRaceSelected: function (raceId) {
      return this.selectedRaces[raceId]
    }
  }
}
</script>

<style scoped>

</style>
