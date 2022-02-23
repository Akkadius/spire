<template>
  <eq-window-complex
    title="Scale Stats by Percentage"
    class="minified-inputs mb-0"
  >
    <div class="text-center mb-4">
      <span style="color: #7b714a">Increase by Multiplier from Original Stats</span>
      <br>
      <br>
      Use <b>Enter Key</b>, <b>Arrow Up/Down Keys</b>
    </div>

    <div class="row">
      <div class="col-4 text-right m-0 p-0 mt-2">
        All Stats
      </div>
      <div class="col-5">
        <b-form-input
          v-model.number="increaseAllStatsBy"
          @change="increaseAllStats"
          @focus="highlightAllFields"
          @blur="clearAllHighlights"
          @mouseover="highlightAllFields"
          @mouseleave="clearAllHighlights"
          type="number"
          step=".1"
        />
      </div>
      <div class="col-2 p-0 m-0 mt-2">
        {{ Math.round(increaseAllStatsBy * 100) }}%
      </div>
    </div>

    <div class="row mt-1">
      <div class="col-4 text-right m-0 p-0 mt-2">
        Top Stats
      </div>
      <div class="col-5">
        <b-form-input
          v-model.number="increaseAllTopStatsBy"
          @change="increaseAllTopStats"
          @focus="highlightTopStatFields"
          @blur="clearAllHighlights"
          @mouseover="highlightTopStatFields"
          @mouseleave="clearAllHighlights"
          type="number"
          step=".1"
        />
      </div>
      <div class="col-2 p-0 m-0 mt-2">
        {{ Math.round(increaseAllTopStatsBy * 100) }}%
      </div>
    </div>

    <div class="row mt-1">
      <div class="col-4 text-right m-0 p-0 mt-2">
        Damage
      </div>
      <div class="col-5">
        <b-form-input
          v-model.number="increaseAllDamageBy"
          @change="increaseAllDamage"
          @focus="highlightDamageFields"
          @blur="clearAllHighlights"
          @mouseover="highlightDamageFields"
          @mouseleave="clearAllHighlights"
          type="number"
          step=".1"
        />
      </div>
      <div class="col-2 p-0 m-0 mt-2">
        {{ Math.round(increaseAllDamageBy * 100) }}%
      </div>
    </div>

    <div class="row mt-1">
      <div class="col-4 text-right m-0 p-0 mt-2">
        Basic Stats
      </div>
      <div class="col-5">
        <b-form-input
          v-model.number="increaseAllBasicStatsBy"
          @change="increaseAllBasicStats"
          @focus="highlightBasicStatFields"
          @blur="clearAllHighlights"
          @mouseover="highlightBasicStatFields"
          @mouseleave="clearAllHighlights"
          type="number"
          step=".1"
        />
      </div>
      <div class="col-2 p-0 m-0 mt-2">
        {{ Math.round(increaseAllBasicStatsBy * 100) }}%
      </div>
    </div>

    <div class="row mt-1">
      <div class="col-4 text-right m-0 p-0 mt-2">
        Resists
      </div>
      <div class="col-5">
        <b-form-input
          v-model.number="increaseAllResistsBy"
          @change="increaseAllResists"
          @focus="highlightResistFields"
          @blur="clearAllHighlights"
          @mouseover="highlightResistFields"
          @mouseleave="clearAllHighlights"
          type="number"
          step=".1"
        />
      </div>
      <div class="col-2 p-0 m-0 mt-2">
        {{ Math.round(increaseAllResistsBy * 100) }}%
      </div>
    </div>

    <div class="row mt-1">
      <div class="col-4 text-right m-0 p-0 mt-2">
        Mods
      </div>
      <div class="col-5">
        <b-form-input
          v-model.number="increaseAllModsBy"
          @change="increaseAllMods"
          @focus="highlightModFields"
          @blur="clearAllHighlights"
          @mouseover="highlightModFields"
          @mouseleave="clearAllHighlights"
          type="number"
          step=".1"
        />
      </div>
      <div class="col-2 p-0 m-0 mt-2">
        {{ Math.round(increaseAllModsBy * 100) }}%
      </div>
    </div>

  </eq-window-complex>
</template>

<script>
import {Items}         from "@/app/items";
import EqWindowComplex from "@/components/eq-ui/EQWindowComplex";

