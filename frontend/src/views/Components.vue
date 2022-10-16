<template>
  <!--        <page-header title="Components" pre-title="Preview Components"/>-->
  <content-area class="content-pop" style="width: 90%">

    <!-- Faces -->
    <div class="header mt-md-1">
      <div class="header-body">
        <h1 class="header-title" id="facial-appearance">
          Facial Appearance
        </h1>

        <p class="header-subtitle">
          Facial appearance selection component
        </p>
      </div>
    </div>

    <div class="row">
      <div class="col-12 text-center">
        <eq-window title="Facial Appearance">
          <facial-appearance-selector/>
        </eq-window>
      </div>
    </div>


      <!-- Progress Bars -->
    <div class="header mt-md-1">
      <div class="header-body">
        <h1 class="header-title" id="progress-bars">
          Progress Bars
        </h1>

        <p class="header-subtitle">
          Progress bars with different style examples, hp, mana etc.
        </p>
      </div>
    </div>

    <div class="row">
      <div class="col-4 text-center">
        <!-- Window and Progress Bars -->
        <eq-window title="Progress Bars" style="height: 90%">
          Endurance
          <eq-progress-bar :percent="yellowProgress"/>
          Mana
          <eq-progress-bar :percent="blueProgress" color="blue"/>
          Health
          <eq-progress-bar :percent="redProgress" color="red"/>
          Orange
          <eq-progress-bar :percent="orangeProgress" color="orange"/>

        </eq-window>
      </div>

    </div>

    <div class="header mt-md-1">
      <div class="header-body">
        <h1 class="header-title" id="page-headers">
          Page Headers
        </h1>
        <p class="header-subtitle">
          Page headers use Menuetto font along with custom CSS styling
        </p>
      </div>
    </div>

    <div class="row">
      <div class="col-4">
        <!-- Headers -->
        <eq-window style="text-align:center; height: 90%">
          <h1 class="eq-header">Header H1</h1>
          <h2 class="eq-header">Header H2</h2>
          <h3 class="eq-header">Header H3</h3>
          <h4 class="eq-header">Header H4</h4>
          <h5 class="eq-header">Header H5</h5>
          <h6 class="eq-header">Header H6</h6>
        </eq-window>
      </div>
    </div>

    <div class="header mt-md-1">
      <div class="header-body">
        <h1 class="header-title" id="range-visual">
          Range Visual
        </h1>
        <p class="header-subtitle">
          Visualizes in game range units
        </p>
      </div>
    </div>

    <eq-window-simple>
      Range Input
      <b-input v-model.number="rangeVisual" type="number" min="0" max="1000"/>

      <input
        type="range"
        min="0"
        max="1000"
        step="1"
        class="p-0 m-0 mt-2 form-control"
        v-model.number="rangeVisual"
      >
    </eq-window-simple>

    <range-visualizer :unit-marker="rangeVisual" class="mt-3"/>

    <!-- Item Preview -->
    <div class="header mt-md-1">
      <div class="header-body">
        <h1 class="header-title" id="item-preview">
          Item Preview
        </h1>

        <p class="header-subtitle">
          Item Preview can be used anywhere, it takes in an item data payload and properly formats the item
          display. This component requires API calls to be made for fetching Spell data
        </p>
      </div>
    </div>

    <div class="row">
      <!-- Item Preview -->
      <div
        v-for="(item, index) in items"
        class="mb-3 col-5"
        :key="index"
        style="display: inline-block; vertical-align: top"
      >
        <eq-window style="width: auto; height: 100%">
          <eq-item-card-preview :item-data="item"/>
        </eq-window>
      </div>
    </div>

    <!-- Spell Icons -->
    <div class="header mt-md-1">
      <div class="header-body">
        <h1 class="header-title" id="spell-icons">
          Spell Icons
        </h1>

      </div>
    </div>

    <eq-window-simple>
      <div class="row">
        <div class="col-1 text-center" v-for="n in 12">
          <div v-for="icon in spellIcons">
            {{ icon.iconSize }}

            <span
              :style="'width: ' + icon.iconSize + 'px; height: ' + icon.iconSize + 'px; border: 1px solid; border-radius: 3px; display: inline-block'"
              :class="'spell-' + (parseInt(icon.icon) + n) + '-' + icon.iconSize + ' mr-1 mt-1'"
            />
          </div>
        </div>
      </div>

    </eq-window-simple>

    <!-- Spell Preview -->
    <div class="header mt-md-1">
      <div class="header-body">
        <h1 class="header-title" id="spell-preview">
          Spell Preview
        </h1>

        <p class="header-subtitle">
          Spell Preview can be used anywhere, it takes in an spells_new data payload and properly formats the
          item display
        </p>
      </div>
    </div>

    <div class="row">
      <!-- Spell Preview -->
      <div
        v-for="(spell, index) in spells"
        :key="index"
        style="display: inline-block; vertical-align: top"
        class="col-5 mb-3"
      >
        <eq-window style="width: auto; height: 100%">
          <eq-spell-preview :spell-data="spell"/>
        </eq-window>
      </div>
    </div>

    <div class="header mt-md-1">
      <div class="header-body">
        <h1 class="header-title" id="npc-card-preview">
          NPC Card Preview
        </h1>

        <p class="header-subtitle">
          NPC Preview Component
        </p>
      </div>
    </div>

    <div class="row">
      <div
        v-for="npc in npcs"
        :key="npc.id"
        style="display: inline-block; vertical-align: top"
        class="col-6 mb-5"
      >
        <eq-window
          style="margin-right: 10px; width: auto; height: 100%"
        >
          <eq-npc-card-preview
            :npc="npc"
          />
        </eq-window>
      </div>
    </div>

    <div class="header mt-md-1">
      <div class="header-body">
        <h1 class="header-title" id="npc-special-abilities">
          NPC Special Abilities
        </h1>

        <p class="header-subtitle">
          NPC Special Abilities editor component
        </p>
      </div>
    </div>

    <div class="row">
      <div class="col-12">
        <eq-window style="width: 100%">
          <npc-special-abilities
            :show-special-abilities-result="true"
            :inputData.sync="specialAbilityInput"
            :abilities="specialAbilityInput"
          />

          <h4 class="eq-header mt-3">Special Abilities Input</h4>
          <input
            type="text"
            class="form-control mb-3"
            v-model="specialAbilityInput"
          >

          The input is two-way bound with the special abilities editor
        </eq-window>
      </div>
    </div>

    <!-- Table -->
    <div class="header mt-md-1">
      <div class="header-body">
        <h1 class="header-title" id="table">
          Table
        </h1>
      </div>
    </div>

    <!-- Table Example -->
    <eq-window title="Table Example">
      <div class='eq-window-nested-blue'>
        <table id="tabbox1" class="eq-table eq-highlight-rows" style="display: table;">
          <thead>
          <tr>
            <th>Title</th>
            <th>Cur/Max</th>
            <th>Cost</th>
          </tr>
          </thead>
          <tbody>
          <tr>
            <td>Baking Mastery</td>
            <td>0 / 3</td>
            <td>3</td>
          </tr>
          <tr>
            <td>Bandage Wounds</td>
            <td>0 / 5</td>
            <td>3</td>
          </tr>
          <tr>
            <td>Battle Ready</td>
            <td>0 / 4</td>
            <td>5</td>
          </tr>
          <tr>
            <td>Blacksmithing Mastery</td>
            <td>0 / 3</td>
            <td>3</td>
          </tr>
          <tr>
            <td>Brewing Mastery</td>
            <td>0 / 3</td>
            <td>3</td>
          </tr>
          <tr>
            <td>Combat Agility</td>
            <td>0 / 40</td>
            <td>2</td>
          </tr>
          <tr>
            <td>Combat Medic</td>
            <td>0 / 5</td>
            <td>3</td>
          </tr>
          <tr>
            <td>Combat Stability</td>
            <td>0 / 39</td>
            <td>2</td>
          </tr>
          <tr>
            <td>Delay Death</td>
            <td>0 / 31</td>
            <td>3</td>
          </tr>
          <tr>
            <td>Discordant Defiance</td>
            <td>0 / 16</td>
            <td>5</td>
          </tr>
          <tr>
            <td>Earthen Alacrity</td>
            <td>0 / 16</td>
            <td>5</td>
          </tr>
          <tr>
            <td>Earthen Allure</td>
            <td>0 / 16</td>
            <td>5</td>
          </tr>
          <tr>
            <td>Earthen Artistry</td>
            <td>0 / 16</td>
            <td>5</td>
          </tr>
          <tr>
            <td>Earthen Brawn</td>
            <td>0 / 16</td>
            <td>5</td>
          </tr>
          </tbody>
        </table>
      </div>
    </eq-window>

    <!-- Tabs -->
    <div class="header mt-md-1">
      <div class="header-body">
        <h1 class="header-title" id="tabs">
          Tabs
        </h1>
      </div>
    </div>

    <!-- Tabs -->
    <eq-window title="Tabs">
      <eq-tabs>
        <eq-tab name="Kunark" :selected="true">
          <h2>Has Kunark Things</h2>

          Like frogs... lots of them!
        </eq-tab>
        <eq-tab name="Velious">
          <h2>Has Velious Things</h2>

          Like dragons, giants and coldain!
        </eq-tab>
        <eq-tab name="Luclin">
          <h2>Has Luclin Things</h2>

          We're on the moon baby!
        </eq-tab>
      </eq-tabs>
    </eq-window>

    <!-- Tabs -->
    <div class="header mt-md-1">
      <div class="header-body">
        <h1 class="header-title" id="form-elements">
          Form Elements
        </h1>
      </div>
    </div>

    <eq-window title="Form Elements">

      <table style="align-content: center; margin: auto">
        <tr>
          <td style="vertical-align: middle">
            Checkbox 1
          </td>
          <td>
            <input type="checkbox" id="checkbox-1" checked class='eq-checkbox'>
            <label for="checkbox-1"></label>
          </td>
        </tr>
        <tr>
          <td style="vertical-align: middle">
            Checkbox 2
          </td>
          <td>
            <input type="checkbox" id="checkbox-2" class='eq-checkbox'>
            <label for="checkbox-2"></label>
          </td>
        </tr>
        <tr>
          <td style="vertical-align: middle">
            Checkbox 3
          </td>
          <td>
            <input type="checkbox" id="checkbox-3" class='eq-checkbox'>
            <label for="checkbox-3"></label>
          </td>
        </tr>
      </table>

      <div class="form-group">
        <input id="name" type="text" name="name" class="eq-input form-control" value="Input value">
      </div>

      <div class="form-group">
        <select class="eq-input form-control">
          <option>Option 1</option>
          <option>Option 2</option>
          <option>Option 3</option>
        </select>
      </div>

      <div class="form-group">
        <button class='eq-button-wrap' onclick="alert('click')">Large</button>
      </div>

      <div class="form-group">
        <button class='eq-button' onclick="alert('click')">Centered</button>
      </div>

    </eq-window>

    <!-- Tabs -->
    <div class="header mt-md-1">
      <div class="header-body">
        <h1 class="header-title" id="windows">
          Windows
        </h1>
      </div>
    </div>

    <!-- Temp-->

    <eq-window title="Some Title">With title</eq-window>
    <eq-window>No title</eq-window>

    <eq-window-complex title="EQ Window Complex">
      Window body
    </eq-window-complex>

    <eq-window-complex>
      Window body no title
    </eq-window-complex>

    <eq-window-fancy title="EQ Window Fancy">
      Window body
    </eq-window-fancy>

    <eq-window-fancy>
      Window body no title
    </eq-window-fancy>

    <eq-window-fancy title="Sub windows">
      <div class='eq-window-nested-tan'>
        <div class='WindowNestedTanTitleBar'>
          Sub Window Title #1
        </div>
        Sub Window Body
      </div>
      <div class='eq-window-nested-tan'>
        <div class='WindowNestedTanTitleBar'>
          Sub Window Title #2
        </div>
        Sub Window Body
      </div>
      <div class='eq-window-nested-blue'>
        Blue Box Body
      </div>
    </eq-window-fancy>

    <eq-window-simple title="Window Simple Title">
      Window body
    </eq-window-simple>

    <eq-window-simple>
      Window body simple no title
    </eq-window-simple>
  </content-area>
