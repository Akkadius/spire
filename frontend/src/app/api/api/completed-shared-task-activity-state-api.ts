import globalAxios, { AxiosPromise, AxiosInstance } from 'axios';
import { Configuration } from '../configuration';
// Some imports not used depending on template conditions
// @ts-ignore
import { BASE_PATH, COLLECTION_FORMATS, RequestArgs, BaseAPI, RequiredError } from '../base';
// @ts-ignore
import { CrudcontrollersBulkFetchByIdsGetRequest } from '../models';
// @ts-ignore
import { ModelsCompletedSharedTaskActivityState } from '../models';
export const CompletedSharedTaskActivityStateApiAxiosParamCreator = function (configuration?: Configuration) {
    return {
        createCompletedSharedTaskActivityState: async (completedSharedTaskActivityState: ModelsCompletedSharedTaskActivityState, options: any = {}): Promise<RequestArgs> => {
            if (completedSharedTaskActivityState === null || completedSharedTaskActivityState === undefined) {
                throw new RequiredError('completedSharedTaskActivityState','Required parameter completedSharedTaskActivityState was null or undefined when calling createCompletedSharedTaskActivityState.');
            }
            const localVarPath = `/completed_shared_task_activity_state`;
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
            const nonString = typeof completedSharedTaskActivityState !== 'string';
            const needsSerialization = nonString && configuration && configuration.isJsonMime
                ? configuration.isJsonMime(localVarRequestOptions.headers['Content-Type'])
                : nonString;
            localVarRequestOptions.data =  needsSerialization
                ? JSON.stringify(completedSharedTaskActivityState !== undefined ? completedSharedTaskActivityState : {})
                : (completedSharedTaskActivityState || "");
            return {
                url: localVarUrlObj.pathname + localVarUrlObj.search + localVarUrlObj.hash,
                options: localVarRequestOptions,
            };
        },
        deleteCompletedSharedTaskActivityState: async (id: number, options: any = {}): Promise<RequestArgs> => {
            if (id === null || id === undefined) {
                throw new RequiredError('id','Required parameter id was null or undefined when calling deleteCompletedSharedTaskActivityState.');
            }
            const localVarPath = `/completed_shared_task_activity_state/{id}`
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
        getCompletedSharedTaskActivityState: async (id: number, includes?: string, select?: string, options: any = {}): Promise<RequestArgs> => {
            if (id === null || id === undefined) {
                throw new RequiredError('id','Required parameter id was null or undefined when calling getCompletedSharedTaskActivityState.');
            }
            const localVarPath = `/completed_shared_task_activity_state/{id}`
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
        getCompletedSharedTaskActivityStatesBulk: async (body: CrudcontrollersBulkFetchByIdsGetRequest, options: any = {}): Promise<RequestArgs> => {
            if (body === null || body === undefined) {
                throw new RequiredError('body','Required parameter body was null or undefined when calling getCompletedSharedTaskActivityStatesBulk.');
            }
            const localVarPath = `/completed_shared_task_activity_states/bulk`;
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
        getCompletedSharedTaskActivityStatesCount: async (includes?: string, where?: string, whereOr?: string, groupBy?: string, limit?: string, page?: number, orderBy?: string, orderDirection?: string, select?: string, options: any = {}): Promise<RequestArgs> => {
            const localVarPath = `/completed_shared_task_activity_states/count`;
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
        listCompletedSharedTaskActivityStates: async (includes?: string, where?: string, whereOr?: string, groupBy?: string, limit?: string, page?: number, orderBy?: string, orderDirection?: string, select?: string, options: any = {}): Promise<RequestArgs> => {
            const localVarPath = `/completed_shared_task_activity_states`;
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
        updateCompletedSharedTaskActivityState: async (id: number, completedSharedTaskActivityState: ModelsCompletedSharedTaskActivityState, options: any = {}): Promise<RequestArgs> => {
            if (id === null || id === undefined) {
                throw new RequiredError('id','Required parameter id was null or undefined when calling updateCompletedSharedTaskActivityState.');
            }
            if (completedSharedTaskActivityState === null || completedSharedTaskActivityState === undefined) {
                throw new RequiredError('completedSharedTaskActivityState','Required parameter completedSharedTaskActivityState was null or undefined when calling updateCompletedSharedTaskActivityState.');
            }
            const localVarPath = `/completed_shared_task_activity_state/{id}`
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
            const nonString = typeof completedSharedTaskActivityState !== 'string';
            const needsSerialization = nonString && configuration && configuration.isJsonMime
                ? configuration.isJsonMime(localVarRequestOptions.headers['Content-Type'])
                : nonString;
            localVarRequestOptions.data =  needsSerialization
                ? JSON.stringify(completedSharedTaskActivityState !== undefined ? completedSharedTaskActivityState : {})
                : (completedSharedTaskActivityState || "");
            return {
                url: localVarUrlObj.pathname + localVarUrlObj.search + localVarUrlObj.hash,
                options: localVarRequestOptions,
            };
        },
    }
};
export const CompletedSharedTaskActivityStateApiFp = function(configuration?: Configuration) {
    return {
        async createCompletedSharedTaskActivityState(completedSharedTaskActivityState: ModelsCompletedSharedTaskActivityState, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsCompletedSharedTaskActivityState>>> {
            const localVarAxiosArgs = await CompletedSharedTaskActivityStateApiAxiosParamCreator(configuration).createCompletedSharedTaskActivityState(completedSharedTaskActivityState, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        async deleteCompletedSharedTaskActivityState(id: number, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<string>> {
            const localVarAxiosArgs = await CompletedSharedTaskActivityStateApiAxiosParamCreator(configuration).deleteCompletedSharedTaskActivityState(id, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        async getCompletedSharedTaskActivityState(id: number, includes?: string, select?: string, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsCompletedSharedTaskActivityState>>> {
            const localVarAxiosArgs = await CompletedSharedTaskActivityStateApiAxiosParamCreator(configuration).getCompletedSharedTaskActivityState(id, includes, select, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        async getCompletedSharedTaskActivityStatesBulk(body: CrudcontrollersBulkFetchByIdsGetRequest, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsCompletedSharedTaskActivityState>>> {
            const localVarAxiosArgs = await CompletedSharedTaskActivityStateApiAxiosParamCreator(configuration).getCompletedSharedTaskActivityStatesBulk(body, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        async getCompletedSharedTaskActivityStatesCount(includes?: string, where?: string, whereOr?: string, groupBy?: string, limit?: string, page?: number, orderBy?: string, orderDirection?: string, select?: string, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsCompletedSharedTaskActivityState>>> {
            const localVarAxiosArgs = await CompletedSharedTaskActivityStateApiAxiosParamCreator(configuration).getCompletedSharedTaskActivityStatesCount(includes, where, whereOr, groupBy, limit, page, orderBy, orderDirection, select, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        async listCompletedSharedTaskActivityStates(includes?: string, where?: string, whereOr?: string, groupBy?: string, limit?: string, page?: number, orderBy?: string, orderDirection?: string, select?: string, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsCompletedSharedTaskActivityState>>> {
            const localVarAxiosArgs = await CompletedSharedTaskActivityStateApiAxiosParamCreator(configuration).listCompletedSharedTaskActivityStates(includes, where, whereOr, groupBy, limit, page, orderBy, orderDirection, select, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        async updateCompletedSharedTaskActivityState(id: number, completedSharedTaskActivityState: ModelsCompletedSharedTaskActivityState, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsCompletedSharedTaskActivityState>>> {
            const localVarAxiosArgs = await CompletedSharedTaskActivityStateApiAxiosParamCreator(configuration).updateCompletedSharedTaskActivityState(id, completedSharedTaskActivityState, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
    }
};
export const CompletedSharedTaskActivityStateApiFactory = function (configuration?: Configuration, basePath?: string, axios?: AxiosInstance) {
    return {
        createCompletedSharedTaskActivityState(completedSharedTaskActivityState: ModelsCompletedSharedTaskActivityState, options?: any): AxiosPromise<Array<ModelsCompletedSharedTaskActivityState>> {
            return CompletedSharedTaskActivityStateApiFp(configuration).createCompletedSharedTaskActivityState(completedSharedTaskActivityState, options).then((request) => request(axios, basePath));
        },
        deleteCompletedSharedTaskActivityState(id: number, options?: any): AxiosPromise<string> {
            return CompletedSharedTaskActivityStateApiFp(configuration).deleteCompletedSharedTaskActivityState(id, options).then((request) => request(axios, basePath));
        },
        getCompletedSharedTaskActivityState(id: number, includes?: string, select?: string, options?: any): AxiosPromise<Array<ModelsCompletedSharedTaskActivityState>> {
            return CompletedSharedTaskActivityStateApiFp(configuration).getCompletedSharedTaskActivityState(id, includes, select, options).then((request) => request(axios, basePath));
        },
        getCompletedSharedTaskActivityStatesBulk(body: CrudcontrollersBulkFetchByIdsGetRequest, options?: any): AxiosPromise<Array<ModelsCompletedSharedTaskActivityState>> {
            return CompletedSharedTaskActivityStateApiFp(configuration).getCompletedSharedTaskActivityStatesBulk(body, options).then((request) => request(axios, basePath));
        },
        getCompletedSharedTaskActivityStatesCount(includes?: string, where?: string, whereOr?: string, groupBy?: string, limit?: string, page?: number, orderBy?: string, orderDirection?: string, select?: string, options?: any): AxiosPromise<Array<ModelsCompletedSharedTaskActivityState>> {
            return CompletedSharedTaskActivityStateApiFp(configuration).getCompletedSharedTaskActivityStatesCount(includes, where, whereOr, groupBy, limit, page, orderBy, orderDirection, select, options).then((request) => request(axios, basePath));
        },
        listCompletedSharedTaskActivityStates(includes?: string, where?: string, whereOr?: string, groupBy?: string, limit?: string, page?: number, orderBy?: string, orderDirection?: string, select?: string, options?: any): AxiosPromise<Array<ModelsCompletedSharedTaskActivityState>> {
            return CompletedSharedTaskActivityStateApiFp(configuration).listCompletedSharedTaskActivityStates(includes, where, whereOr, groupBy, limit, page, orderBy, orderDirection, select, options).then((request) => request(axios, basePath));
        },
        updateCompletedSharedTaskActivityState(id: number, completedSharedTaskActivityState: ModelsCompletedSharedTaskActivityState, options?: any): AxiosPromise<Array<ModelsCompletedSharedTaskActivityState>> {
            return CompletedSharedTaskActivityStateApiFp(configuration).updateCompletedSharedTaskActivityState(id, completedSharedTaskActivityState, options).then((request) => request(axios, basePath));
        },
    };
};
export interface CompletedSharedTaskActivityStateApiCreateCompletedSharedTaskActivityStateRequest {
    readonly completedSharedTaskActivityState: ModelsCompletedSharedTaskActivityState
}
export interface CompletedSharedTaskActivityStateApiDeleteCompletedSharedTaskActivityStateRequest {
    readonly id: number
}
export interface CompletedSharedTaskActivityStateApiGetCompletedSharedTaskActivityStateRequest {
    readonly id: number
    readonly includes?: string
    readonly select?: string
}
export interface CompletedSharedTaskActivityStateApiGetCompletedSharedTaskActivityStatesBulkRequest {
    readonly body: CrudcontrollersBulkFetchByIdsGetRequest
}
export interface CompletedSharedTaskActivityStateApiGetCompletedSharedTaskActivityStatesCountRequest {
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
export interface CompletedSharedTaskActivityStateApiListCompletedSharedTaskActivityStatesRequest {
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
export interface CompletedSharedTaskActivityStateApiUpdateCompletedSharedTaskActivityStateRequest {
    readonly id: number
    readonly completedSharedTaskActivityState: ModelsCompletedSharedTaskActivityState
}
export class CompletedSharedTaskActivityStateApi extends BaseAPI {
    public createCompletedSharedTaskActivityState(requestParameters: CompletedSharedTaskActivityStateApiCreateCompletedSharedTaskActivityStateRequest, options?: any) {
        return CompletedSharedTaskActivityStateApiFp(this.configuration).createCompletedSharedTaskActivityState(requestParameters.completedSharedTaskActivityState, options).then((request) => request(this.axios, this.basePath));
    }
    public deleteCompletedSharedTaskActivityState(requestParameters: CompletedSharedTaskActivityStateApiDeleteCompletedSharedTaskActivityStateRequest, options?: any) {
        return CompletedSharedTaskActivityStateApiFp(this.configuration).deleteCompletedSharedTaskActivityState(requestParameters.id, options).then((request) => request(this.axios, this.basePath));
    }
    public getCompletedSharedTaskActivityState(requestParameters: CompletedSharedTaskActivityStateApiGetCompletedSharedTaskActivityStateRequest, options?: any) {
        return CompletedSharedTaskActivityStateApiFp(this.configuration).getCompletedSharedTaskActivityState(requestParameters.id, requestParameters.includes, requestParameters.select, options).then((request) => request(this.axios, this.basePath));
    }
    public getCompletedSharedTaskActivityStatesBulk(requestParameters: CompletedSharedTaskActivityStateApiGetCompletedSharedTaskActivityStatesBulkRequest, options?: any) {
        return CompletedSharedTaskActivityStateApiFp(this.configuration).getCompletedSharedTaskActivityStatesBulk(requestParameters.body, options).then((request) => request(this.axios, this.basePath));
    }
    public getCompletedSharedTaskActivityStatesCount(requestParameters: CompletedSharedTaskActivityStateApiGetCompletedSharedTaskActivityStatesCountRequest = {}, options?: any) {
        return CompletedSharedTaskActivityStateApiFp(this.configuration).getCompletedSharedTaskActivityStatesCount(requestParameters.includes, requestParameters.where, requestParameters.whereOr, requestParameters.groupBy, requestParameters.limit, requestParameters.page, requestParameters.orderBy, requestParameters.orderDirection, requestParameters.select, options).then((request) => request(this.axios, this.basePath));
    }
    public listCompletedSharedTaskActivityStates(requestParameters: CompletedSharedTaskActivityStateApiListCompletedSharedTaskActivityStatesRequest = {}, options?: any) {
        return CompletedSharedTaskActivityStateApiFp(this.configuration).listCompletedSharedTaskActivityStates(requestParameters.includes, requestParameters.where, requestParameters.whereOr, requestParameters.groupBy, requestParameters.limit, requestParameters.page, requestParameters.orderBy, requestParameters.orderDirection, requestParameters.select, options).then((request) => request(this.axios, this.basePath));
    }
    public updateCompletedSharedTaskActivityState(requestParameters: CompletedSharedTaskActivityStateApiUpdateCompletedSharedTaskActivityStateRequest, options?: any) {
        return CompletedSharedTaskActivityStateApiFp(this.configuration).updateCompletedSharedTaskActivityState(requestParameters.id, requestParameters.completedSharedTaskActivityState, options).then((request) => request(this.axios, this.basePath));
    }
}
