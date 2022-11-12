import { ModelsNpcSpell } from './models-npc-spell';
import { ModelsSpellsNew } from './models-spells-new';
export interface ModelsBotSpellsEntry {
    bucket_comparison?: number;
    bucket_name?: string;
    bucket_value?: string;
    id?: number;
    manacost?: number;
    max_hp?: number;
    maxlevel?: number;
    min_hp?: number;
    minlevel?: number;
    npc_spell?: ModelsNpcSpell;
    npc_spells_id?: number;
    priority?: number;
    recast_delay?: number;
    resist_adjust?: number;
    spellid?: number;
    spells_new?: ModelsSpellsNew;
    type?: number;
}
