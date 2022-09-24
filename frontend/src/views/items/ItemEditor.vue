<template>
  <div>
    <div class="row">
      <div class="col-7">
        <eq-window>

          <div
            v-if="notification"
            :class="'text-center mt-2 btn-xs eq-header fade-in'"
            style="width: 100%; font-size: 30px"
            @click="notification = ''"
          >
            <i class="ra ra-shield mr-1"></i>
            {{ notification }}
          </div>

          <b-alert show dismissable variant="danger" v-if="error" class="mt-2">
            <i class="fa fa-warning"></i> {{ error }}
          </b-alert>

          <app-loader :is-loading="!item" class="mt-3 mb-3"/>

          <eq-tabs
            v-if="item"
            id="item-edit-card"
            class="item-edit-card"
            @mouseover.native="previewItem"
          >
            <eq-tab
              name="General"
              :selected="true"
            >
              <div class="row">
                <div
                  class="col-2"
                  @click="drawFreeIdSelector"
                  v-b-tooltip.hover.v-dark.right :title="getFieldDescription('id')"
                >
                  Id
                  <b-form-input
                    id="id" v-model.number="item.id"
                  />
                </div>
                <div class="col-4">
                  Name
                  <b-form-input
                    v-b-tooltip.hover.v-dark.right :title="getFieldDescription('name')"
                    :value="item.name" @change="v => item.name = v"
                  />
                </div>

                <!-- Item Type -->
                <div class="col-3">
                  Item Type
                  <select
                    v-b-tooltip.hover.v-dark.right :title="getFieldDescription('itemtype')"
                    v-model.number="item['itemtype']" class="form-control"
                  >
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
                <div class="col-3">
                  Item Class
                  <select
                    v-b-tooltip.hover.v-dark.right :title="getFieldDescription('itemclass')"
                    v-model.number="item['itemclass']" class="form-control"
                  >
                    <option
                      v-for="(description, index) in DB_ITEM_CLASS"
                      :key="index"
                      :value="parseInt(index)"
                    >
                      {{ index }}) {{ description }}
                    </option>
                  </select>
                </div>

              </div>

              <div class="row">

                <!-- Lore -->
                <div class="col-6">
                  Lore
                  <b-form-input
                    :value="item.lore" @change="v => item.lore = v"
                    v-b-tooltip.hover.v-dark.right :title="getFieldDescription('lore')"
                  />
                </div>

                <!-- Lore Group -->
                <div class="col-2">
                  Lore Group
                  <b-form-input
                    v-model.number="item.loregroup"
                    v-b-tooltip.hover.v-dark.right
                    :title="getFieldDescription('loregroup')"
                  />
                </div>

                <!-- Min Status -->
                <div class="col-2">
                  Min Status
                  <b-form-input
                    :style="(item['minstatus'] === 0 ? 'opacity: .5' : '')"
                    v-model.number="item.minstatus"
                    v-b-tooltip.hover.v-dark.right
                    :title="getFieldDescription('minstatus')"
                  />
                </div>

                <!-- Script File ID -->
                <div class="col-2">
                  Script File ID
                  <b-form-input
                    :style="(item['scriptfileid'] === 0 ? 'opacity: .5' : '')"
                    v-model.number="item.scriptfileid"
                    v-b-tooltip.hover.v-dark.right
                    :title="getFieldDescription('scriptfileid')"
                  />
                </div>

              </div>

              <div class="row">

                <!-- Stack Size -->
                <div class="col-2" :style="(item['stackable'] === 0 ? 'opacity: .5' : '')">
                  Stack Size
                  <b-form-input
                    v-model.number="item.stacksize"
                    v-b-tooltip.hover.v-dark.right
                    :title="getFieldDescription('stacksize')"
                  />
                </div>

                <!-- Item Size -->
                <div class="col-2">
                  Item Size
                  <b-form-select
                    v-model.number="item.size"
                    v-b-tooltip.hover.v-dark.right :title="getFieldDescription('size')"
                  >
                    <b-form-select-option
                      variant="outline-warning"
                      :value="parseInt(index)"
                      v-for="(value, index) in ITEM_SIZE"
                      :key="index"
                    >{{ index }})
                      {{ value }}
                    </b-form-select-option>
                  </b-form-select>
                </div>

                <!-- Weight -->
                <div class="col-2">
                  Weight
                  <b-form-input
                    v-model.number="item.weight"
                    v-b-tooltip.hover.v-dark.right
                    :title="getFieldDescription('weight')"
                  />
                </div>

                <!-- Recommended Level -->
                <div class="col-2" :style="(item['reclevel'] === 0 ? 'opacity: .5' : '')">
                  Rec. Level
                  <b-form-input
                    v-model.number="item.reclevel"
                    v-b-tooltip.hover.v-dark.right
                    :title="getFieldDescription('reclevel')"
                  />
                </div>

                <!-- Required Level -->
                <div class="col-2" :style="(item['reqlevel'] === 0 ? 'opacity: .5' : '')">
                  Required Level
                  <b-form-input
                    v-model.number="item.reqlevel"
                    v-b-tooltip.hover.v-dark.right
                    :title="getFieldDescription('reqlevel')"
                  />
                </div>

                <!-- Recommended Skill -->
                <div class="col-2" :style="(item['recskill'] === 0 ? 'opacity: .5' : '')">
                  Rec Skill

                  <b-form-input
                    v-model.number="item.recskill"
                    v-b-tooltip.hover.v-dark.right
                    :title="getFieldDescription('reqskill')"
                  />
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
                             {
                               description: 'Epic Item',
                               field: 'epicitem'
                             },
                             {
                               description: 'Arrow Expend',
                               field: 'expendablearrow'
                             },
                             {
                               description: 'Heirloom',
                               field: 'heirloom'
                             },
                           ]"
                      v-b-tooltip.hover.v-dark.right :title="getFieldDescription(field.field)"
                    >
                      <div class="col-9 text-right p-0 pr-2 m-0">
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
                      Classes
                    </h4>
                    <class-bitmask-calculator
                      style="border-radius: 15px; min-height: 400px;"
                      class="text-center mt-3"
                      :show-text-top="false"
                      :show-text-side="true"
                      :imageSize="38"
                      :centered-buttons="true"
                      @input="item.classes = parseInt($event); setFieldModifiedById('classes')"
                      :mask="item.classes"
                    />
                    <div class="text-center mt-3">
                      Classes
                      <b-input
                        id="classes"
                        v-model.number="item.classes"
                        v-b-tooltip.hover.v-dark.right
                        :title="getFieldDescription('classes')"
                      />
                    </div>
                  </div>
                  <div class="col-2">
                    <h4 class="eq-header text-center">
                      Races
                    </h4>
                    <race-bitmask-calculator
                      style="min-height: 400px;"
                      :imageSize="37"
                      class="mt-3"
                      :show-text-top="false"
                      :centered-buttons="true"
                      @input="item.races = parseInt($event); setFieldModifiedById('races')"
                      :mask="item.races"
                    />

                    <div class="text-center mt-3">
                      Races
                      <b-input
                        id="races" v-model.number="item.races"
                        v-b-tooltip.hover.v-dark.right :title="getFieldDescription('races')"
                      />
                    </div>
                  </div>
                  <div class="col-2">
                    <h4 class="eq-header text-center">
                      Deities
                    </h4>
                    <deity-bitmask-calculator
                      style="min-height: 400px;"
                      class="mt-3"
                      :imageSize="37"
                      :show-names="false"
                      :centered-buttons="true"
                      @input="item.deity = parseInt($event); setFieldModifiedById('deity')"
                      :mask="item.deity"
                    />

                    <div class="text-center mt-3">
                      Deities
                      <b-input
                        id="deity" v-model.number="item.deity"
                        v-b-tooltip.hover.v-dark.right :title="getFieldDescription('deity')"
                      />
                    </div>
                  </div>

                  <!-- Slots -->
                  <div class="col-2">
                    <h4 class="eq-header text-center">
                      Slots
                    </h4>
                    <inventory-slot-calculator
                      style="min-height: 400px;"
                      class="mt-1"
                      :imageSize="37"
                      :show-text-top="false"
                      :centered-buttons="false"
                      @input="item.slots = parseInt($event); setFieldModifiedById('slots')"
                      :mask="item.slots"
                    />

                    <div class="text-center mt-3">
                      Slots
                      <b-input
                        id="slots"
                        v-b-tooltip.hover.v-dark.right :title="getFieldDescription('slots')"
                        v-model.number="item.slots"
                      />
                    </div>
                  </div>

                  <!-- Model Preview -->
                  <div
                    class="col-2 minified-inputs"
                    style="text-align: center"
                  >
                    <h4 class="eq-header text-center">
                      Visuals
                    </h4>

                    <!-- item model -->
                    <div
                      @click="drawItemModelSelector()"
                    >
                      <item-model-preview
                        :id="item.idfile"
                        class="mb-2"
                      />

                      Item Model
                      <b-form-input
                        id="idfile"
                        v-model.number="item.idfile"
                      />
                    </div>

                    <!-- icon -->
                    Icon
                    <div
                      @click="drawIconSelector()"
                      class="row"
                    >
                      <div class="col-4">
                        <div class="d-inline-block" style="width: 50px">
                            <span
                              :class="'fade-in item-' + item.icon"
                              style="border: 1px solid rgb(218 218 218 / 30%); border-radius: 7px;"
                            />
                        </div>
                      </div>
                      <div class="col-8">
                        <b-form-input
                          style="height: 40px"
                          id="icon"
                          v-model.number="item.icon"
                        />
                      </div>
                    </div>

                    <!-- color -->
                    Color
                    <div
                      @click="drawColorSelector()"
                      class="row"
                    >
                      <div class="col-2">
                        <div
                          class="mr-3"
                          :style="'width: 25px; height: 25px; margin-top: 3px; border-radius: 5px; background-color: ' + hexColor"
                        />
                      </div>
                      <div class="col-10">
                        <b-form-input
                          disabled
                          id="color"
                          style="height: 30px; margin-top: 0px; margin-left: 5px"
                          v-model.number="item.color"
                        />
                      </div>
                    </div>

                    <!-- Material -->
                    <div
                      class="row"
                      @click="drawRaceMaterialPreview"
                    >
                      <div class="col-12">
                        Material
                        <select
                          v-model.number="item['material']"
                          @change="materialChange"
                          class="form-control"
                          id="material"
                        >
                          <option
                            v-for="(description, index) in DB_ITEM_MATERIAL"
                            :key="index"
                            :value="parseInt(index)"
                          >
                            {{ index }}) {{ description }}
                          </option>
                        </select>
                      </div>
                    </div>

                    <!-- Elite Material -->
                    <div class="row">
                      <div class="col-12">
                        Elite Material
                        <b-form-input
                          id="elitematerial"
                          v-model.number="item.elitematerial"
                        />
                      </div>
                    </div>

                    <!-- Light Emission -->
                    <div class="row">
                      <div class="col-12">
                        Light Emission
                        <b-form-select
                          v-model.number="item.light"
                        >
                          <b-form-select-option
                            :value="i"
                            v-for="(i) in [0,1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23,255]"
                            :key="i"
                          >{{ i }}
                          </b-form-select-option>
                        </b-form-select>
                      </div>
                    </div>

                    <!-- Hero Forge Model -->
                    <div class="row">
                      <div class="col-12">
                        Hero Forge Model
                        <b-form-input
                          id="herosforgemodel"
                          v-model.number="item.herosforgemodel"
                        />
                      </div>
                    </div>

                  </div>
                </div>
              </div>


            </eq-tab>

            <eq-tab name="Stats" class="minified-inputs">
              <div class="row">
                <div class="col-4">

                  <!-- Stats -->
                  <div class="col-12 text-center">
                    <div
                      class="row"
                      :key="field.field"
                      @mouseover="drawStatScaleTool"
                      style="border: 1px solid rgb(0,0,0, 0%)"
                      v-for="field in
                       [
                         {
                           description: 'AC',
                           field: 'ac'
                         },
                         {
                           description: 'HP',
                           field: 'hp',
                         },
                         {
                           description: 'Mana',
                           field: 'mana',
                         },
                         {
                           description: 'Endur',
                           field: 'endur',
                         },
                         {
                           description: 'Purity',
                           field: 'purity',
                         },
                       ]"
                    >
                      <div class="col-4 text-right mr-3 p-0 mt-2">
                        {{ field.description }}
                      </div>
                      <div class="col-7 p-0 m-0" :style="(item[field.field] === 0 ? 'opacity: .5' : '')">
                        <b-form-input
                          :id="field.field"
                          v-model.number="item[field.field]"
                        />
                      </div>
                    </div>

                    <div
                      v-for="(stat, description) in stats"
                      :key="stat.stat"
                      style="border: 1px solid rgb(0,0,0, 0%)"
                      class="row text-center"
                    >
                      <div class="col-4 text-right mr-3 p-0 mt-2">
                        {{ description }}
                      </div>
                      <div class="col-3 p-0 m-0" :style="(item[stat.stat] === 0 ? 'opacity: .5' : '')">
                        <b-form-input
                          :id="stat.stat"
                          v-model.number="item[stat.stat]"
                        />
                      </div>
                      <div class="col-1 p-0 m-0 mt-2">
                        +
                      </div>
                      <div class="col-3 p-0 m-0" :style="(item[stat.heroic] === 0 ? 'opacity: .5' : '')">
                        <b-form-input
                          :id="stat.heroic"
                          v-model.number="item[stat.heroic]"
                        />
                      </div>
                    </div>
                  </div>

                </div>

                <div class="col-4">
                  <div class="col-12">
                    <div
                      class="row"
                      :key="field.field"
                      @mouseover="drawStatScaleTool"
                      style="border: 1px solid rgb(0,0,0, 0%)"
                      v-for="field in damageStats"
                    >
                      <div class="col-8 text-right mt-1 mb-1 p-0 pr-3">
                        {{ field.description }}
                      </div>
                      <div
                        v-b-tooltip.hover.v-dark.right :title="getFieldDescription(field.field)"
                        class="col-4 m-0 p-0" :style="(item[field.field] <= 0 ? 'opacity: .5' : '')"
                      >

                        <b-form-input
                          :id="field.field"
                          v-model.number="item[field.field]"
                          v-if="!field.selectData"
                        />

                        <b-form-select
                          v-if="field.selectData"
                          v-model.number="item[field.field]"
                          class="form-control"
                        >
                          <option
                            v-for="(value, key) in field.selectData"
                            :key="key"
                            :value="parseInt(key)"
                          >
                            {{ key }}) {{ value }}
                          </option>
                        </b-form-select>

                      </div>
                    </div>
                  </div>
                </div>

                <div class="col-4">
                  <div
                    v-for="(field, description) in mod3"
                    :key="field"
                    @mouseover="drawStatScaleTool"
                    style="border: 1px solid rgb(0,0,0, 0%)"
                    class="row text-center"
                  >
                    <div class="col-7 text-right mt-1 mb-1 p-0 pr-3">
                      {{ description }}
                    </div>
                    <div
                      class="col-3 m-0 p-0"
                      v-b-tooltip.hover.v-dark.right :title="getFieldDescription(field)"
                      :style="(parseInt(item[field]) === 0 ? 'opacity: .5' : '')"
                    >
                      <b-form-input :id="field" v-model.number="item[field]" v-if="field !== 'combateffects'"/>
                      <!-- For some reason combateffects is a varchar field -->
                      <b-form-input :id="field" v-model="item[field]" v-if="field === 'combateffects'"/>
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
                    v-b-tooltip.hover.v-dark.right :title="getFieldDescription(field.effectField)"
                    style="width: 100px"
                    @mouseover="drawEffectSelector"
                    :id="field.effectField"
                    v-model.number="item[field.effectField]"
                  />
                  <b-form-input
                    class="m-0"
                    placeholder="As Level"
                    @change="syncEffects(field.asLevelField, field.reqLevelField)"
                    v-b-tooltip.hover.v-dark.right :title="getFieldDescription(field.asLevelField)"
                    v-model.number="item[field.asLevelField]"
                  />
                  <b-form-input
                    class="m-0"
                    placeholder="Required Level"
                    @change="syncEffects(field.reqLevelField, field.asLevelField)"
                    v-b-tooltip.hover.v-dark.right :title="getFieldDescription(field.reqLevelField)"
                    v-model.number="item[field.reqLevelField]"
                  />
                </b-input-group>

                <!-- Proc -->
                <div class="row mt-3" v-if="item.proceffect > 0">
                  <div class="col-1">
                    <h6 class="eq-header mt-5 text-right">Proc</h6>
                  </div>

                  <div
                    v-b-tooltip.hover.v-dark.right :title="getFieldDescription('procrate')"
                    class="col-4 m-0 p-0 pl-3 text-center"
                  >
                    Proc Rate
                    <b-input-group>
                      <b-input-group-append>
                        <b-form-input
                          v-model.number="item.procrate"
                          id="procrate"
                          class="mt-3"
                        />
                        <b-form-select v-model.number="item.procrate" class="form-control mt-3">
                          <option
                            v-for="(e, index) in [
                                    {desc: '100%', value: 0},
                                    {desc: '150%', value: 50},
                                    {desc: '200%', value: 100},
                                    {desc: '250%', value: 150},
                                    {desc: '300%', value: 200},
                                    {desc: '350%', value: 250},
                                    {desc: '400%', value: 300},
                                    {desc: '450%', value: 350},
                                    {desc: '500%', value: 400},
                                    {desc: '550%', value: 450},
                                    {desc: '600%', value: 500},
                                    {desc: '650%', value: 550},
                                    {desc: '700%', value: 600},
                                    {desc: '750%', value: 650},
                                    {desc: '800%', value: 700},
                                    {desc: '850%', value: 750},
                                    {desc: '900%', value: 800},
                                    {desc: '950%', value: 850},
                                    {desc: '1000%', value: 900},
                                  ]"
                            :key="e.value"
                            :value="parseInt(e.value)"
                          >
                            {{ e.value }}) {{ e.desc }}
                          </option>
                        </b-form-select>
                      </b-input-group-append>
                    </b-input-group>

                  </div>
                </div>

                <!-- Click -->
                <div class="row mt-3" v-if="item.clickeffect > 0">
                  <div class="col-1">
                    <h6 class="eq-header mt-5 text-right">Click</h6>
                  </div>

                  <div class="col-3 text-center">
                    Item Click Charges
                    <b-input-group class="mt-3">
                      <b-input-group-append style="height: 38px">
                        <b-form-input v-model.number="item.maxcharges" id="maxcharges" style="width: 70px"/>

                        <b-form-select
                          v-model.number="item.maxcharges"
                          id="maxcharges"
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

                    <loader-cast-bar-timer
                      class="mt-3 mb-3"
                      style="margin-left: 20px"
                      color="#FF00FF"
                      :time-ms="item.casttime"
                    />

                    ({{ msToTime(item.casttime) }})
                  </div>

                  <div class="col-2 m-0 p-0 pl-3 text-center">
                    Recast Time (seconds)
                    <b-form-input
                      v-model.number="item.recastdelay"
                      id="recastdelay"
                      class="mt-3"
                    />

                    <loader-cast-bar-timer
                      class="mt-3 mb-3"
                      style="margin-left: 20px"
                      color="#FF00FF"
                      :time-ms="(item.recastdelay * 1000)"
                    />

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

                    <select v-model.number="item.clicktype" id="clicktype" class="form-control mt-3">
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
              name="Aug"
              class="minified-inputs"
            >

              <div class="text-center mt-3 mb-3 font-weight-bold">Item Is Augment</div>

              <!-- Aug Type -->
              <div class="row" @mouseover="drawAugmentTypeCalculator">
                <div class="col-1">

                </div>
                <div class="col-4 text-right p-0 m-0 mt-2">
                  Augment Type
                </div>
                <div class="col-3">
                  <b-form-input
                    id="augtype"
                    v-b-tooltip.hover.v-dark.right
                    :title="getFieldDescription('augtype')"
                    :style="(item['augtype'] === 0 ? 'opacity: .5' : '')"
                    v-model.number="item['augtype']"
                  />
                </div>
              </div>

              <div class="row">
                <div class="col-1">

                </div>
                <div class="col-4 text-right p-0 m-0 mt-2">
                  Augment Restriction
                </div>
                <div class="col-3">
                  <select
                    v-model.number="item['augrestrict']"
                    class="form-control"
                    v-b-tooltip.hover.v-dark.right :title="getFieldDescription('augrestrict')"
                    :style="(item['augrestrict'] === 0 ? 'opacity: .5' : '')"
                  >
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

              <div class="text-center mt-3 mb-3 font-weight-bold">Item Can Have Augments</div>

              <!-- Aug Distiller Type -->
              <div class="row">
                <div class="col-1">

                </div>
                <div class="col-4 text-right p-0 m-0 mt-2">
                  Augment Distiller Item ID
                </div>
                <div class="col-3">
                  <b-form-input
                    :style="(item['augdistiller'] === 0 ? 'opacity: .5' : '')"
                    v-model.number="item['augdistiller']"
                    v-b-tooltip.hover.v-dark.right :title="getFieldDescription('augdistiller')"
                  />
                </div>
              </div>

              <!-- Aug Type -->
              <div class="row" v-for="i in 5" :key="i">
                <div class="col-1">

                </div>
                <div class="col-4 text-right p-0 m-0 mt-2">
                  Augment Slot {{ i }} Type
                </div>
                <div class="col-3">
                  <select
                    v-model.number="item['augslot_' + i + '_type']"
                    :style="(item['augslot_' + i + '_type'] === 0 ? 'opacity: .5' : '')"
                    class="form-control"
                    v-b-tooltip.hover.v-dark.right :title="getFieldDescription('augslot_' + i + '_type')"
                  >
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

            <eq-tab name="Bag" class="minified-inputs">
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
                           selectData: DB_BAG_TYPES,

                         },
                         {
                           description: 'Bag Weight Reduction',
                           field: 'bagwr',
                         },
                       ]"
              >
                <div class="col-5 text-right mr-3 p-0 mt-2">
                  {{ field.description }}
                </div>
                <div
                  class="col-3 p-0 m-0"
                  v-b-tooltip.hover.v-dark.right :title="getFieldDescription(field.field)"
                  :style="(item[field.field] === 0 ? 'opacity: .5' : '')"
                >
                  <b-form-input
                    v-if="!field.selectData"
                    :id="field.field"
                    v-model.number="item[field.field]"
                  />

                  <select
                    v-model.number="item[field.field]"
                    class="form-control"
                    v-if="field.selectData"
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

            <eq-tab name="Book" class="minified-inputs">
              <div
                class="row"
                :key="field.field"
                v-for="field in
                       [
                         {
                           description: 'Is Book',
                           field: 'book',
                           type: 'bool',
                         },
                         {
                           description: 'Book Type',
                           field: 'booktype',
                           selectData: BOOK_TYPES
                         },
                         {
                           description: 'Book (File) Name',
                           field: 'filename',
                         },
                       ]"
              >
                <div class="col-5 text-right mr-3 p-0 mt-2">
                  {{ field.description }}
                </div>
                <div class="col-3 p-0 m-0" :style="(item[field.field] === 0 ? 'opacity: .5' : '')">
                  <b-form-input
                    v-b-tooltip.hover.v-dark.right :title="getFieldDescription(field.field)"
                    v-if="!field.selectData && !field.type"
                    :id="field.field"
                    v-model.number="item[field.field]"
                  />

                  <eq-checkbox
                    v-b-tooltip.hover.v-dark.right :title="getFieldDescription(field.field)"
                    class="d-inline-block mt-2 mb-2"
                    :true-value="(typeof field.true !== 'undefined' ? field.true : 1)"
                    :false-value="(typeof field.false !== 'undefined' ? field.false : 0)"
                    v-model.number="item[field.field]"
                    @input="item[field.field] = $event"
                    v-if="field.type === 'bool'"
                  />

                  <select
                    v-b-tooltip.hover.v-dark.right :title="getFieldDescription(field.field)"
                    v-model.number="item[field.field]"
                    class="form-control"
                    v-if="field.selectData"
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

            <eq-tab name="Charm" class="minified-inputs">
              <div
                class="row"
                :key="field.field"
                v-for="field in
                       [
                         {
                           description: 'Charm File',
                           field: 'charmfile',
                         },
                         {
                           description: 'Charm File ID',
                           field: 'charmfileid'
                         },
                       ]"
              >
                <div class="col-5 text-right mr-3 p-0 mt-2">
                  {{ field.description }}
                </div>
                <div class="col-3 p-0 m-0" :style="(item[field.field] === 0 ? 'opacity: .5' : '')">
                  <b-form-input
                    v-b-tooltip.hover.v-dark.right :title="getFieldDescription(field.field)"
                    v-if="!field.selectData && !field.type"
                    :id="field.field"
                    v-model.number="item[field.field]"
                  />

                  <eq-checkbox
                    v-b-tooltip.hover.v-dark.right :title="getFieldDescription(field.field)"
                    class="d-inline-block mt-2 mb-2"
                    :true-value="(typeof field.true !== 'undefined' ? field.true : 1)"
                    :false-value="(typeof field.false !== 'undefined' ? field.false : 0)"
                    v-model.number="item[field.field]"
                    @input="item[field.field] = $event"
                    v-if="field.type === 'bool'"
                  />

                  <select
                    v-b-tooltip.hover.v-dark.right :title="getFieldDescription(field.field)"
                    v-model.number="item[field.field]"
                    class="form-control"
                    v-if="field.selectData"
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

            <eq-tab name="Faction" class="minified-inputs">
              <!-- Aug Type -->
              <div class="row" v-for="i in 4" :key="i">
                <div class="col-4 text-right mt-2 m-0 p-0">
                  Faction Mod {{ i }}
                </div>
                <div class="col-2 mr-0 pr-0" :style="(item['factionmod_' + i] === 0 ? 'opacity: .5' : '')">
                  <b-form-input
                    v-b-tooltip.hover.v-dark.right :title="getFieldDescription('factionmod_' + i)"
                    v-model.number="item['factionmod_' + i]"
                  />
                </div>
                <div class="col-1 text-center mt-2 m-0 p-0">
                  Amount
                </div>
                <div class="col-2" :style="(item['factionamt_' + i] === 0 ? 'opacity: .5' : '')">
                  <b-form-input
                    v-b-tooltip.hover.v-dark.right :title="getFieldDescription('factionamt_' + i)"
                    v-model.number="item['factionamt_' + i]"
                  />
                </div>
              </div>
            </eq-tab>

            <eq-tab name="Pricing" class="minified-inputs">
              <div
                v-for="(field, description) in pricingFields"
                :key="field"
                class="row text-center"
              >
                <div class="col-1">

                </div>
                <div class="col-4 text-right mt-1 mr-0 pr-0">
                  {{ description }}
                </div>
                <div
                  v-b-tooltip.hover.v-dark.right :title="getFieldDescription(field)"
                  class="col-2" :style="(item[field] === 0 ? 'opacity: .5' : '')"
                >
                  <b-form-input v-model.number="item[field]"/>
                  <!--                      <b-form-input v-model="item[field]"/>-->
                </div>
              </div>
            </eq-tab>

            <eq-tab name="Meta">
              <div
                class="row"
                :key="field.field"
                v-for="field in
                       [
                         {
                           description: 'Item Creation Time',
                           field: 'created'
                         },
                         {
                           description: 'Item Updated Time',
                           field: 'updated',
                         },
                         {
                           description: 'Item Verified Time',
                           field: 'verified',
                         },
                         {
                           description: 'Data Source',
                           field: 'source',
                         },
                         {
                           description: 'Comment',
                           field: 'comment',
                         },
                         {
                           description: 'serialization',
                           field: 'serialization',
                         },
                         {
                           description: 'serialized',
                           field: 'serialized',
                         },
                       ]"
                v-b-tooltip.hover.v-dark.right :title="getFieldDescription(field.field)"
              >
                <div class="col-5 text-right mr-3 p-0 mt-2">
                  {{ field.description }}
                </div>
                <div class="col-3 p-0 m-0" :style="(item[field.field] === 0 ? 'opacity: .5' : '')">
                  <b-form-input
                    :id="field.field"
                    v-model.number="item[field.field]"
                  />
                </div>
              </div>
            </eq-tab>

            <eq-tab name="Unsup." v-if="showUnknown">
              <div class="text-center font-weight-bold mt-5 mb-5">
                These are fields that are identified but unsupported or unimplemented by the EverQuest Emulator
                Server Source (yet)
              </div>

              <div class="row">
                <div
                  class="col-2"
                  :key="field.field"
                  v-for="field in
                       [
                          'evoid',
                          'evoitem',
                          'evolvinglevel',
                          'evomax',
                          'benefitflag',
                          'pendingloreflag',
                          'lorefile',
                       ]"
                >
                  {{ field }}
                  <b-form-input v-model.number="item[field]"/>
                </div>
              </div>
            </eq-tab>

            <eq-tab name="?" v-if="showUnknown">
              <div class="text-center font-weight-bold mt-3 mb-3">
                These are fields that are unknown
              </div>

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


          <div class="text-center mt-3" v-if="item">

            <div
              :class="'text-center mt-2 btn-xs eq-button-fancy'"
              @click="saveItem()"
            >
              <i class="ra ra-shield mr-1"></i>
              Save Item
            </div>

            <!--                <b-button-->
            <!--                  class="btn-dark btn-sm btn-outline-warning"-->
            <!--                  @click="saveItem"-->
            <!--                >-->
            <!--                  <i class="ra ra-book mr-1"></i>-->
            <!--                  Save Item-->
            <!--                </b-button>-->
          </div>

          <div class="row" v-if="item">
            <div class="col-10"></div>
            <div class="col-2 text-right" title="Show unknown fields">
              Unknown
              <eq-checkbox
                class="mb-2 d-inline-block"
                v-model.number="showUnknown"
              />
            </div>
          </div>

        </eq-window>
      </div>

      <!-- Preview / Selector Pane -->
      <div class="col-5">

        <!-- Stat Scale Tool-->
        <div
          class="fade-in"
          style="margin-bottom: 55px"
          v-if="drawStatScaleToolActive"
        >
          <div class="row">
            <div class="col-6">
              <item-stat-scale-percentage
                v-if="originalItem"
                style="height: 100%"
                :original-item-data="originalItem"
                @field="item[$event.field] = $event.value; setFieldModifiedById($event.field)"
              />
            </div>
            <div class="col-6">
              <item-stat-scale-range
                v-if="originalItem"
                style="height: 100%"
                :original-item-data="originalItem"
                @field="item[$event.field] = $event.value; setFieldModifiedById($event.field)"
              />
            </div>
          </div>
        </div>

        <!-- preview item -->
        <eq-window
          style="margin-right: 10px; width: auto;"
          :key="item.updatedAt"
          v-if="previewItemActive && item && item.id > 0"
        >
          <eq-item-preview
            :item-data="item"
            :show-related-data="true"
          />
        </eq-window>

        <!-- item model selector -->
        <div
          class="fade-in"
          v-if="itemModelSelectorActive && item"
        >
          <item-model-selector
            :selected-model="item.idfile"
            :find-by-icon="item.icon"
            @input="item.idfile = $event; setFieldModifiedById('idfile')"
          />
        </div>

        <!-- augment type calculator -->
        <eq-window
          style=" margin-right: 10px; width: auto;"
          class="fade-in"
          title="Augment Type Selector"
          v-if="drawAugmentTypeCalculatorActive && item"
        >
          <aug-bitmask-calculator
            :inputData.sync="item.augtype"
            :mask="item.augtype"
          />
        </eq-window>

        <!-- race material preview -->
        <eq-window-simple
          style=" margin-right: 10px; width: auto;"
          class="fade-in text-center"
          v-if="drawRaceMaterialPreviewActive && item"
        >

          <item-material-preview
            v-if="item.material >= 0"
            :selected-material="item.material"
            @input="item.material = $event; setFieldModifiedById('material')"
          />
        </eq-window-simple>

        <!-- color selector -->
        <div
          style="margin-top: 600px; margin-right: 10px; width: auto;"
          class="fade-in"
          v-if="drawColorSelectorActive && item"
        >

          <item-color-selector
            style="width: 100%; margin: auto; text-align: center"
            :color="hexColor"
            @input="item.color = $event.color; hexColor = $event.hexColor; setFieldModifiedById('color')"
          />

        </div>

        <!-- icon selector -->
        <div
          class="fade-in"
          v-if="iconSelectorActive"
        >
          <item-icon-selector
            :selected-icon="item.icon"
            :find-by-model="item.idfile"
            @input="item.icon = $event; setFieldModifiedById('icon')"
          />
        </div>

        <!-- free id selector -->
        <eq-window-simple
          title="Free Item Ids"
          style="margin-right: 10px; width: auto;"
          class="fade-in"
          v-if="freeIdSelectorActive"
        >

          <free-id-selector
            table-name="items"
            id-name="id"
            name-label="name"
            :with-reserved="true"
            @input="item.id = $event; setFieldModifiedById('id')"
          />
        </eq-window-simple>

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
</template>

