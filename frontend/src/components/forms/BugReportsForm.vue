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
              "help": "\u003cp\u003eUnique Bug Report Identifier\u003c/p\u003e\n"
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
              "label": "Client Version Id",
              "model": "client_version_id",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"https://eqemu.gitbook.io/server/categories/player/client-version-bitmasks\" target=\"_blank\"\u003eClient Version Identifier\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Client Version Name",
              "model": "client_version_name",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"https://eqemu.gitbook.io/server/categories/player/client-version-bitmasks\" target=\"_blank\"\u003eClient Version Name\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Account Id",
              "model": "account_id",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"account.md\" target=\"_blank\"\u003eAccount Identifier\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Character Id",
              "model": "character_id",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"character_data.md\" target=\"_blank\"\u003eCharacter Identifier\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Character Name",
              "model": "character_name",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"character_data.md\" target=\"_blank\"\u003eCharacter Name\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Reporter Spoof",
              "model": "reporter_spoof",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eReporter Spoof\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Category Id",
              "model": "category_id",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eCategory Identifier\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Category Name",
              "model": "category_name",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eCategory Name\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Reporter Name",
              "model": "reporter_name",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eReporter Name\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Ui Path",
              "model": "ui_path",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eUI Path\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Pos X",
              "model": "pos_x",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003ePosition X Coordinate\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Pos Y",
              "model": "pos_y",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003ePosition Y Coordinate\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Pos Z",
              "model": "pos_z",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003ePosition Z Coordinate\u003c/p\u003e\n"
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
              "label": "Time Played",
              "model": "time_played",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eTime Played in Seconds\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Target Id",
              "model": "target_id",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eTarget Identifier\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Target Name",
              "model": "target_name",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eTarget Name\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Optional Info Mask",
              "model": "optional_info_mask",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eOptional Info Mask: 0 = False, 1 = True\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": " Can Duplicate",
              "model": "_can_duplicate",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eCan Duplicate: 0 = False, 1 = True\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": " Crash Bug",
              "model": "_crash_bug",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eCrash Bug\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": " Target Info",
              "model": "_target_info",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eTarget Info\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": " Character Flags",
              "model": "_character_flags",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eCharacter Flags\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": " Unknown Value",
              "model": "_unknown_value",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eUnknown\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Bug Report",
              "model": "bug_report",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eBug Report\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "System Info",
              "model": "system_info",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eSystem Information\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Report Datetime",
              "model": "report_datetime",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eReport Datetime\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Bug Status",
              "model": "bug_status",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eBug Status\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Last Review",
              "model": "last_review",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eLast Review Datetime\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Last Reviewer",
              "model": "last_reviewer",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eLast Reviewer\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Reviewer Notes",
              "model": "reviewer_notes",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eReviewer Notes\u003c/p\u003e\n"
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
