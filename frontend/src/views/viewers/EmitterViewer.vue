<template>
  <div :class="isComponent ? '' : 'container-fluid'">
    <app-loader :is-loading="!loaded" padding="8"/>

    <eq-window-simple
      title="Environment Emitters"
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
          :data-src="animBaseUrl + preview + '.mp4'"
          class="video-preview emitter-preview"
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
import Emitters       from "@/app/asset-maps/emitters.json";
import EqWindowSimple from "../../components/eq-ui/EQWindowSimple";
import VideoViewer    from "../../app/video-viewer/video-viewer";

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


let renderEventListener = null
let previewExists       = {}

export default {
  components: { EqWindowSimple, EqWindow, PageHeader },
  data() {
    return {
      loaded: false,
      previews: [],
      filteredPreviews: [],
      search: "",
      animBaseUrl: App.ASSET_EMITTER_CLIPS,
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
      Emitters[0].contents.forEach((row) => {
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
      }, 100);
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
.emitter-preview {
  /*height: 270px;*/
  /*width: 480px;*/

  /*height: 180px;*/
  /*width: 320px;*/

  height: 135px;
  width: 240px;

  border-radius: 10px;
  margin: 1px;
}

.overlay {
  position: absolute;
  bottom: 2px;
  left: 9px;
}
</style>
