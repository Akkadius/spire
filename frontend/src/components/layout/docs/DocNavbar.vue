<template>
  <nav class="navbar navbar-vertical fixed-left navbar-expand-md navbar-dark navbar-vibrant" id="sidebar">
    <div class="container-fluid">

      <!-- Toggler -->
      <button class="navbar-toggler" type="button" data-toggle="collapse" data-target="#sidebarCollapse"
              aria-controls="sidebarCollapse" aria-expanded="false" aria-label="Toggle navigation">
        <span class="navbar-toggler-icon"></span>
      </button>

      <!-- Brand -->
      <a class="ml-3 mt-3" href="./index.html">
        <img src="~@/assets/img/eqemu-logo-1.png" class="navbar-brand-img mx-auto d-none d-sm-block mb-3" alt="..."
             style="max-height: 7rem">
      </a>

      <hr class="dropdown-divider">

      <h4 class=" text-center menuetto-header small-mobile">
        Spire
      </h4>

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
        <form class="mt-4 mb-3 d-md-none">
          <div class="input-group input-group-rounded input-group-merge">
            <input type="search" class="form-control form-control-rounded form-control-prepended" placeholder="Search"
                   aria-label="Search">
            <div class="input-group-prepend">
              <div class="input-group-text">
                <span class="fe fe-search"></span>
              </div>
            </div>
          </div>
        </form>

        <v-runtime-template :template="docNav"/>

        <!-- Navigation -->
        <ul class="navbar-nav mb-md-4">
          <li class="nav-item">
            <a class="nav-link " :href="backendBaseUrl + '/swagger/index.html'" target="swagger">
              <i class="ra ra-book mr-2"></i> Spire API
            </a>
          </li>
        </ul>

        <!-- Push content down -->
        <div class="mt-auto"></div>

        <!-- User (md) -->
        <div class="navbar-user d-none d-md-flex" id="sidebarUser">

          <!-- Icon -->
          <a href="#sidebarModalActivity" class="navbar-user-link" data-toggle="modal">
              <span class="icon">
                <i class="fe fe-bell"></i>
              </span>
          </a>

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

import {App}              from "@/constants/app";
import NavbarDropdownMenu from "@/components/layout/NavbarDropdownMenu";
import UserContext from "@/app/user/UserContext";
import {SpireApi}  from "../../../app/api/spire-api";

export default {
  components: {
    NavbarDropdownMenu,
    "v-runtime-template": () => import("v-runtime-template")
  },
  data() {
    return {
      backendBaseUrl: "",
      user: null,
      docNav: null,
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
      ]
    }
  },

  async mounted() {
    this.backendBaseUrl = App.BACKEND_BASE_URL
    this.user           = await UserContext.getUser()

    SpireApi.v1().get(`/doc/SUMMARY`).then((response) => {
      if (response.data && response.data.data) {

        const md = require("markdown-it")({
          html: true,
          xhtmlOut: false,
          breaks: true,
          typographer: false
        });

        let result = response.data.data

        result = md.render(result);

        result = result.replaceAll("<h1>", "<h1 class='navbar-heading'>")
        result = result.replaceAll("<h2>", "<h2 class='navbar-heading' style='margin-bottom: 5px; margin-top: 20px'>")
        result = result.replaceAll("<ul>", "<ul class='nav nav-sm flex-column4'>")
        result = result.replaceAll("<li>", "<li class='nav-item'>")
        result = result.replaceAll("<a", "<a class='nav-link'")
        result = result.replaceAll("<a", "<router-link")
        result = result.replaceAll(".md", "")
        result = result.replaceAll("href=\"", "to=\"/doc/")
        result = result.replaceAll("</a", "</router-link")

        console.log(result)

        // doc
        this.docNav = "<div>" + result + "</div>"

      }
    })

  },

  methods: {
    hasRoute: function (partial) {
      return (this.$route.path.indexOf(partial) > -1)
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
