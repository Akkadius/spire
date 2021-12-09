<template>
  <div class="container-fluid">
    <div class="panel-body">
      <div class="panel panel-default">
        <div class="row">
          <div class="col-7">
            <eq-window style="margin-top: 30px" title="Edit Item">

              <div v-if="notification">
                <b-button class="btn-dark btn-outline-warning form-control" @click="notification = ''">
                  <i class="ra ra-book mr-1"></i>{{ notification }}
                </b-button>
              </div>

              <b-alert show dismissable variant="danger" v-if="error">
                <i class="fa fa-warning"></i> {{ error }}
              </b-alert>

              <eq-tabs
                v-if="item"
                id="item-edit-card"
                class="item-edit-card"
                :hover-open="true"
                @mouseover.native="previewItem"
              >
                <eq-tab
                  name="General"
                  :selected="true"
                >
                  <div class="row">
                    <div class="col-2" @mouseover="drawFreeIdSelector">
                      Id
                      <b-form-input v-model.number="item.id"/>
                    </div>
                    <div class="col-7">
                      Name
                      <b-form-input
                        :value="item.name" @change="v => item.name = v"
                      />
                    </div>

                    <div class="col-2" @mouseover="drawIconSelector">
                      Icon
                      <b-form-input v-model.number="item.icon"/>
                    </div>

                    <div
                      class="col-1" v-if="item.icon > 0"
                      style="margin-top: 18px"
                      @mouseover="drawIconSelector"
                    >
                      <span
                        :class="'fade-in item-' + item.icon"
                        style="border: 1px solid rgb(218 218 218 / 30%); border-radius: 7px;"/>
                    </div>

                  </div>

                  <div class="row">
                    <div class="col-8">

                      <div class="row">


                        <!-- Lore -->
                        <div class="col-10">
                          Lore
                          <b-form-input
                            :value="item.lore" @change="v => item.lore = v"
                          />
                        </div>

                        <!-- Lore Group-->
                        <div class="col-2">
                          Lore Group
                          <b-form-input v-model.number="item.loregroup"/>
                        </div>
                      </div>

                      <div class="row">

                        <!-- Item Type -->
                        <div class="col-3">
                          Item Type
                          <select v-model.number="item['itemtype']" class="form-control">
                            <option
                              v-for="(description, index) in DB_ITEM_TYPES"
                              :key="index"
                              :value="parseInt(index)">
                              {{ index }}) {{ description }}
                            </option>
                          </select>
                        </div>

                        <!-- Item Class -->
                        <div class="col-4">
                          Item Class
                          <select v-model.number="item['itemclass']" class="form-control">
                            <option
                              v-for="(description, index) in DB_ITEM_CLASS"
                              :key="index"
                              :value="parseInt(index)">
                              {{ index }}) {{ description }}
                            </option>
                          </select>
                        </div>

                        <!-- Material -->
                        <div class="col-3">
                          Material
                          <select v-model.number="item['material']" class="form-control">
                            <option
                              v-for="(description, index) in DB_ITEM_MATERIAL"
                              :key="index"
                              :value="parseInt(index)">
                              {{ index }}) {{ description }}
                            </option>
                          </select>
                        </div>
                      </div>

                      <div class="row">
                        <!-- Is Magic -->
                        <div class="col-2 text-center">
                          Is Magic
                          <eq-checkbox
                            class="mt-3 mb-2"
                            v-model.number="item.magic"
                            @input="item.magic = $event"
                          />
                        </div>

                        <!-- No Drop -->
                        <div class="col-2 text-center">
                          No Drop
                          <eq-checkbox
                            class="mt-3 mb-2"
                            :true-value="0"
                            :false-value="1"
                            v-model.number="item.nodrop"
                            @input="item.nodrop = $event"
                          />
                        </div>

                        <!-- No Rent -->
                        <div class="col-2 text-center">
                          No Rent
                          <eq-checkbox
                            class="mt-3 mb-2"
                            :true-value="0"
                            :false-value="1"
                            v-model.number="item.norent"
                            @input="item.norent = $event"
                          />
                        </div>

                        <!-- Tradeskills -->
                        <div class="col-3 text-center">
                          Tradeskill Item
                          <eq-checkbox
                            class="mt-3 mb-2"
                            v-model.number="item.tradeskills"
                            @input="item.tradeskills = $event"
                          />
                        </div>

                      </div>

                    </div>

                    <!-- Model Preview -->
                    <div
                      class="col-4"
                      style="text-align: center"
                      @mouseover="drawItemModelSelector"
                    >

                      <item-model-preview :id="item.idfile"/>

                      Item Model
                      <b-form-input v-model.number="item.idfile"/>
                    </div>
                  </div>

                  <div class="mt-3 mb-3">
                    <class-bitmask-calculator
                      class="text-center mt-3"
                      :imageSize="43"
                      :centered-buttons="true"
                      @input="item.classes = parseInt($event)"
                      :mask="item.classes"
                    />

                    <race-bitmask-calculator
                      :imageSize="40"
                      class="mt-3"
                      :centered-buttons="true"
                      @input="item.races = parseInt($event)"
                      :mask="item.races"
                    />

                    <deity-bitmask-calculator
                      class="mt-3"
                      :imageSize="37"
                      :show-names="true"
                      :centered-buttons="true"
                      @input="item.deity = parseInt($event)"
                      :mask="item.deity"
                    />
                  </div>

                  <div class="mt-3 mb-3">
                    <inventory-slot-calculator
                      class="mt-1"
                      :imageSize="45"
                      :centered-buttons="false"
                      @input="item.slots = parseInt($event)"
                      :mask="item.slots"
                    />
                  </div>
                </eq-tab>

                <eq-tab name="Damage">

                  <div class="row">
                    <div class="col-6">
                      <h6 class="eq-header text-center mb-3 mt-3">Damage / Delay / Haste</h6>

                      <div class="row"
                           :key="field.field"
                           v-for="field in
                       [
                         {
                           description: 'Damage',
                           field: 'damage'
                         },
                         {
                           description: 'Delay',
                           field: 'delay'
                         },
                         {
                           description: 'Haste',
                           field: 'haste'
                         },
                       ]">
                        <div class="col-5 text-right">
                          {{ field.description }}
                        </div>
                        <div class="col-4">
                          <b-form-input v-model.number="item[field.field]"/>
                        </div>
                      </div>

                      <h6 class="eq-header text-center mb-3 mt-3">Extra Damage</h6>

                      <div class="row"
                           :key="field.field"
                           v-for="field in
                       [
                         {
                           description: 'Extra Damage Skill',
                           field: 'extradmgskill'
                         },
                         {
                           description: 'Extra Damage Amount',
                           field: 'extradmgamt'
                         },
                       ]">
                        <div class="col-5 text-right">
                          {{ field.description }}
                        </div>
                        <div class="col-4">
                          <b-form-input v-model.number="item[field.field]"/>
                        </div>
                      </div>

                      <h6 class="eq-header text-center mb-3 mt-3">Weapon</h6>

                      <div class="row"
                           :key="field.field"
                           v-for="field in
                       [
                         {
                           description: 'Backstab Damage',
                           field: 'backstabdmg'
                         },
                         {
                           description: 'Range',
                           field: 'range'
                         },
                         {
                           description: 'Spell Damage',
                           field: 'spelldmg'
                         },
                       ]">
                        <div class="col-5 text-right">
                          {{ field.description }}
                        </div>
                        <div class="col-4">
                          <b-form-input v-model.number="item[field.field]"/>
                        </div>
                      </div>
                    </div>
                    <div class="col-6">
                      <h6 class="eq-header text-center mb-3 mt-3">Bane</h6>

                      <div class="row"
                           :key="field.field"
                           v-for="field in
                       [
                         {
                           description: 'Bane Damage Amount',
                           field: 'banedmgamt'
                         },
                         {
                           description: 'Bane Damage Body',
                           field: 'banedmgbody'
                         },
                         {
                           description: 'Bane Damage Race',
                           field: 'banedmgrace'
                         },
                         {
                           description: 'Bane Damage Race Amount',
                           field: 'banedmgraceamt'
                         },
                       ]">
                        <div class="col-5 text-right">
                          {{ field.description }}
                        </div>
                        <div class="col-4">
                          <b-form-input v-model.number="item[field.field]"/>
                        </div>
                      </div>

                      <h6 class="eq-header text-center mb-3 mt-3">Elemental</h6>

                      <div class="row"
                           :key="field.field"
                           v-for="field in
                       [
                         {
                           description: 'Elemental Damage Amount',
                           field: 'elemdmgamt'
                         },
                         {
                           description: 'Element Damage Type',
                           field: 'elemdmgtype'
                         },
                       ]">
                        <div class="col-5 text-right">
                          {{ field.description }}
                        </div>
                        <div class="col-4">
                          <b-form-input v-model.number="item[field.field]"/>
                        </div>
                      </div>
                    </div>
                  </div>
                </eq-tab>

                <eq-tab
                  name="Augmentation">

                  <h6 class="eq-header text-center mt-3 mb-3">Item Is Augment</h6>

                  <!-- Aug Type -->
                  <div class="row">
                    <div class="col-1">

                    </div>
                    <div class="col-4 text-right">
                      Augment Type
                    </div>
                    <div class="col-3">
                      <select v-model.number="item['augtype']" class="form-control">
                        <option
                          v-for="(value, index) in AUG_TYPES"
                          :key="index"
                          :value="parseInt(index)">
                          {{ index }}) {{ value.name }}
                        </option>
                      </select>

                    </div>
                  </div>

                  <div class="row">
                    <div class="col-1">

                    </div>
                    <div class="col-4 text-right">
                      Augment Restriction
                    </div>
                    <div class="col-3">
                      <select v-model.number="item['augrestrict']" class="form-control">
                        <option
                          v-for="(value, index) in DB_ITEM_AUG_RESTRICT"
                          :key="index"
                          :value="parseInt(index)">
                          {{ index }}) {{ value }}
                        </option>
                      </select>
                    </div>
                  </div>

                  <h6 class="eq-header text-center mt-3 mb-3">Item Has Augments</h6>

                  <!-- Aug Distiller Type -->
                  <div class="row">
                    <div class="col-1">

                    </div>
                    <div class="col-4 text-right">
                      Augment Distiller Type
                    </div>
                    <div class="col-3">
                      <b-form-input v-model.number="item['augdistiller']"/>
                    </div>
                  </div>

                  <!-- Aug Type -->
                  <div class="row" v-for="i in 5" :key="i">
                    <div class="col-1">

                    </div>
                    <div class="col-4 text-right">
                      Augment Slot {{ i }} Type
                    </div>
                    <div class="col-3">
                      <select v-model.number="item['augslot_' + i + '_type']" class="form-control">
                        <option
                          v-for="(value, index) in AUG_TYPES"
                          :key="index"
                          :value="parseInt(index)">
                          {{ index }}) {{ value.name }}
                        </option>
                      </select>
                    </div>
                  </div>
                </eq-tab>

                <eq-tab
                  name="Resists & Stats"
                >

                  <div class="row">

                    <!-- Resists -->
                    <div class="col-6 text-center">
                      <h6 class="eq-header">
                        Resists
                      </h6>
                      <div
                        v-for="(resist, description) in resists"
                        :key="resist.stat"
                        class="row text-center"
                      >
                        <div class="col-2 text-right">
                          {{ description }}
                        </div>
                        <div class="col-4">
                          <b-form-input v-model.number="item[resist.stat]"/>
                        </div>
                        <div class="col-2 text-right">
                          Heroic
                        </div>
                        <div class="col-4">
                          <b-form-input v-model.number="item[resist.heroic]"/>
                        </div>
                      </div>
                    </div>

                    <!-- Stats -->
                    <div class="col-6 text-center">

                      <h6 class="eq-header">
                        Stats
                      </h6>
                      <div
                        v-for="(stat, description) in stats"
                        :key="stat.stat"
                        class="row text-center"
                      >
                        <div class="col-2 text-right">
                          {{ description }}
                        </div>
                        <div class="col-4">
                          <b-form-input v-model.number="item[stat.stat]"/>
                        </div>
                        <div class="col-2 text-right">
                          Heroic
                        </div>
                        <div class="col-4">
                          <b-form-input v-model.number="item[stat.heroic]"/>
                        </div>
                      </div>
                    </div>

                  </div>
                </eq-tab>

                <eq-tab name="Mods">
                  <div v-for="(field, description) in mod3" :key="field" class="row text-center">
                    <div class="col-1">

                    </div>
                    <div class="col-4 text-right">
                      {{ description }}
                    </div>
                    <div class="col-2">
                      <b-form-input v-model.number="item[field]" v-if="field !== 'combateffects'"/>
                      <!-- For some reason combateffects is a varchar field -->
                      <b-form-input v-model="item[field]" v-if="field === 'combateffects'"/>
                    </div>
                  </div>
                </eq-tab>

                <eq-tab name="Pricing">
                  <div v-for="(field, description) in pricingFields" :key="field" class="row text-center">
                    <div class="col-1">

                    </div>
                    <div class="col-4 text-right">
                      {{ description }}
                    </div>
                    <div class="col-2">
                      <b-form-input v-model.number="item[field]"/>
                      <!--                      <b-form-input v-model="item[field]"/>-->
                    </div>
                  </div>
                </eq-tab>

              </eq-tabs>

              <div class="text-center mt-3">
                <b-button
                  class="btn-dark btn-sm btn-outline-warning"
                  @click="saveItem"
                >
                  <i class="ra ra-book mr-1"></i>
                  Save Item
                </b-button>
              </div>
            </eq-window>
          </div>

          <div class="col-5">

            <!-- preview item -->
            <eq-window
              style="margin-top: 30px; margin-right: 10px; width: auto;"
              :key="item.updatedAt"
              v-if="previewItemActive && item && item.id > 0">
              <eq-item-preview
                :item-data="item"/>
            </eq-window>

            <!-- item model selector -->
            <eq-window
              style="margin-top: 30px; margin-right: 10px; width: auto;"
              class="fade-in"
              v-if="itemModelSelectorActive && item">

              <item-model-selector
                :selected-model="item.idfile"
                @input="item.idfile = $event"
              />
            </eq-window>

            <!-- icon selector -->
            <eq-window
              style="margin-top: 30px; margin-right: 10px; width: auto;"
              class="fade-in"
              v-if="iconSelectorActive">

              <item-icon-selector
                :selected-icon="item.icon"
                @input="item.icon = $event"
              />
            </eq-window>

            <!-- free id selector -->
            <eq-window
              title="Free Item Ids"
              style="margin-top: 30px; margin-right: 10px; width: auto;"
              class="fade-in"
              v-if="freeIdSelectorActive">

              <free-id-selector
                table-name="items"
                id-name="id"
                name-label="name"
                :with-reserved="true"
                @input="item.id = $event"/>
            </eq-window>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import EqWindowFancy           from "../../components/eq-ui/EQWindowFancy";
