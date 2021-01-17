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
              "help": "\u003cp\u003eUnique Spell Effect Entry Identifier\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Npc Spells Effects Id",
              "model": "npc_spells_effects_id",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"npc_spells_effects.md\" target=\"_blank\"\u003eNPC Spells Effects Identifier\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Spell Effect Id",
              "model": "spell_effect_id",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"https://eqemu.gitbook.io/server/categories/spells/spell-effect-ids\" target=\"_blank\"\u003eSpell Effect Identifier\u003c/a\u003e\u003c/p\u003e\n"
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
              "label": "Se Base",
              "model": "se_base",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eSpell Effect Base\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Se Limit",
              "model": "se_limit",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eSpell Effect Limit\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Se Max",
              "model": "se_max",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eSpell Effect Maximum\u003c/p\u003e\n"
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
