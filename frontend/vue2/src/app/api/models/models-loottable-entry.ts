import { ModelsLootdrop } from './models-lootdrop';
import { ModelsLoottable } from './models-loottable';
export interface ModelsLoottableEntry {
    droplimit?: number;
    lootdrop?: ModelsLootdrop;
    lootdrop_id?: number;
    loottable?: ModelsLoottable;
    loottable_id?: number;
    mindrop?: number;
    multiplier?: number;
    probability?: number;
}
