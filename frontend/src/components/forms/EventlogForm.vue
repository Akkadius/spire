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
              "help": "\u003cp\u003eUnique Event Identifier\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Accountname",
              "model": "accountname",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"account.md\" target=\"_blank\"\u003eAccount Name\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Accountid",
              "model": "accountid",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"account.md\" target=\"_blank\"\u003eAccount Identifier\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Status",
              "model": "status",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"https://eqemu.gitbook.io/server/categories/player/status-levels\" target=\"_blank\"\u003eStatus\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Charname",
              "model": "charname",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"character_data.md\" target=\"_blank\"\u003eCharacter Name\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Target",
              "model": "target",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eTarget\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Time",
              "model": "time",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eTIme Timestamp\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Descriptiontype",
              "model": "descriptiontype",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eDescription Type\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Description",
              "model": "description",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eDescription\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Event Nid",
              "model": "event_nid",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eEvent Identifier\u003c/p\u003e\n"
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
