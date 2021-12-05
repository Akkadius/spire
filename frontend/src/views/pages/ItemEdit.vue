<template>
  <div class="container-fluid">
    <div class="panel-body">
      <div class="panel panel-default">
        <div class="row">
          <div class="col-7">
            <eq-window style="margin-top: 30px" title="Edit Item">

              <div v-if="notification">
                <b-button class="btn-dark btn-outline-warning form-control" @click="notification = ''">
                  <i class="ra ra-book mr-1"></i>{{ notification }}
                </b-button>
              </div>

              <b-alert show dismissable variant="danger" v-if="error">
                <i class="fa fa-warning"></i> {{ error }}
              </b-alert>

              <eq-tabs
                v-if="item"
                id="item-edit-card"
                class="item-edit-card"
                :hover-open="true"
                @mouseover.native="previewItem"
              >
                <eq-tab
                  name="General"
                  :selected="true"
                >
                  <div class="row">
                    <div class="col-2" @mouseover="drawFreeIdSelector">
                      Id
                      <b-form-input v-model.number="item.id"/>
                    </div>
                    <div class="col-7">
                      Name
                      <b-form-input
                        :value="item.name" @change="v => item.name = v"
                      />
                    </div>

                    <div class="col-2" @mouseover="drawIconSelector">
                      Icon
                      <b-form-input v-model.number="item.icon"/>
                    </div>

                    <div
                      class="col-1" v-if="item.icon > 0"
                      style="margin-top: 7px"
                      @mouseover="drawIconSelector"
                    >
                      <span
                        :class="'fade-in item-' + item.icon"
                        style="border: 1px solid rgb(218 218 218 / 30%); border-radius: 7px;"/>
                    </div>

                  </div>

                  <!--                  <item-class-selector :spell="spell" @input="spell = $event"/>-->
                  <!--                  <item-deity-selector :spell="spell" @input="spell = $event"/>-->

                  <div class="row">
                    <div class="col-8">

                      Stuff goes here...

                    </div>
                    <div
                      class="col-4"
                      style="text-align: center"
                      @mouseover="drawItemModelSelector"
                    >

                      <item-model-preview :id="item.idfile"/>

                      <!--                      <item-animation-preview-->
                      <!--                        class="mt-4"-->
                      <!--                        :id="item.itemanim"/>-->
                      <!---->
                      Item Model
                      <b-form-input v-model.number="item.idfile"/>
                    </div>
                  </div>


                </eq-tab>
              </eq-tabs>

              <div class="text-center mt-3">
                <b-button
                  class="btn-dark btn-sm btn-outline-warning"
                  @click="saveItem"
                >
                  <i class="ra ra-book mr-1"></i>
                  Save Item
                </b-button>
              </div>
            </eq-window>
          </div>

          <div class="col-5">

            <!-- preview item -->
            <eq-window
              style="margin-top: 30px; margin-right: 10px; width: auto;"
              v-if="previewItemActive && item">
              <eq-item-preview
                :item-data="item"/>
            </eq-window>

            <!-- icon selector -->
            <eq-window
              style="margin-top: 30px; margin-right: 10px; width: auto;"
              class="fade-in"
              v-if="iconSelectorActive">
              <!--              <item-icon-selector-->
              <!--                :selected-icon="item.new_icon"-->
              <!--                :inputData.sync="item.new_icon"-->
              <!--              />-->
            </eq-window>

            <!-- free id selector -->
            <eq-window
              title="Free Item Ids"
              style="margin-top: 30px; margin-right: 10px; width: auto;"
              class="fade-in"
              v-if="freeIdSelectorActive">

              <free-id-selector
                table-name="items"
                id-name="id"
                name-label="name"
                :with-reserved="true"
                @input="item.id = $event"/>
            </eq-window>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import EqWindowFancy    from "../../components/eq-ui/EQWindowFancy";
import EqWindow         from "../../components/eq-ui/EQWindow";
import EqTabs           from "../../components/eq-ui/EQTabs";
import EqTab            from "../../components/eq-ui/EQTab";
import EqItemPreview    from "../../components/eq-ui/EQItemCardPreview";
import {App}            from "../../constants/app";
import EqCheckbox       from "../../components/eq-ui/EQCheckbox";
import {SpireApiClient} from "../../app/api/spire-api-client";
import * as util        from "util";
import FreeIdSelector   from "../../components/tools/FreeIdSelector";
import {Items}          from "../../app/items";
import {ItemApi}        from "../../app/api";
import ItemModelPreview from "../../components/tools/ItemModelPreview";

