<template>
  <div class="header">
<!--    <header-image-component/>-->

    <div class="container-fluid mt-6">

      <!-- Body -->
      <div class="header-body mt-n5 mt-md-n6">
        <div class="row align-items-end">
          <div class="col mb-3 ml-n3 ml-md-n2">

            <!-- Pretitle -->
            <h6 class="header-pretitle">
              {{ preTitle }}
            </h6>

            <!-- Title -->
            <h1 class="header-title">
              {{ title}}
            </h1>

          </div>
        </div>
      </div>

    </div>
  </div>
</template>

<script>


export default {
  name: "PageHeader",
  props: {
    preTitle: String,
    title: String,
  },
}
</script>

<style scoped>

</style>
