<template>
  <transition name="modal">
    <div class="eq-modal-mask">
      <div class="eq-modal-wrapper" @click.self="dismiss()">
        <eq-window class="eq-modal-container" style="width: 1000px">

          <div class="eq-modal-header">
            <slot name="header">
              default header
            </slot>
          </div>

          <div class="eq-modal-body">
            <slot name="body">
              default body
            </slot>
          </div>

          <div class="eq-modal-footer">
            <slot name="footer">
              default footer
              <button class="eq-modal-default-button" @click="$emit('close')">
                OK
              </button>
            </slot>
          </div>
        </eq-window>
      </div>
    </div>
  </transition>
</template>

<script>
import EqWindow from "@/components/eq-ui/EQWindow.vue";

export default {
  name: 'EqModal',
  components: { EqWindow },
  props: {
  },
  methods: {
    dismiss() {
      this.$emit('close');
    }
  }
}
</script>

<style>
.eq-modal-mask {
  position: fixed;
  z-index: 9998;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background-color: rgba(0, 0, 0, 0.5);
  display: table;
}

.eq-modal-wrapper {
  display: table-cell;
  vertical-align: middle;
}

.eq-modal-container {
  width: 300px;
  margin: 0px auto;
  padding: 20px 30px;
  background-color: #fff;
  border-radius: 2px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.33);
  transition: all 0.3s ease;
  font-family: Helvetica, Arial, sans-serif;
  z-index: 99999999999 !important;
}

.eq-modal-body {
  margin: 20px 0;
}

.eq-modal-default-button {
  float: right;
}

/*
 * The following styles are auto-applied to elements with
 * transition="modal" when their visibility is toggled
 * by Vue.js.
 *
 * You can easily play with the modal transition by editing
 * these styles.
 */

.eq-modal-enter {
  opacity: 0;
}

.eq-modal-leave-active {
  opacity: 0;
}

.eq-modal-enter .eq-modal-container,
.eq-modal-leave-active .eq-modal-container {
  -webkit-transform: scale(1.1);
  transform: scale(1.1);
}

</style>
