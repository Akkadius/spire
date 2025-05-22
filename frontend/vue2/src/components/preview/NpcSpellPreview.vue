<template>
  <div v-if="spells">
    <div
      class="mt-3 text-center font-weight-bold p-3"
      v-if="spells && !spells.npc_spells_entries"
    >
      No spells found in table
    </div>

    <table
      class="eq-table bordered eq-highlight-rows"
      style="font-size: 14px; "
      v-if="spells && spells.npc_spells_entries"
      id="npc-spell-preview-list-table"
    >
      <thead class="eq-table-floating-header">
      <tr>
        <th v-if="editButtons"></th>
        <th>Spell</th>
        <th>Type</th>
        <th>Mana</th>
        <th>Priority</th>
        <th>Recast</th>
        <th>Min Lvl</th>
        <th>Max Lvl</th>
        <th>Min HP</th>
        <th>Max HP</th>
      </tr>
      </thead>
      <tbody>
      <tr
        v-for="e in spellsList"
        :id="'spell-list-entry-' + e.id"
        :key="'spell-' + e.id"
        :class="(highlightedSpell && e.id === highlightedSpell ? 'pulsate-highlight-white' : '')"
        :style="(e.is_parented ? 'background-color: rgb(0 255 255 / 20%);' : '')"
      >
        <td
          class="text-center pl-0 pr-0"
          style="min-width: 80px;"
          v-if="editButtons"
        >
          {{(e.is_parented ? `Parent List (${e.npc_spells_id})` : '')}}

          <b-button
            variant="primary"
            class="btn-dark btn-sm"
            style="padding: 0px 6px;"
            title="Edit spell entry"
            @click="editSpellListEntry(e)"
            v-if="!e.is_parented"
          >
            <i class="fa fa-pencil-square"></i>
          </b-button>

          <b-button
            variant="primary"
            class="btn btn-dark btn-sm btn-outline-danger ml-1 btn-primary"
            style="padding: 0px 6px;"
            title="Delete spell entry"
            @click="deleteSpellListEntry(e)"
            v-if="!e.is_parented"
          >
            <i class="fa fa-trash"></i>
          </b-button>
        </td>

        <td style="min-width: 220px">
          <spell-popover
            :spell="e.spells_new"
            :size="20"
            v-if="Object.keys(e.spells_new).length > 0 && e.spells_new"
            class="mt-2"
          />
        </td>
        <td>{{ (NPC_SPELL_TYPES[e.type] ? NPC_SPELL_TYPES[e.type] : "") }}</td>
        <td>{{ getManacost(e) }}</td>
        <td>{{ e.priority }}</td>
        <td>{{ getRecastDelay(e) }}s</td>
        <td>{{ e.minlevel }}</td>
        <td>{{ e.maxlevel }}</td>
        <td>{{ e.min_hp }}</td>
        <td>{{ e.max_hp }}</td>
      </tr>
      </tbody>
    </table>

  </div>
</template>

<script>
import SpellPopover        from "../SpellPopover";
import EqDebug             from "../eq-ui/EQDebug";
import {NPC_SPELL_TYPES}   from "../../app/constants/eq-npc-spells";
import {SpireQueryBuilder} from "../../app/api/spire-query-builder";
import {NpcSpellApi} from "../../app/api";
import {SpireApi}    from "../../app/api/spire-api";
import Tablesort     from "../../app/utility/tablesort";
import {NpcSpellsEntryApi} from "../../app/api/api/npc-spells-entry-api";

export default {
  name: "NpcSpellPreview",
  components: { EqDebug, SpellPopover },
  data() {
    return {
      NPC_SPELL_TYPES: NPC_SPELL_TYPES,

      spellsList: [],

      fields: {
        "type": 0,
        "minlevel": 1,
        "maxlevel": 255,
        "manacost": -1,
        "recast_delay": -1,
        "priority": 0,
        "resist_adjust": null,
        "min_hp": 0,
        "max_hp": 0,
      }
    }
  },
  watch: {
    spells: {
      deep: true,
      async handler() {
        this.init()
      }
    },
    highlightedSpell: {
      handler() {

      }
    }
  },
  props: {
    spells: {
      type: Object,
      required: false
    },
    editButtons: {
      type: Boolean,
      required: false,
      default: false,
    },
    highlightedSpell: {
      type: [Number, String],
      required: false,
      default: 0
    }
  },
  mounted() {
    this.init()
  },
  methods: {

    editSpellListEntry(e) {
      console.log("edit spell", e)
      this.$emit("edit-spell", e);
    },

    async deleteSpellListEntry(e) {
      if (confirm(`Are you sure you want to delete this spell? \n\n(${e.spells_new.name}) (${e.spells_new.id})?`)) {

        try {
          await (new NpcSpellsEntryApi(...SpireApi.cfg()))
            .deleteNpcSpellsEntry({ id: e.id })
            .then(() => {
              this.$emit("reload-parent", true);
              this.$emit("notification", "Spell deleted successfully!");
            })
        } catch (err) {
          if (err.response.data.error) {
            if (err.response && err.response.data && err.response.data.error) {
              this.$emit("error", err.response.data.error);
            }
          }
        }
      }
    },

    async init() {
      console.log("spell preview watcher")

      let spellsList = this.spells && this.spells.npc_spells_entries && this.spells.npc_spells_entries.length > 0
        ? JSON.parse(JSON.stringify(this.spells.npc_spells_entries))
        : []

      if (this.spells.parent_list > 0) {
        try {
          const NpcSpellsClient = (new NpcSpellApi(...SpireApi.cfg()))
          const r               = await NpcSpellsClient.getNpcSpell({ id: this.spells.parent_list },
            {
              query:
                (new SpireQueryBuilder())
                  .includes([
                    "NpcSpellsEntries.SpellsNew",
                  ])
                  .limit(100000)
                  .get()
            }
          )

          for (const [i, e] of r.data.npc_spells_entries.entries()) {
            r.data.npc_spells_entries[i].is_parented = true
          }

          if (r.status === 200 && r.data && r.data.npc_spells_entries) {
            spellsList = spellsList.concat(r.data.npc_spells_entries)
          }
        } catch (err) {
          if (err.response.data.error) {
            console.log("Error fetching parent list", err)
          }
        }
      }

      this.spellsList = spellsList
        .sort((a, b) => {
          if (a.minlevel === b.minlevel) {
            return a.spells_new &&
              b.spells_new &&
              b.spells_new.name &&
              a.spells_new.name &&
              a.spells_new.name.localeCompare(b.spells_new.name)
          }

          return (b.minlevel < a.minlevel) ? 1 : -1
        })

      // table sort
      const target = document.getElementById('npc-spell-preview-list-table')
      if (target) {
        setTimeout(() => {
          new Tablesort(target);
        }, 100)
      }

    },
    getRecastDelay(e) {
      if (e.recast_delay === -2) {
        return 0
      }
      if (e.recast_delay === -1) {
        return e.spells_new.recast_time / 1000
      }

      return e.recast_delay
    },
    getManacost(e) {
      if (e.manacost === -1) {
        return e.spells_new.mana
      }

      return e.manacost
    },
    toTitleCase(str) {
      return str.replace(
        /\w\S*/g,
        function (txt) {
          return txt.charAt(0).toUpperCase() + txt.substr(1).toLowerCase();
        }
      );
    }
  }
}
</script>

<style scoped>

</style>
