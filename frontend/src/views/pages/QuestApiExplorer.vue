<template>
  <!-- CONTENT -->
  <div class="container-fluid">
    <div class="panel-body">
      <div class="panel panel-default">
        <div class="row">
          <div class="col-12">
            <eq-window style="margin-top: 30px">

              <h1 class="eq-header text-center mt-5" v-if="!loaded">LOADING... PLEASE WAIT...</h1>

              <app-loader :is-loading="!loaded" padding="4"/>

              <div v-if="loaded">

                <div class="row">
                  <div class="col-2 text-center">
                    Language
                    <b-form-select
                      v-model="languageSelection"
                      :options="languageOptions"
                      @change="formChange(); languageSelect()"/>
                  </div>
                  <div class="col-2 text-center">
                    Types
                    <b-form-select
                      v-model="methodTypeSelection"
                      @change="formChange(); methodTypeSelect()"
                      :options="methodTypeOptions"/>
                  </div>

                </div>
                <div class="row mt-2" v-if="apiMethods.length > 0">
                  <div class="col-12">
                    <pre
                      class="highlight html bg-dark hljs mb-4 code-display"
                      style="padding-left: 20px !important; padding-top: 10px !important; padding-bottom: 10px !important"
                      v-if="apiMethods.length > 0"><div
                      v-for="(method, index) in apiMethods"
                      :key="index"><span v-if="method.methodPrefix" style="color:#9CDCFE">{{
                        method.methodPrefix
                      }}</span><span v-if="!linkedExamples[method.method + '(']">{{ method.method }}</span><a
                      @click="loadExamples(method.method)" href="javascript:void(0);"
                      v-if="linkedExamples[method.method + '(']">{{ method.method }}</a>({{
                        method.params.join(", ")
                      }}); <span
                        v-if="method.comments" style="color: #57A64A">{{ method.comments }}</span></div></pre>
                  </div>

                  <div class="col-6 example-preview">

                    <eq-window title="Examples" v-if="displayExamples">

                      <div class="example-preview-inner">
                        <div v-for="example in displayExamples.slice(0,50)">
                          <span style="font-weight: bold">{{ example.file_name }}</span>
                          <pre
                            class="highlight html bg-dark hljs mb-4 code-display"
                            style="padding-left: 20px !important; padding-top: 10px !important; padding-bottom: 10px !important"
                          >{{ example.full_example.trim() }}</pre>
                        </div>
                      </div>


                    </eq-window>

                  </div>

                </div>
              </div>


            </eq-window>
          </div>
        </div>
      </div>
    </div>

    <!--        <debug-display-component :data="database"/>-->

  </div>

</template>

<script type="ts">
import EqWindow              from "@/components/eq-ui/EQWindow.vue";
import DebugDisplayComponent from "@/components/DebugDisplayComponent.vue";
import {SpireApiClient}      from "@/app/api/spire-api-client";
import axios                 from "axios";
import EqWindowSimple        from "@/components/eq-ui/EQWindowSimple.vue";
import EqTabs                from "@/components/eq-ui/EQTabs.vue";
import EqTab                 from "@/components/eq-ui/EQTab.vue";