<script>
import EqWindowFancy           from "../../components/eq-ui/EQWindowFancy";
import EqWindow                from "../../components/eq-ui/EQWindow";
import EqTabs                  from "../../components/eq-ui/EQTabs";
import EqTab                   from "../../components/eq-ui/EQTab";
import EqItemPreview           from "../../components/preview/EQItemCardPreview";
import EqCheckbox              from "../../components/eq-ui/EQCheckbox";
import {
  SpireApiClient
}                              from "../../app/api/spire-api-client";
import FreeIdSelector          from "../../components/tools/FreeIdSelector";
import {
  Items
}                              from "../../app/items";
import {
  ItemApi
}                              from "../../app/api";
import ItemModelPreview        from "./components/ItemModelPreview";
import ItemModelSelector       from "../../components/selectors/ItemModelSelector";
import ItemIconSelector        from "./components/ItemIconSelector";
import ClassBitmaskCalculator  from "../../components/tools/ClassBitmaskCalculator";
import RaceBitmaskCalculator   from "../../components/tools/RaceBitmaskCalculator";
import DeityBitmaskCalculator  from "../../components/tools/DeityCalculator";
import {
  BOOK_TYPES,
  DB_BAG_TYPES,
  DB_ITEM_AUG_RESTRICT,
  DB_ITEM_CLASS,
  DB_ITEM_MATERIAL,
  DB_ITEM_TYPES,
  ITEM_SIZE
}                              from "../../app/constants/eq-item-constants";
import {
  AUG_TYPES
}                              from "../../app/constants/eq-aug-constants";
import InventorySlotCalculator from "../../components/tools/InventorySlotCalculator";