import EqWindow                from "../../components/eq-ui/EQWindow";
import EqTabs                  from "../../components/eq-ui/EQTabs";
import EqTab                   from "../../components/eq-ui/EQTab";
import EqItemPreview           from "../../components/eq-ui/EQItemCardPreview";
import {App}                   from "../../constants/app";
import EqCheckbox              from "../../components/eq-ui/EQCheckbox";
import {SpireApiClient}        from "../../app/api/spire-api-client";
import * as util               from "util";
import FreeIdSelector          from "../../components/tools/FreeIdSelector";
import {Items}                 from "../../app/items";
import {ItemApi}               from "../../app/api";
import ItemModelPreview        from "../../components/tools/ItemModelPreview";
import ItemModelSelector       from "../../components/tools/ItemModelSelector";
import ItemIconSelector        from "../../components/tools/ItemIconSelector";
import ClassBitmaskCalculator  from "../../components/tools/ClassBitmaskCalculator";
import RaceBitmaskCalculator   from "../../components/tools/RaceBitmaskCalculator";
import DeityBitmaskCalculator  from "../../components/tools/DeityCalculator";
import {
  DB_ITEM_AUG_RESTRICT,
  DB_ITEM_CLASS,
  DB_ITEM_MATERIAL,
  DB_ITEM_TYPES
}                              from "../../app/constants/eq-item-constants";
import {AUG_TYPES}             from "../../app/constants/eq-aug-constants";
import InventorySlotCalculator from "../../components/tools/InventorySlotCalculator";

