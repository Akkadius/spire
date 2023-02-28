<template>
  <div>
    <eq-window>
      <ninja-keys
        class="dark ninja-icon"
        ref="ninjaKeys"
        :placeholder="placeholder"
        @selected="selected"
        :noAutoLoadMdIcons="true"
        @change="change"
      ></ninja-keys>
    </eq-window>
  </div>
</template>

<script>
import EqWindow from "@/components/eq-ui/EQWindow.vue";
import "ninja-keys";
import {ROUTE}  from "@/routes";

export default {
  name: "Test",
  components: {
    EqWindow,
  },
  data() {
    return {
      placeholder: "Placeholder prop test",

      adminNavs: [
        { label: "Backups", labelIcon: "fa fa-download mr-1", to: ROUTE.ADMIN_BACKUPS },
        { label: "Client Files", labelIcon: "fa fa-download mr-1", to: ROUTE.ADMIN_CLIENT_FILE_DOWNLOADS },
        {
          label: "Configuration",
          labelIcon: "fa fa-cog mr-1",
          routePrefixMatch: "admin/configuration",
          navs: [
            { title: "Server Config", to: ROUTE.ADMIN_SERVER_CONFIG, icon: "fa fa-cog mr-1" },
            { title: "Crash Webhooks", to: ROUTE.ADMIN_CONFIG_DISCORD_CRASH_WEBHOOK, icon: "ra ra-fire mr-1", isOcculus: true },
            { title: "MOTD", to: ROUTE.ADMIN_CONFIG_MOTD, icon: "ra ra-wooden-sign mr-1" },
            { title: "Quest Hot Reload", to: ROUTE.ADMIN_CONFIG_QUEST_HOT_RELOAD, icon: "ra ra-alien-fire mr-1", isOcculus: true },
            { title: "Server Rules", to: ROUTE.ADMIN_CONFIG_SERVER_RULES, icon: "ra ra-interdiction mr-1" },
            { title: "UCS", to: ROUTE.ADMIN_SERVER_CONFIG + '?s=UCS', icon: "ra ra-speech-bubbles mr-1", exact: true },
            { title: "World Server", to: ROUTE.ADMIN_SERVER_CONFIG + '?s=World+Server', icon: "ra ra-double-team mr-1", exact: true },
            { title: "Zone Server", to: ROUTE.ADMIN_SERVER_CONFIG + '?s=Zone+Server', icon: "ra ra-player mr-1", exact: true },
          ]
        },
        {
          label: "Database", labelIcon: "fa fa-database mr-1", routePrefixMatch: "admin/database",
          navs: [
            { title: "Database Config", to: ROUTE.ADMIN_SERVER_CONFIG + '?s=Database', icon: "fa fa-cog mr-1", exact: true},
            { title: "Database Backups", to: ROUTE.ADMIN_DATABASE_BACKUP, icon: "fa fa-download mr-1" },
          ]
        },
        {
          label: "Logs",
          labelIcon: "ra ra-telescope mr-1",
          routePrefixMatch: "admin/tools/player",
          navs: [
            { title: "Discord Webhooks", to: ROUTE.ADMIN_DISCORD_WEBHOOK_SETTINGS, icon: "fa fa-cog mr-1" },
            { title: "File Logs", to: ROUTE.ADMIN_FILE_LOGS, icon: "fe fe-book mr-1" },
            { title: "Log Settings", to: ROUTE.ADMIN_LOG_SETTINGS, icon: "ra ra-scroll-unfurled mr-1" },
            { title: "Player Event Settings", to: ROUTE.ADMIN_CONFIG_PLAYER_EVENT_LOGS, icon: "fa fa-cog mr-1" },
            { title: "Player Event Log Viewer", to: ROUTE.ADMIN_TOOL_PLAYER_EVENT_LOGS, icon: "ra ra-telescope mr-1" },
          ]
        },
        { label: "Quests", labelIcon: "fa fa-code-fork mr-1", to: ROUTE.ADMIN_TOOL_SERVER_QUESTS, isOcculus: true },
        { label: "Reloading (Global)", labelIcon: "fa fa-refresh mr-1", routePrefixMatch: "admin/tools/player", to: ROUTE.ADMIN_RELOAD },
        { label: "Server Update", labelIcon: "fa fa-upload mr-1", to: ROUTE.ADMIN_SERVER_UPDATE },

      ],

      hotkeys: [
        {
          id: "Home",
          title: "Open Home",
          hotkey: "cmd+h",
          handler: () => {
            console.log("navigation to home");
          },
        },
        {
          id: "Open Projects",
          icon: "<i class='ra ra-compass ninja-icon mr-1'></i>",
          title: "Open Projects",
          hotkey: "cmd+p",
          handler: () => {
            console.log("navigation to projects");
          },
        },
        {
          id: "Theme",
          title: "Change theme...",
          children: [
            {
              id: "Light Theme",
              title: "Change theme to Light",
              mdIcon: "light_mode",
              handler: () => {
                console.log("theme light");
              },
            },
            {
              id: "Dark Theme",
              title: "Change theme to Dark",
              mdIcon: "dark_mode",
              keywords: "lol",
              handler: () => {
                console.log("theme dark");
              },
            },
          ],
        },
      ],
    };
  },
  methods: {
    selected(event) {
      console.log("selected", event.detail);
    },
    change(event) {
      console.log("change", event.detail);
    },
  },
  mounted() {
    let keys = []
    for (let n of this.adminNavs) {
      if (n.label && n.to) {
        keys.push({
          id: n.label,
          title: n.label,
          handler: () => {
            this.$router.push(n.to).catch((e) => {
            })
          }
        })
      }
    }




    this.$refs.ninjaKeys.data = keys;
  },
};
</script>
