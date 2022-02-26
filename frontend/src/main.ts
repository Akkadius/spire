import Vue          from 'vue'
import router       from './router'
import App          from './App.vue'
import store        from './store'
import BootstrapVue from 'bootstrap-vue'

// Bootstrap
// import 'bootstrap/dist/css/bootstrap.css'
import 'bootstrap-vue/dist/bootstrap-vue.css'
import '@/assets/css/custom.css'

// Dashkit
import './assets/css/theme.min.css'
import './assets/fonts/feather/feather.min.css'

// EQ Assets - These should be moved into the assets themselves
import '@/components/eq-ui/styles/eq-ui.css'

// global custom
import './assets/css/global.css'

// font awesome
import 'fontawesome-4.7'

// rpg awesome icons
import "rpg-awesome/css/rpg-awesome.min.css";

// highlight js css
import "@/assets/css/highlight-js-vs2015.css";

/* ag-grid */
import "../node_modules/ag-grid-community/dist/styles/ag-grid.css";
import "../node_modules/ag-grid-community/dist/styles/ag-theme-balham-dark.css";

/* spritesheet assets */
// import "../public/eq-asset-preview-master/assets/sprites/item-icons.css";
// import "../public/eq-asset-preview-master/assets/sprites/objects.css";
// import "../public/eq-asset-preview-master/assets/sprites/race-models.css";

// Ag grid enterprisec
import 'ag-grid-enterprise';

// vue-tree

import 'sl-vue-tree/src/sl-vue-tree-dark.css'
import 'sl-vue-tree/src/sl-vue-tree.js'

// import {BVConfigPlugin} from 'bootstrap-vue';

// Vue.use(BVConfigPlugin, {
//   BTooltip: {
//     delay: {
//       show: 0,
//       hide: 0,
//     },
//   },
// });

import "@exuanbo/file-icons-js/dist/css/file-icons.min.css"

Vue.use(BootstrapVue)

/**
 * App loader
 */
Vue.component('app-loader', () => import('@/components/LoaderComponent.vue'));

Vue.config.productionTip = false

new Vue({
  router,
  store,
  render: h => h(App)
}).$mount('#app')
