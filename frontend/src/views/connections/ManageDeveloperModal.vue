<template>
  <b-modal
    id="manage-developer-modal"
    v-if="user && Object.keys(user).length > 0"
    centered
    :title="`Manage Developer [${user.user_name}] for connection [${connection.database_connection.name}]`"
    size="lg"
    @show="init()"
    @hidden="reset()"
  >

    <template #modal-header>
      <div class="">

        <b-avatar
          :src="user.avatar"
          variant="info"
          v-b-tooltip.hover.v-dark.top
          :title="user.user_name"
          size="30"
          class="mr-1"
        />

        {{ `Manage Developer [${user.user_name}] for connection [${connection.database_connection.name}]` }}

        <b-button
          variant="danger"
          class="btn-sm ml-3"
          style="padding: 0px 6px;"
          @click="deleteUserFromConn()"
          v-if="!isSelectedUserOwnerOfConnection() && isCurrentUserOwnerOfConnection()"
        >
          <i class="fa fa-trash"></i>
          Remove User
        </b-button>

      </div>

      <a @click="closeModal">
        <i class="fa fa-times"></i>
      </a>
    </template>

    <info-error-banner
      :slim="true"
      :notification="notification"
      :error="error"
      @dismiss-error="error = ''"
      @dismiss-notification="notification = ''"
      class="mb-3"
    />

    <div v-if="isSelectedUserOwnerOfConnection()">
      User is owner of connection and has no limitations
    </div>

    <div v-if="!isSelectedUserOwnerOfConnection()">

      <div v-if="!canViewPermissions()">
        Can not view user permissions
      </div>

      <div id="permissions" v-if="canViewPermissions() && permissionsLoaded">
        <!-- Header -->
        <div class="row mt-1 mb-3">
          <div class="col-5 text-right font-weight-bold">Resource</div>
          <div class="col-7 text-muted">
            Permission
          </div>
        </div>

        <!-- All -->
        <div class="row mt-1 mb-3">
          <div class="col-5 text-right">ALL</div>
          <div class="col-7">
            <b-form-checkbox
              switch
              v-for="option in options"
              :disabled="!isCurrentUserOwnerOfConnection()"
              :key="option.value"
              v-model="selectedAllToggle[option.value]"
              :aria-describedby="option.text"
              @change="toggleAll(option.value)"
              name="flavour-4a"
              inline
            >
              {{ option.text }}
            </b-form-checkbox>
          </div>
        </div>

        <div
          class="row mt-1" v-for="p in permissions"
          :key="p.name"
        >
          <div class="col-5 text-right">{{ p.name }}</div>
          <div class="col-7">
            <b-form-checkbox
              switch
              v-for="option in options"
              :key="option.value"
              :disabled="!isCurrentUserOwnerOfConnection()"
              v-model="selectedPermissions[p.identifier][option.value]"
              :aria-describedby="option.text"
              @change="showChanges(p.identifier)"
              name="flavour-4a"
              inline
            >
              {{ option.text }}
            </b-form-checkbox>
          </div>
        </div>
      </div>

    </div>

    <template #modal-footer>
      <div class="">

      </div>
    </template>
  </b-modal>
</template>

<script>
import {SpireApi}      from "../../app/api/spire-api";
import InfoErrorBanner from "@/components/InfoErrorBanner";
import UserContext     from "@/app/user/UserContext";

