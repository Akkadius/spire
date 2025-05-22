<template>
  <div>
    <EqModal
      v-show="isModalVisible"
      @close="closeModal"
    >
      <template v-slot:header>
        <h3 class="eq-header text-center">Modal Header</h3>
      </template>
      <template v-slot:body>

        <div class="text-center">
          <p>This is the body of the modal</p>
          <img src="https://www.giantbomb.com/a/uploads/original/0/1992/2118741-sleeper_fight2.jpg">
        </div>
      </template>
    </EqModal>

    <eq-window>

      <button
        type="button"
        class="eq-btn"
        @click="showModal"

      >
        <i class="fa fa-eye"></i>
        Open Modal
      </button>

      <div class="row mt-3">
        <div class="col-10">
          <input
            type="text"
            class="form-control"
            title="This is the lore description that will show en an item inspect window when it is identified."
          >
        </div>
        <div class="col-2">

          <button
            type="button"
            class="eq-btn form-control"
          >
            <i class="fa fa-save"></i> Save
          </button>
        </div>

      </div>

      <h3 class="eq-header mt-3">Buttons</h3>

      <div class="row mt-3">
        <div class="col-12">
          <button type="button" class="btn btn-primary">Primary</button>
          <button type="button" class="btn btn-secondary">Secondary</button>
          <button type="button" class="btn btn-success">Success</button>
          <button type="button" class="btn btn-danger">Danger</button>
          <button type="button" class="btn btn-warning">Warning</button>
          <button type="button" class="btn btn-info">Info</button>
          <button type="button" class="btn btn-light">Light</button>
          <button type="button" class="btn btn-white">White</button>
          <button type="button" class="btn btn-dark">Dark</button>
          <button type="button" class="btn btn-link">Link</button>
        </div>
      </div>

      <div class="row mt-3">
        <div class="col-12">
          <button type="button" class="btn btn-outline-primary">Primary</button>
          <button type="button" class="btn btn-outline-secondary">Secondary</button>
          <button type="button" class="btn btn-outline-success">Success</button>
          <button type="button" class="btn btn-outline-danger">Danger</button>
          <button type="button" class="btn btn-dark">Warning</button>
          <button type="button" class="btn btn-outline-info">Info</button>
          <button type="button" class="btn btn-outline-light">Light</button>
          <button type="button" class="btn btn-outline-white">White</button>
          <button type="button" class="btn btn-outline-dark">Dark</button>
        </div>
      </div>

    </eq-window>
  </div>
</template>

<script>
import EqWindow from "@/components/eq-ui/EQWindow.vue";
import EqModal  from "@/components/eq-ui/EQModal.vue";

export default {
  name: "Test",
  components: { EqModal, EqWindow },
  data() {
    return {
      options: ['something', 'something else', 'something else again'],

      isModalVisible: false,
    };
  },
  methods: {
    showModal() {
      this.isModalVisible = true;
    },
    closeModal() {
      this.isModalVisible = false;
    },
  }
}
</script>

<style scoped>


</style>
