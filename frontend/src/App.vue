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

<style lang="scss">
@import "../node_modules/ag-grid-community/dist/styles/ag-grid.css";
@import "../node_modules/ag-grid-community/dist/styles/ag-theme-balham-dark.css";

.card-slim {
  padding: 18px !important;
}

/*akkadius:experimental*/
body {
  /* background-repeat: no-repeat; */
  background:            -webkit-radial-gradient(circle, transparent 40%, var(--color-v) 75%), linear-gradient(to right, var(--color), var(--color)), var(--image2) !important;
  background:            radial-gradient(circle, transparent 40%, var(--color-v) 75%), linear-gradient(to right, var(--color), var(--color)), var(--image2);


  background-position:   center center;
  background-size:       cover;
  background-repeat:     no-repeat;
  background-attachment: fixed;
  background-blend-mode: normal, saturation, normal;
  --image:               url(~@/assets/img/champions-of-norrath-wallpaper.jpg);
  --image2:              url(~@/assets/img/champions-of-norrath-wallpaper.jpg);
  --color-v:             black;
  --color:               grey;
}



</style>
