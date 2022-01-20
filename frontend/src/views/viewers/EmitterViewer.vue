<template>
  <div :class="isComponent ? '' : 'container-fluid'" v-if="loaded">
    <app-loader :is-loading="!loaded" padding="8"/>

    <eq-window-simple
      title="Emitter Viewer"
      class="mt-3 text-center"
    >
      <div v-if="filteredPreviews && filteredPreviews.length === 0">
        No previews found...
      </div>

      <div
        v-on:scroll.passive="videoRender"
        style="height: 90vh; overflow-y: scroll;"
        class="row"
      >
        <div class="col-12">
          <div
            v-for="(preview) in filteredPreviews"
            style="display:inline-block; position: relative;"
          >
            <video
              muted
              loop
              :id="'preview-' + preview"
              :data-src="animBaseUrl + preview + '.mp4'"
              class="video-preview emitter-preview"
            >
            </video>
            <div class="emitter-overlay">
              <h6 class="eq-header">{{ preview }}</h6>
            </div>
          </div>
        </div>

        <div class="col-12 mt-3">Videos Credits @DeadZergling</div>
      </div>

    </eq-window-simple>
  </div>
</template>

<script>
import PageHeader     from "@/components/layout/PageHeader";
import {App}          from "@/constants/app";
import EqWindow       from "@/components/eq-ui/EQWindow";
import EqWindowSimple from "../../components/eq-ui/EQWindowSimple";
import VideoViewer    from "../../app/video-viewer/video-viewer";
import EqAssets       from "../../app/eq-assets/eq-assets";

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
    videoRender() {
      VideoViewer.handleRender();
    },
    render: function () {
      this.previews = EqAssets.getEmitterPreviewFileIds()
      this.loaded   = true

      setTimeout(() => {
        VideoViewer.handleRender()
      }, 100);
    },
    previewAnimSearch: function () {
      this.loaded           = false
      this.filteredPreviews = this.previews
      this.loaded           = true

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

  /*height: 162px;*/
  /*width: 288px;*/

  height: 262px;
  width: 464px;

  /*height: 135px;*/
  /*width: 240px;*/

  border-radius: 5px !important;
  margin: 1px;
  margin-right: 10px;
}

.emitter-overlay {
  position: absolute;
  bottom: 2px;
  left: 9px;
}
</style>