</template>

<script>
// @ is an alias to /src

import EqWindowSimple       from "@/components/eq-ui/EQWindowSimple";
import EqWindowFancy        from "@/components/eq-ui/EQWindowFancy";
import EqWindowComplex      from "@/components/eq-ui/EQWindowComplex";
import EqWindow             from "@/components/eq-ui/EQWindow";
import EqTabs               from "@/components/eq-ui/EQTabs";
import EqTab                from "@/components/eq-ui/EQTab";
import EqProgressBar        from "@/components/eq-ui/EQProgressBar";
import PageHeader           from "@/components/layout/PageHeader";
import EqItemCardPreview        from "@/components/preview/EQItemCardPreview";
import {EXAMPLE_ITEM_DATA}      from "@/app/constants/eq-example-item-data";
import EqSpellPreview           from "@/components/preview/EQSpellCardPreview";
import {EXAMPLE_SPELL_DATA}     from "@/app/constants/eq-example-spell-data";
import NpcSpecialAbilities      from "@/components/tools/NpcSpecialAbilities";
import RangeVisualizer          from "../components/tools/RangeVisualizer";
import EqNpcCardPreview         from "../components/preview/EQNpcCardPreview";
import {NpcTypeApi}        from "../app/api";
import {SpireApi}          from "../app/api/spire-api";
import {SpireQueryBuilder} from "../app/api/spire-query-builder";
import ContentArea              from "../components/layout/ContentArea";
import FacialAppearanceSelector from "../components/selectors/FacialAppearanceSelector";

