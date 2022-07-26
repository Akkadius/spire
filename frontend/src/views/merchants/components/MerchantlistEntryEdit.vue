<template>
  <div class="col-5" v-if="(Object.keys(editMerchantEntry).length > 0)">
    <eq-window :title="`Edit Merchant List Entry (${editMerchantEntrySlot})`">
      <div
        v-for="field in editMerchantEntryFields"
        :key="field.field"
        :class="'row minified-inputs'"
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

          <content-flag-selector
            v-if="field.fType === 'content-flag'"
            :value="editMerchantEntry[field.field]"
            @input="editMerchantEntry[field.field] = $event; rerender = Date.now()"
            :key="rerender"
          />

        </div>
      </div>

      <!-- Save -->
      <div class="row">
        <div class="col-12 text-center mt-3">
          <b-button
            @click="save()"
            size="sm"
            variant="outline-warning"
          >
            <i class="ra ra-save"></i>
            Save
          </b-button>
        </div>
      </div>

      <!-- Notification / Error -->
      <info-error-banner
        :notification="notification"
        :error="error"
        @dismiss-error="error = ''"
        @dismiss-notification="notification = ''"
      />

      <eq-debug :data="editMerchantEntry"/>
    </eq-window>
  </div>
</template>

<script>
import EqDebug             from "../../../components/eq-ui/EQDebug";
import EqWindow            from "../../../components/eq-ui/EQWindow";
import EqCheckbox          from "../../../components/eq-ui/EQCheckbox";
import {Merchants}         from "../../../app/merchants";
import InfoErrorBanner     from "../../../components/InfoErrorBanner";
import ContentFlagSelector from "../../../components/selectors/ContentFlagSelector";

export default {
  name: "MerchantlistEntryEdit",
  components: { ContentFlagSelector, InfoErrorBanner, EqCheckbox, EqWindow, EqDebug },
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

      // local notification / error state
      notification: "",
      error: "",

      // rerender property
      rerender: 0,

      // fields
      editMerchantEntryFields: [
        { desc: "Faction Requirement", field: "faction_required", fType: "text" },
        { desc: "Level Requirement", field: "level_required", fType: "text" },
        { desc: "Alternate Currency Cost", field: "alt_currency_cost", fType: "text" },
        { desc: "Classes Required", field: "classes_required", fType: "text" },
        { desc: "Probability", field: "probability", fType: "text" },
        { desc: "Min Expansion", field: "min_expansion", fType: "text" },
        { desc: "Max Expansion", field: "max_expansion", fType: "text" },
        { desc: "Enabled on Content Flag(s)", field: "content_flags", fType: "content-flag" },
        { desc: "Disabled on Content Flag(s)", field: "content_flags_disabled", fType: "content-flag" },
      ],
    }
  },
  methods: {

    // need to force update to re-propogate object update to the component
    handleContentFlagUpdate(field, e) {
      console.log("[handleContentFlagUpdate] field [%s] val [%s]", field, e)
      this.editMerchantEntry[field] = e
      this.$forceUpdate()
    },

    /**
     * Misc
     */
    getFieldDescription(field) {
      return ""
    },


    /**
     * Save
     */
    async save() {
      let e = this.editMerchantEntry

      // test error
      // e.classes_required = 909032540354

      try {
        const r = await Merchants.updateSlotForEntry(
          e.merchantid,
          e.slot,
          e
        )
        if (r.status === 200) {
          this.notification = `Saved!`
        }
      } catch (err) {
        console.log(err)
        if (err.response && err.response.data && err.response.data.error) {
          this.error = err.response.data.error
        }
      }
    }

  }
}
</script>

<style scoped>

</style>
