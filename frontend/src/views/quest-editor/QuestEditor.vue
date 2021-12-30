<template>
  <div>
    <MonacoEditor
      style="width:100%;height:100vh;"
      class="editor"
      :options="options"
      v-model="code"
      language="lua"
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

let gSuggestions = []

export default {
  components: {
    MonacoEditor
  },

  async mounted() {
    SpireApiClient.v1().get('/quest-api/definitions').then((response) => {
      if (response.data && response.data.data) {
        let api         = response.data.data
        let suggestions = [];
        let language    = 'lua'

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

          // this is label friendly, doesn't include completion indexes
          let methodLabel = util.format(
            "%s.%s(%s)",
            m.methodPrefix,
            m.method,
            m.params.join(", ")
          )

          // method snippet contains completion indexes for user to fill out param values
          let methodSnippet = util.format(
            "%s.%s(%s)",
            m.methodPrefix,
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

        gSuggestions = suggestions

        monaco.languages.registerCompletionItemProvider('lua', {
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

      gSuggestions = gSuggestions.concat(suggestions)

      // console.log(gSuggestions)
    }

  },

  methods: {
    getSuggestions(range) {
      // console.log(range)

      let suggestions = gSuggestions
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
      code: `-- same for vxed and tipt
local compass     = { zone="barindu", x=-1249.24, y=575.142, z=-148.257 }
local safereturn  = { zone="barindu", x=-1242.00, y=456.00, z=-121.76, h=0.0 }
local vxed_zonein = { x=-427.0, y=-3552.0, z=14.99, h=463.0 }
local tipt_zonein = { x=-448.0, y=-2374.0, z=13.16, h=395.0 }

local vxed = {
  expedition = { name="Vxed, The Crumbling Caverns", min_players=1, max_players=6 },
  instance   = { zone="vxed", version=0, duration=eq.seconds("6h") },
  compass    = compass,
  safereturn = safereturn,
  zonein     = vxed_zonein
}

local tipt = {
  expedition = { name="Tipt, Treacherous Crags", min_players=1, max_players=6 },
  instance   = { zone="tipt", version=0, duration=eq.seconds("6h") },
  compass    = compass,
  safereturn = safereturn,
  zonein     = tipt_zonein
}

local paladin_epic = {
  expedition = { name="Vxed, The Crumbling Caverns", min_players=6, max_players=24 },
  instance   = { zone="vxed", version=1, duration=eq.seconds("6h") },
  compass    = compass,
  safereturn = safereturn,
  zonein     = vxed_zonein
}

local function create_expedition(client, expedition_info)
  local dz = client:CreateExpedition(expedition_info)
  if dz.valid then
    dz:AddReplayLockout(eq.seconds("5m"))
  end
end

function event_say(e)
  local qglobals = eq.get_qglobals(e.other)
  local has_vxed_access = (qglobals.god_vxed_access and qglobals.god_vxed_access == "1") -- sewers or rondo complete
  local has_tipt_access = (qglobals.god_tipt_access and qglobals.god_tipt_access == "1") -- has_permanent_vxed

  local sewers_flag = tonumber(eq.get_data(e.other:CharacterID() .. "-god_sewers")) or 0
  eq.debug(string.format("sewers[%s] vxed_access[%s] tipt_access[%s]", sewers_flag, tostring(has_vxed_access), tostring(has_tipt_access)))

  local is_gm = (e.other:Admin() > 80 and e.other:GetGM())

  -- Smith Rondo skip flag was not tested on live, will just bypass sewers progression dialogue here
  if e.message:findi("hail") then
    if e.other:GetClass() == Class.PALADIN and e.other:HasItem(69933) then -- Item: Seal of Enic
      e.other:Message(MT.NPCQuestSay, string.format("I heard you released Reiya from his tourture, %s. I have seen muramites gathering in Vxed and I fear this may have to do with Reiya and the creatures responsible. Go there now and investigate, Noble Knight.", e.other:GetCleanName()))
      create_expedition(e.other, paladin_epic)
    elseif sewers_flag == 4 and not has_tipt_access then -- sewers progression dialogue
      -- note: live dialogue is always here if finished sewers, checking tipt access flag here avoids this live "bug"
      e.other:Message(MT.NPCQuestSay, "Udranda looks around sheepishly.  'Greetings.  A messenger told me you would come.  I believe we owe you our thanks -- though I'm not sure allowing you to face the terror in the mountains is a reward.  I will have my stone worker allow you to pass through, but you must [" .. eq.say_link("take heed") .. "] first and be as quiet as you can.'")
    else
      e.other:Message(MT.NPCQuestSay, "Udranda tells you, 'Greetings.  I have a duty here to stand guard at the mountain pass and allow the Muramites to pass through by moving the rocks with my magic.'  Udranda looks around for anyone listening and whispers to you. 'If you want to go into [" .. eq.say_link("Tipt") .. "] or [" .. eq.say_link("Vxed") .. "] I will have my stone worker open the passage for you.  If you want to progress past the mountains, I would ask that you first prove your worth with High Priest Diru.'")
    end
  elseif e.message:findi("take heed") and not has_tipt_access then -- no longer responds to this after receiving permanent vxed completed flag
    if sewers_flag == 4 or has_vxed_access then
      e.other:Message(MT.NPCQuestSay, "Udranda tells you, 'Good.  There are many secrets I can't tell you right now for fear the Muramites might hear them, but you must do as I say and trust in me.  In the mountain pass, Vxed, you must seek out Stonespiritist Ekikoa.  He is my mentor and will help you in your quest to find passage to the peaks.  We have created a very unique magic that we will share with you, but you must speak with him.  It will all be clear to you when you find him.  Tell me when you are [" .. eq.say_link("ready to pass") .. "] through.  I cannot leave a way open for too long.'")
      create_expedition(e.other, vxed) -- no emote when creating here
    else -- same message as hailing without progression
      e.other:Message(MT.NPCQuestSay, "Udranda tells you, 'Greetings.  I have a duty here to stand guard at the mountain pass and allow the Muramites to pass through by moving the rocks with my magic.'  Udranda looks around for anyone listening and whispers to you. 'If you want to go into [" .. eq.say_link("Tipt") .. "] or [" .. eq.say_link("Vxed") .. "] I will have my stone worker open the passage for you.  If you want to progress past the mountains, I would ask that you first prove your worth with High Priest Diru.'")
    end
  elseif e.message:findi("ready to pass") and not has_tipt_access then
    -- only responds without permanent vxed flag, possibly old dialogue to request tipt that was removed
    e.other:Message(MT.NPCQuestSay, "Udranda tells you, 'I cannot allow you to pass with your companions.  You seem heroic enough, but you must prove your worth to High Priest Diru and our stonespiritist in the Vxed mountain pass before I can allow you through.'")
  elseif e.message:findi("tipt") and (is_gm or has_tipt_access) then
    eq.get_entity_list():MessageClose(e.self, true, 100, MT.SayEcho, e.self:GetCleanName() .. " subtly commands her stone worker to open the passage for you before any invaders can take notice.")
    create_expedition(e.other, tipt)
  elseif e.message:findi("vxed") and (is_gm or has_vxed_access) then
    eq.get_entity_list():MessageClose(e.self, true, 100, MT.SayEcho, e.self:GetCleanName() .. " subtly commands her stone worker to open the passage for you before any invaders can take notice.")
    create_expedition(e.other, vxed)
  end
end

function event_trade(e)
  local item_lib = require("items")
  item_lib.return_items(e.self, e.other, e.trade)
end
      `,
      options: {
        theme: 'vs-dark'
      }
    }
  }
}
</script>
