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
              "help": "\u003cp\u003eUnique Bug Identifier\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Zone",
              "model": "zone",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"https://eqemu.gitbook.io/server/categories/zones/zone-list\" target=\"_blank\"\u003eZone Short Name\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Name",
              "model": "name",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003ePlayer Name\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Ui",
              "model": "ui",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eUI\u003c/p\u003e\n"
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
              "label": "Type",
              "model": "type",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eType\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Flag",
              "model": "flag",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eFlag\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Target",
              "model": "target",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eTarget when reported\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Bug",
              "model": "bug",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eBug reported\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Date",
              "model": "date",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eDate when reported\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Status",
              "model": "status",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e[Account Status of Reporter](\u003ca href=\"https://eqemu.gitbook.io/server/categories/player/status-levels\" target=\"_blank\"\u003ehttps://eqemu.gitbook.io/server/categories/player/status-levels\u003c/a\u003e\u003c/p\u003e\n"
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
