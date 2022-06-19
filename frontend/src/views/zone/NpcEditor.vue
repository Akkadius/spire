<template>
  <div>
    <div class="row">
      <div class="col-7">
        <eq-window>

          <app-loader :is-loading="!npc" class="mt-3 mb-3"/>

          <eq-tabs
            v-if="npc"
            id="npc-edit-card"
            class="npc-edit-card minified-inputs"
            @mouseover.native="previewMain()"
          >
            <eq-tab
              :name="tab.name"
              :selected="(index === 0)"
              :key="tab.name"
              v-for="(tab, index) in tabs"
            >
              <div class="row">
                <div class="col-12">
                  <div
                    v-for="field in tab.fields"
                    :key="field.field"
                    :class="'row'"
                  >

                    <div
                      class="col-6 text-right p-0 m-0 mr-1 mt-3"
                      style="position: relative; bottom: 6px;"
                      v-if="field.fType === 'checkbox'"
                    >
                      <span v-if="field.category" class="font-weight-bold">{{ field.category }}</span>
                      {{ field.desc }}
                    </div>
                    <div
                      class="col-6 text-right p-0 m-0 mr-3"
                      v-if="field.fType !== 'checkbox'"
                      style="margin-top: 10px !important"
                    >
                      <span v-if="field.category" class="font-weight-bold">{{ field.category }}</span>
                      {{ field.desc }}
                    </div>

                    <!--                  <div class="text-center" v-if="field.fType !== 'checkbox'">-->
                    <!--                    <span-->
                    <!--                      v-if="field.itemIcon"-->
                    <!--                      :class="'item-' + field.itemIcon + '-sm'"-->
                    <!--                      style="display: inline-block"-->
                    <!--                    />-->
                    <!--                    {{ field.desc }}-->
                    <!--                  </div>-->

                    <div class="col-3 text-left p-0 mt-2">

                      <!-- checkbox -->
                      <div :class="'text-left ml-2 mt-1'" v-if="field.fType === 'checkbox'">
                        <!--                        <div class="d-inline-block" style="bottom: 2px; position: relative; margin-right: 1px">-->
                        <!--                          {{ field.desc }}-->
                        <!--                        </div>-->
                        <eq-checkbox
                          v-b-tooltip.hover.v-dark.right :title="getFieldDescription(field.field)"
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
                        v-on="field.e ? getEventHandlers(field.e, field.field) : {}"
                        v-b-tooltip.hover.v-dark.right :title="getFieldDescription(field.field)"
                        :style="(npc[field.field] === 0 ? 'opacity: .5' : '')"
                      />

                      <!-- input text -->
                      <b-form-input
                        v-if="field.fType === 'text'"
                        :id="field.field"
                        v-model.number="npc[field.field]"
                        class="m-0 mt-1"
                        v-on="field.e ? getEventHandlers(field.e, field.field) : {}"
                        v-b-tooltip.hover.v-dark.right :title="getFieldDescription(field.field)"
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
                        v-on="field.e ? getEventHandlers(field.e, field.field) : {}"
                        v-b-tooltip.hover.v-dark.right :title="getFieldDescription(field.field)"
                        :style="(npc[field.field] === '' ? 'opacity: .5' : '') + ';'"
                      ></b-textarea>

                      <!-- select -->
                      <select
                        v-model.number="npc[field.field]"
                        :id="field.field"
                        class="form-control m-0 mt-1"
                        v-if="field.selectData"
                        v-on="field.e ? getEventHandlers(field.e, field.field) : {}"
                        v-b-tooltip.hover.v-dark.right :title="getFieldDescription(field.field)"
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
                </div>
              </div>

            </eq-tab>

          </eq-tabs>
        </eq-window>
      </div>

      <!-- Preview / Selector Pane -->
      <div class="col-5">

        <eq-window v-if="npc && !isAnySelectorActive()" style="max-height: 95vh; overflow-y: scroll; overflow-x: hidden">
          <eq-npc-card-preview :npc="npc"/>
        </eq-window>

        <eq-window v-if="selectorActive['special_abilities']">
          <npc-special-abilities
            :abilities="npc.special_abilities"
            :inputData.sync="npc.special_abilities"
          />
        </eq-window>

        <item-model-selector
          v-if="selectorActive['d_melee_texture_1']"
          :selected-model="npc.d_melee_texture_1"
          @input="npc.d_melee_texture_1 = $event.replaceAll('IT', ''); setFieldModifiedById('d_melee_texture_1')"
        />

        <item-model-selector
          v-if="selectorActive['d_melee_texture_2']"
          :selected-model="npc.d_melee_texture_2"
          @input="npc.d_melee_texture_2 = $event.replaceAll('IT', ''); setFieldModifiedById('d_melee_texture_2')"
        />

        <item-model-selector
          v-if="selectorActive['ammo_idfile']"
          :selected-model="npc.ammo_idfile"
          @input="npc.ammo_idfile = $event; setFieldModifiedById('ammo_idfile')"
        />

        <race-selector
          v-if="(selectorActive['race'] || selectorActive['gender'] || selectorActive['texture'] || selectorActive['helmtexture']) && npc"
          :race="npc.race"
          :gender="npc.gender"
          :texture="npc.texture"
          :helm-texture="npc.helmtexture"
          @selected="npc.race = $event.race; npc.gender = $event.gender; npc.texture = $event.texture; npc.helmtexture = $event.helmtexture"
        />

        <eq-window
          v-if="(isFacialSelectorActive()) && npc"
          title="Facial Appearance Selector"
        >
          <facial-appearance-selector
            :in-race="npc.race"
            :in-gender="npc.gender"
            :in-face="npc.face"
            :in-hair="npc.luclin_hairstyle"
            :in-hair-color="npc.luclin_haircolor"
            :in-beard="npc.luclin_beard"
            :in-beard-color="npc.luclin_beardcolor"
            @input="handleFacialAppearanceUpdate($event)"
          />
        </eq-window>

        <loot-sub-editor
          v-if="selectorActive['loottable_id']"
        />

        <merchant-sub-editor
          v-if="selectorActive['merchant_id']"
          @input="npc.merchant_id = $event; setFieldModifiedById('merchant_id')"
        />

      </div>
    </div>
  </div>
