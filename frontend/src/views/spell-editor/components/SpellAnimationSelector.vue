<template>
  <div class="text-center">
    <app-loader :is-loading="!loaded" padding="8"/>

    <input
      type="text"
      class="form-control ml-2 mb-4"
      v-model="search"
      v-on:keyup="triggerSearch"
      style="width: 95%"
      placeholder="Search for spell names to find animations"
    >

    <div
      style="height: 85vh; overflow-y: scroll"
      v-on:scroll.passive="render"
      id="spell-video-view-port"
    >

      <div v-if="filteredAnimations && filteredAnimations.length === 0">
        No animations found...
      </div>

      <div
        v-for="(animationId) in filteredAnimations"
        :key="animationId"
        class="d-inline-block"
      >
        <video
          muted
          loop
          style="width: 160px; height: 230px; border-radius: 10px; border: 1px solid;"
          :id="'spell-' + animationId"
          :data-src="animBaseUrl + animationId + '.mp4'"
          @mousedown="selectSpellAnim(animationId)"
          :class="'video-preview spell-preview ' + classIsPulsating(animationId)"
        >
        </video>
      </div>
    </div>

  </div>
</template>

<script>
import PageHeader        from "@/components/layout/PageHeader";
import {App}             from "@/constants/app";
import EqWindow          from "@/components/eq-ui/EQWindow";
import SpellAnimations   from "@/app/eq-assets/spell-animations-map.json";
import spellAnimMappings from "@/app/data-maps/spell-icon-anim-name-map.json";
import * as util         from "util";
import VideoViewer       from "../../../app/video-viewer/video-viewer";

let animationPreviewExists = {}

export default {
  name: "SpellAnimationSelector",
  components: { EqWindow, PageHeader },
  data() {
    return {
      loaded: false,
      spellAnimations: [],
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
      if (!this.$route.query.q) {
        this.search        = ""
        this.filteredRaces = []
      }

      this.render()
      this.spellAnimSearch()

      // bring focus to the selected video
      if (this.selectedAnimation > 0) {
        // we need 100ms delay because the videos haven't been rendered yet
        setTimeout(() => {
          const container = document.getElementById("spell-video-view-port");
          const target    = document.getElementById(util.format("spell-%s", this.selectedAnimation))

          // 230 is height of video to offset
          if (container && target) {
            container.scrollTop = target.offsetTop - 80;
          }
        }, 100)
      }

    },
    render: function () {
      // Preload model files
      let modelFiles = [];
      SpellAnimations[0].contents.forEach((row) => {
        const pieces      = row.name.split(/\//);
        const fileName    = pieces[pieces.length - 1].replace(".mp4", "");
        const animationId = parseInt(fileName)

        modelFiles.push(animationId)

        animationPreviewExists[animationId] = 1
      })

      // Sort by spell animation number
      modelFiles.sort(function (a, b) {
        return a - b;
      });

      this.spellAnimations = modelFiles
      this.loaded          = true

      VideoViewer.handleRender()
    },
    classIsPulsating(animation) {
      return animation === this.selectedAnimation ? 'pulsate' : ''
    },
    triggerSearch() {
      this.spellAnimSearch();
    },
    spellAnimSearch: function () {
      this.loaded = false

      let foundAnim          = {};
      let filteredAnimations = []

      for (let spellAnimMapping of spellAnimMappings) {
        const spellName   = spellAnimMapping[0].toLowerCase().trim()
        const spellAnimId = spellAnimMapping[2]

        if (spellName.includes(this.search.toLowerCase())) {
          if (!foundAnim[spellAnimId] && animationPreviewExists[spellAnimId]) {
            filteredAnimations.push(spellAnimId)
            foundAnim[spellAnimId] = 1
          }
        }
      }

      // Sort by spell animation number
      filteredAnimations.sort(function (a, b) {
        return a - b;
      });

      this.filteredAnimations = filteredAnimations
      this.loaded             = true

      setTimeout(() => {
        VideoViewer.handleRender()
      }, 100);

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
.spell-preview {
  height: 250px;
  min-width: 150px;
  max-width: 200px;
  border-radius: 10px;
  margin: 3px;
}
</style>
