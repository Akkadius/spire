<template>
  <b-modal
    v-if="connection && Object.keys(connection).length > 0"
    id="add-developer-server-connection-modal"
    centered
    :title="`Add Developer to Database Connection [${connection.database_connection.name}]`"
    size="lg"
  >
    <b-form-input
      v-model="search"
      id="user-search"
      v-on:keyup="searchUser()"
      placeholder="Search for user"
    />

    <div class="mt-3" v-if="usersToAdd.length === 0">
      No users found, you've added them all, or please refine search criteria
    </div>

    <div class="mt-3" v-if="usersToAdd.length > 0">
      Available users to add
    </div>

    <!-- Users to add -->
    <b-list-group class="mt-3">
      <b-list-group-item
        class="d-flex align-items-center"
        v-for="user in usersToAdd"
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

        <a class="btn btn-sm btn-white btn-danger ml-3" @click="addUser(user)">
          <i class="fa fa-plus"></i>
          Add User
        </a>
      </b-list-group-item>
    </b-list-group>

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
import {SpireApi}      from "../../app/api/spire-api";
import InfoErrorBanner from "@/components/InfoErrorBanner";

export default {
  name: "AddDeveloperModal",
  components: { InfoErrorBanner },
  props: {
    connection: {
      type: Object,
    },
  },
  data() {
    return {

      // search
      search: "",

      // user list
      usersToAdd: [],

      // notification / errors
      notification: "",
      error: "",
    }
  },
  watch: {
    connection: {
      handler() {
        this.init()
      },
      deep: true
    },
  },

  mounted() {
    this.init()
  },
  methods: {
    init() {
      this.usersToAdd = []
      this.search     = ""
      this.searchUser()
    },

    async addUser(user) {
      try {
        const r = await SpireApi.v1().post(`connection/${this.connection.id}/add-user/${user.id}`)
        if (r.status === 200) {
          this.$bvModal.hide('add-developer-server-connection-modal')
          this.$emit("reload-connections", true);

          // modal success
          this.$bvToast.toast(`User [${user.user_name}] added to [${this.connection.database_connection.name}]`, {
            title: `User added to connection`,
            autoHideDelay: 2000,
            solid: true
          })
        }
      } catch (err) {
        // error notify
        if (err.response && err.response.data && err.response.data.error) {
          this.error = err.response.data.error
        }
      }
    },

    searchUser() {
      if (Object.keys(this.connection).length === 0) {
        return;
      }

      const existingUsers = this.connection.database_connection.user_server_database_connections.filter((e) => {
        return e.user.id
      })

      SpireApi.v1().get('users', { params: { q: this.search } }).then((r) => {
        if (r.status === 200) {

          // only show users that aren't already on the connection
          this.usersToAdd = r.data.filter((u) => {
            return !existingUsers.some((e) => {
              // console.log("e", e)
              // console.log("u", u)
              // console.log(e.user_id, u.id)
              // console.log(e.user_id === u.id)
              return e.user_id === u.id
            })
          })
        }
      })
    },
    toTitleCase(str) {
      return str.replace(
        /\w\S*/g,
        function (txt) {
          return txt.charAt(0).toUpperCase() + txt.substr(1).toLowerCase();
        }
      );
    },
  }
}
</script>

<style scoped>

</style>
