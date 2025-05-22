import { ModelsLoottable } from './models-loottable';
export interface ModelsGlobalLoot {
    bodytype?: string;
    _class?: string;
    content_flags?: string;
    content_flags_disabled?: string;
    description?: string;
    enabled?: number;
    hot_zone?: number;
    id?: number;
    loottable?: ModelsLoottable;
    loottable_id?: number;
    max_expansion?: number;
    max_level?: number;
    min_expansion?: number;
    min_level?: number;
    race?: string;
    raid?: number;
    rare?: number;
    zone?: string;
}
