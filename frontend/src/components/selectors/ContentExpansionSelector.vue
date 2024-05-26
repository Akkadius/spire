<template>
  <div>
    <div v-for="(expansion, expansionId) in EXPANSIONS_FULL" class="d-inline-block">
      <img
        :alt="getExpansionName(expansionId)"
        v-if="!getExpansionIcon(expansionId).includes('base64')"
        :title="getExpansionName(expansionId) + ' (' + expansionId + ')'"
        :src="getExpansionIcon(expansionId)"
        @click="selectExpansion(expansionId)"
        :style="'width: 60px; ' + (isExpansionSelected(expansionId) ? 'border: 2px solid #dadada; border-radius: 7px;' : 'border: 2px solid rgb(218 218 218 / 30%); border-radius: 7px; opacity: .3')"
        class="mr-2 p-1 mt-1 hover-highlight-inner"
      >
      <span
        v-if="showNames"
        :style="isExpansionSelected(expansionId) ? 'font-weight: bold; opacity: 1;' : 'font-weight: normal; opacity: .7;'"
      >
        {{ expansion.name }}
      </span>
    </div>

    <div
      :style="'width: 56px; font-size: 12px; font-weight: bold; padding: 0px !important; opacity: 1; ' + (isExpansionSelected(-1) ? 'border: 2px solid #dadada; border-radius: 7px;' : 'border: 2px solid rgb(218 218 218 / 30%); border-radius: 7px;')"
      class="mr-2 mt-2 hover-highlight-inner text-center"
      @click="selectExpansion(-1)"
    >
      All
    </div>

    <div style="font-weight: bold; font-size: 14px;" class="mt-3">
      Selected ({{ selectedExpansion }}) {{ getExpansionName(selectedExpansion) }}

    </div>


  </div>
</template>

<script>
import {EXPANSIONS_FULL} from "../../app/constants/eq-expansions";
import Expansions        from "../../app/utility/expansions";

export default {
  name: "ContentExpansionSelector",
  data() {
    return {
      selectedExpansion: -1,

      EXPANSIONS_FULL: EXPANSIONS_FULL,
    }
  },
  watch: {
    value: function (newVal) {
      this.selectedExpansion = newVal
    }
  },
  props: {
    value: [Number, Array],
    showNames: {
      type: Boolean,
      default: false
    }
  },
  mounted() {
    this.selectedExpansion = parseInt(this.value)
  },
  methods: {
    selectExpansion(expansion) {
      console.log("select expansion [%s] value [%s]", expansion, this.value)
      expansion = parseInt(expansion)
      // loop avoidance - keep from circular updating
      if (this.value !== expansion) {
        this.selectedExpansion = expansion;
        this.$emit('input', expansion);
      }
    },
    isExpansionSelected(expansion) {
      return parseInt(expansion) === this.selectedExpansion
    },
    getExpansionIcon(expansion) {
      return Expansions.getExpansionIconUrlSmall(expansion)
    },
    getExpansionName(expansion) {
      return Expansions.getExpansionName(expansion)
    },
  }
}
</script>

<style scoped>

</style>
