import { ModelsNpcSpellsEntry } from './models-npc-spells-entry';
export interface ModelsNpcSpell {
    attack_proc?: number;
    defensive_proc?: number;
    dproc_chance?: number;
    engaged_b_other_chance?: number;
    engaged_b_self_chance?: number;
    engaged_d_chance?: number;
    engaged_no_sp_recast_max?: number;
    engaged_no_sp_recast_min?: number;
    fail_recast?: number;
    id?: number;
    idle_b_chance?: number;
    idle_no_sp_recast_max?: number;
    idle_no_sp_recast_min?: number;
    name?: string;
    npc_spell?: ModelsNpcSpell;
    npc_spells_entries?: Array<ModelsNpcSpellsEntry>;
    parent_list?: number;
    proc_chance?: number;
    pursue_d_chance?: number;
    pursue_no_sp_recast_max?: number;
    pursue_no_sp_recast_min?: number;
    range_proc?: number;
    rproc_chance?: number;
}
