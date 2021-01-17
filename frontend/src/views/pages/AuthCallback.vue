<template>
  <div></div>
</template>

<script>
import UserContext from '@/app/user/UserContext'
export default {
  name: 'AuthCallback',
  mounted() {
    const jwt = this.$route.query.jwt;
    UserContext.reset()
    UserContext.storeAccessToken(jwt)
    window.close()
  }
}
</script>
