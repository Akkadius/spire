<template>
  <div>
    <eq-window-simple
      title="Quest API Explorer"
      class="p-2"
    >
      <div v-if="loaded">

        <!-- Form -->
        <div class="row">
          <div class="col-lg-1 col-sm-12 text-center">
            Language
            <b-form-select
              v-model="languageSelection"
              :options="languageOptions"
              @change="languageReset(); updateQueryState();"
            />
          </div>
          <div class="col-lg-2 col-sm-12 text-center">
            Types
            <b-form-select
              v-model="methodTypeSelection"
              @change="methodTypeSelectReset(); updateQueryState(); "
              :options="methodTypeOptions"
            />
          </div>
          <div class="col-lg-2 col-sm-12 text-center">
            Events
            <b-form-select
              v-model="eventSelection"
              @change="eventSelectReset(); updateQueryState();"
              :options="eventOptions"
            />
          </div>
          <div class="col-lg-2 col-sm-12 text-center">
            Constants
            <b-form-select
              v-model="constantSelection"
              @change="constantSelectReset(); updateQueryState();"
              :options="constantOptions"
            />
          </div>
          <div :class="'col-lg-' + (appEnvLocal ? '3' : '4') + ' col-sm-12 text-center'">
            Search
            <b-input
              id="quest-explorer-search"
              v-model="search"
              v-on:keyup="optionLoaded = false; onSearch(); onSearchMethodExampleLoad()"
              placeholder="Search for methods, events (soon constants)..."
              autofocus
            />
          </div>
          <div class="col-lg-1 col-sm-12 text-center" v-if="appEnvLocal">
            <b-button
              variant="outline-warning"
              @click="refreshDefinitions"
              size="sm"
              style="margin-top: 20px"
            >
              <i class="fa fa-refresh"></i> Refresh
            </b-button>
          </div>
          <div class="col-lg-1 col-sm-12 text-center mt-3">
            <span class="font-weight-bold">Last Updated</span>
            <div>{{ fromNow(api.last_refreshed) }}</div>
          </div>

        </div>

      </div>

      <app-loader :is-loading="!loaded" padding="4"/>

      <eq-debug :data="api" v-if="Object.keys(api).length"/>

    </eq-window-simple>

    <eq-window-simple
      v-if="eventSelection || apiMethods.length > 0 || (languageSelection && constantSelection && api) || (search.length > 0 && optionLoaded)"
      class="p-0"
    >
      <!-- Display Quest API Methods -->
      <div class="row mt-2" v-if="apiMethods.length > 0">

        <!-- Methods -->
        <div class="col-12">

          <div
            class="mb-4 code-display"
            style="padding-left: 10px !important; padding-top: 10px !important; padding-bottom: 10px !important; white-space:nowrap"
            v-if="apiMethods.length > 0"
          >

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
          <quest-api-display-events
            :language-selection="languageSelection"
            :events="[eventSelection]"
            :linked-examples="linkedExamples"
            :api="api"
            @load-quest-example="loadQuestExample"
          />
        </div>
      </div>

      <!-- Constants -->
      <div
        class="row mt-2 pl-2"
        v-if="languageSelection && constantSelection && api && Object.keys(api).length > 0"
        style="display:inline-block"
      >
        <div class="col-12">
          <quest-api-display-constants
            :language-selection="languageSelection"
            :constant-selection="constantSelection"
            :linked-examples="linkedExamples"
            :api="api"
            @load-quest-example="loadQuestExample"
          />
        </div>
      </div>

      <!-- Search Results -->
      <div v-if="search.length > 0 && optionLoaded">
        <div
          class="mt-3"
          v-if="search.length > 0 && searchApiResultMethods.length === 0 && searchEventsResult.length === 0"
        >
          No results found...
        </div>

        <!-- Search: Methods -->
        <div class="row mt-2" v-if="searchApiResultMethods.length > 0">

          <!-- Methods -->
          <div class="col-12">

            <h5 class="eq-header">Methods ({{ (searchApiResultMethods.length) }})</h5>

            <div
              class="mb-4 code-display"
              style="padding-left: 10px !important; padding-top: 10px !important; padding-bottom: 10px !important; white-space:nowrap"
            >

              <quest-api-display-methods
                :language-selection="languageSelection"
                :linked-examples="linkedExamples"
                :api-methods="searchApiResultMethods"
                @load-quest-example="loadQuestExample"
              />

            </div>
          </div>
        </div>

        <!-- Search: Events -->
        <div class="row mt-2" v-if="searchEventsResult.length > 0">

          <!-- Events -->
          <div class="col-12">

            <h5 class="eq-header">Events ({{ (searchEventsResult.length) }})</h5>

            <quest-api-display-events
              :language-selection="languageSelection"
              :events="searchEventsResult"
              :linked-examples="linkedExamples"
              :api="api"
              @load-quest-example="loadQuestExample"
            />

          </div>

        </div>

      </div>

      <app-loader :is-loading="!optionLoaded" padding="4"/>
    </eq-window-simple>

    <!-- Example Previews -->
    <div class="row mt-2 pl-2">
      <div class="col-lg-8 col-sm-12 example-preview" v-if="displayExamples.length > 0">

        <eq-window-simple
          :title="'Examples (' + displayExamples.slice(0,50).length + ')'"
          class="mt-3"
          @click="closeExample"
        >
          <div style="right: 30px; top: -10px; position: absolute; z-index: 999999">
            <a
              style="color: white"
              @click="closeExample"
            >(Esc) X
            </a>
          </div>

          <div class="example-preview-inner ">
            <div
              v-for="(example, index) in displayExamples.slice(0,50)"
              :key="example.file_name + index + example.line_number"
              class="mr-6"
            >
              <div class="row ">
                <div class="col">
                  <a
                    href="javascript:;"
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
                      class="btn btn-primary btn-sm mr-2 "
                    >
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

        </eq-window-simple>

      </div>
    </div>

    <div class="card mt-3" v-if="isZeroState()">
      <div class="card-body">

        <div class="header mt-md-1 mb-3">
          <h1>
            <img src="@/assets/img/vscode-logo.png" style="height: 50px; width: auto" class="mr-3">
            Spire Quest API Completions - VSCode Plugin
          </h1>

          <p class="header-subtitle">
            Like the Quest API explorer? Have all of the EverQuest Emulator server Quest API readily available and always up to date within
            your editor by using Visual Studio Code and using Spire's plugin extension all completely for free!
          </p>

        </div>

        <a href="https://marketplace.visualstudio.com/items?itemName=Akkadius.eqemu-spire-quest-api&ssr=false#overview" target="_blank">
          <img src="https://github.com/EQEmu/Server/assets/3319450/55d8a923-47d8-4584-aede-f3aa9220fbaa" style="width: 100%; border-radius: 10px">
        </a>

