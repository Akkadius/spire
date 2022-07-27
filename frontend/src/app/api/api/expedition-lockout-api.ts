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
import { ModelsExpeditionLockout } from '../models';
/**
 * ExpeditionLockoutApi - axios parameter creator
 * @export
 */
export const ExpeditionLockoutApiAxiosParamCreator = function (configuration?: Configuration) {
    return {
        /**
         * 
         * @summary Creates ExpeditionLockout
         * @param {ModelsExpeditionLockout} expeditionLockout ExpeditionLockout
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        createExpeditionLockout: async (expeditionLockout: ModelsExpeditionLockout, options: any = {}): Promise<RequestArgs> => {
            // verify required parameter 'expeditionLockout' is not null or undefined
            if (expeditionLockout === null || expeditionLockout === undefined) {
                throw new RequiredError('expeditionLockout','Required parameter expeditionLockout was null or undefined when calling createExpeditionLockout.');
            }
            const localVarPath = `/expedition_lockout`;
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
            const nonString = typeof expeditionLockout !== 'string';
            const needsSerialization = nonString && configuration && configuration.isJsonMime
                ? configuration.isJsonMime(localVarRequestOptions.headers['Content-Type'])
                : nonString;
            localVarRequestOptions.data =  needsSerialization
                ? JSON.stringify(expeditionLockout !== undefined ? expeditionLockout : {})
                : (expeditionLockout || "");

            return {
                url: localVarUrlObj.pathname + localVarUrlObj.search + localVarUrlObj.hash,
                options: localVarRequestOptions,
            };
        },
        /**
         * 
         * @summary Deletes ExpeditionLockout
         * @param {number} id id
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        deleteExpeditionLockout: async (id: number, options: any = {}): Promise<RequestArgs> => {
            // verify required parameter 'id' is not null or undefined
            if (id === null || id === undefined) {
                throw new RequiredError('id','Required parameter id was null or undefined when calling deleteExpeditionLockout.');
            }
            const localVarPath = `/expedition_lockout/{id}`
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
         * @summary Gets ExpeditionLockout
         * @param {number} id Id
         * @param {string} [includes] Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names 
         * @param {string} [select] Column names [.] separated to fetch specific fields in response
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        getExpeditionLockout: async (id: number, includes?: string, select?: string, options: any = {}): Promise<RequestArgs> => {
            // verify required parameter 'id' is not null or undefined
            if (id === null || id === undefined) {
                throw new RequiredError('id','Required parameter id was null or undefined when calling getExpeditionLockout.');
            }
            const localVarPath = `/expedition_lockout/{id}`
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
         * @summary Gets ExpeditionLockouts in bulk
         * @param {CrudcontrollersBulkFetchByIdsGetRequest} body body
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        getExpeditionLockoutsBulk: async (body: CrudcontrollersBulkFetchByIdsGetRequest, options: any = {}): Promise<RequestArgs> => {
            // verify required parameter 'body' is not null or undefined
            if (body === null || body === undefined) {
                throw new RequiredError('body','Required parameter body was null or undefined when calling getExpeditionLockoutsBulk.');
            }
            const localVarPath = `/expedition_lockouts/bulk`;
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
         * @summary Lists ExpeditionLockouts
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
        listExpeditionLockouts: async (includes?: string, where?: string, whereOr?: string, groupBy?: string, limit?: string, page?: number, orderBy?: string, orderDirection?: string, select?: string, options: any = {}): Promise<RequestArgs> => {
            const localVarPath = `/expedition_lockouts`;
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
         * @summary Updates ExpeditionLockout
         * @param {number} id Id
         * @param {ModelsExpeditionLockout} expeditionLockout ExpeditionLockout
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        updateExpeditionLockout: async (id: number, expeditionLockout: ModelsExpeditionLockout, options: any = {}): Promise<RequestArgs> => {
            // verify required parameter 'id' is not null or undefined
            if (id === null || id === undefined) {
                throw new RequiredError('id','Required parameter id was null or undefined when calling updateExpeditionLockout.');
            }
            // verify required parameter 'expeditionLockout' is not null or undefined
            if (expeditionLockout === null || expeditionLockout === undefined) {
                throw new RequiredError('expeditionLockout','Required parameter expeditionLockout was null or undefined when calling updateExpeditionLockout.');
            }
            const localVarPath = `/expedition_lockout/{id}`
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
            const nonString = typeof expeditionLockout !== 'string';
            const needsSerialization = nonString && configuration && configuration.isJsonMime
                ? configuration.isJsonMime(localVarRequestOptions.headers['Content-Type'])
                : nonString;
            localVarRequestOptions.data =  needsSerialization
                ? JSON.stringify(expeditionLockout !== undefined ? expeditionLockout : {})
                : (expeditionLockout || "");

            return {
                url: localVarUrlObj.pathname + localVarUrlObj.search + localVarUrlObj.hash,
                options: localVarRequestOptions,
            };
        },
    }
};

/**
 * ExpeditionLockoutApi - functional programming interface
 * @export
 */
export const ExpeditionLockoutApiFp = function(configuration?: Configuration) {
    return {
        /**
         * 
         * @summary Creates ExpeditionLockout
         * @param {ModelsExpeditionLockout} expeditionLockout ExpeditionLockout
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async createExpeditionLockout(expeditionLockout: ModelsExpeditionLockout, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsExpeditionLockout>>> {
            const localVarAxiosArgs = await ExpeditionLockoutApiAxiosParamCreator(configuration).createExpeditionLockout(expeditionLockout, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        /**
         * 
         * @summary Deletes ExpeditionLockout
         * @param {number} id id
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async deleteExpeditionLockout(id: number, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<string>> {
            const localVarAxiosArgs = await ExpeditionLockoutApiAxiosParamCreator(configuration).deleteExpeditionLockout(id, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        /**
         * 
         * @summary Gets ExpeditionLockout
         * @param {number} id Id
         * @param {string} [includes] Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names 
         * @param {string} [select] Column names [.] separated to fetch specific fields in response
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async getExpeditionLockout(id: number, includes?: string, select?: string, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsExpeditionLockout>>> {
            const localVarAxiosArgs = await ExpeditionLockoutApiAxiosParamCreator(configuration).getExpeditionLockout(id, includes, select, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        /**
         * 
         * @summary Gets ExpeditionLockouts in bulk
         * @param {CrudcontrollersBulkFetchByIdsGetRequest} body body
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async getExpeditionLockoutsBulk(body: CrudcontrollersBulkFetchByIdsGetRequest, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsExpeditionLockout>>> {
            const localVarAxiosArgs = await ExpeditionLockoutApiAxiosParamCreator(configuration).getExpeditionLockoutsBulk(body, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        /**
         * 
         * @summary Lists ExpeditionLockouts
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
        async listExpeditionLockouts(includes?: string, where?: string, whereOr?: string, groupBy?: string, limit?: string, page?: number, orderBy?: string, orderDirection?: string, select?: string, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsExpeditionLockout>>> {
            const localVarAxiosArgs = await ExpeditionLockoutApiAxiosParamCreator(configuration).listExpeditionLockouts(includes, where, whereOr, groupBy, limit, page, orderBy, orderDirection, select, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        /**
         * 
         * @summary Updates ExpeditionLockout
         * @param {number} id Id
         * @param {ModelsExpeditionLockout} expeditionLockout ExpeditionLockout
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async updateExpeditionLockout(id: number, expeditionLockout: ModelsExpeditionLockout, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsExpeditionLockout>>> {
            const localVarAxiosArgs = await ExpeditionLockoutApiAxiosParamCreator(configuration).updateExpeditionLockout(id, expeditionLockout, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
    }
};

/**
 * ExpeditionLockoutApi - factory interface
 * @export
 */
export const ExpeditionLockoutApiFactory = function (configuration?: Configuration, basePath?: string, axios?: AxiosInstance) {
    return {
        /**
         * 
         * @summary Creates ExpeditionLockout
         * @param {ModelsExpeditionLockout} expeditionLockout ExpeditionLockout
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        createExpeditionLockout(expeditionLockout: ModelsExpeditionLockout, options?: any): AxiosPromise<Array<ModelsExpeditionLockout>> {
            return ExpeditionLockoutApiFp(configuration).createExpeditionLockout(expeditionLockout, options).then((request) => request(axios, basePath));
        },
        /**
         * 
         * @summary Deletes ExpeditionLockout
         * @param {number} id id
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        deleteExpeditionLockout(id: number, options?: any): AxiosPromise<string> {
            return ExpeditionLockoutApiFp(configuration).deleteExpeditionLockout(id, options).then((request) => request(axios, basePath));
        },
        /**
         * 
         * @summary Gets ExpeditionLockout
         * @param {number} id Id
         * @param {string} [includes] Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names 
         * @param {string} [select] Column names [.] separated to fetch specific fields in response
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        getExpeditionLockout(id: number, includes?: string, select?: string, options?: any): AxiosPromise<Array<ModelsExpeditionLockout>> {
            return ExpeditionLockoutApiFp(configuration).getExpeditionLockout(id, includes, select, options).then((request) => request(axios, basePath));
        },
        /**
         * 
         * @summary Gets ExpeditionLockouts in bulk
         * @param {CrudcontrollersBulkFetchByIdsGetRequest} body body
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        getExpeditionLockoutsBulk(body: CrudcontrollersBulkFetchByIdsGetRequest, options?: any): AxiosPromise<Array<ModelsExpeditionLockout>> {
            return ExpeditionLockoutApiFp(configuration).getExpeditionLockoutsBulk(body, options).then((request) => request(axios, basePath));
        },
        /**
         * 
         * @summary Lists ExpeditionLockouts
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
        listExpeditionLockouts(includes?: string, where?: string, whereOr?: string, groupBy?: string, limit?: string, page?: number, orderBy?: string, orderDirection?: string, select?: string, options?: any): AxiosPromise<Array<ModelsExpeditionLockout>> {
            return ExpeditionLockoutApiFp(configuration).listExpeditionLockouts(includes, where, whereOr, groupBy, limit, page, orderBy, orderDirection, select, options).then((request) => request(axios, basePath));
        },
        /**
         * 
         * @summary Updates ExpeditionLockout
         * @param {number} id Id
         * @param {ModelsExpeditionLockout} expeditionLockout ExpeditionLockout
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        updateExpeditionLockout(id: number, expeditionLockout: ModelsExpeditionLockout, options?: any): AxiosPromise<Array<ModelsExpeditionLockout>> {
            return ExpeditionLockoutApiFp(configuration).updateExpeditionLockout(id, expeditionLockout, options).then((request) => request(axios, basePath));
        },
    };
};

/**
 * Request parameters for createExpeditionLockout operation in ExpeditionLockoutApi.
 * @export
 * @interface ExpeditionLockoutApiCreateExpeditionLockoutRequest
 */
export interface ExpeditionLockoutApiCreateExpeditionLockoutRequest {
    /**
     * ExpeditionLockout
     * @type {ModelsExpeditionLockout}
     * @memberof ExpeditionLockoutApiCreateExpeditionLockout
     */
    readonly expeditionLockout: ModelsExpeditionLockout
}

/**
 * Request parameters for deleteExpeditionLockout operation in ExpeditionLockoutApi.
 * @export
 * @interface ExpeditionLockoutApiDeleteExpeditionLockoutRequest
 */
export interface ExpeditionLockoutApiDeleteExpeditionLockoutRequest {
    /**
     * id
     * @type {number}
     * @memberof ExpeditionLockoutApiDeleteExpeditionLockout
     */
    readonly id: number
}

/**
 * Request parameters for getExpeditionLockout operation in ExpeditionLockoutApi.
 * @export
 * @interface ExpeditionLockoutApiGetExpeditionLockoutRequest
 */
export interface ExpeditionLockoutApiGetExpeditionLockoutRequest {
    /**
     * Id
     * @type {number}
     * @memberof ExpeditionLockoutApiGetExpeditionLockout
     */
    readonly id: number

    /**
     * Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names 
     * @type {string}
     * @memberof ExpeditionLockoutApiGetExpeditionLockout
     */
    readonly includes?: string

    /**
     * Column names [.] separated to fetch specific fields in response
     * @type {string}
     * @memberof ExpeditionLockoutApiGetExpeditionLockout
     */
    readonly select?: string
}

/**
 * Request parameters for getExpeditionLockoutsBulk operation in ExpeditionLockoutApi.
 * @export
 * @interface ExpeditionLockoutApiGetExpeditionLockoutsBulkRequest
 */
export interface ExpeditionLockoutApiGetExpeditionLockoutsBulkRequest {
    /**
     * body
     * @type {CrudcontrollersBulkFetchByIdsGetRequest}
     * @memberof ExpeditionLockoutApiGetExpeditionLockoutsBulk
     */
    readonly body: CrudcontrollersBulkFetchByIdsGetRequest
}

/**
 * Request parameters for listExpeditionLockouts operation in ExpeditionLockoutApi.
 * @export
 * @interface ExpeditionLockoutApiListExpeditionLockoutsRequest
 */
export interface ExpeditionLockoutApiListExpeditionLockoutsRequest {
    /**
     * Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names 
     * @type {string}
     * @memberof ExpeditionLockoutApiListExpeditionLockouts
     */
    readonly includes?: string

    /**
     * Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2
     * @type {string}
     * @memberof ExpeditionLockoutApiListExpeditionLockouts
     */
    readonly where?: string

    /**
     * Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2
     * @type {string}
     * @memberof ExpeditionLockoutApiListExpeditionLockouts
     */
    readonly whereOr?: string

    /**
     * Group by field. Multiple conditions [.] separated Example: field1.field2
     * @type {string}
     * @memberof ExpeditionLockoutApiListExpeditionLockouts
     */
    readonly groupBy?: string

    /**
     * Rows to limit in response (Default: 10,000)
     * @type {string}
     * @memberof ExpeditionLockoutApiListExpeditionLockouts
     */
    readonly limit?: string

    /**
     * Pagination page
     * @type {number}
     * @memberof ExpeditionLockoutApiListExpeditionLockouts
     */
    readonly page?: number

    /**
     * Order by [field]
     * @type {string}
     * @memberof ExpeditionLockoutApiListExpeditionLockouts
     */
    readonly orderBy?: string

    /**
     * Order by field direction
     * @type {string}
     * @memberof ExpeditionLockoutApiListExpeditionLockouts
     */
    readonly orderDirection?: string

    /**
     * Column names [.] separated to fetch specific fields in response
     * @type {string}
     * @memberof ExpeditionLockoutApiListExpeditionLockouts
     */
    readonly select?: string
}

/**
 * Request parameters for updateExpeditionLockout operation in ExpeditionLockoutApi.
 * @export
 * @interface ExpeditionLockoutApiUpdateExpeditionLockoutRequest
 */
export interface ExpeditionLockoutApiUpdateExpeditionLockoutRequest {
    /**
     * Id
     * @type {number}
     * @memberof ExpeditionLockoutApiUpdateExpeditionLockout
     */
    readonly id: number

    /**
     * ExpeditionLockout
     * @type {ModelsExpeditionLockout}
     * @memberof ExpeditionLockoutApiUpdateExpeditionLockout
     */
    readonly expeditionLockout: ModelsExpeditionLockout
}

/**
 * ExpeditionLockoutApi - object-oriented interface
 * @export
 * @class ExpeditionLockoutApi
 * @extends {BaseAPI}
 */
export class ExpeditionLockoutApi extends BaseAPI {
    /**
     * 
     * @summary Creates ExpeditionLockout
     * @param {ExpeditionLockoutApiCreateExpeditionLockoutRequest} requestParameters Request parameters.
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof ExpeditionLockoutApi
     */
    public createExpeditionLockout(requestParameters: ExpeditionLockoutApiCreateExpeditionLockoutRequest, options?: any) {
        return ExpeditionLockoutApiFp(this.configuration).createExpeditionLockout(requestParameters.expeditionLockout, options).then((request) => request(this.axios, this.basePath));
    }

    /**
     * 
     * @summary Deletes ExpeditionLockout
     * @param {ExpeditionLockoutApiDeleteExpeditionLockoutRequest} requestParameters Request parameters.
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof ExpeditionLockoutApi
     */
    public deleteExpeditionLockout(requestParameters: ExpeditionLockoutApiDeleteExpeditionLockoutRequest, options?: any) {
        return ExpeditionLockoutApiFp(this.configuration).deleteExpeditionLockout(requestParameters.id, options).then((request) => request(this.axios, this.basePath));
    }

    /**
     * 
     * @summary Gets ExpeditionLockout
     * @param {ExpeditionLockoutApiGetExpeditionLockoutRequest} requestParameters Request parameters.
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof ExpeditionLockoutApi
     */
    public getExpeditionLockout(requestParameters: ExpeditionLockoutApiGetExpeditionLockoutRequest, options?: any) {
        return ExpeditionLockoutApiFp(this.configuration).getExpeditionLockout(requestParameters.id, requestParameters.includes, requestParameters.select, options).then((request) => request(this.axios, this.basePath));
    }

    /**
     * 
     * @summary Gets ExpeditionLockouts in bulk
     * @param {ExpeditionLockoutApiGetExpeditionLockoutsBulkRequest} requestParameters Request parameters.
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof ExpeditionLockoutApi
     */
    public getExpeditionLockoutsBulk(requestParameters: ExpeditionLockoutApiGetExpeditionLockoutsBulkRequest, options?: any) {
        return ExpeditionLockoutApiFp(this.configuration).getExpeditionLockoutsBulk(requestParameters.body, options).then((request) => request(this.axios, this.basePath));
    }

    /**
     * 
     * @summary Lists ExpeditionLockouts
     * @param {ExpeditionLockoutApiListExpeditionLockoutsRequest} requestParameters Request parameters.
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof ExpeditionLockoutApi
     */
    public listExpeditionLockouts(requestParameters: ExpeditionLockoutApiListExpeditionLockoutsRequest = {}, options?: any) {
        return ExpeditionLockoutApiFp(this.configuration).listExpeditionLockouts(requestParameters.includes, requestParameters.where, requestParameters.whereOr, requestParameters.groupBy, requestParameters.limit, requestParameters.page, requestParameters.orderBy, requestParameters.orderDirection, requestParameters.select, options).then((request) => request(this.axios, this.basePath));
    }

    /**
     * 
     * @summary Updates ExpeditionLockout
     * @param {ExpeditionLockoutApiUpdateExpeditionLockoutRequest} requestParameters Request parameters.
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof ExpeditionLockoutApi
     */
    public updateExpeditionLockout(requestParameters: ExpeditionLockoutApiUpdateExpeditionLockoutRequest, options?: any) {
        return ExpeditionLockoutApiFp(this.configuration).updateExpeditionLockout(requestParameters.id, requestParameters.expeditionLockout, options).then((request) => request(this.axios, this.basePath));
    }
}
