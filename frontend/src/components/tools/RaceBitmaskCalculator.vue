<template>
  <div class="row">
    <div v-for="(race, index) in races" class="mb-3 text-center">
      <div class="text-center p-1 col-lg-12 col-sm-12">
        {{ race.race }}
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
</template>

<script>
import {DB_PLAYER_RACES} from "../../app/constants/eq-races-constants";
import {App} from "../../constants/app";

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
    }
  },
  watch: {
    mask: {
      // the callback will be called immediately after the start of the observation
      immediate: true,
      handler (val, oldVal) {
        this.currentMask = parseInt(this.mask)
        this.calculateFromBitmask();
      }
    },
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
    calculateFromBitmask() {
      Object.keys(this.races).reverse().forEach((raceId) => {
        const race = this.races[raceId];
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
