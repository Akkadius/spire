<template>
  <div class="col-12 mr-0 mt-3 mb-3 text-center">
    <div
      v-for="(gClass, classId) in classes"
      class="mb-1 d-inline-block text-center"
      style="padding-right: 5px"
    >
      <div class="text-center">
        {{ gClass.short }}
        <div class="text-center">
          <span
            :style="(isClassSelected(classId) ? 'border-radius: 3px;' : 'border-radius: 3px; opacity: .6')"
            :class="'item-' + gClass.icon + ' ' + (isClassSelected(classId) ? 'highlight-selected-inner' : '')"
            class="mt-1 mb-2"
          />
          <b-form-input
            v-model.number="spell['classes_' + classId]"
            @change="updateParent"
            :id="'classes_' + classId"
            style="width: 45px; display: block; padding: 0; padding-left: 9px !important"
          />
        </div>
      </div>
    </div>
  </div>
</template>

<script>
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
