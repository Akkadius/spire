<template>
  <div>
    <dashboard-navbar></dashboard-navbar>

    <div class="main-content">
      <keep-alive>
      <router-view></router-view>
      </keep-alive>
    </div>

    <dashboard-footer></dashboard-footer>

    <back-to-top bottom="50px" right="50px">
      <button type="button" class="btn btn-white btn-to-top"><i class="fa fa-chevron-up"></i></button>
    </back-to-top>
  </div>
</template>

<script lang="ts">
  import DashboardFooter from './Footer.vue'
  import DashboardHeader from './Header.vue'
  import DashboardNavbar from './Navbar.vue'
  import BackToTop from 'vue-backtotop'

  export default {
    components: {
      DashboardFooter,
      DashboardHeader,
      DashboardNavbar,
      BackToTop
    },
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
