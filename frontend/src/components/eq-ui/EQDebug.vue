<template>
  <pre
    v-if="debug"
    class="highlight bg-dark hljs"
    style="max-width: 400px; padding-top: 10px !important">{{ data }}</pre>
</template>

<script>

import {App} from "@/constants/app";

export default {
  name: "EqDebug",
  props: {
    data: Object
  },
  data() {
    return {
      debug: App.DEBUG
    }
  }
}
</script>

<style scoped>

</style>
