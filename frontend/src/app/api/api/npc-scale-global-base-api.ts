import globalAxios, { AxiosPromise, AxiosInstance } from 'axios';
import { Configuration } from '../configuration';
// Some imports not used depending on template conditions
// @ts-ignore
import { BASE_PATH, COLLECTION_FORMATS, RequestArgs, BaseAPI, RequiredError } from '../base';
// @ts-ignore
import { CrudcontrollersBulkFetchByIdsGetRequest } from '../models';
// @ts-ignore
import { ModelsNpcScaleGlobalBase } from '../models';
export const NpcScaleGlobalBaseApiAxiosParamCreator = function (configuration?: Configuration) {
    return {
        createNpcScaleGlobalBase: async (npcScaleGlobalBase: ModelsNpcScaleGlobalBase, options: any = {}): Promise<RequestArgs> => {
            if (npcScaleGlobalBase === null || npcScaleGlobalBase === undefined) {
                throw new RequiredError('npcScaleGlobalBase','Required parameter npcScaleGlobalBase was null or undefined when calling createNpcScaleGlobalBase.');
            }
            const localVarPath = `/npc_scale_global_base`;
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
            const nonString = typeof npcScaleGlobalBase !== 'string';
            const needsSerialization = nonString && configuration && configuration.isJsonMime
                ? configuration.isJsonMime(localVarRequestOptions.headers['Content-Type'])
                : nonString;
            localVarRequestOptions.data =  needsSerialization
                ? JSON.stringify(npcScaleGlobalBase !== undefined ? npcScaleGlobalBase : {})
                : (npcScaleGlobalBase || "");
            return {
                url: localVarUrlObj.pathname + localVarUrlObj.search + localVarUrlObj.hash,
                options: localVarRequestOptions,
            };
        },
        deleteNpcScaleGlobalBase: async (id: number, options: any = {}): Promise<RequestArgs> => {
            if (id === null || id === undefined) {
                throw new RequiredError('id','Required parameter id was null or undefined when calling deleteNpcScaleGlobalBase.');
            }
            const localVarPath = `/npc_scale_global_base/{id}`
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
        getNpcScaleGlobalBase: async (id: number, includes?: string, select?: string, options: any = {}): Promise<RequestArgs> => {
            if (id === null || id === undefined) {
                throw new RequiredError('id','Required parameter id was null or undefined when calling getNpcScaleGlobalBase.');
            }
            const localVarPath = `/npc_scale_global_base/{id}`
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
        getNpcScaleGlobalBasesBulk: async (body: CrudcontrollersBulkFetchByIdsGetRequest, options: any = {}): Promise<RequestArgs> => {
            if (body === null || body === undefined) {
                throw new RequiredError('body','Required parameter body was null or undefined when calling getNpcScaleGlobalBasesBulk.');
            }
            const localVarPath = `/npc_scale_global_bases/bulk`;
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
        listNpcScaleGlobalBases: async (includes?: string, where?: string, whereOr?: string, groupBy?: string, limit?: string, page?: number, orderBy?: string, orderDirection?: string, select?: string, options: any = {}): Promise<RequestArgs> => {
            const localVarPath = `/npc_scale_global_bases`;
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
        updateNpcScaleGlobalBase: async (id: number, npcScaleGlobalBase: ModelsNpcScaleGlobalBase, options: any = {}): Promise<RequestArgs> => {
            if (id === null || id === undefined) {
                throw new RequiredError('id','Required parameter id was null or undefined when calling updateNpcScaleGlobalBase.');
            }
            if (npcScaleGlobalBase === null || npcScaleGlobalBase === undefined) {
                throw new RequiredError('npcScaleGlobalBase','Required parameter npcScaleGlobalBase was null or undefined when calling updateNpcScaleGlobalBase.');
            }
            const localVarPath = `/npc_scale_global_base/{id}`
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
            const nonString = typeof npcScaleGlobalBase !== 'string';
            const needsSerialization = nonString && configuration && configuration.isJsonMime
                ? configuration.isJsonMime(localVarRequestOptions.headers['Content-Type'])
                : nonString;
            localVarRequestOptions.data =  needsSerialization
                ? JSON.stringify(npcScaleGlobalBase !== undefined ? npcScaleGlobalBase : {})
                : (npcScaleGlobalBase || "");
            return {
                url: localVarUrlObj.pathname + localVarUrlObj.search + localVarUrlObj.hash,
                options: localVarRequestOptions,
            };
        },
    }
};
export const NpcScaleGlobalBaseApiFp = function(configuration?: Configuration) {
    return {
        async createNpcScaleGlobalBase(npcScaleGlobalBase: ModelsNpcScaleGlobalBase, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsNpcScaleGlobalBase>>> {
            const localVarAxiosArgs = await NpcScaleGlobalBaseApiAxiosParamCreator(configuration).createNpcScaleGlobalBase(npcScaleGlobalBase, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        async deleteNpcScaleGlobalBase(id: number, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<string>> {
            const localVarAxiosArgs = await NpcScaleGlobalBaseApiAxiosParamCreator(configuration).deleteNpcScaleGlobalBase(id, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        async getNpcScaleGlobalBase(id: number, includes?: string, select?: string, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsNpcScaleGlobalBase>>> {
            const localVarAxiosArgs = await NpcScaleGlobalBaseApiAxiosParamCreator(configuration).getNpcScaleGlobalBase(id, includes, select, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        async getNpcScaleGlobalBasesBulk(body: CrudcontrollersBulkFetchByIdsGetRequest, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsNpcScaleGlobalBase>>> {
            const localVarAxiosArgs = await NpcScaleGlobalBaseApiAxiosParamCreator(configuration).getNpcScaleGlobalBasesBulk(body, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        async listNpcScaleGlobalBases(includes?: string, where?: string, whereOr?: string, groupBy?: string, limit?: string, page?: number, orderBy?: string, orderDirection?: string, select?: string, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsNpcScaleGlobalBase>>> {
            const localVarAxiosArgs = await NpcScaleGlobalBaseApiAxiosParamCreator(configuration).listNpcScaleGlobalBases(includes, where, whereOr, groupBy, limit, page, orderBy, orderDirection, select, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        async updateNpcScaleGlobalBase(id: number, npcScaleGlobalBase: ModelsNpcScaleGlobalBase, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsNpcScaleGlobalBase>>> {
            const localVarAxiosArgs = await NpcScaleGlobalBaseApiAxiosParamCreator(configuration).updateNpcScaleGlobalBase(id, npcScaleGlobalBase, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
    }
};
export const NpcScaleGlobalBaseApiFactory = function (configuration?: Configuration, basePath?: string, axios?: AxiosInstance) {
    return {
        createNpcScaleGlobalBase(npcScaleGlobalBase: ModelsNpcScaleGlobalBase, options?: any): AxiosPromise<Array<ModelsNpcScaleGlobalBase>> {
            return NpcScaleGlobalBaseApiFp(configuration).createNpcScaleGlobalBase(npcScaleGlobalBase, options).then((request) => request(axios, basePath));
        },
        deleteNpcScaleGlobalBase(id: number, options?: any): AxiosPromise<string> {
            return NpcScaleGlobalBaseApiFp(configuration).deleteNpcScaleGlobalBase(id, options).then((request) => request(axios, basePath));
        },
        getNpcScaleGlobalBase(id: number, includes?: string, select?: string, options?: any): AxiosPromise<Array<ModelsNpcScaleGlobalBase>> {
            return NpcScaleGlobalBaseApiFp(configuration).getNpcScaleGlobalBase(id, includes, select, options).then((request) => request(axios, basePath));
        },
        getNpcScaleGlobalBasesBulk(body: CrudcontrollersBulkFetchByIdsGetRequest, options?: any): AxiosPromise<Array<ModelsNpcScaleGlobalBase>> {
            return NpcScaleGlobalBaseApiFp(configuration).getNpcScaleGlobalBasesBulk(body, options).then((request) => request(axios, basePath));
        },
        listNpcScaleGlobalBases(includes?: string, where?: string, whereOr?: string, groupBy?: string, limit?: string, page?: number, orderBy?: string, orderDirection?: string, select?: string, options?: any): AxiosPromise<Array<ModelsNpcScaleGlobalBase>> {
            return NpcScaleGlobalBaseApiFp(configuration).listNpcScaleGlobalBases(includes, where, whereOr, groupBy, limit, page, orderBy, orderDirection, select, options).then((request) => request(axios, basePath));
        },
        updateNpcScaleGlobalBase(id: number, npcScaleGlobalBase: ModelsNpcScaleGlobalBase, options?: any): AxiosPromise<Array<ModelsNpcScaleGlobalBase>> {
            return NpcScaleGlobalBaseApiFp(configuration).updateNpcScaleGlobalBase(id, npcScaleGlobalBase, options).then((request) => request(axios, basePath));
        },
    };
};
export interface NpcScaleGlobalBaseApiCreateNpcScaleGlobalBaseRequest {
    readonly npcScaleGlobalBase: ModelsNpcScaleGlobalBase
}
export interface NpcScaleGlobalBaseApiDeleteNpcScaleGlobalBaseRequest {
    readonly id: number
}
export interface NpcScaleGlobalBaseApiGetNpcScaleGlobalBaseRequest {
    readonly id: number
    readonly includes?: string
    readonly select?: string
}
export interface NpcScaleGlobalBaseApiGetNpcScaleGlobalBasesBulkRequest {
    readonly body: CrudcontrollersBulkFetchByIdsGetRequest
}
export interface NpcScaleGlobalBaseApiListNpcScaleGlobalBasesRequest {
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
export interface NpcScaleGlobalBaseApiUpdateNpcScaleGlobalBaseRequest {
    readonly id: number
    readonly npcScaleGlobalBase: ModelsNpcScaleGlobalBase
}
export class NpcScaleGlobalBaseApi extends BaseAPI {
    public createNpcScaleGlobalBase(requestParameters: NpcScaleGlobalBaseApiCreateNpcScaleGlobalBaseRequest, options?: any) {
        return NpcScaleGlobalBaseApiFp(this.configuration).createNpcScaleGlobalBase(requestParameters.npcScaleGlobalBase, options).then((request) => request(this.axios, this.basePath));
    }
    public deleteNpcScaleGlobalBase(requestParameters: NpcScaleGlobalBaseApiDeleteNpcScaleGlobalBaseRequest, options?: any) {
        return NpcScaleGlobalBaseApiFp(this.configuration).deleteNpcScaleGlobalBase(requestParameters.id, options).then((request) => request(this.axios, this.basePath));
    }
    public getNpcScaleGlobalBase(requestParameters: NpcScaleGlobalBaseApiGetNpcScaleGlobalBaseRequest, options?: any) {
        return NpcScaleGlobalBaseApiFp(this.configuration).getNpcScaleGlobalBase(requestParameters.id, requestParameters.includes, requestParameters.select, options).then((request) => request(this.axios, this.basePath));
    }
    public getNpcScaleGlobalBasesBulk(requestParameters: NpcScaleGlobalBaseApiGetNpcScaleGlobalBasesBulkRequest, options?: any) {
        return NpcScaleGlobalBaseApiFp(this.configuration).getNpcScaleGlobalBasesBulk(requestParameters.body, options).then((request) => request(this.axios, this.basePath));
    }
    public listNpcScaleGlobalBases(requestParameters: NpcScaleGlobalBaseApiListNpcScaleGlobalBasesRequest = {}, options?: any) {
        return NpcScaleGlobalBaseApiFp(this.configuration).listNpcScaleGlobalBases(requestParameters.includes, requestParameters.where, requestParameters.whereOr, requestParameters.groupBy, requestParameters.limit, requestParameters.page, requestParameters.orderBy, requestParameters.orderDirection, requestParameters.select, options).then((request) => request(this.axios, this.basePath));
    }
    public updateNpcScaleGlobalBase(requestParameters: NpcScaleGlobalBaseApiUpdateNpcScaleGlobalBaseRequest, options?: any) {
        return NpcScaleGlobalBaseApiFp(this.configuration).updateNpcScaleGlobalBase(requestParameters.id, requestParameters.npcScaleGlobalBase, options).then((request) => request(this.axios, this.basePath));
    }
}
