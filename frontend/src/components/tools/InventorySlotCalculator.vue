<template>
  <div class="row">
    <div class="mr-3 d-inline-block" style="display: inline-block">
      <div v-for="(slot, slotId) in slots" class="mb-1 text-center d-inline-block">
        <div class="text-center p-1 col-lg-12 col-sm-12" v-if="!isSlotSkipped(slotId)">
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
    <div class="mt-4 d-inline-block" v-if="displayAllNone">
      <button class='eq-button mr-3' @click="selectAll()" style="display: inline-block; width: 80px">All</button>
      <button class='eq-button' @click="selectNone()" style="display: inline-block; width: 80px">None</button>
    </div>
  </div>
</template>

<script>
import {PLAYER_INVENTORY_SLOT, PLAYER_INVENTORY_SLOTS} from "@/app/constants/eq-inventory-constants";
import {App} from "@/constants/app";

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
    },
    displayAllNone: {
      type: Boolean,
      required: false,
      default: true
    },
    skipDuplicateSlots: {
      type: Boolean,
      required: false,
      default: false
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
      handler(val, oldVal) {
        this.currentMask = parseInt(this.mask)
        this.calculateFromBitmask();
      }
    }
  },
  methods: {
    isSlotSkipped(slot) {
      if (!this.skipDuplicateSlots){
        return false;
      }

      if (this.skipDuplicateSlots) {
        // this comes in as a string
        slot = parseInt(slot)

        if (slot === PLAYER_INVENTORY_SLOT.EAR_2) {
          return true;
        }
        if (slot === PLAYER_INVENTORY_SLOT.BRACER_2) {
          return true;
        }
        if (slot === PLAYER_INVENTORY_SLOT.RING_2) {
          return true;
        }
      }
    },
    selectAll() {
      Object.keys(this.slots).reverse().forEach((slotId) => {
        this.selectedSlots[slotId] = true;
      });
      this.$forceUpdate();
      this.calculateToBitmask();
    },
    selectNone() {
      Object.keys(this.slots).reverse().forEach((slotId) => {
        this.selectedSlots[slotId] = false;
      });
      this.$forceUpdate();
      this.calculateToBitmask();
    },
    selectSlot: function (slotId) {
      slotId = parseInt(slotId)

      this.selectedSlots[slotId] = !this.selectedSlots[slotId];

      // dual updates
      const dualUpdates                           = {}
      dualUpdates[PLAYER_INVENTORY_SLOT.EAR_1]    = PLAYER_INVENTORY_SLOT.EAR_2
      dualUpdates[PLAYER_INVENTORY_SLOT.EAR_2]    = PLAYER_INVENTORY_SLOT.EAR_1
      dualUpdates[PLAYER_INVENTORY_SLOT.BRACER_1] = PLAYER_INVENTORY_SLOT.BRACER_2
      dualUpdates[PLAYER_INVENTORY_SLOT.BRACER_2] = PLAYER_INVENTORY_SLOT.BRACER_1
      dualUpdates[PLAYER_INVENTORY_SLOT.RING_1]   = PLAYER_INVENTORY_SLOT.RING_2
      dualUpdates[PLAYER_INVENTORY_SLOT.RING_2]   = PLAYER_INVENTORY_SLOT.RING_1

      if (dualUpdates[slotId]) {
        this.selectedSlots[dualUpdates[slotId]] = this.selectedSlots[slotId];
      }

      this.$forceUpdate()
      this.calculateToBitmask();
    },
    isSlotSelected: function (slotId) {
      return this.selectedSlots[slotId]
    },
    calculateFromBitmask() {
      Object.keys(this.slots).reverse().forEach((slot) => {
        const inventorySlot      = this.slots[slot];
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
      this.$emit("fired", "true");
    }
  }
}
</script>

<style>

</style>
