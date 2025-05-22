<template>
  <eq-progress-bar :percent="progress"/>
</template>

<script>
import EqProgressBar from "./eq-ui/EQProgressBar";

export default {
  name: 'loader-fake-progress',
  data() {
    return {
      progress: 0,
      internalProgress: 0,
      interval: null,
    }
  },
  components: { EqProgressBar },
  props: {
    intervalMs: {
      type: Number,
      required: false,
      default: 10
    },
  },
  mounted() {
    this.interval = setInterval(this.incrementLoader, this.intervalMs)
  },
  beforeDestroy() {
    if (this.interval) {
      clearInterval(this.interval)
    }
  },
  methods: {
    incrementLoader() {
      let progress = this.internalProgress

      if (progress < 25) {
        progress += .5;
      } else if (progress < 50) {
        progress += .1;
      } else if (progress < 75) {
        progress += .05;
      } else if (progress < 85) {
        progress += .025;
      } else if (progress < 100) {
        progress += .01;
      }

      if (progress > 100) {
        clearInterval(this.interval)
      }

      this.internalProgress = progress;
      this.progress         = Math.round(progress)
    }
  }
}
</script>
