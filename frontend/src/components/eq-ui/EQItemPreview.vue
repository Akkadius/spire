<template>
  <div class="item-bg" style="max-width: 400px; padding: 5px">

    <div class="row">
      <div class="col-1">
        <img :src="cdnUrl + 'assets/item_icons/item_' + itemData.icon + '.png'" style="width:40px;height:auto;">
      </div>
      <div class="col-11 pl-5">
        <h6 class="eq-header" style="margin: 0px; margin-bottom: 10px">
          {{ itemData.name }}
        </h6>

        <!-- Top level -->
        <div class="mb-3 mt-3">
          <table>
            <tbody>

            <tr>
              <td colspan="2" nowrap="1"> {{ getItemTags() }}</td>
            </tr>

            <tr v-for="(value, stat) in toplevel">
              <td v-if="value !== '' && value !== 0">
                {{ stat }}: {{ value }}
              </td>
            </tr>

            <tr>
              <td colspan="2" nowrap="1"> {{ getSlots() }}</td>
            </tr>

            </tbody>
          </table>
        </div>

      </div>
    </div>

    <div>

      <!-- 2nd level -->
      <div class="mb-3 row">
        <!-- First Section -->
        <div class="stat-section col-4" style="padding-left: 0px">
          <table style="width: 125px;">
            <tbody>

            <tr v-for="(value, stat) in secondlevel1">
              <!-- Label -->
              <td style="font-weight: bold" v-if="value !== '' && value !== 0">
                {{ stat }}
              </td>

              <!-- Regular stat -->
              <td style="text-align: right" v-if="value !== '' && value !== 0">
                {{ value }}
              </td>

            </tr>
            </tbody>
          </table>
        </div>

        <!-- Second Section (AC / HP / Mana / End) -->
        <div class="stat-section col-4">
          <table style="width: 125px;">
            <tbody>

            <tr v-for="(value, stat) in secondlevel2">
              <!-- Label -->
              <td style="font-weight: bold" v-if="value !== '' && value !== 0">
                {{ stat }}
              </td>

              <!-- Regular stat -->
              <td style="text-align: right" v-if="value !== '' && value !== 0">
                {{ value }}
              </td>

            </tr>
            </tbody>
          </table>
        </div>

        <!-- Third Section (Weapon Damage) -->
        <div class="stat-section col-4">
          <table style="width: 125px;">
            <tbody>

            <tr v-for="(value, stat) in secondlevel3">
              <!-- Label -->
              <td style="font-weight: bold" v-if="value !== '' && value !== 0">
                {{ stat }}
              </td>

              <!-- Regular stat -->
              <td style="text-align: right" v-if="value !== '' && value !== 0">
                {{ value }}
              </td>

            </tr>
            </tbody>
          </table>
        </div>
      </div>

      <!-- 3rd level -->
      <div class="mb-3 row">

        <!-- Stats -->
        <div class="stat-section col-4" style="padding-left: 0px">
          <table style="width:100%">
            <tbody>
            <tr v-for="(data, stat) in stats">

              <!-- Label -->
              <td style="font-weight: bold" v-if="itemData[data.stat] > 0 || itemData[data.heroic] > 0">
                {{ stat }}
              </td>

              <!-- Regular stat -->
              <td style="text-align: right" v-if="itemData[data.stat] > 0 || itemData[data.heroic] > 0">
                {{ itemData[data.stat] }}
              </td>

              <!-- Heroic -->
              <td style="text-align: right" v-if="itemData[data.heroic] > 0">
                          <span style="color: #ffecca" v-if="itemData[data.heroic]">
                            {{ itemData[data.heroic] > 0 ? "+" + itemData[data.heroic] : itemData[data.heroic] }}
                          </span>
              </td>
            </tr>
            </tbody>
          </table>
        </div>

        <div class="stat-section col-4">
          <table style="width:100%">
            <tbody>
            <tr v-for="(data, stat) in resists">

              <!-- Label -->
              <td style="font-weight: bold; min-width: 95px" v-if="itemData[data.stat] > 0 || itemData[data.heroic] > 0">
                {{ stat }}
              </td>

              <!-- Regular stat -->
              <td style="text-align: right" v-if="itemData[data.stat] > 0 || itemData[data.heroic] > 0">
                {{ itemData[data.stat] }}
              </td>

              <!-- Heroic -->
              <td style="text-align: right" v-if="itemData[data.heroic] > 0">
                <span style="color: #ffecca" v-if="itemData[data.heroic]">
                  {{ itemData[data.heroic] > 0 ? "+" + itemData[data.heroic] : itemData[data.heroic] }}
                </span>
              </td>
            </tr>
            </tbody>
          </table>
        </div>

        <div class="stat-section col-4">
          <table style="width:100%">
            <tbody>
            <tr v-for="(field, stat) in mod3">

              <!-- Label -->
              <td style="font-weight: bold" v-if="itemData[field] > 0">
                {{ stat }}
              </td>

              <!-- Regular stat -->
              <td style="text-align: right" v-if="itemData[field] > 0">
                {{ itemData[field] }}
              </td>

            </tr>
            </tbody>
          </table>
        </div>
      </div>
    </div>

    <!-- TODO: Price -->

    <!-- Extra Damage Amount -->
    <div v-if="itemData['extradmgamt'] > 0" class="mt-3 mb-3">
      <div class="row">
        <span style="font-weight: bold" class="pr-2">{{ getExtraDmgSkill() }} Damage </span> +{{ itemData.extradmgamt }}
      </div>
    </div>

    <!-- Bard Skill -->
    <div v-if="itemData['bardtype'] > 22 && itemData['bardtype'] < 65535" class="mt-3 mb-3">
      <div class="row">
        <span style="font-weight: bold" class="pr-2">Bard Skill ({{ getBardSkill() }}) </span>
        {{ ((itemData.bardvalue * 10) - 100) }}%
      </div>
    </div>

    <!-- Skill Mod Type -->
    <div v-if="itemData['skillmodtype'] > 0 && itemData['skillmodvalue'] !== 0" class="mt-3 mb-3">
      <div class="row">
        <span style="font-weight: bold" class="pr-2">Skill Mod ({{ getSkillModSkill() }}) </span>
        +{{ itemData.skillmodvalue }}
      </div>
    </div>

    <!-- Augmentation Type -->
    <div v-if="itemData['itemtype'] === 54" class="mt-3 mb-3">
      <div class="row">
        <span style="font-weight: bold" class="pr-2">Augmentation Slot Type(s) </span> {{ getAugSlotTypes() }}
      </div>
      <div class="row" v-if="itemData.augrestrict > 0">
        <span style="font-weight: bold" class="pr-2">Augmentation Restriction </span> {{ getAugRestriction() }}
      </div>
    </div>

    <!-- Augment Slots -->
    <div class="pl-0">
      <div v-for="(n, i) in 5" class="row">
        <div v-if="itemData['augslot_' + n + '_type'] > 0">
          <img
            src='~@/assets/img/icons/inventory/blank_slot.gif'
            class="pr-3"
            style='width:auto;height:15px'>
          <span style="font-weight: bold">Slot {{ n }}</span> Type {{ itemData["augslot_" + n + "_type"] }}
          {{ getAugTypeDescription(itemData["augslot_" + n + "_type"]) }}
        </div>
      </div>
    </div>

    <!-- Effects -->
    <div class="row col-12 pl-0 pt-3 pb-3">
      <div v-for="effect in effects" :key="effect.field">

        <!-- Click Effect -->
        <div v-if="itemData[effect.field] > 0 && effectData[effect.field]">

          <!-- Target -->
          <div :id="itemData[effect.field] + '-' + componentId">

            <div style="display: inline-block">
              <img
                :src="spellCdnUrl + effectData[effect.field].new_icon + '.gif'"
                style="height:15px; border-radius: 25px; width:auto;"
                class="mr-3">
              <span style="font-weight: bold" class="mr-2">Effect</span>
            </div>

            <div style="display: inline-block">

              {{ effectData[effect.field].name }} ({{ effect.name }})
            </div>

          </div>

          <!-- Popover -->
          <b-popover
            :target="itemData[effect.field] + '-' + componentId"
            placement="auto"
            custom-class="no-bg"
            delay="1"
            triggers="hover focus"
            style="width: 500px !important"
          >
            <eq-window style="margin-right: 10px; width: auto; height: 90%">
              <eq-spell-preview :spell-data="effectData[effect.field]"/>
            </eq-window>
          </b-popover>
        </div>

      </div>
    </div>

    <div class="pb-4"></div>

    <eq-debug :data="itemData"/>
  </div>
