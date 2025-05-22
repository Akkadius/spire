<template>
  <div>
    <eq-window-simple title="Description Selector">
      <div class="row" v-if="fields">
        <div
          v-for="field in fields"
          :class="field.col + ' mb-3 pl-2 pr-2'"
        >
          <div class="mb-1">
            <span
              v-if="field.itemIcon"
              :class="'item-' + field.itemIcon + '-sm'"
              style="display: inline-block"
            />
            {{ field.description }}
          </div>

          <!-- Activity -->

          <!-- input text -->

          <div class="row" v-if="field.description === 'Activities' && activityDescriptions[field.index]">
            <div class="col-3 text-right">
              <b-button
                class="btn-dark btn-sm btn-outline-danger mt-1"
                @click="deleteActivityDescription(field.index)"
              >
                <i class="fa fa-trash"></i>
              </b-button>
            </div>
            <div class="col-9">
              <b-form-input
                v-if="field.fieldType === 'text' && field.activity"
                v-model.number="activityDescriptions[field.index].multiKey"
                @keyup="sendDescriptionToParentDebounce"
                @change="sendDescriptionToParent"
                class="m-0 mt-1"
              />
            </div>
          </div>

          <!-- textarea -->
          <b-textarea
            v-if="field.fieldType === 'textarea' && field.activity && activityDescriptions[field.index]"
            v-model="activityDescriptions[field.index].description"
            class="m-0 mt-1"
            @keyup="sendDescriptionToParentDebounce"
            @change="sendDescriptionToParent"
            rows="3"
            max-rows="6"
          ></b-textarea>

          <!-- Global -->

          <!-- textarea -->
          <b-textarea
            v-if="field.fieldType === 'textarea' && !field.activity"
            v-model="globalDescription"
            class="m-0 mt-1"
            @keyup="sendDescriptionToParentDebounce"
            @change="sendDescriptionToParent"
            rows="2"
            max-rows="6"
          ></b-textarea>

          <div
            class="mt-3"
            v-if="field.fieldType === 'textarea' && activityDescriptions[field.index] && activityDescriptions[field.index].multiKey && activityDescriptions[field.index].multiKey.length > 0"
          >
            <div v-for="activityId in activityDescriptions[field.index].multiKey.split(',')">
              <span class="font-weight-bold">
                {{ buildActivityDescriptionFromIndex(activityId).prefix }}
              </span>
              {{ buildActivityDescriptionFromIndex(activityId).description }}
            </div>
          </div>

        </div>
      </div>

      <div class="row">
        <div class="col-12">
          <b-button
            class="btn-dark btn-sm btn-dark mr-3"
            @click="addActivityDescriptionRow"
          >
            <i class="fa fa-plus"></i>
          </b-button>

          Add activity level description
        </div>
      </div>

    </eq-window-simple>
  </div>
</template>

<script>
import EqWindowSimple from "@/components/eq-ui/EQWindowSimple";
import EqCheckbox     from "@/components/eq-ui/EQCheckbox";
import {Tasks}        from "@/app/tasks";
import {debounce} from "@/app/utility/debounce";

