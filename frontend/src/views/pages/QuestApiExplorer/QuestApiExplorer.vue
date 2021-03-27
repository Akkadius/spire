<template>
  <!-- CONTENT -->
  <div class="container-fluid">
    <div class="panel-body">
      <div class="panel panel-default">
        <div class="row">
          <div class="col-12">
            <eq-window style="margin-top: 30px" title="Quest API Explorer">

              <h1 class="eq-header text-center mt-5" v-if="!loaded">LOADING... PLEASE WAIT...</h1>

              <app-loader :is-loading="!loaded" padding="4"/>

              <div v-if="loaded">

                <!-- Form -->
                <div class="row">
                  <div class="col-1 text-center">
                    Language
                    <b-form-select
                      v-model="languageSelection"
                      :options="languageOptions"
                      @change="languageSelect();"/>
                  </div>
                  <div class="col-2 text-center">
                    Types
                    <b-form-select
                      v-model="methodTypeSelection"
                      @change="formChange();"
                      :options="methodTypeOptions"/>
                  </div>
                  <div class="col-3 text-center">
                    Events
                    <b-form-select
                      v-model="eventSelection"
                      @change="eventSelect();"
                      :options="eventOptions"/>
                  </div>
                  <div class="col-2 text-center">
                    Constants
                    <b-form-select
                      v-model="constantSelection"
                      @change="formChange();"
                      :options="constantOptions"/>
                  </div>

                  <div class="col text-center align-middle ">
                    <div class="float-right">
                      Last Updated: {{ fromNow(api.last_refreshed) }}
                    </div>
                  </div>
                </div>

                <!-- Display Quest API Methods -->
                <div class="row mt-2" v-if="apiMethods.length > 0">

                  <!-- Methods -->
                  <div class="col-12">

                    <div
                      class="mb-4 code-display"
                      style="padding-left: 10px !important; padding-top: 10px !important; padding-bottom: 10px !important; white-space:nowrap"
                      v-if="apiMethods.length > 0">

                      <quest-api-display-methods
                        :language-selection="languageSelection"
                        :linked-examples="linkedExamples"
                        :api-methods="apiMethods"
                        @load-quest-example="loadQuestExample"
                      />

                    </div>
                  </div>

                </div>

                <!-- Events -->
                <div
                  class="row mt-2 pl-2"
                  v-if="eventSelection"
                  style="display:inline-block"
                >
                  <div class="col-12">

                    <!-- Lua -->
                    <pre
                      v-if="languageSelection === 'lua'"
                      class="highlight html bg-dark ml-0 mb-4 code-display"
                      style="width: 50vw; display: inline-block; padding-left: 20px !important; padding-top: 10px !important; padding-bottom: 10px !important"
                    ><span style="color: #57A64A">-- {{ eventSelectionFormatName() }}</span>
function {{ getSelectedEvent().event_identifier }}(e)
	<span style="color: #57A64A">-- Exported event variables</span>
<span v-for="(e, index) in eventVars()" :key="index">	eq.debug("{{ e }} " .. e.{{ e }});
</span>end</pre>
                    <!-- Perl -->
                    <pre
                      v-if="languageSelection === 'perl'"
                      class="highlight html bg-dark ml-0 mb-4 code-display"
                      style="width: 50vw; display: inline-block; padding-left: 20px !important; padding-top: 10px !important; padding-bottom: 10px !important"
                    ><span style="color: #57A64A"># {{ eventSelectionFormatName() }}</span>
