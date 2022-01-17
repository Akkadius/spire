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
import { ModelsAdventureDetail } from '../models';
/**
 * AdventureDetailApi - axios parameter creator
 * @export
 */
export const AdventureDetailApiAxiosParamCreator = function (configuration?: Configuration) {
    return {
        /**
         * 
         * @summary Creates AdventureDetail
         * @param {ModelsAdventureDetail} adventureDetail AdventureDetail
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        createAdventureDetail: async (adventureDetail: ModelsAdventureDetail, options: any = {}): Promise<RequestArgs> => {
            // verify required parameter 'adventureDetail' is not null or undefined
            if (adventureDetail === null || adventureDetail === undefined) {
                throw new RequiredError('adventureDetail','Required parameter adventureDetail was null or undefined when calling createAdventureDetail.');
            }
            const localVarPath = `/adventure_detail`;
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
            const nonString = typeof adventureDetail !== 'string';
            const needsSerialization = nonString && configuration && configuration.isJsonMime
                ? configuration.isJsonMime(localVarRequestOptions.headers['Content-Type'])
                : nonString;
            localVarRequestOptions.data =  needsSerialization
                ? JSON.stringify(adventureDetail !== undefined ? adventureDetail : {})
                : (adventureDetail || "");

            return {
                url: localVarUrlObj.pathname + localVarUrlObj.search + localVarUrlObj.hash,
                options: localVarRequestOptions,
            };
        },
        /**
         * 
         * @summary Deletes AdventureDetail
         * @param {number} id Id
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        deleteAdventureDetail: async (id: number, options: any = {}): Promise<RequestArgs> => {
            // verify required parameter 'id' is not null or undefined
            if (id === null || id === undefined) {
                throw new RequiredError('id','Required parameter id was null or undefined when calling deleteAdventureDetail.');
            }
            const localVarPath = `/adventure_detail/{id}`
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
         * @summary Gets AdventureDetail
         * @param {number} id Id
         * @param {string} [includes] Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names 
         * @param {string} [select] Column names [.] separated to fetch specific fields in response
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        getAdventureDetail: async (id: number, includes?: string, select?: string, options: any = {}): Promise<RequestArgs> => {
            // verify required parameter 'id' is not null or undefined
            if (id === null || id === undefined) {
                throw new RequiredError('id','Required parameter id was null or undefined when calling getAdventureDetail.');
            }
            const localVarPath = `/adventure_detail/{id}`
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
         * @summary Gets AdventureDetails in bulk
         * @param {CrudcontrollersBulkFetchByIdsGetRequest} body body
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        getAdventureDetailsBulk: async (body: CrudcontrollersBulkFetchByIdsGetRequest, options: any = {}): Promise<RequestArgs> => {
            // verify required parameter 'body' is not null or undefined
            if (body === null || body === undefined) {
                throw new RequiredError('body','Required parameter body was null or undefined when calling getAdventureDetailsBulk.');
            }
            const localVarPath = `/adventure_details/bulk`;
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
         * @summary Lists AdventureDetails
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
        listAdventureDetails: async (includes?: string, where?: string, whereOr?: string, groupBy?: string, limit?: string, orderBy?: string, orderDirection?: string, select?: string, options: any = {}): Promise<RequestArgs> => {
            const localVarPath = `/adventure_details`;
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
         * @summary Updates AdventureDetail
         * @param {number} id Id
         * @param {ModelsAdventureDetail} adventureDetail AdventureDetail
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        updateAdventureDetail: async (id: number, adventureDetail: ModelsAdventureDetail, options: any = {}): Promise<RequestArgs> => {
            // verify required parameter 'id' is not null or undefined
            if (id === null || id === undefined) {
                throw new RequiredError('id','Required parameter id was null or undefined when calling updateAdventureDetail.');
            }
            // verify required parameter 'adventureDetail' is not null or undefined
            if (adventureDetail === null || adventureDetail === undefined) {
                throw new RequiredError('adventureDetail','Required parameter adventureDetail was null or undefined when calling updateAdventureDetail.');
            }
            const localVarPath = `/adventure_detail/{id}`
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
            const nonString = typeof adventureDetail !== 'string';
            const needsSerialization = nonString && configuration && configuration.isJsonMime
                ? configuration.isJsonMime(localVarRequestOptions.headers['Content-Type'])
                : nonString;
            localVarRequestOptions.data =  needsSerialization
                ? JSON.stringify(adventureDetail !== undefined ? adventureDetail : {})
                : (adventureDetail || "");

            return {
                url: localVarUrlObj.pathname + localVarUrlObj.search + localVarUrlObj.hash,
                options: localVarRequestOptions,
            };
        },
    }
};

/**
 * AdventureDetailApi - functional programming interface
 * @export
 */
export const AdventureDetailApiFp = function(configuration?: Configuration) {
    return {
        /**
         * 
         * @summary Creates AdventureDetail
         * @param {ModelsAdventureDetail} adventureDetail AdventureDetail
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async createAdventureDetail(adventureDetail: ModelsAdventureDetail, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsAdventureDetail>>> {
            const localVarAxiosArgs = await AdventureDetailApiAxiosParamCreator(configuration).createAdventureDetail(adventureDetail, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        /**
         * 
         * @summary Deletes AdventureDetail
         * @param {number} id Id
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async deleteAdventureDetail(id: number, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<string>> {
            const localVarAxiosArgs = await AdventureDetailApiAxiosParamCreator(configuration).deleteAdventureDetail(id, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        /**
         * 
         * @summary Gets AdventureDetail
         * @param {number} id Id
         * @param {string} [includes] Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names 
         * @param {string} [select] Column names [.] separated to fetch specific fields in response
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async getAdventureDetail(id: number, includes?: string, select?: string, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsAdventureDetail>>> {
            const localVarAxiosArgs = await AdventureDetailApiAxiosParamCreator(configuration).getAdventureDetail(id, includes, select, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        /**
         * 
         * @summary Gets AdventureDetails in bulk
         * @param {CrudcontrollersBulkFetchByIdsGetRequest} body body
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async getAdventureDetailsBulk(body: CrudcontrollersBulkFetchByIdsGetRequest, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsAdventureDetail>>> {
            const localVarAxiosArgs = await AdventureDetailApiAxiosParamCreator(configuration).getAdventureDetailsBulk(body, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        /**
         * 
         * @summary Lists AdventureDetails
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
        async listAdventureDetails(includes?: string, where?: string, whereOr?: string, groupBy?: string, limit?: string, orderBy?: string, orderDirection?: string, select?: string, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsAdventureDetail>>> {
            const localVarAxiosArgs = await AdventureDetailApiAxiosParamCreator(configuration).listAdventureDetails(includes, where, whereOr, groupBy, limit, orderBy, orderDirection, select, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        /**
         * 
         * @summary Updates AdventureDetail
         * @param {number} id Id
         * @param {ModelsAdventureDetail} adventureDetail AdventureDetail
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async updateAdventureDetail(id: number, adventureDetail: ModelsAdventureDetail, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsAdventureDetail>>> {
            const localVarAxiosArgs = await AdventureDetailApiAxiosParamCreator(configuration).updateAdventureDetail(id, adventureDetail, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
    }
};

/**
 * AdventureDetailApi - factory interface
 * @export
 */
export const AdventureDetailApiFactory = function (configuration?: Configuration, basePath?: string, axios?: AxiosInstance) {
    return {
        /**
         * 
         * @summary Creates AdventureDetail
         * @param {ModelsAdventureDetail} adventureDetail AdventureDetail
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        createAdventureDetail(adventureDetail: ModelsAdventureDetail, options?: any): AxiosPromise<Array<ModelsAdventureDetail>> {
            return AdventureDetailApiFp(configuration).createAdventureDetail(adventureDetail, options).then((request) => request(axios, basePath));
        },
        /**
         * 
         * @summary Deletes AdventureDetail
         * @param {number} id Id
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        deleteAdventureDetail(id: number, options?: any): AxiosPromise<string> {
            return AdventureDetailApiFp(configuration).deleteAdventureDetail(id, options).then((request) => request(axios, basePath));
        },
        /**
         * 
         * @summary Gets AdventureDetail
         * @param {number} id Id
         * @param {string} [includes] Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names 
         * @param {string} [select] Column names [.] separated to fetch specific fields in response
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        getAdventureDetail(id: number, includes?: string, select?: string, options?: any): AxiosPromise<Array<ModelsAdventureDetail>> {
            return AdventureDetailApiFp(configuration).getAdventureDetail(id, includes, select, options).then((request) => request(axios, basePath));
        },
        /**
         * 
         * @summary Gets AdventureDetails in bulk
         * @param {CrudcontrollersBulkFetchByIdsGetRequest} body body
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        getAdventureDetailsBulk(body: CrudcontrollersBulkFetchByIdsGetRequest, options?: any): AxiosPromise<Array<ModelsAdventureDetail>> {
            return AdventureDetailApiFp(configuration).getAdventureDetailsBulk(body, options).then((request) => request(axios, basePath));
        },
        /**
         * 
         * @summary Lists AdventureDetails
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
        listAdventureDetails(includes?: string, where?: string, whereOr?: string, groupBy?: string, limit?: string, orderBy?: string, orderDirection?: string, select?: string, options?: any): AxiosPromise<Array<ModelsAdventureDetail>> {
            return AdventureDetailApiFp(configuration).listAdventureDetails(includes, where, whereOr, groupBy, limit, orderBy, orderDirection, select, options).then((request) => request(axios, basePath));
        },
        /**
         * 
         * @summary Updates AdventureDetail
         * @param {number} id Id
         * @param {ModelsAdventureDetail} adventureDetail AdventureDetail
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        updateAdventureDetail(id: number, adventureDetail: ModelsAdventureDetail, options?: any): AxiosPromise<Array<ModelsAdventureDetail>> {
            return AdventureDetailApiFp(configuration).updateAdventureDetail(id, adventureDetail, options).then((request) => request(axios, basePath));
        },
    };
};

/**
 * Request parameters for createAdventureDetail operation in AdventureDetailApi.
 * @export
 * @interface AdventureDetailApiCreateAdventureDetailRequest
 */
export interface AdventureDetailApiCreateAdventureDetailRequest {
    /**
     * AdventureDetail
     * @type {ModelsAdventureDetail}
     * @memberof AdventureDetailApiCreateAdventureDetail
     */
    readonly adventureDetail: ModelsAdventureDetail
}

/**
 * Request parameters for deleteAdventureDetail operation in AdventureDetailApi.
 * @export
 * @interface AdventureDetailApiDeleteAdventureDetailRequest
 */
export interface AdventureDetailApiDeleteAdventureDetailRequest {
    /**
     * Id
     * @type {number}
     * @memberof AdventureDetailApiDeleteAdventureDetail
     */
    readonly id: number
}

/**
 * Request parameters for getAdventureDetail operation in AdventureDetailApi.
 * @export
 * @interface AdventureDetailApiGetAdventureDetailRequest
 */
export interface AdventureDetailApiGetAdventureDetailRequest {
    /**
     * Id
     * @type {number}
     * @memberof AdventureDetailApiGetAdventureDetail
     */
    readonly id: number

    /**
     * Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names 
     * @type {string}
     * @memberof AdventureDetailApiGetAdventureDetail
     */
    readonly includes?: string

    /**
     * Column names [.] separated to fetch specific fields in response
     * @type {string}
     * @memberof AdventureDetailApiGetAdventureDetail
     */
    readonly select?: string
}

/**
 * Request parameters for getAdventureDetailsBulk operation in AdventureDetailApi.
 * @export
 * @interface AdventureDetailApiGetAdventureDetailsBulkRequest
 */
export interface AdventureDetailApiGetAdventureDetailsBulkRequest {
    /**
     * body
     * @type {CrudcontrollersBulkFetchByIdsGetRequest}
     * @memberof AdventureDetailApiGetAdventureDetailsBulk
     */
    readonly body: CrudcontrollersBulkFetchByIdsGetRequest
}

/**
 * Request parameters for listAdventureDetails operation in AdventureDetailApi.
 * @export
 * @interface AdventureDetailApiListAdventureDetailsRequest
 */
export interface AdventureDetailApiListAdventureDetailsRequest {
    /**
     * Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names 
     * @type {string}
     * @memberof AdventureDetailApiListAdventureDetails
     */
    readonly includes?: string

    /**
     * Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2
     * @type {string}
     * @memberof AdventureDetailApiListAdventureDetails
     */
    readonly where?: string

    /**
     * Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2
     * @type {string}
     * @memberof AdventureDetailApiListAdventureDetails
     */
    readonly whereOr?: string

    /**
     * Group by field. Multiple conditions [.] separated Example: field1.field2
     * @type {string}
     * @memberof AdventureDetailApiListAdventureDetails
     */
    readonly groupBy?: string

    /**
     * Rows to limit in response (Default: 10,000)
     * @type {string}
     * @memberof AdventureDetailApiListAdventureDetails
     */
    readonly limit?: string

    /**
     * Order by [field]
     * @type {string}
     * @memberof AdventureDetailApiListAdventureDetails
     */
    readonly orderBy?: string

    /**
     * Order by field direction
     * @type {string}
     * @memberof AdventureDetailApiListAdventureDetails
     */
    readonly orderDirection?: string

    /**
     * Column names [.] separated to fetch specific fields in response
     * @type {string}
     * @memberof AdventureDetailApiListAdventureDetails
     */
    readonly select?: string
}

/**
 * Request parameters for updateAdventureDetail operation in AdventureDetailApi.
 * @export
 * @interface AdventureDetailApiUpdateAdventureDetailRequest
 */
export interface AdventureDetailApiUpdateAdventureDetailRequest {
    /**
     * Id
     * @type {number}
     * @memberof AdventureDetailApiUpdateAdventureDetail
     */
    readonly id: number

    /**
     * AdventureDetail
     * @type {ModelsAdventureDetail}
     * @memberof AdventureDetailApiUpdateAdventureDetail
     */
    readonly adventureDetail: ModelsAdventureDetail
}

/**
 * AdventureDetailApi - object-oriented interface
 * @export
 * @class AdventureDetailApi
 * @extends {BaseAPI}
 */
export class AdventureDetailApi extends BaseAPI {
    /**
     * 
     * @summary Creates AdventureDetail
     * @param {AdventureDetailApiCreateAdventureDetailRequest} requestParameters Request parameters.
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof AdventureDetailApi
     */
    public createAdventureDetail(requestParameters: AdventureDetailApiCreateAdventureDetailRequest, options?: any) {
        return AdventureDetailApiFp(this.configuration).createAdventureDetail(requestParameters.adventureDetail, options).then((request) => request(this.axios, this.basePath));
    }

    /**
     * 
     * @summary Deletes AdventureDetail
     * @param {AdventureDetailApiDeleteAdventureDetailRequest} requestParameters Request parameters.
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof AdventureDetailApi
     */
    public deleteAdventureDetail(requestParameters: AdventureDetailApiDeleteAdventureDetailRequest, options?: any) {
        return AdventureDetailApiFp(this.configuration).deleteAdventureDetail(requestParameters.id, options).then((request) => request(this.axios, this.basePath));
    }

    /**
     * 
     * @summary Gets AdventureDetail
     * @param {AdventureDetailApiGetAdventureDetailRequest} requestParameters Request parameters.
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof AdventureDetailApi
     */
    public getAdventureDetail(requestParameters: AdventureDetailApiGetAdventureDetailRequest, options?: any) {
        return AdventureDetailApiFp(this.configuration).getAdventureDetail(requestParameters.id, requestParameters.includes, requestParameters.select, options).then((request) => request(this.axios, this.basePath));
    }

    /**
     * 
     * @summary Gets AdventureDetails in bulk
     * @param {AdventureDetailApiGetAdventureDetailsBulkRequest} requestParameters Request parameters.
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof AdventureDetailApi
     */
    public getAdventureDetailsBulk(requestParameters: AdventureDetailApiGetAdventureDetailsBulkRequest, options?: any) {
        return AdventureDetailApiFp(this.configuration).getAdventureDetailsBulk(requestParameters.body, options).then((request) => request(this.axios, this.basePath));
    }

    /**
     * 
     * @summary Lists AdventureDetails
     * @param {AdventureDetailApiListAdventureDetailsRequest} requestParameters Request parameters.
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof AdventureDetailApi
     */
    public listAdventureDetails(requestParameters: AdventureDetailApiListAdventureDetailsRequest = {}, options?: any) {
        return AdventureDetailApiFp(this.configuration).listAdventureDetails(requestParameters.includes, requestParameters.where, requestParameters.whereOr, requestParameters.groupBy, requestParameters.limit, requestParameters.orderBy, requestParameters.orderDirection, requestParameters.select, options).then((request) => request(this.axios, this.basePath));
    }

    /**
     * 
     * @summary Updates AdventureDetail
     * @param {AdventureDetailApiUpdateAdventureDetailRequest} requestParameters Request parameters.
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof AdventureDetailApi
     */
    public updateAdventureDetail(requestParameters: AdventureDetailApiUpdateAdventureDetailRequest, options?: any) {
        return AdventureDetailApiFp(this.configuration).updateAdventureDetail(requestParameters.id, requestParameters.adventureDetail, options).then((request) => request(this.axios, this.basePath));
    }
}
