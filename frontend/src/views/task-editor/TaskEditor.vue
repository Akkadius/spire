<template>
  <content-area>
    <div class="row">
      <div class="col-7">
        <eq-window-simple
          title="Task Editor"
          v-if="tasks"
          class="eq-window-hybrid"
          style="margin-top: 30px"
        >
          <div class="row">
            <div :class="(task ? 'col-4' : 'col-12') + ' p-0'">
              <!-- Task List -->
              <div style="" class="">

                <input
                  type="text"
                  placeholder="Filter results by name..."
                  v-model="taskSearchFilter"
                  @keyup="filterResultsByName"
                  class="form-control"
                >

                <select
                  id="task-list"
                  size="2"
                  v-model="selectedTask"
                  @change="updateQueryState"
                  class="form-control eq-input eq p-1 mt-3 eq-dark-background"
                  style="overflow-x: scroll; height: 80vh; overflow-y: scroll"
                >
                  <option
                    v-for="task in filteredTasks"
                    :id="'task-entry-' + task.id"
                    :value="task.id"
                  >
                    ({{ task.id }}) {{ task.title }}
                  </option>
                </select>

              </div>
            </div>

            <!-- Task -->
            <div class="col-8 fade-in" id="my-form" v-if="task">

              <eq-tabs
                id="task-edit-window"
              >
                <eq-tab
                  name="Task"
                  selected="true"
                >
                  <div class="row">
                    <div
                      v-for="field in
                 [
                   {
                     description: 'Task ID',
                     field: 'id',
                     col: 'col-2',
                   },
                   {
                     description: 'Title',
                     field: 'title',
                     col: 'col-6',
                   },
                   {
                     description: 'Task Type',
                     field: 'type',
                     fieldType: 'select',
                     col: 'col-4',
                     selectData: TASK_TYPES,
                     zeroValue: -1,
                   },
                   {
                     description: 'Task Description',
                     field: 'description',
                     fieldType: 'textarea',
                     col: 'col-12',
                   },
                   {
                     description: 'Duration Code',
                     field: 'duration_code',
                     fieldType: 'select',
                     selectData: TASK_DURATION_TYPES,
                     col: 'col-4',
                     zeroValue: -1,
                   },
                   {
                     description: 'Duration',
                     field: 'duration',
                     fieldType: 'text',
                     col: 'col-4',
                     zeroValue: -1,
                   },
                   {
                     description: 'Duration (Selector)',
                     field: 'duration',
                     fieldType: 'select',
                     selectData: TASK_DURATION_HUMAN,
                     col: 'col-4',
                     zeroValue: -1,
                   },
                   {
                     description: 'Min Level',
                     field: 'minlevel',
                     fieldType: 'text',
                     col: 'col-2',
                   },
                   {
                     description: 'Max Level',
                     field: 'maxlevel',
                     fieldType: 'text',
                     col: 'col-2',
                   },
                   {
                     description: 'Min Players',
                     field: 'min_players',
                     fieldType: 'text',
                     col: 'col-2',
                   },
                   {
                     description: 'Max Players',
                     field: 'max_players',
                     fieldType: 'text',
                     col: 'col-2',
                   },
                   {
                     description: 'Lvl Spread',
                     field: 'level_spread',
                     fieldType: 'text',
                     col: 'col-2',
                   },
                   {
                     description: 'Repeatable',
                     field: 'repeatable',
                     fieldType: 'checkbox',
                     col: 'col-2',
                   },
                   {
                     description: 'Completion Emote',
                     field: 'completion_emote',
                     fieldType: 'textarea',
                     col: 'col-12',
                   },
                   {
                     description: 'Reward',
                     field: 'reward',
                     fieldType: 'text',
                     col: 'col-4',
                   },
                   {
                     description: 'Reward ID',
                     field: 'rewardid',
                     fieldType: 'text',
                     col: 'col-4',
                   },
                   {
                     description: 'EXP Reward',
                     field: 'xpreward',
                     fieldType: 'text',
                     col: 'col-4'
                   },
                   {
                     description: 'Cash Reward',
                     field: 'cashreward',
                     fieldType: 'text',
                     col: 'col-4',
                   },
                   {
                     description: 'Faction Reward',
                     field: 'faction_reward',
                     fieldType: 'text',
                     col: 'col-4',
                   },
                   {
                     description: 'Reward Ebon Crystals',
                     field: 'reward_ebon_crystals',
                     fieldType: 'text',
                     col: 'col-4',
                   },
                   {
                     description: 'Reward Radiant Crystals',
                     field: 'reward_radiant_crystals',
                     fieldType: 'text',
                     col: 'col-4',
                   },
                   {
                     description: 'Replay Timer Seconds',
                     field: 'replay_timer_seconds',
                     fieldType: 'text',
                     col: 'col-4',
                   },
                   {
                     description: 'Request Timer Seconds',
                     field: 'request_timer_seconds',
                     fieldType: 'text',
                     col: 'col-4',
                   },

                 ]"
                      :class="field.col + ' mb-3'"
                    >

                      <div>
                        {{ field.description }}
                      </div>

                      <!-- checkbox -->
                      <eq-checkbox
                        v-b-tooltip.hover.v-dark.right :title="getFieldDescription(field.field)"
                        class="mb-1 mt-3 d-inline-block"
                        :true-value="(typeof field.true !== 'undefined' ? field.true : 1)"
                        :false-value="(typeof field.false !== 'undefined' ? field.false : 0)"
                        v-model.number="task[field.field]"
                        @input="task[field.field] = $event"
                        v-if="field.fieldType === 'checkbox'"
                      />

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
                        max-rows="12"
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

                    <select
                      size="2"
                      v-model="selectedActivity"
                      v-bind="task.task_activities"
                      @change="updateQueryState"
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

                    <div v-if="selectedActivity !== null && task && task.task_activities[selectedActivity]">
                      <div class="row mt-3">
                        <div
                          v-for="field in
                     [
                       {
                         description: 'Activity ID',
                         field: 'activityid',
                         col: 'col-6',
                       },
                       {
                         description: 'Task Step',
                         field: 'step',
                         col: 'col-6',
                       },
                       {
                         description: 'Activity Type',
                         field: 'activitytype',
                         fieldType: 'select',
                         selectData: TASK_ACTIVITY_TYPES,
                         col: 'col-6',
                       },
                       {
                         description: 'Activity Target',
                         field: 'target_name',
                         col: 'col-6',
                       },
                       {
                         description: 'Description Override',
                         field: 'description_override',
                         col: 'col-12',
                       },
                       {
                         description: 'Goal ID',
                         field: 'goalid',
                         col: 'col-2',
                       },
                       {
                         description: 'Goal Method',
                         field: 'goalmethod',
                         fieldType: 'select',
                         selectData: TASK_GOAL_METHOD_TYPE,
                         zeroValue: -1,
                         col: 'col-4',
                       },
                       {
                         description: 'Goal Count',
                         field: 'goalcount',
                         col: 'col-3',
                       },
                       {
                         description: 'Activity Optional',
                         field: 'optional',
                         fieldType: 'checkbox',
                         col: 'col-3',
                       },
                       {
                         description: 'Deliver to NPC',
                         field: 'delivertonpc',
                         col: 'col-6',
                         zeroValue: 0,
                       },
                       {
                         description: 'Zones',
                         field: 'zones',
                         col: 'col-6',
                       },
                       {
                         description: 'Item List',
                         field: 'item_list',
                         col: 'col-6',
                       },
                       {
                         description: 'Skill List',
                         field: 'skill_list',
                         col: 'col-6',
                       },
                       {
                         description: 'Spell List',
                         field: 'spell_list',
                         col: 'col-12',
                       },
                     ]"
                          :class="field.col + ' mb-3'"
                        >
                          <div>
                            {{ field.description }}
                          </div>

                          <!-- checkbox -->
                          <eq-checkbox
                            v-b-tooltip.hover.v-dark.right :title="getFieldDescription(field.field)"
                            class="mb-1 mt-3 d-inline-block"
                            :true-value="(typeof field.true !== 'undefined' ? field.true : 1)"
                            :false-value="(typeof field.false !== 'undefined' ? field.false : 0)"
                            v-model.number="task.task_activities[selectedActivity][field.field]"
                            @input="task.task_activities[selectedActivity][field.field] = $event"
                            v-if="field.fieldType === 'checkbox'"
                          />

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
                            v-model.number="task.task_activities[selectedActivity][field.field]"
                            class="m-0 mt-1"
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


            </div>

          </div>

          <eq-debug :data="task"/>

        </eq-window-simple>
      </div>

      <div class="col-5">
        <task-preview
          :task="task"
          :selected-activity="selectedActivity"
          v-if="task"
        />
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

