import { ModelsRaidLeader } from './models-raid-leader';
import { ModelsRaidMember } from './models-raid-member';
export interface ModelsRaidDetail {
    locked?: number;
    loottype?: number;
    marked_npc_1?: number;
    marked_npc_2?: number;
    marked_npc_3?: number;
    motd?: string;
    raid_leaders?: Array<ModelsRaidLeader>;
    raid_members?: Array<ModelsRaidMember>;
    raidid?: number;
}
