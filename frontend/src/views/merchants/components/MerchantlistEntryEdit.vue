<template>
  <div class="col-5" v-if="(Object.keys(editMerchantEntry).length > 0)">
    <eq-window :title="`Edit Merchant List Entry (${editMerchantEntrySlot})`">
      <div
        v-for="field in editMerchantEntryFields"
        :key="field.field"
        :class="'row'"
      >
        <div
          class="col-4 text-right p-0 m-0 mr-1 mt-3"
          style="position: relative; bottom: 6px;"
          v-if="field.fType === 'checkbox'"
        >
          <span v-if="field.category" class="font-weight-bold">{{ field.category }}</span>
          {{ field.desc }}
        </div>
        <div
          class="col-4 text-right p-0 m-0 mr-3"
          v-if="field.fType !== 'checkbox'"
          style="margin-top: 10px !important"
        >
          <span v-if="field.category" class="font-weight-bold">{{ field.category }}</span>
          {{ field.desc }}
        </div>

        <!--                  <div class="text-center" v-if="field.fType !== 'checkbox'">-->
        <!--                    <span-->
        <!--                      v-if="field.itemIcon"-->
        <!--                      :class="'item-' + field.itemIcon + '-sm'"-->
        <!--                      style="display: inline-block"-->
        <!--                    />-->
        <!--                    {{ field.desc }}-->
        <!--                  </div>-->

        <div class="col-7 text-left p-0 mt-2">

          <!-- checkbox -->
          <div :class="'text-left ml-2 mt-1'" v-if="field.fType === 'checkbox'">
            <!--                        <div class="d-inline-block" style="bottom: 2px; position: relative; margin-right: 1px">-->
            <!--                          {{ field.desc }}-->
            <!--                        </div>-->
            <eq-checkbox
              v-b-tooltip.hover.v-dark.right :title="getFieldDescription(field.field)"
              class="d-inline-block text-center"
              :true-value="(typeof field.true !== 'undefined' ? field.true : 1)"
              :false-value="(typeof field.false !== 'undefined' ? field.false : 0)"
              v-model.number="editMerchantEntry[field.field]"
              @input="editMerchantEntry[field.field] = $event"

            />
          </div>

          <!-- input number -->
          <b-form-input
            v-if="field.fType === 'number'"
            :id="field.field"
            v-model.number="editMerchantEntry[field.field]"
            class="m-0 mt-1"
            v-on="field.e ? getEventHandlers(field.e, field.field) : {}"
            v-b-tooltip.hover.v-dark.right :title="getFieldDescription(field.field)"
            :style="(editMerchantEntry[field.field] === 0 ? 'opacity: .5' : '')"
          />

          <!-- input text -->
          <b-form-input
            v-if="field.fType === 'text'"
            :id="field.field"
            v-model.number="editMerchantEntry[field.field]"
            class="m-0 mt-1"
            v-on="field.e ? getEventHandlers(field.e, field.field) : {}"
            v-b-tooltip.hover.v-dark.right :title="getFieldDescription(field.field)"
            :style="(editMerchantEntry[field.field] <= (typeof field.zeroValue !== 'undefined' ? field.zeroValue : 0) ? 'opacity: .5' : '')"
          />

          <!-- textarea -->
          <b-textarea
            v-if="field.fType === 'textarea'"
            :id="field.field"
            v-model="editMerchantEntry[field.field]"
            class="m-0 mt-1"
            rows="2"
            max-rows="6"
            v-on="field.e ? getEventHandlers(field.e, field.field) : {}"
            v-b-tooltip.hover.v-dark.right :title="getFieldDescription(field.field)"
            :style="(editMerchantEntry[field.field] === '' ? 'opacity: .5' : '') + ';'"
          ></b-textarea>

          <!-- select -->
          <select
            v-model.number="editMerchantEntry[field.field]"
            :id="field.field"
            class="form-control m-0 mt-1"
            v-if="field.selectData"
            v-on="field.e ? getEventHandlers(field.e, field.field) : {}"
            v-b-tooltip.hover.v-dark.right :title="getFieldDescription(field.field)"
            :style="(editMerchantEntry[field.field] <= (typeof field.zeroValue !== 'undefined' ? field.zeroValue : 0) ? 'opacity: .5' : '')"
          >
            <option
              v-for="(desc, index) in field.selectData"
              :key="index"
              :value="parseInt(index)"
            >
              {{ index }}) {{ desc }}
            </option>
          </select>

        </div>
      </div>

      <eq-debug :data="editMerchantEntry"/>
    </eq-window>
  </div>
</template>

<script>
import EqDebug    from "../../../components/eq-ui/EQDebug";
import EqWindow   from "../../../components/eq-ui/EQWindow";
import EqCheckbox from "../../../components/eq-ui/EQCheckbox";
export default {
  name: "MerchantlistEntryEdit",
  components: { EqCheckbox, EqWindow, EqDebug },
  props: {
    editMerchantEntry: {
      type: Object,
      required: true
    },
    editMerchantEntrySlot: {
      type: Number,
      required: true
    },
  },
  data() {
    return {
      editMerchantEntryFields: [
        { desc: "faction_required", field: "faction_required", fType: "text" },
        { desc: "level_required", field: "level_required", fType: "text" },
        { desc: "alt_currency_cost", field: "alt_currency_cost", fType: "text" },
        { desc: "classes_required", field: "classes_required", fType: "text" },
        { desc: "probability", field: "probability", fType: "text" },
        { desc: "min_expansion", field: "min_expansion", fType: "text" },
        { desc: "max_expansion", field: "max_expansion", fType: "text" },
        { desc: "content_flags", field: "content_flags", fType: "text" },
        { desc: "content_flags_disabled", field: "content_flags_disabled", fType: "text" },
      ],
    }
  },
  methods: {
    /**
     * Misc
     */
    getFieldDescription(field) {
      return ""
    },

  }
}
</script>

<style scoped>

</style>
