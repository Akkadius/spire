import Vue          from 'vue'
import router       from './router'
import App          from './App.vue'
import store        from './store'
import BootstrapVue from 'bootstrap-vue'

// Bootstrap
// import 'bootstrap/dist/css/bootstrap.css'
import 'bootstrap-vue/dist/bootstrap-vue.css'
import '@/assets/css/custom.css'
// iziToast
import 'izitoast/dist/css/iziToast.css'
import VueIziToast  from 'vue-izitoast'
// Dashkit
import './assets/css/theme.min.css'
import './assets/fonts/feather/feather.min.css'
import './assets/css/global.css'

// EQ Assets - These should be moved into the assets themselves
import '@/components/eq-ui/styles/eq-ui.css'

// font awesome
import '@/assets/font-awesome-4.7.0/css/font-awesome.min.css'

// Lazyload
import VueLazyload      from 'vue-lazyload'
// rpg awesome icons
import "rpg-awesome/css/rpg-awesome.min.css";

// highlight js css
import "@/assets/css/highlight-js-vs2015.css";
// Vue Form Generator
import VueFormGenerator from 'vue-form-generator'
import 'vue-form-generator/dist/vfg.css'


// Ag grid enterprisec
import 'ag-grid-enterprise';

Vue.use(BootstrapVue)

Vue.use(VueIziToast)

Vue.use(VueLazyload)

// Custom components
Vue.component('field-eq-text-input', () => import('@/components/form-components/FieldEqTextInput.vue'));
Vue.component('field-eq-text-area', () => import('@/components/form-components/FieldEqTextArea.vue'));
Vue.component('field-task-type', () => import('@/components/form-components/FieldTaskType.vue'));
Vue.component('field-task-duration-code', () => import('@/components/form-components/FieldTaskDurationCode.vue'));
Vue.component('field-task-duration', () => import('@/components/form-components/FieldTaskDuration.vue'));
Vue.component('field-eq-yes-no', () => import('@/components/form-components/FieldEqYesNo.vue'));
Vue.component('field-eq-cash', () => import('@/components/form-components/FieldEqCash.vue'));
Vue.component('eq-item-preview', () => import('@/components/eq-ui/EQItemPreview.vue'));


// vue form generator
Vue.use(VueFormGenerator)
Vue.component('vue-form-generator', VueFormGenerator.component);

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
