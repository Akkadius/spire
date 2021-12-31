<template>
  <div class="row" v-if="mask >= 0">
    <div
      class="ml-1 mr-3 d-inline-block text-center"
      :style="(centeredButtons ? 'width: 100%' : '')"
    >
      <div v-for="(race, index) in races" class="mb-1 text-center d-inline-block">
        <div class="text-center p-0 mr-1 col-lg-12 col-sm-12">
          <span v-if="showTextTop">{{ race.short }}</span>
          <div class="text-center">
            <img
              :title="race.race"
              @click="selectRace(index)"
              :src="itemCdnUrl + 'item_' + race.icon + '.png'"
              :style="getImageSize() + (isRaceSelected(index) ? 'border-radius: 7px; box-shadow: 0px 0px 1px 1px white;' : 'border-radius: 7px; opacity: .5')"
              class="hover-highlight">
          </div>
        </div>
      </div>

      <!-- Select All / None -->
      <div class="d-inline-block" v-if="displayAllNone">
        <div
          :class="'text-center mt-2 btn-xs eq-button-fancy ' + (parseInt(mask) >= 65535 ? 'eq-button-fancy-highlighted' : '')"
          @click="selectAll()"
        >
          All
        </div>
        <div
          :class="'text-center mt-2 btn-xs eq-button-fancy ' + (parseInt(mask) === 0 ? 'eq-button-fancy-highlighted' : '')"
          @click="selectNone()"
        >
          None
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import {DB_PLAYER_RACES} from "@/app/constants/eq-races-constants";
import {App}             from "@/constants/app";
import util              from "util";

export default {
  name: "RaceBitmaskCalculator",
  props: {
    debug: {
      type: Boolean,
      required: false
    },
    mask: {
      type: Number,
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
    },
    showTextTop: {
      type: Boolean,
      required: false,
      default: true
    },
    imageSize: {
      type: Number,
      required: false,
      default: 50,
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
    getImageSize() {
      return util.format("width: %spx; height %spx;", this.imageSize, this.imageSize)
    },
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

      this.$emit("update:inputData", parseInt(bitmask));
      this.$emit("input", parseInt(bitmask));
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