sub {{ getSelectedEvent().event_identifier }} {
	<span style="color: #57A64A"># Exported event variables</span>
<span v-for="(e, index) in eventVars()" :key="index">	quest::debug("{{ e }} " . ${{ e }});
</span>}</pre>
                  </div>

                </div>

                <!-- Constants -->
                <div
                  class="row mt-2 pl-2"
                  v-if="languageSelection && constantSelection && api && Object.keys(api).length > 0"
                  style="display:inline-block"
                >
                  <div class="col-12">

                    <!-- Perl -->
                    <pre
                      v-if="languageSelection === 'perl'"
                      class="highlight html bg-dark ml-0 mb-4 code-display"
                      style="width: 50vw; display: inline-block; padding-left: 20px !important; padding-top: 10px !important; padding-bottom: 10px !important"
                    ><div
                      v-for="(constant, index) in api[getLanguageKey()].constants[constantSelection]"
                      :key="index">{{ constant.constant }}</div></pre>

                    <!-- Lua -->
                    <pre
                      v-if="languageSelection === 'lua'"
                      class="highlight html bg-dark ml-0 mb-4 code-display"
                      style="min-width: 30vw; display: inline-block; padding-left: 20px !important; padding-top: 10px !important; padding-bottom: 10px !important"
                    ><div
                      v-for="(constant, index) in api[getLanguageKey()].constants[constantSelection]"
                      :key="index"><a
                      @click="loadConstantExamples(formatConstant(constant))" href="javascript:void(0);"
                      v-if="linkedExamples[languageSelection][formatConstant(constant)]">{{
                        formatConstant(constant)
                      }}</a><span
                      v-if="!linkedExamples[languageSelection][formatConstant(constant)]">{{
                        formatConstant(constant)
                      }}</span></div></pre>
                  </div>
                </div>

                <!-- Example Previews -->
                <div class="row mt-2 pl-2">
                  <div class="col-6 example-preview" v-if="displayExamples.length > 0">

                    <eq-window :title="'Examples (' + displayExamples.slice(0,50).length + ')'">

                      <button class="btn btn-white btn-sm" @click="closeExample"
                              style="right: 30px; top: 10px; position: absolute">X
                      </button>

                      <div class="example-preview-inner">

                        <div v-for="(example, index) in displayExamples.slice(0,50)"
                             :key="example.file_name + index + example.line_number">
                          <div class="row ">
                            <div class="col">
                              <a href="javascript:;"
                                 class="ml-5"
                                 @click="navigateTo('https://github.com/' + example.org + '/' + example.repo + '/blob/' + example.branch + '/' + encodeURIComponent(example.file_name) + '#L' + example.line_number, example.file_name)"
                                 style="color: #b9b194"
                              >
                                <i class="fe fe-github"></i>
                                {{ example.file_name + ":" + example.line_number }}
                              </a>
                            </div>

                            <div class="col">
                              <div class="float-right">
                                <button
                                  @click="navigateTo('https://github.com/' + example.org + '/' + example.repo)"
                                  style="line-height: 1.40;"
                                  class="btn btn-primary btn-sm mr-2 ">
                                  <i class="fe fe-github"></i>
                                  {{ example.org + " / " + example.repo }}
                                </button>
                              </div>
                            </div>
                          </div>

                          <editor
                            v-model="example.full_contents"
                            @init="editorInit(slug(example.file_name), example.line_number)"
                            :id="slug(example.file_name)"
                            :lang="languageSelection"
                            theme="terminal"
                            width="100%"
                            :ref="slug(example.file_name)"
                            height="275px"
                            class="mt-3 mb-3 ml-5 pr-6"
                          ></editor>

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
import EqWindow               from "@/components/eq-ui/EQWindow.vue";
import DebugDisplayComponent  from "@/components/DebugDisplayComponent.vue";
import {SpireApiClient}       from "@/app/api/spire-api-client";
import axios                  from "axios";
import EqWindowSimple         from "@/components/eq-ui/EQWindowSimple.vue";
import EqTabs                 from "@/components/eq-ui/EQTabs.vue";
import EqTab                  from "@/components/eq-ui/EQTab.vue";
import slugify                from "slugify";
import moment                 from "moment";
import * as util              from "util";
import QuestApiDisplayMethods from "@/views/pages/QuestApiExplorer/components/QuestApiDisplayMethods.vue";
import Analytics              from "@/app/analytics/analytics";


