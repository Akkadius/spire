import { ModelsItem } from './models-item';
import { ModelsLootdrop } from './models-lootdrop';
export interface ModelsLootdropEntry {
    chance?: number;
    content_flags?: string;
    content_flags_disabled?: string;
    disabled_chance?: number;
    equip_item?: number;
    item?: ModelsItem;
    item_charges?: number;
    item_id?: number;
    lootdrop?: ModelsLootdrop;
    lootdrop_id?: number;
    max_expansion?: number;
    min_expansion?: number;
    multiplier?: number;
    npc_max_level?: number;
    npc_min_level?: number;
    trivial_max_level?: number;
    trivial_min_level?: number;
}
