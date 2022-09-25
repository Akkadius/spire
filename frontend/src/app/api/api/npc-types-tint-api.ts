import globalAxios, { AxiosPromise, AxiosInstance } from 'axios';
import { Configuration } from '../configuration';
// Some imports not used depending on template conditions
// @ts-ignore
import { BASE_PATH, COLLECTION_FORMATS, RequestArgs, BaseAPI, RequiredError } from '../base';
// @ts-ignore
import { CrudcontrollersBulkFetchByIdsGetRequest } from '../models';
// @ts-ignore
import { ModelsNpcTypesTint } from '../models';
export const NpcTypesTintApiAxiosParamCreator = function (configuration?: Configuration) {
    return {
        createNpcTypesTint: async (npcTypesTint: ModelsNpcTypesTint, options: any = {}): Promise<RequestArgs> => {
            if (npcTypesTint === null || npcTypesTint === undefined) {
                throw new RequiredError('npcTypesTint','Required parameter npcTypesTint was null or undefined when calling createNpcTypesTint.');
            }
            const localVarPath = `/npc_types_tint`;
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
            const nonString = typeof npcTypesTint !== 'string';
            const needsSerialization = nonString && configuration && configuration.isJsonMime
                ? configuration.isJsonMime(localVarRequestOptions.headers['Content-Type'])
                : nonString;
            localVarRequestOptions.data =  needsSerialization
                ? JSON.stringify(npcTypesTint !== undefined ? npcTypesTint : {})
                : (npcTypesTint || "");
            return {
                url: localVarUrlObj.pathname + localVarUrlObj.search + localVarUrlObj.hash,
                options: localVarRequestOptions,
            };
        },
        deleteNpcTypesTint: async (id: number, options: any = {}): Promise<RequestArgs> => {
            if (id === null || id === undefined) {
                throw new RequiredError('id','Required parameter id was null or undefined when calling deleteNpcTypesTint.');
            }
            const localVarPath = `/npc_types_tint/{id}`
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
        getNpcTypesTint: async (id: number, includes?: string, select?: string, options: any = {}): Promise<RequestArgs> => {
            if (id === null || id === undefined) {
                throw new RequiredError('id','Required parameter id was null or undefined when calling getNpcTypesTint.');
            }
            const localVarPath = `/npc_types_tint/{id}`
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
        getNpcTypesTintsBulk: async (body: CrudcontrollersBulkFetchByIdsGetRequest, options: any = {}): Promise<RequestArgs> => {
            if (body === null || body === undefined) {
                throw new RequiredError('body','Required parameter body was null or undefined when calling getNpcTypesTintsBulk.');
            }
            const localVarPath = `/npc_types_tints/bulk`;
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
        listNpcTypesTints: async (includes?: string, where?: string, whereOr?: string, groupBy?: string, limit?: string, page?: number, orderBy?: string, orderDirection?: string, select?: string, options: any = {}): Promise<RequestArgs> => {
            const localVarPath = `/npc_types_tints`;
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
        updateNpcTypesTint: async (id: number, npcTypesTint: ModelsNpcTypesTint, options: any = {}): Promise<RequestArgs> => {
            if (id === null || id === undefined) {
                throw new RequiredError('id','Required parameter id was null or undefined when calling updateNpcTypesTint.');
            }
            if (npcTypesTint === null || npcTypesTint === undefined) {
                throw new RequiredError('npcTypesTint','Required parameter npcTypesTint was null or undefined when calling updateNpcTypesTint.');
            }
            const localVarPath = `/npc_types_tint/{id}`
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
            const nonString = typeof npcTypesTint !== 'string';
            const needsSerialization = nonString && configuration && configuration.isJsonMime
                ? configuration.isJsonMime(localVarRequestOptions.headers['Content-Type'])
                : nonString;
            localVarRequestOptions.data =  needsSerialization
                ? JSON.stringify(npcTypesTint !== undefined ? npcTypesTint : {})
                : (npcTypesTint || "");
            return {
                url: localVarUrlObj.pathname + localVarUrlObj.search + localVarUrlObj.hash,
                options: localVarRequestOptions,
            };
        },
    }
};
export const NpcTypesTintApiFp = function(configuration?: Configuration) {
    return {
        async createNpcTypesTint(npcTypesTint: ModelsNpcTypesTint, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsNpcTypesTint>>> {
            const localVarAxiosArgs = await NpcTypesTintApiAxiosParamCreator(configuration).createNpcTypesTint(npcTypesTint, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        async deleteNpcTypesTint(id: number, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<string>> {
            const localVarAxiosArgs = await NpcTypesTintApiAxiosParamCreator(configuration).deleteNpcTypesTint(id, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        async getNpcTypesTint(id: number, includes?: string, select?: string, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsNpcTypesTint>>> {
            const localVarAxiosArgs = await NpcTypesTintApiAxiosParamCreator(configuration).getNpcTypesTint(id, includes, select, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        async getNpcTypesTintsBulk(body: CrudcontrollersBulkFetchByIdsGetRequest, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsNpcTypesTint>>> {
            const localVarAxiosArgs = await NpcTypesTintApiAxiosParamCreator(configuration).getNpcTypesTintsBulk(body, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        async listNpcTypesTints(includes?: string, where?: string, whereOr?: string, groupBy?: string, limit?: string, page?: number, orderBy?: string, orderDirection?: string, select?: string, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsNpcTypesTint>>> {
            const localVarAxiosArgs = await NpcTypesTintApiAxiosParamCreator(configuration).listNpcTypesTints(includes, where, whereOr, groupBy, limit, page, orderBy, orderDirection, select, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        async updateNpcTypesTint(id: number, npcTypesTint: ModelsNpcTypesTint, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsNpcTypesTint>>> {
            const localVarAxiosArgs = await NpcTypesTintApiAxiosParamCreator(configuration).updateNpcTypesTint(id, npcTypesTint, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
    }
};
export const NpcTypesTintApiFactory = function (configuration?: Configuration, basePath?: string, axios?: AxiosInstance) {
    return {
        createNpcTypesTint(npcTypesTint: ModelsNpcTypesTint, options?: any): AxiosPromise<Array<ModelsNpcTypesTint>> {
            return NpcTypesTintApiFp(configuration).createNpcTypesTint(npcTypesTint, options).then((request) => request(axios, basePath));
        },
        deleteNpcTypesTint(id: number, options?: any): AxiosPromise<string> {
            return NpcTypesTintApiFp(configuration).deleteNpcTypesTint(id, options).then((request) => request(axios, basePath));
        },
        getNpcTypesTint(id: number, includes?: string, select?: string, options?: any): AxiosPromise<Array<ModelsNpcTypesTint>> {
            return NpcTypesTintApiFp(configuration).getNpcTypesTint(id, includes, select, options).then((request) => request(axios, basePath));
        },
        getNpcTypesTintsBulk(body: CrudcontrollersBulkFetchByIdsGetRequest, options?: any): AxiosPromise<Array<ModelsNpcTypesTint>> {
            return NpcTypesTintApiFp(configuration).getNpcTypesTintsBulk(body, options).then((request) => request(axios, basePath));
        },
        listNpcTypesTints(includes?: string, where?: string, whereOr?: string, groupBy?: string, limit?: string, page?: number, orderBy?: string, orderDirection?: string, select?: string, options?: any): AxiosPromise<Array<ModelsNpcTypesTint>> {
            return NpcTypesTintApiFp(configuration).listNpcTypesTints(includes, where, whereOr, groupBy, limit, page, orderBy, orderDirection, select, options).then((request) => request(axios, basePath));
        },
        updateNpcTypesTint(id: number, npcTypesTint: ModelsNpcTypesTint, options?: any): AxiosPromise<Array<ModelsNpcTypesTint>> {
            return NpcTypesTintApiFp(configuration).updateNpcTypesTint(id, npcTypesTint, options).then((request) => request(axios, basePath));
        },
    };
};
export interface NpcTypesTintApiCreateNpcTypesTintRequest {
    readonly npcTypesTint: ModelsNpcTypesTint
}
export interface NpcTypesTintApiDeleteNpcTypesTintRequest {
    readonly id: number
}
export interface NpcTypesTintApiGetNpcTypesTintRequest {
    readonly id: number
    readonly includes?: string
    readonly select?: string
}
export interface NpcTypesTintApiGetNpcTypesTintsBulkRequest {
    readonly body: CrudcontrollersBulkFetchByIdsGetRequest
}
export interface NpcTypesTintApiListNpcTypesTintsRequest {
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
export interface NpcTypesTintApiUpdateNpcTypesTintRequest {
    readonly id: number
    readonly npcTypesTint: ModelsNpcTypesTint
}
export class NpcTypesTintApi extends BaseAPI {
    public createNpcTypesTint(requestParameters: NpcTypesTintApiCreateNpcTypesTintRequest, options?: any) {
        return NpcTypesTintApiFp(this.configuration).createNpcTypesTint(requestParameters.npcTypesTint, options).then((request) => request(this.axios, this.basePath));
    }
    public deleteNpcTypesTint(requestParameters: NpcTypesTintApiDeleteNpcTypesTintRequest, options?: any) {
        return NpcTypesTintApiFp(this.configuration).deleteNpcTypesTint(requestParameters.id, options).then((request) => request(this.axios, this.basePath));
    }
    public getNpcTypesTint(requestParameters: NpcTypesTintApiGetNpcTypesTintRequest, options?: any) {
        return NpcTypesTintApiFp(this.configuration).getNpcTypesTint(requestParameters.id, requestParameters.includes, requestParameters.select, options).then((request) => request(this.axios, this.basePath));
    }
    public getNpcTypesTintsBulk(requestParameters: NpcTypesTintApiGetNpcTypesTintsBulkRequest, options?: any) {
        return NpcTypesTintApiFp(this.configuration).getNpcTypesTintsBulk(requestParameters.body, options).then((request) => request(this.axios, this.basePath));
    }
    public listNpcTypesTints(requestParameters: NpcTypesTintApiListNpcTypesTintsRequest = {}, options?: any) {
        return NpcTypesTintApiFp(this.configuration).listNpcTypesTints(requestParameters.includes, requestParameters.where, requestParameters.whereOr, requestParameters.groupBy, requestParameters.limit, requestParameters.page, requestParameters.orderBy, requestParameters.orderDirection, requestParameters.select, options).then((request) => request(this.axios, this.basePath));
    }
    public updateNpcTypesTint(requestParameters: NpcTypesTintApiUpdateNpcTypesTintRequest, options?: any) {
        return NpcTypesTintApiFp(this.configuration).updateNpcTypesTint(requestParameters.id, requestParameters.npcTypesTint, options).then((request) => request(this.axios, this.basePath));
    }
}
