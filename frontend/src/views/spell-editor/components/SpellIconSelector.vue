<template>
  <div
    v-if="icons"
    class="align-content-center"
    id="spell-icon-view-port"
    style="height: 90vh; overflow-y: scroll"
  >
    <div
      @click="selectIcon(icon)"
      v-for="icon in icons"
      :key="icon"
      :class="'d-inline-block '"
      :id="'spell-icon-' + icon"
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
import EqAssets from "../../../app/eq-assets/eq-assets";
import util     from "util";

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

    if (this.selectedIcon > 0) {
      // we need 100ms delay because the videos haven't been rendered yet
      setTimeout(() => {
        console.time("[SpellIconSelector] scrollTo");
        const container = document.getElementById("spell-icon-view-port");
        const target    = document.getElementById(util.format("spell-icon-%s", this.selectedIcon))

        // 230 is height of video to offset
        if (container && target) {
          container.scrollTop = container.scrollTop + target.getBoundingClientRect().top - 250;
        }
        console.timeEnd("[SpellIconSelector] scrollTo");
      }, 1)
    }

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
