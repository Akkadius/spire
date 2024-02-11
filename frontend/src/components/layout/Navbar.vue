<template>
  <nav
    class="navbar navbar-vertical fade-in fixed-left navbar-expand-md navbar-dark navbar-vibrant"
    style="overflow-y: overlay"
    id="sidebar"
    @click.self="expandNavbar()"
    v-if="!hideNavbar"
  >
    <div
      style="position: inherit; top: 50%; left: 10px; display: none"
      @click="expandNavbar()"
      id="collapse-nav-chevron"
    >
      <i
        class="fa fa-chevron-right navbar-chevron"
        style="font-size: 40px; color: white; position: relative;"
      ></i>
    </div>

    <div class="container-fluid navbar-contents">

      <!-- Toggler -->
      <button
        class="navbar-toggler" type="button" data-toggle="collapse" data-target="#sidebarCollapse"
        aria-controls="sidebarCollapse" aria-expanded="false" aria-label="Toggle navigation"
      >
        <span class="navbar-toggler-icon"></span>
      </button>

      <router-link class="ml-3 mt-3" to="/">
        <h1 class="text-center eq-header small-mobile">
          Spire
          <!--          <h3 class="text-center eq-header small-mobile d-inline" style="font-size: 40px">-->
          <!--            [Admin]-->
          <!--          </h3>-->
        </h1>
      </router-link>

      <div class="navbar-user d-md-none">
        <div class="dropdown">
          <a
            href="#" id="sidebarIcon" class="dropdown-toggle" role="button" data-toggle="dropdown" aria-haspopup="true"
            aria-expanded="false"
          >
            <div :class="'avatar avatar-sm ' + (isUserLoggedIn() ? 'avatar-online' : '')">
              <img
                :src="getUserAvatar()"
                class="avatar-img rounded-circle" alt="..."
              >
            </div>
          </a>
          <navbar-dropdown-menu menu-right="1"/>
        </div>
      </div>

      <div class="collapse navbar-collapse" id="sidebarCollapse">
        <div v-if="isAppLocal()">
          <h6 class="navbar-heading mt-3">
            Admin
          </h6>

          <ul class="navbar-nav mb-md-3">
            <li class="nav-item" v-if="!isInAdmin()">
              <router-link class="nav-link" :to="ROUTE.ADMIN_ROOT" exact>
                <i class="ra ra-eye-shield mr-1"></i> Server Admin
                <b-badge class="ml-3" variant="primary">NEW!</b-badge>
              </router-link>
            </li>

            <nav-section-component
              v-for="nav in adminNavs"
              :key="nav.label"
              :config="nav"
              v-if="isInAdmin()"
            />

          </ul>
        </div>

        <div v-if="!isInAdmin()">
          <h6 class="navbar-heading">
            Content Tools
          </h6>

          <ul class="navbar-nav mb-md-3">
            <nav-section-component :config="botNav"/>
            <nav-section-component :config="calculatorNav"/>

            <li class="nav-item">
              <router-link class="nav-link" to="/client-files">
                <i class="ra ra-cycle mr-1"></i> Client File
                <b-badge class="ml-3" variant="primary">NEW!</b-badge>
              </router-link>
            </li>

            <li class="nav-item">
              <router-link class="nav-link " to="/items">
                <i class="ra ra-relic-blade mr-1"></i> Items
                <b-badge class="ml-3" variant="primary">NEW!</b-badge>
              </router-link>
            </li>

            <nav-section-component :config="npcNav"/>

            <li class="nav-item">
              <router-link class="nav-link " to="/quest-api-explorer">
                <i class="ra ra-compass mr-1"></i> Quest API Explorer
              </router-link>
            </li>

            <li class="nav-item">
              <router-link class="nav-link " to="/sage">
                <i class="ra ra-crystal-ball mr-1"></i> Sage
                <b-badge class="ml-3" variant="primary">NEW!</b-badge>
              </router-link>
            </li>

            <li class="nav-item">
              <router-link class="nav-link " to="/strings-database">
                <i class="ra  ra-scroll-unfurled mr-1"></i> Strings DB
                <b-badge class="ml-3" variant="primary">NEW!</b-badge>
              </router-link>
            </li>
            <li class="nav-item">
              <router-link class="nav-link " to="/spells">
                <i class="ra ra-book mr-1"></i> Spells
                <b-badge class="ml-3" variant="primary">NEW!</b-badge>
              </router-link>
            </li>
            <li class="nav-item">
              <router-link class="nav-link " to="/tasks">
                <i class="ra ra-zebra-shield mr-1"></i> Tasks
                <b-badge class="ml-3" variant="primary">BETA</b-badge>
                <b-badge class="ml-3" variant="primary">NEW!</b-badge>
              </router-link>
            </li>
            <nav-section-component :config="viewerNav"/>

            <li class="nav-item">
              <router-link class="nav-link " to="/zones">
                <i class="ra ra-wooden-sign mr-2"></i> Zones
                <b-badge class="ml-3" variant="primary">ALPHA</b-badge>
                <b-badge class="ml-3" variant="primary">NEW!</b-badge>
              </router-link>
            </li>

          </ul>

          <!-- Heading -->
          <h6 class="navbar-heading">
            Spire Docs
          </h6>

          <!-- Navigation -->
          <ul class="navbar-nav mb-md-3">

            <nav-section-component :config="spireApiNav"/>

            <!-- Components -->
            <li class="nav-item">
              <a
                :class="'nav-link collapse ' + (hasRoute('components') ? 'active' : 'collapsed')"
                href="#sidebarComponents" data-toggle="collapse" role="button"
                aria-expanded="false" aria-controls="sidebarComponents"
              >
                <i class="ra ra-burst-blob mr-1"></i> Components
              </a>
              <div :class="'collapse ' + (hasRoute('components') ? 'show' : '')" id="sidebarComponents">
                <ul class="nav nav-sm flex-column">
                  <li v-for="nav in componentNavs">
                    <router-link class="nav-link" :to="nav.to">{{ nav.title }}</router-link>
                  </li>
                </ul>
              </div>
            </li>

          </ul>
        </div>

        <h6 class="navbar-heading" v-if="appVersion">
          Version ({{ appEnv }}) {{ appVersion }}
        </h6>

        <ul class="navbar-nav mb-md-3">
          <li class="nav-item" v-if="hasUpdate">
            <a
              href="#"
              style="color: yellow"
              class="nav-link pulsate-highlight-white"
              data-toggle="modal"
              @click="checkForSpireUpdate()"
            >
              <i class="fe fe-check-circle mr-2"></i>
              Spire Update Available
            </a>
          </li>
          <li class="nav-item" v-if="!hasUpdate">
            <a href="#" class="nav-link" data-toggle="modal" @click="checkForSpireUpdate()">
              <i class="fe fe-check-circle mr-2"></i>
              Spire Update Check
            </a>
          </li>
          <li class="nav-item">
            <a href="#" class="nav-link" data-toggle="modal" @click="openSearch()">
              <i class="fe fe-search mr-2"></i>
              Nav Search (Ctrl + K)
            </a>
          </li>
        </ul>

        <!-- Push content down -->
        <div class="mt-auto"></div>

        <!-- User (md) -->
        <div class="navbar-user d-none d-md-flex" id="sidebarUser">
          <navbar-user-settings-cog/>

          <!-- Dropup -->
          <div class="dropup">

            <!-- Toggle -->
            <a
              href="#" id="sidebarIconCopy" class="dropdown-toggle" role="button" data-toggle="dropdown"
              aria-haspopup="true" aria-expanded="false"
            >
              <div class="avatar avatar-sm avatar-online">
                <img
                  :src="getUserAvatar()"
                  class="avatar-img rounded-circle" alt="..."
                >
              </div>
            </a>

            <!-- Menu -->
            <navbar-dropdown-menu/>

          </div>

          <!-- Icon -->
          <a
            href="javascript:void(0)"
            class="navbar-user-link"
            @click="collapseNavbar"
          >
              <span class="icon">
                <i class="fe fe-menu"></i>
              </span>
          </a>

        </div>

        <!-- Active Database Connection Status -->
        <db-connection-status-pill/>

      </div> <!-- / .navbar-collapse -->
    </div>
  </nav>
