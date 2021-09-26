<template>
  <router-view></router-view>
</template>

<script>

export default {
  name: "App",
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

    const backgrounds = [
      require("@/assets/img/wallpaper/champions-of-norrath-wallpaper.jpg"),
      require("@/assets/img/wallpaper/everquest-wallpaper-1.jpeg"),
      require("@/assets/img/wallpaper/faydark.png"),
      require("@/assets/img/wallpaper/freeport.png"),
      require("@/assets/img/wallpaper/lavastorm.png"),
      require("@/assets/img/wallpaper/rivervale.png"),
      require("@/assets/img/wallpaper/qeynos.png"),
      require("@/assets/img/wallpaper/soldungb.png"),
      require("@/assets/img/wallpaper/spire-wall.jpg"),
      require("@/assets/img/wallpaper/void.jpeg"),
    ];

    console.log(backgrounds);

    const background = backgrounds[Math.floor(Math.random() * backgrounds.length)].trim();

    document.body.style.setProperty("--image", "url(" + background + ")");
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
  background:            -webkit-radial-gradient(circle, transparent 40%, var(--color-v) 95%), linear-gradient(to right, var(--color), var(--color)), var(--image) !important;
  background:            radial-gradient(circle, transparent 40%, var(--color-v) 95%), linear-gradient(to right, var(--color), var(--color)), var(--image);


  background-position:   center center;
  background-size:       cover;
  background-repeat:     no-repeat;
  background-attachment: fixed;
  background-blend-mode: normal, saturation, normal;
  /*--image:               url(~@/assets/img/wallpaper/champions-of-norrath-wallpaper.jpg);*/
  /*--image2:              url(https://preview.redd.it/z7oyx36pvc871.png?width=2560&format=png&auto=webp&s=96b8b9471ee1392e7817f8df93fc8119a27fc14a);*/
  --color-v:             black;
  --color:               grey;
}


</style>
