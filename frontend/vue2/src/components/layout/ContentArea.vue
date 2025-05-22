<template>
  <div class="container-fluid content-area">
    <div class="panel-body">
      <div class="panel panel-default">
        <div class="row">
          <div class="col-12 mt-3">
            <slot></slot>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: "ContentArea"
}
</script>

<style scoped>

</style>
