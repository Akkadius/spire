<template>
  <content-area>
    <div class="row">
      <div class="col-7">
        <eq-zone-map
          v-if="zone && version"
          :zone="zone"
          :version="version"
          @npc-marker-hover="processNpcMarkerHover"
        />
      </div>
      <div class="col-5">
        <eq-window
          class="fade-in"
          style="max-height: 95vh; overflow-y: scroll; overflow-x: hidden">
          <eq-npc-card-preview v-if="selectorActive['npc-hover'] && npc"
            :npc="npc"
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

export default {
  name: "Zone",
  components: { EqNpcCardPreview, EqZoneMap, EqWindow, ContentArea },
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
  mounted() {
    Navbar.collapse()

    // pull from router
    this.zone    = this.$route.params.zone
    this.version = this.$route.query.v

  },
  methods: {
    setSelectorActive(selector) {
      for (const [k, v] of Object.entries(this.selectorActive)) {
        this.selectorActive[k] = false
      }

      // this.resetPreviewComponents()
      // this.previewTaskActive        = false;
      // this.lastResetTime            = Date.now()
      this.selectorActive[selector] = true
      this.$forceUpdate()

      // EditFormFieldUtil.setFieldSubEditorHighlightedById(selector)
    },

    processNpcMarkerHover(n) {
      this.npc = {}
      this.setSelectorActive("npc-hover")
      this.npc = n
    }
  }
}
</script>

<style scoped>

</style>
