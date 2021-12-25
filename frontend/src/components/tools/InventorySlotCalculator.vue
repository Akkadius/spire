<template>
  <div class="row" v-if="mask >= 0">
    <div
      class="mr-3 d-inline-block text-center"
      :style="'display: inline-block ' + (centeredButtons ? 'width: 100%' : '')"
    >
      <div
        v-for="(slot, slotId) in slots"
        class="mb-1 text-center d-inline-block"
      >
        <div
          class="text-center p-0 mr-1 col-lg-12 col-sm-12"
          v-if="!isSlotSkipped(slotId)"
        >
          <span
            v-if="showTextTop"
            class="d-inline-block"
            :style="'font-size: 12px; white-space: nowrap; overflow: hidden;text-overflow: ellipsis; max-width: ' + (imageSize) + 'px;'">
            {{ slot.name }}
          </span>

          <div class="text-center">
            <img
              :title="slot.name"
              @click="selectSlot(slotId)"
              :src="slotUrl + 'old_slot_' + slotId + '.gif'"
              class="hover-highlight"
              :style="getImageSize() + (isSlotSelected(slotId) ? 'border: 2px solid #dadada; border-radius: 7px;' : 'border-radius: 7px; ')"
            >
          </div>
        </div>
      </div>

      <!-- Select All / None -->
      <div class="d-inline-block" v-if="displayAllNone">
        <div
          :class="'text-center mt-4 btn-xs eq-button-fancy ' + (parseInt(mask) >= 8388607 ? 'eq-button-fancy-highlighted' : '')"
          @click="selectAll()"
        >
          All
        </div>
        <div
          :class="'text-center mt-4 btn-xs eq-button-fancy ' + (parseInt(mask) === 0 ? 'eq-button-fancy-highlighted' : '')"
          @click="selectNone()"
        >
          None
        </div>
      </div>

    </div>

  </div>
</template>

<script>
import {PLAYER_INVENTORY_SLOT, PLAYER_INVENTORY_SLOTS} from "@/app/constants/eq-inventory-constants";
import {App}                                           from "@/constants/app";
import util                                            from "util";

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
      type: Number,
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
    },
    showTextTop: {
      type: Boolean,
      required: false,
      default: true
    },
    centeredButtons: {
      type: Boolean,
      required: false,
      default: true
    },
    imageSize: {
      type: Number,
      required: false,
      default: 50,
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
    getImageSize() {
      return util.format("width: %spx; height %spx;", this.imageSize, this.imageSize)
    },

    isSlotSkipped(slot) {
      if (!this.skipDuplicateSlots) {
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

      this.$emit("update:inputData", parseInt(bitmask));
      this.$emit("input", parseInt(bitmask));
      this.$emit("fired", "true");
    }
  }
}
</script>

<style>

</style>
