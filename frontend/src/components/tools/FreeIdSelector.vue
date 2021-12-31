<template>
  <div>
    <eq-tabs
      v-if="freeIds"
    >
      <eq-tab
        :name="'Free ID Ranges (' + freeIds.length + ')'"
        selected="true"
      >
        <div class="col-12 text-center" style="height: 80vh; overflow-y: scroll">
          <div class="mb-3 mt-3">These are unused IDs with contiguous blocks of 10 at minimum</div>
          <div v-for="range in freeIds" :key="range.start_id" class="mt-2 row">
            <div class="col-5 text-right pr-0">
              <b-button
                class="btn-dark btn-sm btn-outline-warning"
                @click="selectId(range.start_id)"
              >
                Use {{ range.start_id }}
              </b-button>
            </div>
            <div class="col-2 pl-0 pr-0">
              {{ range.end_id - range.start_id }} Free
            </div>
            <div class="col-5 text-left pl-0">
              <b-button
                class="btn-dark btn-sm btn-outline-warning"
                @click="selectId(range.end_id)"
              >
                Use {{ range.end_id }}
              </b-button>
            </div>
          </div>
        </div>

      </eq-tab>
      <eq-tab
        v-if="withReserved"
        :name="'Reserved ID Ranges (' + freeIdsReserved.length + ')'"
      >
        <div class="col-12 text-center" style="height: 80vh; overflow-y: scroll">
          <div class="mb-3 mt-3">These are reserved IDs with "placeholder" or "reserved" in the name</div>
          <div v-for="id in freeIdsReserved" :key="id.start_id" class="mt-2 row">
            <div class="col-12 text-center pr-0">
              <b-button
                class="btn-dark btn-sm btn-outline-warning"
                @click="selectId(id.id)"
              >
                Use {{ id.id }} ({{id.name}})
              </b-button>
            </div>
          </div>
        </div>

      </eq-tab>
    </eq-tabs>
  </div>
</template>

<script>
import {SpireApiClient} from "../../app/api/spire-api-client";
import * as util        from "util";
import EqTabs           from "../eq-ui/EQTabs";
import EqTab            from "../eq-ui/EQTab";

export default {
  name: "FreeIdSelector",
  components: { EqTab, EqTabs },
  data() {
    return {
      freeIds: [],
      freeIdsReserved: []
    }
  },
  props: {
    tableName: {
      type: String,
      required: true
    },
    idName: {
      type: String,
      required: true,
    },
    nameLabel: {
      type: String,
      required: false,
    },
    withReserved: {
      type: Boolean,
      default: false,
      required: false,
    }
  },
  methods: {
    selectId(id) {
      this.$emit('input', parseInt(id));
    }
  },
  mounted() {
    SpireApiClient.v1().get(
      util.format(
        `query/free-id-ranges/%s/%s`,
        this.tableName,
        this.idName
      )
    ).then((response) => {
      if (response.data && response.data.data) {
        this.freeIds = response.data.data
      }
    });

    if (this.withReserved) {
      SpireApiClient.v1().get(
        util.format(
          `query/free-ids-reserved/%s/%s/%s`,
          this.tableName,
          this.idName,
          this.nameLabel,
        )
      ).then((response) => {
        if (response.data && response.data.data) {
          this.freeIdsReserved = response.data.data
        }
      });
    }
  }
}
</script>

<style scoped>

</style>
