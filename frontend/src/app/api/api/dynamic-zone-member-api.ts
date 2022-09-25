import globalAxios, { AxiosPromise, AxiosInstance } from 'axios';
import { Configuration } from '../configuration';
// Some imports not used depending on template conditions
// @ts-ignore
import { BASE_PATH, COLLECTION_FORMATS, RequestArgs, BaseAPI, RequiredError } from '../base';
// @ts-ignore
import { CrudcontrollersBulkFetchByIdsGetRequest } from '../models';
// @ts-ignore
import { ModelsDynamicZoneMember } from '../models';
export const DynamicZoneMemberApiAxiosParamCreator = function (configuration?: Configuration) {
    return {
        createDynamicZoneMember: async (dynamicZoneMember: ModelsDynamicZoneMember, options: any = {}): Promise<RequestArgs> => {
            if (dynamicZoneMember === null || dynamicZoneMember === undefined) {
                throw new RequiredError('dynamicZoneMember','Required parameter dynamicZoneMember was null or undefined when calling createDynamicZoneMember.');
            }
            const localVarPath = `/dynamic_zone_member`;
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
        deleteDynamicZoneMember: async (id: number, options: any = {}): Promise<RequestArgs> => {
            if (id === null || id === undefined) {
                throw new RequiredError('id','Required parameter id was null or undefined when calling deleteDynamicZoneMember.');
            }
            const localVarPath = `/dynamic_zone_member/{id}`
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
        getDynamicZoneMember: async (id: number, includes?: string, select?: string, options: any = {}): Promise<RequestArgs> => {
            if (id === null || id === undefined) {
                throw new RequiredError('id','Required parameter id was null or undefined when calling getDynamicZoneMember.');
            }
            const localVarPath = `/dynamic_zone_member/{id}`
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
        getDynamicZoneMembersBulk: async (body: CrudcontrollersBulkFetchByIdsGetRequest, options: any = {}): Promise<RequestArgs> => {
            if (body === null || body === undefined) {
                throw new RequiredError('body','Required parameter body was null or undefined when calling getDynamicZoneMembersBulk.');
            }
            const localVarPath = `/dynamic_zone_members/bulk`;
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
        listDynamicZoneMembers: async (includes?: string, where?: string, whereOr?: string, groupBy?: string, limit?: string, page?: number, orderBy?: string, orderDirection?: string, select?: string, options: any = {}): Promise<RequestArgs> => {
            const localVarPath = `/dynamic_zone_members`;
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
        updateDynamicZoneMember: async (id: number, dynamicZoneMember: ModelsDynamicZoneMember, options: any = {}): Promise<RequestArgs> => {
            if (id === null || id === undefined) {
                throw new RequiredError('id','Required parameter id was null or undefined when calling updateDynamicZoneMember.');
            }
            if (dynamicZoneMember === null || dynamicZoneMember === undefined) {
                throw new RequiredError('dynamicZoneMember','Required parameter dynamicZoneMember was null or undefined when calling updateDynamicZoneMember.');
            }
            const localVarPath = `/dynamic_zone_member/{id}`
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
export const DynamicZoneMemberApiFp = function(configuration?: Configuration) {
    return {
        async createDynamicZoneMember(dynamicZoneMember: ModelsDynamicZoneMember, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsDynamicZoneMember>>> {
            const localVarAxiosArgs = await DynamicZoneMemberApiAxiosParamCreator(configuration).createDynamicZoneMember(dynamicZoneMember, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        async deleteDynamicZoneMember(id: number, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<string>> {
            const localVarAxiosArgs = await DynamicZoneMemberApiAxiosParamCreator(configuration).deleteDynamicZoneMember(id, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        async getDynamicZoneMember(id: number, includes?: string, select?: string, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsDynamicZoneMember>>> {
            const localVarAxiosArgs = await DynamicZoneMemberApiAxiosParamCreator(configuration).getDynamicZoneMember(id, includes, select, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        async getDynamicZoneMembersBulk(body: CrudcontrollersBulkFetchByIdsGetRequest, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsDynamicZoneMember>>> {
            const localVarAxiosArgs = await DynamicZoneMemberApiAxiosParamCreator(configuration).getDynamicZoneMembersBulk(body, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        async listDynamicZoneMembers(includes?: string, where?: string, whereOr?: string, groupBy?: string, limit?: string, page?: number, orderBy?: string, orderDirection?: string, select?: string, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsDynamicZoneMember>>> {
            const localVarAxiosArgs = await DynamicZoneMemberApiAxiosParamCreator(configuration).listDynamicZoneMembers(includes, where, whereOr, groupBy, limit, page, orderBy, orderDirection, select, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        async updateDynamicZoneMember(id: number, dynamicZoneMember: ModelsDynamicZoneMember, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsDynamicZoneMember>>> {
            const localVarAxiosArgs = await DynamicZoneMemberApiAxiosParamCreator(configuration).updateDynamicZoneMember(id, dynamicZoneMember, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
    }
};
export const DynamicZoneMemberApiFactory = function (configuration?: Configuration, basePath?: string, axios?: AxiosInstance) {
    return {
        createDynamicZoneMember(dynamicZoneMember: ModelsDynamicZoneMember, options?: any): AxiosPromise<Array<ModelsDynamicZoneMember>> {
            return DynamicZoneMemberApiFp(configuration).createDynamicZoneMember(dynamicZoneMember, options).then((request) => request(axios, basePath));
        },
        deleteDynamicZoneMember(id: number, options?: any): AxiosPromise<string> {
            return DynamicZoneMemberApiFp(configuration).deleteDynamicZoneMember(id, options).then((request) => request(axios, basePath));
        },
        getDynamicZoneMember(id: number, includes?: string, select?: string, options?: any): AxiosPromise<Array<ModelsDynamicZoneMember>> {
            return DynamicZoneMemberApiFp(configuration).getDynamicZoneMember(id, includes, select, options).then((request) => request(axios, basePath));
        },
        getDynamicZoneMembersBulk(body: CrudcontrollersBulkFetchByIdsGetRequest, options?: any): AxiosPromise<Array<ModelsDynamicZoneMember>> {
            return DynamicZoneMemberApiFp(configuration).getDynamicZoneMembersBulk(body, options).then((request) => request(axios, basePath));
        },
        listDynamicZoneMembers(includes?: string, where?: string, whereOr?: string, groupBy?: string, limit?: string, page?: number, orderBy?: string, orderDirection?: string, select?: string, options?: any): AxiosPromise<Array<ModelsDynamicZoneMember>> {
            return DynamicZoneMemberApiFp(configuration).listDynamicZoneMembers(includes, where, whereOr, groupBy, limit, page, orderBy, orderDirection, select, options).then((request) => request(axios, basePath));
        },
        updateDynamicZoneMember(id: number, dynamicZoneMember: ModelsDynamicZoneMember, options?: any): AxiosPromise<Array<ModelsDynamicZoneMember>> {
            return DynamicZoneMemberApiFp(configuration).updateDynamicZoneMember(id, dynamicZoneMember, options).then((request) => request(axios, basePath));
        },
    };
};
export interface DynamicZoneMemberApiCreateDynamicZoneMemberRequest {
    readonly dynamicZoneMember: ModelsDynamicZoneMember
}
export interface DynamicZoneMemberApiDeleteDynamicZoneMemberRequest {
    readonly id: number
}
export interface DynamicZoneMemberApiGetDynamicZoneMemberRequest {
    readonly id: number
    readonly includes?: string
    readonly select?: string
}
export interface DynamicZoneMemberApiGetDynamicZoneMembersBulkRequest {
    readonly body: CrudcontrollersBulkFetchByIdsGetRequest
}
export interface DynamicZoneMemberApiListDynamicZoneMembersRequest {
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
export interface DynamicZoneMemberApiUpdateDynamicZoneMemberRequest {
    readonly id: number
    readonly dynamicZoneMember: ModelsDynamicZoneMember
}
export class DynamicZoneMemberApi extends BaseAPI {
    public createDynamicZoneMember(requestParameters: DynamicZoneMemberApiCreateDynamicZoneMemberRequest, options?: any) {
        return DynamicZoneMemberApiFp(this.configuration).createDynamicZoneMember(requestParameters.dynamicZoneMember, options).then((request) => request(this.axios, this.basePath));
    }
    public deleteDynamicZoneMember(requestParameters: DynamicZoneMemberApiDeleteDynamicZoneMemberRequest, options?: any) {
        return DynamicZoneMemberApiFp(this.configuration).deleteDynamicZoneMember(requestParameters.id, options).then((request) => request(this.axios, this.basePath));
    }
    public getDynamicZoneMember(requestParameters: DynamicZoneMemberApiGetDynamicZoneMemberRequest, options?: any) {
        return DynamicZoneMemberApiFp(this.configuration).getDynamicZoneMember(requestParameters.id, requestParameters.includes, requestParameters.select, options).then((request) => request(this.axios, this.basePath));
    }
    public getDynamicZoneMembersBulk(requestParameters: DynamicZoneMemberApiGetDynamicZoneMembersBulkRequest, options?: any) {
        return DynamicZoneMemberApiFp(this.configuration).getDynamicZoneMembersBulk(requestParameters.body, options).then((request) => request(this.axios, this.basePath));
    }
    public listDynamicZoneMembers(requestParameters: DynamicZoneMemberApiListDynamicZoneMembersRequest = {}, options?: any) {
        return DynamicZoneMemberApiFp(this.configuration).listDynamicZoneMembers(requestParameters.includes, requestParameters.where, requestParameters.whereOr, requestParameters.groupBy, requestParameters.limit, requestParameters.page, requestParameters.orderBy, requestParameters.orderDirection, requestParameters.select, options).then((request) => request(this.axios, this.basePath));
    }
    public updateDynamicZoneMember(requestParameters: DynamicZoneMemberApiUpdateDynamicZoneMemberRequest, options?: any) {
        return DynamicZoneMemberApiFp(this.configuration).updateDynamicZoneMember(requestParameters.id, requestParameters.dynamicZoneMember, options).then((request) => request(this.axios, this.basePath));
    }
}
