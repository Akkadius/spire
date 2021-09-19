export class ItemStore {
  public static data = {}

  public static setItem(itemId, item) {
    this.data[itemId] = item;
  }

  public static getItems() {
    return this.data;
  }
}
