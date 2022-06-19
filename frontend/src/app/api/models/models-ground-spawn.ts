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


import { ModelsZone } from './models-zone';

/**
 * 
 * @export
 * @interface ModelsGroundSpawn
 */
export interface ModelsGroundSpawn {
    /**
     * 
     * @type {string}
     * @memberof ModelsGroundSpawn
     */
    comment?: string;
    /**
     * 
     * @type {string}
     * @memberof ModelsGroundSpawn
     */
    content_flags?: string;
    /**
     * 
     * @type {string}
     * @memberof ModelsGroundSpawn
     */
    content_flags_disabled?: string;
    /**
     * 
     * @type {number}
     * @memberof ModelsGroundSpawn
     */
    heading?: number;
    /**
     * 
     * @type {number}
     * @memberof ModelsGroundSpawn
     */
    id?: number;
    /**
     * 
     * @type {number}
     * @memberof ModelsGroundSpawn
     */
    item?: number;
    /**
     * 
     * @type {number}
     * @memberof ModelsGroundSpawn
     */
    max_allowed?: number;
    /**
     * 
     * @type {number}
     * @memberof ModelsGroundSpawn
     */
    max_expansion?: number;
    /**
     * 
     * @type {number}
     * @memberof ModelsGroundSpawn
     */
    max_x?: number;
    /**
     * 
     * @type {number}
     * @memberof ModelsGroundSpawn
     */
    max_y?: number;
    /**
     * 
     * @type {number}
     * @memberof ModelsGroundSpawn
     */
    max_z?: number;
    /**
     * 
     * @type {number}
     * @memberof ModelsGroundSpawn
     */
    min_expansion?: number;
    /**
     * 
     * @type {number}
     * @memberof ModelsGroundSpawn
     */
    min_x?: number;
    /**
     * 
     * @type {number}
     * @memberof ModelsGroundSpawn
     */
    min_y?: number;
    /**
     * 
     * @type {string}
     * @memberof ModelsGroundSpawn
     */
    name?: string;
    /**
     * 
     * @type {number}
     * @memberof ModelsGroundSpawn
     */
    respawn_timer?: number;
    /**
     * 
     * @type {number}
     * @memberof ModelsGroundSpawn
     */
    version?: number;
    /**
     * 
     * @type {ModelsZone}
     * @memberof ModelsGroundSpawn
     */
    zone?: ModelsZone;
    /**
     * 
     * @type {number}
     * @memberof ModelsGroundSpawn
     */
    zoneid?: number;
}


