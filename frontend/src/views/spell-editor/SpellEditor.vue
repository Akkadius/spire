<template>
  <div class="container-fluid">
    <div class="panel-body">
      <div class="panel panel-default">
        <div class="row">
          <div class="col-7">
            <eq-window style="margin-top: 30px" title="Edit Spell">

              <div
                v-if="notification"
                :class="'text-center mt-2 btn-xs eq-header fade-in'"
                style="width: 100%; font-size: 30px"
                @click="notification = ''"
              >
                <i class="ra ra-book mr-1"></i>
                {{ notification }}
              </div>

              <b-alert show dismissable variant="danger" v-if="error">
                <i class="fa fa-warning"></i> {{ error }}
              </b-alert>

              <eq-tabs
                v-if="spell"
                id="spell-edit-card"
                class="spell-edit-card"
                @mouseover.native="previewSpell"
              >
                <eq-tab
                  name="Basic"
                  :selected="true"
                >

                  <div class="row">
                    <div class="col-2" @mouseover="drawFreeIdSelector">
                      Id
                      <b-form-input v-model.number="spell.id"/>
                    </div>
                    <div class="col-7">
                      Name
                      <b-form-input
                        :value="spell.name" @change="v => spell.name = v"
                      />
                    </div>

                    <div class="col-2" @mouseover="drawIconSelector">
                      Icon
                      <b-form-input v-model.number="spell.new_icon"/>
                    </div>

                    <div
                      class="col-1" v-if="spell.new_icon > 0"
                      style="margin-top: 7px"
                      @mouseover="drawIconSelector"
                    >
                      <img
                        :src="spellCdnUrl + spell.new_icon + '.gif'"
                        class="mt-3"
                        :style="'width:35px;height:auto;border-radius: 10px; ' + 'border: 2px solid ' + getTargetTypeColor(this.spell['targettype']) + '; border-radius: 7px;'"
                      >
                    </div>

                  </div>

                  <spell-class-selector :spell="spell" @input="spell = $event"/>
                  <!--                  <spell-deity-selector :spell="spell" @input="spell = $event"/>-->

                  <div class="row">
                    <div class="col-8">
                      <div class="row">
                        <div class="col-6">
                          You Cast
                          <b-form-input
                            :value="spell.you_cast" @change="v => spell.you_cast = v"
                          />
                        </div>
                        <div class="col-6">
                          Other Casts
                          <b-form-input
                            :value="spell.other_casts" @change="v => spell.other_casts = v"
                          />
                        </div>
                      </div>

                      <div class="row">
                        <div class="col-6">
                          Cast On You
                          <b-form-input
                            :value="spell.cast_on_you" @change="v => spell.cast_on_you = v"
                          />
                        </div>
                        <div class="col-6">
                          Cast On Other
                          <b-form-input
                            :value="spell.cast_on_other" @change="v => spell.cast_on_other = v"
                          />
                        </div>
                      </div>

                      <div class="row">
                        <div class="col-6">
                          Spell Fades
                          <b-form-input
                            :value="spell.spell_fades" @change="v => spell.spell_fades = v"
                          />
                        </div>
                        <div class="col-6">
                          ID File
                          <b-form-input
                            :value="spell.player_1" @change="v => spell.player_1 = v"
                          />
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
                        :id="spell.spellanim"
                      />

                      Spell Animation
                      <b-form-input v-model.number="spell.spellanim"/>
                    </div>
                  </div>


                </eq-tab>
                <eq-tab name="Effects" class="effect-tab">
                  <div>
                    <b-input-group style="height: 30px; margin-bottom: 8px">
                      <template #prepend>
                        <b-input-group-text
                          style="width: 40px; "
                        >#
                        </b-input-group-text>
                      </template>

                      <b-form-input placeholder="Effect" disabled style="width: 150px"/>
                      <b-form-input placeholder="Base" disabled/>
                      <b-form-input placeholder="Limit" disabled/>
                      <b-form-input placeholder="Max" disabled/>
                      <b-form-input placeholder="Formula" disabled/>
                    </b-input-group>

                    <b-input-group v-for="i in 12" :key="i" style="margin-top: -1px">
                      <template #prepend>
                        <b-input-group-text style="width: 40px; ">{{ i }}</b-input-group-text>
                      </template>

                      <b-form-select
                        @mouseover.native="drawSpaDetailPane(spell['effectid_' + i], i)"
                        v-model.number="spell['effectid_' + i]"
                        style="width: 150px"
                      >
                        <b-form-select-option
                          v-for="(effect, id) in DB_SPA"
                          :key="id"
                          :value="parseInt(id)"
                        >{{ id }}) {{ effect }}
                        </b-form-select-option>
                      </b-form-select>

                      <b-form-input v-model.number="spell['effect_base_value_' + i]"/>
                      <b-form-input v-model.number="spell['effect_limit_value_' + i]"/>
                      <b-form-input v-model.number="spell['max_' + i]"/>
                      <b-form-input v-model.number="spell['formula_' + i]"/>
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
                      <b-form-input v-model.number="spell.pushback"/>
                    </div>
                    <div class="col-4">
                      Push Back
                      <b-form-input v-model.number="spell.pushup"/>
                    </div>
                  </div>

                  <!-- Hate -->
                  <div class="row">
                    <div class="col-2 text-right">
                      <h6 class="eq-header mt-3">Hate</h6>
                    </div>
                    <div class="col-3">
                      Hate Modifier
                      <b-form-input v-model.number="spell.bonushate"/>
                    </div>
                    <div class="col-3">
                      Spell Hate Given
                      <b-form-input v-model.number="spell.hate_added"/>
                    </div>
                    <div class="col-3 text-center">
                      No Detrimental Spell Aggro
                      <eq-checkbox
                        class="mt-2 mb-2"
                        v-model.number="spell.field_198"
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
                      <b-form-input v-model.number="spell.viral_range"/>
                    </div>
                    <div class="col-3">
                      Viral Targets
                      <b-form-input v-model.number="spell.viral_targets"/>
                    </div>
                    <div class="col-3">
                      Viral Timer
                      <b-form-input v-model.number="spell.viral_timer"/>
                    </div>
                  </div>

                  <!-- Focus -->
                  <div class="row">
                    <div class="col-2 text-right">
                      <h6 class="eq-header mt-3">Focus</h6>
                    </div>
                    <div class="col-3">
                      Max Targets
                      <b-form-input v-model.number="spell.maxtargets"/>
                    </div>
                    <div class="col-3">
                      Song Base Effect Cap
                      <b-form-input v-model.number="spell.songcap"/>
                    </div>
                    <div class="col-3 text-center">
                      Not Focusable
                      <eq-checkbox
                        class="mt-2 mb-2" v-model.number="spell.not_extendable"
                        @input="spell.not_extendable = $event"
                      />
                    </div>
                  </div>

                  <!-- Spell Group -->
                  <div class="row">
                    <div class="col-2 text-right">
                      <h6 class="eq-header mt-3">Spell Group</h6>
                    </div>
                    <div class="col-4">
                      Spell Group
                      <b-form-input v-model.number="spell.spellgroup"/>
                    </div>
                    <div class="col-4">
                      Rank
                      <b-form-input v-model.number="spell.rank"/>
                    </div>
                  </div>

                  <!-- Misc -->
                  <div class="row">
                    <div class="col-2 text-right">
                      <h6 class="eq-header mt-3">Misc</h6>
                    </div>
                    <div class="col-3">
                      Max Critical Chance
                      <b-form-input v-model.number="spell.field_217"/>
                    </div>
                    <div class="col-3">
                      Nimbus Type
                      <b-form-input v-model.number="spell.nimbuseffect"/>
                    </div>
                    <div class="col-3">
                      Recourse Spell ID
                      <b-form-input v-model.number="spell.recourse_link"/>
                    </div>
                  </div>
                </eq-tab>

                <eq-tab name="General">
                  <div class="row">
                    <div class="col-6">
                      Skill
                      <b-form-select v-model.number="spell.skill" v-if="DB_SKILLS">
                        <b-form-select-option
                          :value="parseInt(id)" v-for="(skill, id) in DB_SKILLS"
                          :key="id"
                        >{{ id }}) {{ skill }}
                        </b-form-select-option>
                      </b-form-select>
                    </div>
                    <div class="col-6">
                      Good Effect
                      <b-form-input v-model.number="spell.good_effect"/>
                    </div>
                  </div>

                  <div class="row">
                    <div class="col-3">
                      Mana Cost
                      <b-form-input v-model.number="spell.mana"/>
                    </div>
                    <div class="col-3">
                      Endurance Cost
                      <b-form-input v-model.number="spell.endur_cost"/>
                    </div>
                    <div class="col-3">
                      Endurance Upkeep
                      <b-form-input v-model.number="spell.endur_upkeep"/>
                    </div>
                    <div class="col-3 text-center">
                      Use Discipline Window
                      <eq-checkbox
                        class="mt-2 mb-2"
                        v-model.number="spell.is_discipline"
                        @input="spell.is_discipline = $event"
                      />
                    </div>
                  </div>

                  <div class="row">
                    <div class="col-6">
                      pcnpc_only_flag (???)
                      <b-form-input v-model.number="spell.pcnpc_only_flag"/>
                    </div>
                    <div class="col-6">
                      Teleport Zone / Pet DbaseID / ItemGraphic for Bolt Spells
                      <b-form-input v-model.number="spell.teleport_zone"/>
                    </div>
                  </div>
                </eq-tab>

                <eq-tab name="Restrictions">

                  <div class="row">
                    <div class="col-3">
                      <div
                        class="row" v-for="field in
                         [
                           {
                             description: 'Can Cast out of Combat',
                             field: 'outof_combat'
                           },
                           {
                             description: 'Can Cast in Combat',
                             field: 'in_combat',
                           },
                           {
                             description: 'Only During Fast Regen',
                             field: 'allowrest',
                           },
                           {
                             description: 'Cancel on Sit',
                             field: 'disallow_sit',
                           },
                           {
                             description: 'Must be Sneaking',
                             field: 'sneaking',
                           },
                         ]"
                      >
                        <div class="col-9 text-right p-0 pr-2 m-0">
                          {{ field.description }}
                        </div>
                        <div class="col-3 text-left p-0">
                          <eq-checkbox
                            v-b-tooltip.hover.v-dark.right :title="getFieldDescription(field.field)"
                            class="mb-2 d-inline-block"
                            :true-value="(typeof field.true !== 'undefined' ? field.true : 1)"
                            :false-value="(typeof field.false !== 'undefined' ? field.false : 0)"
                            v-model.number="spell[field.field]"
                            @input="spell[field.field] = $event"
                          />
                        </div>
                      </div>
                    </div>

                    <div class="col-9">
                      <div class="row">
                        <div class="col-4">
                          Target Restriction
                          <b-form-input v-model.number="spell.cast_restriction"/>
                        </div>
                        <div class="col-4">
                          Caster Restriction
                          <b-form-input v-model.number="spell.field_220"/>
                        </div>
                        <div class="col-4">
                          Zone Type (select)
                          <b-form-input v-model.number="spell.zonetype"/>
                        </div>

                        <div class="col-4">
                          Environment Type (???)
                          <b-form-input v-model.number="spell.environment_type"/>
                        </div>
                        <div class="col-4">
                          Time of Day (???)
                          <b-form-input v-model.number="spell.time_of_day"/>
                        </div>
                      </div>
                    </div>

                  </div>
                </eq-tab>
                <eq-tab name="Casting">

                  <div class="row">
                    <div class="col-3">
                      Cast Time (Clarify)
                      <b-form-input v-model.number="spell.cast_time"/>
                    </div>
                    <div class="col-3">
                      Recovery Time (Clarify)
                      <b-form-input v-model.number="spell.recovery_time"/>
                    </div>
                    <div class="col-3">
                      Recast Time (Clarify)
                      <b-form-input v-model.number="spell.recast_time"/>
                    </div>
                    <div class="col-3">
                      Timer Index (???)
                      <b-form-input v-model.number="spell.endur_timer_index"/>
                    </div>
                  </div>

                  <div class="row">
                    <div class="col-4">
                      Fizzle Adjustment
                      <b-form-input v-model.number="spell.basediff"/>
                    </div>
                    <div class="col-4">
                      Cast Not Standing
                      <b-form-input v-model.number="spell.cast_not_standing"/>
                    </div>
                    <div class="col-4 text-center">
                      Uninterruptable
                      <eq-checkbox
                        class="mt-2 mb-2"
                        v-model.number="spell.uninterruptable"
                        @input="spell.uninterruptable = $event"
                      />
                    </div>
                  </div>

                </eq-tab>
                <eq-tab name="Buffing">

                  <div class="row">
                    <div class="col-3">
                      <div
                        class="row" v-for="field in
                         [
                           {
                             description: 'Can Not Dispell',
                             field: 'nodispell'
                           },
                           {
                             description: 'Can Not Click Off',
                             field: 'field_232'
                           },
                           {
                             description: 'Persist After Death',
                             field: 'persistdeath'
                           },
                           {
                             description: 'Suspendable',
                             field: 'suspendable'
                           },
                           {
                             description: 'Can MGB',
                             field: 'can_mgb'
                           },
                           {
                             description: 'Appear In Short Buff Box',
                             field: 'short_buff_box'
                           },
                           {
                             description: 'No Buff Block',
                             field: 'no_block'
                           },
                           {
                             description: 'DOT Not Stackable',
                             field: 'dot_stacking_exempt'
                           },
                         ]"
                      >
                        <div class="col-9 text-right p-0 pr-2 m-0">
                          {{ field.description }}
                        </div>
                        <div class="col-3 text-left p-0">
                          <eq-checkbox
                            v-b-tooltip.hover.v-dark.right :title="getFieldDescription(field.field)"
                            class="mb-2 d-inline-block"
                            :true-value="(typeof field.true !== 'undefined' ? field.true : 1)"
                            :false-value="(typeof field.false !== 'undefined' ? field.false : 0)"
                            v-model.number="spell[field.field]"
                            @input="spell[field.field] = $event"
                          />
                        </div>
                      </div>
                    </div>

                    <div class="col-9">
                      <div class="row">
                        <div class="col-3">
                          Buff Duration
                          <b-form-input v-model.number="spell.buffduration"/>
                        </div>
                        <div class="col-3">
                          Duration Formula
                          <b-form-input v-model.number="spell.buffdurationformula"/>
                        </div>
                        <div class="col-3">
                          PVP Duration
                          <b-form-input v-model.number="spell.pvp_duration"/>

                        </div>
                        <div class="col-3">
                          PVP Duration Cap
                          <b-form-input v-model.number="spell.pvp_duration_cap"/>
                        </div>
                      </div>
                    </div>

                  </div>
                </eq-tab>
                <eq-tab name="Range">
                  <div class="row">
                    <div class="col-3">
                      Spell Range
                      <b-form-input v-model.number="spell.range"/>
                    </div>
                    <div class="col-3">
                      Target Type
                      <b-form-select v-model.number="spell.targettype" v-if="DB_SPELL_TARGETS">
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
                      <eq-checkbox
                        class="mt-2 mb-2" v-model.number="spell.npc_no_los"
                        @input="spell.npc_no_los = $event"
                      />
                    </div>
                  </div>

                  <h6 class="eq-header">Area of Effect (AOE)</h6>

                  <div class="row">
                    <div class="col-4">
                      AOE Range
                      <b-form-input v-model.number="spell.aoerange"/>
                    </div>
                    <div class="col-4">
                      AOE Rain Waves
                      <b-form-input v-model.number="spell.ae_duration"/>
                    </div>
                    <div class="col-4">
                      AOE Max Targets
                      <b-form-input v-model.number="spell.aemaxtargets"/>
                    </div>
                  </div>

                  <div class="row">
                    <div class="col-2">
                      Min Range
                      <b-form-input v-model.number="spell.min_range"/>
                    </div>
                    <div class="col-2">
                      Min Distance for Mod
                      <b-form-input v-model.number="spell.min_dist"/>
                    </div>
                    <div class="col-2">
                      Min Distance Mod
                      <b-form-input v-model.number="spell.min_dist_mod"/>
                    </div>
                    <div class="col-3">
                      Max Distance for Mod
                      <b-form-input v-model.number="spell.max_dist"/>
                    </div>
                    <div class="col-3">
                      Max Distance Mod
                      <b-form-input v-model.number="spell.max_dist_mod"/>
                    </div>

                  </div>

                  <div class="row">
                    <div class="col-6">
                      Max Hits Type
                      <b-form-input v-model.number="spell.numhitstype"/>
                    </div>
                    <div class="col-6">
                      Max Hits Allowed
                      <b-form-input v-model.number="spell.numhits"/>
                    </div>
                  </div>

                  <h6 class="eq-header">Cone</h6>

                  <div class="row">
                    <div class="col-6">
                      Cone Angle Start
                      <b-form-input v-model.number="spell.cone_start_angle"/>
                    </div>
                    <div class="col-6">
                      Cone Angle End
                      <b-form-input v-model.number="spell.cone_stop_angle"/>
                    </div>
                  </div>

                </eq-tab>
                <eq-tab name="Resist">

                  <div class="row">
                    <div class="col-3">
                      Resist Type
                      <b-form-select
                        v-model.number="spell.resisttype"
                        v-if="DB_SPELL_RESISTS"
                      >
                        <b-form-select-option
                          :value="parseInt(id)" v-for="(value, id) in DB_SPELL_RESISTS"
                          :key="id"
                        >{{ id }})
                          {{ value }}
                        </b-form-select-option>
                      </b-form-select>
                    </div>
                    <div class="col-3">
                      Resist Diff
                      <b-form-input v-model.number="spell.resist_diff"/>
                    </div>
                    <div class="col-3 text-center">
                      Unresistable
                      <eq-checkbox
                        class="mt-2 mb-2" v-model.number="spell.field_209"
                        @input="spell.field_209 = $event"
                      />
                    </div>
                    <div class="col-3 text-center">
                      No Partial Resists
                      <eq-checkbox
                        class="mt-2 mb-2"
                        v-model.number="spell.no_partial_resist"
                        @input="spell.no_partial_resist = $event"
                      />
                    </div>
                  </div>

                  <div class="row">
                    <div class="col-4">
                      PVP Resist Mod
                      <b-form-input v-model.number="spell.pvpresistbase"/>
                    </div>
                    <div class="col-4">
                      PVP Resist Per Level
                      <b-form-input v-model.number="spell.pvpresistcalc"/>
                    </div>
                    <div class="col-4">
                      PVP Resist Cap
                      <b-form-input v-model.number="spell.pvpresistcap"/>
                    </div>
                  </div>

                  Resist Chance Limits: Max Chance (Actual in game chance is divided by 2)
                  <b-form-input v-model.number="spell.max_resist"/>
                  Resist Chance Limits: Min Chance (Actual in game chance is divided by 2)
                  <b-form-input v-model.number="spell.min_resist"/>

                  Reflectable
                  <b-form-input v-model.number="spell.reflectable"/>
                  Feedbackable
                  <b-form-input v-model.number="spell.field_160"/>
                </eq-tab>
              </eq-tabs>

              <div class="text-center align-content-center mt-3">

                <div
                  :class="'text-center mt-2 btn-xs eq-button-fancy'"
                  @click="saveSpell()"
                >
                  <i class="ra ra-book mr-1"></i>
                  Save Spell
                </div>
              </div>


            </eq-window>

          </div>

          <!-- Preview Pane -->
          <div class="col-5">

            <!-- preview spell -->
            <eq-window
              style="margin-top: 30px; margin-right: 10px; width: auto;"
              v-if="previewSpellActive"
            >
              <eq-spell-preview
                :spell-data="spell"
              />
            </eq-window>

            <!-- icon selector -->
            <eq-window
              style="margin-top: 30px; margin-right: 10px; width: auto;"
              class="fade-in"
              v-if="iconSelectorActive"
            >
              <spell-icon-selector
                :selected-icon="spell.new_icon"
                :inputData.sync="spell.new_icon"
              />
            </eq-window>

            <!-- spell anim selector -->
            <eq-window
              style="margin-top: 30px; margin-right: 10px; width: auto;"
              class="fade-in"
              v-if="spellAnimSelectorActive"
            >

              <spell-animation-selector
                :selected-animation="spell.spellanim"
                :inputData.sync="spell.spellanim"
              />
            </eq-window>

            <!-- free id selector -->
            <eq-window
              title="Free Spell Ids"
              style="margin-top: 30px; margin-right: 10px; width: auto;"
              class="fade-in"
              v-if="freeIdSelectorActive"
            >

              <free-id-selector
                table-name="spells_new"
                id-name="id"
                name-label="name"
                :with-reserved="true"
                @input="spell.id = $event"
              />
            </eq-window>

            <!-- SPA Detail Pane -->
            <eq-window
              style="margin-top: 30px; margin-right: 10px; width: auto;"
              class="fade-in"
              v-if="spaDetailPaneActive"
            >
              <spell-spa-preview-pane
                :spa="spaPreviewNumber"
                :spell="spell"
                :effect-index="spaEffectIndex"
                v-if="spaPreviewNumber"
              />

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
import SpellIconSelector                                              from "./components/SpellIconSelector";
import SpellAnimationPreview                                          from "./components/SpellAnimationPreview";
import SpellAnimationViewer                                           from "../viewers/SpellAnimationViewer";
import SpellAnimationSelector                                         from "./components/SpellAnimationSelector";
import EqCheckbox                                                     from "../../components/eq-ui/EQCheckbox";
import {SpellsNewApi}                                                 from "../../app/api";
import {SpireApiClient}                                               from "../../app/api/spire-api-client";
import * as util                                                      from "util";
import SpellClassSelector                                             from "./components/SpellClassSelector";
import SpellDeitySelector                                             from "./components/SpellDeitySelector";
import FreeIdSelector                                                 from "../../components/tools/FreeIdSelector";
import SpellSpaPreviewPane                                            from "./components/SpellSpaPreviewPane";

