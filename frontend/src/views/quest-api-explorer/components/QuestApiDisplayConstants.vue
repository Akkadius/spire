<template>
  <div
    class="ml-0 mb-4 code-display"
    style="width: 50vw; display: inline-block; padding-left: 10px !important; padding-top: 10px !important; padding-bottom: 10px !important"
    v-if="api"
  >
    <div
      v-for="(constant, index) in api[this.languageSelection].constants[constantSelection]"
      :style="(highlightedConstant === formatConstant(constant) ? 'background-color: rgba(106, 76, 50, 0.5);' : '')"
      :class="'method-scroll-' + formatConstant(constant)"
      @click="highlightConstant(formatConstant(constant)); loadConstantExamples(formatConstant(constant))"
      :key="index">

      <div class="d-inline-block">
        <button
          class='btn btn-sm btn-dark mb-1 mr-3'
          @click="copyToClip(formatConstant(constant))"
          style="font-size: 8px; padding: 0.125rem 0.4rem; opacity: .6">
          <i class="fa fa-clipboard"></i>
        </button>
      </div>

      <div class="d-inline-block">
        <!-- linked -->
        <a
          @click="loadConstantExamples(formatConstant(constant))" href="javascript:void(0);"
          v-if="linkedExamples[languageSelection][formatConstant(constant)]">
          {{ formatConstant(constant) }}
        </a>

        <!-- Non-linked -->
        <div v-if="!linkedExamples[languageSelection][formatConstant(constant)]"
        >
          {{ formatConstant(constant) }}
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import ClipBoard from "@/app/clipboard/clipboard";
import Analytics from "@/app/analytics/analytics";
import {ROUTE}   from "@/routes";
import {Notify}  from "@/app/Notify";

export default {
  name: "QuestApiDisplayConstants",
  props: {
    api: { type: Object },
    linkedExamples: { type: Object },
    languageSelection: { type: String },
    constantSelection: { type: String }
  },
  data() {
    return {
      highlightedConstant: ""
    }
  },
  created() {
    if (this.$route.query.h) {
      this.highlightedConstant = this.$route.query.h

      setTimeout(this.scrollHighlightedIntoView, 100)
      setTimeout(() => {
        this.loadConstantExamples(this.highlightedConstant)
      }, 500)

    }
  },
  methods: {
    scrollHighlightedIntoView() {
      let highlighted = document.getElementsByClassName("method-scroll-" + this.$route.query.h);
      if (highlighted && highlighted[0]) {
        window.scrollTo({ top: highlighted[0].offsetTop + 70, behavior: "smooth" });
      }
    },

    loadConstantExamples: function (search) {
      this.$emit("load-quest-example",
        {
          language: this.languageSelection,
          search: search
        }
      );
    },

    copyToClip(s) {
      ClipBoard.copyFromText(s)

      Analytics.trackCountsEvent("clipboard_copy_constant", s)

      Notify.toast("Copied to clipboard!");
    },
    formatConstant(constant) {
      return this.languageSelection === "lua" ? this.constantSelection + "." + constant.constant : constant.constant;
    },

    highlightConstant(constant) {
      this.highlightedConstant = constant;

      let queryState = {};
      Object.assign(queryState, this.$route.query)
      queryState.h = constant

      this.$router.push(
        {
          path: ROUTE.QUEST_API_EXPLORER,
          query: queryState
        }
      ).catch(() => {
      })

      // console.log(this.$route.query)
    }
  }
}
</script>

<style scoped>

</style>
