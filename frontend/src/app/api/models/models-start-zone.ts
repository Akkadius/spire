import { ModelsZone } from './models-zone';
export interface ModelsStartZone {
    bind_id?: number;
    bind_x?: number;
    bind_y?: number;
    bind_z?: number;
    content_flags?: string;
    content_flags_disabled?: string;
    heading?: number;
    max_expansion?: number;
    min_expansion?: number;
    player_choice?: number;
    player_class?: number;
    player_deity?: number;
    player_race?: number;
    select_rank?: number;
    start_zone?: number;
    x?: number;
    y?: number;
    z?: number;
    zone?: ModelsZone;
    zone_id?: number;
}