export default {
  name: "TaskDescriptionSelector",
  components: { EqCheckbox, EqWindowSimple },
  data() {
    return {
      globalDescription: "",
      activityDescriptions: [],
      fields: []
    }
  },
  props: {
    task: {
      type: Object,
      required: false
    },
    description: {
      type: String,
      required: true,
    },
  },
  mounted() {
    this.init()
  },
  watch: {
    task: {
      handler(newVal) {
        // this.init()
      },
      deep: true
    },
    description: {
      handler(newVal) {
        // this.init()
      },
      deep: true
    },
    activityDescriptions: {
      handler(newVal) {
      },
      deep: true
    }
  },
  methods: {
    buildActivityDescriptionFromIndex(index) {
      if (this.task.task_activities[index - 1]) {
        const activity = this.task.task_activities[index - 1]
        const prefix   = `Step [${activity.step}] Activity [${activity.activityid}] `
        return {
          prefix: prefix,
          description: this.buildActivityDescription(activity)
        }

      }
      return {}
    },

    buildActivityDescription(activity) {
      return Tasks.buildActivityDescription(activity)
    },

    init() {
      this.getDescription()
    },

    deleteActivityDescription(index) {

      // delete from form fields first
      let fields = []
      for (let i in this.fields) {
        if (this.fields[i].activity && this.fields[i].index === index) {
          continue
        }
        fields.push(this.fields[i])
      }

      this.fields = fields
      this.$forceUpdate()

      // delete from activity description v-model binding after
      let activityDesc = []
      for (let i in this.activityDescriptions) {
        if (parseInt(i) === parseInt(index)) {
          continue
        }
        activityDesc[i] = this.activityDescriptions[i]
      }

      this.activityDescriptions = activityDesc
      this.sendDescriptionToParent()
    },

    addActivityDescriptionRow() {
      let newIndex         = 0;
      let nextActivityStep = 1
      for (let key in this.activityDescriptions) {
        newIndex = key + 1

        // best effort, this could be comma separated
        if (parseInt(this.activityDescriptions[key].multiKey) >= 0) {
          nextActivityStep = parseInt(this.activityDescriptions[key].multiKey) + 1
        }
      }

      // add new activity row
      this.activityDescriptions[newIndex] =
        {
          multiKey: nextActivityStep,
          description: ""
        }

      this.fields.push(
        {
          description: 'Activities',
          itemIcon: '5739',
          fieldType: 'text',
          col: 'col-3',
          activity: true,
          index: newIndex,
        },
        {
          description: 'Description',
          fieldType: 'textarea',
          itemIcon: '2275',
          col: 'col-9',
          activity: true,
          index: newIndex,
        },
      )
    },

    sendDescriptionToParentDebounce: debounce(function () {
      this.sendDescriptionToParent()
    }, 500),

    sendDescriptionToParent() {
      let finalDescription = `${this.globalDescription}`
      for (let key in this.activityDescriptions) {
        const multiKey    = this.activityDescriptions[key].multiKey
        const description = this.activityDescriptions[key].description
        finalDescription += `[${multiKey},${description}]`
      }

      finalDescription = finalDescription.replaceAll("\n", "<BR>")

      this.$forceUpdate()
      this.$emit('input', finalDescription);
    },

    replaceDescriptionContent(s) {
      let d = s

      d = d.replaceAll("]", "")
      d = d.replaceAll("<BR>", "<br />")
      d = d.replaceAll("<br>", "<br />")
      d = d.replaceAll("<c \"", "<span style=\"color: ")
      d = d.replaceAll("</c>", "</span>")
      d = d.replace(/^\s+|\s+$/g, "")
      return d
    },
    getDescription() {
      // set global description
      this.globalDescription = this.description.split("[")[0]

      // push fields
      this.fields.push(
        {
          description: 'Global Description (Available in all activities)',
          field: 'title',
          fieldType: 'textarea',
          itemIcon: '6840',
          col: 'col-12',
          model: this.globalDescription,
        },
      )

      const descriptionParts = this.description.split("[")

      let index                = 0
      let activityDescriptions = []
      descriptionParts.forEach((p) => {
        // 2,3,Your first clue is that the...
        let indexes = []
        p.split(",").forEach((s) => {
          if (parseInt(s)) {
            indexes.push(s)
          }
        })

        // get description in part
        let description = p
        description     = description.replaceAll(indexes.join(",") + ",", "")
        description     = this.replaceDescriptionContent(description)

        // set desc on multi-key
        const multiKey = indexes.join(",")
        if (multiKey !== "") {
          activityDescriptions[index] =
            {
              multiKey: multiKey,
              description: description.trim(),
            }

          this.fields.push(
            {
              description: 'Activities',
              itemIcon: '5739',
              fieldType: 'text',
              col: 'col-3',
              activity: true,
              index: index,
            },
            {
              description: 'Description',
              fieldType: 'textarea',
              itemIcon: '2275',
              col: 'col-9',
              activity: true,
              index: index,
            },
          )

        }

        index++;
      })

      this.activityDescriptions = activityDescriptions
    },
  }
}
</script>

<style scoped>
</style>
