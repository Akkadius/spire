<template>
  <div>
    <players-online-component :full-list="true"/>
  </div>
</template>

<script>
import PlayersOnlineComponent from "@/views/admin/components/PlayersOnlineComponent.vue";

export default {
  name: "PlayersOnline",
  components: { PlayersOnlineComponent }
}
</script>

<style scoped>

</style>
