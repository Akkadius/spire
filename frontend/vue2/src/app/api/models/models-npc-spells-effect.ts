import { ModelsNpcSpellsEffectsEntry } from './models-npc-spells-effects-entry';
export interface ModelsNpcSpellsEffect {
    id?: number;
    name?: string;
    npc_spells_effects_entries?: Array<ModelsNpcSpellsEffectsEntry>;
    parent_list?: number;
}
