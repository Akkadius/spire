<template>
  <div class="container-fluid">
    <div class="panel-body">
      <div class="panel panel-default">
        <div class="row">
          <div class="col-7">
            <eq-window style="margin-top: 30px" title="Edit Spell">

              <eq-tabs
                v-if="spell && tabSelected"
                class="spell-edit-card"
                :hover-open="true"
                @mouseover.native="previewSpell"
              >
                <eq-tab
                  name="Basic"
                  :selected="tabSelected['basic']"
                >

                  <div class="row">
                    <div class="col-2">
                      Id
                      <b-form-input v-model="spell.id"/>
                    </div>
                    <div class="col-7">
                      Name
                      <b-form-input v-model="spell.name"/>
                    </div>

                    <div class="col-2" @mouseover="drawIconSelector">
                      Icon
                      <b-form-input v-model="spell.new_icon"/>
                    </div>

                    <div class="col-1" v-if="spell.new_icon > 0" @mouseover="drawIconSelector">
                      <img
                        :src="spellCdnUrl + spell.new_icon + '.gif'"
                        class="mt-3"
                        :style="'width:35px;height:auto;border-radius: 10px; ' + 'border: 2px solid ' + getTargetTypeColor(this.spell['targettype']) + '; border-radius: 7px;'">
                    </div>

                  </div>

                  <div class="row">
                    <div class="col-8">
                      <div class="row">
                        <div class="col-6">
                          You Cast
                          <b-form-input v-model="spell.you_cast"/>
                        </div>
                        <div class="col-6">
                          Other Casts
                          <b-form-input v-model="spell.other_casts"/>
                        </div>
                      </div>

                      <div class="row">
                        <div class="col-6">
                          Cast On You
                          <b-form-input v-model="spell.cast_on_you"/>
                        </div>
                        <div class="col-6">
                          Cast On Other
                          <b-form-input v-model="spell.cast_on_other"/>
                        </div>
                      </div>

                      <div class="row">
                        <div class="col-6">
                          Spell Fades
                          <b-form-input v-model="spell.spell_fades"/>
                        </div>
                        <div class="col-6">
                          ID File
                          <b-form-input v-model="spell.player_1"/>
                        </div>
                      </div>

                    </div>
                    <div
                      class="col-4"
                      style="text-align: center"
                      @mouseover="drawSpellAnimationSelector"
                    >
                      <spell-animation-preview
                        class="mt-4"
                        :id="spell.spellanim"/>

                      Spell Animation
                      <b-form-input v-model.number="spell.spellanim"/>
                    </div>
                  </div>
                </eq-tab>
                <eq-tab name="Effects" class="effect-tab">
                  <div>
                    <b-input-group style="height: 30px; margin-bottom: 15px">
                      <template #prepend>
                        <b-input-group-text style="width: 40px;">#</b-input-group-text>
                      </template>

                      <b-form-input placeholder="Effect" disabled style="width: 150px"/>
                      <b-form-input placeholder="Base" disabled/>
                      <b-form-input placeholder="Limit" disabled/>
                      <b-form-input placeholder="Max" disabled/>
                      <b-form-input placeholder="Formula" disabled/>
                    </b-input-group>

                    <b-input-group v-for="i in 12" :key="i">
                      <template #prepend>
                        <b-input-group-text style="width: 40px;">{{ i }}</b-input-group-text>
                      </template>

                      <b-form-select v-model="spell['effectid_' + i]" style="width: 150px">
                        <b-form-select-option v-for="(effect, id) in DB_SPA" :key="id" :value="parseInt(id)">{{ id }})
                          {{
                            effect
                          }}
                        </b-form-select-option>
                      </b-form-select>

                      <b-form-input v-model="spell['effect_base_value_' + i]"/>
                      <b-form-input v-model="spell['max_' + i]"/>
                      <b-form-input v-model="spell['effect_limit_value_' + i]"/>
                      <b-form-input v-model="spell['formula_' + i]"/>
                    </b-input-group>

                  </div>
                </eq-tab>

                <eq-tab name="Effects+">

                  <!-- Knockback -->
                  <div class="row">
                    <div class="col-2 text-right">
                      <h6 class="eq-header mt-3">Knockback</h6>
                    </div>
                    <div class="col-4">
                      Push Up
                      <b-form-input v-model="spell.pushback"/>
                    </div>
                    <div class="col-4">
                      Push Back
                      <b-form-input v-model="spell.pushup"/>
                    </div>
                  </div>

                  <!-- Hate -->
                  <div class="row">
                    <div class="col-2 text-right">
                      <h6 class="eq-header mt-3">Hate</h6>
                    </div>
                    <div class="col-3">
                      Hate Modifier
                      <b-form-input v-model="spell.bonushate"/>
                    </div>
                    <div class="col-3">
                      Spell Hate Given
                      <b-form-input v-model="spell.hate_added"/>
                    </div>
                    <div class="col-3 text-center">
                      No Detrimental Spell Aggro
                      <eq-checkbox
                        class="mt-2 mb-2"
                        v-model="spell.field_198"
                        @input="spell.field_198 = $event"
                      />
                    </div>
                  </div>

                  <!-- Viral Spells -->
                  <div class="row">
                    <div class="col-2 text-right">
                      <h6 class="eq-header mt-3">Viral Spells</h6>
                    </div>
                    <div class="col-3">
                      Viral Range
                      <b-form-input v-model="spell.viral_range"/>
                    </div>
                    <div class="col-3">
                      Viral Targets
                      <b-form-input v-model="spell.viral_targets"/>
                    </div>
                    <div class="col-3">
                      Viral Timer
                      <b-form-input v-model="spell.viral_timer"/>
                    </div>
                  </div>

                  <!-- Focus -->
                  <div class="row">
                    <div class="col-2 text-right">
                      <h6 class="eq-header mt-3">Focus</h6>
                    </div>
                    <div class="col-3">
                      Max Targets
                      <b-form-input v-model="spell.maxtargets"/>
                    </div>
                    <div class="col-3">
                      Song Base Effect Cap
                      <b-form-input v-model="spell.songcap"/>
                    </div>
                    <div class="col-3 text-center">
                      Not Focusable
                      <eq-checkbox class="mt-2 mb-2" v-model="spell.not_extendable" @input="spell.not_extendable = $event"/>
                    </div>
                  </div>

                  <!-- Spell Group -->
                  <div class="row">
                    <div class="col-2 text-right">
                      <h6 class="eq-header mt-3">Spell Group</h6>
                    </div>
                    <div class="col-4">
                      Spell Group
                      <b-form-input v-model="spell.spellgroup"/>
                    </div>
                    <div class="col-4">
                      Rank
                      <b-form-input v-model="spell.rank"/>
                    </div>
                  </div>

                  Max Critical Chance
                  <b-form-input v-model="spell.field_217"/>

                  Nimbus Type
                  <b-form-input v-model="spell.nimbuseffect"/>
                  Max Hits Type
                  <b-form-input v-model="spell.numhitstype"/>
                  Max Hits Allowed
                  <b-form-input v-model="spell.numhits"/>
                  Recourse Spell ID
                  <b-form-input v-model="spell.recourse_link"/>

                </eq-tab>

                <eq-tab name="General">
                  <div class="row">
                    <div class="col-6">
                      Skill
                      <b-form-select v-model="spell.skill" v-if="DB_SKILLS">
                        <b-form-select-option
                          :value="parseInt(id)" v-for="(skill, id) in DB_SKILLS"
                          :key="id"
                        >{{ id }}) {{ skill }}
                        </b-form-select-option>
                      </b-form-select>
                    </div>
                    <div class="col-6">
                      Good Effect
                      <b-form-input v-model="spell.good_effect"/>
                    </div>
                  </div>

                  <div class="row">
                    <div class="col-3">
                      Mana Cost
                      <b-form-input v-model="spell.mana"/>
                    </div>
                    <div class="col-3">
                      Endurance Cost
                      <b-form-input v-model="spell.endur_cost"/>
                    </div>
                    <div class="col-3">
                      Endurance Upkeep
                      <b-form-input v-model="spell.endur_upkeep"/>
                    </div>
                    <div class="col-3 text-center">
                      Use Discipline Window
                      <eq-checkbox
                        class="mt-2 mb-2"
                        v-model="spell.is_discipline"
                        @input="spell.is_discipline = $event"
                      />
                    </div>
                  </div>

                  <div class="row">
                    <div class="col-6">
                      pcnpc_only_flag
                      <b-form-input v-model="spell.pcnpc_only_flag"/>
                    </div>
                    <div class="col-6">
                      Teleport Zone / Pet DbaseID / ItemGraphic for Bolt Spells
                      <b-form-input v-model="spell.teleport_zone"/>
                    </div>
                  </div>
                </eq-tab>

                <eq-tab name="Restrictions">
                  <div class="row">
                    <div class="col-3" style="text-align:center">
                      Must Be Out of Combat
                      <eq-checkbox class="mt-2 mb-2" v-model="spell.outof_combat" @input="spell.outof_combat = $event"/>
                    </div>
                    <div class="col-3" style="text-align:center">
                      Must Be In Combat
                      <eq-checkbox class="mt-2 mb-2" v-model="spell.in_combat" @input="spell.in_combat = $event"/>
                    </div>
                    <div class="col-3" style="text-align:center">
                      Only During Fast Regen
                      <eq-checkbox class="mt-2 mb-2" v-model="spell.allowrest" @input="spell.allowrest = $event"/>
                    </div>
                    <div class="col-3" style="text-align:center">
                      Cancel On Sit
                      <eq-checkbox class="mt-2 mb-2" v-model="spell.disallow_sit" @input="spell.disallow_sit = $event"/>
                    </div>
                    <div class="col-3" style="text-align:center">
                      Must be Sneaking
                      <eq-checkbox class="mt-2 mb-2" v-model="spell.sneaking" @input="spell.sneaking = $event"/>
                    </div>
                  </div>

                  Target Restriction
                  <b-form-input v-model="spell.cast_restriction"/>
                  Caster Restriction
                  <b-form-input v-model="spell.field_220"/>

                  Zone Type (select)
                  <b-form-input v-model="spell.zonetype"/>
                  Environment Type
                  <b-form-input v-model="spell.environment_type"/>
                  Time of Day
                  <b-form-input v-model="spell.time_of_day"/>
                </eq-tab>
                <eq-tab name="Casting">
                  Cast Time
                  <b-form-input v-model="spell.cast_time"/>
                  Recovery Time
                  <b-form-input v-model="spell.recovery_time"/>
                  Recast Time
                  <b-form-input v-model="spell.recast_time"/>
                  Timer Index (Timer Max 19, also seen few discs with -1)
                  <b-form-input v-model="spell.endur_timer_index"/>
                  Uninterruptable
                  <eq-checkbox class="mt-2 mb-2" v-model="spell.uninterruptable"
                               @input="spell.uninterruptable = $event"/>

                  Fizzle Adjustment
                  <b-form-input v-model="spell.basediff"/>
                  Cast Not Standing (Can Cast from Sitting position, Can cast on invulnerable Targets, Can not be
                  interrupted by SE_InterruptCasting)
                  <b-form-input v-model="spell.cast_not_standing"/>
                </eq-tab>
                <eq-tab name="Buffing">

                  <div class="row">
                    <div class="col-3" style="text-align:center">
                      Can Not Dispell
                      <eq-checkbox class="mt-2 mb-2" v-model="spell.nodispell" @input="spell.nodispell = $event"/>
                    </div>
                    <div class="col-3" style="text-align:center">
                      Can Not Click Off
                      <eq-checkbox class="mt-2 mb-2" v-model="spell.field_232" @input="spell.field_232 = $event"/>
                    </div>
                    <div class="col-3" style="text-align:center">
                      Persist After Death
                      <eq-checkbox class="mt-2 mb-2" v-model="spell.persistdeath" @input="spell.persistdeath = $event"/>
                    </div>
                    <div class="col-3" style="text-align:center">
                      Suspendable
                      <eq-checkbox class="mt-2 mb-2" v-model="spell.suspendable" @input="spell.suspendable = $event"/>
                    </div>
                  </div>

                  <div class="row">
                    <div class="col-3" style="text-align:center">
                      Can MGB
                      <eq-checkbox class="mt-2 mb-2" v-model="spell.can_mgb" @input="spell.can_mgb = $event"/>
                    </div>

                    <div class="col-3" style="text-align:center">
                      Short Duration Buff
                      <eq-checkbox
                        class="mt-2 mb-2"
                        v-model="spell.short_buff_box"
                        @input="spell.short_buff_box = $event"
                      />
                    </div>
                    <div class="col-3" style="text-align:center">
                      No Buff Block
                      <eq-checkbox
                        class="mt-2 mb-2"
                        v-model="spell.no_block"
                        @input="spell.no_block = $event"
                      />
                    </div>
                    <div class="col-3" style="text-align:center">
                      DOT Not Stackable
                      <eq-checkbox
                        class="mt-2 mb-2"
                        v-model="spell.dot_stacking_exempt"
                        @input="spell.dot_stacking_exempt = $event"
                      />
                    </div>
                  </div>

                  <div class="row">
                    <div class="col-3">
                      Buff Duration
                      <b-form-input v-model="spell.buffduration"/>
                    </div>
                    <div class="col-3">
                      Duration Formula
                      <b-form-input v-model="spell.buffdurationformula"/>
                    </div>
                    <div class="col-3">
                      PVP Duration
                      <b-form-input v-model="spell.field_181"/>

                    </div>
                    <div class="col-3">
                      PVP Duration Cap
                      <b-form-input v-model="spell.field_182"/>
                    </div>
                  </div>


                </eq-tab>
                <eq-tab name="Range">

                  <div class="row">
                    <div class="col-3">
                      Spell Range
                      <b-form-input v-model="spell.range"/>
                    </div>
                    <div class="col-3">
                      Target Type
                      <b-form-select v-model="spell.targettype" v-if="DB_SPELL_TARGETS">
                        <b-form-select-option
                          :value="parseInt(id)"
                          v-for="(value, id) in DB_SPELL_TARGETS"
                          :key="id"
                        >{{ id }})
                          {{ value }}
                        </b-form-select-option>
                      </b-form-select>
                    </div>
                    <div class="col-6 text-center">
                      NPC Line of Sight Not Required to Cast
                      <eq-checkbox class="mt-2 mb-2" v-model="spell.npc_no_los" @input="spell.npc_no_los = $event"/>
                    </div>
                  </div>

                  <h6 class="eq-header">Area of Effect (AOE)</h6>

                  <div class="row">
                    <div class="col-4">
                      AOE Range
                      <b-form-input v-model="spell.aoerange"/>
                    </div>
                    <div class="col-4">
                      AOE Rain Waves
                      <b-form-input v-model="spell.ae_duration"/>
                    </div>
                    <div class="col-4">
                      AOE Max Targets
                      <b-form-input v-model="spell.aemaxtargets"/>
                    </div>
                  </div>

                  <div class="row">
                    <div class="col-2">
                      Min Range
                      <b-form-input v-model="spell.min_range"/>
                    </div>
                    <div class="col-2">
                      Min Distance for Mod
                      <b-form-input v-model="spell.min_dist"/>
                    </div>
                    <div class="col-2">
                      Min Distance Mod
                      <b-form-input v-model="spell.min_dist_mod"/>
                    </div>
                    <div class="col-3">
                      Max Distance for Mod
                      <b-form-input v-model="spell.max_dist"/>
                    </div>
                    <div class="col-3">
                      Max Distance Mod
                      <b-form-input v-model="spell.max_dist_mod"/>
                    </div>
                  </div>

                  <h6 class="eq-header">Cone</h6>

                  <div class="row">
                    <div class="col-6">
                      Cone Angle Start
                      <b-form-input v-model="spell.cone_start_angle"/>
                    </div>
                    <div class="col-6">
                      Cone Angle End
                      <b-form-input v-model="spell.cone_stop_angle"/>
                    </div>
                  </div>

                </eq-tab>
                <eq-tab name="Resist">
                  Resist Type
                  <b-form-select
                    v-model="spell.resisttype"
                    v-if="DB_SPELL_RESISTS"
                  >
                    <b-form-select-option
                      :value="parseInt(id)" v-for="(value, id) in DB_SPELL_RESISTS"
                      :key="id"
                    >{{ id }})
                      {{ value }}
                    </b-form-select-option>
                  </b-form-select>

                  Unresistable
                  <eq-checkbox class="mt-2 mb-2" v-model="spell.field_209" @input="spell.field_209 = $event"/>

                  Resist Diff
                  <b-form-input v-model="spell.resist_diff"/>
                  No Partial Resists
                  <eq-checkbox
                    class="mt-2 mb-2"
                    v-model="spell.no_partial_resist"
                    @input="spell.no_partial_resist = $event"/>

                  Resist Chance Limits: Max Chance (Actual in game chance is divided by 2)
                  <b-form-input v-model="spell.max_resist"/>
                  Resist Chance Limits: Min Chance (Actual in game chance is divided by 2)
                  <b-form-input v-model="spell.min_resist"/>
                  PVP Resist Mod
                  <b-form-input v-model="spell.pvpresistbase"/>
                  PVP Resist Per Level
                  <b-form-input v-model="spell.pvpresistcalc"/>
                  PVP Resist Cap
                  <b-form-input v-model="spell.pvpresistcap"/>
                  Reflectable
                  <b-form-input v-model="spell.reflectable"/>
                  Feedbackable
                  <b-form-input v-model="spell.field_160"/>
                </eq-tab>
              </eq-tabs>

            </eq-window>
          </div>

          <div class="col-5">

            <!-- preview spell -->
            <eq-window
              style="margin-top: 30px; margin-right: 10px; width: auto;"
              v-if="previewSpellActive">
              <eq-spell-preview
                :spell-data="spell"/>
            </eq-window>

            <!-- icon selector -->
            <eq-window
              style="margin-top: 30px; margin-right: 10px; width: auto;"
              class="fade-in"
              v-if="iconSelectorActive">
              <spell-icon-selector
                :inputData.sync="spell.new_icon"
              />
            </eq-window>

            <!-- spell anim selector -->
            <eq-window
              style="margin-top: 30px; margin-right: 10px; width: auto;"
              class="fade-in"
              v-if="spellAnimSelectorActive">

              <!--              <spell-animation-viewer :is-component="true"/>-->
              <spell-animation-selector
                :inputData.sync="spell.spellanim"
              />

              <!--              <spell-icon-selector-->
              <!--                :inputData.sync="spell.new_icon"-->
              <!--              />-->
            </eq-window>


          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import EqWindowFancy                                                  from "../../components/eq-ui/EQWindowFancy";
