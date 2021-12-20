<template>
  <div>
    Increase Stats by % from original stats
    <b-form-input
      v-model.number="increaseStatBy"
      @change="syncStats"
      type="number"
    />
  </div>
</template>

<script>
export default {
  name: "ItemStatScaleTool",
  data() {
    return {
      increaseStatBy: 100,

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

      stats: {
        "Strength": { stat: "astr", heroic: "heroic_str" },
        "Stamina": { stat: "asta", heroic: "heroic_sta" },
        "Intelligence": { stat: "aint", heroic: "heroic_int" },
        "Wisdom": { stat: "awis", heroic: "heroic_wis" },
        "Agility": { stat: "aagi", heroic: "heroic_agi" },
        "Dexterity": { stat: "adex", heroic: "heroic_dex" },
        "Charisma": { stat: "acha", heroic: "heroic_cha" },
        "Magic Resist": { stat: "mr", heroic: "heroic_mr" },
        "Fire Resists": { stat: "fr", heroic: "heroic_fr" },
        "Cold Resist": { stat: "cr", heroic: "heroic_cr" },
        "Disease Resist": { stat: "dr", heroic: "heroic_dr" },
        "Poison Resist": { stat: "pr", heroic: "heroic_pr" },
        "Corruption": { stat: "svcorruption", heroic: "heroic_svcorrup" }
      },
      mod3: {
        "Attack": "attack",
        "HP Regen": "regen",
        "Mana Regen": "manaregen",
        "Endurance Regen": "enduranceregen",
        "Accuracy": "accuracy",
        "Avoidance": "avoidance",
        "Clairvoyance": "clairvoyance",
        "Combat Effects": "combateffects",
        "Damage Shield": "damageshield",
        "Damage Shield Mitigation": "dsmitigation",
        "DoT Shielding": "dotshielding",
        "Heal Amount": "healamt",
        "Shielding": "shielding",
        "Spell Shielding": "spellshield",
        "Strikethrough": "strikethrough",
        "Stun Resist": "stunresist",
      },
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
          description: 'Extra Damage Skill',
          field: 'extradmgskill'
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
          description: 'Bane Damage Body',
          field: 'banedmgbody'
        },
        {
          description: 'Bane Damage Race',
          field: 'banedmgrace'
        },
        {
          description: 'Bane Damage Race Amount',
          field: 'banedmgraceamt'
        },
        {
          description: 'Elemental Damage Amount',
          field: 'elemdmgamt'
        },
        {
          description: 'Element Damage Type',
          field: 'elemdmgtype'
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
    syncStats() {
      for (let key in this.mod3) {
        const field = this.mod3[key]
        let update  = {
          field: field,
          value: Math.round(this.originalItemData[field] * (this.increaseStatBy / 100))
        }

        if (update.value > 0) {
          this.$emit("field", update);
        }
      }

      for (let key in this.damageStats) {
        const entry = this.damageStats[key]
        const field = entry.field
        let update  = {
          field: field,
          value: Math.round(this.originalItemData[field] * (this.increaseStatBy / 100))
        }

        if (update.value > 0) {
          this.$emit("field", update);
        }
      }

      for (let key in this.topStats) {
        const entry = this.topStats[key]
        const field = entry.field
        let update  = {
          field: field,
          value: Math.round(this.originalItemData[field] * (this.increaseStatBy / 100))
        }

        if (update.value > 0) {
          this.$emit("field", update);
        }
      }

      for (let key in this.stats) {
        const entry  = this.stats[key]
        const stat   = entry.stat
        const heroic = entry.heroic
        let update   = {
          field: stat,
          value: Math.round(this.originalItemData[stat] * (this.increaseStatBy / 100))
        }

        if (update.value > 0) {
          this.$emit("field", update);
        }

        let heroicUpdate = {
          field: heroic,
          value: Math.round(this.originalItemData[heroic] * (this.increaseStatBy / 100))
        }

        if (heroicUpdate.value > 0) {
          this.$emit("field", heroicUpdate);
        }
      }

    }
  }
}
</script>

<style scoped>

</style>
