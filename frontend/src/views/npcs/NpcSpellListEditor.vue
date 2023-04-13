<template>
  <div>
    <div class="row">
      <div :class="(!areSelectorsActive() ? 'col-12' : 'col-7')">
        <eq-window
          v-if="spellSet"
          :title="`NPC Spells Editor ID (${spellSet.id}) [${spellSet.name}] Count (${(spellSet.npc_spells_entries ? spellSet.npc_spells_entries.length : 0)})`"
          class="p-0"
        >
          <div class="row">
            <div class="col-4">
              <div class="btn-group d-inline-block mt-4 ml-3" role="group">
                <b-button
                  size="sm"
                  variant="outline-success"
                  @click="goBack"
                >
                  <i class="fa fa-arrow-left mr-1"></i>
                  Go back to Spell Lists
                </b-button>

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

            <div class="col-8">
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

          <div class="row ml-1 mt-3">
            <div class="col-lg-3">
              Spell Set Name
              <input
                type="text"
                class="form-control"
                placeholder="Spell Set Name"
                @change="saveSpellSet"
                v-model="spellSet.name"
              >
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
      <div :class="(areSelectorsActive() ? 'col-5' : '')">
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
          <eq-window :title="`Editing NPC Spell List (${editingSpellEntryId})`">
            <div
              v-for="field in editingSpellEntryFields"
              :key="field.field"
              :class="'row minified-inputs'"
            >
              <div
                class="col-4 text-right p-0 m-0 mr-1 mt-3"
                style="position: relative; bottom: 6px;"
                v-if="field.fType === 'checkbox'"
              >
                <span v-if="field.category" class="font-weight-bold">{{ field.category }}</span>
                {{ field.desc }}
              </div>
              <div
                class="col-4 text-right p-0 m-0 mr-3"
                v-if="field.fType !== 'checkbox'"
                style="margin-top: 10px !important"
              >
                <span v-if="field.category" class="font-weight-bold">{{ field.category }}</span>
                {{ field.desc }}
              </div>

              <div class="col-7 text-left p-0 mt-2">

                <div
                  class="text-left"
                  v-if="field.field === 'spellid' && editingSpellEntry.spells_new && editingSpellEntry.spells_new.id > 0"
                >

                  <spell-popover
                    :spell="editingSpellEntry.spells_new"
                    :size="20"
                    v-if="Object.keys(editingSpellEntry.spells_new).length > 0 && editingSpellEntry.spells_new"
                    class="mt-2"
                  />

                </div>

                <!-- checkbox -->
                <div :class="'text-left ml-2 mt-1'" v-if="field.fType === 'checkbox'">
                  <eq-checkbox
                    v-b-tooltip.hover.v-dark.right :title="getFieldDescription(field.field)"
                    class="d-inline-block text-center"
                    :true-value="(typeof field.true !== 'undefined' ? field.true : 1)"
                    :false-value="(typeof field.false !== 'undefined' ? field.false : 0)"
                    v-model.number="editingSpellEntry[field.field]"
                    @input="editingSpellEntry[field.field] = $event"

                  />
                </div>

                <!-- input number -->
                <b-form-input
                  v-if="field.fType === 'number'"
                  :id="field.field"
                  v-model.number="editingSpellEntry[field.field]"
                  class="m-0 mt-1"
                  v-on="field.e ? getEventHandlers(field.e, field.field) : {}"
                  v-b-tooltip.hover.v-dark.right :title="getFieldDescription(field.field)"
                  :style="(editingSpellEntry[field.field] === 0 ? 'opacity: .5' : '')"
                />

                <!-- range -->
                <b-form-input
                  v-if="field.fType === 'range'"
                  type="range"
                  :id="field.field"
                  :min="field.min"
                  :max="field.max"
                  style="width: 80%"
                  v-model.number="editingSpellEntry[field.field]"
                  class="m-0 mt-1 d-inline-block"
                  v-on="field.e ? getEventHandlers(field.e, field.field) : {}"
                  v-b-tooltip.hover.v-dark.right :title="getFieldDescription(field.field)"
                  :style="(editingSpellEntry[field.field] <= (typeof field.zeroValue !== 'undefined' ? field.zeroValue : 0) ? 'opacity: .5' : '')"
                  @change="rerenderProbability = Date.now()"
                  :key="rerenderProbability"
                />

                <div v-if="field.fType === 'range'" class="d-inline-block ml-3" :key="rerenderProbability + '-visual'">
                  ({{ editingSpellEntry[field.field] }})
                </div>

                <!-- input text -->
                <b-form-input
                  v-if="field.fType === 'text'"
                  :id="field.field"
                  v-model="editingSpellEntry[field.field]"
                  class="m-0 mt-1"
                  v-on="field.e ? getEventHandlers(field.e, field.field) : {}"
                  v-b-tooltip.hover.v-dark.right :title="getFieldDescription(field.field)"
                  :style="(editingSpellEntry[field.field] <= (typeof field.zeroValue !== 'undefined' ? field.zeroValue : 0) ? 'opacity: .5' : '')"
                />

                <div
                  class="text-center"
                  v-if="['manacost', 'recast_delay'].includes(field.field) && parseInt(editingSpellEntry[field.field]) === -1 && editingSpellEntry.spells_new"
                >
                  <span class="font-weight-bold">(Spell) {{ field.desc }}</span>
                  {{ editingSpellEntry.spells_new[field.spellField] }}
                </div>
                <!--                {{editingSpellEntry}}-->

                <!-- textarea -->
                <b-textarea
                  v-if="field.fType === 'textarea'"
                  :id="field.field"
                  v-model="editingSpellEntry[field.field]"
                  class="m-0 mt-1"
                  rows="2"
                  max-rows="6"
                  v-on="field.e ? getEventHandlers(field.e, field.field) : {}"
                  v-b-tooltip.hover.v-dark.right :title="getFieldDescription(field.field)"
                  :style="(editingSpellEntry[field.field] === '' ? 'opacity: .5' : '') + ';'"
                ></b-textarea>

                <!-- select -->
                <select
                  v-model.number="editingSpellEntry[field.field]"
                  :id="field.field"
                  class="form-control m-0 mt-1"
                  v-if="field.selectData"
                  v-on="field.e ? getEventHandlers(field.e, field.field) : {}"
                  v-b-tooltip.hover.v-dark.right :title="getFieldDescription(field.field)"
                  :style="(editingSpellEntry[field.field] <= (typeof field.zeroValue !== 'undefined' ? field.zeroValue : 0) ? 'opacity: .5' : '')"
                >
                  <option
                    v-for="(desc, index) in field.selectData"
                    :key="index"
                    :value="parseInt(index)"
                  >
                    {{ index }}) {{ desc }}
                  </option>
                </select>

              </div>
            </div>

            <div class="text-center">
              <b-button
                @click="saveSpellListEntry()"
                size="sm"
                class="mt-3"
                variant="outline-warning"
              >
                <i class="fa fa-save"></i>
                Save
              </b-button>
            </div>


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
import {SpireApi}          from "../../app/api/spire-api";
import SpellSelector       from "../../components/selectors/SpellSelector";
import {NpcSpellsEntryApi} from "../../app/api/api/npc-spells-entry-api";
import {scrollToTarget}    from "../../app/utility/scrollToTarget";
import InfoErrorBanner     from "../../components/InfoErrorBanner";
import {ROUTE}             from "../../routes";
import util                from "util";
import EqCheckbox          from "../../components/eq-ui/EQCheckbox";
import SpellPopover        from "../../components/SpellPopover";
import {NPC_SPELL_TYPES}   from "../../app/constants/eq-npc-spells";
import {Spells}            from "../../app/spells";

