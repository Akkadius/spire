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
                    <div class="col-4">
                      Name
                      <b-form-input
                        :value="item.name" @change="v => item.name = v"
                      />
                    </div>

                    <!-- Lore -->
                    <div class="col-4">
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
                    <div class="col-12">
                      <div class="row">

                        <!-- Item Type -->
                        <div class="col-3">
                          Item Type
                          <select v-model.number="item['itemtype']" class="form-control">
                            <option
                              v-for="(description, index) in DB_ITEM_TYPES"
                              :key="index"
                              :value="parseInt(index)"
                            >
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
                              :value="parseInt(index)"
                            >
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
                              :value="parseInt(index)"
                            >
                              {{ index }}) {{ description }}
                            </option>
                          </select>
                        </div>

                        <!-- Stack Size -->
                        <div class="col-2">
                          Stack Size
                          <b-form-input v-model.number="item.stacksize"/>
                        </div>
                      </div>


                    </div>


                  </div>

                  <div class="mt-3 mb-3">
                    <div class="row">
                      <div class="col-2 text-center">

                        <div
                          class="row" v-for="field in
                       [
                         {
                           description: 'Is Magic',
                           field: 'magic'
                         },
                         {
                           description: 'No Drop',
                           field: 'nodrop',
                           true: 0,
                           false: 1,
                         },
                         {
                           description: 'FV No Drop',
                           field: 'fvnodrop',
                         },
                         {
                           description: 'No Rent',
                           field: 'norent',
                           true: 0,
                           false: 1,
                         },
                         {
                           description: 'Tradeskill Item',
                           field: 'tradeskills'
                         },
                         {
                           description: 'Book',
                           field: 'book'
                         },
                         {
                           description: 'No Transfer',
                           field: 'notransfer'
                         },
                         {
                           description: 'Summoned',
                           field: 'summonedflag'
                         },
                         {
                           description: 'Quest',
                           field: 'questitemflag'
                         },
                         {
                           description: 'Artifact',
                           field: 'artifactflag'
                         },
                         {
                           description: 'No Pet',
                           field: 'nopet'
                         },
                         {
                           description: 'Attuneable',
                           field: 'attuneable'
                         },
                         {
                           description: 'Stackable',
                           field: 'stackable'
                         },
                         {
                           description: 'Potion Belt',
                           field: 'potionbelt'
                         },
                         {
                           description: 'Placeable',
                           field: 'placeable'
                         },
                       ]"
                        >
                          <div class="col-9 text-right p-0 pr-2">
                            {{ field.description }}
                          </div>
                          <div class="col-3 text-left p-0">
                            <eq-checkbox
                              class="mb-2 d-inline-block"
                              :true-value="(typeof field.true !== 'undefined' ? field.true : 1)"
                              :false-value="(typeof field.false !== 'undefined' ? field.false : 0)"
                              v-model.number="item[field.field]"
                              @input="item[field.field] = $event"
                            />
                          </div>
                        </div>

                      </div>

                      <div class="col-2">
                        <h4 class="eq-header text-center">
                          Class
                        </h4>
                        <class-bitmask-calculator
                          class="text-center mt-3"
                          :show-text-top="false"
                          :show-text-side="true"
                          :imageSize="38"
                          :centered-buttons="true"
                          @input="item.classes = parseInt($event)"
                          :mask="item.classes"
                        />
                      </div>
                      <div class="col-2">
                        <h4 class="eq-header text-center">
                          Race
                        </h4>
                        <race-bitmask-calculator
                          :imageSize="37"
                          class="mt-3"
                          :show-text-top="false"
                          :centered-buttons="true"
                          @input="item.races = parseInt($event)"
                          :mask="item.races"
                        />
                      </div>
                      <div class="col-2">
                        <h4 class="eq-header text-center">
                          Deity
                        </h4>
                        <deity-bitmask-calculator
                          class="mt-3"
                          :imageSize="37"
                          :show-names="false"
                          :centered-buttons="true"
                          @input="item.deity = parseInt($event)"
                          :mask="item.deity"
                        />
                      </div>

                      <!-- Slots -->
                      <div class="col-2">
                        <h4 class="eq-header text-center">
                          Slots
                        </h4>
                        <inventory-slot-calculator
                          class="mt-1"
                          :imageSize="37"
                          :show-text-top="false"
                          :centered-buttons="false"
                          @input="item.slots = parseInt($event)"
                          :mask="item.slots"
                        />
                      </div>

                      <!-- Model Preview -->
                      <div
                        class="col-2"
                        style="text-align: center"
                      >
                        <h4 class="eq-header text-center">
                          Model
                        </h4>

                        <div @mouseover="drawItemModelSelector">
                          <item-model-preview
                            :id="item.idfile"
                          />

                          Item Model
                          <b-form-input
                            v-model.number="item.idfile"
                          />
                        </div>

                        <div @mouseover="drawIconSelector" class="mt-3">
                          <div>
                          <span
                            :class="'fade-in item-' + item.icon"
                            style="border: 1px solid rgb(218 218 218 / 30%); border-radius: 7px;"
                          />
                          </div>

                          Icon
                          <b-form-input v-model.number="item.icon"/>
                        </div>

                      </div>

                    </div>
                  </div>
                </eq-tab>

                <eq-tab name="Stats" class="stats-tab">

                  <div class="row">

                    <div class="col-4">

                      <!-- Stats -->
                      <div class="col-12 text-center">

                        <div
                          class="row"
                          :key="field.field"
                          v-for="field in
                       [
                         {
                           description: 'AC',
                           field: 'ac'
                         },
                         {
                           description: 'HP',
                           field: 'hp',
                           regen: 'regen'
                         },
                         {
                           description: 'Mana',
                           field: 'mana',
                           regen: 'manaregen',
                         },
                         {
                           description: 'Endur',
                           field: 'endur',
                           regen: 'enduranceregen'
                         },
                       ]"
                        >
                          <div class="col-4 text-right mr-3 p-0 mt-2">
                            {{ field.description }}
                          </div>
                          <div class="col-7 p-0 m-0" :style="(item[field.field] === 0 ? 'opacity: .5' : '')">
                            <b-form-input v-model.number="item[field.field]"/>
                          </div>
                        </div>

                        <div
                          v-for="(stat, description) in stats"
                          :key="stat.stat"
                          class="row text-center"
                        >
                          <div class="col-4 text-right mr-3 p-0 mt-2">
                            {{ description }}
                          </div>
                          <div class="col-3 p-0 m-0" :style="(item[stat.stat] === 0 ? 'opacity: .5' : '')">
                            <b-form-input v-model.number="item[stat.stat]"/>
                          </div>
                          <div class="col-1 p-0 m-0 mt-2">
                            +
                          </div>
                          <div class="col-3 p-0 m-0" :style="(item[stat.heroic] === 0 ? 'opacity: .5' : '')">
                            <b-form-input v-model.number="item[stat.heroic]"/>
                          </div>
                        </div>
                      </div>

                    </div>

                    <div class="col-4">
                      <div class="col-12">
                        <div
                          class="row"
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
                         {
                           description: 'Extra Damage Skill',
                           field: 'extradmgskill'
                         },
                         {
                           description: 'Extra Damage Amount',
                           field: 'extradmgamt'
                         },
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
                         {
                           description: 'Elemental Damage Amount',
                           field: 'elemdmgamt'
                         },
                         {
                           description: 'Element Damage Type',
                           field: 'elemdmgtype'
                         },
                       ]"
                        >
                          <div class="col-8 text-right mt-2 p-0 pr-3">
                            {{ field.description }}
                          </div>
                          <div class="col-4 m-0 p-0" :style="(item[field.field] === 0 ? 'opacity: .5' : '')">
                            <b-form-input v-model.number="item[field.field]"/>
                          </div>
                        </div>
                      </div>
                    </div>

                    <div class="col-4">
                      <div v-for="(field, description) in mod3" :key="field" class="row text-center">
                        <div class="col-7 text-right mt-2 p-0 pr-3">
                          {{ description }}
                        </div>
                        <div class="col-3 m-0 p-0" :style="(parseInt(item[field]) === 0 ? 'opacity: .5' : '')">
                          <b-form-input v-model.number="item[field]" v-if="field !== 'combateffects'"/>
                          <!-- For some reason combateffects is a varchar field -->
                          <b-form-input v-model="item[field]" v-if="field === 'combateffects'"/>
                        </div>
                      </div>
                    </div>

                  </div>


                </eq-tab>

                <eq-tab name="Effects">
                  <div>
                    <b-input-group style="height: 30px; margin-bottom: 15px">
                      <template #prepend>
                        <b-input-group-text class="m-0" style="width: 100px; height: 38px">Effect Type
                        </b-input-group-text>
                      </template>

                      <!--                      <b-input-group-text class="m-0" placeholder="Spell" disabled style="width: 200px">Spell</b-input-group-text>-->
                      <b-form-input class="m-0" placeholder="Spell ID" disabled style="width: 100px"/>
                      <b-form-input class="m-0" placeholder="As Level" disabled/>
                      <b-form-input class="m-0" placeholder="Required Level" disabled/>
                    </b-input-group>

                    <b-input-group
                      :key="field.field"
                      :style="(item[field.effectField] <= 0 ? 'opacity: .5' : '')"
                      v-for="field in
                       [
                         {
                           effectType: 'Scroll',
                           effectField: 'scrolleffect',
                           asLevelField: 'scrolllevel',
                           reqLevelField: 'scrolllevel_2',
                         },
                         {
                           effectType: 'Click',
                           effectField: 'clickeffect',
                           asLevelField: 'clicklevel',
                           reqLevelField: 'clicklevel_2',
                         },
                         {
                           effectType: 'Proc',
                           effectField: 'proceffect',
                           asLevelField: 'proclevel',
                           reqLevelField: 'proclevel_2',
                         },
                         {
                           effectType: 'Focus',
                           effectField: 'focuseffect',
                           asLevelField: 'focuslevel',
                           reqLevelField: 'focuslevel_2',
                         },
                         {
                           effectType: 'Worn',
                           effectField: 'worneffect',
                           asLevelField: 'wornlevel',
                           reqLevelField: 'wornlevel_2',
                         },
                         {
                           effectType: 'Bard',
                           effectField: 'bardeffect',
                           asLevelField: 'bardlevel',
                           reqLevelField: 'bardlevel_2',
                         },
                       ]"
                    >

                      <template #prepend>
                        <b-input-group-text style="width: 100px; height: 38px; text-align: right;">{{
                            field.effectType
                          }}
                        </b-input-group-text>
                      </template>

                      <!--                      <b-input-group-text style="width: 200px; height: 38px">Spell</b-input-group-text>-->
                      <b-form-input
                        class="m-0"
                        placeholder="Spell ID"
                        style="width: 100px"
                        @mouseover="drawEffectSelector"
                        :id="field.effectField"
                        v-model.number="item[field.effectField]"
                      />
                      <b-form-input
                        class="m-0"
                        placeholder="As Level"
                        @change="syncEffects(field.asLevelField, field.reqLevelField)"
                        v-model.number="item[field.asLevelField]"
                      />
                      <b-form-input
                        class="m-0"
                        placeholder="Required Level"
                        @change="syncEffects(field.reqLevelField, field.asLevelField)"
                        v-model.number="item[field.reqLevelField]"
                      />
                    </b-input-group>

                    <div class="row mt-3" v-if="item.clickeffect > 0">
                      <div class="col-1">
                        <h6 class="eq-header mt-4 text-center">Click</h6>
                      </div>

                      <div class="col-3 text-center">
                        Item Click Charges
                        <b-input-group class="mt-3">
                          <b-input-group-append style="height: 38px">
                            <b-form-input v-model.number="item.maxcharges" id="maxcharges" style="width: 70px"/>

                            <b-form-select
                              v-model.number="item.maxcharges"
                              style="width: 140px"
                              @change="setFieldModifiedById('maxcharges')"
                            >
                              <b-form-select-option
                                variant="outline-warning"
                                :value="parseInt(e.value)"
                                v-for="(e, id) in [
                                  {desc: 'Infinite', value: -1},
                                  {desc: 'None', value: 0},
                                ]"
                                :key="e.value"
                              >{{ e.value }}
                                ({{ e.desc }})
                              </b-form-select-option>
                              <b-form-select-option
                                :value="i"
                                v-for="(i) in [1,2,3,4,5,6,7,8,9,10,12,13,14,15,17,18,20,30,40,50,60,100,250,500,750,1000]"
                                :key="i"
                              >{{ i }}
                              </b-form-select-option>
                            </b-form-select>
                          </b-input-group-append>
                        </b-input-group>
                      </div>

                      <div class="col-2 m-0 p-0 pl-3 text-center">
                        Cast Time (ms)
                        <b-form-input
                          v-model.number="item.casttime"
                          id="casttime"
                          class="mt-3"
                        />
                        ({{ msToTime(item.casttime) }})
                      </div>

                      <div class="col-2 m-0 p-0 pl-3 text-center">
                        Recast Time (ms)
                        <b-form-input
                          v-model.number="item.recastdelay"
                          id="recastdelay"
                          class="mt-3"
                        />
                        ({{ msToTime(item.recastdelay) }})
                      </div>

                      <div class="col-2 text-center">
                        Recast Type
                        <b-form-input
                          v-model.number="item.recasttype"
                          id="recasttype"
                          class="mt-3"
                        />
                      </div>
                      <div class="col-2 text-center">
                        Click Type

                        <select v-model.number="item.clicktype" class="form-control mt-3">
                          <option
                            v-for="(e, index) in [
                              {desc: 'None', value: 0},
                              {desc: 'Clickable from Inventory', value: 1},
                              {desc: 'Clickable from Inventory', value: 3},
                              {desc: 'Must Equip to Cast', value: 4},
                              {desc: 'Clickable from Inventory', value: 5},
                            ]"
                            :key="e.value"
                            :value="parseInt(e.value)"
                          >
                            {{ e.value }}) {{ e.desc }}
                          </option>
                        </select>

                      </div>
                    </div>
                  </div>
                </eq-tab>

                <eq-tab
                  name="Augmentation"
                >

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
                          :value="parseInt(index)"
                        >
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
                          :value="parseInt(index)"
                        >
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
                          :value="parseInt(index)"
                        >
                          {{ index }}) {{ value.name }}
                        </option>
                      </select>
                    </div>
                  </div>
                </eq-tab>

                <eq-tab name="Bag">
                  <div
                    class="row"
                    :key="field.field"
                    v-for="field in
                       [
                         {
                           description: 'Bag Size',
                           field: 'bagsize'
                         },
                         {
                           description: 'Bag Slots',
                           field: 'bagslots',
                         },
                         {
                           description: 'Bag Type',
                           field: 'bagtype',
                         },
                         {
                           description: 'Bag Weight Restriction',
                           field: 'bagwr',
                         },
                       ]"
                  >
                    <div class="col-5 text-right mr-3 p-0 mt-2">
                      {{ field.description }}
                    </div>
                    <div class="col-3 p-0 m-0" :style="(item[field.field] === 0 ? 'opacity: .5' : '')">
                      <b-form-input v-model.number="item[field.field]"/>
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

                <eq-tab name="Unimplemented" v-if="showUnknown">
                  <div class="row">
                    <div
                      class="col-2"
                      :key="field.field"
                      v-for="field in
                       [
                          'benefitflag',
                          'pendingloreflag',
                       ]"
                    >
                      {{ field }}
                      <b-form-input v-model.number="item[field]"/>
                    </div>
                  </div>
                </eq-tab>

                <eq-tab name="?" v-if="showUnknown">
                  <div class="row">
                    <div
                      class="col-2"
                      :key="field.field"
                      v-for="field in
                       [
                          'unk_012',
                          'unk_013',
                          'unk_014',
                          'unk_033',
                          'unk_054',
                          'unk_059',
                          'unk_060',
                          'unk_120',
                          'unk_121',
                          'unk_123',
                          'unk_124',
                          'unk_127',
                          'unk_132',
                          'unk_134',
                          'unk_137',
                          'unk_142',
                          'unk_147',
                          'unk_152',
                          'unk_157',
                          'unk_193',
                          'unk_214',
                          'unk_220',
                          'unk_221',
                          'unk_223',
                          'unk_224',
                          'unk_225',
                          'unk_226',
                          'unk_227',
                          'unk_228',
                          'unk_229',
                          'unk_230',
                          'unk_231',
                          'unk_232',
                          'unk_233',
                          'unk_234',
                          'unk_236',
                          'unk_237',
                          'unk_238',
                          'unk_239',
                          'unk_240',
                          'unk_241',
                          'wornunk_1',
                          'wornunk_2',
                          'wornunk_3',
                          'wornunk_4',
                          'wornunk_5',
                          'wornunk_6',
                          'wornunk_7',
                          'procunk_1',
                          'procunk_2',
                          'procunk_3',
                          'procunk_4',
                          'procunk_6',
                          'procunk_7',
                          'scrollunk_1',
                          'scrollunk_2',
                          'scrollunk_3',
                          'scrollunk_4',
                          'scrollunk_5',
                          'scrollunk_6',
                          'focusunk_1',
                          'focusunk_2',
                          'focusunk_3',
                          'focusunk_4',
                          'focusunk_5',
                          'focusunk_6',
                          'focusunk_7',
                          'clickunk_5',
                          'clickunk_6',
                          'clickunk_7',
                          'bardunk_1',
                          'bardunk_2',
                          'bardunk_3',
                          'bardunk_4',
                          'bardunk_5',
                          'bardunk_7',
                       ]"
                    >
                      {{ field }}
                      <b-form-input v-model.number="item[field]"/>
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

              <div class="row">
                <div class="col-10"></div>
                <div class="col-2 text-right">
                  Unknown
                  <eq-checkbox
                    class="mb-2 d-inline-block"
                    v-model="showUnknown"
                  />
                </div>
              </div>

            </eq-window>
          </div>

          <div class="col-5">

            <!-- preview item -->
            <eq-window
              style="margin-top: 30px; margin-right: 10px; width: auto;"
              :key="item.updatedAt"
              v-if="previewItemActive && item && item.id > 0"
            >
              <eq-item-preview
                :item-data="item"
              />
            </eq-window>

            <!-- item model selector -->
            <eq-window
              style="margin-top: 30px; margin-right: 10px; width: auto;"
              class="fade-in"
              v-if="itemModelSelectorActive && item"
            >

              <item-model-selector
                :selected-model="item.idfile"
                @input="item.idfile = $event"
              />
            </eq-window>

            <!-- icon selector -->
            <eq-window
              style="margin-top: 30px; margin-right: 10px; width: auto;"
              class="fade-in"
              v-if="iconSelectorActive"
            >

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
              v-if="freeIdSelectorActive"
            >

              <free-id-selector
                table-name="items"
                id-name="id"
                name-label="name"
                :with-reserved="true"
                @input="item.id = $event"
              />
            </eq-window>

            <!-- spell effect selector -->
            <div
              class="fade-in"
              v-if="spellEffectSelectorActive"
            >

              <spell-effect-selector
                @input="item[$event.field] = $event.spell.id; setFieldModifiedById($event.field)"
              />

            </div>

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

