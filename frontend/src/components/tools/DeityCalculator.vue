<template>
  <div class="pl-1 row" v-if="mask >= 0">
    <div
      class="mr-3 d-inline-block text-center"
      :style="(centeredButtons ? 'width: 100%' : '')"
    >
      <div v-for="(deity, deityId) in deities" class="mb-1 text-center d-inline-block">
        <div class="text-center pl-0 pr-0 mr-1 col-lg-12 col-sm-12">
          <small
            :style="(deity.short.length > 8 ? 'font-size: 9px' : 'font-size: 11px')"
            v-if="showNames"
          >{{ deity.short }}</small>
          <div class="text-center">
            <img
              :title="deity.name"
              @click="selectDeity(deityId)"
              :src="itemCdnUrl + 'item_' + deity.icon + '.png'"
              :style="getImageSize() + ' ' + (isDeitySelected(deityId) ? 'border: 2px solid #dadada; border-radius: 7px;' : 'border-radius: 7px; opacity: .6')"
              class="mt-1 hover-highlight" alt=""
            >
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
import {App}             from "@/constants/app";
import {DB_DIETIES_FULL} from "@/app/constants/eq-deities-constants";
import * as util         from "util";

export default {
  name: "DeityBitmaskCalculator",
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
    showNames: {
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
      deities: DB_DIETIES_FULL,
      itemCdnUrl: App.ASSET_ITEM_ICON_BASE_URL,
      selectedDeityes: {},
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
      Object.keys(this.deities).reverse().forEach((deityId) => {
        this.selectedDeityes[deityId] = true;
      });
      this.$forceUpdate();
      this.calculateToBitmask();
    },
    selectNone() {
      Object.keys(this.deities).reverse().forEach((deityId) => {
        this.selectedDeityes[deityId] = false;
      });
      this.$forceUpdate();
      this.calculateToBitmask();
    },
    calculateFromBitmask() {
      Object.keys(this.deities).reverse().forEach((deityId) => {
        const gameDeity               = this.deities[deityId];
        this.selectedDeityes[deityId] = false
        if (this.currentMask >= gameDeity.mask) {
          this.currentMask -= gameDeity.mask;
          this.selectedDeityes[deityId] = true;
        }
      });
      this.$forceUpdate()
    },
    calculateToBitmask() {
      let bitmask = 0;

      Object.keys(this.deities).reverse().forEach((deityId) => {
        const gameDeity = this.deities[deityId];
        if (this.selectedDeityes[deityId]) {
          bitmask += parseInt(gameDeity.mask);
        }
      });

      this.$emit("update:inputData", parseInt(bitmask));
      this.$emit("input", parseInt(bitmask));
      this.$emit("fired", "true");
    },
    selectDeity: function (deityId) {
      this.selectedDeityes[deityId] = !this.selectedDeityes[deityId];

      this.$forceUpdate()
      this.calculateToBitmask();
    },
    isDeitySelected: function (deityId) {
      return this.selectedDeityes[deityId]
    }
  }
}
</script>

<style scoped>

</style>
