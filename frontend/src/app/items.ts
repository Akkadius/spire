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

    const api = (new ItemApi(SpireApiClient.getOpenApiConfig()))
    try {
      const result = await api.getItem({id: itemId})
      if (result.status === 200 && result.data) {
        this.setItem(itemId, result.data);
        return result.data
      }
    } catch (err) {
      console.log("items.ts %s", err)
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
      "Accuracy": "accuracy",
      "Attack": "attack",
      "Avoidance": "avoidance",
      "Clairvoyance": "clairvoyance",
      "Combat Effects": "combateffects",
      "Damage Shield Mitigation": "dsmitigation",
      "Damage Shield": "damageshield",
      "DoT Shielding": "dotshielding",
      "Endurance Regen": "enduranceregen",
      "Heal Amount": "healamt",
      "HP Regen": "regen",
      "Mana Regen": "manaregen",
      "Shielding": "shielding",
      "Spell Damage": "spelldmg",
      "Spell Shielding": "spellshield",
      "Strikethrough": "strikethrough",
      "Stun Resist": "stunresist",
    };
  }

  // these were pulled from EOC and should be refined
  public static getFieldDescriptions() {
    return {
      "id": "This is the unique item ID in the database",
      "minstatus": "Minimum GM status to summon this item",
      "Name": "The name of the Item itself",
      "aagi": "Agility Statistic - This determines how often you get hit or missed by an attack and how much damage you take when you get hit.  It also affects how quickly you learn defensive skills.",
      "ac": "",
      "accuracy": "Improved chance to hit. This adds to the \"hit chance\" part of your Attack. 15 accuracy would equal to 1% more landed hits on the mob.",
      "acha": "This stat determines how charismatic you are. It is important for charming and pacify abilities. (enchanter, bard, cleric, ranger). In the raiding game, it will also affect how often Divine Intervention saves you as a tank. However, I highly suggest you ignore this ability at character creation unless you are a bard or enchanter.",
      "adex": "Determines your skills with weapons. Higher dexterity means that if your weapons have effects on them, they will activate (or proc) more frequently. High Dexterity also determines how often you will critical strike with weapons.",
      "aint": "Determines the starting mana for Bards, Enchanters, Magicians, Necromancers, Shadowknights, and Wizards. The higher this stat is for those classes, the more mana they will have.",
      "artifactflag": "",
      "asta": "Determines your base health along with your class and race. The higher your stamina, the more health you will have",
      "astr": "Determines how much you can carry. Also adds to the damage you do with melee weapons.",
      "attack": "This used to be called Vengeance wich no longer exists, it's now just \"+X attack\". 10 attack = around 1% more dps.",
      "augrestrict": "This restricts the augment to only be inserted in these types of items",
      "augslot1type": "When making an item this is the augmentation type that will fit into slot 1 (0 for none)",
      "augslot1visible": "This is either 1 (Visible) or 0 (Not Visible)",
      "augslot2type": "When making an item this is the augmentation type that will fit into slot 2 (0 for none)",
      "augslot2visible": "This is either 1 (Visible) or 0 (Not Visible)",
      "augslot3type": "When making an item this is the augmentation type that will fit into slot 3 (0 for none)",
      "augslot3visible": "This is either 1 (Visible) or 0 (Not Visible)",
      "augslot4type": "When making an item this is the augmentation type that will fit into slot 4 (0 for none)",
      "augslot4visible": "This is either 1 (Visible) or 0 (Not Visible)",
      "augslot5type": "When making an item this is the augmentation type that will fit into slot 5 (0 for none)",
      "augslot5visible": "This is either 1 (Visible) or 0 (Not Visible)",
      "augtype": "These is the arugment type(s) the augment will fit into",
      "avoidance": "Amount of Avoidance on the Item",
      "awis": "The same effect as intelligence for the deity-based cc users:  Cleric, Shaman Druid, Paladin and Ranger.  The higher your wisdom, the more mana you get for each level.",
      "bagsize": "The max size an item can be and still fit into the container. 1 = Small, 2 = Medium, 3 = Large, 4 = Giant, 5 = Giant (Assembly Kit only?)",
      "bagslots": "The number of slots the bag has.",
      "bagtype": "",
      "bagwr": "Bag Weight Reduction. Interestingly, you can also have bags that reduce coin weight. However, in order to set bags to reduce coin weight, you MUST set the Item ID somewhere between 17201 and 17230.",
      "banedmgamt": "Valid values: 0 -> 255 (-127 -> 127)",
      "banedmgraceamt": "Valid values: 0 -> 4294967295",
      "banedmgbody": "Valid values: 0 -> 4294967295",
      "banedmgrace": "The Race that the Bane Damage will affect (0-255)",
      "bardtype": "The type of instrument modified by bardvalue",
      "bardvalue": "How much the instrument type (bardtype) is modified when equipped",
      "book": "",
      "casttime": "For items that have a castable effect, the cast time of the spell (in milliseconds). On food/drink items, this determines how long they last before they are consumed:, -1 - 5: This is a snack. | This is a whistle wetter., 6 - 20: This is a meal. | This is a drink., 21 - 30: This is a hearty meal. | This is a refreshing drink., 31 - 40: This is a banquet size meal. | This is a lasting drink., 41 - 50: This meal is a feast! | This is a flowing drink!, 51 - 60: This is an enduring meal! | This is an enduring drink!, 70+: This is a miraculous meal! | This is a miraculous drink!",
      "casttime_": "",
      "charmfile": "This field corresponds to the quests/items/ folder.  The name set in this field will be the script loaded for this item such as 'CharmTest' would use the file 'CharmTest.pl' or 'CharmTest.lua'",
      "charmfileid": "Needs to be non-zero if item is to be scaled as a charm",
      "classes": "The class or classes can use the item",
      "color": "This is the tint/color that will show up for this item if it is a visible armor piece.",
      "combateffects": "",
      "extradmgskill": "Increases the amount of damage a skill can do. This defines the skill. Values can be 0 - 255.",
      "extradmgamt": "Increases the amount of damage a skill can do. This defines how much.",
      "price": "The price the item should cost from a vendor (In copper currency)",
      "cr": "The amount of Cold Resist provided by this item.",
      "damage": "",
      "damageshield": "",
      "deity": "",
      "delay": "This is the delay for a melee or ranged weapon.  It is set in 10ths of a second, so a value of 25 would be a 2.5 second delay.",
      "augdistiller": "",
      "dotshielding": "",
      "dr": "The amount of Disease Resist provided by this item.",
      "clicktype": "(1 or 3 = Clickable from Inventory), (4 = Must Equip to Cast), (5 = Rogue Poison)",
      "clicklevel2": "This is the level that the click effect will be cast as.",
      "elemdmgtype": "",
      "elemdmgamt": "",
      "endur": "The amount of Endurance provided by this item.",
      "factionamt1": "The amount of faction that will be modified.  This can be positive or negative values.",
      "factionmod1": "The faction ID to be modified by Faction Amount 1 (Table faction_list)",
      "factionamt2": "The amount of faction that will be modified.  This can be positive or negative values.",
      "factionmod2": "The faction ID to be modified by Faction Amount 2 (Table faction_list)",
      "factionamt3": "The amount of faction that will be modified.  This can be positive or negative values.",
      "factionmod3": "The faction ID to be modified by Faction Amount 3 (Table faction_list)",
      "factionamt4": "The amount of faction that will be modified.  This can be positive or negative values.",
      "factionmod4": "The faction ID to be modified by Faction Amount 4 (Table faction_list)",
      "filename": "Books filename which is linked to the `books` table under column `name`",
      "focuseffect": "The spell ID used for the focus effect on this item. (-1 for none)",
      "fr": "The amount of Fire Resist provided by this item.",
      "fvnodrop": "Sets item to NO DROP under the FV ruleset",
      "haste": "",
      "clicklevel": "The required level to use the click effect on this item.",
      "hp": "The amount of Hit Points provided by this item.",
      "regen": "The amount of Hit Point Regeneration provided by this item.",
      "icon": "This is the item icon that will be displayed when inspecting an item.",
      "idfile": "This defines how an item will look when equiped or dropped on the ground.  The model number must start with 'IT' and end with numbers such as 'IT63' which is the default bag model.",
      "itemclass": "",
      "itemtype": "",
      "ldonprice": "The price of an item in LDoN points when purchased from an LDoN merchant.",
      "ldontheme": "The LDoN Theme that this item is sold for.  This correlates to the LDoN Price.",
      "ldonsold": "This defines if an item can be sold to an LDoN merchant or not.",
      "light": "The amount of light given off when this item is equipped or put into a normal inventory slot (not inside a bag)",
      "lore": "This is the lore description that will show en an item inspect window when it is identified.",
      "loregroup": "Characters can only have 1 item from any lore group that is set to something other than 0.  Epic Lore is defined as Lore Group 1, but you can create any number of other Lore Groups.",
      "magic": "Sets item to be magical, you like your items magical, don't you?",
      "mana": "The amount of Mana provided by this item.",
      "manaregen": "The amount of Mana Regeneration provided by this item.",
      "enduranceregen": "The amount of Endurance Regeneration provided by this item.",
      "material": "This is the texture used for the item.  Only worn armor pieces require this setting.",
      "maxcharges": "The maximum charges for the Click Effect on this item.  Setting a value of '-1' will make charges unlimited.",
      "mr": "The amount of Magic Resist provided by this item.",
      "nodrop": "Sets item to not be droppable.",
      "norent": "Sets item to no rent.",
      "pendingloreflag": "This field is unused.  It was used on Live to identify items that will soon be changed to Lore.",
      "pr": "The amount of Poison Resist provided by this item.",
      "procrate": "The percentage that a weapon will proc. (0 = Normal, 50 = 150%)",
      "races": "",
      "range": "This is the range that an ammo or ranged weapon will use.",
      "reclevel": "This is the Recommended Level to use an item.  The item's stats will be scaled down if below this level.",
      "recskill": "",
      "reqlevel": "This is the Required Level to use an item.  No stats will be gained from this item if below this level.",
      "sellrate": "The adjusted rate that merchants will buy an item for.  This is a percentage of the price where 1 is 100%.  Do not set this above 1 unless you want an item to be able to be sold for more than it costs to buy.",
      "shielding": "",
      "size": "This is the size that will be displayed on an item.  Item size is used primarily for determining if the item can be stored in a specific bag depending on the size allowed by the bag.",
      "skillmodtype": "",
      "skillmodvalue": "",
      "slots": "",
      "clickeffect": "The spell ID used for the clickable effect on this item. (-1 for none)",
      "spellshield": "",
      "strikethrough": "",
      "stunresist": "",
      "summonedflag": "",
      "tradeskills": "",
      "favor": "Amount of Person Favor this item should give when turned in to favor NPC",
      "weight": "Weight of the item (multiplied by 10) IE, Cloth Cap weighs 0.2, is stored as 2.",
      "UNK012": "",
      "UNK013": "",
      "benefitflag": "",
      "UNK054": "",
      "UNK059": "",
      "elitematerial": "If set to 1, this will display a slightly different texture material on armor.  This only works for Drakkin Race.",
      "booktype": "Determines the visual appearance of the book in game, example scrolls, book, parchment etc.",
      "recastdelay": "This is the required time to wait (in ms) between using the click effect.",
      "recasttype": "(-1 = None) This is the group that the recast delay will be used in.  All clickable items in this same group will also be required to wait until the recast delay is done.",
      "guildfavor": "Amount of Guild Favor this item should give when turned in to guild favor NPC *This value should match the favor column*",
      "UNK123": "",
      "UNK124": "",
      "attuneable": "Once enabled, item becomes NO DROP when equipped for the first time.",
      "nopet": "If enabled, this item will not be able to be equipped by pets.",
      "updated": "This is the timestamp that the item was last updated.",
      "comment": "This is an optional field that is used to make notes/comments for the item and is not used by the server or client in any way.",
      "UNK127": "",
      "pointtype": "",
      "potionbelt": "If set, this item can be used from the potion belt.",
      "potionbeltslots": "",
      "stacksize": "This is the maximum stack size used for stackable items.",
      "notransfer": "",
      "stackable": "If enabled, this item will be stackable. Stack size must be set to the desired maximum stack size.",
      "UNK134": "",
      "UNK137": "",
      "proceffect": "The spell ID used for the Proc effect on this item. (-1 for none)",
      "proctype": "",
      "proclevel2": "",
      "proclevel": "",
      "UNK142": "",
      "worneffect": "The spell ID used for the Worn effect on this item. (-1 for none)",
      "worntype": "0 - None, 4 - Worn",
      "wornlevel2": "Level the effect will reach Maximum effect",
      "wornlevel": "Minimum Level required for the effect to work",
      "UNK147": "",
      "focustype": "0 = None, 6 = Focus",
      "focuslevel2": "Level Focus will reach Max effect",
      "focuslevel": "Level Required",
      "UNK152": "",
      "scrolleffect": "The spell ID used for the Scroll effect on this item. (-1 for none)",
      "scrolltype": "The Spell ID in which the scroll will scribe",
      "scrolllevel2": "",
      "scrolllevel": "",
      "UNK157": "",
      "bardeffect": "The spell ID used for the Bard effect on this item. (-1 for none)",
      "bardeffecttype": "",
      "bardlevel": "",
      "bardlevel2": "",
      "serialized": "",
      "verified": "This is the date/time that the item was manually verified to be valid.",
      "serialization": "",
      "source": "This is a notes field for defining where the item stats came from (IE. 13th Floor, Custom, ZAM, etc.)",
      "UNK033": "",
      "lorefile": "",
      "UNK014": "",
      "svcorruption": "",
      "UNK038": "",
      "UNK060": "",
      "augslot1unk2": "",
      "augslot2unk2": "",
      "augslot3unk2": "",
      "augslot4unk2": "",
      "augslot5unk2": "",
      "ldonsellbackrate": "This is the percentage at which an item can be sold back to an LDoN merchant.  Do not set above 100 (100%).",
      "UNK120": "",
      "UNK121": "",
      "questitemflag": "Item flag that simply signifies its use in quests.",
      "expendablearrow": "",
      "UNK132": "",
      "clickunk5": "",
      "clickunk6": "",
      "clickunk7": "",
      "procunk1": "",
      "procunk2": "",
      "procunk3": "",
      "procunk4": "",
      "procunk6": "",
      "procunk7": "",
      "wornunk1": "",
      "wornunk2": "",
      "wornunk3": "",
      "wornunk4": "",
      "wornunk5": "",
      "wornunk6": "",
      "wornunk7": "",
      "focusunk1": "",
      "focusunk2": "",
      "focusunk3": "",
      "focusunk4": "",
      "focusunk5": "",
      "focusunk6": "",
      "focusunk7": "",
      "scrollunk1": "",
      "scrollunk2": "",
      "scrollunk3": "",
      "scrollunk4": "",
      "scrollunk5": "",
      "scrollunk6": "",
      "scrollunk7": "",
      "UNK193": "",
      "purity": "",
      "evolvinglevel": "",
      "scriptfileid": "This is the ID (numeric value) of a script that is called from the 'quests/items/' folder when this item is right clicked. The number set in this field will be the script loaded for this item such as '123456' would use the file 'script_123456.pl' or 'script_123456.lua'",
      "clickname": "If a name is defined here, it will override the spell name of the set click effect (only works for SoF+ clients)",
      "procname": "If a name is defined here, it will override the spell name of the set proc effect (only works for SoF+ clients)",
      "wornname": "If a name is defined here, it will override the spell name of the set worn effect (only works for SoF+ clients)",
      "focusname": "If a name is defined here, it will override the spell name of the set focus effect (only works for SoF+ clients)",
      "scrollname": "If a name is defined here, it will override the spell name of the set scroll effect (only works for SoF+ clients)",
      "bardname": "If a name is defined here, it will override the spell name of the set bard effect (only works for SoF+ clients)",
      "dsmitigation": "",
      "heroic_str": "",
      "heroic_int": "",
      "heroic_wis": "",
      "heroic_agi": "",
      "heroic_dex": "",
      "heroic_sta": "",
      "heroic_cha": "",
      "heroic_pr": "",
      "heroic_dr": "",
      "heroic_fr": "",
      "heroic_cr": "",
      "heroic_mr": "",
      "heroic_svcorrup": "",
      "healamt": "",
      "spelldmg": "",
      "clairvoyance": "",
      "backstabdmg": "",
      "created": "This is the date/time that the item was created."
    }
  }

  public static getFieldDescription(field: string) {
    // we do this because the payload we get back from spire API is
    // formatted slightly different
    let fieldLookup = field.toLowerCase().replace("_", "")

    for (let key in this.getFieldDescriptions()) {
      let keyLookup = key.toLowerCase().replace("_", "")
      if (keyLookup === fieldLookup) {
        return this.getFieldDescriptions()[key]
      }
    }

    return ''
  }

}
