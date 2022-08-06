<template>
  <eq-window
    :title="'Replay Selection (' + replayField + ')'"
    v-if="replayField"
  >
    <div
      style="overflow-y: scroll; max-height: 95vh"
      id="replay-selection-viewport"
    >
      <div class="font-weight-bold text-center">
        Task replay ID(s) are sourced by the field itself. <br>
        ID's are not referenced by another table
      </div>

      <div class="mt-3">
        <button
          class='btn btn-sm btn-outline-warning mb-1 mr-2'
          @click="setUnusedId()"
        >
          <i class="fa fa-arrow-left"></i>
          Select Unused Group ID
        </button>
      </div>

      <table
        class="eq-table eq-highlight-rows row-table"
        style="display: table; font-size: 14px; overflow-x: scroll"
        v-if="replayGroupIds && replayGroupIds.length > 0"
      >
        <thead
          class="eq-table-floating-header"
        >
        <tr>
          <th style="width: 50px"></th>
          <th style="width: 40px">ID</th>
          <th>Task(s)</th>

        </tr>
        </thead>
        <tbody>
        <tr
          :id="'row-' + e"
          :class="(isRowSelected(e) ? 'pulsate-highlight-white' : '')"
          v-for="(e, index) in replayGroupIds"
          :key="e"
          style="height: 50px"
        >
          <td>
            <b-button
              class="btn-dark btn-sm btn-outline-warning"
              title="Select"
              @click="selectRow(e)"
            >
              <i class="fa fa-arrow-left"></i>
            </b-button>
          </td>

          <td>
            {{ e }}
          </td>

          <td>
            <div v-for="t in getTasksByReplayId(e)">
              ({{t.id}}) {{t.title}}
            </div>
          </td>

        </tr>
        </tbody>
      </table>
    </div>

  </eq-window>
</template>

<script>
import EqWindow            from "../../../components/eq-ui/EQWindow";
import {TaskApi}           from "@/app/api";
import {SpireApiClient}    from "@/app/api/spire-api-client";
import {SpireQueryBuilder} from "@/app/api/spire-query-builder";
import {scrollToTarget}    from "@/app/utility/scrollToTarget";

export default {
  name: "TaskReplayRequestGroupSelector",
  components: { EqWindow },
  data() {
    return {
      replayGroupIds: []
    }
  },
  props: {
    selectedId: {
      type: Number,
      required: false
    },
    replayField: {
      type: String,
      required: true
    },
  },
  watch: {
    replayField: {
      handler() {
        this.init()
      }
    },
  },
  created() {
    this.tasks = []
  },
  async mounted() {
    this.init()
  },
  methods: {
    async init() {
      const r = await (new TaskApi(SpireApiClient.getOpenApiConfig())).listTasks(
        (new SpireQueryBuilder())
          .orderBy([this.replayField])
          .get()
      )
      if (r.status === 200) {
        this.tasks         = r.data
        let replayGroupIds = []
        for (const t of r.data) {
          if (t[this.replayField] > 0) {
            if (!replayGroupIds.includes(t[this.replayField])) {
              replayGroupIds.push(t[this.replayField])
            }
          }
        }

        this.replayGroupIds = replayGroupIds
      }

      if (this.selectedId > 0) {
        scrollToTarget(
          'replay-selection-viewport',
          'row-' + this.selectedId
        )
      }
    },
    selectRow(entry) {
      this.$emit('input', entry);
    },
    isRowSelected(e) {
      return this.selectedId &&
        this.selectedId > 0 &&
        e === this.selectedId;
    },
    getTasksByReplayId(id) {
      let tasks = []
      for (const t of this.tasks) {
        if (t[this.replayField] > 0 && t[this.replayField] === id) {
          tasks.push(t)
        }
      }

      return tasks
    },
    setUnusedId() {
      let groupId = 0
      for (const t of this.tasks) {
        if (t[this.replayField] > 0 && t[this.replayField] > groupId) {
          groupId = t[this.replayField]
        }
      }

      this.$emit('input', (groupId + 1))
    }
  }
}
</script>

<style scoped>

</style>