</template>

<script>

import {
  DB_ITEM_AUG_RESTRICT,
  DB_ITEM_TYPES,
  ITEM_DB_SLOTS,
  ITEM_ELEMENTS,
  ITEM_SIZE
} from "@/app/constants/eq-item-constants";
import {BODYTYPES}                     from "@/app/constants/eq-bodytype-constants";
import {DB_CLASSES_WEAR_SHORT}         from "@/app/constants/eq-classes-constants";
import {DB_RACE_NAMES, DB_RACES_SHORT} from "@/app/constants/eq-races-constants";
import {DB_DIETIES} from "@/app/constants/eq-deities-constants";
import EqDebug from "@/components/eq-ui/EQDebug";
import {App} from "@/constants/app";
import EqSpellPreview from "@/components/eq-ui/EQSpellPreview";
import {EXAMPLE_SPELL_DATA} from "@/app/constants/eq-example-spell-data";
import EqWindow from "@/components/eq-ui/EQWindow";
import {SpireApiClient} from "@/app/api/spire-api-client";
import {SpellsNewApi} from "@/app/api";
import {DB_BARD_SKILLS, DB_SKILLS} from "@/app/constants/eq-skill-constants";
import {AUG_TYPES} from "@/app/constants/eq-aug-constants";

export default {
  name: "EqItemPreview",
  components: { EqWindow, EqSpellPreview, EqDebug },
  data() {
    return {
      spells: EXAMPLE_SPELL_DATA,
      cdnUrl: App.ASSET_CDN_BASE_URL,
      spellCdnUrl: App.ASSET_SPELL_ICONS_BASE_URL,
      componentId: "",
      stats: {
        "Strength": { stat: "astr", heroic: "heroic_str" },
        "Stamina": { stat: "asta", heroic: "heroic_sta" },
        "Intelligence": { stat: "aint", heroic: "heroic_int" },
        "Wisdom": { stat: "awis", heroic: "heroic_wis" },
        "Agility": { stat: "aagi", heroic: "heroic_agi" },
        "Dexterity": { stat: "adex", heroic: "heroic_dex" },
        "Charisma": { stat: "acha", heroic: "heroic_cha" }
      },
      resists: {
        "Magic Resist": { stat: "mr", heroic: "heroic_mr" },
        "Fire Resists": { stat: "fr", heroic: "heroic_fr" },
        "Cold Resist": { stat: "cr", heroic: "heroic_cr" },
        "Disease Resist": { stat: "dr", heroic: "heroic_dr" },
        "Poison Resist": { stat: "pr", heroic: "heroic_pr" }
      },
      mod3: {
        "Attack": "attack",
        "HP Regen": "regen",
        "Mana Regen": "manaregen",
        "Endurance Regen": "enduranceregen",
        "Spell Shielding": "spellshield",
        "Combat Effects": "combateffects",
        "Shielding": "shielding",
        "DoT Shielding": "dotshielding",
        "Avoidance": "avoidance",
        "Accuracy": "accuracy",
        "Stun Resist": "stunresist",
        "Strikethrough": "strikethrough",
        "Damage Shield": "damageshield"
        // TODO: extradmgamt
      },
      toplevel: {
        "Class": this.getClasses(),
        "Race": this.getRaces(),
        "Deity": this.getDeity()
      },
      secondlevel1: {
        "Size": this.getItemSize(this.itemData.size).toUpperCase(),
        "Weight": this.itemData.weight / 10,
        "Item Type": this.getItemType(),
        "Rec Level": this.itemData.reclevel,
        "Req Level": this.itemData.reqlevel
      },
      secondlevel2: {
        "AC": this.itemData.ac,
        "HP": this.itemData.hp,
        "Hana": this.itemData.mana,
        "End": this.itemData.endur,
        "Haste": this.itemData.haste > 0 ? (this.itemData.haste + "%") : this.itemData.haste
      },
      secondlevel3: {},

      effectData: {}, // stores effect data when loaded from API
      effects: [
        { field: "proceffect", name: "Proc" },
        { field: "worneffect", name: "Worn Effect" },
        { field: "focuseffect", name: "Focus Effect" },
        { field: "scrolleffect", name: "Scroll Effect" },
        { field: "clickeffect", name: "Click Effect" },
        { field: "bardeffect", name: "Bard Effect" }
      ]
      // augslots: {}
    }
  },
  methods: {
    getItemSize: function (size) {
      return ITEM_SIZE[size] ? ITEM_SIZE[size] : "N/A";
    },
    getItemTags: function () {
      let tags = [];
      if (this.itemData.itemtype === 54) {
        tags.push("Augment");
      }
      if (this.itemData.magic === 1) {
        tags.push("Magic");
      }
      if (this.itemData.loregroup === -1) {
        tags.push("Lore");
      }
      if (this.itemData.nodrop === 0) {
        tags.push("No Trade");
      }
      if (this.itemData.norent === 0) {
        tags.push("No Rent");
      }

      return tags.join(", ")
    },
    getItemType: function () {
      return DB_ITEM_TYPES[this.itemData.itemtype] ? DB_ITEM_TYPES[this.itemData.itemtype] : "";
    },
    getClasses: function () {
      let classes      = []
      let classesValue = this.itemData.classes
      for (const [key, value] of Object.entries(DB_CLASSES_WEAR_SHORT).reverse()) {

        if (key <= classesValue) {
          classesValue -= key;
          classes.push(value)
        }
      }

      return classes.join(", ").trim()
    },
    getRaces: function () {
      let races      = []
      let racesValue = this.itemData.races
      for (const [key, value] of Object.entries(DB_RACES_SHORT).reverse()) {

        if (key <= racesValue) {
          racesValue -= key;
          races.push(value)
        }
      }

      return races.join(", ").trim()
    },
    getDeity: function () {
      let deities      = []
      let deitiesValue = this.itemData.deities
      for (const [key, value] of Object.entries(DB_DIETIES).reverse()) {

        if (key <= deitiesValue) {
          deitiesValue -= key;
          deities.push(value)
        }
      }

      return deities.join(", ").trim()
    },
    getSlots: function () {
      let slots      = []
      let slotsValue = this.itemData.slots
      for (const [key, value] of Object.entries(ITEM_DB_SLOTS).reverse()) {

        if (key <= slotsValue) {
          slotsValue -= key;
          slots.push(value)
        }
      }

      return slots.join(", ").trim()
    },
    getAugSlotTypes: function () {
      let augType  = this.itemData.augtype
      let augSlots = []
      let bit      = 1;
      for (let i = 1; i < 25; i++) {
        if (bit <= augType && bit & augType) {
          augSlots.push(i + " " + this.getAugTypeDescription(i))
        }
        bit *= 2;
      }

      return this.itemData.augtype > 0 ? augSlots.join(", ").trim() : "All Slots"
    },
    getExtraDmgSkill: function () {
      return DB_SKILLS[this.itemData.extradmgskill] ? this.title(DB_SKILLS[this.itemData.extradmgskill].replace("_", " ").toLowerCase()) : ""
    },
    getAugRestriction: function () {
      return DB_ITEM_AUG_RESTRICT[this.itemData.augrestrict] ? DB_ITEM_AUG_RESTRICT[this.itemData.augrestrict] : "Unknown Type (" + this.itemData.augrestrict + ")"
    },
    getBardSkill: function () {
      return DB_BARD_SKILLS[this.itemData.bardtype] ? DB_BARD_SKILLS[this.itemData.bardtype] : "Unknown Bardtype (" + this.itemData.bardtype + ")"
    },
    getSkillModSkill: function () {
      return DB_SKILLS[this.itemData.skillmodtype] ? this.title(DB_SKILLS[this.itemData.skillmodtype].replace("_", " ").toLowerCase()) : ""
    },
    getElementDamageName: function () {
      return ITEM_ELEMENTS[this.itemData.elemdmgtype] ? ITEM_ELEMENTS[this.itemData.elemdmgtype] : ""
    },
    getBaneDamageName: function () {
      return DB_RACE_NAMES[this.itemData.banedmgrace] ? DB_RACE_NAMES[this.itemData.banedmgrace] : ""
    },
    getBaneDamageBodyName: function () {
      return BODYTYPES[this.itemData.banedmgbody] ? BODYTYPES[this.itemData.banedmgbody] : ""
    },
    getAugTypeDescription: function (type) {
      return AUG_TYPES[type] ? AUG_TYPES[type] : ""
    },
    title: function (str) {
      return str.replace(
        /\w\S*/g,
        function (txt) {
          return txt.charAt(0).toUpperCase() + txt.substr(1).toLowerCase();
        }
      );
    }
  },
  created: function () {
    const uuidv4     = require("uuid/v4")
    this.componentId = uuidv4()


    // dynamic section builder
    this.secondlevel3   = {}
    let data            = {};
    data["Base Damage"] = this.itemData.damage
    if (this.itemData.elemdmgamt > 0) {
      data[this.getElementDamageName() + " Damage"] = this.itemData.elemdmgamt
    }
    if (this.itemData.banedmgrace > 0 && this.itemData.banedmgamt !== 0) {
      data["Bane Damage (" + this.getBaneDamageName() + ")"] = this.itemData.banedmgamt
    }
    if (this.itemData.banedmgbody > 0 && this.itemData.banedmgamt !== 0) {
      data[this.getBaneDamageBodyName()] = this.itemData.banedmgamt
    }

    data["Backstab Damage"] = this.itemData.backstabdmg
    data["Delay"]           = this.itemData.delay
    if (this.itemData.damage > 0) {
      data["Ratio"] = Math.round(this.itemData.damage / this.itemData.delay * 100) / 100
    }
    // TODO: Damage bonus
    data["Range"] = this.itemData.range

    // spell loading
    this.effects.forEach((effect) => {
      if (this.itemData[effect.field] > 0) {
        (new SpellsNewApi(SpireApiClient.getOpenApiConfig())).getSpellsNew({ id: this.itemData[effect.field] }).then((result) => {
          if (result.status === 200) {
            this.effectData[effect.field] = result.data
            this.$forceUpdate()
          }
        })
      }
    })

    this.secondlevel3 = data
  },
  props: {
    itemData: Object
  }
}
</script>

<style>
.stat-section {
  padding-left:   10px;
  display:        inline-block;
  vertical-align: top;
}

</style>
