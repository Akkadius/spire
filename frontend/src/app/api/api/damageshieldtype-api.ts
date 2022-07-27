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
import { ModelsDamageshieldtype } from '../models';
/**
 * DamageshieldtypeApi - axios parameter creator
 * @export
 */
export const DamageshieldtypeApiAxiosParamCreator = function (configuration?: Configuration) {
    return {
        /**
         * 
         * @summary Creates Damageshieldtype
         * @param {ModelsDamageshieldtype} damageshieldtype Damageshieldtype
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        createDamageshieldtype: async (damageshieldtype: ModelsDamageshieldtype, options: any = {}): Promise<RequestArgs> => {
            // verify required parameter 'damageshieldtype' is not null or undefined
            if (damageshieldtype === null || damageshieldtype === undefined) {
                throw new RequiredError('damageshieldtype','Required parameter damageshieldtype was null or undefined when calling createDamageshieldtype.');
            }
            const localVarPath = `/damageshieldtype`;
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
            const nonString = typeof damageshieldtype !== 'string';
            const needsSerialization = nonString && configuration && configuration.isJsonMime
                ? configuration.isJsonMime(localVarRequestOptions.headers['Content-Type'])
                : nonString;
            localVarRequestOptions.data =  needsSerialization
                ? JSON.stringify(damageshieldtype !== undefined ? damageshieldtype : {})
                : (damageshieldtype || "");

            return {
                url: localVarUrlObj.pathname + localVarUrlObj.search + localVarUrlObj.hash,
                options: localVarRequestOptions,
            };
        },
        /**
         * 
         * @summary Deletes Damageshieldtype
         * @param {number} id spellid
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        deleteDamageshieldtype: async (id: number, options: any = {}): Promise<RequestArgs> => {
            // verify required parameter 'id' is not null or undefined
            if (id === null || id === undefined) {
                throw new RequiredError('id','Required parameter id was null or undefined when calling deleteDamageshieldtype.');
            }
            const localVarPath = `/damageshieldtype/{id}`
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
         * @summary Gets Damageshieldtype
         * @param {number} id Id
         * @param {string} [includes] Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names 
         * @param {string} [select] Column names [.] separated to fetch specific fields in response
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        getDamageshieldtype: async (id: number, includes?: string, select?: string, options: any = {}): Promise<RequestArgs> => {
            // verify required parameter 'id' is not null or undefined
            if (id === null || id === undefined) {
                throw new RequiredError('id','Required parameter id was null or undefined when calling getDamageshieldtype.');
            }
            const localVarPath = `/damageshieldtype/{id}`
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
         * @summary Gets Damageshieldtypes in bulk
         * @param {CrudcontrollersBulkFetchByIdsGetRequest} body body
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        getDamageshieldtypesBulk: async (body: CrudcontrollersBulkFetchByIdsGetRequest, options: any = {}): Promise<RequestArgs> => {
            // verify required parameter 'body' is not null or undefined
            if (body === null || body === undefined) {
                throw new RequiredError('body','Required parameter body was null or undefined when calling getDamageshieldtypesBulk.');
            }
            const localVarPath = `/damageshieldtypes/bulk`;
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
         * @summary Lists Damageshieldtypes
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
        listDamageshieldtypes: async (includes?: string, where?: string, whereOr?: string, groupBy?: string, limit?: string, page?: number, orderBy?: string, orderDirection?: string, select?: string, options: any = {}): Promise<RequestArgs> => {
            const localVarPath = `/damageshieldtypes`;
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
         * @summary Updates Damageshieldtype
         * @param {number} id Id
         * @param {ModelsDamageshieldtype} damageshieldtype Damageshieldtype
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        updateDamageshieldtype: async (id: number, damageshieldtype: ModelsDamageshieldtype, options: any = {}): Promise<RequestArgs> => {
            // verify required parameter 'id' is not null or undefined
            if (id === null || id === undefined) {
                throw new RequiredError('id','Required parameter id was null or undefined when calling updateDamageshieldtype.');
            }
            // verify required parameter 'damageshieldtype' is not null or undefined
            if (damageshieldtype === null || damageshieldtype === undefined) {
                throw new RequiredError('damageshieldtype','Required parameter damageshieldtype was null or undefined when calling updateDamageshieldtype.');
            }
            const localVarPath = `/damageshieldtype/{id}`
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
            const nonString = typeof damageshieldtype !== 'string';
            const needsSerialization = nonString && configuration && configuration.isJsonMime
                ? configuration.isJsonMime(localVarRequestOptions.headers['Content-Type'])
                : nonString;
            localVarRequestOptions.data =  needsSerialization
                ? JSON.stringify(damageshieldtype !== undefined ? damageshieldtype : {})
                : (damageshieldtype || "");

            return {
                url: localVarUrlObj.pathname + localVarUrlObj.search + localVarUrlObj.hash,
                options: localVarRequestOptions,
            };
        },
    }
};

/**
 * DamageshieldtypeApi - functional programming interface
 * @export
 */
export const DamageshieldtypeApiFp = function(configuration?: Configuration) {
    return {
        /**
         * 
         * @summary Creates Damageshieldtype
         * @param {ModelsDamageshieldtype} damageshieldtype Damageshieldtype
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async createDamageshieldtype(damageshieldtype: ModelsDamageshieldtype, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsDamageshieldtype>>> {
            const localVarAxiosArgs = await DamageshieldtypeApiAxiosParamCreator(configuration).createDamageshieldtype(damageshieldtype, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        /**
         * 
         * @summary Deletes Damageshieldtype
         * @param {number} id spellid
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async deleteDamageshieldtype(id: number, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<string>> {
            const localVarAxiosArgs = await DamageshieldtypeApiAxiosParamCreator(configuration).deleteDamageshieldtype(id, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        /**
         * 
         * @summary Gets Damageshieldtype
         * @param {number} id Id
         * @param {string} [includes] Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names 
         * @param {string} [select] Column names [.] separated to fetch specific fields in response
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async getDamageshieldtype(id: number, includes?: string, select?: string, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsDamageshieldtype>>> {
            const localVarAxiosArgs = await DamageshieldtypeApiAxiosParamCreator(configuration).getDamageshieldtype(id, includes, select, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        /**
         * 
         * @summary Gets Damageshieldtypes in bulk
         * @param {CrudcontrollersBulkFetchByIdsGetRequest} body body
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async getDamageshieldtypesBulk(body: CrudcontrollersBulkFetchByIdsGetRequest, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsDamageshieldtype>>> {
            const localVarAxiosArgs = await DamageshieldtypeApiAxiosParamCreator(configuration).getDamageshieldtypesBulk(body, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        /**
         * 
         * @summary Lists Damageshieldtypes
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
        async listDamageshieldtypes(includes?: string, where?: string, whereOr?: string, groupBy?: string, limit?: string, page?: number, orderBy?: string, orderDirection?: string, select?: string, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsDamageshieldtype>>> {
            const localVarAxiosArgs = await DamageshieldtypeApiAxiosParamCreator(configuration).listDamageshieldtypes(includes, where, whereOr, groupBy, limit, page, orderBy, orderDirection, select, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        /**
         * 
         * @summary Updates Damageshieldtype
         * @param {number} id Id
         * @param {ModelsDamageshieldtype} damageshieldtype Damageshieldtype
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async updateDamageshieldtype(id: number, damageshieldtype: ModelsDamageshieldtype, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsDamageshieldtype>>> {
            const localVarAxiosArgs = await DamageshieldtypeApiAxiosParamCreator(configuration).updateDamageshieldtype(id, damageshieldtype, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
    }
};

/**
 * DamageshieldtypeApi - factory interface
 * @export
 */
export const DamageshieldtypeApiFactory = function (configuration?: Configuration, basePath?: string, axios?: AxiosInstance) {
    return {
        /**
         * 
         * @summary Creates Damageshieldtype
         * @param {ModelsDamageshieldtype} damageshieldtype Damageshieldtype
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        createDamageshieldtype(damageshieldtype: ModelsDamageshieldtype, options?: any): AxiosPromise<Array<ModelsDamageshieldtype>> {
            return DamageshieldtypeApiFp(configuration).createDamageshieldtype(damageshieldtype, options).then((request) => request(axios, basePath));
        },
        /**
         * 
         * @summary Deletes Damageshieldtype
         * @param {number} id spellid
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        deleteDamageshieldtype(id: number, options?: any): AxiosPromise<string> {
            return DamageshieldtypeApiFp(configuration).deleteDamageshieldtype(id, options).then((request) => request(axios, basePath));
        },
        /**
         * 
         * @summary Gets Damageshieldtype
         * @param {number} id Id
         * @param {string} [includes] Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names 
         * @param {string} [select] Column names [.] separated to fetch specific fields in response
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        getDamageshieldtype(id: number, includes?: string, select?: string, options?: any): AxiosPromise<Array<ModelsDamageshieldtype>> {
            return DamageshieldtypeApiFp(configuration).getDamageshieldtype(id, includes, select, options).then((request) => request(axios, basePath));
        },
        /**
         * 
         * @summary Gets Damageshieldtypes in bulk
         * @param {CrudcontrollersBulkFetchByIdsGetRequest} body body
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        getDamageshieldtypesBulk(body: CrudcontrollersBulkFetchByIdsGetRequest, options?: any): AxiosPromise<Array<ModelsDamageshieldtype>> {
            return DamageshieldtypeApiFp(configuration).getDamageshieldtypesBulk(body, options).then((request) => request(axios, basePath));
        },
        /**
         * 
         * @summary Lists Damageshieldtypes
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
        listDamageshieldtypes(includes?: string, where?: string, whereOr?: string, groupBy?: string, limit?: string, page?: number, orderBy?: string, orderDirection?: string, select?: string, options?: any): AxiosPromise<Array<ModelsDamageshieldtype>> {
            return DamageshieldtypeApiFp(configuration).listDamageshieldtypes(includes, where, whereOr, groupBy, limit, page, orderBy, orderDirection, select, options).then((request) => request(axios, basePath));
        },
        /**
         * 
         * @summary Updates Damageshieldtype
         * @param {number} id Id
         * @param {ModelsDamageshieldtype} damageshieldtype Damageshieldtype
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        updateDamageshieldtype(id: number, damageshieldtype: ModelsDamageshieldtype, options?: any): AxiosPromise<Array<ModelsDamageshieldtype>> {
            return DamageshieldtypeApiFp(configuration).updateDamageshieldtype(id, damageshieldtype, options).then((request) => request(axios, basePath));
        },
    };
};

/**
 * Request parameters for createDamageshieldtype operation in DamageshieldtypeApi.
 * @export
 * @interface DamageshieldtypeApiCreateDamageshieldtypeRequest
 */
export interface DamageshieldtypeApiCreateDamageshieldtypeRequest {
    /**
     * Damageshieldtype
     * @type {ModelsDamageshieldtype}
     * @memberof DamageshieldtypeApiCreateDamageshieldtype
     */
    readonly damageshieldtype: ModelsDamageshieldtype
}

/**
 * Request parameters for deleteDamageshieldtype operation in DamageshieldtypeApi.
 * @export
 * @interface DamageshieldtypeApiDeleteDamageshieldtypeRequest
 */
export interface DamageshieldtypeApiDeleteDamageshieldtypeRequest {
    /**
     * spellid
     * @type {number}
     * @memberof DamageshieldtypeApiDeleteDamageshieldtype
     */
    readonly id: number
}

/**
 * Request parameters for getDamageshieldtype operation in DamageshieldtypeApi.
 * @export
 * @interface DamageshieldtypeApiGetDamageshieldtypeRequest
 */
export interface DamageshieldtypeApiGetDamageshieldtypeRequest {
    /**
     * Id
     * @type {number}
     * @memberof DamageshieldtypeApiGetDamageshieldtype
     */
    readonly id: number

    /**
     * Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names 
     * @type {string}
     * @memberof DamageshieldtypeApiGetDamageshieldtype
     */
    readonly includes?: string

    /**
     * Column names [.] separated to fetch specific fields in response
     * @type {string}
     * @memberof DamageshieldtypeApiGetDamageshieldtype
     */
    readonly select?: string
}

/**
 * Request parameters for getDamageshieldtypesBulk operation in DamageshieldtypeApi.
 * @export
 * @interface DamageshieldtypeApiGetDamageshieldtypesBulkRequest
 */
export interface DamageshieldtypeApiGetDamageshieldtypesBulkRequest {
    /**
     * body
     * @type {CrudcontrollersBulkFetchByIdsGetRequest}
     * @memberof DamageshieldtypeApiGetDamageshieldtypesBulk
     */
    readonly body: CrudcontrollersBulkFetchByIdsGetRequest
}

/**
 * Request parameters for listDamageshieldtypes operation in DamageshieldtypeApi.
 * @export
 * @interface DamageshieldtypeApiListDamageshieldtypesRequest
 */
export interface DamageshieldtypeApiListDamageshieldtypesRequest {
    /**
     * Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names 
     * @type {string}
     * @memberof DamageshieldtypeApiListDamageshieldtypes
     */
    readonly includes?: string

    /**
     * Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2
     * @type {string}
     * @memberof DamageshieldtypeApiListDamageshieldtypes
     */
    readonly where?: string

    /**
     * Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2
     * @type {string}
     * @memberof DamageshieldtypeApiListDamageshieldtypes
     */
    readonly whereOr?: string

    /**
     * Group by field. Multiple conditions [.] separated Example: field1.field2
     * @type {string}
     * @memberof DamageshieldtypeApiListDamageshieldtypes
     */
    readonly groupBy?: string

    /**
     * Rows to limit in response (Default: 10,000)
     * @type {string}
     * @memberof DamageshieldtypeApiListDamageshieldtypes
     */
    readonly limit?: string

    /**
     * Pagination page
     * @type {number}
     * @memberof DamageshieldtypeApiListDamageshieldtypes
     */
    readonly page?: number

    /**
     * Order by [field]
     * @type {string}
     * @memberof DamageshieldtypeApiListDamageshieldtypes
     */
    readonly orderBy?: string

    /**
     * Order by field direction
     * @type {string}
     * @memberof DamageshieldtypeApiListDamageshieldtypes
     */
    readonly orderDirection?: string

    /**
     * Column names [.] separated to fetch specific fields in response
     * @type {string}
     * @memberof DamageshieldtypeApiListDamageshieldtypes
     */
    readonly select?: string
}

/**
 * Request parameters for updateDamageshieldtype operation in DamageshieldtypeApi.
 * @export
 * @interface DamageshieldtypeApiUpdateDamageshieldtypeRequest
 */
export interface DamageshieldtypeApiUpdateDamageshieldtypeRequest {
    /**
     * Id
     * @type {number}
     * @memberof DamageshieldtypeApiUpdateDamageshieldtype
     */
    readonly id: number

    /**
     * Damageshieldtype
     * @type {ModelsDamageshieldtype}
     * @memberof DamageshieldtypeApiUpdateDamageshieldtype
     */
    readonly damageshieldtype: ModelsDamageshieldtype
}

/**
 * DamageshieldtypeApi - object-oriented interface
 * @export
 * @class DamageshieldtypeApi
 * @extends {BaseAPI}
 */
export class DamageshieldtypeApi extends BaseAPI {
    /**
     * 
     * @summary Creates Damageshieldtype
     * @param {DamageshieldtypeApiCreateDamageshieldtypeRequest} requestParameters Request parameters.
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof DamageshieldtypeApi
     */
    public createDamageshieldtype(requestParameters: DamageshieldtypeApiCreateDamageshieldtypeRequest, options?: any) {
        return DamageshieldtypeApiFp(this.configuration).createDamageshieldtype(requestParameters.damageshieldtype, options).then((request) => request(this.axios, this.basePath));
    }

    /**
     * 
     * @summary Deletes Damageshieldtype
     * @param {DamageshieldtypeApiDeleteDamageshieldtypeRequest} requestParameters Request parameters.
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof DamageshieldtypeApi
     */
    public deleteDamageshieldtype(requestParameters: DamageshieldtypeApiDeleteDamageshieldtypeRequest, options?: any) {
        return DamageshieldtypeApiFp(this.configuration).deleteDamageshieldtype(requestParameters.id, options).then((request) => request(this.axios, this.basePath));
    }

    /**
     * 
     * @summary Gets Damageshieldtype
     * @param {DamageshieldtypeApiGetDamageshieldtypeRequest} requestParameters Request parameters.
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof DamageshieldtypeApi
     */
    public getDamageshieldtype(requestParameters: DamageshieldtypeApiGetDamageshieldtypeRequest, options?: any) {
        return DamageshieldtypeApiFp(this.configuration).getDamageshieldtype(requestParameters.id, requestParameters.includes, requestParameters.select, options).then((request) => request(this.axios, this.basePath));
    }

    /**
     * 
     * @summary Gets Damageshieldtypes in bulk
     * @param {DamageshieldtypeApiGetDamageshieldtypesBulkRequest} requestParameters Request parameters.
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof DamageshieldtypeApi
     */
    public getDamageshieldtypesBulk(requestParameters: DamageshieldtypeApiGetDamageshieldtypesBulkRequest, options?: any) {
        return DamageshieldtypeApiFp(this.configuration).getDamageshieldtypesBulk(requestParameters.body, options).then((request) => request(this.axios, this.basePath));
    }

    /**
     * 
     * @summary Lists Damageshieldtypes
     * @param {DamageshieldtypeApiListDamageshieldtypesRequest} requestParameters Request parameters.
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof DamageshieldtypeApi
     */
    public listDamageshieldtypes(requestParameters: DamageshieldtypeApiListDamageshieldtypesRequest = {}, options?: any) {
        return DamageshieldtypeApiFp(this.configuration).listDamageshieldtypes(requestParameters.includes, requestParameters.where, requestParameters.whereOr, requestParameters.groupBy, requestParameters.limit, requestParameters.page, requestParameters.orderBy, requestParameters.orderDirection, requestParameters.select, options).then((request) => request(this.axios, this.basePath));
    }

    /**
     * 
     * @summary Updates Damageshieldtype
     * @param {DamageshieldtypeApiUpdateDamageshieldtypeRequest} requestParameters Request parameters.
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof DamageshieldtypeApi
     */
    public updateDamageshieldtype(requestParameters: DamageshieldtypeApiUpdateDamageshieldtypeRequest, options?: any) {
        return DamageshieldtypeApiFp(this.configuration).updateDamageshieldtype(requestParameters.id, requestParameters.damageshieldtype, options).then((request) => request(this.axios, this.basePath));
    }
}