const MILLISECONDS_BEFORE_WINDOW_RESET = 3000;

export default {
  name: "ItemEdit",
  components: {
    InventorySlotCalculator,
    DeityBitmaskCalculator,
    RaceBitmaskCalculator,
    ClassBitmaskCalculator,
    ItemIconSelector,
    ItemModelSelector,
    ItemModelPreview,
    FreeIdSelector,
    EqCheckbox,
    EqItemPreview,
    EqTab,
    EqTabs,
    EqWindow,
    EqWindowFancy
  },
  data() {
    return {
      item: null, // item record data
      spellCdnUrl: App.ASSET_SPELL_ICONS_BASE_URL,
      loaded: true,

      previewItemActive: true,
      iconSelectorActive: false,
      itemModelSelectorActive: false,
      freeIdSelectorActive: false,

      lastResetTime: Date.now(),

      notification: "",
      error: "",

      DB_ITEM_MATERIAL: DB_ITEM_MATERIAL,
      DB_ITEM_AUG_RESTRICT: DB_ITEM_AUG_RESTRICT,
      DB_ITEM_CLASS: DB_ITEM_CLASS,
      DB_ITEM_TYPES: DB_ITEM_TYPES,
      AUG_TYPES: AUG_TYPES,

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

      pricingFields: {
        "Price": "price",
        "Sell Rate": "sellrate",
        "Favor": "favor",
        "Guild Favor": "guildfavor",
        "Point Type": "pointtype",
        "LDON Price": "ldonprice",
        "LDON Theme": "ldontheme",
        "LDON Sold": "ldonsold",
        "LDON Sell Back Rate": "ldonsellbackrate",
      },
    }
  },
  watch: {
    '$route'() {
      // reset state vars when we navigate away
      this.notification = ""
      // reload
      this.load()
    },
    item: {
      handler(val, oldVal) {
        if (this.item) {
          this.item.updatedAt = Date.now()
        }
      },
      deep: true,
      immediate: true,
    },
  },
  async created() {

    setTimeout(() => {
      document.getElementById("item-edit-card").removeEventListener('input', this.setFieldModified, true);
      document.getElementById("item-edit-card").addEventListener('input', this.setFieldModified)
    }, 1000)

    this.load()
  },
  methods: {

    setFieldModified(evt) {
      // border: 2px #555555 solid !important;
      evt.target.style.setProperty('border-color', 'orange', 'important');
    },

    resetFieldEditedStatus() {
      // reset elements
      const elements = document.getElementById("item-edit-card").querySelectorAll("input, select");
      elements.forEach((element) => {
        element.style.setProperty('border-color', '#555555', 'important');
      });
    },

    async saveItem() {
      this.error        = ""
      this.notification = ""

      const api = (new ItemApi(SpireApiClient.getOpenApiConfig()))
      api.updateItem({
        id: this.item.id,
        item: this.item
      }).then((result) => {
        if (result.status === 200) {
          this.notification = util.format("Item updated successfully! (%s) %s", this.item.id, this.item.name)
          this.resetFieldEditedStatus()
        }

        if (result.data.error) {
          this.notification = result.data.error
        }

      }).catch(async (error) => {

        // marshalling error
        if (error.response.data && error.response.data.error.includes("marshal")) {
          this.error = error.response.data.error
          return
        }

        const createRes = await api.createItem({
          item: this.item
        })

        if (createRes.status === 200) {
          this.notification = util.format("Created new Item! (%s) %s", this.item.id, this.item.name)
          this.resetFieldEditedStatus()
        }
      })
    },

    load() {
      if (this.$route.params.id > 0) {
        Items.getItem(this.$route.params.id).then(result => {
          this.item      = result
          this.updatedAt = Date.now()
        })
      }
    },

    /**
     * Selector / previewers
     */
    resetPreviewComponents() {
      this.previewItemActive       = false;
      this.iconSelectorActive      = false;
      this.itemModelSelectorActive = false;
      this.freeIdSelectorActive    = false;
    },
    previewItem() {
      let shouldReset = Date.now() - this.lastResetTime > MILLISECONDS_BEFORE_WINDOW_RESET;
      // SECONDS_BEFORE_WINDOW_RESET

      if (!this.previewItemActive && shouldReset) {
        this.resetPreviewComponents()
        this.previewItemActive = true;
        this.lastResetTime     = Date.now()
      }
    },
    drawItemModelSelector() {
      this.resetPreviewComponents()
      this.itemModelSelectorActive = true
    },
    drawIconSelector() {
      if (!this.freeIdSelectorActive) {
        this.resetPreviewComponents()
        this.iconSelectorActive = true;
      }
    },
    drawFreeIdSelector() {
      this.resetPreviewComponents()
      this.lastResetTime        = Date.now()
      this.freeIdSelectorActive = true
    },
    getTargetTypeColor(targetType) {
      return Items.getTargetTypeColor(targetType);
    },
  }
}
</script>

<style scoped>
.item-edit-card input, .item-edit-card select {
  margin-bottom: 10px;
}

.effect-tab input, .effect-tab select {
  margin-bottom: 0;
}
</style>
