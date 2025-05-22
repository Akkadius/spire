<template>
  <div class='eq-window-simple'>
    <div class='eq-window-title-bar' v-if="title">{{title}}</div>
    <div>
      <slot></slot>
    </div>
  </div>
</template>

<script>
export default {
  name: 'EqWindow',
  props: {
    title: {
      type: String,
      required: false
    }
  }
}
</script>

<style scoped>

</style>
