<template>
  <div>
    <eq-window title="Merchant Editor">
      <div class="row">
        <div class="col-lg-5">
          <input
            type="text"
            class="form-control ml-2"
            placeholder="Search Merchants by Item"
            v-model="search"
            @keyup.enter="doSearch()"
          >
        </div>

        <div class="col-lg-5">
          <select
            name="class"
            id="Class"
            class="form-control"
            v-model="zoneSelection"
            @change="doSearch()"
          >
            <option value="0">-- Select --</option>
            <option v-for="z in zones" v-bind:value="z">
              {{ z.short_name }} ({{ z.version }}) ({{ z.zoneidnumber }}) {{ z.long_name }}
            </option>
          </select>
        </div>

        <div class="col-lg-2 text-center p-0">
          <div class="btn-group" role="group" aria-label="Basic example">
            <b-button title="Search" @click="doSearch()" size="sm" variant="outline-warning">
              <i class="fa fa-search"></i>
            </b-button>
            <b-button title="Reset" @click="reset()" size="sm" variant="outline-danger">
              <i class="fa fa-eraser"></i>
            </b-button>
          </div>
        </div>
      </div>

    </eq-window>

    <eq-window v-if="loading">
      <div class="text-center">
        Loading
        <loader-fake-progress class="mt-3"/>
      </div>
    </eq-window>

    <eq-window v-if="!loading && zoneSelection !== 0 && mlz && mlz.length === 0">
      No merchants found in this zone...
    </eq-window>

    <eq-window
      v-if="!loading && mlz && mlz.length > 0"
      class="p-2 mt-5"
      :title="'Merchants (' + mlz.length + ')'"
    >
      <div style="overflow-y: scroll; max-height: 80vh">
        <table
          id="merchant-list-table"
          class="eq-table eq-highlight-rows"
          style="font-size: 14px; "
        >
          <thead class="eq-table-floating-header">
          <tr>
            <th></th>
            <th class="text-center" style="width: 80px">Merchant ID</th>
            <th>Merchant (NPC)</th>
          </tr>
          </thead>
          <tbody>
          <tr
            v-for="(n, index) in mlz"
            :id="'mlz-' + n.id"
            :key="'ml-' + n.id + '-' + n.merchant_id"
          >
            <td
              class="text-center"
              style="width: 100px"
            >

              <b-button
                class="btn-dark btn-sm btn-dark"
                title="Select Merchant List"
                @click="selectMerchantList(n.merchant_id)"
              >
                <i class="fa fa-arrow-left"></i>
              </b-button>

              <b-button
                class="btn-dark btn-sm btn-dark ml-3"
                @click="editMerchantList(n.merchant_id)"
                title="Edit Merchant List"
              >
                <i class="fa fa-edit"></i>
              </b-button>
            </td>
            <td
              class="text-center"
              style="vertical-align: middle"
            >{{ n.id }}
            </td>
            <td>
              <npc-popover
                :limit-entries="25"
                :no-stats="true"
                :npc="n"
              />
            </td>
          </tr>
          </tbody>
        </table>
      </div>

    </eq-window>
  </div>
</template>

<script>
import EqWindow           from "../eq-ui/EQWindow";
import {Zones}            from "../../app/zones";
import {Npcs}             from "../../app/npcs";
import NpcPopover         from "../NpcPopover";
import LoaderFakeProgress from "../LoaderFakeProgress";

export default {
  name: "MerchantSubEditor",
  components: { LoaderFakeProgress, NpcPopover, EqWindow },
  async created() {
    this.mlz   = [] // by zone
    this.zones = await Zones.getZones()
  },
  data() {
    return {
      // selection
      search: "",
      zoneSelection: 0,

      loading: false,

      // feeds selection
      zones: [],

      // merchant list
      ml: [],
    }
  },
  methods: {

    editMerchantList(merchantId) {
      console.log("[MerchantSubEditor] Editing [%s]", merchantId)

    },
    selectMerchantList(merchantId) {
      console.log("[MerchantSubEditor] Selecting [%s]", merchantId)

      this.$emit('input', merchantId);
    },

    reset() {
      console.log("reset")
    },
    spliceIntoChunks(arr, chunkSize) {
      const res = [];
      while (arr.length > 0) {
        const chunk = arr.splice(0, chunkSize);
        res.push(chunk);
      }
      return res;
    },
    async doSearch() {
      const z = this.zoneSelection
      const r = (await Npcs.getNpcsByZone(
        z.short_name,
        z.version,
        {
          relations: ["Spawnentries.NpcType.Merchantlists.Items"]
        }
      )).filter((e) => {
        let hasItems = false
        if (e.merchantlists) {
          for (const e of e.merchantlists) {
            if (e.items && e.items.length > 0) {
              console.log("has items")
              console.log(e.items)
              hasItems = true;
            }
            // console.log(e)
          }
        }

        return e.merchant_id > 0
      })

      // console.log(r)

      const withItems = r.filter((e) => {
        let hasItems = false
        if (e.merchantlists) {
          for (const e of e.merchantlists) {
            if (e.items && e.items.length > 0) {
              // console.log("has items")
              // console.log(e.items)
              hasItems = true;
            }
            // console.log(e)
          }
        }

        return hasItems
      })

      this.loading = true

      // edge case, if we loaded too much data and failed to load items, load each npc
      let npcs = []
      if (withItems.length === 0) {
        // chunk requests
        for (let chunk of this.spliceIntoChunks(r, 10)) {
          let npcIds = chunk.map((e) => {
            return e.id
          })

          const b = await Npcs.getNpcsBulk(npcIds, ["Merchantlists.Items"])

          npcs = [...npcs, ...b]
        }
      }

      // if our bulk loading didn't work the first time, use the second list
      this.mlz = withItems.length === 0 ? npcs : r
      this.$forceUpdate()

      this.loading = false
    },
  }
}
</script>

<style>

#merchant-list-table td {
  vertical-align: middle;
  padding: 10px;
  height: 60px;
}
</style>
