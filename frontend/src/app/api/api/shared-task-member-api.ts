import globalAxios, { AxiosPromise, AxiosInstance } from 'axios';
import { Configuration } from '../configuration';
// Some imports not used depending on template conditions
// @ts-ignore
import { BASE_PATH, COLLECTION_FORMATS, RequestArgs, BaseAPI, RequiredError } from '../base';
// @ts-ignore
import { CrudcontrollersBulkFetchByIdsGetRequest } from '../models';
// @ts-ignore
import { ModelsSharedTaskMember } from '../models';
export const SharedTaskMemberApiAxiosParamCreator = function (configuration?: Configuration) {
    return {
        createSharedTaskMember: async (sharedTaskMember: ModelsSharedTaskMember, options: any = {}): Promise<RequestArgs> => {
            if (sharedTaskMember === null || sharedTaskMember === undefined) {
                throw new RequiredError('sharedTaskMember','Required parameter sharedTaskMember was null or undefined when calling createSharedTaskMember.');
            }
            const localVarPath = `/shared_task_member`;
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
            const nonString = typeof sharedTaskMember !== 'string';
            const needsSerialization = nonString && configuration && configuration.isJsonMime
                ? configuration.isJsonMime(localVarRequestOptions.headers['Content-Type'])
                : nonString;
            localVarRequestOptions.data =  needsSerialization
                ? JSON.stringify(sharedTaskMember !== undefined ? sharedTaskMember : {})
                : (sharedTaskMember || "");
            return {
                url: localVarUrlObj.pathname + localVarUrlObj.search + localVarUrlObj.hash,
                options: localVarRequestOptions,
            };
        },
        deleteSharedTaskMember: async (id: number, options: any = {}): Promise<RequestArgs> => {
            if (id === null || id === undefined) {
                throw new RequiredError('id','Required parameter id was null or undefined when calling deleteSharedTaskMember.');
            }
            const localVarPath = `/shared_task_member/{id}`
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
        getSharedTaskMember: async (id: number, includes?: string, select?: string, options: any = {}): Promise<RequestArgs> => {
            if (id === null || id === undefined) {
                throw new RequiredError('id','Required parameter id was null or undefined when calling getSharedTaskMember.');
            }
            const localVarPath = `/shared_task_member/{id}`
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
        getSharedTaskMembersBulk: async (body: CrudcontrollersBulkFetchByIdsGetRequest, options: any = {}): Promise<RequestArgs> => {
            if (body === null || body === undefined) {
                throw new RequiredError('body','Required parameter body was null or undefined when calling getSharedTaskMembersBulk.');
            }
            const localVarPath = `/shared_task_members/bulk`;
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
        listSharedTaskMembers: async (includes?: string, where?: string, whereOr?: string, groupBy?: string, limit?: string, page?: number, orderBy?: string, orderDirection?: string, select?: string, options: any = {}): Promise<RequestArgs> => {
            const localVarPath = `/shared_task_members`;
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
        updateSharedTaskMember: async (id: number, sharedTaskMember: ModelsSharedTaskMember, options: any = {}): Promise<RequestArgs> => {
            if (id === null || id === undefined) {
                throw new RequiredError('id','Required parameter id was null or undefined when calling updateSharedTaskMember.');
            }
            if (sharedTaskMember === null || sharedTaskMember === undefined) {
                throw new RequiredError('sharedTaskMember','Required parameter sharedTaskMember was null or undefined when calling updateSharedTaskMember.');
            }
            const localVarPath = `/shared_task_member/{id}`
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
            const nonString = typeof sharedTaskMember !== 'string';
            const needsSerialization = nonString && configuration && configuration.isJsonMime
                ? configuration.isJsonMime(localVarRequestOptions.headers['Content-Type'])
                : nonString;
            localVarRequestOptions.data =  needsSerialization
                ? JSON.stringify(sharedTaskMember !== undefined ? sharedTaskMember : {})
                : (sharedTaskMember || "");
            return {
                url: localVarUrlObj.pathname + localVarUrlObj.search + localVarUrlObj.hash,
                options: localVarRequestOptions,
            };
        },
    }
};
export const SharedTaskMemberApiFp = function(configuration?: Configuration) {
    return {
        async createSharedTaskMember(sharedTaskMember: ModelsSharedTaskMember, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsSharedTaskMember>>> {
            const localVarAxiosArgs = await SharedTaskMemberApiAxiosParamCreator(configuration).createSharedTaskMember(sharedTaskMember, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        async deleteSharedTaskMember(id: number, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<string>> {
            const localVarAxiosArgs = await SharedTaskMemberApiAxiosParamCreator(configuration).deleteSharedTaskMember(id, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        async getSharedTaskMember(id: number, includes?: string, select?: string, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsSharedTaskMember>>> {
            const localVarAxiosArgs = await SharedTaskMemberApiAxiosParamCreator(configuration).getSharedTaskMember(id, includes, select, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        async getSharedTaskMembersBulk(body: CrudcontrollersBulkFetchByIdsGetRequest, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsSharedTaskMember>>> {
            const localVarAxiosArgs = await SharedTaskMemberApiAxiosParamCreator(configuration).getSharedTaskMembersBulk(body, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        async listSharedTaskMembers(includes?: string, where?: string, whereOr?: string, groupBy?: string, limit?: string, page?: number, orderBy?: string, orderDirection?: string, select?: string, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsSharedTaskMember>>> {
            const localVarAxiosArgs = await SharedTaskMemberApiAxiosParamCreator(configuration).listSharedTaskMembers(includes, where, whereOr, groupBy, limit, page, orderBy, orderDirection, select, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        async updateSharedTaskMember(id: number, sharedTaskMember: ModelsSharedTaskMember, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsSharedTaskMember>>> {
            const localVarAxiosArgs = await SharedTaskMemberApiAxiosParamCreator(configuration).updateSharedTaskMember(id, sharedTaskMember, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
    }
};
export const SharedTaskMemberApiFactory = function (configuration?: Configuration, basePath?: string, axios?: AxiosInstance) {
    return {
        createSharedTaskMember(sharedTaskMember: ModelsSharedTaskMember, options?: any): AxiosPromise<Array<ModelsSharedTaskMember>> {
            return SharedTaskMemberApiFp(configuration).createSharedTaskMember(sharedTaskMember, options).then((request) => request(axios, basePath));
        },
        deleteSharedTaskMember(id: number, options?: any): AxiosPromise<string> {
            return SharedTaskMemberApiFp(configuration).deleteSharedTaskMember(id, options).then((request) => request(axios, basePath));
        },
        getSharedTaskMember(id: number, includes?: string, select?: string, options?: any): AxiosPromise<Array<ModelsSharedTaskMember>> {
            return SharedTaskMemberApiFp(configuration).getSharedTaskMember(id, includes, select, options).then((request) => request(axios, basePath));
        },
        getSharedTaskMembersBulk(body: CrudcontrollersBulkFetchByIdsGetRequest, options?: any): AxiosPromise<Array<ModelsSharedTaskMember>> {
            return SharedTaskMemberApiFp(configuration).getSharedTaskMembersBulk(body, options).then((request) => request(axios, basePath));
        },
        listSharedTaskMembers(includes?: string, where?: string, whereOr?: string, groupBy?: string, limit?: string, page?: number, orderBy?: string, orderDirection?: string, select?: string, options?: any): AxiosPromise<Array<ModelsSharedTaskMember>> {
            return SharedTaskMemberApiFp(configuration).listSharedTaskMembers(includes, where, whereOr, groupBy, limit, page, orderBy, orderDirection, select, options).then((request) => request(axios, basePath));
        },
        updateSharedTaskMember(id: number, sharedTaskMember: ModelsSharedTaskMember, options?: any): AxiosPromise<Array<ModelsSharedTaskMember>> {
            return SharedTaskMemberApiFp(configuration).updateSharedTaskMember(id, sharedTaskMember, options).then((request) => request(axios, basePath));
        },
    };
};
export interface SharedTaskMemberApiCreateSharedTaskMemberRequest {
    readonly sharedTaskMember: ModelsSharedTaskMember
}
export interface SharedTaskMemberApiDeleteSharedTaskMemberRequest {
    readonly id: number
}
export interface SharedTaskMemberApiGetSharedTaskMemberRequest {
    readonly id: number
    readonly includes?: string
    readonly select?: string
}
export interface SharedTaskMemberApiGetSharedTaskMembersBulkRequest {
    readonly body: CrudcontrollersBulkFetchByIdsGetRequest
}
export interface SharedTaskMemberApiListSharedTaskMembersRequest {
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
export interface SharedTaskMemberApiUpdateSharedTaskMemberRequest {
    readonly id: number
    readonly sharedTaskMember: ModelsSharedTaskMember
}
export class SharedTaskMemberApi extends BaseAPI {
    public createSharedTaskMember(requestParameters: SharedTaskMemberApiCreateSharedTaskMemberRequest, options?: any) {
        return SharedTaskMemberApiFp(this.configuration).createSharedTaskMember(requestParameters.sharedTaskMember, options).then((request) => request(this.axios, this.basePath));
    }
    public deleteSharedTaskMember(requestParameters: SharedTaskMemberApiDeleteSharedTaskMemberRequest, options?: any) {
        return SharedTaskMemberApiFp(this.configuration).deleteSharedTaskMember(requestParameters.id, options).then((request) => request(this.axios, this.basePath));
    }
    public getSharedTaskMember(requestParameters: SharedTaskMemberApiGetSharedTaskMemberRequest, options?: any) {
        return SharedTaskMemberApiFp(this.configuration).getSharedTaskMember(requestParameters.id, requestParameters.includes, requestParameters.select, options).then((request) => request(this.axios, this.basePath));
    }
    public getSharedTaskMembersBulk(requestParameters: SharedTaskMemberApiGetSharedTaskMembersBulkRequest, options?: any) {
        return SharedTaskMemberApiFp(this.configuration).getSharedTaskMembersBulk(requestParameters.body, options).then((request) => request(this.axios, this.basePath));
    }
    public listSharedTaskMembers(requestParameters: SharedTaskMemberApiListSharedTaskMembersRequest = {}, options?: any) {
        return SharedTaskMemberApiFp(this.configuration).listSharedTaskMembers(requestParameters.includes, requestParameters.where, requestParameters.whereOr, requestParameters.groupBy, requestParameters.limit, requestParameters.page, requestParameters.orderBy, requestParameters.orderDirection, requestParameters.select, options).then((request) => request(this.axios, this.basePath));
    }
    public updateSharedTaskMember(requestParameters: SharedTaskMemberApiUpdateSharedTaskMemberRequest, options?: any) {
        return SharedTaskMemberApiFp(this.configuration).updateSharedTaskMember(requestParameters.id, requestParameters.sharedTaskMember, options).then((request) => request(this.axios, this.basePath));
    }
}
