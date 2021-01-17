<template>
  <div class="card mt-5" v-if="debug">
    <div class="card-body">
      Debug
    </div>
    <div class="card-footer bg-dark pb-0">
      <pre class="html highlight bg-dark hljs xml" style="color: #569CD6">{{ data }}</pre>
    </div>
  </div>
</template>

<script>

import {App} from "@/constants/app";

export default {
  name: "DebugDisplayComponent",
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
