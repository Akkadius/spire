import { ModelsZone } from './models-zone';
export interface ModelsStartingItem {
    augment_five?: number;
    augment_four?: number;
    augment_one?: number;
    augment_six?: number;
    augment_three?: number;
    augment_two?: number;
    class_list?: string;
    content_flags?: string;
    content_flags_disabled?: string;
    deity_list?: string;
    id?: number;
    inventory_slot?: number;
    item_charges?: number;
    item_id?: number;
    max_expansion?: number;
    min_expansion?: number;
    race_list?: string;
    status?: number;
    zone?: ModelsZone;
    zone_id_list?: string;
}
