<template>
  <div v-if="icons" class="align-content-center">
    <div
      @click="selectIcon(icon)"
      v-for="icon in icons"
      :key="icon"
      :class="'d-inline-block '"
      style="margin: 2px"
    >
      <span
        :style="'width: 40px; height: 40px; border: 1px solid; border-radius: 10px; display: inline-block'"
        :class="'spell-' + icon + '-40 ' + classIsPulsating(icon)"
        :title="icon"
      />
    </div>
  </div>
</template>

<script>
import {App}    from "../../../constants/app";
import EqAssets from "../../../app/eq-assets/eq-assets";

export default {
  name: "SpellIconSelector",
  data() {
    return {
      icons: [],
    }
  },
  props: {
    selectedIcon: {
      type: Number,
      default: 0,
      required: true
    },
  },
  created() {
    this.icons = EqAssets.getSpellIcons()
  },
  methods: {
    selectIcon(icon) {
      this.$emit("update:inputData", parseInt(icon));
    },
    classIsPulsating(icon) {
      return parseInt(icon) === parseInt(this.selectedIcon) ? 'pulsate' : ''
    },
  }
}
</script>

<style scoped>

</style>
