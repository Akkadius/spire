<template>
  <div class="container-fluid">
    <div class="panel-body">
      <div class="panel panel-default">
        <div class="row">
          <div class="col-7">
            <eq-window style="margin-top: 30px" title="Edit Spell">

              <eq-tabs v-if="spell">
                <eq-tab name="Basic" :selected="true">

                  Id
                  <b-form-input v-model="spell.id"/>
                  Name
                  <b-form-input v-model="spell.name"/>
                  You Cast
                  <b-form-input v-model="spell.you_cast"/>
                  Other Casts
                  <b-form-input v-model="spell.other_casts"/>
                  Cast On You
                  <b-form-input v-model="spell.cast_on_you"/>
                  Cast On Other
                  <b-form-input v-model="spell.cast_on_other"/>
                  Spell Fades
                  <b-form-input v-model="spell.spell_fades"/>
                  ID File
                  <b-form-input v-model="spell.player_1"/>


                </eq-tab>
                <eq-tab name="Effects">

                  <h4 class="eq-header">Effects</h4>
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

                  <h4 class="eq-header mt-3">Added Effects</h4>

                  Push Up
                  <b-form-input v-model="spell.pushback"/>
                  Push Back
                  <b-form-input v-model="spell.pushup"/>
                  Hate Modifier
                  <b-form-input v-model="spell.bonushate"/>
                  Spell Hate Given
                  <b-form-input v-model="spell.hate_added"/>
                  No Detrimental Spell Aggro
                  <b-form-input v-model="spell.field_198"/>
                  Max Critical Chance
                  <b-form-input v-model="spell.field_217"/>
                  Viral Range
                  <b-form-input v-model="spell.viral_range"/>
                  Viral Targets
                  <b-form-input v-model="spell.viral_targets"/>
                  Viral Timer
                  <b-form-input v-model="spell.viral_timer"/>
                  Nimbus Type
                  <b-form-input v-model="spell.nimbuseffect"/>
                  Max Hits Type
                  <b-form-input v-model="spell.numhitstype"/>
                  Max Hits Allowed
                  <b-form-input v-model="spell.numhits"/>
                  Recourse Spell ID
                  <b-form-input v-model="spell.recourse_link"/>

                  <h4 class="eq-header mt-3">Focus</h4>

                  Not Focusable
                  <b-form-input v-model="spell.not_extendable"/>
                  Do not allow Heal or Dmg Item Mods (Maxtargets)
                  <b-form-input v-model="spell.maxtargets"/>
                  Song Base Effect Cap
                  <b-form-input v-model="spell.songcap"/>

                  <h4 class="eq-header mt-3">Spell Group</h4>

                  Spell Group
                  <b-form-input v-model="spell.spellgroup"/>
                  Rank
                  <b-form-input v-model="spell.rank"/>

                </eq-tab>
                <eq-tab name="General">
                  Skill
                  <b-form-select v-model="spell.skill" class="mb-3" v-if="DB_SKILLS">
                    <b-form-select-option :value="parseInt(id)" v-for="(skill, id) in DB_SKILLS">{{ id }}) {{ skill }}
                    </b-form-select-option>
                  </b-form-select>

                  Good Effect
                  <b-form-input v-model="spell.good_effect"/>
                  Mana
                  <b-form-input v-model="spell.mana"/>
                  Endurance Cost
                  <b-form-input v-model="spell.endur_cost"/>
                  Endurance Upkeep
                  <b-form-input v-model="spell.endur_upkeep"/>
                  Use Discipline Window
                  <b-form-input v-model="spell.is_discipline"/>
                  pcnpc_only_flag
                  <b-form-input v-model="spell.pcnpc_only_flag"/>
                  Teleport Zone / Pet DbaseID / ItemGraphic for Bolt Spells
                  <b-form-input v-model="spell.teleport_zone"/>
                </eq-tab>
                <eq-tab name="Restrictions">
                  Target Restriction
                  <b-form-input v-model="spell.cast_restriction"/>
                  Caster Restriction
                  <b-form-input v-model="spell.field_220"/>
                  Must Be In Combat
                  <b-form-input v-model="spell.in_combat"/>
                  Must Be Out of Combat
                  <b-form-input v-model="spell.outof_combat"/>
                  Zone Type (select)
                  <b-form-input v-model="spell.zonetype"/>
                  Environment Type
                  <b-form-input v-model="spell.environment_type"/>
                  Time of Day
                  <b-form-input v-model="spell.time_of_day"/>
                  Only During Fast Regen
                  <b-form-input v-model="spell.allowrest"/>
                  Cancel On Sit
                  <b-form-input v-model="spell.disallow_sit"/>
                  Must be Sneaking
                  <b-form-input v-model="spell.sneaking"/>
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
                  <b-form-input v-model="spell.uninterruptable"/>
                  Fizzle Adjustment
                  <b-form-input v-model="spell.basediff"/>
                  Cast Not Standing (Can Cast from Sitting position, Can cast on invulnerable Targets, Can not be
                  interrupted by SE_InterruptCasting)
                  <b-form-input v-model="spell.cast_not_standing"/>
                </eq-tab>
                <eq-tab name="Buffing">
                  Buff Duration
                  <b-form-input v-model="spell.buffduration"/>
                  Duration Formula
                  <b-form-input v-model="spell.buffdurationformula"/>
                  Can Not Dispell
                  <b-form-input v-model="spell.nodispell"/>
                  Can Not Click Off
                  <b-form-input v-model="spell.field_232"/>
                  Persist After Death
                  <b-form-input v-model="spell.persistdeath"/>
                  Suspendable
                  <b-form-input v-model="spell.suspendable"/>
                  Short Duration Buff Box
                  <b-form-input v-model="spell.short_buff_box"/>
                  Can MGB
                  <b-form-input v-model="spell.can_mgb"/>
                  No Buff Block
                  <b-form-input v-model="spell.no_block"/>
                  DOT not stackable
                  <b-form-input v-model="spell.dot_stacking_exempt"/>
                  PVP Duration
                  <b-form-input v-model="spell.field_181"/>
                  PVP Duration Cap
                  <b-form-input v-model="spell.field_182"/>
                </eq-tab>
                <eq-tab name="Range">
                  Target Type
                  <b-form-select v-model="spell.targettype" class="mb-3" v-if="DB_SPELL_TARGETS">
                    <b-form-select-option :value="parseInt(id)" v-for="(skill, id) in DB_SPELL_TARGETS">{{ id }})
                      {{ skill }}
                    </b-form-select-option>
                  </b-form-select>

                  Range
                  <b-form-input v-model="spell.range"/>
                  AOE Range
                  <b-form-input v-model="spell.aoerange"/>
                  AOE Rain Waves
                  <b-form-input v-model="spell.ae_duration"/>
                  AOE Max Targets
                  <b-form-input v-model="spell.aemaxtargets"/>
                  Min Range
                  <b-form-input v-model="spell.min_range"/>
                  Min Distance for Modifier
                  <b-form-input v-model="spell.min_dist"/>
                  Min Distance Modifier
                  <b-form-input v-model="spell.min_dist_mod"/>
                  Max Distance for Modifier
                  <b-form-input v-model="spell.max_dist"/>
                  Max Distance Modifier
                  <b-form-input v-model="spell.max_dist_mod"/>
                  Cone Angle Start
                  <b-form-input v-model="spell.cone_start_angle"/>
                  Cone Angle End
                  <b-form-input v-model="spell.cone_stop_angle"/>
                  NPC does not need Light of Sight to cast
                  <b-form-input v-model="spell.npc_no_los"/>
                </eq-tab>
                <eq-tab name="Resist">
                  Resist Type
                  <b-form-input v-model="spell.resisttype"/>
                  Unresistable
                  <b-form-input v-model="spell.field_209"/>
                  Resist Diff
                  <b-form-input v-model="spell.resist_diff"/>
                  No Partial Resists
                  <b-form-input v-model="spell.no_partial_resist"/>
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
            <eq-window style="margin-top: 30px; margin-right: 10px; width: auto;">
              <eq-spell-preview
                :spell-data="spell"/>
            </eq-window>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import EqWindowFancy                                from "../../components/eq-ui/EQWindowFancy";