<!--        <iframe-->
<!--          frameborder="0"-->
<!--          allowtransparency="true"-->
<!--          style="width: 100%; height: 75vh; border-radius: 10px"-->
<!--          src="https://marketplace.visualstudio.com/items?itemName=Akkadius.eqemu-spire-quest-api&ssr=false#overview"-->
<!--        />-->

      </div>
    </div>

  </div>

</template>

<script type="ts">
import EqWindow from "@/components/eq-ui/EQWindow.vue";
import DebugDisplayComponent from "@/components/DebugDisplayComponent.vue";
import {SpireApi} from "../../app/api/spire-api";
import axios from "axios";
import EqWindowSimple from "@/components/eq-ui/EQWindowSimple.vue";
import EqTabs from "@/components/eq-ui/EQTabs.vue";
import EqTab from "@/components/eq-ui/EQTab.vue";
import slugify from "slugify";
import moment from "moment";
import * as util from "util";
import QuestApiDisplayMethods from "@/views/quest-api-explorer/components/QuestApiDisplayMethods.vue";
import Analytics from "@/app/analytics/analytics";
import QuestApiDisplayConstants from "@/views/quest-api-explorer/components/QuestApiDisplayConstants.vue";
import QuestApiDisplayEvents from "@/views/quest-api-explorer/components/QuestApiDisplayEvents.vue";
import EqDebug from "@/components/eq-ui/EQDebug.vue";
import Debug from "@/app/debug/debug";
import {debounce} from "@/app/utility/debounce";
import {ROUTE} from "@/routes";
import {AppEnv} from "@/app/env/app-env";
import ContentArea from "@/components/layout/ContentArea.vue";
import {LocalSettings, Setting} from "@/app/local-settings/localsettings";

