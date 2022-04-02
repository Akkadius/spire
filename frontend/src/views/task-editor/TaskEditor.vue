<template>
  <content-area>
    <div class="row">
      <div class="col-12">

        <eq-window title="Task Editor" v-if="tasks" style="margin-top: 30px">
          <div class="row">
            <div :class="task ? 'col-4' : 'col-12' + ''">
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
              <!--                <test-form :task="task" class="eq-input"/>-->
            </div>

            <!-- Task Activities -->
            <div class="col-3" v-if="task">
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

export default {
  components: {
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
    }
  },

  watch: {
    $route(to, from) {
      this.loadEntity()
    }
  },
  methods: {
    async selectTask(task) {
      this.taskSelected = task.id
      this.$router.push({path: util.format(ROUTE.TASK_EDIT, this.taskSelected)})
        .catch(() => {
        })
    },

    async filterResultsByName() {
      let filteredTasks = [];
      filteredTasks = this.tasks.filter((task) => {
        return task.title.toLowerCase().includes(this.taskSearchFilter.toLowerCase())
      })

      this.filteredTasks = filteredTasks;
    },

    async loadEntity() {
      this.task = null

      if (this.$route.params.id > 0) {
        this.task            = (await Tasks.getTask(this.$route.params.id))
        this.taskSelected     = this.$route.params.id
        this.activitySelected = null
      }
    },
    async onChange() {
      if (typeof this.taskSelected === "undefined") {
        return
      }

      this.$router.push({path: util.format(ROUTE.TASK_EDIT, this.taskSelected)})
        .catch(() => {})
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
