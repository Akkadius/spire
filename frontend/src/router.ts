import Vue from 'vue'
import Router from 'vue-router'
import {ROUTE} from "@/routes";
import * as util from "util";
import {AppEnv} from "@/app/env/app-env";
import {EventBus} from "@/app/event-bus/event-bus";

Vue.use(Router)

const router = new Router({
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
          component: () => import('./views/asset-viewers/RaceViewer.vue'),
          meta: {title: "Race Viewer"},
        },
        {
          path: ROUTE.ITEM_VIEWER,
          component: () => import('./views/asset-viewers/ItemViewer.vue'),
          meta: {title: "Item Viewer"},
        },
        {
          path: ROUTE.ITEM_ICON_VIEWER,
          component: () => import('./views/asset-viewers/ItemIconViewer.vue'),
          meta: {title: "Item Icon Viewer"},
        },
        {
          path: ROUTE.SPELL_ANIMATION_VIEWER,
          component: () => import('./views/asset-viewers/SpellAnimationViewer.vue'),
          meta: {title: "Spell Animations"},
        },
        {
          path: ROUTE.PLAYER_ANIMATION_VIEWER,
          component: () => import('./views/asset-viewers/PlayerAnimationViewer.vue'),
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
          component: () => import('./views/asset-viewers/EmitterViewer.vue'),
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
          component: () => import('./views/tasks/TaskEditor.vue'),
          meta: {title: "Task Editor"},
        },
        {
          path: '/tasks/:id',
          component: () => import('./views/tasks/TaskEditor.vue'),
          meta: {title: "Task Editor"},
        },
        {
          path: ROUTE.SPELLS_LIST,
          component: () => import('./views/spells/Spells.vue'),
          meta: {title: "Spells Browser"},
        },
        {
          path: util.format(ROUTE.SPELL_EDIT, ":id"),
          component: () => import('./views/spells/SpellEditor.vue'),
          meta: {title: "Spell Edit"},
        },
        {
          path: ROUTE.ITEMS_LIST,
          component: () => import('./views/items/Items.vue'),
          meta: {title: "Items Browser"},
        },
        {
          path: util.format(ROUTE.ITEM_EDIT, ":id"),
          component: () => import('./views/items/ItemEditor.vue'),
          meta: {title: "Item Edit"},
        },
        {
          path: ROUTE.LOOT,
          component: () => import('./views/loot/Loot.vue'),
          meta: {title: "Loot Edit"},
        },
        {
          path: ROUTE.ZONES,
          component: () => import('./views/zone/Zones.vue'),
          meta: {title: "Zones"},
        },
        {
          path: ROUTE.RELEASES,
          component: () => import('./views/server-developer/Releases.vue'),
          meta: {title: "Releases"},
        },
        {
          path: util.format(ROUTE.RELEASE, ":version"),
          component: () => import('./views/server-developer/Release.vue'),
          meta: {title: "Release"},
        },
        {
          path: ROUTE.CHANGELOG,
          component: () => import('./views/server-developer/Changelog.vue'),
          meta: {title: "Changelog"},
        },
        {
          path: ROUTE.MERCHANTS,
          component: () => import('./views/merchants/Merchants.vue'),
          meta: {title: "Merchants"},
        },
        {
          path: util.format(ROUTE.MERCHANT_EDIT, ":id"),
          component: () => import('./views/merchants/MerchantEdit.vue'),
          meta: {title: "Merchant Edit"},
        },
        {
          path: ROUTE.NPCS_EDIT,
          component: () => import('./views/npcs/NPCs.vue'),
          meta: {title: "NPC Grid Editor"},
        },
        {
          path: ROUTE.NPC_EDIT,
          component: () => import('./views/npcs/NpcEditor.vue'),
          meta: {title: "NPC Editor"},
        },
        {
          path: ROUTE.NPC_EMOTES_EDIT,
          component: () => import('./views/npcs/NpcEmotesEditor.vue'),
          meta: {title: "NPC Emotes Editor"},
        },
        {
          path: ROUTE.NPC_SPELLS_EDIT,
          component: () => import('./views/npcs/NpcSpellsEditor.vue'),
          meta: {title: "NPC Spells Editor"},
        },
        {
          path: util.format(ROUTE.NPC_SPELL_EDIT, ":id"),
          component: () => import('./views/npcs/NpcSpellListEditor.vue'),
          meta: {title: "NPC Spells List Editor"},
        },
        {
          path: '/zone/:zone',
          component: () => import('./views/zone/Zone.vue'),
          meta: {title: "Zone"},
        },
        {
          path: '/connections',
          component: () => import('./views/connections/Connections.vue'),
          meta: {title: "Manage Database Connections"},
        },
        {
          path: ROUTE.USER_MANAGEMENT,
          component: () => import('./views/user/UserManagement.vue'),
          meta: {title: "Manage Spire Users"},
        },
        {
          path: util.format(ROUTE.DATABASE_CONNECTION_AUDIT_LOG, ":connection"),
          component: () => import('./views/connections/AuditLog.vue'),
          meta: {title: "Audit Log"},
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
        {
          path: ROUTE.API_MODEL_RELATIONSHIP_EXPLORER,
          component: () => import('./views/api/ModelRelationshipExplorer.vue'),
          meta: {title: "API Model Relationship Explorer"},
        },
        {
          path: '/expansions',
          component: () => import('./views/Expansion.vue'),
          meta: {title: "Expansions"},
        },
        {
          path: ROUTE.BOT_SPELLS_EDIT,
          component: () => import('./views/bots/BotSpellsEditor.vue'),
          meta: {title: "Bot Spells List Editor"},
        },
        {
          path: util.format(ROUTE.BOT_SPELL_EDIT, ":id"),
          component: () => import('./views/bots/BotSpellListEditor.vue'),
          meta: {title: "Bot Spells List Editor"},
        },
      ]
    },
    {
      path: ROUTE.ADMIN_ROOT,
      component: () => import('./views/admin/layout/AdminLayout.vue'),
      children: [
        {
          path: '/',
          component: () => import('./views/admin/Dashboard.vue'),
          meta: {title: "Dashboard"},
        },
        {
          path: 'players-online',
          component: () => import('./views/admin/PlayersOnline.vue'),
          meta: {title: "Players Online"},
        },
        {
          path: 'zones',
          component: () => import('./views/admin/ZoneServers.vue'),
          meta: {title: "Zone Servers"},
        },
        {
          path: 'zoneservers/:port/netstats',
          component: () => import('./views/admin/NetStats.vue'),
          meta: {title: "Netstats"},
        },
        {
          path: 'configuration',
          meta: {title: "Configuration"},
          component: () => import('./views/admin/configuration/Configuration.vue'),
          children: [
            {
              path: '/',
              component: () => import('./views/admin/configuration/WorldServerSettings.vue'),
              meta: {title: "World Server Settings"},
            },
            {
              path: 'settings/zoneserver',
              component: () => import('./views/admin/configuration/ZoneServerSettings.vue'),
              meta: {title: "Zone Server Settings"},
            },
            {
              path: 'settings/ucs',
              component: () => import('./views/admin/configuration/UcsSettings.vue'),
              meta: {title: "UCS Config"},
            },
            {
              path: 'settings/discord',
              component: () => import('./views/admin/configuration/DiscordSettings.vue'),
              meta: {title: "Discord Config"},
            },
            {
              path: 'settings/database',
              component: () => import('./views/admin/configuration/Database.vue'),
              meta: {title: "Database Config"},
            },
            {
              path: 'server-rules',
              component: () => import('./views/admin/configuration/ServerRules.vue'),
              meta: {title: "Server Rules"},
            },
            {
              path: 'motd',
              component: () => import('./views/admin/configuration/Motd.vue'),
              meta: {title: "Message of the Day"},
            }
          ]
        },
        {
          path: 'tools',
          meta: {title: "Tools"},
          component: () => import('./views/admin/tools/Tools.vue'),
          children: [
            {
              path: 'logs',
              component: () => import('./views/admin/tools/Logs.vue'),
              meta: {title: "Server Logs"},
            },
            {
              path: 'backups',
              component: () => import('./views/admin/tools/Backups.vue'),
              meta: {title: "Backups"},
            },
            {
              path: 'server-code',
              component: () => import('./views/admin/tools/ServerCode.vue'),
              meta: {title: "Code Management"},
            },
            {
              path: 'server-quests',
              component: () => import('./views/admin/tools/ServerQuests.vue'),
              meta: {title: "Quests Management"},
            },
            {
              path: 'client-assets',
              component: () => import('./views/admin/tools/ClientAssets.vue'),
              meta: {title: "Client Asset Management"},
            },
          ]
        }
      ]
    },

    {
      path: '/break/',
      component: () => import('./views/Break.vue')
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
      path: ROUTE.SPIRE_INITIALIZE,
      component: () => import('./views/SpireInitialize.vue'),
      meta: {title: "Spire Setup"},
    },
    {
      path: '/fe-auth-callback',
      component: () => import('./views/AuthCallback.vue'),
      meta: {title: "Authentication Callback"},
    },
  ]
})

// router middleware
router.beforeEach(async (to, from, next) => {
  // console.log("Middleware from [%s] to [%s]", from, to)
  // console.log(from)
  // console.log(to)
  // console.log("Is spire initialized", AppEnv.isSpireInitialized())

  // ensure appenv has been initialized first
  if (typeof AppEnv.isSpireInitialized() !== 'undefined') {
    // capture user under conditions
    if (!AppEnv.isSpireInitialized() && to.fullPath !== ROUTE.SPIRE_INITIALIZE) {
      console.log("hello")
      // re-route to spire setup if not setup yet
      router.push(ROUTE.SPIRE_INITIALIZE).catch((e) => {
      })
    } else if (AppEnv.isSpireInitialized() && to.fullPath === ROUTE.SPIRE_INITIALIZE) {
      router.push(ROUTE.HOME).catch((e) => {
      })
    }
  }

  EventBus.$emit('ROUTE_CHANGE', to);

  next()
})

export default router;