export default {
  name: "ManageDeveloperModal",
  components: { InfoErrorBanner },
  data() {
    return {

      userContext: {},

      permissionsLoaded: false,

      // permissions
      selectedPermissions: {},
      permissions: [],
      selectedAllToggle: {},
      options: [
        // { text: 'Read and Write', value: 'Read and Write' },
        { text: 'Read', value: 'read' },
        { text: 'Write', value: 'write' },
      ],

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
  methods: {
    toggleAll(type) {
      console.log("toggle all", type, this.selectedAllToggle[type])

      for (let p of this.permissions) {
        this.selectedPermissions[p.identifier][type] = this.selectedAllToggle[type]
      }

      this.submitPermissions()
    },

    showChanges(identifier) {
      this.submitPermissions()
    },

    async submitPermissions() {
      let payload = []

      // read/write all's
      let canReadAll  = true
      let canWriteAll = true
      for (let p of this.permissions) {
        if (!this.selectedPermissions[p.identifier].read) {
          canReadAll = false
        }
        if (!this.selectedPermissions[p.identifier].write) {
          canWriteAll = false
        }
      }

      if (canReadAll || canWriteAll) {
        let e        = {}
        e.permission = "ALL"
        e.read       = canReadAll
        e.write      = canWriteAll
        payload.push(e)
      }

      // individual perms
      for (let p of this.permissions) {
        let e = {}

        e.permission = p.identifier
        if (!canReadAll) {
          e.read = this.selectedPermissions[p.identifier].read
        }
        if (!canWriteAll) {
          e.write = this.selectedPermissions[p.identifier].write
        }

        if (e.read || e.write) {
          payload.push(e)
        }
      }
      try {
        const r = await SpireApi.v1().post(
          `connection-permissions/${this.connection.server_database_connection_id}/user/${this.user.id}`,
          payload
        )
      } catch (err) {
        // error notify
        if (err.response && err.response.data && err.response.data.error) {
          this.error = err.response.data.error
        }
      }

    },

    canViewPermissions() {
      return this.isCurrentUserOwnerOfConnection() || this.userContext.id === this.user.id
    },

    isCurrentUserOwnerOfConnection() {
      return this.userContext.id === this.connection.database_connection.created_by
    },

    isSelectedUserOwnerOfConnection() {
      return this.user.id === this.connection.database_connection.created_by
    },

    async deleteUserFromConn() {
      if (confirm(`Are you sure you want to remove this user from this connection?`)) {
        try {
          const r = await SpireApi.v1().delete(`connection/${this.connection.server_database_connection_id}/add-user/${this.user.id}`)
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

    reset() {
      this.notification        = ""
      this.error               = ""
      this.selectedPermissions = {};
      this.permissions         = []
      this.selectedAllToggle   = {}
    },

    async init() {
      // console.log("manage developer modal init")
      this.userContext = await (UserContext.getUser())

      // list permissions
      if (!this.isSelectedUserOwnerOfConnection()) {

        // get resources mapping list
        const r = await SpireApi.v1().get(`permissions/resources`)
        if (r.status === 200) {
          this.permissions = r.data

          // loop through permissions
          let permissions = {}
          for (let p of r.data) {
            if (typeof permissions[p.identifier] === "undefined") {
              permissions[p.identifier] = {}
            }
            permissions[p.identifier].read  = false
            permissions[p.identifier].write = false
          }

          // set permission selection
          this.selectedPermissions = permissions
        }

        try {
          this.permissionsLoaded = false;
          // user perms
          const userPerms = await SpireApi.v1().get(`connection-permissions/${this.connection.server_database_connection_id}/user/${this.user.id}`)
          if (userPerms.status === 200) {
            let permissions = {}

            for (let p of r.data) {
              if (typeof permissions[p.identifier] === "undefined") {
                permissions[p.identifier] = {}
              }
            }

            for (let p of userPerms.data) {
              if (typeof permissions[p.resource_name] === "undefined") {
                permissions[p.resource_name] = {}
              }
              // permissions[p.identifier].read  = false
              // permissions[p.identifier].write = false

              permissions[p.resource_name].read  = p.can_read === 1
              permissions[p.resource_name].write = p.can_write === 1
            }

            // for ALL top level permissions
            // set all permissions when set
            for (let p of this.permissions) {
              if (typeof permissions[p.identifier] === "undefined") {
                permissions[p.identifier] = {}
              }
            }

            for (let up of userPerms.data) {
              if (up.resource_name === 'ALL' && up.can_read) {
                this.selectedAllToggle['read'] = true
                for (let p of this.permissions) {
                  permissions[p.identifier].read = true
                }
              }
              if (up.resource_name === 'ALL' && up.can_write) {
                this.selectedAllToggle['write'] = true
                for (let p of this.permissions) {
                  permissions[p.identifier].write = true
                }
              }
            }

            this.selectedPermissions = permissions
            this.permissionsLoaded = true;
          }

        } catch (err) {
          console.log(err)

          // error notify
          if (err.response && err.response.data && err.response.data.error) {
            this.error = err.response.data.error
          }
        }

      }
    }
  }
}
</script>

<style scoped>

</style>
