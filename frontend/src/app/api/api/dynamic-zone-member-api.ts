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
import { ModelsDynamicZoneMember } from '../models';
/**
 * DynamicZoneMemberApi - axios parameter creator
 * @export
 */
export const DynamicZoneMemberApiAxiosParamCreator = function (configuration?: Configuration) {
    return {
        /**
         * 
         * @summary Creates DynamicZoneMember
         * @param {ModelsDynamicZoneMember} dynamicZoneMember DynamicZoneMember
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        createDynamicZoneMember: async (dynamicZoneMember: ModelsDynamicZoneMember, options: any = {}): Promise<RequestArgs> => {
            // verify required parameter 'dynamicZoneMember' is not null or undefined
            if (dynamicZoneMember === null || dynamicZoneMember === undefined) {
                throw new RequiredError('dynamicZoneMember','Required parameter dynamicZoneMember was null or undefined when calling createDynamicZoneMember.');
            }
            const localVarPath = `/dynamic_zone_member`;
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
            const nonString = typeof dynamicZoneMember !== 'string';
            const needsSerialization = nonString && configuration && configuration.isJsonMime
                ? configuration.isJsonMime(localVarRequestOptions.headers['Content-Type'])
                : nonString;
            localVarRequestOptions.data =  needsSerialization
                ? JSON.stringify(dynamicZoneMember !== undefined ? dynamicZoneMember : {})
                : (dynamicZoneMember || "");

            return {
                url: localVarUrlObj.pathname + localVarUrlObj.search + localVarUrlObj.hash,
                options: localVarRequestOptions,
            };
        },
        /**
         * 
         * @summary Deletes DynamicZoneMember
         * @param {number} id id
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        deleteDynamicZoneMember: async (id: number, options: any = {}): Promise<RequestArgs> => {
            // verify required parameter 'id' is not null or undefined
            if (id === null || id === undefined) {
                throw new RequiredError('id','Required parameter id was null or undefined when calling deleteDynamicZoneMember.');
            }
            const localVarPath = `/dynamic_zone_member/{id}`
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
         * @summary Gets DynamicZoneMember
         * @param {number} id Id
         * @param {string} [includes] Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names 
         * @param {string} [select] Column names [.] separated to fetch specific fields in response
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        getDynamicZoneMember: async (id: number, includes?: string, select?: string, options: any = {}): Promise<RequestArgs> => {
            // verify required parameter 'id' is not null or undefined
            if (id === null || id === undefined) {
                throw new RequiredError('id','Required parameter id was null or undefined when calling getDynamicZoneMember.');
            }
            const localVarPath = `/dynamic_zone_member/{id}`
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
         * @summary Gets DynamicZoneMembers in bulk
         * @param {CrudcontrollersBulkFetchByIdsGetRequest} body body
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        getDynamicZoneMembersBulk: async (body: CrudcontrollersBulkFetchByIdsGetRequest, options: any = {}): Promise<RequestArgs> => {
            // verify required parameter 'body' is not null or undefined
            if (body === null || body === undefined) {
                throw new RequiredError('body','Required parameter body was null or undefined when calling getDynamicZoneMembersBulk.');
            }
            const localVarPath = `/dynamic_zone_members/bulk`;
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
         * @summary Lists DynamicZoneMembers
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
        listDynamicZoneMembers: async (includes?: string, where?: string, whereOr?: string, groupBy?: string, limit?: string, page?: number, orderBy?: string, orderDirection?: string, select?: string, options: any = {}): Promise<RequestArgs> => {
            const localVarPath = `/dynamic_zone_members`;
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
         * @summary Updates DynamicZoneMember
         * @param {number} id Id
         * @param {ModelsDynamicZoneMember} dynamicZoneMember DynamicZoneMember
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        updateDynamicZoneMember: async (id: number, dynamicZoneMember: ModelsDynamicZoneMember, options: any = {}): Promise<RequestArgs> => {
            // verify required parameter 'id' is not null or undefined
            if (id === null || id === undefined) {
                throw new RequiredError('id','Required parameter id was null or undefined when calling updateDynamicZoneMember.');
            }
            // verify required parameter 'dynamicZoneMember' is not null or undefined
            if (dynamicZoneMember === null || dynamicZoneMember === undefined) {
                throw new RequiredError('dynamicZoneMember','Required parameter dynamicZoneMember was null or undefined when calling updateDynamicZoneMember.');
            }
            const localVarPath = `/dynamic_zone_member/{id}`
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
            const nonString = typeof dynamicZoneMember !== 'string';
            const needsSerialization = nonString && configuration && configuration.isJsonMime
                ? configuration.isJsonMime(localVarRequestOptions.headers['Content-Type'])
                : nonString;
            localVarRequestOptions.data =  needsSerialization
                ? JSON.stringify(dynamicZoneMember !== undefined ? dynamicZoneMember : {})
                : (dynamicZoneMember || "");

            return {
                url: localVarUrlObj.pathname + localVarUrlObj.search + localVarUrlObj.hash,
                options: localVarRequestOptions,
            };
        },
    }
};

/**
 * DynamicZoneMemberApi - functional programming interface
 * @export
 */
export const DynamicZoneMemberApiFp = function(configuration?: Configuration) {
    return {
        /**
         * 
         * @summary Creates DynamicZoneMember
         * @param {ModelsDynamicZoneMember} dynamicZoneMember DynamicZoneMember
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async createDynamicZoneMember(dynamicZoneMember: ModelsDynamicZoneMember, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsDynamicZoneMember>>> {
            const localVarAxiosArgs = await DynamicZoneMemberApiAxiosParamCreator(configuration).createDynamicZoneMember(dynamicZoneMember, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        /**
         * 
         * @summary Deletes DynamicZoneMember
         * @param {number} id id
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async deleteDynamicZoneMember(id: number, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<string>> {
            const localVarAxiosArgs = await DynamicZoneMemberApiAxiosParamCreator(configuration).deleteDynamicZoneMember(id, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        /**
         * 
         * @summary Gets DynamicZoneMember
         * @param {number} id Id
         * @param {string} [includes] Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names 
         * @param {string} [select] Column names [.] separated to fetch specific fields in response
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async getDynamicZoneMember(id: number, includes?: string, select?: string, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsDynamicZoneMember>>> {
            const localVarAxiosArgs = await DynamicZoneMemberApiAxiosParamCreator(configuration).getDynamicZoneMember(id, includes, select, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        /**
         * 
         * @summary Gets DynamicZoneMembers in bulk
         * @param {CrudcontrollersBulkFetchByIdsGetRequest} body body
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async getDynamicZoneMembersBulk(body: CrudcontrollersBulkFetchByIdsGetRequest, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsDynamicZoneMember>>> {
            const localVarAxiosArgs = await DynamicZoneMemberApiAxiosParamCreator(configuration).getDynamicZoneMembersBulk(body, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        /**
         * 
         * @summary Lists DynamicZoneMembers
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
        async listDynamicZoneMembers(includes?: string, where?: string, whereOr?: string, groupBy?: string, limit?: string, page?: number, orderBy?: string, orderDirection?: string, select?: string, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsDynamicZoneMember>>> {
            const localVarAxiosArgs = await DynamicZoneMemberApiAxiosParamCreator(configuration).listDynamicZoneMembers(includes, where, whereOr, groupBy, limit, page, orderBy, orderDirection, select, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        /**
         * 
         * @summary Updates DynamicZoneMember
         * @param {number} id Id
         * @param {ModelsDynamicZoneMember} dynamicZoneMember DynamicZoneMember
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async updateDynamicZoneMember(id: number, dynamicZoneMember: ModelsDynamicZoneMember, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsDynamicZoneMember>>> {
            const localVarAxiosArgs = await DynamicZoneMemberApiAxiosParamCreator(configuration).updateDynamicZoneMember(id, dynamicZoneMember, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
    }
};

/**
 * DynamicZoneMemberApi - factory interface
 * @export
 */
export const DynamicZoneMemberApiFactory = function (configuration?: Configuration, basePath?: string, axios?: AxiosInstance) {
    return {
        /**
         * 
         * @summary Creates DynamicZoneMember
         * @param {ModelsDynamicZoneMember} dynamicZoneMember DynamicZoneMember
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        createDynamicZoneMember(dynamicZoneMember: ModelsDynamicZoneMember, options?: any): AxiosPromise<Array<ModelsDynamicZoneMember>> {
            return DynamicZoneMemberApiFp(configuration).createDynamicZoneMember(dynamicZoneMember, options).then((request) => request(axios, basePath));
        },
        /**
         * 
         * @summary Deletes DynamicZoneMember
         * @param {number} id id
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        deleteDynamicZoneMember(id: number, options?: any): AxiosPromise<string> {
            return DynamicZoneMemberApiFp(configuration).deleteDynamicZoneMember(id, options).then((request) => request(axios, basePath));
        },
        /**
         * 
         * @summary Gets DynamicZoneMember
         * @param {number} id Id
         * @param {string} [includes] Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names 
         * @param {string} [select] Column names [.] separated to fetch specific fields in response
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        getDynamicZoneMember(id: number, includes?: string, select?: string, options?: any): AxiosPromise<Array<ModelsDynamicZoneMember>> {
            return DynamicZoneMemberApiFp(configuration).getDynamicZoneMember(id, includes, select, options).then((request) => request(axios, basePath));
        },
        /**
         * 
         * @summary Gets DynamicZoneMembers in bulk
         * @param {CrudcontrollersBulkFetchByIdsGetRequest} body body
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        getDynamicZoneMembersBulk(body: CrudcontrollersBulkFetchByIdsGetRequest, options?: any): AxiosPromise<Array<ModelsDynamicZoneMember>> {
            return DynamicZoneMemberApiFp(configuration).getDynamicZoneMembersBulk(body, options).then((request) => request(axios, basePath));
        },
        /**
         * 
         * @summary Lists DynamicZoneMembers
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
        listDynamicZoneMembers(includes?: string, where?: string, whereOr?: string, groupBy?: string, limit?: string, page?: number, orderBy?: string, orderDirection?: string, select?: string, options?: any): AxiosPromise<Array<ModelsDynamicZoneMember>> {
            return DynamicZoneMemberApiFp(configuration).listDynamicZoneMembers(includes, where, whereOr, groupBy, limit, page, orderBy, orderDirection, select, options).then((request) => request(axios, basePath));
        },
        /**
         * 
         * @summary Updates DynamicZoneMember
         * @param {number} id Id
         * @param {ModelsDynamicZoneMember} dynamicZoneMember DynamicZoneMember
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        updateDynamicZoneMember(id: number, dynamicZoneMember: ModelsDynamicZoneMember, options?: any): AxiosPromise<Array<ModelsDynamicZoneMember>> {
            return DynamicZoneMemberApiFp(configuration).updateDynamicZoneMember(id, dynamicZoneMember, options).then((request) => request(axios, basePath));
        },
    };
};

/**
 * Request parameters for createDynamicZoneMember operation in DynamicZoneMemberApi.
 * @export
 * @interface DynamicZoneMemberApiCreateDynamicZoneMemberRequest
 */
export interface DynamicZoneMemberApiCreateDynamicZoneMemberRequest {
    /**
     * DynamicZoneMember
     * @type {ModelsDynamicZoneMember}
     * @memberof DynamicZoneMemberApiCreateDynamicZoneMember
     */
    readonly dynamicZoneMember: ModelsDynamicZoneMember
}

/**
 * Request parameters for deleteDynamicZoneMember operation in DynamicZoneMemberApi.
 * @export
 * @interface DynamicZoneMemberApiDeleteDynamicZoneMemberRequest
 */
export interface DynamicZoneMemberApiDeleteDynamicZoneMemberRequest {
    /**
     * id
     * @type {number}
     * @memberof DynamicZoneMemberApiDeleteDynamicZoneMember
     */
    readonly id: number
}

/**
 * Request parameters for getDynamicZoneMember operation in DynamicZoneMemberApi.
 * @export
 * @interface DynamicZoneMemberApiGetDynamicZoneMemberRequest
 */
export interface DynamicZoneMemberApiGetDynamicZoneMemberRequest {
    /**
     * Id
     * @type {number}
     * @memberof DynamicZoneMemberApiGetDynamicZoneMember
     */
    readonly id: number

    /**
     * Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names 
     * @type {string}
     * @memberof DynamicZoneMemberApiGetDynamicZoneMember
     */
    readonly includes?: string

    /**
     * Column names [.] separated to fetch specific fields in response
     * @type {string}
     * @memberof DynamicZoneMemberApiGetDynamicZoneMember
     */
    readonly select?: string
}

/**
 * Request parameters for getDynamicZoneMembersBulk operation in DynamicZoneMemberApi.
 * @export
 * @interface DynamicZoneMemberApiGetDynamicZoneMembersBulkRequest
 */
export interface DynamicZoneMemberApiGetDynamicZoneMembersBulkRequest {
    /**
     * body
     * @type {CrudcontrollersBulkFetchByIdsGetRequest}
     * @memberof DynamicZoneMemberApiGetDynamicZoneMembersBulk
     */
    readonly body: CrudcontrollersBulkFetchByIdsGetRequest
}

/**
 * Request parameters for listDynamicZoneMembers operation in DynamicZoneMemberApi.
 * @export
 * @interface DynamicZoneMemberApiListDynamicZoneMembersRequest
 */
export interface DynamicZoneMemberApiListDynamicZoneMembersRequest {
    /**
     * Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names 
     * @type {string}
     * @memberof DynamicZoneMemberApiListDynamicZoneMembers
     */
    readonly includes?: string

    /**
     * Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2
     * @type {string}
     * @memberof DynamicZoneMemberApiListDynamicZoneMembers
     */
    readonly where?: string

    /**
     * Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2
     * @type {string}
     * @memberof DynamicZoneMemberApiListDynamicZoneMembers
     */
    readonly whereOr?: string

    /**
     * Group by field. Multiple conditions [.] separated Example: field1.field2
     * @type {string}
     * @memberof DynamicZoneMemberApiListDynamicZoneMembers
     */
    readonly groupBy?: string

    /**
     * Rows to limit in response (Default: 10,000)
     * @type {string}
     * @memberof DynamicZoneMemberApiListDynamicZoneMembers
     */
    readonly limit?: string

    /**
     * Pagination page
     * @type {number}
     * @memberof DynamicZoneMemberApiListDynamicZoneMembers
     */
    readonly page?: number

    /**
     * Order by [field]
     * @type {string}
     * @memberof DynamicZoneMemberApiListDynamicZoneMembers
     */
    readonly orderBy?: string

    /**
     * Order by field direction
     * @type {string}
     * @memberof DynamicZoneMemberApiListDynamicZoneMembers
     */
    readonly orderDirection?: string

    /**
     * Column names [.] separated to fetch specific fields in response
     * @type {string}
     * @memberof DynamicZoneMemberApiListDynamicZoneMembers
     */
    readonly select?: string
}

/**
 * Request parameters for updateDynamicZoneMember operation in DynamicZoneMemberApi.
 * @export
 * @interface DynamicZoneMemberApiUpdateDynamicZoneMemberRequest
 */
export interface DynamicZoneMemberApiUpdateDynamicZoneMemberRequest {
    /**
     * Id
     * @type {number}
     * @memberof DynamicZoneMemberApiUpdateDynamicZoneMember
     */
    readonly id: number

    /**
     * DynamicZoneMember
     * @type {ModelsDynamicZoneMember}
     * @memberof DynamicZoneMemberApiUpdateDynamicZoneMember
     */
    readonly dynamicZoneMember: ModelsDynamicZoneMember
}

/**
 * DynamicZoneMemberApi - object-oriented interface
 * @export
 * @class DynamicZoneMemberApi
 * @extends {BaseAPI}
 */
export class DynamicZoneMemberApi extends BaseAPI {
    /**
     * 
     * @summary Creates DynamicZoneMember
     * @param {DynamicZoneMemberApiCreateDynamicZoneMemberRequest} requestParameters Request parameters.
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof DynamicZoneMemberApi
     */
    public createDynamicZoneMember(requestParameters: DynamicZoneMemberApiCreateDynamicZoneMemberRequest, options?: any) {
        return DynamicZoneMemberApiFp(this.configuration).createDynamicZoneMember(requestParameters.dynamicZoneMember, options).then((request) => request(this.axios, this.basePath));
    }

    /**
     * 
     * @summary Deletes DynamicZoneMember
     * @param {DynamicZoneMemberApiDeleteDynamicZoneMemberRequest} requestParameters Request parameters.
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof DynamicZoneMemberApi
     */
    public deleteDynamicZoneMember(requestParameters: DynamicZoneMemberApiDeleteDynamicZoneMemberRequest, options?: any) {
        return DynamicZoneMemberApiFp(this.configuration).deleteDynamicZoneMember(requestParameters.id, options).then((request) => request(this.axios, this.basePath));
    }

    /**
     * 
     * @summary Gets DynamicZoneMember
     * @param {DynamicZoneMemberApiGetDynamicZoneMemberRequest} requestParameters Request parameters.
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof DynamicZoneMemberApi
     */
    public getDynamicZoneMember(requestParameters: DynamicZoneMemberApiGetDynamicZoneMemberRequest, options?: any) {
        return DynamicZoneMemberApiFp(this.configuration).getDynamicZoneMember(requestParameters.id, requestParameters.includes, requestParameters.select, options).then((request) => request(this.axios, this.basePath));
    }

    /**
     * 
     * @summary Gets DynamicZoneMembers in bulk
     * @param {DynamicZoneMemberApiGetDynamicZoneMembersBulkRequest} requestParameters Request parameters.
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof DynamicZoneMemberApi
     */
    public getDynamicZoneMembersBulk(requestParameters: DynamicZoneMemberApiGetDynamicZoneMembersBulkRequest, options?: any) {
        return DynamicZoneMemberApiFp(this.configuration).getDynamicZoneMembersBulk(requestParameters.body, options).then((request) => request(this.axios, this.basePath));
    }

    /**
     * 
     * @summary Lists DynamicZoneMembers
     * @param {DynamicZoneMemberApiListDynamicZoneMembersRequest} requestParameters Request parameters.
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof DynamicZoneMemberApi
     */
    public listDynamicZoneMembers(requestParameters: DynamicZoneMemberApiListDynamicZoneMembersRequest = {}, options?: any) {
        return DynamicZoneMemberApiFp(this.configuration).listDynamicZoneMembers(requestParameters.includes, requestParameters.where, requestParameters.whereOr, requestParameters.groupBy, requestParameters.limit, requestParameters.page, requestParameters.orderBy, requestParameters.orderDirection, requestParameters.select, options).then((request) => request(this.axios, this.basePath));
    }

    /**
     * 
     * @summary Updates DynamicZoneMember
     * @param {DynamicZoneMemberApiUpdateDynamicZoneMemberRequest} requestParameters Request parameters.
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof DynamicZoneMemberApi
     */
    public updateDynamicZoneMember(requestParameters: DynamicZoneMemberApiUpdateDynamicZoneMemberRequest, options?: any) {
        return DynamicZoneMemberApiFp(this.configuration).updateDynamicZoneMember(requestParameters.id, requestParameters.dynamicZoneMember, options).then((request) => request(this.axios, this.basePath));
    }
}
