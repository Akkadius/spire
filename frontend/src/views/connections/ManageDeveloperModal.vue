<template>
  <b-modal
    id="manage-developer-modal"
    v-if="user && Object.keys(user).length > 0"
    centered
    :title="`Manage Developer [${user.user_name}] for connection [${connection.database_connection.name}]`"
    size="lg"
  >

    this is where developer managed

    <template #modal-footer>
      <div class="">

      </div>
    </template>
  </b-modal>
</template>

<script>
export default {
  name: "ManageDeveloperModal",
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
    init() {
      console.log("manage developer modal init")
    }
  }
}
</script>

<style scoped>

</style>
