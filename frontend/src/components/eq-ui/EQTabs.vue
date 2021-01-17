<template>
    <div>
        <nav class='eq-tab-box-fancy'>
            <ul>
                <li v-for="tab in tabs" :class="{ 'eq-tab-open': tab.isActive }"  @click="selectTab(tab)">
                    <a :href="tab.href" @click="selectTab(tab)" style="color:white">{{ tab.name }}</a>
                </li>
            </ul>
        </nav>

        <div class="tabs-details">
            <slot></slot>
        </div>
    </div>
</template>

<script>
  export default {
    name: 'EqTabs',
    data() {
      return { tabs: [] };
    },
    created() {
      this.tabs = this.$children;
    },
    methods: {
      selectTab(selectedTab) {
        this.tabs.forEach(tab => {
          tab.isActive = (tab.name === selectedTab.name);
        });
      }
    }
  }
</script>

<style scoped>

</style>
