<template>
  <div v-if="taskTimerDisplay">
    {{ taskTimerDisplay }} <span v-if="taskDuration !== ''">{{ taskDuration }}</span>
  </div>
</template>

<script>
import {debounce} from "@/app/utility/debounce";

export default {
  name: "TaskTimerCountdown",
  data() {
    return {
      taskTimer: null,

      // this keeps track of the last loaded state so we are not forcing
      // re-renders every time updates to the "task" object are made
      lastTaskDuration: -1,

      // local display
      taskTimerDisplay: "Unlimited",
      taskDuration: "",
    }
  },
  props: {
    task: Object,
  },
  watch: {
    task: {
      deep: true,
      handler() {
        this.debouncedLoad()
        console.log("task timer watch")
      }
    },
    selectedActivity: {
      handler() {
        this.description = this.getDescription()
      }
    },
  },
  destroyed() {
    clearInterval(this.taskTimer)
  },
  mounted() {
    this.load()

    this.debouncedLoad = debounce(() => {
      this.load();
    }, 500);
  },
  methods: {
    load() {
      console.log("task timer countdown")

      if (this.task.duration && this.task.duration !== this.lastTaskDuration) {
        this.setCountDownTimer()
        this.lastTaskDuration = this.task.duration
      }

      if (this.task.duration > 0) {
        let countDownDate = new Date();
        countDownDate.setSeconds(countDownDate.getSeconds() + this.task.duration)
        this.taskDuration     = "(" + this.getFormattedTime(countDownDate) + ")"
        this.taskTimerDisplay = this.getFormattedTime(countDownDate)

        this.calculateTime()

        if (this.taskTimer) {
          clearInterval(this.taskTimer);
        }

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
    },

    calculateTime() {
      let countDownDate = new Date();
      countDownDate.setSeconds(countDownDate.getSeconds() + this.task.duration)
      this.taskDuration     = "(" + this.getFormattedTime(countDownDate) + ")"
      this.taskTimerDisplay = this.getFormattedTime(countDownDate)
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
        this.calculateTime()
      }
    }
  }
}
</script>

<style scoped>

</style>
