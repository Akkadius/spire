<template>
  <eq-window
    id="loot-card-preview"
    :title="`Loot Table (${loot.id}) [${loot.name}]`"
    v-if="loot"
    class="p-0"
  >
    <div class="pb-3 mt-3 p-3" style="max-height: 94vh; overflow-y: scroll; overflow-x: hidden">

      <!--      {{loot}}-->

      <!--      <div class="mb-3 mt-3">-->
      <!--        <span class="font-weight-bold">Loot Table</span> ({{ loot.id }}) ({{ loot.name }})-->
      <!--        <div class="mt-3" v-if="loot.mincash && loot.maxcash">-->
      <!--          <span class="font-weight-bold mr-1">Cash</span>-->
      <!--          <eq-cash-display :price="loot.mincash"/>-->
      <!--          <span class="font-weight-bold mr-2 ml-2">-->
      <!--            <i class="fa fa-chevron-left"/>-->
      <!--            <i class="fa fa-chevron-right"/>-->
      <!--          </span>-->
      <!--          <eq-cash-display :price="loot.maxcash"/>-->
      <!--        </div>-->
      <!--      </div>-->
      <div v-for="le in loot.loottable_entries" class="mt-3">
        <div class="mb-1">
          <span class="font-weight-bold mr-1">Loot Drop Table</span>
          <i class="fa fa-chevron-right mr-1"/>
          {{ le.lootdrop.id }}
          <i class="fa fa-chevron-right mr-1"/>
          {{ le.lootdrop.name }}
          <div class="ml-3 mt-1">
            <i class="fa fa-chevron-down mr-1"/>
            <span class="font-weight-bold">Probability</span> ({{ le.probability }}%)
            <span class="font-weight-bold">Multiplier</span> ({{ le.multiplier }}x)
            <span class="font-weight-bold">Drop Limit</span> ({{ le.droplimit }})
            <span class="font-weight-bold">Total Items</span> ({{ le.lootdrop.lootdrop_entries.length }})
          </div>
        </div>

        <table class="eq-table eq-highlight-rows mt-3">
          <tbody>
          <tr
            v-for="lde in le.lootdrop.lootdrop_entries"
          >
            <td>
              <item-popover
                :item="lde.item"
                size="regular"
                style="min-width: 200px"
                class="d-inline-block ml-1"
              />
            </td>
            <td><span class="font-weight-bold">Chance</span> {{ lde.chance }}%</td>
            <td><span class="font-weight-bold">Multiplier</span> {{ lde.multiplier }}x</td>
          </tr>
          </tbody>
        </table>

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
