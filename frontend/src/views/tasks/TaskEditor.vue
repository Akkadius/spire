<template>
  <div>
    <div class="row">
      <div :class="(task ? 'col-7' : 'col-12') + ' p-0'">
        <eq-window-simple
          title="Task Editor"
          v-if="tasks"
          @mouseover.native="previewTask()"
        >
          <div class="row">
            <div :class="(task ? 'col-4' : 'col-12') + ' p-0'">
              <!-- Task List -->
              <div style="" class="">

                <div class="text-center mt-3">
                  <div class="btn-group" role="group">
                    <b-button
                      @click="createTask()"
                      size="sm"
                      variant="outline-warning"
                    >
                      <i class="fa fa-plus mr-1"></i>
                      New
                    </b-button>

                    <b-button
                      @click="cloneTask()"
                      size="sm"
                      variant="outline-light"
                      v-if="selectedTask"
                    >
                      <i class="ra ra-double-team"></i>
                      Clone
                    </b-button>

                    <b-button
                      @click="deleteTask()"

                      size="sm"
                      variant="outline-danger"
                      v-if="selectedTask"
                    >
                      <i class="fa fa-trash"></i>
                      Delete
                    </b-button>
                  </div>
                </div>

                <b-input-group class="mt-3">
                  <b-form-input
                    type="text"
                    placeholder="Filter results by name..."
                    v-model="taskSearchFilter"
                    @keyup="filterResultsByName"
                    class="form-control"
                  />
                  <b-input-group-append>
                    <b-button
                      variant="warning"
                      style="padding: 0.3rem 0.75rem;"
                      @click="resetFilter()"
                    ><i class="fa fa-refresh"></i>
                    </b-button>
                  </b-input-group-append>
                </b-input-group>

                <select
                  id="task-list"
                  size="2"
                  v-model="selectedTask"
                  @change="selectTask()"
                  class="form-control eq-input eq p-1 mt-3 eq-dark-background"
                  style="overflow-x: scroll; height: 75vh; overflow-y: scroll"
                >
                  <option
                    v-for="task in filteredTasks"
                    :id="'task-entry-' + task.id"
                    :value="task.id"
                  >
                    ({{ task.id }}) {{ task.title }}
                  </option>
                </select>

                <div class="mt-3 text-center">
                  Showing {{ filteredTasks.length }} out of {{ tasks.length }} tasks
                </div>

              </div>
            </div>

            <!-- Task -->
            <div class="col-8 fade-in" id="my-form" v-if="task">

              <!--              <div-->
              <!--                v-if="notification"-->
              <!--                :class="'text-center mt-2 btn-xs eq-header fade-in'"-->
              <!--                style="width: 100%; font-size: 30px"-->
              <!--                @click="notification = ''"-->
              <!--              >-->
              <!--                <i class="ra ra-book mr-1"></i>-->
              <!--                {{ notification }}-->
              <!--              </div>-->

              <div style="height: 40px">
                <b-alert show dismissable variant="danger" v-if="error">
                  <div class="row" @click="error = ''">
                    <div class="col-11">
                      <i class="fa fa-warning"></i> {{ error }}
                    </div>
                    <div class="col-1 text-right">
                      <i class="fa fa-remove"></i>
                    </div>
                  </div>
                </b-alert>

                <b-alert show dismissable variant="warning" v-if="notification">
                  <div class="row" @click="notification = ''">
                    <div class="col-11">
                      <i class="fa fa-info-circle mr-3"></i> {{ notification }}
                    </div>
                    <div class="col-1 text-right">
                      <i class="fa fa-remove"></i>
                    </div>
                  </div>
                </b-alert>
              </div>

              <eq-tabs
                id="task-edit-window"
                class="minified-inputs ml-2"
              >
                <eq-tab
                  name="Task"
                  selected="true"
                >
                  <div class="row">
                    <div
                      v-for="field in
                     [
                       // {
                       //   description: 'Task ID',
                       //   field: 'id',
                       //   col: 'col-2',
                       //   itemIcon: '2275',
                       //   onclick: setSelectorActive,
                       // },
                       {
                         description: 'Title',
                         field: 'title',
                         itemIcon: '6840',
                         col: 'col-6',
                       },
                       {
                         description: 'Task Type',
                         field: 'type',
                         fieldType: 'select',
                         itemIcon: '2275',
                         col: 'col-6',
                         selectData: TASK_TYPES,
                         zeroValue: -1,
                       },
                       {
                         description: 'Task Description',
                         field: 'description',
                         itemIcon: '2275',
                         fieldType: 'textarea',
                         col: 'col-12',
                         onclick: setSelectorActive,
                       },
                       {
                         description: 'Duration Code',
                         field: 'duration_code',
                         itemIcon: '750',
                         fieldType: 'select',
                         selectData: TASK_DURATION_TYPES,
                         col: 'col-4',
                         zeroValue: -1,
                       },
                       {
                         description: 'Duration',
                         field: 'duration',
                         fieldType: 'text',
                         itemIcon: '750',
                         col: 'col-4',
                         zeroValue: -1,
                       },
                       {
                         description: 'Duration',
                         field: 'duration',
                         fieldType: 'select',
                         selectData: TASK_DURATION_HUMAN,
                         itemIcon: '750',
                         col: 'col-4',
                         zeroValue: -1,
                       },
                       {
                         description: 'Min Level',
                         field: 'minlevel',
                         itemIcon: '5885',
                         fieldType: 'text',
                         col: 'col-2',
                       },
                       {
                         description: 'Max Level',
                         itemIcon: '5885',
                         field: 'maxlevel',
                         fieldType: 'text',
                         col: 'col-2',
                       },
                       {
                         description: 'Lvl Spread',
                         itemIcon: '5885',
                         field: 'level_spread',
                         fieldType: 'text',
                         col: 'col-2',
                       },
                       {
                         description: 'Min Players',
                         field: 'min_players',
                         itemIcon: '6018',
                         fieldType: 'text',
                         col: 'col-2',
                       },
                       {
                         description: 'Max Players',
                         field: 'max_players',
                         itemIcon: '6018',
                         fieldType: 'text',
                         col: 'col-2',
                       },

                       {
                         description: 'Repeatable',
                         field: 'repeatable',
                         itemIcon: '2903',
                         fieldType: 'checkbox',
                         col: 'col-2',
                       },
                       {
                         description: 'Completion Emote',
                         field: 'completion_emote',
                         itemIcon: '653',
                         fieldType: 'textarea',
                         col: 'col-12',
                       },
                       {
                         description: 'Reward Type',
                         field: 'rewardmethod',
                         fieldType: 'select',
                         itemIcon: '3366',
                         selectData: TASK_REWARD_METHOD_TYPE,
                         zeroValue: -1,
                         col: 'col-3',
                       },
                       {
                         description: 'Reward Text',
                         field: 'reward',
                         fieldType: 'text',
                         itemIcon: '3366',
                         col: 'col-5',
                       },
                       {
                         description: 'Reward Item ID',
                         field: 'rewardid',
                         fieldType: 'text',
                         itemIcon: '3366',
                         col: 'col-4',
                         onclick: setSelectorActive,
                       },
                       {
                         description: 'EXP Reward',
                         field: 'xpreward',
                         itemIcon: '2045',
                         fieldType: 'text',
                         col: 'col-4'
                       },
                       {
                         description: 'Cash Reward',
                         field: 'cashreward',
                         itemIcon: '646',
                         fieldType: 'text',
                         col: 'col-4',
                       },
                       {
                         description: 'Faction Reward',
                         field: 'faction_reward',
                         itemIcon: '528',
                         fieldType: 'text',
                         col: 'col-4',
                       },
                       {
                         description: 'Reward Ebon Crystals',
                         field: 'reward_ebon_crystals',
                         fieldType: 'text',
                         itemIcon: '1535',
                         col: 'col-6',
                       },
                       {
                         description: 'Replay Timer Seconds',
                         field: 'replay_timer_seconds',
                         fieldType: 'text',
                         itemIcon: '750',
                         col: 'col-6',
                       },

                       {
                         description: 'Reward Radiant Crystals',
                         field: 'reward_radiant_crystals',
                         itemIcon: '1536',
                         fieldType: 'text',
                         col: 'col-6',
                       },
                       {
                         description: 'Request Timer Seconds',
                         field: 'request_timer_seconds',
                         fieldType: 'text',
                         itemIcon: '750',
                         col: 'col-6',
                       },

                     ]"
                      :class="field.col + ' mb-3 pl-1 pr-1'"
                    >
                      <!--                      <div class="text-center">-->
                      <div>
                        <span
                          v-if="field.itemIcon"
                          :class="'item-' + field.itemIcon + '-sm'"
                          style="display: inline-block"
                        />
                        {{ field.description }}
                        <span
                          v-if="(field.field.includes('timer_seconds') || field.field.includes('duration')) && task[field.field] > 0"
                          class="font-weight-bold"
                        >({{ secondsToHumanTime(task[field.field]) }})</span>
                      </div>

                      <!-- checkbox -->
                      <div class="text-center">
                        <eq-checkbox
                          v-b-tooltip.hover.v-dark.right :title="getFieldDescription(field.field)"
                          class="mb-1 mt-3 d-inline-block text-center"
                          :true-value="(typeof field.true !== 'undefined' ? field.true : 1)"
                          :false-value="(typeof field.false !== 'undefined' ? field.false : 0)"
                          v-model.number="task[field.field]"
                          @input="task[field.field] = $event"
                          v-if="field.fieldType === 'checkbox'"
                        />
                      </div>


                      <!-- input number -->
                      <b-form-input
                        v-if="field.fieldType === 'number'"
                        :id="field.field"
                        v-model.number="task[field.field]"
                        class="m-0 mt-1"
                        v-on="typeof field.onclick !== 'undefined' ? { click: () => field.onclick(field.field) } : {}"
                        v-b-tooltip.hover.v-dark.right :title="getFieldDescription(field.field)"
                        :style="(task[field.field] === 0 ? 'opacity: .5' : '')"
                      />

                      <!-- input text -->
                      <b-form-input
                        v-if="field.fieldType === 'text' || !field.fieldType"
                        :id="field.field"
                        v-model.number="task[field.field]"
                        class="m-0 mt-1"
                        v-on="typeof field.onclick !== 'undefined' ? { click: () => field.onclick(field.field) } : {}"
                        v-b-tooltip.hover.v-dark.right :title="getFieldDescription(field.field)"
                        :style="(task[field.field] <= (typeof field.zeroValue !== 'undefined' ? field.zeroValue : 0) ? 'opacity: .5' : '')"
                      />

                      <!-- textarea -->
                      <b-textarea
                        v-if="field.fieldType === 'textarea'"
                        :id="field.field"
                        v-model="task[field.field]"
                        class="m-0 mt-1"
                        rows="2"
                        max-rows="6"
                        v-on="typeof field.onclick !== 'undefined' ? { click: () => field.onclick(field.field) } : {}"
                        v-b-tooltip.hover.v-dark.right :title="getFieldDescription(field.field)"
                        :style="(task[field.field] === '' ? 'opacity: .5' : '') + ';'"
                      ></b-textarea>

                      <!-- select -->
                      <select
                        v-model.number="task[field.field]"
                        class="form-control m-0 mt-1"
                        v-if="field.selectData"
                        v-b-tooltip.hover.v-dark.right :title="getFieldDescription(field.field)"
                        :style="(task[field.field] <= (typeof field.zeroValue !== 'undefined' ? field.zeroValue : 0) ? 'opacity: .5' : '')"
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
                <eq-tab
                  name="Activities"
                >

                  <div>
                    <span class="font-weight-bold">Activities</span>

                    <div class="d-inline-block">
                      <b-button @click="createActivity()" size="sm" variant="outline-warning" class="ml-2">
                        <i class="fa fa-plus"></i>
                      </b-button>
                      <b-button @click="cloneActivity()" size="sm" variant="outline-white" class="ml-2">
                        <i class="ra ra-double-team"></i>
                      </b-button>
                      <b-button @click="deleteActivity()" size="sm" variant="outline-danger" class="ml-2">
                        <i class="fa fa-trash"></i>
                      </b-button>
                      <b-button @click="moveActivityUp()" size="sm" variant="outline-primary" class="ml-2">
                        <i class="fa fa-arrow-up"></i>
                      </b-button>
                      <b-button @click="moveActivityDown()" size="sm" variant="outline-primary" class="ml-2">
                        <i class="fa fa-arrow-down"></i>
                      </b-button>
                    </div>

                    <select
                      size="2"
                      id="activities-scroll"
                      v-model="selectedActivity"
                      v-bind="task.task_activities"
                      @mouseover="scrollToActivity"
                      @change="updateQueryState"
                      ignore-input-change="1"
                      class="form-control eq-input"
                      style="overflow-x: scroll; min-height: 20vh; overflow-y: scroll"
                    >
                      <option
                        v-for="activity in task.task_activities"
                        :value="activity.activityid"
                        :id="'activities-entry' + activity.activityid"
                      >
                        Step [{{ activity.step }}] Activity [{{ activity.activityid }}]
                        {{ buildActivityDescription(activity) }}
                      </option>
                    </select>

                    <div v-if="selectedActivity !== null && task && task.task_activities && task.task_activities[selectedActivity]">
                      <div class="row mt-3">
                        <div
                          v-for="field in
                         [
                           // User should probably not be able to manipulate this
                           // {
                           //   description: 'Activity ID',
                           //   field: 'activityid',
                           //   col: 'col-6',
                           //   zeroValue: -1
                           // },
                           {
                             description: 'Task Step',
                             field: 'step',
                             fieldType: 'step',
                             itemIcon: '5739',
                             col: 'col-2',
                             zeroValue: -1
                           },
                           {
                             description: 'Activity Type',
                             field: 'activitytype',
                             itemIcon: '5739',
                             fieldType: 'select',
                             selectData: TASK_ACTIVITY_TYPES,
                             col: 'col-4',
                             onchange: activityTypeChange,
                           },
                           {
                             description: 'Activity Target',
                             itemIcon: '5739',
                             fieldType: 'text',
                             field: 'target_name',
                             col: 'col-6',
                             showIf: isActivityTargetVisible()
                           },
                           {
                             description: 'Item List',
                             itemIcon: '5739',
                             fieldType: 'text',
                             field: 'item_list',
                             col: 'col-12',
                             showIf: isItemListVisible()
                           },
                           {
                             description: 'Description Override',
                             field: 'description_override',
                             fieldType: 'text',
                             itemIcon: '2275',
                             col: 'col-12',
                           },
                           {
                             description: 'Goal Method',
                             field: 'goalmethod',
                             fieldType: 'select',
                             itemIcon: '3196',
                             selectData: TASK_GOAL_METHOD_TYPE,
                             zeroValue: -1,
                             col: 'col-4',
                           },
                           {
                             description: 'Goal Count',
                             itemIcon: '3196',
                             field: 'goalcount',
                             fieldType: 'number',
                             col: 'col-3',
                           },
                           {
                             description: 'Goal ID' + renderGoalIdDescriptor(),
                             field: 'goalid',
                             fieldType: 'number',
                             itemIcon: '3196',
                             col: 'col-3',
                             showIf: isGoalIdSelectorActive(),
                             onclick: isGoalIdSelectorActive() ? setSelectorActive : () => {},
                           },
                           {
                             description: 'Optional',
                             field: 'optional',
                             itemIcon: '4493',
                             fieldType: 'checkbox',
                             col: 'col-2',
                           },
                           {
                             description: 'Goal Match List ' + renderGoalMatchListDescription() + ' Multiple entries separated by |',
                             field: 'goal_match_list',
                             fieldType: 'text',
                             itemIcon: '3196',
                             col: 'col-12',
                             showIf: isGoalMatchListActive(),
                             onclick: setSelectorActive,
                           },
                           {
                             description: 'Quest Example (Quest Controlled)',
                             itemIcon: '3196',
                             field: 'quest_example',
                             fieldType: 'popout',
                             showIf: isActivityQuestControlled(),
                             col: 'col-12',
                           },
                           {
                             description: 'Deliver to NPC',
                             field: 'delivertonpc',
                             fieldType: 'number',
                             itemIcon: '5742',
                             col: 'col-6',
                             zeroValue: 0,
                             showIf: isDeliverToNPCActive(),
                             onclick: setSelectorActive,
                           },
                           {
                             description: 'Zone',
                             itemIcon: '3133',
                             field: 'zones',
                             fieldType: 'text',
                             type: 'text',
                             col: 'col-6',
                             onclick: setSelectorActive,
                           },
                           // Removed until fully implemented
                           // {
                           //   description: 'Skill List',
                           //   field: 'skill_list',
                           //   col: 'col-6',
                           // },
                           // {
                           //   description: 'Spell List',
                           //   field: 'spell_list',
                           //   col: 'col-12',
                           // },
                         ]"
                          :class="field.col + ' mb-3 pl-2 pr-2'"
                          v-if="(typeof field.showIf !== 'undefined' && field.showIf) || typeof field.showIf === 'undefined'"
                        >
                          <div>
                            <span
                              v-if="field.itemIcon"
                              :class="'item-' + field.itemIcon + '-sm'"
                              style="display: inline-block"
                            />
                            {{ field.description }}
                          </div>

                          <div v-if="field.fieldType === 'popout' && field.field === 'quest_example'">
                            <div>

                              <eq-tabs :bottom-tab-margin="10">
                                <eq-tab name="Perl" selected="true" class="mb-0">
                                  <div
                                    class="ml-0 code-display pl-0"
                                    style="width: 100%; display: inline-block; padding-top: 10px !important; padding-bottom: 10px !important; border-radius: 5px"
                                    v-b-tooltip.hover.v-dark.left
                                    :title="'$client->UpdateTaskActivity(int task_id, int activity_id, int count);'"
                                  >
                                    <button
                                      class='btn btn-sm btn-outline-warning mb-1 mr-2'
                                      @click="copyToClip(`$client->UpdateTaskActivity(${buildQuestUpdateTaskActivityParams()});`)"
                                      style="font-size: 8px; padding: 0.125rem 0.4rem; opacity: .6"
                                    >
                                      <i class="fa fa-clipboard"></i>
                                    </button>
                                    <span style="color: rgb(156, 220, 254);">$client-></span>UpdateTaskActivity({{
                                      buildQuestUpdateTaskActivityParams()
                                    }});
                                  </div>
                                </eq-tab>

                                <eq-tab name="Lua">
                                  <div
                                    class="ml-0 code-display pl-0"
                                    style="width: 100%; display: inline-block; padding-top: 10px !important; padding-bottom: 10px !important; border-radius: 5px"
                                    v-b-tooltip.hover.v-dark.left
                                    :title="'client:UpdateTaskActivity(int task, int activity, int count);'"
                                  >
                                    <button
                                      class='btn btn-sm btn-outline-warning mb-1 mr-2'
                                      @click="copyToClip(`client:UpdateTaskActivity(${buildQuestUpdateTaskActivityParams()});`)"
                                      style="font-size: 8px; padding: 0.125rem 0.4rem; opacity: .6"
                                    >
                                      <i class="fa fa-clipboard"></i>
                                    </button>
                                    <span style="color: rgb(156, 220, 254);">client:</span>UpdateTaskActivity({{
                                      buildQuestUpdateTaskActivityParams()
                                    }});
                                  </div>
                                </eq-tab>
                              </eq-tabs>

                            </div>
                          </div>

                          <!-- checkbox -->
                          <div class="text-center">
                            <eq-checkbox
                              v-b-tooltip.hover.v-dark.right :title="getFieldDescription(field.field)"
                              class="mb-1 mt-3 d-inline-block"
                              :true-value="(typeof field.true !== 'undefined' ? field.true : 1)"
                              :false-value="(typeof field.false !== 'undefined' ? field.false : 0)"
                              v-model.number="task.task_activities[selectedActivity][field.field]"
                              @input="task.task_activities[selectedActivity][field.field] = $event"
                              v-if="field.fieldType === 'checkbox'"
                            />
                          </div>

                          <!-- input number -->
                          <b-form-input
                            v-if="field.fieldType === 'number'"
                            :id="field.field"
                            v-model.number="task.task_activities[selectedActivity][field.field]"
                            class="m-0 mt-1"
                            v-on="typeof field.onclick !== 'undefined' ? { click: () => field.onclick(field.field) } : {}"
                            v-b-tooltip.hover.v-dark.right :title="getFieldDescription(field.field)"
                            :style="(task.task_activities[selectedActivity][field.field] === 0 ? 'opacity: .5' : '')"
                          />

                          <!-- input number step -->
                          <b-form-input
                            v-if="field.fieldType === 'step'"
                            type="number"
                            :id="field.field"
                            v-model.number="task.task_activities[selectedActivity][field.field]"
                            class="m-0 mt-1"
                            v-on="typeof field.onclick !== 'undefined' ? { click: () => field.onclick(field.field) } : {}"
                            v-b-tooltip.hover.v-dark.right :title="getFieldDescription(field.field)"
                            :style="(task.task_activities[selectedActivity][field.field] === 0 ? 'opacity: .5' : '')"
                          />

                          <!-- input text -->
                          <b-form-input
                            v-if="field.fieldType === 'text'"
                            :id="field.field"
                            v-model="task.task_activities[selectedActivity][field.field]"
                            class="m-0 mt-1"
                            v-on="typeof field.onclick !== 'undefined' ? { click: () => field.onclick(field.field) } : {}"
                            v-b-tooltip.hover.v-dark.right :title="getFieldDescription(field.field)"
                            :style="(task.task_activities[selectedActivity][field.field] <= (typeof field.zeroValue !== 'undefined' ? field.zeroValue : 0) ? 'opacity: .5' : '')"
                          />

                          <!-- textarea -->
                          <b-textarea
                            v-if="field.fieldType === 'textarea'"
                            :id="field.field"
                            v-model="task.task_activities[selectedActivity][field.field]"
                            class="m-0 mt-1"
                            rows="2"
                            max-rows="12"
                            v-b-tooltip.hover.v-dark.right :title="getFieldDescription(field.field)"
                            :style="(task.task_activities[selectedActivity][field.field] === '' ? 'opacity: .5' : '') + ';'"
                          ></b-textarea>

                          <!-- select -->
                          <select
                            v-model.number="task.task_activities[selectedActivity][field.field]"
                            class="form-control m-0 mt-1"
                            v-if="field.selectData"
                            v-b-tooltip.hover.v-dark.right :title="getFieldDescription(field.field)"
                            v-on="typeof field.onchange !== 'undefined' ? { change: () => field.onchange(field.field) } : {}"
                            :style="(task.task_activities[selectedActivity][field.field] <= (typeof field.zeroValue !== 'undefined' ? field.zeroValue : 0) ? 'opacity: .5' : '')"
                          >
                            <option
                              v-for="(description, index) in field.selectData"
                              :key="index"
                              :value="parseInt(index)"
                            >
                              {{ index }}) {{ description }}
                            </option>
                          </select>

                          <!-- Show zone name underneath zone field -->
                          <div
                            v-if="field.field === 'zones' && task.task_activities[selectedActivity][field.field]"
                            class="font-weight-bold mt-1 text-center"
                          >
                            {{ getZoneNames() }}
                          </div>

                        </div>
                      </div>
                      <eq-debug :data="task.task_activities[selectedActivity]"/>
                    </div>
                  </div>

                </eq-tab>
              </eq-tabs>

              <div class="text-center mt-3">
                <div class="btn-group" role="group">
                  <b-button
                    @click="saveTask()"
                    size="sm"
                    variant="outline-warning"
                  >
                    <i class="fa fa-save mr-1"></i>
                    Save
                  </b-button>
                </div>
              </div>

            </div>
          </div>

          <eq-debug :data="task"/>

        </eq-window-simple>
      </div>

      <div class="col-5 fade-in" v-if="task">

        <!-- goal match list previewer -->
        <div
          style="width: auto;"
          class="fade-in"
          v-if="selectorActive['goal_match_list'] && task.task_activities[selectedActivity] && typeof task.task_activities[selectedActivity].goal_match_list !== 'undefined'"
        >
          <task-goal-match-list-previewer
            :goal-match-list="task.task_activities[selectedActivity].goal_match_list"
            :activityType="task.task_activities[selectedActivity].activitytype"
            :zone-ids="task.task_activities[selectedActivity].zones.toString()"
          />
        </div>


        <!-- description selector -->
        <div
          style="width: auto;"
          class="fade-in"
          v-if="selectorActive['description']"
        >
          <task-description-selector
            :task="task"
            :description="task.description"
            @input="task['description'] = $event; setFieldModifiedById('description');"
          ></task-description-selector>
        </div>

        <!-- (rewardid) item selector -->
        <div
          style="width: auto;"
          class="fade-in"
          v-if="selectorActive['rewardid'] || selectorActive['reward']"
        >
          <task-item-selector
            :selected-item-id="task.rewardid > 0 ? task.rewardid : 0"
            @input="task['rewardid'] = $event.id; task['reward'] = $event.name; setFieldModifiedById('rewardid'); setFieldModifiedById('reward')"
          />
        </div>

        <!-- (goalid) item selector -->
        <div
          style="width: auto;"
          class="fade-in"
          v-if="selectorActive['goalid'] && isGoalIdItemSelectorActive() && task.task_activities[selectedActivity].goalmethod === 0"
        >
          <task-item-selector
            :selected-item-id="task.task_activities[selectedActivity].goalid ? task.task_activities[selectedActivity].goalid : 0"
            @input="task.task_activities[selectedActivity].goalid = $event.id; setFieldModifiedById('goalid'); postTargetNameUpdateProcessor($event, 'goalid')"
          />
        </div>

        <!-- (goalid) npc selector -->
        <div
          style="width: auto;"
          class="fade-in"
          v-if="selectorActive['goalid'] && isGoalIdNpcSelectorActive() && task.task_activities[selectedActivity].goalmethod === 0"
        >
          <task-npc-selector
            :selected-npc-id="task.task_activities[selectedActivity].goalid"
            @input="task.task_activities[selectedActivity].goalid = $event.npcId; setFieldModifiedById('goalid'); postTargetNameUpdateProcessor($event, 'goalid')"
          />
        </div>

        <!-- (goalid) explore selector -->
        <div
          style="width: auto;"
          class="fade-in"
          v-if="selectorActive['goalid'] && isGoalIdExploreActive() && task.task_activities[selectedActivity].goalmethod === 0"
        >
          <task-explore-selector
            :selected-explore-id="task.task_activities[selectedActivity].goalid"
            @input="task.task_activities[selectedActivity].goalid = $event.id; setFieldModifiedById('goalid'); postTargetNameUpdateProcessor($event, 'goalid')"
          />
        </div>

        <!-- (zones) Zone Selector -->
        <task-zone-selector
          :selected-zone-id="parseInt(task.task_activities[selectedActivity].zones)"
          v-if="task && task.task_activities && task.task_activities[selectedActivity] && selectorActive['zones']"
          @input="task.task_activities[selectedActivity].zones = $event.zoneId; setFieldModifiedById('zones')"
        />

        <!-- (delivertonpc) NPC Selector -->
        <task-npc-selector
          :selected-npc-id="task.task_activities[selectedActivity].delivertonpc"
          v-if="task && task.task_activities && task.task_activities[selectedActivity] && selectorActive['delivertonpc']"
          @input="task.task_activities[selectedActivity].delivertonpc = $event.npcId; setFieldModifiedById('delivertonpc'); postTargetNameUpdateProcessor($event, 'delivertonpc')"
        />

        <!-- (id) free id selector -->
        <eq-window-simple
          title="Free Item Ids"
          style="margin-right: 10px; width: auto;"
          class="fade-in"
          v-if="selectorActive['id']"
        >
          <free-id-selector
            table-name="tasks"
            id-name="id"
            name-label="title"
            :with-reserved="true"
            @input="task.id = $event; setFieldModifiedById('id')"
          />
        </eq-window-simple>

        <div v-if="previewTaskActive">
          <task-preview
            :task="task"
            :selected-activity="selectedActivity"
          />

          <eq-window-simple
            title="Chat (Completion Emote)"
            class="p-0"
            v-if="task.completion_emote !== ''"
          >
            <div
              class="mt-3 eq-background-dark p-2"
              style="border: rgba(122, 134, 183, 0.5) 1px solid; ;"
            >
            <span style="color: yellow">
              {{ task.completion_emote }}
            </span>
            </div>
          </eq-window-simple>

          <task-quest-example-preview
            :task="task"
            :selected-activity="selectedActivity"
          />
        </div>


      </div>
    </div>
  </div>
