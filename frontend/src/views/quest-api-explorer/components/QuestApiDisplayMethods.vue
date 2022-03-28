<template>
  <div>
    <div
      v-for="(method, index) in apiMethods"
      :key="index"
      :style="(highlightedMethod === method.method ? 'background-color: rgba(106, 76, 50, 0.5);' : '') + '; overflow-y: scroll;'"
      :class="'method-scroll-' + method.method"
      @click="highlightMethod(method); loadExamples(method.method)"
    >

      <div class="d-inline-block">
        <button
          class='btn btn-sm btn-outline-warning mb-1 mr-2'
          @click="copyToClip(buildFullMethod(method))"
          style="font-size: 8px; padding: 0.125rem 0.4rem; opacity: .6"
        >
          <i class="fa fa-clipboard"></i>
        </button>
      </div>

      <!-- Method Prefix: eg $client quest:: eq. etc. -->
      <div v-if="method.methodPrefix" style="color:#9CDCFE;" class="d-inline-block">
        {{ method.methodPrefix }}
      </div>

      <!-- Method signature: Display if no linked example -->
      <div v-if="!linkedExamples[languageSelection][method.method + '(']" class="d-inline-block">
        {{ method.method }}
      </div>

      <!-- Method signature: With linked examples -->
      <div class="d-inline-block">
        <a
          @click="loadExamples(method.method)" href="javascript:void(0);"
          class="d-inline-block"
          v-if="linkedExamples[languageSelection][method.method + '(']"
        >
          {{ method.method }}
        </a>

        <!-- Show method signature eg: (uint8 exp_percentage, uint8 max_level = 0, bool ignore_mods = false) -->
        <div class="d-inline-block">({{ method.params.join(", ") }});</div>

        <!-- Code Comments -->
        <div v-if="method.comments" style="color: #57A64A" class="d-inline-block ml-3">
          {{ method.comments }}
        </div>

        <!-- Display how many examples in comments -->
        <div
          style="color: #57A64A"
          class="d-inline-block ml-3"
          v-if="linkedExamples[languageSelection][method.method + '('] && linkedExamples[languageSelection][method.method + '('].length > 0"
        >
          {{ (languageSelection === "perl" ? "#" : "--") }}
          {{ linkedExamples[languageSelection][method.method + "("].length }} Example(s)
        </div>

        <!-- Display categories -->
        <div
          style="color: #57A64A"
          class="d-inline-block ml-3"
          v-if="method.categories && method.categories.length > 0"
        >
          {{ (languageSelection === "perl" ? "#" : "--") }}
          ({{ method.categories.join(", ") }})
        </div>

      </div>
    </div>
  </div>
</template>

<script>
import util      from "util";
import ClipBoard from "@/app/clipboard/clipboard";
import Analytics from "@/app/analytics/analytics";
import {ROUTE}   from "@/routes";

export default {
  name: "QuestApiDisplayMethods",
  data() {
    return {
      highlightedMethod: ""
    }
  },
  props: {
    apiMethods: { type: Array },
    linkedExamples: { type: Object },
    languageSelection: { type: String }
  },
  created() {
    if (this.$route.query.h) {
      this.highlightedMethod = this.$route.query.h

      setTimeout(this.scrollHighlightedIntoView, 100)
    }
  },
  methods: {
    loadExamples: function (method) {
      this.$emit("load-quest-example",
        {
          language: this.languageSelection,
          search: method + "("
        }
      );
    },
    scrollHighlightedIntoView() {
      let highlighted = document.getElementsByClassName("method-scroll-" + this.$route.query.h);
      if (highlighted && highlighted[0]) {
        window.scrollTo({ top: highlighted[0].offsetTop + 70, behavior: "smooth" });
      }
    },
    copyToClip(s) {
      ClipBoard.copyFromText(s)

      Analytics.trackCountsEvent("clipboard_copy_method", s)

      this.$bvToast.toast(s, {
        title: "Copied to Clipboard!",
        autoHideDelay: 2000,
        solid: true
      })
    },
    buildFullMethod(method) {
      return util.format(
        "%s%s(%s);",
        method.methodPrefix,
        method.method,
        method.params.join(", ")
      )
    },
    highlightMethod(method) {
      this.highlightedMethod = method.method;

      let queryState = {};
      Object.assign(queryState, this.$route.query)
      queryState.h = method.method

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
