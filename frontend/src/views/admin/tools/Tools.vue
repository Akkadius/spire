<template>
  <div>
    <div class="card">
      <div class="card-body pt-0 mb-1">
        <ul class="nav nav-tabs nav-overflow header-tabs">
          <li class="nav-item">
            <router-link
              :to="'/admin/tools/logs'"
              active-class=""
              exact-active-class="active"
              :class="'nav-link ' + (hasRoute('/admin/tools/logs') ? 'active' : '')"
            >
              <i class="fe fe-book mr-2"></i>
              Logs
            </router-link>
          </li>

          <li class="nav-item">
            <router-link
              :to="'/admin/tools/server-code'"
              active-class=""
              exact-active-class="active"
              class="nav-link"
            >
              <i class="fa fa-code-fork mr-2"></i>
              Server Code (Linux)
            </router-link>
          </li>


          <li class="nav-item">
            <router-link
              :to="'/admin/tools/server-quests'"
              active-class=""
              exact-active-class="active"
              class="nav-link"
            >
              <i class="fa fa-code-fork mr-2"></i>
              Quests
            </router-link>
          </li>

          <li class="nav-item">
            <router-link
              :to="'/admin/tools/backups'"
              active-class=""
              exact-active-class="active"
              class="nav-link"
            >
              <i class="fe fe-download-cloud mr-2"></i>
              Manual Backup
            </router-link>
          </li>

          <li class="nav-item">
            <router-link
              :to="'/admin/tools/client-assets'"
              active-class=""
              exact-active-class="active"
              class="nav-link"
            >
              <span class="icon mr-3"><i class="fa fa-th"></i></span>
              Client Assets
            </router-link>
          </li>

        </ul>
      </div>
    </div>

    <div class="row">
      <div class="col-lg-12">
        <router-view></router-view>
      </div>
    </div>

  </div>
</template>

<script>
export default {
  name: 'Tools',
  data() {
    return {}
  },
  methods: {
    hasRoute: function (partial) {
      return (this.$route.path.indexOf(partial) > -1)
    },
  }
}
</script>
