<template>
  <content-area>
    <div class="row">
      <div :class="(task ? 'col-7' : 'col-12') + ' p-0'">
        <eq-window-simple
          title="Task Editor"
          v-if="tasks"
          class="eq-window-hybrid"
          style="margin-top: 30px"
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
                      @click="createTask(true)"
                      size="sm"
                      variant="outline-light"
                    >
                      <i class="ra ra-double-team"></i>
                      Clone
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

              <b-alert show dismissable variant="warning" v-if="notification" class="mt-2">
                <div class="row">
                  <div class="col-11">
                    <i class="fa fa-info-circle mr-3"></i> {{ notification }}
                  </div>
                  <div class="col-1 text-right">
                    <i class="fa fa-remove"></i>
                  </div>
                </div>
              </b-alert>

              <b-alert show dismissable variant="danger" v-if="error" class="mt-2" @click="error = ''">
                <div class="row">
                  <div class="col-11">
                    <i class="fa fa-warning"></i> {{ error }}
                  </div>
                  <div class="col-1 text-right">
                    <i class="fa fa-remove"></i>
                  </div>
                </div>
              </b-alert>

              <eq-tabs
                id="task-edit-window"
                class="minified-inputs"
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
                       //   onclick: drawFreeIdSelector,
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
                         description: 'Reward Text',
                         field: 'reward',
                         fieldType: 'text',
                         itemIcon: '3366',
                         col: 'col-8',
                       },
                       {
                         description: 'Reward Item ID',
                         field: 'rewardid',
                         fieldType: 'text',
                         itemIcon: '3366',
                         col: 'col-4',
                         onclick: drawItemSelector,
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
                      :class="field.col + ' mb-3 pl-2 pr-2'"
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
                        v-b-tooltip.hover.v-dark.right :title="getFieldDescription(field.field)"
                        :style="(task[field.field] === 0 ? 'opacity: .5' : '')"
                      />

                      <!-- input text -->
                      <b-form-input
                        v-if="field.fieldType === 'text' || !field.fieldType"
                        :id="field.field"
                        v-model.number="task[field.field]"
                        class="m-0 mt-1"
                        v-on="typeof field.onclick !== 'undefined' ? { click: () => field.onclick() } : {}"
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
                      <b-button @click="createActivity()" size="sm" variant="outline-warning" class="ml-3">
                        <i class="fa fa-plus"></i></b-button>
                      <b-button @click="deleteActivity()" size="sm" variant="outline-danger" class="ml-3">
                        <i class="fa fa-trash"></i></b-button>
                    </div>


                    <select
                      size="2"
                      v-model="selectedActivity"
                      v-bind="task.task_activities"
                      @change="updateQueryState"
                      ignore-input-change="1"
                      class="form-control eq-input"
                      style="overflow-x: scroll; min-height: 20vh; overflow-y: scroll"
                    >
                      <option
                        v-for="activity in task.task_activities"
                        :value="activity.activityid"
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
                             fieldType: 'number',
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
                           },
                           {
                             description: 'Activity Target',
                             itemIcon: '5739',
                             field: 'target_name',
                             col: 'col-6',
                           },
                           {
                             description: 'Description Override',
                             field: 'description_override',
                             itemIcon: '2275',
                             col: 'col-12',
                           },
                           {
                             description: 'Goal ID',
                             field: 'goalid',
                             fieldType: 'number',
                             itemIcon: '3196',
                             col: 'col-2',
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
                             description: 'Activity Optional',
                             field: 'optional',
                             itemIcon: '4493',
                             fieldType: 'checkbox',
                             col: 'col-3',
                           },
                           {
                             description: 'Deliver to NPC',
                             field: 'delivertonpc',
                             fieldType: 'number',
                             itemIcon: '5742',
                             col: 'col-6',
                             zeroValue: 0,
                             onclick: drawNpcSelector,
                           },
                           {
                             description: 'Zone',
                             itemIcon: '3133',
                             field: 'zones',
                             col: 'col-6',
                             onclick: drawZoneSelector,
                           },
                           // Removed until fully implemented
                           // {
                           //   description: 'Item List',
                           //   field: 'item_list',
                           //   col: 'col-6',
                           // },
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
                        >
                          <div>
                            <span
                              v-if="field.itemIcon"
                              :class="'item-' + field.itemIcon + '-sm'"
                              style="display: inline-block"
                            />
                            {{ field.description }}
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
                            v-b-tooltip.hover.v-dark.right :title="getFieldDescription(field.field)"
                            :style="(task.task_activities[selectedActivity][field.field] === 0 ? 'opacity: .5' : '')"
                          />

                          <!-- input text -->
                          <b-form-input
                            v-if="field.fieldType === 'text' || !field.fieldType"
                            :id="field.field"
                            v-model="task.task_activities[selectedActivity][field.field]"
                            class="m-0 mt-1"
                            v-on="typeof field.onclick !== 'undefined' ? { click: () => field.onclick(selectedActivity) } : {}"
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
                    size="sm"
                    variant="outline-warning"
                  >
                    <i class="fa fa-save mr-1"></i>
                    Save
                  </b-button>

                  <b-button
                    @click="deleteTask()"

                    size="sm"
                    variant="outline-danger"
                  >
                    <i class="fa fa-trash"></i>
                    Delete
                  </b-button>
                </div>
              </div>
            </div>
          </div>

          <eq-debug :data="task"/>

        </eq-window-simple>
      </div>

      <div class="col-5 fade-in" v-if="task">

        <!-- item selector -->
        <div
          style="margin-top: 20px; width: auto;"
          class="fade-in"
          v-if="itemSelectorActive"
        >
          <task-item-selector
            @input="task['rewardid'] = $event.id; task['reward'] = $event.name; setFieldModifiedById('rewardid'); setFieldModifiedById('reward')"
          />
        </div>

        <!-- Zone Selector -->
        <task-zone-selector
          :selected-zone-id="parseInt(task.task_activities[selectedActivity].zones)"
          v-if="task && task.task_activities && task.task_activities[selectedActivity] && zoneSelectorActive"
          @input="task.task_activities[selectedActivity].zones = $event.zoneId; setFieldModifiedById('zones')"
        />

        <!-- NPC Selector -->
        <task-npc-selector
          :selected-npc-id="task.task_activities[selectedActivity].delivertonpc"
          v-if="task && task.task_activities && task.task_activities[selectedActivity] && npcSelectorActive"
          @input="task.task_activities[selectedActivity].delivertonpc = $event.npcId; setFieldModifiedById('delivertonpc')"
        />

        <!-- free id selector -->
        <eq-window-simple
          title="Free Item Ids"
          style="margin-top: 30px; margin-right: 10px; width: auto;"
          class="fade-in"
          v-if="freeIdSelectorActive"
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
        </div>


      </div>
    </div>
  </content-area>
