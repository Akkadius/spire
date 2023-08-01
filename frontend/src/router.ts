import Vue from 'vue'
import Router from 'vue-router'
import {ROUTE} from "@/routes";
import * as util from "util";
import {AppEnv} from "@/app/env/app-env";
import {EventBus} from "@/app/event-bus/event-bus";
import qs from "qs";
import {scrollToHash} from "@/app/utility/scrollToTarget";

Vue.use(Router)

const router = new Router({
  mode: 'history',
  linkActiveClass: 'active',
  linkExactActiveClass: 'active',
  stringifyQuery: query => {
    let result = qs.stringify(query, {format: 'RFC1738'})
    return result ? ('?' + result) : ''
  },
  scrollBehavior(to, from, savedPosition) {
    // console.log("[scrollBehavior] to, from, savedPosition", to, from, savedPosition)

    // if title is passed
    if (to.meta && to.meta.title) {
      document.title = "[Spire] " + to.meta.title || "Spire"
    }

    // if link contains a hash target
    if (to.hash) {
      // for initial loads
      if (to.fullPath === from.fullPath) {
        setTimeout(() => {
          scrollToHash(to)
        }, 1000);
      }
      else if (to.fullPath !== from.fullPath) {
        return null;
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
          path: ROUTE.COFFEE,
          component: () => import('./views/Coffee.vue'),
          meta: {title: "Coffee"},
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
          path: '/test',
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

    // Admin
    {
      path: ROUTE.ADMIN_OCCULUS_REQUIRED,
      component: () => import('./views/admin/OcculusRequire.vue'),
      meta: {title: "Occulus Required"},
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
          path: 'zoneservers/:port/logs',
          component: () => import('./views/admin/ZoneLogs.vue'),
          meta: {title: "Zone Logs"},
        },
        {
          path: ROUTE.ADMIN_CONFIG_DISCORD_CRASH_WEBHOOK,
          component: () => import('./views/admin/configuration/DiscordCrashWebhook.vue'),
          meta: {title: "Discord Crash Webhook", occulus: true},
        },
        {
          path: ROUTE.ADMIN_DISCORD_WEBHOOK_SETTINGS,
          component: () => import('./views/admin/configuration/DiscordWebhooks.vue'),
          meta: {title: "Discord Webhook Settings"},
        },
        {
          path: ROUTE.ADMIN_CONFIG_SERVER_RULES,
          component: () => import('./views/admin/configuration/ServerRules.vue'),
          meta: {title: "Server Rules"},
        },
        {
          path: ROUTE.ADMIN_CONFIG_QUEST_HOT_RELOAD,
          component: () => import('./views/admin/configuration/QuestHotReload.vue'),
          meta: {title: "Quest Hot Reload Settings", occulus: true},
        },
        {
          path: ROUTE.ADMIN_DATABASE_BACKUP,
          component: () => import('./views/admin/tools/DatabaseBackup.vue'),
          meta: {title: "Database Backup"},
        },
        {
          path: ROUTE.ADMIN_RELOAD,
          component: () => import('./views/admin/tools/ServerReload.vue'),
          meta: {title: "Server Reload API"},
        },
        {
          path: ROUTE.ADMIN_CONFIG_MOTD,
          component: () => import('./views/admin/configuration/Motd.vue'),
          meta: {title: "Message of the Day", occulus: true},
        },
        {
          path: ROUTE.ADMIN_LOG_SETTINGS,
          component: () => import('./views/admin/configuration/LogSettings.vue'),
          meta: {title: "Log Settings"},
        },
        {
          path: ROUTE.ADMIN_CONFIG_PLAYER_EVENT_LOGS,
          component: () => import('./views/admin/player-event-logs/PlayerEventLogSettings.vue'),
          meta: {title: "Player Event Log Settings"},
        },
        {
          path: ROUTE.ADMIN_SERVER_CONFIG,
          component: () => import('./views/admin/configuration/ServerConfig.vue'),
          meta: {title: "Server Configuration"},
        },

        // tools
        {
          path: ROUTE.ADMIN_FILE_LOGS,
          component: () => import('./views/admin/FileLogs.vue'),
          meta: {title: "Server Logs"},
        },

        {
          path: ROUTE.ADMIN_BACKUPS,
          component: () => import('./views/admin/Backups.vue'),
          meta: {title: "Manual Backups"},
        },
        {
          path: ROUTE.ADMIN_SERVER_UPDATE,
          component: () => import('./views/admin/server-update/ServerUpdate.vue'),
          meta: {title: "Server Update"},
        },
        {
          path: ROUTE.ADMIN_TOOL_SERVER_QUESTS,
          component: () => import('./views/admin/tools/ServerQuests.vue'),
          meta: {title: "Quests Management", occulus: true},
        },
        {
          path: ROUTE.ADMIN_CLIENT_FILE_DOWNLOADS,
          component: () => import('./views/admin/tools/ClientAssets.vue'),
          meta: {title: "Client Asset Management"},
        },

        {
          path: ROUTE.ADMIN_TOOL_PLAYER_EVENT_LOGS,
          component: () => import('./views/admin/player-event-logs/PlayerEventLogs.vue'),
          meta: {title: "Player Event Logs Explorer"},
        },
        {
          path: ROUTE.ADMIN_TOOL_PLAYER_EVENT_LOGS,
          component: () => import('./views/admin/player-event-logs/PlayerEventLogs.vue'),
          meta: {title: "Player Event Logs Explorer"},
        },
        {
          path: '/admin/ws-poc',
          component: () => import('./views/admin/WebsocketPoc.vue'),
          meta: {title: "Websocket POC"},
        },
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

  if (typeof AppEnv.getEnv() !== 'undefined') {
    AppEnv.routeCheckOcculus(to, router)
    AppEnv.routeCheckSpireInitialized(to, router)
  }

  EventBus.$emit('ROUTE_CHANGE', to);

  next()
})

router.afterEach(async (to, from) => {
  setTimeout(() => {
    const inputs = document.getElementsByTagName("input");
    for (let i = 0; i < inputs.length; i++) {
      if (inputs[i].getAttribute("autofocus")) {
        inputs[i].focus();
        break;
      }
    }
  }, 1)
})

export default router;
