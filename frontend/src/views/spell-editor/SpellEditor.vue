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

                  <!-- Spell Group -->
                  <div class="row">
                    <div class="col-2">
                      Spell Group
                      <b-form-input v-model.number="spell.spellgroup"/>
                    </div>
                    <div class="col-2">
                      Spell Group Rank
                      <b-form-input v-model.number="spell.rank"/>
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

                <eq-tab name="Effects+" class="minified-inputs">
                  <div class="row" v-for="field in
                     [
                       {
                         description: '(Knockback) Push Back',
                         field: 'pushback'
                       },
                       {
                         description: '(Knockback) Push Up',
                         field: 'pushup'
                       },
                       {
                         description: '(Recourse) Recourse ID',
                         field: 'recourse_link'
                       },
                       {
                         description: '(Hate) Hate Modifier',
                         field: 'bonushate'
                       },
                       {
                         description: '(Hate) Spell Hate Given',
                         field: 'hate_added'
                       },
                       {
                         description: '(Hate) No Detrimental Spell Aggro',
                         field: 'field_198',
                         bool: true
                       },
                       {
                         description: '(Viral) Viral Range',
                         field: 'viral_range'
                       },
                       {
                         description: '(Viral) Viral Targets',
                         field: 'viral_targets'
                       },
                       {
                         description: '(Viral) Viral Timer',
                         field: 'viral_timer'
                       },
                       {
                         description: '(Focus) Max Targets',
                         field: 'maxtargets'
                       },
                       {
                         description: '(Focus) Song Base Effect Cap',
                         field: 'songcap'
                       },
                       {
                         description: '(Focus) Not Focusable',
                         field: 'not_extendable',
                         bool: true,
                       },
                       {
                         description: 'Max Critical Chance',
                         field: 'field_217'
                       },
                       {
                         description: 'Nimbus Type',
                         field: 'nimbuseffect'
                       },
                       {
                         description: 'Teleport Zone / Pet DB ID / Item Graphic for Bolt Spells',
                         field: 'teleport_zone',
                         text: true,
                       },
                     ]"
                  >
                    <div class="col-6 text-right p-0 m-0 mr-3 mt-3" v-if="field.bool">
                      {{ field.description }}
                    </div>
                    <div class="col-6 text-right p-0 m-0 mr-3" v-if="!field.bool" style="margin-top: 10px !important">
                      {{ field.description }}
                    </div>
                    <div class="col-3 text-left p-0 mt-1">

                      <!-- checkbox -->
                      <eq-checkbox
                        v-b-tooltip.hover.v-dark.right :title="getFieldDescription(field.field)"
                        class="mb-1 mt-3 d-inline-block"
                        :true-value="(typeof field.true !== 'undefined' ? field.true : 1)"
                        :false-value="(typeof field.false !== 'undefined' ? field.false : 0)"
                        v-model.number="spell[field.field]"
                        @input="spell[field.field] = $event"
                        v-if="field.bool"
                      />

                      <!-- input number -->
                      <b-form-input
                        v-if="!field.selectData && !field.bool && !field.text"
                        :id="field.field"
                        v-model.number="spell[field.field]"
                        class="m-0 mt-1"
                        v-b-tooltip.hover.v-dark.right :title="getFieldDescription(field.field)"
                        :style="(spell[field.field] === 0 ? 'opacity: .5' : '')"
                      />

                      <!-- input text -->
                      <b-form-input
                        v-if="!field.selectData && !field.bool && field.text"
                        :id="field.field"
                        v-model.number="spell[field.field]"
                        class="m-0 mt-1"
                        v-b-tooltip.hover.v-dark.right :title="getFieldDescription(field.field)"
                        :style="(spell[field.field] === '' ? 'opacity: .5' : '')"
                      />

                      <!-- select -->
                      <select
                        v-model.number="spell[field.field]"
                        class="form-control m-0 mt-1"
                        v-if="field.selectData"
                        v-b-tooltip.hover.v-dark.right :title="getFieldDescription(field.field)"
                        :style="(spell[field.field] <= 0 ? 'opacity: .5' : '')"
                      >
                        <option
                          v-for="(description, index) in field.selectData"
                          :key="index"
                          :value="parseInt(index)"
                        >
                          {{ index }}) {{ description }}
                        </option>
                      </select>
                    </div>
                  </div>

                </eq-tab>

                <eq-tab name="General" class="minified-inputs">
                  <div class="row">
                    <div class="col-12">
                      <div
                        class="row" v-for="field in
                         [
                           {
                             description: 'Is Discipline',
                             field: 'is_discipline',
                             bool: true
                           },
                           {
                             description: 'Skills',
                             field: 'skill',
                             selectData: DB_SKILLS
                           },
                           {
                             description: 'Good Effect',
                             field: 'good_effect',
                           },
                           {
                             description: 'Mana Cost',
                             field: 'mana',
                           },
                           {
                             description: 'Endurance Cost',
                             field: 'endur_cost',
                           },
                           {
                             description: 'Endurance Upkeep',
                             field: 'endur_upkeep',
                           },
                         ]"
                      >
                        <div class="col-6 text-right p-0 m-0 mr-3" v-if="field.bool">
                          {{ field.description }}
                        </div>
                        <div
                          class="col-6 text-right p-0 m-0 mr-3"
                          v-if="!field.bool"
                          style="margin-top: 10px !important"
                        >
                          {{ field.description }}
                        </div>
                        <div class="col-3 text-left p-0 mt-1">

                          <!-- checkbox -->
                          <eq-checkbox
                            v-b-tooltip.hover.v-dark.right :title="getFieldDescription(field.field)"
                            class="mb-2 d-inline-block"
                            :true-value="(typeof field.true !== 'undefined' ? field.true : 1)"
                            :false-value="(typeof field.false !== 'undefined' ? field.false : 0)"
                            v-model.number="spell[field.field]"
                            @input="spell[field.field] = $event"
                            v-if="field.bool"
                          />

                          <!-- input -->
                          <b-form-input
                            v-if="!field.selectData && !field.bool"
                            :id="field.field"
                            v-b-tooltip.hover.v-dark.right :title="getFieldDescription(field.field)"
                            v-model.number="spell[field.field]"
                            class="m-0 mt-1"
                            :style="(spell[field.field] <= 0 ? 'opacity: .5' : '')"
                          />

                          <!-- select -->
                          <select
                            v-model.number="spell[field.field]"
                            class="form-control m-0 mt-1"
                            v-if="field.selectData"
                            v-b-tooltip.hover.v-dark.right :title="getFieldDescription(field.field)"
                            :style="(spell[field.field] <= 0 ? 'opacity: .5' : '')"
                          >
                            <option
                              v-for="(description, index) in field.selectData"
                              :key="index"
                              :value="parseInt(index)"
                            >
                              {{ index }}) {{ description }}
                            </option>
                          </select>
                        </div>
                      </div>

                    </div>
                  </div>

                </eq-tab>

                <eq-tab name="Restrictions" class="minified-inputs">
                  <div class="row">
                    <div class="col-12">
                      <div
                        class="row" v-for="field in
                         [
                           {
                             description: 'Can Cast out of Combat',
                             field: 'outof_combat',
                             bool: true
                           },
                           {
                             description: 'Can Cast in Combat',
                             field: 'in_combat',
                             bool: true
                           },
                           {
                             description: 'Only During Fast Regen',
                             field: 'allowrest',
                             bool: true
                           },
                           {
                             description: 'Cancel on Sit',
                             field: 'disallow_sit',
                             bool: true
                           },
                           {
                             description: 'Must be Sneaking',
                             field: 'sneaking',
                             bool: true
                           },
                           {
                             description: 'Target Restriction',
                             field: 'cast_restriction',
                             selectData: DB_SPELL_TARGET_RESTRICTION,
                           },
                           {
                             description: 'Caster Restriction',
                             field: 'field_220',
                             selectData: DB_SPELL_TARGET_RESTRICTION,
                           },
                           {
                             description: 'Zone Type',
                             field: 'zonetype',
                             selectData: DB_SPELL_ZONE_TYPE,
                           },
                           {
                             description: 'Environment Type',
                             field: 'environment_type',
                           },
                           {
                             description: 'Time of Day',
                             field: 'time_of_day',
                           },
                           {
                             description: 'PC or NPC Only',
                             field: 'pcnpc_only_flag',
                             selectData: DB_PC_NPC_ONLY_FLAG,
                           },
                         ]"
                      >
                        <div class="col-6 text-right p-0 m-0 mr-3" v-if="field.bool">
                          {{ field.description }}
                        </div>
                        <div
                          class="col-6 text-right p-0 m-0 mr-3"
                          v-if="!field.bool"
                          style="margin-top: 10px !important"
                        >
                          {{ field.description }}
                        </div>
                        <div class="col-3 text-left p-0 mt-1">

                          <!-- checkbox -->
                          <eq-checkbox
                            v-b-tooltip.hover.v-dark.right :title="getFieldDescription(field.field)"
                            class="mb-2 d-inline-block"
                            :true-value="(typeof field.true !== 'undefined' ? field.true : 1)"
                            :false-value="(typeof field.false !== 'undefined' ? field.false : 0)"
                            v-model.number="spell[field.field]"
                            @input="spell[field.field] = $event"
                            v-if="field.bool"
                          />

                          <!-- input -->
                          <b-form-input
                            v-if="!field.selectData && !field.bool"
                            :id="field.field"
                            v-model.number="spell[field.field]"
                            class="m-0 mt-1"
                            v-b-tooltip.hover.v-dark.right :title="getFieldDescription(field.field)"
                            :style="(spell[field.field] <= 0 ? 'opacity: .5' : '')"
                          />

                          <!-- select -->
                          <select
                            v-model.number="spell[field.field]"
                            class="form-control m-0 mt-1"
                            v-if="field.selectData"
                            :style="(spell[field.field] <= 0 ? 'opacity: .5' : '')"
                            v-b-tooltip.hover.v-dark.right :title="getFieldDescription(field.field)"
                          >
                            <option
                              v-for="(description, index) in field.selectData"
                              :key="index"
                              :value="parseInt(index)"
                            >
                              {{ index }}) {{ description }}
                            </option>
                          </select>
                        </div>
                      </div>

                    </div>
                  </div>
                </eq-tab>

                <eq-tab name="Range" class="minified-inputs">

                  <div class="row" v-for="field in
                     [
                       {
                         description: 'Spell Range',
                         field: 'range'
                       },
                       {
                         description: 'Target Type',
                         field: 'targettype',
                         selectData: DB_SPELL_TARGETS,
                       },
                       {
                         description: 'NPC Line of Sight Not Required to Cast',
                         field: 'npc_no_los',
                         bool: true
                       },
                       {
                         description: 'AOE Range',
                         field: 'aoerange'
                       },
                       {
                         description: 'AOE Rain Waves',
                         field: 'ae_duration'
                       },
                       {
                         description: 'AOE Max Targets',
                         field: 'aemaxtargets'
                       },
                       {
                         description: 'Min Range',
                         field: 'min_range'
                       },
                       {
                         description: 'Min Distance for Mod',
                         field: 'min_dist'
                       },
                       {
                         description: 'Min Distance Mod',
                         field: 'min_dist_mod'
                       },
                       {
                         description: 'Max Distance for Mod',
                         field: 'max_dist'
                       },
                       {
                         description: 'Max Distance Mod',
                         field: 'max_dist_mod'
                       },
                       {
                         description: 'Max Hits Type',
                         field: 'numhitstype'
                       },
                       {
                         description: 'Max Hits Allowed',
                         field: 'numhits'
                       },
                       {
                         description: 'Cone Angle Start',
                         field: 'cone_start_angle'
                       },
                       {
                         description: 'Cone Angle End',
                         field: 'cone_stop_angle'
                       },
                     ]"
                  >
                    <div class="col-6 text-right p-0 m-0 mr-3 mt-3" v-if="field.bool">
                      {{ field.description }}
                    </div>
                    <div class="col-6 text-right p-0 m-0 mr-3" v-if="!field.bool" style="margin-top: 10px !important">
                      {{ field.description }}
                    </div>
                    <div class="col-3 text-left p-0 mt-1">

                      <!-- checkbox -->
                      <eq-checkbox
                        v-b-tooltip.hover.v-dark.right :title="getFieldDescription(field.field)"
                        class="mb-1 mt-3 d-inline-block"
                        :true-value="(typeof field.true !== 'undefined' ? field.true : 1)"
                        :false-value="(typeof field.false !== 'undefined' ? field.false : 0)"
                        v-model.number="spell[field.field]"
                        @input="spell[field.field] = $event"
                        v-if="field.bool"
                      />

                      <!-- input number -->
                      <b-form-input
                        v-if="!field.selectData && !field.bool && !field.text"
                        :id="field.field"
                        v-model.number="spell[field.field]"
                        class="m-0 mt-1"
                        v-b-tooltip.hover.v-dark.right :title="getFieldDescription(field.field)"
                        :style="(spell[field.field] === 0 ? 'opacity: .5' : '')"
                      />

                      <!-- input text -->
                      <b-form-input
                        v-if="!field.selectData && !field.bool && field.text"
                        :id="field.field"
                        v-model.number="spell[field.field]"
                        class="m-0 mt-1"
                        v-b-tooltip.hover.v-dark.right :title="getFieldDescription(field.field)"
                        :style="(spell[field.field] === '' ? 'opacity: .5' : '')"
                      />

                      <!-- select -->
                      <select
                        v-model.number="spell[field.field]"
                        class="form-control m-0 mt-1"
                        v-if="field.selectData"
                        v-b-tooltip.hover.v-dark.right :title="getFieldDescription(field.field)"
                        :style="(spell[field.field] <= 0 ? 'opacity: .5' : '')"
                      >
                        <option
                          v-for="(description, index) in field.selectData"
                          :key="index"
                          :value="parseInt(index)"
                        >
                          {{ index }}) {{ description }}
                        </option>
                      </select>
                    </div>
                  </div>

                </eq-tab>


                <eq-tab name="Casting" class="minified-inputs">
                  <div
                    class="row" v-for="field in
                     [
                       {
                         description: 'Uninterruptable',
                         field: 'uninterruptable',
                         bool: true
                       },
                       {
                         description: 'Cast Time',
                         field: 'cast_time'
                       },
                       {
                         description: 'Recovery Time',
                         field: 'recovery_time'
                       },
                       {
                         description: 'Recast Time',
                         field: 'recast_time'
                       },
                       {
                         description: 'Timer Index (???)',
                         field: 'endur_timer_index'
                       },
                       {
                         description: 'Fizzle Adjustment',
                         field: 'basediff'
                       },
                       {
                         description: 'Cast Not Standing',
                         field: 'cast_not_standing'
                       },
                     ]"
                  >
                    <div class="col-6 text-right p-0 m-0 mr-3" v-if="field.bool">
                      {{ field.description }}
                    </div>
                    <div class="col-6 text-right p-0 m-0 mr-3" v-if="!field.bool" style="margin-top: 10px !important">
                      {{ field.description }}
                    </div>
                    <div class="col-3 text-left p-0 mt-1">

                      <!-- checkbox -->
                      <eq-checkbox
                        v-b-tooltip.hover.v-dark.right :title="getFieldDescription(field.field)"
                        class="mb-2 d-inline-block"
                        :true-value="(typeof field.true !== 'undefined' ? field.true : 1)"
                        :false-value="(typeof field.false !== 'undefined' ? field.false : 0)"
                        v-model.number="spell[field.field]"
                        @input="spell[field.field] = $event"
                        v-if="field.bool"
                      />

                      <!-- input -->
                      <b-form-input
                        v-if="!field.selectData && !field.bool"
                        :id="field.field"
                        v-b-tooltip.hover.v-dark.right :title="getFieldDescription(field.field)"
                        v-model.number="spell[field.field]"
                        class="m-0 mt-1"
                        :style="(spell[field.field] <= 0 ? 'opacity: .5' : '')"
                      />

                      <!-- select -->
                      <select
                        v-model.number="spell[field.field]"
                        class="form-control m-0 mt-1"
                        v-if="field.selectData"
                        v-b-tooltip.hover.v-dark.right :title="getFieldDescription(field.field)"
                        :style="(spell[field.field] <= 0 ? 'opacity: .5' : '')"
                      >
                        <option
                          v-for="(description, index) in field.selectData"
                          :key="index"
                          :value="parseInt(index)"
                        >
                          {{ index }}) {{ description }}
                        </option>
                      </select>
                    </div>
                  </div>

                </eq-tab>
                <eq-tab name="Buffing" class="minified-inputs">
                  <div
                    class="row" v-for="field in
                     [
                       {
                         description: 'Can Not Dispell',
                         field: 'nodispell',
                         bool: true
                       },
                       {
                         description: 'Can Not Click Off',
                         field: 'field_232',
                         bool: true
                       },
                       {
                         description: 'Persist After Death',
                         field: 'persistdeath',
                         bool: true
                       },
                       {
                         description: 'Suspendable',
                         field: 'suspendable',
                         bool: true
                       },
                       {
                         description: 'Can MGB',
                         field: 'can_mgb',
                         bool: true
                       },
                       {
                         description: 'Appear In Short Buff Box',
                         field: 'short_buff_box',
                         bool: true
                       },
                       {
                         description: 'No Buff Block',
                         field: 'no_block',
                         bool: true
                       },
                       {
                         description: 'DOT Not Stackable',
                         field: 'dot_stacking_exempt',
                         bool: true
                       },
                       {
                         description: 'Buff Duration',
                         field: 'buffduration'
                       },
                       {
                         description: 'Buff Duration Formula',
                         field: 'buffdurationformula'
                       },
                       {
                         description: 'PVP Duration',
                         field: 'pvp_duration'
                       },
                       {
                         description: 'PVP Duration Cap',
                         field: 'pvp_duration_cap'
                       },
                     ]"
                  >
                    <div class="col-6 text-right p-0 m-0 mr-3" v-if="field.bool">
                      {{ field.description }}
                    </div>
                    <div class="col-6 text-right p-0 m-0 mr-3" v-if="!field.bool" style="margin-top: 10px !important">
                      {{ field.description }}
                    </div>
                    <div class="col-3 text-left p-0 mt-1">

                      <!-- checkbox -->
                      <eq-checkbox
                        v-b-tooltip.hover.v-dark.right :title="getFieldDescription(field.field)"
                        class="mb-2 d-inline-block"
                        :true-value="(typeof field.true !== 'undefined' ? field.true : 1)"
                        :false-value="(typeof field.false !== 'undefined' ? field.false : 0)"
                        v-model.number="spell[field.field]"
                        @input="spell[field.field] = $event"
                        v-if="field.bool"
                      />

                      <!-- input -->
                      <b-form-input
                        v-if="!field.selectData && !field.bool"
                        :id="field.field"
                        v-model.number="spell[field.field]"
                        class="m-0 mt-1"
                        v-b-tooltip.hover.v-dark.right :title="getFieldDescription(field.field)"
                        :style="(spell[field.field] <= 0 ? 'opacity: .5' : '')"
                      />

                      <!-- select -->
                      <select
                        v-model.number="spell[field.field]"
                        class="form-control m-0 mt-1"
                        v-if="field.selectData"
                        v-b-tooltip.hover.v-dark.right :title="getFieldDescription(field.field)"
                        :style="(spell[field.field] <= 0 ? 'opacity: .5' : '')"
                      >
                        <option
                          v-for="(description, index) in field.selectData"
                          :key="index"
                          :value="parseInt(index)"
                        >
                          {{ index }}) {{ description }}
                        </option>
                      </select>
                    </div>
                  </div>
                </eq-tab>
                <eq-tab name="Resist" class="minified-inputs">
                  <div
                    class="row" v-for="field in
                     [
                       {
                         description: 'Unresistable',
                         field: 'field_209',
                         bool: true
                       },
                       {
                         description: 'No Partial Resists',
                         field: 'no_partial_resist',
                         bool: true
                       },
                       {
                         description: 'Resist Type',
                         field: 'resisttype',
                         selectData: DB_SPELL_RESISTS,
                       },
                       {
                         description: 'Resist Diff',
                         field: 'resist_diff'
                       },
                       {
                         description: 'PVP Resist Mod',
                         field: 'pvpresistbase'
                       },
                       {
                         description: 'PVP Resist Per Level',
                         field: 'pvpresistcalc'
                       },
                       {
                         description: 'PVP Resist Cap',
                         field: 'pvpresistcap'
                       },
                       {
                         description: 'Resist Chance Limits: Min Chance',
                         field: 'min_resist'
                       },
                       {
                         description: 'Resist Chance Limits: Max Chance',
                         field: 'max_resist'
                       },
                       {
                         description: 'Reflectable',
                         field: 'reflectable'
                       },
                       {
                         description: 'Feedbackable',
                         field: 'field_160'
                       },
                     ]"
                  >
                    <div class="col-6 text-right p-0 m-0 mr-3" v-if="field.bool">
                      {{ field.description }}
                    </div>
                    <div class="col-6 text-right p-0 m-0 mr-3" v-if="!field.bool" style="margin-top: 10px !important">
                      {{ field.description }}
                    </div>
                    <div class="col-3 text-left p-0 mt-1">

                      <!-- checkbox -->
                      <eq-checkbox
                        v-b-tooltip.hover.v-dark.right :title="getFieldDescription(field.field)"
                        class="mb-2 d-inline-block"
                        :true-value="(typeof field.true !== 'undefined' ? field.true : 1)"
                        :false-value="(typeof field.false !== 'undefined' ? field.false : 0)"
                        v-model.number="spell[field.field]"
                        @input="spell[field.field] = $event"
                        v-if="field.bool"
                      />

                      <!-- input -->
                      <b-form-input
                        v-if="!field.selectData && !field.bool"
                        :id="field.field"
                        v-model.number="spell[field.field]"
                        class="m-0 mt-1"
                        v-b-tooltip.hover.v-dark.right :title="getFieldDescription(field.field)"
                        :style="(spell[field.field] === 0 ? 'opacity: .5' : '')"
                      />

                      <!-- select -->
                      <select
                        v-model.number="spell[field.field]"
                        class="form-control m-0 mt-1"
                        v-if="field.selectData"
                        v-b-tooltip.hover.v-dark.right :title="getFieldDescription(field.field)"
                        :style="(spell[field.field] <= 0 ? 'opacity: .5' : '')"
                      >
                        <option
                          v-for="(description, index) in field.selectData"
                          :key="index"
                          :value="parseInt(index)"
                        >
                          {{ index }}) {{ description }}
                        </option>
                      </select>
                    </div>
                  </div>

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
                v-if="spaPreviewNumber >= 0"
              />

            </eq-window>

          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import EqWindowFancy          from "../../components/eq-ui/EQWindowFancy";
