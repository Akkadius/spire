<template>
  <div>
    <img
      :alt="getExpansionName(expansionId)"
      v-for="(expansion, expansionId) in EXPANSIONS_FULL"
      v-if="!getExpansionIcon(expansionId).includes('base64')"
      :title="getExpansionName(expansionId) + ' (' + expansionId + ')'"
      :src="getExpansionIcon(expansionId)"
      @click="selectExpansion(expansionId)"
      :style="'width: 56px; opacity: .7; ' + (isExpansionSelected(expansionId) ? 'border: 2px solid #dadada; border-radius: 7px;' : 'border: 2px solid rgb(218 218 218 / 30%); border-radius: 7px;')"
      class="mr-2 p-1 mt-1 hover-highlight-inner"
    >

    <div
      :style="'width: 56px; font-size: 12px; font-weight: bold; padding: 0px !important; opacity: .7; ' + (isExpansionSelected(-1) ? 'border: 2px solid #dadada; border-radius: 7px;' : 'border: 2px solid rgb(218 218 218 / 30%); border-radius: 7px;')"
      class="mr-2 mt-2 hover-highlight-inner d-inline-block text-center"
      @click="selectExpansion(-1)"
    >
      All
    </div>

    ({{ selectedExpansion }}) {{ getExpansionName(selectedExpansion) }}

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
  props: {
    value: [Number, Array]
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
