<template>
  <eq-window-simple
    title="Quest Script Examples"
    class="eq-window-hybrid p-0 pl-3 pr-3"
    style="margin-top: 30px"
    v-if="task"
  >
    <eq-tabs>
      <eq-tab name="Perl" selected="true">

        <span class="font-weight-bold">Task Assignment</span>

        <div
          class="ml-0 code-display pl-0"
          style="width: 100%; display: inline-block; padding-top: 10px !important; padding-bottom: 10px !important; border-radius: 5px"
        >
          <div :title="'$client->AssignTask(int task_id, int npc_id);'" v-b-tooltip.hover.v-dark.left>
            <button
              class='btn btn-sm btn-outline-warning mb-1 mr-2'
              @click="copyToClip(`$client->AssignTask(${task.id}, 0);`)"
              style="font-size: 8px; padding: 0.125rem 0.4rem; opacity: .6"
            >
              <i class="fa fa-clipboard"></i>
            </button>
            <span style="color: rgb(156, 220, 254);">$client-></span>AssignTask({{ task.id }}, 0);
          </div>
          <div :title="'quest::taskselector(task1, task2, task3);'" v-b-tooltip.hover.v-dark.left>
            <button
              class='btn btn-sm btn-outline-warning mb-1 mr-2'
              @click="copyToClip(`quest::taskselector(${task.id});`)"
              style="font-size: 8px; padding: 0.125rem 0.4rem; opacity: .6"
            >
              <i class="fa fa-clipboard"></i>
            </button>
            <span style="color: rgb(156, 220, 254);">quest::</span>taskselector({{ task.id }});
          </div>
        </div>

        <div class="mt-3">
          To see more Task methods, events, examples, see
          <router-link to="/quest-api-explorer?lang=perl&q=task">here</router-link>
        </div>

      </eq-tab>
      <eq-tab name="Lua">

        <span class="font-weight-bold">Task Assignment</span>

        <div
          class="ml-0 code-display pl-0"
          style="width: 100%; display: inline-block; padding-top: 10px !important; padding-bottom: 10px !important; border-radius: 5px"
        >
          <div :title="'client:AssignTask(int task, int npc_id);'" v-b-tooltip.hover.v-dark.left>
            <button
              class='btn btn-sm btn-outline-warning mb-1 mr-2'
              @click="copyToClip(`client:AssignTask(${task.id}, 0);`)"
              style="font-size: 8px; padding: 0.125rem 0.4rem; opacity: .6"
            >
              <i class="fa fa-clipboard"></i>
            </button>
            <span style="color: rgb(156, 220, 254);">client:</span>AssignTask({{ task.id }}, 0);
          </div>
          <div :title="'eq.task_selector({task1, task2, task3, etc.});'" v-b-tooltip.hover.v-dark.left>
            <button
              class='btn btn-sm btn-outline-warning mb-1 mr-2'
              @click="copyToClip(`eq.task_selector({${task.id}});`)"
              style="font-size: 8px; padding: 0.125rem 0.4rem; opacity: .6"
            >
              <i class="fa fa-clipboard"></i>
            </button>
            <span style="color: rgb(156, 220, 254);">eq.</span>task_selector({ {{ task.id }} });
          </div>
        </div>

        <div class="mt-3">
          To see more Task methods, events, examples, see
          <router-link to="/quest-api-explorer?lang=lua&q=task">here</router-link>
        </div>

      </eq-tab>
    </eq-tabs>

  </eq-window-simple>
</template>

<script>
import EqWindowSimple from "../../../components/eq-ui/EQWindowSimple";
import EqTabs         from "../../../components/eq-ui/EQTabs";
import EqTab          from "../../../components/eq-ui/EQTab";
import ClipBoard      from "@/app/clipboard/clipboard";
export default {
  name: "TaskQuestExamplePreview",
  components: { EqTab, EqTabs, EqWindowSimple },
  props: {
    task: Object,
    selectedActivity: {
      type: Number,
      required: false
    }
  },
  methods: {
    copyToClip(s) {
      ClipBoard.copyFromText(s)

      this.$bvToast.toast(s, {
        title: "Copied to Clipboard!",
        autoHideDelay: 2000,
        solid: true
      })
    },
  }
}
</script>

<style scoped>

</style>
