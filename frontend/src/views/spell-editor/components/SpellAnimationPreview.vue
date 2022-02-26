<template>
  <div>
    <div
      v-if="previewId === 0"
      :style="'width: 100%; height: 235px; border-radius:5px; border: 1px solid rgba(255, 255, 255, .3)'"
      class="text-center mb-2"
    >
      <div
        class="pt-5 pl-4 pr-4 p-7"
        style="min-height: 111px; width: 100%; background-color: black;"
      >
        Spell casting animation preview not found...
      </div>
    </div>

    <div v-if="previewId > 0">
      <video
        muted
        loop
        autoplay
        class="video-preview"
        :style="'width: 100%; min-height: 235px; height: auto; border-radius:5px; border: 1px solid rgba(255, 255, 255, .3); background-color: black;'"
        :src="videoSource"
      >
      </video>
    </div>
  </div>
</template>

<script>
import {App}       from "../../../constants/app";
import VideoViewer from "@/app/video-viewer/video-viewer";
import EqAssets    from "@/app/eq-assets/eq-assets";

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
        this.render()
        VideoViewer.handleRender();
      },
    },
  },
  methods: {
    render() {
      this.previewId = 0;

      EqAssets.getSpellAnimationFileIds().forEach((animationId) => {
        if (this.id === animationId) {
          this.previewId = this.id;

          this.videoSource = this.spellAnimationUrl + this.previewId + '.mp4#t=3';
          return false
        }
      })
    }
  },
  activated() {
    VideoViewer.handleRender();
  },
  created() {
    this.render()
  },
  props: {
    id: {
      required: true,
      type: Number,
    },
  }
}
</script>

<style scoped>

</style>