const MILLISECONDS_BEFORE_WINDOW_RESET = 3000;

export default {
  name: "SpellEdit",
  components: {
    SpellSpaPreviewPane,
    FreeIdSelector,
    SpellDeitySelector,
    SpellClassSelector,
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


      // constants
      DB_SPELL_EFFECTS: DB_SPELL_EFFECTS,
      DB_SPA: DB_SPA,
      DB_SKILLS: DB_SKILLS,
      DB_SPELL_TARGETS: DB_SPELL_TARGETS,
      DB_SPELL_RESISTS: DB_SPELL_RESISTS,
      loaded: true,

      // preview / selectors
      previewSpellActive: true,
      iconSelectorActive: false,
      spellAnimSelectorActive: false,
      freeIdSelectorActive: false,
      spaDetailPaneActive: false,

      spaPreviewNumber: 0,
      spaEffectIndex: 0,

      lastResetTime: Date.now(),

      notification: "",
      error: "",
    }
  },
  watch: {
    '$route'() {
      // reset state vars when we navigate away
      this.notification = ""
      // reload
      this.load()
    }
  },
  async created() {

    setTimeout(() => {
      document.getElementById("spell-edit-card").removeEventListener('input', this.setFieldModified, true);
      document.getElementById("spell-edit-card").addEventListener('input', this.setFieldModified)
    }, 300)

    this.load()
  },
  methods: {

    getFieldDescription(field) {
      return Spells.getFieldDescription(field);
    },

    setFieldModified(evt) {
      // border: 2px #555555 solid !important;
      evt.target.style.setProperty('border-color', 'orange', 'important');
    },

    resetFieldEditedStatus() {
      // reset elements
      const elements = document.getElementById("spell-edit-card").querySelectorAll("input, select");
      elements.forEach((element) => {
        element.style.setProperty('border-color', '#555555', 'important');
      });
    },

    dismissNotification() {
      setTimeout(() => {
        this.notification = ""
      }, 5000)
    },

    sendNotification(message) {
      this.notification = message
      this.dismissNotification()
    },

    async saveSpell() {
      this.error        = ""
      this.notification = ""

      const api = (new SpellsNewApi(SpireApiClient.getOpenApiConfig()))
      api.updateSpellsNew({
        id: this.spell.id,
        spellsNew: this.spell
      }).then((result) => {
        if (result.status === 200) {
          this.sendNotification("Spell updated successfully!")
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

        const createRes = await api.createSpellsNew({
          spellsNew: this.spell
        })

        if (createRes.status === 200) {
          this.sendNotification("Created new Spell!")
          this.resetFieldEditedStatus()
        }
      })
    },

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
      this.freeIdSelectorActive    = false;
      this.spaDetailPaneActive     = false;
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
    drawSpaDetailPane(spa, index) {
      this.resetPreviewComponents()
      this.lastResetTime       = Date.now()
      this.spaDetailPaneActive = true
      this.spaPreviewNumber    = spa
      this.spaEffectIndex      = index
    },
    getTargetTypeColor(targetType) {
      return Spells.getTargetTypeColor(targetType);
    },
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
