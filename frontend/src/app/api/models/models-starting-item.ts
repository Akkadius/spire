import { ModelsItem } from './models-item';
import { ModelsZone } from './models-zone';
export interface ModelsStartingItem {
    _class?: number;
    content_flags?: string;
    content_flags_disabled?: string;
    deityid?: number;
    gm?: number;
    id?: number;
    item?: ModelsItem;
    item_charges?: number;
    itemid?: number;
    max_expansion?: number;
    min_expansion?: number;
    race?: number;
    slot?: number;
    zone?: ModelsZone;
    zoneid?: number;
}
