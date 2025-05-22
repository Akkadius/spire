import Vue from 'vue'
import App from './App.vue'

import VueRouter from 'vue-router'

const originalPush = VueRouter.prototype.push

const VUE3_HANDLED_ROUTES = [
    // '/items'
]

function isVue3Route(location) {
    return VUE3_HANDLED_ROUTES.includes(typeof location === 'string' ? location : location.path)
}

VueRouter.prototype.push = function push(location, onResolve, onReject) {
    const current = this.currentRoute
    const target = this.resolve(location).route

    if (target.fullPath === current.fullPath) {
        if (onResolve || onReject) return
        return Promise.resolve(current)
    }

    if (onResolve || onReject) {
        return originalPush.call(this, location, onResolve, onReject)
    }

    return originalPush.call(this, location).catch(err => {
        if (err && err._isRouter && err.type === 8) {
            // Suppress redundant nav errors
            return current
        }
        throw err
    })
}

import router from './router'

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
import BootstrapVue from "bootstrap-vue";

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

// Vue 3 route paths

router.beforeEach((to, from, next) => {
    if (isVue3Route(to.fullPath)) {
        window.location.href = to.fullPath
        return // stop Vue 2 navigation
    }

    next()
})


let vue2Instance = null

window.Vue2App = {
    mountLegacyApp(el, initialPath) {
        if (initialPath) {
            router.push(initialPath)
        }

        vue2Instance = new Vue({
            router,
            render: h => h(App)
        }).$mount(el)

        return vue2Instance
    },
}
