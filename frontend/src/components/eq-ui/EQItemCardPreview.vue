<template>
  <div class="item-bg" style="min-width: 450px; max-width: 500px;" v-if="itemData">
    <span
      v-if="!['IT63', 'IT64'].includes(itemData.idfile)"
      style="position: absolute; right: 7%; opacity: .5;"
      :class="'mt-2 mb-2 object-ctn-' + itemData.idfile.replace('IT', '')"
    />

    <div class="row">
      <div class="col-1">
        <span :class="'item-' + itemData.icon" :title="itemData.icon">
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
        <div class="mt-3">
          <table>
            <tbody>

            <tr>
              <td colspan="2" v-if="showEdit"><a
                href="javascript:void(0)"
                @click="editItem(itemData.id)"
              >
                Edit
              </a></td>
            </tr>
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

    <div class="row pl-2 pr-2">
      <div class="col-12">

        <!-- 2nd level -->
        <div class="row mt-3">
          <!-- First Section -->
          <div class="stat-section col-4">
            <table class="item-preview-table">
              <tbody>

              <tr v-for="(value, stat) in secondlevel1">
                <!-- Label -->
                <td style="font-weight: bold" v-if="value !== '' && value !== 0">
                  {{ stat }}
                </td>

                <!-- Regular stat -->
                <td style="text-align: right" v-if="value !== '' && value !== 0">
                  {{ commify(value) }}
                </td>

              </tr>
              </tbody>
            </table>
          </div>

          <!-- Second Section (AC / HP / Mana / End) -->
          <div class="stat-section col-4">
            <table style="width: 100%">
              <tbody>

              <tr v-for="(value, stat) in secondlevel2">
                <!-- Label -->
                <td style="font-weight: bold" v-if="value !== '' && value !== 0">
                  {{ stat }}
                </td>

                <!-- Regular stat -->
                <td style="text-align: right" v-if="value !== '' && value !== 0">
                  {{ commify(value) }}
                </td>

              </tr>
              </tbody>
            </table>
          </div>

          <!-- Third Section (Weapon Damage) -->
          <div class="stat-section col-4">
            <table style="width: 100%">
              <tbody>

              <tr v-for="(value, stat) in secondlevel3">
                <!-- Label -->
                <td style="font-weight: bold" v-if="value !== '' && value !== 0">
                  {{ stat }}
                </td>

                <!-- Regular stat -->
                <td style="text-align: right" v-if="value !== '' && value !== 0">
                  {{ commify(value) }}
                </td>

              </tr>
              </tbody>
            </table>
          </div>
        </div>

        <!-- 3rd level -->
        <div class="row mt-3">

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
                  {{ commify(itemData[data.stat]) }}
                </td>

                <!-- Heroic -->
                <td style="text-align: right" v-if="itemData[data.heroic] > 0">
                    <span style="color: #ffecca" v-if="itemData[data.heroic]">
                      {{ commify(itemData[data.heroic] > 0 ? "+" + itemData[data.heroic] : itemData[data.heroic]) }}
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
    <div v-if="itemData['extradmgamt'] > 0" class="mt-3 row">
      <div class="col-12">
        <span style="font-weight: bold" class="pr-2">{{ getExtraDmgSkill() }} Damage </span> +{{ itemData.extradmgamt }}
      </div>
    </div>

    <!-- Bard Skill -->
    <div v-if="itemData['bardtype'] > 22 && itemData['bardtype'] < 65535" class="mt-3 row">
      <div class="col-12">
        <span style="font-weight: bold" class="pr-2">Bard Skill ({{ getBardSkill() }}) </span>
        {{ ((itemData.bardvalue * 10) - 100) }}%
      </div>
    </div>

    <!-- Skill Mod Type -->
    <div class="row" v-if="itemData['skillmodtype'] > 0 && itemData['skillmodvalue'] !== 0">
      <div class="mt-3 col-12">
        <span style="font-weight: bold" class="pr-2">Skill Mod ({{ getSkillModSkill() }}) </span>
        +{{ itemData.skillmodvalue }}
      </div>
    </div>

    <!-- Augmentation Type -->
    <div class="row mt-3" v-if="itemData['itemtype'] === 54">
      <div class="col-12">
        <div>
          <div style="font-weight: bold" class="pr-2">Augmentation Slot Type(s)</div>
          <div v-for="(augType, index) in getAugSlotTypes()" :key="augType">
            {{ augType }}
          </div>
          <div class="mt-3" v-if="itemData.augrestrict > 0">
            <span style="font-weight: bold" class="pr-2">Augmentation Restriction </span> {{ getAugRestriction() }}
          </div>
        </div>
      </div>
    </div>

    <!-- Augment Slots -->
    <div class="mt-3" v-if="hasAugs">
      <div v-for="(n, i) in 5">
        <div class="pl-0 row mt-1 " v-if="itemData['augslot_' + n + '_type'] > 0">
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
    <div class="mt-4" v-if="hasEffects">
      <div v-for="effect in effects" :key="effect.field" class="col-12">
        <div v-if="itemData[effect.field] > 0 && effectData[effect.field]" class="row col-12 pl-0 mt-1">

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
            <eq-window style="margin-right: 10px; width: auto; height: 100%">
              <eq-spell-preview :spell-data="effectData[effect.field]"/>
            </eq-window>
          </b-popover>
        </div>
      </div>
    </div>

    <!-- Bag Weight Reduction -->
    <div v-if="itemData.bagwr" class="mt-3 row">
      <div class="col-12">
        <span style="font-weight: bold" class="mr-1">Bag Weight Reduction</span> {{ itemData.bagwr }}%
      </div>
    </div>

    <!-- Lore -->
    <div v-if="itemData.lore" class="mt-3 row">
      <div class="col-12">
        <span style="font-weight: bold" class="mr-1">Lore</span> {{ itemData.lore }}
      </div>
    </div>

    <!-- Faction -->
    <div v-for="i in 4" :key="i">
      <div class="mt-1 row" v-if="itemData['factionmod_' + i] > 0 && factionNames[itemData['factionmod_' + i]]">
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

    <!-- Related Data -->
    <div v-if="showRelatedData" class="mt-3">

      <!-- Unlocks Doors -->
      <div v-if="unlocksDoors.length > 0" class="font-weight-bold mt-3">
        Unlocks Doors
      </div>

      <div class="mt-3" v-if="unlocksDoors && unlocksDoors.length > 0">
        <li v-for="door in unlocksDoors">
          <span class="font-weight-bold">{{ door.name }}</span>
          in <span class="font-weight-bold">{{ door.zone }}</span>
          @ {{ door.x }}, {{ door.y }}, {{ door.z }}
        </li>
      </div>

      <!-- Dropped By -->
      <div v-if="droppedBy.length > 0" class="font-weight-bold mt-3">
        Dropped By
      </div>

      <div class="mt-3" v-if="droppedBy && droppedBy.length > 0">
        <li v-for="drop in droppedBy">
          <span class="font-weight-bold">{{ drop.name }}</span> in <span class="font-weight-bold">{{ drop.zone }}</span>
        </li>
      </div>

      <!-- Fished In -->
      <div v-if="fishedIn.length > 0" class="font-weight-bold mt-3">
        Can be fished in
      </div>

      <div class="mt-3" v-if="fishedIn && fishedIn.length > 0">
        <li v-for="e in fishedIn">
          <span class="font-weight-bold">{{ e.zone.long_name }}</span>
          with skill <span class="font-weight-bold">{{ e.skill }}</span>
        </li>
      </div>

      <!-- Foraged In -->
      <div v-if="foragedIn.length > 0" class="font-weight-bold mt-3">
        Can be foraged in
      </div>

      <div class="mt-3" v-if="foragedIn && foragedIn.length > 0">
        <li v-for="e in foragedIn">
          <span class="font-weight-bold">{{ e.zone.long_name }}</span>
          with skill <span class="font-weight-bold">{{ e.skill }}</span>
          chance <span class="font-weight-bold">{{ e.chance }}</span>
        </li>
      </div>

      <!-- Starting Item -->
      <div v-if="startingItems.length > 0" class="font-weight-bold mt-3">
        Is a starting item for
      </div>

      <div class="mt-3" v-if="startingItems && startingItems.length > 0">
        <li v-for="e in startingItems">
          <span class="font-weight-bold" v-if="e.class !== 0">{{ e.class }}(s)</span>
          <span v-if="e.race !== 0" class="ml-1">race
            <span class="font-weight-bold">{{ e.race }}</span>
          </span>
          <span v-if="e.deity !== 0" class="ml-1">deity
            <span class="font-weight-bold">{{ e.deity }}</span>
          </span>
          <span v-if="e.data.charges !== 0" class="ml-1">count
            <span class="font-weight-bold">({{ e.data.item_charges }})</span>
          </span>
          <span v-if="e.zone !== 0" class="ml-1">
            <span class="font-weight-bold">({{ e.zone.long_name }})</span>
          </span>
        </li>
      </div>

      <!-- Ground Spawns -->
      <div v-if="groundSpawns.length > 0" class="font-weight-bold mt-3">
        Is found as a ground spawn
      </div>

      <div class="mt-3" v-if="groundSpawns && groundSpawns.length > 0">
        <li v-for="e in groundSpawns">
          In <span class="font-weight-bold" v-if="e.zone !== 0">{{ e.zone.long_name }}</span>
          <span v-if="e.data.max_x !== 0" class="ml-1">@
            {{ e.data.max_x }}, {{ e.data.max_y }}, {{ e.data.max_z }}
          </span>
          <span v-if="e.data.respawn_timer !== 0" class="ml-1">respawns every
            <span class="font-weight-bold">({{ Math.round(e.data.respawn_timer / 60) }} minute(s))</span>
          </span>
        </li>
      </div>

      <!-- Tasks -->
      <div v-if="taskRewards.length > 0" class="font-weight-bold mt-3">
        Is a reward in task(s)
      </div>

      <div class="mt-3" v-if="taskRewards && taskRewards.length > 0">
        <li v-for="e in taskRewards">
          <span v-if="e.data.title !== ''" class="font-weight-bold">
            {{ e.data.title }} ({{ e.data.id }})
          </span>
        </li>
      </div>

      <!-- Merchants -->
      <div v-if="merchants.length > 0" class="font-weight-bold mt-3">
        Is sold by merchant(s)
      </div>

      <div class="mt-3" v-if="merchants && merchants.length > 0">
        <li v-for="e in merchants">
          <span v-if="e.merchantName !== ''" class="font-weight-bold">
            {{ e.merchantName }}
          </span>
          <span v-if="e.merchantZone !== ''" class="">in zone
            <span class="font-weight-bold">({{ e.merchantZone }})</span>
          </span>
        </li>
      </div>

      <!-- Tradeskill result -->
      <div v-if="tradeskillResult.length > 0" class="font-weight-bold mt-3">
        Is the result of tradeskill recipe(s)
      </div>

      <div class="mt-3" v-if="tradeskillResult && tradeskillResult.length > 0">
        <li v-for="e in tradeskillResult">
          <span v-if="e.name !== ''" class="font-weight-bold">
            {{ e.name }}
          </span>
          <span v-if="e.tradeskill !== 0">
            <span class="font-weight-bold">({{ e.tradeskill }})</span>
          </span>
          <span v-if="e.trivial !== 0" class="ml-1">Trivial
            <span class="font-weight-bold">({{ e.trivial }})</span>
          </span>
          <span v-if="e.id !== 0" class="ml-1">ID
            <span class="font-weight-bold">({{ e.id }})</span>
          </span>
        </li>
      </div>


    </div>

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
}                                          from "@/app/constants/eq-item-constants";
import {BODYTYPES}                         from "@/app/constants/eq-bodytype-constants";
import {DB_CLASSES, DB_CLASSES_WEAR_SHORT} from "@/app/constants/eq-classes-constants";
import {DB_RACE_NAMES, DB_RACES_SHORT}     from "@/app/constants/eq-races-constants";
import {DB_DIETIES, DB_DIETIES_FULL}       from "@/app/constants/eq-deities-constants";
import EqDebug                             from "@/components/eq-ui/EQDebug";
import {App}                               from "@/constants/app";
import EqSpellPreview                      from "@/components/eq-ui/EQSpellCardPreview";
import {EXAMPLE_SPELL_DATA}                from "@/app/constants/eq-example-spell-data";
import EqWindow                    from "@/components/eq-ui/EQWindow";
import {DB_BARD_SKILLS, DB_SKILLS} from "@/app/constants/eq-skill-constants";
import {AUG_TYPES}                 from "@/app/constants/eq-aug-constants";
import {Spells}                    from "@/app/spells";
import util                        from "util";
import {ROUTE}                     from "@/routes";
import EqCashDisplay               from "@/components/eq-ui/EqCashDisplay";
import {Items}                     from "@/app/items";
import {FactionListApi}            from "@/app/api";
import {SpireApiClient}            from "@/app/api/spire-api-client";
import {Zones}                     from "@/app/zones";
import {TRADESKILLS}               from "@/app/constants/eq-tradeskill-constants";

