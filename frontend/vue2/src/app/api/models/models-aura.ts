import { ModelsSpellsNew } from './models-spells-new';
export interface ModelsAura {
    aura_type?: number;
    cast_time?: number;
    distance?: number;
    duration?: number;
    icon?: number;
    movement?: number;
    name?: string;
    npc_type?: number;
    spawn_type?: number;
    spell_id?: number;
    spells_new?: ModelsSpellsNew;
    type?: number;
}
