<template>
  <nav class="navbar navbar-vertical fixed-left navbar-expand-md navbar-dark navbar-vibrant" id="sidebar">
    <div class="container-fluid">

      <!-- Toggler -->
      <button class="navbar-toggler" type="button" data-toggle="collapse" data-target="#sidebarCollapse"
              aria-controls="sidebarCollapse" aria-expanded="false" aria-label="Toggle navigation">
        <span class="navbar-toggler-icon"></span>
      </button>

      <!-- Brand -->
      <router-link class="ml-3 mt-3 d-none d-lg-block" to="/">
        <img src="~@/assets/img/eqemu-logo-1.png"
             class="navbar-brand-img mx-auto d-none d-sm-block mb-3" alt="..."
             style="max-height: 6rem">
      </router-link>


      <router-link class="ml-3 mt-3" to="/">
        <h1 class="text-center eq-header small-mobile">
          Spire
        </h1>

      </router-link>

      <!--      <hr class="dropdown-divider">-->


      <!--      <hr class="dropdown-divider">-->

      <!--      <h4 class=" text-center menuetto-header mt-2" style="font-size: 34px">-->
      <!--        Tools-->
      <!--      </h4>-->

      <!-- User (xs) -->
      <div class="navbar-user d-md-none">

        <!-- Dropdown -->
        <div class="dropdown">

          <!-- Toggle -->
          <a href="#" id="sidebarIcon" class="dropdown-toggle" role="button" data-toggle="dropdown" aria-haspopup="true"
             aria-expanded="false">
            <div class="avatar avatar-sm avatar-online">
              <img :src="user ? user.avatar : require('@/assets/img/eqemu-avatar.png')"
                   class="avatar-img rounded-circle" alt="...">
            </div>
          </a>

          <navbar-dropdown-menu menu-right="1"/>

        </div>

      </div>

      <!-- Collapse -->
      <div class="collapse navbar-collapse" id="sidebarCollapse">

        <!-- Form -->
        <!--        <form class="mt-4 mb-3 d-md-none">-->
        <!--          <div class="input-group input-group-rounded input-group-merge">-->
        <!--            <input type="search" class="form-control form-control-rounded form-control-prepended" placeholder="Search"-->
        <!--                   aria-label="Search">-->
        <!--            <div class="input-group-prepend">-->
        <!--              <div class="input-group-text">-->
        <!--                <span class="fe fe-search"></span>-->
        <!--              </div>-->
        <!--            </div>-->
        <!--          </div>-->
        <!--        </form>-->

        <!-- Heading -->
        <h6 class="navbar-heading">
          Tools
        </h6>

        <!-- Navigation -->
        <ul class="navbar-nav mb-md-4">
          <nav-section-component :config="viewerNav"/>
          <nav-section-component :config="calculatorNav"/>
        </ul>

        <h6 class="navbar-heading">
          Tools
        </h6>

        <!-- Navigation -->
        <ul class="navbar-nav mb-md-4">
          <li class="nav-item">
            <router-link class="nav-link " to="/spells">
              <i class="ra ra-book mr-1"></i> Spells
            </router-link>
          </li>
          <li class="nav-item">
            <router-link class="nav-link " to="/items">
              <i class="ra ra-relic-blade mr-1"></i> Items
            </router-link>
          </li>
        </ul>

        <!-- Heading -->
        <h6 class="navbar-heading">
          Documentation
        </h6>

        <!-- Navigation -->
        <ul class="navbar-nav mb-md-4">

          <li class="nav-item">
            <a class="nav-link " :href="backendBaseUrl + '/swagger/index.html'" target="swagger">
              <i class="ra ra-book mr-2"></i> Spire API
            </a>
          </li>

          <li class="nav-item">
            <router-link class="nav-link " to="/quest-api-explorer">
              <i class="ra ra-compass mr-2"></i> Quest API Explorer
            </router-link>
          </li>

          <!-- Components -->
          <li class="nav-item">
            <a :class="'nav-link collapse ' + (hasRoute('components') ? 'active' : 'collapsed')"
               href="#sidebarComponents" data-toggle="collapse" role="button"
               aria-expanded="false" aria-controls="sidebarComponents">
              <i class="fe fe-book-open mr-1"></i> Components
            </a>
            <div :class="'collapse ' + (hasRoute('components') ? 'show' : '')" id="sidebarComponents">
              <ul class="nav nav-sm flex-column">
                <li v-for="nav in componentNavs">
                  <router-link class="nav-link" :to="nav.to">{{ nav.title }}</router-link>
                </li>
              </ul>
            </div>
          </li>

          <!-- Test Pages -->
          <li class="nav-item">
            <a :class="'nav-link collapse ' + (hasRoute('-test') ? 'active' : 'collapsed')"
               href="#test-pages" data-toggle="collapse" role="button"
               aria-expanded="false" aria-controls="test-pages">
              <i class="fe fe-box mr-1"></i> Test Pages
            </a>
            <div :class="'collapse ' + (hasRoute('components') ? 'show' : '')" id="test-pages">
              <ul class="nav nav-sm flex-column">
                <li v-for="nav in testPageNavs">
                  <router-link class="nav-link" :to="nav.to">{{ nav.title }}</router-link>
                </li>
              </ul>
            </div>
          </li>

        </ul>

        <!--        &lt;!&ndash; Heading &ndash;&gt;-->
        <!--        <h6 class="navbar-heading">-->
        <!--          Editors-->
        <!--        </h6>-->

        <!--        &lt;!&ndash; Navigation &ndash;&gt;-->
        <!--        <ul class="navbar-nav mb-md-4">-->
        <!--          <li class="nav-item">-->
        <!--            <router-link class="nav-link " to="/tasks/">-->
        <!--              <i class="fe fe-list"></i> Task Editor-->
        <!--            </router-link>-->
        <!--          </li>-->
        <!--        </ul>-->

        <!-- Push content down -->
        <div class="mt-auto"></div>

        <!-- User (md) -->
        <div class="navbar-user d-none d-md-flex" id="sidebarUser">

          <navbar-user-settings-cog/>

          <!-- Icon -->


          <!-- Dropup -->
          <div class="dropup">

            <!-- Toggle -->
            <a href="#" id="sidebarIconCopy" class="dropdown-toggle" role="button" data-toggle="dropdown"
               aria-haspopup="true" aria-expanded="false">
              <div class="avatar avatar-sm avatar-online">
                <img :src="user ? user.avatar : require('@/assets/img/eqemu-avatar.png')"
                     class="avatar-img rounded-circle" alt="...">
              </div>
            </a>

            <!-- Menu -->
            <navbar-dropdown-menu/>

          </div>

          <!-- Icon -->
          <a href="#sidebarModalSearch" class="navbar-user-link" data-toggle="modal">
              <span class="icon">
                <i class="fe fe-search"></i>
              </span>
          </a>

        </div>


      </div> <!-- / .navbar-collapse -->

    </div>
  </nav>
