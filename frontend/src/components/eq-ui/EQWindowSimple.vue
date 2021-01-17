<template>
    <div class='eq-simple-window'>
        <div class='eq-window-title-bar' v-if="title">{{title}}</div>
        <div style="padding: 10px">
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
      }
    }
  }
</script>

<style scoped>

</style>
