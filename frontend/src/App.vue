<template>
  <div>
    <ninja-keys
      class="dark ninja-icon"
      ref="ninjaKeys"
      placeholder="Where would you like to go?"
    />
    <keypress-commands-modal/>
    <router-view></router-view>
  </div>
</template>

<script>

import "ninja-keys";
import * as util                from "util";
import {App}                    from "@/constants/app";
import {EventBus}               from "@/app/event-bus/event-bus";
import {AppEnv}                 from "@/app/env/app-env";
import {LocalSettings, Setting} from "@/app/local-settings/localsettings";
import {ROUTE}                  from "@/routes";
import UserContext              from "@/app/user/UserContext";
import KeypressCommandsModal    from "@/components/modals/KeypressCommandsModal.vue";

export default {
  name: "App",
  components: { KeypressCommandsModal },
  async beforeMount() {
    await AppEnv.init()
  },

  async mounted() {

    this.loadKeypressBindings();
    this.loadWallpaper();
    this.loadSpellIconSettings();

    this.user = await UserContext.getUser()
    if (typeof AppEnv.getOS() === "undefined") {
      await AppEnv.init()
    }

    EventBus.$emit('APP_ENV_LOADED', true);
    AppEnv.routeCheckOcculus(this.$route, this.$router)
    AppEnv.routeCheckSpireInitialized(this.$route, this.$router)
    this.checkIfUserNeedsToAuth()

    setTimeout(() => {
      AppEnv.routeCheckSpireInitialized(this.$route, this.$router)
    }, 1)
  },

  created() {
    EventBus.$on("SPELL_LEGACY_ICONS_ENABLED", this.loadSpellIconSettings);
  },
  destroyed() {
    EventBus.$off("SPELL_LEGACY_ICONS_ENABLED", this.loadSpellIconSettings);
  },

  methods: {

    async checkIfUserNeedsToAuth() {
      if (AppEnv.isSpireInitialized() &&
        AppEnv.isLocalAuthEnabled() &&
        UserContext.getAccessToken().length === 0 &&
        this.$route.fullPath !== ROUTE.LOGIN) {
        await this.$router.push(ROUTE.LOGIN).catch((e) => {
        })
      }
    },

    loadKeypressBindings() {
      let controlPressed = false;

      document.onkeydown = (e) => {
        const tagName = e.srcElement.tagName

        e = e || window.event; //Get event
        if (!e.ctrlKey && e.key !== "Control") return;
        let code = e.which || e.keyCode; //Get key code

        // if we only press the control key, reject further action if we are inside
        // things like inputs
        if (!["BODY", "A", "NINJA-KEYS", "DIV"].includes(tagName) && code === 17) {
          return
        }

        // if command is accompanied by another key, dismiss commands modal
        if (code !== 17) {
          this.$bvModal.hide('keypress-commands-modal')
        }

        switch (code) {
          case 17: // just command
            if (controlPressed) {
              this.$bvModal.hide('keypress-commands-modal')
              return
            }

            // don't show modal help window if search box is open
            const n = document.querySelector('ninja-keys');
            if (n.visible) {
              break;
            }
            this.$bvModal.show('keypress-commands-modal');
            break;
          case 191: // Ctrl + /
          case 75: // Ctrl + K
            e.preventDefault();

            const ninja = document.querySelector('ninja-keys');
            setTimeout(() => {
              ninja.open()
            }, 1)

            e.stopPropagation();

            break;
          case 87://Block Ctrl+W
          case 83://Block Ctrl+S
            e.preventDefault();
            // e.stopPropagation();
            break;
        }

        controlPressed = true;
      };

      document.onkeyup = (e) => {
        e = e || window.event; //Get event
        if (!e.ctrlKey && e.key !== "Control") return;
        let code = e.which || e.keyCode; //Get key code

        setTimeout(() => {
          if (document.getElementById('keypress-commands-modal')) {
            this.$bvModal.hide('keypress-commands-modal')
          }
          controlPressed = false;
        }, 1)
      };

      window.onblur  = () => {
        this.$bvModal.hide('keypress-commands-modal')
      }
      window.onfocus = () => {
        this.$bvModal.hide('keypress-commands-modal')
      }

      window.onkeyup = (e) => {
        e = e || window.event; //Get event
        if (!e.ctrlKey && e.key !== "Control") return;

        this.$bvModal.hide('keypress-commands-modal')
      };

      window.addEventListener("keypress", (e) => {
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
          case 'b':
            this.$router.push({ path: '/break' })
            break
          case '`':
            this.$router.push({ path: '/' })
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
