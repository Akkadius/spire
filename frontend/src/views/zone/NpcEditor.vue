<template>
  <content-area>
    <div class="row">
      <div class="col-7">
        <eq-window>

          <app-loader :is-loading="!npc" class="mt-3 mb-3"/>

          <eq-tabs
            v-if="npc"
            id="npc-edit-card"
            class="npc-edit-card minified-inputs"
          >
            <eq-tab
              :name="tab.name"
              :selected="(index === 0)"
              :key="tab.name"
              v-for="(tab, index) in tabs"
            >
              <div class="row">
                <div
                  v-for="field in tab.fields"
                  :key="field.field"
                  :class="field.col + ' mb-3 pl-2 pr-2'"
                >
                  <div class="text-center" v-if="field.fType !== 'checkbox'">
                    <span
                      v-if="field.itemIcon"
                      :class="'item-' + field.itemIcon + '-sm'"
                      style="display: inline-block"
                    />
                    {{ field.desc }}
                  </div>

                  <!-- checkbox -->
                  <div :class="'text-right ' + (field.inline ? 'mt-4' : '')" v-if="field.fType === 'checkbox'">
                    <div class="d-inline-block" style="bottom: 2px; position: relative; margin-right: 1px">
                      {{field.desc}}
                    </div>
                    <eq-checkbox
                      v-b-tooltip.hover.v-dark.right :title="getFielddesc(field.field)"
                      class="d-inline-block text-center"
                      :true-value="(typeof field.true !== 'undefined' ? field.true : 1)"
                      :false-value="(typeof field.false !== 'undefined' ? field.false : 0)"
                      v-model.number="npc[field.field]"
                      @input="npc[field.field] = $event"

                    />
                  </div>

                  <!-- input number -->
                  <b-form-input
                    v-if="field.fType === 'number'"
                    :id="field.field"
                    v-model.number="npc[field.field]"
                    class="m-0 mt-1"
                    v-on="typeof field.onclick !== 'undefined' ? { click: () => field.onclick(field.field) } : {}"
                    v-b-tooltip.hover.v-dark.right :title="getFielddesc(field.field)"
                    :style="(npc[field.field] === 0 ? 'opacity: .5' : '')"
                  />

                  <!-- input text -->
                  <b-form-input
                    v-if="field.fType === 'text'"
                    :id="field.field"
                    v-model.number="npc[field.field]"
                    class="m-0 mt-1"
                    v-on="typeof field.onclick !== 'undefined' ? { click: () => field.onclick(field.field) } : {}"
                    v-b-tooltip.hover.v-dark.right :title="getFielddesc(field.field)"
                    :style="(npc[field.field] <= (typeof field.zeroValue !== 'undefined' ? field.zeroValue : 0) ? 'opacity: .5' : '')"
                  />

                  <!-- textarea -->
                  <b-textarea
                    v-if="field.fType === 'textarea'"
                    :id="field.field"
                    v-model="npc[field.field]"
                    class="m-0 mt-1"
                    rows="2"
                    max-rows="6"
                    v-on="typeof field.onclick !== 'undefined' ? { click: () => field.onclick(field.field) } : {}"
                    v-b-tooltip.hover.v-dark.right :title="getFielddesc(field.field)"
                    :style="(npc[field.field] === '' ? 'opacity: .5' : '') + ';'"
                  ></b-textarea>

                  <!-- select -->
                  <select
                    v-model.number="npc[field.field]"
                    class="form-control m-0 mt-1"
                    v-if="field.selectData"
                    v-b-tooltip.hover.v-dark.right :title="getFielddesc(field.field)"
                    :style="(npc[field.field] <= (typeof field.zeroValue !== 'undefined' ? field.zeroValue : 0) ? 'opacity: .5' : '')"
                  >
                    <option
                      v-for="(desc, index) in field.selectData"
                      :key="index"
                      :value="parseInt(index)"
                    >
                      {{ index }}) {{ desc }}
                    </option>
                  </select>
                </div>
              </div>

            </eq-tab>

          </eq-tabs>
        </eq-window>
      </div>

      <!-- Preview / Selector Pane -->
      <div class="col-5">

        <eq-window v-if="selectorActive['special_abilities']">
          <npc-special-abilities
            :abilities="npc.special_abilities"
            :inputData.sync="npc.special_abilities"
          />
        </eq-window>

        <eq-window v-if="npc && Object.keys(selectorActive).length === 0">
          <eq-npc-card-preview :npc="npc"/>
        </eq-window>

      </div>
    </div>
  </content-area>
