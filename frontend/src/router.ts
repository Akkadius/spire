import Vue    from 'vue'
import Router from 'vue-router'
import {ROUTE} from "@/routes";
import * as util from "util";

Vue.use(Router)

export default new Router({
  mode: 'history',
  linkActiveClass: 'active',
  linkExactActiveClass: 'active',
  routes: [
    {
      path: '/',
      component: () => import('./components/layout/MainLayout.vue'),
      children: [
        {
          path: '/',
          component: () => import('./views/Home.vue'),
          meta: {title: "Home"},
        },
        {
          path: ROUTE.RACE_VIEWER,
          component: () => import('./views/viewers/RaceViewer.vue'),
          meta: {title: "Race Viewer"},
        },
        {
          path: ROUTE.ITEM_VIEWER,
          component: () => import('./views/viewers/ItemViewer.vue'),
          meta: {title: "Item Viewer"},
        },
        {
          path: ROUTE.ITEM_ICON_VIEWER,
          component: () => import('./views/viewers/ItemIconViewer.vue'),
          meta: {title: "Item Icon Viewer"},
        },
        {
          path: ROUTE.SPELL_ANIMATION_VIEWER,
          component: () => import('./views/viewers/SpellAnimationViewer.vue'),
          meta: {title: "Spell Animation Viewer"},
        },
        {
          path: '/test/:zone',
          component: () => import('./views/Test.vue'),
          meta: {title: "Test"},
        },
        {
          path: '/components',
          component: () => import('./views/Components.vue'),
          meta: {title: "Component Documentation"},
        },
        {
          path: '/tasks/',
          component: () => import('./views/Tasks.vue'),
          meta: {title: "Task Editor"},
        },
        {
          path: '/tasks/:id',
          component: () => import('./views/Tasks.vue'),
          meta: {title: "Task Editor"},
        },
        {
          path: '/items-test',
          component: () => import('./views/ItemsTest.vue'),
          meta: {title: "Items Test"},
        },
        {
          path: ROUTE.SPELLS_LIST,
          component: () => import('./views/Spells.vue'),
          meta: {title: "Spells Browser"},
        },
        {
          path: util.format(ROUTE.SPELL_EDIT, ":id"),
          component: () => import('./views/spell-editor/SpellEditor.vue'),
          meta: {title: "Spell Edit"},
        },
        {
          path: ROUTE.ITEMS_LIST,
          component: () => import('./views/items/Items.vue'),
          meta: {title: "Items Browser"},
        },
        {
          path: util.format(ROUTE.ITEM_EDIT, ":id"),
          component: () => import('./views/item-editor/ItemEditor.vue'),
          meta: {title: "Item Edit"},
        },
        {
          path: '/zones',
          component: () => import('./views/Zones.vue'),
          meta: {title: "Zones"},
        },
        {
          path: '/zone/:zoneId',
          component: () => import('./views/Zone.vue'),
          meta: {title: "Zone"},
        },
        {
          path: '/connections',
          component: () => import('./views/Connections.vue'),
          meta: {title: "Manage Database Connections"},
        },
        {
          path: ROUTE.QUEST_API_EXPLORER,
          component: () => import('./views/quest-api-explorer/QuestApiExplorer.vue'),
          meta: {title: "Quest API Explorer"},
        },
        {
          path: '/calculators',
          component: () => import('./views/Calculators.vue'),
          meta: {title: "Calculators"},
        },
      ]
    },
    {
      path: '/editor',
      component: () => import('./components/layout/QuestEditorLayout.vue'),
      meta: {title: "Editor"},
      children: [
        {
          path: '/',
          component: () => import('./views/quest-editor/QuestEditor.vue'),
          meta: {title: "Editor"},
        },
      ]
    },
    {
      path: '/docs',
      component: () => import('./components/layout/docs/DocLayout.vue'),
      children: [
        {
          path: '/doc/:doc*',
          component: () => import('./views/Doc.vue'),
          meta: {title: "Doc"},
        },
      ]
    },
    {
      path: '/logout',
      component: () => import('./views/Logout.vue')
    },
    {
      path: '/login',
      component: () => import('./views/Login.vue'),
      meta: {title: "Login"},
    },
    {
      path: '/fe-auth-callback',
      component: () => import('./views/AuthCallback.vue'),
      meta: {title: "Authentication Callback"},
    },
  ]
})
