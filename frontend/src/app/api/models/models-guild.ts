import { ModelsGuildBank } from './models-guild-bank';
import { ModelsGuildMember } from './models-guild-member';
import { ModelsGuildRank } from './models-guild-rank';
export interface ModelsGuild {
    channel?: string;
    favor?: number;
    guild_banks?: Array<ModelsGuildBank>;
    guild_members?: Array<ModelsGuildMember>;
    guild_ranks?: Array<ModelsGuildRank>;
    id?: number;
    leader?: number;
    minstatus?: number;
    motd?: string;
    motd_setter?: string;
    name?: string;
    tribute?: number;
    url?: string;
}
