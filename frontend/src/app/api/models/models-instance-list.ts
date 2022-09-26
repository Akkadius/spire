import { ModelsInstanceListPlayer } from './models-instance-list-player';
import { ModelsZone } from './models-zone';
export interface ModelsInstanceList {
    duration?: number;
    id?: number;
    instance_list_players?: Array<ModelsInstanceListPlayer>;
    is_global?: number;
    never_expires?: number;
    start_time?: number;
    version?: number;
    zone?: number;
    zones?: Array<ModelsZone>;
}
