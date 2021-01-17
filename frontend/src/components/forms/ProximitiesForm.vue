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
              "label": "Zoneid",
              "model": "zoneid",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"https://eqemu.gitbook.io/server/categories/zones/zone-list\" target=\"_blank\"\u003eZone Identifier\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Exploreid",
              "model": "exploreid",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"goallists.md\" target=\"_blank\"\u003eExplore Identifier\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Minx",
              "model": "minx",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eMinimum X Coordinate\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Maxx",
              "model": "maxx",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eMaximum X Coordinate\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Miny",
              "model": "miny",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eMinimum Y Coordinate\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Maxy",
              "model": "maxy",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eMaximum Y Coordinate\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Minz",
              "model": "minz",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eMinimum Z Coordinate\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Maxz",
              "model": "maxz",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eMaximum Z Coordinate\u003c/p\u003e\n"
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
