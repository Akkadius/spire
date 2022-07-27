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
import { ModelsQuestGlobal } from '../models';
/**
 * QuestGlobalApi - axios parameter creator
 * @export
 */
export const QuestGlobalApiAxiosParamCreator = function (configuration?: Configuration) {
    return {
        /**
         * 
         * @summary Creates QuestGlobal
         * @param {ModelsQuestGlobal} questGlobal QuestGlobal
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        createQuestGlobal: async (questGlobal: ModelsQuestGlobal, options: any = {}): Promise<RequestArgs> => {
            // verify required parameter 'questGlobal' is not null or undefined
            if (questGlobal === null || questGlobal === undefined) {
                throw new RequiredError('questGlobal','Required parameter questGlobal was null or undefined when calling createQuestGlobal.');
            }
            const localVarPath = `/quest_global`;
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
            const nonString = typeof questGlobal !== 'string';
            const needsSerialization = nonString && configuration && configuration.isJsonMime
                ? configuration.isJsonMime(localVarRequestOptions.headers['Content-Type'])
                : nonString;
            localVarRequestOptions.data =  needsSerialization
                ? JSON.stringify(questGlobal !== undefined ? questGlobal : {})
                : (questGlobal || "");

            return {
                url: localVarUrlObj.pathname + localVarUrlObj.search + localVarUrlObj.hash,
                options: localVarRequestOptions,
            };
        },
        /**
         * 
         * @summary Deletes QuestGlobal
         * @param {number} id charid
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        deleteQuestGlobal: async (id: number, options: any = {}): Promise<RequestArgs> => {
            // verify required parameter 'id' is not null or undefined
            if (id === null || id === undefined) {
                throw new RequiredError('id','Required parameter id was null or undefined when calling deleteQuestGlobal.');
            }
            const localVarPath = `/quest_global/{id}`
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
         * @summary Gets QuestGlobal
         * @param {number} id Id
         * @param {string} [includes] Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names 
         * @param {string} [select] Column names [.] separated to fetch specific fields in response
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        getQuestGlobal: async (id: number, includes?: string, select?: string, options: any = {}): Promise<RequestArgs> => {
            // verify required parameter 'id' is not null or undefined
            if (id === null || id === undefined) {
                throw new RequiredError('id','Required parameter id was null or undefined when calling getQuestGlobal.');
            }
            const localVarPath = `/quest_global/{id}`
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
         * @summary Gets QuestGlobals in bulk
         * @param {CrudcontrollersBulkFetchByIdsGetRequest} body body
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        getQuestGlobalsBulk: async (body: CrudcontrollersBulkFetchByIdsGetRequest, options: any = {}): Promise<RequestArgs> => {
            // verify required parameter 'body' is not null or undefined
            if (body === null || body === undefined) {
                throw new RequiredError('body','Required parameter body was null or undefined when calling getQuestGlobalsBulk.');
            }
            const localVarPath = `/quest_globals/bulk`;
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
         * @summary Lists QuestGlobals
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
        listQuestGlobals: async (includes?: string, where?: string, whereOr?: string, groupBy?: string, limit?: string, page?: number, orderBy?: string, orderDirection?: string, select?: string, options: any = {}): Promise<RequestArgs> => {
            const localVarPath = `/quest_globals`;
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
         * @summary Updates QuestGlobal
         * @param {number} id Id
         * @param {ModelsQuestGlobal} questGlobal QuestGlobal
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        updateQuestGlobal: async (id: number, questGlobal: ModelsQuestGlobal, options: any = {}): Promise<RequestArgs> => {
            // verify required parameter 'id' is not null or undefined
            if (id === null || id === undefined) {
                throw new RequiredError('id','Required parameter id was null or undefined when calling updateQuestGlobal.');
            }
            // verify required parameter 'questGlobal' is not null or undefined
            if (questGlobal === null || questGlobal === undefined) {
                throw new RequiredError('questGlobal','Required parameter questGlobal was null or undefined when calling updateQuestGlobal.');
            }
            const localVarPath = `/quest_global/{id}`
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
            const nonString = typeof questGlobal !== 'string';
            const needsSerialization = nonString && configuration && configuration.isJsonMime
                ? configuration.isJsonMime(localVarRequestOptions.headers['Content-Type'])
                : nonString;
            localVarRequestOptions.data =  needsSerialization
                ? JSON.stringify(questGlobal !== undefined ? questGlobal : {})
                : (questGlobal || "");

            return {
                url: localVarUrlObj.pathname + localVarUrlObj.search + localVarUrlObj.hash,
                options: localVarRequestOptions,
            };
        },
    }
};

/**
 * QuestGlobalApi - functional programming interface
 * @export
 */
export const QuestGlobalApiFp = function(configuration?: Configuration) {
    return {
        /**
         * 
         * @summary Creates QuestGlobal
         * @param {ModelsQuestGlobal} questGlobal QuestGlobal
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async createQuestGlobal(questGlobal: ModelsQuestGlobal, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsQuestGlobal>>> {
            const localVarAxiosArgs = await QuestGlobalApiAxiosParamCreator(configuration).createQuestGlobal(questGlobal, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        /**
         * 
         * @summary Deletes QuestGlobal
         * @param {number} id charid
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async deleteQuestGlobal(id: number, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<string>> {
            const localVarAxiosArgs = await QuestGlobalApiAxiosParamCreator(configuration).deleteQuestGlobal(id, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        /**
         * 
         * @summary Gets QuestGlobal
         * @param {number} id Id
         * @param {string} [includes] Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names 
         * @param {string} [select] Column names [.] separated to fetch specific fields in response
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async getQuestGlobal(id: number, includes?: string, select?: string, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsQuestGlobal>>> {
            const localVarAxiosArgs = await QuestGlobalApiAxiosParamCreator(configuration).getQuestGlobal(id, includes, select, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        /**
         * 
         * @summary Gets QuestGlobals in bulk
         * @param {CrudcontrollersBulkFetchByIdsGetRequest} body body
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async getQuestGlobalsBulk(body: CrudcontrollersBulkFetchByIdsGetRequest, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsQuestGlobal>>> {
            const localVarAxiosArgs = await QuestGlobalApiAxiosParamCreator(configuration).getQuestGlobalsBulk(body, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        /**
         * 
         * @summary Lists QuestGlobals
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
        async listQuestGlobals(includes?: string, where?: string, whereOr?: string, groupBy?: string, limit?: string, page?: number, orderBy?: string, orderDirection?: string, select?: string, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsQuestGlobal>>> {
            const localVarAxiosArgs = await QuestGlobalApiAxiosParamCreator(configuration).listQuestGlobals(includes, where, whereOr, groupBy, limit, page, orderBy, orderDirection, select, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        /**
         * 
         * @summary Updates QuestGlobal
         * @param {number} id Id
         * @param {ModelsQuestGlobal} questGlobal QuestGlobal
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async updateQuestGlobal(id: number, questGlobal: ModelsQuestGlobal, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsQuestGlobal>>> {
            const localVarAxiosArgs = await QuestGlobalApiAxiosParamCreator(configuration).updateQuestGlobal(id, questGlobal, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
    }
};

/**
 * QuestGlobalApi - factory interface
 * @export
 */
export const QuestGlobalApiFactory = function (configuration?: Configuration, basePath?: string, axios?: AxiosInstance) {
    return {
        /**
         * 
         * @summary Creates QuestGlobal
         * @param {ModelsQuestGlobal} questGlobal QuestGlobal
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        createQuestGlobal(questGlobal: ModelsQuestGlobal, options?: any): AxiosPromise<Array<ModelsQuestGlobal>> {
            return QuestGlobalApiFp(configuration).createQuestGlobal(questGlobal, options).then((request) => request(axios, basePath));
        },
        /**
         * 
         * @summary Deletes QuestGlobal
         * @param {number} id charid
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        deleteQuestGlobal(id: number, options?: any): AxiosPromise<string> {
            return QuestGlobalApiFp(configuration).deleteQuestGlobal(id, options).then((request) => request(axios, basePath));
        },
        /**
         * 
         * @summary Gets QuestGlobal
         * @param {number} id Id
         * @param {string} [includes] Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names 
         * @param {string} [select] Column names [.] separated to fetch specific fields in response
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        getQuestGlobal(id: number, includes?: string, select?: string, options?: any): AxiosPromise<Array<ModelsQuestGlobal>> {
            return QuestGlobalApiFp(configuration).getQuestGlobal(id, includes, select, options).then((request) => request(axios, basePath));
        },
        /**
         * 
         * @summary Gets QuestGlobals in bulk
         * @param {CrudcontrollersBulkFetchByIdsGetRequest} body body
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        getQuestGlobalsBulk(body: CrudcontrollersBulkFetchByIdsGetRequest, options?: any): AxiosPromise<Array<ModelsQuestGlobal>> {
            return QuestGlobalApiFp(configuration).getQuestGlobalsBulk(body, options).then((request) => request(axios, basePath));
        },
        /**
         * 
         * @summary Lists QuestGlobals
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
        listQuestGlobals(includes?: string, where?: string, whereOr?: string, groupBy?: string, limit?: string, page?: number, orderBy?: string, orderDirection?: string, select?: string, options?: any): AxiosPromise<Array<ModelsQuestGlobal>> {
            return QuestGlobalApiFp(configuration).listQuestGlobals(includes, where, whereOr, groupBy, limit, page, orderBy, orderDirection, select, options).then((request) => request(axios, basePath));
        },
        /**
         * 
         * @summary Updates QuestGlobal
         * @param {number} id Id
         * @param {ModelsQuestGlobal} questGlobal QuestGlobal
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        updateQuestGlobal(id: number, questGlobal: ModelsQuestGlobal, options?: any): AxiosPromise<Array<ModelsQuestGlobal>> {
            return QuestGlobalApiFp(configuration).updateQuestGlobal(id, questGlobal, options).then((request) => request(axios, basePath));
        },
    };
};

/**
 * Request parameters for createQuestGlobal operation in QuestGlobalApi.
 * @export
 * @interface QuestGlobalApiCreateQuestGlobalRequest
 */
export interface QuestGlobalApiCreateQuestGlobalRequest {
    /**
     * QuestGlobal
     * @type {ModelsQuestGlobal}
     * @memberof QuestGlobalApiCreateQuestGlobal
     */
    readonly questGlobal: ModelsQuestGlobal
}

/**
 * Request parameters for deleteQuestGlobal operation in QuestGlobalApi.
 * @export
 * @interface QuestGlobalApiDeleteQuestGlobalRequest
 */
export interface QuestGlobalApiDeleteQuestGlobalRequest {
    /**
     * charid
     * @type {number}
     * @memberof QuestGlobalApiDeleteQuestGlobal
     */
    readonly id: number
}

/**
 * Request parameters for getQuestGlobal operation in QuestGlobalApi.
 * @export
 * @interface QuestGlobalApiGetQuestGlobalRequest
 */
export interface QuestGlobalApiGetQuestGlobalRequest {
    /**
     * Id
     * @type {number}
     * @memberof QuestGlobalApiGetQuestGlobal
     */
    readonly id: number

    /**
     * Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names 
     * @type {string}
     * @memberof QuestGlobalApiGetQuestGlobal
     */
    readonly includes?: string

    /**
     * Column names [.] separated to fetch specific fields in response
     * @type {string}
     * @memberof QuestGlobalApiGetQuestGlobal
     */
    readonly select?: string
}

/**
 * Request parameters for getQuestGlobalsBulk operation in QuestGlobalApi.
 * @export
 * @interface QuestGlobalApiGetQuestGlobalsBulkRequest
 */
export interface QuestGlobalApiGetQuestGlobalsBulkRequest {
    /**
     * body
     * @type {CrudcontrollersBulkFetchByIdsGetRequest}
     * @memberof QuestGlobalApiGetQuestGlobalsBulk
     */
    readonly body: CrudcontrollersBulkFetchByIdsGetRequest
}

/**
 * Request parameters for listQuestGlobals operation in QuestGlobalApi.
 * @export
 * @interface QuestGlobalApiListQuestGlobalsRequest
 */
export interface QuestGlobalApiListQuestGlobalsRequest {
    /**
     * Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names 
     * @type {string}
     * @memberof QuestGlobalApiListQuestGlobals
     */
    readonly includes?: string

    /**
     * Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2
     * @type {string}
     * @memberof QuestGlobalApiListQuestGlobals
     */
    readonly where?: string

    /**
     * Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2
     * @type {string}
     * @memberof QuestGlobalApiListQuestGlobals
     */
    readonly whereOr?: string

    /**
     * Group by field. Multiple conditions [.] separated Example: field1.field2
     * @type {string}
     * @memberof QuestGlobalApiListQuestGlobals
     */
    readonly groupBy?: string

    /**
     * Rows to limit in response (Default: 10,000)
     * @type {string}
     * @memberof QuestGlobalApiListQuestGlobals
     */
    readonly limit?: string

    /**
     * Pagination page
     * @type {number}
     * @memberof QuestGlobalApiListQuestGlobals
     */
    readonly page?: number

    /**
     * Order by [field]
     * @type {string}
     * @memberof QuestGlobalApiListQuestGlobals
     */
    readonly orderBy?: string

    /**
     * Order by field direction
     * @type {string}
     * @memberof QuestGlobalApiListQuestGlobals
     */
    readonly orderDirection?: string

    /**
     * Column names [.] separated to fetch specific fields in response
     * @type {string}
     * @memberof QuestGlobalApiListQuestGlobals
     */
    readonly select?: string
}

/**
 * Request parameters for updateQuestGlobal operation in QuestGlobalApi.
 * @export
 * @interface QuestGlobalApiUpdateQuestGlobalRequest
 */
export interface QuestGlobalApiUpdateQuestGlobalRequest {
    /**
     * Id
     * @type {number}
     * @memberof QuestGlobalApiUpdateQuestGlobal
     */
    readonly id: number

    /**
     * QuestGlobal
     * @type {ModelsQuestGlobal}
     * @memberof QuestGlobalApiUpdateQuestGlobal
     */
    readonly questGlobal: ModelsQuestGlobal
}

/**
 * QuestGlobalApi - object-oriented interface
 * @export
 * @class QuestGlobalApi
 * @extends {BaseAPI}
 */
export class QuestGlobalApi extends BaseAPI {
    /**
     * 
     * @summary Creates QuestGlobal
     * @param {QuestGlobalApiCreateQuestGlobalRequest} requestParameters Request parameters.
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof QuestGlobalApi
     */
    public createQuestGlobal(requestParameters: QuestGlobalApiCreateQuestGlobalRequest, options?: any) {
        return QuestGlobalApiFp(this.configuration).createQuestGlobal(requestParameters.questGlobal, options).then((request) => request(this.axios, this.basePath));
    }

    /**
     * 
     * @summary Deletes QuestGlobal
     * @param {QuestGlobalApiDeleteQuestGlobalRequest} requestParameters Request parameters.
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof QuestGlobalApi
     */
    public deleteQuestGlobal(requestParameters: QuestGlobalApiDeleteQuestGlobalRequest, options?: any) {
        return QuestGlobalApiFp(this.configuration).deleteQuestGlobal(requestParameters.id, options).then((request) => request(this.axios, this.basePath));
    }

    /**
     * 
     * @summary Gets QuestGlobal
     * @param {QuestGlobalApiGetQuestGlobalRequest} requestParameters Request parameters.
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof QuestGlobalApi
     */
    public getQuestGlobal(requestParameters: QuestGlobalApiGetQuestGlobalRequest, options?: any) {
        return QuestGlobalApiFp(this.configuration).getQuestGlobal(requestParameters.id, requestParameters.includes, requestParameters.select, options).then((request) => request(this.axios, this.basePath));
    }

    /**
     * 
     * @summary Gets QuestGlobals in bulk
     * @param {QuestGlobalApiGetQuestGlobalsBulkRequest} requestParameters Request parameters.
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof QuestGlobalApi
     */
    public getQuestGlobalsBulk(requestParameters: QuestGlobalApiGetQuestGlobalsBulkRequest, options?: any) {
        return QuestGlobalApiFp(this.configuration).getQuestGlobalsBulk(requestParameters.body, options).then((request) => request(this.axios, this.basePath));
    }

    /**
     * 
     * @summary Lists QuestGlobals
     * @param {QuestGlobalApiListQuestGlobalsRequest} requestParameters Request parameters.
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof QuestGlobalApi
     */
    public listQuestGlobals(requestParameters: QuestGlobalApiListQuestGlobalsRequest = {}, options?: any) {
        return QuestGlobalApiFp(this.configuration).listQuestGlobals(requestParameters.includes, requestParameters.where, requestParameters.whereOr, requestParameters.groupBy, requestParameters.limit, requestParameters.page, requestParameters.orderBy, requestParameters.orderDirection, requestParameters.select, options).then((request) => request(this.axios, this.basePath));
    }

    /**
     * 
     * @summary Updates QuestGlobal
     * @param {QuestGlobalApiUpdateQuestGlobalRequest} requestParameters Request parameters.
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof QuestGlobalApi
     */
    public updateQuestGlobal(requestParameters: QuestGlobalApiUpdateQuestGlobalRequest, options?: any) {
        return QuestGlobalApiFp(this.configuration).updateQuestGlobal(requestParameters.id, requestParameters.questGlobal, options).then((request) => request(this.axios, this.basePath));
    }
}
