<template>
  <div :class="isComponent ? '' : 'container-fluid'">
    <app-loader :is-loading="!loaded" padding="8"/>

    <eq-window-simple
      title="Environment Previews"
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
          class="player-anim-preview"
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
import PageHeader      from "@/components/layout/PageHeader";
import {App}           from "@/constants/app";
import EqWindow        from "@/components/eq-ui/EQWindow";
import Previews from "@/app/asset-maps/player-animations.json";
import {Listeners}     from "@/app/listeners/listeners";
import {ROUTE}         from "../../routes";
import EqWindowSimple  from "../../components/eq-ui/EQWindowSimple";

let itemModels = [];

function handleRender() {
  let playing  = []
  let stopping = []
  let videos   = document.getElementsByClassName("player-anim-preview");
  for (let i = 0; i < videos.length; i++) {

    let video   = videos.item(i)
    let source  = document.createElement("source");
    let dataSrc = video.getAttribute("data-src")

    // Toggle playing
    if (elementInViewport(video)) {
      if (dataSrc) {

        // video.setAttribute("src", dataSrc);
        video.removeAttribute("data-src");
        video.pause()
        video.innerHTML = "";
        video.removeAttribute("src");

        source.setAttribute("src", dataSrc);
        source.setAttribute("type", "video/mp4");
        video.appendChild(source);
        video.load();
        video.play();
      }

      if (!videoPlaying(video) && videoLoaded(video)) {
        video.play()
        playing.push(video.getAttribute("id"))
      }
    } else {
      if (videoPlaying(video) && videoLoaded(video)) {
        video.pause()
        stopping.push(video.getAttribute("id"))
      }
    }
  }

  console.log("Playing", playing)
  console.log("Stopping", stopping)
}

function elementInViewport(el) {
  let top    = el.offsetTop;
  let left   = el.offsetLeft;
  let width  = el.offsetWidth;
  let height = el.offsetHeight;

  while (el.offsetParent) {
    el = el.offsetParent;
    top += el.offsetTop;
    left += el.offsetLeft;
  }

  return (
    top < (window.pageYOffset + window.innerHeight) &&
    left < (window.pageXOffset + window.innerWidth) &&
    (top + height) > window.pageYOffset &&
    (left + width) > window.pageXOffset
  );
}

function debounce(func, delay) {
  let debounceTimer;
  return function () {
    const context = this;
    const args    = arguments;
    clearTimeout(debounceTimer);
    debounceTimer = setTimeout(() => func.apply(context, args), delay);
  };
}

function videoPlaying(el) {
  return !!(el.currentTime > 0 && !el.paused && !el.ended && el.readyState > 2);
}

function videoLoaded(el) {
  return el.readyState === 4
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

      // render scroll listener
      if (Listeners.EmitterViewerRenderListener) {
        window.removeEventListener("scroll", Listeners.EmitterViewerRenderListener)
      }

      Listeners.EmitterViewerRenderListener = debounce(handleRender, 100)
      window.addEventListener("scroll", Listeners.EmitterViewerRenderListener);
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
        handleRender()
      }, 500);
    },
    triggerSearch: debounce(function () {
      this.$router.push(
        {
          path: ROUTE.EMITTER_VIEWER,
          query: {
            q: this.search
          }
        }
      ).catch(err => err)
    }, 1000),
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
        handleRender()
      }, 100);
    }
  },
  activated() {
    this.init()
  },
  deactivated() {
    if (Listeners.EmitterViewerRenderListener) {
      console.log("Removing listener")
      window.removeEventListener("scroll", Listeners.EmitterViewerRenderListener, true)
      Listeners.EmitterViewerRenderListener = null
    }

    // remove route watcher
    this.routeWatcher()
  },
  beforeDestroy() {
    if (Listeners.EmitterViewerRenderListener) {
      console.log("Removing listener2")

      window.removeEventListener("scroll", Listeners.EmitterViewerRenderListener, true)
      Listeners.EmitterViewerRenderListener = null
    }
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
