<template>
  <div>
    <div v-if="previewId === 0" class="mb-3">
      Spell casting animation preview not found...
    </div>
    <div v-if="previewId > 0">
      <video
        muted
        loop
        autoplay
        :style="'width: ' + this.width + 'px; height: ' + this.height + 'px; border-radius:5px; border: 1px solid rgba(255, 255, 255, .3)'"
        :src="videoSource"
      >
      </video>
    </div>
  </div>
</template>

<script>
import SpellAnimations from "@/app/eq-assets/spell-animations-map.json";
import {App}           from "../../../constants/app";

export default {
  name: "SpellAnimationPreview",
  data() {
    return {
      spellAnimationUrl: App.ASSET_SPELL_ANIMATIONS,
      previewId: 0,
      videoSource: ""
    }
  },
  watch: {
    id: {
      handler: function (val, oldVal) {
        console.log("trigger")
        this.render()
      },
    },
  },
  methods: {
    render() {
      this.previewId = 0;
      SpellAnimations[0].contents.forEach((row) => {
        const pieces      = row.name.split(/\//);
        const fileName    = pieces[pieces.length - 1].replace(".mp4", "");
        const animationId = parseInt(fileName)

        if (this.id === animationId) {
          this.previewId = this.id;

          this.videoSource = this.spellAnimationUrl + this.previewId + '.mp4';
          return false
        }
      })
    }
  },
  created() {
    this.render()
  },
  props: {
    id: {
      required: true,
      type: Number,
    },
    width: {
      default: 384,
      required: false,
      type: Number,
    },
    height: {
      default: 216,
      required: false,
      type: Number,
    },
  }
}
</script>

<style scoped>

</style>
