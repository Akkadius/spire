<template>
  <div class='eq-window-simple'>
    <div
      class='eq-window-title-bar'
      v-if="title"
      :style="'font-weight: ' + (boldTitle ? 'bold' : 'normal') + ';'"
    >{{ title }}
    </div>
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
    },
    boldTitle: {
      type: Boolean,
      required: false,
      default: false
    }
  }
}
</script>

<style scoped>

</style>