import SpellEffectSelector     from "./components/ItemSpellEffectSelector";
import {DB_SKILLS}             from "../../app/constants/eq-skill-constants";
import ItemStatScaleTool       from "./components/ItemStatScalePercentage";
import ItemStatScalePercentage from "./components/ItemStatScalePercentage";
import ItemStatScaleRange      from "./components/ItemStatScaleRange";
import ItemColorSelector       from "./components/ItemColorSelector";
import * as util               from "util";
import {RACES}                 from "../../app/constants/eq-race-constants";
import ItemMaterialPreview     from "./components/ItemMaterialPreview";
import {BODYTYPES}             from "../../app/constants/eq-bodytype-constants";
import {DB_SPELL_RESISTS}      from "../../app/constants/eq-spell-constants";
import {DB_ITEM_BARD_TYPE}     from "../../app/constants/eq-bard-types";
import AugBitmaskCalculator    from "../../components/tools/AugmentTypeCalculator";
import EqWindowSimple          from "../../components/eq-ui/EQWindowSimple";
import LoaderCastBarTimer      from "../../components/LoaderCastBarTimer";
import {EditFormFieldUtil}     from "../../app/forms/edit-form-field-util";
import {FreeIdFetcher}         from "../../app/free-id-fetcher";
import ContentArea             from "../../components/layout/ContentArea";

