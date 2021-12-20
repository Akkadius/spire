<template>
  <div :class="'dropdown-menu ' + (menuRight ? 'dropdown-menu-right' : '')" :aria-labelledby="'sidebarIcon' + menuRight">
    <a @click="loginGithub" class="dropdown-item" v-if="!user">
      <i class="fe fe-github"></i> Github Login
    </a>
    <span class="dropdown-item" v-if="user">{{ user.user_name }} ({{ user.provider }})</span>
    <hr class="dropdown-divider" v-if="user">
    <router-link to="/connections" class="dropdown-item" v-if="user">
      <i class="fe fe-database"></i> Manage Connections
    </router-link>
    <hr class="dropdown-divider" v-if="user">
    <router-link to="/logout" class="dropdown-item" v-if="user">Logout</router-link>
  </div>
</template>

<script>
import UserContext      from "@/app/user/UserContext";
import {SpireApiClient} from "@/app/api/spire-api-client";

export default {
  name: "NavbarDropdownMenu",
  data() {
    return {
      user: null
    }
  },
  props: {
    menuRight: String
  },
  async mounted() {
    this.user = await (UserContext.getUser())
  },
  methods: {
    loginGithub: function () {
      const width  = 800
      const height = 800
      const left   = (screen.width / 2) - (width / 2);
      const top    = (screen.height / 2) - (height / 2);
      const url    = SpireApiClient.getBasePath() + "/auth/github";
      const title  = "Github"
      const win    = window.open(url, title, "toolbar=no, location=no, directories=no, status=no, menubar=no, scrollbars=no, resizable=no, copyhistory=no, width=" + width + ", height=" + height + ", top=" + top + ", left=" + left);
      var timer    = setInterval(async () => {
        if (win.closed) {
          clearInterval(timer);
          this.$router.go(0);
          this.user = await (UserContext.getUser())
        }
      }, 100);
    }
  }
}
</script>
