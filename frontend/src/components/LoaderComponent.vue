<template>
  <b-container :class="[ 'bv-example-row', 'p-' + padding ]" v-show="isLoading">
    <b-row class="justify-content-md-center">
      <div class="spinner-border" role="status">
        <span class="sr-only">Loading...</span>
      </div>
    </b-row>
  </b-container>
</template>

<script>
  export default {
    name: 'app-loader',
    props: {
      isLoading: {
        default: null,
        required: true
      },
      padding: {
        default: 2,
        required: false
      }
    }
  }
</script>

