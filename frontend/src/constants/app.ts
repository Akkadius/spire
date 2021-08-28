// export const ASSET_CDN_BASE_URL = "https://rawcdn.githack.com/Akkadius/eq-asset-preview/36903141f431d94766eb83782e9445a856f23514/"

import LocalSettings from "@/app/local-settings/localsettings";

const CDN_VERSION_HASH = 'be32443283a960c9dfb6a6425fce26cddf36adb8'

// List of different git CDN's
// const ASSET_CDN_BASE_URL_INT = `https://gitcdn.xyz/repo/Akkadius/eq-asset-preview/${CDN_VERSION_HASH}/`

// fast
const ASSET_CDN_BASE_URL_INT = `https://rawcdn.githack.com/Akkadius/eq-asset-preview/${CDN_VERSION_HASH}/`
// const ASSET_CDN_BASE_URL_INT = `https://ghcdn.rawgit.org/Akkadius/eq-asset-preview/${CDN_VERSION_HASH}/`

export const App = {
  ASSET_CDN_BASE_URL: ASSET_CDN_BASE_URL_INT,
  ASSET_ITEM_ICON_BASE_URL: ASSET_CDN_BASE_URL_INT + 'assets/item_icons/',
  ASSET_MONOGRAMS_BASE_URL: ASSET_CDN_BASE_URL_INT + 'assets/monograms/',
  ASSET_NPC_MODEL_BASE_URL: ASSET_CDN_BASE_URL_INT + 'assets/npc_models/',
  ASSET_OBJECTS_BASE_URL: ASSET_CDN_BASE_URL_INT + 'assets/objects/',
  ASSET_SPELL_ICONS_BASE_URL: ASSET_CDN_BASE_URL_INT + 'assets/spell_icons/',
  ASSET_SPELL_ANIMATIONS: ASSET_CDN_BASE_URL_INT + 'assets/spell_animations/',
  ASSET_EXPANSION_ICON_SMALL_URL: ASSET_CDN_BASE_URL_INT + 'assets/expansion-icons-small/',
  BACKEND_BASE_URL: (process.env.VUE_APP_BACKEND_BASE_URL ? process.env.VUE_APP_BACKEND_BASE_URL : window.location.origin),
  DEBUG: LocalSettings.get("debug-mode") === "true",
  // DEBUG: true,
}
