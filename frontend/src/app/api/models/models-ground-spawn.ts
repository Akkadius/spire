import { ModelsZone } from './models-zone';
export interface ModelsGroundSpawn {
    comment?: string;
    content_flags?: string;
    content_flags_disabled?: string;
    heading?: number;
    id?: number;
    item?: number;
    max_allowed?: number;
    max_expansion?: number;
    max_x?: number;
    max_y?: number;
    max_z?: number;
    min_expansion?: number;
    min_x?: number;
    min_y?: number;
    name?: string;
    respawn_timer?: number;
    version?: number;
    zone?: ModelsZone;
    zoneid?: number;
}