import EqWindow                                                       from "../../components/eq-ui/EQWindow";
import EqTabs                                                         from "../../components/eq-ui/EQTabs";
import EqTab                                                          from "../../components/eq-ui/EQTab";
import EqSpellPreview                                                 from "../../components/eq-ui/EQSpellCardPreview";
import {Spells}                                                       from "../../app/spells";
import {DB_SPA, DB_SPELL_EFFECTS, DB_SPELL_RESISTS, DB_SPELL_TARGETS} from "../../app/constants/eq-spell-constants";
import {DB_SKILLS}                                                    from "../../app/constants/eq-skill-constants";
import {App}                                                          from "../../constants/app";
import SpellIconSelector                                              from "../../components/tools/SpellIconSelector";
import SpellAnimationPreview
                                                                      from "../../components/tools/SpellAnimationPreview";
import SpellAnimationViewer                                           from "./SpellAnimationViewer";
import SpellAnimationSelector
                                                                      from "../../components/tools/SpellAnimationSelector";
import EqCheckbox                                                     from "../../components/eq-ui/EQCheckbox";

const MILLISECONDS_BEFORE_WINDOW_RESET = 3000;

export default {
  name: "SpellEdit",
  components: {
    EqCheckbox,
    SpellAnimationSelector,
    SpellAnimationViewer,
    SpellAnimationPreview,
    SpellIconSelector,
    EqSpellPreview,
    EqTab,
    EqTabs,
    EqWindow,
    EqWindowFancy
  },
  data() {
    return {
      spell: null, // spell record data
      spellCdnUrl: App.ASSET_SPELL_ICONS_BASE_URL,
      DB_SPELL_EFFECTS: DB_SPELL_EFFECTS,
      DB_SPA: DB_SPA,
      DB_SKILLS: DB_SKILLS,
      DB_SPELL_TARGETS: DB_SPELL_TARGETS,
      DB_SPELL_RESISTS: DB_SPELL_RESISTS,
      loaded: true,

      previewSpellActive: true,
      iconSelectorActive: false,
      spellAnimSelectorActive: false,

      tabSelected: { 'basic': true },

      lastResetTime: Date.now(),
    }
  },
  watch: {
    '$route'() {
      this.load()
    }
  },
  async created() {
    this.load()
  },
  methods: {
    load() {
      if (this.$route.params.id > 0) {
        Spells.getSpell(this.$route.params.id).then(result => {
          this.spell = result
        })
      }
    },

    /**
     * Selector / previewers
     */
    resetPreviewComponents() {
      this.previewSpellActive      = false;
      this.iconSelectorActive      = false;
      this.spellAnimSelectorActive = false;
    },
    previewSpell() {
      let shouldReset = Date.now() - this.lastResetTime > MILLISECONDS_BEFORE_WINDOW_RESET;
      // SECONDS_BEFORE_WINDOW_RESET

      if (!this.previewSpellActive && shouldReset) {
        this.resetPreviewComponents()
        this.previewSpellActive = true;
        this.lastResetTime      = Date.now()
      }
    },
    drawSpellAnimationSelector() {
      this.resetPreviewComponents()
      this.spellAnimSelectorActive = true
    },
    drawIconSelector() {
      this.resetPreviewComponents()
      this.iconSelectorActive = true;
    },


    getTargetTypeColor(targetType) {
      return Spells.getTargetTypeColor(targetType);
    },
    redrawCard() {
      this.loaded = false

      setTimeout(() => {
        this.loaded = true
      }, 5)
    }
  }
}
</script>

<style scoped>
.spell-edit-card input, .spell-edit-card select {
  margin-bottom: 10px;
}

.effect-tab input, .effect-tab select {
  margin-bottom: 0;
}
</style>
