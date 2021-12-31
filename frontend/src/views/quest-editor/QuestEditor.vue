<template>
  <div>
    <MonacoEditor
      style="width:100%;height:100vh;"
      class="editor"
      :options="options"
      v-model="code"
      :language="language"
      @editorDidMount="editorDidMount"
    />

  </div>
</template>

<script>
import MonacoEditor                   from 'vue-monaco'
import {SpireApiClient}               from "../../app/api/spire-api-client";
import axios                          from "axios";
import * as util                      from "util";
import {NpcTypeApi, TaskApi, ZoneApi} from "../../app/api";
import {EventBus}                     from "../../app/event-bus/event-bus";

let gSuggestions     = {}
gSuggestions["perl"] = []
gSuggestions["lua"]  = []

export default {
  components: {
    MonacoEditor
  },

  created() {
    EventBus.$on("EDITOR_OPEN_FILE", this.openFile);
  },
  destroyed() {
    EventBus.$off("EDITOR_OPEN_FILE", this.openFile);
  },

  async mounted() {
    let languages = ['lua', 'perl']

    SpireApiClient.v1().get('/quest-api/definitions').then((response) => {
      if (response.data && response.data.data) {
        let api = response.data.data
        languages.forEach((language) => {
          let suggestions = [];

          // constants
          let constants = []
          for (const key in api[language].constants) {
            const entries = api[language].constants[key]

            entries.forEach((e) => {
              constants.push(
                util.format(
                  "%s.%s",
                  key,
                  e.constant
                )
              )
            })
          }

          constants.forEach((c) => {
            let completionSnippet = {
              label: c,
              kind: monaco.languages.CompletionItemKind.Constant,
              insertText: c
            }

            suggestions.push(completionSnippet)
          })

          // methods
          let apiMethods = []
          for (const [key, value] of Object.entries(api[language].methods)) {
            api[language].methods[key].forEach((method) => {
              apiMethods.push(this.getFormattedApiMethod(method))
            })
          }

          apiMethods.forEach((m) => {
            let completionParams    = []
            let autoCompletionIndex = 1;
            m.params.forEach((p) => {
              completionParams.push(
                util.format("${%s:%s}", autoCompletionIndex, p)
              )
              autoCompletionIndex++;
            })

            let languageQuestPrepend = "."
            if (language === "perl") {
              languageQuestPrepend = "::"
            }

            // this is label friendly, doesn't include completion indexes
            let methodLabel = util.format(
              "%s%s%s(%s)",
              m.methodPrefix,
              languageQuestPrepend,
              m.method,
              m.params.join(", ")
            )

            // method snippet contains completion indexes for user to fill out param values
            let methodSnippet = util.format(
              "%s%s%s(%s)",
              m.methodPrefix,
              languageQuestPrepend,
              m.method,
              completionParams.join(", ")
            )

            let completionSnippet = {
              label: methodLabel,
              // filterText: methodLabel,
              kind: monaco.languages.CompletionItemKind.Function,
              // filterText: methodLabel,
              insertText: methodSnippet,
              insertTextRules: monaco.languages.CompletionItemInsertTextRule.InsertAsSnippet
            }

            suggestions.push(completionSnippet)
          })

          // console.log(suggestions)

          gSuggestions[language] = suggestions

          monaco.languages.registerCompletionItemProvider(language, {
            // triggerCharacters: ['.', ':', ' '],
            provideCompletionItems: (model, position, context) => {
              // find out if we are completing a property in the 'dependencies' object.
              // var textUntilPosition = model.getValueInRange({
              //   startLineNumber: 1,
              //   startColumn: 1,
              //   endLineNumber: position.lineNumber,
              //   endColumn: position.column
              // });
              // var match = textUntilPosition.match(
              //   /"dependencies"\s*:\s*\{\s*("[^"]*"\s*:\s*"[^"]*"\s*,\s*)*([^"]*)?$/
              // );
              // if (!match) {
              //   return { suggestions: [] };
              // }
              var word  = model.getWordUntilPosition(position);
              var range = {
                startLineNumber: position.lineNumber,
                endLineNumber: position.lineNumber,
                startColumn: word.startColumn,
                endColumn: word.endColumn
              };
              return {
                suggestions: this.getSuggestions(range)
              };
            },
          });
        })


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

    const loadDatabaseSugggestions = true

    if (loadDatabaseSugggestions) {
      let result      = {}
      let suggestions = []

      result = await (new ZoneApi(SpireApiClient.getOpenApiConfig())).listZones({ orderBy: "zoneidnumber" })
      if (result.status === 200) {
        // console.log(result.data)

        result.data.forEach((z) => {
          suggestions.push({
            label: "Zone." + z.short_name + " (Returns Zone ID)",
            kind: monaco.languages.CompletionItemKind.Constant,
            insertText: z.zoneidnumber.toString(),
          })
          suggestions.push({
            label: "Zone." + z.long_name + " (Returns Zone ID)",
            kind: monaco.languages.CompletionItemKind.Constant,
            insertText: z.zoneidnumber.toString(),
          })
          suggestions.push({
            label: "Zone." + z.short_name + " (Returns Zone Short Name)",
            kind: monaco.languages.CompletionItemKind.Constant,
            insertText: '"' + z.short_name.toString() + '"',
          })
        })
      }

      result = await (new NpcTypeApi(SpireApiClient.getOpenApiConfig())).listNpcTypes({ orderBy: "id" })
      if (result.status === 200) {
        result.data.forEach((n) => {
          suggestions.push({
            label: "NPC." + n.name + " (Returns NPC ID)",
            kind: monaco.languages.CompletionItemKind.Constant,
            insertText: n.id.toString(),
          })
        })
      }

      result = await (new TaskApi(SpireApiClient.getOpenApiConfig())).listTasks({ orderBy: "id" })
      if (result.status === 200) {
        result.data.forEach((t) => {
          suggestions.push({
            label: "Task." + t.title + " (Returns Task ID)",
            kind: monaco.languages.CompletionItemKind.Constant,
            insertText: t.id.toString(),
          })
        })
      }

      languages.forEach((language) => {
        gSuggestions[language] = gSuggestions[language].concat(suggestions)
      });

      // console.log(gSuggestions)
    }

  },

  methods: {
    openFile(e) {
      // console.log("received file open event")
      // console.log(e)
      this.code = e.contents

      if (e.fullFileName.includes(".pl")) {
        this.language = "perl"
      }
      if (e.fullFileName.includes(".lua")) {
        this.language = "lua"
      }
    },

    getSuggestions(range) {
      // console.log(range)
      // console.log("language is " + this.language)
      // console.log(gSuggestions[this.language])

      let suggestions = gSuggestions[this.language]
      for (let key in suggestions) {
        suggestions[key].range = range
      }

      // console.log(suggestions)

      return suggestions
    },
    editorDidMount(editor) {
      // console.log("editor mounted")
      // console.log(editor)

      // editor.trigger('source - use any string you like', 'editor.action.triggerSuggest', {});

      editor.onDidType(e => {
        // console.log("trigger")
        // editor.getAction('editor.action.formatDocument').run()
        // editor.trigger("anyString", 'editor.action.formatDocument');

        // console.log("model change")
        // console.log(e)

        // console.log("[event] onDidType")
        //
        // const currentLine       = editor.getModel().getLineContent(editor.getPosition().lineNumber)
        // const line              = editor.getPosition().lineNumber;
        // const col               = editor.getPosition().column + 1;
        // const textUntilPosition = editor.getModel().getValueInRange({
        //   startLineNumber: editor.getPosition().lineNumber,
        //   startColumn: 0,
        //   endLineNumber: line,
        //   endColumn: col
        // });
        //
        // const currentPos = textUntilPosition.length;
        //
        // const wordSplit = textUntilPosition.split(" ")
        // if (wordSplit.length > 0) {
        //   let lastWord = wordSplit[wordSplit.length - 1]
        //
        //   if (lastWord.includes("eq.")) {
        //     console.log("manual trigger")
        //     // editor.trigger(lastWord.replace("eq."), 'editor.action.triggerSuggest', {});
        //     console.log(this.suggestions)
        //
        //     editor.trigger(lastWord.trim() +  ' ', 'editor.action.triggerSuggest', this.suggestions);
        //   }
        //
        //   console.log("last word is [%s]", lastWord)
        // }
        //
        // console.log("[currentLine] [%s]", currentLine)
        // console.log("[textUntilPosition] [%s]", textUntilPosition)
      })


      // Listen to `scroll` event
      editor.onDidScrollChange(e => {

        // console.log(e)
        // editor.trigger('source - use any string you like', 'editor.action.triggerSuggest', {});
      })
    },

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
        if (prefix === "quest") {
          prefix = prefix.replace(prefix, "quest::")
        } else {
          prefix = prefix.replace(prefix, "$" + snakeCase(prefix).toLowerCase() + "->")
        }
      }

      // lua
      if (this.languageSelection === "lua") {
        if (prefix === "eq") {
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
  },

  data() {
    return {
      language: "lua",
      code: `Select a file...`,
      options: {
        theme: 'vs-dark',
        autoClosingBrackets: false,
        autoIndent: true,
      }
    }
  }
}
</script>
