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
              "help": "\u003cp\u003eUnique Title Identifier\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Skill Id",
              "model": "skill_id",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"https://eqemu.gitbook.io/server/categories/player/skills\" target=\"_blank\"\u003eSkill Identifier\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Min Skill Value",
              "model": "min_skill_value",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eMinimum Skill Value\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Max Skill Value",
              "model": "max_skill_value",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eMaximum Skill Value\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Min Aa Points",
              "model": "min_aa_points",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eMinimum AA Points\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Max Aa Points",
              "model": "max_aa_points",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eMaximum AA Points\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Class",
              "model": "class",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"https://eqemu.gitbook.io/server/categories/player/class-list\" target=\"_blank\"\u003eClass\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Gender",
              "model": "gender",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"https://eqemu.gitbook.io/server/categories/npc/genders\" target=\"_blank\"\u003eGender\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Char Id",
              "model": "char_id",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"character_data.md\" target=\"_blank\"\u003eUnique Character Identifier\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Status",
              "model": "status",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"https://eqemu.gitbook.io/server/categories/player/status-levels\" target=\"_blank\"\u003eRequired Status\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Item Id",
              "model": "item_id",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"items.md\" target=\"_blank\"\u003eItem Identifier\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Prefix",
              "model": "prefix",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003ePrefix\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Suffix",
              "model": "suffix",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eSuffix\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Title Set",
              "model": "title_set",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eTitle Set Identifier\u003c/p\u003e\n"
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
