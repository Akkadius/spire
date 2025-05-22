import { ModelsItem } from './models-item';
import { ModelsNpcType } from './models-npc-type';
export interface ModelsMerchantlist {
    alt_currency_cost?: number;
    bucket_comparison?: number;
    bucket_name?: string;
    bucket_value?: string;
    classes_required?: number;
    content_flags?: string;
    content_flags_disabled?: string;
    faction_required?: number;
    item?: number;
    items?: Array<ModelsItem>;
    level_required?: number;
    max_expansion?: number;
    max_status?: number;
    merchantid?: number;
    min_expansion?: number;
    min_status?: number;
    npc_types?: Array<ModelsNpcType>;
    probability?: number;
    slot?: number;
}