</template>

<script>
import EqWindowFancy           from "../../components/eq-ui/EQWindowFancy";
import EqWindow                from "../../components/eq-ui/EQWindow";
import EqTabs                  from "../../components/eq-ui/EQTabs";
import EqTab                   from "../../components/eq-ui/EQTab";
import EqItemPreview           from "../../components/preview/EQItemCardPreview";
import EqCheckbox              from "../../components/eq-ui/EQCheckbox";
import ClassBitmaskCalculator  from "../../components/tools/ClassBitmaskCalculator";
import RaceBitmaskCalculator   from "../../components/tools/RaceBitmaskCalculator";
import DeityBitmaskCalculator  from "../../components/tools/DeityCalculator";
import InventorySlotCalculator from "../../components/tools/InventorySlotCalculator";
import AugBitmaskCalculator from "../../components/tools/AugmentTypeCalculator";
import EqWindowSimple       from "../../components/eq-ui/EQWindowSimple";
import LoaderCastBarTimer   from "../../components/LoaderCastBarTimer";
import ContentArea          from "../../components/layout/ContentArea";
import {Npcs}               from "../../app/npcs";
import EqDebug              from "../../components/eq-ui/EQDebug";
import EqNpcCardPreview     from "../../components/preview/EQNpcCardPreview";
import {DB_CLASSES}         from "../../app/constants/eq-classes-constants";
import {DB_RACE_NAMES}      from "../../app/constants/eq-races-constants";
import {BODYTYPES}          from "../../app/constants/eq-bodytype-constants";
import {EditFormFieldUtil}  from "../../app/forms/edit-form-field-util";
import NpcSpecialAbilities  from "../../components/tools/NpcSpecialAbilities";

const MILLISECONDS_BEFORE_WINDOW_RESET = 5000;

