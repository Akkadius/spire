<template>
  <div
    :class="'dropdown-menu ' + (menuRight ? 'dropdown-menu-right' : '')"
    :aria-labelledby="'sidebarIcon' + menuRight"
  >
    <router-link :to="ROUTE.LOGIN" class="dropdown-item" v-if="!user && githubAuthEnabled">
      <i class="fe fe-user"></i> Login
    </router-link>
    <span class="dropdown-item" v-if="user">{{ user.user_name }} ({{ user.provider }})</span>
    <hr class="dropdown-divider" v-if="user">
    <router-link :to="ROUTE.DATABASE_CONNECTIONS" class="dropdown-item" v-if="user">
      <i class="fe fe-database"></i> Manage Connections
    </router-link>
    <router-link :to="ROUTE.USER_MANAGEMENT" class="dropdown-item" v-if="user && isAdmin()">
      <i class="fe fe-user"></i> Manage Users
    </router-link>
    <hr class="dropdown-divider" v-if="user">
    <router-link :to="ROUTE.LOGOUT" class="dropdown-item" v-if="user">Logout</router-link>
  </div>
</template>

<script>
import UserContext from "@/app/user/UserContext";
import {AppEnv}    from "@/app/env/app-env";
import {ROUTE}     from "@/routes";

export default {
  name: "NavbarDropdownMenu",
  data() {
    return {
      user: null,
      githubAuthEnabled: AppEnv.isGithubAuthEnabled(),
      ROUTE: ROUTE
    }
  },
  props: {
    menuRight: String
  },
  methods: {
    isAdmin() {
      return this.user.is_admin
    }
  },
  async mounted() {
    this.user = await (UserContext.getUser())

    setTimeout(() => {
      this.githubAuthEnabled = AppEnv.isGithubAuthEnabled()
    }, 1000)
  },
}
</script>
