<template>
  <div class="row">
    <div v-for="(aug, index) in augs" class="row col-12">
      <div class="p-1 col-lg-12 col-sm-12" @click="selectAug(index)">
        <img
          :src="slotUrl + 'blankslot3.gif'"
          :style="'width:auto;height: 20px; ' + (isAugSelected(index) ? 'border: 2px solid #dadada; border-radius: 7px;' : 'border: 2px solid rgb(218 218 218 / 0%); border-radius: 7px;')"
          class="">
        ({{ index }})
        {{ aug.name }}
      </div>
      <br>
    </div>
    <div class="form-group text-center">
      <button class='eq-button mr-3' @click="selectAll()" style="display: inline-block; width: 80px">All</button>
      <button class='eq-button' @click="selectNone()" style="display: inline-block; width: 80px">None</button>
    </div>
  </div>
</template>

<script>
import {App} from "../../constants/app";
import {AUG_TYPES} from "../../app/constants/eq-aug-constants";
import itemTypesIconMapping from "@/constants/item-type-icon-mapping.json"

export default {
  name: "AugBitmaskCalculator",
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
      augs: AUG_TYPES,
      slotUrl: App.ASSET_INVENTORY_SLOT_URL,
      selectedAugs: {},
      currentMask: 0
    }
  },
  mounted() {
    this.currentMask = parseInt(this.mask)
    this.calculateFromBitmask();
  },
  methods: {
    selectAll() {
      Object.keys(this.augs).reverse().forEach((augId) => {
        this.selectedAugs[augId] = true;
      });
      this.$forceUpdate();
      this.calculateToBitmask();
    },
    selectNone() {
      Object.keys(this.augs).reverse().forEach((augId) => {
        this.selectedAugs[augId] = false;
      });
      this.$forceUpdate();
      this.calculateToBitmask();
    },
    calculateFromBitmask() {
      Object.keys(this.augs).reverse().forEach((augId) => {
        const aug                = this.augs[augId];
        this.selectedAugs[augId] = false
        if (this.currentMask >= aug.mask) {
          this.currentMask -= aug.mask;
          this.selectedAugs[augId] = true;
        }
      });
      this.$forceUpdate()
    },
    calculateToBitmask() {
      let bitmask = 0;

      Object.keys(this.augs).reverse().forEach((augId) => {
        const aug = this.augs[augId];
        if (this.selectedAugs[augId]) {
          bitmask += parseInt(aug.mask);
        }
      });

      this.$emit("update:inputData", bitmask.toString());
    },
    selectAug: function (augId) {
      this.selectedAugs[augId] = !this.selectedAugs[augId];

      this.$forceUpdate()
      this.calculateToBitmask();
    },
    isAugSelected: function (augId) {
      return this.selectedAugs[augId]
    }
  }
}
</script>

<style scoped>

</style>
