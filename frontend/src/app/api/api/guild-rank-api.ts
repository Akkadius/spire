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
import { ModelsGuildRank } from '../models';
/**
 * GuildRankApi - axios parameter creator
 * @export
 */
export const GuildRankApiAxiosParamCreator = function (configuration?: Configuration) {
    return {
        /**
         * 
         * @summary Creates GuildRank
         * @param {ModelsGuildRank} guildRank GuildRank
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        createGuildRank: async (guildRank: ModelsGuildRank, options: any = {}): Promise<RequestArgs> => {
            // verify required parameter 'guildRank' is not null or undefined
            if (guildRank === null || guildRank === undefined) {
                throw new RequiredError('guildRank','Required parameter guildRank was null or undefined when calling createGuildRank.');
            }
            const localVarPath = `/guild_rank`;
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
            const nonString = typeof guildRank !== 'string';
            const needsSerialization = nonString && configuration && configuration.isJsonMime
                ? configuration.isJsonMime(localVarRequestOptions.headers['Content-Type'])
                : nonString;
            localVarRequestOptions.data =  needsSerialization
                ? JSON.stringify(guildRank !== undefined ? guildRank : {})
                : (guildRank || "");

            return {
                url: localVarUrlObj.pathname + localVarUrlObj.search + localVarUrlObj.hash,
                options: localVarRequestOptions,
            };
        },
        /**
         * 
         * @summary Deletes GuildRank
         * @param {number} id guildId
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        deleteGuildRank: async (id: number, options: any = {}): Promise<RequestArgs> => {
            // verify required parameter 'id' is not null or undefined
            if (id === null || id === undefined) {
                throw new RequiredError('id','Required parameter id was null or undefined when calling deleteGuildRank.');
            }
            const localVarPath = `/guild_rank/{id}`
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
         * @summary Gets GuildRank
         * @param {number} id Id
         * @param {string} [includes] Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names 
         * @param {string} [select] Column names [.] separated to fetch specific fields in response
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        getGuildRank: async (id: number, includes?: string, select?: string, options: any = {}): Promise<RequestArgs> => {
            // verify required parameter 'id' is not null or undefined
            if (id === null || id === undefined) {
                throw new RequiredError('id','Required parameter id was null or undefined when calling getGuildRank.');
            }
            const localVarPath = `/guild_rank/{id}`
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
         * @summary Gets GuildRanks in bulk
         * @param {CrudcontrollersBulkFetchByIdsGetRequest} body body
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        getGuildRanksBulk: async (body: CrudcontrollersBulkFetchByIdsGetRequest, options: any = {}): Promise<RequestArgs> => {
            // verify required parameter 'body' is not null or undefined
            if (body === null || body === undefined) {
                throw new RequiredError('body','Required parameter body was null or undefined when calling getGuildRanksBulk.');
            }
            const localVarPath = `/guild_ranks/bulk`;
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
         * @summary Lists GuildRanks
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
        listGuildRanks: async (includes?: string, where?: string, whereOr?: string, groupBy?: string, limit?: string, page?: number, orderBy?: string, orderDirection?: string, select?: string, options: any = {}): Promise<RequestArgs> => {
            const localVarPath = `/guild_ranks`;
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
         * @summary Updates GuildRank
         * @param {number} id Id
         * @param {ModelsGuildRank} guildRank GuildRank
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        updateGuildRank: async (id: number, guildRank: ModelsGuildRank, options: any = {}): Promise<RequestArgs> => {
            // verify required parameter 'id' is not null or undefined
            if (id === null || id === undefined) {
                throw new RequiredError('id','Required parameter id was null or undefined when calling updateGuildRank.');
            }
            // verify required parameter 'guildRank' is not null or undefined
            if (guildRank === null || guildRank === undefined) {
                throw new RequiredError('guildRank','Required parameter guildRank was null or undefined when calling updateGuildRank.');
            }
            const localVarPath = `/guild_rank/{id}`
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
            const nonString = typeof guildRank !== 'string';
            const needsSerialization = nonString && configuration && configuration.isJsonMime
                ? configuration.isJsonMime(localVarRequestOptions.headers['Content-Type'])
                : nonString;
            localVarRequestOptions.data =  needsSerialization
                ? JSON.stringify(guildRank !== undefined ? guildRank : {})
                : (guildRank || "");

            return {
                url: localVarUrlObj.pathname + localVarUrlObj.search + localVarUrlObj.hash,
                options: localVarRequestOptions,
            };
        },
    }
};

/**
 * GuildRankApi - functional programming interface
 * @export
 */
export const GuildRankApiFp = function(configuration?: Configuration) {
    return {
        /**
         * 
         * @summary Creates GuildRank
         * @param {ModelsGuildRank} guildRank GuildRank
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async createGuildRank(guildRank: ModelsGuildRank, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsGuildRank>>> {
            const localVarAxiosArgs = await GuildRankApiAxiosParamCreator(configuration).createGuildRank(guildRank, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        /**
         * 
         * @summary Deletes GuildRank
         * @param {number} id guildId
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async deleteGuildRank(id: number, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<string>> {
            const localVarAxiosArgs = await GuildRankApiAxiosParamCreator(configuration).deleteGuildRank(id, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        /**
         * 
         * @summary Gets GuildRank
         * @param {number} id Id
         * @param {string} [includes] Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names 
         * @param {string} [select] Column names [.] separated to fetch specific fields in response
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async getGuildRank(id: number, includes?: string, select?: string, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsGuildRank>>> {
            const localVarAxiosArgs = await GuildRankApiAxiosParamCreator(configuration).getGuildRank(id, includes, select, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        /**
         * 
         * @summary Gets GuildRanks in bulk
         * @param {CrudcontrollersBulkFetchByIdsGetRequest} body body
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async getGuildRanksBulk(body: CrudcontrollersBulkFetchByIdsGetRequest, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsGuildRank>>> {
            const localVarAxiosArgs = await GuildRankApiAxiosParamCreator(configuration).getGuildRanksBulk(body, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        /**
         * 
         * @summary Lists GuildRanks
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
        async listGuildRanks(includes?: string, where?: string, whereOr?: string, groupBy?: string, limit?: string, page?: number, orderBy?: string, orderDirection?: string, select?: string, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsGuildRank>>> {
            const localVarAxiosArgs = await GuildRankApiAxiosParamCreator(configuration).listGuildRanks(includes, where, whereOr, groupBy, limit, page, orderBy, orderDirection, select, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        /**
         * 
         * @summary Updates GuildRank
         * @param {number} id Id
         * @param {ModelsGuildRank} guildRank GuildRank
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async updateGuildRank(id: number, guildRank: ModelsGuildRank, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsGuildRank>>> {
            const localVarAxiosArgs = await GuildRankApiAxiosParamCreator(configuration).updateGuildRank(id, guildRank, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
    }
};

/**
 * GuildRankApi - factory interface
 * @export
 */
export const GuildRankApiFactory = function (configuration?: Configuration, basePath?: string, axios?: AxiosInstance) {
    return {
        /**
         * 
         * @summary Creates GuildRank
         * @param {ModelsGuildRank} guildRank GuildRank
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        createGuildRank(guildRank: ModelsGuildRank, options?: any): AxiosPromise<Array<ModelsGuildRank>> {
            return GuildRankApiFp(configuration).createGuildRank(guildRank, options).then((request) => request(axios, basePath));
        },
        /**
         * 
         * @summary Deletes GuildRank
         * @param {number} id guildId
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        deleteGuildRank(id: number, options?: any): AxiosPromise<string> {
            return GuildRankApiFp(configuration).deleteGuildRank(id, options).then((request) => request(axios, basePath));
        },
        /**
         * 
         * @summary Gets GuildRank
         * @param {number} id Id
         * @param {string} [includes] Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names 
         * @param {string} [select] Column names [.] separated to fetch specific fields in response
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        getGuildRank(id: number, includes?: string, select?: string, options?: any): AxiosPromise<Array<ModelsGuildRank>> {
            return GuildRankApiFp(configuration).getGuildRank(id, includes, select, options).then((request) => request(axios, basePath));
        },
        /**
         * 
         * @summary Gets GuildRanks in bulk
         * @param {CrudcontrollersBulkFetchByIdsGetRequest} body body
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        getGuildRanksBulk(body: CrudcontrollersBulkFetchByIdsGetRequest, options?: any): AxiosPromise<Array<ModelsGuildRank>> {
            return GuildRankApiFp(configuration).getGuildRanksBulk(body, options).then((request) => request(axios, basePath));
        },
        /**
         * 
         * @summary Lists GuildRanks
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
        listGuildRanks(includes?: string, where?: string, whereOr?: string, groupBy?: string, limit?: string, page?: number, orderBy?: string, orderDirection?: string, select?: string, options?: any): AxiosPromise<Array<ModelsGuildRank>> {
            return GuildRankApiFp(configuration).listGuildRanks(includes, where, whereOr, groupBy, limit, page, orderBy, orderDirection, select, options).then((request) => request(axios, basePath));
        },
        /**
         * 
         * @summary Updates GuildRank
         * @param {number} id Id
         * @param {ModelsGuildRank} guildRank GuildRank
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        updateGuildRank(id: number, guildRank: ModelsGuildRank, options?: any): AxiosPromise<Array<ModelsGuildRank>> {
            return GuildRankApiFp(configuration).updateGuildRank(id, guildRank, options).then((request) => request(axios, basePath));
        },
    };
};

/**
 * Request parameters for createGuildRank operation in GuildRankApi.
 * @export
 * @interface GuildRankApiCreateGuildRankRequest
 */
export interface GuildRankApiCreateGuildRankRequest {
    /**
     * GuildRank
     * @type {ModelsGuildRank}
     * @memberof GuildRankApiCreateGuildRank
     */
    readonly guildRank: ModelsGuildRank
}

/**
 * Request parameters for deleteGuildRank operation in GuildRankApi.
 * @export
 * @interface GuildRankApiDeleteGuildRankRequest
 */
export interface GuildRankApiDeleteGuildRankRequest {
    /**
     * guildId
     * @type {number}
     * @memberof GuildRankApiDeleteGuildRank
     */
    readonly id: number
}

/**
 * Request parameters for getGuildRank operation in GuildRankApi.
 * @export
 * @interface GuildRankApiGetGuildRankRequest
 */
export interface GuildRankApiGetGuildRankRequest {
    /**
     * Id
     * @type {number}
     * @memberof GuildRankApiGetGuildRank
     */
    readonly id: number

    /**
     * Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names 
     * @type {string}
     * @memberof GuildRankApiGetGuildRank
     */
    readonly includes?: string

    /**
     * Column names [.] separated to fetch specific fields in response
     * @type {string}
     * @memberof GuildRankApiGetGuildRank
     */
    readonly select?: string
}

/**
 * Request parameters for getGuildRanksBulk operation in GuildRankApi.
 * @export
 * @interface GuildRankApiGetGuildRanksBulkRequest
 */
export interface GuildRankApiGetGuildRanksBulkRequest {
    /**
     * body
     * @type {CrudcontrollersBulkFetchByIdsGetRequest}
     * @memberof GuildRankApiGetGuildRanksBulk
     */
    readonly body: CrudcontrollersBulkFetchByIdsGetRequest
}

/**
 * Request parameters for listGuildRanks operation in GuildRankApi.
 * @export
 * @interface GuildRankApiListGuildRanksRequest
 */
export interface GuildRankApiListGuildRanksRequest {
    /**
     * Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names 
     * @type {string}
     * @memberof GuildRankApiListGuildRanks
     */
    readonly includes?: string

    /**
     * Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2
     * @type {string}
     * @memberof GuildRankApiListGuildRanks
     */
    readonly where?: string

    /**
     * Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2
     * @type {string}
     * @memberof GuildRankApiListGuildRanks
     */
    readonly whereOr?: string

    /**
     * Group by field. Multiple conditions [.] separated Example: field1.field2
     * @type {string}
     * @memberof GuildRankApiListGuildRanks
     */
    readonly groupBy?: string

    /**
     * Rows to limit in response (Default: 10,000)
     * @type {string}
     * @memberof GuildRankApiListGuildRanks
     */
    readonly limit?: string

    /**
     * Pagination page
     * @type {number}
     * @memberof GuildRankApiListGuildRanks
     */
    readonly page?: number

    /**
     * Order by [field]
     * @type {string}
     * @memberof GuildRankApiListGuildRanks
     */
    readonly orderBy?: string

    /**
     * Order by field direction
     * @type {string}
     * @memberof GuildRankApiListGuildRanks
     */
    readonly orderDirection?: string

    /**
     * Column names [.] separated to fetch specific fields in response
     * @type {string}
     * @memberof GuildRankApiListGuildRanks
     */
    readonly select?: string
}

/**
 * Request parameters for updateGuildRank operation in GuildRankApi.
 * @export
 * @interface GuildRankApiUpdateGuildRankRequest
 */
export interface GuildRankApiUpdateGuildRankRequest {
    /**
     * Id
     * @type {number}
     * @memberof GuildRankApiUpdateGuildRank
     */
    readonly id: number

    /**
     * GuildRank
     * @type {ModelsGuildRank}
     * @memberof GuildRankApiUpdateGuildRank
     */
    readonly guildRank: ModelsGuildRank
}

/**
 * GuildRankApi - object-oriented interface
 * @export
 * @class GuildRankApi
 * @extends {BaseAPI}
 */
export class GuildRankApi extends BaseAPI {
    /**
     * 
     * @summary Creates GuildRank
     * @param {GuildRankApiCreateGuildRankRequest} requestParameters Request parameters.
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof GuildRankApi
     */
    public createGuildRank(requestParameters: GuildRankApiCreateGuildRankRequest, options?: any) {
        return GuildRankApiFp(this.configuration).createGuildRank(requestParameters.guildRank, options).then((request) => request(this.axios, this.basePath));
    }

    /**
     * 
     * @summary Deletes GuildRank
     * @param {GuildRankApiDeleteGuildRankRequest} requestParameters Request parameters.
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof GuildRankApi
     */
    public deleteGuildRank(requestParameters: GuildRankApiDeleteGuildRankRequest, options?: any) {
        return GuildRankApiFp(this.configuration).deleteGuildRank(requestParameters.id, options).then((request) => request(this.axios, this.basePath));
    }

    /**
     * 
     * @summary Gets GuildRank
     * @param {GuildRankApiGetGuildRankRequest} requestParameters Request parameters.
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof GuildRankApi
     */
    public getGuildRank(requestParameters: GuildRankApiGetGuildRankRequest, options?: any) {
        return GuildRankApiFp(this.configuration).getGuildRank(requestParameters.id, requestParameters.includes, requestParameters.select, options).then((request) => request(this.axios, this.basePath));
    }

    /**
     * 
     * @summary Gets GuildRanks in bulk
     * @param {GuildRankApiGetGuildRanksBulkRequest} requestParameters Request parameters.
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof GuildRankApi
     */
    public getGuildRanksBulk(requestParameters: GuildRankApiGetGuildRanksBulkRequest, options?: any) {
        return GuildRankApiFp(this.configuration).getGuildRanksBulk(requestParameters.body, options).then((request) => request(this.axios, this.basePath));
    }

    /**
     * 
     * @summary Lists GuildRanks
     * @param {GuildRankApiListGuildRanksRequest} requestParameters Request parameters.
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof GuildRankApi
     */
    public listGuildRanks(requestParameters: GuildRankApiListGuildRanksRequest = {}, options?: any) {
        return GuildRankApiFp(this.configuration).listGuildRanks(requestParameters.includes, requestParameters.where, requestParameters.whereOr, requestParameters.groupBy, requestParameters.limit, requestParameters.page, requestParameters.orderBy, requestParameters.orderDirection, requestParameters.select, options).then((request) => request(this.axios, this.basePath));
    }

    /**
     * 
     * @summary Updates GuildRank
     * @param {GuildRankApiUpdateGuildRankRequest} requestParameters Request parameters.
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof GuildRankApi
     */
    public updateGuildRank(requestParameters: GuildRankApiUpdateGuildRankRequest, options?: any) {
        return GuildRankApiFp(this.configuration).updateGuildRank(requestParameters.id, requestParameters.guildRank, options).then((request) => request(this.axios, this.basePath));
    }
}