export default {
  name: "ItemStatScalePercentage",
  components: { EqWindowComplex },
  watch: {
    'increaseAllStatsBy': function (newVal, oldVal) {
      this.increaseAllResistsBy    = 1
      this.increaseAllDamageBy     = 1
      this.increaseAllTopStatsBy   = 1
      this.increaseAllBasicStatsBy = 1
      this.increaseAllModsBy       = 1
    },
    // 'increaseAllResistsBy': function (newVal, oldVal) {
    //   this.increaseAllStatsBy = 1
    // },
  },
  data() {
    return {
      increaseAllStatsBy: 1,
      increaseAllTopStatsBy: 1,
      increaseAllDamageBy: 1,
      increaseAllResistsBy: 1,
      increaseAllBasicStatsBy: 1,
      increaseAllModsBy: 1,

      topStats: [
        {
          description: 'AC',
          field: 'ac'
        },
        {
          description: 'HP',
          field: 'hp',
        },
        {
          description: 'Mana',
          field: 'mana',
        },
        {
          description: 'Endur',
          field: 'endur',
        },
        {
          description: 'Purity',
          field: 'purity',
        },
      ],

      stats: Items.getBasicStatFields(),
      mod3: Items.getMod3Fields(),

      damageStats: [
        {
          description: 'Damage',
          field: 'damage'
        },
        {
          description: 'Haste',
          field: 'haste'
        },
        {
          description: 'Extra Damage Amount',
          field: 'extradmgamt'
        },
        {
          description: 'Backstab Damage',
          field: 'backstabdmg'
        },
        // {
        //   description: 'Range',
        //   field: 'range'
        // },
        {
          description: 'Spell Damage',
          field: 'spelldmg'
        },
        {
          description: 'Bane Damage Amount',
          field: 'banedmgamt'
        },
        {
          description: 'Bane Damage Race Amount',
          field: 'banedmgraceamt'
        },
        {
          description: 'Elemental Damage Amount',
          field: 'elemdmgamt'
        },
      ],

    }
  },
  props: {
    originalItemData: {
      type: Object,
      required: true
    },
  },
  methods: {
    // get field functions
    getMod3Fields() {
      let fields = [];
      for (let key in this.mod3) {
        const field = this.mod3[key]
        fields.push({ field: field, value: this.originalItemData[field] })
      }
      return fields;
    },

    getDamageStatFields() {
      let fields = [];
      for (let key in this.damageStats) {
        const entry = this.damageStats[key]
        const field = entry.field
        fields.push({ field: field, value: this.originalItemData[field] })
      }
      return fields;
    },

    getResistFields() {
      let fields = [];
      for (let key in Items.getResistFields()) {
        const entry = Items.getResistFields()[key]
        const field = entry.stat
        const heroic = entry.heroic
        fields.push({ field: field, value: this.originalItemData[field] })
        fields.push({ field: heroic, value: this.originalItemData[field] })
      }
      return fields;
    },

    getTopStatFields() {
      let fields = [];
      for (let key in this.topStats) {
        const entry = this.topStats[key]
        const field = entry.field
        fields.push({ field: field, value: this.originalItemData[field] })
      }
      return fields;
    },

    getBasicStatFields() {
      let fields = [];
      for (let key in this.stats) {
        const entry  = this.stats[key]
        const stat   = entry.stat
        const heroic = entry.heroic
        fields.push({ field: stat, value: this.originalItemData[stat] })
        fields.push({ field: heroic, value: this.originalItemData[heroic] })
      }

      return fields;
    },

    getAllFields() {
      let fields = [];

      fields = fields.concat(
        this.getMod3Fields(),
        this.getDamageStatFields(),
        this.getTopStatFields(),
        this.getBasicStatFields(),
        this.getResistFields()
      )

      return fields;
    },

    // calculator functions
    // percentage increase functions
    increaseAllStats() {
      this.getAllFields().forEach((entry) => {
        const field = entry.field
        const value = entry.value

        let update = {
          field: field,
          value: Math.round(value * this.increaseAllStatsBy)
        }

        if (update.value > 0) {
          this.$emit("field", update);
        }
      })
    },
    increaseAllTopStats() {
      this.getTopStatFields().forEach((entry) => {
        const field = entry.field
        const value = entry.value

        let update = {
          field: field,
          value: Math.round(value * this.increaseAllTopStatsBy)
        }

        if (update.value > 0) {
          this.$emit("field", update);
        }
      })
    },
    increaseAllDamage() {
      this.getDamageStatFields().forEach((entry) => {
        const field = entry.field
        const value = entry.value

        let update = {
          field: field,
          value: Math.round(value * this.increaseAllDamageBy)
        }

        if (update.value > 0) {
          this.$emit("field", update);
        }
      })
    },
    increaseAllBasicStats() {
      this.getBasicStatFields().forEach((entry) => {
        const field = entry.field
        const value = entry.value

        let update = {
          field: field,
          value: Math.round(value * this.increaseAllBasicStatsBy)
        }

        if (update.value > 0) {
          this.$emit("field", update);
        }
      })
    },
    increaseAllResists() {
      this.getResistFields().forEach((entry) => {
        const field = entry.field
        const value = entry.value

        let update = {
          field: field,
          value: Math.round(value * this.increaseAllResistsBy)
        }

        if (update.value > 0) {
          this.$emit("field", update);
        }
      })
    },
    increaseAllMods() {
      this.getMod3Fields().forEach((entry) => {
        const field = entry.field
        const value = entry.value

        let update = {
          field: field,
          value: Math.round(value * this.increaseAllModsBy)
        }

        if (update.value > 0) {
          this.$emit("field", update);
        }
      })
    },

    // highlight functions
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

    highlightAllFields() {
      this.getAllFields().forEach((entry) => {
        if (entry.value > 0) {
          this.highlightField(entry.field)
        }
      })
    },
    highlightTopStatFields() {
      this.getTopStatFields().forEach((entry) => {
        if (entry.value > 0) {
          this.highlightField(entry.field)
        }
      })
    },
    highlightDamageFields() {
      this.getDamageStatFields().forEach((entry) => {
        if (entry.value > 0) {
          this.highlightField(entry.field)
        }
      })
    },
    highlightBasicStatFields() {
      this.getBasicStatFields().forEach((entry) => {
        if (entry.value > 0) {
          this.highlightField(entry.field)
        }
      })
    },
    highlightResistFields() {
      this.getResistFields().forEach((entry) => {
        if (entry.value > 0) {
          this.highlightField(entry.field)
        }
      })
    },
    highlightModFields() {
      this.getMod3Fields().forEach((entry) => {
        if (entry.value > 0) {
          this.highlightField(entry.field)
        }
      })
    },

    clearAllHighlights() {
      this.getAllFields().forEach((entry) => {
        this.unhighlightField(entry.field)
      })
    }
  }
}
</script>

<style scoped>

</style>
