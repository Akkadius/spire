<template>
  <div>
    <eq-window-simple title="Aura Selector">
      <b-input
        v-model="auraSearch"
        class="form-control"
        v-on:keyup="searchAura"
        placeholder="Search by aura name..."
      />
    </eq-window-simple>

    <eq-window-simple
      id="aura-view-container"
      style="height: 85vh; overflow-y: scroll;" class="p-0"
    >
      <table
        id="auratable"
        class="eq-table eq-highlight-rows"
        style="display: table; font-size: 14px; overflow-x: scroll;"
      >
        <thead
          class="eq-table-floating-header"
        >
        <tr>
          <th style="width: 30px"></th>
          <th>Name</th>
          <th>NPC ID</th>
          <th>Spell ID</th>
          <th>Distance</th>
        </tr>
        </thead>
        <tbody>
        <tr
          :id="'aura-' + stripNonAlpha(aura.name)"
          :class="(isAuraSelected(aura) ? 'pulsate-highlight-white' : '')"
          v-for="(aura, index) in filteredAuras"
          :key="aura.id"
        >
          <td>
            <b-button
              class="btn-dark btn-sm btn-outline-warning"
              @click="selectAura(aura)"
            >
              <i class="fa fa-arrow-left"></i>
            </b-button>
          </td>
          <td>{{ aura.name }}</td>
          <td>{{ aura.npc_type }}</td>
          <td>{{ aura.spell_id }}</td>
          <td>{{ aura.distance }}</td>
        </tr>
        </tbody>
      </table>
    </eq-window-simple>
  </div>
</template>

<script>
import {TELEPORT_ZONE_SELECTOR_TYPE} from "@/app/constants/eq-spell-constants";
import EqWindowSimple                from "@/components/eq-ui/EQWindowSimple";
import {AuraApi}                     from "@/app/api";
import {SpireApiClient}              from "@/app/api/spire-api-client";
import util                          from "util";
import Expansions                    from "@/app/utility/expansions";
import EqCheckbox                    from "@/components/eq-ui/EQCheckbox";

let auras = {}

export default {
  name: "SpellTeleportZoneSelectorAura",
  components: { EqCheckbox, EqWindowSimple },
  data() {
    return {
      TELEPORT_ZONE_SELECTOR_TYPE: TELEPORT_ZONE_SELECTOR_TYPE,

      // filtered content
      filteredAuras: {},

      // search
      auraSearch: "",

      // model we work with after the prop is passed so we can manipulate it ourselves
      // props should not be mutated
      selectedAura: "",
    }
  },
  props: {
    selectedAuraName: {
      type: String,
      required: true,
    },
  },
  methods: {

    stripNonAlpha(string) {
      return string.replace(/[\W_]+/g, " ");
    },

    isAuraSelected(aura) {
      return aura.name.trim() === this.selectedAura
    },

    selectAura(aura) {
      this.$emit('input', {
        aura: aura,
      });

      this.selectedAura = aura.name
    },

    searchAura() {
      const searchString = this.auraSearch.toLowerCase().trim()
      let filteredAuras  = []
      auras.forEach((aura) => {
        if (this.auraSearch.trim() !== '' && aura.name.toLowerCase().includes(searchString)) {
          filteredAuras.push(aura)
        }
      });
      this.filteredAuras = filteredAuras

      if (filteredAuras.length === 0) {
        this.filteredAuras = auras;
      }
    },

    getExpansionIcon(expansion) {
      return Expansions.getExpansionIconUrlSmall(expansion - 1) // aura table is offset by 1
    },
    getExpansionName(expansion) {
      return Expansions.getExpansionName(expansion - 1) // aura table is offset by 1
    },

    async loadAuras() {
      const api    = (new AuraApi(SpireApiClient.getOpenApiConfig()))
      const result = await api.listAuras({
        groupBy: "name",
      })

      if (result.status === 200) {
        auras              = result.data
        this.filteredAuras = auras
      }
    },

    init() {
      this.loadAuras()
    }
  },
  mounted() {
    // model we work with after the prop is passed - we can manipulate it ourselves
    this.selectedAura = this.selectedAuraName
    this.init()

    setTimeout(() => {
      const container = document.getElementById("aura-view-container");
      const target    = document.getElementById(util.format("aura-%s", this.selectedAura))
      if (container && target) {
        const top           = target.getBoundingClientRect().top
        container.scrollTop = container.scrollTop + top - 300;
      }
    }, 1000)
  }
}
</script>

<style scoped>
#auratable td {
  vertical-align: middle !important;
}
</style>