export default {
  name: "Home",
  components: {
    FacialAppearanceSelector,
    ContentArea,
    EqNpcCardPreview,
    RangeVisualizer,
    NpcSpecialAbilities,
    EqSpellPreview,
    EqItemCardPreview,
    PageHeader,
    EqProgressBar,
    EqTab,
    EqTabs,
    EqWindow,
    EqWindowComplex,
    EqWindowFancy,
    EqWindowSimple
  },
  data() {
    return {
      blueProgress: 50,
      orangeProgress: 50,
      redProgress: 50,
      yellowProgress: 50,
      rangeVisual: 157,
      items: EXAMPLE_ITEM_DATA,
      spells: EXAMPLE_SPELL_DATA,
      // npcs: {},
      specialAbilityInput: "1,1,3000,50^2,1,1,1000,2340^3,1,20,0,0,0,0,100,0^4,1,0,100,0,0,0,100,0^11,1,4,150,0,0,5^29,1,50^40,1,10,10,100^7,1^10,1^14,1^19,1^22,1^25,1^26,1",

      spellIcons: [
        { icon: 50, iconSize: 12, },
        { icon: 50, iconSize: 20, },
        { icon: 50, iconSize: 30, },
        { icon: 50, iconSize: 40, },
      ],
    }
  },
  created() {
    this.npcs = {}
  },
  async mounted() {
    setInterval(() => {
      this.blueProgress   = Math.floor(Math.random() * 100 + 1)
      this.redProgress    = Math.floor(Math.random() * 100 + 1)
      this.orangeProgress = Math.floor(Math.random() * 100 + 1)
      this.yellowProgress = Math.floor(Math.random() * 100 + 1)
    }, 1000);

    // load NPCS with relationships
    let npcIDs = [
      32040,
      189009,
      30073,
      159691,
      702065,
      511,
      68135,
      150206,
    ]

    let builder = (new SpireQueryBuilder())

    builder.includes(
      [
        "NpcSpell.NpcSpellsEntries.SpellsNew",
        "NpcFactions.NpcFactionEntries.FactionList",
        "NpcFactions",
        "NpcEmotes",
        "Merchantlists.Items",
        "Loottable.LoottableEntries.LootdropEntries.Item"
      ]
    )
    builder.orderBy(["id"])

    const response = await (new NpcTypeApi(...SpireApi.cfg())).getNpcTypesBulk({
        body: {
          ids: npcIDs
        },
      },
      {
        query: builder.get()
      }
    )

    if (response.status === 200 && response.data && parseInt(response.data.length) > 0) {
      this.npcs = response.data
    }

    this.$forceUpdate()

  },
}
</script>

<style>
.header {
  margin-bottom: 0px;
}
</style>
