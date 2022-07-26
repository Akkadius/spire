<template>
  <div class="row mt-3" v-if="localNotification || localError">
    <div class="col-12">

      <!-- Notification -->
      <b-alert show dismissable variant="warning" class="mb-0" v-if="localNotification">
        <div class="row" @click="dismissNotification()">
          <div class="col-11">
            <i class="fa fa-info-circle mr-3"></i> {{ localNotification }}
          </div>
          <div class="col-1 text-right">
            <i class="fa fa-remove"></i>
          </div>
        </div>
      </b-alert>

      <!-- Error -->
      <b-alert show dismissable variant="danger" v-if="localError" class="mb-0">
        <div class="row" @click="dismissError()">
          <div class="col-11">
            <i class="fa fa-warning mr-3"></i> {{ localError }}
          </div>
          <div class="col-1 text-right">
            <i class="fa fa-remove"></i>
          </div>
        </div>
      </b-alert>

    </div>
  </div>
</template>

<script>
export default {
  name: "InfoErrorBanner",
  props: {
    notification: {
      type: String,
      required: false
    },
    error: {
      type: String,
      required: false
    },
  },
  data() {
    return {
      localNotification: "",
      localError: "",
    }
  },
  watch: {
    notification: {
      handler(newVal) {
        console.log("[InfoErrorBanner] notification watcher [%s]", this.notification)
        this.sendNotification(this.notification, 5000)
      },
    },
    error: {
      handler(newVal) {
        console.log("[InfoErrorBanner] error watcher [%s]", this.error)

        this.localError = this.error
      },
    },
  },
  methods: {
    dismissError() {
      this.localError = ''
      this.$emit("dismiss-error", true);
    },
    dismissNotification() {
      this.localNotification = ''
      this.$emit("dismiss-notification", true);
    },

    sendNotification(message) {
      this.localNotification = message

      // dismiss in interval
      setTimeout(() => {
        this.dismissNotification()
      }, 5000)
    },
  }
}
</script>

<style scoped>

</style>