import EqWindow               from "../../components/eq-ui/EQWindow";
import EqTabs                 from "../../components/eq-ui/EQTabs";
import EqTab                  from "../../components/eq-ui/EQTab";
import EqSpellPreview         from "../../components/eq-ui/EQSpellCardPreview";
import {Spells}               from "../../app/spells";
import {
  DB_PC_NPC_ONLY_FLAG,
  DB_SPA,
  DB_SPELL_EFFECTS,
  DB_SPELL_RESISTS,
  DB_SPELL_TARGET_RESTRICTION,
  DB_SPELL_TARGETS,
  DB_SPELL_ZONE_TYPE
}                             from "../../app/constants/eq-spell-constants";
import {DB_SKILLS}            from "../../app/constants/eq-skill-constants";
import {App}                  from "../../constants/app";
import SpellIconSelector      from "./components/SpellIconSelector";
import SpellAnimationPreview  from "./components/SpellAnimationPreview";
import SpellAnimationViewer   from "../viewers/SpellAnimationViewer";
import SpellAnimationSelector from "./components/SpellAnimationSelector";
import EqCheckbox             from "../../components/eq-ui/EQCheckbox";
import {SpellsNewApi}         from "../../app/api";
import {SpireApiClient}       from "../../app/api/spire-api-client";
import SpellClassSelector     from "./components/SpellClassSelector";
import SpellDeitySelector     from "./components/SpellDeitySelector";
import FreeIdSelector         from "../../components/tools/FreeIdSelector";
import SpellSpaPreviewPane    from "./components/SpellSpaPreviewPane";

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
      DB_SPELL_TARGET_RESTRICTION: DB_SPELL_TARGET_RESTRICTION,
      DB_SPELL_ZONE_TYPE: DB_SPELL_ZONE_TYPE,
      DB_PC_NPC_ONLY_FLAG: DB_PC_NPC_ONLY_FLAG,
      loaded: true,

      // preview / selectors
      previewSpellActive: true,
      iconSelectorActive: false,
      spellAnimSelectorActive: false,
      freeIdSelectorActive: false,
      spaDetailPaneActive: false,

      spaPreviewNumber: -1,
      spaEffectIndex: -1,

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