</template>

<script type="ts">
import EqWindow from "@/components/eq-ui/EQWindow.vue";
import ContentArea from "@/components/layout/ContentArea.vue";
import util from "util";
import {ROUTE} from "@/routes";
import {Tasks} from "@/app/tasks";
import EqCheckbox from "@/components/eq-ui/EQCheckbox.vue";
import {
  TASK_ACTIVITY_TYPES,
  TASK_DURATION_HUMAN,
  TASK_DURATION_TYPES,
  TASK_GOAL_METHOD_TYPE,
  TASK_TYPES
} from "@/app/constants/eq-task-constants";
import EqWindowSimple from "@/components/eq-ui/EQWindowSimple.vue";
import EqDebug from "@/components/eq-ui/EQDebug.vue";
import EqTabs from "@/components/eq-ui/EQTabs.vue";
import EqTab from "@/components/eq-ui/EQTab.vue";
import {EditFormFieldUtil} from "@/app/forms/edit-form-field-util";
import TaskPreview from "@/views/task-editor/components/TaskPreview.vue";
import TaskZoneSelector from "@/views/task-editor/components/TaskZoneSelector.vue";
import TaskItemSelector from "@/views/task-editor/components/TaskItemSelector.vue";
import FreeIdSelector from "@/components/tools/FreeIdSelector.vue";
import TaskNpcSelector from "@/views/task-editor/components/TaskNpcSelector.vue";
import {FreeIdFetcher} from "@/app/free-id-fetcher";

const MILLISECONDS_BEFORE_WINDOW_RESET = 5000;

