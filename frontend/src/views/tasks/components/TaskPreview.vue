<template>
  <div>
    <eq-window-simple
      title="Quest Journal"
      v-if="task"
    >
      <div class="row">
        <div class="mb-3 col-9">
          Tasks
        </div>
        <div class="col-3">
          Request Timer: 00:00
        </div>
      </div>

      <table class="col-12 task-window">
        <thead class="task-window-header">
        <tr>
          <td style="width: 30px"></td>
          <td style="width: 300px">Task Title</td>
          <td v-if="taskTimerDisplay">Time Left</td>
        </tr>
        </thead>
        <tbody>
        <tr>
          <td :style="'text-align: center; color: ' + getTaskColor()">{{ getTaskTypeDisplayCode() }}</td>
          <td :style="'color: ' + getTaskColor()">{{ task.title }}</td>
          <td v-if="taskTimerDisplay">{{ taskTimerDisplay }} <span v-if="taskDuration !== ''">{{ taskDuration }}</span>
          </td>
        </tr>
        </tbody>
      </table>

      <div class="row">
        <div class="mt-3 col-9">
          Task Progression
        </div>
        <div class="col-3">
          <!--              <button class='eq-button' onclick="alert('click')">Remove</button>-->
        </div>
      </div>

      <table class="col-12 mt-3 task-window">
        <thead class="task-window-header">
        <tr>
          <td>Objective Instructions</td>
          <td style="width: 50px">Status</td>
          <td style="width: 200px">Zone</td>
        </tr>
        </thead>
        <tbody>
        <tr
          v-for="activity in task.task_activities"
        >
          <td
            :style="'color: ' + getTaskActivityStepColor(activity)"
            :title="buildActivityDescription(activity)"
          >
            {{ truncate(buildActivityDescription(activity), 50) }}
          </td>
          <td>{{ renderTaskActivityProgress(activity) }}</td>
          <td :style="'color: ' + getTaskActivityStepColor(activity)">
            {{ isActivityStepActive(activity) ? getZone(activity) : '???' }}
          </td>
        </tr>
        <tr class="this-is-a-space-row-dont-remove">
          <td>&nbsp;</td>
          <td></td>
          <td></td>
        </tr>
        </tbody>
      </table>

      <div
        class="mt-3 eq-background-dark p-2"
        v-if="description !== '' || hasReward()"
        style="border: rgba(122, 134, 183, 0.5) 1px solid;"
      >
        <v-runtime-template :template="'<div>' + description + '</div>'"/>

        <div
          :class="(description !== '' ? 'mt-3' : '')"
          v-if="hasReward()"
        >
          <span class="font-weight-bold">Reward(s)</span>

          <div v-if="task.reward_text" style="color: magenta" class="mt-1">
            {{ task.reward_text }}
          </div>

          <div
            v-for="item in rewardItems"
            :key="item.id"
          >
            <item-popover
              :item="item"
              v-if="Object.keys(item).length > 0 && item"
              size="regular"
              class="mt-1"
            />
          </div>
          <div v-if="task.alternate_currency && task.alternate_currency.item">
            <item-popover
              :item="task.alternate_currency.item"
              v-if="Object.keys(task.alternate_currency.item).length > 0"
              size="regular"
              class="mt-1"
              :annotation="' (' + task.reward_points.toString().replace(/\B(?=(\d{3})+(?!\d))/g, ',') + ')'"
            />
          </div>

          <div v-if="task.cash_reward > 0" class="mt-3">
            Cash
            <eq-cash-display
              class="d-inline-block"
              :price="task.cash_reward"
            />
          </div>

          <div v-if="task.exp_reward > 0" class="mt-3">
            Experience ({{ task.exp_reward.toString().replace(/\B(?=(\d{3})+(?!\d))/g, ",") }})
          </div>
        </div>

      </div>

    </eq-window-simple>


  </div>

</template>

<script>
import EqWindowSimple from "../../../components/eq-ui/EQWindowSimple";
import {TASK_TYPE}    from "@/app/constants/eq-task-constants";
import {Tasks}        from "@/app/tasks";
import {Zones}        from "@/app/zones";
import util           from "util";
import ItemPopover    from "@/components/ItemPopover";
import {Items}        from "@/app/items";
import EqCashDisplay  from "@/components/eq-ui/EqCashDisplay";
import EqTabs         from "@/components/eq-ui/EQTabs";
import EqTab          from "@/components/eq-ui/EQTab";

