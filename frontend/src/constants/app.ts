// export convst ASSET_CDN_BASE_URL = "https://rawcdn.githack.com/Akkadius/eq-asset-preview/36903141f431d94766eb83782e9445a856f23514/"

import {LocalSettings} from "@/app/local-settings/localsettings";

// const CDN_VERSION_HASH = '4dfd41553f0f30e4061ca63d7dba589b978c40b8'
// List of different git CDN's
// const ASSET_CDN_BASE_URL_INT = `https://gitcdn.xyz/repo/Akkadius/eq-asset-preview/${CDN_VERSION_HASH}/`
// fast
// `https://ghcdn.rawgit.org/Akkadius/eq-asset-preview/${CDN_VERSION_HASH}/`
// let ASSET_CDN_BASE_URL_INT = `https://rawcdn.githack.com/Akkadius/eq-asset-preview/${CDN_VERSION_HASH}/`
// let ASSET_CDN_BASE_URL_INT = `https://ghcdn.rawgit.org/Akkadius/eq-asset-preview/${CDN_VERSION_HASH}/`
// use local assets if desktop build
// if (process.env.VUE_APP_ASSETS && process.env.VUE_APP_ASSETS.trim() === "local") {
//   ASSET_CDN_BASE_URL_INT = "/eq-asset-preview-master/";
// }

// @akkadius - I was using a CDN to host these files and decided for simplicty to require them as part
// of how the application gets bootstrapped and bundled
// These assets get pulled down during install and during release build but do not get checked in "ever"
let ASSET_CDN_BASE_URL_INT = "/eq-asset-preview-master/";

export const App = {
  ASSET_CDN_BASE_URL: ASSET_CDN_BASE_URL_INT,
  ASSET_ITEM_ICON_BASE_URL: ASSET_CDN_BASE_URL_INT + 'assets/item_icons/',
  ASSET_MONOGRAMS_BASE_URL: ASSET_CDN_BASE_URL_INT + 'assets/monograms/',
  ASSET_NPC_MODEL_BASE_URL: ASSET_CDN_BASE_URL_INT + 'assets/npc_models/',
  ASSET_OBJECTS_BASE_URL: ASSET_CDN_BASE_URL_INT + 'assets/objects/',
  ASSET_SPELL_ICONS_BASE_URL: ASSET_CDN_BASE_URL_INT + 'assets/spell_icons/',
  ASSET_SPELL_ANIMATIONS: ASSET_CDN_BASE_URL_INT + 'assets/spell_animations/',
  ASSET_EMITTER_CLIPS: ASSET_CDN_BASE_URL_INT + 'assets/emitters/',
  ASSET_EXPANSION_ICON_SMALL_URL: ASSET_CDN_BASE_URL_INT + 'assets/expansion-icons-small/',
  ASSET_WALLPAPER_URL: ASSET_CDN_BASE_URL_INT + 'assets/wallpaper/',
  ASSET_INVENTORY_SLOT_URL: ASSET_CDN_BASE_URL_INT + 'assets/inventory/',
  ASSET_SPRITE_ITEM_ICONS_URL: ASSET_CDN_BASE_URL_INT + 'assets/sprites/item-icons.css',
  BACKEND_BASE_URL: (process.env.VUE_APP_BACKEND_BASE_URL ? process.env.VUE_APP_BACKEND_BASE_URL : window.location.origin),
  DEBUG: LocalSettings.get("debug-mode") === "true",
  // DEBUG: true,
}
