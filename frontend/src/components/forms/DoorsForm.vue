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
              "help": "\u003cp\u003eUnique Entry Identifier\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Doorid",
              "model": "doorid",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eUnique Door Identifier\u003c/p\u003e\n"
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
              "help": "\u003cp\u003eZone Version: -1 For All\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Name",
              "model": "name",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eThis is the name of the door, such as \u0026lsquo;IT11161\u0026rsquo; or \u0026lsquo;POPCREATE501\u0026rsquo;, for names of objects you can see.\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Pos Y",
              "model": "pos_y",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eDoor Y Coordinate\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Pos X",
              "model": "pos_x",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eDoor X Coordinate\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Pos Z",
              "model": "pos_z",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eDoor Z Coordinate\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Heading",
              "model": "heading",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eDoor Heading Coordinate\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Opentype",
              "model": "opentype",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"https://eqemu.gitbook.io/server/categories/zones/door-open-types\" target=\"_blank\"\u003eDoor Open Type\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Guild",
              "model": "guild",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"guilds.md\" target=\"_blank\"\u003eGuild Identifier\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Lockpick",
              "model": "lockpick",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eLockpicking Skill Required: -1 = Unpickable\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Keyitem",
              "model": "keyitem",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"items.md\" target=\"_blank\"\u003eItem Identifier\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Nokeyring",
              "model": "nokeyring",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eNo Key Ring: 0 = False, 1 = True\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Triggerdoor",
              "model": "triggerdoor",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eTrigger Door: 0 For Current Door or use a Unique Door Identifier\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Triggertype",
              "model": "triggertype",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eTrigger Type: 1 = Open a Type 255 door, 255 = Will Not Open\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Disable Timer",
              "model": "disable_timer",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eDisable Timer\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Doorisopen",
              "model": "doorisopen",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eDoor Is Open: 0 = False, 1 = True\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Door Param",
              "model": "door_param",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eDoor Parameter\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Dest Zone",
              "model": "dest_zone",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"https://eqemu.gitbook.io/server/categories/zones/zone-list\" target=\"_blank\"\u003eZone Short Name\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Dest Instance",
              "model": "dest_instance",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"instance_list.md\" target=\"_blank\"\u003eDestination Instance\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Dest X",
              "model": "dest_x",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eDestination X Coordinate\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Dest Y",
              "model": "dest_y",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eDestination Y Coordinate\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Dest Z",
              "model": "dest_z",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eDestination Z Coordinate\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Dest Heading",
              "model": "dest_heading",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eDestination Heading Coordinate\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Invert State",
              "model": "invert_state",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eThis column will basically behave like such: if the door has a click type and it is to raise up like a door, it will be raised on spawn of the door. Meaning it is inverted. Another example: If a \u003ca href=\"https://eqemu.gitbook.io/server/categories/zones/door-open-types\" target=\"_blank\"\u003eDoor Open Type\u003c/a\u003e is set to a spinning object on click, you could set this to 1 to have the door be spinning on spawn.\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Incline",
              "model": "incline",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eIncline\u003c/p\u003e\n"
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
              "label": "Buffer",
              "model": "buffer",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eUnused\u003c/p\u003e\n"
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
              "label": "Is Ldon Door",
              "model": "is_ldon_door",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eIs LDoN Door: 0 = False, 1 = True\u003c/p\u003e\n"
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
