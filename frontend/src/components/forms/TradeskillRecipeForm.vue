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
              "help": "\u003cp\u003eUnique Tradeskill Recipe Identifier\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Name",
              "model": "name",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eRecipe Name\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Tradeskill",
              "model": "tradeskill",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"https://eqemu.gitbook.io/server/categories/player/skills\" target=\"_blank\"\u003eTradeskill\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Skillneeded",
              "model": "skillneeded",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eSkill Level Needed\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Trivial",
              "model": "trivial",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eTrivial Skill Level\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Nofail",
              "model": "nofail",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eNo Fail: 0 = False, 1 = True\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Replace Container",
              "model": "replace_container",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eReplace Container: 0 = False, 1 = True\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Notes",
              "model": "notes",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eNotes\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Must Learn",
              "model": "must_learn",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eMust Learn: 0 = False, 1 = True\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Quest",
              "model": "quest",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eQuest: 0 = False, 1 = True\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Enabled",
              "model": "enabled",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eEnabled: 0 = False, 1 = True\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Min Expansion",
              "model": "min_expansion",
              "styleClasses": "col-12",
              "type": "input"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Max Expansion",
              "model": "max_expansion",
              "styleClasses": "col-12",
              "type": "input"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Content Flags",
              "model": "content_flags",
              "styleClasses": "col-12",
              "type": "input"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Content Flags Disabled",
              "model": "content_flags_disabled",
              "styleClasses": "col-12",
              "type": "input"
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
