<template>
  <div class="row">
    <div v-for="(gClass, classId) in classes" class="mb-3 text-center">
      <div class="text-center p-1 col-lg-12 col-sm-12">
        {{ gClass.short }}
        <div class="text-center">
          <img
            @click="selectClass(classId)"
            :src="itemCdnUrl + 'item_' + gClass.icon + '.png'"
            :style="'width:auto;' + (isClassSelected(classId) ? 'border: 2px solid #dadada; border-radius: 7px;' : 'border: 2px solid rgb(218 218 218 / 0%); border-radius: 7px;')"
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
import {DB_PLAYER_CLASSES_ALL} from "../../app/constants/eq-classes-constants";

export default {
  name: "ClassBitmaskCalculator",
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
      classes: DB_PLAYER_CLASSES_ALL,
      itemCdnUrl: App.ASSET_ITEM_ICON_BASE_URL,
      selectedClasses: {},
      currentMask: 0
    }
  },
  mounted() {
    this.currentMask = parseInt(this.mask)
    this.calculateFromBitmask();
  },
  methods: {
    selectAll() {
      Object.keys(this.classes).reverse().forEach((classId) => {
        this.selectedClasses[classId] = true;
      });
      this.$forceUpdate();
      this.calculateToBitmask();
    },
    selectNone() {
      Object.keys(this.classes).reverse().forEach((classId) => {
        this.selectedClasses[classId] = false;
      });
      this.$forceUpdate();
      this.calculateToBitmask();
    },
    calculateFromBitmask() {
      Object.keys(this.classes).reverse().forEach((classId) => {
        const gameClass               = this.classes[classId];
        this.selectedClasses[classId] = false
        if (this.currentMask >= gameClass.mask) {
          this.currentMask -= gameClass.mask;
          this.selectedClasses[classId] = true;
        }
      });
      this.$forceUpdate()
    },
    calculateToBitmask() {
      let bitmask = 0;

      Object.keys(this.classes).reverse().forEach((classId) => {
        const gameClass = this.classes[classId];
        if (this.selectedClasses[classId]) {
          bitmask += parseInt(gameClass.mask);
        }
      });

      this.$emit("update:inputData", bitmask.toString());
    },
    selectClass: function (classId) {
      this.selectedClasses[classId] = !this.selectedClasses[classId];

      this.$forceUpdate()
      this.calculateToBitmask();
    },
    isClassSelected: function (classId) {
      return this.selectedClasses[classId]
    }
  }
}
</script>

<style scoped>

</style>
