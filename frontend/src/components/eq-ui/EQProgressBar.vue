<template>
    <div class='eq-progress-bar'>
        <div :style=' {width: percent + "%", "background-color": color }'>
            <span v-if="percent > 0" style="text-shadow: -2px 2px 3px #000;">{{percent}}%</span>
        </div>
    </div>
</template>

<script>
  export default {
    name: 'EqProgressBar',
    props: {
      title: {
        type: String,
        required: false,
        default: 'Window Title'
      },
      percent: {
        type: Number,
        required: false,
        default: 0
      },
      color: {
        type: String,
        required: false,
        default: 'yellow'
      }
    }
  }
</script>

<style scoped>

</style>
