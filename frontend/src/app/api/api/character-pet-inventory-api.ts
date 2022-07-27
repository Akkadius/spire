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


import globalAxios, { AxiosPromise, AxiosInstance } from 'axios';
import { Configuration } from '../configuration';
// Some imports not used depending on template conditions
// @ts-ignore
import { BASE_PATH, COLLECTION_FORMATS, RequestArgs, BaseAPI, RequiredError } from '../base';
// @ts-ignore
import { CrudcontrollersBulkFetchByIdsGetRequest } from '../models';
// @ts-ignore
import { ModelsCharacterPetInventory } from '../models';
/**
 * CharacterPetInventoryApi - axios parameter creator
 * @export
 */
export const CharacterPetInventoryApiAxiosParamCreator = function (configuration?: Configuration) {
    return {
        /**
         * 
         * @summary Creates CharacterPetInventory
         * @param {ModelsCharacterPetInventory} characterPetInventory CharacterPetInventory
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        createCharacterPetInventory: async (characterPetInventory: ModelsCharacterPetInventory, options: any = {}): Promise<RequestArgs> => {
            // verify required parameter 'characterPetInventory' is not null or undefined
            if (characterPetInventory === null || characterPetInventory === undefined) {
                throw new RequiredError('characterPetInventory','Required parameter characterPetInventory was null or undefined when calling createCharacterPetInventory.');
            }
            const localVarPath = `/character_pet_inventory`;
            // use dummy base URL string because the URL constructor only accepts absolute URLs.
            const localVarUrlObj = new URL(localVarPath, 'https://example.com');
            let baseOptions;
            if (configuration) {
                baseOptions = configuration.baseOptions;
            }

            const localVarRequestOptions = { method: 'PUT', ...baseOptions, ...options};
            const localVarHeaderParameter = {} as any;
            const localVarQueryParameter = {} as any;


    
            localVarHeaderParameter['Content-Type'] = 'application/json';

            const queryParameters = new URLSearchParams(localVarUrlObj.search);
            for (const key in localVarQueryParameter) {
                queryParameters.set(key, localVarQueryParameter[key]);
            }
            for (const key in options.query) {
                queryParameters.set(key, options.query[key]);
            }
            localVarUrlObj.search = (new URLSearchParams(queryParameters)).toString();
            let headersFromBaseOptions = baseOptions && baseOptions.headers ? baseOptions.headers : {};
            localVarRequestOptions.headers = {...localVarHeaderParameter, ...headersFromBaseOptions, ...options.headers};
            const nonString = typeof characterPetInventory !== 'string';
            const needsSerialization = nonString && configuration && configuration.isJsonMime
                ? configuration.isJsonMime(localVarRequestOptions.headers['Content-Type'])
                : nonString;
            localVarRequestOptions.data =  needsSerialization
                ? JSON.stringify(characterPetInventory !== undefined ? characterPetInventory : {})
                : (characterPetInventory || "");

            return {
                url: localVarUrlObj.pathname + localVarUrlObj.search + localVarUrlObj.hash,
                options: localVarRequestOptions,
            };
        },
        /**
         * 
         * @summary Deletes CharacterPetInventory
         * @param {number} id charId
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        deleteCharacterPetInventory: async (id: number, options: any = {}): Promise<RequestArgs> => {
            // verify required parameter 'id' is not null or undefined
            if (id === null || id === undefined) {
                throw new RequiredError('id','Required parameter id was null or undefined when calling deleteCharacterPetInventory.');
            }
            const localVarPath = `/character_pet_inventory/{id}`
                .replace(`{${"id"}}`, encodeURIComponent(String(id)));
            // use dummy base URL string because the URL constructor only accepts absolute URLs.
            const localVarUrlObj = new URL(localVarPath, 'https://example.com');
            let baseOptions;
            if (configuration) {
                baseOptions = configuration.baseOptions;
            }

            const localVarRequestOptions = { method: 'DELETE', ...baseOptions, ...options};
            const localVarHeaderParameter = {} as any;
            const localVarQueryParameter = {} as any;


    
            const queryParameters = new URLSearchParams(localVarUrlObj.search);
            for (const key in localVarQueryParameter) {
                queryParameters.set(key, localVarQueryParameter[key]);
            }
            for (const key in options.query) {
                queryParameters.set(key, options.query[key]);
            }
            localVarUrlObj.search = (new URLSearchParams(queryParameters)).toString();
            let headersFromBaseOptions = baseOptions && baseOptions.headers ? baseOptions.headers : {};
            localVarRequestOptions.headers = {...localVarHeaderParameter, ...headersFromBaseOptions, ...options.headers};

            return {
                url: localVarUrlObj.pathname + localVarUrlObj.search + localVarUrlObj.hash,
                options: localVarRequestOptions,
            };
        },
        /**
         * 
         * @summary Gets CharacterPetInventories in bulk
         * @param {CrudcontrollersBulkFetchByIdsGetRequest} body body
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        getCharacterPetInventoriesBulk: async (body: CrudcontrollersBulkFetchByIdsGetRequest, options: any = {}): Promise<RequestArgs> => {
            // verify required parameter 'body' is not null or undefined
            if (body === null || body === undefined) {
                throw new RequiredError('body','Required parameter body was null or undefined when calling getCharacterPetInventoriesBulk.');
            }
            const localVarPath = `/character_pet_inventories/bulk`;
            // use dummy base URL string because the URL constructor only accepts absolute URLs.
            const localVarUrlObj = new URL(localVarPath, 'https://example.com');
            let baseOptions;
            if (configuration) {
                baseOptions = configuration.baseOptions;
            }

            const localVarRequestOptions = { method: 'POST', ...baseOptions, ...options};
            const localVarHeaderParameter = {} as any;
            const localVarQueryParameter = {} as any;


    
            localVarHeaderParameter['Content-Type'] = 'application/json';

            const queryParameters = new URLSearchParams(localVarUrlObj.search);
            for (const key in localVarQueryParameter) {
                queryParameters.set(key, localVarQueryParameter[key]);
            }
            for (const key in options.query) {
                queryParameters.set(key, options.query[key]);
            }
            localVarUrlObj.search = (new URLSearchParams(queryParameters)).toString();
            let headersFromBaseOptions = baseOptions && baseOptions.headers ? baseOptions.headers : {};
            localVarRequestOptions.headers = {...localVarHeaderParameter, ...headersFromBaseOptions, ...options.headers};
            const nonString = typeof body !== 'string';
            const needsSerialization = nonString && configuration && configuration.isJsonMime
                ? configuration.isJsonMime(localVarRequestOptions.headers['Content-Type'])
                : nonString;
            localVarRequestOptions.data =  needsSerialization
                ? JSON.stringify(body !== undefined ? body : {})
                : (body || "");

            return {
                url: localVarUrlObj.pathname + localVarUrlObj.search + localVarUrlObj.hash,
                options: localVarRequestOptions,
            };
        },
        /**
         * 
         * @summary Gets CharacterPetInventory
         * @param {number} id Id
         * @param {string} [includes] Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names 
         * @param {string} [select] Column names [.] separated to fetch specific fields in response
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        getCharacterPetInventory: async (id: number, includes?: string, select?: string, options: any = {}): Promise<RequestArgs> => {
            // verify required parameter 'id' is not null or undefined
            if (id === null || id === undefined) {
                throw new RequiredError('id','Required parameter id was null or undefined when calling getCharacterPetInventory.');
            }
            const localVarPath = `/character_pet_inventory/{id}`
                .replace(`{${"id"}}`, encodeURIComponent(String(id)));
            // use dummy base URL string because the URL constructor only accepts absolute URLs.
            const localVarUrlObj = new URL(localVarPath, 'https://example.com');
            let baseOptions;
            if (configuration) {
                baseOptions = configuration.baseOptions;
            }

            const localVarRequestOptions = { method: 'GET', ...baseOptions, ...options};
            const localVarHeaderParameter = {} as any;
            const localVarQueryParameter = {} as any;

            if (includes !== undefined) {
                localVarQueryParameter['includes'] = includes;
            }

            if (select !== undefined) {
                localVarQueryParameter['select'] = select;
            }


    
            const queryParameters = new URLSearchParams(localVarUrlObj.search);
            for (const key in localVarQueryParameter) {
                queryParameters.set(key, localVarQueryParameter[key]);
            }
            for (const key in options.query) {
                queryParameters.set(key, options.query[key]);
            }
            localVarUrlObj.search = (new URLSearchParams(queryParameters)).toString();
            let headersFromBaseOptions = baseOptions && baseOptions.headers ? baseOptions.headers : {};
            localVarRequestOptions.headers = {...localVarHeaderParameter, ...headersFromBaseOptions, ...options.headers};

            return {
                url: localVarUrlObj.pathname + localVarUrlObj.search + localVarUrlObj.hash,
                options: localVarRequestOptions,
            };
        },
        /**
         * 
         * @summary Lists CharacterPetInventories
         * @param {string} [includes] Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names 
         * @param {string} [where] Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2
         * @param {string} [whereOr] Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2
         * @param {string} [groupBy] Group by field. Multiple conditions [.] separated Example: field1.field2
         * @param {string} [limit] Rows to limit in response (Default: 10,000)
         * @param {number} [page] Pagination page
         * @param {string} [orderBy] Order by [field]
         * @param {string} [orderDirection] Order by field direction
         * @param {string} [select] Column names [.] separated to fetch specific fields in response
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        listCharacterPetInventories: async (includes?: string, where?: string, whereOr?: string, groupBy?: string, limit?: string, page?: number, orderBy?: string, orderDirection?: string, select?: string, options: any = {}): Promise<RequestArgs> => {
            const localVarPath = `/character_pet_inventories`;
            // use dummy base URL string because the URL constructor only accepts absolute URLs.
            const localVarUrlObj = new URL(localVarPath, 'https://example.com');
            let baseOptions;
            if (configuration) {
                baseOptions = configuration.baseOptions;
            }

            const localVarRequestOptions = { method: 'GET', ...baseOptions, ...options};
            const localVarHeaderParameter = {} as any;
            const localVarQueryParameter = {} as any;

            if (includes !== undefined) {
                localVarQueryParameter['includes'] = includes;
            }

            if (where !== undefined) {
                localVarQueryParameter['where'] = where;
            }

            if (whereOr !== undefined) {
                localVarQueryParameter['whereOr'] = whereOr;
            }

            if (groupBy !== undefined) {
                localVarQueryParameter['groupBy'] = groupBy;
            }

            if (limit !== undefined) {
                localVarQueryParameter['limit'] = limit;
            }

            if (page !== undefined) {
                localVarQueryParameter['page'] = page;
            }

            if (orderBy !== undefined) {
                localVarQueryParameter['orderBy'] = orderBy;
            }

            if (orderDirection !== undefined) {
                localVarQueryParameter['orderDirection'] = orderDirection;
            }

            if (select !== undefined) {
                localVarQueryParameter['select'] = select;
            }


    
            const queryParameters = new URLSearchParams(localVarUrlObj.search);
            for (const key in localVarQueryParameter) {
                queryParameters.set(key, localVarQueryParameter[key]);
            }
            for (const key in options.query) {
                queryParameters.set(key, options.query[key]);
            }
            localVarUrlObj.search = (new URLSearchParams(queryParameters)).toString();
            let headersFromBaseOptions = baseOptions && baseOptions.headers ? baseOptions.headers : {};
            localVarRequestOptions.headers = {...localVarHeaderParameter, ...headersFromBaseOptions, ...options.headers};

            return {
                url: localVarUrlObj.pathname + localVarUrlObj.search + localVarUrlObj.hash,
                options: localVarRequestOptions,
            };
        },
        /**
         * 
         * @summary Updates CharacterPetInventory
         * @param {number} id Id
         * @param {ModelsCharacterPetInventory} characterPetInventory CharacterPetInventory
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        updateCharacterPetInventory: async (id: number, characterPetInventory: ModelsCharacterPetInventory, options: any = {}): Promise<RequestArgs> => {
            // verify required parameter 'id' is not null or undefined
            if (id === null || id === undefined) {
                throw new RequiredError('id','Required parameter id was null or undefined when calling updateCharacterPetInventory.');
            }
            // verify required parameter 'characterPetInventory' is not null or undefined
            if (characterPetInventory === null || characterPetInventory === undefined) {
                throw new RequiredError('characterPetInventory','Required parameter characterPetInventory was null or undefined when calling updateCharacterPetInventory.');
            }
            const localVarPath = `/character_pet_inventory/{id}`
                .replace(`{${"id"}}`, encodeURIComponent(String(id)));
            // use dummy base URL string because the URL constructor only accepts absolute URLs.
            const localVarUrlObj = new URL(localVarPath, 'https://example.com');
            let baseOptions;
            if (configuration) {
                baseOptions = configuration.baseOptions;
            }

            const localVarRequestOptions = { method: 'PATCH', ...baseOptions, ...options};
            const localVarHeaderParameter = {} as any;
            const localVarQueryParameter = {} as any;


    
            localVarHeaderParameter['Content-Type'] = 'application/json';

            const queryParameters = new URLSearchParams(localVarUrlObj.search);
            for (const key in localVarQueryParameter) {
                queryParameters.set(key, localVarQueryParameter[key]);
            }
            for (const key in options.query) {
                queryParameters.set(key, options.query[key]);
            }
            localVarUrlObj.search = (new URLSearchParams(queryParameters)).toString();
            let headersFromBaseOptions = baseOptions && baseOptions.headers ? baseOptions.headers : {};
            localVarRequestOptions.headers = {...localVarHeaderParameter, ...headersFromBaseOptions, ...options.headers};
            const nonString = typeof characterPetInventory !== 'string';
            const needsSerialization = nonString && configuration && configuration.isJsonMime
                ? configuration.isJsonMime(localVarRequestOptions.headers['Content-Type'])
                : nonString;
            localVarRequestOptions.data =  needsSerialization
                ? JSON.stringify(characterPetInventory !== undefined ? characterPetInventory : {})
                : (characterPetInventory || "");

            return {
                url: localVarUrlObj.pathname + localVarUrlObj.search + localVarUrlObj.hash,
                options: localVarRequestOptions,
            };
        },
    }
};

/**
 * CharacterPetInventoryApi - functional programming interface
 * @export
 */
export const CharacterPetInventoryApiFp = function(configuration?: Configuration) {
    return {
        /**
         * 
         * @summary Creates CharacterPetInventory
         * @param {ModelsCharacterPetInventory} characterPetInventory CharacterPetInventory
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async createCharacterPetInventory(characterPetInventory: ModelsCharacterPetInventory, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsCharacterPetInventory>>> {
            const localVarAxiosArgs = await CharacterPetInventoryApiAxiosParamCreator(configuration).createCharacterPetInventory(characterPetInventory, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        /**
         * 
         * @summary Deletes CharacterPetInventory
         * @param {number} id charId
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async deleteCharacterPetInventory(id: number, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<string>> {
            const localVarAxiosArgs = await CharacterPetInventoryApiAxiosParamCreator(configuration).deleteCharacterPetInventory(id, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        /**
         * 
         * @summary Gets CharacterPetInventories in bulk
         * @param {CrudcontrollersBulkFetchByIdsGetRequest} body body
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async getCharacterPetInventoriesBulk(body: CrudcontrollersBulkFetchByIdsGetRequest, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsCharacterPetInventory>>> {
            const localVarAxiosArgs = await CharacterPetInventoryApiAxiosParamCreator(configuration).getCharacterPetInventoriesBulk(body, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        /**
         * 
         * @summary Gets CharacterPetInventory
         * @param {number} id Id
         * @param {string} [includes] Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names 
         * @param {string} [select] Column names [.] separated to fetch specific fields in response
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async getCharacterPetInventory(id: number, includes?: string, select?: string, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsCharacterPetInventory>>> {
            const localVarAxiosArgs = await CharacterPetInventoryApiAxiosParamCreator(configuration).getCharacterPetInventory(id, includes, select, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        /**
         * 
         * @summary Lists CharacterPetInventories
         * @param {string} [includes] Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names 
         * @param {string} [where] Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2
         * @param {string} [whereOr] Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2
         * @param {string} [groupBy] Group by field. Multiple conditions [.] separated Example: field1.field2
         * @param {string} [limit] Rows to limit in response (Default: 10,000)
         * @param {number} [page] Pagination page
         * @param {string} [orderBy] Order by [field]
         * @param {string} [orderDirection] Order by field direction
         * @param {string} [select] Column names [.] separated to fetch specific fields in response
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async listCharacterPetInventories(includes?: string, where?: string, whereOr?: string, groupBy?: string, limit?: string, page?: number, orderBy?: string, orderDirection?: string, select?: string, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsCharacterPetInventory>>> {
            const localVarAxiosArgs = await CharacterPetInventoryApiAxiosParamCreator(configuration).listCharacterPetInventories(includes, where, whereOr, groupBy, limit, page, orderBy, orderDirection, select, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        /**
         * 
         * @summary Updates CharacterPetInventory
         * @param {number} id Id
         * @param {ModelsCharacterPetInventory} characterPetInventory CharacterPetInventory
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async updateCharacterPetInventory(id: number, characterPetInventory: ModelsCharacterPetInventory, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsCharacterPetInventory>>> {
            const localVarAxiosArgs = await CharacterPetInventoryApiAxiosParamCreator(configuration).updateCharacterPetInventory(id, characterPetInventory, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
    }
};

/**
 * CharacterPetInventoryApi - factory interface
 * @export
 */
export const CharacterPetInventoryApiFactory = function (configuration?: Configuration, basePath?: string, axios?: AxiosInstance) {
    return {
        /**
         * 
         * @summary Creates CharacterPetInventory
         * @param {ModelsCharacterPetInventory} characterPetInventory CharacterPetInventory
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        createCharacterPetInventory(characterPetInventory: ModelsCharacterPetInventory, options?: any): AxiosPromise<Array<ModelsCharacterPetInventory>> {
            return CharacterPetInventoryApiFp(configuration).createCharacterPetInventory(characterPetInventory, options).then((request) => request(axios, basePath));
        },
        /**
         * 
         * @summary Deletes CharacterPetInventory
         * @param {number} id charId
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        deleteCharacterPetInventory(id: number, options?: any): AxiosPromise<string> {
            return CharacterPetInventoryApiFp(configuration).deleteCharacterPetInventory(id, options).then((request) => request(axios, basePath));
        },
        /**
         * 
         * @summary Gets CharacterPetInventories in bulk
         * @param {CrudcontrollersBulkFetchByIdsGetRequest} body body
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        getCharacterPetInventoriesBulk(body: CrudcontrollersBulkFetchByIdsGetRequest, options?: any): AxiosPromise<Array<ModelsCharacterPetInventory>> {
            return CharacterPetInventoryApiFp(configuration).getCharacterPetInventoriesBulk(body, options).then((request) => request(axios, basePath));
        },
        /**
         * 
         * @summary Gets CharacterPetInventory
         * @param {number} id Id
         * @param {string} [includes] Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names 
         * @param {string} [select] Column names [.] separated to fetch specific fields in response
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        getCharacterPetInventory(id: number, includes?: string, select?: string, options?: any): AxiosPromise<Array<ModelsCharacterPetInventory>> {
            return CharacterPetInventoryApiFp(configuration).getCharacterPetInventory(id, includes, select, options).then((request) => request(axios, basePath));
        },
        /**
         * 
         * @summary Lists CharacterPetInventories
         * @param {string} [includes] Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names 
         * @param {string} [where] Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2
         * @param {string} [whereOr] Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2
         * @param {string} [groupBy] Group by field. Multiple conditions [.] separated Example: field1.field2
         * @param {string} [limit] Rows to limit in response (Default: 10,000)
         * @param {number} [page] Pagination page
         * @param {string} [orderBy] Order by [field]
         * @param {string} [orderDirection] Order by field direction
         * @param {string} [select] Column names [.] separated to fetch specific fields in response
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        listCharacterPetInventories(includes?: string, where?: string, whereOr?: string, groupBy?: string, limit?: string, page?: number, orderBy?: string, orderDirection?: string, select?: string, options?: any): AxiosPromise<Array<ModelsCharacterPetInventory>> {
            return CharacterPetInventoryApiFp(configuration).listCharacterPetInventories(includes, where, whereOr, groupBy, limit, page, orderBy, orderDirection, select, options).then((request) => request(axios, basePath));
        },
        /**
         * 
         * @summary Updates CharacterPetInventory
         * @param {number} id Id
         * @param {ModelsCharacterPetInventory} characterPetInventory CharacterPetInventory
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        updateCharacterPetInventory(id: number, characterPetInventory: ModelsCharacterPetInventory, options?: any): AxiosPromise<Array<ModelsCharacterPetInventory>> {
            return CharacterPetInventoryApiFp(configuration).updateCharacterPetInventory(id, characterPetInventory, options).then((request) => request(axios, basePath));
        },
    };
};

/**
 * Request parameters for createCharacterPetInventory operation in CharacterPetInventoryApi.
 * @export
 * @interface CharacterPetInventoryApiCreateCharacterPetInventoryRequest
 */
export interface CharacterPetInventoryApiCreateCharacterPetInventoryRequest {
    /**
     * CharacterPetInventory
     * @type {ModelsCharacterPetInventory}
     * @memberof CharacterPetInventoryApiCreateCharacterPetInventory
     */
    readonly characterPetInventory: ModelsCharacterPetInventory
}

/**
 * Request parameters for deleteCharacterPetInventory operation in CharacterPetInventoryApi.
 * @export
 * @interface CharacterPetInventoryApiDeleteCharacterPetInventoryRequest
 */
export interface CharacterPetInventoryApiDeleteCharacterPetInventoryRequest {
    /**
     * charId
     * @type {number}
     * @memberof CharacterPetInventoryApiDeleteCharacterPetInventory
     */
    readonly id: number
}

/**
 * Request parameters for getCharacterPetInventoriesBulk operation in CharacterPetInventoryApi.
 * @export
 * @interface CharacterPetInventoryApiGetCharacterPetInventoriesBulkRequest
 */
export interface CharacterPetInventoryApiGetCharacterPetInventoriesBulkRequest {
    /**
     * body
     * @type {CrudcontrollersBulkFetchByIdsGetRequest}
     * @memberof CharacterPetInventoryApiGetCharacterPetInventoriesBulk
     */
    readonly body: CrudcontrollersBulkFetchByIdsGetRequest
}

/**
 * Request parameters for getCharacterPetInventory operation in CharacterPetInventoryApi.
 * @export
 * @interface CharacterPetInventoryApiGetCharacterPetInventoryRequest
 */
export interface CharacterPetInventoryApiGetCharacterPetInventoryRequest {
    /**
     * Id
     * @type {number}
     * @memberof CharacterPetInventoryApiGetCharacterPetInventory
     */
    readonly id: number

    /**
     * Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names 
     * @type {string}
     * @memberof CharacterPetInventoryApiGetCharacterPetInventory
     */
    readonly includes?: string

    /**
     * Column names [.] separated to fetch specific fields in response
     * @type {string}
     * @memberof CharacterPetInventoryApiGetCharacterPetInventory
     */
    readonly select?: string
}

/**
 * Request parameters for listCharacterPetInventories operation in CharacterPetInventoryApi.
 * @export
 * @interface CharacterPetInventoryApiListCharacterPetInventoriesRequest
 */
export interface CharacterPetInventoryApiListCharacterPetInventoriesRequest {
    /**
     * Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names 
     * @type {string}
     * @memberof CharacterPetInventoryApiListCharacterPetInventories
     */
    readonly includes?: string

    /**
     * Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2
     * @type {string}
     * @memberof CharacterPetInventoryApiListCharacterPetInventories
     */
    readonly where?: string

    /**
     * Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2
     * @type {string}
     * @memberof CharacterPetInventoryApiListCharacterPetInventories
     */
    readonly whereOr?: string

    /**
     * Group by field. Multiple conditions [.] separated Example: field1.field2
     * @type {string}
     * @memberof CharacterPetInventoryApiListCharacterPetInventories
     */
    readonly groupBy?: string

    /**
     * Rows to limit in response (Default: 10,000)
     * @type {string}
     * @memberof CharacterPetInventoryApiListCharacterPetInventories
     */
    readonly limit?: string

    /**
     * Pagination page
     * @type {number}
     * @memberof CharacterPetInventoryApiListCharacterPetInventories
     */
    readonly page?: number

    /**
     * Order by [field]
     * @type {string}
     * @memberof CharacterPetInventoryApiListCharacterPetInventories
     */
    readonly orderBy?: string

    /**
     * Order by field direction
     * @type {string}
     * @memberof CharacterPetInventoryApiListCharacterPetInventories
     */
    readonly orderDirection?: string

    /**
     * Column names [.] separated to fetch specific fields in response
     * @type {string}
     * @memberof CharacterPetInventoryApiListCharacterPetInventories
     */
    readonly select?: string
}

/**
 * Request parameters for updateCharacterPetInventory operation in CharacterPetInventoryApi.
 * @export
 * @interface CharacterPetInventoryApiUpdateCharacterPetInventoryRequest
 */
export interface CharacterPetInventoryApiUpdateCharacterPetInventoryRequest {
    /**
     * Id
     * @type {number}
     * @memberof CharacterPetInventoryApiUpdateCharacterPetInventory
     */
    readonly id: number

    /**
     * CharacterPetInventory
     * @type {ModelsCharacterPetInventory}
     * @memberof CharacterPetInventoryApiUpdateCharacterPetInventory
     */
    readonly characterPetInventory: ModelsCharacterPetInventory
}

/**
 * CharacterPetInventoryApi - object-oriented interface
 * @export
 * @class CharacterPetInventoryApi
 * @extends {BaseAPI}
 */
export class CharacterPetInventoryApi extends BaseAPI {
    /**
     * 
     * @summary Creates CharacterPetInventory
     * @param {CharacterPetInventoryApiCreateCharacterPetInventoryRequest} requestParameters Request parameters.
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof CharacterPetInventoryApi
     */
    public createCharacterPetInventory(requestParameters: CharacterPetInventoryApiCreateCharacterPetInventoryRequest, options?: any) {
        return CharacterPetInventoryApiFp(this.configuration).createCharacterPetInventory(requestParameters.characterPetInventory, options).then((request) => request(this.axios, this.basePath));
    }

    /**
     * 
     * @summary Deletes CharacterPetInventory
     * @param {CharacterPetInventoryApiDeleteCharacterPetInventoryRequest} requestParameters Request parameters.
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof CharacterPetInventoryApi
     */
    public deleteCharacterPetInventory(requestParameters: CharacterPetInventoryApiDeleteCharacterPetInventoryRequest, options?: any) {
        return CharacterPetInventoryApiFp(this.configuration).deleteCharacterPetInventory(requestParameters.id, options).then((request) => request(this.axios, this.basePath));
    }

    /**
     * 
     * @summary Gets CharacterPetInventories in bulk
     * @param {CharacterPetInventoryApiGetCharacterPetInventoriesBulkRequest} requestParameters Request parameters.
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof CharacterPetInventoryApi
     */
    public getCharacterPetInventoriesBulk(requestParameters: CharacterPetInventoryApiGetCharacterPetInventoriesBulkRequest, options?: any) {
        return CharacterPetInventoryApiFp(this.configuration).getCharacterPetInventoriesBulk(requestParameters.body, options).then((request) => request(this.axios, this.basePath));
    }

    /**
     * 
     * @summary Gets CharacterPetInventory
     * @param {CharacterPetInventoryApiGetCharacterPetInventoryRequest} requestParameters Request parameters.
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof CharacterPetInventoryApi
     */
    public getCharacterPetInventory(requestParameters: CharacterPetInventoryApiGetCharacterPetInventoryRequest, options?: any) {
        return CharacterPetInventoryApiFp(this.configuration).getCharacterPetInventory(requestParameters.id, requestParameters.includes, requestParameters.select, options).then((request) => request(this.axios, this.basePath));
    }

    /**
     * 
     * @summary Lists CharacterPetInventories
     * @param {CharacterPetInventoryApiListCharacterPetInventoriesRequest} requestParameters Request parameters.
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof CharacterPetInventoryApi
     */
    public listCharacterPetInventories(requestParameters: CharacterPetInventoryApiListCharacterPetInventoriesRequest = {}, options?: any) {
        return CharacterPetInventoryApiFp(this.configuration).listCharacterPetInventories(requestParameters.includes, requestParameters.where, requestParameters.whereOr, requestParameters.groupBy, requestParameters.limit, requestParameters.page, requestParameters.orderBy, requestParameters.orderDirection, requestParameters.select, options).then((request) => request(this.axios, this.basePath));
    }

    /**
     * 
     * @summary Updates CharacterPetInventory
     * @param {CharacterPetInventoryApiUpdateCharacterPetInventoryRequest} requestParameters Request parameters.
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof CharacterPetInventoryApi
     */
    public updateCharacterPetInventory(requestParameters: CharacterPetInventoryApiUpdateCharacterPetInventoryRequest, options?: any) {
        return CharacterPetInventoryApiFp(this.configuration).updateCharacterPetInventory(requestParameters.id, requestParameters.characterPetInventory, options).then((request) => request(this.axios, this.basePath));
    }
}
