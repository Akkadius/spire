<template>
  <div class="text-center">
    <app-loader :is-loading="!loaded" padding="8"/>

    <input
      type="text"
      class="form-control ml-2 mb-4"
      v-model="search"
      v-on:keyup="triggerSearch"
      style="width: 95%"
      placeholder="Search for spell names to find animations"
    >

    <div
      style="height: 85vh; overflow-y: scroll"
      v-on:scroll.passive="render"
      id="spell-video-view-port"
    >

      <div v-if="filteredAnimations && filteredAnimations.length === 0">
        No animations found...
      </div>

      <div
        v-for="(animationId) in filteredAnimations"
        :key="animationId"
        class="d-inline-block"
      >
        <video
          muted
          loop
          style="width: 160px; height: 230px; border-radius: 10px; border: 1px solid;"
          :id="'spell-' + animationId"
          :data-src="animBaseUrl + animationId + '.mp4'"
          @mousedown="selectSpellAnim(animationId)"
          :class="'spell-preview ' + classIsPulsating(animationId)"
        >
        </video>
      </div>
    </div>

  </div>
</template>

<script>
import PageHeader        from "@/components/layout/PageHeader";
import {App}             from "@/constants/app";
import EqWindow          from "@/components/eq-ui/EQWindow";
import SpellAnimations   from "@/app/asset-maps/spell-animations-map.json";
import spellAnimMappings from "@/app/data-maps/spell-icon-anim-name-map.json";
import * as util         from "util";

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

  // console.log("Playing", playing)
  // console.log("Stopping", stopping)
}

function elementInViewport(elem) {
  if (!(elem instanceof Element)) throw Error('DomUtil: elem is not an element.');
  const style = getComputedStyle(elem);
  if (style.display === 'none') return false;
  if (style.visibility !== 'visible') return false;
  if (style.opacity < 0.1) return false;
  if (elem.offsetWidth + elem.offsetHeight + elem.getBoundingClientRect().height +
    elem.getBoundingClientRect().width === 0) {
    return false;
  }
  const elemCenter = {
    x: elem.getBoundingClientRect().left + elem.offsetWidth / 2,
    y: elem.getBoundingClientRect().top + elem.offsetHeight / 2
  };
  if (elemCenter.x < 0) return false;
  if (elemCenter.x > (document.documentElement.clientWidth || window.innerWidth)) return false;
  if (elemCenter.y < 0) return false;
  if (elemCenter.y > (document.documentElement.clientHeight || window.innerHeight)) return false;
  let pointContainer = document.elementFromPoint(elemCenter.x, elemCenter.y);
  do {
    if (pointContainer === elem) return true;
  } while (pointContainer === pointContainer.parentNode);
  return false;
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

let animationPreviewExists = {}

export default {
  name: "SpellAnimationSelector",
  components: { EqWindow, PageHeader },
  data() {
    return {
      loaded: false,
      spellAnimations: [],
      filteredAnimations: [],
      search: "",
      animBaseUrl: App.ASSET_SPELL_ANIMATIONS,
    }
  },
  props: {
    selectedAnimation: {
      type: Number,
      default: 0,
      required: true
    },
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

      this.render()
      this.spellAnimSearch()

      // bring focus to the selected video
      if (this.selectedAnimation > 0) {
        // we need 100ms delay because the videos haven't been rendered yet
        setTimeout(() => {
          const container = document.getElementById("spell-video-view-port");
          const target    = document.getElementById(util.format("spell-%s", this.selectedAnimation))

          // 230 is height of video to offset
          if (container && target) {
            container.scrollTop = target.offsetTop - 230;
          }
        }, 100)
      }

    },
    render: function () {
      // Preload model files
      let modelFiles = [];
      SpellAnimations[0].contents.forEach((row) => {
        const pieces      = row.name.split(/\//);
        const fileName    = pieces[pieces.length - 1].replace(".mp4", "");
        const animationId = parseInt(fileName)

        modelFiles.push(animationId)

        animationPreviewExists[animationId] = 1
      })

      // Sort by spell animation number
      modelFiles.sort(function (a, b) {
        return a - b;
      });

      this.spellAnimations = modelFiles
      this.loaded          = true

      handleRender()
    },
    classIsPulsating(animation) {
      return animation === this.selectedAnimation ? 'pulsate' : ''
    },
    triggerSearch() {
      this.spellAnimSearch();
    },
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

    },
    selectSpellAnim(anim) {
      this.$emit("update:inputData", anim);
    }
  },
  activated() {
    this.init()
  },
}
</script>

<style>
.spell-preview {
  height: 250px;
  min-width: 150px;
  max-width: 200px;
  border-radius: 10px;
  margin: 3px;
}
</style>
