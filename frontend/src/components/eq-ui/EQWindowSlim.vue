<template>
    <div class='eq-window-fancy'>
        <div class='eq-window-title-bar' v-if="title">{{title}}</div>
        <div style="padding: 10px">
            <slot></slot>
        </div>
    </div>
</template>

<script>
  export default {
    name: 'EqWindowSlim',
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