import {Sketch}            from 'vue-color'
import SpellEffectSelector from "../../components/tools/SpellEffectSelector";

const MILLISECONDS_BEFORE_WINDOW_RESET = 5000;

export default {
  name: "ItemEdit",
  components: {
    SpellEffectSelector,
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
    EqWindowFancy,
    Sketch
  },
  data() {
    return {
      item: null, // item record data
      originalItem: null, // item record data; used to reference original values in tools

      spellCdnUrl: App.ASSET_SPELL_ICONS_BASE_URL,

      // state, loaded or not
      loaded: true,

      // preview toggle bools
      previewItemActive: true,
      iconSelectorActive: false,
      itemModelSelectorActive: false,
      spellEffectSelectorActive: false,
      freeIdSelectorActive: false,

      // show unknown fields
      showUnknown: false,

      // used to track when the subselector tool window has last spawned a tool
      // this keeps from a subsequent hover redrawing another tool within a grace period defined by
      // MILLISECONDS_BEFORE_WINDOW_RESET
      lastResetTime: Date.now(),

      // notifications and errors during save
      notification: "",
      error: "",

      // constants
      DB_ITEM_MATERIAL: DB_ITEM_MATERIAL,
      DB_ITEM_AUG_RESTRICT: DB_ITEM_AUG_RESTRICT,
      DB_ITEM_CLASS: DB_ITEM_CLASS,
      DB_ITEM_TYPES: DB_ITEM_TYPES,
      AUG_TYPES: AUG_TYPES,

      // fields used in forms
      stats: {
        "Strength": { stat: "astr", heroic: "heroic_str" },
        "Stamina": { stat: "asta", heroic: "heroic_sta" },
        "Intelligence": { stat: "aint", heroic: "heroic_int" },
        "Wisdom": { stat: "awis", heroic: "heroic_wis" },
        "Agility": { stat: "aagi", heroic: "heroic_agi" },
        "Dexterity": { stat: "adex", heroic: "heroic_dex" },
        "Charisma": { stat: "acha", heroic: "heroic_cha" },
        "Magic Resist": { stat: "mr", heroic: "heroic_mr" },
        "Fire Resists": { stat: "fr", heroic: "heroic_fr" },
        "Cold Resist": { stat: "cr", heroic: "heroic_cr" },
        "Disease Resist": { stat: "dr", heroic: "heroic_dr" },
        "Poison Resist": { stat: "pr", heroic: "heroic_pr" },
        "Corruption": { stat: "svcorruption", heroic: "heroic_svcorrup" }
      },
      mod3: {
        "Attack": "attack",
        "HP Regen": "regen",
        "Mana Regen": "manaregen",
        "Endurance Regen": "enduranceregen",
        "Accuracy": "accuracy",
        "Avoidance": "avoidance",
        "Clairvoyance": "clairvoyance",
        "Combat Effects": "combateffects",
        "Damage Shield": "damageshield",
        "Damage Shield Mitigation": "dsmitigation",
        "DoT Shielding": "dotshielding",
        "Heal Amount": "healamt",
        "Shielding": "shielding",
        "Spell Shielding": "spellshield",
        "Strikethrough": "strikethrough",
        "Stun Resist": "stunresist",
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


    // reset state vars when we navigate away
    '$route'() {
      this.item         = null;
      this.originalItem = null;

      // reset state vars when we navigate away
      this.notification = ""
      this.resetFieldEditedStatus()
      this.resetPreviewComponents()

      // reload
      this.load()
    },

    // some item effect types have defaults when the effect is set
    // they only appear to use that type when the effect is non-zero
    'item.worneffect': function (newVal, oldVal) {
      if (newVal !== oldVal && this.item) {
        this.item.worntype = newVal > 0 ? 2 : 0
        console.log("worn type is [%s]", this.item.worntype)
      }
    },
    'item.focuseffect': function (newVal, oldVal) {
      if (newVal !== oldVal && this.item) {
        this.item.focustype = newVal > 0 ? 6 : 0
        console.log("focus type is [%s]", this.item.focustype)
      }
    },
    'item.scrolleffect': function (newVal, oldVal) {
      if (newVal !== oldVal && this.item) {
        this.item.scrolltype = newVal > 0 ? 7 : 0
        console.log("this.item.scrolltype type is [%s]", this.item.scrolltype)
      }
    },
    'item.bardeffect': function (newVal, oldVal) {
      if (newVal !== oldVal && this.item) {
        this.item.bardeffecttype = newVal > 0 ? 8 : 0
        console.log("this.item.bardeffecttype type is [%s]", this.item.bardeffecttype)
      }
    },
    'item.casttime': function (newVal, oldVal) {
      if (newVal !== oldVal && this.item) {
        this.item.casttime_ = this.item.casttime
        console.log("casttime_ is [%s]", this.item.casttime_)
      }
    },

    // when aug slot type is non-zero, set visible to 1
    'item.augslot_1_type': function (newVal, oldVal) {
      if (newVal !== oldVal && this.item) {
        this.item.augslot_1_visible = newVal > 0 ? 1 : 0
        console.log("this.item.augslot_1_visible is [%s]", this.item.augslot_1_visible)
      }
    },
    'item.augslot_2_type': function (newVal, oldVal) {
      if (newVal !== oldVal && this.item) {
        this.item.augslot_2_visible = newVal > 0 ? 1 : 0
        console.log("this.item.augslot_2_visible is [%s]", this.item.augslot_2_visible)
      }
    },
    'item.augslot_3_type': function (newVal, oldVal) {
      if (newVal !== oldVal && this.item) {
        this.item.augslot_3_visible = newVal > 0 ? 1 : 0
        console.log("this.item.augslot_3_visible is [%s]", this.item.augslot_3_visible)
      }
    },
    'item.augslot_4_type': function (newVal, oldVal) {
      if (newVal !== oldVal && this.item) {
        this.item.augslot_4_visible = newVal > 0 ? 1 : 0
        console.log("this.item.augslot_4_visible is [%s]", this.item.augslot_4_visible)
      }
    },
    'item.augslot_5_type': function (newVal, oldVal) {
      if (newVal !== oldVal && this.item) {
        this.item.augslot_5_visible = newVal > 0 ? 1 : 0
        console.log("this.item.augslot_5_visible is [%s]", this.item.augslot_5_visible)
      }
    },

    // setting updatedAt tricks the component into re-rendering
    item: {
      handler(val, oldVal) {
        if (this.item) {
          this.item.updatedAt = Date.now()

          // console.log(JSON.stringify(this.item))
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

    msToTime(ms) {
      let seconds = (ms / 1000).toFixed(1);
      let minutes = (ms / (1000 * 60)).toFixed(1);
      let hours   = (ms / (1000 * 60 * 60)).toFixed(1);
      let days    = (ms / (1000 * 60 * 60 * 24)).toFixed(1);
      if (seconds < 60) return seconds + " Sec";
      else if (minutes < 60) return minutes + " Min";
      else if (hours < 24) return hours + " Hrs";
      else return days + " Days"
    },

    // prefills required and as level if the other value is 0 to save extra effort
    syncEffects(source, destination) {
      if (this.item[destination] === 0) {
        this.item[destination] = this.item[source]
      }
    },

    setFieldModified(evt) {
      // border: 2px #555555 solid !important;
      evt.target.style.setProperty('border-color', 'orange', 'important');
    },

    setFieldModifiedById(id) {
      const target = document.getElementById(id)
      if (target) {
        target.style.setProperty('border-color', 'orange', 'important');
      }
    },

    resetFieldEditedStatus() {
      // reset elements
      const itemEditCard = document.getElementById("item-edit-card")

      if (itemEditCard) {
        const elements = itemEditCard.querySelectorAll("input, select");
        elements.forEach((element) => {
          if (element) {
            element.style.setProperty('border-color', '#555555', 'important');
          }
        });
      }
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
          this.item              = result
          this.updatedAt         = Date.now()
          this.previewItemActive = true
        })
      }
    },

    /**
     * Selector / previewers
     */
    resetPreviewComponents() {
      this.freeIdSelectorActive      = false;
      this.iconSelectorActive        = false;
      this.itemModelSelectorActive   = false;
      this.previewItemActive         = false;
      this.spellEffectSelectorActive = false;
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
    drawEffectSelector() {
      this.resetPreviewComponents()
      this.spellEffectSelectorActive = true
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

.stats-tab input, .stats-tab select {
  margin-bottom: 7px;
  height: 30px;
}
</style>
