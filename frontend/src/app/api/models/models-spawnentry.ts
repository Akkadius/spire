import { ModelsNpcType } from './models-npc-type';
import { ModelsSpawngroup } from './models-spawngroup';
export interface ModelsSpawnentry {
    chance?: number;
    condition_value_filter?: number;
    content_flags?: string;
    content_flags_disabled?: string;
    max_expansion?: number;
    min_expansion?: number;
    npc_id?: number;
    npc_type?: ModelsNpcType;
    spawngroup?: ModelsSpawngroup;
    spawngroup_id?: number;
}
