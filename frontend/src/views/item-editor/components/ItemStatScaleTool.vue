<template>
  <div>
    Increase All (Non-Zero) Stats by Multiplier from Original Stats <br>(1 = 100%, 1.3 = 130%, 0.5 = 50%) Hit Enter Key or Move out of Input to see affect
    <b-form-input
      v-model.number="increaseAllStatsBy"
      @change="increaseAllStats"
      @mouseover="highlightAllFields"
      @mouseleave="clearAllHighlights"
      class="mt-3"
      type="number"
      step=".1"
    />
  </div>
</template>

<script>
import {Items} from "@/app/items";

export default {
  name: "ItemStatScaleTool",
  data() {
    return {
      increaseAllStatsBy: 1,

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
      ],

      stats: Items.getBasicStatAndResistFields(),
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
        this.getBasicStatFields()
      )

      return fields;
    },

    // calculator functions
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
        // this.$emit("highlight", entry.field);
        if (entry.value > 0) {
          this.highlightField(entry.field)
        }
      })

      console.log("highlight all")
    },
    clearAllHighlights() {
      this.getAllFields().forEach((entry) => {
        // this.$emit("highlight", entry.field);
        this.unhighlightField(entry.field)
      })

      console.log("clear all")
    }
  }
}
</script>

<style scoped>

</style>
