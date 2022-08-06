<template>
  <eq-window title="Alternate Currency Selector" class="p-0">
    <div style="overflow-y: scroll; max-height: 95vh" id="alternate-currency-viewport">

      <div class="mt-3 p-3 text-center" style="padding-bottom: 0px !important;">
        Currencies below are loaded from the `alternate_currency` table
      </div>

      <table
        class="eq-table eq-highlight-rows currency-table"
        style="display: table; font-size: 14px; overflow-x: scroll"
        v-if="currencies"
      >
        <thead
          class="eq-table-floating-header"
        >
        <tr>
          <th style="width: 50px"></th>
          <th style="width: 40px">ID</th>
          <th>Currency</th>
        </tr>
        </thead>
        <tbody>
        <tr
          :id="'currency-' + e.id"
          :class="(isCurrencySelected(e) ? 'pulsate-highlight-white' : '')"
          v-for="(e, index) in currencies"
          :key="e.id"
          style="height: 50px"
        >
          <td>
            <b-button
              class="btn-dark btn-sm btn-outline-warning"
              title="Select"
              @click="selectAlternateCurrency(e)"
            >
              <i class="fa fa-arrow-left"></i>
            </b-button>
          </td>

          <td>
            {{ e.id }}
          </td>

          <td class="text-left" style="vertical-align: middle">
            <item-popover
              :item="e.item"
              v-if="e.item"
              size="regular"
            />
          </td>

        </tr>
        </tbody>
      </table>
    </div>
  </eq-window>
</template>

<script>
import EqWindow               from "@/components/eq-ui/EQWindow";
import {AlternateCurrencyApi} from "@/app/api";
import {SpireApiClient}       from "@/app/api/spire-api-client";
import ItemPopover            from "@/components/ItemPopover";
import {SpireQueryBuilder} from "@/app/api/spire-query-builder";
import {scrollToTarget}    from "@/app/utility/scrollToTarget";

export default {
  name: "AlternateCurrencySelector",
  components: { ItemPopover, EqWindow },
  props: {
    selectedCurrency: {
      type: Number,
      required: false
    }
  },
  data() {
    return {
      currencies: []
    }
  },
  async mounted() {
    const api = (new AlternateCurrencyApi(SpireApiClient.getOpenApiConfig()))
    const r   = await api.listAlternateCurrencies(
      (new SpireQueryBuilder())
        .includes(['Item'])
        .get()
    )
    if (r.status === 200) {
      this.currencies = r.data
    }

    if (this.selectedCurrency > 0) {
      scrollToTarget(
        "alternate-currency-viewport",
        'currency-' + this.selectedCurrency
      )
    }
  },
  methods: {
    selectAlternateCurrency(currency) {
      this.$emit('input', currency.id);
    },
    isCurrencySelected(e) {
      return this.selectedCurrency &&
        this.selectedCurrency > 0 &&
        e.id === this.selectedCurrency;
    }
  }
}
</script>

<style>
.currency-table td {
  vertical-align: middle !important;
}
</style>
