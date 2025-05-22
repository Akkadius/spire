<template>
  <div></div>
</template>

<script>

import UserContext from "@/app/user/UserContext";
import {ROUTE}     from "@/routes";

export default {
  name: "Logout.vue",
  async created() {
    UserContext.reset()
    UserContext.deleteAccessToken()
    this.$router.push({ path: ROUTE.LOGIN })
  }
}
</script>
