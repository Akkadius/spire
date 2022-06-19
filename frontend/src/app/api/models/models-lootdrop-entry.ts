/* tslint:disable */
/* eslint-disable */
/**
 * Spire
 * Spire API documentation
 *
 * The version of the OpenAPI document: 3.0
 * Contact: akkadius1@gmail.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


import { ModelsItem } from './models-item';
import { ModelsLootdrop } from './models-lootdrop';

/**
 * 
 * @export
 * @interface ModelsLootdropEntry
 */
export interface ModelsLootdropEntry {
    /**
     * 
     * @type {number}
     * @memberof ModelsLootdropEntry
     */
    chance?: number;
    /**
     * 
     * @type {number}
     * @memberof ModelsLootdropEntry
     */
    disabled_chance?: number;
    /**
     * 
     * @type {number}
     * @memberof ModelsLootdropEntry
     */
    equip_item?: number;
    /**
     * 
     * @type {ModelsItem}
     * @memberof ModelsLootdropEntry
     */
    item?: ModelsItem;
    /**
     * 
     * @type {number}
     * @memberof ModelsLootdropEntry
     */
    item_charges?: number;
    /**
     * 
     * @type {number}
     * @memberof ModelsLootdropEntry
     */
    item_id?: number;
    /**
     * 
     * @type {ModelsLootdrop}
     * @memberof ModelsLootdropEntry
     */
    lootdrop?: ModelsLootdrop;
    /**
     * 
     * @type {number}
     * @memberof ModelsLootdropEntry
     */
    lootdrop_id?: number;
    /**
     * 
     * @type {number}
     * @memberof ModelsLootdropEntry
     */
    multiplier?: number;
    /**
     * 
     * @type {number}
     * @memberof ModelsLootdropEntry
     */
    npc_max_level?: number;
    /**
     * 
     * @type {number}
     * @memberof ModelsLootdropEntry
     */
    npc_min_level?: number;
    /**
     * 
     * @type {number}
     * @memberof ModelsLootdropEntry
     */
    trivial_max_level?: number;
    /**
     * 
     * @type {number}
     * @memberof ModelsLootdropEntry
     */
    trivial_min_level?: number;
}


