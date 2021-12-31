<template>
  <div v-if="icons" class="align-content-center">
    <div
      @click="selectIcon(icon)"
      v-for="icon in icons"
      :key="icon"
      :class="'d-inline-block '"
      style="margin: 2px">
      <img
        :src="spellCdnUrl + icon + '.gif'"
        :class="classIsPulsating(icon)"
        :style="'width:40px; height: 40px; border-radius:10px; border: 1px solid;'"
      >
    </div>
  </div>
</template>

<script>
import SpellIcons from "@/app/asset-maps/spell-icons-map.json";
import {App}      from "../../../constants/app";

export default {
  name: "SpellIconSelector",
  data() {
    return {
      spellCdnUrl: App.ASSET_SPELL_ICONS_BASE_URL,
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
    SpellIcons[0].contents.forEach((row) => {
      const pieces   = row.name.split(/\//);
      const fileName = pieces[pieces.length - 1];
      const iconId   = fileName.replace(".gif", "")

      this.icons.push(iconId);
    })
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
