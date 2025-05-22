import { ModelsNpcFactionEntry } from './models-npc-faction-entry';
export interface ModelsNpcFaction {
    id?: number;
    ignore_primary_assist?: number;
    name?: string;
    npc_faction_entries?: Array<ModelsNpcFactionEntry>;
    primaryfaction?: number;
}
