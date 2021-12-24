import {ItemApi} from "@/app/api";
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

  public static getBasicStatFields() {
    return {
      "Strength": {stat: "astr", heroic: "heroic_str"},
      "Stamina": {stat: "asta", heroic: "heroic_sta"},
      "Intelligence": {stat: "aint", heroic: "heroic_int"},
      "Wisdom": {stat: "awis", heroic: "heroic_wis"},
      "Agility": {stat: "aagi", heroic: "heroic_agi"},
      "Dexterity": {stat: "adex", heroic: "heroic_dex"},
      "Charisma": {stat: "acha", heroic: "heroic_cha"},
    }
  };

  public static getResistFields() {
    return {
      "Magic Resist": {stat: "mr", heroic: "heroic_mr"},
      "Fire Resists": {stat: "fr", heroic: "heroic_fr"},
      "Cold Resist": {stat: "cr", heroic: "heroic_cr"},
      "Disease Resist": {stat: "dr", heroic: "heroic_dr"},
      "Poison Resist": {stat: "pr", heroic: "heroic_pr"},
      "Corruption": {stat: "svcorruption", heroic: "heroic_svcorrup"}
    }
  };

  public static getBasicStatAndResistFields() {
    let fields = {};

    Object.assign(fields, this.getBasicStatFields());
    Object.assign(fields, this.getResistFields());

    return fields;
  }

  public static getMod3Fields() {
    return {
      "Attack": "attack",
      "HP Regen": "regen",
      "Mana Regen": "manaregen",
      "Endurance Regen": "enduranceregen",
      "Accuracy": "accuracy",
      "Avoidance": "avoidance",
      "Clairvoyance": "clairvoyance",
      "Combat Effects": "combateffects",
      "Damage Shield": "damageshield",
      "Damage Shield Mitigation": "dsmitigation",
      "DoT Shielding": "dotshielding",
      "Heal Amount": "healamt",
      "Shielding": "shielding",
      "Spell Shielding": "spellshield",
      "Strikethrough": "strikethrough",
      "Stun Resist": "stunresist",
    };
  }
}
