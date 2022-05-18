<template>
  <div class='eq-window' :style="'margin-bottom: 40px; min-height: 275px; ' + (title ? 'padding-top: 30px' : 'padding-top: 0px !important')">
    <div class='eq-window-title-bar' v-if="title">{{ title }}</div>
    <div :style="'padding: 10px; ' + (title ? 'margin-top: 10px' : '') ">
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
/*@import './eq-window-styles.css';*/
</style>
