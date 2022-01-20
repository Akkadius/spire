<template>
  <div>
    <div :class="isComponent ? '' : 'container-fluid'">
      <app-loader :is-loading="!loaded" padding="8"/>

      <eq-window-simple title="Spell Animations" style="margin-bottom: 1px">
        <div class="row">
          <div class="col-12">
            <input
              type="text"
              class="form-control ml-2"
              v-model="search"
              v-on:keyup="triggerSearch"
              @enter="triggerSearch"
              placeholder="Search for spell names to find animations"
            >
          </div>
        </div>
      </eq-window-simple>

      <eq-window-simple
        v-if="loaded"
        class="text-center mt-3"
      >
        <div v-if="filteredAnimations && filteredAnimations.length === 0">
          No animations found...
        </div>

        <div
          class="row "
          v-on:scroll.passive="videoRender"
          style="height: 79vh; overflow-y: scroll;"
        >
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

          <div class="col-12 mt-3">Videos Credits @DeadZergling</div>
        </div>
      </eq-window-simple>
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
import EqWindowSimple    from "../../components/eq-ui/EQWindowSimple";

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
  components: { EqWindowSimple, EqWindow, PageHeader },
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
    videoRender() {
      VideoViewer.handleRender();
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
      console.log("trigger")

      this.loaded = false

      let foundAnim          = {};
      let filteredAnimations = EqAssets.getSpellAnimationFileIds()

      if (this.search !== "") {
        filteredAnimations = []
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

  /*height: 262px;*/
  /*width: 464px;*/

  height: 25vh;
  width: 44vh;

  border-radius: 5px !important;
  margin-right: 10px;
}

.overlay {
  position: absolute;
  bottom: 2px;
  left: 9px;
}
</style>
