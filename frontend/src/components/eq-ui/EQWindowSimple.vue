<template>
  <div class='eq-window-simple'>
    <div class='eq-window-title-bar' v-if="title">{{ title }}</div>
    <div
      :style="'padding: ' + (typeof padding !== 'undefined' ? padding : 10) + 'px'"
      :class="classes"
    >
      <slot></slot>
    </div>
  </div>
</template>

<script>
export default {
  name: 'EqWindowSimple',
  props: {
    title: {
      type: String,
      required: false
    },
    padding: {
      type: Number,
      required: false,
      default: 10
    },
    classes: {
      type: String,
      required: false,
      default: ''
    }
  }
}
</script>

<style scoped>

</style>
