<template>
  <content-area>

    <create-user-modal
      @reload-users="listUsers()"
    />

    <div class="row justify-content-center">
      <div class="col-12 col-lg-10 col-xl-8 content-pop">
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
                v-for="user in users"
                :key="user.id"
              >
                <b-avatar
                  :src="user.avatar"
                  size="30"
                  variant="info"
                  v-b-tooltip.hover.v-dark.top
                  :title="user.user_name"
                  class="mr-3"
                />
                <span class="mr-auto">{{ user.user_name }}</span>
                <b-badge>{{ toTitleCase(user.provider) }}</b-badge>

                <a class="btn btn-sm btn-danger ml-3" @click="deleteUser(user)">
                  <i class="fa fa-trash"></i>
                  Delete
                </a>
              </b-list-group-item>
            </b-list-group>

          </div>

        </div>
      </div>
    </div>

  </content-area>
</template>

<script>
import ContentArea     from "@/components/layout/ContentArea.vue";
import {SpireApi}      from "@/app/api/spire-api";
import CreateUserModal from "@/views/user/CreateUserModal.vue";
import InfoErrorBanner from "@/components/InfoErrorBanner.vue";

export default {
  name: "UserManagement",
  components: { InfoErrorBanner, CreateUserModal, ContentArea },
  data() {
    return {
      // user list
      users: [],

      // notification / errors
      notification: "",
      error: "",
    }
  },
  methods: {

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
    this.listUsers()
  }
}
</script>

<style scoped>

</style>
