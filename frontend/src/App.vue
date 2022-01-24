<template>
  <router-view></router-view>
</template>

<script>

import * as util  from "util";
import {App}      from "@/constants/app";
import {EventBus} from "@/app/event-bus/event-bus";
import {AppEnv}   from "@/app/env/app-env";

export default {
  name: "App",
  mounted() {
    var self = this
    window.addEventListener("keypress", function (e) {
      if (e.srcElement.tagName !== "BODY" && e.srcElement.tagName !== "A") {
        return
      }

      if (window.location.pathname === "/login") {
        return
      }

      switch (String.fromCharCode(e.keyCode)) {
        case 'h':
          EventBus.$emit('HIDE_NAVBAR', true);
          break
      }
    })

    this.loadWallpaper();

    // init app env / version
    AppEnv.init()
  },

  updated() {
    this.scrollToHashIfExists();
  },
  methods: {
    scrollToHashIfExists() {
      // for some reason initial load there are router things that are conflicting here
      // delay a little longer than the watcher
      setTimeout(() => this.scrollFix(this.$route.hash), 1000);
    },

    scrollFix: function (hashbang) {
      if (hashbang) {
        console.log(location.hash);

        location.hash = hashbang;
        const hashTarget = hashbang.replace("#", "");
        if (document.getElementById(hashTarget)) {
          console.log("scrolling to", hashTarget)
          document.getElementById(hashTarget).scrollIntoView();
        }
      }
    },

    loadWallpaper() {
      const backgrounds = [
        "champions-of-norrath-wallpaper.jpg",
        "faydark.png",
        "freeport.png",
        "lavastorm.png",
        "rivervale.png",
        "qeynos.png",
        "soldungb.png",
        "spire-wall.jpg",
        "void.jpeg",
      ]

      const background = util.format("%s%s",
        App.ASSET_WALLPAPER_URL,
        backgrounds[Math.floor(Math.random() * backgrounds.length)].trim()
      )

      let curImg = new Image();
      curImg.src = background;
      curImg.onload = function(){
        // do whatever here, add it to the background, append the image ect.
        document.body.style.setProperty("--image", "url(" + background + ")");
      }
    }
  },
  watch: {
    "$route"(to, from) {
      document.title = "[Spire] " + to.meta.title || "Spire"
    },
    "$route.hash": function () {
      setTimeout(() => this.scrollFix(this.$route.hash), 100);
    }
  }
}
</script>

<style lang="scss">

.card-slim {
  padding: 18px !important;
}

/*akkadius:experimental*/
body {
  /* background-repeat: no-repeat; */
  background:            -webkit-radial-gradient(circle, transparent 40%, var(--color-v) 95%), linear-gradient(to right, var(--color), var(--color)), var(--image) !important;
  background:            radial-gradient(circle, transparent 40%, var(--color-v) 95%), linear-gradient(to right, var(--color), var(--color)), var(--image);

  background-position:   center center;
  background-size:       cover;
  background-repeat:     no-repeat;
  background-attachment: fixed;
  background-blend-mode: normal, saturation, normal;
  /*--image:               url(~@/assets/img/eq-wallpaper-1.jpg);*/
  --color-v:             black;
  --color:               grey;
}


</style>
