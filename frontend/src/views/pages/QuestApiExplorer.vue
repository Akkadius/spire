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
                      @change="formChange(); languageSelect();"/>
                  </div>
                  <div class="col-2 text-center">
                    Types
                    <b-form-select
                      v-model="methodTypeSelection"
                      @change="formChange(); methodTypeSelect();"
                      :options="methodTypeOptions"/>
                  </div>
                  <div class="col-3 text-center">
                    Events
                    <b-form-select
                      v-model="eventSelection"
                      @change="formChange(); eventSelect();"
                      :options="eventOptions"/>
                  </div>

                  <div class="col text-center align-middle ">
                    <div class="float-right">
                      Last Updated: {{ fromNow(methods.last_refreshed) }}
                    </div>
                  </div>
                </div>

                <!-- Display Quest API Methods -->
                <div class="row mt-2" v-if="apiMethods.length > 0">
                  <div class="col-12">

                    <!-- Everthing below will look a bit messy because things need -->
                    <!-- to be on the same line for proper pre formatting -->
                    <pre
                      class="highlight html bg-dark hljs mb-4 code-display"
                      style="padding-left: 20px !important; padding-top: 10px !important; padding-bottom: 10px !important"
                      v-if="apiMethods.length > 0"><div
                      v-for="(method, index) in apiMethods"
                      :key="index"><span v-if="method.methodPrefix" style="color:#9CDCFE">{{
                        method.methodPrefix
                      }}</span><span v-if="!linkedExamples[languageSelection][method.method + '(']">{{
                        method.method
                      }}</span><a
                      @click="loadExamples(method.method)" href="javascript:void(0);"
                      v-if="linkedExamples[languageSelection][method.method + '(']">{{ method.method }}</a>({{
                        method.params.join(", ")
                      }}); <span
                        v-if="method.comments" style="color: #57A64A">{{ method.comments }}</span><span
                        style="color: #57A64A"
                        v-if="linkedExamples[languageSelection][method.method + '('] && linkedExamples[languageSelection][method.method + '('].length > 0">{{
                          (languageSelection === "perl" ? "#" : "--")
                        }} {{
                          linkedExamples[languageSelection][method.method + "("].length
                        }} Example(s)</span></div></pre>
                  </div>

                  <div class="col-6 example-preview" v-if="displayExamples.length > 0">

                    <eq-window title="Examples">

                      <button class="btn btn-white btn-sm" @click="closeExample"
                              style="right: 30px; top: 10px; position: absolute">X
                      </button>

                      <div class="example-preview-inner">

                        <div v-for="(example, index) in displayExamples.slice(0,50)"
                             :key="example.file_name + index + example.line_number">
                          <div class="row">
                            <div class="col">
                              <a href="javascript:;"
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
                                  class="btn btn-primary btn-sm mr-2">
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
                            height="250px"
                            class="mt-3 mb-3"
                          ></editor>

                        </div>
                      </div>

                    </eq-window>

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
import slugify               from "slugify";
import moment                from "moment";
import * as util             from "util";

export default {
  components: {
    EqTab,
    EqTabs,
    EqWindowSimple,
    DebugDisplayComponent,
    EqWindow,
    editor: require('vue2-ace-editor'),
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

      // select
      eventSelection: null,
      eventOptions: [],

      apiMethods: [],


      // route watcher
      routeWatcher: null,

      loaded: false,

      // linked method examples
      linkedExamples: {
        perl: {},
        lua: {}
      },
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
    fromNow(time) {
      return moment(time).fromNow()
    },
    init() {
      this.routeWatcher = this.$watch('$route.query', () => {
        this.loadQueryParams()
      });

      this.loaded = false

      SpireApiClient.v1().get('/quest-api/methods', this.database).then((response) => {
        if (response.data && response.data.data) {
          this.methods = response.data.data

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
      if (this.eventSelection) {
        query.event = this.eventSelection
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
      let events   = this.methods[this.getLanguageKey()].events
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
    languageSelect: function () {

      // methods
      if (this.methods[this.getLanguageKey()].methods) {
        let options  = []
        let types    = Object.keys(this.methods[this.getLanguageKey()].methods).sort().filter((item) => {
          return !item.includes("Deprecated") && !item.includes("EQDB")
        })
        let typeSize = types.length
        options.push({value: null, text: "--- Select a Type (" + typeSize + ") ---"})
        types.forEach((option) => {
          let methodCount = this.methods[this.getLanguageKey()].methods[option].length
          options.push({value: option, text: option + ' (' + methodCount + ')'})
        })

        this.methodTypeOptions = options
        if (!this.methodTypeSelection) {
          this.methodTypeSelection = null
        }
      }

      // events
      if (this.methods[this.getLanguageKey()].events) {
        let events       = this.methods[this.getLanguageKey()].events
        let eventSize    = events.length
        let selectOption = '--- Select Event (' + eventSize + ') ---'
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

    },
    slug: function (toSlug) {
      return slugify(toSlug.replace(/[&\/\\#, +()$~%.'":*?<>{}]/g, "-"))
    },
    loadExamples: function (method) {
      // destroy old editors
      // for (let refsKey in this.$refs) {
      //   console.log(refsKey)
      //   console.log(this.$refs[refsKey][0])
      //   if (this.$refs[refsKey] && this.$refs[refsKey][0].editor) {
      //     this.$refs[refsKey][0].editor.destroy()
      //   }
      // }

      this.displayExamples = this.linkedExamples[this.languageSelection][method + '(']

    },
    editorInit: async function (slug, lineNumber) {

      this.$refs[slug][0].editor.setFontSize(14)

      setTimeout(() => {
        // console.log(slug)
        // console.log(this.$refs[slug][0].editor)
        // console.log(lineNumber)

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
    methodTypeSelect: function () {

      // reset other displays
      this.apiMethods     = []
      let apiMethods      = []
      this.eventSelection = null
      this.formChange()

      // used to search sources for examples
      let methodSearchTerms  = []
      // keeps track of duplicate search terms
      let declaredSearchTerm = {}

      // methods
      if (this.methods[this.getLanguageKey()].methods) {
        this.codeClass  = this.getLanguageKey()
        const methods   = this.methods[this.getLanguageKey()].methods[this.methodTypeSelection]
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

        // quest-api/source-examples/org/:org/repo/:repo/:branch

        // fixed for now, but designed to be able to source from multiple locations
        const org    = "ProjectEQ"
        const repo   = "projecteqquests"
        const branch = "master"

        // reset
        this.linkedExamples = {
          perl: {},
          lua: {}
        }

        SpireApiClient.v1().post(util.format('/quest-api/source-examples/org/%s/repo/%s/branch/%s', org, repo, branch), {
          "search_terms": methodSearchTerms,
          "language": this.languageSelection
        }).then((response) => {
          if (response.data && response.data.data) {
            // console.log(response.data.data)

            response.data.data.forEach((result) => {
              if (typeof this.linkedExamples[this.languageSelection][result.search_term] === "undefined") {
                this.linkedExamples[this.languageSelection][result.search_term] = []
              }

              result.org    = org
              result.repo   = repo
              result.branch = branch

              this.linkedExamples[this.languageSelection][result.search_term].push(result)
            })
          }
          this.$forceUpdate()
        });


        this.displayExamples = []

      }
    },
    eventSelect: function () {

      // reset other displays
      this.apiMethods          = []
      this.methodTypeSelection = null
      this.formChange()


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