export default {
  components: {
    QuestApiDisplayMethods,
    EqTab,
    EqTabs,
    EqWindowSimple,
    DebugDisplayComponent,
    EqWindow,
    editor: require('vue2-ace-editor'),
    "page-header": () => import("@/views/layout/PageHeader.vue")
  },
  data() {
    return {
      codeClass: "",

      // languages:select
      languageSelection: null,
      languageOptions: [
        {value: null, text: ''},
        {value: 'perl', text: 'Perl'},
        {value: 'lua', text: 'Lua'},
      ],

      // methods:select
      methodTypeSelection: null,
      methodTypeOptions: [],

      // events:select
      eventSelection: null,
      eventOptions: [],

      // constants:select
      constantSelection: null,
      constantOptions: [],

      // response data
      api: {},

      // intermediary data for display
      apiMethods: {},
      apiConstants: {},

      // route watcher
      routeWatcher: null,

      loaded: false,

      // linked method examples
      linkedExamples: {
        perl: {},
        lua: {}
      },
      // examples being displayed for current method
      displayExamples: [],

    }
  },
  deactivated() {
    // remove route watcher
    this.routeWatcher()

    document.body.removeEventListener('keyup', this.closeExampleKeyHandler)
  },
  activated() {
    this.init()
  },
  methods: {
    fromNow(time) {
      return moment(time).fromNow()
    },
    closeExampleKeyHandler(e) {
      if (e.keyCode === 27) {
        this.closeExample()
      }
    },
    init() {

      // reset
      this.languageSelection   = null
      this.methodTypeSelection = null
      this.methodTypeOptions   = []
      this.eventSelection      = null
      this.eventOptions        = []
      this.constantSelection   = null
      this.constantOptions     = []
      this.displayExamples     = []
      this.loaded              = false

      // route watcher
      this.routeWatcher = this.$watch('$route.query', () => {
        this.loadQueryParams()
      });

      // escape handler
      document.body.addEventListener('keyup', this.closeExampleKeyHandler)

      // load data from api
      SpireApiClient.v1().get('/quest-api/definitions').then((response) => {
        if (response.data && response.data.data) {
          this.api    = response.data.data
          this.loaded = true
          this.loadQueryParams()
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
    },
    loadQueryParams() {
      this.languageSelection   = null
      this.methodTypeSelection = null
      this.eventSelection      = null
      this.constantSelection   = null

      if (this.$route.query.lang) {
        this.languageSelection = this.$route.query.lang
        this.languageSelect()
      }
      if (this.$route.query.type) {
        this.methodTypeSelection = this.$route.query.type
        this.methodTypeSelect()
      }
      if (this.$route.query.event) {
        this.eventSelection = this.$route.query.event
      }
      if (this.$route.query.constant) {
        this.constantSelection = this.$route.query.constant
        this.constantSelect()
      }

      if (!this.$route.query.lang && !this.$route.query.type) {
        this.apiMethods = []
      }

      this.lastQueryParamState = this.$route.query
    },
    formChange() {
      let query = {}
      if (this.languageSelection) {
        query.lang = this.languageSelection
      }
      if (this.methodTypeSelection) {
        query.type = this.methodTypeSelection
      }
      if (this.eventSelection) {
        query.event = this.eventSelection
      }
      if (this.constantSelection) {
        query.constant = this.constantSelection
      }
      if (this.$route.query.m) {
        query.m = this.$route.query.m
      }

      this.$router.push(
        {
          path: '/quest-api-explorer',
          query: query
        }
      ).catch(() => {
      })
    },
    navigateTo(url, name) {
      window.open(url, name);
    },
    formatConstant(constant) {
      return this.languageSelection == "lua" ? this.constantSelection + '.' + constant.constant : constant.constant;
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
    closeExample: function () {
      this.displayExamples = []
    },
    eventSelectionFormatName() {
      const entity = this.eventSelection.split("-")[0]
      const event  = this.eventSelection.split("-")[1]
      return util.format("[%s] %s", entity, event)
    },
    getSelectedEvent() {
      const entity = this.eventSelection.split("-")[0]
      const event  = this.eventSelection.split("-")[1]
      let events   = this.api[this.getLanguageKey()].events
      let e        = []
      events.forEach(row => {
        if (row.event_identifier === "event_") {
          return
        }

        if (row.event_identifier == event && row.entity_type == entity) {
          e = row
        }

      })

      return e
    },
    eventVars() {
      return this.getSelectedEvent().event_vars
    },

    // when language is selected
    languageSelect: function () {
      this.constantSelection = null

      // methods
      if (this.api[this.getLanguageKey()].methods) {
        let options  = []
        let types    = Object.keys(this.api[this.getLanguageKey()].methods).sort().filter((item) => {
          return !item.includes("Deprecated") && !item.includes("EQDB")
        })
        let typeSize = types.length
        options.push({value: null, text: "--- Types (" + typeSize + ") ---"})
        types.forEach((option) => {
          let methodCount = this.api[this.getLanguageKey()].methods[option].length
          options.push({value: option, text: option + ' (' + methodCount + ')'})
        })

        this.methodTypeOptions = options
        if (!this.methodTypeSelection) {
          this.methodTypeSelection = null
        }
      }

      // events
      if (this.api[this.getLanguageKey()].events) {
        let events       = this.api[this.getLanguageKey()].events
        let eventSize    = events.length
        let selectOption = '--- Events   (' + eventSize + ') ---'
        let options      = [{value: null, text: selectOption}]
        events.forEach((option, index) => {
          if (option.event_identifier === "event_") {
            return
          }

          let text = "[" + option.entity_type + "] " + option.event_identifier;
          options.push({value: option.entity_type + '-' + option.event_identifier, text: text})
        })
        this.eventOptions   = options
        this.eventSelection = null
      }

      // constants
      if (this.api[this.getLanguageKey()].constants) {
        let apiConstants = this.api[this.getLanguageKey()].constants
        let options      = []
        let constants    = Object.keys(apiConstants).sort().filter((constant) => {
          return !constant.includes("Deprecated") && !constant.includes("EQDB")
        })
        let typeSize     = constants.length
        options.push({value: null, text: "--- Select a constant group (" + typeSize + ") ---"})
        constants.forEach((category) => {
          let methodCount = apiConstants[category].length
          options.push({value: category, text: category + ' (' + methodCount + ')'})
        })

        this.constantOptions = options
        if (!this.constantSelection) {
          this.constantSelection = null
        }
      }

      // update browser / route state
      setTimeout(() => {
        this.formChange(), 100
      });
    },
    slug: function (toSlug) {
      return slugify(toSlug.replace(/[&\/\\#, +()$~%.'":*?<>{}]/g, "-"))
    },
    // received from component event
    loadQuestExample(event) {
      this.loadExamples(event.search)
    },
    // when method is clicked; loads editor examples
    loadExamples: function (method) {
      this.closeExample()
      if (this.linkedExamples[this.languageSelection][method + '(']) {
        this.displayExamples = this.linkedExamples[this.languageSelection][method + '(']

        this.sendAllAnalytics("load_quest_example_" + this.languageSelection, method)
      }
    },
    // when method is clicked; loads editor examples
    loadConstantExamples: function (search) {
      this.displayExamples = this.linkedExamples[this.languageSelection][search]

      this.sendAllAnalytics("load_quest_constant_example_" + this.languageSelection, search)
    },
    editorInit: async function (slug, lineNumber) {


      setTimeout(() => {
        // console.log(slug)
        // console.log(this.$refs[slug][0].editor)
        // console.log(lineNumber)

        // console.log(this.$refs)

        this.$refs[slug][0].editor.setFontSize(13)
        this.$refs[slug][0].editor.setReadOnly(true);
        this.$refs[slug][0].editor.scrollToLine(lineNumber, true, true, function () {
        });
        this.$refs[slug][0].editor.gotoLine(lineNumber, 0, true);

        const Range = ace.acequire('ace/range').Range;

        this.$refs[slug][0].editor.selection.setRange(
          new Range(
            lineNumber - 1,
            0,
            lineNumber,
            0
          )
        );
      }, 10)


      // this.loaded = false;
      require('brace/ext/language_tools')
      require('brace/theme/terminal')
      require('brace/mode/lua')
      require('brace/mode/perl')
      // this.$refs.myEditor.editor.setFontSize(16)
      // this.$refs.editorslug.editor.setReadOnly(true);
      // this.$refs.editorslug.editor.gotoLine(1, lineNumber, true);
    },
    // when a method type is selected
    methodTypeSelect: function () {

      // reset other displays
      this.apiMethods        = []
      let apiMethods         = []
      this.constantSelection = null
      this.eventSelection    = null
      this.formChange()

      // used to search sources for examples
      let methodSearchTerms  = []
      // keeps track of duplicate search terms
      let declaredSearchTerm = {}

      // methods
      if (this.api[this.getLanguageKey()].methods) {
        this.codeClass  = this.getLanguageKey()
        const methods   = this.api[this.getLanguageKey()].methods[this.methodTypeSelection]
        const snakeCase = string => {
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
                comment = "-- @return " + method.return_type
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
            if (!declaredSearchTerm[method.method]) {
              methodSearchTerms.push(method.method + "(")
              declaredSearchTerm[method.method] = true
            }
          })
        }

        this.apiMethods = apiMethods

        this.loadSearchExamples(methodSearchTerms)

        this.sendAllAnalytics("quest_navigate_methods_" + this.languageSelection, this.methodTypeSelection)
      }

    },
    // when an event type is selected
    eventSelect: function () {

      // reset other displays
      this.apiMethods          = []
      this.methodTypeSelection = null
      this.constantSelection   = null

      // update browser / route state
      setTimeout(() => {
        this.formChange(), 100
      });

      this.sendAllAnalytics("quest_navigate_events_" + this.languageSelection, this.eventSelection)
    },
    // when an constant is selected
    constantSelect: function () {

      // reset other displays
      this.apiMethods          = []
      this.eventSelection      = null
      this.methodTypeSelection = null

      // // update browser / route state
      // setTimeout(() => {
      //   this.formChange(), 100
      // });

      if (this.constantSelection) {
        let apiConstants = this.api[this.getLanguageKey()].constants[this.constantSelection]
        let searchTerms  = []

        apiConstants.forEach((constant) => {
          const constantVal = this.languageSelection == "lua" ? this.constantSelection + '.' + constant.constant : constant.constant;
          searchTerms.push(constantVal)
        })

        this.loadSearchExamples(searchTerms)

        this.sendAllAnalytics("quest_navigate_constants_" + this.languageSelection, this.constantSelection)
      }
    },
    loadSearchExamples: function (searchTerms) {

      // examples
      // reset
      let linkedExamples = {
        perl: {},
        lua: {}
      }

      // fixed for now, but designed to be able to source from multiple locations
      const org         = "ProjectEQ"
      const repo        = "projecteqquests"
      const branch      = "master"
      const exampleRepo = util.format('/quest-api/source-examples/org/%s/repo/%s/branch/%s', org, repo, branch)
      SpireApiClient.v1().post(exampleRepo, {
        "search_terms": searchTerms,
        "language": this.languageSelection
      }).then((response) => {
        if (response.data && response.data.data) {
          // console.log(response.data.data)

          response.data.data.forEach((result) => {
            if (typeof linkedExamples[this.languageSelection][result.search_term] === "undefined") {
              linkedExamples[this.languageSelection][result.search_term] = []
            }

            result.org    = org
            result.repo   = repo
            result.branch = branch

            linkedExamples[this.languageSelection][result.search_term].push(result)
          })
        }
        this.linkedExamples = linkedExamples
        this.$forceUpdate()
      });

    },
    sendAllAnalytics(name, value) {
      this.sendAnalyticsEvent(name, value)
      this.sendAnalyticsCountEvent(name, value)
    },
    sendAnalyticsEvent(name, value) {
      Analytics.trackEvent(name, value)
    },
    sendAnalyticsCountEvent(name, key) {
      Analytics.trackCountsEvent(name, key)
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
