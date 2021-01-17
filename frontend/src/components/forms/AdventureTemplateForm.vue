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
              "help": "\u003cp\u003e\u003ca href=\"adventure_details.md\" target=\"_blank\"\u003eAdventure Identifier\u003c/a\u003e\u003c/p\u003e\n"
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
              "label": "Zone Version",
              "model": "zone_version",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eZone Version\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Is Hard",
              "model": "is_hard",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eHard: 0 = False, 1 = True\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Is Raid",
              "model": "is_raid",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eRaid: 0 = False, 1 = True\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Min Level",
              "model": "min_level",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eMinimum Level\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Max Level",
              "model": "max_level",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eMaximum Level\u003c/p\u003e\n"
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
              "label": "Type Data",
              "model": "type_data",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eType Data\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Type Count",
              "model": "type_count",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eType Count\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Assa X",
              "model": "assa_x",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eAssassination X Coordinate\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Assa Y",
              "model": "assa_y",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eAssassination Y Coordinate\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Assa Z",
              "model": "assa_z",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eAssassination Z Coordinate\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Assa H",
              "model": "assa_h",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eAssassination Heading Coordinate\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Text",
              "model": "text",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eText\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Duration",
              "model": "duration",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eDuration\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Zone In Time",
              "model": "zone_in_time",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eZone In Duration\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Win Points",
              "model": "win_points",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eLDoN Points for Winning\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Lose Points",
              "model": "lose_points",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eLDoN Points for Losing\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Theme",
              "model": "theme",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"https://eqemu.gitbook.io/server/categories/zones/ldon-themes\" target=\"_blank\"\u003eLDoN Theme\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Zone In Zone Id",
              "model": "zone_in_zone_id",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"https://eqemu.gitbook.io/server/categories/zones/zone-list\" target=\"_blank\"\u003eZone In Zone Identifier\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Zone In X",
              "model": "zone_in_x",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eZone In X Coordinate\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Zone In Y",
              "model": "zone_in_y",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eZone In Y Coordinate\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Zone In Object Id",
              "model": "zone_in_object_id",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eZone In Object Identifier\u003c/p\u003e\n"
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
              "label": "Dest H",
              "model": "dest_h",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eDestination Heading Coordinate\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Graveyard Zone Id",
              "model": "graveyard_zone_id",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"https://eqemu.gitbook.io/server/categories/zones/zone-list\" target=\"_blank\"\u003eZone Identifier\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Graveyard X",
              "model": "graveyard_x",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eGraveyard X Coordinate\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Graveyard Y",
              "model": "graveyard_y",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eGraveyard Y Coordinate\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Graveyard Z",
              "model": "graveyard_z",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eGraveyard Z Coordinate\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Graveyard Radius",
              "model": "graveyard_radius",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eGraveyard Radius\u003c/p\u003e\n"
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
