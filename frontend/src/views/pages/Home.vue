<template>
  <div class="container-fluid">

    <div class="row justify-content-center">
      <div class="col-12 col-lg-10 col-xl-10 content-pop">
        <!--        <page-header title="Components" pre-title="Preview Components"/>-->

        <div class="container-fluid">

          <div class="header mt-md-1">
            <div class="header-body">
              <h1 class="header-title" id="progress-bars">
                Changelog & News
              </h1>

              <p class="header-subtitle">
                See the latest Spire changes here!
              </p>
            </div>
          </div>

          <div class="row">
            <v-runtime-template :template="changelog"/>
          </div>

        </div>
      </div>
    </div>
  </div>
</template>

<script>

import EqWindow         from "@/components/eq-ui/EQWindow";
import UserContext      from "@/app/user/UserContext";
import {SpireApiClient} from "../../app/api/spire-api-client";

export default {
  components: {
    EqWindow,
    "v-runtime-template": () => import("v-runtime-template")
  },
  data() {
    return {
      userContext: null,
      changelog: "",
    }
  },
  async mounted() {
    this.userContext = await (UserContext.getUser())

    SpireApiClient.v1().get(`/app/changelog`).then((response) => {
      if (response.data && response.data.data) {

        const md = require("markdown-it")({
          html: true,
          xhtmlOut: false,
          breaks: true,
          typographer: false
        });

        let result = response.data.data

        result = md.render(result);

        console.log(result)

        // doc
        this.changelog = "<div>" + result + "</div>"

      }
    })

  }


}
</script>

