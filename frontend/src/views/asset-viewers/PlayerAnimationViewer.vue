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
        class="row "
        v-on:scroll.passive="videoRender"
        style="height: 90vh; overflow-y: scroll;"
      >
        <div class="col-12">
      <div v-for="(preview) in filteredPreviews" style="display:inline-block; position: relative;">
        <video
          muted
          loop
          :id="'preview-' + preview"
          style="background-color: black;"
          :data-src="previewBaseUrl + preview + '.mp4'"
          class="video-preview player-anim-preview"
        >
        </video>
        <div class="overlay">
          <h6 class="eq-header">{{ preview }}</h6>
        </div>
      </div>
      <div class="mt-3">Videos Credits @DeadZergling</div>
        </div></div>
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

<style>
.player-anim-preview {
  /*height: 270px;*/
  /*width: 480px;*/

  /*height: 180px;*/
  /*width: 320px;*/

  /*height: 144px;*/
  /*width: 256px;*/

  /*height: 135px;*/
  /*width: 240px;*/

  height: 25vh;
  width: 44vh;

  border-radius: 5px;
  margin-right: 10px;
}

.overlay {
  position: absolute;
  bottom: 6px;
  left: 9px;
}
</style>
