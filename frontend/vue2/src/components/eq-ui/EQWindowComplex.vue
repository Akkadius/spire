<template>
    <div class='eq-window-complex'>
        <div class='eq-window-title-bar' v-if="title">{{title}}</div>
        <div style="padding: 10px">
            <slot></slot>
        </div>
    </div>
</template>

<script>
  export default {
    name: 'EqWindowComplex',
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
