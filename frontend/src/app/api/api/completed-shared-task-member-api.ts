import globalAxios, { AxiosPromise, AxiosInstance } from 'axios';
import { Configuration } from '../configuration';
// Some imports not used depending on template conditions
// @ts-ignore
import { BASE_PATH, COLLECTION_FORMATS, RequestArgs, BaseAPI, RequiredError } from '../base';
// @ts-ignore
import { CrudcontrollersBulkFetchByIdsGetRequest } from '../models';
// @ts-ignore
import { ModelsCompletedSharedTaskMember } from '../models';
export const CompletedSharedTaskMemberApiAxiosParamCreator = function (configuration?: Configuration) {
    return {
        createCompletedSharedTaskMember: async (completedSharedTaskMember: ModelsCompletedSharedTaskMember, options: any = {}): Promise<RequestArgs> => {
            if (completedSharedTaskMember === null || completedSharedTaskMember === undefined) {
                throw new RequiredError('completedSharedTaskMember','Required parameter completedSharedTaskMember was null or undefined when calling createCompletedSharedTaskMember.');
            }
            const localVarPath = `/completed_shared_task_member`;
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
            const nonString = typeof completedSharedTaskMember !== 'string';
            const needsSerialization = nonString && configuration && configuration.isJsonMime
                ? configuration.isJsonMime(localVarRequestOptions.headers['Content-Type'])
                : nonString;
            localVarRequestOptions.data =  needsSerialization
                ? JSON.stringify(completedSharedTaskMember !== undefined ? completedSharedTaskMember : {})
                : (completedSharedTaskMember || "");
            return {
                url: localVarUrlObj.pathname + localVarUrlObj.search + localVarUrlObj.hash,
                options: localVarRequestOptions,
            };
        },
        deleteCompletedSharedTaskMember: async (id: number, options: any = {}): Promise<RequestArgs> => {
            if (id === null || id === undefined) {
                throw new RequiredError('id','Required parameter id was null or undefined when calling deleteCompletedSharedTaskMember.');
            }
            const localVarPath = `/completed_shared_task_member/{id}`
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
        getCompletedSharedTaskMember: async (id: number, includes?: string, select?: string, options: any = {}): Promise<RequestArgs> => {
            if (id === null || id === undefined) {
                throw new RequiredError('id','Required parameter id was null or undefined when calling getCompletedSharedTaskMember.');
            }
            const localVarPath = `/completed_shared_task_member/{id}`
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
        getCompletedSharedTaskMembersBulk: async (body: CrudcontrollersBulkFetchByIdsGetRequest, options: any = {}): Promise<RequestArgs> => {
            if (body === null || body === undefined) {
                throw new RequiredError('body','Required parameter body was null or undefined when calling getCompletedSharedTaskMembersBulk.');
            }
            const localVarPath = `/completed_shared_task_members/bulk`;
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
        getCompletedSharedTaskMembersCount: async (includes?: string, where?: string, whereOr?: string, groupBy?: string, limit?: string, page?: number, orderBy?: string, orderDirection?: string, select?: string, options: any = {}): Promise<RequestArgs> => {
            const localVarPath = `/completed_shared_task_members/count`;
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
        listCompletedSharedTaskMembers: async (includes?: string, where?: string, whereOr?: string, groupBy?: string, limit?: string, page?: number, orderBy?: string, orderDirection?: string, select?: string, options: any = {}): Promise<RequestArgs> => {
            const localVarPath = `/completed_shared_task_members`;
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
        updateCompletedSharedTaskMember: async (id: number, completedSharedTaskMember: ModelsCompletedSharedTaskMember, options: any = {}): Promise<RequestArgs> => {
            if (id === null || id === undefined) {
                throw new RequiredError('id','Required parameter id was null or undefined when calling updateCompletedSharedTaskMember.');
            }
            if (completedSharedTaskMember === null || completedSharedTaskMember === undefined) {
                throw new RequiredError('completedSharedTaskMember','Required parameter completedSharedTaskMember was null or undefined when calling updateCompletedSharedTaskMember.');
            }
            const localVarPath = `/completed_shared_task_member/{id}`
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
            const nonString = typeof completedSharedTaskMember !== 'string';
            const needsSerialization = nonString && configuration && configuration.isJsonMime
                ? configuration.isJsonMime(localVarRequestOptions.headers['Content-Type'])
                : nonString;
            localVarRequestOptions.data =  needsSerialization
                ? JSON.stringify(completedSharedTaskMember !== undefined ? completedSharedTaskMember : {})
                : (completedSharedTaskMember || "");
            return {
                url: localVarUrlObj.pathname + localVarUrlObj.search + localVarUrlObj.hash,
                options: localVarRequestOptions,
            };
        },
    }
};
export const CompletedSharedTaskMemberApiFp = function(configuration?: Configuration) {
    return {
        async createCompletedSharedTaskMember(completedSharedTaskMember: ModelsCompletedSharedTaskMember, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsCompletedSharedTaskMember>>> {
            const localVarAxiosArgs = await CompletedSharedTaskMemberApiAxiosParamCreator(configuration).createCompletedSharedTaskMember(completedSharedTaskMember, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        async deleteCompletedSharedTaskMember(id: number, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<string>> {
            const localVarAxiosArgs = await CompletedSharedTaskMemberApiAxiosParamCreator(configuration).deleteCompletedSharedTaskMember(id, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        async getCompletedSharedTaskMember(id: number, includes?: string, select?: string, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsCompletedSharedTaskMember>>> {
            const localVarAxiosArgs = await CompletedSharedTaskMemberApiAxiosParamCreator(configuration).getCompletedSharedTaskMember(id, includes, select, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        async getCompletedSharedTaskMembersBulk(body: CrudcontrollersBulkFetchByIdsGetRequest, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsCompletedSharedTaskMember>>> {
            const localVarAxiosArgs = await CompletedSharedTaskMemberApiAxiosParamCreator(configuration).getCompletedSharedTaskMembersBulk(body, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        async getCompletedSharedTaskMembersCount(includes?: string, where?: string, whereOr?: string, groupBy?: string, limit?: string, page?: number, orderBy?: string, orderDirection?: string, select?: string, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsCompletedSharedTaskMember>>> {
            const localVarAxiosArgs = await CompletedSharedTaskMemberApiAxiosParamCreator(configuration).getCompletedSharedTaskMembersCount(includes, where, whereOr, groupBy, limit, page, orderBy, orderDirection, select, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        async listCompletedSharedTaskMembers(includes?: string, where?: string, whereOr?: string, groupBy?: string, limit?: string, page?: number, orderBy?: string, orderDirection?: string, select?: string, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsCompletedSharedTaskMember>>> {
            const localVarAxiosArgs = await CompletedSharedTaskMemberApiAxiosParamCreator(configuration).listCompletedSharedTaskMembers(includes, where, whereOr, groupBy, limit, page, orderBy, orderDirection, select, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        async updateCompletedSharedTaskMember(id: number, completedSharedTaskMember: ModelsCompletedSharedTaskMember, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsCompletedSharedTaskMember>>> {
            const localVarAxiosArgs = await CompletedSharedTaskMemberApiAxiosParamCreator(configuration).updateCompletedSharedTaskMember(id, completedSharedTaskMember, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
    }
};
export const CompletedSharedTaskMemberApiFactory = function (configuration?: Configuration, basePath?: string, axios?: AxiosInstance) {
    return {
        createCompletedSharedTaskMember(completedSharedTaskMember: ModelsCompletedSharedTaskMember, options?: any): AxiosPromise<Array<ModelsCompletedSharedTaskMember>> {
            return CompletedSharedTaskMemberApiFp(configuration).createCompletedSharedTaskMember(completedSharedTaskMember, options).then((request) => request(axios, basePath));
        },
        deleteCompletedSharedTaskMember(id: number, options?: any): AxiosPromise<string> {
            return CompletedSharedTaskMemberApiFp(configuration).deleteCompletedSharedTaskMember(id, options).then((request) => request(axios, basePath));
        },
        getCompletedSharedTaskMember(id: number, includes?: string, select?: string, options?: any): AxiosPromise<Array<ModelsCompletedSharedTaskMember>> {
            return CompletedSharedTaskMemberApiFp(configuration).getCompletedSharedTaskMember(id, includes, select, options).then((request) => request(axios, basePath));
        },
        getCompletedSharedTaskMembersBulk(body: CrudcontrollersBulkFetchByIdsGetRequest, options?: any): AxiosPromise<Array<ModelsCompletedSharedTaskMember>> {
            return CompletedSharedTaskMemberApiFp(configuration).getCompletedSharedTaskMembersBulk(body, options).then((request) => request(axios, basePath));
        },
        getCompletedSharedTaskMembersCount(includes?: string, where?: string, whereOr?: string, groupBy?: string, limit?: string, page?: number, orderBy?: string, orderDirection?: string, select?: string, options?: any): AxiosPromise<Array<ModelsCompletedSharedTaskMember>> {
            return CompletedSharedTaskMemberApiFp(configuration).getCompletedSharedTaskMembersCount(includes, where, whereOr, groupBy, limit, page, orderBy, orderDirection, select, options).then((request) => request(axios, basePath));
        },
        listCompletedSharedTaskMembers(includes?: string, where?: string, whereOr?: string, groupBy?: string, limit?: string, page?: number, orderBy?: string, orderDirection?: string, select?: string, options?: any): AxiosPromise<Array<ModelsCompletedSharedTaskMember>> {
            return CompletedSharedTaskMemberApiFp(configuration).listCompletedSharedTaskMembers(includes, where, whereOr, groupBy, limit, page, orderBy, orderDirection, select, options).then((request) => request(axios, basePath));
        },
        updateCompletedSharedTaskMember(id: number, completedSharedTaskMember: ModelsCompletedSharedTaskMember, options?: any): AxiosPromise<Array<ModelsCompletedSharedTaskMember>> {
            return CompletedSharedTaskMemberApiFp(configuration).updateCompletedSharedTaskMember(id, completedSharedTaskMember, options).then((request) => request(axios, basePath));
        },
    };
};
export interface CompletedSharedTaskMemberApiCreateCompletedSharedTaskMemberRequest {
    readonly completedSharedTaskMember: ModelsCompletedSharedTaskMember
}
export interface CompletedSharedTaskMemberApiDeleteCompletedSharedTaskMemberRequest {
    readonly id: number
}
export interface CompletedSharedTaskMemberApiGetCompletedSharedTaskMemberRequest {
    readonly id: number
    readonly includes?: string
    readonly select?: string
}
export interface CompletedSharedTaskMemberApiGetCompletedSharedTaskMembersBulkRequest {
    readonly body: CrudcontrollersBulkFetchByIdsGetRequest
}
export interface CompletedSharedTaskMemberApiGetCompletedSharedTaskMembersCountRequest {
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
export interface CompletedSharedTaskMemberApiListCompletedSharedTaskMembersRequest {
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
export interface CompletedSharedTaskMemberApiUpdateCompletedSharedTaskMemberRequest {
    readonly id: number
    readonly completedSharedTaskMember: ModelsCompletedSharedTaskMember
}
export class CompletedSharedTaskMemberApi extends BaseAPI {
    public createCompletedSharedTaskMember(requestParameters: CompletedSharedTaskMemberApiCreateCompletedSharedTaskMemberRequest, options?: any) {
        return CompletedSharedTaskMemberApiFp(this.configuration).createCompletedSharedTaskMember(requestParameters.completedSharedTaskMember, options).then((request) => request(this.axios, this.basePath));
    }
    public deleteCompletedSharedTaskMember(requestParameters: CompletedSharedTaskMemberApiDeleteCompletedSharedTaskMemberRequest, options?: any) {
        return CompletedSharedTaskMemberApiFp(this.configuration).deleteCompletedSharedTaskMember(requestParameters.id, options).then((request) => request(this.axios, this.basePath));
    }
    public getCompletedSharedTaskMember(requestParameters: CompletedSharedTaskMemberApiGetCompletedSharedTaskMemberRequest, options?: any) {
        return CompletedSharedTaskMemberApiFp(this.configuration).getCompletedSharedTaskMember(requestParameters.id, requestParameters.includes, requestParameters.select, options).then((request) => request(this.axios, this.basePath));
    }
    public getCompletedSharedTaskMembersBulk(requestParameters: CompletedSharedTaskMemberApiGetCompletedSharedTaskMembersBulkRequest, options?: any) {
        return CompletedSharedTaskMemberApiFp(this.configuration).getCompletedSharedTaskMembersBulk(requestParameters.body, options).then((request) => request(this.axios, this.basePath));
    }
    public getCompletedSharedTaskMembersCount(requestParameters: CompletedSharedTaskMemberApiGetCompletedSharedTaskMembersCountRequest = {}, options?: any) {
        return CompletedSharedTaskMemberApiFp(this.configuration).getCompletedSharedTaskMembersCount(requestParameters.includes, requestParameters.where, requestParameters.whereOr, requestParameters.groupBy, requestParameters.limit, requestParameters.page, requestParameters.orderBy, requestParameters.orderDirection, requestParameters.select, options).then((request) => request(this.axios, this.basePath));
    }
    public listCompletedSharedTaskMembers(requestParameters: CompletedSharedTaskMemberApiListCompletedSharedTaskMembersRequest = {}, options?: any) {
        return CompletedSharedTaskMemberApiFp(this.configuration).listCompletedSharedTaskMembers(requestParameters.includes, requestParameters.where, requestParameters.whereOr, requestParameters.groupBy, requestParameters.limit, requestParameters.page, requestParameters.orderBy, requestParameters.orderDirection, requestParameters.select, options).then((request) => request(this.axios, this.basePath));
    }
    public updateCompletedSharedTaskMember(requestParameters: CompletedSharedTaskMemberApiUpdateCompletedSharedTaskMemberRequest, options?: any) {
        return CompletedSharedTaskMemberApiFp(this.configuration).updateCompletedSharedTaskMember(requestParameters.id, requestParameters.completedSharedTaskMember, options).then((request) => request(this.axios, this.basePath));
    }
}
