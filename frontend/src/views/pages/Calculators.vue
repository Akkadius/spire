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
import EqSpellPreview from "@/components/eq-ui/EQSpellPreview";
import NpcSpecialAbilities from "@/components/tools/NpcSpecialAbilities";
import HeaderComponent from "../layout/HeaderComponent";
import RaceBitmaskCalculator from "../../components/tools/RaceBitmaskCalculator";
import ClassBitmaskCalculator from "../../components/tools/ClassBitmaskCalculator";

export default {
  name: "Calculators",
  components: {
    ClassBitmaskCalculator,
    RaceBitmaskCalculator,
    HeaderComponent,
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
      racesBitmask: "1859",
      classesBitmask: "37349"
    }
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
