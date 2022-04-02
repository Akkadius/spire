<template>
  <content-area>
    <div class="row">
      <div class="col-12">

        <eq-window title="Task Editor" v-if="tasks" style="margin-top: 30px">
          <div class="row">
            <div :class="task ? 'col-3' : 'col-12' + ''">
              <h3 class="eq-header" style="text-align: center">Task Selection</h3>

              <!-- Task List -->
              <div style="" class="">

                <input
                  type="text"
                  placeholder="Filter results by name..."
                  v-model="taskSearchFilter"
                  @keyup="filterResultsByName"
                  class="form-control ml-3" style="width: 95%"
                >

                <ul
                  style="overflow-y: scroll; height: 70vh; overflow-x: hidden; white-space:nowrap"
                  class="eq p-1 m-3 eq-dark-background"
                >

                  <li
                    @click="selectTask(task)"
                    :style="(parseInt(taskSelected) === parseInt(task.id) ? 'background-color: rgba(106, 76, 50, 0.5);' : '')"
                    v-for="task in filteredTasks"
                  >
                    <router-link
                      :to="'/tasks/' + task.id"
                      :style="'color: #FFFFFF !important; '"
                    >({{ task.id }}) {{ task.title }}

                    </router-link>
                  </li>
                </ul>
              </div>
            </div>

            <!-- Task -->
            <div class="col-5" id="my-form" v-if="task">
              <h3 class="eq-header" style="text-align: center">Task</h3>

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
                     selectData: TASK_TYPES
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
                     description: 'Level Spread',
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
                     fieldType: 'text',
                     col: 'col-12',
                   },
                   {
                     description: 'Reward',
                     field: 'reward',
                     fieldType: 'text',
                     col: 'col-2',
                   },
                   {
                     description: 'Reward ID',
                     field: 'rewardid',
                     fieldType: 'text',
                     col: 'col-2',
                   },
                   {
                     description: 'EXP Reward',
                     field: 'xpreward',
                     fieldType: 'text',
                     col: 'col-2',
                   },
                   {
                     description: 'Cash Reward',
                     field: 'cashreward',
                     fieldType: 'text',
                     col: 'col-3',
                   },
                   {
                     description: 'Faction Reward',
                     field: 'faction_reward',
                     fieldType: 'text',
                     col: 'col-3',
                   },
                   {
                     description: 'Reward Ebon Crystals',
                     field: 'reward_ebon_crystals',
                     fieldType: 'text',
                     col: 'col-6',
                   },
                   {
                     description: 'Reward Radiant Crystals',
                     field: 'reward_radiant_crystals',
                     fieldType: 'text',
                     col: 'col-6',
                   },
                   {
                     description: 'Replay Timer Seconds',
                     field: 'replay_timer_seconds',
                     fieldType: 'text',
                     col: 'col-6',
                   },
                   {
                     description: 'Request Timer Seconds',
                     field: 'request_timer_seconds',
                     fieldType: 'text',
                     col: 'col-6',
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
                    :style="(task[field.field] === '' ? 'opacity: .5' : '')"
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

            </div>

            <!-- Task Activities -->
            <div class="col-4" v-if="task">
              <h3 class="eq-header" style="text-align: center">Task Activities</h3>

              <select
                size="2"
                v-model="activitySelected"
                v-bind="task.task_activities"
                class="form-control eq-input"
                style="overflow-x: scroll; min-height: 20vh; overflow-y: scroll"
              >
                <option
                  v-for="activity in task.task_activities"
                  :value="activity.activityid"
                >
                  [{{ activity.step }}-{{ activity.activityid }}]
                  {{ buildActivityDescription(activity) }}
                </option>
              </select>

              <div v-if="activitySelected !== null">
                <pre class="eq">{{ task.task_activities[activitySelected] }}</pre>
              </div>

            </div>
          </div>

          <pre class="eq p-1 m-3">{{ task }}</pre>

        </eq-window>
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
import {TASK_DURATION_HUMAN, TASK_DURATION_TYPES, TASK_TYPES} from "@/app/constants/eq-task-constants";

export default {
  components: {
    EqCheckbox,
    ContentArea,
    EqWindow,
  },

  data() {
    return {
      task: null,
      tasks: [],
      filteredTasks: [],
      taskSelected: null,
      activitySelected: null,
      taskSearchFilter: "",

      TASK_TYPES: TASK_TYPES,
      TASK_DURATION_TYPES: TASK_DURATION_TYPES,
      TASK_DURATION_HUMAN: TASK_DURATION_HUMAN,
    }
  },

  watch: {
    $route(to, from) {
      this.loadEntity()
    }
  },
  methods: {

    getFieldDescription(field) {
      return Tasks.getFieldDescription(field);
    },

    async selectTask(task) {
      this.taskSelected = task.id
      this.$router.push({path: util.format(ROUTE.TASK_EDIT, this.taskSelected)})
        .catch(() => {
        })
    },

    async filterResultsByName() {
      let filteredTasks = [];
      filteredTasks     = this.tasks.filter((task) => {
        return task.title.toLowerCase().includes(this.taskSearchFilter.toLowerCase())
      })

      this.filteredTasks = filteredTasks;
    },

    async loadEntity() {
      if (this.$route.params.id > 0) {
        this.task = (await Tasks.getTask(this.$route.params.id))
        if (Object.keys(this.task).length > 0) {
          this.taskSelected     = this.$route.params.id
          this.activitySelected = null
          return
        }
        // if no task found...
        this.task = null
      }
    },
    async onChange() {
      if (typeof this.taskSelected === "undefined") {
        return
      }

      this.$router.push({path: util.format(ROUTE.TASK_EDIT, this.taskSelected)})
        .catch(() => {
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
      await this.loadTasks()
      await this.loadEntity()
    }
  },
  async mounted() {
    await this.init()
  },
  activated() {
    this.init()
  }
}

</script>
