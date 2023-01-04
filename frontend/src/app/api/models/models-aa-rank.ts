import { ModelsSpellsNew } from './models-spells-new';
export interface ModelsAaRank {
    cost?: number;
    desc_sid?: number;
    expansion?: number;
    id?: number;
    level_req?: number;
    lower_hotkey_sid?: number;
    next_id?: number;
    prev_id?: number;
    recast_time?: number;
    spell?: number;
    spell_type?: number;
    spells_new?: ModelsSpellsNew;
    title_sid?: number;
    upper_hotkey_sid?: number;
}
