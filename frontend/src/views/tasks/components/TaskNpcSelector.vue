<template>
  <div>
    <eq-window-simple title="NPC Selector">
      <b-input
        v-model="npcSearch"
        class="form-control"
        id="search-selector"
        v-on:keyup.enter="searchNpc"
        placeholder="Search by npc name, id..."
      />
    </eq-window-simple>

    <eq-window-simple
      id="npc-view-container"
      v-if="filteredNpcs && filteredNpcs.length > 0"
      style="height: 85vh; overflow-y: scroll;" class="p-0"
    >
      <table
        id="npctable"
        class="eq-table eq-highlight-rows"
        style="display: table; font-size: 14px; overflow-x: scroll"
      >
        <thead
          class="eq-table-floating-header"
        >
        <tr>
          <th style="width: 30px"></th>
          <th style="width: 30px">ID</th>
          <th>Name</th>
          <th>Zones</th>
          <th></th>
        </tr>
        </thead>
        <tbody>
        <tr
          :id="'npc-' + npc.short_name"
          :class="(isNpcSelected(npc) ? 'pulsate-highlight-white' : '')"
          v-for="(npc, index) in filteredNpcs"
          :key="npc.id"
        >
          <td>
            <b-button
              class="btn-dark btn-sm btn-outline-warning"
              @click="selectNpc(npc)"
            >
              <i class="fa fa-arrow-left"></i>
            </b-button>
          </td>
          <td style="text-align: center" class="p-0">{{ npc.id }}</td>
          <td>{{ npc.name }} <span v-if="npc.lastname && npc.lastname.length > 0">({{ npc.lastname }})</span></td>
          <td><span v-if="npcZones[npc.id]">{{ npcZones[npc.id].join(",") }}</span></td>
          <td class="text-center"><span :class="'race-models-ctn-' + npc.race + '-' + npc.gender + '-' + npc.texture + '-' + npc.helmtexture + ''" style="zoom: 75  %;"></span></td>
        </tr>
        </tbody>
      </table>
    </eq-window-simple>
  </div>
</template>

<script>
import EqWindowSimple      from "@/components/eq-ui/EQWindowSimple";
import {NpcTypeApi}        from "@/app/api";
import {SpireApiClient}    from "@/app/api/spire-api-client";
import util                from "util";
import Expansions          from "@/app/utility/expansions";
import EqCheckbox          from "@/components/eq-ui/EQCheckbox";
import {SpireQueryBuilder} from "@/app/api/spire-query-builder";
import {Zones}             from "@/app/zones";

export default {
  name: "TaskNpcSelector",
  components: { EqCheckbox, EqWindowSimple },
  data() {
    return {
      // filtered content
      filteredNpcs: {},

      // search
      npcSearch: "",

      // model we work with after the prop is passed we can manipulate it ourselves props should not be mutated
      selectedNpcIdAttr: 0,

      // key by npcid
      npcZones: {},
    }
  },
  props: {
    selectedNpcId: {
      type: Number,
      required: true,
    },
  },
  methods: {
    isNpcSelected(npc) {
      return parseInt(npc.id) === parseInt(this.selectedNpcIdAttr)
    },

    selectNpc(npc) {
      this.$emit('input', {
        npcId: npc.id,
        npc: npc,
      });

      this.selectedNpcIdAttr = npc.id
    },

    async searchNpc() {
      if (this.npcSearch === "") {
        return ""
      }

      const api   = (new NpcTypeApi(SpireApiClient.getOpenApiConfig()))
      let builder = (new SpireQueryBuilder())
      if (this.npcSearch !== "") {
        if (parseInt(this.npcSearch) > 0) {
          builder.whereOr("id", "=", this.npcSearch)
        }
        else {
          builder.where("name", "like", this.npcSearch)
          builder.whereOr("lastname", "like", this.npcSearch)
        }
      }

      builder.includes([
        "Spawnentries",
        "Spawnentries.Spawngroup",
        "Spawnentries.Spawngroup.Spawn2",
        "Spawnentries.Spawngroup.Spawn2.Spawnentries",
        "Spawnentries.Spawngroup.Spawn2.Spawngroup",
      ])

      const result = await api.listNpcTypes(
        builder.orderBy(["id", "name"]).get()
      )

      if (result.status === 200) {
        this.filteredNpcs = result.data

        // get zones where npc's reside in
        let npcZones = {}
        for (const npc of result.data) {
          if (
            npc.spawnentries &&
            npc.spawnentries[0].spawngroup &&
            npc.spawnentries[0].spawngroup.spawn_2
          ) {
            const zoneName = await Zones.getZoneLongNameByShortName(npc.spawnentries[0].spawngroup.spawn_2.zone.toLowerCase())
            if (typeof this.npcZones[npc.id] === "undefined") {
              npcZones[npc.id] = []
            }

            npcZones[npc.id].push(`${zoneName} (${npc.spawnentries[0].spawngroup.spawn_2.version})`)
          }
        }

        // if (this.npcSearch !== "") {
        //   let filtered = []
        //   for (let index in this.filteredNpcs) {
        //     const r = this.filteredNpcs[index]
        //     if (npcZones[npc.id].join("").toLowerCase().includes(this.npcSearch.toLowerCase())) {
        //       filtered.
        //     }
        //   }
        // }

        // update npcZones
        this.npcZones = npcZones
      }
    },

    getExpansionIcon(expansion) {
      return Expansions.getExpansionIconUrlSmall(expansion - 1) // npc table is offset by 1
    },
    getExpansionName(expansion) {
      return Expansions.getExpansionName(expansion - 1) // npc table is offset by 1
    },
  },
  mounted() {
    const t = document.getElementById("search-selector")
    if (t) {
      t.focus()
    }

    this.selectedNpcIdAttr = this.selectedNpcId

    if (this.selectedNpcIdAttr !== 0) {
      this.npcSearch = this.selectedNpcIdAttr
    }

    this.searchNpc()

    setTimeout(() => {
      const container = document.getElementById("npc-view-container");
      const target    = document.getElementById(util.format("npc-%s", this.selectedNpcIdAttr))
      if (container && target) {
        const top           = target.getBoundingClientRect().top
        container.scrollTop = container.scrollTop + top - 300;
      }
    }, 100)

    // pre-select text box input
    setTimeout(() => {
      const input = document.getElementById('search-selector');
      if (input) {
        input.focus();
        input.select();
      }
    }, 100)
  }
}
</script>

<style scoped>
#npctable td {
  vertical-align: middle !important;
}
</style>