</template>

<script>
import EqWindowFancy            from "../../components/eq-ui/EQWindowFancy";
import EqWindow                 from "../../components/eq-ui/EQWindow";
import EqTabs                   from "../../components/eq-ui/EQTabs";
import EqTab                    from "../../components/eq-ui/EQTab";
import EqItemPreview            from "../../components/preview/EQItemCardPreview";
import EqCheckbox               from "../../components/eq-ui/EQCheckbox";
import EqWindowSimple           from "../../components/eq-ui/EQWindowSimple";
import LoaderCastBarTimer       from "../../components/LoaderCastBarTimer";
import ContentArea              from "../../components/layout/ContentArea";
import {Npcs}                   from "@/app/npcs";
import EqDebug                  from "../../components/eq-ui/EQDebug";
import EqNpcCardPreview         from "../../components/preview/EQNpcCardPreview";
import {DB_CLASSES}             from "@/app/constants/eq-classes-constants";
import {DB_RACE_NAMES}          from "@/app/constants/eq-races-constants";
import {BODYTYPES}              from "@/app/constants/eq-bodytype-constants";
import {EditFormFieldUtil}      from "@/app/forms/edit-form-field-util";
import NpcSpecialAbilities      from "../../components/tools/NpcSpecialAbilities";
import {DB_SKILLS}              from "@/app/constants/eq-skill-constants";
import {FLYMODE}                from "@/app/constants/eq-flymode-constants";
import ItemModelSelector        from "../../components/selectors/ItemModelSelector";
import {GENDER}                 from "@/app/constants/eq-gender-constants";
import {DB_ITEM_MATERIAL}       from "@/app/constants/eq-item-constants";
import RaceSelector             from "../../components/selectors/RaceSelector";
import FacialAppearanceSelector from "../../components/selectors/FacialAppearanceSelector";
import MerchantSubEditor        from "../../components/subeditors/MerchantSubEditor";
import LootSubEditor            from "../../components/subeditors/LootSubEditor";

