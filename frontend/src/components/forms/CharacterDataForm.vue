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
              "help": "\u003cp\u003eUnique Character Identifier\u003c/p\u003e\n"
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
              "label": "Name",
              "model": "name",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eName\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Last Name",
              "model": "last_name",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eLast Name\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Title",
              "model": "title",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eTitle\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Suffix",
              "model": "suffix",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eSuffix\u003c/p\u003e\n"
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
              "label": "Zone Instance",
              "model": "zone_instance",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eZone Instance Identifier\u003c/p\u003e\n"
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
              "label": "Gender",
              "model": "gender",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"https://eqemu.gitbook.io/server/categories/npc/genders\" target=\"_blank\"\u003eGender\u003c/a\u003e\u003c/p\u003e\n"
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
              "label": "Class",
              "model": "class",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"https://eqemu.gitbook.io/server/categories/player/class-list\" target=\"_blank\"\u003eClass\u003c/a\u003e\u003c/p\u003e\n"
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
              "label": "Deity",
              "model": "deity",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"https://eqemu.gitbook.io/server/categories/player/deity-list\" target=\"_blank\"\u003eDeity\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Birthday",
              "model": "birthday",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eUNIX Timestamp of Birthday\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Last Login",
              "model": "last_login",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eUNIX Timestamp of Last Login\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Time Played",
              "model": "time_played",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eTime Played\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Level2",
              "model": "level2",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eLevel 2\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Anon",
              "model": "anon",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eAnon: 0 = False, 1 = Anonymous, 2 = Roleplaying\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Gm",
              "model": "gm",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eGM: 0 = False, 1 = True\u003c/p\u003e\n"
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
              "label": "Hair Color",
              "model": "hair_color",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eHair Color\u003c/p\u003e\n"
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
              "label": "Beard",
              "model": "beard",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eBeard\u003c/p\u003e\n"
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
              "label": "Ability Time Seconds",
              "model": "ability_time_seconds",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eAbility Timer in Seconds\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Ability Number",
              "model": "ability_number",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"aa_ability.md\" target=\"_blank\"\u003eAbility Number\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Ability Time Minutes",
              "model": "ability_time_minutes",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eAbility Timer in Minutes\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Ability Time Hours",
              "model": "ability_time_hours",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eAbility Timer in Hours\u003c/p\u003e\n"
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
              "label": "Aa Points Spent",
              "model": "aa_points_spent",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eAA Points Spent\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Aa Exp",
              "model": "aa_exp",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eAA Experience\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Aa Points",
              "model": "aa_points",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eAA Points\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Group Leadership Exp",
              "model": "group_leadership_exp",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eGroup Leadership Experience\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Raid Leadership Exp",
              "model": "raid_leadership_exp",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eRaid Leadership Experience\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Group Leadership Points",
              "model": "group_leadership_points",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eGroup Leadership Points\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Raid Leadership Points",
              "model": "raid_leadership_points",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eRaid Leadership Points\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Points",
              "model": "points",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003ePoints\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Cur Hp",
              "model": "cur_hp",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eHealth\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Mana",
              "model": "mana",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eMana\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Endurance",
              "model": "endurance",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eEndurance\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Intoxication",
              "model": "intoxication",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eIntoxication\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Str",
              "model": "str",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eStrength\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Sta",
              "model": "sta",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eStamina\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Cha",
              "model": "cha",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eCharisma\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Dex",
              "model": "dex",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eDexterity\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Int",
              "model": "int",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eIntelligence\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Agi",
              "model": "agi",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eAgility\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Wis",
              "model": "wis",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eWisdom\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Zone Change Count",
              "model": "zone_change_count",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eZone Change Count\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Toxicity",
              "model": "toxicity",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eToxicity\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Hunger Level",
              "model": "hunger_level",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eHunger Level\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Thirst Level",
              "model": "thirst_level",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eThirst Level\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Ability Up",
              "model": "ability_up",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eAbility Up\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Ldon Points Guk",
              "model": "ldon_points_guk",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eLDoN Points - \u003ca href=\"https://eqemu.gitbook.io/server/categories/zones/ldon-themes\" target=\"_blank\"\u003eDeepest Guk\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Ldon Points Mir",
              "model": "ldon_points_mir",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eLDoN Points - \u003ca href=\"https://eqemu.gitbook.io/server/categories/zones/ldon-themes\" target=\"_blank\"\u003eMiragul\u0026rsquo;s Menagerie\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Ldon Points Mmc",
              "model": "ldon_points_mmc",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eLDoN Points - \u003ca href=\"https://eqemu.gitbook.io/server/categories/zones/ldon-themes\" target=\"_blank\"\u003eMistmoore Catacombs\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Ldon Points Ruj",
              "model": "ldon_points_ruj",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eLDoN Points - \u003ca href=\"https://eqemu.gitbook.io/server/categories/zones/ldon-themes\" target=\"_blank\"\u003eRujarkian Hills\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Ldon Points Tak",
              "model": "ldon_points_tak",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eLDoN Points - \u003ca href=\"https://eqemu.gitbook.io/server/categories/zones/ldon-themes\" target=\"_blank\"\u003eTakish-Hiz\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Ldon Points Available",
              "model": "ldon_points_available",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eLDoN Points Available\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Tribute Time Remaining",
              "model": "tribute_time_remaining",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eTribute Time Remaining\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Career Tribute Points",
              "model": "career_tribute_points",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eCareer Tribute Points\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Tribute Points",
              "model": "tribute_points",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eTribute Points\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Tribute Active",
              "model": "tribute_active",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eTribute Active: 0 = False, 1 = True\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Pvp Status",
              "model": "pvp_status",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003ePVP Status: 0 = False, 1 = True\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Pvp Kills",
              "model": "pvp_kills",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003ePVP Kills\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Pvp Deaths",
              "model": "pvp_deaths",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003ePVP Deaths\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Pvp Current Points",
              "model": "pvp_current_points",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003ePVP Points\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Pvp Career Points",
              "model": "pvp_career_points",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eCareer PVP Points\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Pvp Best Kill Streak",
              "model": "pvp_best_kill_streak",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eBest Kill Streak\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Pvp Worst Death Streak",
              "model": "pvp_worst_death_streak",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eWorse Death Streak\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Pvp Current Kill Streak",
              "model": "pvp_current_kill_streak",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eCurrent Kill Streak\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Pvp2",
              "model": "pvp2",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003ePVP Status: 0 = False, 1 = True\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Pvp Type",
              "model": "pvp_type",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003ePVP Type\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Show Helm",
              "model": "show_helm",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eShow Helm: 0 = False, 1 = True\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Group Auto Consent",
              "model": "group_auto_consent",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eGroup Auto Consent: 0 = False, 1 = True\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Raid Auto Consent",
              "model": "raid_auto_consent",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eRaid Auto Consent: 0 = False, 1 = True\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Guild Auto Consent",
              "model": "guild_auto_consent",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eGuild Auto Consent: 0 = False, 1 = True\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Leadership Exp On",
              "model": "leadership_exp_on",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eLeadership Experience On : 0 = False, 1 = True\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "RestTimer",
              "model": "RestTimer",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eRest Timer\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Air Remaining",
              "model": "air_remaining",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eAir Remaining\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Autosplit Enabled",
              "model": "autosplit_enabled",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eAutosplit Enabled: 0 = False, 1 = True\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Lfp",
              "model": "lfp",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eLooking For Party: 0 = False, 1 = True\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Lfg",
              "model": "lfg",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eLooking For Guild: 0 = False, 1 = True\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Mailkey",
              "model": "mailkey",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eMail Key\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Xtargets",
              "model": "xtargets",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eXTargets\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Firstlogon",
              "model": "firstlogon",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eFirst Logon\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "E Aa Effects",
              "model": "e_aa_effects",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eExperience AA Effects\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "E Percent To Aa",
              "model": "e_percent_to_aa",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eExperience Percentage to AA Points\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "E Expended Aa Spent",
              "model": "e_expended_aa_spent",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eExpended AA Points Spent\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Aa Points Spent Old",
              "model": "aa_points_spent_old",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eAA Points Spent Old\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Aa Points Old",
              "model": "aa_points_old",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eAA Points Old\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "E Last Invsnapshot",
              "model": "e_last_invsnapshot",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eLast Inventory Snapshot\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Deleted At",
              "model": "deleted_at",
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
