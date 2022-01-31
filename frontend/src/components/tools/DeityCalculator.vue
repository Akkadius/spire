<template>
  <div class="pl-1 row" v-if="mask >= 0">
    <div
      class="mr-3 d-inline-block text-center"
      :style="(centeredButtons ? 'width: 100%' : '')"
    >
      <div v-for="(deity, deityId) in deities" class="text-center d-inline-block">
        <div class="text-center pl-0 pr-0 mr-1 col-lg-12 col-sm-12">
          <small
            :style="(deity.short.length > 8 ? 'font-size: 9px' : 'font-size: 11px')"
            v-if="showNames"
          >{{ deity.short }}</small>
          <div class="text-center">
            <span
              :title="deity.name"
              @click="selectDeity(deityId)"
              :style="(isDeitySelected(deityId) ? 'border-radius: 3px;' : 'border-radius: 3px; opacity: .6')"
              :class="'hover-highlight-inner item-' + deity.icon + ' ' + (isDeitySelected(deityId) ? 'highlight-selected-inner' : '')"
            />
          </div>
        </div>
      </div>

      <!-- Select All / None -->
      <div
        class="d-inline-block"
        v-if="displayAllNone"
        :style="'line-height: 25px; bottom: ' + (centeredButtons ? -10 : 15) + 'px; position: relative;'"
      >
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
      selectedDeityes: {},
      currentMask: 0
    }
  },
  mounted() {
    this.currentMask = parseInt(this.mask)
    this.calculateFromBitmask();
  },
  methods: {
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