export default {
  name: "EqItemCardPreview",
  components: { EqCashDisplay, EqWindow, EqSpellPreview, EqDebug },
  data() {
    return {
      hasAugs: false,
      hasEffects: false,
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
        "Fire Resist": { stat: "fr", heroic: "heroic_fr" },
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
        "Mana": this.itemData.mana,
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
      ],

      // constants
      DB_CLASSES: DB_CLASSES,
      DB_RACE_NAMES: DB_RACE_NAMES,
      DB_DIETIES_FULL: DB_DIETIES_FULL,
      TRADESKILLS: TRADESKILLS,

      // related data
      droppedBy: [],
      unlocksDoors: [],
      fishedIn: [],
      foragedIn: [],
      groundSpawns: [],
      taskRewards: [],
      merchants: [],
      tradeskillResult: [],
      startingItems: []
      // augslots: {}
    }
  },
  methods: {

    commify(x) {
      return x.toString().replace(/\B(?=(\d{3})+(?!\d))/g, ",");
    },

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

      for (let i = 1; i <= 5; i++) {
        if (this.itemData['augslot_' + i + '_type']) {
          this.hasAugs = true
        }
      }

      for (let e of this.effects) {
        if (this.itemData[e]) {
          this.hasEffects = true;
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

      if (this.showRelatedData) {
        this.renderRelatedData()
      }
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
    },
    async renderRelatedData() {

      const d = this.itemData

      // doors
      let unlocksDoors = []
      if (d.doors) {
        for (const e of d.doors) {
          unlocksDoors.push(
            {
              name: e.name,
              zone: await Zones.getZoneLongNameByShortName(e.zone),
              x: e.pos_x,
              y: e.pos_y,
              z: e.pos_z,
            }
          )
        }
      }
      this.unlocksDoors = unlocksDoors

      // loot
      let droppedBy = []
      if (d.lootdrop_entries) {
        let droppedByNpc = []
        for (const lootdropEntry of d.lootdrop_entries) {
          if (lootdropEntry.lootdrop && lootdropEntry.lootdrop.loottable_entries) {
            for (const loottableEntry of lootdropEntry.lootdrop.loottable_entries) {
              if (loottableEntry.loottable && loottableEntry.loottable.npc_types) {
                for (const npcType of loottableEntry.loottable.npc_types) {
                  const npcName = npcType.name.replaceAll("_", " ").replaceAll("#", "").trim()

                  if (
                    npcType.spawnentries &&
                    npcType.spawnentries[0].spawngroup &&
                    npcType.spawnentries[0].spawngroup.spawn_2
                  ) {

                    if (!droppedByNpc.includes(npcName)) {
                      const zoneName = await Zones.getZoneLongNameByShortName(npcType.spawnentries[0].spawngroup.spawn_2.zone.toLowerCase())

                      droppedBy.push(
                        {
                          name: npcName,
                          zone: zoneName,
                        }
                      )

                      droppedByNpc.push(npcName)
                    }
                  }
                }
              }
            }
          }
        }
      }

      this.droppedBy = droppedBy.sort((a, b) => (a.name > b.name) ? 1 : -1)

      // fishing
      let fishedIn = []
      if (d.fishings) {
        for (const e of d.fishings) {
          fishedIn.push(
            {
              zone: e.zone,
              skill: e.skill_level,
            }
          )
        }
      }
      this.fishedIn = fishedIn

      // forage
      let foragedIn = []
      if (d.forages) {
        for (const e of d.forages) {
          foragedIn.push(
            {
              zone: e.zone,
              skill: e.level,
              chance: e.chance,
            }
          )
        }
      }
      this.foragedIn = foragedIn

      // starting items
      let startingItems = []
      if (d.starting_items) {
        for (const e of d.starting_items) {
          startingItems.push(
            {
              zone: (e.zoneid > 0 ? e.zone : 0),
              class: (e.class > 0 && DB_CLASSES[e.class] ? DB_CLASSES[e.class] : 0),
              race: (e.race > 0 && DB_RACE_NAMES[e.race] ? DB_RACE_NAMES[e.race] : 0),
              deity: (e.deityid > 0 && DB_DIETIES_FULL[e.deityid] ? DB_DIETIES_FULL[e.deityid].name : 0),
              data: e,
            }
          )
        }
      }
      this.startingItems = startingItems

      // ground spawns
      let groundSpawns = []
      if (d.ground_spawns) {
        for (const e of d.ground_spawns) {
          groundSpawns.push(
            {
              zone: (e.zoneid > 0 ? e.zone : 0),
              data: e,
            }
          )
        }
      }
      this.groundSpawns = groundSpawns

      // tasks
      let taskRewards = []
      if (d.tasks) {
        for (const e of d.tasks) {
          taskRewards.push(
            {
              data: e,
            }
          )
        }
      }
      this.taskRewards = taskRewards

      // merchants
      let merchants = []
      if (d.merchantlists) {
        for (const e of d.merchantlists) {
          let merchantZone = ""
          if (
            e.npc_type
            && e.npc_type.spawnentries
            && e.npc_type.spawnentries[0]
            && e.npc_type.spawnentries[0].spawngroup
            && e.npc_type.spawnentries[0].spawngroup.spawn_2
          ) {
            merchantZone = await Zones.getZoneLongNameByShortName(e.npc_type.spawnentries[0].spawngroup.spawn_2.zone)
          }

          const npcName = (e.npc_type && e.npc_type.name ? e.npc_type.name : '').replaceAll("_", " ").replaceAll("#", "").trim()

          if (npcName !== '' && merchantZone !== '') {
            merchants.push(
              {
                merchantName: npcName,
                merchantZone: merchantZone,
                data: e,
              }
            )
          }
        }
      }
      this.merchants = merchants

      // tradeskill result
      let tradeskillResult = []
      if (d.tradeskill_recipe_entries) {
        for (const e of d.tradeskill_recipe_entries) {

          if (
            this.itemData.id === e.item_id
            && e.componentcount === 0
            && e.successcount === 1
          ) {
            tradeskillResult.push(
              {
                name: e.tradeskill_recipe ? e.tradeskill_recipe.name : '',
                trivial: e.tradeskill_recipe ? e.tradeskill_recipe.trivial : 0,
                tradeskill: e.tradeskill_recipe ? TRADESKILLS[e.tradeskill_recipe.tradeskill] : 0,
                id: e.id,
                data: e,
              }
            )

          }

        }
      }
      this.tradeskillResult = tradeskillResult
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
    },
    showRelatedData: {
      type: Boolean,
      default: false,
      required: false,
    }
  }
}
</script>

<style>
.stat-section {
  padding: 0px;
  padding-left: 5px;
  padding-right: 5px;
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
