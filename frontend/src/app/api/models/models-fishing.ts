import { ModelsItem } from './models-item';
import { ModelsNpcType } from './models-npc-type';
import { ModelsZone } from './models-zone';
export interface ModelsFishing {
    chance?: number;
    content_flags?: string;
    content_flags_disabled?: string;
    id?: number;
    item?: ModelsItem;
    itemid?: number;
    max_expansion?: number;
    min_expansion?: number;
    npc_chance?: number;
    npc_id?: number;
    npc_type?: ModelsNpcType;
    skill_level?: number;
    zone?: ModelsZone;
    zoneid?: number;
}