</template>

<script>

import {App}                 from "@/constants/app";
import NavbarDropdownMenu    from "@/views/layout/NavbarDropdownMenu";
import NavbarUserSettingsCog from "@/views/layout/NavbarUserSettingsCog";
import UserContext           from "@/app/user/UserContext";
import NavSectionComponent   from "@/views/layout/NavSectionComponent";
import {ROUTE}               from "@/routes";

export default {
  components: { NavSectionComponent, NavbarDropdownMenu, NavbarUserSettingsCog },
  data() {
    return {
      backendBaseUrl: "",
      user: null,
      componentNavs: [
        { title: "Progress Bars", to: "/components#progress-bars" },
        { title: "Page Headers", to: "/components#page-headers" },
        { title: "Item Preview", to: "/components#item-preview" },
        { title: "Spell Preview", to: "/components#spell-preview" },
        { title: "NPC Special Abilities", to: "/components#npc-special-abilities" },
        { title: "Table", to: "/components#table" },
        { title: "Tabs", to: "/components#tabs" },
        { title: "Form Elements", to: "/components#form-elements" },
        { title: "Windows", to: "/components#windows" }
      ],
      viewerNav: {
        label: "Viewers",
        labelIcon: "ra ra-bleeding-eye mr-1",
        routePrefixMatch: "viewer",
        navs: [
          { title: "Race Viewer", to: ROUTE.RACE_VIEWER, icon: "ra ra-monster-skull mr-1" },
          { title: "Item Model Viewer", to: ROUTE.ITEM_VIEWER, icon: "ra ra-crossed-swords mr-1" },
          { title: "Item Icon Viewer", to: ROUTE.ITEM_ICON_VIEWER, icon: "ra ra-burning-eye mr-1" },
          { title: "Spell Animation Viewer", to: ROUTE.SPELL_ANIMATION_VIEWER, icon: "ra ra-dragon mr-1" }
        ]
      },
      calculatorNav: {
        label: "Calculators",
        labelIcon: "ra ra-cog mr-1",
        routePrefixMatch: "calculator",
        navs: [
          { title: "Race Calculator", to: "/calculators#race-bitmask-calculator", icon: "ra ra-eye-monster mr-1" },
          { title: "Class Calculator", to: "/calculators#class-bitmask-calculator", icon: "ra ra-lion mr-1" },
          { title: "Deity Calculator", to: "/calculators#deity-bitmask-calculator", icon: "ra ra-venomous-snake mr-1" },
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
      testPageNavs: [
        { title: "Items Test", to: "/items-test" },
        { title: "Task Editor (Non Functional)", to: "/tasks" }
      ]
    }
  },

  async mounted() {
    this.backendBaseUrl = App.BACKEND_BASE_URL
    this.user           = await UserContext.getUser()
  },

  methods: {
    hasRoute: function (partial) {
      return (this.$route.path.indexOf(partial) > -1)
    },
    hideNavbarAfterClick() {
      const sidebar = document.getElementById("sidebarCollapse")
      if (sidebar) {
        sidebar.classList.remove("show");
      }
    },
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

</style>
