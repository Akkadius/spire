import { ModelsItem } from './models-item';
import { ModelsZone } from './models-zone';
export interface ModelsForage {
    chance?: number;
    content_flags?: string;
    content_flags_disabled?: string;
    id?: number;
    item?: ModelsItem;
    itemid?: number;
    level?: number;
    max_expansion?: number;
    min_expansion?: number;
    zone?: ModelsZone;
    zoneid?: number;
}
