<template>
  <div>
    <navbar/>

    <div class="main-content">
      <content-area class="text-center fade-in">
        <h1 class="eq-header mt-8" style="font-size: 120px">
          Occulus
        </h1>

        <div class="row justify-content-center">
          <div class="card col-6">
            <div class="card-body" style="font-size: 16px">
              A local version of Spire is required to use Occulus features
            </div>
          </div>
        </div>

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

export default {
  name: "OcculusRequire",
  components: {
    AdminHeader,
    Navbar,
    ContentArea,
    Header,
    Footer,
    BackToTop
  },
}
</script>
