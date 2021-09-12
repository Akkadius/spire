<template>
  <div class="container-fluid">

    <div class="row justify-content-center">
      <div class="col-12 col-lg-10 col-xl-10 content-pop">
        <!--        <page-header title="Components" pre-title="Preview Components"/>-->

        <div class="container-fluid">

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

            <div class="panel panel-default ">
              <div class="row">


                <div class="col-6">
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
                <div v-for="(item, index) in items" :key="index" style="display: inline-block; vertical-align: top">
                  <eq-window style="margin-right: 10px; width: auto; height: 90%">
                    <eq-item-preview :item-data="item"/>
                  </eq-window>
                </div>
              </div>

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
                <div v-for="(spell, index) in spells" :key="index" style="display: inline-block; vertical-align: top">
                  <eq-window style="margin-right: 10px; width: auto; height: 90%">
                    <eq-spell-preview :spell-data="spell"/>
                  </eq-window>
                </div>
              </div>


              <!-- Spell Preview -->
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
                <npc-special-abilities
                  show-special-abilities-result="1"
                  :inputData.sync="specialAbilityInput"
                  :abilities="specialAbilityInput"/>

                <eq-window class="eq-window" style="width: 100%">
                  <h4 class="eq-header">Special Abilities Input</h4>
                  <input
                    type="text"
                    class="form-control mb-5 mt-4"
                    v-model="specialAbilityInput">

                  The input is two-way bound with the special abilities editor
                </eq-window>
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
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
// @ is an alias to /src

import EqWindowSimple from "@/components/eq-ui/EQWindowSimple";
import EqWindowFancy from "@/components/eq-ui/EQWindowFancy";
import EqWindowComplex from "@/components/eq-ui/EQWindowComplex";
import EqWindow from "@/components/eq-ui/EQWindow";
import EqTabs from "@/components/eq-ui/EQTabs";
import EqTab from "@/components/eq-ui/EQTab";
import EqProgressBar from "@/components/eq-ui/EQProgressBar";
import PageHeader from "@/views/layout/PageHeader";
import EqItemPreview from "@/components/eq-ui/EQItemPreview";
import {EXAMPLE_ITEM_DATA} from "@/app/constants/eq-example-item-data";
import EqSpellPreview from "@/components/eq-ui/EQSpellPreview";
import {EXAMPLE_SPELL_DATA} from "@/app/constants/eq-example-spell-data";
import NpcSpecialAbilities from "@/components/tools/NpcSpecialAbilities";

export default {
  name: "Home",
  components: {
    NpcSpecialAbilities,
    EqSpellPreview,
    EqItemPreview,
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
      items: EXAMPLE_ITEM_DATA,
      spells: EXAMPLE_SPELL_DATA,
      specialAbilityInput: "1,1,3000,50^2,1,1,1000,2340^3,1,20,0,0,0,0,100,0^4,1,0,100,0,0,0,100,0^11,1,4,150,0,0,5^29,1,50^40,1,10,10,100^7,1^10,1^14,1^19,1^22,1^25,1^26,1"
    }
  },
  mounted() {
    setInterval(() => {
      this.blueProgress   = Math.floor(Math.random() * 100 + 1)
      this.redProgress    = Math.floor(Math.random() * 100 + 1)
      this.orangeProgress = Math.floor(Math.random() * 100 + 1)
      this.yellowProgress = Math.floor(Math.random() * 100 + 1)
    }, 1000);

    setTimeout(() => this.scrollFix(this.$route.hash), 1);
  },
  methods: {
    scrollFix: function (hashbang) {
      location.hash = hashbang;
      document.getElementById(hashbang.replace("#", "")).scrollIntoView();
    }

  },
  watch: {
    "$route.hash": function () {
      setTimeout(() => this.scrollFix(this.$route.hash), 1);
    }
  }
}
</script>