</template>

<script type="ts">
import EqWindow from "@/components/eq-ui/EQWindow.vue";
import ContentArea from "@/components/layout/ContentArea.vue";
import util from "util";
import {ROUTE} from "@/routes";
import {Tasks} from "@/app/tasks";
import EqCheckbox from "@/components/eq-ui/EQCheckbox.vue";
import {
  TASK_ACTIVITY_TYPE,
  TASK_ACTIVITY_TYPES,
  TASK_DURATION_HUMAN,
  TASK_DURATION_TYPES,
  TASK_GOAL_METHOD_TYPE,
  TASK_GOAL_METHOD_TYPES,
  TASK_REWARD_METHOD_TYPE,
  TASK_TYPES
} from "@/app/constants/eq-task-constants";
import EqWindowSimple from "@/components/eq-ui/EQWindowSimple.vue";
import EqDebug from "@/components/eq-ui/EQDebug.vue";
import EqTabs from "@/components/eq-ui/EQTabs.vue";
import EqTab from "@/components/eq-ui/EQTab.vue";
import {EditFormFieldUtil} from "@/app/forms/edit-form-field-util";
import TaskPreview from "@/views/tasks/components/TaskPreview.vue";
import TaskZoneSelector from "@/views/tasks/components/TaskZoneSelector.vue";
import TaskItemSelector from "@/views/tasks/components/TaskItemSelector.vue";
import FreeIdSelector from "@/components/tools/FreeIdSelector.vue";
import TaskNpcSelector from "@/views/tasks/components/TaskNpcSelector.vue";
import {FreeIdFetcher} from "@/app/free-id-fetcher";
import {Npcs} from "@/app/npcs";
import TaskExploreSelector from "@/views/tasks/components/TaskExploreSelector.vue";
import TaskDescriptionSelector from "@/views/tasks/components/TaskDescriptionSelector.vue";
import TaskGoalMatchListPreviewer from "@/views/tasks/components/TaskGoalMatchListPreviewer.vue";
import {Zones} from "@/app/zones";
import ClipBoard from "@/app/clipboard/clipboard";
import TaskQuestExamplePreview from "@/views/tasks/components/TaskQuestExamplePreview.vue";

