<template>
  <eq-window-simple
    title="Quest Journal"
    class="eq-window-hybrid"
    style="margin-top: 30px"
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
        <td style="width: 200px">Task Title</td>
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
        <td :style="'color: ' + getTaskActivityStepColor(activity)">{{ buildActivityDescription(activity) }}</td>
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

    <div class="mt-3 eq-background-dark p-2" style="border: rgba(122, 134, 183, 0.5) 1px solid; height: 300px;">
      <v-runtime-template :template="getDescription()"/>

      <div
        class="mt-3"
        v-if="task.rewardid > 0 || task.reward_ebon_crystals > 0 || task.reward_radiant_crystals > 0"
      >
        Reward(s)

        <div v-if="task.rewardid > 0 && rewardItem">
          <item-popover
            :item="rewardItem"
            v-if="Object.keys(rewardItem).length > 0"
            size="regular"
            class="mt-3"
          />
        </div>
        <div v-if="task.reward_ebon_crystals > 0 && ebonCrystalItem">
          <item-popover
            :item="ebonCrystalItem"
            v-if="Object.keys(ebonCrystalItem).length > 0"
            size="regular"
            class="mt-1"
            :annotation="' (' + task.reward_ebon_crystals + ')'"
          />
        </div>
        <div v-if="task.reward_radiant_crystals > 0 && radiantCrystalItem" class="mt-1">
          <item-popover
            :item="radiantCrystalItem"
            v-if="Object.keys(radiantCrystalItem).length > 0"
            size="regular"
            :annotation="' (' + task.reward_radiant_crystals + ')'"
          />
        </div>
      </div>

    </div>

  </eq-window-simple>
</template>

<script>
import EqWindowSimple from "../../../components/eq-ui/EQWindowSimple";
import {TASK_TYPE}    from "@/app/constants/eq-task-constants";
import {Tasks}        from "@/app/tasks";
import {Zones}        from "@/app/zones";
import util           from "util";
import ItemPopover    from "@/components/ItemPopover";
import {Items}        from "@/app/items";

export default {
  name: "TaskPreview",
  components: {
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
      taskTimer: null,
      taskTimerDisplay: "Unlimited",
      taskDuration: "",

      zones: {},

      rewardItem: null,
      radiantCrystalItem: null,
      ebonCrystalItem: null,
    }
  },
  destroyed() {
    console.log("Destroyed")
    clearInterval(this.taskTimer)
  },

  watch: {
    task: {
      // This will let Vue know to look inside the array
      deep: true,

      // We have to move our method to a handler field
      handler() {
        console.log("task watch")
        clearInterval(this.taskTimer)
        this.load()
      }
    }
  },

  async mounted() {
    this.load()

    // preload zones
    this.zones = await Zones.getZones()
  },
  methods: {

    load() {
      this.setCountDownTimer()
      this.getDescription()

      this.rewardItem         = null
      this.radiantCrystalItem = null
      this.ebonCrystalItem    = null

      if (this.task.rewardid) {
        Items.getItem(this.task.rewardid).then((r) => {
          this.rewardItem = r
        })
        Items.getItem(40903).then((r) => {
          this.radiantCrystalItem = r
        })
        Items.getItem(40902).then((r) => {
          this.ebonCrystalItem = r
        })
      }
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
        let description = p
        description     = description.replaceAll(indexes.join(",") + ",", "")
        description     = description.replaceAll("]", "")
        description     = description.replaceAll("<BR>", "<br />")

        // load indexes with description part
        indexes.forEach((i) => {
          parts[i - 1] = description
        })
      })

      let finalDescription = globalDescription.replaceAll("<BR>", "<br />")
      if (this.selectedActivity >= 0 && parts[this.selectedActivity]) {
        finalDescription += parts[this.selectedActivity]
      }

      return '<div>' + finalDescription + '</div>'
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
      return activity.step
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

<style scoped>

</style>
