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
              "label": "Npc Faction Id",
              "model": "npc_faction_id",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"npc_faction.md\" target=\"_blank\"\u003eNPC Faction Identifier\u003c/a\u003e\u003c/p\u003e\n"
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
              "label": "Value",
              "model": "value",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eValue\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Npc Value",
              "model": "npc_value",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eNPC Value: -1 = Attack, 0 = Neutral, 1 = Assist\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Temp",
              "model": "temp",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eTemeporary: 0 = Faction is permanent, player recieves a message, 1 = Faction is temporary, player does not recieve a message, 2 = Faction is temporary, player recieves a message, 3 = Faction is permanent, but player does not recieve a message.\u003c/p\u003e\n"
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
