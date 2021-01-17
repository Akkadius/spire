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
              "help": "\u003cp\u003eUnique Zone Point Identifier\u003c/p\u003e\n"
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
              "label": "Version",
              "model": "version",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eVersion\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Number",
              "model": "number",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eRepresents the iterator field sent in the struct ZonePoint_Entry, zone points for the current zone are sent when client zones in (during Client::Handle_Connect_OP_ReqClientSpawn in client_packet.cpp). This number field must be unique and also could have a hardcoded equivalent in the client, eg. client is expecting a specific number value for a zone point or teleport/object pad, such as in Erudin (erudnext).\u003c/p\u003e\n"
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
              "label": "X",
              "model": "x",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eX Coordinate\u003c/p\u003e\n"
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
              "label": "Target Y",
              "model": "target_y",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eTarget Y Coordinate\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Target X",
              "model": "target_x",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eTarget X Coordinate\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Target Z",
              "model": "target_z",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eTarget Z Coordinate\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Target Heading",
              "model": "target_heading",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eTarget Heading Coordinate\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Zoneinst",
              "model": "zoneinst",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"instance_list.d\" target=\"_blank\"\u003eInstance Identifier\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Target Zone Id",
              "model": "target_zone_id",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"https://eqemu.gitbook.io/server/categories/zones/zone-list\" target=\"_blank\"\u003eTarget Zone Identifier\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Target Instance",
              "model": "target_instance",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eTarget Instance Identifier\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Buffer",
              "model": "buffer",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eZone Point Buffer\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Client Version Mask",
              "model": "client_version_mask",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"https://eqemu.gitbook.io/server/categories/player/client-version-bitmasks\" target=\"_blank\"\u003eClient Version Mask\u003c/a\u003e\u003c/p\u003e\n"
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
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Is Virtual",
              "model": "is_virtual",
              "styleClasses": "col-12",
              "type": "input"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Height",
              "model": "height",
              "styleClasses": "col-12",
              "type": "input"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Width",
              "model": "width",
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
