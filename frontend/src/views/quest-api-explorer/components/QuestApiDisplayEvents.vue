<template>
  <div>
    <div v-for="event in events" :key="event">
      <div class="d-inline-block mt-2" style="vertical-align: top"
      >
        <button
          class='btn btn-sm btn-dark mb-1 mr-2'
          @click="copyToClip(event)"
          style="font-size: 8px; padding: 0.125rem 0.4rem; opacity: .6">
          <i class="fa fa-clipboard"></i>
        </button>
      </div>

      <pre
        :id="event"
        class="ml-0 mb-4 code-display"
        style="width: 100%; display: inline-block; padding-left: 20px !important; padding-top: 10px !important; padding-bottom: 10px !important"
      >
<span style="color:#ff7b72;">{{ getLangEventPrefix() }}</span> <span
        style="color: #d2a8ff">{{ getSelectedEvent(event).event_identifier }}</span>{{ getLangSubPostfix() }}
	<span style="color: #57A64A">{{ getLangCommentCharacters() }} {{event}}</span>
	<span style="color: #57A64A">{{ getLangCommentCharacters() }} Exported event variables</span>
<span v-for="(e, index) in getSelectedEvent(event).event_vars" :key="index"><span
  style="color:#9CDCFE;">	{{ getLangQuestPrefix() }}</span>debug("{{ e }} " {{ getLangConcatenate() }} <span
  style="color: rgb(252 199 33);">{{ getLangVariablePrefix() }}{{ e }}</span>);
</span><span style="color:#ff7b72">{{ getLangEventPostfix() }}</span></pre>
    </div>
  </div>
</template>

<script>
import util from "util";
import ClipBoard from "@/app/clipboard/clipboard";
import Analytics from "@/app/analytics/analytics";
import {Notify} from "@/app/Notify";

export default {
  name: "QuestApiDisplayEvents",
  data() {
    return {}
  },
  props: {
    api: { type: Object },
    linkedExamples: { type: Object },
    languageSelection: { type: String },
    events: { type: Array }
  },
  methods: {
    loadExamples: function (search) {
      this.$emit("load-quest-example",
        {
          language: this.languageSelection,
          search: search + "("
        }
      );
    },
    copyToClip(e) {
      ClipBoard.copyFromElement(e)

      const event = e.split("-")[1]
      Analytics.trackCountsEvent("clipboard_copy_content", event)

      Notify.toast("Copied [" + event + "] to clipboard!");
    },
    getSelectedEvent(ev) {
      const entity = ev.split("-")[0]
      const event  = ev.split("-")[1]
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
    eventSelectionFormatName(ev) {
      const entity = ev.split("-")[0]
      const event  = ev.split("-")[1]
      return util.format("[%s] %s", entity, event)
    },
    getLangEventPrefix() {
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
    getLangEventPostfix() {
      return this.languageSelection === "perl" ? "}" : "end"
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
