<template>
  <div class="row text-center mt-3 mb-3">
    <div class="col-12">
      <div v-for="(gClass, classId) in classes" class="mb-1 d-inline-block text-center mr-2">
        <div class="text-center">
          {{ gClass.short }}
          <div class="text-center">
            <img
              @click="selectClass(classId)"
              :src="itemCdnUrl + 'item_' + gClass.icon + '.png'"
              :style="'height: 45px; width:auto;' + selectedStyling(classId)"
              class="mt-1 p-1 mb-2">
            <b-form-input
              v-model.number="spell['classes_' + classId]"
              @change="updateParent"
              style="width: 45px; display: block; padding: 0; padding-left: 9px !important"/>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import {App}                   from "@/constants/app";
import {DB_PLAYER_CLASSES_ALL} from "@/app/constants/eq-classes-constants";

export default {
  name: "SpellClassSelector",
  props: {
    debug: {
      type: Boolean,
      required: false
    },
    spell: {
      type: Object,
      required: true
    },
  },
  data() {
    return {
      classes: DB_PLAYER_CLASSES_ALL,
      itemCdnUrl: App.ASSET_ITEM_ICON_BASE_URL,
    }
  },
  methods: {
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
    updateParent() {
      this.$emit('input', this.spell);
    },
    isClassSelected: function (classId) {
      return this.spell['classes_' + classId] > 0 && this.spell['classes_' + classId] < 255;
    },
    selectedStyling(classId) {
      return (
        this.isClassSelected(classId) ?
          'border: 2px solid #dadada; border-radius: 7px;' :
          'border: 2px solid rgb(218 218 218 / 0%); border-radius: 7px; opacity: .3'
      )
    }
  }
}
</script>

<style scoped>

</style>
