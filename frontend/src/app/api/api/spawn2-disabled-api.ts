import globalAxios, { AxiosPromise, AxiosInstance } from 'axios';
import { Configuration } from '../configuration';
// Some imports not used depending on template conditions
// @ts-ignore
import { BASE_PATH, COLLECTION_FORMATS, RequestArgs, BaseAPI, RequiredError } from '../base';
// @ts-ignore
import { CrudcontrollersBulkFetchByIdsGetRequest } from '../models';
// @ts-ignore
import { ModelsSpawn2Disabled } from '../models';
export const Spawn2DisabledApiAxiosParamCreator = function (configuration?: Configuration) {
    return {
        createSpawn2Disabled: async (spawn2Disabled: ModelsSpawn2Disabled, options: any = {}): Promise<RequestArgs> => {
            if (spawn2Disabled === null || spawn2Disabled === undefined) {
                throw new RequiredError('spawn2Disabled','Required parameter spawn2Disabled was null or undefined when calling createSpawn2Disabled.');
            }
            const localVarPath = `/spawn_2_disabled`;
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
            const nonString = typeof spawn2Disabled !== 'string';
            const needsSerialization = nonString && configuration && configuration.isJsonMime
                ? configuration.isJsonMime(localVarRequestOptions.headers['Content-Type'])
                : nonString;
            localVarRequestOptions.data =  needsSerialization
                ? JSON.stringify(spawn2Disabled !== undefined ? spawn2Disabled : {})
                : (spawn2Disabled || "");
            return {
                url: localVarUrlObj.pathname + localVarUrlObj.search + localVarUrlObj.hash,
                options: localVarRequestOptions,
            };
        },
        deleteSpawn2Disabled: async (id: number, options: any = {}): Promise<RequestArgs> => {
            if (id === null || id === undefined) {
                throw new RequiredError('id','Required parameter id was null or undefined when calling deleteSpawn2Disabled.');
            }
            const localVarPath = `/spawn_2_disabled/{id}`
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
        getSpawn2Disabled: async (id: number, includes?: string, select?: string, options: any = {}): Promise<RequestArgs> => {
            if (id === null || id === undefined) {
                throw new RequiredError('id','Required parameter id was null or undefined when calling getSpawn2Disabled.');
            }
            const localVarPath = `/spawn_2_disabled/{id}`
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
        getSpawn2DisabledsBulk: async (body: CrudcontrollersBulkFetchByIdsGetRequest, options: any = {}): Promise<RequestArgs> => {
            if (body === null || body === undefined) {
                throw new RequiredError('body','Required parameter body was null or undefined when calling getSpawn2DisabledsBulk.');
            }
            const localVarPath = `/spawn_2_disableds/bulk`;
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
        getSpawn2DisabledsCount: async (includes?: string, where?: string, whereOr?: string, groupBy?: string, limit?: string, page?: number, orderBy?: string, orderDirection?: string, select?: string, options: any = {}): Promise<RequestArgs> => {
            const localVarPath = `/spawn_2_disableds/count`;
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
        listSpawn2Disableds: async (includes?: string, where?: string, whereOr?: string, groupBy?: string, limit?: string, page?: number, orderBy?: string, orderDirection?: string, select?: string, options: any = {}): Promise<RequestArgs> => {
            const localVarPath = `/spawn_2_disableds`;
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
        updateSpawn2Disabled: async (id: number, spawn2Disabled: ModelsSpawn2Disabled, options: any = {}): Promise<RequestArgs> => {
            if (id === null || id === undefined) {
                throw new RequiredError('id','Required parameter id was null or undefined when calling updateSpawn2Disabled.');
            }
            if (spawn2Disabled === null || spawn2Disabled === undefined) {
                throw new RequiredError('spawn2Disabled','Required parameter spawn2Disabled was null or undefined when calling updateSpawn2Disabled.');
            }
            const localVarPath = `/spawn_2_disabled/{id}`
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
            const nonString = typeof spawn2Disabled !== 'string';
            const needsSerialization = nonString && configuration && configuration.isJsonMime
                ? configuration.isJsonMime(localVarRequestOptions.headers['Content-Type'])
                : nonString;
            localVarRequestOptions.data =  needsSerialization
                ? JSON.stringify(spawn2Disabled !== undefined ? spawn2Disabled : {})
                : (spawn2Disabled || "");
            return {
                url: localVarUrlObj.pathname + localVarUrlObj.search + localVarUrlObj.hash,
                options: localVarRequestOptions,
            };
        },
    }
};
export const Spawn2DisabledApiFp = function(configuration?: Configuration) {
    return {
        async createSpawn2Disabled(spawn2Disabled: ModelsSpawn2Disabled, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsSpawn2Disabled>>> {
            const localVarAxiosArgs = await Spawn2DisabledApiAxiosParamCreator(configuration).createSpawn2Disabled(spawn2Disabled, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        async deleteSpawn2Disabled(id: number, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<string>> {
            const localVarAxiosArgs = await Spawn2DisabledApiAxiosParamCreator(configuration).deleteSpawn2Disabled(id, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        async getSpawn2Disabled(id: number, includes?: string, select?: string, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsSpawn2Disabled>>> {
            const localVarAxiosArgs = await Spawn2DisabledApiAxiosParamCreator(configuration).getSpawn2Disabled(id, includes, select, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        async getSpawn2DisabledsBulk(body: CrudcontrollersBulkFetchByIdsGetRequest, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsSpawn2Disabled>>> {
            const localVarAxiosArgs = await Spawn2DisabledApiAxiosParamCreator(configuration).getSpawn2DisabledsBulk(body, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        async getSpawn2DisabledsCount(includes?: string, where?: string, whereOr?: string, groupBy?: string, limit?: string, page?: number, orderBy?: string, orderDirection?: string, select?: string, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsSpawn2Disabled>>> {
            const localVarAxiosArgs = await Spawn2DisabledApiAxiosParamCreator(configuration).getSpawn2DisabledsCount(includes, where, whereOr, groupBy, limit, page, orderBy, orderDirection, select, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        async listSpawn2Disableds(includes?: string, where?: string, whereOr?: string, groupBy?: string, limit?: string, page?: number, orderBy?: string, orderDirection?: string, select?: string, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsSpawn2Disabled>>> {
            const localVarAxiosArgs = await Spawn2DisabledApiAxiosParamCreator(configuration).listSpawn2Disableds(includes, where, whereOr, groupBy, limit, page, orderBy, orderDirection, select, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        async updateSpawn2Disabled(id: number, spawn2Disabled: ModelsSpawn2Disabled, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsSpawn2Disabled>>> {
            const localVarAxiosArgs = await Spawn2DisabledApiAxiosParamCreator(configuration).updateSpawn2Disabled(id, spawn2Disabled, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
    }
};
export const Spawn2DisabledApiFactory = function (configuration?: Configuration, basePath?: string, axios?: AxiosInstance) {
    return {
        createSpawn2Disabled(spawn2Disabled: ModelsSpawn2Disabled, options?: any): AxiosPromise<Array<ModelsSpawn2Disabled>> {
            return Spawn2DisabledApiFp(configuration).createSpawn2Disabled(spawn2Disabled, options).then((request) => request(axios, basePath));
        },
        deleteSpawn2Disabled(id: number, options?: any): AxiosPromise<string> {
            return Spawn2DisabledApiFp(configuration).deleteSpawn2Disabled(id, options).then((request) => request(axios, basePath));
        },
        getSpawn2Disabled(id: number, includes?: string, select?: string, options?: any): AxiosPromise<Array<ModelsSpawn2Disabled>> {
            return Spawn2DisabledApiFp(configuration).getSpawn2Disabled(id, includes, select, options).then((request) => request(axios, basePath));
        },
        getSpawn2DisabledsBulk(body: CrudcontrollersBulkFetchByIdsGetRequest, options?: any): AxiosPromise<Array<ModelsSpawn2Disabled>> {
            return Spawn2DisabledApiFp(configuration).getSpawn2DisabledsBulk(body, options).then((request) => request(axios, basePath));
        },
        getSpawn2DisabledsCount(includes?: string, where?: string, whereOr?: string, groupBy?: string, limit?: string, page?: number, orderBy?: string, orderDirection?: string, select?: string, options?: any): AxiosPromise<Array<ModelsSpawn2Disabled>> {
            return Spawn2DisabledApiFp(configuration).getSpawn2DisabledsCount(includes, where, whereOr, groupBy, limit, page, orderBy, orderDirection, select, options).then((request) => request(axios, basePath));
        },
        listSpawn2Disableds(includes?: string, where?: string, whereOr?: string, groupBy?: string, limit?: string, page?: number, orderBy?: string, orderDirection?: string, select?: string, options?: any): AxiosPromise<Array<ModelsSpawn2Disabled>> {
            return Spawn2DisabledApiFp(configuration).listSpawn2Disableds(includes, where, whereOr, groupBy, limit, page, orderBy, orderDirection, select, options).then((request) => request(axios, basePath));
        },
        updateSpawn2Disabled(id: number, spawn2Disabled: ModelsSpawn2Disabled, options?: any): AxiosPromise<Array<ModelsSpawn2Disabled>> {
            return Spawn2DisabledApiFp(configuration).updateSpawn2Disabled(id, spawn2Disabled, options).then((request) => request(axios, basePath));
        },
    };
};
export interface Spawn2DisabledApiCreateSpawn2DisabledRequest {
    readonly spawn2Disabled: ModelsSpawn2Disabled
}
export interface Spawn2DisabledApiDeleteSpawn2DisabledRequest {
    readonly id: number
}
export interface Spawn2DisabledApiGetSpawn2DisabledRequest {
    readonly id: number
    readonly includes?: string
    readonly select?: string
}
export interface Spawn2DisabledApiGetSpawn2DisabledsBulkRequest {
    readonly body: CrudcontrollersBulkFetchByIdsGetRequest
}
export interface Spawn2DisabledApiGetSpawn2DisabledsCountRequest {
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
export interface Spawn2DisabledApiListSpawn2DisabledsRequest {
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
export interface Spawn2DisabledApiUpdateSpawn2DisabledRequest {
    readonly id: number
    readonly spawn2Disabled: ModelsSpawn2Disabled
}
export class Spawn2DisabledApi extends BaseAPI {
    public createSpawn2Disabled(requestParameters: Spawn2DisabledApiCreateSpawn2DisabledRequest, options?: any) {
        return Spawn2DisabledApiFp(this.configuration).createSpawn2Disabled(requestParameters.spawn2Disabled, options).then((request) => request(this.axios, this.basePath));
    }
    public deleteSpawn2Disabled(requestParameters: Spawn2DisabledApiDeleteSpawn2DisabledRequest, options?: any) {
        return Spawn2DisabledApiFp(this.configuration).deleteSpawn2Disabled(requestParameters.id, options).then((request) => request(this.axios, this.basePath));
    }
    public getSpawn2Disabled(requestParameters: Spawn2DisabledApiGetSpawn2DisabledRequest, options?: any) {
        return Spawn2DisabledApiFp(this.configuration).getSpawn2Disabled(requestParameters.id, requestParameters.includes, requestParameters.select, options).then((request) => request(this.axios, this.basePath));
    }
    public getSpawn2DisabledsBulk(requestParameters: Spawn2DisabledApiGetSpawn2DisabledsBulkRequest, options?: any) {
        return Spawn2DisabledApiFp(this.configuration).getSpawn2DisabledsBulk(requestParameters.body, options).then((request) => request(this.axios, this.basePath));
    }
    public getSpawn2DisabledsCount(requestParameters: Spawn2DisabledApiGetSpawn2DisabledsCountRequest = {}, options?: any) {
        return Spawn2DisabledApiFp(this.configuration).getSpawn2DisabledsCount(requestParameters.includes, requestParameters.where, requestParameters.whereOr, requestParameters.groupBy, requestParameters.limit, requestParameters.page, requestParameters.orderBy, requestParameters.orderDirection, requestParameters.select, options).then((request) => request(this.axios, this.basePath));
    }
    public listSpawn2Disableds(requestParameters: Spawn2DisabledApiListSpawn2DisabledsRequest = {}, options?: any) {
        return Spawn2DisabledApiFp(this.configuration).listSpawn2Disableds(requestParameters.includes, requestParameters.where, requestParameters.whereOr, requestParameters.groupBy, requestParameters.limit, requestParameters.page, requestParameters.orderBy, requestParameters.orderDirection, requestParameters.select, options).then((request) => request(this.axios, this.basePath));
    }
    public updateSpawn2Disabled(requestParameters: Spawn2DisabledApiUpdateSpawn2DisabledRequest, options?: any) {
        return Spawn2DisabledApiFp(this.configuration).updateSpawn2Disabled(requestParameters.id, requestParameters.spawn2Disabled, options).then((request) => request(this.axios, this.basePath));
    }
}
