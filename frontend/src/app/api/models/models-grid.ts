import { ModelsGridEntry } from './models-grid-entry';
import { ModelsZone } from './models-zone';
export interface ModelsGrid {
    grid_entries?: Array<ModelsGridEntry>;
    id?: number;
    type?: number;
    type_2?: number;
    zone?: ModelsZone;
    zoneid?: number;
}