export default {
  name: "NpcSpellListEditor",
  components: { SpellPopover, EqCheckbox, InfoErrorBanner, SpellSelector, NpcSpellPreview, EqWindow },
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
      rerenderProbability: 0,

      highlightedSpellId: 0,

      // fields
      editingSpellEntryFields: [
        { desc: "Spell Type", field: "type", fType: "select", selectData: NPC_SPELL_TYPES },
        { desc: "Priority", field: "priority", fType: "range", min: 0, max: 100 },
        { desc: "Mana Cost", field: "manacost", fType: "number", spellField: 'mana' },
        { desc: "Recast Delay", field: "recast_delay", fType: "number", spellField: 'recast_time' },
        { desc: "Minimum Level", field: "minlevel", fType: "number" },
        { desc: "Maximum Level", field: "maxlevel", fType: "number" },
        { desc: "Min HP", field: "min_hp", fType: "number" },
        { desc: "Max HP", field: "max_hp", fType: "number" },
        { desc: "Resist Adjust", field: "resist_adjust", fType: "text" },
      ],
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

    async saveSpellSet() {

      // clone
      let payload = JSON.parse(JSON.stringify(this.spellSet))
      delete payload.npc_spells_entries

      // save
      const api = (new NpcSpellApi(...SpireApi.cfg()))

      try {
        const r = await api.updateNpcSpell(
          {
            id: payload.id,
            npcSpell: payload
          }
        )

        if (r.status === 200) {
          this.notification = "Updated spell set!"
        }
      } catch (e) {
        console.log(e)

        if (e.response.data.error) {
          if (e.response && e.response.data && e.response.data.error) {
            this.error = "Error! " + e.response.data.error
          }
        }
      }

    },

    goBack() {
      this.$router.push(
        {
          path: ROUTE.NPC_SPELLS_EDIT
        }
      ).catch(() => {
      })
    },

    hasSpellEntries() {
      return this.spellSet.npc_spells_entries && this.spellSet.npc_spells_entries.length
    },

    areSelectorsActive() {
      return this.editingSpellEntryId > 0 || this.selectorActive['spell-selector']
    },

    async saveSpellListEntry() {
      try {
        const api = (new NpcSpellsEntryApi(...SpireApi.cfg()))
        let entry = JSON.parse(JSON.stringify(this.editingSpellEntry))
        delete entry.spells_new;

        console.log("entry is ", entry)

        const r = await api.updateNpcSpellsEntry(
          {
            id: entry.id,
            npcSpellsEntry: entry
          })

        if (r.status === 200) {
          this.notification = "Updated spell list entry!"
          this.init()
        }

      } catch (err) {
        if (err.response.data.error) {
          if (err.response && err.response.data && err.response.data.error) {
            this.error = "Error! " + err.response.data.error
          }
        }
      }
    },

    getFieldDescription(field) {
      const descriptions = {
        "priority": "Higher this number, the more likely-hood the NPC will choose to cast this Spell",
        "minlevel": "Spell only casts if NPC is above this level (default 1)",
        "maxlevel": "Spell only casts if NPC is below this level (default 255)",
        "manacost": "How much this spell costs when the NPC casts (-2 no cast time) (-1 use mana cost from spell data)",
        "recast_delay": "Seconds before NPC can recast spell (-1 - Spell default)",
        "resist_adjust": "",
        "min_hp": "Spell only casts if above this number if non-zero",
        "max_hp": "Spell only casts if below this number if non-zero",
      }

      // we do this because the payload we get back from spire API is
      // formatted slightly different
      let fieldLookup = field.toLowerCase().replace("_", "")

      for (let key in descriptions) {
        let keyLookup = key.toLowerCase().replace("_", "")
        if (keyLookup === fieldLookup) {
          return descriptions[key]
        }
      }
    },

    /**
     * Tabs / fields
     */
    getEventHandlers(e, field) {
      let handlers = {}
      if (e.onclick) {
        handlers.click = () => e.onclick(field)
      }
      if (e.onmouseover) {
        handlers.mouseover = () => e.onmouseover(field)
      }

      return handlers
    },

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
      this.error        = ""
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
          this.editingSpellEntry = selectedSpellEntry
        }

        this.scrollToSpellListEntry(this.editingSpellEntryId)
      }
    },

    addSpellListEntry() {
      this.reset()
      this.resetNotifications()
      this.updateQueryState()
      setTimeout(() => {
        this.setSelectorActive('spell-selector')
      }, 50)
    },

    calculateNpcSpellTypeFromSpell(e) {
      // console.log("calculateNpcSpellTypeFromSpell", e)

      for (let effectIndex = 1; effectIndex <= 12; effectIndex++) {

        // pets
        if ([33, 106, 71].includes(e["effectid_" + effectIndex])) {
          console.log("[npc-spell-list-editor] detected pet spell SPA [%s]", e["effectid_" + effectIndex])
          return 32;
        }

        // charm
        if ([22].includes(e["effectid_" + effectIndex])) {
          console.log("[npc-spell-list-editor] detected charm spell SPA [%s]", e["effectid_" + effectIndex])
          return 4096;
        }

        // snare
        if ([128].includes(e["effectid_" + effectIndex])) {
          console.log("[npc-spell-list-editor] detected snare spell SPA [%s]", e["effectid_" + effectIndex])
          return 128;
        }

        // dispell
        if ([209].includes(e["effectid_" + effectIndex])) {
          console.log("[npc-spell-list-editor] detected dispell spell SPA [%s]", e["effectid_" + effectIndex])
          return 512;
        }

        // mez
        if ([31].includes(e["effectid_" + effectIndex])) {
          console.log("[npc-spell-list-editor] detected mez spell SPA [%s]", e["effectid_" + effectIndex])
          return 2048;
        }

        // heal
        if (e["effectid_" + effectIndex] === 0 && e["effect_base_value_" + effectIndex] > 0 && e.good_effect === 1) {
          console.log("[npc-spell-list-editor] detected heal SPA [%s]", e["effectid_" + effectIndex])
          return 2;
        }

        // nuke
        if (e["effectid_" + effectIndex] === 0 && e["effect_base_value_" + effectIndex] > 0 && e.good_effect === 0) {
          console.log("[npc-spell-list-editor] detected nuke SPA [%s]", e["effectid_" + effectIndex])
          return 1;
        }

        // DOT
        if (e["effectid_" + effectIndex] === 0 && e["effect_base_value_" + effectIndex] !== 0 && e.good_effect === 0 && e.buffduration > 0) {
          console.log("[npc-spell-list-editor] detected DOT SPA [%s]", e["effectid_" + effectIndex])
          return 256;
        }

        // buff
        if (e.good_effect === 1) {
          console.log("[npc-spell-list-editor] detected buff SPA [%s]", e["effectid_" + effectIndex])
          return 8;
        }
      }


      // 1: "Nuke",
      //   2: "Heal",
      //   4: "Root",
      //   8: "Buff",
      //   16: "Escape",
      //   32: "Pet",
      //   64: "Lifetap",
      //   128: "Snare",
      //   256: "DOT",
      //   512: "Dispel",
      //   1024: "In-Combat Buff",
      //   2048: "Mez",
      //   4096: "Charm"

      return 1;
    },

    async addSpellToList(spellId) {
      this.resetNotifications()
      console.log(this.$route.params)

      const spell = await Spells.getSpell(spellId)

      try {
        const npcSpellsId = parseInt(this.$route.params.id)
        const api         = (new NpcSpellsEntryApi(...SpireApi.cfg()))
        const r           = await api.createNpcSpellsEntry(
          {
            npcSpellsEntry: {
              "npc_spells_id": npcSpellsId,
              "spellid": parseInt(spellId),
              "type": this.calculateNpcSpellTypeFromSpell(spell),
              "minlevel": 0,
              "maxlevel": 0,
              "manacost": -1,
              "recast_delay": -1,
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
      }, 100)
    },

    async getNpcSpell() {
      let builder = (new SpireQueryBuilder())
        .includes([
          "NpcSpellsEntries.SpellsNew",
        ])

      try {
        const r = await (new NpcSpellApi(...SpireApi.cfg()))
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
