<template>
  <div>
    <eq-window title="Spell Animations" style="margin-bottom: 1px">
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

      <div class="row">
        <div class="col-12 text-center mt-3">

          <div class="btn-group ml-3" role="group" aria-label="Basic example">
            <b-button
              @click="filterNimbuses = true; triggerSearch()"
              size="sm"
              :variant="(filterNimbuses ? 'warning' : 'outline-warning')"
            >Nimbuses
            </b-button>
            <b-button
              @click="reset(); triggerSearch()"
              size="sm"
              :variant="(!filterNimbuses ? 'warning' : 'outline-warning')"
            >All
            </b-button>
          </div>

          <div class="btn-group ml-3" role="group" aria-label="Basic example">
            <b-button size="sm" variant="outline-warning"><i class="fa fa-clock-o"></i> Start Preview @</b-button>
            <b-button
              @click="startVideoTime = 0; triggerSearch()"
              size="sm"
              :variant="(parseInt(startVideoTime) === 0 ? 'warning' : 'outline-warning')"
            >0s
            </b-button>
            <b-button
              @click="startVideoTime = 3; triggerSearch()"
              size="sm"
              :variant="(parseInt(startVideoTime) === 3 ? 'warning' : 'outline-warning')"
            >3s
            </b-button>
          </div>

          <b-button
            @click="reset(); triggerSearch()"
            size="sm"
            class="ml-3"
            variant="outline-warning"
          ><i class="fa fa-refresh"></i> Reset
          </b-button>

        </div>
      </div>
    </eq-window>

    <eq-window-simple
      v-if="loaded"
      class="text-center mt-3"
    >
      <div v-if="filteredAnimations && filteredAnimations.length === 0">
        No animations found...
      </div>

      <div
        class="row justify-content-center"
        v-on:scroll.passive="videoRender"
        style="height: 74vh; overflow-y: scroll; overflow-x: hidden; box-sizing: border-box;"
      >
        <div
          v-for="(spell) in filteredAnimations"
          :key="spell"
          class="col-sm-12 col-xs-12 col-md-12 col-lg-6 col-xl-4"
        >
          <div style="position: relative; width: 100%;">
            <video
              muted
              loop
              :id="'spell-' + spell"
              :data-video-id="spell"
              :data-src="animBaseUrl + spell + '.mp4#t=' + startVideoTime"
              class="video-preview spell-preview-video"
              style="border-radius: 5px; height: auto; width: 100%; background-color: black"
            >
            </video>
            <div
              :id="'overlay-' + spell"
              class="fade-in"
              style="position: absolute; bottom: 2px; width: 100%; display: none"
            >
              <h6 class="eq-header">{{ spell }}</h6>
            </div>
          </div>
        </div>

        <div class="col-12 mt-3">Videos Credits @DeadZergling</div>
      </div>
    </eq-window-simple>
  </div>
</template>

<script>
import PageHeader       from "@/components/layout/PageHeader";
import {App}            from "@/constants/app";
import EqWindow         from "@/components/eq-ui/EQWindow";
import {ROUTE}          from "../../routes";
import VideoViewer      from "../../app/video-viewer/video-viewer";
import EqAssets         from "../../app/eq-assets/eq-assets";
import EqWindowSimple   from "../../components/eq-ui/EQWindowSimple";
import {SPELL_NIMBUSES} from "../../app/constants/eq-spell-constants";
import ContentArea      from "../../components/layout/ContentArea";

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
  components: { ContentArea, EqWindowSimple, EqWindow, PageHeader },
  data() {
    return {
      loaded: false,
      spellAnimations: [],
      filteredAnimations: [],
      search: "",
      animBaseUrl: App.ASSET_SPELL_ANIMATIONS,
      // routeWatcher: null,
      filterNimbuses: false,

      SPELL_NIMBUSES: SPELL_NIMBUSES,

      startVideoTime: 3,
    }
  },

  watch: {
    // reset state vars when we navigate away
    '$route'() {
      this.loadQueryState()
      this.spellAnimSearch()
    },
  },

  created() {
    this.loadQueryState()

    this.init()
  },
  methods: {

    reset() {
      this.startVideoTime = 0;
      this.filterNimbuses = false;
    },

    loadQueryState: function () {
      if (this.$route.query.q && this.$route.query.q !== "") {
        this.search = this.$route.query.q
      }
      if (this.$route.query.nimbus && this.$route.query.nimbus === "true") {
        this.filterNimbuses = true
      }
      if (this.$route.query.startVideoTime && this.$route.query.nimbus !== 0) {
        this.startVideoTime = this.$route.query.startVideoTime
      }
    },

    updateQueryState: function () {
      let queryState = {};
      if (this.search !== "") {
        queryState.q = this.search
      }
      if (this.filterNimbuses) {
        queryState.nimbus = this.filterNimbuses
      }
      if (this.startVideoTime) {
        queryState.startVideoTime = this.startVideoTime
      }

      this.$router.push(
        {
          path: ROUTE.SPELL_ANIMATION_VIEWER,
          query: queryState
        }
      ).catch(() => {
      })
    },

    async init() {
      if (!this.$route.query.q) {
        this.search        = ""
        this.filteredRaces = []
      }

      await this.render()
      this.spellAnimSearch()

      // hook video viewer scroll listener
      VideoViewer.addScrollListener()
    },
    videoRender() {
      VideoViewer.handleRender();
    },
    render: async function () {
      const r = await EqAssets.getSpellAnimationFileIds()
      r.forEach((animationId) => {
        animationPreviewExists[animationId] = 1
      })

      this.spellAnimations = r
      this.loaded          = true

      setTimeout(() => {
        VideoViewer.handleRender()
      }, 500);
    },
    triggerSearch: debounce(function () {
      this.updateQueryState()
    }, 1000),
    spellAnimSearch: async function () {
      this.loaded = false

      let foundAnim          = {};
      let filteredAnimations = await EqAssets.getSpellAnimationFileIds()

      if (this.search !== "") {
        filteredAnimations      = []
        const spellAnimMappings = await EqAssets.getSpellAnimNameMappings()
        for (let spellAnimMapping of spellAnimMappings) {
          const spellName   = spellAnimMapping[0].toLowerCase().trim()
          const spellAnimId = spellAnimMapping[2]

          // spell name search filter
          if (spellName.includes(this.search.toLowerCase())) {
            if (!foundAnim[spellAnimId] && animationPreviewExists[spellAnimId]) {
              filteredAnimations.push(spellAnimId)
              foundAnim[spellAnimId] = 1
            }
          }
        }
      }

      // filter on nimbuses if filter is set
      if (this.filterNimbuses) {
        filteredAnimations = []

        this.spellAnimations.forEach((animationId) => {
          if (SPELL_NIMBUSES.includes(animationId)) {
            filteredAnimations.push(animationId)
            foundAnim[animationId] = 1
          }
        })
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
  beforeDestroy() {
    VideoViewer.destroyScrollListener()
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