const MILLISECONDS_BEFORE_WINDOW_RESET = 10000;

export default {
  components: {
    TaskQuestExamplePreview,
    TaskGoalMatchListPreviewer,
    TaskDescriptionSelector,
    TaskExploreSelector,
    TaskNpcSelector,
    FreeIdSelector,
    TaskItemSelector,
    TaskZoneSelector,
    TaskPreview,
    EqTab,
    EqTabs,
    EqDebug,
    EqWindowSimple,
    EqCheckbox,
    ContentArea,
    EqWindow,
  },

  data() {
    return {
      task: null,
      originalTask: {}, // hold original task data so we can determine when certain data has changed
      tasks: [],
      filteredTasks: [],
      selectedTask: null,
      selectedActivity: null,
      taskSearchFilter: "",

      // preview / selectors
      selectorActive: {},
      previewTaskActive: true,

      lastResetTime: Date.now(),

      scrolledToActivity: false,

      notification: "",
      error: "",

      // cached zone names for visual reference within the editor
      zoneNames: {},

      // constants
      TASK_ACTIVITY_TYPE: TASK_ACTIVITY_TYPE,
      TASK_REWARD_METHOD_TYPE: TASK_REWARD_METHOD_TYPE,
      TASK_TYPES: TASK_TYPES,
      TASK_DURATION_TYPES: TASK_DURATION_TYPES,
      TASK_DURATION_HUMAN: TASK_DURATION_HUMAN,
      TASK_ACTIVITY_TYPES: TASK_ACTIVITY_TYPES,
      TASK_GOAL_METHOD_TYPE: TASK_GOAL_METHOD_TYPE,
      TASK_GOAL_METHOD_TYPES: TASK_GOAL_METHOD_TYPES,
    }
  },

  watch: {
    $route(to, from) {
      if (Object.keys(this.$route.params).length === 0) {
        this.resetState()
      }
      this.init()
    },
  },

  methods: {

    isActivityQuestControlled() {
      return this.task.task_activities[this.selectedActivity].goalmethod === TASK_GOAL_METHOD_TYPES.QUEST_CONTROLLED
    },

    buildQuestUpdateTaskActivityParams() {
      return `${this.task.id}, ${this.selectedActivity}, 1`
    },

    copyToClip(s) {
      ClipBoard.copyFromText(s)

      this.$bvToast.toast(s, {
        title: "Copied to Clipboard!",
        autoHideDelay: 2000,
        solid: true
      })
    },

    getZoneNames() {
      const zones   = this.task.task_activities[this.selectedActivity].zones ? this.task.task_activities[this.selectedActivity].zones.toString() : ""
      let zoneNames = []
      for (let z of zones.split(",")) {
        zoneNames.push(
          this.zoneNames[z]
        )
      }

      return zoneNames.join(", ")
    },

    async fetchZoneNames() {
      let zoneNames = {}
      for (const z of (await Zones.getZones())) {
        zoneNames[z.zoneidnumber] = z.long_name
      }

      this.zoneNames = zoneNames
    },

    activityTypeChange() {
      if (this.task) {
        // console.log("type change fire")

        // reset certain fields when activity type is changed
        if (this.task.task_activities[this.selectedActivity]) {
          this.task.task_activities[this.selectedActivity].zones                = "0"
          this.task.task_activities[this.selectedActivity].goalid               = 0
          this.task.task_activities[this.selectedActivity].goalmethod           = 0
          this.task.task_activities[this.selectedActivity].goalcount            = 1
          this.task.task_activities[this.selectedActivity].optional             = 0
          this.task.task_activities[this.selectedActivity].item_list            = ""
          this.task.task_activities[this.selectedActivity].spell_list           = ""
          this.task.task_activities[this.selectedActivity].target_name          = ""
          this.task.task_activities[this.selectedActivity].description_override = ""
        }
      }
    },

    scrollToActivity() {
      // console.log("event firing")

      // activitiy list scroll to activity
      if (!this.scrolledToActivity) {
        if (this.selectedActivity) {
          const container = document.getElementById("activities-scroll");
          const target    = document.getElementById(util.format("task-entry-%s", this.selectedActivity))
          if (container && target) {
            container.scrollTo({top: target.offsetTop - 150, behavior: "smooth"});
          }
        }

        this.scrolledToActivity = true
      }
    },

    renderGoalIdDescriptor() {
      if (this.isGoalIdItemSelectorActive()) {
        return ' (Item)'
      }
      if (this.isGoalIdNpcSelectorActive()) {
        return ' (NPC)'
      }

      return ""
    },

    renderGoalMatchListDescription() {
      if (this.isGoalIdItemSelectorActive()) {
        return ' (Item) (List of item ids)'
      }
      if (this.isGoalIdNpcSelectorActive()) {
        return ' (NPC) (Can be NPC ID\'s or names)'
      }

      return ""
    },

    isGoalIdSelectorActive() {
      return this.task.task_activities && this.task.task_activities[this.selectedActivity] ? this.task.task_activities[this.selectedActivity].goalmethod === 0 : false
    },

    isGoalMatchListActive() {
      return this.task.task_activities[this.selectedActivity].goalmethod === TASK_GOAL_METHOD_TYPES.LIST
    },

    isDeliverToNPCActive() {
      return [
        TASK_ACTIVITY_TYPE.DELIVER,
        TASK_ACTIVITY_TYPE.GIVE,
      ].includes(
        parseInt(this.task.task_activities[this.selectedActivity].activitytype)
      )
    },

    isGoalIdItemSelectorActive() {
      return [
        TASK_ACTIVITY_TYPE.LOOT,
        TASK_ACTIVITY_TYPE.TRADESKILL,
        TASK_ACTIVITY_TYPE.DELIVER,
        TASK_ACTIVITY_TYPE.FISH,
        TASK_ACTIVITY_TYPE.FORAGE
      ].includes(
        parseInt(this.task.task_activities[this.selectedActivity].activitytype)
      )
    },

    isActivityTargetVisible() {
      return [
        TASK_ACTIVITY_TYPE.KILL,
        TASK_ACTIVITY_TYPE.LOOT,
        TASK_ACTIVITY_TYPE.DELIVER,
        TASK_ACTIVITY_TYPE.SPEAK_WITH,
        TASK_ACTIVITY_TYPE.EXPLORE,
      ].includes(
        parseInt(this.task.task_activities[this.selectedActivity].activitytype)
      )
    },

    isGoalIdNpcSelectorActive() {
      return [
        TASK_ACTIVITY_TYPE.KILL,
        TASK_ACTIVITY_TYPE.SPEAK_WITH,
        TASK_ACTIVITY_TYPE.GIVE
      ].includes(
        parseInt(this.task.task_activities[this.selectedActivity].activitytype)
      )
    },

    isGoalIdExploreActive() {
      return [
        TASK_ACTIVITY_TYPE.EXPLORE,
      ].includes(
        parseInt(this.task.task_activities[this.selectedActivity].activitytype)
      )
    },

    isItemListVisible() {
      return this.isGoalIdItemSelectorActive()
    },

    postTargetNameUpdateProcessor(event, fieldId) {
      const selectedActivity           = this.selectedActivity
      const updateType                 = this.task.task_activities[selectedActivity].activitytype ? this.task.task_activities[selectedActivity].activitytype : -1
      const isTargetNameEmpty          = typeof this.task.task_activities[selectedActivity].target_name !== "undefined" && this.task.task_activities[selectedActivity].target_name.trim().length === 0
      const isDescriptionOverrideEmpty = typeof this.task.task_activities[selectedActivity].description_override !== "undefined" && this.task.task_activities[selectedActivity].description_override.trim().length === 0

      if (isTargetNameEmpty) {
        //
      }

      if (updateType === TASK_ACTIVITY_TYPE.DELIVER && fieldId === 'delivertonpc') {
        this.task.task_activities[selectedActivity].target_name = Npcs.getCleanName(event.npc.name)
        EditFormFieldUtil.setFieldModifiedById('target_name');
      }

      if (isDescriptionOverrideEmpty && fieldId === 'goalid') {
        if (this.isGoalIdItemSelectorActive()) {
          this.task.task_activities[selectedActivity].item_list = event.name
          EditFormFieldUtil.setFieldModifiedById('item_list');
        }
        if (this.isGoalIdNpcSelectorActive()) {
          this.task.task_activities[selectedActivity].target_name = Npcs.getCleanName(event.npc.name)
          EditFormFieldUtil.setFieldModifiedById('target_name');
        }
      }
    },

    async moveActivityUp() {
      const selectedActivity = this.selectedActivity
      let activities         = []
      if (this.task.task_activities && this.task.task_activities.length > 0) {
        activities = JSON.parse(JSON.stringify(this.task.task_activities))
      }

      let matchedIndex = -1
      for (const index in activities) {
        if (parseInt(selectedActivity) === parseInt(index) && activities[index - 1]) {
          matchedIndex = parseInt(index)
        }
      }

      if (matchedIndex >= 0) {
        let currentActivity      = activities[matchedIndex]
        let previousActivity     = activities[matchedIndex - 1]
        const currentActivityId  = activities[matchedIndex].activityid
        const previousActivityId = activities[matchedIndex - 1].activityid
        const currentStep        = activities[matchedIndex].step
        const previousStep       = activities[matchedIndex - 1].step

        // shift activities between current and previous
        currentActivity.activityid  = previousActivityId
        previousActivity.activityid = currentActivityId

        // shift steps between current and previous
        currentActivity.step  = previousStep
        previousActivity.step = currentStep

        // write current activity using the previous activity object
        activities[matchedIndex]     = previousActivity
        // write previous activity using the current activity object
        activities[matchedIndex - 1] = currentActivity

        // set current selected to previous since we are shifting up
        this.selectedActivity = previousActivityId
      }

      // update current reactive property
      this.task.task_activities = JSON.parse(JSON.stringify(activities))

      await this.saveTask()

      // update the query state
      this.updateQueryState()
    },

    async cloneActivity() {
      try {
        const r = await Tasks.cloneTaskActivity(this.getBackendFormattedTask(), this.selectedActivity)
        if (r.status === 200) {
          this.task             = (await Tasks.getTask(this.$route.params.id))
          this.selectedActivity = r.data.activityid
          this.sendNotification(`Task activity (${this.selectedActivity}) successfully cloned`)
        }
      } catch (err) {
        console.log(err)
        if (err.response && err.response.data && err.response.data.error) {
          this.error = err.response.data.error
        }
      }
    },

    async moveActivityDown() {
      const selectedActivity = this.selectedActivity
      let activities         = []
      if (this.task.task_activities && this.task.task_activities.length > 0) {
        activities = JSON.parse(JSON.stringify(this.task.task_activities))
      }

      let matchedIndex = -1
      for (const index in activities) {
        console.log(index)
        if (parseInt(selectedActivity) === parseInt(index) && activities[parseInt(index) + 1]) {
          console.log("found matched index", index)
          matchedIndex = parseInt(index)
        }
      }

      if (matchedIndex >= 0) {
        let currentActivity     = activities[matchedIndex]
        let nextActivity        = activities[matchedIndex + 1]
        const currentActivityId = activities[matchedIndex].activityid
        const nextActivityId    = activities[matchedIndex + 1].activityid
        const currentStep       = activities[matchedIndex].step
        const nextStep          = activities[matchedIndex + 1].step

        // shift activities between current and previous
        currentActivity.activityid = nextActivityId
        nextActivity.activityid    = currentActivityId

        // shift steps between current and previous
        currentActivity.step = nextStep
        nextActivity.step    = currentStep

        // write current activity using the previous activity object
        activities[matchedIndex]     = nextActivity
        // write previous activity using the current activity object
        activities[matchedIndex + 1] = currentActivity

        // set current selected to previous since we are shifting up
        this.selectedActivity = nextActivityId
      }

      // update current reactive property
      this.task.task_activities = JSON.parse(JSON.stringify(activities))

      await this.saveTask()

      // update the query state
      this.updateQueryState()
    },

    async createActivity() {
      try {
        const r = await Tasks.createNewTaskActivity(this.task)
        if (r.status === 200) {
          this.task             = (await Tasks.getTask(this.$route.params.id))
          this.selectedActivity = r.data.activityid
          this.sendNotification("Task activity successfully created")
        }
      } catch (err) {
        console.log(err)
        if (err.response && err.response.data && err.response.data.error) {
          this.error = err.response.data.error
        }
      }
    },

    async deleteActivity() {
      if (!this.task.task_activities[this.selectedActivity]) {
        return
      }

      const activity = this.task.task_activities[this.selectedActivity]
      if (confirm(`Are you sure you want to delete this task activity?\n\n(Step ${activity.step}) Activity ${activity.activityid} \n\n` + this.buildActivityDescription(activity))) {
        try {
          let deletedSuccessfully = false
          if (this.task.task_activities) {

            for (const a of this.task.task_activities) {
              if (parseInt(a.activityid) === parseInt(this.selectedActivity)) {
                const r = await Tasks.deleteTaskActivity(a)
                if (r.status === 200) {
                  this.task             = (await Tasks.getTask(this.$route.params.id))
                  this.selectedActivity = a.activityid - 1
                  this.sendNotification("Task activity successfully deleted")
                  deletedSuccessfully = true
                }
              }
            }

            // // reorder task activities and then save again
            if (deletedSuccessfully) {
              let activityId = 0;
              for (const index in this.task.task_activities) {
                if (parseInt(this.task.task_activities[index].activityid) != activityId) {
                  const previousId                            = parseInt(this.task.task_activities[index].activityid)
                  this.task.task_activities[index].activityid = activityId
                  await Tasks.updateTaskActivityId(this.task.task_activities[index], previousId)

                  // console.log("Updating [%s] from [%s]", this.task.task_activities[index].activityid, previousId)
                }

                activityId++;
              }
            }

          }
        } catch (err) {
          if (err.response && err.response.data && err.response.data.error) {
            this.error = err.response.data.error
          }
          console.log(err)
        }
      }
    },

    sendNotification(message) {
      this.notification = message
      setTimeout(() => {
        this.notification = ""
      }, 5000)
    },

    getBackendFormattedTask() {
      let t          = this.task
      let activities = []
      if (t.task_activities) {
        t.task_activities.forEach((a) => {
          let activity   = a
          activity.zones = activity.zones.toString()
          activities.push(activity)
        })

        t.task_activities = activities
      }

      return t
    },

    async saveTask() {
      try {
        const r = await Tasks.updateTask(this.getBackendFormattedTask())
        if (r.status === 200) {
          EditFormFieldUtil.resetFieldEditedStatus()
          this.sendNotification("Task updated!");
        }
      } catch (err) {
        if (err.response && err.response.data && err.response.data.error) {
          this.error = err.response.data.error
        }
      }

      console.log("updated task")
    },

    async deleteTask() {
      if (confirm(`Are you sure you want to delete this task?\n\n(${this.task.id}) ${this.task.title} `)) {
        const r = await Tasks.deleteTaskWithActivities(this.task)
        if (r !== "") {
          this.error = r
        }

        this.resetStateAll()
        this.updateQueryState()
      }
    },

    async createTask() {
      const id = await FreeIdFetcher.get("tasks", "id", "title")
      if (id > 0) {
        EditFormFieldUtil.setFieldModifiedById('id')

        // relink id at the task level
        let newTask = Tasks.getExampleTask()
        newTask.id  = id

        // relink id at the activities level
        let activities = []
        if (newTask.task_activities) {
          newTask.task_activities.forEach((a) => {
            let activity    = a
            activity.taskid = id
            activities.push(activity)
          })

          newTask.task_activities = activities
        }

        try {
          const r = await Tasks.createTask(newTask)
          if (r.status === 200) {
            this.resetState()
            this.selectedTask     = r.data.id
            this.selectedActivity = 0
            this.tasks            = []
            setTimeout(() => {
              this.updateQueryState()
              this.sendNotification("New task created successfully!")
            }, 100)
          }
        } catch (err) {
          if (err.response && err.response.data && err.response.data.error) {
            this.error = err.response.data.error
          }
        }
      }
    },

    async cloneTask() {
      if (confirm(`Are you sure you want to copy this task?\n\n(${this.task.id}) ${this.task.title} `)) {
        const id = await FreeIdFetcher.get("tasks", "id", "title")
        if (id > 0) {
          EditFormFieldUtil.setFieldModifiedById('id')

          // fetch task only with activities
          const task = await Tasks.getTaskWithActivities(this.selectedTask)
          if (task.id > 0) {

            // relink id at the task level
            let newTask = task
            newTask.id  = id

            // relink id at the activities level
            let activities = []
            if (newTask.task_activities) {
              newTask.task_activities.forEach((a) => {
                let activity    = a
                activity.taskid = id
                activities.push(activity)
              })

              newTask.task_activities = activities
            }

            try {
              const r = await Tasks.createTask(newTask)
              if (r.status === 200) {
                this.resetState()
                this.selectedTask     = r.data.id
                this.selectedActivity = 0
                this.tasks            = []

                setTimeout(() => {
                  this.updateQueryState()
                  this.sendNotification("New task cloned successfully!")
                }, 100)
              }
            } catch (err) {
              if (err.response && err.response.data && err.response.data.error) {
                this.error = err.response.data.error
              }
            }
          }
        }
      }
    },

    secondsToHumanTime(seconds) {
      let levels     = [
        [Math.floor(seconds / 31536000), 'y'],
        [Math.floor((seconds % 31536000) / 86400), 'd'],
        [Math.floor(((seconds % 31536000) % 86400) / 3600), 'h'],
        [Math.floor((((seconds % 31536000) % 86400) % 3600) / 60), 'm'],
        [(((seconds % 31536000) % 86400) % 3600) % 60, 's'],
      ];
      let returntext = '';
      for (let i = 0, max = levels.length; i < max; i++) {
        if (levels[i][0] === 0) continue;
        returntext += ' ' + levels[i][0] + '' + (levels[i][1]);
      }
      return returntext.trim();
    },

    resetState() {
      this.task             = null;
      this.selectedTask     = null;
      this.selectedActivity = null;
      this.taskSearchFilter = "";
    },

    resetStateAll() {
      this.resetState()
      this.selectedActivity = 0
      this.tasks            = []
      this.notification     = ""
      this.error            = ""
    },

    setFieldModifiedById(id) {
      EditFormFieldUtil.setFieldModifiedById(id)
    },

    resetFilter() {
      this.filteredTasks    = this.tasks
      this.taskSearchFilter = ""
      this.updateQueryState()
    },

    getFieldDescription(field) {
      return Tasks.getFieldDescription(field);
    },

    async selectTask() {
      this.selectedActivity = 0
      this.updateQueryState()
    },

    async filterResultsByName() {
      let filteredTasks = [];
      filteredTasks     = this.tasks.filter((task) => {
        return task.title.toLowerCase().includes(this.taskSearchFilter.toLowerCase())
      })

      this.filteredTasks = filteredTasks;
    },

    async loadTask() {
      if (this.$route.params.id > 0) {

        // if we're trying to load the same selected task, bail out
        // keeps us from propagating tons of unnecessary updates (slow)
        if (this.task !== null && this.task && this.task.id && parseInt(this.task.id) === parseInt(this.$route.params.id)) {
          return
        }

        this.task = (await Tasks.getTask(this.$route.params.id))

        // only load these once
        if (Object.keys(this.zoneNames).length === 0) {
          this.fetchZoneNames().then(() => {
            this.$forceUpdate()
          })
        }

        // copy to original
        Object.assign(this.originalTask, this.task);

        setTimeout(() => {

          // task list scroll to task
          const container = document.getElementById("task-list");
          const target    = document.getElementById(util.format("task-entry-%s", this.$route.params.id))
          if (container && target) {
            // container.scrollTop = target.offsetTop - 100;
            container.scrollTo({top: target.offsetTop - 150, behavior: "smooth"});
          }

          // hooks
          setTimeout(() => {
            const target = document.getElementById("task-edit-window")
            if (target) {
              target.removeEventListener('input', EditFormFieldUtil.setFieldModified, true);
              target.addEventListener('input', EditFormFieldUtil.setFieldModified)
            }

            EditFormFieldUtil.resetFieldEditedStatus()
          }, 1)
        }, 1)


        // if no task found...
        // this.task = null
      }
    },
    setFieldHighlights() {
      let hasSubEditorFields = [
        "id",
        "description",
        "delivertonpc",
        "goal_match_list",
        "zones",
        // "item_list",
        // "skill_list",
        // "spell_list"
      ];
      hasSubEditorFields.forEach((field) => {
        EditFormFieldUtil.setFieldHighlightHasSubEditor(field)
      })

      if (this.task.rewardmethod === 0) {
        EditFormFieldUtil.setFieldHighlightHasSubEditor("rewardid")
        // EditFormFieldUtil.setFieldHighlightHasSubEditor("reward")
      }

      if (this.isGoalIdSelectorActive()) {
        EditFormFieldUtil.setFieldHighlightHasSubEditor("goalid")
      }
    },
    buildActivityDescription(activity) {
      return Tasks.buildActivityDescription(activity)
    },
    async loadTasks() {
      if (this.tasks.length > 0) {
        return
      }

      const tasks = await Tasks.getTasks()
      if (tasks.length > 0) {
        this.tasks         = tasks
        this.filteredTasks = tasks
        this.filterResultsByName()
      }
    },
    async init() {
      if (Object.keys(this.$route.query).length !== 0) {
        this.loadQueryState()
      }
      await this.loadTasks()
      this.loadTask().then(() => {
        this.previewTask(true)
        EditFormFieldUtil.resetFieldEditedStatus()
        EditFormFieldUtil.resetFieldSubEditorHighlightedStatus()
        EditFormFieldUtil.resetFieldHighlightHasSubEditorStatus()
        this.setFieldHighlights()
      })
    },

    updateQueryState() {
      // query params
      let queryState = {};
      if (this.selectedActivity >= 0) {
        queryState.activity = this.selectedActivity
      }
      if (this.taskSearchFilter !== "") {
        queryState.q = this.taskSearchFilter
      }

      // navigation
      this.$router.push(
        {
          path: this.selectedTask > 0 ? util.format(ROUTE.TASK_EDIT, this.selectedTask) : ROUTE.TASKS,
          query: queryState
        }
      ).catch(() => {
      })
    },

    loadQueryState() {
      if (parseInt(this.$route.params.id) >= 0) {
        this.selectedTask = parseInt(this.$route.params.id);
      }
      if (typeof this.$route.query.activity !== 'undefined' && parseInt(this.$route.query.activity) >= 0) {
        this.selectedActivity = parseInt(this.$route.query.activity);
      }
      if (typeof this.$route.query.q !== 'undefined' && this.$route.query.q) {
        this.taskSearchFilter = this.$route.query.q
      }
    },

    shouldReset() {
      return (Date.now() - this.lastResetTime) > MILLISECONDS_BEFORE_WINDOW_RESET
    },

    previewTask(force = false) {
      if ((this.shouldReset() && !this.previewTaskActive) || force) {
        this.resetPreviewComponents()
        this.previewTaskActive = true
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
      this.previewTaskActive        = false;
      this.lastResetTime            = Date.now()
      this.selectorActive[selector] = true
      this.$forceUpdate()

      EditFormFieldUtil.setFieldSubEditorHighlightedById(selector)
    }
  },
  async mounted() {
    await this.init()
  },
}

</script>

<style>
#task-list {
  overflow-y: scroll;
  height: 80vh;
  overflow-x: hidden;
  white-space: nowrap;
  border-radius: 5px;
}

</style>
