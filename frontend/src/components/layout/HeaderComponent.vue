<template>
  <div class="row">
    <div class="header mt-md-1">
      <div class="header-body">
        <h1 class="header-title" :id="headerId">
          {{ header }}
        </h1>

        <p class="header-subtitle" v-if="subHeader">
          {{ subHeader }}
        </p>
      </div>
    </div>
  </div>
</template>

<script>
import slugify from "slugify";

export default {
  name: "HeaderComponent",
  data() {
    return {
      headerId: ""
    }
  },
  props: {
    header: {
      type: String,
      required: true
    },
    subHeader: {
      type: String,
      required: false
    }
  },
  mounted() {
    this.headerId = slugify(this.header.toLowerCase())
  }
}
</script>

<style scoped>

</style>