export default {
  name: "ItemEdit",
  components: {
    NpcSpecialAbilities,
    EqNpcCardPreview,
    EqDebug,
    ContentArea,
    LoaderCastBarTimer,
    EqWindowSimple,
    AugBitmaskCalculator,
    InventorySlotCalculator,
    DeityBitmaskCalculator,
    RaceBitmaskCalculator,
    ClassBitmaskCalculator,
    EqCheckbox,
    EqItemPreview,
    EqTab,
    EqTabs,
    EqWindow,
    EqWindowFancy
  },
  data() {
    return {
      npc: null,
      originalItem: {}, // item record data; used to reference original values in tools

      // selectors
      selectorActive: {},

      // state, loaded or not
      loaded: true,

      tabs: this.getTabs()
    }
  },
  watch: {

    // reset state vars when we navigate away
    '$route'() {
      // this.item         = null;
      // this.originalItem = {};

      // reset state vars when we navigate away
      // this.notification = ""
      // EditFormFieldUtil.resetFieldEditedStatus()
      // this.resetPreviewComponents()

      // reload
      // this.load()
    },

  },
  async created() {
    this.load()
  },
  methods: {

    /**
     * Selectors
     */
    resetPreviewComponents() {
      for (const [k, v] of Object.entries(this.selectorActive)) {
        this.selectorActive[k] = false
      }

      EditFormFieldUtil.resetFieldSubEditorHighlightedStatus()
    },
    setSelectorActive(selector) {
      this.resetPreviewComponents()
      this.lastResetTime            = Date.now()
      this.selectorActive[selector] = true
      this.$forceUpdate()

      EditFormFieldUtil.setFieldSubEditorHighlightedById(selector)
    },

    /**
     * Load
     */
    async load() {
      this.npc = await Npcs.getNpc(this.$route.params.npc)

      this.loaded = true

      setTimeout(() => {
        let hasSubEditorFields = [
          "npc_spells_id",
          "npc_spells_effects_id",
          "merchant_id",
          "emoteid",
          "special_abilities",
          "loottable_id",
          "alt_currency_id",
          "ammo_idfile",
          "npc_faction_id"
        ];
        hasSubEditorFields.forEach((field) => {
          EditFormFieldUtil.setFieldHighlightHasSubEditor(field)
        })
      }, 1)

    },

    getFielddesc() {
      return ""
    },

    getTabs() {
      return [
        {
          name: 'General',
          fields: [
            { desc: 'ID', field: 'id', fType: 'text', itemIcon: '6840', col: 'col-2', },
            { desc: 'Name', field: 'name', fType: 'text', itemIcon: '6840', col: 'col-2', },
            { desc: 'Last Name', field: 'lastname', fType: 'text', itemIcon: '6840', col: 'col-2', },
            { desc: 'Level', field: 'level', fType: 'text', itemIcon: '6840', col: 'col-2', },
            { desc: 'Class', field: 'class', selectData: DB_CLASSES, col: 'col-2', },
            { desc: 'Race', field: 'race', selectData: DB_RACE_NAMES, col: 'col-2', },

            { desc: 'Bodytype', field: 'bodytype', selectData: BODYTYPES, col: 'col-2', },


            // {desc: "id", field: "id", fType: "text", col: 'col-2' },
            // {desc: "name", field: "name", fType: "text", col: 'col-2' },
            // {desc: "lastname", field: "lastname", fType: "text", col: 'col-2' },
            // {desc: "level", field: "level", fType: "text", col: 'col-2' },
            // {desc: "race", field: "race", fType: "text", col: 'col-2' },
            // {desc: "class", field: "class", fType: "text", col: 'col-2' },
            // {desc: "bodytype", field: "bodytype", fType: "text", col: 'col-2' },


            { desc: "loottable_id", field: "loottable_id", fType: "text", col: 'col-2' },
            { desc: "merchant_id", field: "merchant_id", fType: "text", col: 'col-2' },
            { desc: "alt_currency_id", field: "alt_currency_id", fType: "text", col: 'col-2' },
            { desc: "npc_spells_id", field: "npc_spells_id", fType: "text", col: 'col-2' },
            { desc: "npc_spells_effects_id", field: "npc_spells_effects_id", fType: "text", col: 'col-2' },
            { desc: "npc_faction_id", field: "npc_faction_id", fType: "text", col: 'col-2' },
            { desc: "adventure_template_id", field: "adventure_template_id", fType: "text", col: 'col-2' },
            { desc: "trap_template", field: "trap_template", fType: "text", col: 'col-2' },


            // { desc: "see_invis", field: "see_invis", fType: "text", col: 'col-2' },
            // { desc: "see_invis_undead", field: "see_invis_undead", fType: "text", col: 'col-2' },
            // { desc: "qglobal", field: "qglobal", fType: "text", col: 'col-2' },

            // possible checkboxes
            { desc: "spawn_limit", field: "spawn_limit", fType: "text", col: 'col-2' },

            // { desc: "findable", field: "findable", fType: "text", col: 'col-2' },

            // { desc: "see_hide", field: "see_hide", fType: "text", col: 'col-2' },
            // { desc: "see_improved_hide", field: "see_improved_hide", fType: "text", col: 'col-2' },
            // { desc: "trackable", field: "trackable", fType: "text", col: 'col-2' },
            { desc: "isbot", field: "isbot", fType: "text", col: 'col-2' },
            { desc: "exclude", field: "exclude", fType: "text", col: 'col-2' },

            // version is not really used outside of the PEQ editor
            // { desc: "version", field: "version", fType: "text", col: 'col-2' },

            { desc: "private_corpse", field: "private_corpse", fType: "text", col: 'col-2' },
            // {desc: "unique_spawn_by_name", field: "unique_spawn_by_name", fType: "text", col: 'col-2' },
            // {desc: "underwater", field: "underwater", fType: "text", col: 'col-2' },
            // {desc: "isquest", field: "isquest", fType: "text", col: 'col-2' },
            { desc: "emoteid", field: "emoteid", fType: "text", col: 'col-2' },


            { desc: "peqid", field: "peqid", fType: "text", col: 'col-2' },
            { desc: "unique_", field: "unique_", fType: "text", col: 'col-2' },
            { desc: "fixed", field: "fixed", fType: "text", col: 'col-2' },
            // {desc: "ignore_despawn", field: "ignore_despawn", fType: "text", col: 'col-2' },
            // {desc: "show_name", field: "show_name", fType: "text", col: 'col-2' },
            // {desc: "untargetable", field: "untargetable", fType: "text", col: 'col-2' },

            // { desc: "skip_global_loot", field: "skip_global_loot", fType: "text", col: 'col-2' },
            // { desc: "rare_spawn", field: "rare_spawn", fType: "text", col: 'col-2' },
            { desc: "stuck_behavior", field: "stuck_behavior", fType: "text", col: 'col-2' },
            { desc: "model", field: "model", fType: "text", col: 'col-2' },
            { desc: "flymode", field: "flymode", fType: "text", col: 'col-2' },
            // { desc: "always_aggro", field: "always_aggro", fType: "text", col: 'col-2' },
            { desc: "exp_mod", field: "exp_mod", fType: "text", col: 'col-2' },

          ],
        },
        {
          name: 'Weapon',
          fields: [
            { desc: "ammo_idfile", field: "ammo_idfile", fType: "text", col: 'col-2' },
            { desc: "prim_melee_type", field: "prim_melee_type", fType: "text", col: 'col-2' },
            { desc: "sec_melee_type", field: "sec_melee_type", fType: "text", col: 'col-2' },
            { desc: "ranged_type", field: "ranged_type", fType: "text", col: 'col-2' },
            { desc: "d_melee_texture_1", field: "d_melee_texture_1", fType: "text", col: 'col-2' },
            { desc: "d_melee_texture_2", field: "d_melee_texture_2", fType: "text", col: 'col-2' },
          ]
        },
        {
          name: 'Aggro',
          fields: [
            { desc: "aggroradius", field: "aggroradius", fType: "text", col: 'col-2' },
            { desc: "assistradius", field: "assistradius", fType: "text", col: 'col-2' },

            { desc: 'Always Aggro', field: 'always_aggro', fType: 'checkbox', col: 'col-2', inline: true },
            { desc: "NPC Aggro", field: "npc_aggro", fType: "checkbox", col: 'col-2', inline: true },
          ]
        },
        {
          name: 'Appearance',
          fields: [
            { desc: "gender", field: "gender", fType: "text", col: 'col-2' },
            { desc: "texture", field: "texture", fType: "text", col: 'col-2' },
            { desc: "helmtexture", field: "helmtexture", fType: "text", col: 'col-2' },
            { desc: "herosforgemodel", field: "herosforgemodel", fType: "text", col: 'col-2' },
            { desc: "size", field: "size", fType: "text", col: 'col-2' },

            { desc: "face", field: "face", fType: "text", col: 'col-2' },
            { desc: "luclin_hairstyle", field: "luclin_hairstyle", fType: "text", col: 'col-2' },
            { desc: "luclin_haircolor", field: "luclin_haircolor", fType: "text", col: 'col-2' },
            { desc: "luclin_eyecolor", field: "luclin_eyecolor", fType: "text", col: 'col-2' },
            { desc: "luclin_eyecolor_2", field: "luclin_eyecolor_2", fType: "text", col: 'col-2' },
            { desc: "luclin_beardcolor", field: "luclin_beardcolor", fType: "text", col: 'col-2' },
            { desc: "luclin_beard", field: "luclin_beard", fType: "text", col: 'col-2' },
            { desc: "drakkin_heritage", field: "drakkin_heritage", fType: "text", col: 'col-2' },
            { desc: "drakkin_tattoo", field: "drakkin_tattoo", fType: "text", col: 'col-2' },
            { desc: "drakkin_details", field: "drakkin_details", fType: "text", col: 'col-2' },
            { desc: "armortint_id", field: "armortint_id", fType: "text", col: 'col-2' },
            { desc: "armortint_red", field: "armortint_red", fType: "text", col: 'col-2' },
            { desc: "armortint_green", field: "armortint_green", fType: "text", col: 'col-2' },
            { desc: "armortint_blue", field: "armortint_blue", fType: "text", col: 'col-2' },

            { desc: "armtexture", field: "armtexture", fType: "text", col: 'col-2' },
            { desc: "bracertexture", field: "bracertexture", fType: "text", col: 'col-2' },
            { desc: "handtexture", field: "handtexture", fType: "text", col: 'col-2' },
            { desc: "legtexture", field: "legtexture", fType: "text", col: 'col-2' },
            { desc: "feettexture", field: "feettexture", fType: "text", col: 'col-2' },
            { desc: "light", field: "light", fType: "text", col: 'col-2' },
          ]
        },
        {
          name: 'Stats',
          fields: [
            { desc: "mindmg", field: "mindmg", fType: "text", col: 'col-2' },
            { desc: "maxdmg", field: "maxdmg", fType: "text", col: 'col-2' },
            { desc: "attack_count", field: "attack_count", fType: "text", col: 'col-2' },
            // deprecated
            // { desc: "npcspecialattks", field: "npcspecialattks", fType: "text", col: 'col-2' },

            { desc: "ac", field: "ac", fType: "text", col: 'col-2' },
            { desc: "hp", field: "hp", fType: "text", col: 'col-2' },
            { desc: "mana", field: "mana", fType: "text", col: 'col-2' },
            { desc: "hp_regen_rate", field: "hp_regen_rate", fType: "text", col: 'col-2' },
            { desc: "hp_regen_per_second", field: "hp_regen_per_second", fType: "text", col: 'col-2' },
            { desc: "mana_regen_rate", field: "mana_regen_rate", fType: "text", col: 'col-2' },

            { desc: "attack_speed", field: "attack_speed", fType: "text", col: 'col-2' },
            { desc: "attack_delay", field: "attack_delay", fType: "text", col: 'col-2' },
            { desc: "atk", field: "atk", fType: "text", col: 'col-2' },
            { desc: "accuracy", field: "accuracy", fType: "text", col: 'col-2' },
            { desc: "avoidance", field: "avoidance", fType: "text", col: 'col-2' },
            { desc: "slow_mitigation", field: "slow_mitigation", fType: "text", col: 'col-2' },

            { desc: "runspeed", field: "runspeed", fType: "text", col: 'col-2' },
            { desc: "walkspeed", field: "walkspeed", fType: "text", col: 'col-2' },

            { desc: "str", field: "str", fType: "text", col: 'col-2' },
            { desc: "sta", field: "sta", fType: "text", col: 'col-2' },
            { desc: "dex", field: "dex", fType: "text", col: 'col-2' },
            { desc: "agi", field: "agi", fType: "text", col: 'col-2' },
            { desc: "_int", field: "_int", fType: "text", col: 'col-2' },
            { desc: "wis", field: "wis", fType: "text", col: 'col-2' },
            { desc: "cha", field: "cha", fType: "text", col: 'col-2' },

            { desc: "mr", field: "mr", fType: "text", col: 'col-2' },
            { desc: "cr", field: "cr", fType: "text", col: 'col-2' },
            { desc: "dr", field: "dr", fType: "text", col: 'col-2' },
            { desc: "fr", field: "fr", fType: "text", col: 'col-2' },
            { desc: "pr", field: "pr", fType: "text", col: 'col-2' },
            { desc: "corrup", field: "corrup", fType: "text", col: 'col-2' },
            { desc: "ph_r", field: "ph_r", fType: "text", col: 'col-2' },

            { desc: "spellscale", field: "spellscale", fType: "text", col: 'col-2' },
            { desc: "healscale", field: "healscale", fType: "text", col: 'col-2' },

            { desc: "scalerate", field: "scalerate", fType: "text", col: 'col-2' },
            { desc: "maxlevel", field: "maxlevel", fType: "text", col: 'col-2' },

            { desc: "special_abilities", field: "special_abilities", fType: "text", col: 'col-12', onclick: this.setSelectorActive, },


          ]
        },
        {
          name: 'Charm Stats',
          fields: [
            { desc: "charm_ac", field: "charm_ac", fType: "text", col: 'col-2' },
            { desc: "charm_min_dmg", field: "charm_min_dmg", fType: "text", col: 'col-2' },
            { desc: "charm_max_dmg", field: "charm_max_dmg", fType: "text", col: 'col-2' },
            { desc: "charm_attack_delay", field: "charm_attack_delay", fType: "text", col: 'col-2' },
            { desc: "charm_accuracy_rating", field: "charm_accuracy_rating", fType: "text", col: 'col-2' },
            { desc: "charm_avoidance_rating", field: "charm_avoidance_rating", fType: "text", col: 'col-2' },
            { desc: "charm_atk", field: "charm_atk", fType: "text", col: 'col-2' },
          ]
        },
        {
          name: 'Settings',
          fields: [
            // checkboxes
            { desc: 'See Hide', field: 'see_hide', fType: 'checkbox', col: 'col-2', },
            { desc: 'See Improved Hide', field: 'see_improved_hide', fType: 'checkbox', col: 'col-2', },
            { desc: 'See Invisible', field: 'see_invis', fType: 'checkbox', col: 'col-2', },
            { desc: 'See Invis Undead', field: 'see_invis_undead', fType: 'checkbox', col: 'col-2', },
            { desc: 'Show Name', field: 'show_name', fType: 'checkbox', col: 'col-2', },
            { desc: 'Trackable', field: 'trackable', fType: 'checkbox', col: 'col-2', },
            { desc: 'Skip Global Loot', field: 'skip_global_loot', fType: 'checkbox', col: 'col-2', },
            { desc: 'No Target Hotkey', field: 'no_target_hotkey', fType: 'checkbox', col: 'col-2', },
            { desc: 'Findable', field: 'findable', fType: 'checkbox', col: 'col-2', },
            { desc: 'Untargetable', field: 'untargetable', fType: 'checkbox', col: 'col-2', },
            { desc: 'Underwater', field: 'underwater', fType: 'checkbox', col: 'col-2', },
            { desc: 'QGlobal', field: 'qglobal', fType: 'checkbox', col: 'col-2', },
            { desc: 'Ignore Despawn', field: 'ignore_despawn', fType: 'checkbox', col: 'col-2', },
            { desc: 'Quest NPC', field: 'isquest', fType: 'checkbox', col: 'col-2', },
            { desc: 'Unique Spawn', field: 'unique_spawn_by_name', fType: 'checkbox', col: 'col-2', },
            { desc: 'Rare Spawn', field: 'rare_spawn', fType: 'checkbox', col: 'col-2', },
            { desc: 'Always Aggro', field: 'always_aggro', fType: 'checkbox', col: 'col-2', },
            { desc: "NPC Aggro", field: "npc_aggro", fType: "checkbox", col: 'col-2' },
            { desc: "Raid Target", field: "raid_target", fType: "checkbox", col: 'col-2' },

          ]
        },
      ]
    }
  }
}
</script>

<style scoped>

.effect-tab input, .effect-tab select {
  margin-bottom: 0;
}
</style>
