import globalAxios, { AxiosPromise, AxiosInstance } from 'axios';
import { Configuration } from '../configuration';
// Some imports not used depending on template conditions
// @ts-ignore
import { BASE_PATH, COLLECTION_FORMATS, RequestArgs, BaseAPI, RequiredError } from '../base';
// @ts-ignore
import { CrudcontrollersBulkFetchByIdsGetRequest } from '../models';
// @ts-ignore
import { ModelsNpcSpellsEffectsEntry } from '../models';
export const NpcSpellsEffectsEntryApiAxiosParamCreator = function (configuration?: Configuration) {
    return {
        createNpcSpellsEffectsEntry: async (npcSpellsEffectsEntry: ModelsNpcSpellsEffectsEntry, options: any = {}): Promise<RequestArgs> => {
            if (npcSpellsEffectsEntry === null || npcSpellsEffectsEntry === undefined) {
                throw new RequiredError('npcSpellsEffectsEntry','Required parameter npcSpellsEffectsEntry was null or undefined when calling createNpcSpellsEffectsEntry.');
            }
            const localVarPath = `/npc_spells_effects_entry`;
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
            const nonString = typeof npcSpellsEffectsEntry !== 'string';
            const needsSerialization = nonString && configuration && configuration.isJsonMime
                ? configuration.isJsonMime(localVarRequestOptions.headers['Content-Type'])
                : nonString;
            localVarRequestOptions.data =  needsSerialization
                ? JSON.stringify(npcSpellsEffectsEntry !== undefined ? npcSpellsEffectsEntry : {})
                : (npcSpellsEffectsEntry || "");
            return {
                url: localVarUrlObj.pathname + localVarUrlObj.search + localVarUrlObj.hash,
                options: localVarRequestOptions,
            };
        },
        deleteNpcSpellsEffectsEntry: async (id: number, options: any = {}): Promise<RequestArgs> => {
            if (id === null || id === undefined) {
                throw new RequiredError('id','Required parameter id was null or undefined when calling deleteNpcSpellsEffectsEntry.');
            }
            const localVarPath = `/npc_spells_effects_entry/{id}`
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
        getNpcSpellsEffectsEntriesBulk: async (body: CrudcontrollersBulkFetchByIdsGetRequest, options: any = {}): Promise<RequestArgs> => {
            if (body === null || body === undefined) {
                throw new RequiredError('body','Required parameter body was null or undefined when calling getNpcSpellsEffectsEntriesBulk.');
            }
            const localVarPath = `/npc_spells_effects_entries/bulk`;
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
        getNpcSpellsEffectsEntry: async (id: number, includes?: string, select?: string, options: any = {}): Promise<RequestArgs> => {
            if (id === null || id === undefined) {
                throw new RequiredError('id','Required parameter id was null or undefined when calling getNpcSpellsEffectsEntry.');
            }
            const localVarPath = `/npc_spells_effects_entry/{id}`
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
        listNpcSpellsEffectsEntries: async (includes?: string, where?: string, whereOr?: string, groupBy?: string, limit?: string, page?: number, orderBy?: string, orderDirection?: string, select?: string, options: any = {}): Promise<RequestArgs> => {
            const localVarPath = `/npc_spells_effects_entries`;
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
        updateNpcSpellsEffectsEntry: async (id: number, npcSpellsEffectsEntry: ModelsNpcSpellsEffectsEntry, options: any = {}): Promise<RequestArgs> => {
            if (id === null || id === undefined) {
                throw new RequiredError('id','Required parameter id was null or undefined when calling updateNpcSpellsEffectsEntry.');
            }
            if (npcSpellsEffectsEntry === null || npcSpellsEffectsEntry === undefined) {
                throw new RequiredError('npcSpellsEffectsEntry','Required parameter npcSpellsEffectsEntry was null or undefined when calling updateNpcSpellsEffectsEntry.');
            }
            const localVarPath = `/npc_spells_effects_entry/{id}`
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
            const nonString = typeof npcSpellsEffectsEntry !== 'string';
            const needsSerialization = nonString && configuration && configuration.isJsonMime
                ? configuration.isJsonMime(localVarRequestOptions.headers['Content-Type'])
                : nonString;
            localVarRequestOptions.data =  needsSerialization
                ? JSON.stringify(npcSpellsEffectsEntry !== undefined ? npcSpellsEffectsEntry : {})
                : (npcSpellsEffectsEntry || "");
            return {
                url: localVarUrlObj.pathname + localVarUrlObj.search + localVarUrlObj.hash,
                options: localVarRequestOptions,
            };
        },
    }
};
export const NpcSpellsEffectsEntryApiFp = function(configuration?: Configuration) {
    return {
        async createNpcSpellsEffectsEntry(npcSpellsEffectsEntry: ModelsNpcSpellsEffectsEntry, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsNpcSpellsEffectsEntry>>> {
            const localVarAxiosArgs = await NpcSpellsEffectsEntryApiAxiosParamCreator(configuration).createNpcSpellsEffectsEntry(npcSpellsEffectsEntry, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        async deleteNpcSpellsEffectsEntry(id: number, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<string>> {
            const localVarAxiosArgs = await NpcSpellsEffectsEntryApiAxiosParamCreator(configuration).deleteNpcSpellsEffectsEntry(id, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        async getNpcSpellsEffectsEntriesBulk(body: CrudcontrollersBulkFetchByIdsGetRequest, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsNpcSpellsEffectsEntry>>> {
            const localVarAxiosArgs = await NpcSpellsEffectsEntryApiAxiosParamCreator(configuration).getNpcSpellsEffectsEntriesBulk(body, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        async getNpcSpellsEffectsEntry(id: number, includes?: string, select?: string, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsNpcSpellsEffectsEntry>>> {
            const localVarAxiosArgs = await NpcSpellsEffectsEntryApiAxiosParamCreator(configuration).getNpcSpellsEffectsEntry(id, includes, select, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        async listNpcSpellsEffectsEntries(includes?: string, where?: string, whereOr?: string, groupBy?: string, limit?: string, page?: number, orderBy?: string, orderDirection?: string, select?: string, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsNpcSpellsEffectsEntry>>> {
            const localVarAxiosArgs = await NpcSpellsEffectsEntryApiAxiosParamCreator(configuration).listNpcSpellsEffectsEntries(includes, where, whereOr, groupBy, limit, page, orderBy, orderDirection, select, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        async updateNpcSpellsEffectsEntry(id: number, npcSpellsEffectsEntry: ModelsNpcSpellsEffectsEntry, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsNpcSpellsEffectsEntry>>> {
            const localVarAxiosArgs = await NpcSpellsEffectsEntryApiAxiosParamCreator(configuration).updateNpcSpellsEffectsEntry(id, npcSpellsEffectsEntry, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
    }
};
export const NpcSpellsEffectsEntryApiFactory = function (configuration?: Configuration, basePath?: string, axios?: AxiosInstance) {
    return {
        createNpcSpellsEffectsEntry(npcSpellsEffectsEntry: ModelsNpcSpellsEffectsEntry, options?: any): AxiosPromise<Array<ModelsNpcSpellsEffectsEntry>> {
            return NpcSpellsEffectsEntryApiFp(configuration).createNpcSpellsEffectsEntry(npcSpellsEffectsEntry, options).then((request) => request(axios, basePath));
        },
        deleteNpcSpellsEffectsEntry(id: number, options?: any): AxiosPromise<string> {
            return NpcSpellsEffectsEntryApiFp(configuration).deleteNpcSpellsEffectsEntry(id, options).then((request) => request(axios, basePath));
        },
        getNpcSpellsEffectsEntriesBulk(body: CrudcontrollersBulkFetchByIdsGetRequest, options?: any): AxiosPromise<Array<ModelsNpcSpellsEffectsEntry>> {
            return NpcSpellsEffectsEntryApiFp(configuration).getNpcSpellsEffectsEntriesBulk(body, options).then((request) => request(axios, basePath));
        },
        getNpcSpellsEffectsEntry(id: number, includes?: string, select?: string, options?: any): AxiosPromise<Array<ModelsNpcSpellsEffectsEntry>> {
            return NpcSpellsEffectsEntryApiFp(configuration).getNpcSpellsEffectsEntry(id, includes, select, options).then((request) => request(axios, basePath));
        },
        listNpcSpellsEffectsEntries(includes?: string, where?: string, whereOr?: string, groupBy?: string, limit?: string, page?: number, orderBy?: string, orderDirection?: string, select?: string, options?: any): AxiosPromise<Array<ModelsNpcSpellsEffectsEntry>> {
            return NpcSpellsEffectsEntryApiFp(configuration).listNpcSpellsEffectsEntries(includes, where, whereOr, groupBy, limit, page, orderBy, orderDirection, select, options).then((request) => request(axios, basePath));
        },
        updateNpcSpellsEffectsEntry(id: number, npcSpellsEffectsEntry: ModelsNpcSpellsEffectsEntry, options?: any): AxiosPromise<Array<ModelsNpcSpellsEffectsEntry>> {
            return NpcSpellsEffectsEntryApiFp(configuration).updateNpcSpellsEffectsEntry(id, npcSpellsEffectsEntry, options).then((request) => request(axios, basePath));
        },
    };
};
export interface NpcSpellsEffectsEntryApiCreateNpcSpellsEffectsEntryRequest {
    readonly npcSpellsEffectsEntry: ModelsNpcSpellsEffectsEntry
}
export interface NpcSpellsEffectsEntryApiDeleteNpcSpellsEffectsEntryRequest {
    readonly id: number
}
export interface NpcSpellsEffectsEntryApiGetNpcSpellsEffectsEntriesBulkRequest {
    readonly body: CrudcontrollersBulkFetchByIdsGetRequest
}
export interface NpcSpellsEffectsEntryApiGetNpcSpellsEffectsEntryRequest {
    readonly id: number
    readonly includes?: string
    readonly select?: string
}
export interface NpcSpellsEffectsEntryApiListNpcSpellsEffectsEntriesRequest {
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
export interface NpcSpellsEffectsEntryApiUpdateNpcSpellsEffectsEntryRequest {
    readonly id: number
    readonly npcSpellsEffectsEntry: ModelsNpcSpellsEffectsEntry
}
export class NpcSpellsEffectsEntryApi extends BaseAPI {
    public createNpcSpellsEffectsEntry(requestParameters: NpcSpellsEffectsEntryApiCreateNpcSpellsEffectsEntryRequest, options?: any) {
        return NpcSpellsEffectsEntryApiFp(this.configuration).createNpcSpellsEffectsEntry(requestParameters.npcSpellsEffectsEntry, options).then((request) => request(this.axios, this.basePath));
    }
    public deleteNpcSpellsEffectsEntry(requestParameters: NpcSpellsEffectsEntryApiDeleteNpcSpellsEffectsEntryRequest, options?: any) {
        return NpcSpellsEffectsEntryApiFp(this.configuration).deleteNpcSpellsEffectsEntry(requestParameters.id, options).then((request) => request(this.axios, this.basePath));
    }
    public getNpcSpellsEffectsEntriesBulk(requestParameters: NpcSpellsEffectsEntryApiGetNpcSpellsEffectsEntriesBulkRequest, options?: any) {
        return NpcSpellsEffectsEntryApiFp(this.configuration).getNpcSpellsEffectsEntriesBulk(requestParameters.body, options).then((request) => request(this.axios, this.basePath));
    }
    public getNpcSpellsEffectsEntry(requestParameters: NpcSpellsEffectsEntryApiGetNpcSpellsEffectsEntryRequest, options?: any) {
        return NpcSpellsEffectsEntryApiFp(this.configuration).getNpcSpellsEffectsEntry(requestParameters.id, requestParameters.includes, requestParameters.select, options).then((request) => request(this.axios, this.basePath));
    }
    public listNpcSpellsEffectsEntries(requestParameters: NpcSpellsEffectsEntryApiListNpcSpellsEffectsEntriesRequest = {}, options?: any) {
        return NpcSpellsEffectsEntryApiFp(this.configuration).listNpcSpellsEffectsEntries(requestParameters.includes, requestParameters.where, requestParameters.whereOr, requestParameters.groupBy, requestParameters.limit, requestParameters.page, requestParameters.orderBy, requestParameters.orderDirection, requestParameters.select, options).then((request) => request(this.axios, this.basePath));
    }
    public updateNpcSpellsEffectsEntry(requestParameters: NpcSpellsEffectsEntryApiUpdateNpcSpellsEffectsEntryRequest, options?: any) {
        return NpcSpellsEffectsEntryApiFp(this.configuration).updateNpcSpellsEffectsEntry(requestParameters.id, requestParameters.npcSpellsEffectsEntry, options).then((request) => request(this.axios, this.basePath));
    }
}