export default {
  components: {
    ContentArea,
    EqDebug,
    QuestApiDisplayEvents,
    QuestApiDisplayConstants,
    QuestApiDisplayMethods,
    EqTab,
    EqTabs,
    EqWindowSimple,
    DebugDisplayComponent,
    EqWindow,
    editor: require('vue2-ace-editor'),
    "page-header": () => import("@/components/layout/PageHeader.vue")
  },
  data() {
    return {
      codeClass: "",

      appEnvLocal: false,

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
      optionLoaded: true,

      // linked method examples
      linkedExamples: {
        perl: {},
        lua: {}
      },
      // examples being displayed for current method
      displayExamples: [],

      // keep last query params so when our watcher detects state
      // we don't trigger load methods that don't need to load when
      // certain query param values haven't changed
      lastQueryParamState: {},

      // search
      search: "",
      searchApiResultMethods: [],
      searchEventsResult: [],
      searchConstantsResult: []
    }
  },
  watch: {
    $route(to, from) {
      this.loadQueryState()
    }
  },

  destroyed() {
    Debug.log("[deactivated]")

    document.body.removeEventListener('keyup', this.closeExampleKeyHandler)
  },
  async mounted() {
    Debug.log("[activated]")

    this.reset()
    await this.loadDefinitions();
    this.loadQueryState()
    this.init()
  },
  methods: {
    isZeroState() {
      return Object.keys(this.$route.query).length === 0
    },

    refreshDefinitions() {
      this.loaded = false

      SpireApi.v1().post('quest-api/refresh-definitions').then((response) => {
        if (response.status === 200) {
          this.init()
        }
      });
    },

    onSearch: debounce(function () {
      Debug.log("[onSearch] [%s]", this.search)

      if (this.search.trim() === "") {
        this.optionLoaded = true
        return
      }

      this.doSearch()
    }, 500),

    doSearch: function () {
      // reset
      this.apiMethods          = []
      this.methodTypeSelection = null
      this.eventSelection      = null
      this.constantSelection   = null
      this.displayExamples     = []

      if (this.search.trim() === "") {
        return
      }

      Debug.log("[search] [%s]", this.search)

      this.optionLoaded = false

      this.updateQueryState()

      let apiMethods = []
      for (const [key, value] of Object.entries(this.api[this.languageSelection].methods)) {
        this.api[this.languageSelection].methods[key].forEach((method) => {

          let search = this.search.toLowerCase()

          // search in params
          let foundInParam = false
          method.params.forEach((param) => {
            if (foundInParam) {
              return true
            }

            if (param.includes(search)) {
              foundInParam = true
            }
          })

          if (method.method.toLowerCase().includes(search) || foundInParam) {
            // console.log()
            apiMethods.push(this.getFormattedApiMethod(method))
          }
        })
      }

      let searchEventsResult = []
      this.api[this.languageSelection].events.forEach((event) => {
        if (event.event_identifier.toLowerCase().includes(this.search.toLowerCase())) {
          const eventIdentifier = event.entity_type + "-" + event.event_identifier
          searchEventsResult.push(eventIdentifier)
        }
      })

      this.searchEventsResult     = searchEventsResult
      this.searchApiResultMethods = apiMethods

      if (this.search.length >= 3) {
        Analytics.trackCountsEvent("quest_api_search", this.search)
      }

      this.optionLoaded = true
    },

    onSearchMethodExampleLoad: debounce(function () {
      Debug.log("[onSearchMethodExampleLoad] [%s]", this.search)

      // used to search sources for examples
      let methodSearchTerms  = []
      // keeps track of duplicate search terms
      let declaredSearchTerm = {}

      // methods
      this.searchApiResultMethods.forEach((method) => {

        // we use the bracket after the method to grep in searches because
        // we can pick up false positives in comments and other unrelated things
        if (!declaredSearchTerm[method.method]) {
          methodSearchTerms.push(method.method + "(")
          declaredSearchTerm[method.method] = true
        }
      })

      this.loadSearchExamples(methodSearchTerms)
    }, 500),

    fromNow(time) {
      return moment(time).fromNow()
    },
    closeExampleKeyHandler(e) {
      if (e.keyCode === 27) {
        this.closeExample()
      }
    },
    languageReset() {
      delete this.$route.query.h
      this.methodTypeSelection    = null
      this.methodTypeOptions      = []
      this.eventSelection         = null
      this.eventOptions           = []
      this.constantSelection      = null
      this.constantOptions        = []
      this.displayExamples        = []
      this.search                 = ""
      this.searchApiResultMethods = []
    },

    reset() {
      this.languageSelection      = null
      this.methodTypeSelection    = null
      this.methodTypeOptions      = []
      this.eventSelection         = null
      this.eventOptions           = []
      this.constantSelection      = null
      this.constantOptions        = []
      this.displayExamples        = []
      this.search                 = ""
      this.searchApiResultMethods = []
      this.loaded                 = false
    },

    async loadDefinitions() {
      const r = await SpireApi.v1().get('/quest-api/definitions')
      if (r.data && r.data.data) {
        this.api    = r.data.data
        this.loaded = true
      }
    },

    init() {
      Debug.blankLine()
      Debug.log("[init]")

      // get app env
      this.appEnvLocal = AppEnv.isAppLocal()

      // escape handler
      document.body.addEventListener('keyup', this.closeExampleKeyHandler)

      this.loadDefinitions()
    },
    loadQueryState() {
      Debug.log("[loadQueryState] trigger")

      this.languageSelection   = null
      this.methodTypeSelection = null
      this.eventSelection      = null
      this.constantSelection   = null
      this.search              = ""

      if (!this.$route.query.lang) {
        this.methodTypeOptions = []
        this.eventOptions      = []
        this.constantOptions   = []
        this.displayExamples   = []
      }

      const langSetting = LocalSettings.get(Setting.DEFAULT_LANGUAGE_PREFERENCE);
      if (this.$route.query.lang) {
        this.languageSelection = this.$route.query.lang
        if (this.lastQueryParamState.lang !== this.$route.query.lang) {
          setTimeout(this.methodTypeSelect, 50)
          setTimeout(() => {
            let search = document.getElementById("quest-explorer-search")
            if (search) {
              search.focus()
            }
          }, 100)
        }
        this.languageSelect()
      }
      // load default user preference if query string not set
      else if (langSetting !== "") {
        this.languageSelection = langSetting
        this.languageSelect()
      }

      if (this.$route.query.type) {
        this.methodTypeSelection = this.$route.query.type
        if (this.lastQueryParamState.type !== this.$route.query.type) {
          this.methodTypeSelect()
        }
      }
      if (this.$route.query.event) {
        this.eventSelection = this.$route.query.event
        if (this.lastQueryParamState.event !== this.$route.query.event) {
          Debug.log("[loadQueryState] eventSelect")
          this.eventSelect()
        }
      }
      if (this.$route.query.constant) {
        this.constantSelection = this.$route.query.constant

        if (this.lastQueryParamState.constant !== this.$route.query.constant) {
          Debug.log("[loadQueryState] constantSelect")
          this.constantSelect()
        }
      }

      if (this.$route.query.q) {
        this.search = this.$route.query.q

        if (this.lastQueryParamState.q !== this.$route.query.q) {
          Debug.log("[loadQueryState] search")
          this.doSearch()
          this.onSearchMethodExampleLoad()
        }
      }

      if (!this.$route.query.lang && !this.$route.query.type) {
        this.apiMethods = []
      }

      this.lastQueryParamState = this.$route.query
    },
    updateQueryState() {
      Debug.blankLine()
      Debug.log("[updateQueryState] trigger")
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
      if (this.search) {
        query.q = this.search
      }
      if (this.search.length === 0) {
        delete query.q
      }
      if (this.$route.query.h) {
        query.h = this.$route.query.h
      }

      this.$router.push(
        {
          path: ROUTE.QUEST_API_EXPLORER,
          query: query
        }
      ).catch(() => {
      })
    },
    navigateTo(url, name) {
      window.open(url, name);
    },
    closeExample: function () {
      this.displayExamples = []
    },

    // when language is selected
    languageSelect: function () {
      this.constantSelection = null
      this.search            = ""

      if (this.languageSelection.length !== 0) {
        LocalSettings.set(Setting.DEFAULT_LANGUAGE_PREFERENCE, this.languageSelection)
      }

      // methods
      if (this.api[this.languageSelection].methods) {
        let options  = []
        let types    = Object.keys(this.api[this.languageSelection].methods).sort().filter((item) => {
          return !item.includes("Deprecated") && !item.includes("EQDB")
        })
        let typeSize = types.length
        options.push({value: null, text: "--- Types (" + typeSize + ") ---"})
        types.forEach((option) => {
          let methodCount = this.api[this.languageSelection].methods[option].length
          options.push({value: option, text: option + ' (' + methodCount + ')'})
        })

        this.methodTypeOptions = options
        if (!this.methodTypeSelection) {
          this.methodTypeSelection = null
        }
      }

      // events
      if (this.api[this.languageSelection].events) {
        let events       = this.api[this.languageSelection].events
        let eventSize    = events.length
        let selectOption = '--- Events   (' + eventSize + ') ---'
        let options      = [{value: null, text: selectOption}]
        events.forEach((option, index) => {
          if (option.event_identifier === "event_") {
            return
          }

          if (option.event_identifier != "") {
            let text = "[" + option.entity_type + "] " + option.event_identifier;
            options.push({value: option.entity_type + '-' + option.event_identifier, text: text})
          }
        })
        this.eventOptions   = options
        this.eventSelection = null
      }

      // constants
      if (this.api[this.languageSelection].constants) {
        let apiConstants = this.api[this.languageSelection].constants
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
    },
    slug: function (toSlug) {
      return slugify(toSlug.replace(/[&\/\\#, +()$~%.'":*?<>{}]/g, "-"))
    },
    // received from component event
    loadQuestExample(event) {
      console.log(event)
      this.closeExample()
      if (this.linkedExamples[this.languageSelection][event.search]) {
        this.displayExamples = this.linkedExamples[this.languageSelection][event.search]

        this.trackAnalytics("load_quest_example_" + this.languageSelection, event.search.replace("(", ""))
      }

    },
    // when method is clicked; loads editor examples
    loadConstantExamples: function (search) {
      this.displayExamples = this.linkedExamples[this.languageSelection][search]

      this.trackAnalytics("load_quest_constant_example_" + this.languageSelection, search)
    },
    editorInit: async function (slug, lineNumber) {
      Debug.log("[editorInit] trigger")

      // ace
      require('brace/ext/language_tools')
      require('brace/theme/terminal')
      require('brace/mode/lua')
      require('brace/mode/perl')

      setTimeout(() => {
        // console.log(slug)
        // console.log(this.$refs[slug][0].editor)
        // console.log(lineNumber)

        // console.log(this.$refs)

        this.$refs[slug][0].editor.setFontSize(16)
        this.$refs[slug][0].editor.setReadOnly(true);
        this.$refs[slug][0].editor.scrollToLine(lineNumber, true, true, () => {
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

      // this.$refs.myEditor.editor.setFontSize(16)
      // this.$refs.editorslug.editor.setReadOnly(true);
      // this.$refs.editorslug.editor.gotoLine(1, lineNumber, true);
    },

    methodTypeSelectReset: function () {
      Debug.log("[methodTypeSelectReset] trigger")
      this.constantSelection = null
      this.eventSelection    = null
      this.search            = ""
    },
    eventSelectReset: function () {
      Debug.log("[eventSelectReset] trigger")
      this.methodTypeSelection = null
      this.constantSelection   = null
      this.search              = ""
    },
    constantSelectReset: function () {
      Debug.log("[constantSelectReset] trigger")
      this.methodTypeSelection = null
      this.eventSelection      = null
      this.search              = ""
    },

    // formats expected api rendered format from the API response
    getFormattedApiMethod(method) {
      const snakeCase = string => {
        return string.replace(/\W+/g, " ")
          .split(/ |\B(?=[A-Z])/)
          .map(word => word.toLowerCase())
          .join('_');
      };

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

      return {
        method: method.method,
        methodPrefix: prefix,
        params: method.params,
        comments: comment,
        categories: method.categories,
      }
    },

    // when a method type is selected
    methodTypeSelect: function () {
      Debug.log("[methodTypeSelect] trigger")

      // reset other displays
      let apiMethods  = []
      this.apiMethods = []

      // this.updateQueryState()

      // used to search sources for examples
      let methodSearchTerms  = []
      // keeps track of duplicate search terms
      let declaredSearchTerm = {}

      // methods
      if (this.api[this.languageSelection].methods) {
        this.codeClass = this.languageSelection
        const methods  = this.api[this.languageSelection].methods[this.methodTypeSelection]
        if (methods && methods.length > 0) {
          methods.forEach((method) => {
            apiMethods.push(this.getFormattedApiMethod(method))

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
        this.trackAnalytics("quest_navigate_methods_" + this.languageSelection, this.methodTypeSelection)
      }

    },
    // when an event type is selected
    eventSelect: async function () {
      Debug.log("[eventSelect] trigger")

      // reset other displays
      this.apiMethods          = []
      this.methodTypeSelection = null
      this.constantSelection   = null

      // update browser / route state
      // setTimeout(() => {
      //   this.updateQueryState(), 100
      // });

      const event        = this.eventSelection.split("-")[1]
      const searchString = (event.trim() + "(")
      await this.loadSearchExamples([searchString])

      let searchEvent = {
        search: searchString
      }

      this.loadQuestExample(searchEvent)

      this.trackAnalytics("quest_navigate_events_" + this.languageSelection, this.eventSelection)
    },
    // when an constant is selected
    constantSelect: function () {
      Debug.log("[constantSelect] trigger")

      // reset other displays
      this.apiMethods          = []
      this.eventSelection      = null
      this.methodTypeSelection = null

      // // update browser / route state
      // setTimeout(() => {
      //   this.updateQueryState(), 100
      // });

      if (this.constantSelection) {
        let apiConstants = this.api[this.languageSelection].constants[this.constantSelection]
        let searchTerms  = []

        apiConstants.forEach((constant) => {
          const constantVal = this.languageSelection == "lua" ? this.constantSelection + '.' + constant.constant : constant.constant;
          searchTerms.push(constantVal)
        })

        this.loadSearchExamples(searchTerms)

        this.trackAnalytics("quest_navigate_constants_" + this.languageSelection, this.constantSelection)
      }
    },
    loadSearchExamples: async function (searchTerms) {
      Debug.log("[loadSearchExamples] trigger")

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
      const response    = await SpireApi.v1().post(exampleRepo, {
        "search_terms": searchTerms,
        "language": this.languageSelection
      })

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
    },
    trackAnalytics(name, value) {
      Debug.log("[trackAnalytics] trigger")
      Analytics.trackAllEvents(name, value)
    },
  }
}

</script>

<style>
.code-display {
  font-size: 14px !important;
  max-width: 100% !important;
  width: 100%;
}

.example-preview {
  position: fixed;
  left: 20%;
  top: 2%;
  z-index: 9999999
}

.example-preview-inner {
  max-height: 90vh;
  overflow-y: scroll;
}
</style>
