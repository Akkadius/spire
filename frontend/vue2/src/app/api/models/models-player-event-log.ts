import { ModelsAccount } from './models-account';
import { ModelsCharacterDatum } from './models-character-datum';
import { ModelsZone } from './models-zone';
export interface ModelsPlayerEventLog {
    account?: ModelsAccount;
    account_id?: number;
    character_datum?: ModelsCharacterDatum;
    character_id?: number;
    created_at?: string;
    event_data?: string;
    event_type_id?: number;
    event_type_name?: string;
    heading?: number;
    id?: number;
    instance_id?: number;
    x?: number;
    y?: number;
    z?: number;
    zone?: ModelsZone;
    zone_id?: number;
}
