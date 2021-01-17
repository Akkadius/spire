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
              "help": "\u003cp\u003eUnique NPC Type Identifier\u003c/p\u003e\n"
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
              "label": "Lastname",
              "model": "lastname",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eLast Name\u003c/p\u003e\n"
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
              "label": "Class",
              "model": "class",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"https://eqemu.gitbook.io/server/categories/player/class-list\" target=\"_blank\"\u003eClass\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Bodytype",
              "model": "bodytype",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"https://eqemu.gitbook.io/server/categories/npc/body-types\" target=\"_blank\"\u003eBody Type\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Hp",
              "model": "hp",
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
              "label": "Gender",
              "model": "gender",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"https://eqemu.gitbook.io/server/categories/npc/genders\" target=\"_blank\"\u003eGender\u003c/a\u003e\u003c/p\u003e\n"
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
              "label": "Helmtexture",
              "model": "helmtexture",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"https://eqemu.gitbook.io/server/categories/npc/textures\" target=\"_blank\"\u003eHelmet Texture\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Herosforgemodel",
              "model": "herosforgemodel",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eHero\u0026rsquo;s Forge Model\u003c/p\u003e\n"
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
              "label": "Hp Regen Rate",
              "model": "hp_regen_rate",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eHealth Regeneration\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Mana Regen Rate",
              "model": "mana_regen_rate",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eMana Regeneration\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Loottable Id",
              "model": "loottable_id",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"loottable.md\" target=\"_blank\"\u003eLoottable Identifier\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Merchant Id",
              "model": "merchant_id",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"merchantlist.md\" target=\"_blank\"\u003eMerchant Identifier\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Alt Currency Id",
              "model": "alt_currency_id",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"alternate_currency.md\" target=\"_blank\"\u003eAlternate Currency Identifier\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Npc Spells Id",
              "model": "npc_spells_id",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"npc_spells.md\" target=\"_blank\"\u003eNPC Spell Set Identifier\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Npc Spells Effects Id",
              "model": "npc_spells_effects_id",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"npc_spells_effects.md\" target=\"_blank\"\u003eNPC Spell Effects Identifier\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Npc Faction Id",
              "model": "npc_faction_id",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"faction_list.md\" target=\"_blank\"\u003eNPC Faction Identifier\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Adventure Template Id",
              "model": "adventure_template_id",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"adventure_template.md\" target=\"_blank\"\u003eAdventure Template Identifier\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Trap Template",
              "model": "trap_template",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"traps.md\" target=\"_blank\"\u003eTrap Template Identifier\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Mindmg",
              "model": "mindmg",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eMinimum Damage\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Maxdmg",
              "model": "maxdmg",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eMaximum Damage\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Attack Count",
              "model": "attack_count",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eAttack Count\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Npcspecialattks",
              "model": "npcspecialattks",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eNPC Special Attacks (Deprecated)\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Special Abilities",
              "model": "special_abilities",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eNPC Special Abilities\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Aggroradius",
              "model": "aggroradius",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eAggro Radius\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Assistradius",
              "model": "assistradius",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eAssist Radius\u003c/p\u003e\n"
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
              "label": "Luclin Hairstyle",
              "model": "luclin_hairstyle",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eHair Style\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Luclin Haircolor",
              "model": "luclin_haircolor",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eHair Color\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Luclin Eyecolor",
              "model": "luclin_eyecolor",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eEye Color 1\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Luclin Eyecolor2",
              "model": "luclin_eyecolor2",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eEye Color 2\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Luclin Beardcolor",
              "model": "luclin_beardcolor",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eBeard Color\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Luclin Beard",
              "model": "luclin_beard",
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
              "label": "Armortint Id",
              "model": "armortint_id",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"npc_types_tint.md\" target=\"_blank\"\u003eArmor Tint Identifier\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Armortint Red",
              "model": "armortint_red",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eArmor Tint Red: 0 = None, 255 = Max\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Armortint Green",
              "model": "armortint_green",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eArmor Tint Green: 0 = None, 255 = Max\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Armortint Blue",
              "model": "armortint_blue",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eArmor Tint Blue: 0 = None, 255 = Max\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "D Melee Texture1",
              "model": "d_melee_texture1",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003ePrimary Weapon Texture\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "D Melee Texture2",
              "model": "d_melee_texture2",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eSecondary Weapon Texture\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Ammo Idfile",
              "model": "ammo_idfile",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eAmmo Texture\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Prim Melee Type",
              "model": "prim_melee_type",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"https://eqemu.gitbook.io/server/categories/player/skills\" target=\"_blank\"\u003ePrimary Melee Type\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Sec Melee Type",
              "model": "sec_melee_type",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"https://eqemu.gitbook.io/server/categories/player/skills\" target=\"_blank\"\u003eSecondary Melee Type\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Ranged Type",
              "model": "ranged_type",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"https://eqemu.gitbook.io/server/categories/player/skills\" target=\"_blank\"\u003eRanged Type\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Runspeed",
              "model": "runspeed",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eRun Speed\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "MR",
              "model": "MR",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eMagic Resistance\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "CR",
              "model": "CR",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eCold Resistance\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "DR",
              "model": "DR",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eDisease Resistance\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "FR",
              "model": "FR",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eFire Resistance\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "PR",
              "model": "PR",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003ePoison Resistance\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Corrup",
              "model": "Corrup",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eCorruption Resistance\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "PhR",
              "model": "PhR",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003ePhysical Resistance\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "See Invis",
              "model": "see_invis",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eSee Invisible: 0 = False, 1 = True\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "See Invis Undead",
              "model": "see_invis_undead",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eSee Invisible vs. Undread: 0 = False, 1 = True\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Qglobal",
              "model": "qglobal",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eQuest Globals: 0 = Disabled, 1 = Enabled (Deprecated)\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "AC",
              "model": "AC",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eArmor Class\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Npc Aggro",
              "model": "npc_aggro",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eNPC Aggro: 0 = False, 1 = True\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Spawn Limit",
              "model": "spawn_limit",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eSpawn Limit\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Attack Speed",
              "model": "attack_speed",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eAttack Speed: The lower the number, the faster the NPC hits. (Deprecated)\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Attack Delay",
              "model": "attack_delay",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eAttack Delay: Delay between the attack arounds in 10ths of a second.\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Findable",
              "model": "findable",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eFindable: 0 = False, 1 = True\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "STR",
              "model": "STR",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eStrength\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "STA",
              "model": "STA",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eStamina\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "DEX",
              "model": "DEX",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eDexterity\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "AGI",
              "model": "AGI",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eAgility\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": " INT",
              "model": "_INT",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eIntelligence\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "WIS",
              "model": "WIS",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eWisdom\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "CHA",
              "model": "CHA",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eCharisma\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "See Hide",
              "model": "see_hide",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eSee Hide: 0 = False, 1 = True\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "See Improved Hide",
              "model": "see_improved_hide",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eSee Improved Hide: 0 = False, 1 = True\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Trackable",
              "model": "trackable",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eTrackable: 0 = False, 1 = True\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Isbot",
              "model": "isbot",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eIs Bot: 0 = False, 1 = True\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Exclude",
              "model": "exclude",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eExclude: 0 = False, 1 = True\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "ATK",
              "model": "ATK",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eAttack\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Accuracy",
              "model": "Accuracy",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eAccuracy\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Avoidance",
              "model": "Avoidance",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eAvoidance\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Slow Mitigation",
              "model": "slow_mitigation",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eSlow Mitigation\u003c/p\u003e\n"
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
              "label": "Maxlevel",
              "model": "maxlevel",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eMaximum Level\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Scalerate",
              "model": "scalerate",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eScale Rate\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Private Corpse",
              "model": "private_corpse",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003ePrivate Corpse: 0 = False, 1 = True\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Unique Spawn By Name",
              "model": "unique_spawn_by_name",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eUnique Spawn By Name: 0 = False, 1 = True\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Underwater",
              "model": "underwater",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eUnderwater: 0 = False, 1 = True\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Isquest",
              "model": "isquest",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eIs Quest: 0 = False, 1 = True\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Emoteid",
              "model": "emoteid",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"npc_emotes.md\" target=\"_blank\"\u003eEmote Identifier\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Spellscale",
              "model": "spellscale",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eSpell Scale: 25 = 25%, 50 = 50%, 100 = 100%\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Healscale",
              "model": "healscale",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eHeal Scale: 25 = 25%, 50 = 50%, 100 = 100%\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "No Target Hotkey",
              "model": "no_target_hotkey",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eNo Target Hotkey: 0 = False, 1 = True\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Raid Target",
              "model": "raid_target",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eRaid Target: 0 = False, 1 = True\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Armtexture",
              "model": "armtexture",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"https://eqemu.gitbook.io/server/categories/npc/textures\" target=\"_blank\"\u003eArm Texture\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Bracertexture",
              "model": "bracertexture",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"https://eqemu.gitbook.io/server/categories/npc/textures\" target=\"_blank\"\u003eBracer Texture\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Handtexture",
              "model": "handtexture",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"https://eqemu.gitbook.io/server/categories/npc/textures\" target=\"_blank\"\u003eHand Texture\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Legtexture",
              "model": "legtexture",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"https://eqemu.gitbook.io/server/categories/npc/textures\" target=\"_blank\"\u003eLeg Texture\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Feettexture",
              "model": "feettexture",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"https://eqemu.gitbook.io/server/categories/npc/textures\" target=\"_blank\"\u003eFeet Texture\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Light",
              "model": "light",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eLight\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Walkspeed",
              "model": "walkspeed",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eWalk Speed\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Peqid",
              "model": "peqid",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003ePEQ Identifier\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Unique ",
              "model": "unique_",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eUnique\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Fixed",
              "model": "fixed",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eFixed\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Ignore Despawn",
              "model": "ignore_despawn",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eIgnore Despawn: 0 = False, 1 = True\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Show Name",
              "model": "show_name",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eShow Name: 0 = False, 1 = True\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Untargetable",
              "model": "untargetable",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eUntargetable: 0 = False, 1 = True\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Charm Ac",
              "model": "charm_ac",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eCharmed Armor Class\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Charm Min Dmg",
              "model": "charm_min_dmg",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eCharmed Minimum Damage\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Charm Max Dmg",
              "model": "charm_max_dmg",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eCharmed Maximum Damage\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Charm Attack Delay",
              "model": "charm_attack_delay",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eCharmed Attack Delay\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Charm Accuracy Rating",
              "model": "charm_accuracy_rating",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eCharmed Accuracy\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Charm Avoidance Rating",
              "model": "charm_avoidance_rating",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eCharmed Avoidance\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Charm Atk",
              "model": "charm_atk",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eCharmed Attack\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Skip Global Loot",
              "model": "skip_global_loot",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eSkip Global Loot: 0 = False, 1 = True\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Rare Spawn",
              "model": "rare_spawn",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eRare Spawn: 0 = False, 1 = True\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Stuck Behavior",
              "model": "stuck_behavior",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eStuck Behavior\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Model",
              "model": "model",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eModel\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Flymode",
              "model": "flymode",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"https://eqemu.gitbook.io/server/categories/npc/fly-modes\" target=\"_blank\"\u003eFly Mode\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Always Aggro",
              "model": "always_aggro",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eAggro regardless of _int or level : 0 = False, 1 = True\u003c/p\u003e\n"
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
