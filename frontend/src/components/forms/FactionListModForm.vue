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
              "help": "\u003cp\u003eUnique Faction List Modifier Identifier\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Faction Id",
              "model": "faction_id",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"faction_list.md\" target=\"_blank\"\u003eFaction Identifier\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Mod",
              "model": "mod",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eModifier\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Mod Name",
              "model": "mod_name",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eModifier Name: r# = \u003ca href=\"https://eqemu.gitbook.io/server/categories/npc/race-list\" target=\"_blank\"\u003eRace Identifier\u003c/a\u003e, c# = \u003ca href=\"https://eqemu.gitbook.io/server/categories/player/class-list\" target=\"_blank\"\u003eClass Identifier\u003c/a\u003e, d# = \u003ca href=\"https://eqemu.gitbook.io/server/categories/player/deity-list\" target=\"_blank\"\u003eDeity Identifier\u003c/a\u003e\u003c/p\u003e\n"
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
