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


import { ModelsRaidLeader } from './models-raid-leader';
import { ModelsRaidMember } from './models-raid-member';

/**
 * 
 * @export
 * @interface ModelsRaidDetail
 */
export interface ModelsRaidDetail {
    /**
     * 
     * @type {number}
     * @memberof ModelsRaidDetail
     */
    locked?: number;
    /**
     * 
     * @type {number}
     * @memberof ModelsRaidDetail
     */
    loottype?: number;
    /**
     * 
     * @type {string}
     * @memberof ModelsRaidDetail
     */
    motd?: string;
    /**
     * 
     * @type {Array<ModelsRaidLeader>}
     * @memberof ModelsRaidDetail
     */
    raid_leaders?: Array<ModelsRaidLeader>;
    /**
     * 
     * @type {Array<ModelsRaidMember>}
     * @memberof ModelsRaidDetail
     */
    raid_members?: Array<ModelsRaidMember>;
    /**
     * 
     * @type {number}
     * @memberof ModelsRaidDetail
     */
    raidid?: number;
}


