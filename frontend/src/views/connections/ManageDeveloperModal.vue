<template>
  <b-modal
    id="manage-developer-modal"
    v-if="user && Object.keys(user).length > 0"
    centered
    :title="`Manage Developer [${user.user_name}] for connection [${connection.database_connection.name}]`"
    size="lg"
  >

    <template #modal-header>
      <div class="">
        {{ `Manage Developer [${user.user_name}] for connection [${connection.database_connection.name}]` }}

        <b-button
          variant="danger"
          class="btn-sm ml-1"
          style="padding: 0px 6px;"
          @click="deleteUserFromConn()"
        >
          <i class="fa fa-trash"></i>
          Remove User
        </b-button>

      </div>

      <a @click="closeModal">
        <i class="fa fa-times"></i>
      </a>
    </template>

    this is where developer managed

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
import {SpireApiClient} from "@/app/api/spire-api-client";
import InfoErrorBanner  from "@/components/InfoErrorBanner";

export default {
  name: "ManageDeveloperModal",
  components: { InfoErrorBanner },
  data() {
    return {

      // notification / errors
      notification: "",
      error: "",
    }
  },
  props: {
    user: {
      type: Object,
      required: true
    },
    connection: {
      type: Object,
      required: true
    },
  },
  watch: {
    user: {
      handler() {
        this.init()
      },
      deep: true
    },
    connection: {
      handler() {
        this.init()
      },
      deep: true
    },
  },
  methods: {
    async deleteUserFromConn() {
      if (confirm(`Are you sure you want to remove this user from this connection?`)) {
        try {
          const r = await SpireApiClient.v1().delete(`connection/${this.connection.id}/add-user/${this.user.id}`)
          if (r.status === 200) {

            // modal success
            this.$bvToast.toast(`User [${this.user.user_name}] removed from [${this.connection.database_connection.name}]`, {
              title: `User removed from connection`,
              autoHideDelay: 2000,
              solid: true
            })

            this.$emit("reload-connections", true);
            this.$bvModal.hide('manage-developer-modal')
          }
        } catch (err) {
          // error notify
          if (err.response && err.response.data && err.response.data.error) {
            this.error = err.response.data.error
          }
        }
      }
    },

    closeModal() {
      this.$bvModal.hide('manage-developer-modal')
    },

    init() {
      // console.log("manage developer modal init")
    }
  }
}
</script>

<style scoped>

</style>
