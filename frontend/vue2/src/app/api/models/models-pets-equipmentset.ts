import { ModelsPetsEquipmentsetEntry } from './models-pets-equipmentset-entry';
export interface ModelsPetsEquipmentset {
    nested_set?: number;
    pets_equipmentset_entries?: Array<ModelsPetsEquipmentsetEntry>;
    set_id?: number;
    setname?: string;
}
