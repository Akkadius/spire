<template>
  <div>

    <!-- CONTENT -->
    <div class="container-fluid">
      <div class="panel-body">
        <div class="panel panel-default">
          <eq-window title="Task Editor" v-if="tasks" style="margin-top: 30px">
            <div class="row">
              <div :class="model ? 'col-4' : 'col-12' + ''">
                <h3 class="eq-header" style="text-align: center">Task Selection</h3>

                <!-- Task List -->
                <div style="" class="">

                  <input type="text"
                         placeholder="Filter results by name..."
                         v-model="taskSearchFilter"
                         @keyup="filterResultsByName"
                         class="form-control ml-3" style="width: 95%">

                  <ul
                    style="overflow-y: scroll; height: 70vh; overflow-y: scroll; overflow-x: hidden; white-space:nowrap"
                    class="eq p-1 m-3 eq-dark-background">

                    <li
                      @click="selectTask(task)"
                      :style="(parseInt(taskSelected) === parseInt(task.id) ? 'background-color: rgba(106, 76, 50, 0.5);' : '')"
                      v-for="task in filteredTasks">
                      <router-link
                        :to="'/tasks/' + task.id"
                        :style="'color: #FFFFFF !important; '">({{ task.id }}) {{ task.title }}

                      </router-link>
                    </li>
                  </ul>
                </div>
              </div>

              <!-- Task -->
              <div class="col-5" id="my-form" v-if="model">
                <h3 class="eq-header" style="text-align: center">Task</h3>
                <test-form :model="model" class="eq-input"/>
              </div>

              <!-- Task Activities -->
              <div class="col-3" v-if="model">
                <h3 class="eq-header" style="text-align: center">Task Activities</h3>

                <select
                  size="2"
                  v-model="activitySelected"
                  v-bind="model.task_activities"
                  class="form-control eq-input"
                  style="overflow-x: scroll; min-height: 20vh; overflow-y: scroll">
                  <option
                    v-for="activity in model.task_activities"
                    :value="activity.activityid">
                    [{{ activity.step }}-{{ activity.activityid }}]
                    {{ getActivityDescription(activity) }}
                  </option>
                </select>

                <div v-if="activitySelected !== null">
                  <pre class="eq">{{ model.task_activities[activitySelected] }}</pre>
                </div>

              </div>
            </div>

            <pre class="eq p-1 m-3">{{ model }}</pre>

          </eq-window>
        </div>

      </div>
    </div>

  </div>
</template>

<script type="ts">
import {TaskApi} from "@/app/api/api";
import EqWindow from "@/components/eq-ui/EQWindow.vue";
import {SpireApiClient} from "@/app/api/spire-api-client";

export default {
  components: {
    EqWindow,
    "test-form": () => import("@/components/forms/TasksForm"),
    "task-activity": () => import("@/components/forms/TaskActivitiesForm"),
    "page-header": () => import("@/views/layout/PageHeader")
  },

  data() {
    return {
      model: null,
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

      this.$router.push({path: '/tasks/' + this.taskSelected}).catch(() => {
      })
    },

    async filterResultsByName() {
      // console.log(this.taskSearchFilter)
      // console.log(this.tasks)

      let filteredTasks = [];
      this.tasks.forEach((task, index) => {
        if (task.title.toLowerCase().includes(this.taskSearchFilter.toLowerCase())) {
          filteredTasks.push(task);
        }
      });

      this.filteredTasks = filteredTasks;
    },

    async loadEntity() {
      if (typeof this.$route.params.id === "undefined") {
        this.model = null
        return
      }

      (new TaskApi(SpireApiClient.getOpenApiConfig())).getTask({
        id: this.$route.params.id,
        includes: "3"
      }).then((result) => {
        if (result.status === 200) {
          this.model = result.data
        }
      })

      this.taskSelected = this.$route.params.id
      this.activitySelected = null
    },
    async onChange() {
      if (typeof this.taskSelected === "undefined") {
        return
      }

      this.$router.push({path: '/tasks/' + this.taskSelected})
    },
    getActivityDescription(activity) {
      if (activity.description_override !== "") {
        return activity.description_override;
      }

      switch (activity.activitytype) {
        case 1:
          if (activity.item_list !== "") {
            return activity.item_list
          }

          // Deliver
          return "Deliver " + activity.goalcount + " to " + activity.target_name;
        case 2:
          // Kill
          return "Kill " + activity.goalcount + " " + activity.target_name;
        case 3:
          // Loot
          let string = "Loot " + activity.goalcount;

          if (activity.item_list !== "") {
            string += " " + activity.item_list
          }

          if (activity.target_name !== "") {
            string += " from " + activity.target_name
          }

          return string
        case 4:
          // SpeakWith
          return "Speak with " + activity.target_name;
        case 5:
          // Explore
          return "Explore " + activity.target_name;
        case 6:
          // TradeSkill
          return "Create " + activity.goalcount + " " + activity.target_name;
        case 7:
          // Fish
          return "Fish " + activity.goalcount;
        case 8:
          // Forage
          return "Forage " + activity.goalcount;
        case 9:
          // ActivityUse1
          return "Use " + activity.goalcount;
        case 10:
          // ActivityUse2
          return "Use " + activity.goalcount;
        case 11:
          // ActivityTouch
          return "Touch " + activity.target_name;
        case 100:
          // ActivityGiveCash
          return "Give " + activity.goalcount + " to " + activity.target_name;
        case 255:
          // Custom Task Activity Type
          return "None";
        default:
          // Custom Task Activity Type
          return "None";

      }
    },
    async init() {
      const result = await (new TaskApi(SpireApiClient.getOpenApiConfig())).listTasks()
      if (result.status === 200) {
        this.tasks = result.data
        this.filteredTasks = this.tasks;
      }

      this.loadEntity()
    }
  },
  async mounted() {
    this.init()
  },
  activated() {
    this.init()
  }
}

</script>
