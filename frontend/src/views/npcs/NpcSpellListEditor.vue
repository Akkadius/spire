<template>
  <div>
    <div class="row">
      <div class="col-7">
        <eq-window
          v-if="spellSet && spellSet.npc_spells_entries && spellSet.npc_spells_entries.length"
          :title="`NPC Spells Editor ID (${spellSet.id}) [${spellSet.name}] Count (${spellSet.npc_spells_entries.length})`"
          class="p-0"
        >
          <div class="row mb-2">
            <div class="col-6">
              <div class="mt-4 ml-3">
                <b-button
                  @click="addSpellListEntry()"
                  size="sm"
                  variant="outline-warning"
                >
                  <i class="fa fa-plus mr-1"></i>
                  Add Spell
                </b-button>
              </div>
            </div>

            <div class="col-6">
              <!-- Notification / Error -->
              <info-error-banner
                class="mr-1"
                style="margin-top: 20px"
                :notification="notification"
                :error="error"
                :slim="true"
                @dismiss-error="error = ''"
                @dismiss-notification="notification = ''"
              />
            </div>

          </div>

          <div
            id="npc-spell-preview-list-viewport"
            style="max-height: 90vh; overflow-y: scroll; overflow-x: hidden"
          >
            <npc-spell-preview
              :spells="spellSet"
              :edit-buttons="true"
              :highlighted-spell="highlightedSpellId"
              @reload-parent="init()"
              @edit-spell="editSpell($event)"
              @error="error = $event"
              @notification="notification = $event"
            />
          </div>
        </eq-window>
      </div>
      <div class="col-5">
        <div
          style="width: auto;"
          class="fade-in"
          v-if="selectorActive['spell-selector']"
        >
          <spell-selector
            @input="addSpellToList($event.spellId)"
          />
        </div>

        <div v-if="editingSpellEntryId > 0">
          <eq-window>
            We're editing!
          </eq-window>
        </div>
      </div>
    </div>

  </div>
</template>

<script>
import EqWindow            from "../../components/eq-ui/EQWindow";
import NpcSpellPreview     from "../../components/preview/NpcSpellPreview";
import {SpireQueryBuilder} from "../../app/api/spire-query-builder";
import {NpcSpellApi}       from "../../app/api";
import {SpireApiClient}    from "../../app/api/spire-api-client";
import SpellSelector       from "../../components/selectors/SpellSelector";
import {NpcSpellsEntryApi} from "../../app/api/api/npc-spells-entry-api";
import {scrollToTarget}    from "../../app/utility/scrollToTarget";
import InfoErrorBanner     from "../../components/InfoErrorBanner";
import {ROUTE}             from "../../routes";
import util                from "util";

