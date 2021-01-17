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
              "label": "Id",
              "model": "id",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eUnique NPC Spell Entry Identifier\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Npc Spells Id",
              "model": "npc_spells_id",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"npc_spells.md\" target=\"_blank\"\u003eUnique NPC Spell Set Identifier\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Spellid",
              "model": "spellid",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"spells_new.md\" target=\"_blank\"\u003eSpell Identifier\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Type",
              "model": "type",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"https://eqemu.gitbook.io/server/categories/spells/spell-types\" target=\"_blank\"\u003eSpell Type Bitmask\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Minlevel",
              "model": "minlevel",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eMinimum Level\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Maxlevel",
              "model": "maxlevel",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eMaximum Level\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Manacost",
              "model": "manacost",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eMana Cost\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Recast Delay",
              "model": "recast_delay",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eRecast Delay\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Priority",
              "model": "priority",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003ePriority: 0 = Innate, 1 = Highest Priority, 5 = Lower Priority, 10 = Even Lower Priority\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Resist Adjust",
              "model": "resist_adjust",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eResist Adjustment\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Min Hp",
              "model": "min_hp",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eMinimum Health Percentage\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Max Hp",
              "model": "max_hp",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eMaximum Health Percentage\u003c/p\u003e\n"
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
