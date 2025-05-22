import globalAxios, { AxiosPromise, AxiosInstance } from 'axios';
import { Configuration } from '../configuration';
// Some imports not used depending on template conditions
// @ts-ignore
import { BASE_PATH, COLLECTION_FORMATS, RequestArgs, BaseAPI, RequiredError } from '../base';
// @ts-ignore
import { CrudcontrollersBulkFetchByIdsGetRequest } from '../models';
// @ts-ignore
import { ModelsSharedTaskActivityState } from '../models';
export const SharedTaskActivityStateApiAxiosParamCreator = function (configuration?: Configuration) {
    return {
        createSharedTaskActivityState: async (sharedTaskActivityState: ModelsSharedTaskActivityState, options: any = {}): Promise<RequestArgs> => {
            if (sharedTaskActivityState === null || sharedTaskActivityState === undefined) {
                throw new RequiredError('sharedTaskActivityState','Required parameter sharedTaskActivityState was null or undefined when calling createSharedTaskActivityState.');
            }
            const localVarPath = `/shared_task_activity_state`;
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
            const nonString = typeof sharedTaskActivityState !== 'string';
            const needsSerialization = nonString && configuration && configuration.isJsonMime
                ? configuration.isJsonMime(localVarRequestOptions.headers['Content-Type'])
                : nonString;
            localVarRequestOptions.data =  needsSerialization
                ? JSON.stringify(sharedTaskActivityState !== undefined ? sharedTaskActivityState : {})
                : (sharedTaskActivityState || "");
            return {
                url: localVarUrlObj.pathname + localVarUrlObj.search + localVarUrlObj.hash,
                options: localVarRequestOptions,
            };
        },
        deleteSharedTaskActivityState: async (id: number, options: any = {}): Promise<RequestArgs> => {
            if (id === null || id === undefined) {
                throw new RequiredError('id','Required parameter id was null or undefined when calling deleteSharedTaskActivityState.');
            }
            const localVarPath = `/shared_task_activity_state/{id}`
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
        getSharedTaskActivityState: async (id: number, includes?: string, select?: string, options: any = {}): Promise<RequestArgs> => {
            if (id === null || id === undefined) {
                throw new RequiredError('id','Required parameter id was null or undefined when calling getSharedTaskActivityState.');
            }
            const localVarPath = `/shared_task_activity_state/{id}`
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
        getSharedTaskActivityStatesBulk: async (body: CrudcontrollersBulkFetchByIdsGetRequest, options: any = {}): Promise<RequestArgs> => {
            if (body === null || body === undefined) {
                throw new RequiredError('body','Required parameter body was null or undefined when calling getSharedTaskActivityStatesBulk.');
            }
            const localVarPath = `/shared_task_activity_states/bulk`;
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
        getSharedTaskActivityStatesCount: async (includes?: string, where?: string, whereOr?: string, groupBy?: string, limit?: string, page?: number, orderBy?: string, orderDirection?: string, select?: string, options: any = {}): Promise<RequestArgs> => {
            const localVarPath = `/shared_task_activity_states/count`;
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
        listSharedTaskActivityStates: async (includes?: string, where?: string, whereOr?: string, groupBy?: string, limit?: string, page?: number, orderBy?: string, orderDirection?: string, select?: string, options: any = {}): Promise<RequestArgs> => {
            const localVarPath = `/shared_task_activity_states`;
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
        updateSharedTaskActivityState: async (id: number, sharedTaskActivityState: ModelsSharedTaskActivityState, options: any = {}): Promise<RequestArgs> => {
            if (id === null || id === undefined) {
                throw new RequiredError('id','Required parameter id was null or undefined when calling updateSharedTaskActivityState.');
            }
            if (sharedTaskActivityState === null || sharedTaskActivityState === undefined) {
                throw new RequiredError('sharedTaskActivityState','Required parameter sharedTaskActivityState was null or undefined when calling updateSharedTaskActivityState.');
            }
            const localVarPath = `/shared_task_activity_state/{id}`
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
            const nonString = typeof sharedTaskActivityState !== 'string';
            const needsSerialization = nonString && configuration && configuration.isJsonMime
                ? configuration.isJsonMime(localVarRequestOptions.headers['Content-Type'])
                : nonString;
            localVarRequestOptions.data =  needsSerialization
                ? JSON.stringify(sharedTaskActivityState !== undefined ? sharedTaskActivityState : {})
                : (sharedTaskActivityState || "");
            return {
                url: localVarUrlObj.pathname + localVarUrlObj.search + localVarUrlObj.hash,
                options: localVarRequestOptions,
            };
        },
    }
};
export const SharedTaskActivityStateApiFp = function(configuration?: Configuration) {
    return {
        async createSharedTaskActivityState(sharedTaskActivityState: ModelsSharedTaskActivityState, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsSharedTaskActivityState>>> {
            const localVarAxiosArgs = await SharedTaskActivityStateApiAxiosParamCreator(configuration).createSharedTaskActivityState(sharedTaskActivityState, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        async deleteSharedTaskActivityState(id: number, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<string>> {
            const localVarAxiosArgs = await SharedTaskActivityStateApiAxiosParamCreator(configuration).deleteSharedTaskActivityState(id, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        async getSharedTaskActivityState(id: number, includes?: string, select?: string, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsSharedTaskActivityState>>> {
            const localVarAxiosArgs = await SharedTaskActivityStateApiAxiosParamCreator(configuration).getSharedTaskActivityState(id, includes, select, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        async getSharedTaskActivityStatesBulk(body: CrudcontrollersBulkFetchByIdsGetRequest, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsSharedTaskActivityState>>> {
            const localVarAxiosArgs = await SharedTaskActivityStateApiAxiosParamCreator(configuration).getSharedTaskActivityStatesBulk(body, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        async getSharedTaskActivityStatesCount(includes?: string, where?: string, whereOr?: string, groupBy?: string, limit?: string, page?: number, orderBy?: string, orderDirection?: string, select?: string, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsSharedTaskActivityState>>> {
            const localVarAxiosArgs = await SharedTaskActivityStateApiAxiosParamCreator(configuration).getSharedTaskActivityStatesCount(includes, where, whereOr, groupBy, limit, page, orderBy, orderDirection, select, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        async listSharedTaskActivityStates(includes?: string, where?: string, whereOr?: string, groupBy?: string, limit?: string, page?: number, orderBy?: string, orderDirection?: string, select?: string, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsSharedTaskActivityState>>> {
            const localVarAxiosArgs = await SharedTaskActivityStateApiAxiosParamCreator(configuration).listSharedTaskActivityStates(includes, where, whereOr, groupBy, limit, page, orderBy, orderDirection, select, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        async updateSharedTaskActivityState(id: number, sharedTaskActivityState: ModelsSharedTaskActivityState, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsSharedTaskActivityState>>> {
            const localVarAxiosArgs = await SharedTaskActivityStateApiAxiosParamCreator(configuration).updateSharedTaskActivityState(id, sharedTaskActivityState, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
    }
};
export const SharedTaskActivityStateApiFactory = function (configuration?: Configuration, basePath?: string, axios?: AxiosInstance) {
    return {
        createSharedTaskActivityState(sharedTaskActivityState: ModelsSharedTaskActivityState, options?: any): AxiosPromise<Array<ModelsSharedTaskActivityState>> {
            return SharedTaskActivityStateApiFp(configuration).createSharedTaskActivityState(sharedTaskActivityState, options).then((request) => request(axios, basePath));
        },
        deleteSharedTaskActivityState(id: number, options?: any): AxiosPromise<string> {
            return SharedTaskActivityStateApiFp(configuration).deleteSharedTaskActivityState(id, options).then((request) => request(axios, basePath));
        },
        getSharedTaskActivityState(id: number, includes?: string, select?: string, options?: any): AxiosPromise<Array<ModelsSharedTaskActivityState>> {
            return SharedTaskActivityStateApiFp(configuration).getSharedTaskActivityState(id, includes, select, options).then((request) => request(axios, basePath));
        },
        getSharedTaskActivityStatesBulk(body: CrudcontrollersBulkFetchByIdsGetRequest, options?: any): AxiosPromise<Array<ModelsSharedTaskActivityState>> {
            return SharedTaskActivityStateApiFp(configuration).getSharedTaskActivityStatesBulk(body, options).then((request) => request(axios, basePath));
        },
        getSharedTaskActivityStatesCount(includes?: string, where?: string, whereOr?: string, groupBy?: string, limit?: string, page?: number, orderBy?: string, orderDirection?: string, select?: string, options?: any): AxiosPromise<Array<ModelsSharedTaskActivityState>> {
            return SharedTaskActivityStateApiFp(configuration).getSharedTaskActivityStatesCount(includes, where, whereOr, groupBy, limit, page, orderBy, orderDirection, select, options).then((request) => request(axios, basePath));
        },
        listSharedTaskActivityStates(includes?: string, where?: string, whereOr?: string, groupBy?: string, limit?: string, page?: number, orderBy?: string, orderDirection?: string, select?: string, options?: any): AxiosPromise<Array<ModelsSharedTaskActivityState>> {
            return SharedTaskActivityStateApiFp(configuration).listSharedTaskActivityStates(includes, where, whereOr, groupBy, limit, page, orderBy, orderDirection, select, options).then((request) => request(axios, basePath));
        },
        updateSharedTaskActivityState(id: number, sharedTaskActivityState: ModelsSharedTaskActivityState, options?: any): AxiosPromise<Array<ModelsSharedTaskActivityState>> {
            return SharedTaskActivityStateApiFp(configuration).updateSharedTaskActivityState(id, sharedTaskActivityState, options).then((request) => request(axios, basePath));
        },
    };
};
export interface SharedTaskActivityStateApiCreateSharedTaskActivityStateRequest {
    readonly sharedTaskActivityState: ModelsSharedTaskActivityState
}
export interface SharedTaskActivityStateApiDeleteSharedTaskActivityStateRequest {
    readonly id: number
}
export interface SharedTaskActivityStateApiGetSharedTaskActivityStateRequest {
    readonly id: number
    readonly includes?: string
    readonly select?: string
}
export interface SharedTaskActivityStateApiGetSharedTaskActivityStatesBulkRequest {
    readonly body: CrudcontrollersBulkFetchByIdsGetRequest
}
export interface SharedTaskActivityStateApiGetSharedTaskActivityStatesCountRequest {
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
export interface SharedTaskActivityStateApiListSharedTaskActivityStatesRequest {
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
export interface SharedTaskActivityStateApiUpdateSharedTaskActivityStateRequest {
    readonly id: number
    readonly sharedTaskActivityState: ModelsSharedTaskActivityState
}
export class SharedTaskActivityStateApi extends BaseAPI {
    public createSharedTaskActivityState(requestParameters: SharedTaskActivityStateApiCreateSharedTaskActivityStateRequest, options?: any) {
        return SharedTaskActivityStateApiFp(this.configuration).createSharedTaskActivityState(requestParameters.sharedTaskActivityState, options).then((request) => request(this.axios, this.basePath));
    }
    public deleteSharedTaskActivityState(requestParameters: SharedTaskActivityStateApiDeleteSharedTaskActivityStateRequest, options?: any) {
        return SharedTaskActivityStateApiFp(this.configuration).deleteSharedTaskActivityState(requestParameters.id, options).then((request) => request(this.axios, this.basePath));
    }
    public getSharedTaskActivityState(requestParameters: SharedTaskActivityStateApiGetSharedTaskActivityStateRequest, options?: any) {
        return SharedTaskActivityStateApiFp(this.configuration).getSharedTaskActivityState(requestParameters.id, requestParameters.includes, requestParameters.select, options).then((request) => request(this.axios, this.basePath));
    }
    public getSharedTaskActivityStatesBulk(requestParameters: SharedTaskActivityStateApiGetSharedTaskActivityStatesBulkRequest, options?: any) {
        return SharedTaskActivityStateApiFp(this.configuration).getSharedTaskActivityStatesBulk(requestParameters.body, options).then((request) => request(this.axios, this.basePath));
    }
    public getSharedTaskActivityStatesCount(requestParameters: SharedTaskActivityStateApiGetSharedTaskActivityStatesCountRequest = {}, options?: any) {
        return SharedTaskActivityStateApiFp(this.configuration).getSharedTaskActivityStatesCount(requestParameters.includes, requestParameters.where, requestParameters.whereOr, requestParameters.groupBy, requestParameters.limit, requestParameters.page, requestParameters.orderBy, requestParameters.orderDirection, requestParameters.select, options).then((request) => request(this.axios, this.basePath));
    }
    public listSharedTaskActivityStates(requestParameters: SharedTaskActivityStateApiListSharedTaskActivityStatesRequest = {}, options?: any) {
        return SharedTaskActivityStateApiFp(this.configuration).listSharedTaskActivityStates(requestParameters.includes, requestParameters.where, requestParameters.whereOr, requestParameters.groupBy, requestParameters.limit, requestParameters.page, requestParameters.orderBy, requestParameters.orderDirection, requestParameters.select, options).then((request) => request(this.axios, this.basePath));
    }
    public updateSharedTaskActivityState(requestParameters: SharedTaskActivityStateApiUpdateSharedTaskActivityStateRequest, options?: any) {
        return SharedTaskActivityStateApiFp(this.configuration).updateSharedTaskActivityState(requestParameters.id, requestParameters.sharedTaskActivityState, options).then((request) => request(this.axios, this.basePath));
    }
}