export default {
  name: "NpcSpellListEditor",
  components: { InfoErrorBanner, SpellSelector, NpcSpellPreview, EqWindow },
  data() {
    return {
      spellSet: {},

      // local notification / error state
      notification: "",
      error: "",

      // preview / selectors
      selectorActive: {},

      // editing
      editingSpellEntryId: 0,
      editingSpellEntry: {}, // object itself

      lastResetTime: Date.now(),

      highlightedSpellId: 0,
    }
  },
  async mounted() {
    this.init()
  },
  watch: {
    '$route'() {
      console.log("route trigger")
      this.reset()
      this.init()
    },
  },

  methods: {
    /**
     * State
     */
    updateQueryState: function () {
      let queryState = {};

      if (this.editingSpellEntryId !== 0) {
        queryState.id = this.editingSpellEntryId
      }

      const npcSpellsId = parseInt(this.$route.params.id)

      this.$router.push(
        {
          path: util.format(ROUTE.NPC_SPELL_EDIT, npcSpellsId),
          query: queryState
        }
      ).catch(() => {
      })
    },

    loadQueryState() {
      console.log("loading query state")
      if (this.$route.query.id >= 0) {
        this.editingSpellEntryId = parseInt(this.$route.query.id);
      }

      console.log("query params", this.$route.query)
    },

    editSpell(e) {
      console.log("we got it!", e)
      this.editingSpellEntryId = e.id
      this.updateQueryState()
    },

    reset() {
      this.resetPreviewComponents()
      this.resetNotifications()
      this.highlightedSpellId  = 0
      this.editingSpellEntryId = 0
    },

    resetNotifications() {
      this.error = ""
      this.notification = ""
    },

    async init() {
      this.loadQueryState()
      this.spellSet = await this.getNpcSpell()

      if (this.editingSpellEntryId > 0) {
        const selectedSpellEntry = this.spellSet.npc_spells_entries.find((e) => {
          return e.id === this.editingSpellEntryId
        })

        if (Object.keys(selectedSpellEntry).length > 0) {
          this.editingSpellEntsry = selectedSpellEntry
        }
      }
    },

    addSpellListEntry() {
      console.log("add spell list entry")
      this.setSelectorActive('spell-selector')
    },
    async addSpellToList(spellId) {
     this.resetNotifications()
      console.log(this.$route.params)

      try {
        const npcSpellsId = parseInt(this.$route.params.id)
        const api         = (new NpcSpellsEntryApi(SpireApiClient.getOpenApiConfig()))
        const r           = await api.createNpcSpellsEntry(
          {
            npcSpellsEntry: {
              "npc_spells_id": npcSpellsId,
              "spellid": parseInt(spellId),
              "type": 2,
              "minlevel": 51,
              "maxlevel": 59,
              "manacost": -1,
              "recast_delay": 60,
              "priority": 1,
              "resist_adjust": null,
              "min_hp": 0,
              "max_hp": 0,
            }
          })

        // pass to the child component the highlighted spell id
        this.scrollToSpellListEntry(r.data.id)
        this.notification = "Added spell to list!"
        this.init()

        console.log("highlighted spell id", this.highlightedSpellId)
      } catch (err) {
        if (err.response.data.error) {
          if (err.response && err.response.data && err.response.data.error) {
            if (err.response.data.error.includes("Duplicate")) {
              console.log("We're getting spell", spellId)

              // find the spell entry in the stack by spell ID and then scroll to it
              const selectedSpellEntry = this.spellSet.npc_spells_entries.find((e) => {
                return e.spells_new && e.spells_new.id === spellId
              })

              if (Object.keys(selectedSpellEntry).length > 0) {
                this.scrollToSpellListEntry(selectedSpellEntry.id)
              }

              this.error = "This spell is already in this list!"
              return;
            }
            this.error = "Error! " + err.response.data.error
          }
        }
      }
    },

    scrollToSpellListEntry(entryId) {
      this.highlightedSpellId = entryId

      setTimeout(() => {
        scrollToTarget(
          "npc-spell-preview-list-viewport",
          'spell-list-entry-' + entryId
        )
      }, 500)
    },

    async getNpcSpell() {
      let builder = (new SpireQueryBuilder())
        .includes([
          "NpcSpellsEntries.SpellsNew",
        ])

      try {
        const r = await (new NpcSpellApi(SpireApiClient.getOpenApiConfig()))
          .getNpcSpell(
            {
              id: this.$route.params.id,
            },
            {
              query: builder.get()
            }
          )
        if (r.status === 200 && r.data) {
          return r.data
        }
      } catch (err) {
        if (err.response.data.error) {
          if (err.response && err.response.data && err.response.data.error) {
            this.error = "Error! " + err.response.data.error
          }
        }
      }
    },

    // selectors
    resetPreviewComponents() {
      for (const [k, v] of Object.entries(this.selectorActive)) {
        this.selectorActive[k] = false
      }
    },
    setSelectorActive(selector) {
      this.resetPreviewComponents()
      this.lastResetTime            = Date.now()
      this.selectorActive[selector] = true
      this.$forceUpdate()
    }
  },
}
</script>

<style scoped>

</style>
