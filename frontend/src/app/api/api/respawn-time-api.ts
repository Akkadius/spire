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
import { ModelsRespawnTime } from '../models';
/**
 * RespawnTimeApi - axios parameter creator
 * @export
 */
export const RespawnTimeApiAxiosParamCreator = function (configuration?: Configuration) {
    return {
        /**
         * 
         * @summary Creates RespawnTime
         * @param {ModelsRespawnTime} respawnTime RespawnTime
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        createRespawnTime: async (respawnTime: ModelsRespawnTime, options: any = {}): Promise<RequestArgs> => {
            // verify required parameter 'respawnTime' is not null or undefined
            if (respawnTime === null || respawnTime === undefined) {
                throw new RequiredError('respawnTime','Required parameter respawnTime was null or undefined when calling createRespawnTime.');
            }
            const localVarPath = `/respawn_time`;
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
            const nonString = typeof respawnTime !== 'string';
            const needsSerialization = nonString && configuration && configuration.isJsonMime
                ? configuration.isJsonMime(localVarRequestOptions.headers['Content-Type'])
                : nonString;
            localVarRequestOptions.data =  needsSerialization
                ? JSON.stringify(respawnTime !== undefined ? respawnTime : {})
                : (respawnTime || "");

            return {
                url: localVarUrlObj.pathname + localVarUrlObj.search + localVarUrlObj.hash,
                options: localVarRequestOptions,
            };
        },
        /**
         * 
         * @summary Deletes RespawnTime
         * @param {number} id Id
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        deleteRespawnTime: async (id: number, options: any = {}): Promise<RequestArgs> => {
            // verify required parameter 'id' is not null or undefined
            if (id === null || id === undefined) {
                throw new RequiredError('id','Required parameter id was null or undefined when calling deleteRespawnTime.');
            }
            const localVarPath = `/respawn_time/{id}`
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
         * @summary Gets RespawnTime
         * @param {number} id Id
         * @param {string} [includes] Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names 
         * @param {string} [select] Column names [.] separated to fetch specific fields in response
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        getRespawnTime: async (id: number, includes?: string, select?: string, options: any = {}): Promise<RequestArgs> => {
            // verify required parameter 'id' is not null or undefined
            if (id === null || id === undefined) {
                throw new RequiredError('id','Required parameter id was null or undefined when calling getRespawnTime.');
            }
            const localVarPath = `/respawn_time/{id}`
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
         * @summary Gets RespawnTimes in bulk
         * @param {CrudcontrollersBulkFetchByIdsGetRequest} body body
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        getRespawnTimesBulk: async (body: CrudcontrollersBulkFetchByIdsGetRequest, options: any = {}): Promise<RequestArgs> => {
            // verify required parameter 'body' is not null or undefined
            if (body === null || body === undefined) {
                throw new RequiredError('body','Required parameter body was null or undefined when calling getRespawnTimesBulk.');
            }
            const localVarPath = `/respawn_times/bulk`;
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
         * @summary Lists RespawnTimes
         * @param {string} [includes] Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names 
         * @param {string} [where] Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2
         * @param {string} [whereOr] Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2
         * @param {string} [limit] Rows to limit in response (Default: 10,000)
         * @param {string} [orderBy] Order by [field]
         * @param {string} [orderDirection] Order by field direction
         * @param {string} [select] Column names [.] separated to fetch specific fields in response
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        listRespawnTimes: async (includes?: string, where?: string, whereOr?: string, limit?: string, orderBy?: string, orderDirection?: string, select?: string, options: any = {}): Promise<RequestArgs> => {
            const localVarPath = `/respawn_times`;
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
         * @summary Updates RespawnTime
         * @param {number} id Id
         * @param {ModelsRespawnTime} respawnTime RespawnTime
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        updateRespawnTime: async (id: number, respawnTime: ModelsRespawnTime, options: any = {}): Promise<RequestArgs> => {
            // verify required parameter 'id' is not null or undefined
            if (id === null || id === undefined) {
                throw new RequiredError('id','Required parameter id was null or undefined when calling updateRespawnTime.');
            }
            // verify required parameter 'respawnTime' is not null or undefined
            if (respawnTime === null || respawnTime === undefined) {
                throw new RequiredError('respawnTime','Required parameter respawnTime was null or undefined when calling updateRespawnTime.');
            }
            const localVarPath = `/respawn_time/{id}`
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
            const nonString = typeof respawnTime !== 'string';
            const needsSerialization = nonString && configuration && configuration.isJsonMime
                ? configuration.isJsonMime(localVarRequestOptions.headers['Content-Type'])
                : nonString;
            localVarRequestOptions.data =  needsSerialization
                ? JSON.stringify(respawnTime !== undefined ? respawnTime : {})
                : (respawnTime || "");

            return {
                url: localVarUrlObj.pathname + localVarUrlObj.search + localVarUrlObj.hash,
                options: localVarRequestOptions,
            };
        },
    }
};

/**
 * RespawnTimeApi - functional programming interface
 * @export
 */
export const RespawnTimeApiFp = function(configuration?: Configuration) {
    return {
        /**
         * 
         * @summary Creates RespawnTime
         * @param {ModelsRespawnTime} respawnTime RespawnTime
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async createRespawnTime(respawnTime: ModelsRespawnTime, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsRespawnTime>>> {
            const localVarAxiosArgs = await RespawnTimeApiAxiosParamCreator(configuration).createRespawnTime(respawnTime, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        /**
         * 
         * @summary Deletes RespawnTime
         * @param {number} id Id
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async deleteRespawnTime(id: number, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<string>> {
            const localVarAxiosArgs = await RespawnTimeApiAxiosParamCreator(configuration).deleteRespawnTime(id, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        /**
         * 
         * @summary Gets RespawnTime
         * @param {number} id Id
         * @param {string} [includes] Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names 
         * @param {string} [select] Column names [.] separated to fetch specific fields in response
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async getRespawnTime(id: number, includes?: string, select?: string, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsRespawnTime>>> {
            const localVarAxiosArgs = await RespawnTimeApiAxiosParamCreator(configuration).getRespawnTime(id, includes, select, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        /**
         * 
         * @summary Gets RespawnTimes in bulk
         * @param {CrudcontrollersBulkFetchByIdsGetRequest} body body
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async getRespawnTimesBulk(body: CrudcontrollersBulkFetchByIdsGetRequest, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsRespawnTime>>> {
            const localVarAxiosArgs = await RespawnTimeApiAxiosParamCreator(configuration).getRespawnTimesBulk(body, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        /**
         * 
         * @summary Lists RespawnTimes
         * @param {string} [includes] Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names 
         * @param {string} [where] Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2
         * @param {string} [whereOr] Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2
         * @param {string} [limit] Rows to limit in response (Default: 10,000)
         * @param {string} [orderBy] Order by [field]
         * @param {string} [orderDirection] Order by field direction
         * @param {string} [select] Column names [.] separated to fetch specific fields in response
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async listRespawnTimes(includes?: string, where?: string, whereOr?: string, limit?: string, orderBy?: string, orderDirection?: string, select?: string, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsRespawnTime>>> {
            const localVarAxiosArgs = await RespawnTimeApiAxiosParamCreator(configuration).listRespawnTimes(includes, where, whereOr, limit, orderBy, orderDirection, select, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        /**
         * 
         * @summary Updates RespawnTime
         * @param {number} id Id
         * @param {ModelsRespawnTime} respawnTime RespawnTime
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async updateRespawnTime(id: number, respawnTime: ModelsRespawnTime, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsRespawnTime>>> {
            const localVarAxiosArgs = await RespawnTimeApiAxiosParamCreator(configuration).updateRespawnTime(id, respawnTime, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
    }
};

/**
 * RespawnTimeApi - factory interface
 * @export
 */
export const RespawnTimeApiFactory = function (configuration?: Configuration, basePath?: string, axios?: AxiosInstance) {
    return {
        /**
         * 
         * @summary Creates RespawnTime
         * @param {ModelsRespawnTime} respawnTime RespawnTime
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        createRespawnTime(respawnTime: ModelsRespawnTime, options?: any): AxiosPromise<Array<ModelsRespawnTime>> {
            return RespawnTimeApiFp(configuration).createRespawnTime(respawnTime, options).then((request) => request(axios, basePath));
        },
        /**
         * 
         * @summary Deletes RespawnTime
         * @param {number} id Id
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        deleteRespawnTime(id: number, options?: any): AxiosPromise<string> {
            return RespawnTimeApiFp(configuration).deleteRespawnTime(id, options).then((request) => request(axios, basePath));
        },
        /**
         * 
         * @summary Gets RespawnTime
         * @param {number} id Id
         * @param {string} [includes] Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names 
         * @param {string} [select] Column names [.] separated to fetch specific fields in response
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        getRespawnTime(id: number, includes?: string, select?: string, options?: any): AxiosPromise<Array<ModelsRespawnTime>> {
            return RespawnTimeApiFp(configuration).getRespawnTime(id, includes, select, options).then((request) => request(axios, basePath));
        },
        /**
         * 
         * @summary Gets RespawnTimes in bulk
         * @param {CrudcontrollersBulkFetchByIdsGetRequest} body body
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        getRespawnTimesBulk(body: CrudcontrollersBulkFetchByIdsGetRequest, options?: any): AxiosPromise<Array<ModelsRespawnTime>> {
            return RespawnTimeApiFp(configuration).getRespawnTimesBulk(body, options).then((request) => request(axios, basePath));
        },
        /**
         * 
         * @summary Lists RespawnTimes
         * @param {string} [includes] Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names 
         * @param {string} [where] Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2
         * @param {string} [whereOr] Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2
         * @param {string} [limit] Rows to limit in response (Default: 10,000)
         * @param {string} [orderBy] Order by [field]
         * @param {string} [orderDirection] Order by field direction
         * @param {string} [select] Column names [.] separated to fetch specific fields in response
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        listRespawnTimes(includes?: string, where?: string, whereOr?: string, limit?: string, orderBy?: string, orderDirection?: string, select?: string, options?: any): AxiosPromise<Array<ModelsRespawnTime>> {
            return RespawnTimeApiFp(configuration).listRespawnTimes(includes, where, whereOr, limit, orderBy, orderDirection, select, options).then((request) => request(axios, basePath));
        },
        /**
         * 
         * @summary Updates RespawnTime
         * @param {number} id Id
         * @param {ModelsRespawnTime} respawnTime RespawnTime
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        updateRespawnTime(id: number, respawnTime: ModelsRespawnTime, options?: any): AxiosPromise<Array<ModelsRespawnTime>> {
            return RespawnTimeApiFp(configuration).updateRespawnTime(id, respawnTime, options).then((request) => request(axios, basePath));
        },
    };
};

/**
 * Request parameters for createRespawnTime operation in RespawnTimeApi.
 * @export
 * @interface RespawnTimeApiCreateRespawnTimeRequest
 */
export interface RespawnTimeApiCreateRespawnTimeRequest {
    /**
     * RespawnTime
     * @type {ModelsRespawnTime}
     * @memberof RespawnTimeApiCreateRespawnTime
     */
    readonly respawnTime: ModelsRespawnTime
}

/**
 * Request parameters for deleteRespawnTime operation in RespawnTimeApi.
 * @export
 * @interface RespawnTimeApiDeleteRespawnTimeRequest
 */
export interface RespawnTimeApiDeleteRespawnTimeRequest {
    /**
     * Id
     * @type {number}
     * @memberof RespawnTimeApiDeleteRespawnTime
     */
    readonly id: number
}

/**
 * Request parameters for getRespawnTime operation in RespawnTimeApi.
 * @export
 * @interface RespawnTimeApiGetRespawnTimeRequest
 */
export interface RespawnTimeApiGetRespawnTimeRequest {
    /**
     * Id
     * @type {number}
     * @memberof RespawnTimeApiGetRespawnTime
     */
    readonly id: number

    /**
     * Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names 
     * @type {string}
     * @memberof RespawnTimeApiGetRespawnTime
     */
    readonly includes?: string

    /**
     * Column names [.] separated to fetch specific fields in response
     * @type {string}
     * @memberof RespawnTimeApiGetRespawnTime
     */
    readonly select?: string
}

/**
 * Request parameters for getRespawnTimesBulk operation in RespawnTimeApi.
 * @export
 * @interface RespawnTimeApiGetRespawnTimesBulkRequest
 */
export interface RespawnTimeApiGetRespawnTimesBulkRequest {
    /**
     * body
     * @type {CrudcontrollersBulkFetchByIdsGetRequest}
     * @memberof RespawnTimeApiGetRespawnTimesBulk
     */
    readonly body: CrudcontrollersBulkFetchByIdsGetRequest
}

/**
 * Request parameters for listRespawnTimes operation in RespawnTimeApi.
 * @export
 * @interface RespawnTimeApiListRespawnTimesRequest
 */
export interface RespawnTimeApiListRespawnTimesRequest {
    /**
     * Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names 
     * @type {string}
     * @memberof RespawnTimeApiListRespawnTimes
     */
    readonly includes?: string

    /**
     * Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2
     * @type {string}
     * @memberof RespawnTimeApiListRespawnTimes
     */
    readonly where?: string

    /**
     * Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2
     * @type {string}
     * @memberof RespawnTimeApiListRespawnTimes
     */
    readonly whereOr?: string

    /**
     * Rows to limit in response (Default: 10,000)
     * @type {string}
     * @memberof RespawnTimeApiListRespawnTimes
     */
    readonly limit?: string

    /**
     * Order by [field]
     * @type {string}
     * @memberof RespawnTimeApiListRespawnTimes
     */
    readonly orderBy?: string

    /**
     * Order by field direction
     * @type {string}
     * @memberof RespawnTimeApiListRespawnTimes
     */
    readonly orderDirection?: string

    /**
     * Column names [.] separated to fetch specific fields in response
     * @type {string}
     * @memberof RespawnTimeApiListRespawnTimes
     */
    readonly select?: string
}

/**
 * Request parameters for updateRespawnTime operation in RespawnTimeApi.
 * @export
 * @interface RespawnTimeApiUpdateRespawnTimeRequest
 */
export interface RespawnTimeApiUpdateRespawnTimeRequest {
    /**
     * Id
     * @type {number}
     * @memberof RespawnTimeApiUpdateRespawnTime
     */
    readonly id: number

    /**
     * RespawnTime
     * @type {ModelsRespawnTime}
     * @memberof RespawnTimeApiUpdateRespawnTime
     */
    readonly respawnTime: ModelsRespawnTime
}

/**
 * RespawnTimeApi - object-oriented interface
 * @export
 * @class RespawnTimeApi
 * @extends {BaseAPI}
 */
export class RespawnTimeApi extends BaseAPI {
    /**
     * 
     * @summary Creates RespawnTime
     * @param {RespawnTimeApiCreateRespawnTimeRequest} requestParameters Request parameters.
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof RespawnTimeApi
     */
    public createRespawnTime(requestParameters: RespawnTimeApiCreateRespawnTimeRequest, options?: any) {
        return RespawnTimeApiFp(this.configuration).createRespawnTime(requestParameters.respawnTime, options).then((request) => request(this.axios, this.basePath));
    }

    /**
     * 
     * @summary Deletes RespawnTime
     * @param {RespawnTimeApiDeleteRespawnTimeRequest} requestParameters Request parameters.
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof RespawnTimeApi
     */
    public deleteRespawnTime(requestParameters: RespawnTimeApiDeleteRespawnTimeRequest, options?: any) {
        return RespawnTimeApiFp(this.configuration).deleteRespawnTime(requestParameters.id, options).then((request) => request(this.axios, this.basePath));
    }

    /**
     * 
     * @summary Gets RespawnTime
     * @param {RespawnTimeApiGetRespawnTimeRequest} requestParameters Request parameters.
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof RespawnTimeApi
     */
    public getRespawnTime(requestParameters: RespawnTimeApiGetRespawnTimeRequest, options?: any) {
        return RespawnTimeApiFp(this.configuration).getRespawnTime(requestParameters.id, requestParameters.includes, requestParameters.select, options).then((request) => request(this.axios, this.basePath));
    }

    /**
     * 
     * @summary Gets RespawnTimes in bulk
     * @param {RespawnTimeApiGetRespawnTimesBulkRequest} requestParameters Request parameters.
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof RespawnTimeApi
     */
    public getRespawnTimesBulk(requestParameters: RespawnTimeApiGetRespawnTimesBulkRequest, options?: any) {
        return RespawnTimeApiFp(this.configuration).getRespawnTimesBulk(requestParameters.body, options).then((request) => request(this.axios, this.basePath));
    }

    /**
     * 
     * @summary Lists RespawnTimes
     * @param {RespawnTimeApiListRespawnTimesRequest} requestParameters Request parameters.
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof RespawnTimeApi
     */
    public listRespawnTimes(requestParameters: RespawnTimeApiListRespawnTimesRequest = {}, options?: any) {
        return RespawnTimeApiFp(this.configuration).listRespawnTimes(requestParameters.includes, requestParameters.where, requestParameters.whereOr, requestParameters.limit, requestParameters.orderBy, requestParameters.orderDirection, requestParameters.select, options).then((request) => request(this.axios, this.basePath));
    }

    /**
     * 
     * @summary Updates RespawnTime
     * @param {RespawnTimeApiUpdateRespawnTimeRequest} requestParameters Request parameters.
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof RespawnTimeApi
     */
    public updateRespawnTime(requestParameters: RespawnTimeApiUpdateRespawnTimeRequest, options?: any) {
        return RespawnTimeApiFp(this.configuration).updateRespawnTime(requestParameters.id, requestParameters.respawnTime, options).then((request) => request(this.axios, this.basePath));
    }
}
