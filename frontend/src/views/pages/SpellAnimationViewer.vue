<template>
  <div>
    <div class="container-fluid">
      <app-loader :is-loading="!loaded" padding="8"/>

      <eq-window
        title="Spell Animations"
        v-if="loaded"
        class="mt-5 text-center"
      >

        <input
          type="text"
          class="form-control ml-2 mb-4"
          v-model="search"
          v-on:keyup="triggerSearch"
          style="width: 95%"
          placeholder="Search for spell names to find animations">

        <div v-if="filteredAnimations && filteredAnimations.length === 0">
          No animations found...
        </div>

        <div v-for="(spell) in filteredAnimations" style="display:inline-block">
          <video
            muted
            loop
            :id="'spell-' + spell"
            :data-src="animBaseUrl + spell + '.mp4'"
            class="spell-preview">
          </video>
          <div class="overlay">
            <h6 class="eq-header">{{ spell }}</h6>
          </div>
        </div>

        <div class="mt-5">Videos courtesy of Georges <3</div>
      </eq-window>
    </div>
  </div>
</template>

<script>
import PageHeader from "@/views/layout/PageHeader";
import {App} from "@/constants/app";
import EqWindow from "@/components/eq-ui/EQWindow";
import SpellAnimations from "@/app/asset-maps/spell-animations-map.json";
import spellAnimMappings from "@/app/data-maps/spell-icon-anim-name-map.json";
import {Listeners} from "@/app/listeners/listeners";

let itemModels = [];

function handleRender() {
  let playing  = []
  let stopping = []
  let videos   = document.getElementsByClassName("spell-preview");
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

let renderEventListener    = null
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
  methods: {
    init: function () {
      // Preload model files
      let modelFiles = [];
      SpellAnimations[0].contents.forEach((row) => {
        const pieces      = row.name.split(/\//);
        const fileName    = pieces[pieces.length - 1].replace(".mp4", "");
        const animationId = parseInt(fileName)

        modelFiles.push(animationId)

        animationPreviewExists[animationId] = 1
      })

      console.log(animationPreviewExists)

      // Sort by spell animation number
      modelFiles.sort(function (a, b) {
        return a - b;
      });

      this.spellAnimations = modelFiles
      this.loaded          = true

      setTimeout(() => {
        handleRender()
      }, 500);
    },
    triggerSearch: debounce(function () {
      this.$router.push({ path: "/spell-animation-viewer", query: { q: this.search } }).catch(err => err)
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
        handleRender()
      }, 100);

    }
  },
  activated() {
    if (!this.$route.query.q) {
      this.search        = ""
      this.filteredRaces = []
    }

    // create route watcher
    this.routeWatcher = this.$watch('$route.query', () => {
      this.search = this.$route.query.q
      this.spellAnimSearch();
    });

    this.init()
    this.spellAnimSearch()

    // render scroll listener
    if (Listeners.SpellAnimViewerRenderListener) {
      window.removeEventListener("scroll", Listeners.SpellAnimViewerRenderListener)
    }

    Listeners.SpellAnimViewerRenderListener = debounce(handleRender, 100)
    window.addEventListener("scroll", Listeners.SpellAnimViewerRenderListener);
  },
  deactivated() {
    if (Listeners.SpellAnimViewerRenderListener) {
      console.log("Removing listener")
      window.removeEventListener("scroll", Listeners.SpellAnimViewerRenderListener, true)
      Listeners.SpellAnimViewerRenderListener = null
    }

    // remove route watcher
    this.routeWatcher()
  },
  beforeDestroy() {
    if (Listeners.SpellAnimViewerRenderListener) {
      console.log("Removing listener2")

      window.removeEventListener("scroll", Listeners.SpellAnimViewerRenderListener, true)
      Listeners.SpellAnimViewerRenderListener = null
    }
  }
}
</script>

<style>
.spell-preview {
  height:        250px;
  min-width:     150px;
  max-width:     200px;
  border-radius: 10px;
  margin:        3px;
}
</style>
