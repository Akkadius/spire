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
              "help": "\u003cp\u003eUnique Item Identifier\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Minstatus",
              "model": "minstatus",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"https://eqemu.gitbook.io/server/categories/player/status-levels\" target=\"_blank\"\u003eMinimum Status\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Name",
              "model": "Name",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eName\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Aagi",
              "model": "aagi",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eAgility: -128 to 127\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Ac",
              "model": "ac",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eArmor Class\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Accuracy",
              "model": "accuracy",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eAccuracy\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Acha",
              "model": "acha",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eCharisma: -128 to 127\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Adex",
              "model": "adex",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eDexterity: -128 to 127\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Aint",
              "model": "aint",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eIntelligence: -128 to 127\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Artifactflag",
              "model": "artifactflag",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eArtifact: 0 = False, 1 = True\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Asta",
              "model": "asta",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eStamina: -128 to 127\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Astr",
              "model": "astr",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eStrenght: -128 to 127\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Attack",
              "model": "attack",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eAttack\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Augrestrict",
              "model": "augrestrict",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"https://eqemu.gitbook.io/server/categories/items/augment-restrictions\" target=\"_blank\"\u003eAugment Restriction\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Augslot1type",
              "model": "augslot1type",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"https://eqemu.gitbook.io/server/categories/items/augment-types\" target=\"_blank\"\u003eAugment Slot 1 Type\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Augslot1visible",
              "model": "augslot1visible",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eAugment Slot 1 Visible: 0 = False, 1 = True\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Augslot2type",
              "model": "augslot2type",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"https://eqemu.gitbook.io/server/categories/items/augment-types\" target=\"_blank\"\u003eAugment Slot 2 Type\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Augslot2visible",
              "model": "augslot2visible",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eAugment Slot 2 Visible: 0 = False, 1 = True\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Augslot3type",
              "model": "augslot3type",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"https://eqemu.gitbook.io/server/categories/items/augment-types\" target=\"_blank\"\u003eAugment Slot 3 Type\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Augslot3visible",
              "model": "augslot3visible",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eAugment Slot 3 Visible: 0 = False, 1 = True\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Augslot4type",
              "model": "augslot4type",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"https://eqemu.gitbook.io/server/categories/items/augment-types\" target=\"_blank\"\u003eAugment Slot 4 Type\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Augslot4visible",
              "model": "augslot4visible",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eAugment Slot 4 Visible: 0 = False, 1 = True\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Augslot5type",
              "model": "augslot5type",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"https://eqemu.gitbook.io/server/categories/items/augment-types\" target=\"_blank\"\u003eAugment Slot 5 Type\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Augslot5visible",
              "model": "augslot5visible",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eAugment Slot 5 Visible: 0 = False, 1 = True\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Augslot6type",
              "model": "augslot6type",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"https://eqemu.gitbook.io/server/categories/items/augment-types\" target=\"_blank\"\u003eAugment Slot 6 Type\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Augslot6visible",
              "model": "augslot6visible",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eAugment Slot 6 Visible: 0 = False, 1 = True\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Augtype",
              "model": "augtype",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"https://eqemu.gitbook.io/server/categories/items/augment-types\" target=\"_blank\"\u003eAugment Type\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Avoidance",
              "model": "avoidance",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eAvoidance\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Awis",
              "model": "awis",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eWisdom: -128 to 127\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Bagsize",
              "model": "bagsize",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"https://eqemu.gitbook.io/server/categories/inventory/bag-sizes\" target=\"_blank\"\u003eBag Size\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Bagslots",
              "model": "bagslots",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eBag Slots: 1 = Minimum, 10 = Maximum\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Bagtype",
              "model": "bagtype",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"https://eqemu.gitbook.io/server/categories/inventory/bag-types\" target=\"_blank\"\u003eBag Type\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Bagwr",
              "model": "bagwr",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eBag Weight Reduction: 0 = 0%, 100 = 100%\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Banedmgamt",
              "model": "banedmgamt",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eBane Damage Amount\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Banedmgraceamt",
              "model": "banedmgraceamt",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eBane Damage Race Amount\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Banedmgbody",
              "model": "banedmgbody",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"https://eqemu.gitbook.io/server/categories/npc/body-types\" target=\"_blank\"\u003eBane Damage Body Type\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Banedmgrace",
              "model": "banedmgrace",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"https://eqemu.gitbook.io/server/categories/npc/race-list\" target=\"_blank\"\u003eBane Damage Race\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Bardtype",
              "model": "bardtype",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"https://eqemu.gitbook.io/server/categories/items/bard-types\" target=\"_blank\"\u003eBard Type\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Bardvalue",
              "model": "bardvalue",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eBard Value\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Book",
              "model": "book",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"books.md\" target=\"_blank\"\u003eBook\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Casttime",
              "model": "casttime",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eCast Time in Seconds\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Casttime ",
              "model": "casttime_",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eCast Time in Seconds\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Charmfile",
              "model": "charmfile",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eCharm File\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Charmfileid",
              "model": "charmfileid",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eCharm File Identifier\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Classes",
              "model": "classes",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"https://eqemu.gitbook.io/server/categories/player/class-list\" target=\"_blank\"\u003eClasses\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Color",
              "model": "color",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eColor\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Combateffects",
              "model": "combateffects",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eCombat Effects\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Extradmgskill",
              "model": "extradmgskill",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"https://eqemu.gitbook.io/server/categories/player/skills\" target=\"_blank\"\u003eExtra Damage Skill\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Extradmgamt",
              "model": "extradmgamt",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eExtra Damage Amount\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Price",
              "model": "price",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003ePrice in Copper\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Cr",
              "model": "cr",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eCold Resistance: -128 to 127\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Damage",
              "model": "damage",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eDamage\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Damageshield",
              "model": "damageshield",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eDamage Shield\u003c/p\u003e\n"
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
              "label": "Delay",
              "model": "delay",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eDelay\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Augdistiller",
              "model": "augdistiller",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eAugment Distiller Item Identifier\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Dotshielding",
              "model": "dotshielding",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eDamage Over Time Shielding\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Dr",
              "model": "dr",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eDisease Resistance: -128 to 127\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Clicktype",
              "model": "clicktype",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"https://eqemu.gitbook.io/server/categories/items/click-types\" target=\"_blank\"\u003eClick Type\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Clicklevel2",
              "model": "clicklevel2",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eClick Level 2\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Elemdmgtype",
              "model": "elemdmgtype",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"https://eqemu.gitbook.io/server/categories/items/item-element-types\" target=\"_blank\"\u003eElemental Damage Type\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Elemdmgamt",
              "model": "elemdmgamt",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eElemental Damage Amount\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Endur",
              "model": "endur",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eEndurance\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Factionamt1",
              "model": "factionamt1",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eFaction Amount 1\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Factionamt2",
              "model": "factionamt2",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eFaction Amount 2\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Factionamt3",
              "model": "factionamt3",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eFaction Amount 3\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Factionamt4",
              "model": "factionamt4",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eFaction Amount 4\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Factionmod1",
              "model": "factionmod1",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eFaction Modifier 1\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Factionmod2",
              "model": "factionmod2",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eFaction Modifier 2\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Factionmod3",
              "model": "factionmod3",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eFaction Modifier 3\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Factionmod4",
              "model": "factionmod4",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eFaction Modifier 4\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Filename",
              "model": "filename",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eFile Name\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Focuseffect",
              "model": "focuseffect",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"spells_new.md\" target=\"_blank\"\u003eFocus Effect Identifier\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Fr",
              "model": "fr",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eFire Resistance: -128 to 127\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Fvnodrop",
              "model": "fvnodrop",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eFiriona Vie No Drop: 0 = False, 1 = True\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Haste",
              "model": "haste",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eHaste: 0 = 0%, 255 = 255%\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Clicklevel",
              "model": "clicklevel",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eClick Level\u003c/p\u003e\n"
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
              "label": "Regen",
              "model": "regen",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eHealth Regeneration\u003c/p\u003e\n"
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
              "label": "Idfile",
              "model": "idfile",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eItem Texture\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Itemclass",
              "model": "itemclass",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"https://eqemu.gitbook.io/server/categories/items/item-class\" target=\"_blank\"\u003eItem Class\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Itemtype",
              "model": "itemtype",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"https://eqemu.gitbook.io/server/categories/items/item-types\" target=\"_blank\"\u003eItem Type\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Ldonprice",
              "model": "ldonprice",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eLDoN Price\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Ldontheme",
              "model": "ldontheme",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"https://eqemu.gitbook.io/server/categories/zones/ldon-themes\" target=\"_blank\"\u003eLDoN Theme\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Ldonsold",
              "model": "ldonsold",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eLDoN Sold: 0 = False, 1 = True\u003c/p\u003e\n"
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
              "label": "Lore",
              "model": "lore",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eLore Description\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Loregroup",
              "model": "loregroup",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"https://eqemu.gitbook.io/server/categories/items/item-lore-groups\" target=\"_blank\"\u003eLore Group\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Magic",
              "model": "magic",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eMagic: 0 = False, 1 = True\u003c/p\u003e\n"
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
              "label": "Manaregen",
              "model": "manaregen",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eMana Regeneration\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Enduranceregen",
              "model": "enduranceregen",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eEndurance Regeneration\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Material",
              "model": "material",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"https://eqemu.gitbook.io/server/categories/npc/textures\" target=\"_blank\"\u003eMaterial\u003c/a\u003e\u003c/p\u003e\n"
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
              "label": "Maxcharges",
              "model": "maxcharges",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eMaximum Charges\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Mr",
              "model": "mr",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eMagic Resistance: -128 to 127\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Nodrop",
              "model": "nodrop",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eNo Drop: 0 = True, 1 = False\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Norent",
              "model": "norent",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eNo Rent: 0 = True, 1 = False\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Pendingloreflag",
              "model": "pendingloreflag",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003ePending Lore Flag: 0 = False, 1 = True\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Pr",
              "model": "pr",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003ePoison Resistance: -128 to 127\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Procrate",
              "model": "procrate",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eProc Rate: 0 = 100%, 50 = 150%, 100 = 200%\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Races",
              "model": "races",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"https://eqemu.gitbook.io/server/categories/npc/race-list\" target=\"_blank\"\u003eRaces\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Range",
              "model": "range",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eRange: 0 to 255\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Reclevel",
              "model": "reclevel",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eRecommended Level\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Recskill",
              "model": "recskill",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eRecommended Skill Level\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Reqlevel",
              "model": "reqlevel",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eRequired Level\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Sellrate",
              "model": "sellrate",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eSell Rate\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Shielding",
              "model": "shielding",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eShielding: 5 = 5%, 20 = 20%, 50 = 50%\u003c/p\u003e\n"
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
              "label": "Skillmodtype",
              "model": "skillmodtype",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"https://eqemu.gitbook.io/server/categories/player/skills\" target=\"_blank\"\u003eSkill Modifier Type\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Skillmodvalue",
              "model": "skillmodvalue",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eSkill Modifier Value\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Slots",
              "model": "slots",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"https://eqemu.gitbook.io/server/categories/inventory/item-slots\" target=\"_blank\"\u003eSlots\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Clickeffect",
              "model": "clickeffect",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"spells_new.md\" target=\"_blank\"\u003eClick Effect Identifier\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Spellshield",
              "model": "spellshield",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eSpell Shielding\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Strikethrough",
              "model": "strikethrough",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eStrikethrough\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Stunresist",
              "model": "stunresist",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eStun Resist\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Summonedflag",
              "model": "summonedflag",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eUnknown\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Tradeskills",
              "model": "tradeskills",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eTradeskill Item: 0 = False, 1= True\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Favor",
              "model": "favor",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eFavor\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Weight",
              "model": "weight",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eWeight: 10 = 1.0, 25 = 2.5, 100 = 10.0\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "UNK012",
              "model": "UNK012",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eUnknown\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "UNK013",
              "model": "UNK013",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eUnknown\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Benefitflag",
              "model": "benefitflag",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eUnknown\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "UNK054",
              "model": "UNK054",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eUnknown\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "UNK059",
              "model": "UNK059",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eUnknown\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Booktype",
              "model": "booktype",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"https://eqemu.gitbook.io/server/categories/player/languages\" target=\"_blank\"\u003eBook Language\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Recastdelay",
              "model": "recastdelay",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eRecast Delay in Seconds\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Recasttype",
              "model": "recasttype",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eRecast Type: -1 = None, \u0026gt;0 = Recast Type used across all items\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Guildfavor",
              "model": "guildfavor",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eGuild Favor\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "UNK123",
              "model": "UNK123",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eUnknown\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "UNK124",
              "model": "UNK124",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eUnknown\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Attuneable",
              "model": "attuneable",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eAttuneable: 0 = False, 1 = True\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Nopet",
              "model": "nopet",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eNo Pet: 0 = False, 1 = True\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Updated",
              "model": "updated",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eUpdated Datetime\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Comment",
              "model": "comment",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eComment\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "UNK127",
              "model": "UNK127",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eUnknown\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Pointtype",
              "model": "pointtype",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eUnknown\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Potionbelt",
              "model": "potionbelt",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003ePotion Belt: 0 = False, 1 = True\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Potionbeltslots",
              "model": "potionbeltslots",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003ePotion Belt Slots\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Stacksize",
              "model": "stacksize",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eStack Size\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Notransfer",
              "model": "notransfer",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eNo Transfer: 0 = False, 1 = True\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Stackable",
              "model": "stackable",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eStackable: 0 = False, 1 = True\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "UNK134",
              "model": "UNK134",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eUnknown\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "UNK137",
              "model": "UNK137",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eUnknown\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Proceffect",
              "model": "proceffect",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"spells_new.md\" target=\"_blank\"\u003eProc Effect Identifier\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Proctype",
              "model": "proctype",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eProc Type: 0\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Proclevel2",
              "model": "proclevel2",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eProc Level 2\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Proclevel",
              "model": "proclevel",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eProc Level\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "UNK142",
              "model": "UNK142",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eUnknown\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Worneffect",
              "model": "worneffect",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"spells_new.md\" target=\"_blank\"\u003eWorn Effect Identifier\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Worntype",
              "model": "worntype",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eWorn Type: 2 = Worn\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Wornlevel2",
              "model": "wornlevel2",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eWorn Level 2\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Wornlevel",
              "model": "wornlevel",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eWorn Level\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "UNK147",
              "model": "UNK147",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eUnknown\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Focustype",
              "model": "focustype",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eFocus Type: 6 = Focus\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Focuslevel2",
              "model": "focuslevel2",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eFocus Level 2\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Focuslevel",
              "model": "focuslevel",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eFocus Level\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "UNK152",
              "model": "UNK152",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eUnknown\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Scrolleffect",
              "model": "scrolleffect",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"spells_new.md\" target=\"_blank\"\u003eScroll Effect Identifier\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Scrolltype",
              "model": "scrolltype",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eScroll Type: 7 = Scroll\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Scrolllevel2",
              "model": "scrolllevel2",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eScroll Level 2\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Scrolllevel",
              "model": "scrolllevel",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eScroll Level\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "UNK157",
              "model": "UNK157",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eUnknown\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Serialized",
              "model": "serialized",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eSerialized Datetime\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Verified",
              "model": "verified",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eVerified Datetime\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Serialization",
              "model": "serialization",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eSerialization\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Source",
              "model": "source",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eSource\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "UNK033",
              "model": "UNK033",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eUnknown\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Lorefile",
              "model": "lorefile",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eLore File\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "UNK014",
              "model": "UNK014",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eUnknown\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Svcorruption",
              "model": "svcorruption",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eCorruption Resistance: -128 to 127\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Skillmodmax",
              "model": "skillmodmax",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eSkill Modifier Max\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "UNK060",
              "model": "UNK060",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eUnknown\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Augslot1unk2",
              "model": "augslot1unk2",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eUnknown\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Augslot2unk2",
              "model": "augslot2unk2",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eUnknown\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Augslot3unk2",
              "model": "augslot3unk2",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eUnknown\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Augslot4unk2",
              "model": "augslot4unk2",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eUnknown\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Augslot5unk2",
              "model": "augslot5unk2",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eUnknown\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Augslot6unk2",
              "model": "augslot6unk2",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eUnknown\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "UNK120",
              "model": "UNK120",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eUnknown\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "UNK121",
              "model": "UNK121",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eUnknown\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Questitemflag",
              "model": "questitemflag",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eQuest Item: 0 = False, 1 = True\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "UNK132",
              "model": "UNK132",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eUnknown\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Clickunk5",
              "model": "clickunk5",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eUnknown\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Clickunk6",
              "model": "clickunk6",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eUnknown\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Clickunk7",
              "model": "clickunk7",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eUnknown\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Procunk1",
              "model": "procunk1",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eUnknown\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Procunk2",
              "model": "procunk2",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eUnknown\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Procunk3",
              "model": "procunk3",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eUnknown\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Procunk4",
              "model": "procunk4",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eUnknown\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Procunk6",
              "model": "procunk6",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eUnknown\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Procunk7",
              "model": "procunk7",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eUnknown\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Wornunk1",
              "model": "wornunk1",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eUnknown\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Wornunk2",
              "model": "wornunk2",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eUnknown\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Wornunk3",
              "model": "wornunk3",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eUnknown\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Wornunk4",
              "model": "wornunk4",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eUnknown\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Wornunk5",
              "model": "wornunk5",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eUnknown\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Wornunk6",
              "model": "wornunk6",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eUnknown\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Wornunk7",
              "model": "wornunk7",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eUnknown\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Focusunk1",
              "model": "focusunk1",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eUnknown\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Focusunk2",
              "model": "focusunk2",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eUnknown\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Focusunk3",
              "model": "focusunk3",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eUnknown\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Focusunk4",
              "model": "focusunk4",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eUnknown\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Focusunk5",
              "model": "focusunk5",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eUnknown\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Focusunk6",
              "model": "focusunk6",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eUnknown\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Focusunk7",
              "model": "focusunk7",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eUnknown\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Scrollunk1",
              "model": "scrollunk1",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eUnknown\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Scrollunk2",
              "model": "scrollunk2",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eUnknown\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Scrollunk3",
              "model": "scrollunk3",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eUnknown\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Scrollunk4",
              "model": "scrollunk4",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eUnknown\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Scrollunk5",
              "model": "scrollunk5",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eUnknown\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Scrollunk6",
              "model": "scrollunk6",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eUnknown\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Scrollunk7",
              "model": "scrollunk7",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eUnknown\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "UNK193",
              "model": "UNK193",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eUnknown\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Purity",
              "model": "purity",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003ePurity\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Evoitem",
              "model": "evoitem",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eEvolving Item: 0 = False, 1 = True\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Evoid",
              "model": "evoid",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"items.md\" target=\"_blank\"\u003eEvolving Identifier\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Evolvinglevel",
              "model": "evolvinglevel",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eEvolving Level\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Evomax",
              "model": "evomax",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eEvolving Max\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Clickname",
              "model": "clickname",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eClick Name\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Procname",
              "model": "procname",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eProc Name\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Wornname",
              "model": "wornname",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eWorn Name\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Focusname",
              "model": "focusname",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eFocus Name\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Scrollname",
              "model": "scrollname",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eScroll Name\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Dsmitigation",
              "model": "dsmitigation",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eDamage Shield Mitigation\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Heroic Str",
              "model": "heroic_str",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eHeroic Strength\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Heroic Int",
              "model": "heroic_int",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eHeroic Intelligence\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Heroic Wis",
              "model": "heroic_wis",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eHeroic Wisdom\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Heroic Agi",
              "model": "heroic_agi",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eHeroic Agility\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Heroic Dex",
              "model": "heroic_dex",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eHeroic Dexterity\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Heroic Sta",
              "model": "heroic_sta",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eHeroic Stamina\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Heroic Cha",
              "model": "heroic_cha",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eHeroic Charisma\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Heroic Pr",
              "model": "heroic_pr",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eHeroic Poison Resistance\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Heroic Dr",
              "model": "heroic_dr",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eHeroic Disease Resistance\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Heroic Fr",
              "model": "heroic_fr",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eHeroic Fire Resistance\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Heroic Cr",
              "model": "heroic_cr",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eHeroic Cold Resistance\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Heroic Mr",
              "model": "heroic_mr",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eHeroic Magic Resistance\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Heroic Svcorrup",
              "model": "heroic_svcorrup",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eHeroic Corruption Resistance\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Healamt",
              "model": "healamt",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eHeal Amount: 0 to 32767\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Spelldmg",
              "model": "spelldmg",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eSpell Damage: 0 to 32767\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Clairvoyance",
              "model": "clairvoyance",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eClairvoyance\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Backstabdmg",
              "model": "backstabdmg",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eBackstab Damage\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Created",
              "model": "created",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eCreated\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Elitematerial",
              "model": "elitematerial",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eElite Material\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Ldonsellbackrate",
              "model": "ldonsellbackrate",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eLDoN Sellback Rate\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Scriptfileid",
              "model": "scriptfileid",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eScript File Name\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Expendablearrow",
              "model": "expendablearrow",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eExpendable Arrow: 0 = False, 1 = True\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Powersourcecapacity",
              "model": "powersourcecapacity",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003ePowersource Capacity\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Bardeffect",
              "model": "bardeffect",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"spells_new.md\" target=\"_blank\"\u003eBard Effect Identifier\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Bardeffecttype",
              "model": "bardeffecttype",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003e\u003ca href=\"https://eqemu.gitbook.io/server/categories/items/bard-types\" target=\"_blank\"\u003eBard Effect Type\u003c/a\u003e\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Bardlevel2",
              "model": "bardlevel2",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eBard Level 2\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Bardlevel",
              "model": "bardlevel",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eBard Level\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Bardunk1",
              "model": "bardunk1",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eUnknown\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Bardunk2",
              "model": "bardunk2",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eUnknown\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Bardunk3",
              "model": "bardunk3",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eUnknown\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Bardunk4",
              "model": "bardunk4",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eUnknown\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Bardunk5",
              "model": "bardunk5",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eUnknown\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Bardname",
              "model": "bardname",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eBard Name\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Bardunk7",
              "model": "bardunk7",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eUnknown\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "UNK214",
              "model": "UNK214",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eUnknown\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Subtype",
              "model": "subtype",
              "styleClasses": "col-12",
              "type": "input"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "UNK220",
              "model": "UNK220",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eUnknown\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "UNK221",
              "model": "UNK221",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eUnknown\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Heirloom",
              "model": "heirloom",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eHeirloom: 0 = False, 1 = True\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "UNK223",
              "model": "UNK223",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eUnknown\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "UNK224",
              "model": "UNK224",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eUnknown\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "UNK225",
              "model": "UNK225",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eUnknown\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "UNK226",
              "model": "UNK226",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eUnknown\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "UNK227",
              "model": "UNK227",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eUnknown\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "UNK228",
              "model": "UNK228",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eUnknown\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "UNK229",
              "model": "UNK229",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eUnknown\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "UNK230",
              "model": "UNK230",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eUnknown\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "UNK231",
              "model": "UNK231",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eUnknown\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "UNK232",
              "model": "UNK232",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eUnknown\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "UNK233",
              "model": "UNK233",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eUnknown\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "UNK234",
              "model": "UNK234",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eUnknown\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Placeable",
              "model": "placeable",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003ePlaceable: 0 = False, 1 = True\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "UNK236",
              "model": "UNK236",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eUnknown\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "UNK237",
              "model": "UNK237",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eUnknown\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "UNK238",
              "model": "UNK238",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eUnknown\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "UNK239",
              "model": "UNK239",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eUnknown\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "UNK240",
              "model": "UNK240",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eUnknown\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "UNK241",
              "model": "UNK241",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eUnknown\u003c/p\u003e\n"
            },
            {
              "featured": false,
              "inputType": "text",
              "label": "Epicitem",
              "model": "epicitem",
              "styleClasses": "col-12",
              "type": "input",
              "help": "\u003cp\u003eEpic Item: 0 = False, 1 = True\u003c/p\u003e\n"
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
