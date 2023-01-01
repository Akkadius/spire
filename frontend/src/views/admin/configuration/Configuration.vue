<template>
  <div>

    <div class="card">
      <div class="card-body p-0">
        <!-- Nav -->
        <ul class="nav nav-tabs nav-overflow header-tabs nav-fill m-1 ml-3 mr-3">
          <li class="nav-item">
            <router-link
              to="/admin/configuration"
              active-class=""
              exact-active-class="active"
              class="nav-link"
            >
              <span class="icon mr-1"><i class="fe fe-globe"></i></span>
              World Server
            </router-link>
          </li>

          <li class="nav-item">
            <router-link
              to="/admin/configuration/settings/zoneserver"
              active-class=""
              exact-active-class="active"
              class="nav-link"
            >
              <span class="icon mr-1"><i class="fe fe-box"></i></span>
              Zone Server
            </router-link>
          </li>

          <li class="nav-item">
            <router-link
              to="/admin/configuration/settings/ucs"
              active-class=""
              exact-active-class="active"
              class="nav-link"
            >
              <span class="icon mr-1"><i class="fe fe-users"></i></span>
              Universal Chat Service
            </router-link>
          </li>


          <li class="nav-item">
            <router-link
              to="/admin/configuration/settings/database"
              active-class=""
              exact-active-class="active"
              class="nav-link"
            >
              <span class="icon mr-1"><i class="fe fe-database"></i></span>
              Database
            </router-link>
          </li>

          <li class="nav-item">
            <router-link
              to="/admin/configuration/motd"
              active-class=""
              exact-active-class="active"
              class="nav-link"
            >
              <span class="icon mr-1"><i class="fe fe-bell"></i></span>
              MOTD
            </router-link>
          </li>

          <li class="nav-item">
            <router-link
              to="/admin/configuration/server-rules"
              active-class=""
              exact-active-class="active"
              class="nav-link"
            >
              <span class="icon mr-1"><i class="fe fe-settings"></i></span>
              Rules
            </router-link>
          </li>

          <li class="nav-item">
            <router-link
              to="/admin/configuration/settings/discord"
              active-class=""
              exact-active-class="active"
              class="nav-link"
            >
                  <span class="icon mr-1">
                    <img src="~@/assets/img/icons/discord-logo-small.png" style="height: 25px;">
                  </span>
              Discord
            </router-link>
          </li>
        </ul>
      </div>
    </div>

    <router-view></router-view>

  </div>
</template>
