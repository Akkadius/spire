<template>
  <li class="nav-item" v-show="navId !== '' && config && config.label">
    <a
      :class="'nav-link collapse ' + (hasRoute(config.routePrefixMatch) || hasRouteInArray(config.routePrefixMatches) ? 'active' : 'collapsed')"
      :href="'#sidebar-' + navId"
      data-toggle="collapse"
      role="button"
      :aria-expanded="(hasRoute(config.routePrefixMatch) || hasRouteInArray(config.routePrefixMatches) ? 'true' : 'false')"
      :aria-controls="'sidebar-' + navId"
    >
      <i :class="config.labelIcon" v-if="config.labelIcon"></i>
      {{ config.label }}
    </a>
    <div
      :class="'collapse ' + (hasRoute(config.routePrefixMatch) || hasRouteInArray(config.routePrefixMatches) ? 'show' : '')"
      :id="'sidebar-' + navId"
    >
      <ul class="nav nav-sm flex-column">
        <li v-for="nav in config.navs">

          <!-- internal link -->
          <router-link
            :class="'nav-link collapse ' + (hasRoute(nav.to) || hasRouteInArray(nav.routes) ? 'active' : 'collapsed')"
            :to="nav.to"
            v-if="!nav.to.includes('http')"
          >
            <i :class="nav.icon" v-if="nav.icon"></i>{{ nav.title }}
            <b-badge class="ml-3" variant="primary" v-if="nav.isAlpha">ALPHA</b-badge>
            <b-badge class="ml-3" variant="primary" v-if="nav.isNew">NEW!</b-badge>
          </router-link>

          <!-- external link -->
          <a
            :class="'nav-link collapse ' + (hasRoute(nav.to) || hasRouteInArray(nav.routes) ? 'active' : 'collapsed')"
            :href="nav.to"
            :target="nav.to"
            v-if="nav.to.includes('http')"
          >
            <i :class="nav.icon" v-if="nav.icon"></i>{{ nav.title }}
            <b-badge class="ml-3" variant="primary" v-if="nav.isAlpha">ALPHA</b-badge>
            <b-badge class="ml-3" variant="primary" v-if="nav.isNew">NEW!</b-badge>
          </a>
        </li>
      </ul>
    </div>
  </li>
</template>

<script>
export default {
  name: "NavSectionComponent",
  methods: {
    hasRouteInArray(matches) {
      let matched = false
      if (matches && matches.length > 0) {
        for (const m of matches) {
          if (this.$route.path.includes(m)) {
            matched = true
          }
        }
      }

      return matched
    },
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

    console.log(this.config)
  }
}
</script>

<style scoped>

</style>
