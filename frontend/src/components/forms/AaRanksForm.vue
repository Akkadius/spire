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
              "help": "\u003cp\u003e\u003ca href=\"aa_ability.md\" target=\"_blank\"\u003eAA Identifier\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Upper Hotkey Sid",
              "model": "upper_hotkey_sid",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eUpper Hotkey SID\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Lower Hotkey Sid",
              "model": "lower_hotkey_sid",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eLower Hotkey SID\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Title Sid",
              "model": "title_sid",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eTitle SID\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Desc Sid",
              "model": "desc_sid",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eDescription SID\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Cost",
              "model": "cost",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eCost in AA Points\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Level Req",
              "model": "level_req",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eLevel Required\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Spell",
              "model": "spell",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"spells_new.md\" target=\"_blank\"\u003eSpell Identifier\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Spell Type",
              "model": "spell_type",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"https://eqemu.gitbook.io/server/categories/spells/spell-types\" target=\"_blank\"\u003eSpell Type\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Recast Time",
              "model": "recast_time",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eRecast Timer\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Expansion",
              "model": "expansion",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"https://eqemu.gitbook.io/server/categories/operation/expansion-list\" target=\"_blank\"\u003eExpansion Identifier\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Prev Id",
              "model": "prev_id",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003ePrevious Rank Identifier\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Next Id",
              "model": "next_id",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eNext Rank Identifier\u003c/p\u003e\n"
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
