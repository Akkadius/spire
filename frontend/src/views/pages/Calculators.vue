<template>
  <div class="container-fluid">

    <div class="row justify-content-center">
      <div class="col-12 col-lg-10 col-xl-10 content-pop">
        <div class="container-fluid">

          <!-- Race Bitmask Calculator -->
          <header-component
            header="Race Bitmask Calculator"
            sub-header="Computes the bitmask for selected races..."
          />

          <div class="row">
            <div class="col-12">
              <eq-window style="width: 500px">
                <race-bitmask-calculator
                  :inputData.sync="racesBitmask"
                  :mask="racesBitmask"/>

                <div class="row">
                  <div class="col-12">
                    <h4 class="eq-header mt-2">Race Bitmask</h4>

                    The input is two-way bound with the race selector
                    <input
                      type="text"
                      class="form-control mb-3 mt-3"
                      v-model="racesBitmask">
                  </div>
                </div>
              </eq-window>
            </div>
          </div>

          <!-- Class Bitmask Calculator -->
          <header-component
            header="Class Bitmask Calculator"
            sub-header="Computes the bitmask for selected classes..."
          />

          <div class="row">
            <div class="col-12">
              <eq-window style="width: 500px">
                <class-bitmask-calculator
                  :inputData.sync="classesBitmask"
                  :mask="classesBitmask"/>

                <div class="row">
                  <div class="col-12">
                    <h4 class="eq-header mt-2">Class Bitmask</h4>
                    The input is two-way bound with the class selector
                    <input
                      type="text"
                      class="form-control mb-3 mt-3"
                      v-model="classesBitmask">
                  </div>
                </div>
              </eq-window>
            </div>
          </div>

          <!-- Deity Bitmask Calculator -->
          <header-component
            header="Deity Bitmask Calculator"
            sub-header="Computes the bitmask for selected deities..."
          />

          <div class="row">
            <div class="col-12">
              <eq-window style="width: 500px">
                <deity-bitmask-calculator
                  :inputData.sync="deityBitmask"
                  :mask="deityBitmask"/>

                <div class="row">
                  <div class="col-12">
                    <h4 class="eq-header mt-2">Deity Bitmask</h4>
                    The input is two-way bound with the deity selector
                    <input
                      type="text"
                      class="form-control mb-3 mt-3"
                      v-model="deityBitmask">
                  </div>
                </div>
              </eq-window>
            </div>
          </div>

          <!-- Expansions Bitmask Calculator -->
          <header-component
            header="Expansions Bitmask Calculator"
            sub-header="Computes the bitmask for selected expansions..."
          />

          <div class="row">
            <div class="col-12">
              <eq-window style="width: 500px">
                <expansion-bitmask-calculator
                  :inputData.sync="expansionBitmask"
                  :mask="expansionBitmask"/>

                <div class="row">
                  <div class="col-12">
                    <h4 class="eq-header mt-2">Expansion Bitmask</h4>
                    The input is two-way bound with the expansion selector
                    <input
                      type="text"
                      class="form-control mb-3 mt-3"
                      v-model="expansionBitmask">
                  </div>
                </div>
              </eq-window>
            </div>
          </div>

          <!-- Augment Type Bitmask Calculator -->
          <header-component
            header="Augment Type Calculator"
            sub-header="Computes the bitmask for selected augments..."
          />

          <div class="row">
            <div class="col-12">
              <eq-window style="width: 500px">
                <aug-bitmask-calculator
                  :inputData.sync="augBitmask"
                  :mask="augBitmask"/>

                <div class="row">
                  <div class="col-12">
                    <h4 class="eq-header mt-2">Aug Bitmask</h4>
                    The input is two-way bound with the aug selector
                    <input
                      type="text"
                      class="form-control mb-3 mt-3"
                      v-model="augBitmask">
                  </div>
                </div>
              </eq-window>
            </div>
          </div>

          <!-- Inventory Slot Calculator -->
          <header-component
            header="Inventory Slot Calculator"
            sub-header="Computes outputs based on inputs..."
          />

          <div class="row">
            <div class="col-12">
              <eq-window style="width: 500px">

                <inventory-slot-calculator
                  :inputData.sync="inventorySlotBitmask"
                  :mask="inventorySlotBitmask"
                />

                <div class="row">
                  <div class="col-12">
                    <h4 class="eq-header mt-2">Inventory Slot Bitmask</h4>
                    The input is two-way bound with above
                    <input
                      type="text"
                      class="form-control mb-3 mt-3"
                      v-model="inventorySlotBitmask">
                  </div>
                </div>
              </eq-window>
            </div>
          </div>

          <!-- NPC Special Abilities -->
          <header-component
            header="NPC Special Abilities"
            sub-header="Computes NPC Special abilities..."
          />

          <div class="row">
            <div class="col-12">
              <eq-window style="width: 100%">

                <npc-special-abilities
                  :inputData.sync="specialAbilityInput"
                  :abilities="specialAbilityInput"/>

                <div class="row">
                  <div class="col-12">
                    <h4 class="eq-header mt-2">Special Abilities Code</h4>
                    The input is two-way bound with above
                    <input
                      type="text"
                      class="form-control mb-3 mt-3"
                      v-model="specialAbilityInput">
                  </div>
                </div>
              </eq-window>
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
import EqItemCardPreview from "@/components/eq-ui/EQItemCardPreview";
import EqSpellPreview from "@/components/eq-ui/EQSpellCardPreview";
import NpcSpecialAbilities from "@/components/tools/NpcSpecialAbilities";
import HeaderComponent from "../layout/HeaderComponent";
import RaceBitmaskCalculator from "../../components/tools/RaceBitmaskCalculator";
import ClassBitmaskCalculator from "../../components/tools/ClassBitmaskCalculator";
import InventorySlotCalculator from "../../components/tools/InventorySlotCalculator";
import AugBitmaskCalculator from "../../components/tools/AugmentTypeCalculator";
import DeityBitmaskCalculator from "../../components/tools/DeityCalculator";
import ExpansionBitmaskCalculator from "../../components/tools/ExpansionsCalculator";

export default {
  name: "Calculators",
  components: {
    ExpansionBitmaskCalculator,
    DeityBitmaskCalculator,
    AugBitmaskCalculator,
    InventorySlotCalculator,
    ClassBitmaskCalculator,
    RaceBitmaskCalculator,
    HeaderComponent,
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
      augBitmask: "477396",
      racesBitmask: "1859",
      classesBitmask: "37349",
      inventorySlotBitmask: "1536",
      specialAbilityInput: "1,1,3000,50^2,1,1,1000,2340^3,1,20,0,0,0,0,100,0^4,1,0,100,0,0,0,100,0^11,1,4,150,0,0,5^29,1,50^40,1,10,10,100^7,1^10,1^14,1^19,1^22,1^25,1^26,1",
      deityBitmask: "3079",
      expansionBitmask: "3079"
    }
  },
}
</script>
