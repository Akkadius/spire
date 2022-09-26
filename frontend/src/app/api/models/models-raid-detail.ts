import { ModelsRaidLeader } from './models-raid-leader';
import { ModelsRaidMember } from './models-raid-member';
export interface ModelsRaidDetail {
    locked?: number;
    loottype?: number;
    motd?: string;
    raid_leaders?: Array<ModelsRaidLeader>;
    raid_members?: Array<ModelsRaidMember>;
    raidid?: number;
}
