import Vue from 'vue'
import Router from 'vue-router'
import {ROUTE} from "@/routes";
import * as util from "util";

Vue.use(Router)

export default new Router({
  mode: 'history',
  linkActiveClass: 'active',
  linkExactActiveClass: 'active',
  scrollBehavior(to, from, savedPosition) {
    // console.log("[scrollBehavior] to, from, savedPosition", to, from, savedPosition)

    // if title is passed
    if (to.meta && to.meta.title) {
      document.title = "[Spire] " + to.meta.title || "Spire"
    }

    // if link contains a hash target
    if (to.hash) {
      const hash       = to.hash.replace("#", "");
      const hashTarget = document.getElementById(hash)
      if (hashTarget) {
        setTimeout(() => {
          hashTarget.scrollIntoView();
        }, 400);
        return null
      }
    }

    // otherwise resolve the state of location prior
    return new Promise((resolve, reject) => {
      setTimeout(() => {
        if (savedPosition) {
          resolve(savedPosition)
        } else {
          if (Object.keys(to.query).length === 0) {
            resolve({x: 0, y: 0})
          }
        }
      }, 400)
    })
  },
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
          meta: {title: "Spell Animations"},
        },
        {
          path: ROUTE.PLAYER_ANIMATION_VIEWER,
          component: () => import('./views/viewers/PlayerAnimationViewer.vue'),
          meta: {title: "Player Animation Viewer"},
        },
        {
          path: ROUTE.CLIENT_FILES,
          component: () => import('./views/client-files/ClientFiles.vue'),
          meta: {title: "Client Files"},
        },
        {
          path: ROUTE.STRINGS_DATABASE,
          component: () => import('./views/strings-database/StringsDatabase.vue'),
          meta: {title: "Strings Database (dbstr)"},
        },
        {
          path: ROUTE.EMITTER_VIEWER,
          component: () => import('./views/viewers/EmitterViewer.vue'),
          meta: {title: "Emitter Viewer"},
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
          path: ROUTE.TASKS,
          component: () => import('./views/task-editor/TaskEditor.vue'),
          meta: {title: "Task Editor"},
        },
        {
          path: '/tasks/:id',
          component: () => import('./views/task-editor/TaskEditor.vue'),
          meta: {title: "Task Editor"},
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
          path: ROUTE.ZONES,
          component: () => import('./views/zone/Zones.vue'),
          meta: {title: "Zones"},
        },
        {
          path: ROUTE.MERCHANTS,
          component: () => import('./views/zone/Merchants.vue'),
          meta: {title: "Merchants"},
        },
        {
          path: ROUTE.NPCS_EDIT,
          component: () => import('./views/zone/NPCs.vue'),
          meta: {title: "NPC Grid Editor"},
        },
        {
          path: ROUTE.NPC_EDIT,
          component: () => import('./views/zone/NpcEditor.vue'),
          meta: {title: "NPC Editor"},
        },
        {
          path: '/zone/:zone',
          component: () => import('./views/zone/Zone.vue'),
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
      path: '/break/',
      component: () => import('./views/Break.vue')
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
