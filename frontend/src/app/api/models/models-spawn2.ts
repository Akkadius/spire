import { ModelsSpawnentry } from './models-spawnentry';
import { ModelsSpawngroup } from './models-spawngroup';
export interface ModelsSpawn2 {
    _condition?: number;
    animation?: number;
    cond_value?: number;
    content_flags?: string;
    content_flags_disabled?: string;
    heading?: number;
    id?: number;
    max_expansion?: number;
    min_expansion?: number;
    path_when_zone_idle?: number;
    pathgrid?: number;
    respawntime?: number;
    spawnentries?: Array<ModelsSpawnentry>;
    spawngroup?: ModelsSpawngroup;
    spawngroup_id?: number;
    variance?: number;
    version?: number;
    x?: number;
    y?: number;
    z?: number;
    zone?: string;
}
