<template>
  <b-modal
    id="keypress-commands-modal"
    centered
    no-fade
    title="Keyboard Shortcuts"
    size="lg"
  >
    <div class="row mt-3" v-for="s in shortcuts">
      <div class="col-4 text-right font-weight-bold text-muted">
        {{s.keys}}
      </div>
      <div class="col-8">
        {{s.desc}}
      </div>
    </div>

    <div class="text-center mt-3 text-muted">
      Keyboard shortcuts must be entered while not focused in an input window
    </div>

    <template #modal-footer>
      <div class="">

      </div>
    </template>
  </b-modal>
</template>

<script>
export default {
  name: "KeypressCommandsModal",
  data() {
    return {
      shortcuts: [
        { keys: "Control+K", desc: "Navigation search window"},
        { keys: "H", desc: "Toggles Navigation menu hide/show"},
        { keys: "P", desc: "Start server modal (Admin)"},
        { keys: "S", desc: "Stop server modal (Admin)"},
        { keys: "R", desc: "Restart server modal (Admin)"},
        { keys: "C", desc: "Cancel server modal (Admin)"},
      ]
    }
  }
}
</script>

<style scoped>

</style>
