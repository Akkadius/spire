<template>
  <div>
    <nav class='eq-tab-box-fancy'>
      <ul :style="'margin-bottom: ' + bottomTabMargin + 'px !important;'">
        <li
          v-for="tab in tabs"
          :class="{ 'eq-tab-open': tab.isActive }"
          @mouseover="selectTabHover(tab)"
          @click="selectTab(tab)"
        >
          <a
            @click="selectTab(tab)"
            style="color: white"
          >
            {{ tab.name }}
          </a>
        </li>
      </ul>
    </nav>

    <div class="tabs-details">
      <slot></slot>
    </div>
  </div>
</template>

<script>
import {LocalSettings} from "@/app/local-settings/localsettings";

export default {
  name: 'EqTabs',
  data() {
    return { tabs: [] };
  },
  watch: {
    'selected'() {
      if (this.selected && this.selected.length > 0) {
        this.tabs.forEach(tab => {
          tab.isActive = tab.name === this.selected;
        })
      }
    },
  },
  created() {
    this.tabs = this.$children;
  },
  mounted() {
    if (this.selected && this.selected.length > 0) {
      this.tabs.forEach(tab => {
        tab.isActive = tab.name === this.selected;
      })
    }
  },
  methods: {
    selectTabHover(selectedTab) {
      if (this.hoverOpen || LocalSettings.isTabHoverEnabled()) {
        this.selectTab(selectedTab)
      }
    },
    selectTab(selectedTab) {
      this.tabs.forEach(tab => {
        tab.isActive = (tab.name === selectedTab.name);
        if (tab.isActive) {
          this.$emit("on-selected", tab.name)
        }
      });
    }
  },
  props: {
    selected: {
      default: "",
      required: false,
      type: String,
    },
    hoverOpen: {
      default: false,
      required: false,
      type: Boolean,
    },
    bottomTabMargin: {
      default: 20,
      required: false,
      type: Number
    }
  }
}
</script>

<style scoped>

</style>