import EqWindow                                     from "../../components/eq-ui/EQWindow";
import EqTabs                                       from "../../components/eq-ui/EQTabs";
import EqTab                                        from "../../components/eq-ui/EQTab";
import EqSpellPreview                               from "../../components/eq-ui/EQSpellCardPreview";
import {Spells}                                     from "../../app/spells";
import {DB_SPA, DB_SPELL_EFFECTS, DB_SPELL_TARGETS} from "../../app/constants/eq-spell-constants";
import {DB_SKILLS}                                  from "../../app/constants/eq-skill-constants";

export default {
  name: "SpellEdit",
  components: { EqSpellPreview, EqTab, EqTabs, EqWindow, EqWindowFancy },
  data() {
    return {
      spell: null,
      DB_SPELL_EFFECTS: DB_SPELL_EFFECTS,
      DB_SPA: DB_SPA,
      DB_SKILLS: DB_SKILLS,
      DB_SPELL_TARGETS: DB_SPELL_TARGETS,
      loaded: true,
    }
  },
  watch: {
    '$route'() {
      this.load()
    }
  },
  async created() {
    this.load()
    console.log("created")
  },
  methods: {
    load() {
      if (this.$route.params.id > 0) {
        Spells.getSpell(this.$route.params.id).then(result => {
          console.log(result)
          this.spell = result
        })
      }
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

</style>
