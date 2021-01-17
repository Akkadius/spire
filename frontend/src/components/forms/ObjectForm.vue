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
              "help": "\u003cp\u003eUnique Object Identifier\u003c/p\u003e\n"
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
              "label": "Version",
              "model": "version",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eVersion: -1 For All\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Xpos",
              "model": "xpos",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eX Coordinate\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Ypos",
              "model": "ypos",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eY Coordinate\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Zpos",
              "model": "zpos",
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
              "label": "Itemid",
              "model": "itemid",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"items.md\" target=\"_blank\"\u003eItem Identifier\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Charges",
              "model": "charges",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eCharges\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Objectname",
              "model": "objectname",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eObject Name\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Type",
              "model": "type",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"https://eqemu.gitbook.io/server/categories/zones/object-types\" target=\"_blank\"\u003eType\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Icon",
              "model": "icon",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eIcon\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Unknown08",
              "model": "unknown08",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eUnknown\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Unknown10",
              "model": "unknown10",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eUnknown\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Unknown20",
              "model": "unknown20",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eUnknown\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Unknown24",
              "model": "unknown24",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eUnknown\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Unknown60",
              "model": "unknown60",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eUnknown\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Unknown64",
              "model": "unknown64",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eUnknown\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Unknown68",
              "model": "unknown68",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eUnknown\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Unknown72",
              "model": "unknown72",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eUnknown\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Unknown76",
              "model": "unknown76",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eUnknown\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Unknown84",
              "model": "unknown84",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eUnknown\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Size",
              "model": "size",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eSize\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Tilt X",
              "model": "tilt_x",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eTilt X\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Tilt Y",
              "model": "tilt_y",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eTilt Y\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Display Name",
              "model": "display_name",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eDisplay Name\u003c/p\u003e\n"
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
