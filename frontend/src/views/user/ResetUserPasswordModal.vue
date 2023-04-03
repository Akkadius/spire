<template>
  <b-modal
    id="reset-user-password-modal"
    centered
    :title="`Reset Password For User [${user.user_name}] (${user.id})`"
    size="lg"
    @show="init()"
  >

    <div class="mt-3">
      Password
      <div class="input-group">
        <b-form-input
          v-model="password"
          :type="(showPassword ? 'text' : 'password')"
          id="password"
          name="password"
          autocomplete="off"
          placeholder="Password"
        />
        <div class="input-group-append">
          <button class="btn btn-dark" type="button" @click="showPassword = !showPassword">
            <i class="fe fe-lock"></i>
            Show Password
          </button>
        </div>
      </div>
    </div>

    <div class="mt-3">
      Confirm Password
      <div class="input-group">
        <b-form-input
          v-model="passwordConfirm"
          :type="(showPassword ? 'text' : 'password')"
          id="password-confirm"
          name="password"
          autocomplete="off"
          placeholder="Password"
        />
        <div class="input-group-append">
          <button class="btn btn-dark" type="button" @click="showPassword = !showPassword">
            <i class="fe fe-lock"></i>
            Show Password
          </button>
        </div>
      </div>
    </div>

    <b-button
      class="mt-5 btn-white form-control"
      @click="resetPassword()"
    ><i class="fe fe-lock"></i> Reset Password
    </b-button>

    <info-error-banner
      :slim="true"
      :notification="notification"
      :error="error"
      @dismiss-error="error = ''"
      @dismiss-notification="notification = ''"
      class="mt-3"
    />

    <template #modal-footer>
      <div class="">

      </div>
    </template>
  </b-modal>
</template>

<script>
import InfoErrorBanner from "@/components/InfoErrorBanner";
import {SpireApi}      from "@/app/api/spire-api";

export default {
  name: "ResetUserPasswordModal",
  components: { InfoErrorBanner },
  data() {
    return {

      // form
      password: "",
      passwordConfirm: "",
      showPassword: false,

      // notification / errors
      notification: "",
      error: "",
    }
  },
  props: {
    user: {
      type: Object,
    },
  },
  mounted() {
    this.init()
  },
  methods: {
    async resetPassword() {

      // validation
      if (this.password.length === 0) {
        this.error = "Password is empty!"
        return;
      }
      if (this.passwordConfirm.length === 0) {
        this.error = "Password confirmation is empty!"
        return;
      }
      if (this.password !== this.passwordConfirm) {
        this.error = "Passwords do not match!"
        return;
      }
      if (this.password.length < 8) {
        this.error = "Password must be greater than 8 characters!"
        return;
      }

      try {
        const r = await SpireApi.v1().post(`user/${this.user.id}/password-reset`, {
          password: this.password,
        })
        if (r.status === 200) {
          this.notification = "User password reset successfully!"

          this.$emit("reload-users", true);

          setTimeout(() => {
            this.$bvModal.hide('reset-user-password-modal')
          }, 2000)
        }
      } catch (e) {
        // error notify
        if (e.response && e.response.data && e.response.data.error) {
          this.error = e.response.data.error
        }
      }
    },
    init() {
      // reset
      this.password        = ""
      this.passwordConfirm = ""
      this.showPassword    = false

      this.notification = ""
      this.error        = ""
    },
  }
}
</script>

<style scoped>

</style>
