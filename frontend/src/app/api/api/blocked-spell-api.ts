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
import { ModelsBlockedSpell } from '../models';
/**
 * BlockedSpellApi - axios parameter creator
 * @export
 */
export const BlockedSpellApiAxiosParamCreator = function (configuration?: Configuration) {
    return {
        /**
         * 
         * @summary Creates BlockedSpell
         * @param {ModelsBlockedSpell} blockedSpell BlockedSpell
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        createBlockedSpell: async (blockedSpell: ModelsBlockedSpell, options: any = {}): Promise<RequestArgs> => {
            // verify required parameter 'blockedSpell' is not null or undefined
            if (blockedSpell === null || blockedSpell === undefined) {
                throw new RequiredError('blockedSpell','Required parameter blockedSpell was null or undefined when calling createBlockedSpell.');
            }
            const localVarPath = `/blocked_spell`;
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
            const nonString = typeof blockedSpell !== 'string';
            const needsSerialization = nonString && configuration && configuration.isJsonMime
                ? configuration.isJsonMime(localVarRequestOptions.headers['Content-Type'])
                : nonString;
            localVarRequestOptions.data =  needsSerialization
                ? JSON.stringify(blockedSpell !== undefined ? blockedSpell : {})
                : (blockedSpell || "");

            return {
                url: localVarUrlObj.pathname + localVarUrlObj.search + localVarUrlObj.hash,
                options: localVarRequestOptions,
            };
        },
        /**
         * 
         * @summary Deletes BlockedSpell
         * @param {number} id Id
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        deleteBlockedSpell: async (id: number, options: any = {}): Promise<RequestArgs> => {
            // verify required parameter 'id' is not null or undefined
            if (id === null || id === undefined) {
                throw new RequiredError('id','Required parameter id was null or undefined when calling deleteBlockedSpell.');
            }
            const localVarPath = `/blocked_spell/{id}`
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
         * @summary Gets BlockedSpell
         * @param {number} id Id
         * @param {string} [includes] Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names 
         * @param {string} [select] Column names [.] separated to fetch specific fields in response
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        getBlockedSpell: async (id: number, includes?: string, select?: string, options: any = {}): Promise<RequestArgs> => {
            // verify required parameter 'id' is not null or undefined
            if (id === null || id === undefined) {
                throw new RequiredError('id','Required parameter id was null or undefined when calling getBlockedSpell.');
            }
            const localVarPath = `/blocked_spell/{id}`
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
         * @summary Gets BlockedSpells in bulk
         * @param {CrudcontrollersBulkFetchByIdsGetRequest} body body
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        getBlockedSpellsBulk: async (body: CrudcontrollersBulkFetchByIdsGetRequest, options: any = {}): Promise<RequestArgs> => {
            // verify required parameter 'body' is not null or undefined
            if (body === null || body === undefined) {
                throw new RequiredError('body','Required parameter body was null or undefined when calling getBlockedSpellsBulk.');
            }
            const localVarPath = `/blocked_spells/bulk`;
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
         * @summary Lists BlockedSpells
         * @param {string} [includes] Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names 
         * @param {string} [where] Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2
         * @param {string} [whereOr] Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2
         * @param {string} [groupBy] Group by field. Multiple conditions [.] separated Example: field1.field2
         * @param {string} [limit] Rows to limit in response (Default: 10,000)
         * @param {string} [orderBy] Order by [field]
         * @param {string} [orderDirection] Order by field direction
         * @param {string} [select] Column names [.] separated to fetch specific fields in response
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        listBlockedSpells: async (includes?: string, where?: string, whereOr?: string, groupBy?: string, limit?: string, orderBy?: string, orderDirection?: string, select?: string, options: any = {}): Promise<RequestArgs> => {
            const localVarPath = `/blocked_spells`;
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
         * @summary Updates BlockedSpell
         * @param {number} id Id
         * @param {ModelsBlockedSpell} blockedSpell BlockedSpell
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        updateBlockedSpell: async (id: number, blockedSpell: ModelsBlockedSpell, options: any = {}): Promise<RequestArgs> => {
            // verify required parameter 'id' is not null or undefined
            if (id === null || id === undefined) {
                throw new RequiredError('id','Required parameter id was null or undefined when calling updateBlockedSpell.');
            }
            // verify required parameter 'blockedSpell' is not null or undefined
            if (blockedSpell === null || blockedSpell === undefined) {
                throw new RequiredError('blockedSpell','Required parameter blockedSpell was null or undefined when calling updateBlockedSpell.');
            }
            const localVarPath = `/blocked_spell/{id}`
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
            const nonString = typeof blockedSpell !== 'string';
            const needsSerialization = nonString && configuration && configuration.isJsonMime
                ? configuration.isJsonMime(localVarRequestOptions.headers['Content-Type'])
                : nonString;
            localVarRequestOptions.data =  needsSerialization
                ? JSON.stringify(blockedSpell !== undefined ? blockedSpell : {})
                : (blockedSpell || "");

            return {
                url: localVarUrlObj.pathname + localVarUrlObj.search + localVarUrlObj.hash,
                options: localVarRequestOptions,
            };
        },
    }
};

/**
 * BlockedSpellApi - functional programming interface
 * @export
 */
export const BlockedSpellApiFp = function(configuration?: Configuration) {
    return {
        /**
         * 
         * @summary Creates BlockedSpell
         * @param {ModelsBlockedSpell} blockedSpell BlockedSpell
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async createBlockedSpell(blockedSpell: ModelsBlockedSpell, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsBlockedSpell>>> {
            const localVarAxiosArgs = await BlockedSpellApiAxiosParamCreator(configuration).createBlockedSpell(blockedSpell, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        /**
         * 
         * @summary Deletes BlockedSpell
         * @param {number} id Id
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async deleteBlockedSpell(id: number, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<string>> {
            const localVarAxiosArgs = await BlockedSpellApiAxiosParamCreator(configuration).deleteBlockedSpell(id, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        /**
         * 
         * @summary Gets BlockedSpell
         * @param {number} id Id
         * @param {string} [includes] Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names 
         * @param {string} [select] Column names [.] separated to fetch specific fields in response
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async getBlockedSpell(id: number, includes?: string, select?: string, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsBlockedSpell>>> {
            const localVarAxiosArgs = await BlockedSpellApiAxiosParamCreator(configuration).getBlockedSpell(id, includes, select, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        /**
         * 
         * @summary Gets BlockedSpells in bulk
         * @param {CrudcontrollersBulkFetchByIdsGetRequest} body body
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async getBlockedSpellsBulk(body: CrudcontrollersBulkFetchByIdsGetRequest, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsBlockedSpell>>> {
            const localVarAxiosArgs = await BlockedSpellApiAxiosParamCreator(configuration).getBlockedSpellsBulk(body, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        /**
         * 
         * @summary Lists BlockedSpells
         * @param {string} [includes] Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names 
         * @param {string} [where] Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2
         * @param {string} [whereOr] Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2
         * @param {string} [groupBy] Group by field. Multiple conditions [.] separated Example: field1.field2
         * @param {string} [limit] Rows to limit in response (Default: 10,000)
         * @param {string} [orderBy] Order by [field]
         * @param {string} [orderDirection] Order by field direction
         * @param {string} [select] Column names [.] separated to fetch specific fields in response
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async listBlockedSpells(includes?: string, where?: string, whereOr?: string, groupBy?: string, limit?: string, orderBy?: string, orderDirection?: string, select?: string, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsBlockedSpell>>> {
            const localVarAxiosArgs = await BlockedSpellApiAxiosParamCreator(configuration).listBlockedSpells(includes, where, whereOr, groupBy, limit, orderBy, orderDirection, select, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        /**
         * 
         * @summary Updates BlockedSpell
         * @param {number} id Id
         * @param {ModelsBlockedSpell} blockedSpell BlockedSpell
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async updateBlockedSpell(id: number, blockedSpell: ModelsBlockedSpell, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsBlockedSpell>>> {
            const localVarAxiosArgs = await BlockedSpellApiAxiosParamCreator(configuration).updateBlockedSpell(id, blockedSpell, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
    }
};

/**
 * BlockedSpellApi - factory interface
 * @export
 */
export const BlockedSpellApiFactory = function (configuration?: Configuration, basePath?: string, axios?: AxiosInstance) {
    return {
        /**
         * 
         * @summary Creates BlockedSpell
         * @param {ModelsBlockedSpell} blockedSpell BlockedSpell
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        createBlockedSpell(blockedSpell: ModelsBlockedSpell, options?: any): AxiosPromise<Array<ModelsBlockedSpell>> {
            return BlockedSpellApiFp(configuration).createBlockedSpell(blockedSpell, options).then((request) => request(axios, basePath));
        },
        /**
         * 
         * @summary Deletes BlockedSpell
         * @param {number} id Id
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        deleteBlockedSpell(id: number, options?: any): AxiosPromise<string> {
            return BlockedSpellApiFp(configuration).deleteBlockedSpell(id, options).then((request) => request(axios, basePath));
        },
        /**
         * 
         * @summary Gets BlockedSpell
         * @param {number} id Id
         * @param {string} [includes] Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names 
         * @param {string} [select] Column names [.] separated to fetch specific fields in response
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        getBlockedSpell(id: number, includes?: string, select?: string, options?: any): AxiosPromise<Array<ModelsBlockedSpell>> {
            return BlockedSpellApiFp(configuration).getBlockedSpell(id, includes, select, options).then((request) => request(axios, basePath));
        },
        /**
         * 
         * @summary Gets BlockedSpells in bulk
         * @param {CrudcontrollersBulkFetchByIdsGetRequest} body body
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        getBlockedSpellsBulk(body: CrudcontrollersBulkFetchByIdsGetRequest, options?: any): AxiosPromise<Array<ModelsBlockedSpell>> {
            return BlockedSpellApiFp(configuration).getBlockedSpellsBulk(body, options).then((request) => request(axios, basePath));
        },
        /**
         * 
         * @summary Lists BlockedSpells
         * @param {string} [includes] Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names 
         * @param {string} [where] Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2
         * @param {string} [whereOr] Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2
         * @param {string} [groupBy] Group by field. Multiple conditions [.] separated Example: field1.field2
         * @param {string} [limit] Rows to limit in response (Default: 10,000)
         * @param {string} [orderBy] Order by [field]
         * @param {string} [orderDirection] Order by field direction
         * @param {string} [select] Column names [.] separated to fetch specific fields in response
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        listBlockedSpells(includes?: string, where?: string, whereOr?: string, groupBy?: string, limit?: string, orderBy?: string, orderDirection?: string, select?: string, options?: any): AxiosPromise<Array<ModelsBlockedSpell>> {
            return BlockedSpellApiFp(configuration).listBlockedSpells(includes, where, whereOr, groupBy, limit, orderBy, orderDirection, select, options).then((request) => request(axios, basePath));
        },
        /**
         * 
         * @summary Updates BlockedSpell
         * @param {number} id Id
         * @param {ModelsBlockedSpell} blockedSpell BlockedSpell
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        updateBlockedSpell(id: number, blockedSpell: ModelsBlockedSpell, options?: any): AxiosPromise<Array<ModelsBlockedSpell>> {
            return BlockedSpellApiFp(configuration).updateBlockedSpell(id, blockedSpell, options).then((request) => request(axios, basePath));
        },
    };
};

/**
 * Request parameters for createBlockedSpell operation in BlockedSpellApi.
 * @export
 * @interface BlockedSpellApiCreateBlockedSpellRequest
 */
export interface BlockedSpellApiCreateBlockedSpellRequest {
    /**
     * BlockedSpell
     * @type {ModelsBlockedSpell}
     * @memberof BlockedSpellApiCreateBlockedSpell
     */
    readonly blockedSpell: ModelsBlockedSpell
}

/**
 * Request parameters for deleteBlockedSpell operation in BlockedSpellApi.
 * @export
 * @interface BlockedSpellApiDeleteBlockedSpellRequest
 */
export interface BlockedSpellApiDeleteBlockedSpellRequest {
    /**
     * Id
     * @type {number}
     * @memberof BlockedSpellApiDeleteBlockedSpell
     */
    readonly id: number
}

/**
 * Request parameters for getBlockedSpell operation in BlockedSpellApi.
 * @export
 * @interface BlockedSpellApiGetBlockedSpellRequest
 */
export interface BlockedSpellApiGetBlockedSpellRequest {
    /**
     * Id
     * @type {number}
     * @memberof BlockedSpellApiGetBlockedSpell
     */
    readonly id: number

    /**
     * Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names 
     * @type {string}
     * @memberof BlockedSpellApiGetBlockedSpell
     */
    readonly includes?: string

    /**
     * Column names [.] separated to fetch specific fields in response
     * @type {string}
     * @memberof BlockedSpellApiGetBlockedSpell
     */
    readonly select?: string
}

/**
 * Request parameters for getBlockedSpellsBulk operation in BlockedSpellApi.
 * @export
 * @interface BlockedSpellApiGetBlockedSpellsBulkRequest
 */
export interface BlockedSpellApiGetBlockedSpellsBulkRequest {
    /**
     * body
     * @type {CrudcontrollersBulkFetchByIdsGetRequest}
     * @memberof BlockedSpellApiGetBlockedSpellsBulk
     */
    readonly body: CrudcontrollersBulkFetchByIdsGetRequest
}

/**
 * Request parameters for listBlockedSpells operation in BlockedSpellApi.
 * @export
 * @interface BlockedSpellApiListBlockedSpellsRequest
 */
export interface BlockedSpellApiListBlockedSpellsRequest {
    /**
     * Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names 
     * @type {string}
     * @memberof BlockedSpellApiListBlockedSpells
     */
    readonly includes?: string

    /**
     * Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2
     * @type {string}
     * @memberof BlockedSpellApiListBlockedSpells
     */
    readonly where?: string

    /**
     * Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2
     * @type {string}
     * @memberof BlockedSpellApiListBlockedSpells
     */
    readonly whereOr?: string

    /**
     * Group by field. Multiple conditions [.] separated Example: field1.field2
     * @type {string}
     * @memberof BlockedSpellApiListBlockedSpells
     */
    readonly groupBy?: string

    /**
     * Rows to limit in response (Default: 10,000)
     * @type {string}
     * @memberof BlockedSpellApiListBlockedSpells
     */
    readonly limit?: string

    /**
     * Order by [field]
     * @type {string}
     * @memberof BlockedSpellApiListBlockedSpells
     */
    readonly orderBy?: string

    /**
     * Order by field direction
     * @type {string}
     * @memberof BlockedSpellApiListBlockedSpells
     */
    readonly orderDirection?: string

    /**
     * Column names [.] separated to fetch specific fields in response
     * @type {string}
     * @memberof BlockedSpellApiListBlockedSpells
     */
    readonly select?: string
}

/**
 * Request parameters for updateBlockedSpell operation in BlockedSpellApi.
 * @export
 * @interface BlockedSpellApiUpdateBlockedSpellRequest
 */
export interface BlockedSpellApiUpdateBlockedSpellRequest {
    /**
     * Id
     * @type {number}
     * @memberof BlockedSpellApiUpdateBlockedSpell
     */
    readonly id: number

    /**
     * BlockedSpell
     * @type {ModelsBlockedSpell}
     * @memberof BlockedSpellApiUpdateBlockedSpell
     */
    readonly blockedSpell: ModelsBlockedSpell
}

/**
 * BlockedSpellApi - object-oriented interface
 * @export
 * @class BlockedSpellApi
 * @extends {BaseAPI}
 */
export class BlockedSpellApi extends BaseAPI {
    /**
     * 
     * @summary Creates BlockedSpell
     * @param {BlockedSpellApiCreateBlockedSpellRequest} requestParameters Request parameters.
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof BlockedSpellApi
     */
    public createBlockedSpell(requestParameters: BlockedSpellApiCreateBlockedSpellRequest, options?: any) {
        return BlockedSpellApiFp(this.configuration).createBlockedSpell(requestParameters.blockedSpell, options).then((request) => request(this.axios, this.basePath));
    }

    /**
     * 
     * @summary Deletes BlockedSpell
     * @param {BlockedSpellApiDeleteBlockedSpellRequest} requestParameters Request parameters.
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof BlockedSpellApi
     */
    public deleteBlockedSpell(requestParameters: BlockedSpellApiDeleteBlockedSpellRequest, options?: any) {
        return BlockedSpellApiFp(this.configuration).deleteBlockedSpell(requestParameters.id, options).then((request) => request(this.axios, this.basePath));
    }

    /**
     * 
     * @summary Gets BlockedSpell
     * @param {BlockedSpellApiGetBlockedSpellRequest} requestParameters Request parameters.
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof BlockedSpellApi
     */
    public getBlockedSpell(requestParameters: BlockedSpellApiGetBlockedSpellRequest, options?: any) {
        return BlockedSpellApiFp(this.configuration).getBlockedSpell(requestParameters.id, requestParameters.includes, requestParameters.select, options).then((request) => request(this.axios, this.basePath));
    }

    /**
     * 
     * @summary Gets BlockedSpells in bulk
     * @param {BlockedSpellApiGetBlockedSpellsBulkRequest} requestParameters Request parameters.
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof BlockedSpellApi
     */
    public getBlockedSpellsBulk(requestParameters: BlockedSpellApiGetBlockedSpellsBulkRequest, options?: any) {
        return BlockedSpellApiFp(this.configuration).getBlockedSpellsBulk(requestParameters.body, options).then((request) => request(this.axios, this.basePath));
    }

    /**
     * 
     * @summary Lists BlockedSpells
     * @param {BlockedSpellApiListBlockedSpellsRequest} requestParameters Request parameters.
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof BlockedSpellApi
     */
    public listBlockedSpells(requestParameters: BlockedSpellApiListBlockedSpellsRequest = {}, options?: any) {
        return BlockedSpellApiFp(this.configuration).listBlockedSpells(requestParameters.includes, requestParameters.where, requestParameters.whereOr, requestParameters.groupBy, requestParameters.limit, requestParameters.orderBy, requestParameters.orderDirection, requestParameters.select, options).then((request) => request(this.axios, this.basePath));
    }

    /**
     * 
     * @summary Updates BlockedSpell
     * @param {BlockedSpellApiUpdateBlockedSpellRequest} requestParameters Request parameters.
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof BlockedSpellApi
     */
    public updateBlockedSpell(requestParameters: BlockedSpellApiUpdateBlockedSpellRequest, options?: any) {
        return BlockedSpellApiFp(this.configuration).updateBlockedSpell(requestParameters.id, requestParameters.blockedSpell, options).then((request) => request(this.axios, this.basePath));
    }
}
