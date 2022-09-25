import globalAxios, { AxiosPromise, AxiosInstance } from 'axios';
import { Configuration } from '../configuration';
// Some imports not used depending on template conditions
// @ts-ignore
import { BASE_PATH, COLLECTION_FORMATS, RequestArgs, BaseAPI, RequiredError } from '../base';
// @ts-ignore
import { CrudcontrollersBulkFetchByIdsGetRequest } from '../models';
// @ts-ignore
import { ModelsNpcFactionEntry } from '../models';
export const NpcFactionEntryApiAxiosParamCreator = function (configuration?: Configuration) {
    return {
        createNpcFactionEntry: async (npcFactionEntry: ModelsNpcFactionEntry, options: any = {}): Promise<RequestArgs> => {
            if (npcFactionEntry === null || npcFactionEntry === undefined) {
                throw new RequiredError('npcFactionEntry','Required parameter npcFactionEntry was null or undefined when calling createNpcFactionEntry.');
            }
            const localVarPath = `/npc_faction_entry`;
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
            const nonString = typeof npcFactionEntry !== 'string';
            const needsSerialization = nonString && configuration && configuration.isJsonMime
                ? configuration.isJsonMime(localVarRequestOptions.headers['Content-Type'])
                : nonString;
            localVarRequestOptions.data =  needsSerialization
                ? JSON.stringify(npcFactionEntry !== undefined ? npcFactionEntry : {})
                : (npcFactionEntry || "");
            return {
                url: localVarUrlObj.pathname + localVarUrlObj.search + localVarUrlObj.hash,
                options: localVarRequestOptions,
            };
        },
        deleteNpcFactionEntry: async (id: number, options: any = {}): Promise<RequestArgs> => {
            if (id === null || id === undefined) {
                throw new RequiredError('id','Required parameter id was null or undefined when calling deleteNpcFactionEntry.');
            }
            const localVarPath = `/npc_faction_entry/{id}`
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
        getNpcFactionEntriesBulk: async (body: CrudcontrollersBulkFetchByIdsGetRequest, options: any = {}): Promise<RequestArgs> => {
            if (body === null || body === undefined) {
                throw new RequiredError('body','Required parameter body was null or undefined when calling getNpcFactionEntriesBulk.');
            }
            const localVarPath = `/npc_faction_entries/bulk`;
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
        getNpcFactionEntry: async (id: number, includes?: string, select?: string, options: any = {}): Promise<RequestArgs> => {
            if (id === null || id === undefined) {
                throw new RequiredError('id','Required parameter id was null or undefined when calling getNpcFactionEntry.');
            }
            const localVarPath = `/npc_faction_entry/{id}`
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
        listNpcFactionEntries: async (includes?: string, where?: string, whereOr?: string, groupBy?: string, limit?: string, page?: number, orderBy?: string, orderDirection?: string, select?: string, options: any = {}): Promise<RequestArgs> => {
            const localVarPath = `/npc_faction_entries`;
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
        updateNpcFactionEntry: async (id: number, npcFactionEntry: ModelsNpcFactionEntry, options: any = {}): Promise<RequestArgs> => {
            if (id === null || id === undefined) {
                throw new RequiredError('id','Required parameter id was null or undefined when calling updateNpcFactionEntry.');
            }
            if (npcFactionEntry === null || npcFactionEntry === undefined) {
                throw new RequiredError('npcFactionEntry','Required parameter npcFactionEntry was null or undefined when calling updateNpcFactionEntry.');
            }
            const localVarPath = `/npc_faction_entry/{id}`
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
            const nonString = typeof npcFactionEntry !== 'string';
            const needsSerialization = nonString && configuration && configuration.isJsonMime
                ? configuration.isJsonMime(localVarRequestOptions.headers['Content-Type'])
                : nonString;
            localVarRequestOptions.data =  needsSerialization
                ? JSON.stringify(npcFactionEntry !== undefined ? npcFactionEntry : {})
                : (npcFactionEntry || "");
            return {
                url: localVarUrlObj.pathname + localVarUrlObj.search + localVarUrlObj.hash,
                options: localVarRequestOptions,
            };
        },
    }
};
export const NpcFactionEntryApiFp = function(configuration?: Configuration) {
    return {
        async createNpcFactionEntry(npcFactionEntry: ModelsNpcFactionEntry, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsNpcFactionEntry>>> {
            const localVarAxiosArgs = await NpcFactionEntryApiAxiosParamCreator(configuration).createNpcFactionEntry(npcFactionEntry, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        async deleteNpcFactionEntry(id: number, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<string>> {
            const localVarAxiosArgs = await NpcFactionEntryApiAxiosParamCreator(configuration).deleteNpcFactionEntry(id, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        async getNpcFactionEntriesBulk(body: CrudcontrollersBulkFetchByIdsGetRequest, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsNpcFactionEntry>>> {
            const localVarAxiosArgs = await NpcFactionEntryApiAxiosParamCreator(configuration).getNpcFactionEntriesBulk(body, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        async getNpcFactionEntry(id: number, includes?: string, select?: string, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsNpcFactionEntry>>> {
            const localVarAxiosArgs = await NpcFactionEntryApiAxiosParamCreator(configuration).getNpcFactionEntry(id, includes, select, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        async listNpcFactionEntries(includes?: string, where?: string, whereOr?: string, groupBy?: string, limit?: string, page?: number, orderBy?: string, orderDirection?: string, select?: string, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsNpcFactionEntry>>> {
            const localVarAxiosArgs = await NpcFactionEntryApiAxiosParamCreator(configuration).listNpcFactionEntries(includes, where, whereOr, groupBy, limit, page, orderBy, orderDirection, select, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        async updateNpcFactionEntry(id: number, npcFactionEntry: ModelsNpcFactionEntry, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsNpcFactionEntry>>> {
            const localVarAxiosArgs = await NpcFactionEntryApiAxiosParamCreator(configuration).updateNpcFactionEntry(id, npcFactionEntry, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
    }
};
export const NpcFactionEntryApiFactory = function (configuration?: Configuration, basePath?: string, axios?: AxiosInstance) {
    return {
        createNpcFactionEntry(npcFactionEntry: ModelsNpcFactionEntry, options?: any): AxiosPromise<Array<ModelsNpcFactionEntry>> {
            return NpcFactionEntryApiFp(configuration).createNpcFactionEntry(npcFactionEntry, options).then((request) => request(axios, basePath));
        },
        deleteNpcFactionEntry(id: number, options?: any): AxiosPromise<string> {
            return NpcFactionEntryApiFp(configuration).deleteNpcFactionEntry(id, options).then((request) => request(axios, basePath));
        },
        getNpcFactionEntriesBulk(body: CrudcontrollersBulkFetchByIdsGetRequest, options?: any): AxiosPromise<Array<ModelsNpcFactionEntry>> {
            return NpcFactionEntryApiFp(configuration).getNpcFactionEntriesBulk(body, options).then((request) => request(axios, basePath));
        },
        getNpcFactionEntry(id: number, includes?: string, select?: string, options?: any): AxiosPromise<Array<ModelsNpcFactionEntry>> {
            return NpcFactionEntryApiFp(configuration).getNpcFactionEntry(id, includes, select, options).then((request) => request(axios, basePath));
        },
        listNpcFactionEntries(includes?: string, where?: string, whereOr?: string, groupBy?: string, limit?: string, page?: number, orderBy?: string, orderDirection?: string, select?: string, options?: any): AxiosPromise<Array<ModelsNpcFactionEntry>> {
            return NpcFactionEntryApiFp(configuration).listNpcFactionEntries(includes, where, whereOr, groupBy, limit, page, orderBy, orderDirection, select, options).then((request) => request(axios, basePath));
        },
        updateNpcFactionEntry(id: number, npcFactionEntry: ModelsNpcFactionEntry, options?: any): AxiosPromise<Array<ModelsNpcFactionEntry>> {
            return NpcFactionEntryApiFp(configuration).updateNpcFactionEntry(id, npcFactionEntry, options).then((request) => request(axios, basePath));
        },
    };
};
export interface NpcFactionEntryApiCreateNpcFactionEntryRequest {
    readonly npcFactionEntry: ModelsNpcFactionEntry
}
export interface NpcFactionEntryApiDeleteNpcFactionEntryRequest {
    readonly id: number
}
export interface NpcFactionEntryApiGetNpcFactionEntriesBulkRequest {
    readonly body: CrudcontrollersBulkFetchByIdsGetRequest
}
export interface NpcFactionEntryApiGetNpcFactionEntryRequest {
    readonly id: number
    readonly includes?: string
    readonly select?: string
}
export interface NpcFactionEntryApiListNpcFactionEntriesRequest {
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
export interface NpcFactionEntryApiUpdateNpcFactionEntryRequest {
    readonly id: number
    readonly npcFactionEntry: ModelsNpcFactionEntry
}
export class NpcFactionEntryApi extends BaseAPI {
    public createNpcFactionEntry(requestParameters: NpcFactionEntryApiCreateNpcFactionEntryRequest, options?: any) {
        return NpcFactionEntryApiFp(this.configuration).createNpcFactionEntry(requestParameters.npcFactionEntry, options).then((request) => request(this.axios, this.basePath));
    }
    public deleteNpcFactionEntry(requestParameters: NpcFactionEntryApiDeleteNpcFactionEntryRequest, options?: any) {
        return NpcFactionEntryApiFp(this.configuration).deleteNpcFactionEntry(requestParameters.id, options).then((request) => request(this.axios, this.basePath));
    }
    public getNpcFactionEntriesBulk(requestParameters: NpcFactionEntryApiGetNpcFactionEntriesBulkRequest, options?: any) {
        return NpcFactionEntryApiFp(this.configuration).getNpcFactionEntriesBulk(requestParameters.body, options).then((request) => request(this.axios, this.basePath));
    }
    public getNpcFactionEntry(requestParameters: NpcFactionEntryApiGetNpcFactionEntryRequest, options?: any) {
        return NpcFactionEntryApiFp(this.configuration).getNpcFactionEntry(requestParameters.id, requestParameters.includes, requestParameters.select, options).then((request) => request(this.axios, this.basePath));
    }
    public listNpcFactionEntries(requestParameters: NpcFactionEntryApiListNpcFactionEntriesRequest = {}, options?: any) {
        return NpcFactionEntryApiFp(this.configuration).listNpcFactionEntries(requestParameters.includes, requestParameters.where, requestParameters.whereOr, requestParameters.groupBy, requestParameters.limit, requestParameters.page, requestParameters.orderBy, requestParameters.orderDirection, requestParameters.select, options).then((request) => request(this.axios, this.basePath));
    }
    public updateNpcFactionEntry(requestParameters: NpcFactionEntryApiUpdateNpcFactionEntryRequest, options?: any) {
        return NpcFactionEntryApiFp(this.configuration).updateNpcFactionEntry(requestParameters.id, requestParameters.npcFactionEntry, options).then((request) => request(this.axios, this.basePath));
    }
}
