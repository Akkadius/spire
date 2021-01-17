<template>
  <div>
    <div class="header">
      <header-image-component/>

      <div class="container-fluid">
        <!-- Body -->
        <div class="header-body mt-n5 mt-md-n6">
          <div class="row align-items-end">
            <div class="col-auto">

              <!-- Avatar -->
              <div class="avatar avatar-xxl header-avatar-top">
                <img src="@/assets/img/eqemu-avatar.png" alt="..."
                     class="avatar-img rounded-circle border border-4 border-body">
              </div>

            </div>
            <div class="col mb-3 ml-n3 ml-md-n2">

              <!-- Pretitle -->
              <h6 class="header-pretitle">
                Members
              </h6>

              <!-- Title -->
              <h1 class="header-title">
                Dianna Smiley
              </h1>

            </div>
            <div class="col-12 col-md-auto mt-2 mt-md-0 mb-md-3">

<!--              &lt;!&ndash; Button &ndash;&gt;-->
<!--              <a href="#!" class="btn btn-primary d-block d-md-inline-block lift">-->
<!--                Subscribe-->
<!--              </a>-->

            </div>
          </div>
        </div>

      </div>
    </div>

    <!-- CONTENT -->
    <div>
      <div class="container-fluid">

        Content hello

        <eq-window v-if="userContext">
          <pre>{{userContext}}
          </pre>
        </eq-window>

      </div>
    </div>

  </div>
</template>

<script>

  import HeaderImageComponent from '../../components/HeaderImageComponent';
  import EqWindow from "@/components/eq-ui/EQWindow";
  import UserContext from "@/app/user/UserContext";

  export default {
    components: { EqWindow, HeaderImageComponent },
    data() {
      return {
        userContext: null,
      }
    },
    async mounted() {
      this.userContext = await (UserContext.getUser())
    }


  }
</script>

