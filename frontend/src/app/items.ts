import {ItemApi}        from "@/app/api";
import {SpireApiClient} from "@/app/api/spire-api-client";

export class Items {
  public static items = {}

  public static setItem(itemId, item) {
    this.items[itemId] = item;
  }

  public static async getItem(itemId) {
    if (itemId === 0) {
      return {}
    }

    if (this.items[itemId]) {
      return this.items[itemId]
    }

    const api    = (new ItemApi(SpireApiClient.getOpenApiConfig()))
    const result = await api.getItem({id: itemId})
    if (result.status === 200 && result.data) {
      this.setItem(itemId, result.data);
      return result.data
    }

    return {}
  }

  public static getItems() {
    return this.items;
  }
}
