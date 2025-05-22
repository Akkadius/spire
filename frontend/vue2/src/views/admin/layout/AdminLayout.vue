<template>
  <div>
    <navbar/>

    <div class="main-content">
      <content-area>
        <admin-header v-if="isLocal"/>

        <router-view></router-view>
      </content-area>
    </div>

    <footer/>

    <back-to-top bottom="50px" right="50px">
      <button type="button" class="btn btn-white btn-to-top"><i class="fa fa-chevron-up"></i></button>
    </back-to-top>
  </div>
</template>

<script lang="ts">
import BackToTop from 'vue-backtotop'
import Header from "@/components/layout/Header.vue";
import Footer from "@/components/layout/Footer.vue";
import ContentArea from "@/components/layout/ContentArea.vue";
import Navbar from "@/components/layout/Navbar.vue";
import AdminHeader from "@/views/admin/layout/AdminHeader.vue";
import {AppEnv} from "@/app/env/app-env";

export default {
  components: {
    AdminHeader,
    Navbar,
    ContentArea,
    Header,
    Footer,
    BackToTop
  },
  data() {
    return {
      isLocal: true
    }
  },
  async created() {
    await AppEnv.init()

    // @ts-ignore
    this.isLocal = AppEnv.isAppLocal()
  },
}
</script>
