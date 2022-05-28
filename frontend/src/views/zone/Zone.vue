<template>
  <content-area>
    <div class="row">
      <div class="col-7">
        <eq-zone-map
          v-if="zone && version"
          :zone="zone"
          :version="version"
          @npc-marker-hover="processNpcMarkerHover"
          @spell-marker-hover="processSpellMarkerHover"
        />
      </div>
      <div class="col-5">

        <eq-window
          :style="'height: 96vh; '"
          id="zone-preview"
          v-if="Object.keys(selectorActive).length === 0 && zoneData"
        >
          <h6 class="eq-header">{{ getZoneLongName() }}</h6>
        </eq-window>

        <!-- NPC -->
        <eq-window
          class="fade-in"
          id="preview-pane"
          style="max-height: 95vh; overflow-y: scroll; overflow-x: hidden"
          v-if="selectorActive['npc-hover'] && npc"
        >
          <eq-npc-card-preview
            :npc="npc"
          />
        </eq-window>

        <!-- Spell -->
        <eq-window
          class="fade-in"
          id="preview-pane"
          style="max-height: 95vh; overflow-y: scroll; overflow-x: hidden"
          v-if="selectorActive['spell-hover'] && spell"
        >
          <eq-spell-preview
            :spell-data="spell"
          />
        </eq-window>

      </div>
    </div>
  </content-area>
</template>

<script>
import ContentArea      from "../../components/layout/ContentArea";
import EqWindow         from "../../components/eq-ui/EQWindow";
import {Navbar}         from "../../app/navbar";
import EqZoneMap        from "../../components/EqZoneMap";
import EqNpcCardPreview from "../../components/preview/EQNpcCardPreview";
import EqSpellPreview   from "../../components/preview/EQSpellCardPreview";
import {Zones}          from "../../app/zones";
import {SpireApiClient} from "../../app/api/spire-api-client";

export default {
  name: "Zone",
  components: { EqSpellPreview, EqNpcCardPreview, EqZoneMap, EqWindow, ContentArea },
  data() {
    return {
      zone: "",
      version: "",

      zoneData: {},

      selectorActive: {},
    }
  },
  beforeDestroy() {
    Navbar.expand()

    if (this.interval) {
      clearInterval(this.interval)
    }
  },
  created() {
    this.npc = {}
    this.backgroundImages = []

    // cycle background images
    this.interval = setInterval(this.setBackgroundImage, 10 * 1000)
  },
  watch: {
    '$route'() {
      console.log("route trigger")
      this.init()
    },
  },

  mounted() {
    this.init()
  },

  methods: {
    async init() {
      this.npc   = {}
      this.spell = {}
      this.resetSelectors()

      Navbar.collapse()

      // pull from router
      this.zone    = this.$route.params.zone
      this.version = this.$route.query.v

      // get zone data
      this.zoneData = (await Zones.getZoneByShortName(this.zone))

      // get zone wallpaper
      this.loadBackgroundImages().then(() => {
        this.setBackgroundImage()
      })
    },

    async loadBackgroundImages() {
      document.body.style.setProperty("--zone-background", "none");
      document.body.style.setProperty("--zone-background-size", "auto");

      // get zone wallpaper
      await SpireApiClient.v1().get('/assets/zone-images/' + encodeURIComponent(this.zoneData.long_name)).then((r) => {
        if (r.status === 200) {
          this.backgroundImages = r.data.images
        }
      })
    },

    setBackgroundImage() {
      if (this.backgroundImages) {
        const image = this.backgroundImages[Math.floor(Math.random() * this.backgroundImages.length)];
        console.log("IMAGE ", image)

        if (image.length > 0) {
          let img    = new Image();
          img.src    = image;
          img.onload = () => {
            document.body.style.setProperty("--zone-background", "url(" + image + ")");
            document.body.style.setProperty("--zone-background-size", "cover");
          }
        }
      }
    },

    resetSelectors() {
      for (const [k, v] of Object.entries(this.selectorActive)) {
        this.selectorActive[k] = false
      }
    },

    getZoneLongName() {
      return this.zoneData.long_name
    },

    setSelectorActive(selector) {
      this.resetSelectors()


      // this.resetPreviewComponents()
      // this.previewTaskActive        = false;
      // this.lastResetTime            = Date.now()
      this.selectorActive[selector] = true
      this.$forceUpdate()

      // EditFormFieldUtil.setFieldSubEditorHighlightedById(selector)
    },

    processSpellMarkerHover(s) {
      this.spell = {}
      this.setSelectorActive("spell-hover")
      this.spell = s

      // reset preview pane scroll to top
      const t = document.getElementById("preview-pane");
      if (t) {
        t.scrollTop = 0;
      }
    },

    processNpcMarkerHover(n) {
      this.npc = {}
      this.setSelectorActive("npc-hover")
      this.npc = n

      // reset preview pane scroll to top
      const t = document.getElementById("preview-pane");
      if (t) {
        t.scrollTop = 0;
      }
    }
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
  /*background-size: cover !important;*/
  background-size: var(--zone-background-size) !important;
  background-repeat: no-repeat !important;
  position: absolute;
  z-index: -99999;
  top: 0;
  right: 0;
  bottom: 0;
  left: 0;
  /*background: url(https://everquest.allakhazam.com/scenery/halas-mcdaniels.jpg);*/
  background: var(--zone-background);
  opacity: .2;

  --webkit-transition: background-image 3s ease-in-out;
  transition: background-image 3s ease-in-out;

  /*animation: fadeIn 3s;*/
}

</style>
