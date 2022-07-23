<template>
  <div>
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

        <!-- Zone Card -->
        <eq-zone-card-preview
          style="height: 96vh"
          v-show="selectorActive['zone-preview'] && zoneData"
          :zone="zoneData"
        />

        <eq-window
          v-if="!isZoneCardActive()"
          class="text-center"
        >
          <b-button
            class="btn-dark btn-sm btn-outline-warning"
            @click="setSelectorActive('zone-preview', true)"
          >
            <i class="fa fa-chevron-up"></i> Return to Zone
          </b-button>
        </eq-window>

        <!-- NPC -->
        <eq-window
          class="fade-in"
          id="preview-pane"
          :style="'max-height: ' + (isZoneCardActive() ? '95' : '87') + 'vh; overflow-y: scroll; overflow-x: hidden'"
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
          :style="'max-height: ' + (isZoneCardActive() ? '95' : '87') + 'vh; overflow-y: scroll; overflow-x: hidden'"
          v-if="selectorActive['spell-hover'] && spell"
        >
          <eq-spell-preview
            :spell-data="spell"
          />
        </eq-window>

      </div>
    </div>
  </div>
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
import {EventBus}        from "../../app/event-bus/event-bus";

const MILLISECONDS_BEFORE_WINDOW_RESET = 5000;

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

    // if (this.zonePreviewInterval) {
    //   clearInterval(this.zonePreviewInterval)
    // }

    EventBus.$off("NPC_SHOW_CARD", this.handleNpcShowCardEvent);
  },
  created() {
    this.npc           = {}
    this.lastResetTime = Date.now()

    // this.zonePreviewInterval = setInterval(this.previewZone, 1000)

    EventBus.$on("NPC_SHOW_CARD", this.handleNpcShowCardEvent);
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

    isZoneCardActive() {
      return Object.keys(this.selectorActive).length > 0 && this.selectorActive['zone-preview']
    },

    // from zone preview card -> zone
    handleNpcShowCardEvent(e) {
      this.processNpcMarkerHover(e)
    },

    previewZone() {
      console.log("[Zone] previewZone trigger")
      this.setSelectorActive('zone-preview')
    },

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

      this.setSelectorActive('zone-preview', true)
    },

    shouldReset() {
      return (Date.now() - this.lastResetTime) > MILLISECONDS_BEFORE_WINDOW_RESET
    },

    resetSelectors() {
      for (const [k, v] of Object.entries(this.selectorActive)) {
        this.selectorActive[k] = false
      }
    },

    setSelectorActive(selector, force = false) {
      if (this.selectorActive[selector] && !force) {
        // console.log(
        //   "[Zone] setSelectorActive. Selector [%s] is already active",
        //   selector
        // )
        return
      }

      if (this.shouldReset() || force) {
        this.lastResetTime = Date.now()
        this.resetSelectors()
        this.selectorActive[selector] = true
        this.$forceUpdate()
        return
      }

      console.log(
        "[Zone] Tried to set selector [%s] but reset time was not met (%s) ms remaining",
        selector,
        MILLISECONDS_BEFORE_WINDOW_RESET - (Date.now() - this.lastResetTime)
      )

      // EditFormFieldUtil.setFieldSubEditorHighlightedById(selector)
    },

    processSpellMarkerHover(s) {
      this.spell = {}
      this.setSelectorActive("spell-hover", true)
      this.spell = s

      // reset preview pane scroll to top
      const t = document.getElementById("preview-pane");
      if (t) {
        t.scrollTop = 0;
      }
    },

    processNpcMarkerHover(n) {
      this.npc = {}
      this.setSelectorActive("npc-hover", true)
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
