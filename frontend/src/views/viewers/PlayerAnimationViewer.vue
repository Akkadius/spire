<template>
  <div :class="isComponent ? '' : 'container-fluid'">
    <app-loader :is-loading="!loaded" padding="8"/>

    <eq-window-simple
      title="Player Animation Previews"
      v-if="loaded"
      class="mt-4 text-center"
    >
      <div v-if="filteredPreviews && filteredPreviews.length === 0">
        No previews found...
      </div>

      <div v-for="(preview) in filteredPreviews" style="display:inline-block; position: relative;">
        <video
          muted
          loop
          :id="'preview-' + preview"
          :data-src="previewBaseUrl + preview + '.mp4'"
          class="video-preview player-anim-preview"
        >
        </video>
        <div class="overlay">
          <h6 class="eq-header">{{ preview }}</h6>
        </div>
      </div>
      <div class="mt-3">Videos courtesy of DeadZergling <3</div>
    </eq-window-simple>
  </div>
</template>

<script>
import PageHeader     from "@/components/layout/PageHeader";
import {App}          from "@/constants/app";
import EqWindow       from "@/components/eq-ui/EQWindow";
import Previews       from "@/app/asset-maps/player-animations.json";
import EqWindowSimple from "../../components/eq-ui/EQWindowSimple";
import VideoViewer    from "../../app/video-viewer/video-viewer";

let itemModels = [];
let previewExists = {}

export default {
  components: { EqWindowSimple, EqWindow, PageHeader },
  data() {
    return {
      loaded: false,
      previews: [],
      filteredPreviews: [],
      search: "",
      previewBaseUrl: App.ASSET_PLAYER_ANIMATION_CLIPS,
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
        this.search = this.$route.query.q
        this.previewAnimSearch();
      });

      this.render()
      this.previewAnimSearch()

      // hook video viewer scroll listener
      VideoViewer.addScrollListener()
    },

    render: function () {
      // Preload model files
      let modelFiles = [];
      Previews[0].contents.forEach((row) => {
        const pieces      = row.name.split(/\//);
        const fileName    = pieces[pieces.length - 1].replace(".mp4", "");
        const animationId = parseInt(fileName)

        modelFiles.push(animationId)

        previewExists[animationId] = 1
      })

      // Sort by preview animation number
      modelFiles.sort(function (a, b) {
        return a - b;
      });

      this.previews = modelFiles
      this.loaded   = true

      setTimeout(() => {
        VideoViewer.handleRender()
      }, 500);
    },
    previewAnimSearch: function () {
      this.loaded          = false
      let filteredPreviews = []

      // Sort by preview animation number
      filteredPreviews.sort(function (a, b) {
        return a - b;
      });

      if (filteredPreviews.length > 0) {
        this.filteredPreviews = filteredPreviews
      } else {
        this.filteredPreviews = this.previews
      }

      this.loaded = true

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
.player-anim-preview {
  height: 270px;
  width: 480px;
  border-radius: 10px;
  margin: 3px;
}

.overlay {
  position: absolute;
  bottom: 2px;
  left: 9px;
}
</style>
