<template>
  <div>
    <pre
      class="ml-0 mb-4 code-display"
      style="width: 50vw; display: inline-block; padding-left: 20px !important; padding-top: 10px !important; padding-bottom: 10px !important"
    >
<span style="color:#ff7b72;">{{ getEventPrefix() }}</span> <span style="color: #d2a8ff">{{ getSelectedEvent().event_identifier }}</span>{{ getLangSubPostfix() }}
	<span style="color: #57A64A">{{ getLangCommentCharacters() }} Exported event variables</span>
<span v-for="(e, index) in getSelectedEvent().event_vars" :key="index"><span style="color:#9CDCFE;">	{{ getLangQuestPrefix() }}</span>debug("{{ e }} " {{ getLangConcatenate() }} <span style="color: rgb(252 199 33);">{{ getLangVariablePrefix() }}{{ e }}</span>);
</span>}</pre>
  </div>
</template>

<script>
import util from "util";

export default {
  name: "QuestApiDisplayEvents",
  props: {
    api: { type: Object },
    linkedExamples: { type: Object },
    languageSelection: { type: String },
    eventSelection: { type: String }
  },
  methods: {
    loadExamples: function (search) {
      this.$emit("load-quest-example",
        {
          language: this.languageSelection,
          search: search
        }
      );
    },
    getSelectedEvent() {
      const entity = this.eventSelection.split("-")[0]
      const event  = this.eventSelection.split("-")[1]
      let events   = this.api[this.languageSelection].events
      let e        = []
      events.forEach(row => {
        if (row.event_identifier === "event_") {
          return
        }

        if (row.event_identifier === event && row.entity_type === entity) {
          e = row
        }

      })

      return e
    },
    eventSelectionFormatName() {
      const entity = this.eventSelection.split("-")[0]
      const event  = this.eventSelection.split("-")[1]
      return util.format("[%s] %s", entity, event)
    },
    getEventPrefix() {
      return this.languageSelection === "perl" ? "sub" : "function"
    },
    getLangQuestPrefix() {
      return this.languageSelection === "perl" ? "quest::" : "eq."
    },
    getLangVariablePrefix() {
      return this.languageSelection === "perl" ? "$" : "e."
    },
    getLangSubPostfix() {
      return this.languageSelection === "perl" ? " {" : "(e)"
    },
    getLangConcatenate() {
      return this.languageSelection === "perl" ? "." : ".."
    },
    getLangCommentCharacters() {
      return this.languageSelection === "perl" ? "#" : "--"
    }
  }
}
</script>

<style scoped>

</style>
