<template>
  <div>
    <doc-navbar></doc-navbar>

    <div class="main-content">
      <keep-alive>
        <router-view></router-view>
      </keep-alive>
    </div>

    <doc-footer></doc-footer>

    <back-to-top bottom="50px" right="50px">
      <button type="button" class="btn btn-white btn-to-top"><i class="fa fa-chevron-up"></i></button>
    </back-to-top>
  </div>
</template>

<script lang="ts">
import DocFooter from './DocFooter.vue'
import DocNavbar from './DocNavbar.vue'
import BackToTop from 'vue-backtotop'

export default {
  components: {
    DocFooter: DocFooter,
    DocNavbar: DocNavbar,
    BackToTop
  }
}
</script>

<style>
/*body {*/
/*  background: url('~@/assets/img/spire-wall.jpg');*/
/*  background-repeat: no-repeat;*/
/*  background-size: cover;*/
/*  background-attachment: fixed;*/
/*}*/
</style>
