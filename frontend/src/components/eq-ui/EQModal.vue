<template>
    <div class="eq-modal-mask">
      <div class="eq-modal-wrapper" @click.self="dismiss()">
        <eq-window
          :title="title"
          class="eq-modal-container" >

          <div class="eq-modal-header">
            <slot name="header">

            </slot>
          </div>

          <div class="eq-modal-body">
            <slot name="body">

            </slot>
          </div>

          <div class="eq-modal-footer">
            <slot name="footer">
              <button class="eq-modal-default-button" @click="$emit('close')">
                OK
              </button>
            </slot>
          </div>
        </eq-window>
      </div>
    </div>
</template>

<script>
import EqWindow from "@/components/eq-ui/EQWindow.vue";

export default {
  name: 'EqModal',
  components: { EqWindow },
  props: {
    title: {
      type: String,
      default: '',
      required: false
    },
  },
  methods: {
    dismiss() {
      this.$emit('close');
    },
    keyPress(e) {
      if (e.key === 'Escape') {
        this.dismiss();
      }
    }
  },
  beforeDestroy() {
    // remove keypress listener
    document.removeEventListener('keydown', this.keyPress);
  },
  mounted() {
    // listen keypress for escape key
    document.addEventListener('keydown', this.keyPress);
  },
}
</script>

<style>
.eq-modal-mask {
  position: fixed;
  z-index: 9999999999 !important;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background-color: rgba(0, 0, 0, 0.7);
  display: table;
}

.eq-modal-wrapper {
  display: table-cell;
  vertical-align: middle;
  text-align: center; /* Ensures the modal is centered horizontally */
}

.eq-modal-container {
  display: inline-block; /* Adapts the width to content size */
  padding: 20px 30px;
  background-color: #fff;
  border-radius: 8px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.33);
  transition: all 0.3s ease;
  z-index: 99999999999 !important;

  /* Constraints */
  max-width: 90%; /* Prevents the modal from becoming too wide */
  max-height: 100%; /* Ensures the modal doesn't exceed screen height */
}

.eq-modal-body {
  max-height: 90vh;
  overflow-y: auto; /* Allows scrolling if content exceeds height */
  text-align: left;
  overflow-x: hidden;
}

.eq-modal-default-button {
  float: right;
}

.eq-modal-footer {
  min-height: 25px;
}

.eq-modal-enter .eq-modal-container,
.eq-modal-leave-active .eq-modal-container {
  -webkit-transform: scale(1.1);
  transform: scale(1.1);
}

</style>
