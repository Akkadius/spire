<template>
  <eq-window
    id="loot-card-preview"
    v-if="loot"
    class="p-0"
  >
    <div class="p-3">

      <!--      {{loot}}-->

      <div class="mb-3">
        <span class="font-weight-bold">Loot Table</span> ({{ loot.id }}) ({{ loot.name }})
        <div class="ml-3 mt-3" v-if="loot.mincash && loot.maxcash">
          <div>
            <span class="font-weight-bold mr-1">Min Cash</span>
            <eq-cash-display :price="loot.mincash"/>
          </div>
          <span class="font-weight-bold mr-1">Max Cash</span>
          <eq-cash-display :price="loot.maxcash"/>
        </div>
      </div>
      <div v-for="le in loot.loottable_entries">

        <div class="mt-1 mb-1">
          <span class="font-weight-bold">Loot Drop</span> ({{ le.lootdrop.id }}) ({{ le.lootdrop.name }})
          <div class="ml-3 mt-1">
            <span class="font-weight-bold">Probability</span> ({{ le.probability }}%)
            <span class="font-weight-bold">Multiplier</span> ({{ le.multiplier }}x)
            <span class="font-weight-bold">Drop Limit</span> ({{ le.droplimit }})
          </div>
        </div>

        <div
          class="ml-5"
          v-for="lde in le.lootdrop.lootdrop_entries"
          v-if="le.lootdrop.lootdrop_entries"
        >
          <item-popover
            :item="lde.item"
          >
            <span class="font-weight-bold">Chance</span> {{ lde.chance }}%
            <span class="font-weight-bold">Multiplier</span> {{ lde.multiplier }}x
          </item-popover>
        </div>
      </div>


      <!--      <eq-debug :data="loot"></eq-debug>-->


    </div>

  </eq-window>
</template>

<script>

import EqWindow           from "../eq-ui/EQWindow";
import EqTabs             from "../eq-ui/EQTabs";
import EqTab              from "../eq-ui/EQTab";
import EqCheckbox         from "../eq-ui/EQCheckbox";
import NpcPopover         from "../NpcPopover";
import LoaderFakeProgress from "../LoaderFakeProgress";
import EqDebug            from "../eq-ui/EQDebug";
import ItemPopover        from "../ItemPopover";
import EqCashDisplay      from "../eq-ui/EqCashDisplay";

export default {
  name: "EqLootCardPreview",
  components: {
    EqCashDisplay,
    ItemPopover,
    EqDebug,
    LoaderFakeProgress,
    NpcPopover,
    EqCheckbox,
    EqTab,
    EqTabs,
    EqWindow
  },
  props: {
    loot: Object,
    required: true,
  },

  created() {

  },
  mounted() {
    this.init()
  },
  methods: {
    init() {
      console.log(this.loot)
    }
  }
}
</script>

<style>

</style>