export default {
  name: "TaskPreview",
  components: {
    EqTab,
    EqTabs,
    EqCashDisplay,
    ItemPopover,
    EqWindowSimple,
    "v-runtime-template": () => import("v-runtime-template")
  },
  props: {
    task: Object,
    selectedActivity: {
      type: Number,
      required: false
    }
  },
  data() {
    return {
      // local display
      taskTimer: null,
      taskTimerDisplay: "Unlimited",
      taskDuration: "",
      description: "",
      zones: {},

      // item objects for rendering
      rewardItems: [],

      // this keeps track of the last loaded state so we are not forcing
      // re-renders every time updates to the "task" object are made
      lastTaskDuration: -1,
    }
  },
  destroyed() {
    clearInterval(this.taskTimer)
  },

  watch: {
    task: {
      deep: true,
      handler() {
        clearInterval(this.taskTimer)
        this.load()
      }
    },
    selectedActivity: {
      handler() {
        this.description = this.getDescription()
      }
    },
  },

  async mounted() {
    this.load()

    // preload zones
    this.zones = await Zones.getZones()
  },
  methods: {

    truncate(string, length) {
      if (string.length > length) {
        return string.substring(0, length) + '...';
      }

      return string
    },

    hasReward() {
      return ((this.rewardItems.length > 0)
        || this.task.reward_ebon_crystals > 0
        || this.task.reward_radiant_crystals > 0
        || this.task.exp_reward > 0
        || this.task.cash_reward > 0) && this.task.reward_method !== 2
    },

    created() {
      this.previousRewardItems = ""
    },

    async load() {
      if (this.task.duration && this.task.duration !== this.lastTaskDuration) {
        this.setCountDownTimer()
        this.lastTaskDuration = this.task.duration
      }

      // bust cache if the reward is different when we redraw this component
      // console.log(this.previousRewardItems)
      // console.log(this.task.reward_id_list)
      // console.log(this.previousRewardItems !== this.task.reward_id_list)
      if (this.previousRewardItems !== this.task.reward_id_list.toString()) {
        this.rewardItems = []
      }

      // load multiple rewards
      if (this.task.reward_id_list && this.task.reward_id_list.length > 0 && this.rewardItems.length === 0) {
        const items = this.task.reward_id_list.split("|").map((x) => {
          return parseInt(x, 10);
        });

        // we check to see if items is the same as previous reward items because
        // when changing fields we end up redrawing this area
        // we already loaded
        if (items.length > 0 && this.previousRewardItems !== items) {
          await Items.loadItemsBulk(items);
          let rewardItems = []
          for (let item of items) {
            rewardItems.push(await Items.getItem(item))
          }
          this.rewardItems         = rewardItems
          this.previousRewardItems = this.task.reward_id_list.toString()
        }
      }

      this.description = this.getDescription()
    },

    replaceDescriptionContent(s) {
      let d = s

      d = d.replaceAll("]", "")
      d = d.replaceAll("<BR>", "<br />")
      d = d.replaceAll("<br>", "<br />")
      d = d.replaceAll("<c \"", "<span style=\"color: ")
      d = d.replaceAll("</c>", "</span>")
      return d
    },

    getDescription() {
      const globalDescription = this.task.description.split("[")[0]
      const descriptionParts  = this.task.description.split("[")
      let parts               = {}

      descriptionParts.forEach((p) => {
        // 2,3,Your first clue is that the...
        let indexes = []
        p.split(",").forEach((s) => {
          if (parseInt(s)) {
            indexes.push(s)
          }
        })

        // get description in part
        let desc = p
        desc     = desc.replaceAll(indexes.join(",") + ",", "")
        desc     = this.replaceDescriptionContent(desc)

        // load indexes with description part
        indexes.forEach((i) => {
          parts[i - 1] = desc
        })
      })

      // strip extra page breaks at end of final description
      const pageBreak      = "<br />"
      let finalDescription = this.replaceDescriptionContent(globalDescription)
      if (this.selectedActivity >= 0 && parts[this.selectedActivity]) {
        finalDescription += parts[this.selectedActivity]
      }

      let n = 0
      while (finalDescription.slice(-6) === pageBreak) {
        finalDescription = finalDescription.slice(0, -6)

        if (n > 100) {
          break;
        }
        n++
      }

      return finalDescription.trim()
    },

    getZone(activity) {
      activity.zones = parseInt(activity.zones)
      if (activity.zones > 0) {
        for (let i in this.zones) {
          const zone = this.zones[i]
          if (zone.zoneidnumber === activity.zones) {
            return zone.long_name
          }
        }
      }

      if (activity.zones === 0) {
        return "ALL"
      }

      if (activity.zones === -1) {
        return "Unknown"
      }

      return ""
    },

    buildActivityDescription(activity) {
      if (!this.isActivityStepActive(activity)) {
        return '???'
      }

      return Tasks.buildActivityDescription(activity)
    },

    getTaskTypeDisplayCode() {
      if (parseInt(this.task.type) === TASK_TYPE.SHARED_TASK) {
        return 'S'
      }

      return 'Q'
    },

    getStepFromActivity(activity) {
      return activity ? activity.step : -1
    },

    isTaskActivityComplete(activity) {
      return (activity.activityid < parseInt(this.selectedActivity))
    },

    isActivityStepActive(activity) {
      const selectedActivity    = parseInt(this.selectedActivity)
      const selectedStep        = this.getStepFromActivity(this.task.task_activities[selectedActivity])
      const currentActivityStep = this.getStepFromActivity(activity)

      return selectedStep >= currentActivityStep;
    },

    renderTaskActivityProgress(activity) {
      if (!this.isActivityStepActive(activity)) {
        return ""
      }

      if (this.isTaskActivityComplete(activity)) {
        return "Done"
      }

      return util.format("%s/%s",
        0,
        activity.goalcount
      )
    },

    getTaskActivityStepColor(activity) {
      const selectedActivity = parseInt(this.selectedActivity)
      if (activity.activityid < selectedActivity) {
        return 'rgb(27 247 27)'
      }

      if (!this.isActivityStepActive(activity)) {
        return 'dimgray'
      }

      return 'white'
    },

    getTaskColor() {
      if (parseInt(this.task.type) === TASK_TYPE.SHARED_TASK) {
        return '#d27cd2'
      }

      return 'white'
    },
    getFormattedTime(countDownDate) {
      let now      = new Date().getTime();
      let distance = countDownDate - now;

      // Time calculations for days, hours, minutes and seconds
      let days    = Math.floor(distance / (1000 * 60 * 60 * 24));
      let hours   = Math.floor((distance % (1000 * 60 * 60 * 24)) / (1000 * 60 * 60));
      let minutes = Math.floor((distance % (1000 * 60 * 60)) / (1000 * 60));
      let seconds = Math.floor((distance % (1000 * 60)) / 1000);

      // Display the result in the element with id="demo"
      let timerDisplay = []
      if (days > 0) {
        timerDisplay.push(days)
      }
      if (hours >= 0) {
        timerDisplay.push(hours.toString().padStart(2, '0'))
      }
      if (minutes >= 0) {
        timerDisplay.push(minutes.toString().padStart(2, '0'))
      }
      if (seconds >= 0) {
        timerDisplay.push(seconds.toString().padStart(2, '0'))
      }

      return timerDisplay.join(":");
    },
    setCountDownTimer() {
      this.taskTimerDisplay = "Unlimited"
      this.taskDuration     = ""
      if (this.taskTimer) {
        clearInterval(this.taskTimer)
      }

      if (this.task.duration > 0) {
        let countDownDate = new Date();
        countDownDate.setSeconds(countDownDate.getSeconds() + this.task.duration)

        this.taskDuration     = "(" + this.getFormattedTime(countDownDate) + ")"
        this.taskTimerDisplay = this.getFormattedTime(countDownDate)

        this.taskTimer = setInterval(() => {
          let now               = new Date().getTime();
          let distance          = countDownDate - now;
          this.taskTimerDisplay = this.getFormattedTime(countDownDate)

          // If the count down is finished, write some text
          if (distance < 0) {
            clearInterval(this.taskTimer);
          }
        }, 1000);
      }
    }
  }
}
</script>

<style>
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
