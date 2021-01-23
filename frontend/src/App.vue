<template>
  <router-view></router-view>
</template>

<script>

export default {
  name: "example",
  data() {
    return {}
  },
  mounted() {
    var self = this
    window.addEventListener("keypress", function (e) {
      if (e.srcElement.tagName !== "BODY" && e.srcElement.tagName !== "A") {
        return
      }

      if (window.location.pathname === "/login") {
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
    "$route"(to, from) {
      document.title = "[Spire] " + to.meta.title || "Spire"
    }
  }
}
</script>

<style>
.card-slim {
  padding: 18px !important;
}

/*akkadius:experimental*/
body {
  /* background-repeat: no-repeat; */
  background-size:       cover;

  background:            radial-gradient(circle, transparent 40%, var(--color-v) 75%), linear-gradient(to right, var(--color), var(--color)), var(--image2);
  background-position:   center center;
  background-size:       cover;
  background-blend-mode: var(--blend-top, normal), var(--blend-bottom, saturation), normal;
  --image:               url(~@/assets/img/champions-of-norrath-wallpaper.jpg);
  --image2:              url(~@/assets/img/champions-of-norrath-wallpaper.jpg);
  --color-v:             black;
  --color:               grey;
}
</style>
