<template>
  <div class="item-bg" style="max-width: 475px; padding: 5px" v-if="itemData">

    <span
      style="position: absolute; right: 7%; opacity: .8"
      :class="'mt-2 mb-2 object-ctn-' + itemData.idfile.replace('IT', '')"
    />

    <div
      class="col-1 text-right"
      v-if="showEdit"
      style="position: absolute; left: 4%; top: 70px; z-index: 9999; opacity: .7"
    >
      <b-button
        variant="outline-warning"
        v-if="showEdit"
        @click="editItem(itemData.id)"
        class="mt-2"
        size="sm"
      >
        Edit
      </b-button>
    </div>

    <div class="row">
      <div class="col-1">
        <span :class="'fade-in item-' + itemData.icon" :title="itemData.icon">
<!--          <span-->
<!--            v-if="itemData.stacksize > 1"-->
<!--            style="position:absolute; right: 0px; top:45px; font-size: 10px">-->
<!--            ({{ itemData.stacksize }})-->
<!--          </span>-->
        </span>
      </div>
      <div class="col-8 pl-5">
        <h6 class="eq-header" style="margin: 0px; margin-bottom: 10px">
          {{ itemData.name }}
        </h6>

        <!-- Top level -->
        <div class="mb-3 mt-3">
          <table>
            <tbody>

            <tr>
              <td colspan="2"> {{ getItemTags() }}</td>
            </tr>

            <tr v-for="(value, stat) in toplevel">
              <td v-if="value !== '' && value !== 0">
                {{ stat }}: {{ value }}
              </td>
            </tr>

            <tr>
              <td colspan="2"> {{ getSlots() }}</td>
            </tr>

            </tbody>
          </table>
        </div>

      </div>
    </div>

    <div class="row">
      <div class="col-12">

        <!-- 2nd level -->
        <div class="mb-3 row">
          <!-- First Section -->
          <div class="stat-section col-4">
            <table style="width: 125px;" class="item-preview-table">
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
            <table style="width: 140px;">
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
          <div class="stat-section col-4">
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
                <td
                  style="font-weight: bold; min-width: 95px"
                  v-if="itemData[data.stat] > 0 || itemData[data.heroic] > 0"
                >
                  {{ stat }}
                </td>

                <!-- Regular stat -->
                <td style="text-align: right" v-if="itemData[data.stat] > 0 || itemData[data.heroic] > 0">
                  {{ itemData[data.stat] }}
                </td>

                <!-- Heroic -->
                <td style="text-align: right; color: #ffecca" v-if="itemData[data.heroic] > 0">
                  {{ itemData[data.heroic] > 0 ? "+" + itemData[data.heroic] : itemData[data.heroic] }}
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
    </div>

    <!-- Extra Damage Amount -->
    <div v-if="itemData['extradmgamt'] > 0" class="mb-1 row">
      <div class="col-12">
        <span style="font-weight: bold" class="pr-2">{{ getExtraDmgSkill() }} Damage </span> +{{ itemData.extradmgamt }}
      </div>
    </div>

    <!-- Bard Skill -->
    <div v-if="itemData['bardtype'] > 22 && itemData['bardtype'] < 65535" class="mb-1 row">
      <div class="col-12">
        <span style="font-weight: bold" class="pr-2">Bard Skill ({{ getBardSkill() }}) </span>
        {{ ((itemData.bardvalue * 10) - 100) }}%
      </div>
    </div>

    <!-- Skill Mod Type -->
    <div class="row" v-if="itemData['skillmodtype'] > 0 && itemData['skillmodvalue'] !== 0">
      <div class="mb-1 col-12">
        <span style="font-weight: bold" class="pr-2">Skill Mod ({{ getSkillModSkill() }}) </span>
        +{{ itemData.skillmodvalue }}
      </div>
    </div>

    <!-- Augmentation Type -->
    <div class="row" v-if="itemData['itemtype'] === 54">
      <div class="col-12">
        <div class="mt-3">
          <div style="font-weight: bold" class="pr-2">Augmentation Slot Type(s)</div>
          <div v-for="(augType, index) in getAugSlotTypes()" :key="augType">
            {{ augType }}
          </div>
          <div class="mt-2" v-if="itemData.augrestrict > 0">
            <span style="font-weight: bold" class="pr-2">Augmentation Restriction </span> {{ getAugRestriction() }}
          </div>
        </div>
      </div>
    </div>

    <!-- Augment Slots -->
    <div class="mb-3 mt-3">
      <div v-for="(n, i) in 5">
        <div class="pl-0 row mb-1 " v-if="itemData['augslot_' + n + '_type'] > 0">
          <div class="col-12">
            <img
              src='~@/assets/img/icons/inventory/blank_slot.gif'
              class="pr-3"
              style='width:auto;height:15px'
            >
            <span style="font-weight: bold">Slot {{ n }}</span> Type {{ itemData["augslot_" + n + "_type"] }}
            {{ getAugTypeDescription(itemData["augslot_" + n + "_type"]) }}
          </div>
        </div>
      </div>
    </div>

    <!-- Effects -->
    <div class="mb-3">
      <div v-for="effect in effects" :key="effect.field" class="col-12">
        <div v-if="itemData[effect.field] > 0 && effectData[effect.field]" class="row col-12 pl-0 mb-1">

          <!-- Target -->
          <div
            :id="itemData[effect.field] + '-' + componentId"
          >
            <span
              :style="'width: 20px; height: 20px; border: 1px solid ' + getTargetTypeColor(effectData[effect.field]['targettype']) + '; border-radius: 3px; display: inline-block'"
              :class="'spell-' + effectData[effect.field].new_icon + '-20'"
            />
            <span
              class="ml-1"
              style="color: #f7ff00; position: relative; top: -7px;"
            >
              {{ effectData[effect.field].name }} ({{ effect.name }})
              <div v-if="['clickeffect'].includes(effect.field)" class="d-inline-block">
                Cast ({{ itemData['casttime'] / 1000 }} sec) Recast ({{ itemData['recastdelay'] }} sec)
              </div>
            </span>

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

    <!-- Bag Weight Reduction -->
    <div v-if="itemData.bagwr" class="mb-1 row">
      <div class="col-12">
        <span style="font-weight: bold" class="mr-1">Bag Weight Reduction</span> {{ itemData.bagwr }}%
      </div>
    </div>

    <!-- Lore -->
    <div v-if="itemData.lore" class="mb-1 row">
      <div class="col-12">
        <span style="font-weight: bold" class="mr-1">Lore</span> {{ itemData.lore }}
      </div>
    </div>

    <!-- Faction -->
    <div v-for="i in 4" :key="i">
      <div class="mb-1 row" v-if="itemData['factionmod_' + i] > 0 && factionNames[itemData['factionmod_' + i]]">
        <div class="col-12">
          <span style="font-weight: bold" class="mr-1">Faction</span>
          {{ (parseInt(itemData['factionamt_' + i]) > 0 ? "Increases" : "Decreases") }} your faction of
          <span style="font-weight: bold">{{ factionNames[itemData['factionmod_' + i]] }}</span> by
          {{ Math.abs(itemData['factionamt_' + i]) }} point(s)
        </div>
      </div>
    </div>

    <!-- Price -->
    <div v-if="itemData.price > 0">
      <span style="font-weight: bold" class="mr-2">Price</span>
      <eq-cash-display
        class="d-inline-block"
        :price="itemData.price"
      />
    </div>

    <div class="pb-4 mb-5"></div>

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
}                                      from "@/app/constants/eq-item-constants";
import {BODYTYPES}                     from "@/app/constants/eq-bodytype-constants";
import {DB_CLASSES_WEAR_SHORT}         from "@/app/constants/eq-classes-constants";
import {DB_RACE_NAMES, DB_RACES_SHORT} from "@/app/constants/eq-races-constants";
import {DB_DIETIES}                    from "@/app/constants/eq-deities-constants";
import EqDebug                         from "@/components/eq-ui/EQDebug";
import {App}                           from "@/constants/app";
import EqSpellPreview                  from "@/components/eq-ui/EQSpellCardPreview";
import {EXAMPLE_SPELL_DATA}            from "@/app/constants/eq-example-spell-data";
import EqWindow                        from "@/components/eq-ui/EQWindow";
import {DB_BARD_SKILLS, DB_SKILLS}     from "@/app/constants/eq-skill-constants";
import {AUG_TYPES}                     from "@/app/constants/eq-aug-constants";
import {Spells}                        from "@/app/spells";
import util                            from "util";
import {ROUTE}                         from "@/routes";
import EqCashDisplay                   from "@/components/eq-ui/EqCashDisplay";
import {Items}                         from "@/app/items";
import {FactionListApi}                from "@/app/api";
import {SpireApiClient}                from "@/app/api/spire-api-client";

