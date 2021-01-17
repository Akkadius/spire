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
              "help": "\u003cp\u003eUnique Trap Identifier\u003c/p\u003e\n"
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
              "help": "\u003cp\u003eVersion: -1 For All\u003c/p\u003e\n"
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
              "label": "Chance",
              "model": "chance",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eChance: 0 = None, 100 = Always\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Maxzdiff",
              "model": "maxzdiff",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eMax Z Difference\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Radius",
              "model": "radius",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eTrap Radius\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Effect",
              "model": "effect",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"https://eqemu.gitbook.io/server/categories/zones/trap-types\" target=\"_blank\"\u003eTrap Type\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Effectvalue",
              "model": "effectvalue",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eEffect Value: (Based on Trap Type) 0 = \u003ca href=\"spells_new.md\" target=\"_blank\"\u003eSpell Identifier\u003c/a\u003e, 1 = Radius, 2 = \u003ca href=\"npc_types.md\" target=\"_blank\"\u003eNPC Type Identifier\u003c/a\u003e, 3 = \u003ca href=\"npc_types.md\" target=\"_blank\"\u003eNPC Type Identifier\u003c/a\u003e, 4 = Minimum Damage\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Effectvalue2",
              "model": "effectvalue2",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eEffect Value 2: (Based on Trap Type) 0 = Unused, 1 = (0 = Everything Will Aggro, 1 = Only KoS Will Agro), 2 = Number of NPCs, 3 = Number of NPCs, 4 = Maximum Damage\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Message",
              "model": "message",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eMessage\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Skill",
              "model": "skill",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"https://eqemu.gitbook.io/server/categories/player/skills\" target=\"_blank\"\u003eSkill Required\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Level",
              "model": "level",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eLevel\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Respawn Time",
              "model": "respawn_time",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eRespawn Timer in Seconds\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Respawn Var",
              "model": "respawn_var",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eRandom Respawn Timer Variance in Seconds\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Triggered Number",
              "model": "triggered_number",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eTriggered Member\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Group",
              "model": "group",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eGroup\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Despawn When Triggered",
              "model": "despawn_when_triggered",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eDespawn When Triggered: 0 = False, 1 = True\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Undetectable",
              "model": "undetectable",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eUndetectable: 0 = False, 1= True\u003c/p\u003e\n"
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
