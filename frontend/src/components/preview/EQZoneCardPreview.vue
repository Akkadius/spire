<template>
  <eq-window
    id="zone-preview"
    v-if="zone"
  >
    <h6 class="eq-header">{{ getZoneLongName() }}</h6>

    <eq-tabs>
      <eq-tab name="Tab A"></eq-tab>
      <eq-tab name="Tab B"></eq-tab>
      <eq-tab name="Tab C"></eq-tab>
      <eq-tab name="Tab D"></eq-tab>
    </eq-tabs>
  </eq-window>
</template>

<script>

import EqWindow         from "../eq-ui/EQWindow";
import {SpireApiClient} from "../../app/api/spire-api-client";
import EqTabs           from "../eq-ui/EQTabs";
import EqTab            from "../eq-ui/EQTab";
export default {
  name: "EqZoneCardPreview",
  components: { EqTab, EqTabs, EqWindow },
  props: {
    zone: Object,
    required: true,
  },

  created() {
    this.backgroundImages = []
    this.currentImageIndex = 0

    // cycle background images
    this.interval = setInterval(this.setBackgroundImage, 3 * 1000)
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

    shuffle(array) {
      let currentIndex = array.length,  randomIndex;

      // While there remain elements to shuffle.
      while (currentIndex !== 0) {

        // Pick a remaining element.
        randomIndex = Math.floor(Math.random() * currentIndex);
        currentIndex--;

        // And swap it with the current element.
        [array[currentIndex], array[randomIndex]] = [
          array[randomIndex], array[currentIndex]];
      }

      return array;
    },

    async loadBackgroundImages() {
      console.log("[EQZoneCardPreview] loadBackgroundImages")

      document.body.style.setProperty("--zone-background", "none");
      document.body.style.setProperty("--zone-background-size", "auto");

      // get zone wallpaper
      await SpireApiClient.v1().get('/assets/zone-images/' + encodeURIComponent(this.zone.long_name)).then((r) => {
        if (r.status === 200) {
          this.backgroundImages = this.shuffle(r.data.images)
        }
      })
    },

    setBackgroundImage() {
      if (this.backgroundImages && this.backgroundImages.length > 0) {
        const image = this.backgroundImages[this.currentImageIndex];
        console.log("IMAGE ", image)


        console.log(
          "[EQZoneCardPreview] loadBackgroundImages Playing index [%s] out of [%s]",
          this.currentImageIndex,
          this.backgroundImages.length
        )

        if (image.length > 0) {
          let img    = new Image();
          img.src    = image;
          img.onload = () => {
            // document.body.style.setProperty("--image", "url(" + image + ")");
            document.body.style.setProperty("--zone-background", "url(" + image + ")");
            document.body.style.setProperty("--zone-background-size", "cover");

            // increment
            this.currentImageIndex++;

            // reset if rollover
            if (this.currentImageIndex >= this.backgroundImages.length) {
              console.log("[EQZoneCardPreview] loadBackgroundImages resetting")
              this.currentImageIndex = 0;
            }
          }
          img.onerror = () => {
            console.log(
              "[EQZoneCardPreview] loadBackgroundImages Failed to load index [%s] out of [%s]",
              this.currentImageIndex,
              this.backgroundImages.length
            )

            this.currentImageIndex++
            this.setBackgroundImage()
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

  --webkit-transition: background-image 1s ease-in-out;
  transition: background-image 1s ease-in-out;
}
</style>
