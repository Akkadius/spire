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
    >
      <thead class="eq-table-floating-header">
      <tr>
        <th>Spell</th>
        <th>Type</th>
        <th>Mana</th>
        <th>Priority</th>
        <th>Recast</th>
        <th>Min / Max Lvl</th>
        <th>Min / Max HP</th>
      </tr>
      </thead>
      <tbody>
      <tr
        v-for="e in sortEntries(spells.npc_spells_entries)"
        :id="'spell-' + e.id"
        :key="'spell-' + e.id"
      >
        <td>
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
        <td>{{ e.minlevel }} / {{ e.maxlevel }}</td>
        <td>{{ e.min_hp }} / {{ e.max_hp }}</td>
      </tr>
      </tbody>
    </table>

  </div>
</template>

<script>
import SpellPopover      from "../SpellPopover";
import EqDebug           from "../eq-ui/EQDebug";
import {NPC_SPELL_TYPES} from "../../app/constants/eq-npc-spells";

export default {
  name: "NpcSpellPreview",
  components: { EqDebug, SpellPopover },
  data() {
    return {
      NPC_SPELL_TYPES: NPC_SPELL_TYPES,

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
  props: {
    spells: {
      type: Object,
      required: false
    },
  },
  methods: {
    sortEntries(e) {
      return e.sort((a, b) => (a.priority < b.priority) ? 1 : -1)
    },
    getRecastDelay(e) {
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
