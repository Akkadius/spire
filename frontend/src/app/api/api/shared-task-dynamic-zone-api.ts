import globalAxios, { AxiosPromise, AxiosInstance } from 'axios';
import { Configuration } from '../configuration';
// Some imports not used depending on template conditions
// @ts-ignore
import { BASE_PATH, COLLECTION_FORMATS, RequestArgs, BaseAPI, RequiredError } from '../base';
// @ts-ignore
import { CrudcontrollersBulkFetchByIdsGetRequest } from '../models';
// @ts-ignore
import { ModelsSharedTaskDynamicZone } from '../models';
export const SharedTaskDynamicZoneApiAxiosParamCreator = function (configuration?: Configuration) {
    return {
        createSharedTaskDynamicZone: async (sharedTaskDynamicZone: ModelsSharedTaskDynamicZone, options: any = {}): Promise<RequestArgs> => {
            if (sharedTaskDynamicZone === null || sharedTaskDynamicZone === undefined) {
                throw new RequiredError('sharedTaskDynamicZone','Required parameter sharedTaskDynamicZone was null or undefined when calling createSharedTaskDynamicZone.');
            }
            const localVarPath = `/shared_task_dynamic_zone`;
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
            const nonString = typeof sharedTaskDynamicZone !== 'string';
            const needsSerialization = nonString && configuration && configuration.isJsonMime
                ? configuration.isJsonMime(localVarRequestOptions.headers['Content-Type'])
                : nonString;
            localVarRequestOptions.data =  needsSerialization
                ? JSON.stringify(sharedTaskDynamicZone !== undefined ? sharedTaskDynamicZone : {})
                : (sharedTaskDynamicZone || "");
            return {
                url: localVarUrlObj.pathname + localVarUrlObj.search + localVarUrlObj.hash,
                options: localVarRequestOptions,
            };
        },
        deleteSharedTaskDynamicZone: async (id: number, options: any = {}): Promise<RequestArgs> => {
            if (id === null || id === undefined) {
                throw new RequiredError('id','Required parameter id was null or undefined when calling deleteSharedTaskDynamicZone.');
            }
            const localVarPath = `/shared_task_dynamic_zone/{id}`
                .replace(`{${"id"}}`, encodeURIComponent(String(id)));
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
        getSharedTaskDynamicZone: async (id: number, includes?: string, select?: string, options: any = {}): Promise<RequestArgs> => {
            if (id === null || id === undefined) {
                throw new RequiredError('id','Required parameter id was null or undefined when calling getSharedTaskDynamicZone.');
            }
            const localVarPath = `/shared_task_dynamic_zone/{id}`
                .replace(`{${"id"}}`, encodeURIComponent(String(id)));
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
        getSharedTaskDynamicZonesBulk: async (body: CrudcontrollersBulkFetchByIdsGetRequest, options: any = {}): Promise<RequestArgs> => {
            if (body === null || body === undefined) {
                throw new RequiredError('body','Required parameter body was null or undefined when calling getSharedTaskDynamicZonesBulk.');
            }
            const localVarPath = `/shared_task_dynamic_zones/bulk`;
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
        getSharedTaskDynamicZonesCount: async (includes?: string, where?: string, whereOr?: string, groupBy?: string, limit?: string, page?: number, orderBy?: string, orderDirection?: string, select?: string, options: any = {}): Promise<RequestArgs> => {
            const localVarPath = `/shared_task_dynamic_zones/count`;
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
        listSharedTaskDynamicZones: async (includes?: string, where?: string, whereOr?: string, groupBy?: string, limit?: string, page?: number, orderBy?: string, orderDirection?: string, select?: string, options: any = {}): Promise<RequestArgs> => {
            const localVarPath = `/shared_task_dynamic_zones`;
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
        updateSharedTaskDynamicZone: async (id: number, sharedTaskDynamicZone: ModelsSharedTaskDynamicZone, options: any = {}): Promise<RequestArgs> => {
            if (id === null || id === undefined) {
                throw new RequiredError('id','Required parameter id was null or undefined when calling updateSharedTaskDynamicZone.');
            }
            if (sharedTaskDynamicZone === null || sharedTaskDynamicZone === undefined) {
                throw new RequiredError('sharedTaskDynamicZone','Required parameter sharedTaskDynamicZone was null or undefined when calling updateSharedTaskDynamicZone.');
            }
            const localVarPath = `/shared_task_dynamic_zone/{id}`
                .replace(`{${"id"}}`, encodeURIComponent(String(id)));
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
            const nonString = typeof sharedTaskDynamicZone !== 'string';
            const needsSerialization = nonString && configuration && configuration.isJsonMime
                ? configuration.isJsonMime(localVarRequestOptions.headers['Content-Type'])
                : nonString;
            localVarRequestOptions.data =  needsSerialization
                ? JSON.stringify(sharedTaskDynamicZone !== undefined ? sharedTaskDynamicZone : {})
                : (sharedTaskDynamicZone || "");
            return {
                url: localVarUrlObj.pathname + localVarUrlObj.search + localVarUrlObj.hash,
                options: localVarRequestOptions,
            };
        },
    }
};
export const SharedTaskDynamicZoneApiFp = function(configuration?: Configuration) {
    return {
        async createSharedTaskDynamicZone(sharedTaskDynamicZone: ModelsSharedTaskDynamicZone, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsSharedTaskDynamicZone>>> {
            const localVarAxiosArgs = await SharedTaskDynamicZoneApiAxiosParamCreator(configuration).createSharedTaskDynamicZone(sharedTaskDynamicZone, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        async deleteSharedTaskDynamicZone(id: number, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<string>> {
            const localVarAxiosArgs = await SharedTaskDynamicZoneApiAxiosParamCreator(configuration).deleteSharedTaskDynamicZone(id, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        async getSharedTaskDynamicZone(id: number, includes?: string, select?: string, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsSharedTaskDynamicZone>>> {
            const localVarAxiosArgs = await SharedTaskDynamicZoneApiAxiosParamCreator(configuration).getSharedTaskDynamicZone(id, includes, select, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        async getSharedTaskDynamicZonesBulk(body: CrudcontrollersBulkFetchByIdsGetRequest, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsSharedTaskDynamicZone>>> {
            const localVarAxiosArgs = await SharedTaskDynamicZoneApiAxiosParamCreator(configuration).getSharedTaskDynamicZonesBulk(body, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        async getSharedTaskDynamicZonesCount(includes?: string, where?: string, whereOr?: string, groupBy?: string, limit?: string, page?: number, orderBy?: string, orderDirection?: string, select?: string, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsSharedTaskDynamicZone>>> {
            const localVarAxiosArgs = await SharedTaskDynamicZoneApiAxiosParamCreator(configuration).getSharedTaskDynamicZonesCount(includes, where, whereOr, groupBy, limit, page, orderBy, orderDirection, select, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        async listSharedTaskDynamicZones(includes?: string, where?: string, whereOr?: string, groupBy?: string, limit?: string, page?: number, orderBy?: string, orderDirection?: string, select?: string, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsSharedTaskDynamicZone>>> {
            const localVarAxiosArgs = await SharedTaskDynamicZoneApiAxiosParamCreator(configuration).listSharedTaskDynamicZones(includes, where, whereOr, groupBy, limit, page, orderBy, orderDirection, select, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        async updateSharedTaskDynamicZone(id: number, sharedTaskDynamicZone: ModelsSharedTaskDynamicZone, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsSharedTaskDynamicZone>>> {
            const localVarAxiosArgs = await SharedTaskDynamicZoneApiAxiosParamCreator(configuration).updateSharedTaskDynamicZone(id, sharedTaskDynamicZone, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
    }
};
export const SharedTaskDynamicZoneApiFactory = function (configuration?: Configuration, basePath?: string, axios?: AxiosInstance) {
    return {
        createSharedTaskDynamicZone(sharedTaskDynamicZone: ModelsSharedTaskDynamicZone, options?: any): AxiosPromise<Array<ModelsSharedTaskDynamicZone>> {
            return SharedTaskDynamicZoneApiFp(configuration).createSharedTaskDynamicZone(sharedTaskDynamicZone, options).then((request) => request(axios, basePath));
        },
        deleteSharedTaskDynamicZone(id: number, options?: any): AxiosPromise<string> {
            return SharedTaskDynamicZoneApiFp(configuration).deleteSharedTaskDynamicZone(id, options).then((request) => request(axios, basePath));
        },
        getSharedTaskDynamicZone(id: number, includes?: string, select?: string, options?: any): AxiosPromise<Array<ModelsSharedTaskDynamicZone>> {
            return SharedTaskDynamicZoneApiFp(configuration).getSharedTaskDynamicZone(id, includes, select, options).then((request) => request(axios, basePath));
        },
        getSharedTaskDynamicZonesBulk(body: CrudcontrollersBulkFetchByIdsGetRequest, options?: any): AxiosPromise<Array<ModelsSharedTaskDynamicZone>> {
            return SharedTaskDynamicZoneApiFp(configuration).getSharedTaskDynamicZonesBulk(body, options).then((request) => request(axios, basePath));
        },
        getSharedTaskDynamicZonesCount(includes?: string, where?: string, whereOr?: string, groupBy?: string, limit?: string, page?: number, orderBy?: string, orderDirection?: string, select?: string, options?: any): AxiosPromise<Array<ModelsSharedTaskDynamicZone>> {
            return SharedTaskDynamicZoneApiFp(configuration).getSharedTaskDynamicZonesCount(includes, where, whereOr, groupBy, limit, page, orderBy, orderDirection, select, options).then((request) => request(axios, basePath));
        },
        listSharedTaskDynamicZones(includes?: string, where?: string, whereOr?: string, groupBy?: string, limit?: string, page?: number, orderBy?: string, orderDirection?: string, select?: string, options?: any): AxiosPromise<Array<ModelsSharedTaskDynamicZone>> {
            return SharedTaskDynamicZoneApiFp(configuration).listSharedTaskDynamicZones(includes, where, whereOr, groupBy, limit, page, orderBy, orderDirection, select, options).then((request) => request(axios, basePath));
        },
        updateSharedTaskDynamicZone(id: number, sharedTaskDynamicZone: ModelsSharedTaskDynamicZone, options?: any): AxiosPromise<Array<ModelsSharedTaskDynamicZone>> {
            return SharedTaskDynamicZoneApiFp(configuration).updateSharedTaskDynamicZone(id, sharedTaskDynamicZone, options).then((request) => request(axios, basePath));
        },
    };
};
export interface SharedTaskDynamicZoneApiCreateSharedTaskDynamicZoneRequest {
    readonly sharedTaskDynamicZone: ModelsSharedTaskDynamicZone
}
export interface SharedTaskDynamicZoneApiDeleteSharedTaskDynamicZoneRequest {
    readonly id: number
}
export interface SharedTaskDynamicZoneApiGetSharedTaskDynamicZoneRequest {
    readonly id: number
    readonly includes?: string
    readonly select?: string
}
export interface SharedTaskDynamicZoneApiGetSharedTaskDynamicZonesBulkRequest {
    readonly body: CrudcontrollersBulkFetchByIdsGetRequest
}
export interface SharedTaskDynamicZoneApiGetSharedTaskDynamicZonesCountRequest {
    readonly includes?: string
    readonly where?: string
    readonly whereOr?: string
    readonly groupBy?: string
    readonly limit?: string
    readonly page?: number
    readonly orderBy?: string
    readonly orderDirection?: string
    readonly select?: string
}
export interface SharedTaskDynamicZoneApiListSharedTaskDynamicZonesRequest {
    readonly includes?: string
    readonly where?: string
    readonly whereOr?: string
    readonly groupBy?: string
    readonly limit?: string
    readonly page?: number
    readonly orderBy?: string
    readonly orderDirection?: string
    readonly select?: string
}
export interface SharedTaskDynamicZoneApiUpdateSharedTaskDynamicZoneRequest {
    readonly id: number
    readonly sharedTaskDynamicZone: ModelsSharedTaskDynamicZone
}
export class SharedTaskDynamicZoneApi extends BaseAPI {
    public createSharedTaskDynamicZone(requestParameters: SharedTaskDynamicZoneApiCreateSharedTaskDynamicZoneRequest, options?: any) {
        return SharedTaskDynamicZoneApiFp(this.configuration).createSharedTaskDynamicZone(requestParameters.sharedTaskDynamicZone, options).then((request) => request(this.axios, this.basePath));
    }
    public deleteSharedTaskDynamicZone(requestParameters: SharedTaskDynamicZoneApiDeleteSharedTaskDynamicZoneRequest, options?: any) {
        return SharedTaskDynamicZoneApiFp(this.configuration).deleteSharedTaskDynamicZone(requestParameters.id, options).then((request) => request(this.axios, this.basePath));
    }
    public getSharedTaskDynamicZone(requestParameters: SharedTaskDynamicZoneApiGetSharedTaskDynamicZoneRequest, options?: any) {
        return SharedTaskDynamicZoneApiFp(this.configuration).getSharedTaskDynamicZone(requestParameters.id, requestParameters.includes, requestParameters.select, options).then((request) => request(this.axios, this.basePath));
    }
    public getSharedTaskDynamicZonesBulk(requestParameters: SharedTaskDynamicZoneApiGetSharedTaskDynamicZonesBulkRequest, options?: any) {
        return SharedTaskDynamicZoneApiFp(this.configuration).getSharedTaskDynamicZonesBulk(requestParameters.body, options).then((request) => request(this.axios, this.basePath));
    }
    public getSharedTaskDynamicZonesCount(requestParameters: SharedTaskDynamicZoneApiGetSharedTaskDynamicZonesCountRequest = {}, options?: any) {
        return SharedTaskDynamicZoneApiFp(this.configuration).getSharedTaskDynamicZonesCount(requestParameters.includes, requestParameters.where, requestParameters.whereOr, requestParameters.groupBy, requestParameters.limit, requestParameters.page, requestParameters.orderBy, requestParameters.orderDirection, requestParameters.select, options).then((request) => request(this.axios, this.basePath));
    }
    public listSharedTaskDynamicZones(requestParameters: SharedTaskDynamicZoneApiListSharedTaskDynamicZonesRequest = {}, options?: any) {
        return SharedTaskDynamicZoneApiFp(this.configuration).listSharedTaskDynamicZones(requestParameters.includes, requestParameters.where, requestParameters.whereOr, requestParameters.groupBy, requestParameters.limit, requestParameters.page, requestParameters.orderBy, requestParameters.orderDirection, requestParameters.select, options).then((request) => request(this.axios, this.basePath));
    }
    public updateSharedTaskDynamicZone(requestParameters: SharedTaskDynamicZoneApiUpdateSharedTaskDynamicZoneRequest, options?: any) {
        return SharedTaskDynamicZoneApiFp(this.configuration).updateSharedTaskDynamicZone(requestParameters.id, requestParameters.sharedTaskDynamicZone, options).then((request) => request(this.axios, this.basePath));
    }
}
