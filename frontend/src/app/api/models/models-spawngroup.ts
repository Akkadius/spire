import { ModelsSpawn2 } from './models-spawn2';
export interface ModelsSpawngroup {
    delay?: number;
    despawn?: number;
    despawn_timer?: number;
    dist?: number;
    id?: number;
    max_x?: number;
    max_y?: number;
    min_x?: number;
    min_y?: number;
    mindelay?: number;
    name?: string;
    spawn_2?: ModelsSpawn2;
    spawn_limit?: number;
    wp_spawns?: number;
}
