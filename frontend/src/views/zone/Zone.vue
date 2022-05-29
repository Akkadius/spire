<template>
  <content-area>
    <div class="row" @mouseover="previewZone">
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

        <eq-zone-card-preview
          style="height: 96vh"
          v-if="selectorActive['zone-preview'] && zoneData"
          :zone="zoneData"
        />

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
import ContentArea       from "../../components/layout/ContentArea";
import EqWindow          from "../../components/eq-ui/EQWindow";
import {Navbar}          from "../../app/navbar";
import EqZoneMap         from "../../components/EqZoneMap";
import EqNpcCardPreview  from "../../components/preview/EQNpcCardPreview";
import EqSpellPreview    from "../../components/preview/EQSpellCardPreview";
import {Zones}           from "../../app/zones";
import EqZoneCardPreview from "../../components/preview/EQZoneCardPreview";
import {debounce}        from "../../app/utility/debounce";

export default {
  name: "Zone",
  components: { EqZoneCardPreview, EqSpellPreview, EqNpcCardPreview, EqZoneMap, EqWindow, ContentArea },
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
  },
  created() {
    this.npc              = {}
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
    previewZone: debounce(function() {
      console.log("preview zone trigger")
      this.setSelectorActive('zone-preview')
    }, 500),

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

      this.setSelectorActive('zone-preview')
    },

    resetSelectors() {
      for (const [k, v] of Object.entries(this.selectorActive)) {
        this.selectorActive[k] = false
      }
    },

    setSelectorActive(selector) {
      this.resetSelectors()
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


/*.map-tiles::after {*/
/*  content: "";*/
/*  background-size: cover !important;*/
/*  background-repeat: no-repeat !important;*/
/*  position: absolute;*/
/*  z-index: 1;*/
/*  top: 0;*/
/*  right: 0;*/
/*  bottom: 0;*/
/*  left: 0;*/
/*  background: var(--zone-background) !important;*/
/*  opacity: .1;*/

/*  --webkit-transition: background-image 3s ease-in-out;*/
/*  transition: background-image 3s ease-in-out;*/

/*  !*animation: fadeIn 3s;*!*/
/*}*/

</style>
