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

          <!-- range -->
          <b-form-input
            v-if="field.fType === 'range'"
            type="range"
            :id="field.field"
            :min="field.min"
            :max="field.max"
            style="width: 80%"
            v-model.number="editMerchantEntry[field.field]"
            class="m-0 mt-1 d-inline-block"
            v-on="field.e ? getEventHandlers(field.e, field.field) : {}"
            v-b-tooltip.hover.v-dark.right :title="getFieldDescription(field.field)"
            :style="(editMerchantEntry[field.field] <= (typeof field.zeroValue !== 'undefined' ? field.zeroValue : 0) ? 'opacity: .5' : '')"
            @change="rerenderProbability = Date.now()"
            :key="rerenderProbability"
          />

          <div v-if="field.fType === 'range'" class="d-inline-block ml-3" :key="rerenderProbability + '-visual'">
            ({{editMerchantEntry[field.field]}}) %
          </div>

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
            @input="editMerchantEntry[field.field] = $event; rerenderContentFlags = Date.now()"
            :key="rerenderContentFlags"
            v-b-tooltip.hover.v-dark.right :title="getFieldDescription(field.field)"
          />

          <content-expansion-selector
            v-if="field.fType === 'content-expansion'"
            :value="editMerchantEntry[field.field]"
            @input="editMerchantEntry[field.field] = $event; rerenderExpansion = Date.now()"
            :key="rerenderExpansion"
            v-b-tooltip.hover.v-dark.right :title="getFieldDescription(field.field)"
          />

          <class-bitmask-calculator
            v-if="field.fType === 'classes'"
            style="border-radius: 15px; min-height: 150px;"
            class="text-center mt-3"
            :show-text-top="false"
            :show-text-side="true"
            :imageSize="38"
            :centered-buttons="true"
            @input="editMerchantEntry[field.field] = parseInt($event); rerenderClasses = Date.now()"
            :key="rerenderClasses"
            :mask="editMerchantEntry[field.field]"
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
        class="mt-4"
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
import EqDebug                  from "../../../components/eq-ui/EQDebug";
import EqWindow                 from "../../../components/eq-ui/EQWindow";
import EqCheckbox               from "../../../components/eq-ui/EQCheckbox";
import {Merchants}              from "../../../app/merchants";
import InfoErrorBanner          from "../../../components/InfoErrorBanner";
import ContentFlagSelector      from "../../../components/selectors/ContentFlagSelector";
import ContentExpansionSelector from "../../../components/selectors/ContentExpansionSelector";
import ClassBitmaskCalculator   from "../../../components/tools/ClassBitmaskCalculator";

export default {
  name: "MerchantlistEntryEdit",
  components: { ClassBitmaskCalculator, ContentExpansionSelector, ContentFlagSelector, InfoErrorBanner, EqCheckbox, EqWindow, EqDebug },
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

      // rerender properties
      rerenderContentFlags: 0,
      rerenderExpansion: 0,
      rerenderClasses: 0,
      rerenderProbability: 0,

      // fields
      editMerchantEntryFields: [
        { desc: "Faction Requirement", field: "faction_required", fType: "text" },
        { desc: "Level Requirement", field: "level_required", fType: "text" },
        { desc: "Alternate Currency Cost", field: "alt_currency_cost", fType: "text" },
        { desc: "Enabled with Classes", field: "classes_required", fType: "classes" },
        { desc: "Probability", field: "probability", fType: "range", min: 0, max: 100 },
        { desc: "Min Expansion", field: "min_expansion", fType: "content-expansion" },
        { desc: "Max Expansion", field: "max_expansion", fType: "content-expansion" },
        { desc: "Enabled on Content Flag(s)", field: "content_flags", fType: "content-flag" },
        { desc: "Disabled on Content Flag(s)", field: "content_flags_disabled", fType: "content-flag" },
      ],
    }
  },
  methods: {

    /**
     * Tabs / fields
     */
    getEventHandlers(e, field) {
      let handlers = {}
      if (e.onclick) {
        handlers.click = () => e.onclick(field)
      }
      if (e.onmouseover) {
        handlers.mouseover = () => e.onmouseover(field)
      }

      return handlers
    },

    /**
     * Misc
     */
    getFieldDescription(field) {
      const descriptions = {
        "classes_required": "This Merchant entry is enabled the client falls within the classes selected",
        "min_expansion": "This Merchant entry is enabled when the server's current expansion is above this value if not -1",
        "max_expansion": "This Merchant entry is enabled when the server's current expansion is below this value if not -1",
        "content_flags": "This Merchant entry is enabled when these content flags are enabled on the server",
        "content_flags_disabled": "This Merchant entry is enabled when these content flags are disabled on the server",
        "probability": "This Merchant entry is enabled on successful roll on the spawn of the Merchant (0-100%)",
      }

      // we do this because the payload we get back from spire API is
      // formatted slightly different
      let fieldLookup = field.toLowerCase().replace("_", "")

      for (let key in descriptions) {
        let keyLookup = key.toLowerCase().replace("_", "")
        if (keyLookup === fieldLookup) {
          return descriptions[key]
        }
      }
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
