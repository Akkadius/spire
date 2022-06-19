<template>
  <content-area>
    <div class="row">
      <div :class="(Object.keys(activeMerchant).length > 0 ? 'col-7' : 'col-12')">
        <eq-window title="Merchant Editor">
          <div class="row">
            <div class="col-lg-5">
              <input
                type="text"
                class="form-control ml-2"
                placeholder="Search Merchants by Name"
                v-model="search"
                @keyup.enter="zoneSelection = 0; updateQueryState()"
              >
            </div>

            <div class="col-lg-5">
              <select
                name="class"
                id="Class"
                class="form-control"
                v-model="zoneSelection"
                @change="search = ''; updateQueryState()"
              >
                <option value="0">-- Select --</option>
                <option v-for="z in zones" v-bind:value="{z: z.short_name, v: z.version}">
                  {{ z.short_name }} ({{ z.version }}) ({{ z.zoneidnumber }}) {{ z.long_name }}
                </option>
              </select>
            </div>

            <div class="col-lg-2 text-center p-0 mt-1">
              <div class="btn-group" role="group" aria-label="Basic example">
                <b-button title="Search" @click="updateQueryState()" size="sm" variant="outline-warning">
                  <i class="fa fa-search"></i> Search
                </b-button>
                <b-button title="Reset" @click="reset()" size="sm" variant="outline-danger">
                  <i class="fa fa-eraser"></i> Reset
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
          No merchants found...
        </eq-window>

        <eq-window
          v-if="!loading && mlz && mlz.length > 0"
          class="p-2 mt-5"
          :title="'Merchants (' + mlz.length + ')'"
        >
          <div style="overflow-y: scroll; max-height: 83vh">
            <table
              id="merchant-list-table"
              class="eq-table bordered eq-highlight-rows"
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
                :key="index"
                @mouseover="activeMerchant = n"
              >
                <td
                  class="text-center"
                  style="width: 100px"
                >

                  <b-button
                    class="btn-dark btn-sm btn-outline-warning"
                    title="Select Merchant List"
                    @click="selectMerchantList(n.merchant_id);"
                  >
                    <i class="fa fa-arrow-left"></i>
                  </b-button>

                  <b-button
                    class="btn-dark btn-sm btn-outline-warning ml-3"
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
                <td class="text-left">
                  <npc-popover
                    :popover-enabled="false"
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

      <div :class="(Object.keys(activeMerchant).length > 0 ? 'col-5' : 'col-12')">
        <eq-window
          v-if="Object.keys(activeMerchant).length > 0"
          style="max-height: 95vh; overflow-y: scroll; overflow-x: hidden"
        >
          <eq-npc-card-preview
            :npc="activeMerchant"
            :no-stats="true"
            :limit-entries="1000"
          />
        </eq-window>
      </div>
    </div>

  </content-area>
</template>

<script>
import {Zones}            from "../../app/zones";
import ContentArea        from "../../components/layout/ContentArea";
import EqWindow           from "../../components/eq-ui/EQWindow";
import LoaderFakeProgress from "../../components/LoaderFakeProgress";
import NpcPopover         from "../../components/NpcPopover";
import {ROUTE}            from "../../routes";
import EqNpcCardPreview   from "../../components/preview/EQNpcCardPreview";
import {Merchants}        from "../../app/merchants";

const MILLISECONDS_BEFORE_WINDOW_RESET = 5000;

export default {
  name: "MerchantSubEditor",
  components: { EqNpcCardPreview, NpcPopover, LoaderFakeProgress, EqWindow, ContentArea },
  async created() {
    this.loadQueryState()
    this.mlz   = [] // by zone
    this.zones = await Zones.getZones()

    this.init()
  },
  data() {
    return {
      // selection
      search: "",
      zoneSelection: 0,

      loading: false,

      // selectors
      selectorActive: {},
      lastResetTime: Date.now(),

      // preview
      activeMerchant: {},

      // feeds selection
      zones: [],

      // merchant list
      ml: [],
    }
  },
  watch: {
    '$route'() {
      this.loadQueryState()
      this.init()
    },
  },
  methods: {

    /**
     * State
     */
    updateQueryState: function () {
      let queryState = {};

      if (this.zoneSelection !== 0) {
        queryState.zone = JSON.stringify(this.zoneSelection)
      }
      if (this.search !== '') {
        queryState.q = JSON.stringify(this.search)
      }

      this.$router.push(
        {
          path: ROUTE.MERCHANTS,
          query: queryState
        }
      ).catch(() => {
      })
    },

    loadQueryState() {
      if (typeof this.$route.query.zone !== 'undefined' && this.$route.query.zone !== 0) {
        this.zoneSelection = JSON.parse(this.$route.query.zone);
      }
      if (typeof this.$route.query.q !== 'undefined' && this.$route.query.q !== '') {
        this.search = JSON.parse(this.$route.query.q);
      }
    },

    /**
     * Init
     */
    async init() {
      this.activeMerchant = {}
      this.mlz            = []
      this.loading        = true
      const z             = this.zoneSelection

      if (Object.keys(this.zoneSelection).length > 0) {
        this.mlz = (await Merchants.getMerchantsByZone(z.z, z.v))
      }
      if (this.search.length > 0) {
        this.mlz = (await Merchants.getMerchantsByName(this.search))
      }

      this.$forceUpdate()

      this.loading = false
    },

    /**
     * Controls
     */
    editMerchantList(merchantId) {
      console.log("[MerchantSubEditor] Editing [%s]", merchantId)
    },
    selectMerchantList(merchantId) {
      console.log("[MerchantSubEditor] Selecting [%s]", merchantId)

      this.$emit('input', merchantId);
    },

    reset() {
      this.search         = ""
      this.zoneSelection  = 0
      this.activeMerchant = {}
      this.mlz            = []
      this.updateQueryState()
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