</template>

<script>

import {App}                  from "@/constants/app";
import NavbarDropdownMenu     from "@/components/layout/NavbarDropdownMenu";
import NavbarUserSettingsCog  from "@/components/layout/NavbarUserSettingsCog";
import UserContext            from "@/app/user/UserContext";
import NavSectionComponent    from "@/components/layout/NavSectionComponent";
import {ROUTE}                from "@/routes";
import {EventBus}             from "@/app/event-bus/event-bus";
import {AppEnv}               from "@/app/env/app-env";
import {Navbar}               from "@/app/navbar";
import DbConnectionStatusPill from "@/components/DbConnectionStatusPill";
import {SpireApi}             from "@/app/api/spire-api";
import {LocalSettings}        from "@/app/local-settings/localsettings";
import semver                 from "semver";

export default {
  computed: {
    ROUTE() {
      return ROUTE
    },
    hasUpdate() {
      if (!this.latestAppVersion) {
        return false
      }

      if (!this.appVersion) {
        return false
      }

      return semver.gt(this.latestAppVersion, this.appVersion)
    }
  },
  components: { DbConnectionStatusPill, NavSectionComponent, NavbarDropdownMenu, NavbarUserSettingsCog },
  data() {
    return {

      // state
      lastPartition: "default",

      backendBaseUrl: "",
      user: null,
      hideNavbar: false,
      appEnv: AppEnv.getEnv(),
      appVersion: AppEnv.getVersion(),
      latestAppVersion: LocalSettings.getLatestUpdateVersion(),
      appFeatures: AppEnv.getFeatures(),
      botNav: {
        label: "Bots",
        labelIcon: "ra ra-reactor mr-1",
        routePrefixMatches: ["bot"],
        navs: [
          {
            title: "Spells",
            to: ROUTE.BOT_SPELLS_EDIT,
            icon: "ra ra-regeneration mr-1",
            isAlpha: true,
            isNew: true,
            routes: ['bot-spells']
          },
        ]
      },
      npcNav: {
        label: "NPCs",
        labelIcon: "ra ra-dragon mr-1",
        routePrefixMatches: ["npc", "merchant", "loot"],
        navs: [
          {
            title: "Emotes",
            to: ROUTE.NPC_EMOTES_EDIT,
            icon: "ra ra-death-skull mr-1",
            isAlpha: true,
            isNew: true,
            routes: ['npc-emotes']
          },
          {
            title: "Merchants",
            to: ROUTE.MERCHANTS,
            icon: "ra ra-emerald mr-1",
            isAlpha: true,
            isNew: true,
            routes: ['merchant', 'merchants']
          },
          {
            title: "Spells",
            to: ROUTE.NPC_SPELLS_EDIT,
            icon: "ra ra-flame-symbol mr-1",
            isAlpha: true,
            isNew: true,
            routes: ['npc-spells']
          },
          // {
          //   title: "Loot",
          //   to: ROUTE.LOOT,
          //   icon: "ra ra-sword mr-1",
          //   isAlpha: true,
          //   isNew: true,
          //   routes: ['loot']
          // },
        ]
      },
      adminNavs: [
        { label: "Server Admin", labelIcon: "ra ra-eye-shield mr-1", to: ROUTE.ADMIN_ROOT, exact: true },
        { label: "Players Online", labelIcon: "ra ra-double-team mr-1", to: ROUTE.ADMIN_PLAYERS_ONLINE },
        { label: "Zone Servers", labelIcon: "ra ra-tower mr-1", to: ROUTE.ADMIN_ZONE_SERVERS },
        { label: "Backups", labelIcon: "fa fa-download mr-1", to: ROUTE.ADMIN_BACKUPS },
        { label: "Client Files", labelIcon: "fa fa-download mr-1", to: ROUTE.ADMIN_CLIENT_FILE_DOWNLOADS },
        {
          label: "Configuration",
          labelIcon: "fa fa-cog mr-1",
          routePrefixMatch: "admin/configuration",
          navs: [
            { title: "Server Config", to: ROUTE.ADMIN_SERVER_CONFIG, icon: "fa fa-cog mr-1" },
            {
              title: "Crash Webhooks",
              to: ROUTE.ADMIN_CONFIG_DISCORD_CRASH_WEBHOOK,
              icon: "ra ra-fire mr-1",
              isOcculus: true
            },
            { title: "MOTD", to: ROUTE.ADMIN_CONFIG_MOTD, icon: "ra ra-wooden-sign mr-1" },
            {
              title: "Quest Hot Reload",
              to: ROUTE.ADMIN_CONFIG_QUEST_HOT_RELOAD,
              icon: "ra ra-alien-fire mr-1",
              isOcculus: true
            },
            { title: "Server Rules", to: ROUTE.ADMIN_CONFIG_SERVER_RULES, icon: "ra ra-interdiction mr-1" },
            { title: "UCS", to: ROUTE.ADMIN_SERVER_CONFIG + '?s=UCS', icon: "ra ra-speech-bubbles mr-1", exact: true },
            {
              title: "World Server",
              to: ROUTE.ADMIN_SERVER_CONFIG + '?s=World+Server',
              icon: "ra ra-double-team mr-1",
              exact: true
            },
            {
              title: "Zone Server",
              to: ROUTE.ADMIN_SERVER_CONFIG + '?s=Zone+Server',
              icon: "ra ra-player mr-1",
              exact: true
            },
          ]
        },
        {
          label: "Database", labelIcon: "fa fa-database mr-1", routePrefixMatch: "admin/database",
          navs: [
            {
              title: "Database Config",
              to: ROUTE.ADMIN_SERVER_CONFIG + '?s=Database',
              icon: "fa fa-cog mr-1",
              exact: true
            },
            { title: "Database Backups", to: ROUTE.ADMIN_DATABASE_BACKUP, icon: "fa fa-download mr-1" },
          ]
        },
        {
          label: "Logs",
          labelIcon: "ra ra-telescope mr-1",
          routePrefixMatch: "admin/tools/player",
          navs: [
            { title: "Discord Webhooks", to: ROUTE.ADMIN_DISCORD_WEBHOOK_SETTINGS, icon: "fa fa-cog mr-1" },
            { title: "File Logs", to: ROUTE.ADMIN_FILE_LOGS, icon: "fe fe-book mr-1" },
            { title: "Log Settings", to: ROUTE.ADMIN_LOG_SETTINGS, icon: "ra ra-scroll-unfurled mr-1" },
            { title: "Player Event Settings", to: ROUTE.ADMIN_CONFIG_PLAYER_EVENT_LOGS, icon: "fa fa-cog mr-1" },
            { title: "Player Event Log Viewer", to: ROUTE.ADMIN_TOOL_PLAYER_EVENT_LOGS, icon: "ra ra-telescope mr-1" },
          ]
        },
        { label: "Quests", labelIcon: "fa fa-code-fork mr-1", to: ROUTE.ADMIN_TOOL_SERVER_QUESTS, isOcculus: true },
        {
          label: "Reloading (Global)",
          labelIcon: "fa fa-refresh mr-1",
          routePrefixMatch: "admin/tools/player",
          to: ROUTE.ADMIN_RELOAD
        },
        { label: "Server Update", labelIcon: "fa fa-upload mr-1", to: ROUTE.ADMIN_SERVER_UPDATE },

      ],
      viewerNav: {
        label: "Viewers",
        labelIcon: "ra ra-bleeding-eye mr-1",
        routePrefixMatch: "viewer",
        navs: [
          { title: "Race Viewer", to: ROUTE.RACE_VIEWER, icon: "ra ra-monster-skull mr-1" },
          { title: "Item Model Viewer", to: ROUTE.ITEM_VIEWER, icon: "ra ra-crossed-swords mr-1" },
          { title: "Item Icon Viewer", to: ROUTE.ITEM_ICON_VIEWER, icon: "ra ra-burning-eye mr-1" },
          {
            title: "Player Animations",
            to: ROUTE.PLAYER_ANIMATION_VIEWER,
            icon: "ra ra-player-dodge mr-1",
            isNew: true
          },
          { title: "Emitter Viewer", to: ROUTE.EMITTER_VIEWER, icon: "ra  ra-droplet-splash mr-1", isNew: true },
          { title: "Spell Animations", to: ROUTE.SPELL_ANIMATION_VIEWER, icon: "ra ra-dragon mr-1", isNew: true }
        ]
      },
      spireApiNav: {
        label: "Spire API Docs",
        labelIcon: "ra ra-book mr-1",
        routePrefixMatches: ["swagger", "model-viewer"],
        navs: [
          {
            title: "Swagger API",
            to: SpireApi.getBasePath() + '/swagger/index.html',
            icon: "ra ra-monster-skull mr-1"
          },
          {
            title: "Model Relationship Explorer",
            to: ROUTE.API_MODEL_RELATIONSHIP_EXPLORER,
            icon: "ra ra-kaleidoscope mr-1"
          },
        ]
      },
      componentNavs: [
        { title: "Facial Appearance", to: "/components#facial-appearance" },
        { title: "Progress Bars", to: "/components#progress-bars" },
        { title: "Page Headers", to: "/components#page-headers" },
        { title: "Range Visual", to: "/components#range-visual" },
        { title: "Item Preview", to: "/components#item-preview" },
        { title: "Spell Icons", to: "/components#spell-icons" },
        { title: "Spell Preview", to: "/components#spell-preview" },
        { title: "NPC Card Preview", to: "/components#npc-card-preview" },
        { title: "NPC Special Abilities", to: "/components#npc-special-abilities" },
        { title: "Table", to: "/components#table" },
        { title: "Tabs", to: "/components#tabs" },
        { title: "Form Elements", to: "/components#form-elements" },
        { title: "Windows", to: "/components#windows" }
      ],
      calculatorNav: {
        label: "Calculators",
        labelIcon: "ra ra-cog mr-1",
        routePrefixMatch: "calculator",
        navs: [
          { title: "Race Calculator", to: "/calculators#race-bitmask-calculator", icon: "ra ra-eye-monster mr-1" },
          { title: "Class Calculator", to: "/calculators#class-bitmask-calculator", icon: "ra ra-lion mr-1" },
          { title: "Deity Calculator", to: "/calculators#deity-bitmask-calculator", icon: "ra ra-venomous-snake mr-1" },
          {
            title: "Client Version Calculator",
            to: "/calculators#client-version-calculator",
            icon: "ra ra-lever mr-1"
          },
          {
            title: "Expansions Calculator",
            to: "/calculators#expansions-bitmask-calculator",
            icon: "ra ra-lever mr-1"
          },
          { title: "Augment Type Calculator", to: "/calculators#augment-type-calculator", icon: "ra ra-sapphire mr-1" },
          {
            title: "Inventory Slot Calculator",
            to: "/calculators#inventory-slot-calculator",
            icon: "ra ra-eye-shield mr-1"
          },
          { title: "Special Abilities Calculator", to: "/calculators#npc-special-abilities", icon: "ra ra-lion mr-1" },
        ]
      },


    }
  },
  created() {
    EventBus.$on("HIDE_NAVBAR", this.toggleNavbarCollapse);
    EventBus.$on("APP_ENV_LOADED", this.handleAppEnvLoaded);
    EventBus.$on("ROUTE_CHANGE", this.handleRouteChange);
  },
  destroyed() {
    EventBus.$off("HIDE_NAVBAR", this.toggleNavbarCollapse);
    EventBus.$off("APP_ENV_LOADED", this.handleAppEnvLoaded);
    EventBus.$off("ROUTE_CHANGE", this.handleRouteChange);
  },

  async mounted() {
    this.backendBaseUrl = App.BACKEND_BASE_URL
    this.user           = await UserContext.getUser()

    // sidebar
    this.setSidebarStyle()

    this.parseNinjaKeys()

    setTimeout(() => {
      this.latestAppVersion = LocalSettings.getLatestUpdateVersion()
      this.$forceUpdate()
    }, 1000)
  },

  methods: {

    updateSpire() {
      this.$bvModal.show('app-update-modal')
    },

    openSearch() {
      const ninja = document.querySelector('ninja-keys')
      setTimeout(() => {
        ninja.open();
      }, 1)
    },

    parseNinjaNav(nav) {
      let keys = []
      for (let n of nav) {
        if (n.label && n.to) {
          let adminPanelRouteEnabled = false
          if (n.to.includes(ROUTE.ADMIN_ROOT)) {
            if (!AppEnv.isAppLocal()) {
              continue;
            }
            adminPanelRouteEnabled = true
          }

          keys.push({
            id: n.label,
            title: (adminPanelRouteEnabled ? '[Admin] ' : '') + n.label,
            handler: () => {
              this.$router.push(n.to).catch((e) => {
              })
            }
          })
        }

        // children
        if (n.navs) {
          for (let c of n.navs) {
            let adminPanelRouteEnabled = false
            if (c.to.includes(ROUTE.ADMIN_ROOT)) {
              if (!AppEnv.isAppLocal()) {
                continue;
              }
              adminPanelRouteEnabled = true
            }

            keys.push({
              id: `[${n.label}] ${c.title}`,
              title: (adminPanelRouteEnabled ? '[Admin] ' : '') + `[${n.label}] ${c.title}`,
              handler: () => {
                this.$router.push(c.to).catch((e) => {
                })
              }
            })
          }
        }
      }

      return keys
    },

    parseNinjaKeys() {
      let keys = []

      let navs = [
        this.adminNavs,
        [this.botNav],
        [this.npcNav],
        [this.viewerNav],
        [this.spireApiNav],
        [this.componentNavs],
        [this.calculatorNav],
      ]

      for (let n of navs) {
        keys = keys.concat(this.parseNinjaNav(n))
      }

      let manualRoutes = [
        { name: "Coffee", route: ROUTE.COFFEE },
        { name: "Tasks", route: ROUTE.TASKS },
        { name: "Items", route: ROUTE.ITEMS_LIST },
        { name: "Spells", route: ROUTE.SPELLS_LIST },
        { name: "[Quest API] Explorer", route: ROUTE.QUEST_API_EXPLORER },
        { name: "[Quest API] Explorer (Perl)", route: `${ROUTE.QUEST_API_EXPLORER}?lang=perl` },
        { name: "[Quest API] Explorer (Lua)", route: `${ROUTE.QUEST_API_EXPLORER}?lang=lua` },
        { name: "Zones", route: ROUTE.ZONES },
      ]

      for (let m of manualRoutes) {
        keys.push({
          id: m.name, title: m.name, handler: () => {
            this.$router.push(m.route).catch((e) => {
            })
          }
        })
      }

      keys = keys.sort((a, b) => {
        return a.title.localeCompare(b.title);
      });


      const ninja = document.querySelector('ninja-keys')
      ninja.data  = keys
    },

    isInAdmin() {
      return this.$route.path.includes("/admin")
    },

    getPartitionName() {
      if (this.isInAdmin()) {
        return "admin"
      }

      return "default"
    },

    isAppLocal() {
      return AppEnv.isAppLocal()
    },

    isUserLoggedIn() {
      return this.user && this.user.avatar && this.user.avatar.length > 0
    },

    getUserAvatar() {
      if (this.user && this.user.avatar && this.user.avatar.length > 0) {
        return this.user.avatar
      }

      return require('@/assets/img/spire.png')
    },

    setSidebarStyle() {
      setTimeout(() => {
        if (this.lastPartition !== this.getPartitionName()) {
          const sidebar = document.getElementById("sidebar")

          if (sidebar) {
            // remove all classes
            sidebar.classList.remove("sidebar-admin")
            sidebar.classList.remove("sidebar-normal")

            // trigger reset class
            // this tricks color animation to bring opacity down before transitioning back
            // up with the other sidebar classes
            sidebar.classList.add("sidebar-reset")

            setTimeout(() => {
              if (this.isInAdmin()) {
                sidebar.classList.remove("sidebar-reset")
                sidebar.classList.add("sidebar-admin")
              } else {
                sidebar.classList.remove("sidebar-reset")
                sidebar.classList.add("sidebar-normal")
              }
            }, 300)
          }

          this.lastPartition = this.getPartitionName()
        }
      }, 10)
    },

    handleRouteChange() {
      this.setSidebarStyle()
    },
    handleAppEnvLoaded() {
      this.appEnv      = AppEnv.getEnv();
      this.appVersion  = AppEnv.getVersion();
      this.appFeatures = AppEnv.getFeatures();
    },
    expandNavbar() {
      Navbar.expand()
    },
    collapseNavbar() {
      Navbar.collapse()
    },
    toggleNavbarCollapse() {
      Navbar.toggleCollapse()
    },
    hideNavBar() {
      this.hideNavbar = !this.hideNavbar
    },
    hasRoute: function (partial) {
      return (this.$route.path.indexOf(partial) > -1)
    },
    hideNavbarAfterClick() {
      const sidebar = document.getElementById("sidebarCollapse")
      if (sidebar) {
        sidebar.classList.remove("show");
      }
    },
    checkForSpireUpdate() {
      EventBus.$emit("CHECK_SPIRE_UPDATE", true)
    }
  },
  watch: {
    $route(to, from) {
      this.hideNavbarAfterClick()
    }
  }
}
</script>

<style scoped>
.navbar-nav .nav-link > .fe {
  min-width: 1.25rem;
}

.navbar-dark.navbar-vibrant {
  background-image: linear-gradient(to bottom right, rgb(21 21 21 / 90%), rgb(191 126 70 / 90%)), url(~@/assets/img/eq-wallpaper-1.jpg);
}

@media only screen and (max-device-width: 640px) {
  .small-mobile {
    font-size: 40px !important;
  }
}

.connection-status-box {
  display: block; /* or inline-block */
  text-overflow: ellipsis;
  word-wrap: break-word;
  overflow: hidden;
}

</style>
