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
import EqNpcCardPreview from "../../components/eq-ui/EQNpcCardPreview";
import EqSpellPreview   from "../../components/eq-ui/EQSpellCardPreview";

export default {
  name: "Zone",
  components: { EqSpellPreview, EqNpcCardPreview, EqZoneMap, EqWindow, ContentArea },
  data() {
    return {
      zone: "",
      version: "",

      selectorActive: {},
    }
  },
  beforeDestroy() {
    Navbar.expand()
  },
  created() {
    this.npc = {}
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
    init() {
      this.npc = {}
      this.spell = {}
      this.resetSelectors()

      Navbar.collapse()

      // pull from router
      this.zone    = this.$route.params.zone
      this.version = this.$route.query.v
    },

    resetSelectors() {
      for (const [k, v] of Object.entries(this.selectorActive)) {
        this.selectorActive[k] = false
      }
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

<style scoped>

</style>
