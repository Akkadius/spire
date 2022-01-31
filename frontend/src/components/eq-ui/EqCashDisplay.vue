<template>
  <div>
    <div v-if="platinum > 0" class="d-inline-block mr-2">
      <div :class="'ml-1 item-644-sm'" title="Silver"/>
      {{ platinum }}
    </div>

    <div v-if="gold > 0" class="d-inline-block mr-2">
      <div :class="'ml-1 item-645-sm'" title="Silver"/>
      {{ gold }}
    </div>

    <div v-if="silver > 0" class="d-inline-block mr-2">
      <div :class="'ml-1 item-646-sm'" title="Silver"/>
      {{ silver }}
    </div>

    <div v-if="copper > 0" class="d-inline-block mr-2">
      <div :class="'ml-1 item-647-sm'" title="Copper"/>
      {{ copper }}
    </div>
  </div>
</template>

<script>
import {App} from "@/constants/app";

export default {
  name: "EqCashDisplay",
  data() {
    return {
      cdnUrl: App.ASSET_CDN_BASE_URL,
      platinum: 0,
      gold: 0,
      silver: 0,
      copper: 0
    }
  },
  props: {
    price: {
      type: Number,
      required: true
    },
    size: {
      type: Number,
      required: false,
      default: 15,
    },
  },
  methods: {
    calc: function () {
      this.value = parseInt(this.platinum + this.gold + this.silver + this.copper)
    },
    init: function () {
      let initialValue = this.price.toString()
      this.copper      = 0;
      this.silver      = 0;
      this.gold        = 0;
      this.platinum    = 0;

      const copper   = initialValue.charAt(initialValue.length - 1)
      const silver   = initialValue.charAt(initialValue.length - 2)
      const gold     = initialValue.charAt(initialValue.length - 3)
      const platinum = initialValue.substr(0, initialValue.length - 3)

      if (copper > 0) {
        this.copper = copper
      }
      if (silver > 0) {
        this.silver = silver
      }
      if (gold > 0) {
        this.gold = gold
      }
      if (platinum > 0) {
        this.platinum = platinum
      }
    }
  },
  watch: {
    value: function () {
      this.init()
    }
  },

  mounted() {
    this.init()
  }
}
</script>
