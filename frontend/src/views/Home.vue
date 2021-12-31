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

          <div class="row" id="changelog">
            <div class="col-12">
              <v-runtime-template class="changelog" :template="changelog"/>
            </div>
          </div>

        </div>
      </div>
    </div>
  </div>
</template>

<script>

import EqWindow         from "@/components/eq-ui/EQWindow";
import UserContext      from "@/app/user/UserContext";
import {SpireApiClient} from "../app/api/spire-api-client";

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

        // doc
        this.changelog = "<div>" + result + "</div>"

        setTimeout(() => {
          const anchors = document.getElementById('changelog').getElementsByTagName('a');
          for (var i = 0; i < anchors.length; i++) {
            anchors[i].setAttribute('target', '_blank');
          }

          document.querySelectorAll('#changelog h1, #changelog h2, #changelog h3, #changelog h4').forEach($heading => {

            //create id from heading text
            const id = $heading.getAttribute("id") || $heading.innerText.toLowerCase().replace(/[`~!@#$%^&*()_|+\-=?;:'",.<>\{\}\[\]\\\/]/gi, '').replace(/ +/g, '-');

            //add id to heading
            $heading.setAttribute('id', id);

            //append parent class to heading
            $heading.classList.add('anchor-heading');

            //create anchor
            let $anchor         = document.createElement('a');
            $anchor.className   = 'anchor-link';
            $anchor.href        = '#' + id;
            $anchor.innerText   = ' # ';
            $anchor.style.color = '#666';

            //append anchor after heading text
            $heading.append($anchor);
          });
        }, 100)

      }
    })

  }


}
</script>

<style>
.changelog {
  font-size: 16px;
  line-height: 1.5;
  word-wrap: break-word;
}

.container {
  position: relative;
  width: 100%;
  height: 0;
  padding-bottom: 56.25%;
}

.anchor-link:hover {
  display: initial;
}

.anchor-link {
  display: none;
}

.video {
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  padding-bottom: 30px;
  padding-top: 15px;
}
</style>
