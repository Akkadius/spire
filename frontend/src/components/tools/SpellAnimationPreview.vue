<template>
  <div>
    <div v-if="previewId === 0">
      Spell casting animation preview not found...
    </div>
    <div v-if="previewId > 0">
      <video
        muted
        loop
        autoplay
        :style="'height: ' + this.height + 'px; border-radius:10px; border: 1px solid;'"
        class="spell-preview">
        <source :src="spellAnimationUrl + previewId + '.mp4'" type="video/mp4">
      </video>
    </div>
  </div>
</template>

<script>
import SpellAnimations from "@/app/asset-maps/spell-animations-map.json";
import {App}           from "../../constants/app";

export default {
  name: "SpellAnimationPreview",
  data() {
    return {
      spellAnimationUrl: App.ASSET_SPELL_ANIMATIONS,
      previewId: 0,
    }
  },
  created() {
    SpellAnimations[0].contents.forEach((row) => {
      const pieces      = row.name.split(/\//);
      const fileName    = pieces[pieces.length - 1].replace(".mp4", "");
      const animationId = parseInt(fileName)

      console.log(row)
      console.log(animationId)
      console.log(this.id)

      if (this.id === animationId) {
        this.previewId = this.id;
        return false
      }
    })
  },
  props: {
    id: {
      required: true,
      type: Number,
    },
    height: {
      default: 300,
      required: false,
      type: Number,
    },
  }
}
</script>

<style scoped>

</style>
