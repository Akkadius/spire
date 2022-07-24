<template>
  <div>
    <div
      v-if="previewId === 0"
      :style="'width: 100%; height: auto; min-height: 111px; border-radius:5px; border: 1px solid rgba(255, 255, 255, .3)'"
      class="text-center mb-2"
    >
      <div
        class="pt-5 pl-4 pr-4"
        style="min-height: 111px; width: 100%; background-color: black;"
      >
        Spell casting animation preview not found...
      </div>
    </div>
    <div v-if="previewId > 0">
      <video
        class="video-preview"
        muted
        loop
        autoplay
        :style="'width: 100%; height: auto; min-height: 111px; border-radius:5px; border: 1px solid rgba(255, 255, 255, .3); background-color: black;'"
        :src="videoSource"
      />
    </div>
  </div>
</template>

<script>
import {App}       from "../../../constants/app";
import EqAssets    from "@/app/eq-assets/eq-assets";
import VideoViewer from "@/app/video-viewer/video-viewer";

export default {
  name: "SpellCastingAnimationPreview",
  data() {
    return {
      playerAnimationClipUrl: App.ASSET_PLAYER_ANIMATION_CLIPS,
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
    async render() {
      this.previewId = 0;

      const r = await EqAssets.getPlayerAnimationFileIds()
      r.forEach((animationId) => {
        if (this.id === animationId) {
          this.previewId   = this.id;
          this.videoSource = this.playerAnimationClipUrl + this.previewId + '.mp4';
          return false
        }
      })
    }
  },
  async created() {
    await this.render()
    VideoViewer.handleRender();
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
