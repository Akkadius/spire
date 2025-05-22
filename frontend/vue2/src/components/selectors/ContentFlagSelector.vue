<template>
  <div v-if="inputVal">
    <b-form-group label-for="tags-component-select" class="mb-0">
      <!-- Prop `add-on-change` is needed to enable adding tags vie the `change` event -->

      <!--      {{inputVal}}-->
      <b-form-tags
        id="tags-component-select"
        v-model="inputVal"
        class="mb-2 p-0"
        add-on-change
        no-outer-focus
        style="background-color: transparent; border: transparent"
      >
        <template v-slot="{ tags, inputAttrs, inputHandlers, disabled, removeTag }">
          <ul v-if="tags.length > 0" class="list-inline d-inline-block mb-2">
            <li v-for="tag in tags" :key="tag" class="list-inline-item">
              <b-form-tag
                @remove="removeTag(tag)"
                :title="tag"
                :disabled="disabled"
                variant="warning"
              >{{ tag }}
              </b-form-tag>
            </li>
          </ul>
          <b-form-select
            v-bind="inputAttrs"
            v-on="inputHandlers"
            :disabled="disabled || availableOptions.length === 0"
            :options="availableOptions"
          >
            <template #first>
              <!-- This is required to prevent bugs with Safari -->
              <option disabled value="">{{ availableOptions.length === 0 ? 'All Flags Selected' : 'Add a flag...' }}
              </option>
            </template>
          </b-form-select>
        </template>
      </b-form-tags>
    </b-form-group>
  </div>
</template>

<script>
import {ContentFlags} from "../../app/content-flags";

export default {
  name: "ContentFlagSelector",
  props: {
    value: [String, Array]
  },
  data() {
    return {
      options: [],
      availableOptions: []
    }
  },
  async mounted() {
    // @ts-ignore
    let flags = []

    for (let f of await ContentFlags.get()) {
      flags.push(f.flag_name)
    }
    this.options = flags
    this.calcAvailableOptions()
  },
  methods: {
    calcAvailableOptions() {
      console.log("calcing options", this.inputVal)
      this.availableOptions = this.options.filter(opt => this.inputVal.indexOf(opt) === -1)
      console.log(this.availableOptions)
    },
  },

  watch: {
    value: {
      handler(newVal) {
        console.log("[ContentFlagSelector] value watcher [%s]", this.value)
        setTimeout(() => {
          console.log("firing options")
          this.calcAvailableOptions()
        }, 100)
      },
    },
    deep: true
  },

  computed: {
    inputVal: {
      get() {
        let val = this.value

        console.log("[ContentFlagSelector] value computed [%s]", val)

        if (typeof val === 'undefined' || val === null) {
          return []
        }

        if (Array.isArray(val)) {
          return val && val.length > 0 ? val : []
        }

        if (typeof val === 'string' || val instanceof String) {
          return val.split(",")
        }

        return []
      },
      set(val) {
        const string = val ? val.join(",") : ""

        console.log("[ContentFlagSelector] value computed set [%s]", string)

        // loop avoidance - keep from circular updating
        if (this.value !== string) {
          this.$emit('input', string);
        }
      }
    }
  }
}
</script>
