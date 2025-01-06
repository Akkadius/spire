<template>
  <content-area>

    <create-user-modal
      @reload-users="listUsers()"
    />

    <reset-user-password-modal
      :user="resetPasswordUser"
    />

    <div class="row justify-content-center">
      <div class="col-12 col-lg-10 col-xl-8 content-pop card">
        <div class="container-fluid">

          <div class="row justify-content-between align-items-center mt-5">
            <div class="col-12">
              <h2 class="mb-2">
                Spire User Management
              </h2>
              <p class="text-muted mb-md-0">
                Manage local database users
              </p>
            </div>
          </div>

          <hr>

          <div class="alert alert-warning" role="alert">
            <i class="fa fa-info-circle mr-1"></i> After user creation, you will need to set the users permissions in the <router-link to="/connections">connections</router-link> page. New users have no permissions by default.
          </div>

          <div class="mt-3">

            <info-error-banner
              :slim="true"
              :notification="notification"
              :error="error"
              @dismiss-error="error = ''"
              @dismiss-notification="notification = ''"
              class="mt-3"
            />

            <h3 class="mt-3">
              Local Users
              <a class="btn btn-sm btn-white btn-danger ml-3" @click="createUser(user)">
                <i class="fa fa-plus"></i>
                Create New User
              </a>
            </h3>

            <b-list-group class="mt-3">
              <b-list-group-item
                class="d-flex align-items-center"
                v-for="u in users"
                :key="u.id"
              >
                <b-avatar
                  :src="u.avatar"
                  size="30"
                  variant="info"
                  v-b-tooltip.hover.v-dark.top
                  :title="u.user_name"
                  class="mr-3"
                />
                <span class="mr-auto">{{ u.user_name }}</span>

                <a class="btn btn-sm btn-white mr-3" @click="resetPassword(u)">
                  <i class="fe fe-lock"></i>
                  Reset Password
                </a>

                <a class="btn btn-sm btn-danger mr-3" @click="deleteUser(u)" v-if="u.id !== user.id">
                  <i class="fa fa-trash"></i>
                  Delete
                </a>

                <b-badge>{{ toTitleCase(u.provider) }}</b-badge>


              </b-list-group-item>
            </b-list-group>

          </div>

        </div>
      </div>
    </div>

  </content-area>
</template>

<script>
import ContentArea            from "@/components/layout/ContentArea.vue";
import {SpireApi}             from "@/app/api/spire-api";
import CreateUserModal        from "@/views/user/CreateUserModal.vue";
import InfoErrorBanner        from "@/components/InfoErrorBanner.vue";
import UserContext            from "@/app/user/UserContext";
import ResetUserPasswordModal from "@/views/user/ResetUserPasswordModal.vue";

export default {
  name: "UserManagement",
  components: { ResetUserPasswordModal, InfoErrorBanner, CreateUserModal, ContentArea },
  data() {
    return {

      // current user
      user: {},

      resetPasswordUser: {},

      // user list
      users: [],

      // notification / errors
      notification: "",
      error: "",
    }
  },
  methods: {

    resetPassword(user) {
      this.resetPasswordUser = user
      this.$bvModal.show('reset-user-password-modal')
    },

    createUser() {
      this.$bvModal.show('create-user-modal')
    },

    async deleteUser(user) {
      if (confirm(`Are you sure you want to delete this user? \n\n[${user.user_name}] (${user.id})`)) {
        try {
          const r = await SpireApi.v1().delete('user/' + user.id)
          if (r.status === 200) {
            this.notification = "User deleted successfully"
            this.listUsers()
          }
        } catch (e) {
          // error notify
          if (e.response && e.response.data && e.response.data.error) {
            this.error = e.response.data.error
          }
        }
      }
    },

    toTitleCase(str) {
      return str.replace(
        /\w\S*/g,
        function (txt) {
          return txt.charAt(0).toUpperCase() + txt.substr(1).toLowerCase();
        }
      );
    },
    async listUsers() {
      const r = await SpireApi.v1().get('users')
      if (r.status === 200) {
        this.users = r.data.filter((e) => {
          return e.provider === "local"
        })
      }
    }
  },
  async mounted() {
    this.user = await UserContext.getUser()

    this.listUsers()
  }
}
</script>

<style scoped>

</style>