export default {
  components: {
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
      tasks: [],
      filteredTasks: [],
      selectedTask: null,
      selectedActivity: null,
      taskSearchFilter: "",

      // preview / selectors
      previewTaskActive: true,
      zoneSelectorActive: false,
      npcSelectorActive: false,
      itemSelectorActive: false,
      freeIdSelectorActive: false,

      lastResetTime: Date.now(),

      notification: "",
      error: "",

      TASK_TYPES: TASK_TYPES,
      TASK_DURATION_TYPES: TASK_DURATION_TYPES,
      TASK_DURATION_HUMAN: TASK_DURATION_HUMAN,
      TASK_ACTIVITY_TYPES: TASK_ACTIVITY_TYPES,
      TASK_GOAL_METHOD_TYPE: TASK_GOAL_METHOD_TYPE,
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

    async createActivity() {
      try {
        const r = await Tasks.createNewTaskActivity(this.task)
        if (r.status === 200) {
          this.task             = (await Tasks.getTask(this.$route.params.id))
          this.selectedActivity = r.data.activityid
        }
      } catch (err) {
        console.log(err)
        if (err.response && err.response.data && err.response.data.error) {
          this.error = err.response.data.error
        }
      }
    },

    async deleteActivity() {
      try {
        if (this.task.task_activities) {
          for (const a of this.task.task_activities) {
            if (parseInt(a.activityid) === parseInt(this.selectedActivity)) {
              const r = await Tasks.deleteTaskActivity(a)
              if (r.status === 200) {
                this.task             = (await Tasks.getTask(this.$route.params.id))
                this.selectedActivity = a.activityid - 1
              }
            }
          }
        }
      } catch (err) {
        if (err.response && err.response.data && err.response.data.error) {
          this.error = err.response.data.error
        }
        console.log(err)
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
            this.updateQueryState()
            this.sendNotification("New task created successfully!")
          }
        } catch (err) {
          if (err.response && err.response.data && err.response.data.error) {
            this.error = err.response.data.error
          }
        }
      }
    },

    async cloneTask() {
      if (confirm(`Are you sure you want to copy this task? It will create a whole new copy.\n\n(${this.task.id}) ${this.task.title} `)) {
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
                this.updateQueryState()
                this.sendNotification("New task cloned successfully!")
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

        setTimeout(() => {
          const container = document.getElementById("task-list");
          const target    = document.getElementById(util.format("task-entry-%s", this.$route.params.id))
          if (container && target) {
            // container.scrollTop = target.offsetTop - 100;
            container.scrollTo({top: target.offsetTop - 150, behavior: "smooth"});
          }

          this.setFieldHighlights()

          // hooks
          setTimeout(() => {
            const target = document.getElementById("task-edit-window")
            if (target) {
              target.removeEventListener('input', EditFormFieldUtil.setFieldModified, true);
              target.addEventListener('input', EditFormFieldUtil.setFieldModified)
            }

            EditFormFieldUtil.resetFieldSubEditorHighlightedStatus()
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
        "reward",
        "rewardid",
        "delivertonpc",
        "zones",
        "item_list",
        "skill_list",
        "spell_list",
        "goalid"
      ];
      hasSubEditorFields.forEach((field) => {
        EditFormFieldUtil.setFieldHighlightHasSubEditor(field)
      })
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
      this.loadTasks()
      this.loadTask()
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
      this.previewTaskActive    = false;
      this.itemSelectorActive   = false;
      this.zoneSelectorActive   = false;
      this.npcSelectorActive    = false;
      this.freeIdSelectorActive = false;

      EditFormFieldUtil.resetFieldSubEditorHighlightedStatus()
    },
    drawZoneSelector(selectedActivity) {
      console.log("zone select", selectedActivity)
      this.resetPreviewComponents()
      this.lastResetTime      = Date.now()
      this.zoneSelectorActive = true
      EditFormFieldUtil.setFieldSubEditorHighlightedById("zones")
    },
    drawItemSelector() {
      console.log("item select")
      this.resetPreviewComponents()
      this.lastResetTime      = Date.now()
      this.itemSelectorActive = true
      EditFormFieldUtil.setFieldSubEditorHighlightedById("rewardid")
    },
    drawFreeIdSelector() {
      console.log("free id select")
      this.resetPreviewComponents()
      this.lastResetTime        = Date.now()
      this.freeIdSelectorActive = true
      EditFormFieldUtil.setFieldSubEditorHighlightedById("id")
    },
    drawNpcSelector() {
      console.log("npcSelectorActive id select")
      this.resetPreviewComponents()
      this.lastResetTime     = Date.now()
      this.npcSelectorActive = true
      EditFormFieldUtil.setFieldSubEditorHighlightedById("delivertonpc")
    },
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

.task-window-header td {
  border-right: rgba(122, 134, 183, 0.5) 1px solid;
  padding-left: 3px;
  padding-right: 3px;
  color: lightskyblue;
}

.task-window-header {
  width: 100%;
  margin: 0;
}

.task-window {
  border: rgba(122, 134, 183, 0.5) 1px solid;
}

.task-window td {
  padding-left: 3px;
}

</style>
