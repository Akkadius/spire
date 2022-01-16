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
import * as util        from "util";
import VideoViewer      from "../app/video-viewer/video-viewer";

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

        let markdownRaw = response.data.data

        const youTubeSplit = markdownRaw.split("[![](https://img.youtube.com/vi/")

        youTubeSplit.forEach((e) => {
          if (e.includes("/0.jpg)](https://www.youtube.com")) {
            const videoCodeSplit = e.split("/0.jpg")
            if (videoCodeSplit.length > 0) {
              const videoCode = videoCodeSplit[0].trim()

              // replace markdown code for html
              markdownRaw = markdownRaw.replace(
                util.format("[![](https://img.youtube.com/vi/%s/0.jpg)](https://www.youtube.com/watch?v=%s)", videoCode, videoCode),
                util.format('<div class="container"><iframe allow="autoplay" class="video" src="https://www.youtube.com/embed/%s?mute=1&showinfo=0&controls=0&modestbranding=1&rel=0&loop=1&showsearch=0&iv_load_policy=3&playlist=%s" title="YouTube video player" frameborder="0" allowfullscreen></iframe></div>\n', videoCode, videoCode)
              )

              // console.log("Video code is [%s]", videoCode)
            }
          }
        })

        const md = require("markdown-it")({
          html: true,
          xhtmlOut: false,
          breaks: true,
          typographer: false
        });

        markdownRaw = md.render(markdownRaw);

        // doc
        this.changelog = "<div>" + markdownRaw + "</div>"

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

    // auto play videos that are in the viewport
    window.addEventListener("scroll", this.handleRender);
    setTimeout(() => {
      this.handleRender()
    }, 500)
  },
  methods: {
    handleRender() {
      let videos = document.getElementsByClassName("video");
      for (let i = 0; i < videos.length; i++) {
        let video = videos.item(i)
        if (VideoViewer.elementInViewport(video) && !video.src.includes("autoplay")) {
          video.src = video.src + "&autoplay=1"
        }
      }
    }
  },
  deactivated() {
    window.removeEventListener("scroll", this.handleRender, false)
  }
}
</script>

<style>
.changelog {
  font-size: 16px;
  line-height: 1.5;
  word-wrap: break-word;
  -webkit-text-size-adjust: 100%;
}

.changelog h1, .changelog h2, .changelog h3, .changelog h4, .changelog h5, .changelog h6 {
  font-weight: 600;
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
