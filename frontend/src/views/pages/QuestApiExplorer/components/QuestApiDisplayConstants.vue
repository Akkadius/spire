<template>
  <div
    class="ml-0 mb-4 code-display"
    style="width: 50vw; display: inline-block; padding-left: 10px !important; padding-top: 10px !important; padding-bottom: 10px !important"
    v-if="api"
  >
    <div
      v-for="(constant, index) in api[getLanguageKey()].constants[constantSelection]"
      :key="index">

      <div class="d-inline-block">
        <button
          class='btn btn-sm btn-outline-warning mb-1 mr-3'
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
export default {
  name: "QuestApiDisplayConstants",
  props: {
    api: { type: Object },
    linkedExamples: { type: Object },
    languageSelection: { type: String },
    constantSelection: { type: String }
  },
  data() {
    return {}
  },
  methods: {
    loadConstantExamples: function (search) {
      console.log(search)
      this.$emit("load-quest-example",
        {
          language: this.languageSelection,
          search: search
        }
      );
    },
    copyToClip(s) {
      let tempNode   = document.createElement("input");
      tempNode.type  = "text";
      tempNode.value = s;

      document.body.appendChild(tempNode);

      tempNode.select();
      document.execCommand("Copy");

      document.body.removeChild(tempNode);

      this.$bvToast.toast(s, {
        title: "Copied to Clipboard!",
        autoHideDelay: 2000,
        solid: true
      })

    },
    getLanguageKey() {
      if (this.languageSelection === "perl") {
        return "perl_api"
      }
      if (this.languageSelection === "lua") {
        return "lua_api"
      }
      return ""
    },
    formatConstant(constant) {
      return this.languageSelection === "lua" ? this.constantSelection + "." + constant.constant : constant.constant;
    }
  }
}
</script>

<style scoped>

</style>
