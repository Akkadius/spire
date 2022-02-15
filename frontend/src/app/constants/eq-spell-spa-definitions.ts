export const SPELL_SPA_DEFINITIONS = {
  0: {
    spa: 'SE_CurrentHP',
    effectName: 'Current HP',
    description: 'Hit Points',
    base: 'amount',
    limit: 'target restriction id',
    value: 'max amount (use positive value)',
    bonusCalc: '',
    notes: 'Negative base value for damage | Positive base value for healing'
  },
  1: {
    spa: 'SE_ArmorClass',
    effectName: 'AC',
    description: 'Stat',
    base: 'amount',
    limit: 'none',
    value: 'amount',
    bonusCalc: '',
    notes: ''
  },
  2: {
    spa: 'SE_ATK',
    effectName: 'ATK',
    description: 'Stat',
    base: 'amount',
    limit: 'none',
    value: 'amount',
    bonusCalc: '',
    notes: ''
  },
  3: {
    spa: 'SE_MovementSpeed',
    effectName: 'Movement Rate',
    description: 'Movement',
    base: 'amount',
    limit: 'none',
    value: 'amount',
    bonusCalc: '',
    notes: ''
  },
  4: {
    spa: 'SE_STR',
    effectName: 'STR',
    description: 'Stat',
    base: 'amount',
    limit: 'none',
    value: 'amount',
    bonusCalc: '',
    notes: ''
  },
  5: {
    spa: 'SE_DEX',
    effectName: 'DEX',
    description: 'Stat',
    base: 'amount',
    limit: 'none',
    value: 'amount',
    bonusCalc: '',
    notes: ''
  },
  6: {
    spa: 'SE_AGI',
    effectName: 'AGI',
    description: 'Stat',
    base: 'amount',
    limit: 'none',
    value: 'amount',
    bonusCalc: '',
    notes: ''
  },
  7: {
    spa: 'SE_STA',
    effectName: 'STA',
    description: 'Stat',
    base: 'amount',
    limit: 'none',
    value: 'amount',
    bonusCalc: '',
    notes: ''
  },
  8: {
    spa: 'SE_INT',
    effectName: 'INT',
    description: 'Stat',
    base: 'amount',
    limit: 'none',
    value: 'amount',
    bonusCalc: '',
    notes: ''
  },
  9: {
    spa: 'SE_WIS',
    effectName: 'WIS',
    description: 'Stat',
    base: 'amount',
    limit: 'none',
    value: 'amount',
    bonusCalc: '',
    notes: ''
  },
  10: {
    spa: 'SE_CHA',
    effectName: 'CHA',
    description: 'Stat',
    base: 'amount',
    limit: 'none',
    value: 'amount',
    bonusCalc: '',
    notes: ''
  },
  11: {
    spa: 'SE_AttackSpeed',
    effectName: 'Attack Speed',
    description: 'Attack Speed',
    base: 'percent haste or slow',
    limit: 'none',
    value: 'none',
    bonusCalc: 'Highest',
    notes: 'Base greater than 100 for haste (120 = 20 pct haste) |  Base less than 100 for slow (80 = 20 pct slow)'
  },
  12: {
    spa: 'SE_Invisibility',
    effectName: 'Invisibility: Unstable',
    description: 'Invisibility',
    base: 'invisibility level',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: 'Invisibility level determines what level of see invisible can detect it.'
  },
  13: {
    spa: 'SE_SeeInvis',
    effectName: 'See Invisible',
    description: 'Invisibility',
    base: 'see invisible level',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: 'See Invisible level determines what level of invisible it can see.'
  },
  14: {
    spa: 'SE_WaterBreathing',
    effectName: 'Water Breathing',
    description: 'Beneficial Uncategorized',
    base: '1 (increase for stacking overwrite)',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: ''
  },
  15: {
    spa: 'SE_CurrentMana',
    effectName: 'Mana',
    description: 'Mana',
    base: 'amount',
    limit: 'none',
    value: 'max amount (use positive value)',
    bonusCalc: '',
    notes: ''
  },
  16: {
    spa: 'SE_NPCFrenzy',
    effectName: 'NPC Frenzy',
    description: 'Unused',
    base: '',
    limit: '',
    value: '',
    bonusCalc: '',
    notes: ''
  },
  17: {
    spa: 'SE_NPCAwareness',
    effectName: 'NPC Awareness',
    description: 'Unused',
    base: '',
    limit: '',
    value: '',
    bonusCalc: '',
    notes: ''
  },
  18: {
    spa: 'SE_Lull',
    effectName: 'Pacify',
    description: 'Aggro',
    base: '1',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: ''
  },
  19: {
    spa: 'SE_AddFaction',
    effectName: 'NPC Faction',
    description: 'Faction',
    base: 'amount',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: ''
  },
  20: {
    spa: 'SE_Blind',
    effectName: 'Blindness',
    description: 'Detrimental Uncategorized',
    base: '1 (increase for stacking overwrite)',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: ''
  },
  21: {
    spa: 'SE_Stun',
    effectName: 'Stun',
    description: 'Stun',
    base: 'duration ms',
    limit: 'pvp duration ms',
    value: 'max target level',
    bonusCalc: '',
    notes: ''
  },
  22: {
    spa: 'SE_Charm',
    effectName: 'Charm',
    description: 'Charm',
    base: 'Unknown set to 1',
    limit: 'none',
    value: 'max target level',
    bonusCalc: '',
    notes: ''
  },
  23: {
    spa: 'SE_Fear',
    effectName: 'Fear',
    description: 'Detrimental Uncategorized',
    base: 'Unknown set to 1',
    limit: 'none',
    value: 'max target level',
    bonusCalc: '',
    notes: ''
  },
  24: {
    spa: 'SE_Stamina',
    effectName: 'Stamina Loss',
    description: 'Beneficial Uncategorized',
    base: 'amount',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: 'Positive value will reduce endurance upkeep | Negative value will increase endurance upkeep | Live no longer uses this spell effect for endurance regeneration.'
  },
  25: {
    spa: 'SE_BindAffinity',
    effectName: 'Bind Affinity',
    description: 'Bind',
    base: 'bind id (Set to 1, 2, or 3)',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: 'Bind id allows you set alternate bind points. Bind Point ID (1=Primary, 2=Secondary 3=Tertiary)'
  },
  26: {
    spa: 'SE_Gate',
    effectName: 'Gate',
    description: 'Teleport',
    base: 'success chance',
    limit: 'bind id  (2 or 3)',
    value: 'none',
    bonusCalc: '',
    notes: 'If limit is not set, you will gate to primary bind location. Bind Point ID (1=Primary, 2=Secondary 3=Tertiary)'
  },
  27: {
    spa: 'SE_CancelMagic',
    effectName: 'Dispel Magic',
    description: 'Dispel',
    base: 'chance level modifier',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: 'Success chance is based on level difference of caster and caster of the buff, base value raises the casters level by the base amount.'
  },
  28: {
    spa: 'SE_InvisVsUndead',
    effectName: 'Invisibility to Undead: Unstable',
    description: 'Invisibility',
    base: 'invisibility level',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: 'Invisibility level determines what level of see invisible can detect it.'
  },
  29: {
    spa: 'SE_InvisVsAnimals',
    effectName: 'Invisibility to Animals: Unstable',
    description: 'Invisibility',
    base: 'invisibility level',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: 'Invisibility level determines what level of see invisible can detect it.'
  },
  30: {
    spa: 'SE_ChangeFrenzyRad',
    effectName: 'NPC Aggro Radius',
    description: 'Aggro',
    base: 'amount',
    limit: 'none',
    value: 'max target level',
    bonusCalc: '',
    notes: ''
  },
  31: {
    spa: 'SE_Mez',
    effectName: 'Mesmerize',
    description: 'Detrimental Uncategorized',
    base: '1 (increase for stacking overwrite)',
    limit: 'none',
    value: 'max target level',
    bonusCalc: '',
    notes: 'Higher value of stacking type will always override the lower value. Used if you want one type of mez to overrite another.'
  },
  32: {
    spa: 'SE_SummonItem',
    effectName: 'Summon Item',
    description: 'Summon',
    base: 'item id',
    limit: 'none',
    value: 'stack amount',
    bonusCalc: '',
    notes: ''
  },
  33: {
    spa: 'SE_SummonPet',
    effectName: 'Summon Pet',
    description: 'Pet',
    base: '1',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: 'Set \'Teleport Zone\' field to be the same as \'type\' field in of the pet you want in the pets table.'
  },
  34: {
    spa: 'SE_Confuse',
    effectName: 'Confuse',
    description: 'Unused',
    base: '',
    limit: '',
    value: '',
    bonusCalc: '',
    notes: ''
  },
  35: {
    spa: 'SE_DiseaseCounter',
    effectName: 'Disease Counter',
    description: 'Cures',
    base: 'amount',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: 'Set to positive values for potency of detrimental spells | Set to negative value for potency of cure spells.'
  },
  36: {
    spa: 'SE_PoisonCounter',
    effectName: 'Poison Counter',
    description: 'Cures',
    base: 'amount',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: 'Set to positive values for potency of detrimental spells | Set to negative value for potency of cure spells.'
  },
  37: {
    spa: 'SE_DetectHostile',
    effectName: 'Detect Hostile',
    description: 'Unused',
    base: '',
    limit: '',
    value: '',
    bonusCalc: '',
    notes: ''
  },
  38: {
    spa: 'SE_DetectMagic',
    effectName: 'Detect Magic',
    description: 'Unused',
    base: '',
    limit: '',
    value: '',
    bonusCalc: '',
    notes: ''
  },
  39: {
    spa: 'SE_TwinCastBlocker',
    effectName: 'Twincast Blocker',
    description: 'Stacking',
    base: '1',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: ''
  },
  40: {
    spa: 'SE_DivineAura',
    effectName: 'Invulnerability',
    description: 'Beneficial Uncategorized',
    base: '1',
    limit: 'none',
    value: 'none',
    bonusCalc: 'Bool',
    notes: ''
  },
  41: {
    spa: 'SE_Destroy',
    effectName: 'Destroy',
    description: 'Hit Points',
    base: '1',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: ''
  },
  42: {
    spa: 'SE_ShadowStep',
    effectName: 'Shadow Step',
    description: 'Movement Speed',
    base: 'Unknown (Seen 1 to 50)',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: 'This effect is handled by the client. Changing the base value does not appear to have any affect.'
  },
  43: {
    spa: 'SE_Berserk',
    effectName: 'Berserk',
    description: 'Offensive Bonus',
    base: '1',
    limit: 'none',
    value: 'none',
    bonusCalc: 'Bool',
    notes: 'This is an unused live spell effect. Custom Spell Effect may be subject to change if live reuses the SPA.'
  },
  44: {
    spa: 'SE_Lycanthropy',
    effectName: 'Stacking:  Delayed Heal Marker',
    description: 'Stacking',
    base: 'stacking overwrite value',
    limit: 'none',
    value: 'none',
    bonusCalc: 'none',
    notes: ''
  },
  45: {
    spa: 'SE_Vampirism',
    effectName: 'Vampirism',
    description: 'Offensive Bonus',
    base: 'Percentage',
    limit: 'none',
    value: 'none',
    bonusCalc: 'Additive',
    notes: 'This is an unused live spell effect. Custom Spell Effect may be subject to change if live reuses the SPA.'
  },
  46: {
    spa: 'SE_ResistFire',
    effectName: 'Fire Resist',
    description: 'Resist',
    base: 'amout',
    limit: 'none',
    value: 'max amount',
    bonusCalc: '',
    notes: ''
  },
  47: {
    spa: 'SE_ResistCold',
    effectName: 'Cold Resist',
    description: 'Resist',
    base: 'amout',
    limit: 'none',
    value: 'max amount',
    bonusCalc: '',
    notes: ''
  },
  48: {
    spa: 'SE_ResistPoison',
    effectName: 'Poison Resist',
    description: 'Resist',
    base: 'amout',
    limit: 'none',
    value: 'max amount',
    bonusCalc: '',
    notes: ''
  },
  49: {
    spa: 'SE_ResistDisease',
    effectName: 'Disease Resist',
    description: 'Resist',
    base: 'amout',
    limit: 'none',
    value: 'max amount',
    bonusCalc: '',
    notes: ''
  },
  50: {
    spa: 'SE_ResistMagic',
    effectName: 'Magic Resist',
    description: 'Resist',
    base: 'amout',
    limit: 'none',
    value: 'max amount',
    bonusCalc: '',
    notes: ''
  },
  51: {
    spa: 'SE_DetectTraps',
    effectName: 'Detect Traps',
    description: 'Unused',
    base: '',
    limit: '',
    value: '',
    bonusCalc: '',
    notes: ''
  },
  52: {
    spa: 'SE_SenseDead',
    effectName: 'Detect Undead',
    description: 'Tracking',
    base: '1',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: ''
  },
  53: {
    spa: 'SE_SenseSummoned',
    effectName: 'Detect Summoned',
    description: 'Tracking',
    base: '1',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: ''
  },
  54: {
    spa: 'SE_SenseAnimals',
    effectName: 'Detect Animals',
    description: 'Tracking',
    base: '1',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: ''
  },
  55: {
    spa: 'SE_Rune',
    effectName: 'Rune',
    description: 'Rune',
    base: 'amount',
    limit: 'none',
    value: 'max amount',
    bonusCalc: '',
    notes: ''
  },
  56: {
    spa: 'SE_TrueNorth',
    effectName: 'True North',
    description: 'Tracking',
    base: '1',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: ''
  },
  57: {
    spa: 'SE_Levitate',
    effectName: 'Levitation',
    description: 'Beneficial Uncategorized',
    base: '1 (increase for stacking overwrite)',
    limit: 'levitate while moving (Set to 1)',
    value: 'none',
    bonusCalc: '',
    notes: 'Levitate while moving is seen on Live \'Flying Mounts\''
  },
  58: {
    spa: 'SE_Illusion',
    effectName: 'Illusion',
    description: 'Beneficial Uncategorized',
    base: 'race id or gender id',
    limit: 'texture id (see notes)',
    value: 'helmet id (see notes)',
    bonusCalc: '',
    notes: 'Illusions have complicated rules. See https://docs.eqemu.io/server/categories/spells/illusion-spell-guidelines'
  },
  59: {
    spa: 'SE_DamageShield',
    effectName: 'Damage Shield',
    description: 'Damage Shield',
    base: 'amount',
    limit: 'none',
    value: 'max amount',
    bonusCalc: '',
    notes: ''
  },
  60: {
    spa: 'SE_TransferItem',
    effectName: 'Transfer Item',
    description: 'not used',
    base: '',
    limit: '',
    value: '',
    bonusCalc: '',
    notes: ''
  },
  61: {
    spa: 'SE_Identify',
    effectName: 'Identify',
    description: 'Beneficial Uncategorized',
    base: '1',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: 'To use, hold item on cursor and cast spell on self or target. The \'lore\' field from items table is displayed in chat window.'
  },
  62: {
    spa: 'SE_ItemID',
    effectName: 'Item ID',
    description: 'Unused',
    base: '',
    limit: '',
    value: '',
    bonusCalc: '',
    notes: ''
  },
  63: {
    spa: 'SE_WipeHateList',
    effectName: 'Memblur',
    description: 'Memory Blur',
    base: 'percent chance',
    limit: 'none',
    value: 'unknown',
    bonusCalc: '',
    notes: 'Actual chance to memory blur is much higher than the spells base value, caster level and CHA modifiers are added get the final calculated percent chance'
  },
  64: {
    spa: 'SE_SpinTarget',
    effectName: 'Spin Stun',
    description: 'Stun',
    base: 'duration ms',
    limit: 'pvp duration ms',
    value: 'max target level',
    bonusCalc: '',
    notes: ''
  },
  65: {
    spa: 'SE_InfraVision',
    effectName: 'Infravision',
    description: 'Vision',
    base: '1',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: ''
  },
  66: {
    spa: 'SE_UltraVision',
    effectName: 'Ultravision',
    description: 'Vision',
    base: '1',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: ''
  },
  67: {
    spa: 'SE_EyeOfZomm',
    effectName: 'Eye Of Zomm',
    description: 'Vision',
    base: '1',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: 'Set \'Teleport Zone\' field to be the same as \'type\' field in of the pet you want in the pets table.'
  },
  68: {
    spa: 'SE_ReclaimPet',
    effectName: 'Reclaim Energy',
    description: 'Mana',
    base: '1',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: ''
  },
  69: {
    spa: 'SE_TotalHP',
    effectName: 'Max HP',
    description: 'Hit Points',
    base: 'amount',
    limit: 'none',
    value: 'max amount',
    bonusCalc: '',
    notes: ''
  },
  70: {
    spa: 'SE_CorpseBomb',
    effectName: 'Corpse Bomb',
    description: 'Unused',
    base: '',
    limit: '',
    value: '',
    bonusCalc: '',
    notes: ''
  },
  71: {
    spa: 'SE_NecPet',
    effectName: 'Create Undead Pet',
    description: 'Pet',
    base: '1',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: 'Set \'Teleport Zone\' field to be the same as \'type\' field in of the pet you want in the pets table.'
  },
  72: {
    spa: 'SE_PreserveCorpse',
    effectName: 'Preserve Corpse',
    description: 'Unused',
    base: '',
    limit: '',
    value: '',
    bonusCalc: '',
    notes: ''
  },
  73: {
    spa: 'SE_BindSight',
    effectName: 'Bind Sight',
    description: 'Vision',
    base: '1',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: ''
  },
  74: {
    spa: 'SE_FeignDeath',
    effectName: 'Feign Death',
    description: 'Beneficial Uncategorized',
    base: 'success chance',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: ''
  },
  75: {
    spa: 'SE_VoiceGraft',
    effectName: 'Voice Graft',
    description: 'Utility Uncategorized',
    base: '1',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: ''
  },
  76: {
    spa: 'SE_Sentinel',
    effectName: 'Sentinel',
    description: 'Not Implemented',
    base: '1',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: 'Not implemented on EQEMU'
  },
  77: {
    spa: 'SE_LocateCorpse',
    effectName: 'Locate Corpse',
    description: 'Tracking',
    base: '1',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: ''
  },
  78: {
    spa: 'SE_AbsorbMagicAtt',
    effectName: 'Spell Rune',
    description: 'Rune',
    base: 'amount',
    limit: 'none',
    value: 'max amount',
    bonusCalc: '',
    notes: ''
  },
  79: {
    spa: 'SE_CurrentHPOnce',
    effectName: 'Current HP Once',
    description: 'Hit points',
    base: 'amount',
    limit: 'target restriction id',
    value: 'max amount',
    bonusCalc: '',
    notes: 'Negative base value for damage | Positive base value for healing'
  },
  80: {
    spa: 'SE_EnchantLight',
    effectName: 'Enchant Light',
    description: 'Unused',
    base: '',
    limit: '',
    value: '',
    bonusCalc: '',
    notes: ''
  },
  81: {
    spa: 'SE_Revive',
    effectName: 'Resurrect',
    description: 'Corpse',
    base: 'percentage exp',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: ''
  },
  82: {
    spa: 'SE_SummonPC',
    effectName: 'Summon Player',
    description: 'Utility Uncategorized',
    base: '1',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: ''
  },
  83: {
    spa: 'SE_Teleport',
    effectName: 'Teleport',
    description: 'Teleport',
    base: 'coordinate(x,y,z,h)',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: 'Set \'Teleport Zone\' field to zone short name OR set to \'same\' to teleport within same zone. To set all xyzh cooridinates, you have use the following. Use this effectid  only once in first effect slot . Cooridinates defined as effect_base_value1=x effect_base_value2=y effect_base_value3=z effect_base_value4=h'
  },
  84: {
    spa: 'SE_TossUp',
    effectName: 'Gravity Flux',
    description: 'Detrimental Uncategorized',
    base: 'distance (negative)',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: ''
  },
  85: {
    spa: 'SE_WeaponProc',
    effectName: 'Add Melee Proc',
    description: 'Procs',
    base: 'spellid',
    limit: 'rate modifer',
    value: 'none',
    bonusCalc: '',
    notes: ''
  },
  86: {
    spa: 'SE_Harmony',
    effectName: 'Reaction Radius',
    description: 'Aggro',
    base: 'amount',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: ''
  },
  87: {
    spa: 'SE_MagnifyVision',
    effectName: 'Magnification',
    description: 'Vision',
    base: 'amount',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: ''
  },
  88: {
    spa: 'SE_Succor',
    effectName: 'Evacuate',
    description: 'Teleport',
    base: 'coordinate(x,y,z,h)',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: 'Set \'Teleport Zone\' field to zone short name OR set to \'same\' to evac within same zone. To set all xyzh cooridinates, you have use the following. Use this effectid only once in first effect slot . Cooridinates defined as effect_base_value1=x effect_base_value2=y effect_base_value3=z effect_base_value4=h'
  },
  89: {
    spa: 'SE_ModelSize',
    effectName: 'Player Size',
    description: 'Utility Uncategorized',
    base: 'percent shrink or grow',
    limit: 'model size',
    value: 'unknown',
    bonusCalc: '',
    notes: 'Base greater than 100 for growth (120 = 20 pct growth) |  Base less than 100 for shrink (80 = 20 pct shrink) | To set to a specific model size, set base to 100 or 0 and then set limit to model size value.'
  },
  90: {
    spa: 'SE_Cloak',
    effectName: 'Ignore Pet',
    description: 'Not Implemented',
    base: '1',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: 'Not implemented on EQEMU'
  },
  91: {
    spa: 'SE_SummonCorpse',
    effectName: 'Summon Corpse',
    description: 'Corpse',
    base: 'max target level',
    limit: '',
    value: '',
    bonusCalc: '',
    notes: ''
  },
  92: {
    spa: 'SE_InstantHate',
    effectName: 'Hate',
    description: 'Aggro',
    base: 'amount',
    limit: 'none',
    value: 'max amount',
    bonusCalc: '',
    notes: 'Positive value increases hate. | Negative value decreases hate.'
  },
  93: {
    spa: 'SE_StopRain',
    effectName: 'Control Weather',
    description: 'Utility Uncategorized',
    base: '1',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: ''
  },
  94: {
    spa: 'SE_NegateIfCombat',
    effectName: 'Make Fragile',
    description: 'Restriction',
    base: '1',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: ''
  },
  95: {
    spa: 'SE_Sacrifice',
    effectName: 'Sacrifice',
    description: 'Utility Uncategorized',
    base: '1',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: ''
  },
  96: {
    spa: 'SE_Silence',
    effectName: 'Silence',
    description: 'Silence',
    base: '1 (increase for stacking overwrite)',
    limit: 'none',
    value: 'unknown',
    bonusCalc: '',
    notes: ''
  },
  97: {
    spa: 'SE_ManaPool',
    effectName: 'Max Mana',
    description: 'Mana',
    base: 'amount',
    limit: 'none',
    value: 'max amount',
    bonusCalc: '',
    notes: ''
  },
  98: {
    spa: 'SE_AttackSpeed2',
    effectName: 'Attack Speed: Does not exceed cap',
    description: 'Attack Speed',
    base: 'percent haste or slow',
    limit: 'none',
    value: 'none',
    bonusCalc: 'Highest',
    notes: 'Base greater than 100 for haste (120 = 20 pct haste) |  Base less than 100 for slow (80 = 20 pct slow)'
  },
  99: {
    spa: 'SE_Root',
    effectName: 'Root',
    description: 'Movement Speed',
    base: '-10000',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: ''
  },
  100: {
    spa: 'SE_HealOverTime',
    effectName: 'Heal Over Time',
    description: 'Hit Points',
    base: 'amount',
    limit: 'target restriction id',
    value: 'max amount',
    bonusCalc: '',
    notes: ''
  },
  101: {
    spa: 'SE_CompleteHeal',
    effectName: 'Complete Heal: With Recast Blocker Buff',
    description: 'Hit Points',
    base: 'heal amount multiplier',
    limit: 'none',
    value: 'max heal amount multipler',
    bonusCalc: '',
    notes: ''
  },
  102: {
    spa: 'SE_Fearless',
    effectName: 'Fear Immunity',
    description: 'Resist',
    base: '1',
    limit: 'none',
    value: 'none',
    bonusCalc: 'Bool',
    notes: ''
  },
  103: {
    spa: 'SE_CallPet',
    effectName: 'Summon Pet',
    description: 'Pet',
    base: '1',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: ''
  },
  104: {
    spa: 'SE_Translocate',
    effectName: 'Translocate',
    description: 'Teleport',
    base: 'coordinate(x,y,z,h) or Bind Point ID',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: 'Set \'Teleport Zone\' field to zone short name OR set to \'same\' to evac within same zone. To set all xyzh cooridinates, you have use the following. Use this effectid only once in first effect slot . If \'Teleport_Zone\' field is not set, then will send to  bind point id, set base value to Bind Point ID (1=Primary, 2=Secondary 3=Tertiary)'
  },
  105: {
    spa: 'SE_AntiGate',
    effectName: 'Inhibit Gate',
    description: 'Teleport',
    base: 'Seen 1 to 3',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: 'Unclear what base value determines. May be related to Bind Point IDs.'
  },
  106: {
    spa: 'SE_SummonBSTPet',
    effectName: 'Summon Warder',
    description: 'Pet',
    base: '1',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: 'Set \'Teleport Zone\' field to be the same as \'type\' field in of the pet you want in the pets table.'
  },
  107: {
    spa: 'SE_AlterNPCLevel',
    effectName: 'Alter NPC Level',
    description: 'Utility Uncategorized',
    base: 'amount',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: 'This is a no longer used on live. Custom Spell Effect may be subject to change if live reuses the SPA.'
  },
  108: {
    spa: 'SE_Familiar',
    effectName: 'Summon Familiar',
    description: 'Pet',
    base: '0 or 1',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: 'Set \'Teleport Zone\' field to be the same as \'type\' field in of the pet you want in the pets table.'
  },
  109: {
    spa: 'SE_SummonItemIntoBag',
    effectName: 'Summon into Bag',
    description: 'Summon',
    base: 'item id',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: 'To use this the first effectid must be SPA 32 SE_SummonItem and this must be a bag such as Phantom Satchel ID 17310. Then use this effectid to summon items that will go into that bag.'
  },
  110: {
    spa: 'SE_IncreaseArchery',
    effectName: 'Increase Archery',
    description: 'Not used',
    base: '',
    limit: '',
    value: '',
    bonusCalc: '',
    notes: ''
  },
  111: {
    spa: 'SE_ResistAll',
    effectName: 'All Resists',
    description: 'Resists',
    base: 'amount',
    limit: 'none',
    value: 'max amount',
    bonusCalc: '',
    notes: ''
  },
  112: {
    spa: 'SE_CastingLevel',
    effectName: 'Casting Level',
    description: 'Casting',
    base: 'amount',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: ''
  },
  113: {
    spa: 'SE_SummonHorse',
    effectName: 'Summon Mount',
    description: 'Summon',
    base: '1',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: 'Set \'Teleport Zone\' field to be the same as \'filename\' field in of the mount you want in the horses table.'
  },
  114: {
    spa: 'SE_ChangeAggro',
    effectName: 'Hate Multiplier',
    description: 'Aggro',
    base: 'percent hate modifer',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: ''
  },
  115: {
    spa: 'SE_Hunger',
    effectName: 'Food',
    description: 'Beneficial Uncategorized',
    base: '1',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: 'Set to positive values for potency of detrimental spells | Set to negative value for potency of cure spells.'
  },
  116: {
    spa: 'SE_CurseCounter',
    effectName: 'Curse Counter',
    description: 'Cures',
    base: 'amount',
    limit: '',
    value: '',
    bonusCalc: '',
    notes: ''
  },
  117: {
    spa: 'SE_MagicWeapon',
    effectName: 'Make Weapons Magical',
    description: 'Beneficial Uncategorized',
    base: '1',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: ''
  },
  118: {
    spa: 'SE_Amplification',
    effectName: 'Singing Amplification',
    description: 'Songs',
    base: 'percent',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: 'Recasting this effect will cause it to focus itself, increasing its potency.'
  },
  119: {
    spa: 'SE_AttackSpeed3',
    effectName: 'Attack Speed: Overhaste',
    description: 'Attack Speed',
    base: 'percent haste or slow',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: 'Base greater than 100 for haste (120 = 20 pct haste) |  Base less than 100 for slow (80 = 20 pct slow)'
  },
  120: {
    spa: 'SE_HealRate',
    effectName: 'Incoming Healing Effectiveness',
    description: 'Healing Modifer',
    base: 'percent healing',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: ''
  },
  121: {
    spa: 'SE_ReverseDS',
    effectName: 'Reverse Damage Shield',
    description: 'Damage Shield',
    base: 'amount',
    limit: '',
    value: '',
    bonusCalc: '',
    notes: 'Negative value will cause the reverse damage shield to heal.'
  },
  122: {
    spa: 'SE_ReduceSkill',
    effectName: 'Reduce Skill',
    description: 'Not Implemented',
    base: 'pending',
    limit: 'pending',
    value: 'pending',
    bonusCalc: '',
    notes: 'not implemented'
  },
  123: {
    spa: 'SE_Screech',
    effectName: 'Stacking:  Screech',
    description: 'Stacking',
    base: '1 or -1',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: ''
  },
  124: {
    spa: 'SE_ImprovedDamage',
    effectName: 'Focus: Spell Damage',
    description: 'Focus',
    base: 'min percent',
    limit: 'none',
    value: 'max percent',
    bonusCalc: '',
    notes: 'Use random effectiveness if base and max value are defined, where base is always lower end and max the higher end of the random range. If random value not wanted, then only set a base value.'
  },
  125: {
    spa: 'SE_ImprovedHeal',
    effectName: 'Focus: Healing',
    description: 'Focus',
    base: 'min percent',
    limit: 'none',
    value: 'max percent',
    bonusCalc: '',
    notes: 'Use random effectiveness if base and max value are defined, where base is always lower end and max the higher end of the random range. If random value not wanted, then only set a base value.'
  },
  126: {
    spa: 'SE_SpellResistReduction',
    effectName: 'Focus: Spell Resist Rate',
    description: 'Focus',
    base: 'min percent',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: ''
  },
  127: {
    spa: 'SE_IncreaseSpellHaste',
    effectName: 'Focus: Spell Cast Time',
    description: 'Focus',
    base: 'percent',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: ''
  },
  128: {
    spa: 'SE_IncreaseSpellDuration',
    effectName: 'Focus: Spell Duration',
    description: 'Focus',
    base: 'percent',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: ''
  },
  129: {
    spa: 'SE_IncreaseRange',
    effectName: 'Focus: Spell Range',
    description: 'Focus',
    base: 'percent',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: ''
  },
  130: {
    spa: 'SE_SpellHateMod',
    effectName: 'Focus: Spell and Bash Hate',
    description: 'Focus',
    base: 'min percent',
    limit: 'none',
    value: 'max percent',
    bonusCalc: '',
    notes: 'Use random effectiveness if base and max value are defined, where base is always lower end and max the higher end of the random range. If random value not wanted, then only set a base value. | Special case: For bash hate to be focused, need to add focus limit SPA 137 and set it to 999.'
  },
  131: {
    spa: 'SE_ReduceReagentCost',
    effectName: 'Focus: Chance of Using Reagent',
    description: 'Focus',
    base: 'min percent',
    limit: 'none',
    value: 'max percent',
    bonusCalc: '',
    notes: 'Use random effectiveness if base and max value are defined, where base is always lower end and max the higher end of the random range. If random value not wanted, then only set a base value.'
  },
  132: {
    spa: 'SE_ReduceManaCost',
    effectName: 'Focus: Spell Mana Cost',
    description: 'Focus',
    base: 'min percent',
    limit: 'none',
    value: 'max percent',
    bonusCalc: '',
    notes: 'Use random effectiveness if base and max value are defined, where base is always lower end and max the higher end of the random range. If random value not wanted, then only set a base value.'
  },
  133: {
    spa: 'SE_FcStunTimeMod',
    effectName: 'Focus: Spell Stun Duration',
    description: 'Focus',
    base: 'percent',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: ''
  },
  134: {
    spa: 'SE_LimitMaxLevel',
    effectName: 'Limit:  Max Level',
    description: 'Limit',
    base: 'max level',
    limit: 'effectiviness percent',
    value: 'none',
    bonusCalc: '',
    notes: ''
  },
  135: {
    spa: 'SE_LimitResist',
    effectName: 'Limit:  Resist',
    description: 'Limit',
    base: 'resist type',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: 'Include set value to positive | Exclude set value to negative'
  },
  136: {
    spa: 'SE_LimitTarget',
    effectName: 'Limit:  Target',
    description: 'Limit',
    base: 'target type',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: 'Include set value to positive | Exclude set value to negative'
  },
  137: {
    spa: 'SE_LimitEffect',
    effectName: 'Limit:  Effect',
    description: 'Limit',
    base: 'spell effect ID',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: 'Include set value to positive | Exclude set value to negative'
  },
  138: {
    spa: 'SE_LimitSpellType',
    effectName: 'Limit:  SpellType',
    description: 'Limit',
    base: '0=Detrimental,  1=Beneficial',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: ''
  },
  139: {
    spa: 'SE_LimitSpell',
    effectName: 'Limit:  Spell',
    description: 'Limit',
    base: 'spell ID',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: 'Include set value to positive | Exclude set value to negative'
  },
  140: {
    spa: 'SE_LimitMinDur',
    effectName: 'Limit:  Min Duration',
    description: 'Limit',
    base: 'tics',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: 'Set duration in tics, 1 tick is 6 seconds of game time'
  },
  141: {
    spa: 'SE_LimitInstant',
    effectName: 'Limit:  Instant spells only',
    description: 'Limit',
    base: '0=Exclude if Instant, 1=Allow only if Instant',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: ''
  },
  142: {
    spa: 'SE_LimitMinLevel',
    effectName: 'Limit:  Min Level',
    description: 'Limit',
    base: 'level',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: ''
  },
  143: {
    spa: 'SE_LimitCastTimeMin',
    effectName: 'Limit:  Min Cast Time',
    description: 'Limit',
    base: 'milliseconds',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: ''
  },
  144: {
    spa: 'SE_LimitCastTimeMax',
    effectName: 'Limit:  Max Cast Time',
    description: 'Limit',
    base: 'milliseconds',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: ''
  },
  145: {
    spa: 'SE_Teleport2',
    effectName: 'Banish',
    description: 'Teleport',
    base: 'coordinate(x,y,z,h)',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: 'Set \'Teleport Zone\' field to zone short name  OR set to \'same\' to teleport within same zone. To set all xyzh cooridinates, you have use the following. Use this effectid only once in first effect slot . Cooridinates defined as effect_base_value1=x effect_base_value2=y effect_base_value3=z effect_base_value4=h'
  },
  146: {
    spa: 'SE_ElectricityResist',
    effectName: 'Portal Locations',
    description: 'Not Implementd',
    base: 'pending',
    limit: 'pending',
    value: 'pending',
    bonusCalc: '',
    notes: ''
  },
  147: {
    spa: 'SE_PercentalHeal',
    effectName: 'Percent HP Heal',
    description: 'Hit Points',
    base: 'percentage',
    limit: 'none',
    value: 'max amount of hit points',
    bonusCalc: '',
    notes: 'Negative base value for damage | Positive base value for healing'
  },
  148: {
    spa: 'SE_StackingCommand_Block',
    effectName: 'Stacking:  Block',
    description: 'Stacking',
    base: 'spell effect id',
    limit: 'none',
    value: 'Block if less than this value.',
    bonusCalc: '',
    notes: 'Formula - 201 = Slot to block | Exampe: AC buff with a value of 500. You want to block any other AC spell from going into Slot 1 that is less than 1000 AC. Slot 1 : Effect SE_AC(SPA 1) with Base=500, Slot 2:  Effect SE_StackingCommand_Block (SPA 148) with Base= SE_AC(SPA 1) Max = 1000 and Formula = 201 (apply to the first spell slot--remember that we use base 0)'
  },
  149: {
    spa: 'SE_StackingCommand_Overwrite',
    effectName: 'Stacking:  Overwrite',
    description: 'Stacking',
    base: 'spell effect id',
    limit: 'none',
    value: 'Overwrite if less than this value.',
    bonusCalc: '',
    notes: 'Formula - 201 = Slot to block | Exampe: AC buff with a value of 1200. You want to overwrite any other AC spell in Slot 1 that is less than 1000 AC. Slot 1 : Effect SE_AC(SPA 1) with Base=1200 , Slot 2:  Effect SE_StackingCommand_Overwrite (SPA 149) with Base= SE_AC(SPA 1) Max = 1000 and Formula = 201 (apply to the first spell slot--remember that we use base 0)'
  },
  150: {
    spa: 'SE_DeathSave',
    effectName: 'Death Save',
    description: 'Death Save',
    base: '1 = 300 HP Healed,  2 = 8000 HP Healed',
    limit: 'min target level to apply override heal amount',
    value: 'override heal amount',
    bonusCalc: '',
    notes: 'If max value is set as heal amount this value will be used instead as the heal amount if the owner of the buff is the mininum level specified in limit field. Chance to receive a heal is determined by the owner of the buffs Charisma. [Chance = ((Charisma * 3) +1) / 10) ] . SPA 277 gives a second chance to be healed if you fail.'
  },
  151: {
    spa: 'SE_SuspendPet',
    effectName: 'Suspend Pet',
    description: 'Pet',
    base: 'save type',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: 'Save Types, 0 = save pet with no buffs or equipment, 1 = save pet with no buffs or equipment, 2 = unknown.  SPA 308 allows for suspended pets to be resummoned after zoning.'
  },
  152: {
    spa: 'SE_TemporaryPets',
    effectName: 'Summon a Pet Swarm',
    description: 'Pet',
    base: 'amount of pets',
    limit: 'none',
    value: 'duration seconds',
    bonusCalc: '',
    notes: 'Set \'Teleport Zone\' field to be the same as \'type\' field in of the pet you want in the pets table.'
  },
  153: {
    spa: 'SE_BalanceHP',
    effectName: 'Balance Party HP',
    description: 'Hit Points',
    base: 'percent modifier',
    limit: 'max HP taken from player',
    value: 'none',
    bonusCalc: '',
    notes: 'Positive base value increases damage being  distrubuted | Negative base value decreases the damage being distributed'
  },
  154: {
    spa: 'SE_DispelDetrimental',
    effectName: 'Dispel Detrimental',
    description: 'Dispel',
    base: 'percent chance x 10',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: 'Actual percent chance is calculated as base / 10'
  },
  155: {
    spa: 'SE_SpellCritDmgIncrease',
    effectName: 'Spell Critical Damage',
    description: 'Spell Modifier',
    base: 'percent modifier',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: ''
  },
  156: {
    spa: 'SE_IllusionCopy',
    effectName: 'Illusion:  Target',
    description: 'Illusion',
    base: 'Seen 0,1,30',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: 'Unknown what base values represent.'
  },
  157: {
    spa: 'SE_SpellDamageShield',
    effectName: 'Spell Damage Shield',
    description: 'Damage Shield',
    base: 'amount damage shield (negative)',
    limit: 'none',
    value: 'unknown',
    bonusCalc: '',
    notes: 'Spells must have \'feedbackable\' field set to a value otherwise they will not be affected by spell damage shields.'
  },
  158: {
    spa: 'SE_Reflect',
    effectName: 'Reflect Spell',
    description: 'Beneficial Uncategorized',
    base: 'percent chance',
    limit: 'resist modifier',
    value: 'percent of base damage modifier',
    bonusCalc: '',
    notes: 'Spells must have \'reflectable\' field set to a value otherwise they will not be reflected. Resist modifer, positive value reduces the resist rate, negative value increases the resist rate. Percent of base damage modifer, max greater than 100 for damage mod (120 = 20 pct increase in damage) |  max less than 100 for damage mod (80 = 20 pct decrease in damage)'
  },
  159: {
    spa: 'SE_AllStats',
    effectName: 'All Stats',
    description: 'STATS',
    base: 'amount',
    limit: 'none',
    value: 'max amount',
    bonusCalc: '',
    notes: 'Effect currently handled entirely client side.'
  },
  160: {
    spa: 'SE_MakeDrunk',
    effectName: 'Drunk',
    description: 'Detrimental Uncategorized',
    base: 'amount',
    limit: 'none',
    value: 'max amount',
    bonusCalc: '',
    notes: ''
  },
  161: {
    spa: 'SE_MitigateSpellDamage',
    effectName: 'Mitigate Spell Damage Rune',
    description: 'Rune',
    base: 'percent mitigation',
    limit: 'max damage absorbed per hit',
    value: 'rune amount',
    bonusCalc: '',
    notes: 'Special: If this effect is placed on item as worn effect or as an AA, it will provide stackable percent spell mitigation for the base value.'
  },
  162: {
    spa: 'SE_MitigateMeleeDamage',
    effectName: 'Mitigate Melee Damage Rune',
    description: 'Rune',
    base: 'percent mitigation',
    limit: 'max damage absorbed per hit',
    value: 'rune amount',
    bonusCalc: '',
    notes: ''
  },
  163: {
    spa: 'SE_NegateAttacks',
    effectName: 'Absorb Damage',
    description: 'Rune',
    base: 'amount of blocked hits',
    limit: 'none',
    value: 'max amount of damage blocked per hit',
    bonusCalc: '',
    notes: ''
  },
  164: {
    spa: 'SE_AppraiseLDonChest',
    effectName: 'Sense LDoN Chest',
    description: 'LDON',
    base: '1',
    limit: 'none',
    value: 'skill check value',
    bonusCalc: '',
    notes: ''
  },
  165: {
    spa: 'SE_DisarmLDoNTrap',
    effectName: 'Disarm LDoN Trap',
    description: 'LDON',
    base: '1',
    limit: 'none',
    value: 'skill check value',
    bonusCalc: '',
    notes: ''
  },
  166: {
    spa: 'SE_UnlockLDoNChest',
    effectName: 'Unlock LDoN Chest',
    description: 'LDON',
    base: '1',
    limit: 'none',
    value: 'skill check value',
    bonusCalc: '',
    notes: ''
  },
  167: {
    spa: 'SE_PetPowerIncrease',
    effectName: 'Focus: Pet Power',
    description: 'FOCUS',
    base: 'power value',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: 'Pet power can be scaled automatically if \'petpower\' field in pets table is set to 0 or -1, if the power field is set to anything it will look to find the cooresponding pet in the table with same power for that \'type\'.'
  },
  168: {
    spa: 'SE_MeleeMitigation',
    effectName: 'Defensive',
    description: 'Defensive Bonus',
    base: 'percent modifer',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: 'Negative base value decreases damage taken | Positive base value increases damage taken'
  },
  169: {
    spa: 'SE_CriticalHitChance',
    effectName: 'Critical Melee Chance',
    description: 'Offensive Bonus',
    base: 'percent modifer',
    limit: 'skill type (-1 = all skill types)',
    value: 'none',
    bonusCalc: '',
    notes: ''
  },
  170: {
    spa: 'SE_SpellCritChance',
    effectName: 'Spell Critical Chance',
    description: 'Spell Modifer',
    base: 'percent modifier',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: 'Must have a chance to perform critical hits in order to have a chance to crippling blow.'
  },
  171: {
    spa: 'SE_CrippBlowChance',
    effectName: 'Crippling Blow Chance',
    description: 'Offensive Bonus',
    base: 'percent modifer',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: ''
  },
  172: {
    spa: 'SE_AvoidMeleeChance',
    effectName: 'Evasion',
    description: 'Defensive Bonus',
    base: 'percent modifer',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: ''
  },
  173: {
    spa: 'SE_RiposteChance',
    effectName: 'Riposte',
    description: 'Defensive Bonus',
    base: 'percent modifer',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: ''
  },
  174: {
    spa: 'SE_DodgeChance',
    effectName: 'Dodge',
    description: 'Defensive Bonus',
    base: 'percent modifer',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: ''
  },
  175: {
    spa: 'SE_ParryChance',
    effectName: 'Parry',
    description: 'Defensive Bonus',
    base: 'percent modifer',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: ''
  },
  176: {
    spa: 'SE_DualWieldChance',
    effectName: 'Dual Wield',
    description: 'Offensive Bonus',
    base: 'percent modifer',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: ''
  },
  177: {
    spa: 'SE_DoubleAttackChance',
    effectName: 'Double Attack',
    description: 'Offensive Bonus',
    base: 'percent modifer',
    limit: 'none',
    value: 'none',
    bonusCalc: 'Highest',
    notes: 'Positive value will heal you | Negative value will damage you'
  },
  178: {
    spa: 'SE_MeleeLifetap',
    effectName: 'Melee Lifetap',
    description: 'Offensive Bonus',
    base: 'percentage',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: ''
  },
  179: {
    spa: 'SE_AllInstrumentMod',
    effectName: 'All Instrument Modifier',
    description: 'Songs',
    base: 'modifier percentage',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: ''
  },
  180: {
    spa: 'SE_ResistSpellChance',
    effectName: 'Resist Spell Chance',
    description: 'Resist',
    base: 'percent modifer',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: ''
  },
  181: {
    spa: 'SE_ResistFearChance',
    effectName: 'Resist Fear Spell Chance',
    description: 'Resist',
    base: 'percent modifer',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: ''
  },
  182: {
    spa: 'SE_HundredHands',
    effectName: 'Attack Delay Reducation',
    description: 'Offensive Bonus',
    base: 'percent modifier',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: 'Negative value reduces delay, example -115 is calculated as a 15 percent reduction (-115/100). Positive value increases delay, 300 is calculated as a 30 percent increase in delay'
  },
  183: {
    spa: 'SE_MeleeSkillCheck',
    effectName: 'Melee Skill Chance',
    description: 'Offensive Bonus',
    base: 'percent chance',
    limit: 'skill type (-1 = all skill types)',
    value: 'none',
    bonusCalc: '',
    notes: 'No longer used on live. It provides no benefits on eqemu.'
  },
  184: {
    spa: 'SE_HitChance',
    effectName: 'Chance to Hit',
    description: 'Offensive Bonus',
    base: 'percent modifer',
    limit: 'skill type (-1 = all skill types)',
    value: 'none',
    bonusCalc: '',
    notes: ''
  },
  185: {
    spa: 'SE_DamageModifier',
    effectName: 'Skills Damage Modifier',
    description: 'Offensive Bonus',
    base: 'percent modifer',
    limit: 'skill type (-1 = all skill types)',
    value: 'none',
    bonusCalc: '',
    notes: ''
  },
  186: {
    spa: 'SE_MinDamageModifier',
    effectName: 'Skills Minimum Damage Modifier',
    description: 'Offensive Bonus',
    base: 'percent modifer',
    limit: 'skill type (-1 = all skill types)',
    value: 'none',
    bonusCalc: '',
    notes: ''
  },
  187: {
    spa: 'SE_BalanceMana',
    effectName: 'Balance Party Mana',
    description: 'Mana',
    base: 'percent modifer',
    limit: 'max mana taken from player',
    value: 'none',
    bonusCalc: '',
    notes: 'Positive base value increases damage being distributed | Negative base value decreases the damage being distributed'
  },
  188: {
    spa: 'SE_IncreaseBlockChance',
    effectName: 'Chance to block',
    description: 'Defensive Bonus',
    base: 'percent modifer',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: ''
  },
  189: {
    spa: 'SE_CurrentEndurance',
    effectName: 'Endurance',
    description: 'Endurance',
    base: 'amount',
    limit: 'none',
    value: 'max amount',
    bonusCalc: '',
    notes: ''
  },
  190: {
    spa: 'SE_EndurancePool',
    effectName: 'Max Endurance',
    description: 'Endurance',
    base: 'amount',
    limit: 'none',
    value: 'max amount',
    bonusCalc: '',
    notes: ''
  },
  191: {
    spa: 'SE_Amnesia',
    effectName: 'Amnesia',
    description: 'Detrimental Uncategorized',
    base: '1 (increase for stacking overwrite)',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: ''
  },
  192: {
    spa: 'SE_Hate',
    effectName: 'Hate',
    description: 'Aggro',
    base: 'amount',
    limit: 'none',
    value: 'max amount',
    bonusCalc: '',
    notes: ''
  },
  193: {
    spa: 'SE_SkillAttack',
    effectName: 'Skill Attack',
    description: 'Instant Combat',
    base: 'weapon damage',
    limit: 'chance to hit modifier',
    value: 'unknown',
    bonusCalc: '',
    notes: 'Skill used to perform combat round is determined by the \'skill\' field in spells table.'
  },
  194: {
    spa: 'SE_FadingMemories',
    effectName: 'Fade',
    description: 'Aggro',
    base: 'success chance',
    limit: 'max level (ROF2 era)',
    value: 'max level (modern era)',
    bonusCalc: '',
    notes: 'Support for max level requires Rule (Spells, UseFadingMemoriesMaxLevel) to be true. If used from limit field, then it set as the level, ie. max level of 75 would use limit value of 75. If set from max field, max level 75 would use max value of 1075, if you want to set it so it checks a level range above the spell target then for it to only work on mobs 5 levels or below you set max value to 5.'
  },
  195: {
    spa: 'SE_StunResist',
    effectName: 'Stun Resist',
    description: 'Defensive Bonus',
    base: 'percent modifier',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: ''
  },
  196: {
    spa: 'SE_StrikeThrough',
    effectName: 'Strikethrough',
    description: 'Offensive Bonus',
    base: 'percent modifier',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: ''
  },
  197: {
    spa: 'SE_SkillDamageTaken',
    effectName: 'Skill Damage Taken',
    description: 'Defensive Bonus',
    base: 'percent modifier',
    limit: 'skill type (-1 = all skill types)',
    value: 'none',
    bonusCalc: '',
    notes: ''
  },
  198: {
    spa: 'SE_CurrentEnduranceOnce',
    effectName: 'Instant Endurance',
    description: 'Endurance',
    base: 'amount',
    limit: 'none',
    value: 'max amount',
    bonusCalc: '',
    notes: 'Negative base value decreases damage taken | Positive base value increases damage taken'
  },
  199: {
    spa: 'SE_Taunt',
    effectName: 'Taunt',
    description: 'Aggro',
    base: 'taunt success chance',
    limit: 'amount hate added',
    value: '',
    bonusCalc: '',
    notes: ''
  },
  200: {
    spa: 'SE_ProcChance',
    effectName: 'Worn Proc Chance',
    description: 'Offensive Bonus',
    base: 'percent modifier',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: ''
  },
  201: {
    spa: 'SE_RangedProc',
    effectName: 'Ranged Proc',
    description: 'Procs',
    base: 'spellid',
    limit: 'rate modifer',
    value: 'none',
    bonusCalc: '',
    notes: ''
  },
  202: {
    spa: 'SE_IllusionOther',
    effectName: 'Project Illusion',
    description: 'Illusion',
    base: 'none',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: ''
  },
  203: {
    spa: 'SE_MassGroupBuff',
    effectName: 'Mass Group Buff',
    description: 'Beneficial Uncategorized',
    base: '1',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: ''
  },
  204: {
    spa: 'SE_GroupFearImmunity',
    effectName: 'Group Fear Immunity',
    description: 'Resist',
    base: 'duration',
    limit: 'none',
    value: 'none',
    bonusCalc: '0',
    notes: 'Duration is calculated as base value * 10. Thus, value of 1 would be 10 seconds. This is not a buff and gives no icon.'
  },
  205: {
    spa: 'SE_Rampage',
    effectName: 'AE Rampage',
    description: 'Instant Combat',
    base: 'number of attack rounds',
    limit: 'max entities hit per round',
    value: 'aoe range override',
    bonusCalc: '',
    notes: 'On live base is always set to 1, if more than one attack is a spell it uses this SPA in multiple slots. Limit value can be used to set a max amount of entities able to be attacked per round.'
  },
  206: {
    spa: 'SE_AETaunt',
    effectName: 'AE Taunt',
    description: 'Aggro',
    base: 'added hate',
    limit: 'none',
    value: 'aoe range override',
    bonusCalc: '',
    notes: ''
  },
  207: {
    spa: 'SE_FleshToBone',
    effectName: 'Flesh to Bone',
    description: 'Utility Uncategorized',
    base: '1',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: ''
  },
  208: {
    spa: 'SE_PurgePoison',
    effectName: 'Purge Poison',
    description: 'Not Used',
    base: '',
    limit: '',
    value: '',
    bonusCalc: '',
    notes: ''
  },
  209: {
    spa: 'SE_DispelBeneficial',
    effectName: 'Dispel Beneficial',
    description: 'Dispel',
    base: 'percent chance x 10',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: 'Actual percent chance is calculated as base / 10'
  },
  210: {
    spa: 'SE_PetShield',
    effectName: 'Pet Shield',
    description: 'Pet',
    base: 'Duration multiplier 1=12 seconds, 2=24 ect',
    limit: 'mitigation on pet owner override',
    value: 'mitigation on pet overide',
    bonusCalc: '',
    notes: 'Special: limit and max values are not on live, they can be used to give mitigation penalties or bonuses to shielder or shielded.'
  },
  211: {
    spa: 'SE_AEMelee',
    effectName: 'AE Melee',
    description: 'Instant Combat',
    base: 'Duration multiplier 1=12 seconds, 2=24 ect',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: 'Only implemented for clients.'
  },
  212: {
    spa: 'SE_FrenziedDevastation',
    effectName: 'Frenzied Devastation',
    description: 'Spell Bonus',
    base: '1',
    limit: 'chance modifier',
    value: 'none',
    bonusCalc: '',
    notes: 'Live no longer uses the effect in this way. It is now a focus effect.'
  },
  213: {
    spa: 'SE_PetMaxHP',
    effectName: 'Pet Max HP',
    description: 'Pet Bonuses From Owner',
    base: 'percent modifier',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: 'This effect goes on the pet owner and then the benefit is applied to the pet.'
  },
  214: {
    spa: 'SE_MaxHPChange',
    effectName: 'Change Max HP',
    description: 'HP',
    base: 'percent modifier',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: 'Base value is divided by 100 to get actual percentage. Example, for 10 percent max HP increase, base value should be 1000.'
  },
  215: {
    spa: 'SE_PetAvoidance',
    effectName: 'Pet Avoidance',
    description: 'Pet Bonuses From Owner',
    base: 'percent modifier',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: 'This effect goes on the pet owner and then the benefit is applied to the pet.'
  },
  216: {
    spa: 'SE_Accuracy',
    effectName: 'Accuracy',
    description: 'Offensive Bonus',
    base: 'amount accuracy',
    limit: 'skill type (-1 = all skill types)',
    value: 'none',
    bonusCalc: '',
    notes: 'AA version of this is not skill limited. Differs from SPA 184, which is a multiplier of your total accuracy.'
  },
  217: {
    spa: 'SE_HeadShot',
    effectName: 'Headshot',
    description: 'Fatal Procs',
    base: 'percent chance',
    limit: 'damage amount',
    value: 'none',
    bonusCalc: '',
    notes: 'Used with SPA 346 which limits headshot by level and adds a bonus chance.'
  },
  218: {
    spa: 'SE_PetCriticalHit',
    effectName: 'Pet Crit Melee',
    description: 'Pet Bonuses From Owner',
    base: 'percent modifier',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: 'This effect goes on the pet owner and then the benefit is applied to the pet.'
  },
  219: {
    spa: 'SE_SlayUndead',
    effectName: 'Slay Undead',
    description: 'Offensive Bonus',
    base: 'damage percent modifier',
    limit: 'chance',
    value: 'none',
    bonusCalc: '',
    notes: 'Actual chance will be your limit value / 10. Example a 14 percent chance would require limit value of 140. Damage modifier baseline is 100, Base greater than 100 for increased damage  (120 = 20 pct damage increase) |  Base less than 100 for reduced damage (80 = 20 pct damage reduction).'
  },
  220: {
    spa: 'SE_SkillDamageAmount',
    effectName: 'Skill Damage Bonus',
    description: 'Offensive Bonus',
    base: 'amount',
    limit: 'skill type (-1 = all skill types)',
    value: 'none',
    bonusCalc: '',
    notes: ''
  },
  221: {
    spa: 'SE_Packrat',
    effectName: 'Reduce Weight',
    description: 'Utility Uncategorized',
    base: 'percent modifier',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: ''
  },
  222: {
    spa: 'SE_BlockBehind',
    effectName: 'Block Behind',
    description: 'Defensive Bonus',
    base: 'percent modifer',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: ''
  },
  223: {
    spa: 'SE_DoubleRiposte',
    effectName: 'Double Riposte',
    description: 'Offensive Bonus',
    base: 'percent chance',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: 'No longer used on live.'
  },
  224: {
    spa: 'SE_GiveDoubleRiposte',
    effectName: 'Additional Riposte',
    description: 'Offensive Bonus',
    base: 'percent chance',
    limit: 'skill type',
    value: 'none',
    bonusCalc: '',
    notes: 'If limit value is set you can riposte using a specific special attack skill, like \'flying kick\'. You can not have multiple skills that can riposte, thus limited to use in only one effect.'
  },
  225: {
    spa: 'SE_GiveDoubleAttack',
    effectName: 'Double Attack',
    description: 'Offensive Bonus',
    base: 'percent chance or modifer',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: ''
  },
  226: {
    spa: 'SE_TwoHandBash',
    effectName: 'Two Hand bash',
    description: 'Offensive Bonus',
    base: '1',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: 'Handled client side.'
  },
  227: {
    spa: 'SE_ReduceSkillTimer',
    effectName: 'Base Refresh Timer',
    description: 'Timers',
    base: 'time seconds (positive)',
    limit: 'skill type',
    value: 'none',
    bonusCalc: '',
    notes: ''
  },
  228: {
    spa: 'SE_ReduceFallDamage',
    effectName: 'Reduce Fall Dmg',
    description: 'Beneficial Uncategorized',
    base: 'percent modifer',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: ''
  },
  229: {
    spa: 'SE_PersistantCasting',
    effectName: 'Cast Through Stun',
    description: 'Beneficial Uncategorized',
    base: 'percent chance',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: ''
  },
  230: {
    spa: 'SE_ExtendedShielding',
    effectName: 'Increase Shield Distance',
    description: 'Shield Ability',
    base: 'distance',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: ''
  },
  231: {
    spa: 'SE_StunBashChance',
    effectName: 'Stun Bash Chance',
    description: 'Beneficial Uncategorized',
    base: 'percent modifer',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: ''
  },
  232: {
    spa: 'SE_DivineSave',
    effectName: 'Divine Save',
    description: 'Divine Save',
    base: 'percent chance',
    limit: 'spellid',
    value: 'none',
    bonusCalc: '',
    notes: 'This effect triggers upon death, where base value gives you percent chance to cast Touch of the Divine which is an Invulnerability, heal, HoT and purify effect. Limit value can be used to add an additional spell being applied on death, usually this is a heal.'
  },
  233: {
    spa: 'SE_Metabolism',
    effectName: 'Metabolism',
    description: 'Utility Uncategorized',
    base: 'percent modifer',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: 'Positive value decreases consumption rate | Negative value increase consumption rate.'
  },
  234: {
    spa: 'SE_ReduceApplyPoisonTime',
    effectName: 'Poison Mastery',
    description: 'Timers',
    base: 'time',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: 'Reducation time calculated as base /1000. Example, 2.5 second reduction would be a base value of 2500.'
  },
  235: {
    spa: 'SE_ChannelChanceSpells',
    effectName: 'Focus Channelling',
    description: 'Casting',
    base: 'percent modifer',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: 'No longer used on live.'
  },
  236: {
    spa: 'SE_FreePet',
    effectName: 'Free Pet',
    description: 'Not Used',
    base: '',
    limit: '',
    value: '',
    bonusCalc: '',
    notes: ''
  },
  237: {
    spa: 'SE_GivePetGroupTarget',
    effectName: 'Pet Affinity',
    description: 'Pet Bonuses From Owner',
    base: '1',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: 'This effect goes on the pet owner and then the benefit is applied to the pet.'
  },
  238: {
    spa: 'SE_IllusionPersistence',
    effectName: 'Permanent Illusion',
    description: 'Illusion',
    base: '1 or 2',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: '1=Persist through zoning, 2=Persist through death'
  },
  239: {
    spa: 'SE_FeignedCastOnChance',
    effectName: 'Feign Death Through Spell Hit ',
    description: 'Beneficial Uncategorized',
    base: 'percent chance',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: 'If spell is resisted your chance is multipled by two.'
  },
  240: {
    spa: 'SE_StringUnbreakable',
    effectName: 'String Unbreakable',
    description: 'Not Used',
    base: '',
    limit: '',
    value: '',
    bonusCalc: '',
    notes: ''
  },
  241: {
    spa: 'SE_ImprovedReclaimEnergy',
    effectName: 'Improve Reclaim Energy',
    description: 'Pet',
    base: 'percent modifier',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: ''
  },
  242: {
    spa: 'SE_IncreaseChanceMemwipe',
    effectName: 'Increase Chance Memwipe',
    description: 'Memory Blur',
    base: 'percent modifier',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: 'Actual chance to memory blur is much higher than the memory blurs spells base value, caster level and CHA modifiers are added get the final calculated percent chance. This effect modifiers that final percent chance.'
  },
  243: {
    spa: 'SE_CharmBreakChance',
    effectName: 'Charm Break Chance',
    description: 'Charm',
    base: 'percent modifier',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: ''
  },
  244: {
    spa: 'SE_RootBreakChance',
    effectName: 'Root Break Chance',
    description: 'Movement Speed',
    base: 'percent modifier',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: 'Modifies the base line root break chance. The benefit is given to any player casting on that NPC with the root, opposed to only the caster of the root.'
  },
  245: {
    spa: 'SE_TrapCircumvention',
    effectName: 'Trap Circumvention',
    description: 'Traps',
    base: 'percent modifier',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: ''
  },
  246: {
    spa: 'SE_SetBreathLevel',
    effectName: 'Lung Capacity',
    description: 'Water Breathing',
    base: 'percent modifier',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: 'Should work client side. No server side support.'
  },
  247: {
    spa: 'SE_RaiseSkillCap',
    effectName: 'Increase SkillCap',
    description: 'Skill Caps',
    base: 'amount',
    limit: 'skill type',
    value: 'none',
    bonusCalc: '',
    notes: ''
  },
  248: {
    spa: 'SE_SecondaryForte',
    effectName: 'Extra Specialization',
    description: 'Skill Caps',
    base: '100',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: 'Changing base value will not alter this effect.'
  },
  249: {
    spa: 'SE_SecondaryDmgInc',
    effectName: 'Offhand Min Damage Bonus',
    description: 'Offensive Bonus',
    base: 'percent modifier',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: ''
  },
  250: {
    spa: 'SE_SpellProcChance',
    effectName: 'Spell Proc Chance',
    description: 'Procs',
    base: 'percent modifier',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: ''
  },
  251: {
    spa: 'SE_ConsumeProjectile',
    effectName: 'Endless Quiver',
    description: 'Utility Uncategorized',
    base: 'percent chance',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: ''
  },
  252: {
    spa: 'SE_FrontalBackstabChance',
    effectName: 'Backstab from Front',
    description: 'Backstab',
    base: 'percent chance',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: ''
  },
  253: {
    spa: 'SE_FrontalBackstabMinDmg',
    effectName: 'Chaotic Stab',
    description: 'Backstab',
    base: '1',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: ''
  },
  254: {
    spa: 'SE_Blank',
    effectName: 'No Spell',
    description: 'Utility Uncategorized',
    base: 'none',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: 'Do not replace this effect.'
  },
  255: {
    spa: 'SE_ShieldDuration',
    effectName: 'Shielding Duration',
    description: 'Shield Ability',
    base: 'seconds',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: ''
  },
  256: {
    spa: 'SE_ShroudofStealth',
    effectName: 'Shroud Of Stealth',
    description: 'Invisibility',
    base: '1',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: ''
  },
  257: {
    spa: 'SE_PetDiscipline',
    effectName: 'Give Pet Hold',
    description: 'Pet',
    base: '1',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: 'SPA 267 with a limit value of 15 is required now to obtain pet /hold. '
  },
  258: {
    spa: 'SE_TripleBackstab',
    effectName: 'Triple Backstab',
    description: 'Backstab',
    base: 'percent chance',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: ''
  },
  259: {
    spa: 'SE_CombatStability',
    effectName: 'AC Softcap Limit',
    description: 'Defensive Bonus',
    base: 'amount',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: ''
  },
  260: {
    spa: 'SE_AddSingingMod',
    effectName: 'Instrument Modifier',
    description: 'Songs',
    base: 'percent modifier',
    limit: 'Item Type ID',
    value: 'none',
    bonusCalc: '',
    notes: 'Item Type IDs, 23=Woodwind, 24=Strings, 25=Brass, 26=Percussions, 50=Singing, 51=All instruments'
  },
  261: {
    spa: 'SE_SongModCap',
    effectName: 'Song Cap',
    description: 'Songs',
    base: 'amount',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: 'Song cap is set in spells table as field \'song_cap\'. Not used on live.'
  },
  262: {
    spa: 'SE_RaiseStatCap',
    effectName: 'Raise Stat Cap',
    description: 'Stat',
    base: 'amount',
    limit: 'stat type id',
    value: 'none',
    bonusCalc: '',
    notes: 'Stat type id, STR=0, STA=1, AGI=2, DEX=3, WIS=4, INT=5, CHA=6, MR=7, CR=8, FR=9, PR=10, DR=11, COR=12'
  },
  263: {
    spa: 'SE_TradeSkillMastery',
    effectName: 'Tradeskill Masteries',
    description: 'Skill Caps',
    base: 'amount of skills that can be raised (max=6)',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: ''
  },
  264: {
    spa: 'SE_HastenedAASkill',
    effectName: 'Reduce AA Timer',
    description: 'Timers',
    base: 'reducation amount seconds',
    limit: 'aa id',
    value: 'none',
    bonusCalc: '',
    notes: 'This can be only set as an AA ability. '
  },
  265: {
    spa: 'SE_MasteryofPast',
    effectName: 'No Fizzle',
    description: 'Casting',
    base: 'level',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: ''
  },
  266: {
    spa: 'SE_ExtraAttackChance',
    effectName: 'Add Extra Attack: 2H Primary',
    description: 'Offensive Bonus',
    base: 'percent chance',
    limit: 'number of attacks',
    value: 'none',
    bonusCalc: '',
    notes: ''
  },
  267: {
    spa: 'SE_AddPetCommand',
    effectName: 'Add Pet Commands',
    description: 'Pet',
    base: '1',
    limit: 'pet command type',
    value: 'none',
    bonusCalc: '',
    notes: 'Full list of command types found in common.h'
  },
  268: {
    spa: 'SE_ReduceTradeskillFail',
    effectName: 'Tradeskill Failure Rate',
    description: 'Utility Uncategorized',
    base: 'chance modifier',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: ''
  },
  269: {
    spa: 'SE_MaxBindWound',
    effectName: 'Bandage Percent Limit',
    description: 'Bind Wound',
    base: 'percent modifier',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: ''
  },
  270: {
    spa: 'SE_BardSongRange',
    effectName: 'Bard Song Range',
    description: 'Songs',
    base: 'percent modifier',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: ''
  },
  271: {
    spa: 'SE_BaseMovementSpeed',
    effectName: 'Base Run Speed',
    description: 'Movement Speed',
    base: 'percent modifier',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: 'Does not stack with run speed modifiers.'
  },
  272: {
    spa: 'SE_CastingLevel2',
    effectName: 'Casting Level',
    description: 'Casting',
    base: 'level amount',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: 'Live decription: This affects, spells that get stronger or last longer based on your level, stacking priority on targets, likelihood that spells that dispel effects will succeed, likelihood that spells that cure blindness will succeed, likelihood that spells that sense, disarm, or pick locked traps will succeed.'
  },
  273: {
    spa: 'SE_CriticalDoTChance',
    effectName: 'Critical DoT',
    description: 'Spell Bonus',
    base: 'percent chance',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: ''
  },
  274: {
    spa: 'SE_CriticalHealChance',
    effectName: 'Critical Heal',
    description: 'Spell Bonus',
    base: 'percent chance',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: ''
  },
  275: {
    spa: 'SE_CriticalMend',
    effectName: 'Critical Mend',
    description: 'Beneficial Uncategorized',
    base: 'percent chance',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: ''
  },
  276: {
    spa: 'SE_Ambidexterity',
    effectName: 'Dual Wield Skill Amount',
    description: 'Offensive Bonus',
    base: 'amount',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: ''
  },
  277: {
    spa: 'SE_UnfailingDivinity',
    effectName: 'Extra DI Chance',
    description: 'Death Save',
    base: 'heal modifier percent',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: 'This works with Death Save SPA 150.'
  },
  278: {
    spa: 'SE_FinishingBlow',
    effectName: 'Finishing Blow',
    description: 'Fatal Procs',
    base: 'percent chance',
    limit: 'damage amount',
    value: 'none',
    bonusCalc: '',
    notes: 'Actual chance is calculated as base value / 10. Example for 50 percent chance, set base to 500. Use with SPA 440 to set max level of NPC that can be affected by finishing blow.'
  },
  279: {
    spa: 'SE_Flurry',
    effectName: 'Flurry Chance',
    description: 'Offensive Bonus',
    base: 'percent chance',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: ''
  },
  280: {
    spa: 'SE_PetFlurry',
    effectName: 'Pet Flurry Chance',
    description: 'Pet Bonuses From Owner',
    base: 'percent chance',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: 'This effect goes on the pet owner and then the benefit is applied to the pet.'
  },
  281: {
    spa: 'SE_FeignedMinion',
    effectName: 'Give Pet Feign',
    description: 'Not Implemented',
    base: '',
    limit: '',
    value: '',
    bonusCalc: '',
    notes: ''
  },
  282: {
    spa: 'SE_ImprovedBindWound',
    effectName: 'Bandage Amount',
    description: 'Bind Wound',
    base: 'percent modifier',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: ''
  },
  283: {
    spa: 'SE_DoubleSpecialAttack',
    effectName: 'Special Attack Chain',
    description: 'Offensive Bonus',
    base: 'percent chance',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: ''
  },
  284: {
    spa: 'SE_LoHSetHeal',
    effectName: 'LoH Set Heal',
    description: 'Not Used',
    base: '',
    limit: '',
    value: '',
    bonusCalc: '',
    notes: ''
  },
  285: {
    spa: 'SE_NimbleEvasion',
    effectName: 'NoMove Check Sneak',
    description: 'Not Implemented',
    base: '',
    limit: '',
    value: '',
    bonusCalc: '',
    notes: ''
  },
  286: {
    spa: 'SE_FcDamageAmt',
    effectName: 'Focus: Spell Damage Amount',
    description: 'Focus',
    base: 'amount',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: ''
  },
  287: {
    spa: 'SE_SpellDurationIncByTic',
    effectName: 'Focus: Buff Duration by Tics',
    description: 'Focus',
    base: 'duration tics',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: '1 tic = 6 seconds, set base in tics'
  },
  288: {
    spa: 'SE_SkillAttackProc',
    effectName: 'Add Proc From Skill Attack',
    description: 'Cast On Effect',
    base: 'chance percent',
    limit: 'skill type',
    value: 'none',
    bonusCalc: '',
    notes: 'Chance is calculated as base value / 10, example 20 percent chance would be a value of 200. For AA\'s the proc spell ID is the \'spell\'  field used in the aa_ranks table. '
  },
  289: {
    spa: 'SE_CastOnFadeEffect',
    effectName: 'Cast Spell On Fade',
    description: 'Cast On Effect',
    base: 'spellid',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: 'Typically seen on spells that can be cured.'
  },
  290: {
    spa: 'SE_IncreaseRunSpeedCap',
    effectName: 'Movement Cap',
    description: 'Movement',
    base: 'amount',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: ''
  },
  291: {
    spa: 'SE_Purify',
    effectName: 'Purify',
    description: 'Dispel',
    base: 'amount spells removed',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: ''
  },
  292: {
    spa: 'SE_StrikeThrough2',
    effectName: 'Strikethrough (v292)',
    description: 'Offensive Bonus',
    base: 'percent modifier',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: ''
  },
  293: {
    spa: 'SE_FrontalStunResist',
    effectName: 'Frontal Stun Resist',
    description: 'Defensive Bonus',
    base: 'percent chance',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: ''
  },
  294: {
    spa: 'SE_CriticalSpellChance',
    effectName: 'Spell Crit Chance',
    description: 'Spell Bonus',
    base: 'critical chance',
    limit: 'critical damage percent modifier',
    value: 'none',
    bonusCalc: '',
    notes: ''
  },
  295: {
    spa: 'SE_ReduceTimerSpecial',
    effectName: 'Reduce Timer Special',
    description: 'Not Used',
    base: '',
    limit: '',
    value: '',
    bonusCalc: '',
    notes: ''
  },
  296: {
    spa: 'SE_FcSpellVulnerability',
    effectName: 'Focus: Incoming Spell Damage',
    description: 'FOCUS',
    base: 'min percent modifier',
    limit: 'none',
    value: 'max percent modifier',
    bonusCalc: '',
    notes: 'Use random effectiveness if base and max value are defined, where base is always lower end and max the higher end of the random range. If random value not wanted, then only set a base value.'
  },
  297: {
    spa: 'SE_FcDamageAmtIncoming',
    effectName: 'Focus: Incoming Spell Damage Amt',
    description: 'FOCUS',
    base: 'amount',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: ''
  },
  298: {
    spa: 'SE_ChangeHeight',
    effectName: 'Pet Size',
    description: 'Utility Uncategorized',
    base: 'percent modifier',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: ''
  },
  299: {
    spa: 'SE_WakeTheDead',
    effectName: 'Wake the Dead',
    description: 'Utility Uncategorized',
    base: '1',
    limit: 'none',
    value: 'duration seconds',
    bonusCalc: '',
    notes: 'Maximum range for corpse from caster is 250 units.'
  },
  300: {
    spa: 'SE_Doppelganger',
    effectName: 'Doppelganger',
    description: 'Utility Uncategorized',
    base: 'amount of pets',
    limit: 'none',
    value: 'duration seconds',
    bonusCalc: '',
    notes: ''
  },
  301: {
    spa: 'SE_ArcheryDamageModifier',
    effectName: 'Archery Damage Modifer',
    description: 'Offensive Bonus',
    base: 'percent modifier',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: ''
  },
  302: {
    spa: 'SE_FcDamagePctCrit',
    effectName: 'Focus: Spell Damage (v302 before crit)',
    description: 'FOCUS',
    base: 'min percent modifier',
    limit: 'none',
    value: 'max percent modifier',
    bonusCalc: '',
    notes: 'Use random effectiveness if base and max value are defined, where base is always lower end and max the higher end of the random range. If random value not wanted, then only set a base value.'
  },
  303: {
    spa: 'SE_FcDamageAmtCrit',
    effectName: 'Focus: Spell Damage Amt (v303 before crit)',
    description: 'FOCUS',
    base: 'amount',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: ''
  },
  304: {
    spa: 'SE_OffhandRiposteFail',
    effectName: 'Secondary Riposte',
    description: 'Offensive Bonus',
    base: 'percent modifier',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: ''
  },
  305: {
    spa: 'SE_MitigateDamageShield',
    effectName: 'Damage Shield Mitigation',
    description: 'Defensive Bonus',
    base: 'amount',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: 'For spells/items set to positive value will reduce the damage shield amount, for AA\'s set this value to negative for reducation, this is converted to positive in source code. This is how live has it.'
  },
  306: {
    spa: 'SE_ArmyOfTheDead',
    effectName: 'Army of the Dead',
    description: 'Not Implemented',
    base: 'amount of pets',
    limit: 'none',
    value: 'duration seconds',
    bonusCalc: '',
    notes: 'Maximum range for corpse from caster is 250 units.'
  },
  307: {
    spa: 'SE_Appraisal',
    effectName: 'Appraisal',
    description: 'Not Implemented',
    base: '',
    limit: '',
    value: '',
    bonusCalc: '',
    notes: ''
  },
  308: {
    spa: 'SE_ZoneSuspendMinion',
    effectName: 'Zone Suspend Minion',
    description: 'Pet',
    base: '1',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: ''
  },
  309: {
    spa: 'SE_GateCastersBindpoint',
    effectName: 'Gate Caster\'s Bindpoint',
    description: 'Teleport',
    base: 'bind id (Set to 1, 2, or 3)',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: ''
  },
  310: {
    spa: 'SE_ReduceReuseTimer',
    effectName: 'Decrease Reuse Timer',
    description: 'Focus',
    base: 'time ms',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: 'Positive value reduces reuse timer. Note: You can set to negative to increase reuse timer, but client will not display it properly.'
  },
  311: {
    spa: 'SE_LimitCombatSkills',
    effectName: 'Limit:  Combat Skills Not Allowed',
    description: 'LIMIT',
    base: ' 0=Exclude if proc 1=Allow only if proc',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: ''
  },
  312: {
    spa: 'SE_Sanctuary',
    effectName: 'Sanctuary',
    description: 'Aggro',
    base: '1',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: ''
  },
  313: {
    spa: 'SE_ForageAdditionalItems',
    effectName: 'Forage Master',
    description: 'Utility Uncategorized',
    base: 'percent chance',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: ''
  },
  314: {
    spa: 'SE_Invisibility2',
    effectName: 'Improved Invisibility',
    description: 'Invisibility',
    base: 'invisibility level',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: 'Invisibility level determines what level of see invisible can detect it.'
  },
  315: {
    spa: 'SE_InvisVsUndead2',
    effectName: 'Improved Invisibility Vs Undead',
    description: 'Invisibility',
    base: 'invisibility level',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: 'Invisibility level determines what level of see invisible can detect it.'
  },
  316: {
    spa: 'SE_ImprovedInvisAnimals',
    effectName: 'Improved Invisibility Vs Animals',
    description: 'Invisibility',
    base: 'invisibility level',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: 'Invisibility level determines what level of see invisible can detect it.'
  },
  317: {
    spa: 'SE_ItemHPRegenCapIncrease',
    effectName: 'Worn Regen Cap',
    description: 'Caps',
    base: 'amount',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: ''
  },
  318: {
    spa: 'SE_ItemManaRegenCapIncrease',
    effectName: 'Worn Mana Cap',
    description: 'Caps',
    base: 'amount',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: ''
  },
  319: {
    spa: 'SE_CriticalHealOverTime',
    effectName: 'Critical HP Regen',
    description: 'Healing Modifer',
    base: 'percent modifier',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: ''
  },
  320: {
    spa: 'SE_ShieldBlock',
    effectName: 'Shield Block Chance',
    description: 'Defensive Bonus',
    base: 'percent chance',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: ''
  },
  321: {
    spa: 'SE_ReduceHate',
    effectName: 'Reduce Target Hate',
    description: 'Aggro',
    base: 'amount',
    limit: 'none',
    value: 'max amount',
    bonusCalc: '',
    notes: 'Positive value decreases hate. | Negative value increases hate.'
  },
  322: {
    spa: 'SE_GateToHomeCity',
    effectName: 'Gate to Starting City',
    description: 'Teleport',
    base: '1',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: ''
  },
  323: {
    spa: 'SE_DefensiveProc',
    effectName: 'Add Defensive Proc',
    description: 'Procs',
    base: 'spellid',
    limit: 'rate modifer',
    value: 'none',
    bonusCalc: '',
    notes: ''
  },
  324: {
    spa: 'SE_HPToMana',
    effectName: 'HP for Mana',
    description: 'Conversion',
    base: 'conversion rate percent',
    limit: '',
    value: '',
    bonusCalc: '',
    notes: ''
  },
  325: {
    spa: 'SE_NoBreakAESneak',
    effectName: 'No Break AE Sneak',
    description: 'Rogue',
    base: 'percent chance',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: ''
  },
  326: {
    spa: 'SE_SpellSlotIncrease',
    effectName: 'Spell Slots',
    description: 'Utility Uncategorized',
    base: 'amount',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: 'Client has to support the value you use.'
  },
  327: {
    spa: 'SE_MysticalAttune',
    effectName: 'Buff Slots',
    description: 'Utility Uncategorized',
    base: 'amount',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: 'Client has to support the value you use.'
  },
  328: {
    spa: 'SE_DelayDeath',
    effectName: 'Negative HP Limit',
    description: 'HP',
    base: 'amount',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: 'Positive value to increase effective negative hit points.'
  },
  329: {
    spa: 'SE_ManaAbsorbPercentDamage',
    effectName: 'Mana Shield Absorb Damage',
    description: 'Defensive Bonus',
    base: 'mitigation percent',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: ''
  },
  330: {
    spa: 'SE_CriticalDamageMob',
    effectName: 'Critical Melee Damage',
    description: 'Offensive Bonus',
    base: 'percent modifier',
    limit: 'skill type (-1 = all skill types)',
    value: 'none',
    bonusCalc: '',
    notes: ''
  },
  331: {
    spa: 'SE_Salvage',
    effectName: 'Item Recovery',
    description: 'Utility Uncategorized',
    base: 'percent modifier',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: 'Positive values increase chance to salvage | Negative values decrase chance to salvage.'
  },
  332: {
    spa: 'SE_SummonToCorpse',
    effectName: 'Summon to Corpse',
    description: 'Corpse',
    base: 'seen 0 or 1',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: ''
  },
  333: {
    spa: 'SE_CastOnRuneFadeEffect',
    effectName: 'Trigger Spell On Rune Fade',
    description: 'Cast on Effect',
    base: 'spellid',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: 'This effect needs to go on a spell containing a rune effect.'
  },
  334: {
    spa: 'SE_BardAEDot',
    effectName: 'Bard AE Dot',
    description: 'Song',
    base: 'amount',
    limit: 'none',
    value: 'amount max',
    bonusCalc: '',
    notes: ''
  },
  335: {
    spa: 'SE_BlockNextSpellFocus',
    effectName: 'Focus: Block Next Spell',
    description: 'FOCUS',
    base: 'percent chance',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: ''
  },
  336: {
    spa: 'SE_IllusionaryTarget',
    effectName: 'Illusionary Target',
    description: 'Not Used',
    base: '',
    limit: '',
    value: '',
    bonusCalc: '',
    notes: ''
  },
  337: {
    spa: 'SE_PercentXPIncrease',
    effectName: 'Experience',
    description: 'Experience',
    base: 'percent modifier',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: ''
  },
  338: {
    spa: 'SE_SummonAndResAllCorpses',
    effectName: 'Expedient Recovery',
    description: 'Corpse',
    base: 'Seen at 70',
    limit: '',
    value: '',
    bonusCalc: '',
    notes: 'Unknown what base value represents.'
  },
  339: {
    spa: 'SE_TriggerOnCast',
    effectName: 'Focus: Trigger on Cast',
    description: 'FOCUS',
    base: 'percent chance',
    limit: 'spellid',
    value: 'none',
    bonusCalc: '',
    notes: ''
  },
  340: {
    spa: 'SE_SpellTrigger',
    effectName: 'Spell Trigger: Only One Spell Cast',
    description: 'Cast On Effects',
    base: 'percent chance',
    limit: 'spellid',
    value: 'none',
    bonusCalc: '',
    notes: 'When multiple of this effect exist on the same spell, only one spell will be selected from the list to be cast. For best results, the total percent chance should equal 100%. Example, Slot 1: Cast Ice Nuke 20%, Slot2: Cast Fire Nuke 50%, Slot3 Cast Magic Nuke 30%. When the spell is cast, only one of these spells be triggered on the target.'
  },
  341: {
    spa: 'SE_ItemAttackCapIncrease',
    effectName: 'Worn Attack Cap',
    description: 'Caps',
    base: 'amount',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: ''
  },
  342: {
    spa: 'SE_ImmuneFleeing',
    effectName: 'Prevent Flee on Low Health',
    description: 'Movement',
    base: '1',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: ''
  },
  343: {
    spa: 'SE_InterruptCasting',
    effectName: 'Spell Interrupt',
    description: 'Casting',
    base: 'percent chance',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: ''
  },
  344: {
    spa: 'SE_ChannelChanceItems',
    effectName: 'Item Channeling',
    description: 'Casting',
    base: 'percent modifer',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: ''
  },
  345: {
    spa: 'SE_AssassinateLevel',
    effectName: 'Assassinate Max Level',
    description: 'Fatal Procs',
    base: 'max target level',
    limit: 'proc chance bonus',
    value: 'none',
    bonusCalc: '',
    notes: ''
  },
  346: {
    spa: 'SE_HeadShotLevel',
    effectName: 'Headshot Max Level',
    description: 'Fatal Procs',
    base: 'max target level',
    limit: 'proc chance bonus',
    value: 'none',
    bonusCalc: '',
    notes: ''
  },
  347: {
    spa: 'SE_DoubleRangedAttack',
    effectName: 'Double Ranged Attack',
    description: 'Offensive Bonus',
    base: 'percent chance',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: 'Will consume an additional ammo item.'
  },
  348: {
    spa: 'SE_LimitManaMin',
    effectName: 'Limit:  Min Mana',
    description: 'LIMIT',
    base: 'mana amount',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: ''
  },
  349: {
    spa: 'SE_ShieldEquipDmgMod',
    effectName: 'Damage With Shield',
    description: 'Offensive Bonus',
    base: 'percent modifier',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: ''
  },
  350: {
    spa: 'SE_ManaBurn',
    effectName: 'Manaburn',
    description: 'Conversion',
    base: 'max amount of mana drained',
    limit: 'percent of mana converted to damage',
    value: 'none',
    bonusCalc: '',
    notes: 'Limit value if set to negative will result in damage | Limit value if set to positive will result in a heal. Once this affect is applied, you can not apply another mana burn until the buff with this effect fades. Example, if base is set to 1000 and limit is to -50, spell will drain 1000 mana and do -500 damage to the target. Calc: (1000/-50)= -500.'
  },
  351: {
    spa: 'SE_PersistentEffect',
    effectName: 'Persistent Effect',
    description: 'Aura',
    base: 'unknown',
    limit: 'none',
    value: 'unknown',
    bonusCalc: '',
    notes: 'Set \'Teleport Zone\' field to be the same as \'name\' field in of the pet you want in the auras table.'
  },
  352: {
    spa: 'SE_IncreaseTrapCount',
    effectName: 'Trap Count',
    description: 'Utility Beneficial',
    base: 'amount',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: ''
  },
  353: {
    spa: 'SE_AdditionalAura',
    effectName: 'Aura Count',
    description: 'Aura',
    base: 'amount',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: ''
  },
  354: {
    spa: 'SE_DeactivateAllTraps',
    effectName: 'Deactivate All Traps',
    description: 'Not Used',
    base: '',
    limit: '',
    value: '',
    bonusCalc: '',
    notes: ''
  },
  355: {
    spa: 'SE_LearnTrap',
    effectName: 'Learn Trap',
    description: 'Not Used',
    base: '',
    limit: '',
    value: '',
    bonusCalc: '',
    notes: ''
  },
  356: {
    spa: 'SE_ChangeTriggerType',
    effectName: 'Change Trigger Type',
    description: 'Not Used',
    base: '',
    limit: '',
    value: '',
    bonusCalc: '',
    notes: ''
  },
  357: {
    spa: 'SE_FcMute',
    effectName: 'Focus: Mute',
    description: 'FOCUS',
    base: 'chance',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: ''
  },
  358: {
    spa: 'SE_CurrentManaOnce',
    effectName: 'Mana Once',
    description: 'Mana',
    base: 'amount',
    limit: 'none',
    value: 'max amount (use positive value)',
    bonusCalc: '',
    notes: ''
  },
  359: {
    spa: 'SE_PassiveSenseTrap',
    effectName: 'Passive Sense Trap',
    description: 'Not Implemented',
    base: '',
    limit: '',
    value: '',
    bonusCalc: '',
    notes: ''
  },
  360: {
    spa: 'SE_ProcOnKillShot',
    effectName: 'Trigger Spell On Kill Shot',
    description: 'Cast On Effect',
    base: 'percent chance',
    limit: 'spellid',
    value: 'minimum target level',
    bonusCalc: '',
    notes: 'Typical use case is a self only buff when triggered.'
  },
  361: {
    spa: 'SE_SpellOnDeath',
    effectName: 'Trigger Spell On Death',
    description: 'Cast On Effect',
    base: 'percent chance',
    limit: 'spellid',
    value: 'none',
    bonusCalc: '',
    notes: 'Typical use case is casting an area of effect upon death. Placing self only beneficial spells such as heals will not work due to player already being dead.'
  },
  362: {
    spa: 'SE_PotionBeltSlots',
    effectName: 'Potion Belt Slots',
    description: 'Not Implemented',
    base: '',
    limit: '',
    value: '',
    bonusCalc: '',
    notes: ''
  },
  363: {
    spa: 'SE_BandolierSlots',
    effectName: 'Bandolier Slots',
    description: 'Not Implemented',
    base: '',
    limit: '',
    value: '',
    bonusCalc: '',
    notes: ''
  },
  364: {
    spa: 'SE_TripleAttackChance',
    effectName: 'Triple Attack Chance',
    description: 'Offensive Bonus',
    base: 'percent modifier',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: ''
  },
  365: {
    spa: 'SE_ProcOnSpellKillShot',
    effectName: 'Trigger Spell on Spell Kill Shot',
    description: 'Cast On Effect',
    base: 'percent chance',
    limit: 'spellid',
    value: 'none',
    bonusCalc: '',
    notes: 'This effect is typical found on direct damage spells.'
  },
  366: {
    spa: 'SE_GroupShielding',
    effectName: 'Group Shielding',
    description: 'Not Used',
    base: '',
    limit: '',
    value: '',
    bonusCalc: '',
    notes: ''
  },
  367: {
    spa: 'SE_SetBodyType',
    effectName: 'Modify Body Type',
    description: 'Utility Uncategorized',
    base: 'body type',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: ''
  },
  368: {
    spa: 'SE_FactionMod',
    effectName: 'Modify Faction',
    description: 'Not Implemented',
    base: '',
    limit: '',
    value: '',
    bonusCalc: '',
    notes: ''
  },
  369: {
    spa: 'SE_CorruptionCounter',
    effectName: 'Corruption Counter',
    description: 'Cures',
    base: 'amount',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: 'Set to positive values for potency of detrimental spells | Set to negative value for potency of cure spells.'
  },
  370: {
    spa: 'SE_ResistCorruption',
    effectName: 'Corruption Resist',
    description: 'Resist',
    base: 'amout',
    limit: 'none',
    value: 'max amount',
    bonusCalc: '',
    notes: ''
  },
  371: {
    spa: 'SE_AttackSpeed4',
    effectName: 'Attack Speed: Inhibit Melee',
    description: 'Attack Speed',
    base: 'slow amount',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: 'This effect works differently than other slows. Base should always be positive.  Example: (SPA 11) Sha\'s Legacy 65% slow + (SPA 371) Lassitude 25% slow, Sha\'s Legacy is calculated as 100 - 35 = 65% slow, therefore the remaining attack speed is (35). Lassitude will now decrease the remaining value (35) by 25% = 16.5%. The total slowed value on the target would be 65% + 16.5% = 81.25% slow. If SPA 371 is only slow effect on target, then it will be slowed for full base value.'
  },
  372: {
    spa: 'SE_ForageSkill',
    effectName: 'Grant Foraging',
    description: 'Utility Uncategorized',
    base: '1',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: ''
  },
  373: {
    spa: 'SE_CastOnFadeEffectAlways',
    effectName: 'Cast Spell On Fade (v373)',
    description: 'Cast On Effects',
    base: 'spellid',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: ''
  },
  374: {
    spa: 'SE_ApplyEffect',
    effectName: 'Spell Trigger: Apply Each Spell',
    description: 'Cast On Effect',
    base: 'percent chance',
    limit: 'spellid',
    value: 'none',
    bonusCalc: '',
    notes: ''
  },
  375: {
    spa: 'SE_DotCritDmgIncrease',
    effectName: 'Critical DoT Damage',
    description: 'Spell Bonus',
    base: 'percent modifier',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: ''
  },
  376: {
    spa: 'SE_Fling',
    effectName: 'Fling',
    description: 'Not Implemented',
    base: '',
    limit: '',
    value: '',
    bonusCalc: '',
    notes: ''
  },
  377: {
    spa: 'SE_CastOnFadeEffectNPC',
    effectName: 'Cast Spell On Fade (v377)',
    description: 'Cast On Effect',
    base: 'spellid',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: ''
  },
  378: {
    spa: 'SE_SpellEffectResistChance',
    effectName: 'Spell Effect Resist Chance',
    description: 'Defensive Bonus',
    base: 'chance modifier',
    limit: 'spell effect id',
    value: 'none',
    bonusCalc: '',
    notes: ''
  },
  379: {
    spa: 'SE_ShadowStepDirectional',
    effectName: 'Directional Shadowstep',
    description: 'Teleport',
    base: 'distance unit',
    limit: 'direction degrees',
    value: '',
    bonusCalc: '',
    notes: 'This effect is handled client side. Unclear how base value equates to actual in game distance movement. Limit directional values example,  0: Shadowstep Forward,  90: Shadowstep Right,  180:Shadowstep Back,  270: Shadowstep Left'
  },
  380: {
    spa: 'SE_Knockdown',
    effectName: 'Knockback',
    description: 'Detrimental Uncategorized',
    base: 'push up?',
    limit: 'push back?',
    value: 'none',
    bonusCalc: '',
    notes: 'Handled by client.'
  },
  381: {
    spa: 'SE_KnockTowardCaster',
    effectName: 'Fling Target to Caster',
    description: 'Not Implemented',
    base: '',
    limit: '',
    value: '',
    bonusCalc: '',
    notes: '   '
  },
  382: {
    spa: 'SE_NegateSpellEffect',
    effectName: 'Negate Spell Effect',
    description: 'Detrimental Uncategorized',
    base: 'negate spell effect type',
    limit: 'spell effect id',
    value: 'none',
    bonusCalc: '',
    notes: ''
  },
  383: {
    spa: 'SE_SympatheticProc',
    effectName: 'Focus: Proc on Spell Cast',
    description: 'FOCUS',
    base: 'proc rate modifier',
    limit: 'spellid',
    value: 'none',
    bonusCalc: '',
    notes: 'Typically found on item focus effects. Longer cast time spells are adjusted to have higher proc rates.'
  },
  384: {
    spa: 'SE_Leap',
    effectName: 'Fling Caster to Target',
    description: 'Movement',
    base: 'distance',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: ''
  },
  385: {
    spa: 'SE_LimitSpellGroup',
    effectName: 'Limit:  SpellGroup',
    description: 'LIMIT',
    base: 'spellgroup id',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: 'Include set value to positive | Exclude set value to negative'
  },
  386: {
    spa: 'SE_CastOnCurer',
    effectName: 'Trigger Spell on Curer',
    description: 'Cast On Effect',
    base: 'spellid',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: 'This effect goes on the spell that needs curing.'
  },
  387: {
    spa: 'SE_CastOnCure',
    effectName: 'Trigger Spell on Cure ',
    description: 'Cast On Effect',
    base: 'spellid',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: 'This effect goes on the spell that needs curing.'
  },
  388: {
    spa: 'SE_SummonCorpseZone',
    effectName: 'Summon All Corpses',
    description: 'Corpse',
    base: '1',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: ''
  },
  389: {
    spa: 'SE_FcTimerRefresh',
    effectName: 'Focus: Spell Gem Refresh',
    description: 'FOCUS',
    base: '1',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: 'Applied from casted spells only.'
  },
  390: {
    spa: 'SE_FcTimerLockout',
    effectName: 'Focus: Spell Gem Lockout',
    description: 'FOCUS',
    base: 'recast duration milliseconds',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: 'Applied from casted spells only.'
  },
  391: {
    spa: 'SE_LimitManaMax',
    effectName: 'Limit:  Max Mana',
    description: 'LIMIT',
    base: 'mana amount',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: ''
  },
  392: {
    spa: 'SE_FcHealAmt',
    effectName: 'Focus: Healing Amount (v392)',
    description: 'FOCUS',
    base: 'amount',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: ''
  },
  393: {
    spa: 'SE_FcHealPctIncoming',
    effectName: 'Focus: Incoming Healing Effectiveness (v392)',
    description: 'FOCUS',
    base: 'percent modifier',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: ''
  },
  394: {
    spa: 'SE_FcHealAmtIncoming',
    effectName: 'Focus: Incoming Healing Amount',
    description: 'FOCUS',
    base: 'amount',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: ''
  },
  395: {
    spa: 'SE_FcHealPctCritIncoming',
    effectName: 'Focus: Incoming Healing Effectiveness (v395)',
    description: 'FOCUS',
    base: 'percent modifier',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: ''
  },
  396: {
    spa: 'SE_FcHealAmtCrit',
    effectName: 'Focus: Healing Amount (v396 before crit)',
    description: 'FOCUS',
    base: 'amount',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: ''
  },
  397: {
    spa: 'SE_PetMeleeMitigation',
    effectName: 'Pet Mitigation',
    description: 'Pet Bonuses From Owner',
    base: 'percent modifier',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: ''
  },
  398: {
    spa: 'SE_SwarmPetDuration',
    effectName: 'Focus: Swarm Pet Duration',
    description: 'FOCUS',
    base: 'duration ms',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: ''
  },
  399: {
    spa: 'SE_FcTwincast',
    effectName: 'Focus: Twincast Chance',
    description: 'FOCUS',
    base: 'percent chance',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: ''
  },
  400: {
    spa: 'SE_HealGroupFromMana',
    effectName: 'Heal Group From Mana',
    description: 'Conversions',
    base: 'max amount mana drained',
    limit: 'ratio of HP gain per 1 mana drained',
    value: 'none',
    bonusCalc: '',
    notes: 'Ratio is calculated as value / 10, example if you want to heal 13 HP for every 1 mana, set base to 130.'
  },
  401: {
    spa: 'SE_ManaDrainWithDmg',
    effectName: 'Damage for Amount Mana Drained',
    description: 'Conversions',
    base: 'max amount mana drained',
    limit: 'ratio of HP of damage per 1 mana drained',
    value: '',
    bonusCalc: '',
    notes: 'Ratio is calculated as value / 10, example if you want to damage for l 13 HP for every 1 mana, set base to 130.'
  },
  402: {
    spa: 'SE_EndDrainWithDmg',
    effectName: 'Damage for Amount Endurance Drained ',
    description: 'Conversions',
    base: 'max amount endurance drained',
    limit: 'ratio of HP of damage per 1 endurance drained',
    value: '',
    bonusCalc: '',
    notes: 'Ratio is calculated as value / 10, example if you want to damage for l 13 HP for every 1 endurance, set base to 130.'
  },
  403: {
    spa: 'SE_LimitSpellClass',
    effectName: 'Limit:  SpellClass',
    description: 'LIMIT',
    base: 'spell_class id',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: 'Include set value to positive | Exclude set value to negative. Definations for spell_class are not known.'
  },
  404: {
    spa: 'SE_LimitSpellSubclass',
    effectName: 'Limit:  SpellSubclass',
    description: 'LIMIT',
    base: 'spell_subclass id',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: 'Include set value to positive | Exclude set value to negative. Definations for spell_subclass are not known.'
  },
  405: {
    spa: 'SE_TwoHandBluntBlock',
    effectName: 'Staff Block Chance',
    description: 'Defensive Bonus',
    base: 'percent chance',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: ''
  },
  406: {
    spa: 'SE_CastonNumHitFade',
    effectName: 'Trigger Spell on Hit Count Fade ',
    description: 'Cast On Effect',
    base: 'spellid',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: ''
  },
  407: {
    spa: 'SE_CastonFocusEffect',
    effectName: 'Trigger Spell on Focus Effect Success',
    description: 'Cast On Effect',
    base: 'spellid',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: 'This effect goes on a spell with a focus effect. '
  },
  408: {
    spa: 'SE_LimitHPPercent',
    effectName: 'Heal Up To Percent Limit',
    description: 'Resource Cap',
    base: 'percent HP',
    limit: 'amount HP',
    value: 'none',
    bonusCalc: '',
    notes: ''
  },
  409: {
    spa: 'SE_LimitManaPercent',
    effectName: 'Restore Mana Up To Percent Limit',
    description: 'Resource Cap',
    base: 'percent HP',
    limit: 'amount HP',
    value: 'none',
    bonusCalc: '',
    notes: ''
  },
  410: {
    spa: 'SE_LimitEndPercent',
    effectName: 'Restore Endurance Up To Percent Limit',
    description: 'Resource Cap',
    base: 'percent HP',
    limit: 'amount HP',
    value: 'none',
    bonusCalc: '',
    notes: ''
  },
  411: {
    spa: 'SE_LimitClass',
    effectName: 'Limit:  PlayerClass',
    description: 'LIMIT',
    base: 'class bitmask',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: 'The class value in dbase is +1 in relation to item class value, set as you would item for multiple classes.'
  },
  412: {
    spa: 'SE_LimitRace',
    effectName: 'Limit:  Race',
    description: 'LIMIT',
    base: 'race id',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: 'Not used in any known live spells. Use only single race at a time.'
  },
  413: {
    spa: 'SE_FcBaseEffects',
    effectName: 'Focus: Base Spell Value',
    description: 'FOCUS',
    base: 'percent modifier',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: 'Used to set bard instrument modifiers. Can be used to focus many effects otherwise not able to be focused by other methods. If modifying spells that are buffs, value must be intervals of 10% and starting at greater than 10%. For example, if a 10% increase in rune amount is required, set base value to 11, that will be calculated as base value * 1.10, resulting in a 10% increase. For a 20% increase set to 12.'
  },
  414: {
    spa: 'SE_LimitCastingSkill',
    effectName: 'Limit:  CastingSkill',
    description: 'LIMIT',
    base: 'skill type',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: 'Include set value to positive | Exclude set value to negative.'
  },
  415: {
    spa: 'SE_FFItemClass',
    effectName: 'Limit:  ItemClass',
    description: 'LIMIT',
    base: 'item ItemType',
    limit: 'item SubType',
    value: 'item Slots',
    bonusCalc: '',
    notes: 'Not used on live. Details, base: item ItemType (-1 to include for all ItemTypes, -1000 to exclude clicks from getting the focus, or exclude specific SubTypes or Slots if set), limit: item SubType (-1 for all SubTypes), max: item Slots (bitmask of valid slots, -1 ALL slots), See comments in Mob::CalcFocusEffect for more details. Special: Can be used with SPA 310 reduce item click recast and SPA 127, 500,5001 to reduce item click casting time.'
  },
  416: {
    spa: 'SE_ACv2',
    effectName: 'AC (v416)',
    description: 'Stat',
    base: 'amount',
    limit: 'none',
    value: 'amount',
    bonusCalc: '',
    notes: ''
  },
  417: {
    spa: 'SE_ManaRegen_v2',
    effectName: 'Mana Regen (v417)',
    description: 'Mana',
    base: 'amount',
    limit: 'none',
    value: 'max amount (use positive value)',
    bonusCalc: '',
    notes: ''
  },
  418: {
    spa: 'SE_SkillDamageAmount2',
    effectName: 'Skill Damage Bonus (v418)',
    description: 'Offensive Bonus',
    base: 'amount',
    limit: 'skill type (-1 = all skill types)',
    value: 'none',
    bonusCalc: '',
    notes: ''
  },
  419: {
    spa: 'SE_AddMeleeProc',
    effectName: 'Add Melee Proc (v419) ',
    description: 'Procs',
    base: 'spellid',
    limit: 'rate modifer',
    value: 'none',
    bonusCalc: '',
    notes: ''
  },
  420: {
    spa: 'SE_FcLimitUse',
    effectName: 'Focus: Hit Count',
    description: 'FOCUS',
    base: 'percent modifier',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: 'Not used in any known live spells. Hit count is set using field \'numhits\' in spells new.'
  },
  421: {
    spa: 'SE_FcIncreaseNumHits',
    effectName: 'Focus: Hit Count Amount',
    description: 'FOCUS',
    base: 'hit count amount',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: 'Hit count is set using field \'numhits\' in spells new.'
  },
  422: {
    spa: 'SE_LimitUseMin',
    effectName: 'Limit:  Minimum Hit Count',
    description: 'LIMIT',
    base: 'hit count amount',
    limit: 'nonw',
    value: 'none',
    bonusCalc: '',
    notes: 'Hit count is set using field \'numhits\' in spells new.'
  },
  423: {
    spa: 'SE_LimitUseType',
    effectName: 'Limit:  Hit Count Type',
    description: 'LIMIT',
    base: 'hit count type',
    limit: 'nonw',
    value: 'none',
    bonusCalc: '',
    notes: 'hit type is set using field \'numhitstype\' in spells_new.  1:Incoming Hit Attempts, 2:Outgoing Hit Attempts, 3:Incoming Spells,  4:Outgoing Spells, 5: Outgoing Hit Successes, 6:Incoming Hit Successes, 7:Matching Spells, 8:Incoming Hits Or Spells, 9:Reflected Spells,  10:Defensive Proc Casts,  11: Offensive Proc Casts.'
  },
  424: {
    spa: 'SE_GravityEffect',
    effectName: 'Gravitate',
    description: 'Utility Uncategorized',
    base: 'Force',
    limit: 'Distance',
    value: 'unknown',
    bonusCalc: '',
    notes: 'Negative base value will pull target | Positive base value will push target. Max value on live is a level modifier, unclear what it is in our current spell file.'
  },
  425: {
    spa: 'SE_Display',
    effectName: 'Fly',
    description: 'Not Implemented',
    base: '',
    limit: '',
    value: '',
    bonusCalc: '',
    notes: ''
  },
  426: {
    spa: 'SE_IncreaseExtTargetWindow',
    effectName: 'AddExtTargetSlots',
    description: 'Not Implemented',
    base: '',
    limit: '',
    value: '',
    bonusCalc: '',
    notes: ''
  },
  427: {
    spa: 'SE_SkillProc',
    effectName: 'Skill Proc On Attempt',
    description: 'Procs',
    base: 'spellid',
    limit: 'rate modifer',
    value: 'none',
    bonusCalc: '',
    notes: 'Example, can add a proc to taunt which will have a chance to fire each time you use skill.'
  },
  428: {
    spa: 'SE_LimitToSkill',
    effectName: 'Limit To Skill',
    description: 'Proc Limiter',
    base: 'skill type',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: 'This needs to be placed after the proc SPA in a spell to be checked properly.'
  },
  429: {
    spa: 'SE_SkillProcSuccess',
    effectName: 'Skill Proc On Success',
    description: 'Procs',
    base: 'spellid',
    limit: 'rate modifer',
    value: 'none',
    bonusCalc: '',
    notes: 'Example, can add a proc to taunt which will have a chance to fire if you taunt successsfully.'
  },
  430: {
    spa: 'SE_PostEffect',
    effectName: 'PostEffect',
    description: 'Not Implemented',
    base: '',
    limit: '',
    value: '',
    bonusCalc: '',
    notes: ''
  },
  431: {
    spa: 'SE_PostEffectData',
    effectName: 'PostEffectData',
    description: 'Not Implemented',
    base: '',
    limit: '',
    value: '',
    bonusCalc: '',
    notes: ''
  },
  432: {
    spa: 'SE_ExpandMaxActiveTrophyBen',
    effectName: 'Expand Max Active Trophy Benefits',
    description: 'Not Implemented',
    base: '',
    limit: '',
    value: '',
    bonusCalc: '',
    notes: ''
  },
  433: {
    spa: 'SE_CriticalDotDecay',
    effectName: 'Critical DoT Decay',
    description: 'Spell Bonus',
    base: 'percent chance',
    limit: 'decay modifier',
    value: 'none',
    bonusCalc: '',
    notes: 'Live no longer uses this effect, replaced after ROF2 with different effect. Effects were introduced during VoA.'
  },
  434: {
    spa: 'SE_CriticalHealDecay',
    effectName: 'Critical Heal Decay',
    description: 'Heal Modifiers',
    base: 'percent chance',
    limit: 'decay modifier',
    value: 'none',
    bonusCalc: '',
    notes: 'Live no longer uses this effect, replaced after ROF2 with different effect. Effects were introduced during VoA.'
  },
  435: {
    spa: 'SE_CriticalRegenDecay',
    effectName: 'Critic Heal Over Time Decay',
    description: 'Heal Modifiers',
    base: 'percent chance',
    limit: 'decay modifier',
    value: 'none',
    bonusCalc: '',
    notes: 'Live no longer uses this effect, replaced after ROF2 with different effect. Effects were introduced during VoA.'
  },
  436: {
    spa: 'SE_BeneficialCountDownHold',
    effectName: 'Toggle Freeze Buff Timers',
    description: 'Not Implemented',
    base: '',
    limit: '',
    value: '',
    bonusCalc: '',
    notes: ''
  },
  437: {
    spa: 'SE_TeleporttoAnchor',
    effectName: 'Teleport to Anchor',
    description: 'Not Implemented',
    base: '',
    limit: '',
    value: '',
    bonusCalc: '',
    notes: ''
  },
  438: {
    spa: 'SE_TranslocatetoAnchor',
    effectName: 'Translocate to Anchor',
    description: 'Not Implemented',
    base: '',
    limit: '',
    value: '',
    bonusCalc: '',
    notes: ''
  },
  439: {
    spa: 'SE_Assassinate',
    effectName: 'Assassinate',
    description: 'Fatal Procs',
    base: 'percent chance',
    limit: 'damage amount',
    value: 'none',
    bonusCalc: '',
    notes: 'Use with SPA 345 to set max target level and bonus chance.'
  },
  440: {
    spa: 'SE_FinishingBlowLvl',
    effectName: 'Finishing Blow Max Level',
    description: 'Fatal Procs',
    base: 'max target level',
    limit: 'hit point percent to trigger',
    value: 'none',
    bonusCalc: '',
    notes: 'Limit value is calculated as limit / 10. To set FB to trigger below 10 pct HP, set limit to 100. If multiple version of this affect exist it will use highest HP percent.'
  },
  441: {
    spa: 'SE_DistanceRemoval',
    effectName: 'Distance Buff Removal',
    description: 'Utility Uncategorized',
    base: 'distance amount',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: ''
  },
  442: {
    spa: 'SE_TriggerOnReqTarget',
    effectName: 'Trigger Spell on Target Requirement',
    description: 'Cast On Effect',
    base: 'spellid',
    limit: 'Target Restriction ID',
    value: 'none',
    bonusCalc: '',
    notes: 'See enum SpellRestriction in spdat.h for IDs. This trigger spell is usually cast on a target.'
  },
  443: {
    spa: 'SE_TriggerOnReqCaster',
    effectName: 'Trigger Spell on Caster Requirement',
    description: 'Cast On Effect',
    base: 'spellid',
    limit: 'Target Restriction ID',
    value: 'none',
    bonusCalc: '',
    notes: 'See enum SpellRestriction in spdat.h for IDs. This trigger spell is usually cast on self.'
  },
  444: {
    spa: 'SE_ImprovedTaunt',
    effectName: 'Improved Taunt',
    description: 'Aggro',
    base: '999',
    limit: 'other players percent of hate generation',
    value: 'max target level',
    bonusCalc: '',
    notes: 'Limit greater than 100  increased hate generation on other players  (120 = 20 pct increased hate generation) |  Limit less than 100 for decreased hate generation on other players (80 = 20 pct decrease in hate generation)'
  },
  445: {
    spa: 'SE_AddMercSlot',
    effectName: 'Add Merc Slot',
    description: 'Not Implemented',
    base: '',
    limit: '',
    value: '',
    bonusCalc: '',
    notes: ''
  },
  446: {
    spa: 'SE_AStacker',
    effectName: 'A Stacker',
    description: 'Stacking',
    base: 'stacking value',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: 'Buffs containing these effects can block each other from taking hold via the following. Does not matter what slot the effect is in. (B) Blocks any buffs from taking hold with (A) in it. (C) Blocks any buff from taking hold with (B) in it. (D) Blocks any buff from taking hold with (C) in it. When checking same type (ie A vs A), the higher base effect value will determine which takes hold.'
  },
  447: {
    spa: 'SE_BStacker',
    effectName: 'B Stacker',
    description: 'Stacking',
    base: 'stacking value',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: 'Buffs containing these effects can block each other from taking hold via the following. Does not matter what slot the effect is in. (B) Blocks any buffs from taking hold with (A) in it. (C) Blocks any buff from taking hold with (B) in it. (D) Blocks any buff from taking hold with (C) in it. When checking same type (ie A vs A), the higher base effect value will determine which takes hold.'
  },
  448: {
    spa: 'SE_CStacker',
    effectName: 'C Stacker',
    description: 'Stacking',
    base: 'stacking value',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: 'Buffs containing these effects can block each other from taking hold via the following. Does not matter what slot the effect is in. (B) Blocks any buffs from taking hold with (A) in it. (C) Blocks any buff from taking hold with (B) in it. (D) Blocks any buff from taking hold with (C) in it. When checking same type (ie A vs A), the higher base effect value will determine which takes hold.'
  },
  449: {
    spa: 'SE_DStacker',
    effectName: 'D Stacker',
    description: 'Stacking',
    base: 'stacking value',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: 'Buffs containing these effects can block each other from taking hold via the following. Does not matter what slot the effect is in. (B) Blocks any buffs from taking hold with (A) in it. (C) Blocks any buff from taking hold with (B) in it. (D) Blocks any buff from taking hold with (C) in it. When checking same type (ie A vs A), the higher base effect value will determine which takes hold.'
  },
  450: {
    spa: 'SE_MitigateDotDamage',
    effectName: 'Mitigate Damage Over Time Rune',
    description: 'Rune',
    base: 'percent mitigation',
    limit: 'max damage absorbed per hit',
    value: 'rune amount',
    bonusCalc: '',
    notes: 'Special: If this effect is placed on item as worn effect or as an AA, it will provide stackable percent damage over time mitigation for the base value.'
  },
  451: {
    spa: 'SE_MeleeThresholdGuard',
    effectName: 'Melee Threshold Guard',
    description: 'Rune',
    base: 'percent mitigation',
    limit: 'minimum damage to be lowered',
    value: 'rune amount',
    bonusCalc: '',
    notes: ''
  },
  452: {
    spa: 'SE_SpellThresholdGuard',
    effectName: 'Spell Threshold Guard',
    description: 'Rune',
    base: 'percent mitigation',
    limit: 'minimum damage to be lowered',
    value: 'rune amount',
    bonusCalc: '',
    notes: ''
  },
  453: {
    spa: 'SE_TriggerMeleeThreshold',
    effectName: 'Trigger Spell on Melee Threshold',
    description: 'Cast On Effect',
    base: 'spellid',
    limit: 'amount of melee damage',
    value: 'none',
    bonusCalc: '',
    notes: ''
  },
  454: {
    spa: 'SE_TriggerSpellThreshold',
    effectName: 'Trigger Spell on Spell Threshold',
    description: 'Cast On Effect',
    base: 'spellid',
    limit: 'amount of spell damage',
    value: 'none',
    bonusCalc: '',
    notes: ''
  },
  455: {
    spa: 'SE_AddHatePct',
    effectName: 'Add Hate Percent',
    description: 'Aggro',
    base: 'percent modifier',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: 'Positive value increased hate. | Negative value decreases hate.'
  },
  456: {
    spa: 'SE_AddHateOverTimePct',
    effectName: 'Add Hate Over Time Percent',
    description: 'Aggro',
    base: 'percent modifier',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: 'Positive value increases hate. | Negative value decreases hate.'
  },
  457: {
    spa: 'SE_ResourceTap',
    effectName: 'Resource Tap',
    description: 'Converts',
    base: 'percent coverted',
    limit: '0=HP, 1=Mana,2=Endurance',
    value: 'max amount resource returned',
    bonusCalc: '',
    notes: 'Conversion percent calculated as value / 10, example to convert 85 percent of damage to mana set base value to 850 and limit to 1.'
  },
  458: {
    spa: 'SE_FactionModPct',
    effectName: 'Faction Pct',
    description: 'Faction',
    base: 'percent modifier',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: ''
  },
  459: {
    spa: 'SE_DamageModifier2',
    effectName: 'Skill Damage Modifier (v459)',
    description: 'Offensive Bonus',
    base: 'percent modifer',
    limit: 'skill type (-1 = all skill types)',
    value: 'none',
    bonusCalc: '',
    notes: ''
  },
  460: {
    spa: 'SE_Ff_Override_NotFocusable',
    effectName: 'Limit:  Include Non-Focusable',
    description: 'LIMIT',
    base: '1',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: ''
  },
  461: {
    spa: 'SE_ImprovedDamage2',
    effectName: 'Focus: Spell Damage (v461)',
    description: 'FOCUS',
    base: 'min percent',
    limit: 'none',
    value: 'max percent',
    bonusCalc: '',
    notes: 'Use random effectiveness if base and max value are defined, where base is always lower end and max the higher end of the random range. If random value not wanted, then only set a base value.'
  },
  462: {
    spa: 'SE_FcDamageAmt2',
    effectName: 'Focus: Spell Damage Amount (v462)',
    description: 'FOCUS',
    base: 'amount',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: ''
  },
  463: {
    spa: 'SE_Shield_Target',
    effectName: 'Shield Target',
    description: 'Not Implemented',
    base: '',
    limit: '',
    value: '',
    bonusCalc: '',
    notes: ''
  },
  464: {
    spa: 'SE_PC_Pet_Rampage',
    effectName: 'PC Pet Rampage',
    description: 'Pet Bonuses',
    base: 'percent chance',
    limit: 'damage modifier',
    value: 'none',
    bonusCalc: '',
    notes: 'Limit greater than 100 for increased damage (120 = 20 pct increased damage) |  Limit less than 100 for decreased damage (80 = 20 pct decreased damage)'
  },
  465: {
    spa: 'SE_PC_Pet_AE_Rampage',
    effectName: 'PC Pet AE Rampage',
    description: 'Pet Bonuses',
    base: 'percent chance',
    limit: 'damage modifier',
    value: 'none',
    bonusCalc: '',
    notes: 'Limit greater than 100 for increased damage (120 = 20 pct increased damage) |  Limit less than 100 for decreased damage (80 = 20 pct decreased damage)'
  },
  466: {
    spa: 'SE_PC_Pet_Flurry_Chance',
    effectName: 'PC Pet Flurry Chance',
    description: 'Pet Bonuses',
    base: 'percent chance',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: ''
  },
  467: {
    spa: 'SE_DS_Mitigation_Amount',
    effectName: 'Damage Shield Mitigation Amount',
    description: 'Damage Shield',
    base: 'amount',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: ''
  },
  468: {
    spa: 'SE_DS_Mitigation_Percentage',
    effectName: 'Damage Shield Mitigation Percentage',
    description: 'Damage Shield',
    base: 'percent modifier',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: ''
  },
  469: {
    spa: 'SE_Chance_Best_in_Spell_Grp',
    effectName: 'Spell Trigger: Best in Spell Group: Only One Spell Cast',
    description: 'Cast On Effect',
    base: 'percent chance',
    limit: 'spellgroup id',
    value: '',
    bonusCalc: '',
    notes: 'When multiple of this effect exist on the same spell, only one spell will be selected from the list to be cast. For best results, the total percent chance should equal 100%. Example, Slot 1: Cast Ice Nuke spellgroup 20%, Slot2: Cast Fire Nuke spellgroup 50%, Slot3 Cast Magic Nuke spellgroup 30%. When the spell is cast, only one of these spells be triggered on the target.'
  },
  470: {
    spa: 'SE_Trigger_Best_in_Spell_Grp',
    effectName: 'Spell Trigger: Best in Spell Group: Apply Each',
    description: 'Cast On Effect',
    base: 'percent chance',
    limit: 'spellgroup id',
    value: '',
    bonusCalc: '',
    notes: ''
  },
  471: {
    spa: 'SE_Double_Melee_Round',
    effectName: 'Double Melee Round (PC Only)',
    description: 'Offensive Bonus',
    base: 'percent chance',
    limit: 'percent damage modifier',
    value: 'none',
    bonusCalc: '',
    notes: ''
  },
  472: {
    spa: 'SE_Buy_AA_Rank',
    effectName: 'Toggle Passive AA Rank',
    description: 'Utility Uncategorized',
    base: '1',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: 'Certain AA, like Weapon Stance line use a special toggle Hotkey to enable or disable the AA\'s passive abilities.  This is occurs by doing the following. Each \'rank\' of Weapon Stance is actually 2 actual ranks.   First rank is always the Disabled version which cost X amount of AA, this is the rank that SPA 472 goes in. Second rank is the Enabled version which cost 0 AA.  When you buy the first rank, you make a hotkey that on live say \'Weapon Stance Disabled\', if you clik that it then BUYS the  next rank of AA (cost 0) which switches the hotkey to \'Enabled Weapon Stance\' and you are given the passive buff effects.  If you click the Enabled hotkey, it causes you to lose an AA rank and once again be disabled. Thus, you are switching between  two AA ranks. Thefore when creating an AA using this ability, you need generate both ranks. Follow the same pattern for additional ranks. See aa.cpp for further details.'
  },
  473: {
    spa: 'SE_Double_Backstab_Front',
    effectName: 'Double Backstab From Front',
    description: 'Offensive Bonus',
    base: 'percent chance',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: ''
  },
  474: {
    spa: 'SE_Pet_Crit_Melee_Damage_Pct_Owner',
    effectName: 'Pet Crit Melee Damage',
    description: 'Pet Bonuses From Owner',
    base: 'percent modifier',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: 'This effect goes on the pet owner and then the benefit is applied to the pet.'
  },
  475: {
    spa: 'SE_Trigger_Spell_Non_Item',
    effectName: 'Trigger Spell: Not Cast From Items: Apply Each',
    description: 'Cast On Effect',
    base: 'chance pecent',
    limit: 'spellid',
    value: '',
    bonusCalc: '',
    notes: ''
  },
  476: {
    spa: 'SE_Weapon_Stance',
    effectName: 'Weapon Stance',
    description: 'Beneficial Uncategorized',
    base: 'spellid',
    limit: '0=2H, 1=Shield, 2=DW',
    value: 'none',
    bonusCalc: '',
    notes: 'On live this is an AA, on emu can use as item or spell buff effect. Live AA uses a toggle system to turn on and off weapons stance. which requires use of SPA 472 SE_Buy_AA_Rank within in the AA. Details can be found in aa.cpp'
  },
  477: {
    spa: 'SE_Hatelist_To_Top_Index',
    effectName: 'Move to Top of Rampage List',
    description: 'Aggro',
    base: 'chacne percent',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: ''
  },
  478: {
    spa: 'SE_Hatelist_To_Tail_Index',
    effectName: 'Move to Bottom of Rampage List',
    description: 'Aggro',
    base: 'percent chance',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: ''
  },
  479: {
    spa: 'SE_Ff_Value_Min',
    effectName: 'Limit:  Base Value Min',
    description: 'LIMIT',
    base: 'effect_base_value',
    limit: 'spell effect id',
    value: 'none',
    bonusCalc: '',
    notes: 'Example, only allow focus if spell has Effect ID 0 which is a Heal and the heals effect value is less than 5000.'
  },
  480: {
    spa: 'SE_Ff_Value_Max',
    effectName: 'Limit:  Base Value Max',
    description: 'LIMIT',
    base: 'effect_base_value',
    limit: 'spell effect id',
    value: 'none',
    bonusCalc: '',
    notes: 'Example, only allow focus if spell has Effect ID 0 which is a Heal and the heals effect value is greater than 5000.'
  },
  481: {
    spa: 'SE_Fc_Cast_Spell_On_Land',
    effectName: 'Focus: Trigger Spell on Spell Landing',
    description: 'FOCUS',
    base: 'spellid',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: 'Example, everytime you are hit with a fire nuke, cast a heal on yourself.'
  },
  482: {
    spa: 'SE_Skill_Base_Damage_Mod',
    effectName: 'Base Hit Damage',
    description: 'Offenisve Bonus',
    base: 'percent modifier',
    limit: 'skill type (-1 = ALL skill types)',
    value: 'none',
    bonusCalc: '',
    notes: ''
  },
  483: {
    spa: 'SE_Fc_Spell_Damage_Pct_IncomingPC',
    effectName: 'Focus: Incoming Spell Damage (v483)',
    description: 'FOCUS',
    base: 'min percent modifier',
    limit: 'none',
    value: 'max percent modifier',
    bonusCalc: '',
    notes: 'Use random effectiveness if base and max value are defined, where base is always lower end and max the higher end of the random range. If random value not wanted, then only set a base value.'
  },
  484: {
    spa: 'SE_Fc_Spell_Damage_Amt_IncomingPC',
    effectName: 'Focus: Incoming Spell Damage Amount (v484)',
    description: 'FOCUS',
    base: 'amount',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: ''
  },
  485: {
    spa: 'SE_Ff_CasterClass',
    effectName: 'Limit: Caster Class',
    description: 'LIMIT',
    base: 'class bitmask',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: 'Set multiple classes same as would for items. This is only used with focus effects that check against incoming spells.'
  },
  486: {
    spa: 'SE_Ff_Same_Caster',
    effectName: 'Limit: Caster',
    description: 'LIMIT',
    base: '0=Must be different caster 1=Must be same caster',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: 'This is only used with focus effects that check against incoming spells.'
  },
  487: {
    spa: 'SE_Extend_Tradeskill_Cap',
    effectName: 'Extend Tradeskill Cap',
    description: 'Not Implemented',
    base: '',
    limit: '',
    value: '',
    bonusCalc: '',
    notes: ''
  },
  488: {
    spa: 'SE_Defender_Melee_Force_Pct_PC',
    effectName: 'Push Taken',
    description: 'Not Implemented',
    base: '',
    limit: '',
    value: '',
    bonusCalc: '',
    notes: ''
  },
  489: {
    spa: 'SE_Worn_Endurance_Regen_Cap',
    effectName: 'Worn Endurance Regen Cap',
    description: 'Cap',
    base: 'amount',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: ''
  },
  490: {
    spa: 'SE_Ff_ReuseTimeMin',
    effectName: 'Limit:  Reuse Time Min',
    description: 'LIMIT',
    base: 'recast time',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: ''
  },
  491: {
    spa: 'SE_Ff_ReuseTimeMax',
    effectName: 'Limit:  Reuse Time Max',
    description: 'LIMIT',
    base: 'recast time',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: ''
  },
  492: {
    spa: 'SE_Ff_Endurance_Min',
    effectName: 'Limit:  Endurance Min',
    description: 'LIMIT',
    base: 'endurance cost',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: ''
  },
  493: {
    spa: 'SE_Ff_Endurance_Max',
    effectName: 'Limit:  Endurance Max',
    description: 'LIMIT',
    base: 'endurance cost',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: ''
  },
  494: {
    spa: 'SE_Pet_Add_Atk',
    effectName: 'Pet Add ATK',
    description: 'Pet Bonuses From Owner',
    base: 'amount',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: 'This effect goes on the pet owner and then the benefit is applied to the pet.'
  },
  495: {
    spa: 'SE_Ff_DurationMax',
    effectName: 'Limit:  Duration Max',
    description: 'LIMIT',
    base: 'tics',
    limit: '',
    value: '',
    bonusCalc: '',
    notes: ''
  },
  496: {
    spa: 'SE_Critical_Melee_Damage_Mod_Max',
    effectName: 'Critical Melee Damage: No Stack',
    description: 'Offensive Bonus',
    base: 'percent modifier',
    limit: 'skill type (-1 = all skill types)',
    value: 'none',
    bonusCalc: '',
    notes: ''
  },
  497: {
    spa: 'SE_Ff_FocusCastProcNoBypass',
    effectName: 'Limit:  Proc No Bypass',
    description: 'Not Implemented',
    base: '',
    limit: '',
    value: '',
    bonusCalc: '',
    notes: ''
  },
  498: {
    spa: 'SE_AddExtraAttackPct_1h_Primary',
    effectName: 'Add Extra Attack: 1H Primary',
    description: 'Offensive Bonus',
    base: 'percent chance',
    limit: 'number of attacks',
    value: 'none',
    bonusCalc: '',
    notes: ''
  },
  499: {
    spa: 'SE_AddExtraAttackPct_1h_Secondary',
    effectName: 'Add Extra Attack: 1H Secondary',
    description: 'Offensive Bonus',
    base: 'percent chance',
    limit: 'number of attacks',
    value: 'none',
    bonusCalc: '',
    notes: ''
  },
  500: {
    spa: 'SE_Fc_CastTimeMod2',
    effectName: 'Focus: Spell Haste (v500, no cap) ',
    description: 'FOCUS',
    base: 'percent modifier',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: 'Can reduce cast time to 0.'
  },
  501: {
    spa: 'SE_Fc_CastTimeAmt',
    effectName: 'Focus: Spell Cast Time',
    description: 'FOCUS',
    base: 'time ms',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: 'Can reduce cast time to 0.'
  },
  502: {
    spa: 'SE_Fearstun',
    effectName: 'Fearstun',
    description: 'Stun',
    base: 'duration ms',
    limit: 'PC duration ms',
    value: 'target max level',
    bonusCalc: '',
    notes: 'Max value can be calculated in two ways, to set a max level that target can be affected the formula is value - 1000, example if max level is 80, then set as 1080. If you want to set max level to be relative to caster, example only affects entities that are 3 more less level higher than caster, then set max value to 3.'
  },
  503: {
    spa: 'SE_Melee_Damage_Position_Mod',
    effectName: 'Rear Arc Melee Damage Mod',
    description: 'Offensive Bonuses',
    base: 'percent modifier',
    limit: '0=back, 1=front',
    value: '',
    bonusCalc: '',
    notes: ''
  },
  504: {
    spa: 'SE_Melee_Damage_Position_Amt',
    effectName: 'Rear Arc Melee Damage Amt',
    description: 'Offensive Bonuses',
    base: 'amount',
    limit: '0=back, 1=front',
    value: '',
    bonusCalc: '',
    notes: ''
  },
  505: {
    spa: 'SE_Damage_Taken_Position_Mod',
    effectName: 'Rear Arc Damage Taken Mod',
    description: 'Offensive Bonuses',
    base: 'percent modifier',
    limit: '0=back, 1=front',
    value: '',
    bonusCalc: '',
    notes: ''
  },
  506: {
    spa: 'SE_Damage_Taken_Position_Amt',
    effectName: 'Rear Arc Damage Taken Amt',
    description: 'Offensive Bonuses',
    base: 'amount',
    limit: '0=back, 1=front',
    value: '',
    bonusCalc: '',
    notes: ''
  },
  507: {
    spa: 'SE_Fc_Amplify_Mod',
    effectName: 'Focus: Spell Heal Damage and DoT',
    description: 'FOCUS',
    base: 'percent modifier',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: ''
  },
  508: {
    spa: 'SE_Fc_Amplify_Amt',
    effectName: 'Focus: Spell Heal Damage and DoT Amount',
    description: 'FOCUS',
    base: 'amount',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: ''
  },
  509: {
    spa: 'SE_Health_Transfer',
    effectName: 'Health Transfer',
    description: 'Hit Points',
    base: 'casters percent HP change',
    limit: 'percent of casters HP to damage',
    value: 'none',
    bonusCalc: '',
    notes: 'Used in newer versions of  Lifeburn and Act of Valor. Base is calculated as value / 100, where a reducation of 75 percent of casters HP would be value 750. Limit is calculated as value / 10, if value is set to 1000 then damage will be 100 percent of casters HP. Negative base decreases casters HP, negative limit decreases targets HP, positive limit heals target.'
  },
  510: {
    spa: 'SE_Fc_ResistIncoming',
    effectName: 'Focus: Incoming Resist Modifier',
    description: 'FOCUS',
    base: 'amount',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: 'Positive value will lower resist modifier. Every 2 pts of value equals a resist rate decrease of 1 percent. Example, base -750 and limit -10000, will consumes 75% of your current health and deals 1000% of that health as direct damage.'
  },
  511: {
    spa: 'SE_Ff_FocusTimerMin',
    effectName: 'Limit: Focus Reuse Timer',
    description: 'LIMIT',
    base: '1',
    limit: 'time ms',
    value: 'none',
    bonusCalc: '',
    notes: 'Example, set limit to 1500, then this focus can only trigger once every 1.5 seconds.'
  },
  512: {
    spa: 'SE_Proc_Timer_Modifier',
    effectName: 'Proc Reuse Timer',
    description: 'Proc Limiter',
    base: '1',
    limit: 'time ms',
    value: 'none',
    bonusCalc: '',
    notes: 'Example, set limit to 1500, then this proc can only trigger once every 1.5 seconds.'
  },
  513: {
    spa: 'SE_Mana_Max_Percent',
    effectName: 'Mana Max Percent',
    description: 'Not Implemented',
    base: '',
    limit: '',
    value: '',
    bonusCalc: '',
    notes: ''
  },
  514: {
    spa: 'SE_Endurance_Max_Percent',
    effectName: 'Endurance Max Percent',
    description: 'Not Implemented',
    base: '',
    limit: '',
    value: '',
    bonusCalc: '',
    notes: ''
  },
  515: {
    spa: 'SE_AC_Avoidance_Max_Percent',
    effectName: 'AC Avoidance Max Percent',
    description: 'Offensive Bonuses',
    base: 'percent modifier',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: 'Calculated as base value / 100 is actual percent. Example for 7.14 percent modifier, set to 714.'
  },
  516: {
    spa: 'SE_AC_Mitigation_Max_Percent',
    effectName: 'AC Mitigation Max Percent',
    description: 'Offensive Bonuses',
    base: 'percent modifier',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: 'Calculated as base value / 100 is actual percent. Example for 7.14 percent modifier, set to 714.'
  },
  517: {
    spa: 'SE_Attack_Offense_Max_Percent',
    effectName: 'Attack Offense Max Percent',
    description: 'Not Implemented',
    base: 'percent modifier',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: 'Calculated as base value / 100 is actual percent. Example for 7.14 percent modifier, set to 714.'
  },
  518: {
    spa: 'SE_Attack_Accuracy_Max_Percent',
    effectName: 'Attack Accuracy Max Percent',
    description: 'Offensive Bonuses',
    base: 'percent modifier',
    limit: 'none',
    value: 'none',
    bonusCalc: '',
    notes: 'Calculated as base value / 100 is actual percent. Example for 7.14 percent modifier, set to 714.'
  },
  519: {
    spa: 'SE_Luck_Amount',
    effectName: 'Luck Amount',
    description: 'Not Implemented',
    base: '',
    limit: '',
    value: '',
    bonusCalc: '',
    notes: ''
  },
  520: {
    spa: 'SE_Luck_Percent',
    effectName: 'Luck Percent',
    description: 'Not Implemented',
    base: '',
    limit: '',
    value: '',
    bonusCalc: '',
    notes: ''
  },
  521: {
    spa: 'SE_Endurance_Absorb_Pct_Damage',
    effectName: 'Absorb Damage with Endurance',
    description: 'Defensive Bonus',
    base: 'mitigation percentage',
    limit: 'ratio',
    value: 'none',
    bonusCalc: '',
    notes: 'Base calculated as value / 100,  example if base 2000 then mitigation is 20 percent. Limit is calculated as value / 10000, example if limit 500 then ratio will be 0.05 Endurance reduced per 1 Hit Point of damage.'
  },
  522: {
    spa: 'SE_Instant_Mana_Pct',
    effectName: 'Instant Mana Percent',
    description: 'Mana',
    base: 'percent of max mana',
    limit: 'max amount',
    value: 'none',
    bonusCalc: '',
    notes: 'Negative base value increases mana | Positive base value increases mana. Base is calculated as value / 100. Example if base 150 and max mana 1000, effect will extracts 1.5% of your maximum mana, or 1000, whichever is lower.'
  },
  523: {
    spa: 'SE_Instant_Endurance_Pct',
    effectName: 'Instant Endurance Percent',
    description: 'Endurance',
    base: 'percent of max endurance',
    limit: 'max amount',
    value: 'none',
    bonusCalc: '',
    notes: 'Negative base value increases endurance | Positive base value increases endurance. Base is calculated as value / 100. Example if base 150 and max endurance 1000, effect will extracts 1.5% of your maximum endurance, or 1000, whichever is lower.'
  },
  524: {
    spa: 'SE_Duration_HP_Pct',
    effectName: 'Duration HP Percent',
    description: 'Hit Points',
    base: 'percent of max hit points',
    limit: 'max amount',
    value: 'none',
    bonusCalc: '',
    notes: 'Negative base value increases hit points | Positive base value increases hit points. Base is calculated as value / 100. Example if base 150 and max hit points 1000, effect will extracts 1.5% of your maximum hit points, or 1000, whichever is lower per tic.'
  },
  525: {
    spa: 'SE_Duration_Mana_Pct',
    effectName: 'Duration Mana Percent',
    description: 'Mana',
    base: 'percent of max mana',
    limit: 'max amount',
    value: 'none',
    bonusCalc: '',
    notes: 'Negative base value increases mana | Positive base value increases mana. Base is calculated as value / 100. Example if base 150 and max mana 1000, effect will extracts 1.5% of your maximum mana, or 1000, whichever is lower per tic.'
  },
  526: {
    spa: 'SE_Duration_Endurance_Pct',
    effectName: 'Duration Endurance Percent',
    description: 'Endurance',
    base: 'percent of max endurance',
    limit: 'max amount',
    value: 'none',
    bonusCalc: '',
    notes: 'Negative base value increases endurance | Positive base value increases endurance. Base is calculated as value / 100. Example if base 150 and max endurance 1000, effect will extracts 1.5% of your maximum endurance, or 1000, whichever is lowe per tic.'
  },
}
