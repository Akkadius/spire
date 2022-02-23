<template>
  <eq-window-complex
    title="Scale Stats by Random Range"
    class="minified-inputs mb-0"
  >
    <div class="text-center mb-4">
      <span style="color: #7b714a">Increase by Min / Max Random Range</span>
      <br>
      <br>
      Use <b>Enter Key</b>, <b>Arrow Up/Down Keys</b>
    </div>

    <div class="row">
      <div class="col-4"></div>
      <div class="col-4 text-center">Min</div>
      <div class="col-4 text-center">Max</div>
    </div>

    <div class="row mt-1">
      <div class="col-4 text-center mt-2">
        HP/Mana/End
      </div>
      <div class="col-4 m-0 p-0 pr-1">
        <b-form-input
          v-model.number="hpManaEndMin"
          @change="calcHpManaEnd"
          @focus="highlightHpManaEnd"
          @blur="clearAllHighlights"
          @mouseover="highlightHpManaEnd"
          @mouseleave="clearAllHighlights"
          type="number"
          :step="hpManaEndIncrement"
        />
      </div>
      <div class="col-4 m-0 p-0">
        <b-form-input
          v-model.number="hpManaEndMax"
          @change="calcHpManaEnd"
          @focus="highlightHpManaEnd"
          @blur="clearAllHighlights"
          @mouseover="highlightHpManaEnd"
          @mouseleave="clearAllHighlights"
          type="number"
          :step="hpManaEndIncrement"
        />
      </div>
    </div>

    <div class="row mt-1">
      <div class="col-4 text-center mt-2">
        Stats
      </div>
      <div class="col-4 m-0 p-0 pr-1">
        <b-form-input
          v-model.number="statsMin"
          @change="calcStats"
          @focus="highlightStatFields"
          @blur="clearAllHighlights"
          @mouseover="highlightStatFields"
          @mouseleave="clearAllHighlights"
          type="number"
          :step="statsIncrement"
        />
      </div>
      <div class="col-4 m-0 p-0">
        <b-form-input
          v-model.number="statsMax"
          @change="calcStats"
          @focus="highlightStatFields"
          @blur="clearAllHighlights"
          @mouseover="highlightStatFields"
          @mouseleave="clearAllHighlights"
          type="number"
          :step="statsIncrement"
        />
      </div>
    </div>

    <div class="row mt-1">
      <div class="col-4 text-center mt-2">
        Heroic Stats
      </div>
      <div class="col-4 m-0 p-0 pr-1">
        <b-form-input
          v-model.number="heroicStatsMin"
          @change="calcHeroicStats"
          @focus="highlightHeroicStatsFields"
          @blur="clearAllHighlights"
          @mouseover="highlightHeroicStatsFields"
          @mouseleave="clearAllHighlights"
          type="number"
          :step="heroicStatsIncrement"
        />
      </div>
      <div class="col-4 m-0 p-0">
        <b-form-input
          v-model.number="heroicStatsMax"
          @change="calcHeroicStats"
          @focus="highlightHeroicStatsFields"
          @blur="clearAllHighlights"
          @mouseover="highlightHeroicStatsFields"
          @mouseleave="clearAllHighlights"
          type="number"
          :step="heroicStatsIncrement"
        />
      </div>
    </div>

    <div class="row mt-1">
      <div class="col-4 text-center mt-2">
        Resists
      </div>
      <div class="col-4 m-0 p-0 pr-1">
        <b-form-input
          v-model.number="resistsMin"
          @change="calcResists"
          @focus="highlightResistFields"
          @blur="clearAllHighlights"
          @mouseover="highlightResistFields"
          @mouseleave="clearAllHighlights"
          type="number"
          :step="resistsIncrement"
        />
      </div>
      <div class="col-4 m-0 p-0">
        <b-form-input
          v-model.number="resistsMax"
          @change="calcResists"
          @focus="highlightResistFields"
          @blur="clearAllHighlights"
          @mouseover="highlightResistFields"
          @mouseleave="clearAllHighlights"
          type="number"
          :step="resistsIncrement"
        />
      </div>
    </div>

    <div class="row mt-1">
      <div class="col-4 text-center mt-2">
        Heroic Res.
      </div>
      <div class="col-4 m-0 p-0 pr-1">
        <b-form-input
          v-model.number="heroicResistsMin"
          @change="calcHeroicResists"
          @focus="highlightHeroicResistsFields"
          @blur="clearAllHighlights"
          @mouseover="highlightHeroicResistsFields"
          @mouseleave="clearAllHighlights"
          type="number"
          :step="heroicResistsIncrement"
        />
      </div>
      <div class="col-4 m-0 p-0">
        <b-form-input
          v-model.number="heroicResistsMax"
          @change="calcHeroicResists"
          @focus="highlightHeroicResistsFields"
          @blur="clearAllHighlights"
          @mouseover="highlightHeroicResistsFields"
          @mouseleave="clearAllHighlights"
          type="number"
          :step="heroicResistsIncrement"
        />
      </div>
    </div>

    <div class="row mt-1">
      <div class="col-4 text-center mt-2">
        Mods
      </div>
      <div class="col-4 m-0 p-0 pr-1">
        <b-form-input
          v-model.number="modsMin"
          @change="calcMods"
          @focus="highlightModFields"
          @blur="clearAllHighlights"
          @mouseover="highlightModFields"
          @mouseleave="clearAllHighlights"
          type="number"
          :step="modsIncrement"
        />
      </div>
      <div class="col-4 m-0 p-0">
        <b-form-input
          v-model.number="modsMax"
          @change="calcMods"
          @focus="highlightModFields"
          @blur="clearAllHighlights"
          @mouseover="highlightModFields"
          @mouseleave="clearAllHighlights"
          type="number"
          :step="modsIncrement"
        />
      </div>
    </div>


  </eq-window-complex>

