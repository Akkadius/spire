<template>
  <eq-window
    :style="'height: 96vh; '"
    id="zone-preview"
    v-if="zone"
  >
    <h6 class="eq-header">{{ getZoneLongName() }}</h6>
  </eq-window>
</template>

<script>
import EqWindow         from "../eq-ui/EQWindow";
import {SpireApiClient} from "../../app/api/spire-api-client";
export default {
  name: "EqZoneCardPreview",
  components: { EqWindow },
  props: {
    zone: Object,
    required: true,
  },
  created() {
    this.backgroundImages = []

    // cycle background images
    this.interval = setInterval(this.setBackgroundImage, 10 * 1000)
  },
  beforeDestroy() {
    if (this.interval) {
      clearInterval(this.interval)
    }
  },
  mounted() {
    this.init()
  },
  watch: {
    zone: {
      handler: function (val, oldVal) {
        this.init()
      },
    },
  },
  methods: {
    init() {
      // get zone wallpaper
      this.loadBackgroundImages().then(() => {
        this.setBackgroundImage()
      })
    },

    getZoneLongName() {
      return this.zone.long_name
    },

    async loadBackgroundImages() {
      console.log("[EQZoneCardPreview] loadBackgroundImages")

      document.body.style.setProperty("--zone-background", "none");
      document.body.style.setProperty("--zone-background-size", "auto");

      // get zone wallpaper
      await SpireApiClient.v1().get('/assets/zone-images/' + encodeURIComponent(this.zone.long_name)).then((r) => {
        if (r.status === 200) {
          this.backgroundImages = r.data.images
        }
      })
    },

    setBackgroundImage() {
      if (this.backgroundImages && this.backgroundImages.length > 0) {
        const image = this.backgroundImages[Math.floor(Math.random() * this.backgroundImages.length)];
        console.log("IMAGE ", image)

        if (image.length > 0) {
          let img    = new Image();
          img.src    = image;
          img.onload = () => {
            // document.body.style.setProperty("--image", "url(" + image + ")");
            document.body.style.setProperty("--zone-background", "url(" + image + ")");
            document.body.style.setProperty("--zone-background-size", "cover");
          }
        }
      }
    },
  }
}
</script>

<style>
:root {
  --zone-background-size: auto;
  --zone-background: none;
}

#zone-preview::before {
  content: "";

  background-size: var(--zone-background-size) !important;
  background-repeat: no-repeat !important;
  position: absolute;
  z-index: -99999;
  top: 0;
  right: 0;
  bottom: 0;
  left: 0;

  background: var(--zone-background);
  opacity: .2;

  --webkit-transition: background-image 3s ease-in-out;
  transition: background-image 3s ease-in-out;
}
</style>
