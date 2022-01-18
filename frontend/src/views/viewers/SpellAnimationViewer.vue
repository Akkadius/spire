<template>
  <div>
    <div :class="isComponent ? '' : 'container-fluid'">
      <app-loader :is-loading="!loaded" padding="8"/>

      <eq-window
        title="Spell Animations"
        v-if="loaded"
        class="mt-5 text-center"
      >

        <div class="row">
          <div class="col-12">
            <input
              type="text"
              class="form-control ml-2 mb-4"
              v-model="search"
              v-on:keyup="triggerSearch"
              placeholder="Search for spell names to find animations"
            >
          </div>
        </div>

        <div v-if="filteredAnimations && filteredAnimations.length === 0">
          No animations found...
        </div>

        <div class="row">
          <div class="col-12">
            <div v-for="(spell) in filteredAnimations" style="display:inline-block; position: relative;">
              <video
                muted
                loop
                :id="'spell-' + spell"
                :data-src="animBaseUrl + spell + '.mp4'"
                class="video-preview spell-preview-viewer"
              >
              </video>
              <div class="overlay">
                <h6 class="eq-header">{{ spell }}</h6>
              </div>
            </div>

          </div>
        </div>


        <div class="mt-3">Videos courtesy of DeadZergling <3</div>
      </eq-window>
    </div>
  </div>
</template>

<script>
import PageHeader        from "@/components/layout/PageHeader";
import {App}             from "@/constants/app";
import EqWindow          from "@/components/eq-ui/EQWindow";
import spellAnimMappings from "@/app/data-maps/spell-icon-anim-name-map.json";
import {ROUTE}           from "../../routes";
import VideoViewer       from "../../app/video-viewer/video-viewer";
import EqAssets          from "../../app/eq-assets/eq-assets";

let itemModels = [];

function debounce(func, delay) {
  let debounceTimer;
  return function () {
    const context = this;
    const args    = arguments;
    clearTimeout(debounceTimer);
    debounceTimer = setTimeout(() => func.apply(context, args), delay);
  };
}

let animationPreviewExists = {}

export default {
  components: { EqWindow, PageHeader },
  data() {
    return {
      loaded: false,
      spellAnimations: [],
      filteredAnimations: [],
      search: "",
      animBaseUrl: App.ASSET_SPELL_ANIMATIONS,
      routeWatcher: null,
    }
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

      // create route watcher
      this.routeWatcher = this.$watch('$route.query', () => {
        if (this.$route.query.q && this.$route.query.q !== "") {
          this.search = this.$route.query.q
        }
        this.spellAnimSearch();
      });

      this.render()
      this.spellAnimSearch()

      // hook video viewer scroll listener
      VideoViewer.addScrollListener()
    },
    render: function () {
      EqAssets.getSpellAnimationFileIds().forEach((animationId) => {
        animationPreviewExists[animationId] = 1
      })

      this.spellAnimations = EqAssets.getSpellAnimationFileIds()
      this.loaded          = true

      setTimeout(() => {
        VideoViewer.handleRender()
      }, 500);
    },
    triggerSearch: debounce(function () {
      this.$router.push({ path: ROUTE.SPELL_ANIMATION_VIEWER, query: { q: this.search } }).catch(err => err)
    }, 1000),
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

    }
  },
  activated() {
    this.init()
  },
  deactivated() {
    VideoViewer.destroyScrollListener()

    // remove route watcher
    this.routeWatcher()
  },
  props: {
    isComponent: { // here for now because this viewer wasn't built as a component in mind
      default: false,
      required: false,
      type: Boolean,
    },
  }
}
</script>

<style>
.spell-preview-viewer {
  /*height: 250px;*/
  /*min-width: 150px;*/
  /*max-width: 200px;*/

  height: 162px;
  width: 288px;

  border-radius: 10px;
  margin-right: 5px;
}

.overlay {
  position: absolute;
  bottom: 2px;
  left: 9px;
}
</style>
