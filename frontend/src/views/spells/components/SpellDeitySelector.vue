<template>
  <div class="row text-center mt-3 mb-3">
    <div class="col-12">
      <div v-for="(deity, deityId) in deities" :key="deityId" class="mb-1 d-inline-block text-center mr-3">
        <div class="text-center">
          <div style="width: 30px; height: 30px; font-size: 9px">
          {{ deity.name }}
          </div>
          <div class="text-center">
            <img
              @click="selectDeity(deityId)"
              :src="itemCdnUrl + 'item_' + deity.icon + '.png'"
              :style="'height: 35px; width:auto;' + selectedStyling(deityId)"
              class="mt-1 p-1 mb-2">
            <b-form-input
              v-model.number="spell['deities' + deityId]"
              @change="updateParent"
              style="width: 35px; display: block; padding: 0; padding-left: 9px !important"/>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import {App}                   from "@/constants/app";
import {DB_PLAYER_CLASSES_ALL} from "@/app/constants/eq-classes-constants";
import {DB_DIETIES_FULL}       from "@/app/constants/eq-deities-constants";

export default {
  name: "SpellDeitySelector",
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
      deities: DB_DIETIES_FULL,
      itemCdnUrl: App.ASSET_ITEM_ICON_BASE_URL,
    }
  },
  methods: {
    calculateFromBitmask() {
      Object.keys(this.classes).reverse().forEach((deityId) => {
        const gameDeity               = this.classes[deityId];
        this.selectedDeityes[deityId] = false
        if (this.currentMask >= gameDeity.mask) {
          this.currentMask -= gameDeity.mask;
          this.selectedDeityes[deityId] = true;
        }
      });
      this.$forceUpdate()
    },
    updateParent() {
      this.$emit('input', this.spell);
    },
    isDeitySelected: function (deityId) {
      return this.spell['deities' + deityId] > 0 && this.spell['deities' + deityId] < 255;
    },
    selectedStyling(deityId) {
      return (
        this.isDeitySelected(deityId) ?
          'border: 2px solid #dadada; border-radius: 7px;' :
          'border: 2px solid rgb(218 218 218 / 0%); border-radius: 7px; opacity: .3'
      )
    }
  }
}
</script>

<style scoped>

</style>
