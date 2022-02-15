<template>
  <div>
    <app-loader :is-loading="!loaded" padding="8"/>

    <eq-window-simple
      class="text-center p-0"
    >
      <div
        style="height: 90vh; overflow-y: scroll"
        v-on:scroll="handleRender"
        id="video-view-port"
      >
        <div class="col-12">
          <div v-if="filteredAnimations && filteredAnimations.length === 0">
            No animations found...
          </div>

          <div
            v-for="(animationId) in filteredAnimations"
            :key="animationId"
            style="display:inline-block; position: relative;"
            class="d-inline-block"
          >
            <video
              muted
              loop
              style="height: 146px; width: 259px; border-radius: 5px; border: 1px solid rgba(255, 255, 255, .3);"
              :id="'spell-' + animationId"
              :data-src="animBaseUrl + animationId + '.mp4'"
              @mousedown="selectSpellAnim(animationId)"
              :class="'video-preview ' + classIsPulsating(animationId)"
            >
            </video>

            <div class="overlay-spell-anim-selector">
              <h6 class="eq-header" style="font-size: 21px; ">{{ animationId }}</h6>
            </div>

          </div>
        </div>

      </div>

    </eq-window-simple>

  </div>
</template>

<script>
import PageHeader     from "@/components/layout/PageHeader";
import {App}          from "@/constants/app";
import EqWindow       from "@/components/eq-ui/EQWindow";
import * as util      from "util";
import VideoViewer    from "../../../app/video-viewer/video-viewer";
import EqWindowSimple from "../../../components/eq-ui/EQWindowSimple";
import EqAssets       from "../../../app/eq-assets/eq-assets";

let animationPreviewExists = {}

export default {
  name: "SpellCastingAnimationSelector",
  components: { EqWindowSimple, EqWindow, PageHeader },
  data() {
    return {
      loaded: false,
      animations: [],
      filteredAnimations: [],
      search: "",
      animBaseUrl: App.ASSET_PLAYER_ANIMATION_CLIPS,
    }
  },
  props: {
    selectedAnimation: {
      type: Number,
      default: 0,
      required: true
    },
  },
  methods: {
    init() {
      this.render()

      // bring focus to the selected video
      if (this.selectedAnimation > 0) {
        // we need 100ms delay because the videos haven't been rendered yet
        setTimeout(() => {
          const container = document.getElementById("video-view-port");
          const target    = document.getElementById(util.format("spell-%s", this.selectedAnimation))

          // 230 is height of video to offset
          if (container && target) {
            const top = target.getBoundingClientRect().top

            // hack to get all of the videos to load without top and bottom clipping
            setTimeout(() => {
              container.scrollTop = top - 400;
              VideoViewer.handleRender();
            }, 1)

            setTimeout(() => {
              container.scrollTop = top;
              VideoViewer.handleRender();
            }, 100)

            setTimeout(() => {
              container.scrollTop = top - 150;
              VideoViewer.handleRender();
            }, 300)

            // container.scrollTop = target.getBoundingClientRect().top - 150;
          }
        }, 300)
      }

    },
    handleRender() {
      VideoViewer.handleRender()
    },
    render: function () {
      // Preload model files
      let modelFiles = [];
      EqAssets.getPlayerAnimationFileIds().forEach((animationId) => {
        modelFiles.push(animationId)
        animationPreviewExists[animationId] = 1
      })

      // Sort by spell animation number
      modelFiles.sort(function (a, b) {
        return a - b;
      });

      this.filteredAnimations = modelFiles
      this.loaded             = true

      VideoViewer.handleRender()
    },
    classIsPulsating(animation) {
      return animation === this.selectedAnimation ? 'pulsate' : ''
    },
    triggerSearch() {
      this.spellAnimSearch();
    },
    selectSpellAnim(anim) {
      this.$emit("update:inputData", anim);
    }
  },
  activated() {
    this.init()
  },
}
</script>

<style>
.overlay-spell-anim-selector {
  position: absolute;
  bottom: 1px;
  left: 11px;
}
</style>