</template>

<script>
import {Items}         from "@/app/items";
import EqWindowComplex from "@/components/eq-ui/EQWindowComplex";

export default {
  name: "ItemStatScaleRange",
  components: { EqWindowComplex },
  created() {
    // hp / mana / end
    this.hpManaEndMin       = this.getMinFromFields(this.getHpManaEndFields());
    this.hpManaEndMax       = this.getMaxFromFields(this.getHpManaEndFields());
    this.hpManaEndIncrement = this.getIncrementFromFields(this.getHpManaEndFields());

    // mods
    this.modsMin       = this.getMinFromFields(this.getMod3Fields());
    this.modsMax       = this.getMaxFromFields(this.getMod3Fields());
    this.modsIncrement = this.getIncrementFromFields(this.getMod3Fields());

    // stats
    this.statsMin       = this.getMinFromFields(this.getStatFields());
    this.statsMax       = this.getMaxFromFields(this.getStatFields());
    this.statsIncrement = this.getIncrementFromFields(this.getStatFields());

    // heroic stats
    this.heroicStatsMin       = this.getMinFromFields(this.getStatFieldsHeroic());
    this.heroicStatsMax       = this.getMaxFromFields(this.getStatFieldsHeroic());
    this.heroicStatsIncrement = this.getIncrementFromFields(this.getStatFieldsHeroic());

    // resists
    this.resistsMin       = this.getMinFromFields(this.getResistFields());
    this.resistsMax       = this.getMaxFromFields(this.getResistFields());
    this.resistsIncrement = this.getIncrementFromFields(this.getResistFields());

    // heroic stats
    this.heroicResistsMin       = this.getMinFromFields(this.getResistHeroicFields());
    this.heroicResistsMax       = this.getMaxFromFields(this.getResistHeroicFields());
    this.heroicResistsIncrement = this.getIncrementFromFields(this.getResistHeroicFields());
  },
  data() {
    return {

      // hp / mana / end
      hpManaEndMin: 0,
      hpManaEndMax: 0,
      hpManaEndIncrement: 0,

      // mods
      modsMin: 0,
      modsMax: 0,
      modsIncrement: 0,

      // stats
      statsMin: 0,
      statsMax: 0,
      statsIncrement: 0,

      // heroicStats
      heroicStatsMin: 0,
      heroicStatsMax: 0,
      heroicStatsIncrement: 0,

      // resists
      resistsMin: 0,
      resistsMax: 0,
      resistsIncrement: 0,

      // heroicResists
      heroicResistsMin: 0,
      heroicResistsMax: 0,
      heroicResistsIncrement: 0,

      // fields
      hpManaEndFields: ['hp', 'mana', 'endur'],


    }
  },
  props: {
    originalItemData: {
      type: Object,
      required: true
    },
  },
  methods: {

    //
    // min / max functions
    //
    getMinFromFields(fields) {
      let minValue = 99999999999999999
      fields.forEach((field) => {
        if (field.value < minValue) {
          minValue = field.value;
        }
      })
      return minValue;
    },

    getMaxFromFields(fields) {
      let maxValue = 0
      fields.forEach((field) => {
        if (field.value > maxValue) {
          maxValue = field.value;
        }
      })
      return maxValue;
    },

    getIncrementFromFields(fields) {
      let values = 0
      let count  = 0
      fields.forEach((field) => {
        values += parseInt(field.value);
        count++;
      })

      // 10% of average for increment
      return Math.round((values / count) / 10);
    },

    //
    // get field functions
    //
    getMod3Fields() {
      let fields = [];
      for (let key in Items.getMod3Fields()) {
        const field = Items.getMod3Fields()[key]
        fields.push({ field: field, value: this.originalItemData[field] })
      }
      return fields;
    },

    getResistFields() {
      let fields = [];
      for (let key in Items.getResistFields()) {
        const entry = Items.getResistFields()[key]
        const field = entry.stat
        fields.push({ field: field, value: this.originalItemData[field] })
      }
      return fields;
    },

    getResistHeroicFields() {
      let fields = [];
      for (let key in Items.getResistFields()) {
        const entry = Items.getResistFields()[key]
        const field = entry.heroic
        fields.push({ field: field, value: this.originalItemData[field] })
      }
      return fields;
    },

    getHpManaEndFields() {
      let fields = [];
      this.hpManaEndFields.forEach((field) => {
        fields.push({ field: field, value: this.originalItemData[field] })
      })
      return fields;
    },

    getStatFields() {
      let fields = [];
      for (let key in Items.getBasicStatFields()) {
        const entry = Items.getBasicStatFields()[key]
        const stat  = entry.stat
        fields.push({ field: stat, value: this.originalItemData[stat] })
      }

      return fields;
    },

    getStatFieldsHeroic() {
      let fields = [];
      for (let key in Items.getBasicStatFields()) {
        const entry  = Items.getBasicStatFields()[key]
        const heroic = entry.heroic
        fields.push({ field: heroic, value: this.originalItemData[heroic] })
      }

      return fields;
    },

    getAllFields() {
      let fields = [];

      fields = fields.concat(
        this.getMod3Fields(),
        this.getHpManaEndFields(),
        this.getStatFields(),
        this.getStatFieldsHeroic(),
        this.getResistFields(),
        this.getResistHeroicFields(),
      )

      return fields;
    },

    randomIntFromInterval(min, max) { // min and max included
      return Math.floor(Math.random() * (max - min + 1) + min)
    },

    //
    // calculator functions
    //
    calcHpManaEnd() {
      this.getHpManaEndFields().forEach((field) => {
        let update = {
          field: field.field,
          value: Math.round(this.randomIntFromInterval(
            this.hpManaEndMin,
            this.hpManaEndMax,
          ))
        }

        if (update.value > 0) {
          this.$emit("field", update);
        }
      })
    },
    calcMods() {
      this.getMod3Fields().forEach((field) => {
        let update = {
          field: field.field,
          value: Math.round(this.randomIntFromInterval(
            this.modsMin,
            this.modsMax,
          ))
        }

        if (update.value > 0) {
          this.$emit("field", update);
        }
      })
    },
    calcStats() {
      this.getStatFields().forEach((field) => {
        let update = {
          field: field.field,
          value: Math.round(this.randomIntFromInterval(
            this.statsMin,
            this.statsMax,
          ))
        }

        if (update.value > 0) {
          this.$emit("field", update);
        }
      })
    },
    calcHeroicStats() {
      this.getStatFieldsHeroic().forEach((field) => {
        let update = {
          field: field.field,
          value: Math.round(this.randomIntFromInterval(
            this.heroicStatsMin,
            this.heroicStatsMax,
          ))
        }

        if (update.value > 0) {
          this.$emit("field", update);
        }
      })
    },
    calcResists() {
      this.getResistFields().forEach((field) => {
        let update = {
          field: field.field,
          value: Math.round(this.randomIntFromInterval(
            this.resistsMin,
            this.resistsMax,
          ))
        }

        if (update.value > 0) {
          this.$emit("field", update);
        }
      })
    },
    calcHeroicResists() {
      this.getResistHeroicFields().forEach((field) => {
        let update = {
          field: field.field,
          value: Math.round(this.randomIntFromInterval(
            this.heroicResistsMin,
            this.heroicResistsMax,
          ))
        }

        if (update.value > 0) {
          this.$emit("field", update);
        }
      })
    },

    //
    // highlight functions
    //
    highlightField(name) {
      const target = document.getElementById(name)
      if (target) {
        target.parentElement.parentElement.classList.add("pulsate-highlight");
      }
    },
    unhighlightField(name) {
      const target = document.getElementById(name)
      if (target) {
        target.parentElement.parentElement.classList.remove("pulsate-highlight");
      }
    },

    // for more precision around heroic / regular
    highlightFieldImmediate(name) {
      const target = document.getElementById(name)
      if (target) {
        target.parentElement.classList.add("pulsate-highlight");
      }
    },
    unhighlightFieldImmediate(name) {
      const target = document.getElementById(name)
      if (target) {
        target.parentElement.classList.remove("pulsate-highlight");
      }
    },

    highlightHpManaEnd() {
      this.hpManaEndFields.forEach((entry) => {
        this.highlightField(entry)
      })
    },
    highlightModFields() {
      this.getMod3Fields().forEach((entry) => {
        this.highlightField(entry.field)
      })
    },
    highlightStatFields() {
      this.getStatFields().forEach((entry) => {
        this.highlightFieldImmediate(entry.field)
      })
    },
    highlightHeroicStatsFields() {
      this.getStatFieldsHeroic().forEach((entry) => {
        this.highlightFieldImmediate(entry.field)
      })
    },
    highlightResistFields() {
      this.getResistFields().forEach((entry) => {
        this.highlightFieldImmediate(entry.field)
      })
    },
    highlightHeroicResistsFields() {
      this.getResistHeroicFields().forEach((entry) => {
        this.highlightFieldImmediate(entry.field)
      })
    },

    clearAllHighlights() {
      this.getAllFields().forEach((entry) => {
        this.unhighlightField(entry.field)
        this.unhighlightFieldImmediate(entry.field)
      })
    }
  }
}
</script>

<style scoped>

</style>
