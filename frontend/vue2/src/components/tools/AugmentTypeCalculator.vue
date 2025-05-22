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

    <!-- Select All / None -->
    <div class="d-inline-block text-center">
      <div
        :class="'text-center mt-4 btn-xs eq-button-fancy ' + (parseInt(mask) >= 8388607 ? 'eq-button-fancy-highlighted' : '')"
        @click="selectAll()"
      >
        All
      </div>
      <div
        :class="'text-center mt-4 btn-xs eq-button-fancy ' + (parseInt(mask) === 0 ? 'eq-button-fancy-highlighted' : '')"
        @click="selectNone()"
      >
        None
      </div>
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
      type: [Number, String],
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

      this.$emit("update:inputData", parseInt(bitmask));
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
