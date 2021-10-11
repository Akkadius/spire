import Vue    from 'vue'
import Router from 'vue-router'
import {ROUTE} from "@/routes";

Vue.use(Router)

export default new Router({
  mode: 'history',
  linkActiveClass: 'active',
  linkExactActiveClass: 'active',
  routes: [
    {
      path: '/',
      component: () => import('./views/layout/Layout.vue'),
      children: [
        {
          path: '/',
          component: () => import('./views/pages/Home.vue'),
          meta: {title: "Home"},
        },
        {
          path: ROUTE.RACE_VIEWER,
          component: () => import('./views/pages/RaceViewer.vue'),
          meta: {title: "Race Viewer"},
        },
        {
          path: ROUTE.ITEM_VIEWER,
          component: () => import('./views/pages/ItemViewer.vue'),
          meta: {title: "Item Viewer"},
        },
        {
          path: ROUTE.ITEM_ICON_VIEWER,
          component: () => import('./views/pages/ItemIconViewer.vue'),
          meta: {title: "Item Icon Viewer"},
        },
        {
          path: ROUTE.SPELL_ANIMATION_VIEWER,
          component: () => import('./views/pages/SpellAnimationViewer.vue'),
          meta: {title: "Spell Animation Viewer"},
        },
        {
          path: '/test/:zone',
          component: () => import('./views/pages/Test.vue'),
          meta: {title: "Test"},
        },
        {
          path: '/components',
          component: () => import('./views/pages/Components.vue'),
          meta: {title: "Component Documentation"},
        },
        {
          path: '/tasks/',
          component: () => import('./views/pages/Tasks.vue'),
          meta: {title: "Task Editor"},
        },
        {
          path: '/tasks/:id',
          component: () => import('./views/pages/Tasks.vue'),
          meta: {title: "Task Editor"},
        },
        {
          path: '/items-test',
          component: () => import('./views/pages/ItemsTest.vue'),
          meta: {title: "Items Test"},
        },
        {
          path: ROUTE.SPELLS_LIST,
          component: () => import('./views/pages/Spells.vue'),
          meta: {title: "Spells Browser"},
        },
        {
          path: ROUTE.ITEMS_LIST,
          component: () => import('./views/pages/Items.vue'),
          meta: {title: "Items Browser"},
        },
        {
          path: '/zones',
          component: () => import('./views/pages/Zones.vue'),
          meta: {title: "Zones"},
        },
        {
          path: '/zone/:zoneId',
          component: () => import('./views/pages/Zone.vue'),
          meta: {title: "Zone"},
        },
        {
          path: '/connections',
          component: () => import('./views/pages/Connections.vue'),
          meta: {title: "Manage Database Connections"},
        },
        {
          path: ROUTE.QUEST_API_EXPLORER,
          component: () => import('./views/pages/QuestApiExplorer/QuestApiExplorer.vue'),
          meta: {title: "Quest API Explorer"},
        },
        {
          path: '/calculators',
          component: () => import('./views/pages/Calculators.vue'),
          meta: {title: "Calculators"},
        },
      ]
    },
    {
      path: '/docs',
      component: () => import('./views/layout/docs/DocLayout.vue'),
      children: [
        {
          path: '/doc/:doc*',
          component: () => import('./views/pages/Doc.vue'),
          meta: {title: "Doc"},
        },
      ]
    },
    {
      path: '/logout',
      component: () => import('./views/pages/Logout.vue')
    },
    {
      path: '/login',
      component: () => import('./views/pages/Login.vue'),
      meta: {title: "Login"},
    },
    {
      path: '/fe-auth-callback',
      component: () => import('./views/pages/AuthCallback.vue'),
      meta: {title: "Authentication Callback"},
    },
  ]
})
