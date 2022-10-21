<template>
  <div class="row">
    <div v-for="(expansion, expansionId) in expansions" class="row col-12">
      <div class="col-lg-12 col-sm-12" @click="selectClientVersion(expansionId)">
        <div class="text-center mr-3" style="display: inline-block;">
          <img
            :style="'' + (isClientVersionSelected(expansionId) ? 'border: 2px solid #dadada; border-radius: 7px;' : 'border: 0px solid rgb(218 218 218 / 30%); border-radius: 0px;')"
            :class="'mt-1 p-0 client-version-sm-' + getClientVersionIcon(expansionId)"
          >
        </div>
        ({{ expansionId }})
        {{ expansion.name }}
      </div>
    </div>
    <div class="form-group text-center">
      <button class='eq-button mr-3' @click="selectAll()" style="display: inline-block; width: 80px">All</button>
      <button class='eq-button' @click="selectNone()" style="display: inline-block; width: 80px">None</button>
    </div>
  </div>
</template>

<script>
import {CLIENT_VERSIONS} from "../../app/constants/eq-client-versions";

export default {
  name: "ClientVersionCalculator",
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
      expansions: CLIENT_VERSIONS,
      selectedExpansiones: {},
      currentMask: 0
    }
  },
  mounted() {
    this.currentMask = parseInt(this.mask)
    this.calculateFromBitmask();
  },
  methods: {
    getClientVersionIcon(expansionId) {
      return CLIENT_VERSIONS[expansionId].icon
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
        const gameExpansion                   = this.expansions[expansionId];
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
    selectClientVersion: function (expansionId) {
      this.selectedExpansiones[expansionId] = !this.selectedExpansiones[expansionId];

      this.$forceUpdate()
      this.calculateToBitmask();
    },
    isClientVersionSelected: function (expansionId) {
      return this.selectedExpansiones[expansionId]
    }
  }
}
</script>

<style scoped>

</style>
