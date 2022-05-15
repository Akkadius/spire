<template>
  <div v-if="npcRowData">
    <pre
      v-if="npcStatusMessage" style="width: 100%; margin: 0px"
      class="mb-2"
    >{{ npcStatusMessage }}
    </pre>
    <ag-grid-vue
      v-if="npcRowData"
      style="width: 100%; height: 86vh;"
      class="ag-theme-balham-dark"
      :columnDefs="npcColumnDefs"
      :gridOptions="gridOptions"
      :rowData="npcRowData"
    >
    </ag-grid-vue>
  </div>
</template>

<script>
// import {AgGridVue}      from "ag-grid-vue";
import {NpcTypeApi}     from "@/app/api";
import {SpireApiClient} from "@/app/api/spire-api-client";
import util             from "util";

export default {
  name: "NpcListGrid",
  props: {
    npcs: {
      type: Array
    }
  },
  data() {
    return {

      npcStatusMessage: "",
      npcColumnDefs: null,
      npcRowData: null,

      gridOptions: {
        defaultColDef: {
          resizable: true,
          sortable: true,
          searchable: true,
          editable: true,
          filter: "agTextColumnFilter",
          floatingFilter: true
        },

        enableCellChangeFlash: true,

        // pagination: true,
        rowSelection: "single",

        rowGroupPanelShow: "always",
        pivotPanelShow: "always",

        enablePivot: false,

        onCellValueChanged: (params) => {
          const row           = params.data
          const columnChanged = params.column.colId
          const newValue      = params.newValue
          const npcApi        = (new NpcTypeApi(SpireApiClient.getOpenApiConfig()))

          // make sure we keep the integer types the same before we send to the
          // backend since it is strict
          if (parseInt(newValue) > 0 || parseInt(newValue) < 0) {
            row[columnChanged] = Number(newValue)
          }

          // updated object
          let update            = {}
          update.id             = row.id
          update[columnChanged] = row[columnChanged]

          npcApi.updateNpcType({
            id: row.id,
            npcType: update
          }).then((result) => {
            if (result.status === 200) {
              this.npcStatusMessage = util.format(
                "Updated NPC [%s] [%s] Column [%s] to Value [%s]",
                row.id,
                row.name,
                columnChanged,
                newValue
              )
            }
          }).catch((err) => {
            this.npcStatusMessage = err.response.data.error
          })


          // console.log(params)

          // trigger filtering on cell edits
          params.api.onFilterChanged();
        },

        sideBar: true
      }
    }
  },
  created() {
    this.npcRowData       = null
    this.npcStatusMessage = "";

    let npcColumnDefs = [
      {
        headerName: "NPC",
        children: [
          { width: 100, headerName: "id", field: "id", pinned: true },
          { width: 200, headerName: "Name", field: "name", pinned: true },
          { headerName: "Lastname", field: "lastname" }
        ]
      },
      {
        headerName: "Meta",
        children: [
          { width: 75, field: "level" },
          {
            width: 130,
            field: "class"
          }
        ]
      },
      {
        headerName: "Stats",
        children: [
          { width: 100, headerName: "hp", field: "hp" },
          { width: 100, headerName: "mana", field: "mana" },
          { width: 75, headerName: "ac", field: "ac" },
          { width: 130, headerName: "hp_regen_rate", field: "hp_regen_rate" },
          { width: 150, headerName: "mana_regen_rate", field: "mana_regen_rate" },
          { width: 120, headerName: "mindmg", field: "mindmg" },
          { width: 120, headerName: "maxdmg", field: "maxdmg" },
          { width: 130, headerName: "attack_count", field: "attack_count" },
          { width: 75, headerName: "str", field: "str" },
          { width: 75, headerName: "sta", field: "sta" },
          { width: 75, headerName: "dex", field: "dex" },
          { width: 75, headerName: "agi", field: "agi" },
          { width: 75, headerName: "_int", field: "_int" },
          { width: 75, headerName: "wis", field: "wis" },
          { width: 75, headerName: "cha", field: "cha" },
          { width: 75, headerName: "atk", field: "atk" }
        ]
      },
      {
        headerName: "Relational Data",
        children: [
          { field: "loottable_id", headerName: "loottable_id" },
          { field: "merchant_id", headerName: "merchant_id" },
          { field: "alt_currency_id", headerName: "alt_currency_id" },
          { field: "npc_spells_id", headerName: "npc_spells_id" },
          { field: "npc_spells_effects_id", headerName: "npc_spells_effects_id" },
          { field: "npc_faction_id", headerName: "npc_faction_id" },
          { field: "adventure_template_id", headerName: "adventure_template_id" },
          { field: "trap_template", headerName: "trap_template" },
          { field: "emoteid", headerName: "emoteid" }
        ]
      },
      {
        headerName: "Resists",
        children: [
          { width: 100, headerName: "Magic", field: "mr" },
          { width: 100, headerName: "Cold", field: "cr" },
          { width: 100, headerName: "Disease", field: "dr" },
          { width: 100, headerName: "Fire", field: "fr" },
          { width: 100, headerName: "Poison", field: "pr" },
          { width: 120, headerName: "Corruption", field: "corrup" },
          { width: 100, headerName: "Physical", field: "ph_r" }
        ]
      },
      {
        headerName: "Aggro / Assist",
        children: [
          { field: "aggroradius" },
          { field: "assistradius" },
          { field: "npc_aggro" }
        ]
      },
      {
        headerName: "Appearance",
        children: [
          { headerName: "race", field: "race" },
          { headerName: "gender", field: "gender" },
          { headerName: "texture", field: "texture" },
          { headerName: "bodytype", field: "bodytype" },
          { headerName: "helmtexture", field: "helmtexture" },
          { headerName: "herosforgemodel", field: "herosforgemodel" },
          { headerName: "size", field: "size" }
        ]
      },
      {
        headerName: "Armor Appearance",
        children: [
          { headerName: "armortint_id", field: "armortint_id" },
          { headerName: "armortint_red", field: "armortint_red" },
          { headerName: "armortint_green", field: "armortint_green" },
          { headerName: "armortint_blue", field: "armortint_blue" },
          { headerName: "d_melee_texture_1", field: "d_melee_texture_1" },
          { headerName: "d_melee_texture_2", field: "d_melee_texture_2" },
          { headerName: "armtexture", field: "armtexture" },
          { headerName: "bracertexture", field: "bracertexture" },
          { headerName: "handtexture", field: "handtexture" },
          { headerName: "legtexture", field: "legtexture" },
          { headerName: "feettexture", field: "feettexture" }
        ]
      },
      {
        headerName: "Facial / Tattoo",
        children: [
          { field: "face" },
          { field: "luclin_hairstyle" },
          { field: "luclin_haircolor" },
          { field: "luclin_eyecolor" },
          { field: "luclin_eyecolor_2" },
          { field: "luclin_beardcolor" },
          { field: "luclin_beard" },
          { field: "drakkin_heritage" },
          { field: "drakkin_tattoo" },
          { field: "drakkin_details" }
        ]
      },
      {
        headerName: "Special Abilities",
        children: [
          { width: 200, headerName: "Special Abilities", field: "special_abilities" },
          { headerName: "Special Attacks (Deprecated)", field: "npcspecialattks", hide: true }
        ]
      },
      {
        headerName: "Charm",
        children: [
          { headerName: "charm_ac", field: "charm_ac" },
          { headerName: "charm_min_dmg", field: "charm_min_dmg" },
          { headerName: "charm_max_dmg", field: "charm_max_dmg" },
          { headerName: "charm_attack_delay", field: "charm_attack_delay" },
          { headerName: "charm_accuracy_rating", field: "charm_accuracy_rating" },
          { headerName: "charm_avoidance_rating", field: "charm_avoidance_rating" },
          { headerName: "charm_atk", field: "charm_atk" }
        ]
      },
      {
        headerName: "Attributes",
        children: [
          { width: 120, field: "untargetable", headerName: "untargetable" },
          { width: 120, field: "findable", headerName: "findable" },
          { width: 120, field: "see_hide", headerName: "see_hide" },
          { width: 120, field: "see_improved_hide", headerName: "see_improved_hide" },
          { width: 120, field: "trackable", headerName: "trackable" },
          { width: 120, field: "underwater", headerName: "underwater" },
          { width: 120, field: "isquest", headerName: "isquest" },
          { width: 120, field: "light", headerName: "light" },
          { width: 120, field: "fixed", headerName: "fixed" },
          { width: 120, field: "ignore_despawn", headerName: "ignore_despawn" },
          { width: 120, field: "show_name", headerName: "show_name" },
          { width: 120, field: "untargetable", headerName: "untargetable" },
          { width: 120, field: "rare_spawn", headerName: "rare_spawn" },
          { width: 120, field: "flymode", headerName: "flymode" },
          { width: 120, field: "always_aggro", headerName: "always_aggro" },
          { width: 120, field: "see_invis", headerName: "see_invis" },
          { width: 120, field: "see_invis_undead", headerName: "see-invis_undead" }
        ]
      }
    ]

    if (this.npcs[0]) {
      Object.keys(this.npcs[0]).forEach((key) => {
        if (typeof this.npcs[0][key] !== "object") {

          let hasField = false
          npcColumnDefs.forEach((def) => {
            if (def.children) {
              def.children.forEach((child) => {
                if (key === child.field) {
                  hasField = true
                }
              })
            }
          })

          if (!hasField) {
            npcColumnDefs.push({ headerName: key, field: key })
          }
        }
      })

      this.npcColumnDefs = npcColumnDefs
      this.npcRowData    = this.npcs
    }
  },
  components: {
    AgGridVue
  }
}
</script>

<style scoped>

</style>