export default {
  components: {
    EqTab,
    EqTabs,
    EqWindowSimple,
    DebugDisplayComponent,
    EqWindow,
    "page-header": () => import("@/views/layout/PageHeader")
  },
  data() {
    return {
      codeClass: "",

      // select
      languageSelection: null,
      languageOptions: [
        {value: null, text: 'Select a Language'},
        {value: 'perl', text: 'Perl'},
        {value: 'lua', text: 'Lua'},
      ],

      // select
      methodTypeSelection: null,
      methodTypeOptions: [],

      apiMethods: [],

      // actual code that gets highlighted
      methodDisplay: "",

      // route watcher
      routeWatcher: null,

      loaded: false,

      // linked method examples
      linkedExamples: {},
      // examples being displayed for current method
      displayExamples: []
    }
  },
  deactivated() {
    // remove route watcher
    this.routeWatcher()
  },
  activated() {
    this.init()
  },
  methods: {
    init() {
      this.routeWatcher = this.$watch('$route.query', () => {
        this.loadQueryParams()
      });

      this.loaded = false

      SpireApiClient.v1().get('/quest-api/methods', this.database).then((response) => {
        if (response.data && response.data.data) {
          this.methods = response.data.data

          this.loaded = true
          this.languageSelect()
        }

      }, (error) => {
        // this.errorMessage = "Unknown error trying to contact the database"
        if (error.response && error.response.data) {
          // this.errorMessage = error.response.data.error
        }
      }).catch((error) => {
        if (!axios.isCancel(error)) {
          // console.log(error)
          // this.errorMessage = "Unknown error trying to contact the database"
        }
      });

      this.loadQueryParams()
    },
    loadQueryParams() {
      this.languageSelection   = null
      this.methodTypeSelection = null

      if (this.$route.query.lang) {
        this.languageSelection = this.$route.query.lang
      }
      if (this.$route.query.type) {
        this.methodTypeSelection = this.$route.query.type
      }

      if (!this.$route.query.lang && !this.$route.query.type) {
        this.apiMethods = []
      }
    },
    formChange() {
      let query = {}
      if (this.languageSelection) {
        query.lang = this.languageSelection
      }
      if (this.methodTypeSelection) {
        query.type = this.methodTypeSelection
      }

      this.$router.push(
        {
          path: '/quest-api-explorer',
          query: query
        }
      ).catch(() => {
      })
    },
    getLanguageKey() {
      if (this.languageSelection == "perl") {
        return "perl_api"
      }
      if (this.languageSelection == "lua") {
        return "lua_api"
      }
      return ""
    },
    languageSelect: function () {
      if (this.methods[this.getLanguageKey()].methods) {
        let options = []
        options.push({value: null, text: "Select a type"})
        Object.keys(this.methods[this.getLanguageKey()].methods).sort().filter((item) => {
          return !item.includes("Deprecated")
        }).forEach((option) => {
          options.push({value: option, text: option})
        })

        this.methodTypeOptions = options

        if (!this.methodTypeSelection) {
          this.methodTypeSelection = this.methodTypeOptions[0]
        }
      }
      this.methodDisplay = ""
      this.methodTypeSelect()
    },
    loadExamples: function (method) {
      this.displayExamples = this.linkedExamples[method + '(']
    },
    methodTypeSelect: function () {

      this.apiMethods = []
      let apiMethods  = []

      // used to search sources for examples
      let methodSearchTerms = []

      if (this.methods[this.getLanguageKey()].methods) {
        let methodDisplay = ""
        this.codeClass    = this.getLanguageKey()
        const methods     = this.methods[this.getLanguageKey()].methods[this.methodTypeSelection]
        const snakeCase   = string => {
          return string.replace(/\W+/g, " ")
                       .split(/ |\B(?=[A-Z])/)
                       .map(word => word.toLowerCase())
                       .join('_');
        };

        if (methods && methods.length > 0) {
          methods.forEach((method) => {
            let prefix = method.method_type
            prefix     = prefix.replace("NPC", "Npc")
            let comment;

            // perl
            if (this.languageSelection === "perl") {
              if (prefix == "quest") {
                prefix = prefix.replace(prefix, "quest::")
              } else {
                prefix = prefix.replace(prefix, "$" + snakeCase(prefix).toLowerCase() + "->")
              }
            }

            // lua
            if (this.languageSelection === "lua") {
              if (prefix == "eq") {
                prefix = prefix.replace(prefix, "eq.")
              } else {
                prefix = prefix.replace(prefix, snakeCase(prefix).toLowerCase() + ":")
              }

              if (method.return_type !== "") {
                comment = "-- return " + method.return_type
              }
            }

            apiMethods.push({
              method: method.method,
              methodPrefix: prefix,
              params: method.params,
              comments: comment,
            })

            // we use the bracket after the method to grep in searches because
            // we can pick up false positives in comments and other unrelated things
            methodSearchTerms.push(method.method + "(")

          })
        }

        this.apiMethods    = apiMethods
        this.methodDisplay = methodDisplay

        SpireApiClient.v1().post('/quest-api/examples/search-projecteq', {
          "search_terms": methodSearchTerms,
          "language": this.languageSelection
        }).then((response) => {
          if (response.data && response.data.data) {
            // console.log(response.data.data)

            response.data.data.forEach((result) => {
              if (typeof this.linkedExamples[result.search_term] === "undefined") {
                this.linkedExamples[result.search_term] = []
              }

              result.full_example = " " + result.before_content + result.line_match + result.after_content

              this.linkedExamples[result.search_term].push(result)
            })
          }
          this.$forceUpdate()
        });


      }
    }
  }
}

</script>

<style>
.code-display {
  font-size: 14px !important;
  max-width: 100% !important;
  width:     100%;
}

.example-preview {
  position: fixed;
  right:    30px;
  top:      2%;
}

.example-preview-inner {
  max-height: 90vh;
  overflow-y: scroll;
}
</style>
