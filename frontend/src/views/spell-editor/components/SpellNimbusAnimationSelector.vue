<template>
  <div>
    <app-loader :is-loading="!loaded" padding="8"/>

    <eq-window-simple
      title="Nimbus Selector"
      class="text-center p-0"
      v-if="loaded"
    >
      <div
        style="height: 90vh; overflow-y: scroll"
        v-on:scroll="handleRender"
        id="spell-video-view-port"
      >
        <div class="col-12 mt-3">
          <div v-if="filteredAnimations && filteredAnimations.length === 0">
            No animations found...
          </div>

          <div
            v-for="(animationId) in filteredAnimations"
            :key="animationId"
            style="display:inline-block; position: relative; "
            class="d-inline-block"
          >
            <video
              muted
              loop
              style="height: 146px; width: 259px; border-radius: 5px; border: 1px solid rgba(255, 255, 255, .3); background-color: black;"
              :id="'spell-' + animationId"
              :data-src="animBaseUrl + animationId + '.mp4#t=3'"
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
import PageHeader       from "@/components/layout/PageHeader";
import {App}            from "@/constants/app";
import EqWindow         from "@/components/eq-ui/EQWindow";
import * as util        from "util";
import VideoViewer      from "@/app/video-viewer/video-viewer";
import EqWindowSimple   from "@/components/eq-ui/EQWindowSimple";
import EqAssets         from "@/app/eq-assets/eq-assets";
import {SPELL_NIMBUSES} from "@/app/constants/eq-spell-constants";

export default {
  name: "SpellNimbusAnimationSelector",
  components: { EqWindowSimple, EqWindow, PageHeader },
  data() {
    return {
      loaded: false,
      filteredAnimations: [],
      search: "",
      animBaseUrl: App.ASSET_SPELL_ANIMATIONS,
    }
  },
  props: {
    selectedAnimation: {
      type: Number,
      default: 0,
      required: true
    },
  },
  created() {
    this.init()
  },
  methods: {
    init() {
      console.time("[SpellNimbusAnimationSelector] init");

      this.render()

      // bring focus to the selected video
      if (this.selectedAnimation > 0) {
        // we need 100ms delay because the videos haven't been rendered yet
        setTimeout(() => {
          console.time("[SpellAnimationSelector] scrollTo");
          const container = document.getElementById("spell-video-view-port");
          const target    = document.getElementById(util.format("spell-%s", this.selectedAnimation))

          // 230 is height of video to offset
          if (container && target) {
            const top = target.getBoundingClientRect().top

            container.scrollTop = container.scrollTop + top - 250;
            VideoViewer.handleRender();
          }

          console.timeEnd("[SpellAnimationSelector] scrollTo");
        }, 1)
      }

      console.timeEnd("[SpellAnimationSelector] init");
    },
    handleRender() {
      VideoViewer.handleRender()
    },
    render: function () {
      // Preload model files
      let animations = [];
      EqAssets.getSpellAnimationFileIds().forEach((animationId) => {
        if (SPELL_NIMBUSES.includes(animationId)) {
          animations.push(animationId)
        }
      })

      // Sort by spell animation number
      animations.sort(function (a, b) {
        return a - b;
      });

      this.filteredAnimations = animations
      this.loaded             = true

      VideoViewer.handleRender()
    },
    classIsPulsating(animation) {
      return animation === this.selectedAnimation ? 'pulsate' : ''
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
