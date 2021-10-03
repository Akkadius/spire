<template>
  <div class="row">
    <div v-for="(expansion, expansionId) in expansions" class="row col-12">
      <div class="col-lg-12 col-sm-12" @click="selectExpansion(expansionId)">
        <div class="text-center" style="width: 70px; display: inline-block;">
        <img
          :src="getExpansionIconUrlSmall(expansionId)"
          :style="'width: 56px;' + (isExpansionSelected(expansionId) ? 'border: 2px solid #dadada; border-radius: 7px;' : 'border: 2px solid rgb(218 218 218 / 30%); border-radius: 7px;')"
          class="mt-1 p-1">
        </div>
          ({{expansionId}})
          {{ expansion.name }}s
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
import {EXPANSIONS_FULL} from "../../app/constants/eq-expansions";
import expansions from "../../app/utility/expansions";

export default {
  name: "ExpansionBitmaskCalculator",
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
      expansions: EXPANSIONS_FULL,
      itemCdnUrl: App.ASSET_ITEM_ICON_BASE_URL,
      selectedExpansiones: {},
      currentMask: 0
    }
  },
  mounted() {
    this.currentMask = parseInt(this.mask)
    this.calculateFromBitmask();
  },
  methods: {
    getExpansionIconUrlSmall(expansionId) {
      return expansions.getExpansionIconUrlSmall(expansionId)
    },
    selectAll() {
      Object.keys(this.expansions).reverse().forEach((expansionId) => {
        this.selectedExpansiones[expansionId] = true;
      });
      this.$forceUpdate();
      this.calculateToBitmask();
    },
    selectNone() {
      Object.keys(this.expansions).reverse().forEach((expansionId) => {
        this.selectedExpansiones[expansionId] = false;
      });
      this.$forceUpdate();
      this.calculateToBitmask();
    },
    calculateFromBitmask() {
      Object.keys(this.expansions).reverse().forEach((expansionId) => {
        const gameExpansion               = this.expansions[expansionId];
        this.selectedExpansiones[expansionId] = false
        if (this.currentMask >= gameExpansion.mask) {
          this.currentMask -= gameExpansion.mask;
          this.selectedExpansiones[expansionId] = true;
        }
      });
      this.$forceUpdate()
    },
    calculateToBitmask() {
      let bitmask = 0;

      Object.keys(this.expansions).reverse().forEach((expansionId) => {
        const gameExpansion = this.expansions[expansionId];
        if (this.selectedExpansiones[expansionId]) {
          bitmask += parseInt(gameExpansion.mask);
        }
      });

      this.$emit("update:inputData", bitmask.toString());
    },
    selectExpansion: function (expansionId) {
      this.selectedExpansiones[expansionId] = !this.selectedExpansiones[expansionId];

      this.$forceUpdate()
      this.calculateToBitmask();
    },
    isExpansionSelected: function (expansionId) {
      return this.selectedExpansiones[expansionId]
    }
  }
}
</script>

<style scoped>

</style>
