import {EXPANSION_ICONS_SMALL} from "@/app/constants/eq-expansion-icons";
import {App}                   from "@/constants/app";
import {EXPANSION_NAMES}       from "@/app/constants/eq-expansions";

export default class Expansions {
  static getExpansionIconUrlSmall(expansionId) {
    if (EXPANSION_ICONS_SMALL[expansionId]) {
      return App.ASSET_EXPANSION_ICON_SMALL_URL + EXPANSION_ICONS_SMALL[expansionId]
    }

    // return transparent base64 encoded image if nothing found
    return 'data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7'
  }
  static getExpansionName(expansionId) {
    if (EXPANSION_NAMES[expansionId]) {
      return EXPANSION_NAMES[expansionId]
    }

    // return unknown expansion if not found
    return 'Unknown'
  }
}
