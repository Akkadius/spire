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
import '@/components/eq-ui/styles/eq-ui-buttons.css'

// global custom
import './assets/css/global.css'

// font awesome
import 'fontawesome-4.7'

// rpg awesome icons
import "rpg-awesome/css/rpg-awesome.min.css";

import 'highlight.js/styles/tomorrow-night-bright.css';

import hljs from 'highlight.js/lib/highlight';
import json from 'highlight.js/lib/languages/json.js';
hljs.registerLanguage('json', json);

import "toastify-js/src/toastify.css"

/* spritesheet assets */
// import "../public/eq-asset-preview-master/assets/sprites/item-icons.css";
// import "../public/eq-asset-preview-master/assets/sprites/objects.css";
// import "../public/eq-asset-preview-master/assets/sprites/race-models.css";

import "@exuanbo/file-icons-js/dist/css/file-icons.min.css"

Vue.use(BootstrapVue)

import 'leaflet/dist/leaflet.css';
import { Icon } from 'leaflet';

// @ts-ignore
delete Icon.Default.prototype._getIconUrl;
Icon.Default.mergeOptions({
  iconRetinaUrl: require('leaflet/dist/images/marker-icon-2x.png'),
  iconUrl: require('leaflet/dist/images/marker-icon.png'),
  shadowUrl: require('leaflet/dist/images/marker-shadow.png'),
});

/**
 * App loader
 */
Vue.component('app-loader', () => import('@/components/LoaderComponent.vue'));

Vue.config.productionTip = false

Vue.config.errorHandler = function (err, vm, info) {
  console.error(`Error in ${info}:`, err);
};


new Vue({
  router,
  store,
  render: h => h(App)
}).$mount('#app')
