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
                      :disabled="saveFreeze"
                      variant="outline-warning btn-dark"
                    >
                      <i class="fa fa-plus mr-1"></i>
                      New
                    </b-button>

                    <b-button
                      @click="cloneTask()"
                      size="sm"
                      :disabled="saveFreeze"
                      variant="outline-light btn-dark"
                      v-if="selectedTask"
                    >
                      <i class="ra ra-double-team"></i>
                      Clone
                    </b-button>

                    <b-button
                      @click="deleteTask()"
                      :disabled="saveFreeze"
                      size="sm"
                      variant="outline-danger btn-dark"
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
                    :disabled="saveFreeze"
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
                  :disabled="saveFreeze"
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

                <!-- Notification / Error -->
                <info-error-banner
                  :notification="notification"
                  :error="error"
                  @dismiss-error="error = ''"
                  @dismiss-notification="notification = ''"
                  class="mt-0"
                />
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
                         description: 'Dynamic Zone Template ID',
                         field: 'dz_template_id',
                         onclick: setSelectorActive,
                         itemIcon: '4004',
                         fieldType: 'text',
                         col: 'col-6',
                         zeroValue: -1,
                       },
                       {
                         description: 'Lock Task on Activity ID',
                         field: 'lock_activity_id',
                         itemIcon: '1077',
                         fieldType: 'select',
                         selectData: buildTaskActivitySelection(),
                         col: 'col-6',
                         zeroValue: -1,
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
                         field: 'min_level',
                         itemIcon: '5885',
                         fieldType: 'text',
                         col: 'col-2',
                       },
                       {
                         description: 'Max Level',
                         itemIcon: '5885',
                         field: 'max_level',
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
                         description: 'Quest Controlled',
                         field: 'reward_method',
                         itemIcon: '3366',
                         fieldType: 'checkbox',
                         true: 2,
                         col: 'col-3',
                       },
                       {
                         description: 'Reward Text',
                         field: 'reward_text',
                         fieldType: 'text',
                         itemIcon: '3366',
                         col: 'col-5',
                       },
                       {
                         description: 'Reward Item ID(s)',
                         field: 'reward_id_list',
                         fieldType: 'reward_id_list',
                         itemIcon: '3366',
                         col: 'col-4',
                         onclick: setSelectorActive,
                       },
                       {
                         description: 'EXP Reward',
                         field: 'exp_reward',
                         itemIcon: '2045',
                         fieldType: 'text',
                         col: 'col-4'
                       },
                       {
                         description: 'Cash Reward',
                         field: 'cash_reward',
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
                         description: 'Reward Point Type',
                         field: 'reward_point_type',
                         itemIcon: '1955',
                         fieldType: 'text',
                         onclick: setSelectorActive,
                         col: 'col-6',
                       },
                       {
                         description: 'Reward Points',
                         field: 'reward_points',
                         itemIcon: '1955',
                         fieldType: 'text',
                         col: 'col-6',
                       },
                       {
                         description: 'Replay Timer Group',
                         field: 'replay_timer_group',
                         onclick: setSelectorActive,
                         fieldType: 'text',
                         itemIcon: '750',
                         col: 'col-6',
                       },
                       {
                         description: 'Replay Timer (sec)',
                         field: 'replay_timer_seconds',
                         fieldType: 'text',
                         itemIcon: '750',
                         col: 'col-6',
                       },
                       {
                         description: 'Request Timer Group',
                         field: 'request_timer_group',
                         onclick: setSelectorActive,
                         itemIcon: '750',
                         fieldType: 'text',
                         col: 'col-6',
                       },
                       {
                         description: 'Request Timer (sec)',
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

                      <!-- reward_id_list -->
                      <b-input-group v-if="field.fieldType === 'reward_id_list'">
                        <b-form-input
                          :id="field.field"
                          v-model="task[field.field]"
                          class="m-0 mt-1"
                          v-on="typeof field.onclick !== 'undefined' ? { click: () => field.onclick(field.field) } : {}"
                          v-b-tooltip.hover.v-dark.right :title="getFieldDescription(field.field)"
                          :style="(task[field.field] <= (typeof field.zeroValue !== 'undefined' ? field.zeroValue : 0) ? 'opacity: .5' : '')"
                        />
                        <b-button
                          @click="rewardIdMatchListItemSelector()"
                          size="sm"
                          v-b-tooltip.hover.v-dark.right :title="'Add entries to list'"
                          style="height: 29px; padding-right: 5px; margin-top: 3px;"
                          variant="btn btn btn-dark btn-outline-success btn-secondary btn-sm"
                        >
                          <i class="fa fa-pencil-square mr-1"></i>
                        </b-button>
                      </b-input-group>

                      <!-- textarea -->
                      <b-textarea
                        v-if="field.fieldType === 'textarea'"
                        :id="field.field"
                        v-model="task[field.field]"
                        class="m-0 mt-1"
                        style="max-height: 75px"
                        v-on="typeof field.onclick !== 'undefined' ? { click: () => field.onclick(field.field) } : {}"
                        v-b-tooltip.hover.v-dark.right :title="getFieldDescription(field.field)"
                        :style="(task[field.field] === '' ? 'opacity: .5' : '') + ';'"
                      ></b-textarea>

                      <!-- select -->
                      <select
                        :disabled="saveFreeze"
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
                      <b-button
                        @click="createActivity()"
                        :disabled="saveFreeze"
                        size="sm"
                        variant="outline-warning btn-dark"
                        class="ml-2"
                      >
                        <i class="fa fa-plus"></i>
                      </b-button>
                      <b-button
                        @click="cloneActivity()"
                        :disabled="saveFreeze"
                        size="sm"
                        variant="outline-white btn-dark"
                        class="ml-2"
                      >
                        <i class="ra ra-double-team"></i>
                      </b-button>
                      <b-button
                        @click="deleteActivity()"
                        :disabled="saveFreeze"
                        size="sm"
                        variant="outline-danger btn-dark"
                        class="ml-2"
                      >
                        <i class="fa fa-trash"></i>
                      </b-button>
                      <b-button
                        @click="moveActivityUp()"
                        :disabled="saveFreeze"
                        size="sm"
                        variant="outline-primary btn-dark"
                        class="ml-2"
                      >
                        <i class="fa fa-arrow-up"></i>
                      </b-button>
                      <b-button
                        @click="moveActivityDown()"
                        :disabled="saveFreeze"
                        size="sm"
                        variant="outline-primary btn-dark"
                        class="ml-2"
                      >
                        <i class="fa fa-arrow-down"></i>
                      </b-button>
                      <b-button
                        @click="forceReorderTaskActivities()"
                        :disabled="saveFreeze"
                        size="sm"
                        variant="outline-white btn-dark"
                        class="ml-2"
                        title="Fix and re-order task activities"
                      >
                        <i class="fa fa-wrench"></i>
                      </b-button>
                    </div>

                    <select
                      :disabled="saveFreeze"
                      size="2"
                      id="activities-scroll"
                      v-model="selectedActivity"
                      v-bind="task.task_activities"
                      @mouseover="scrollToActivity"
                      @change="selectActivity"
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
                              fieldType: 'header',
                              text: 'Activity',
                              col: 'col-12'
                           },
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
                             col: 'col-3',
                             onchange: activityTypeChange,
                           },
                           // {
                           //   description: 'Goal Method',
                           //   field: 'goalmethod',
                           //   fieldType: 'select',
                           //   itemIcon: '3196',
                           //   selectData: TASK_GOAL_METHOD_TYPE,
                           //   zeroValue: -1,
                           //   col: 'col-4',
                           // },
                           {
                             description: 'Goal Count',
                             itemIcon: '3196',
                             field: 'goalcount',
                             fieldType: 'number',
                             col: 'col-2',
                           },
                           {
                             description: 'Quest Controlled',
                             field: 'goalmethod',
                             itemIcon: '869',
                             fieldType: 'checkbox',
                             true: 2,
                             col: 'col-3',
                           },
                           {
                             description: 'Optional',
                             field: 'optional',
                             itemIcon: '6696',
                             fieldType: 'checkbox',
                             col: 'col-2',
                           },
                           {
                             description: 'Required Activity ',
                             itemIcon: '3196',
                             field: 'req_activity_id',
                             fieldType: 'number',
                             col: 'col-12',
                             fieldType: 'select',
                             selectData: buildTaskRequiredActivitySelection(),
                             zeroValue: -1,
                           },
                           {
                              fieldType: 'header',
                              text: 'Activity Description',
                              col: 'col-12',
                              info: 'Activity descriptions are built using either name & target fields or using the description override',
                           },
                           {
                             description: 'Item Name(s)',
                             itemIcon: '2275',
                             fieldType: 'text',
                             field: 'item_list',
                             col: 'col-6',
                             style: isDescriptionOverrideSet() ? 'opacity: .2;' : 'opacity: 1;',
                             showIf: isItemListVisible()
                           },
                           {
                             description: 'Activity Target',
                             itemIcon: '2275',
                             fieldType: 'text',
                             field: 'target_name',
                             col: 'col-6',
                             style: isDescriptionOverrideSet() ? 'opacity: .2;' : 'opacity: 1;',
                             showIf: isActivityTargetVisible()
                           },
                           {
                             description: 'Description Override',
                             field: 'description_override',
                             fieldType: 'text',
                             itemIcon: '2275',
                             col: 'col-12',
                           },
                           // {
                           //   description: 'Goal ID' + renderGoalIdDescriptor(),
                           //   field: 'goalid',
                           //   fieldType: 'number',
                           //   itemIcon: '3196',
                           //   col: 'col-3',
                           //   showIf: isGoalIdSelectorActive(),
                           //   onclick: isGoalIdSelectorActive() ? setSelectorActive : () => {},
                           // },
                           {
                              fieldType: 'header',
                              text: 'Activity Filters',
                              col: 'col-12',
                              info: 'Filters determine what the activity update applies to, zone, items, npc\'s etc.',
                           },
                           {
                             description: 'Zone(s)',
                             itemIcon: '6849',
                             field: 'zones',
                             fieldType: 'text',
                             type: 'text',
                             col: 'col-9',
                             onclick: setSelectorActive,
                           },
                           {
                             description: 'Zone Version',
                             itemIcon: '6849',
                             field: 'zone_version',
                             fieldType: 'number',
                             type: 'text',
                             col: 'col-3',
                           },
                           {
                             description: 'NPC Match List (' + npcMatchListSuffixDescription() + ')',
                             itemIcon: '6849',
                             field: 'npc_match_list',
                             fieldType: 'text',
                             showIf: isNpcMatchListSelectorActive(),
                             info: 'Use partial or full NPC names or exact IDs to match for this activity update. Example (10343|orc|gnoll)',
                             col: 'col-12',
                             onclick: setSelectorActive,
                           },
                           {
                             description: 'Item Match List (' + itemMatchListSuffixDescription() + ')',
                             itemIcon: '6849',
                             field: 'item_id_list',
                             fieldType: 'item_id_list',
                             showIf: isItemMatchListSelectorActive(),
                             info: 'Use exact item IDs to match for this activity update',
                             col: 'col-12',
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
                              fieldType: 'header',
                              text: 'Explore',
                              col: 'col-12',
                              info: 'Defines the exploration boundary',
                              showIf: isGoalIdExploreActive(),
                           },
                           {
                             description: 'Min X',
                             field: 'min_x',
                             fieldType: 'number',
                             itemIcon: '6851',
                             col: 'col-2',
                             showIf: isGoalIdExploreActive()
                           },
                           {
                             description: 'Max X',
                             field: 'max_x',
                             fieldType: 'number',
                             itemIcon: '6851',
                             col: 'col-2',
                             showIf: isGoalIdExploreActive()
                           },
                           {
                             description: 'Min Y',
                             field: 'min_y',
                             fieldType: 'number',
                             itemIcon: '6851',
                             col: 'col-2',
                             showIf: isGoalIdExploreActive()
                           },
                           {
                             description: 'Max Y',
                             field: 'max_y',
                             fieldType: 'number',
                             itemIcon: '6851',
                             col: 'col-2',
                             showIf: isGoalIdExploreActive()
                           },
                           {
                             description: 'Min Z',
                             field: 'min_z',
                             fieldType: 'number',
                             itemIcon: '6851',
                             col: 'col-2',
                             showIf: isGoalIdExploreActive()
                           },
                           {
                             description: 'Max Z',
                             field: 'max_z',
                             fieldType: 'number',
                             itemIcon: '6851',
                             col: 'col-2',
                             showIf: isGoalIdExploreActive()
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
                          :style="(typeof field.style !== 'undefined' && field.style) || typeof field.style === 'undefined'"
                        >

                          <!-- Header -->
                          <div
                            v-if="field.fieldType === 'header' && ((typeof field.showIf !== 'undefined' && field.showIf) || typeof field.showIf === 'undefined')"
                          >
                            <span
                              class="font-weight-bold"
                            >
                              {{ field.text }}
                              <i
                                v-b-tooltip.hover.v-dark.topright
                                :title="field.info"
                                v-if="field.info"
                                style="color: #6b614a"
                                class="fa fa-info-circle"
                              />
                            </span>
                          </div>

                          <!-- Description -->
                          <div v-if="field.description">
                            <span
                              v-if="field.itemIcon"
                              :class="'item-' + field.itemIcon + '-sm'"
                              style="display: inline-block"
                            />
                            {{ field.description }}
                            <i
                              v-b-tooltip.hover.v-dark.topright
                              :title="field.info"
                              v-if="field.info"
                              style="color: #6b614a"
                              class="fa fa-info-circle"
                            />

                            <!-- Show zone name -->
                            <div
                              v-if="field.field === 'zones' && task.task_activities[selectedActivity][field.field]"
                              class="font-weight-bold mt-1 text-center d-inline-block ml-1"
                            >
                              {{ getZoneNames() }}
                            </div>
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

                          <!-- input text -->
                          <b-input-group v-if="field.fieldType === 'item_id_list'">
                            <b-form-input
                              :id="field.field"
                              v-model="task.task_activities[selectedActivity][field.field]"
                              class="m-0 mt-1"
                              v-on="typeof field.onclick !== 'undefined' ? { click: () => field.onclick(field.field) } : {}"
                              v-b-tooltip.hover.v-dark.right :title="getFieldDescription(field.field)"
                              :style="(task.task_activities[selectedActivity][field.field] <= (typeof field.zeroValue !== 'undefined' ? field.zeroValue : 0) ? 'opacity: .5' : '')"
                            />
                            <b-button
                              @click="itemMatchListItemSelector()"
                              size="sm"
                              v-b-tooltip.hover.v-dark.right :title="'Add entries to list'"
                              style="height: 29px; padding-right: 5px; margin-top: 3px;"
                              variant="btn btn btn-dark btn-outline-success btn-secondary btn-sm"
                            >
                              <i class="fa fa-pencil-square mr-1"></i>
                            </b-button>
                          </b-input-group>

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
                    :disabled="saveFreeze"
                    size="sm"
                    variant="outline-warning btn-dark"
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

        <!-- npc_match_list preview -->
        <div
          style="width: auto;"
          class="fade-in"
          v-if="selectorActive['npc_match_list'] && task.task_activities[selectedActivity] && typeof task.task_activities[selectedActivity].npc_match_list !== 'undefined'"
        >
          <task-npc-match-list-previewer
            :activity="task.task_activities[selectedActivity]"
          />
        </div>

        <!-- reward_id_list preview -->
        <div
          style="width: auto;"
          class="fade-in"
          v-if="selectorActive['reward_id_list'] && task['reward_id_list']"
        >
          <task-item-match-list-previewer
            :id-list="task.reward_id_list"
            @remove-item="task.reward_id_list = removeItemFromMatchList(task.reward_id_list, $event)"
          />
        </div>

        <!-- item_id_list preview -->
        <div
          style="width: auto;"
          class="fade-in"
          v-if="selectorActive['item_id_list'] && task.task_activities[selectedActivity] && typeof task.task_activities[selectedActivity].item_id_list !== 'undefined'"
        >
          <task-item-match-list-previewer
            :id-list="task.task_activities[selectedActivity].item_id_list"
            @remove-item="task.task_activities[selectedActivity].item_id_list = removeItemFromMatchList(task.task_activities[selectedActivity].item_id_list, $event)"
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
          />
        </div>

        <!-- reward_id_list_select -->
        <div
          style="width: auto;"
          class="fade-in"
          v-if="selectorActive['reward_id_list_select']"
        >
          <task-item-selector
            @input="task['reward_id_list'] = addItemToMatchList(task['reward_id_list'], $event.id); setFieldModifiedById('reward_id_list')"
          />
        </div>

        <!-- item_id_list_select -->
        <div
          style="width: auto;"
          class="fade-in"
          v-if="selectorActive['item_id_list_select']"
        >
          <task-item-selector
            @input="task.task_activities[selectedActivity].item_id_list = addItemToMatchList(task.task_activities[selectedActivity].item_id_list, $event.id); setFieldModifiedById('reward_id_list');"
          />
        </div>

        <!-- (zones) Zone Selector -->
        <task-zone-selector
          :selected-zone-id="parseInt(task.task_activities[selectedActivity].zones)"
          v-if="task && task.task_activities && task.task_activities[selectedActivity] && selectorActive['zones']"
          @input="task.task_activities[selectedActivity].zones = $event.zoneId; task.task_activities[selectedActivity].npc_match_list = ''; setFieldModifiedById('zones')"
        />

        <!-- reward_point_type selector -->
        <alternate-currency-selector
          v-if="task && selectorActive['reward_point_type']"
          :selected-currency="task.reward_point_type"
          @input="task.reward_point_type = $event; setFieldModifiedById('reward_point_type');"
        />

        <!-- dz_template_id selector -->
        <dynamic-zone-template-selector
          v-if="task && selectorActive['dz_template_id']"
          :selected-id="task.dz_template_id"
          @input="task.dz_template_id = $event; setFieldModifiedById('dz_template_id');"
        />

        <!-- replay_timer_group selector -->
        <task-replay-request-group-selector
          v-if="task && selectorActive['replay_timer_group']"
          :selected-id="task.replay_timer_group"
          replay-field="replay_timer_group"
          @input="task.replay_timer_group = $event; setFieldModifiedById('replay_timer_group');"
        />

        <!-- request_timer_group selector -->
        <task-replay-request-group-selector
          v-if="task && selectorActive['request_timer_group']"
          :selected-id="task.request_timer_group"
          replay-field="request_timer_group"
          @input="task.request_timer_group = $event; setFieldModifiedById('request_timer_group');"
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
import {FreeIdFetcher} from "@/app/free-id-fetcher";
import {Npcs} from "@/app/npcs";
import TaskDescriptionSelector from "@/views/tasks/components/TaskDescriptionSelector.vue";
import TaskNpcMatchListPreviewer from "@/views/tasks/components/TaskNpcMatchListPreviewer.vue";
import {Zones} from "@/app/zones";
import ClipBoard from "@/app/clipboard/clipboard";
import TaskQuestExamplePreview from "@/views/tasks/components/TaskQuestExamplePreview.vue";
import InfoErrorBanner from "@/components/InfoErrorBanner.vue";
import AlternateCurrencySelector from "@/components/selectors/AlternateCurrencySelector.vue";
import DynamicZoneTemplateSelector from "@/components/selectors/DynamicZoneTemplateSelector.vue";
import TaskReplayRequestGroupSelector from "@/views/tasks/components/TaskReplayRequestGroupSelector.vue";
import TaskItemMatchListPreviewer from "@/views/tasks/components/TaskItemMatchListPreviewer.vue";

const MILLISECONDS_BEFORE_WINDOW_RESET = 10000;

export default {
  components: {
    TaskItemMatchListPreviewer,
    TaskReplayRequestGroupSelector,
    DynamicZoneTemplateSelector,
    AlternateCurrencySelector,
    InfoErrorBanner,
    TaskQuestExamplePreview,
    TaskNpcMatchListPreviewer,
    TaskDescriptionSelector,
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

      saveFreeze: false, // freezes inputs when actions are occurring

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

    async saveIfModified() {
      if (EditFormFieldUtil.anyFieldsHaveBeenEdited()) {
        await this.saveTask()
      }
    },

    async selectActivity() {
      await this.saveIfModified()
      this.updateQueryState()
    },

    removeItemFromMatchList(list, id) {
      let newList = list.split("|")
      const index = newList.indexOf(id.toString());
      if (index > -1) {
        newList.splice(index, 1);
      }

      return newList.join("|")
    },

    addItemToMatchList(list, itemId) {
      let newList = list.split("|")

      if (!newList.includes(itemId.toString())) {
        newList.push(itemId.toString())
      }

      return newList.filter(Number).join("|")
    },

    rewardIdMatchListItemSelector() {
      this.setSelectorActive('reward_id_list_select')
    },
    itemMatchListItemSelector() {
      this.setSelectorActive('item_id_list_select')
    },

    npcMatchListSuffixDescription() {
      const type = parseInt(this.task.task_activities[this.selectedActivity].activitytype)
      if (type === TASK_ACTIVITY_TYPE.KILL) {
        return 'To be killed'
      } else if (type === TASK_ACTIVITY_TYPE.LOOT) {
        return 'To be looted from (optional)'
      } else if (type === TASK_ACTIVITY_TYPE.SPEAK_WITH) {
        return 'To be spoken with'
      } else if (this.isDeliverToNPCActive()) {
        return 'To be delivered to'
      }

      return '???';
    },
    itemMatchListSuffixDescription() {
      const type = parseInt(this.task.task_activities[this.selectedActivity].activitytype)
      if (type === TASK_ACTIVITY_TYPE.DELIVER) {
        return 'Item(s) to be delivered'
      } else if (type === TASK_ACTIVITY_TYPE.TRADESKILL) {
        return 'Item(s) to be created'
      } else if (type === TASK_ACTIVITY_TYPE.FISH) {
        return 'Item(s) to be fished'
      } else if (type === TASK_ACTIVITY_TYPE.FORAGE) {
        return 'Item(s) to be foraged'
      } else if (type === TASK_ACTIVITY_TYPE.LOOT) {
        return 'Item(s) to be looted'
      }

      return '???';
    },

    buildTaskActivitySelection() {
      let activities = {
        "-1": "None"
      }
      if (this.task && this.task.task_activities) {
        for (const a of this.task.task_activities) {
          activities[a.activityid] = a.activityid + " " + Tasks.buildActivityDescription(a)
        }
      }

      return activities
    },

    buildTaskRequiredActivitySelection() {
      let activities = {
        "-1": "None"
      }
      if (this.task && this.task.task_activities) {
        for (const a of this.task.task_activities) {
          activities[a.activityid] = Tasks.buildActivityDescription(a)
        }
      }

      return activities
    },

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
          // this.task.task_activities[this.selectedActivity].goalid               = 0
          this.task.task_activities[this.selectedActivity].goalmethod           = 0
          this.task.task_activities[this.selectedActivity].goalcount            = 1
          this.task.task_activities[this.selectedActivity].optional             = 0
          this.task.task_activities[this.selectedActivity].item_id_list         = ""
          this.task.task_activities[this.selectedActivity].npc_match_list       = ""
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
      if (this.isNpcMatchListSelectorActive()) {
        return ' (NPC)'
      }

      return ""
    },

    renderGoalMatchListDescription() {
      if (this.isGoalIdItemSelectorActive()) {
        return ' (Item) (List of item ids)'
      }
      if (this.isNpcMatchListSelectorActive()) {
        return ' (NPC) (Can be NPC ID\'s or names)'
      }

      return ""
    },

    isGoalIdSelectorActive() {
      return this.task.task_activities && this.task.task_activities[this.selectedActivity] ? this.task.task_activities[this.selectedActivity].goalmethod === 0 : false
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

    isDescriptionOverrideSet() {
      return this.task.task_activities[this.selectedActivity].description_override.length > 0
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

    isNpcMatchListSelectorActive() {
      return [
        TASK_ACTIVITY_TYPE.KILL,
        TASK_ACTIVITY_TYPE.SPEAK_WITH,
        TASK_ACTIVITY_TYPE.DELIVER,
        TASK_ACTIVITY_TYPE.GIVE
      ].includes(
        parseInt(this.task.task_activities[this.selectedActivity].activitytype)
      )
    },

    isItemMatchListSelectorActive() {
      return [
        TASK_ACTIVITY_TYPE.LOOT,
        TASK_ACTIVITY_TYPE.TRADESKILL,
        TASK_ACTIVITY_TYPE.FISH,
        TASK_ACTIVITY_TYPE.FORAGE,
        TASK_ACTIVITY_TYPE.DELIVER
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

      if (isDescriptionOverrideEmpty && fieldId === 'goalid') {
        if (this.isGoalIdItemSelectorActive()) {
          this.task.task_activities[selectedActivity].item_list = event.name
          EditFormFieldUtil.setFieldModifiedById('item_list');
        }
        if (this.isNpcMatchListSelectorActive()) {
          this.task.task_activities[selectedActivity].target_name = Npcs.getCleanName(event.npc.name)
          EditFormFieldUtil.setFieldModifiedById('target_name');
        }
      }
    },

    async moveActivityUp() {
      this.saveFreeze = true

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

      this.saveFreeze = false
    },

    async cloneActivity() {
      this.saveFreeze = true

      await this.saveIfModified()
      try {
        const r = await Tasks.cloneTaskActivity(this.getBackendFormattedTask(), this.selectedActivity)
        if (r.status === 200) {
          this.task             = (await Tasks.getTask(this.$route.params.id))
          this.selectedActivity = r.data.activityid
          this.notification     = `Task activity (${this.selectedActivity}) successfully cloned`
        }
      } catch (err) {
        console.log(err)
        if (err.response && err.response.data && err.response.data.error) {
          this.error = err.response.data.error
        }
      }

      this.saveFreeze = false
    },

    async moveActivityDown() {
      this.saveFreeze = true

      await this.saveIfModified()
      const selectedActivity = this.selectedActivity
      let activities         = []
      if (this.task.task_activities && this.task.task_activities.length > 0) {
        activities = JSON.parse(JSON.stringify(this.task.task_activities))
      }

      let matchedIndex = -1
      for (const index in activities) {
        // console.log(index)
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

      this.saveFreeze = false
    },

    async createActivity() {
      this.saveFreeze = true

      try {
        await this.saveIfModified()
        const r = await Tasks.createNewTaskActivity(this.task)
        if (r.status === 200) {
          this.task             = (await Tasks.getTask(this.$route.params.id))
          this.selectedActivity = r.data.activityid
          this.notification     = "Task activity successfully created"
        }
      } catch (err) {
        console.log(err)
        if (err.response && err.response.data && err.response.data.error) {
          this.error = err.response.data.error
        }
      }

      this.saveFreeze = false
    },

    async deleteActivity() {
      if (!this.task.task_activities[this.selectedActivity]) {
        return
      }

      const activity = this.task.task_activities[this.selectedActivity]
      if (confirm(`Are you sure you want to delete this task activity?\n\n(Step ${activity.step}) Activity ${activity.activityid} \n\n` + this.buildActivityDescription(activity))) {
        this.saveFreeze = true

        try {
          await this.saveTask()

          let deletedSuccessfully = false
          if (this.task.task_activities) {

            for (const a of this.task.task_activities) {
              if (parseInt(a.activityid) === parseInt(this.selectedActivity)) {
                const r = await Tasks.deleteTaskActivity(a)
                if (r.status === 200) {
                  this.task             = (await Tasks.getTask(this.$route.params.id))
                  this.selectedActivity = a.activityid - 1
                  this.notification     = "Task activity successfully deleted"
                  deletedSuccessfully   = true
                }
              }
            }

            // // reorder task activities and then save again
            if (deletedSuccessfully) {
              this.reorderTaskActivities()
            }

          }
        } catch (err) {
          if (err.response && err.response.data && err.response.data.error) {
            this.error = err.response.data.error
          }
          console.log(err)
        }

        this.saveFreeze = false
      }
    },

    async forceReorderTaskActivities() {
      if (confirm(`Fix and reorder task activities? This can fix activities that have gaps or are not sequential and can cause the server to not load the activities properly.`)) {
        this.saveFreeze = true
        await this.reorderTaskActivities()
        this.saveFreeze = false
      }
    },

    async reorderTaskActivities() {
      let activityId = 0;
      for (const index in this.task.task_activities) {
        if (parseInt(this.task.task_activities[index].activityid) != activityId) {
          const previousId                            = parseInt(this.task.task_activities[index].activityid)
          this.task.task_activities[index].activityid = activityId
          await Tasks.updateTaskActivityId(this.task.task_activities[index], previousId)
        }

        activityId++;
      }
    },

    sendNotification(message) {
      this.notification = message
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
        const t = this.getBackendFormattedTask()
        const r = await Tasks.updateTask(t)

        if (t && t.task_activities) {
          for (let a of t.task_activities) {
            await Tasks.updateTaskActivityId(a, a.activityid)
          }
        }

        if (r.status === 200) {
          EditFormFieldUtil.resetFieldEditedStatus()
          this.notification = "Task updated!";
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
              this.notification = "New task created successfully!"
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
                  this.notification = "New task cloned successfully!"
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
        "reward_id_list",
        "item_id_list",
        "npc_match_list",
        "zones",
        "reward_point_type",
        "dz_template_id",
        "replay_timer_group",
        "request_timer_group",
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
