<template>
  <div>
    <eq-window
      title="Player Animations"
      v-if="loaded"
      class="text-center"
    >
      <div v-if="filteredPreviews && filteredPreviews.length === 0">
        No previews found...
      </div>

      <div
        class="row justify-content-center"
        v-on:scroll.passive="videoRender"
        style="height: 90vh; overflow-y: scroll;"
      >
        <div
          v-for="(preview) in filteredPreviews"
          class="col-sm-12 col-xs-12 col-md-12 col-lg-6 col-xl-4"
        >
          <div style="position: relative; width: 100%;">
            <video
              muted
              loop
              :data-video-id="preview"
              :id="'preview-' + preview"
              style="border-radius: 5px; height: auto; width: 100%; background-color: black"
              :data-src="previewBaseUrl + preview + '.mp4'"
              class="video-preview"
            >
            </video>
            <div
              :id="'overlay-' + preview"
              class="fade-in"
              style="position: absolute; bottom: 2px; width: 100%; display: none"
            >
              <h6 class="eq-header">{{ preview }}</h6>
            </div>
          </div>
        </div>

        <div class="col-12 mt-3 d-block">Videos Credits @DeadZergling</div>
      </div>

    </eq-window>
  </div>
</template>

<script>
import PageHeader     from "@/components/layout/PageHeader";
import {App}          from "@/constants/app";
import EqWindow       from "@/components/eq-ui/EQWindow";
import EqWindowSimple from "../../components/eq-ui/EQWindowSimple";
import VideoViewer    from "../../app/video-viewer/video-viewer";
import EqAssets       from "../../app/eq-assets/eq-assets";
import ContentArea    from "../../components/layout/ContentArea";

export default {
  components: { ContentArea, EqWindowSimple, EqWindow, PageHeader },
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
    async init() {
      if (!this.$route.query.q) {
        this.search        = ""
        this.filteredRaces = []
      }

      // create route watcher
      this.routeWatcher = this.$watch('$route.query', () => {
        this.search = this.$route.query.q
        this.previewAnimSearch();
      });

      await this.render()
      this.previewAnimSearch()

      // hook video viewer scroll listener
      VideoViewer.addScrollListener()
    },

    videoRender() {
      VideoViewer.handleRender();
    },

    render: async function () {
      this.previews = await EqAssets.getPlayerAnimationFileIds()
      this.loaded   = true

      setTimeout(() => {
        VideoViewer.handleRender()
      }, 500);
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
