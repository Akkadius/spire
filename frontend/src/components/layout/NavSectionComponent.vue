<template>
  <li class="nav-item" v-show="navId !== '' && config && config.label">
    <a
      :class="'nav-link collapse ' + (hasRoute(config.routePrefixMatch) ? 'active' : 'collapsed')"
      :href="'#sidebar-' + navId"
      data-toggle="collapse"
      role="button"
      :aria-expanded="(hasRoute(config.routePrefixMatch) ? 'true' : 'false')"
      :aria-controls="'sidebar-' + navId"
    >
      <i :class="config.labelIcon" v-if="config.labelIcon"></i>
      {{ config.label }}
    </a>
    <div
      :class="'collapse ' + (hasRoute(config.routePrefixMatch) ? 'show' : '')"
      :id="'sidebar-' + navId"
    >
      <ul class="nav nav-sm flex-column">
        <li v-for="nav in config.navs">
          <router-link class="nav-link" :to="nav.to">
            <i :class="nav.icon" v-if="nav.icon"></i>{{ nav.title }}
            <b-badge class="ml-3 d-inline-block" variant="primary" v-if="nav.isNew">NEW!</b-badge>
          </router-link>
        </li>
      </ul>
    </div>
  </li>
</template>

<script>
export default {
  name: "NavSectionComponent",
  methods: {
    hasRoute: function (partial) {
      return (this.$route.path.indexOf(partial) > -1)
    } // config.topLevelIcon
  },
  data() {
    return {
      navId: ""
    }
  },
  props: {
    config: {
      type: Object
    }
  },
  created() {
    const uuidv4 = require("uuid/v4")
    this.navId   = uuidv4()
  }
}
</script>

<style scoped>

</style>
