import { ModelsRaidLeader } from './models-raid-leader';
import { ModelsRaidMember } from './models-raid-member';
export interface ModelsRaidDetail {
    locked?: number;
    loottype?: number;
    marked_npc_1_entity_id?: number;
    marked_npc_1_instance_id?: number;
    marked_npc_1_zone_id?: number;
    marked_npc_2_entity_id?: number;
    marked_npc_2_instance_id?: number;
    marked_npc_2_zone_id?: number;
    marked_npc_3_entity_id?: number;
    marked_npc_3_instance_id?: number;
    marked_npc_3_zone_id?: number;
    motd?: string;
    raid_leaders?: Array<ModelsRaidLeader>;
    raid_members?: Array<ModelsRaidMember>;
    raidid?: number;
}