const MILLISECONDS_BEFORE_WINDOW_RESET = 5000;

export default {
  name: "ItemEdit",
  components: {
    ContentArea,
    LoaderCastBarTimer,
    EqWindowSimple,
    AugBitmaskCalculator,
    ItemMaterialPreview,
    ItemColorSelector,
    ItemStatScaleRange,
    ItemStatScalePercentage,
    ItemStatScaleTool,
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
    EqWindowFancy
  },
  data() {
    return {
      item: null, // item record data
      originalItem: {}, // item record data; used to reference original values in tools

      // state, loaded or not
      loaded: true,

      hexColor: "#FFFFFF",

      // preview toggle bools
      previewItemActive: true,
      iconSelectorActive: false,
      itemModelSelectorActive: false,
      spellEffectSelectorActive: false,
      freeIdSelectorActive: false,
      drawStatScaleToolActive: false,
      drawColorSelectorActive: false,
      drawRaceMaterialPreviewActive: false,
      drawAugmentTypeCalculatorActive: false,

      // show unknown fields
      showUnknown: 0,

      // used to track when the subselector tool window has last spawned a tool
      // this keeps from a subsequent hover redrawing another tool within a grace period defined by
      // MILLISECONDS_BEFORE_WINDOW_RESET
      lastResetTime: Date.now() - MILLISECONDS_BEFORE_WINDOW_RESET,

      // notifications and errors during save
      notification: "",
      error: "",

      // constants
      DB_ITEM_MATERIAL: DB_ITEM_MATERIAL,
      DB_ITEM_AUG_RESTRICT: DB_ITEM_AUG_RESTRICT,
      DB_ITEM_CLASS: DB_ITEM_CLASS,
      DB_ITEM_TYPES: DB_ITEM_TYPES,
      DB_ITEM_BARD_TYPE: DB_ITEM_BARD_TYPE,
      DB_SKILLS: DB_SKILLS,
      DB_BAG_TYPES: DB_BAG_TYPES,
      BOOK_TYPES: BOOK_TYPES,
      ITEM_SIZE: ITEM_SIZE,
      AUG_TYPES: AUG_TYPES,

      // fields used in forms
      stats: Items.getBasicStatAndResistFields(),
      mod3: Items.getMod3Fields(),

      damageStats: [
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
          description: 'Range',
          field: 'range'
        },
        {
          description: 'Extra Damage Skill',
          field: 'extradmgskill',
          selectData: DB_SKILLS
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
          description: 'Bane Damage Body',
          field: 'banedmgbody',
          selectData: BODYTYPES,
        },
        {
          description: 'Bane Damage Amount',
          field: 'banedmgamt'
        },
        {
          description: 'Bane Damage Race',
          field: 'banedmgrace',
          selectData: RACES,
        },
        {
          description: 'Bane Damage Race Amount',
          field: 'banedmgraceamt'
        },
        {
          description: 'Element Damage Type',
          field: 'elemdmgtype',
          selectData: DB_SPELL_RESISTS,
        },
        {
          description: 'Elemental Damage Amount',
          field: 'elemdmgamt'
        },
        {
          description: 'Bard Skill Type',
          field: 'bardtype',
          selectData: DB_ITEM_BARD_TYPE,
        },
        {
          description: 'Bard Value',
          field: 'bardvalue'
        },
        {
          description: 'Skill Mod Type',
          field: 'skillmodtype',
          selectData: DB_SKILLS,
        },
        {
          description: 'Skill Mod Value',
          field: 'skillmodvalue'
        },
        {
          description: 'Skill Mod Max',
          field: 'skillmodmax'
        },

      ],

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
      this.originalItem = {};

      // reset state vars when we navigate away
      this.notification = ""
      EditFormFieldUtil.resetFieldEditedStatus()
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
    'item.combateffects': function (newVal, oldVal) {
      if (newVal !== oldVal && this.item) {
        // have to cast this as string because we made this field varchar for wahtever reason
        this.item.combateffects = newVal.toString()
      }
    },
    'item.clickeffect': function (newVal, oldVal) {
      if (newVal !== oldVal && this.item) {

        // setting item to be a clicky for the first time, lets set some sane defaults
        // console.log(oldVal)
        if (newVal > 0 && oldVal === -1) {
          this.item.maxcharges  = -1
          this.item.casttime    = 3000
          this.item.recastdelay = 6
          this.item.clicktype   = 5

          setTimeout(() => {
            EditFormFieldUtil.setFieldModifiedById('maxcharges')
            EditFormFieldUtil.setFieldModifiedById('casttime')
            EditFormFieldUtil.setFieldModifiedById('recastdelay')
            EditFormFieldUtil.setFieldModifiedById('clicktype')

            this.sendNotification("Set clicky defaults...")
          }, 50)
        }
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
    this.load()
  },
  methods: {

    getFieldDescription(field) {
      return Items.getFieldDescription(field);
    },

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

    setFieldModifiedById(id) {
      EditFormFieldUtil.setFieldModifiedById(id)
    },

    setSubEditorFieldHighlights() {
      let hasSubEditorFields = [
        "id",
        "icon",
        "idfile",
        "material",
        "color",
        "augtype",
        "proceffect",
        "worneffect",
        "focuseffect",
        "scrolleffect",
        "clickeffect",
        "bardeffect"
      ]
      hasSubEditorFields.forEach((field) => {
        EditFormFieldUtil.setFieldHighlightHasSubEditor(field)
      })
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

    async saveItem() {
      this.error        = ""
      this.notification = ""

      const api = (new ItemApi(SpireApiClient.getOpenApiConfig()))
      api.updateItem({
        id: this.item.id,
        item: this.item
      }).then((result) => {
        if (result.status === 200) {
          Items.setItem(this.item.id, this.item) // update cache
          this.sendNotification("Item updated successfully!")
          EditFormFieldUtil.resetFieldEditedStatus()
        }

        if (result.data.error) {
          this.error = result.data.error
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

        const createRes = await api.createItem({
          item: this.item
        })

        if (createRes.status === 200) {
          this.sendNotification("Created new Item!")
          EditFormFieldUtil.resetFieldEditedStatus()
        }

        if (createRes.data.error) {
          this.error = result.data.error
        }
      })
    },

    load() {

      if (this.$route.params.id > 0) {
        this.error = ""
        Items.getItem(this.$route.params.id).then(async (result) => {
          this.item              = result
          this.updatedAt         = Date.now()
          this.previewItemActive = true

          Object.assign(this.originalItem, result);

          let hex = util.format("#%s", this.toHex(this.item.color))

          // color has RR GG BB format however it appears that opacity is first
          // opacity is pretty much always 1 or FF
          // FF RR GG BB
          // the selector is popping and pushing based on this

          if (hex.length === 9) {
            hex = hex.replace("#ff", '#')
          }

          console.log(
            "[color] color is [%s] hex is [%s] length [%s] rgba [%s]",
            this.item.color,
            hex,
            hex.length,
            this.hexToRgbA(hex)
          )

          this.hexColor = hex;

          // if we're cloning this item, automatically fetch an ID
          if (this.$route.query.hasOwnProperty("clone")) {
            const id = await FreeIdFetcher.get("items", "id", "name")
            if (id > 0) {
              EditFormFieldUtil.setFieldModifiedById('id')
              this.item.id = id
            }
          }

          // hooks
          setTimeout(() => {
            const target = document.getElementById("item-edit-card")
            if (target) {
              target.removeEventListener('input', EditFormFieldUtil.setFieldModified, true);
              target.addEventListener('input', EditFormFieldUtil.setFieldModified)
            }

            EditFormFieldUtil.resetFieldSubEditorHighlightedStatus()
            this.setSubEditorFieldHighlights()

          }, 100)

        })
      }
    },

    toHex(d, padding) {
      var hex = Number(d).toString(16);
      padding = typeof (padding) === "undefined" || padding === null ? padding = 2 : padding;

      while (hex.length < padding) {
        hex = "0" + hex;
      }

      return hex;
    },

    hexToRgbA(hexCode, opacity = 1) {
      let hex = hexCode.replace('#', '');

      if (hex.length === 3) {
        hex = `${hex[0]}${hex[0]}${hex[1]}${hex[1]}${hex[2]}${hex[2]}`;
      }

      const r = parseInt(hex.substring(0, 2), 16);
      const g = parseInt(hex.substring(2, 4), 16);
      const b = parseInt(hex.substring(4, 6), 16);

      /* Backward compatibility for whole number based opacity values. */
      if (opacity > 1 && opacity <= 100) {
        opacity = opacity / 100;
      }

      return `rgba(${r},${g},${b},${opacity})`;
    },

    /**
     * Selector / previewers
     */
    resetPreviewComponents() {
      EditFormFieldUtil.resetFieldSubEditorHighlightedStatus()
      this.setSubEditorFieldHighlights()

      this.freeIdSelectorActive            = false;
      this.iconSelectorActive              = false;
      this.itemModelSelectorActive         = false;
      this.previewItemActive               = false;
      this.spellEffectSelectorActive       = false;
      this.drawStatScaleToolActive         = false;
      this.drawColorSelectorActive         = false;
      this.drawRaceMaterialPreviewActive   = false;
      this.drawAugmentTypeCalculatorActive = false;
    },
    shouldReset() {
      return Date.now() - this.lastResetTime > MILLISECONDS_BEFORE_WINDOW_RESET;
    },

    previewItem() {
      if (!this.previewItemActive && this.shouldReset()) {
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
      if ((!this.itemModelSelectorActive)) {
        this.resetPreviewComponents()
        this.itemModelSelectorActive = true;
        this.lastResetTime           = Date.now()

        EditFormFieldUtil.setFieldSubEditorHighlightedById("idfile")
      }
    },
    drawIconSelector() {
      if (!this.freeIdSelectorActive) {
        this.resetPreviewComponents()
        this.iconSelectorActive = true;

        EditFormFieldUtil.setFieldSubEditorHighlightedById("icon")
      }
    },
    drawColorSelector() {
      if (!this.drawColorSelectorActive) {
        this.resetPreviewComponents()
        this.drawColorSelectorActive = true;
        this.lastResetTime           = Date.now()

        EditFormFieldUtil.setFieldSubEditorHighlightedById("color")
      }
    },
    drawRaceMaterialPreview() {
      if (!this.drawRaceMaterialPreviewActive) {
        this.resetPreviewComponents()
        this.drawRaceMaterialPreviewActive = true;
        this.lastResetTime                 = Date.now()

        EditFormFieldUtil.setFieldSubEditorHighlightedById("material")
      }
    },
    drawFreeIdSelector() {
      this.resetPreviewComponents()
      this.lastResetTime        = Date.now()
      this.freeIdSelectorActive = true

      EditFormFieldUtil.setFieldSubEditorHighlightedById("id")
    },
    drawStatScaleTool() {
      // this.resetPreviewComponents()
      this.lastResetTime = Date.now()
      this.resetPreviewComponents()
      this.drawStatScaleToolActive = true
      this.previewItemActive       = true
    },
    drawAugmentTypeCalculator() {
      this.lastResetTime = Date.now()
      this.resetPreviewComponents()
      this.drawAugmentTypeCalculatorActive = true

      EditFormFieldUtil.setFieldSubEditorHighlightedById("augtype")
    },

    materialChange() {
      this.lastResetTime                 = Date.now()
      this.drawRaceMaterialPreviewActive = true
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

.minified-inputs input, .minified-inputs select {
  margin-top: 3px;
  margin-bottom: 3px;
  height: 30px;
}
</style>
