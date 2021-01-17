<template>
    <router-view></router-view>
</template>

<script>

export default {
  name: 'example',
  data () {
    return {}
  },
  mounted () {
    var self = this
    window.addEventListener('keypress', function (e) {
      if (e.srcElement.tagName !== 'BODY' && e.srcElement.tagName !== 'A') {
        return
      }

      if (window.location.pathname === '/login') {
        return
      }

      switch (String.fromCharCode(e.keyCode)) {
        // case 'h':
        //   self.$router.push('/servers')
        //   break
      }
    })
  },
  watch: {
    '$route' (to, from) {
      document.title = '[Spire] ' + to.meta.title || 'Spire'
    }
  },
}
</script>

<style>
  .card-slim {
    padding: 18px !important;
  }
</style>
