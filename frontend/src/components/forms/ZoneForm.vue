<template>
  <vue-form-generator
    v-if="model"
    :schema="schema"
    :model="model"
    @validated="onValidated"
    :options="formOptions"/>
</template>

<script>
export default {
  props: {
    model: {
      type: Object
    }
  },
  methods: {
    onValidated(isValid, errors) {
      console.log("Validation result: ", isValid, ", Errors:", errors);
    },
  },
  data() {
    return {
      schema: {
        fields:
          [
            {
              "featured": false,
              "inputType": "text",
              "label": "Short Name",
              "model": "short_name",
              "styleClasses": "col-4",
              "type": "input",
              "help": "\u003cp\u003eShort Name\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Id",
              "model": "id",
              "styleClasses": "col-4",
              "type": "input",
              "help": "\u003cp\u003eUnique Entry Identifier\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "File Name",
              "model": "file_name",
              "styleClasses": "col-4",
              "type": "input",
              "help": "\u003cp\u003eFile Name\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Long Name",
              "model": "long_name",
              "styleClasses": "col-4",
              "type": "input",
              "help": "\u003cp\u003eLong Name\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Map File Name",
              "model": "map_file_name",
              "styleClasses": "col-4",
              "type": "input",
              "help": "\u003cp\u003eMap File Name\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Safe X",
              "model": "safe_x",
              "styleClasses": "col-4",
              "type": "input",
              "help": "\u003cp\u003eSafe X Coordinate\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Safe Y",
              "model": "safe_y",
              "styleClasses": "col-4",
              "type": "input",
              "help": "\u003cp\u003eSafe Y Coordinate\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Safe Z",
              "model": "safe_z",
              "styleClasses": "col-4",
              "type": "input",
              "help": "\u003cp\u003eSafe Z Coordinate\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Graveyard Id",
              "model": "graveyard_id",
              "styleClasses": "col-4",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"graveyard.md\" target=\"_blank\"\u003eGraveyard Identifier\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Min Level",
              "model": "min_level",
              "styleClasses": "col-4",
              "type": "input",
              "help": "\u003cp\u003eMinimum Level\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Min Status",
              "model": "min_status",
              "styleClasses": "col-4",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"https://eqemu.gitbook.io/server/categories/player/status-levels\" target=\"_blank\"\u003eMinimum Status\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Zoneidnumber",
              "model": "zoneidnumber",
              "styleClasses": "col-4",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"https://eqemu.gitbook.io/server/categories/zones/zone-list\" target=\"_blank\"\u003eUnique Zone Identifier\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Version",
              "model": "version",
              "styleClasses": "col-4",
              "type": "input",
              "help": "\u003cp\u003eVersion\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Timezone",
              "model": "timezone",
              "styleClasses": "col-4",
              "type": "input",
              "help": "\u003cp\u003eTimezone\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Maxclients",
              "model": "maxclients",
              "styleClasses": "col-4",
              "type": "input",
              "help": "\u003cp\u003eMaximum Clients\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Ruleset",
              "model": "ruleset",
              "styleClasses": "col-4",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"rule_sets.md\" target=\"_blank\"\u003eRuleset Identifier\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Note",
              "model": "note",
              "styleClasses": "col-4",
              "type": "input",
              "help": "\u003cp\u003eNote\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Underworld",
              "model": "underworld",
              "styleClasses": "col-4",
              "type": "input",
              "help": "\u003cp\u003eBottom Z to represent when the player is under the world\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Minclip",
              "model": "minclip",
              "styleClasses": "col-4",
              "type": "input",
              "help": "\u003cp\u003eMinimum Clipping Distance\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Maxclip",
              "model": "maxclip",
              "styleClasses": "col-4",
              "type": "input",
              "help": "\u003cp\u003eMaximum Clipping Distance\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Fog Minclip",
              "model": "fog_minclip",
              "styleClasses": "col-4",
              "type": "input",
              "help": "\u003cp\u003eFog Minimum Clipping Distance\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Fog Maxclip",
              "model": "fog_maxclip",
              "styleClasses": "col-4",
              "type": "input",
              "help": "\u003cp\u003eFog Maximum Clipping Distance\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Fog Blue",
              "model": "fog_blue",
              "styleClasses": "col-4",
              "type": "input",
              "help": "\u003cp\u003eFog Blue Level: 0 = None, 255 = Max\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Fog Red",
              "model": "fog_red",
              "styleClasses": "col-4",
              "type": "input",
              "help": "\u003cp\u003eFog Red Level: 0 = None, 255 = Max\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Fog Green",
              "model": "fog_green",
              "styleClasses": "col-4",
              "type": "input",
              "help": "\u003cp\u003eFog Green Level: 0 = None, 255 = Max\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Sky",
              "model": "sky",
              "styleClasses": "col-4",
              "type": "input",
              "help": "\u003cp\u003eSky type the client will present as the backdrop\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Ztype",
              "model": "ztype",
              "styleClasses": "col-4",
              "type": "input",
              "help": "\u003cp\u003eThis field is sent directly to the client on zone-in, most zones are set to 0, 1 or 255.\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Zone Exp Multiplier",
              "model": "zone_exp_multiplier",
              "styleClasses": "col-4",
              "type": "input",
              "help": "\u003cp\u003eThis will multiply the XP to this percentage value (decimal based, 100% = 1.0)\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Walkspeed",
              "model": "walkspeed",
              "styleClasses": "col-4",
              "type": "input",
              "help": "\u003cp\u003eWalkspeed in this zone\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Time Type",
              "model": "time_type",
              "styleClasses": "col-4",
              "type": "input",
              "help": "\u003cp\u003eThis value varies depending on the zone but it is sent to the client on zone in. Most starting zones/newbie areas have this value set to 2, dungeons tyically have this set to 0, some zones break from the norm and have values greater than 2, (akanon = 3, blackburrow = 5, cazicthule = 5, crushbone = 5, erudnint = 4, kaladima = 3, etc.)\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Fog Red1",
              "model": "fog_red1",
              "styleClasses": "col-4",
              "type": "input",
              "help": "\u003cp\u003eFog Red Level 1: 0 = None, 255 = Max\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Fog Green1",
              "model": "fog_green1",
              "styleClasses": "col-4",
              "type": "input",
              "help": "\u003cp\u003eFog Green Level 1: 0 = None, 255 = Max\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Fog Blue1",
              "model": "fog_blue1",
              "styleClasses": "col-4",
              "type": "input",
              "help": "\u003cp\u003eFog Blue Level 1: 0 = None, 255 = Max\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Fog Minclip1",
              "model": "fog_minclip1",
              "styleClasses": "col-4",
              "type": "input",
              "help": "\u003cp\u003eFog Minimum Clipping Distance 1\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Fog Maxclip1",
              "model": "fog_maxclip1",
              "styleClasses": "col-4",
              "type": "input",
              "help": "\u003cp\u003eFog Maximum Clipping Distance 1\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Fog Red2",
              "model": "fog_red2",
              "styleClasses": "col-4",
              "type": "input",
              "help": "\u003cp\u003eFog Red Level 2: 0 = None, 255 = Max\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Fog Green2",
              "model": "fog_green2",
              "styleClasses": "col-4",
              "type": "input",
              "help": "\u003cp\u003eFog Green Level 2: 0 = None, 255 = Max\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Fog Blue2",
              "model": "fog_blue2",
              "styleClasses": "col-4",
              "type": "input",
              "help": "\u003cp\u003eFog Blue Level 2: 0 = None, 255 = Max\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Fog Minclip2",
              "model": "fog_minclip2",
              "styleClasses": "col-4",
              "type": "input",
              "help": "\u003cp\u003eFog Minimum Clipping Distance 2\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Fog Maxclip2",
              "model": "fog_maxclip2",
              "styleClasses": "col-4",
              "type": "input",
              "help": "\u003cp\u003eFog Maximum Clipping Distance 2\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Fog Red3",
              "model": "fog_red3",
              "styleClasses": "col-4",
              "type": "input",
              "help": "\u003cp\u003eFog Red Level 3: 0 = None, 255 = Max\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Fog Green3",
              "model": "fog_green3",
              "styleClasses": "col-4",
              "type": "input",
              "help": "\u003cp\u003eFog Green Level 3: 0 = None, 255 = Max\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Fog Blue3",
              "model": "fog_blue3",
              "styleClasses": "col-4",
              "type": "input",
              "help": "\u003cp\u003eFog Blue Level 3: 0 = None, 255 = Max\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Fog Minclip3",
              "model": "fog_minclip3",
              "styleClasses": "col-4",
              "type": "input",
              "help": "\u003cp\u003eFog Minimum Clipping Distance 3\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Fog Maxclip3",
              "model": "fog_maxclip3",
              "styleClasses": "col-4",
              "type": "input",
              "help": "\u003cp\u003eFog Maximum Clipping Distance 4\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Fog Red4",
              "model": "fog_red4",
              "styleClasses": "col-4",
              "type": "input",
              "help": "\u003cp\u003eFog Red Level 4: 0 = None, 255 = Max\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Fog Green4",
              "model": "fog_green4",
              "styleClasses": "col-4",
              "type": "input",
              "help": "\u003cp\u003eFog Green Level 4: 0 = None, 255 = Max\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Fog Blue4",
              "model": "fog_blue4",
              "styleClasses": "col-4",
              "type": "input",
              "help": "\u003cp\u003eFog Blue Level 4: 0 = None, 255 = Max\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Fog Minclip4",
              "model": "fog_minclip4",
              "styleClasses": "col-4",
              "type": "input",
              "help": "\u003cp\u003eFog Minimum Clipping Distance 4\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Fog Maxclip4",
              "model": "fog_maxclip4",
              "styleClasses": "col-4",
              "type": "input",
              "help": "\u003cp\u003eFog Maximum Clipping Distance 4\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Fog Density",
              "model": "fog_density",
              "styleClasses": "col-4",
              "type": "input",
              "help": "\u003cp\u003eThis is the intensity of the fog, this should be a number between 0-1, most commonly used is .1 or .33\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Flag Needed",
              "model": "flag_needed",
              "styleClasses": "col-4",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"zone_flags.md\" target=\"_blank\"\u003eFlag Required\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Canbind",
              "model": "canbind",
              "styleClasses": "col-4",
              "type": "input",
              "help": "\u003cp\u003eCan Bind: 0 = False, 1 = True (for Caster), 2 = True (for All)\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Cancombat",
              "model": "cancombat",
              "styleClasses": "col-4",
              "type": "input",
              "help": "\u003cp\u003eCan Combat: 0 = False, 1 = True\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Canlevitate",
              "model": "canlevitate",
              "styleClasses": "col-4",
              "type": "input",
              "help": "\u003cp\u003eCan Levitate: 0 = False, 1 = True (Does not affect those with #gm on)\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Castoutdoor",
              "model": "castoutdoor",
              "styleClasses": "col-4",
              "type": "input",
              "help": "\u003cp\u003eCast Outdoors: 0 = False, 1 = True\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Hotzone",
              "model": "hotzone",
              "styleClasses": "col-4",
              "type": "input",
              "help": "\u003cp\u003eHotzone: 0 = False, 1 = True\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Insttype",
              "model": "insttype",
              "styleClasses": "col-4",
              "type": "input",
              "help": "\u003cp\u003eInstance Type\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Shutdowndelay",
              "model": "shutdowndelay",
              "styleClasses": "col-4",
              "type": "input",
              "help": "\u003cp\u003eShutdown Delay\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Peqzone",
              "model": "peqzone",
              "styleClasses": "col-4",
              "type": "input",
              "help": "\u003cp\u003e#peqzone: 0 = False, 1 = True\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Expansion",
              "model": "expansion",
              "styleClasses": "col-4",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"https://eqemu.gitbook.io/server/categories/operation/expansion-list\" target=\"_blank\"\u003eExpansion\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Suspendbuffs",
              "model": "suspendbuffs",
              "styleClasses": "col-4",
              "type": "input",
              "help": "\u003cp\u003eSuspend Buffs: 0 = False, 1 = True\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Rain Chance1",
              "model": "rain_chance1",
              "styleClasses": "col-4",
              "type": "input",
              "help": "\u003cp\u003eRain Chance 1\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Rain Chance2",
              "model": "rain_chance2",
              "styleClasses": "col-4",
              "type": "input",
              "help": "\u003cp\u003eRain Chance 2\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Rain Chance3",
              "model": "rain_chance3",
              "styleClasses": "col-4",
              "type": "input",
              "help": "\u003cp\u003eRain Chance 3\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Rain Chance4",
              "model": "rain_chance4",
              "styleClasses": "col-4",
              "type": "input",
              "help": "\u003cp\u003eRain Chance 4\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Rain Duration1",
              "model": "rain_duration1",
              "styleClasses": "col-4",
              "type": "input",
              "help": "\u003cp\u003eRain Duration 1\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Rain Duration2",
              "model": "rain_duration2",
              "styleClasses": "col-4",
              "type": "input",
              "help": "\u003cp\u003eRain Duration 2\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Rain Duration3",
              "model": "rain_duration3",
              "styleClasses": "col-4",
              "type": "input",
              "help": "\u003cp\u003eRain Duration 3\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Rain Duration4",
              "model": "rain_duration4",
              "styleClasses": "col-4",
              "type": "input",
              "help": "\u003cp\u003eRain Duration 4\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Snow Chance1",
              "model": "snow_chance1",
              "styleClasses": "col-4",
              "type": "input",
              "help": "\u003cp\u003eSnow Chance 1\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Snow Chance2",
              "model": "snow_chance2",
              "styleClasses": "col-4",
              "type": "input",
              "help": "\u003cp\u003eSnow Chance 2\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Snow Chance3",
              "model": "snow_chance3",
              "styleClasses": "col-4",
              "type": "input",
              "help": "\u003cp\u003eSnow Chance 3\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Snow Chance4",
              "model": "snow_chance4",
              "styleClasses": "col-4",
              "type": "input",
              "help": "\u003cp\u003eSnow Chance 4\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Snow Duration1",
              "model": "snow_duration1",
              "styleClasses": "col-4",
              "type": "input",
              "help": "\u003cp\u003eSnow Duration 1\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Snow Duration2",
              "model": "snow_duration2",
              "styleClasses": "col-4",
              "type": "input",
              "help": "\u003cp\u003eSnow Duration 2\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Snow Duration3",
              "model": "snow_duration3",
              "styleClasses": "col-4",
              "type": "input",
              "help": "\u003cp\u003eSnow Duration 3\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Snow Duration4",
              "model": "snow_duration4",
              "styleClasses": "col-4",
              "type": "input",
              "help": "\u003cp\u003eSnow Duration 4\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Gravity",
              "model": "gravity",
              "styleClasses": "col-4",
              "type": "input",
              "help": "\u003cp\u003eGravity\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Type",
              "model": "type",
              "styleClasses": "col-4",
              "type": "input",
              "help": "\u003cp\u003eType (0 = Unknown, 1 = Regular, 2 = Instanced, 3 = Hybrid, 4 = Raid, 5 = City)\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Skylock",
              "model": "skylock",
              "styleClasses": "col-4",
              "type": "input",
              "help": "\u003cp\u003eSky Lock\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Fast Regen Hp",
              "model": "fast_regen_hp",
              "styleClasses": "col-4",
              "type": "input",
              "help": "\u003cp\u003eFast Regen Health\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Fast Regen Mana",
              "model": "fast_regen_mana",
              "styleClasses": "col-4",
              "type": "input",
              "help": "\u003cp\u003eFast Regen Mana\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Fast Regen Endurance",
              "model": "fast_regen_endurance",
              "styleClasses": "col-4",
              "type": "input",
              "help": "\u003cp\u003eFast Regen Endurance\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Npc Max Aggro Dist",
              "model": "npc_max_aggro_dist",
              "styleClasses": "col-4",
              "type": "input",
              "help": "\u003cp\u003eNPC Max Aggro Distance\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Max Movement Update Range",
              "model": "max_movement_update_range",
              "styleClasses": "col-4",
              "type": "input",
              "help": "\u003cp\u003eMax Movement Update Range\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Min Expansion",
              "model": "min_expansion",
              "styleClasses": "col-4",
              "type": "input"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Max Expansion",
              "model": "max_expansion",
              "styleClasses": "col-4",
              "type": "input"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Content Flags",
              "model": "content_flags",
              "styleClasses": "col-4",
              "type": "input"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Content Flags Disabled",
              "model": "content_flags_disabled",
              "styleClasses": "col-4",
              "type": "input"
            }
          ]
      },

      formOptions: {
        validateAfterChanged: true,
        validateDebounceTime: 500,

        // validateAfterLoad: true,
        // validateAfterChanged: true
      }
    };
  }

}
</script>
