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
              "label": "Client Faction Id",
              "model": "client_faction_id",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"faction_list.md\" target=\"_blank\"\u003eClient Faction Identifier\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Min",
              "model": "min",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eMinimum Faction\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Max",
              "model": "max",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eMaximum Faction\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Unk Hero1",
              "model": "unk_hero1",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eUnknown\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Unk Hero2",
              "model": "unk_hero2",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eUnknown\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Unk Hero3",
              "model": "unk_hero3",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eUnknown\u003c/p\u003e\n"
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
