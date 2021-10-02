<template>
  <div class="row">
    <div v-for="(deity, deityId) in deities" class="mb-3 text-center">
      <div class="text-center p-2 col-lg-12 col-sm-12">
        <small :style="(deity.name.length > 8 ? 'font-size: 9px' : '')">{{ deity.name }}</small>
        <div class="text-center">
          <img
            @click="selectDeity(deityId)"
            :src="itemCdnUrl + 'item_' + deity.icon + '.png'"
            :style="'width:auto;' + (isDeitySelected(deityId) ? 'border: 2px solid #dadada; border-radius: 7px;' : 'border: 1px solid rgb(218 218 218 / 30%); border-radius: 7px;')"
            class="mt-1 p-1">
        </div>
      </div>
    </div>
    <div class="form-group text-center">
      <button class='eq-button mr-3' @click="selectAll()" style="display: inline-block; width: 80px">All</button>
      <button class='eq-button' @click="selectNone()" style="display: inline-block; width: 80px">None</button>
    </div>
  </div>
</template>

<script>
import {App} from "../../constants/app";
import {DB_DIETIES_FULL} from "../../app/constants/eq-deities-constants";

export default {
  name: "DeityBitmaskCalculator",
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

      this.$emit("update:inputData", bitmask.toString());
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
