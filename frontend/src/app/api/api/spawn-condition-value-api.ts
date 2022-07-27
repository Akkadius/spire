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
import { ModelsSpawnConditionValue } from '../models';
/**
 * SpawnConditionValueApi - axios parameter creator
 * @export
 */
export const SpawnConditionValueApiAxiosParamCreator = function (configuration?: Configuration) {
    return {
        /**
         * 
         * @summary Creates SpawnConditionValue
         * @param {ModelsSpawnConditionValue} spawnConditionValue SpawnConditionValue
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        createSpawnConditionValue: async (spawnConditionValue: ModelsSpawnConditionValue, options: any = {}): Promise<RequestArgs> => {
            // verify required parameter 'spawnConditionValue' is not null or undefined
            if (spawnConditionValue === null || spawnConditionValue === undefined) {
                throw new RequiredError('spawnConditionValue','Required parameter spawnConditionValue was null or undefined when calling createSpawnConditionValue.');
            }
            const localVarPath = `/spawn_condition_value`;
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
            const nonString = typeof spawnConditionValue !== 'string';
            const needsSerialization = nonString && configuration && configuration.isJsonMime
                ? configuration.isJsonMime(localVarRequestOptions.headers['Content-Type'])
                : nonString;
            localVarRequestOptions.data =  needsSerialization
                ? JSON.stringify(spawnConditionValue !== undefined ? spawnConditionValue : {})
                : (spawnConditionValue || "");

            return {
                url: localVarUrlObj.pathname + localVarUrlObj.search + localVarUrlObj.hash,
                options: localVarRequestOptions,
            };
        },
        /**
         * 
         * @summary Deletes SpawnConditionValue
         * @param {number} id id
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        deleteSpawnConditionValue: async (id: number, options: any = {}): Promise<RequestArgs> => {
            // verify required parameter 'id' is not null or undefined
            if (id === null || id === undefined) {
                throw new RequiredError('id','Required parameter id was null or undefined when calling deleteSpawnConditionValue.');
            }
            const localVarPath = `/spawn_condition_value/{id}`
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
         * @summary Gets SpawnConditionValue
         * @param {number} id Id
         * @param {string} [includes] Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names 
         * @param {string} [select] Column names [.] separated to fetch specific fields in response
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        getSpawnConditionValue: async (id: number, includes?: string, select?: string, options: any = {}): Promise<RequestArgs> => {
            // verify required parameter 'id' is not null or undefined
            if (id === null || id === undefined) {
                throw new RequiredError('id','Required parameter id was null or undefined when calling getSpawnConditionValue.');
            }
            const localVarPath = `/spawn_condition_value/{id}`
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
         * @summary Gets SpawnConditionValues in bulk
         * @param {CrudcontrollersBulkFetchByIdsGetRequest} body body
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        getSpawnConditionValuesBulk: async (body: CrudcontrollersBulkFetchByIdsGetRequest, options: any = {}): Promise<RequestArgs> => {
            // verify required parameter 'body' is not null or undefined
            if (body === null || body === undefined) {
                throw new RequiredError('body','Required parameter body was null or undefined when calling getSpawnConditionValuesBulk.');
            }
            const localVarPath = `/spawn_condition_values/bulk`;
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
         * @summary Lists SpawnConditionValues
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
        listSpawnConditionValues: async (includes?: string, where?: string, whereOr?: string, groupBy?: string, limit?: string, page?: number, orderBy?: string, orderDirection?: string, select?: string, options: any = {}): Promise<RequestArgs> => {
            const localVarPath = `/spawn_condition_values`;
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
         * @summary Updates SpawnConditionValue
         * @param {number} id Id
         * @param {ModelsSpawnConditionValue} spawnConditionValue SpawnConditionValue
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        updateSpawnConditionValue: async (id: number, spawnConditionValue: ModelsSpawnConditionValue, options: any = {}): Promise<RequestArgs> => {
            // verify required parameter 'id' is not null or undefined
            if (id === null || id === undefined) {
                throw new RequiredError('id','Required parameter id was null or undefined when calling updateSpawnConditionValue.');
            }
            // verify required parameter 'spawnConditionValue' is not null or undefined
            if (spawnConditionValue === null || spawnConditionValue === undefined) {
                throw new RequiredError('spawnConditionValue','Required parameter spawnConditionValue was null or undefined when calling updateSpawnConditionValue.');
            }
            const localVarPath = `/spawn_condition_value/{id}`
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
            const nonString = typeof spawnConditionValue !== 'string';
            const needsSerialization = nonString && configuration && configuration.isJsonMime
                ? configuration.isJsonMime(localVarRequestOptions.headers['Content-Type'])
                : nonString;
            localVarRequestOptions.data =  needsSerialization
                ? JSON.stringify(spawnConditionValue !== undefined ? spawnConditionValue : {})
                : (spawnConditionValue || "");

            return {
                url: localVarUrlObj.pathname + localVarUrlObj.search + localVarUrlObj.hash,
                options: localVarRequestOptions,
            };
        },
    }
};

/**
 * SpawnConditionValueApi - functional programming interface
 * @export
 */
export const SpawnConditionValueApiFp = function(configuration?: Configuration) {
    return {
        /**
         * 
         * @summary Creates SpawnConditionValue
         * @param {ModelsSpawnConditionValue} spawnConditionValue SpawnConditionValue
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async createSpawnConditionValue(spawnConditionValue: ModelsSpawnConditionValue, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsSpawnConditionValue>>> {
            const localVarAxiosArgs = await SpawnConditionValueApiAxiosParamCreator(configuration).createSpawnConditionValue(spawnConditionValue, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        /**
         * 
         * @summary Deletes SpawnConditionValue
         * @param {number} id id
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async deleteSpawnConditionValue(id: number, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<string>> {
            const localVarAxiosArgs = await SpawnConditionValueApiAxiosParamCreator(configuration).deleteSpawnConditionValue(id, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        /**
         * 
         * @summary Gets SpawnConditionValue
         * @param {number} id Id
         * @param {string} [includes] Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names 
         * @param {string} [select] Column names [.] separated to fetch specific fields in response
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async getSpawnConditionValue(id: number, includes?: string, select?: string, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsSpawnConditionValue>>> {
            const localVarAxiosArgs = await SpawnConditionValueApiAxiosParamCreator(configuration).getSpawnConditionValue(id, includes, select, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        /**
         * 
         * @summary Gets SpawnConditionValues in bulk
         * @param {CrudcontrollersBulkFetchByIdsGetRequest} body body
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async getSpawnConditionValuesBulk(body: CrudcontrollersBulkFetchByIdsGetRequest, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsSpawnConditionValue>>> {
            const localVarAxiosArgs = await SpawnConditionValueApiAxiosParamCreator(configuration).getSpawnConditionValuesBulk(body, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        /**
         * 
         * @summary Lists SpawnConditionValues
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
        async listSpawnConditionValues(includes?: string, where?: string, whereOr?: string, groupBy?: string, limit?: string, page?: number, orderBy?: string, orderDirection?: string, select?: string, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsSpawnConditionValue>>> {
            const localVarAxiosArgs = await SpawnConditionValueApiAxiosParamCreator(configuration).listSpawnConditionValues(includes, where, whereOr, groupBy, limit, page, orderBy, orderDirection, select, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        /**
         * 
         * @summary Updates SpawnConditionValue
         * @param {number} id Id
         * @param {ModelsSpawnConditionValue} spawnConditionValue SpawnConditionValue
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async updateSpawnConditionValue(id: number, spawnConditionValue: ModelsSpawnConditionValue, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsSpawnConditionValue>>> {
            const localVarAxiosArgs = await SpawnConditionValueApiAxiosParamCreator(configuration).updateSpawnConditionValue(id, spawnConditionValue, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
    }
};

/**
 * SpawnConditionValueApi - factory interface
 * @export
 */
export const SpawnConditionValueApiFactory = function (configuration?: Configuration, basePath?: string, axios?: AxiosInstance) {
    return {
        /**
         * 
         * @summary Creates SpawnConditionValue
         * @param {ModelsSpawnConditionValue} spawnConditionValue SpawnConditionValue
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        createSpawnConditionValue(spawnConditionValue: ModelsSpawnConditionValue, options?: any): AxiosPromise<Array<ModelsSpawnConditionValue>> {
            return SpawnConditionValueApiFp(configuration).createSpawnConditionValue(spawnConditionValue, options).then((request) => request(axios, basePath));
        },
        /**
         * 
         * @summary Deletes SpawnConditionValue
         * @param {number} id id
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        deleteSpawnConditionValue(id: number, options?: any): AxiosPromise<string> {
            return SpawnConditionValueApiFp(configuration).deleteSpawnConditionValue(id, options).then((request) => request(axios, basePath));
        },
        /**
         * 
         * @summary Gets SpawnConditionValue
         * @param {number} id Id
         * @param {string} [includes] Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names 
         * @param {string} [select] Column names [.] separated to fetch specific fields in response
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        getSpawnConditionValue(id: number, includes?: string, select?: string, options?: any): AxiosPromise<Array<ModelsSpawnConditionValue>> {
            return SpawnConditionValueApiFp(configuration).getSpawnConditionValue(id, includes, select, options).then((request) => request(axios, basePath));
        },
        /**
         * 
         * @summary Gets SpawnConditionValues in bulk
         * @param {CrudcontrollersBulkFetchByIdsGetRequest} body body
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        getSpawnConditionValuesBulk(body: CrudcontrollersBulkFetchByIdsGetRequest, options?: any): AxiosPromise<Array<ModelsSpawnConditionValue>> {
            return SpawnConditionValueApiFp(configuration).getSpawnConditionValuesBulk(body, options).then((request) => request(axios, basePath));
        },
        /**
         * 
         * @summary Lists SpawnConditionValues
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
        listSpawnConditionValues(includes?: string, where?: string, whereOr?: string, groupBy?: string, limit?: string, page?: number, orderBy?: string, orderDirection?: string, select?: string, options?: any): AxiosPromise<Array<ModelsSpawnConditionValue>> {
            return SpawnConditionValueApiFp(configuration).listSpawnConditionValues(includes, where, whereOr, groupBy, limit, page, orderBy, orderDirection, select, options).then((request) => request(axios, basePath));
        },
        /**
         * 
         * @summary Updates SpawnConditionValue
         * @param {number} id Id
         * @param {ModelsSpawnConditionValue} spawnConditionValue SpawnConditionValue
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        updateSpawnConditionValue(id: number, spawnConditionValue: ModelsSpawnConditionValue, options?: any): AxiosPromise<Array<ModelsSpawnConditionValue>> {
            return SpawnConditionValueApiFp(configuration).updateSpawnConditionValue(id, spawnConditionValue, options).then((request) => request(axios, basePath));
        },
    };
};

/**
 * Request parameters for createSpawnConditionValue operation in SpawnConditionValueApi.
 * @export
 * @interface SpawnConditionValueApiCreateSpawnConditionValueRequest
 */
export interface SpawnConditionValueApiCreateSpawnConditionValueRequest {
    /**
     * SpawnConditionValue
     * @type {ModelsSpawnConditionValue}
     * @memberof SpawnConditionValueApiCreateSpawnConditionValue
     */
    readonly spawnConditionValue: ModelsSpawnConditionValue
}

/**
 * Request parameters for deleteSpawnConditionValue operation in SpawnConditionValueApi.
 * @export
 * @interface SpawnConditionValueApiDeleteSpawnConditionValueRequest
 */
export interface SpawnConditionValueApiDeleteSpawnConditionValueRequest {
    /**
     * id
     * @type {number}
     * @memberof SpawnConditionValueApiDeleteSpawnConditionValue
     */
    readonly id: number
}

/**
 * Request parameters for getSpawnConditionValue operation in SpawnConditionValueApi.
 * @export
 * @interface SpawnConditionValueApiGetSpawnConditionValueRequest
 */
export interface SpawnConditionValueApiGetSpawnConditionValueRequest {
    /**
     * Id
     * @type {number}
     * @memberof SpawnConditionValueApiGetSpawnConditionValue
     */
    readonly id: number

    /**
     * Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names 
     * @type {string}
     * @memberof SpawnConditionValueApiGetSpawnConditionValue
     */
    readonly includes?: string

    /**
     * Column names [.] separated to fetch specific fields in response
     * @type {string}
     * @memberof SpawnConditionValueApiGetSpawnConditionValue
     */
    readonly select?: string
}

/**
 * Request parameters for getSpawnConditionValuesBulk operation in SpawnConditionValueApi.
 * @export
 * @interface SpawnConditionValueApiGetSpawnConditionValuesBulkRequest
 */
export interface SpawnConditionValueApiGetSpawnConditionValuesBulkRequest {
    /**
     * body
     * @type {CrudcontrollersBulkFetchByIdsGetRequest}
     * @memberof SpawnConditionValueApiGetSpawnConditionValuesBulk
     */
    readonly body: CrudcontrollersBulkFetchByIdsGetRequest
}

/**
 * Request parameters for listSpawnConditionValues operation in SpawnConditionValueApi.
 * @export
 * @interface SpawnConditionValueApiListSpawnConditionValuesRequest
 */
export interface SpawnConditionValueApiListSpawnConditionValuesRequest {
    /**
     * Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names 
     * @type {string}
     * @memberof SpawnConditionValueApiListSpawnConditionValues
     */
    readonly includes?: string

    /**
     * Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2
     * @type {string}
     * @memberof SpawnConditionValueApiListSpawnConditionValues
     */
    readonly where?: string

    /**
     * Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2
     * @type {string}
     * @memberof SpawnConditionValueApiListSpawnConditionValues
     */
    readonly whereOr?: string

    /**
     * Group by field. Multiple conditions [.] separated Example: field1.field2
     * @type {string}
     * @memberof SpawnConditionValueApiListSpawnConditionValues
     */
    readonly groupBy?: string

    /**
     * Rows to limit in response (Default: 10,000)
     * @type {string}
     * @memberof SpawnConditionValueApiListSpawnConditionValues
     */
    readonly limit?: string

    /**
     * Pagination page
     * @type {number}
     * @memberof SpawnConditionValueApiListSpawnConditionValues
     */
    readonly page?: number

    /**
     * Order by [field]
     * @type {string}
     * @memberof SpawnConditionValueApiListSpawnConditionValues
     */
    readonly orderBy?: string

    /**
     * Order by field direction
     * @type {string}
     * @memberof SpawnConditionValueApiListSpawnConditionValues
     */
    readonly orderDirection?: string

    /**
     * Column names [.] separated to fetch specific fields in response
     * @type {string}
     * @memberof SpawnConditionValueApiListSpawnConditionValues
     */
    readonly select?: string
}

/**
 * Request parameters for updateSpawnConditionValue operation in SpawnConditionValueApi.
 * @export
 * @interface SpawnConditionValueApiUpdateSpawnConditionValueRequest
 */
export interface SpawnConditionValueApiUpdateSpawnConditionValueRequest {
    /**
     * Id
     * @type {number}
     * @memberof SpawnConditionValueApiUpdateSpawnConditionValue
     */
    readonly id: number

    /**
     * SpawnConditionValue
     * @type {ModelsSpawnConditionValue}
     * @memberof SpawnConditionValueApiUpdateSpawnConditionValue
     */
    readonly spawnConditionValue: ModelsSpawnConditionValue
}

/**
 * SpawnConditionValueApi - object-oriented interface
 * @export
 * @class SpawnConditionValueApi
 * @extends {BaseAPI}
 */
export class SpawnConditionValueApi extends BaseAPI {
    /**
     * 
     * @summary Creates SpawnConditionValue
     * @param {SpawnConditionValueApiCreateSpawnConditionValueRequest} requestParameters Request parameters.
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof SpawnConditionValueApi
     */
    public createSpawnConditionValue(requestParameters: SpawnConditionValueApiCreateSpawnConditionValueRequest, options?: any) {
        return SpawnConditionValueApiFp(this.configuration).createSpawnConditionValue(requestParameters.spawnConditionValue, options).then((request) => request(this.axios, this.basePath));
    }

    /**
     * 
     * @summary Deletes SpawnConditionValue
     * @param {SpawnConditionValueApiDeleteSpawnConditionValueRequest} requestParameters Request parameters.
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof SpawnConditionValueApi
     */
    public deleteSpawnConditionValue(requestParameters: SpawnConditionValueApiDeleteSpawnConditionValueRequest, options?: any) {
        return SpawnConditionValueApiFp(this.configuration).deleteSpawnConditionValue(requestParameters.id, options).then((request) => request(this.axios, this.basePath));
    }

    /**
     * 
     * @summary Gets SpawnConditionValue
     * @param {SpawnConditionValueApiGetSpawnConditionValueRequest} requestParameters Request parameters.
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof SpawnConditionValueApi
     */
    public getSpawnConditionValue(requestParameters: SpawnConditionValueApiGetSpawnConditionValueRequest, options?: any) {
        return SpawnConditionValueApiFp(this.configuration).getSpawnConditionValue(requestParameters.id, requestParameters.includes, requestParameters.select, options).then((request) => request(this.axios, this.basePath));
    }

    /**
     * 
     * @summary Gets SpawnConditionValues in bulk
     * @param {SpawnConditionValueApiGetSpawnConditionValuesBulkRequest} requestParameters Request parameters.
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof SpawnConditionValueApi
     */
    public getSpawnConditionValuesBulk(requestParameters: SpawnConditionValueApiGetSpawnConditionValuesBulkRequest, options?: any) {
        return SpawnConditionValueApiFp(this.configuration).getSpawnConditionValuesBulk(requestParameters.body, options).then((request) => request(this.axios, this.basePath));
    }

    /**
     * 
     * @summary Lists SpawnConditionValues
     * @param {SpawnConditionValueApiListSpawnConditionValuesRequest} requestParameters Request parameters.
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof SpawnConditionValueApi
     */
    public listSpawnConditionValues(requestParameters: SpawnConditionValueApiListSpawnConditionValuesRequest = {}, options?: any) {
        return SpawnConditionValueApiFp(this.configuration).listSpawnConditionValues(requestParameters.includes, requestParameters.where, requestParameters.whereOr, requestParameters.groupBy, requestParameters.limit, requestParameters.page, requestParameters.orderBy, requestParameters.orderDirection, requestParameters.select, options).then((request) => request(this.axios, this.basePath));
    }

    /**
     * 
     * @summary Updates SpawnConditionValue
     * @param {SpawnConditionValueApiUpdateSpawnConditionValueRequest} requestParameters Request parameters.
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof SpawnConditionValueApi
     */
    public updateSpawnConditionValue(requestParameters: SpawnConditionValueApiUpdateSpawnConditionValueRequest, options?: any) {
        return SpawnConditionValueApiFp(this.configuration).updateSpawnConditionValue(requestParameters.id, requestParameters.spawnConditionValue, options).then((request) => request(this.axios, this.basePath));
    }
}