const MILLISECONDS_BEFORE_WINDOW_RESET = 3000;

export default {
  name: "ItemEdit",
  components: {
    ItemModelPreview,
    FreeIdSelector,
    EqCheckbox,
    EqItemPreview,
    EqTab,
    EqTabs,
    EqWindow,
    EqWindowFancy
  },
  data() {
    return {
      item: null, // item record data
      spellCdnUrl: App.ASSET_SPELL_ICONS_BASE_URL,
      loaded: true,

      previewItemActive: true,
      iconSelectorActive: false,
      spellAnimSelectorActive: false,
      freeIdSelectorActive: false,

      lastResetTime: Date.now(),

      notification: "",
      error: "",
    }
  },
  watch: {
    '$route'() {
      // reset state vars when we navigate away
      this.notification = ""
      // reload
      this.load()
    }
  },
  async created() {

    setTimeout(() => {
      document.getElementById("item-edit-card").removeEventListener('input', this.setFieldModified, true);
      document.getElementById("item-edit-card").addEventListener('input', this.setFieldModified)
    }, 1000)

    this.load()
  },
  methods: {

    setFieldModified(evt) {
      // border: 2px #555555 solid !important;
      evt.target.style.setProperty('border-color', 'orange', 'important');
    },

    resetFieldEditedStatus() {
      // reset elements
      const elements = document.getElementById("item-edit-card").querySelectorAll("input, select");
      elements.forEach((element) => {
        element.style.setProperty('border-color', '#555555', 'important');
      });
    },

    async saveItem() {
      this.error        = ""
      this.notification = ""

      const api = (new ItemApi(SpireApiClient.getOpenApiConfig()))
      api.updateItem({
        id: this.item.id,
        item: this.item
      }).then((result) => {
        if (result.status === 200) {
          this.notification = util.format("Item updated successfully! (%s) %s", this.item.id, this.item.name)
          this.resetFieldEditedStatus()
        }

        if (result.data.error) {
          this.notification = result.data.error
        }

      }).catch(async (error) => {

        // marshalling error
        if (error.response.data && error.response.data.error.includes("marshal")) {
          this.error = error.response.data.error
          return
        }

        const createRes = await api.createItem({
          item: this.item
        })

        if (createRes.status === 200) {
          this.notification = util.format("Created new Item! (%s) %s", this.item.id, this.item.name)
          this.resetFieldEditedStatus()
        }
      })
    },

    load() {
      if (this.$route.params.id > 0) {
        Items.getItem(this.$route.params.id).then(result => {
          this.item = result
        })
      }
    },

    /**
     * Selector / previewers
     */
    resetPreviewComponents() {
      this.previewItemActive      = false;
      this.iconSelectorActive     = false;
      this.itemAnimSelectorActive = false;
      this.freeIdSelectorActive   = false;
    },
    previewItem() {
      let shouldReset = Date.now() - this.lastResetTime > MILLISECONDS_BEFORE_WINDOW_RESET;
      // SECONDS_BEFORE_WINDOW_RESET

      if (!this.previewItemActive && shouldReset) {
        this.resetPreviewComponents()
        this.previewItemActive = true;
        this.lastResetTime     = Date.now()
      }
    },
    drawItemModelSelector() {
      this.resetPreviewComponents()
      this.itemAnimSelectorActive = true
    },
    drawIconSelector() {
      if (!this.freeIdSelectorActive) {
        this.resetPreviewComponents()
        this.iconSelectorActive = true;
      }
    },
    drawFreeIdSelector() {
      this.resetPreviewComponents()
      this.lastResetTime        = Date.now()
      this.freeIdSelectorActive = true
    },
    getTargetTypeColor(targetType) {
      return Items.getTargetTypeColor(targetType);
    },
  }
}
</script>

<style scoped>
.item-edit-card input, .item-edit-card select {
  margin-bottom: 10px;
}

.effect-tab input, .effect-tab select {
  margin-bottom: 0;
}
</style>
