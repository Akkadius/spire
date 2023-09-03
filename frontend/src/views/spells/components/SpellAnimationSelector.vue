<template>
  <div>
    <eq-window-simple style="width: 100%">
      <div class="row">
        <div class="col-12">
          <input
            type="text"
            class="form-control"
            v-model="search"
            v-on:keyup="triggerSearch"
            placeholder="Search for spell names to find animations"
          >
        </div>
      </div>

    </eq-window-simple>

    <eq-window-simple
      class="text-center p-0"
    >
      <div
        style="height: 80vh; overflow-y: scroll"
        v-on:scroll="handleRender"
        id="spell-video-view-port"
      >
        <div class="col-12">
          <div v-if="filteredAnimations && filteredAnimations.length === 0">
            No animations found...
          </div>

          <div
            v-for="(animationId) in filteredAnimations"
            :key="animationId"
            style="display:inline-block; position: relative;"
            class="col-sm-12 col-xs-12 col-md-12 col-lg-6 col-xl-6 d-inline-block"
          >
            <video
              muted
              loop
              style="height: auto; width: 100%; border-radius: 5px; border: 1px solid rgba(255, 255, 255, .3); background-color: black;"
              :id="'spell-' + animationId"
              :data-video-id="animationId"
              :data-src="animBaseUrl + animationId + '.mp4#t=3'"
              @mousedown="selectSpellAnim(animationId)"
              :class="'video-preview ' + classIsPulsating(animationId)"
            >
            </video>

            <div
              :id="'overlay-' + animationId"
              class="fade-in"
              style="position: absolute; bottom: 2px; width: 100%; display: none"
            >
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
import VideoViewer    from "@/app/video-viewer/video-viewer";
import EqWindowSimple from "@/components/eq-ui/EQWindowSimple";
import EqAssets       from "@/app/eq-assets/eq-assets";

let animationPreviewExists = {}

export default {
  name: "SpellAnimationSelector",
  components: { EqWindowSimple, EqWindow, PageHeader },
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
    async init() {
      console.time("[SpellAnimationSelector] init");
      if (!this.$route.query.q) {
        this.search        = ""
        this.filteredRaces = []
      }

      await this.render()
      this.spellAnimSearch()

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
        }, 100)
      }

      console.timeEnd("[SpellAnimationSelector] init");
    },
    handleRender() {
      VideoViewer.handleRender()
    },
    render: async function () {
      // Preload model files
      let modelFiles = [];
      const r        = await EqAssets.getSpellAnimationFileIds()
      r.forEach((animationId) => {
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
    spellAnimSearch: async function () {
      this.loaded = false

      let foundAnim          = {};
      let filteredAnimations = []

      const spellAnimMappings = await EqAssets.getSpellAnimNameMappings()
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
}
</script>
