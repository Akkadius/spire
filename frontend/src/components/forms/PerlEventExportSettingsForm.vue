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
              "inputType": "input",
              "label": "Event Id",
              "model": "event_id",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"https://eqemu.gitbook.io/quest-api/events\" target=\"_blank\"\u003eUnique Perl Event Identifier\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "input",
              "label": "Event Description",
              "model": "event_description",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eEvent Description\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "input",
              "label": "Export Qglobals",
              "model": "export_qglobals",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eExport QGlobals: 0 = False, 1 = True\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "input",
              "label": "Export Mob",
              "model": "export_mob",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eExport Mob: 0 = False, 1 = True\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "input",
              "label": "Export Zone",
              "model": "export_zone",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eExport Zone: 0 = False, 1 = True\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "input",
              "label": "Export Item",
              "model": "export_item",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eExport Item: 0 = False, 1 = True\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "input",
              "label": "Export Event",
              "model": "export_event",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eExport Event: 0 = False, 1 = True\u003c/p\u003e\n"
            }
          ]
      },

      formOptions: {
        validateAfterLoad: true,
        validateAfterChanged: true
      }
    };
  }

}
</script>
