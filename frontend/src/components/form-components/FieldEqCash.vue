<template>
  <div>
    <div style="display:inline-block; margin-right: 15px">
      <div class="block"><input type="number" v-model="platinum" @change="calc" class="coin">
        <img :src="cdnUrl + 'assets/item_icons/item_644.png'" style="height: 20px" title="Platinum">
      </div>
      <div class="block"><input type="number" v-model="gold" @change="calc" class="coin">
        <img :src="cdnUrl + 'assets/item_icons/item_645.png'" style="height: 20px" title="Gold">
      </div>
      <div class="block"><input type="number" v-model="silver" @change="calc" class="coin">
        <img :src="cdnUrl + 'assets/item_icons/item_646.png'" style="height: 20px" title="Silver">
      </div>
      <div class="block"><input type="number" v-model="copper" @change="calc" class="coin">
        <img :src="cdnUrl + 'assets/item_icons/item_647.png'" style="height: 20px" title="Copper">
      </div>
    </div>
  </div>
</template>

<script>
import {abstractField} from "vue-form-generator";
import {App} from "@/constants/app";

export default {
  mixins: [abstractField],
  data() {
    return {
      cdnUrl: App.ASSET_CDN_BASE_URL,
      platinum: 0,
      gold: 0,
      silver: 0,
      copper: 0
    }
  },
  methods: {
    calc: function () {
      this.value = parseInt(this.platinum + this.gold + this.silver + this.copper)
    },
    init: function () {
      let initialValue = this.value.toString()
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

<style scoped>
.coin {
  height:       40px;
  padding:      15px !important;
  width:        100px;
  margin:       0;
  margin-right: 5px;
  border: none;
  border-radius: 5px;
}

.block {
  display:      inline-block;
  margin-right: 8px
}
</style>