const MILLISECONDS_BEFORE_WINDOW_RESET = 10000;

export default {
  name: "ItemEdit",
  components: {
    LootSubEditor,
    MerchantSubEditor,
    FacialAppearanceSelector,
    RaceSelector,
    ItemModelSelector,
    NpcSpecialAbilities,
    EqNpcCardPreview,
    EqDebug,
    ContentArea,
    LoaderCastBarTimer,
    EqWindowSimple,
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
      originalNpc: {}, // item record data; used to reference original values in tools

      // selectors
      selectorActive: {},
      lastResetTime: Date.now(),

      // state, loaded or not
      loaded: true,

      // tabs / fields
      tabs: this.getTabs()
    }
  },
  watch: {

    // reset state vars when we navigate away
    '$route'() {
      this.npc = null;
      // this.originalItem = {};

      // reset state vars when we navigate away
      // this.notification = ""
      EditFormFieldUtil.resetFieldEditedStatus()
      this.resetPreviewComponents()

      // reload
      this.load()
    },

  },
  async created() {
    this.load()
  },
  methods: {

    /**
     * Facial
     */
    handleFacialAppearanceUpdate(e) {
      this.npc.face              = e.face;
      this.npc.luclin_hairstyle  = e.hair;
      this.npc.luclin_haircolor  = e.hairColor ? e.hairColor : 0;
      this.npc.luclin_eyecolor   = e.eye;
      this.npc.luclin_eyecolor_2 = e.eye;
      this.npc.luclin_beard      = e.beard;
      this.npc.luclin_beardcolor = e.beardColor;
    },

    facialFields() {
      return [
        "face",
        "luclin_hairstyle",
        "luclin_haircolor",
        "luclin_eyecolor",
        "luclin_eyecolor_2",
        "luclin_beard",
        "luclin_beardcolor"
      ]
    },

    isFacialSelectorActive() {
      for (let f of this.facialFields()) {
        if (this.selectorActive[f]) {
          return true;
        }
      }
      return false
    },

    setFieldModifiedById(field) {
      EditFormFieldUtil.setFieldModifiedById(field)
    },

    /**
     * Selectors
     */
    isAnySelectorActive() {
      for (const [k, v] of Object.entries(this.selectorActive)) {
        if (this.selectorActive[k]) {
          return true;
        }
      }
    },
    shouldReset() {
      return (Date.now() - this.lastResetTime) > MILLISECONDS_BEFORE_WINDOW_RESET
    },
    previewMain(force = false) {
      if ((this.shouldReset() && this.isAnySelectorActive()) || force) {
        this.resetPreviewComponents()
        this.$forceUpdate()
      }
    },
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
      this.npc    = await Npcs.getNpc(this.$route.params.npc)
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
          "npc_faction_id",

          "race",
          "gender",
          "texture",
          "helmtexture",

          "ammo_idfile",
          "d_melee_texture_1",
          "d_melee_texture_2",

          "face",
          "luclin_hairstyle",
          "luclin_haircolor",
          "luclin_eyecolor",
          "luclin_eyecolor_2",
          "luclin_beardcolor",
          "luclin_beard",
        ];
        hasSubEditorFields.forEach((field) => {
          EditFormFieldUtil.setFieldHighlightHasSubEditor(field)
        })
      }, 1)

    },

    getFieldDescription(field) {
      return Npcs.getFieldDescription(field)
    },

    /**
     * Tabs / fields
     */
    getEventHandlers(e, field) {
      let handlers = {}
      if (e.onclick) {
        handlers.click = () => e.onclick(field)
      }
      if (e.onmouseover) {
        handlers.mouseover = () => e.onmouseover(field)
      }

      return handlers
    },

    getTabs() {
      return [
        {
          name: 'General',
          fields: [
            { desc: 'ID', field: 'id', fType: 'text', itemIcon: '6840', },
            { desc: 'Name', field: 'name', fType: 'text', itemIcon: '6840', },
            { desc: 'Last Name', field: 'lastname', fType: 'text', itemIcon: '6840', },
            { desc: 'Level', field: 'level', fType: 'text', itemIcon: '6840', },
            { desc: 'Class', field: 'class', selectData: DB_CLASSES, },

            { desc: 'Bodytype', field: 'bodytype', selectData: BODYTYPES, },

            { desc: "Walk Speed", field: "walkspeed", fType: "text" },
            { desc: "Run Speed", field: "runspeed", fType: "text" },

            { desc: "Loottable ID", field: "loottable_id", fType: "text", e: { onclick: this.setSelectorActive } },
            { desc: "Merchant ID", field: "merchant_id", fType: "text", e: { onclick: this.setSelectorActive } },
            { desc: "Alternate Currency ID", field: "alt_currency_id", fType: "text" },
            { desc: "NPC Spells ID", field: "npc_spells_id", fType: "text" },
            { desc: "NPC Spell Effects ID", field: "npc_spells_effects_id", fType: "text" },
            { desc: "NPC Faction ID", field: "npc_faction_id", fType: "text" },
            { desc: "Adventure Template Id", field: "adventure_template_id", fType: "text" },
            { desc: "Trap Template", field: "trap_template", fType: "text" },
            { desc: "Emote ID", field: "emoteid", fType: "text" },

            // add back later?
            // { desc: "Spawn Limit", field: "spawn_limit", fType: "text" },
            // { desc: "Exclude", field: "exclude", fType: "text" },

            // version is not really used outside of the PEQ editor
            // { desc: "version", field: "version", fType: "text" },

            { desc: "Stuck Behavior", field: "stuck_behavior", fType: "text" },
            { desc: "Flymode", field: "flymode", fType: "select", selectData: FLYMODE },
            // { desc: "always_aggro", field: "always_aggro", fType: "text" },
            { desc: "Experience Modifier", field: "exp_mod", fType: "text" },

          ],
        },
        {
          name: 'Weapon',
          fields: [
            {
              desc: "Primary Melee Weapon Model",
              field: "d_melee_texture_1",
              fType: "text",
              e: { onclick: this.setSelectorActive }
            },
            { desc: "Primary Melee Type", field: "prim_melee_type", fType: "select", selectData: DB_SKILLS },
            {
              desc: "Secondary Melee Weapon Model",
              field: "d_melee_texture_2",
              fType: "text",
              e: { onclick: this.setSelectorActive }
            },
            { desc: "Secondary Melee Type", field: "sec_melee_type", fType: "select", selectData: DB_SKILLS },
            { desc: "Ranged Melee Type", field: "ranged_type", fType: "select", selectData: DB_SKILLS },
            { desc: "Ammo Weapon Model", field: "ammo_idfile", fType: "text", e: { onclick: this.setSelectorActive } },
          ]
        },
        {
          name: 'Aggro',
          fields: [
            { desc: 'Always Aggro', field: 'always_aggro', fType: 'checkbox' },
            { desc: "NPC Aggro", field: "npc_aggro", fType: "checkbox" },

            { desc: "Aggro Radius", field: "aggroradius", fType: "text" },
            { desc: "Assist Radius", field: "assistradius", fType: "text" },
          ]
        },
        {
          name: 'Appearance',
          fields: [
            { desc: 'Race', field: 'race', selectData: DB_RACE_NAMES, e: { onmouseover: this.setSelectorActive } },
            {
              desc: "Gender",
              field: "gender",
              fType: "select",
              selectData: GENDER,
              e: { onmouseover: this.setSelectorActive }
            },
            {
              desc: "Texture",
              field: "texture",
              fType: "select",
              selectData: DB_ITEM_MATERIAL,
              e: { onmouseover: this.setSelectorActive }
            },
            { desc: "Helm Texture", field: "helmtexture", fType: "text", e: { onmouseover: this.setSelectorActive } },
            { desc: "Heros Forge Model", field: "herosforgemodel", fType: "text" },
            { desc: "Size", field: "size", fType: "text" },
            { desc: "Light", field: "light", fType: "text" },
            { desc: "Model?", field: "model", fType: "text" },
          ]
        },
        {
          name: 'Armor',
          fields: [
            { desc: "Armor Tint ID", field: "armortint_id", fType: "text" },
            { desc: "Armor Tint Red", field: "armortint_red", fType: "text" },
            { desc: "Armor Tint Green", field: "armortint_green", fType: "text" },
            { desc: "Armor Tint Blue", field: "armortint_blue", fType: "text" },

            { desc: "Arm Texture", field: "armtexture", fType: "text" },
            { desc: "Bracer Texture", field: "bracertexture", fType: "text" },
            { desc: "Hand Texture", field: "handtexture", fType: "text" },
            { desc: "Leg Texture", field: "legtexture", fType: "text" },
            { desc: "Feet Texture", field: "feettexture", fType: "text" },
          ]
        },
        {
          name: 'Face',
          fields: [
            { desc: "Face", field: "face", fType: "text", e: { onmouseover: this.setSelectorActive } },
            { desc: "Hairstyle", field: "luclin_hairstyle", fType: "text", e: { onmouseover: this.setSelectorActive } },
            { desc: "Haircolor", field: "luclin_haircolor", fType: "text", e: { onmouseover: this.setSelectorActive } },
            { desc: "Eyecolor 1", field: "luclin_eyecolor", fType: "text", e: { onmouseover: this.setSelectorActive } },
            {
              desc: "Eyecolor 2",
              field: "luclin_eyecolor_2",
              fType: "text",
              e: { onmouseover: this.setSelectorActive }
            },
            {
              desc: "Beardcolor",
              field: "luclin_beardcolor",
              fType: "text",
              e: { onmouseover: this.setSelectorActive }
            },
            { desc: "Beard", field: "luclin_beard", fType: "text", e: { onmouseover: this.setSelectorActive } },
            { desc: "(Drakkin) Heritage", field: "drakkin_heritage", fType: "text" },
            { desc: "(Drakkin) Tattoo", field: "drakkin_tattoo", fType: "text" },
            { desc: "(Drakkin) Details", field: "drakkin_details", fType: "text" },
          ]
        },
        {
          name: 'Stats',
          fields: [

            // deprecated
            // { desc: "npcspecialattks", field: "npcspecialattks", fType: "text" },

            { desc: "AC", field: "ac", fType: "text" },
            { desc: "HP", field: "hp", fType: "text" },
            { desc: "Mana", field: "mana", fType: "text" },
            { desc: "HP Regen (Tic)", field: "hp_regen_rate", fType: "text" },
            { desc: "HP Regen (Sec)", field: "hp_regen_per_second", fType: "text" },
            { desc: "Mana Regen (Tic)", field: "mana_regen_rate", fType: "text" },

            { desc: "Strength", field: "str", fType: "text" },
            { desc: "Stamina", field: "sta", fType: "text" },
            { desc: "Dexterity", field: "dex", fType: "text" },
            { desc: "Agility", field: "agi", fType: "text" },
            { desc: "Intelligence", field: "_int", fType: "text" },
            { desc: "Wisdom", field: "wis", fType: "text" },
            { desc: "Charisma", field: "cha", fType: "text" },

            { desc: "Spell Scale", field: "spellscale", fType: "text" },
            { desc: "Heal Scale", field: "healscale", fType: "text" },

            { desc: "Scale Rate", field: "scalerate", fType: "text" },
            { desc: "Max Level", field: "maxlevel", fType: "text" },
          ]
        },
        {
          name: 'Resists',
          fields: [
            { desc: "Magic Resist", field: "mr", fType: "text" },
            { desc: "Cold Resist", field: "cr", fType: "text" },
            { desc: "Disease Resist", field: "dr", fType: "text" },
            { desc: "Fire Resist", field: "fr", fType: "text" },
            { desc: "Poison Resist", field: "pr", fType: "text" },
            { desc: "Corruption Resist", field: "corrup", fType: "text" },
            { desc: "Physical Resist", field: "ph_r", fType: "text" },
          ]
        },
        {
          name: 'Combat',
          fields: [
            { desc: "Minimum Damage", field: "mindmg", fType: "text" },
            { desc: "Maximum Damage", field: "maxdmg", fType: "text" },
            { desc: "Attack Count", field: "attack_count", fType: "text" },

            { desc: "Attack Speed", field: "attack_speed", fType: "text" },
            { desc: "Attack Delay", field: "attack_delay", fType: "text" },
            { desc: "Attack", field: "atk", fType: "text" },
            { desc: "Accuracy", field: "accuracy", fType: "text" },
            { desc: "Avoidance", field: "avoidance", fType: "text" },
            { desc: "Slow Mitigation", field: "slow_mitigation", fType: "text" },

            {
              desc: "Special Abilities",
              field: "special_abilities",
              fType: "textarea",
              col: 'col-12',
              e: { onclick: this.setSelectorActive },
            },
          ]
        },
        {
          name: 'Charm',
          fields: [
            { desc: "AC", field: "charm_ac", fType: "text" },
            { desc: "Minimum Damage", field: "charm_min_dmg", fType: "text" },
            { desc: "Maximum Damage", field: "charm_max_dmg", fType: "text" },
            { desc: "Attack Delay", field: "charm_attack_delay", fType: "text" },
            { desc: "Accuracy Rating", field: "charm_accuracy_rating", fType: "text" },
            { desc: "Avoidance Rating", field: "charm_avoidance_rating", fType: "text" },
            { desc: "Attack", field: "charm_atk", fType: "text" },
          ]
        },
        {
          name: 'Settings',
          fields: [
            // checkboxes
            { desc: 'See Hide', field: 'see_hide', fType: 'checkbox', },
            { desc: 'See Improved Hide', field: 'see_improved_hide', fType: 'checkbox', },
            { desc: 'See Invisible', field: 'see_invis', fType: 'checkbox', },
            { desc: 'See Invis Undead', field: 'see_invis_undead', fType: 'checkbox', },
            { desc: 'Show Name', field: 'show_name', fType: 'checkbox', },
            { desc: 'Trackable', field: 'trackable', fType: 'checkbox', },
            { desc: 'Skip Global Loot', field: 'skip_global_loot', fType: 'checkbox', },
            { desc: 'No Target Hotkey', field: 'no_target_hotkey', fType: 'checkbox', },
            { desc: 'Findable', field: 'findable', fType: 'checkbox', },
            { desc: 'Untargetable', field: 'untargetable', fType: 'checkbox', },
            { desc: 'Underwater', field: 'underwater', fType: 'checkbox', },
            { desc: 'QGlobal', field: 'qglobal', fType: 'checkbox', },
            { desc: 'Ignore Despawn', field: 'ignore_despawn', fType: 'checkbox', },
            { desc: 'Quest NPC', field: 'isquest', fType: 'checkbox', },
            { desc: 'Unique Spawn', field: 'unique_spawn_by_name', fType: 'checkbox', },
            { desc: 'Rare Spawn', field: 'rare_spawn', fType: 'checkbox', },
            { desc: 'Always Aggro', field: 'always_aggro', fType: 'checkbox', },
            { desc: "NPC Aggro", field: "npc_aggro", fType: "checkbox" },
            { desc: "Raid Target", field: "raid_target", fType: "checkbox" },
            { desc: "Private Corpse", field: "private_corpse", fType: "checkbox" },
            { desc: "Is Bot", field: "isbot", fType: "checkbox" },
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
