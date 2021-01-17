<template>
  <vue-form-generator
    v-if="model"
    :schema="schema"
    :model="model"
    :options="formOptions"/>
</template>

<script>
export default {
  props: {
    model: {
      type: Object
    }
  },
  data() {
    return {
      schema: {
        fields:
          [
            {
              "featured": false,
              "inputType": "text",
              "label": "Type",
              "model": "type",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eType: 0 = Trash, 1 = Named, 2 = Raid\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Level",
              "model": "level",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eLevel\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Ac",
              "model": "ac",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eArmor Class\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Hp",
              "model": "hp",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eHealth\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Accuracy",
              "model": "accuracy",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eAccuracy\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Slow Mitigation",
              "model": "slow_mitigation",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eSlow Mitigation\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Attack",
              "model": "attack",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eAttack\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Strength",
              "model": "strength",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eStrength\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Stamina",
              "model": "stamina",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eStamina\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Dexterity",
              "model": "dexterity",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eDexterity\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Agility",
              "model": "agility",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eAgility\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Intelligence",
              "model": "intelligence",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eIntelligence\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Wisdom",
              "model": "wisdom",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eWisdom\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Charisma",
              "model": "charisma",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eCharisma\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Magic Resist",
              "model": "magic_resist",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eMagic Resistance\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Cold Resist",
              "model": "cold_resist",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eCold Resistance\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Fire Resist",
              "model": "fire_resist",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eFire Resistance\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Poison Resist",
              "model": "poison_resist",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003ePoison Resistance\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Disease Resist",
              "model": "disease_resist",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eDisease Resistance\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Corruption Resist",
              "model": "corruption_resist",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eCorruption Resistance\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Physical Resist",
              "model": "physical_resist",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003ePhysical Resistance\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Min Dmg",
              "model": "min_dmg",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eMinimum Damage\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Max Dmg",
              "model": "max_dmg",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eMaximum Damage\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Hp Regen Rate",
              "model": "hp_regen_rate",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eHealth Regeneration Rate\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Attack Delay",
              "model": "attack_delay",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eAttack Delay: Delay between the attack arounds in 10ths of a second.\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Spell Scale",
              "model": "spell_scale",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eSpell Scale: 50 = 50%, 100 = 100%, 150 = 150%\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Heal Scale",
              "model": "heal_scale",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eHeal Scale: 50 = 50%, 100 = 100%, 150 = 150%\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Special Abilities",
              "model": "special_abilities",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eSpecial Abilities\u003c/p\u003e\n"
            }
          ]
      },

      formOptions: {
        // validateAfterLoad: true,
        // validateAfterChanged: true
      }
    };
  }

}
</script>
