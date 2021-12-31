<template>
  <div class="row" v-if="mask >= 0">
    <div
      class="ml-1 mr-3 d-inline-block"
      :style="(centeredButtons ? 'width: 100%; margin: 0;' : '')"
    >
      <div v-for="(gClass, classId) in classes" class="mb-1 d-inline-block">
        <div class="text-center p-0 mr-1 col-lg-12 col-sm-12">
          <span v-if="showTextTop">{{ gClass.short }}</span>
          <div class="text-center">
            <img
              :title="gClass.class"
              @click="selectClass(classId)"
              :src="itemCdnUrl + 'item_' + gClass.icon + '.png'"
              :style="getImageSize() + (isClassSelected(classId) ? 'border: 2px solid #dadada; border-radius: 7px;' : 'border-radius: 7px; opacity: .6')"
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
import {App}                   from "@/constants/app";
import {DB_PLAYER_CLASSES_ALL} from "@/app/constants/eq-classes-constants";
import util                    from "util";

export default {
  name: "ClassBitmaskCalculator",
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
    showTextSide: {
      type: Boolean,
      required: false,
      default: false
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
    getImageSize() {
      return util.format("width: %spx; height %spx;", this.imageSize, this.imageSize)
    },

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

      this.$emit("update:inputData", parseInt(bitmask));
      this.$emit("input", parseInt(bitmask));
      this.$emit("fired", "true");
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
