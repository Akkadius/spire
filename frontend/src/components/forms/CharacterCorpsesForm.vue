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
              "help": "\u003cp\u003eUnique Corpse Identifier\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Charid",
              "model": "charid",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"character_data.md\" target=\"_blank\"\u003eCharacter Identifier\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Charname",
              "model": "charname",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eCharacter Name\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Zone Id",
              "model": "zone_id",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"https://eqemu.gitbook.io/server/categories/zones/zone-list\" target=\"_blank\"\u003eZone Identifier\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Instance Id",
              "model": "instance_id",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eInstance Identifier\u003c/p\u003e\n"
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
              "label": "Time Of Death",
              "model": "time_of_death",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eTime of Death\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Guild Consent Id",
              "model": "guild_consent_id",
              "styleClasses": "col-12",
              "type": "input"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Is Rezzed",
              "model": "is_rezzed",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eIs Rezzed: 0 = False, 1 = True\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Is Buried",
              "model": "is_buried",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eIs Buried: 0 = False, 1 = True\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Was At Graveyard",
              "model": "was_at_graveyard",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eWas At Graveyard: 0 = False, 1 = True\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Is Locked",
              "model": "is_locked",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eIs Locked: 0 = False, 1 = True\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Exp",
              "model": "exp",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eExperience\u003c/p\u003e\n"
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
              "label": "Level",
              "model": "level",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eLevel\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Race",
              "model": "race",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"https://eqemu.gitbook.io/server/categories/npc/race-list\" target=\"_blank\"\u003eRace\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Gender",
              "model": "gender",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"https://eqemu.gitbook.io/server/categories/npc/genders\" target=\"_blank\"\u003eGender\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Class",
              "model": "class",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"https://eqemu.gitbook.io/server/categories/player/class-list\" target=\"_blank\"\u003eClass\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Deity",
              "model": "deity",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"https://eqemu.gitbook.io/server/categories/player/deity-list\" target=\"_blank\"\u003eDeity\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Texture",
              "model": "texture",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"https://eqemu.gitbook.io/server/categories/npc/textures\" target=\"_blank\"\u003eTexture\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Helm Texture",
              "model": "helm_texture",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"https://eqemu.gitbook.io/server/categories/npc/textures\" target=\"_blank\"\u003eHelm Texture\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Copper",
              "model": "copper",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eCopper\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Silver",
              "model": "silver",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eSilver\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Gold",
              "model": "gold",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eGold\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Platinum",
              "model": "platinum",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003ePlatinum\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Hair Color",
              "model": "hair_color",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eHair Color\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Beard Color",
              "model": "beard_color",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eBeard Color\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Eye Color 1",
              "model": "eye_color_1",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eEye Color 1\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Eye Color 2",
              "model": "eye_color_2",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eEye Color 2\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Hair Style",
              "model": "hair_style",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eHair Style\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Face",
              "model": "face",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eFace\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Beard",
              "model": "beard",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eBeard\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Drakkin Heritage",
              "model": "drakkin_heritage",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eDrakkin Heritage\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Drakkin Tattoo",
              "model": "drakkin_tattoo",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eDrakkin Tattoo\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Drakkin Details",
              "model": "drakkin_details",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eDrakkin Details\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Wc 1",
              "model": "wc_1",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eWear Change 1\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Wc 2",
              "model": "wc_2",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eWear Change 2\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Wc 3",
              "model": "wc_3",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eWear Change 3\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Wc 4",
              "model": "wc_4",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eWear Change 4\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Wc 5",
              "model": "wc_5",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eWear Change 5\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Wc 6",
              "model": "wc_6",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eWear Change 6\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Wc 7",
              "model": "wc_7",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eWear Change 7\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Wc 8",
              "model": "wc_8",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eWear Change 8\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Wc 9",
              "model": "wc_9",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eWear Change 9\u003c/p\u003e\n"
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
