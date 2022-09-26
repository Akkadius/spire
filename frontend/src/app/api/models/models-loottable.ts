import { ModelsLoottableEntry } from './models-loottable-entry';
import { ModelsNpcType } from './models-npc-type';
export interface ModelsLoottable {
    avgcoin?: number;
    content_flags?: string;
    content_flags_disabled?: string;
    done?: number;
    id?: number;
    loottable_entries?: Array<ModelsLoottableEntry>;
    max_expansion?: number;
    maxcash?: number;
    min_expansion?: number;
    mincash?: number;
    name?: string;
    npc_types?: Array<ModelsNpcType>;
}
