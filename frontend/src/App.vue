<template>
  <router-view></router-view>
</template>

<script>

import * as util                from "util";
import {App}                    from "@/constants/app";
import {EventBus}               from "@/app/event-bus/event-bus";
import {AppEnv}                 from "@/app/env/app-env";
import {LocalSettings, Setting} from "@/app/local-settings/localsettings";

export default {
  name: "App",
  mounted() {

    this.loadKeypressBindings();
    this.loadWallpaper();
    this.loadSpellIconSettings();

    // init app env / version
    AppEnv.init().then(() => {
      EventBus.$emit('APP_ENV_LOADED', true);
    })
  },

  created() {
    EventBus.$on("SPELL_LEGACY_ICONS_ENABLED", this.loadSpellIconSettings);
  },
  destroyed() {
    EventBus.$off("SPELL_LEGACY_ICONS_ENABLED", this.loadSpellIconSettings);
  },

  methods: {

    loadKeypressBindings() {
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
          case 'd':
            setTimeout(() => {
              LocalSettings.set(Setting.DEBUG_MODE, !LocalSettings.isDebugEnabled())
              App.DEBUG = !LocalSettings.isDebugEnabled()
              EventBus.$emit('DEBUG_UPDATED', true);
            }, 100)
            break
        }
      })
    },

    loadWallpaper() {
      const backgrounds = [
        "faydark.png",
        "freeport.png",
        "lavastorm.png",
        "rivervale.png",
        "oasis.png",
        "sebilis.png",
        "qeynos.png",
        "soldungb.png",
        "spire-wall.jpg",
      ]

      const background = util.format("%s%s",
        App.ASSET_WALLPAPER_URL,
        backgrounds[Math.floor(Math.random() * backgrounds.length)].trim()
      )

      let curImg    = new Image();
      curImg.src    = background;
      curImg.onload = function () {
        // do whatever here, add it to the background, append the image ect.
        document.body.style.setProperty("--image", "url(" + background + ")");
      }
    },

    loadSpellIconSettings() {
      for (let i of [12, 20, 30, 40]) {
        let e = document.getElementById("spell-icons-" + i)
        e.setAttribute('href', '/eq-asset-preview-master/assets/sprites/spell-icons-' + i + '.css')
        if (e && App.SPELL_LEGACY_ICONS_ENABLED) {
          e.setAttribute('href', '/eq-asset-preview-master/assets/sprites/spell-icons-legacy-' + i + '.css')
        }
      }
    }
  },
}
</script>

<style lang="scss">

.card-slim {
  padding: 18px !important;
}

/*akkadius:experimental*/
body {
  /* background-repeat: no-repeat; */
  background: -webkit-radial-gradient(circle, transparent 40%, var(--color-v) 95%), linear-gradient(to right, var(--color), var(--color)), var(--image) !important;
  background: radial-gradient(circle, transparent 40%, var(--color-v) 95%), linear-gradient(to right, var(--color), var(--color)), var(--image);

  background-position: center center;
  background-size: cover;
  background-repeat: no-repeat;
  background-attachment: fixed;
  background-blend-mode: normal, saturation, normal;
  /*--image:               url(~@/assets/img/eq-wallpaper-1.jpg);*/
  --color-v: black;
  --color: grey;
}


</style>
