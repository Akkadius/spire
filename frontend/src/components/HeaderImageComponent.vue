<template>
  <div style="height:193px; text-align:center; background-color: black">
    <transition name="slide-fade">
      <img
        :src="headerImage"
        :key="headerImage"
        class="header-img-top"
        style="height:193px;"
        alt="...">
    </transition>
  </div>
</template>

<script>
import Timer from "@/app/timer/timer"

export default {
  name: "HeaderImageComponent",
  data() {
    return {
      headerImage: require("@/assets/img/banners/1-new.png")
    }
  },
  methods: {
    rotateImage: function () {
      const imageNumber = Math.floor(Math.random() * 34) + 1;
      this.headerImage  = require(`@/assets/img/banners/${imageNumber}-new.png`);
    }
  },
  beforeDestroy() {
    clearInterval(Timer.timer["rotate-header-images"])
  },
  async mounted() {
    this.rotateImage()

    Timer.timer["rotate-header-images"] = setInterval(() => {
      this.rotateImage()
    }, 5000)
  }
}
</script>

<style type="text/css">
.slide-fade-enter-active {
  transition: all 1s ease;
}

.slide-fade-leave-active {
  transition: all .0s cubic-bezier(1.0, 0.5, 0.8, 1.0);
}

.slide-fade-enter, .slide-fade-leave-to
  /* .slide-fade-leave-active below version 2.1.8 */
{
  transform: translateX(10px);
  opacity:   0;
}
</style>
