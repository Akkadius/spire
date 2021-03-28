<template>
  <div>
    <!-- Lua -->
    <pre
      v-if="languageSelection === 'lua'"
      class="ml-0 mb-4 code-display"
      style="width: 50vw; display: inline-block; padding-left: 20px !important; padding-top: 10px !important; padding-bottom: 10px !important"
    ><span style="color: #57A64A">-- {{ eventSelectionFormatName() }}</span>
function {{ getSelectedEvent().event_identifier }}(e)
	<span style="color: #57A64A">-- Exported event variables</span>
<span v-for="(e, index) in getSelectedEvent().event_vars" :key="index">	eq.debug("{{ e }} " .. e.{{ e }});
</span>end</pre>
    <!-- Perl -->
    <pre
      v-if="languageSelection === 'perl'"
      class="ml-0 mb-4 code-display"
      style="width: 50vw; display: inline-block; padding-left: 20px !important; padding-top: 10px !important; padding-bottom: 10px !important"
    ><span style="color: #57A64A"># {{ eventSelectionFormatName() }}</span>
sub {{ getSelectedEvent().event_identifier }} {
	<span style="color: #57A64A"># Exported event variables</span>
<span v-for="(e, index) in getSelectedEvent().event_vars" :key="index">	quest::debug("{{ e }} " . ${{ e }});
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
  }
}
</script>

<style scoped>

</style>
