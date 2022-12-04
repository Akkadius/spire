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
import MonacoEditor from 'vue-monaco'
import {SpireApi}   from "../../app/api/spire-api";
import axios        from "axios";
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

    SpireApi.v1().get('/quest-api/definitions').then((response) => {
      if (response.data && response.data.data) {
        let api = response.data.data
        let languages = ['lua', 'perl']
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
            let methodPrefix         = m.methodPrefix
            if (language === "lua") {
              if (m.methodPrefix !== "eq") {
                methodPrefix         = "";
                languageQuestPrepend = ":"
              }
            }

            if (language === "perl") {
              languageQuestPrepend = "::"

              if (m.methodPrefix !== "quest") {
                methodPrefix         = "";
                // m.methodPrefix = `\$ ` + m.methodPrefix.toLowerCase()
                languageQuestPrepend = ""
              }
            }

            // this is label friendly, doesn't include completion indexes
            let methodLabel = util.format(
              "%s%s%s(%s)",
              methodPrefix,
              languageQuestPrepend,
              m.method,
              m.params.join(", ")
            )

            // method snippet contains completion indexes for user to fill out param values
            let methodSnippet = util.format(
              "%s%s%s(%s)",
              methodPrefix,
              languageQuestPrepend,
              m.method,
              completionParams.join(", ")
            )

            // console.log(methodLabel)

            if (methodSnippet.includes("$")) {
              // console.log(methodSnippet)
            }

            let completionSnippet = {
              label: methodLabel,
              // filterText: methodLabel,
              kind: monaco.languages.CompletionItemKind.Function,
              eqemuClass: m.methodPrefix.toLowerCase(),
              insertText: methodSnippet,

              insertTextRules: monaco.languages.CompletionItemInsertTextRule.InsertAsSnippet
            }

            console.log(completionSnippet)

            suggestions.push(completionSnippet)
          })

          // console.log(suggestions)

          gSuggestions[language] = suggestions

          monaco.languages.registerCompletionItemProvider(language, {
            triggerCharacters: ['>'],
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

          const config = {
            surroundingPairs: [
              { open: '{', close: '}' },
              { open: '[', close: ']' },
              { open: '(', close: ')' },
              { open: '<', close: '>' },
              // { open: "'", close: "'" },
              // { open: '"', close: '"' },
            ],
            autoClosingPairs: [
              { open: '{', close: '}' },
              { open: '[', close: ']' },
              { open: '(', close: ')' },
              // { open: "'", close: "'", notIn: ['string', 'comment'] },
              // { open: '"', close: '"', notIn: ['string', 'comment'] },
            ],
          };
          monaco.languages.setLanguageConfiguration('lua', config);
          monaco.languages.setLanguageConfiguration('perl', config);


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

      result = await (new ZoneApi(...SpireApi.cfg())).listZones({ orderBy: "zoneidnumber" })
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

      result = await (new NpcTypeApi(...SpireApi.cfg())).listNpcTypes({ orderBy: "id" })
      if (result.status === 200) {
        result.data.forEach((n) => {
          suggestions.push({
            label: "NPC." + n.name + " (Returns NPC ID)",
            kind: monaco.languages.CompletionItemKind.Constant,
            insertText: n.id.toString(),
          })
        })
      }

      result = await (new TaskApi(...SpireApi.cfg())).listTasks({ orderBy: "id" })
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

      let editor = this.editor;


      const currentLine       = editor.getModel().getLineContent(editor.getPosition().lineNumber)
      const line              = editor.getPosition().lineNumber;
      const col               = editor.getPosition().column + 1;
      const textUntilPosition = editor.getModel().getValueInRange({
        startLineNumber: editor.getPosition().lineNumber,
        startColumn: 0,
        endLineNumber: line,
        endColumn: col
      });

      const currentPos = textUntilPosition.length;
      let searchFilter = ""

      const wordSplit = textUntilPosition.split(" ")
      if (wordSplit.length > 0) {
        let lastWord = wordSplit[wordSplit.length - 1]

        const lastWordContainsClosingParenthesis =
                lastWord.includes("(") && lastWord.includes(")")

        if (lastWord.includes("$") && !lastWordContainsClosingParenthesis) {
          let wordSplit = lastWord.split("$")
          if (wordSplit.length > 0) {
            let secondSplit = wordSplit[1].trim().split("->")
            if (secondSplit.length > 0) {
              searchFilter = secondSplit[0].trim()
              console.log("search filter is [%s]", searchFilter)
            }
          }
        }

        console.log("last word is [%s]", lastWord)
      }

      console.log("[currentLine] [%s]", currentLine)
      console.log("[textUntilPosition] [%s]", textUntilPosition)

      // sometimes a class inherits others... like mobs
      let searchClasses = [searchFilter]
      if (["client", "npc"].includes(searchFilter)) {
        searchClasses = [searchFilter, "mob"]
      } else if (searchFilter.toLowerCase().includes("entity_list")) {
        searchClasses = ["entitylist"]
      } else if (searchFilter.toLowerCase().includes("door")) {
        searchClasses = ["doors"]
      }

      let suggestions = gSuggestions[this.language]
      for (let key in suggestions) {
        if (searchFilter !== "") {
          // eqemuClass

          // we're searching on class objects, if there doesn't exist any then lets skip
          if (!suggestions[key].eqemuClass) {
            continue;
          }

          // make sure our last word contains our class prefix
          if (!searchClasses.includes(suggestions[key].eqemuClass)) {
            continue;
          }
        }

        suggestions[key].range = range
      }

      // console.log(suggestions)

      return suggestions
    },
    editorDidMount(editor) {
      // console.log("editor mounted")
      // console.log(editor)

      this.editor = editor

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
        // autoClosingBrackets: false,
        autoIndent: true,
      },
      editor: null,
    }
  }
}
</script>
