<template>
  <div>
    <div class="row">
      <div class="col-xl-7 col-lg-12">
        <eq-window style="" title="Edit Spell">

          <div
            v-if="notification"
            :class="'text-center mt-2 btn-xs eq-header fade-in'"
            style="width: 100%; font-size: 30px"
            @click="notification = ''"
          >
            <i class="ra ra-book mr-1"></i>
            {{ notification }}
          </div>

          <b-alert show dismissable variant="danger" v-if="error" class="mt-2">
            <i class="fa fa-warning"></i> {{ error }}
          </b-alert>

          <!-- Loader -->
          <app-loader :is-loading="!spell" padding="5"/>
          <div v-if="!spell" class="mt-3 text-center">
            <loader-fake-progress/>
          </div>

          <eq-tabs
            v-if="spell && spell.id >= 0"
            id="spell-edit-card"
            class="spell-edit-card"
            @mouseover.native="previewSpell(false)"
          >
            <eq-tab
              name="Basic"
              :selected="zeroStateSelected"
            >
              <div class="row">
                <div
                  class="col-2"
                  @click="drawFreeIdSelector(true)"
                  @mouseover="drawFreeIdSelector(false)"
                >
                  Id
                  <b-form-input id="id" v-model.number="spell.id"/>
                </div>
                <div class="col-7">
                  Name
                  <b-form-input
                    :value="spell.name" @change="v => spell.name = v"
                  />
                </div>

                <div
                  class="col-2"
                  @click="drawIconSelector(true)"
                >
                  Icon
                  <b-form-input id="icon" v-model.number="spell.new_icon"/>
                </div>

                <div
                  class="col-1" v-if="spell.new_icon > 0"
                  style="margin-top: 7px"
                  @click="drawIconSelector(true)"
                >

                      <span
                        :style="'width: 40px; height: 40px; border: 1px solid ' + getTargetTypeColor(this.spell['targettype']) + '; border-radius: 7px; display: inline-block'"
                        :class="'spell-' + spell.new_icon + '-40 mt-2'"
                      />

                </div>

              </div>

              <div class="row">
                <div class="col-6 pl-0 mr-0 mt-4 text-center">
                  <spell-class-selector :spell="spell" @input="spell = $event"/>

                  <b-button
                    @click="none()"
                    size="sm"
                    variant="outline-warning"
                    class="mt-1"
                  >
                    <i class="fa fa-remove"></i> None
                  </b-button>
                </div>
                <div class="col-6">

                  <div
                    class="row"
                    @click="drawSpellAnimationSelector(true)"
                  >

                    <div class="col-12 text-center">
                      Spell Animation

                      <spell-animation-preview
                        class="mt-1"
                        :id="spell.spellanim"
                      />
                      <b-form-input id="spellanim" v-model.number="spell.spellanim" class="col-12 mt-3"/>
                    </div>

                  </div>

                </div>
              </div>

              <div class="row">
                <div class="col-6">
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
                  class="col-6"
                  style="text-align: center"
                >

                  <div class="row">

                    <div
                      class="col-6 text-center"
                      @click="castingAnimField = 'casting_anim'; drawCastingAnimationSelector()"
                    >
                      Casting Animation

                      <spell-casting-animation-preview :id="spell.casting_anim"/>
                      <b-form-input id="casting_anim" v-model.number="spell.casting_anim" class="col-12"/>
                    </div>

                    <div
                      class="col-6 text-center"
                      @click="castingAnimField = 'target_anim'; drawCastingAnimationSelector()"
                    >
                      Target Animation

                      <spell-casting-animation-preview :id="spell.target_anim"/>
                      <b-form-input id="target_anim" v-model.number="spell.target_anim" class="col-12"/>
                    </div>

                  </div>

                </div>
              </div>

            </eq-tab>

            <eq-tab name="Effects" class="effect-tab">

              <div class="mb-3">
                <div class="row">
                  <div class="col-12 text-center">
                    <div class="btn-group text-center mb-3" role="group">
                      <b-button size="sm" variant="warning">Effect Slots</b-button>
                      <b-button
                        v-for="i in 12"
                        :key="i"
                        size="sm"
                        :disabled="spell['effectid_' + i] !== 254"
                        :variant="(visibleEffectSlots[i] ? 'warning' : 'outline-warning')"
                        @click="toggleVisibleEffectSlot(i)"
                      >{{ i }}
                      </b-button>
                    </div>
                  </div>
                </div>

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

                <b-input-group
                  v-for="i in 12"
                  :key="i"
                  style="margin-top: -1px"
                  v-if="visibleEffectSlots[i]"
                >
                  <template #prepend>
                    <b-input-group-text style="width: 40px; ">{{ i }}</b-input-group-text>
                  </template>

                  <b-form-select
                    :id="'effectid_' + i"
                    @mouseover.native="drawSpaDetailPane(spell['effectid_' + i], i)"
                    @change="getSpaDefaultValues(spell['effectid_' + i], i); drawSpaDetailPane(spell['effectid_' + i], i)"
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

                  <b-form-input
                    :id="'effect_base_value_' + i"
                    v-model.number="spell['effect_base_value_' + i]"
                    :class="getSpaSpellHighlights(spell['effectid_' + i], 'base')"
                    @click="processSpaFieldAction(i, spell['effectid_' + i], 'base')"
                  />
                  <b-form-input
                    :id="'effect_limit_value_' + i"
                    v-model.number="spell['effect_limit_value_' + i]"
                    :class="getSpaSpellHighlights(spell['effectid_' + i], 'limit')"
                    @click="processSpaFieldAction(i, spell['effectid_' + i], 'limit')"
                  />
                  <b-form-input
                    :id="'max_' + i"
                    v-model.number="spell['max_' + i]"
                    :class="getSpaSpellHighlights(spell['effectid_' + i], 'max')"
                    @click="processSpaFieldAction(i, spell['effectid_' + i], 'max')"
                  />

                  <b-form-select
                    v-model.number="spell['formula_' + i]"
                    :id="'formula_' + i"
                  >
                    <option
                      v-for="(description, index) in BASE_VALUE_FORMULAS"
                      :key="index"
                      :value="parseInt(index)"
                    >
                      {{ index }}) {{ description }}
                    </option>
                  </b-form-select>

                </b-input-group>

              </div>

              <div
                class="row minified-inputs fade-in" v-for="field in
                     [
                       {
                         description: teleportZoneFieldName,
                         field: 'teleport_zone',
                         text: true,
                         showIf: teleportZoneFieldName !== ''
                       },
                     ]"
                v-if="typeof field.showIf === 'undefined' || (typeof field.showIf !== 'undefined' && field.showIf) || showAllFields"
                @click="processClickInputTrigger(field.field)"
              >
                <div class="col-6 text-right p-0 m-0 mr-3 mt-3" v-if="field.bool">
                  <span v-if="field.category" class="font-weight-bold">{{ field.category }}</span>
                  {{ field.description }}
                </div>
                <div class="col-6 text-right p-0 m-0 mr-3" v-if="!field.bool" style="margin-top: 10px !important">
                  <span v-if="field.category" class="font-weight-bold">{{ field.category }}</span>
                  {{ field.description }}
                </div>
                <div class="col-3 text-left p-0 mt-1">

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

                </div>
              </div>


            </eq-tab>

            <eq-tab name="Effects+" class="minified-inputs">
              <div
                class="row" v-for="field in
                     [
                       {
                         description: 'Push Back',
                         field: 'pushback',
                         category: 'Knockback'
                       },
                       {
                         description: 'Push Up',
                         field: 'pushup',
                         category: 'Knockback'
                       },
                       {
                         description: 'Recourse ID',
                         field: 'recourse_link',
                         category: 'Recourse',
                       },
                       {
                         description: 'Hate Modifier',
                         field: 'bonushate',
                         category: 'Hate',
                       },
                       {
                         description: 'Spell Hate Given',
                         field: 'hate_added',
                         category: 'Hate',
                       },
                       {
                         description: 'No Detrimental Spell Aggro',
                         field: 'field_198',
                         bool: true,
                         category: 'Hate',
                       },
                       {
                         description: 'Max Targets',
                         field: 'maxtargets',
                         category: 'Focus'
                       },
                       {
                         description: 'Song Base Effect Cap',
                         field: 'songcap',
                         category: 'Focus'
                       },
                       {
                         description: 'Not Focusable',
                         field: 'not_extendable',
                         bool: true,
                         category: 'Focus'
                       },
                       {
                         description: 'Max Critical Chance',
                         field: 'field_217'
                       },
                       {
                         description: 'Nimbus Type',
                         field: 'nimbuseffect'
                       },
                     ]"
                @click="processClickInputTrigger(field.field)"
              >
                <div class="col-6 text-right p-0 m-0 mr-3 mt-3" v-if="field.bool">
                  <span v-if="field.category" class="font-weight-bold">{{ field.category }}</span>
                  {{ field.description }}
                </div>
                <div class="col-6 text-right p-0 m-0 mr-3" v-if="!field.bool" style="margin-top: 10px !important">
                  <span v-if="field.category" class="font-weight-bold">{{ field.category }}</span>
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
                    class="row"
                    :key="field.field"
                    v-for="field in
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
                             selectData: {
                               0: 'Detrimental',
                               1: 'Beneficial',
                               2: 'Beneficial Group Only',
                               3: 'Beneficial Group Only'
                             }
                           },
                           {
                             description: 'Mana Cost',
                             field: 'mana',
                             category: 'Cost'
                           },
                           {
                             description: 'Endurance Cost',
                             field: 'endur_cost',
                             category: 'Cost'
                           },
                           {
                             description: 'Endurance Upkeep',
                             field: 'endur_upkeep',
                             category: 'Cost'
                           },
                           {
                             description: 'Primary Category #',
                             field: 'typedescnum',
                             category: 'Description',
                             selectData: this.dbStrSelectData[5]
                           },
                           {
                             description: 'Second Category 1 #',
                             field: 'effectdescnum',
                             category: 'Description',
                             selectData: this.dbStrSelectData[5]
                           },
                           {
                             description: 'Second Category 2 #',
                             field: 'effectdescnum_2',
                             category: 'Description',
                             selectData: this.dbStrSelectData[5]
                           },
                           {
                             description: 'Description #',
                             field: 'descnum',
                             category: 'Description'
                           },
                           {
                             description: 'Group ID',
                             field: 'spellgroup',
                             category: 'Spell Groups'
                           },
                           {
                             description: 'Rank',
                             field: 'rank',
                             category: 'Spell Rank'
                           },
                         ]"
                  >
                    <div class="col-6 text-right p-0 m-0 mr-3" v-if="field.bool">
                      <span v-if="field.category" class="font-weight-bold">{{ field.category }}</span>
                      {{ field.description }}
                    </div>
                    <div
                      class="col-6 text-right p-0 m-0 mr-3"
                      v-if="!field.bool"
                      style="margin-top: 10px !important"
                    >
                      <span v-if="field.category" class="font-weight-bold">{{ field.category }}</span>
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
                        :id="field.field"
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
                    <div class="col-2">
                      <div v-if="field.category === 'Description'">
                        <router-link
                          class="btn btn-warning btn-sm mt-2"
                          tag="button"
                          :to="DB_STRING_EDITOR_URL + '?type=' + (field.field === 'descnum' ? 6 : 5) + '&selectedId=' + spell[field.field] "
                        >
                          <i class="ra ra-scroll-unfurled mr-1"></i> Editor
                        </router-link>
                      </div>
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
                             description: 'Detrimental Spell Allows Rest',
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
                           // {
                           //   description: 'Environment Type',
                           //   field: 'environment_type',
                           // },
                           {
                             description: 'Time of Day',
                             field: 'time_of_day',
                             selectData: {
                               0: 'Any Time',
                               1: 'Day Time',
                               2: 'Night Time',
                             }
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

              <div
                class="row fade-in" v-for="field in
                     [
                       {
                         description: 'NPC Line of Sight Not Required to Cast',
                         field: 'npc_no_los',
                         bool: true,
                         showIf: spell['targettype'] !== 6 && spell['targettype'] !== 7 // exclude 'self' spells
                       },
                       {
                         description: 'Spell Range',
                         type: 'range',
                         field: 'range'
                       },
                       {
                         description: 'Spell Min Range',
                         type: 'range',
                         field: 'min_range'
                       },
                       {
                         description: '(Optional Spell Range) Min Distance for Mod',
                         field: 'min_dist'
                       },
                       {
                         description: '(Optional Spell Range) Min Distance Mod',
                         field: 'min_dist_mod',
                         showIf: spell['min_dist'] !== 0
                       },
                       {
                         description: '(Optional Spell Range) Max Distance for Mod',
                         field: 'max_dist'
                       },
                       {
                         description: '(Optional Spell Range) Max Distance Mod',
                         field: 'max_dist_mod',
                         showIf: spell['max_dist'] !== 0
                       },
                       {
                         description: 'Target Type',
                         field: 'targettype',
                         selectData: DB_SPELL_TARGETS,
                       },
                       {
                         description: 'Cone Angle Start',
                         field: 'cone_start_angle',
                         type: 'cone',
                         showIf: spell['targettype'] === 42 // cone spells
                       },
                       {
                         description: 'Cone Angle End',
                         field: 'cone_stop_angle',
                         type: 'cone',
                         showIf: spell['targettype'] === 42 // cone spells
                       },
                       {
                         description: 'AOE Range',
                         field: 'aoerange',
                         type: 'range',
                         showIf: [4, 8, 24, 45].includes(spell['targettype']) // AOE target types
                       },
                       {
                         description: 'AOE Rain Waves',
                         field: 'ae_duration',
                         showIf: [4, 8, 24, 45].includes(spell['targettype']) // AOE target types
                       },
                       {
                         description: 'AOE Max Targets',
                         field: 'aemaxtargets',
                         showIf: [4, 8, 24, 45].includes(spell['targettype']) // AOE target types
                       },
                       {
                         description: 'Max Hits Type',
                         field: 'numhitstype'
                       },
                       {
                         description: 'Max Hits Allowed',
                         field: 'numhits'
                       },
                     ]"
                :key="field.field"
                v-if="typeof field.showIf === 'undefined' || (typeof field.showIf !== 'undefined' && field.showIf) || showAllFields"
              >
                <div class="col-6 text-right p-0 m-0 mr-3" v-if="field.bool">
                  {{ field.description }}
                </div>
                <div class="col-6 text-right p-0 m-0 mr-3" v-if="!field.bool" style="margin-top: 10px !important">
                  {{ field.description }}
                </div>
                <div
                  class="col-3 text-left p-0 mt-1"
                  @click="processClickInputTrigger(field.field)"
                >
                  <!-- checkbox -->
                  <eq-checkbox
                    v-b-tooltip.hover.v-dark.lefttop :title="getFieldDescription(field.field)"
                    class="mb-1 d-inline-block"
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
                    v-b-tooltip.hover.v-dark.lefttop :title="getFieldDescription(field.field)"
                    :style="(spell[field.field] === 0 ? 'opacity: .5' : '')"
                  />

                  <!-- input text -->
                  <b-form-input
                    v-if="!field.selectData && !field.bool && field.text"
                    :id="field.field"
                    v-model.number="spell[field.field]"
                    class="m-0 mt-1"
                    v-b-tooltip.hover.v-dark.lefttop :title="getFieldDescription(field.field)"
                    :style="(spell[field.field] === '' ? 'opacity: .5' : '')"
                  />

                  <!-- select -->
                  <select
                    v-model.number="spell[field.field]"
                    class="form-control m-0 mt-1"
                    v-if="field.selectData"
                    v-b-tooltip.hover.v-dark.lefttop :title="getFieldDescription(field.field)"
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
                <div
                  class="col-2"
                  @mouseover="processClickInputTrigger(field.field)"
                  @click="processClickInputTrigger(field.field)"
                >
                  <input
                    type="range"
                    v-if="field.type === 'range'"
                    min="0"
                    max="1000"
                    step="5"
                    class="p-0 m-0 mt-2"
                    v-model.number="spell[field.field]"
                  >
                  <input
                    type="range"
                    v-if="field.type === 'cone'"
                    min="0"
                    max="360"
                    step="5"
                    class="p-0 m-0 mt-2"
                    v-model.number="spell[field.field]"
                  >
                  <!-- Modifier description -->
                  <div v-if="['min_dist_mod', 'max_dist_mod'].includes(field.field)" style="margin-top: 8px">
                    ({{ Math.round(spell[field.field] * 100) }}%)
                  </div>
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
                         description: 'Cast Not Standing',
                         field: 'cast_not_standing',
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
                         description: 'Timer Index',
                         field: 'endur_timer_index'
                       },
                       {
                         description: 'Fizzle Adjustment',
                         field: 'basediff'
                       },
                       {
                         description: 'Components # 1',
                         field: 'components_1',
                         category: 'Components'
                       },
                       {
                         description: 'Components # 2',
                         field: 'components_2',
                         category: 'Components'
                       },
                       {
                         description: 'Components # 3',
                         field: 'components_3',
                         category: 'Components'
                       },
                       {
                         description: 'Components # 4',
                         field: 'components_4',
                         category: 'Components'
                       },
                       {
                         description: 'Component Count # 1',
                         field: 'component_counts_1',
                         category: 'Components'
                       },
                       {
                         description: 'Components Count # 2',
                         field: 'component_counts_2',
                         category: 'Components'
                       },
                       {
                         description: 'Components Count # 3',
                         field: 'component_counts_3',
                         category: 'Components'
                       },
                       {
                         description: 'Components Count # 4',
                         field: 'component_counts_4',
                         category: 'Components'
                       },
                       {
                         description: 'Reagent # 1',
                         field: 'noexpend_reagent_1',
                         category: 'Reagent'
                       },
                       {
                         description: 'Reagent # 2',
                         field: 'noexpend_reagent_2',
                         category: 'Reagent'
                       },
                       {
                         description: 'Reagent # 3',
                         field: 'noexpend_reagent_3',
                         category: 'Reagent'
                       },
                       {
                         description: 'Reagent # 4',
                         field: 'noexpend_reagent_4',
                         category: 'Reagent'
                       },
                     ]"
                @click="processClickInputTrigger(field.field)"
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
                <div class="col3 pl-3">
                  <div
                    class="ml-5"
                    v-if="field.description.includes('Time') && !field.description.includes('Timer')"
                  >
                    ({{ (Math.round((spell[field.field] / 1000) * 10) / 10) }} sec)
                  </div>

                  <loader-cast-bar-timer
                    class="ml-3"
                    style="margin-top: 5px"
                    color="#FF00FF"
                    v-if="field.description.includes('Time') && !field.description.includes('Timer')"
                    :time-ms="spell[field.field]"
                  />
                </div>
              </div>

            </eq-tab>

            <eq-tab name="Buffs" class="minified-inputs">
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
                         field: 'buffduration',
                       },
                       {
                         description: 'Buff Duration Formula',
                         field: 'buffdurationformula',
                         selectData: BUFF_DURATION_FORMULAS,
                       },
                       {
                         description: 'Viral Range',
                         field: 'viral_range',
                         category: 'Viral'
                       },
                       {
                         description: 'Viral Targets',
                         field: 'viral_targets',
                         category: 'Viral'
                       },
                       {
                         description: 'Viral Timer',
                         field: 'viral_timer',
                         category: 'Viral'
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
                  <span v-if="field.category" class="font-weight-bold">{{ field.category }}</span>
                  {{ field.description }}
                </div>
                <div class="col-6 text-right p-0 m-0 mr-3" v-if="!field.bool" style="margin-top: 10px !important">
                  <span v-if="field.category" class="font-weight-bold">{{ field.category }}</span>
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

                <div
                  class="col-2"
                  @mouseover="processClickInputTrigger(field.field)"
                  @click="processClickInputTrigger(field.field)"
                >
                  <!-- Modifier description -->
                  <div v-if="['buffduration'].includes(field.field)" style="margin-top: 8px">
                    {{ humanTime(getBuffDuration(spell) * 6) }} - {{ getBuffDuration(spell) }} tic(s)
                  </div>
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
                         description: 'Reflectable',
                         field: 'reflectable',
                         bool: true
                       },
                       {
                         description: 'Feedbackable',
                         field: 'field_160',
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

            <eq-tab name="Misc" class="minified-inputs">
              <div class="row">
                <div class="col-12">
                  <div
                    class="row"
                    :key="field.field"
                    v-for="field in
                         [
                           {
                             description: 'Deleteable',
                             field: 'deleteable',
                             bool: true
                           },
                           {
                             description: 'Activated',
                             field: 'activated',
                             bool: true
                           },
                           {
                             description: 'Light Type',
                             field: 'light_type',
                           },
                           {
                             description: 'Travel Type',
                             field: 'travel_type',
                           },
                           {
                             description: 'LDON Trap',
                             field: 'ldon_trap',
                           },
                           {
                             description: 'Spell Category',
                             field: 'spell_category',
                           },
                           {
                             description: 'NPC Category',
                             field: 'npc_category',
                           },
                           {
                             description: 'NPC Usefulness',
                             field: 'npc_usefulness',
                           },
                         ]"
                  >
                    <div class="col-6 text-right p-0 m-0 mr-3" v-if="field.bool">
                      <span v-if="field.category" class="font-weight-bold">{{ field.category }}</span>
                      {{ field.description }}
                    </div>
                    <div
                      class="col-6 text-right p-0 m-0 mr-3"
                      v-if="!field.bool"
                      style="margin-top: 10px !important"
                    >
                      <span v-if="field.category" class="font-weight-bold">{{ field.category }}</span>
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
                        :id="field.field"
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
                    <div class="col-2">
                      <div v-if="field.category === 'Description'">
                        <router-link
                          class="btn btn-warning btn-sm mt-2"
                          tag="button"
                          :to="DB_STRING_EDITOR_URL + '?type=' + (field.field === 'descnum' ? 6 : 5) + '&selectedId=' + spell[field.field] "
                        >
                          <i class="ra ra-scroll-unfurled mr-1"></i> Strings Editor
                        </router-link>
                      </div>
                    </div>
                  </div>

                </div>
              </div>

            </eq-tab>
          </eq-tabs>

          <div class="text-center align-content-center mt-4" v-if="spell && spell.id >= 0">

            <b-button
              @click="saveSpell()"
              size="sm"
              variant="outline-warning"
            >
              <i class="ra ra-book"></i>
              Save Spell
            </b-button>

          </div>

          <div class="row" v-if="spell">
            <div class="col-10"></div>
            <div class="col-2 text-right" title="Show unknown fields">
              Show All Fields
              <eq-checkbox
                class="mb-2 d-inline-block"
                v-model.number="showAllFields"
              />
            </div>
          </div>


        </eq-window>

      </div>

      <!-- Preview Pane -->
      <div class="col-xl-5 col-lg-12">

        <!-- SPA Detail Pane -->
        <eq-window
          style=" margin-right: 10px; width: auto;"
          class="fade-in"
          v-if="spaDetailPaneActive"
          :title="'Effect (' + spaPreviewNumber + ') Details'"
        >
          <spell-spa-preview-pane
            :spa="spaPreviewNumber"
            :spell="spell"
            :effect-index="spaEffectIndex"
            v-if="spaPreviewNumber >= 0"
          />

        </eq-window>

        <!-- preview spell -->
        <eq-window
          style=" margin-right: 10px; width: auto;"
          v-if="previewSpellActive"
        >
          <eq-spell-preview
            :spell-data="spell"
          />
        </eq-window>

        <!-- icon selector -->
        <eq-window
          style=" margin-right: 10px; width: auto;"
          class="fade-in"
          v-if="iconSelectorActive"
        >
          <spell-icon-selector
            :selected-icon="spell.new_icon"
            :inputData.sync="spell.new_icon"
          />
        </eq-window>

        <!-- spell anim selector -->
        <div
          style=" width: auto;"
          class="fade-in"
          v-if="spellAnimSelectorActive"
        >
          <spell-animation-selector
            :selected-animation="spell.spellanim"
            :inputData.sync="spell.spellanim"
          />
        </div>

        <!-- spell nimbus anim selector -->
        <div
          style=" width: auto;"
          class="fade-in"
          v-if="spellNimbusAnimSelectorActive"
        >
          <spell-nimbus-animation-selector
            :selected-animation="spell.nimbuseffect"
            :inputData.sync="spell.nimbuseffect"
          />
        </div>

        <!-- spell effect selector (Used in effectid 1-12)-->
        <div
          style=" width: auto;"
          class="fade-in"
          v-if="spellSelectorActive"
        >
          <spell-spell-effect-selector
            @input="spell[selectedEffectColumn + '_' + selectedEffectIndex] = $event.spellId; setFieldModifiedById(selectedEffectColumn + '_' + selectedEffectIndex)"
          />
        </div>

        <!-- spellitem selector -->
        <div
          style=" width: auto;"
          class="fade-in"
          v-if="itemSelectorActive"
        >
          <spell-item-selector
            @input="spell[selectedItemSelectorField] = $event.id; setFieldModifiedById(selectedItemSelectorField)"
          />
        </div>

        <!-- simple spell effect selector (Used in simple fields) -->
        <div
          style=" width: auto;"
          class="fade-in"
          v-if="simpleSpellSelectorActive"
        >
          <spell-spell-effect-selector
            @input="spell[selectedSimpleSpellSelectorField] = $event.spellId; setFieldModifiedById(selectedSimpleSpellSelectorField)"
          />
        </div>

        <!-- spell casting anim selector -->
        <div
          style=" width: auto;"
          class="fade-in"
          v-if="castingAnimSelectorActive && spell[castingAnimField] >= 0"
        >
          <spell-casting-animation-selector
            :selected-animation="spell[castingAnimField]"
            :inputData.sync="spell[castingAnimField]"
          />
        </div>

        <!-- cone angle visualizer -->
        <eq-window
          style=" height: 500px"
          class="fade-in text-center"
          title="Cone Visualizer"
          v-if="coneVisualizerActive"
        >
          <spell-cone-visualizer
            :cone-start-angle="spell.cone_start_angle"
            :cone-stop-angle="spell.cone_stop_angle"
          />
        </eq-window>

        <!-- range visualizer -->
        <eq-window-simple
          style=""
          class="fade-in text-center"
          title="Range Visualizer"
          v-if="rangeVisualizerActive && spell[activeRangeField] >= 0"
        >
          <range-visualizer :unit-marker="spell[activeRangeField]"/>
        </eq-window-simple>

        <!-- range visualizer -->
        <div
          style="margin-top: 20px"
          class="fade-in text-center"
          v-if="teleportZoneSelectorActive && typeof spell['teleport_zone'] !== 'undefined'"
        >
          <spell-teleport-zone-selector-aura
            :selected-aura-name="spell['teleport_zone']"
            v-if="selectedTeleportZoneSelectorType === TELEPORT_ZONE_SELECTOR_TYPE.AURAS"
            @input="processTeleportZoneAuraSelectUpdate($event)"
          />

          <spell-teleport-zone-selector-horse
            :selected-horse-name="spell['teleport_zone']"
            v-if="selectedTeleportZoneSelectorType === TELEPORT_ZONE_SELECTOR_TYPE.HORSES"
            @input="processTeleportZoneHorseSelectUpdate($event)"
          />

          <spell-teleport-zone-selector-pet
            :selected-pet-name="spell['teleport_zone']"
            v-if="selectedTeleportZoneSelectorType === TELEPORT_ZONE_SELECTOR_TYPE.PETS"
            @input="processTeleportZonePetSelectUpdate($event)"
          />

          <spell-teleport-zone-selector-zone
            :selected-zone-name="spell['teleport_zone']"
            v-if="selectedTeleportZoneSelectorType === TELEPORT_ZONE_SELECTOR_TYPE.ZONES"
            @input="processTeleportZoneZoneSelectUpdate($event)"
          />
        </div>

        <!-- free id selector -->
        <eq-window
          title="Free Spell Ids"
          style=" margin-right: 10px; width: auto;"
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

      </div>
    </div>
  </div>
</template>

<script>
import EqWindowFancy                  from "../../components/eq-ui/EQWindowFancy";
import EqWindow                       from "../../components/eq-ui/EQWindow";
import EqTabs                         from "../../components/eq-ui/EQTabs";
import EqTab          from "../../components/eq-ui/EQTab";
import EqSpellPreview from "../../components/preview/EQSpellCardPreview";
import {Spells}       from "../../app/spells";
import {
  BASE_VALUE_FORMULAS,
  BUFF_DURATION_FORMULAS,
  DB_PC_NPC_ONLY_FLAG,
  DB_SPA,
  DB_SPELL_EFFECTS,
  DB_SPELL_RESISTS,
  DB_SPELL_TARGET_RESTRICTION,
  DB_SPELL_TARGETS,
  DB_SPELL_ZONE_TYPE,
  TELEPORT_ZONE_SELECTOR_TYPE
}                                     from "../../app/constants/eq-spell-constants";
import {DB_SKILLS}                    from "../../app/constants/eq-skill-constants";
import SpellIconSelector              from "./components/SpellIconSelector";
import SpellAnimationPreview  from "./components/SpellAnimationPreview";
import SpellAnimationViewer   from "../asset-viewers/SpellAnimationViewer";
import SpellAnimationSelector from "./components/SpellAnimationSelector";
import EqCheckbox                     from "../../components/eq-ui/EQCheckbox";
import {DbStrApi, SpellsNewApi}       from "../../app/api";
import {SpireApiClient}               from "../../app/api/spire-api-client";
import SpellClassSelector             from "./components/SpellClassSelector";
import SpellDeitySelector             from "./components/SpellDeitySelector";
import FreeIdSelector                 from "../../components/tools/FreeIdSelector";
import SpellSpaPreviewPane            from "./components/SpellSpaPreviewPane";
import SpellCastingAnimationPreview   from "./components/SpellCastingAnimationPreview";
import SpellCastingAnimationSelector  from "./components/SpellCastingAnimationSelector";
import {SPELL_SPA_DEFINITIONS}        from "../../app/constants/eq-spell-spa-definitions";
import LoaderCastBarTimer             from "../../components/LoaderCastBarTimer";
import {EditFormFieldUtil}            from "../../app/forms/edit-form-field-util";
import LoaderFakeProgress             from "../../components/LoaderFakeProgress";
import SpellSpellEffectSelector       from "./components/SpellSpellEffectSelector";
import {debounce}                     from "../../app/utility/debounce";
import EqWindowSimple                 from "../../components/eq-ui/EQWindowSimple";
import SpellConeVisualizer            from "./components/SpellConeVisualizer";
import SpellNimbusAnimationSelector   from "./components/SpellNimbusAnimationSelector";
import util                           from "util";
import RangeVisualizer                from "../../components/tools/RangeVisualizer";
import SpellTeleportZoneSelectorZone  from "./components/SpellTeleportZoneSelectorZone";
import SpellTeleportZoneSelectorPet   from "./components/SpellTeleportZoneSelectorPet";
import SpellTeleportZoneSelectorHorse from "./components/SpellTeleportZoneSelectorHorse";
import SpellTeleportZoneSelectorAura  from "./components/SpellTeleportZoneSelectorAura";
import {ROUTE}                        from "../../routes";
import SpellItemSelector              from "./components/SpellItemSelector";
import {SpireQueryBuilder}            from "../../app/api/spire-query-builder";
import {FreeIdFetcher}                from "../../app/free-id-fetcher";
import ContentArea                    from "../../components/layout/ContentArea";

const MILLISECONDS_BEFORE_WINDOW_RESET = 5000;

export default {
  name: "SpellEdit",
  components: {
    ContentArea,
    SpellItemSelector,
    SpellTeleportZoneSelectorAura,
    SpellTeleportZoneSelectorHorse,
    SpellTeleportZoneSelectorPet,
    SpellTeleportZoneSelectorZone,
    RangeVisualizer,
    SpellNimbusAnimationSelector,
    SpellConeVisualizer,
    EqWindowSimple,
    SpellSpellEffectSelector,
    LoaderFakeProgress,
    LoaderCastBarTimer,
    SpellCastingAnimationSelector,
    SpellCastingAnimationPreview,
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

      // constants
      DB_SPELL_EFFECTS: DB_SPELL_EFFECTS,
      DB_SPA: DB_SPA,
      DB_SKILLS: DB_SKILLS,
      DB_SPELL_TARGETS: DB_SPELL_TARGETS,
      DB_SPELL_RESISTS: DB_SPELL_RESISTS,
      DB_SPELL_TARGET_RESTRICTION: DB_SPELL_TARGET_RESTRICTION,
      DB_SPELL_ZONE_TYPE: DB_SPELL_ZONE_TYPE,
      DB_PC_NPC_ONLY_FLAG: DB_PC_NPC_ONLY_FLAG,
      BUFF_DURATION_FORMULAS: BUFF_DURATION_FORMULAS,
      BASE_VALUE_FORMULAS: BASE_VALUE_FORMULAS,
      TELEPORT_ZONE_SELECTOR_TYPE: TELEPORT_ZONE_SELECTOR_TYPE,
      DB_STRING_EDITOR_URL: ROUTE.STRINGS_DATABASE,
      loaded: true,

      // preview / selectors
      previewSpellActive: true,
      iconSelectorActive: false,
      spellAnimSelectorActive: false,
      spellNimbusAnimSelectorActive: false,
      freeIdSelectorActive: false,
      spaDetailPaneActive: false,
      castingAnimSelectorActive: false,
      spellSelectorActive: false,
      itemSelectorActive: false,
      simpleSpellSelectorActive: false,
      coneVisualizerActive: false,
      rangeVisualizerActive: false,
      teleportZoneSelectorActive: false,

      spaPreviewNumber: -1,
      spaEffectIndex: -1,

      selectedEffectIndex: 0,
      selectedEffectColumn: "",

      visibleEffectSlots: [],
      zeroStateSelected: true,

      selectedSimpleSpellSelectorField: "",
      activeRangeField: "",
      selectedTeleportZoneSelectorType: "",
      selectedItemSelectorField: "",

      // cache
      dbStrSelectData: {},
      teleportZoneFieldName: "",

      castingAnimField: "",

      showAllFields: 0,

      lastResetTime: Date.now(),

      notification: "",
      error: "",
    }
  },
  watch: {
    'spell.cone_start_angle'(newVal, oldVal) {
      if (this.loaded && newVal && oldVal) {
        this.drawConeVisualizer()
      }
    },
    'spell.cone_end_angle'(newVal, oldVal) {
      if (this.loaded && newVal && oldVal) {
        this.drawConeVisualizer()
      }
    },

    '$route'() {
      this.loaded = false

      // reset state vars when we navigate away
      this.visibleEffectSlots = []
      this.notification       = ""
      this.zeroStateSelected  = true
      EditFormFieldUtil.resetFieldEditedStatus()
      this.resetPreviewComponents()
      EditFormFieldUtil.resetFieldHighlightHasSubEditorStatus()

      // reload
      this.load()
    }
  },
  async created() {
    this.load()
  },
  methods: {

    none() {
      for (let i = 1; i <= 16; i++) {
        const classIndex = "classes_" + i
        if (this.spell[classIndex] !== 255) {
          EditFormFieldUtil.setFieldModifiedById(classIndex)
        }

        this.spell[classIndex] = 255
      }
    },

    async getDbStringsSelectData(typeId) {
      if (this.dbStrSelectData[typeId]) {
        return this.dbStrSelectData[typeId]
      }

      const DbStrApiClient = (new DbStrApi(SpireApiClient.getOpenApiConfig()))
      const response       = await DbStrApiClient.listDbStrs(
        (new SpireQueryBuilder()
            .where("type", "=", typeId)
            .limit(100000)
            .get()
        )
      )

      if (response.status === 200 && response.data) {
        this.loading   = false
        const strings  = response.data
        let selectData = {
          0: 'None'
        }
        strings.forEach((string) => {
          selectData[string.id] = string.value;
        })

        this.dbStrSelectData[typeId] = selectData
      }
    },

    processTeleportZoneAuraSelectUpdate(event) {
      this.spell['teleport_zone'] = event.horse.name
    },
    processTeleportZoneHorseSelectUpdate(event) {
      this.spell['teleport_zone'] = event.horse.filename
    },
    processTeleportZonePetSelectUpdate(event) {
      this.spell['teleport_zone'] = event.pet.type
    },
    processTeleportZoneZoneSelectUpdate(event) {
      console.log("[processTeleportZoneZoneSelectUpdate]")
      // console.log(event)

      const selectedZone  = event.zone;
      const TELEPORT_SPAS = [
        83, 88, 104, 145,
      ]

      for (let effectIndex = 1; effectIndex <= 12; effectIndex++) {
        const spaId = parseInt(this.spell['effectid_' + effectIndex])
        if (spaId !== 254 && TELEPORT_SPAS.includes(spaId)) {

          // set zone

          this.spell['teleport_zone'] = selectedZone.short_name
          EditFormFieldUtil.setFieldModifiedById("teleport_zone")

          // make revelant slots visible when a zone is selected
          for (let slot = effectIndex; slot < (effectIndex + 4); slot++) {
            this.visibleEffectSlots[slot] = true
            this.$forceUpdate()
            setTimeout(() => {
              EditFormFieldUtil.setFieldModifiedById("effect_base_value_" + slot)
            }, 100)
          }

          this.spell["effect_base_value_" + effectIndex]       = selectedZone.safe_x
          this.spell["effect_base_value_" + (effectIndex + 1)] = selectedZone.safe_y
          this.spell["effect_base_value_" + (effectIndex + 2)] = selectedZone.safe_z
          this.spell["effect_base_value_" + (effectIndex + 3)] = selectedZone.safe_heading

          console.log("teleport spa is [%s]", spaId)
        }
      }
    },

    toTitleCase(str) {
      return str.replace(
        /\w\S*/g,
        function (txt) {
          return txt.charAt(0).toUpperCase() + txt.substr(1).toLowerCase();
        }
      );
    },

    getTeleportZoneFieldName() {
      const selectorType = this.getTeleportZoneSelectorType()
      if (selectorType === TELEPORT_ZONE_SELECTOR_TYPE.PETS) {
        return "Select Spell Pet"
      }
      if (selectorType === TELEPORT_ZONE_SELECTOR_TYPE.HORSES) {
        return "Select Spell Horse"
      }
      if (selectorType === TELEPORT_ZONE_SELECTOR_TYPE.ZONES) {
        return "Select Zone"
      }
      if (selectorType === TELEPORT_ZONE_SELECTOR_TYPE.AURAS) {
        return "Select Aura"
      }

      return this.toTitleCase(this.getTeleportZoneSelectorType())
    },

    getTeleportZoneSelectorType() {
      let selectorType = ""
      for (let effectIndex = 1; effectIndex <= 12; effectIndex++) {
        const spaId = this.spell['effectid_' + effectIndex]
        if (spaId !== 254) {
          const data  = SPELL_SPA_DEFINITIONS[spaId]
          const notes = data.notes

          if (notes.includes("pets table")) {
            selectorType = TELEPORT_ZONE_SELECTOR_TYPE.PETS
          }
          if (notes.includes("horses")) {
            selectorType = TELEPORT_ZONE_SELECTOR_TYPE.HORSES
          }
          if (notes.includes("zone short name")) {
            selectorType = TELEPORT_ZONE_SELECTOR_TYPE.ZONES
          }
          // if (notes.includes("auras table")) {
          //   selectorType = TELEPORT_ZONE_SELECTOR_TYPE.AURAS
          // }
        }
      }

      console.log("[getTeleportZoneSelectorType] selectorType [%s]", selectorType)

      return selectorType
    },

    toggleVisibleEffectSlot(slot) {
      if (this.spell['effectid_' + slot] !== 254) {
        return
      }

      this.visibleEffectSlots[slot] = !this.visibleEffectSlots[slot]
      this.$forceUpdate()
    },

    processClickInputTrigger(field) {
      if (field === "cone_start_angle" || field === "cone_stop_angle") {
        this.drawConeVisualizer()
      }
      if (field === "recourse_link") {
        this.drawSimpleSpellSelector(field)
      }
      if (field === "nimbuseffect") {
        this.drawSpellNimbusAnimationSelector()
      }
      if (["range", "aoerange", "min_range"].includes(field)) {
        this.drawRangeVisualizer(field)
      }
      if (["teleport_zone"].includes(field)) {
        this.drawTeleportZoneSelector()
      }
      if (
        [
          "components_1",
          "components_2",
          "components_3",
          "components_4",
          "noexpend_reagent_1",
          "noexpend_reagent_2",
          "noexpend_reagent_3",
          "noexpend_reagent_4"
        ].includes(field)) {
        this.drawItemSelector(field);
      }
    },

    setSubEditorFieldHighlights() {
      let hasSubEditorFields = [
        "id",
        "casting_anim",
        "target_anim",
        "icon",
        "spellanim",
        "cone_start_angle",
        "cone_stop_angle",
        "nimbuseffect",
        "aoerange",
        "teleport_zone",
        "range",
        "min_range",
        "typedescnum",
        "effectdescnum",
        "effectdescnum_2",
        "descnum",
        "components_1",
        "components_2",
        "components_3",
        "components_4",
        "noexpend_reagent_1",
        "noexpend_reagent_2",
        "noexpend_reagent_3",
        "noexpend_reagent_4",
        "recourse_link"
      ]
      hasSubEditorFields.forEach((field) => {
        EditFormFieldUtil.setFieldHighlightHasSubEditor(field)
      })
    },

    getBuffDuration: function (spell) {
      return Spells.getBuffDuration(spell)
    },

    humanTime: function (sec) {
      let result = ""
      if (sec === 0) {
        result = "time";
      } else {
        let h  = Math.floor(sec / 3600);
        let m  = Math.floor((sec - h * 3600) / 60);
        let s  = sec - h * 3600 - m * 60;
        result = (h > 1 ? h + " hours " : "") + (h === 1 ? "1 hour " : "") + (m > 0 ? m + " min " : "") + (s > 0 ? s + " sec" : "");
      }

      return result;
    },

    setFieldModifiedById(id) {
      EditFormFieldUtil.setFieldModifiedById(id)
    },

    getSpaSpellHighlights(spaId, field) {
      if (SPELL_SPA_DEFINITIONS[spaId][field]) {

        // highlight when SPA field links to a sub-editor
        const contents = SPELL_SPA_DEFINITIONS[spaId][field]
        if (["spellid", "spell id", "item id", "itemid"].includes(contents)) {
          return 'pulsate-highlight-green'
        }
      }
    },

    processSpaFieldAction(effectIndex, spaId, field) {
      if (SPELL_SPA_DEFINITIONS[spaId][field]) {

        // highlight when SPA field links to a sub-editor
        const contents = SPELL_SPA_DEFINITIONS[spaId][field]

        // activate spell selector
        if (["spellid", "spell id"].includes(contents)) {
          this.drawSpellSelector(effectIndex, field)
        }
        // activate item selector
        if (["itemid", "item id"].includes(contents)) {
          let fieldId = ""
          if (field.includes("max")) {
            fieldId = "max"
          }
          if (field.includes("limit")) {
            fieldId = "effect_limit_value"
          }
          if (field.includes("base")) {
            fieldId = "effect_base_value"
          }

          this.drawItemSelector(fieldId + "_" + effectIndex)
        }
      }
    },

    getFieldDescription(field) {
      return Spells.getFieldDescription(field);
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
          Spells.setSpell(this.spell.id, this.spell) // update cache
          this.sendNotification("Spell updated successfully!")
          EditFormFieldUtil.resetFieldEditedStatus()
        }

        if (result.data.error) {
          this.notification = result.data.error
        }

      }).catch(async (error) => {

        // some sort of validation error, throw error to user
        if (error.response.data && error.response.data.error) {
          const err           = error.response.data.error
          const expectedError = err.includes("Cannot find entity")
          if (!expectedError) {
            this.error = error.response.data.error
            return
          }
        }

        const createRes = await api.createSpellsNew({
          spellsNew: this.spell
        })

        if (createRes.status === 200) {
          this.sendNotification("Created new Spell!")
          EditFormFieldUtil.resetFieldEditedStatus()
        }
      })
    },

    getSpaDefaultValues(spa, index) {
      const api   = (new SpellsNewApi(SpireApiClient.getOpenApiConfig()))
      let whereOr = [];
      for (let effectIndex = 1; effectIndex <= 12; effectIndex++) {
        whereOr.push(["effectid" + effectIndex, "__", spa]);
      }

      let wheresOrs = [];
      whereOr.forEach((filter) => {
        const where = util.format("%s%s%s", filter[0], filter[1], filter[2])
        wheresOrs.push(where)
      })

      let request   = {};
      request.limit = 1;

      if (Object.keys(wheresOrs).length > 0) {
        request.whereOr = wheresOrs.join(".")
      }

      let sourceFields = [
        'effect_base_value_',
        'effect_limit_value_',
        'max_',
        'formula_'
      ]

      // clear
      sourceFields.forEach((field) => {
        this.spell[field + index] = 0
        EditFormFieldUtil.clearFieldModifiedById(field + index)
      })

      api.listSpellsNews(request).then(async (result) => {
        if (result.status === 200) {
          if (result.data.length > 0) {
            const exampleData = result.data[0]

            for (let i = 1; i <= 12; i++) {
              if (exampleData['effectid_' + i] === parseInt(spa)) {


                // set current spell to example values
                sourceFields.forEach((field) => {
                  this.spell[field + index] = exampleData[field + i]
                  EditFormFieldUtil.setFieldModifiedById(field + index)
                })

              }
            }
          }
        }
      });
    },

    load() {

      if (this.$route.params.id > 0) {
        this.error = ""
        this.getDbStringsSelectData(5)

        Spells.getSpell(this.$route.params.id).then(async (result) => {
          this.spell = JSON.parse(JSON.stringify(result))

          // if we're cloning this spell, automatically fetch an ID
          if (this.$route.query.hasOwnProperty("clone")) {
            const id = await FreeIdFetcher.get("spells_new", "id", "name")
            if (id > 0) {
              EditFormFieldUtil.setFieldModifiedById('id')
              this.spell.id = id
            }
          }

          // hooks
          setTimeout(() => {
            const target = document.getElementById("spell-edit-card")
            if (target) {
              target.removeEventListener('input', EditFormFieldUtil.setFieldModified, true);
              target.addEventListener('input', EditFormFieldUtil.setFieldModified)
            }

            this.resetPreviewComponents()
            this.previewSpellActive = true
            this.setSubEditorFieldHighlights()
          }, 300)

          // visible slots effectid 1-12
          for (let i = 1; i <= 12; i++) {
            if (this.spell["effectid_" + i] !== 254) {
              this.visibleEffectSlots[i] = true
            }
          }

          // calc field name
          this.teleportZoneFieldName = this.getTeleportZoneFieldName()
          this.loaded                = true
        })
      }
    },

    /**
     * Selector / previewers
     */
    resetPreviewComponents() {
      this.previewSpellActive            = false;
      this.iconSelectorActive            = false;
      this.spellAnimSelectorActive       = false;
      this.spellNimbusAnimSelectorActive = false;
      this.freeIdSelectorActive          = false;
      this.spaDetailPaneActive           = false;
      this.castingAnimSelectorActive     = false;
      this.spellSelectorActive           = false;
      this.itemSelectorActive            = false;
      this.simpleSpellSelectorActive     = false;
      this.coneVisualizerActive          = false;
      this.rangeVisualizerActive         = false;
      this.teleportZoneSelectorActive    = false;

      EditFormFieldUtil.resetFieldSubEditorHighlightedStatus()
    },
    shouldReset() {
      return (Date.now() - this.lastResetTime) > MILLISECONDS_BEFORE_WINDOW_RESET
    },

    previewSpell: debounce(function () {
      if (this.shouldReset()) {
        this.resetPreviewComponents()
        this.previewSpellActive = true;
      }
    }, 300),

    drawCastingAnimationSelector() {
      this.resetPreviewComponents()
      this.lastResetTime             = Date.now()
      this.castingAnimSelectorActive = false;

      // queue this so that we can "redraw" the selector when toggling between casting and target
      // this makes the scroll to selected animation more consistent
      setTimeout(() => {
        this.castingAnimSelectorActive = true;
      }, 100)
      EditFormFieldUtil.setFieldSubEditorHighlightedById(this.castingAnimField)
    },
    drawSpellAnimationSelector() {
      if (!this.spellAnimSelectorActive) {
        this.resetPreviewComponents()
        this.lastResetTime = Date.now()
        setTimeout(() => {
          this.spellAnimSelectorActive = true
        }, 100)
        EditFormFieldUtil.setFieldSubEditorHighlightedById("spellanim")
      }
    },
    drawSpellNimbusAnimationSelector() {
      if (!this.spellNimbusAnimSelectorActive) {
        this.resetPreviewComponents()
        this.lastResetTime = Date.now()
        setTimeout(() => {
          this.spellNimbusAnimSelectorActive = true
        }, 100)
        EditFormFieldUtil.setFieldSubEditorHighlightedById("nimbuseffect")
      }
    },
    drawIconSelector() {
      if (!this.iconSelectorActive) {
        this.resetPreviewComponents()
        this.lastResetTime      = Date.now()
        this.iconSelectorActive = true;
        EditFormFieldUtil.setFieldSubEditorHighlightedById("icon")
      }
    },
    drawFreeIdSelector() {
      if (!this.freeIdSelectorActive) {
        this.resetPreviewComponents()
        this.lastResetTime        = Date.now()
        this.freeIdSelectorActive = true
        EditFormFieldUtil.setFieldSubEditorHighlightedById("id")
      }
    },
    drawSpellSelector(effectIndex, field) {
      this.resetPreviewComponents()
      this.lastResetTime       = Date.now()
      this.spellSelectorActive = true
      this.selectedEffectIndex = effectIndex

      let fieldId = ""
      if (field.includes("max")) {
        fieldId = "max"
      }
      if (field.includes("limit")) {
        fieldId = "effect_limit_value"
      }
      if (field.includes("base")) {
        fieldId = "effect_base_value"
      }

      this.selectedEffectColumn = fieldId

      EditFormFieldUtil.setFieldSubEditorHighlightedById(fieldId + "_" + effectIndex)
    },
    drawItemSelector(field) {
      this.resetPreviewComponents()
      this.lastResetTime             = Date.now() + 5000
      this.itemSelectorActive        = true
      this.selectedItemSelectorField = field
      EditFormFieldUtil.setFieldSubEditorHighlightedById(field)
    },
    drawSimpleSpellSelector(field) {
      this.resetPreviewComponents()
      this.lastResetTime                    = Date.now()
      this.simpleSpellSelectorActive        = true
      this.selectedSimpleSpellSelectorField = field

      EditFormFieldUtil.setFieldSubEditorHighlightedById(field)
    },
    drawConeVisualizer() {
      if (!this.coneVisualizerActive) {
        this.resetPreviewComponents()
        this.coneVisualizerActive = true
        this.lastResetTime        = Date.now() + 30000
        EditFormFieldUtil.setFieldSubEditorHighlightedById("cone_start_angle")
        EditFormFieldUtil.setFieldSubEditorHighlightedById("cone_stop_angle")
      }
    },
    drawRangeVisualizer(field) {
      this.resetPreviewComponents()
      this.activeRangeField      = field
      this.rangeVisualizerActive = true
      this.lastResetTime         = Date.now() + 5000
      EditFormFieldUtil.setFieldSubEditorHighlightedById(field)
    },
    drawTeleportZoneSelector() {
      this.resetPreviewComponents()
      this.selectedTeleportZoneSelectorType = this.getTeleportZoneSelectorType()
      this.teleportZoneSelectorActive       = true
      this.lastResetTime                    = Date.now() + 5000
      EditFormFieldUtil.setFieldSubEditorHighlightedById("teleport_zone")
    },
    drawSpaDetailPane(spa, index) {
      this.resetPreviewComponents()
      EditFormFieldUtil.setFieldSubEditorHighlightedById("effectid_" + index)
      this.previewSpellActive  = true
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