export default {
  components: {
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

      TASK_TYPES: TASK_TYPES,
      TASK_DURATION_TYPES: TASK_DURATION_TYPES,
      TASK_DURATION_HUMAN: TASK_DURATION_HUMAN,
      TASK_ACTIVITY_TYPES: TASK_ACTIVITY_TYPES,
      TASK_GOAL_METHOD_TYPE: TASK_GOAL_METHOD_TYPE,
    }
  },

  watch: {
    $route(to, from) {
      this.init()
    }
  },

  methods: {
    getFieldDescription(field) {
      return Tasks.getFieldDescription(field);
    },

    async selectTask(task) {
      this.selectedTask     = task.id
      this.selectedActivity = -1

      if (this.task && typeof this.task.task_activities[0] !== "undefined"
        && Object.keys(this.task.task_activities[0]).length > 0) {
        this.selectedActivity = 0
      }
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
        this.task = (await Tasks.getTask(this.$route.params.id))

        setTimeout(() => {
          const container = document.getElementById("task-list");
          const target    = document.getElementById(util.format("task-entry-%s", this.$route.params.id))
          if (container && target) {
            // container.scrollTop = target.offsetTop - 100;
            container.scrollTo({top: target.offsetTop - 150, behavior: "smooth"});
          }

          this.setFieldHighlights()

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
      const tasks = await Tasks.getTasks()
      if (tasks.length > 0) {
        this.tasks         = tasks
        this.filteredTasks = tasks
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

      // navigation
      this.$router.push(
        {
          path: util.format(ROUTE.TASK_EDIT, this.selectedTask),
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