export default {
  name: "EqItemCardPreview",
  components: { EqCashDisplay, EqWindow, EqSpellPreview, EqDebug },
  data() {
    return {
      spells: EXAMPLE_SPELL_DATA,
      cdnUrl: App.ASSET_CDN_BASE_URL,
      componentId: "",
      factionNames: [],
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
        "Poison Resist": { stat: "pr", heroic: "heroic_pr" },
        "Corruption": { stat: "svcorruption", heroic: "heroic_svcorrup" }
      },
      mod3: Items.getMod3Fields(),
      toplevel: {
        "Class": this.getClasses(),
        "Race": this.getRaces(),
        "Deity": this.getDeity()
      },
      secondlevel1: {
        "Size": this.getItemSize(this.itemData.size).toUpperCase(),
        "Weight": this.itemData.weight / 10,
        "Light": this.itemData.light,
        "Item Type": this.getItemType(),
        "Rec Level": this.itemData.reclevel,
        "Req Level": this.itemData.reqlevel
      },
      secondlevel2: {
        "AC": this.itemData.ac,
        "HP": this.itemData.hp,
        "Hana": this.itemData.mana,
        "End": this.itemData.endur,
        "Haste": this.itemData.haste > 0 ? (this.itemData.haste + "%") : this.itemData.haste,
        "Purity": this.itemData.purity
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
    async init() {
      const uuidv4     = require("uuid/v4")
      this.componentId = uuidv4()

      // dynamic section builder
      this.secondlevel3   = {}
      let data            = {};
      data["Base Damage"] = this.itemData.damage
      if (this.itemData.elemdmgamt > 0) {
        data[this.getElementDamageName() + " Damage"] = this.itemData.elemdmgamt
      }
      if (this.itemData.banedmgrace > 0 && this.itemData.banedmgraceamt !== 0) {
        data["Bane (" + this.getBaneDamageName() + ")"] = this.itemData.banedmgraceamt
      }
      if (this.itemData.banedmgbody > 0 && this.itemData.banedmgamt !== 0) {
        data["Bane (" + this.getBaneDamageBodyName() + ")"] = this.itemData.banedmgamt
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
          Spells.getSpell(this.itemData[effect.field]).then((spell) => {
            this.effectData[effect.field] = spell
            this.$forceUpdate()
          })
        }
      })

      let factions = []
      for (let i = 1; i <= 4; i++) {
        if (this.itemData['factionmod_' + i]) {
          factions.push(this.itemData['factionmod_' + i])
        }
      }

      if (factions.length > 0) {
        const response = await (new FactionListApi(SpireApiClient.getOpenApiConfig())).getFactionListsBulk({
          body: {
            ids: factions
          }
        })

        if (response.status === 200 && response.data && response.data.length > 0) {
          response.data.forEach((faction) => {
            this.factionNames[faction.id] = faction.name
          })
        }
      }

      this.secondlevel3 = data
    },
    editItem(itemId) {
      this.$router.push(
        {
          path: util.format(ROUTE.ITEM_EDIT, itemId),
          query: {}
        }
      ).catch(() => {
      })
    },
    getTargetTypeColor(targetType) {
      return Spells.getTargetTypeColor(targetType)
    },
    getItemSize: function (size) {
      return ITEM_SIZE[parseInt(size)] ? ITEM_SIZE[parseInt(size)] : "N/A";
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
      if (this.itemData.fvnodrop === 1) {
        tags.push("FV No Drop");
      }
      if (this.itemData.book === 1) {
        tags.push("Book");
      }
      if (this.itemData.tradeskills > 0) {
        tags.push("Tradeskill Item");
      }
      if (this.itemData.notransfer === 1) {
        tags.push("No Transfer");
      }
      if (this.itemData.summonedflag === 1) {
        tags.push("Summoned");
      }
      if (this.itemData.questitemflag === 1) {
        tags.push("Quest Item");
      }
      if (this.itemData.artifactflag === 1) {
        tags.push("Artifact");
      }
      if (this.itemData.nopet === 1) {
        tags.push("No Pet");
      }
      if (this.itemData.attuneable === 1) {
        tags.push("Attuneable");
      }
      if (this.itemData.stackable === 1) {
        tags.push("Stackable (" + this.itemData.stacksize + ")");
      }
      if (this.itemData.potionbelt === 1) {
        tags.push("Potion Belt");
      }
      if (this.itemData.placeable === 1) {
        tags.push("Placeable");
      }
      if (this.itemData.bardtype > 0) {
        tags.push("Instrument");
      }
      if (this.itemData.epicitem > 0) {
        tags.push("Epic");
      }
      if (this.itemData.expendablearrow > 0) {
        tags.push("Arrow Expendable");
      }
      if (this.itemData.heirloom > 0) {
        tags.push("Heirloom");
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

      return this.itemData.classes >= 65535 ? 'ALL' : classes.join(", ").trim()
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

      return this.itemData.races >= 65535 ? 'ALL' : races.join(", ").trim()
    },
    getDeity: function () {
      let deities      = []
      let deitiesValue = this.itemData.deity

      for (let [key, value] of Object.entries(DB_DIETIES).reverse()) {
        key = parseInt(key)
        if (key <= deitiesValue) {
          deitiesValue -= key;
          deities.push(value)
        }
      }

      return this.itemData.deity >= 65535 ? 'ALL' : deities.join(", ").trim()
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

      return this.itemData.augtype > 0 ? augSlots : ["All Slots"]
    },
    getExtraDmgSkill: function () {
      return DB_SKILLS[this.itemData.extradmgskill] ? (DB_SKILLS[this.itemData.extradmgskill].replace("_", " ")) : ""
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
      type = parseInt(type)
      return AUG_TYPES[type] ? AUG_TYPES[type].name : ""
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
    this.init()
  },
  props: {
    itemData: {
      type: Object,
      default: {},
      required: true
    },
    showEdit: {
      type: Boolean,
      default: false,
      required: false,
    }
  }
}
</script>

<style>
.stat-section {
  padding-left: 10px;
  display: inline-block;
  vertical-align: top;
}

.item-preview-table {
  word-wrap: break-word;
  width: 100%;
}

.item-preview-table th, td {
  word-wrap: break-word;
}

</style>
