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
              "label": "Gridid",
              "model": "gridid",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"grid.md\" target=\"_blank\"\u003eGrid Identifier\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Zoneid",
              "model": "zoneid",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"https://eqemu.gitbook.io/server/categories/zones/zone-list\" target=\"_blank\"\u003eZone Identifier\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Number",
              "model": "number",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eWaypoint Identifier\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "X",
              "model": "x",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eX Coordinate\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Y",
              "model": "y",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eY Coordinate\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Z",
              "model": "z",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eZ Coordinate\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Heading",
              "model": "heading",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eHeading Coordinate\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Pause",
              "model": "pause",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003ePause in Seconds\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Centerpoint",
              "model": "centerpoint",
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
