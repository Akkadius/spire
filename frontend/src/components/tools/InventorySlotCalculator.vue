<template>
  <div class="row">
    <div v-for="(slot, slotId) in slots" class="mb-3 text-center">
      <div class="text-center p-1 col-lg-12 col-sm-12">
        {{ slot.name }}
        <div class="text-center">
          <img
            @click="selectSlot(slotId)"
            :src="slotUrl + 'old_slot_' + slotId + '.gif'"
            :style="'width:auto;' + (isSlotSelected(slotId) ? 'border: 2px solid #dadada; border-radius: 7px;' : 'border: 2px solid rgb(218 218 218 / 0%); border-radius: 7px;')"
            class="mt-1 p-1">
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import {PLAYER_INVENTORY_SLOT, PLAYER_INVENTORY_SLOTS} from "../../app/constants/eq-inventory-constants";
import {App} from "../../constants/app";

export default {
  name: "InventorySlotCalculator",
  data() {
    return {
      slotUrl: App.ASSET_INVENTORY_SLOT_URL,
      slots: PLAYER_INVENTORY_SLOTS,
      selectedSlots: {},
      currentMask: 0
    }
  },
  props: {
    mask: {
      type: String,
      required: false
    }
  },
  mounted() {
    this.currentMask = parseInt(this.mask)
    this.calculateFromBitmask();
  },
  watch: {
    mask: {
      // the callback will be called immediately after the start of the observation
      immediate: true,
      handler (val, oldVal) {
        this.currentMask = parseInt(this.mask)
        this.calculateFromBitmask();
      }
    },
  },
  methods: {
    selectSlot: function(slotId) {
      slotId = parseInt(slotId)

      this.selectedSlots[slotId] = !this.selectedSlots[slotId];

      // dual updates
      const dualUpdates = {}
      dualUpdates[PLAYER_INVENTORY_SLOT.EAR_1] = PLAYER_INVENTORY_SLOT.EAR_2
      dualUpdates[PLAYER_INVENTORY_SLOT.EAR_2] = PLAYER_INVENTORY_SLOT.EAR_1
      dualUpdates[PLAYER_INVENTORY_SLOT.BRACER_1] = PLAYER_INVENTORY_SLOT.BRACER_2
      dualUpdates[PLAYER_INVENTORY_SLOT.BRACER_2] = PLAYER_INVENTORY_SLOT.BRACER_1
      dualUpdates[PLAYER_INVENTORY_SLOT.RING_1] = PLAYER_INVENTORY_SLOT.RING_2
      dualUpdates[PLAYER_INVENTORY_SLOT.RING_2] = PLAYER_INVENTORY_SLOT.RING_1

      if (dualUpdates[slotId]) {
        this.selectedSlots[dualUpdates[slotId]] = this.selectedSlots[slotId];
      }

      this.$forceUpdate()
      this.calculateToBitmask();
    },
    isSlotSelected: function(slotId) {
      return this.selectedSlots[slotId]
    },
    calculateFromBitmask() {
      Object.keys(this.slots).reverse().forEach((slot) => {
        const inventorySlot = this.slots[slot];
        this.selectedSlots[slot] = false
        if (this.currentMask >= inventorySlot.mask) {
          this.currentMask -= inventorySlot.mask;
          this.selectedSlots[slot] = true;
        }
      });
      this.$forceUpdate()
    },
    calculateToBitmask() {
      let bitmask = 0;

      Object.keys(this.slots).reverse().forEach((slot) => {
        const inventorySlot = this.slots[slot];
        if (this.selectedSlots[slot]) {
          bitmask += parseInt(inventorySlot.mask);
        }
      });

      this.$emit("update:inputData", bitmask.toString());
    },
  }
}
</script>

<style>

</style>
