import { ModelsLootdropEntry } from './models-lootdrop-entry';
import { ModelsLoottableEntry } from './models-loottable-entry';
export interface ModelsLootdrop {
    content_flags?: string;
    content_flags_disabled?: string;
    id?: number;
    lootdrop_entries?: Array<ModelsLootdropEntry>;
    loottable_entries?: Array<ModelsLoottableEntry>;
    max_expansion?: number;
    min_expansion?: number;
    name?: string;
}
