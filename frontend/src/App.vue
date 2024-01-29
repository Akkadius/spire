<template>
  <div>
    <ninja-keys
      class="dark ninja-icon"
      ref="ninjaKeys"
      placeholder="Where would you like to go?"
    />
    <keypress-commands-modal/>
    <router-view></router-view>
    <app-update-modal :release="release" :current-version="currentVersion"/>
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
import semver                   from "semver";
import AppUpdateModal           from "@/components/modals/AppUpdateModal.vue";

export default {
  name: "App",
  components: { AppUpdateModal, KeypressCommandsModal },
  async beforeMount() {
    await AppEnv.init()
  },

  data() {
    return {
      release: {},
      currentVersion: "",
    }
  },

  async mounted() {
    this.loadKeypressBindings();
    this.loadWallpaper();
    this.loadSpellIconSettings();

    if (typeof AppEnv.getOS() === "undefined") {
      await AppEnv.init()
    }

    EventBus.$emit('APP_ENV_LOADED', true);

    this.$router.onReady(async () => {
      if (!this.$route.fullPath.includes(ROUTE.LOGIN)) {
        console.log("login route, skipping auth check")
        this.user = await UserContext.getUser()
        await this.checkIfUserNeedsToAuth()
      }

      AppEnv.routeCheckOcculus(this.$route, this.$router)
      AppEnv.routeCheckSpireInitialized(this.$route, this.$router)
    })

    setTimeout(() => {
      AppEnv.routeCheckSpireInitialized(this.$route, this.$router)
    }, 1)

    this.checkForSpireUpdate()
  },

  created() {
    EventBus.$on("SPELL_LEGACY_ICONS_ENABLED", this.loadSpellIconSettings);
    EventBus.$on("CHECK_SPIRE_UPDATE", this.checkSpireUpdate);
  },
  destroyed() {
    EventBus.$off("SPELL_LEGACY_ICONS_ENABLED", this.loadSpellIconSettings);
    EventBus.$off("CHECK_SPIRE_UPDATE", this.checkSpireUpdate);
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
    },

    checkSpireUpdate() {
      this.checkForSpireUpdate(true)
    },

    checkForSpireUpdate(force = false) {
      if (!AppEnv.isAppLocal()) {
        console.log("skipping update check, not local app")
        return
      }

      const current       = AppEnv.getVersion()
      this.currentVersion = current
      const last_checked  = LocalSettings.getLastCheckedUpdateTime()
      const now           = new Date().getTime() / 1000

      // check if we've checked in the last 1 hour
      if (now - last_checked < 3600 && !force) {
        console.log("skipping update check, checked in last hour")
        this.release = JSON.parse(LocalSettings.getLatestReleasePayload())
        return
      }

      let latest = "0.0.0";

      // fetch from github releases akkadius/spire
      const url = 'https://api.github.com/repos/akkadius/spire/releases/latest'
      fetch(url)
        .then(response => response.json())
        .then(data => {
          latest       = data.tag_name.replace("v", "")
          this.release = data
          const ignoredUpdateVersion = LocalSettings.getIgnoredUpdateVersion()

          if (semver.gt(latest, current)) {
            console.log("update available")
            if (ignoredUpdateVersion !== latest) {
              this.$bvModal.show('app-update-modal')
            } else {
              console.log("update [%s] ignored", latest)
            }
          } else if (force) {
            console.log("no update available")
            this.$bvToast.toast("Already up to date", {
              title: "Spire Update",
              autoHideDelay: 2000,
              solid: true,
              toaster: 'b-toaster-bottom-right',
            })
          }

          LocalSettings.setLastCheckedUpdateTime(new Date().getTime() / 1000)
          LocalSettings.setLatestUpdateVersion(latest)
          LocalSettings.setLatestReleasePayload(JSON.stringify(data))

          this.currentVersion = current
        })
        .catch(err => console.error(err))
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
