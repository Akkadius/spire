import { ModelsNpcType } from './models-npc-type';
export interface ModelsPet {
    equipmentset?: number;
    id?: number;
    monsterflag?: number;
    npc_id?: number;
    npc_type?: ModelsNpcType;
    petcontrol?: number;
    petnaming?: number;
    petpower?: number;
    temp?: number;
    type?: string;
}
