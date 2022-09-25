import globalAxios, { AxiosPromise, AxiosInstance } from 'axios';
import { Configuration } from '../configuration';
// Some imports not used depending on template conditions
// @ts-ignore
import { BASE_PATH, COLLECTION_FORMATS, RequestArgs, BaseAPI, RequiredError } from '../base';
// @ts-ignore
import { CrudcontrollersBulkFetchByIdsGetRequest } from '../models';
// @ts-ignore
import { ModelsAdventureTemplateEntry } from '../models';
export const AdventureTemplateEntryApiAxiosParamCreator = function (configuration?: Configuration) {
    return {
        createAdventureTemplateEntry: async (adventureTemplateEntry: ModelsAdventureTemplateEntry, options: any = {}): Promise<RequestArgs> => {
            if (adventureTemplateEntry === null || adventureTemplateEntry === undefined) {
                throw new RequiredError('adventureTemplateEntry','Required parameter adventureTemplateEntry was null or undefined when calling createAdventureTemplateEntry.');
            }
            const localVarPath = `/adventure_template_entry`;
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
            const nonString = typeof adventureTemplateEntry !== 'string';
            const needsSerialization = nonString && configuration && configuration.isJsonMime
                ? configuration.isJsonMime(localVarRequestOptions.headers['Content-Type'])
                : nonString;
            localVarRequestOptions.data =  needsSerialization
                ? JSON.stringify(adventureTemplateEntry !== undefined ? adventureTemplateEntry : {})
                : (adventureTemplateEntry || "");
            return {
                url: localVarUrlObj.pathname + localVarUrlObj.search + localVarUrlObj.hash,
                options: localVarRequestOptions,
            };
        },
        deleteAdventureTemplateEntry: async (id: number, options: any = {}): Promise<RequestArgs> => {
            if (id === null || id === undefined) {
                throw new RequiredError('id','Required parameter id was null or undefined when calling deleteAdventureTemplateEntry.');
            }
            const localVarPath = `/adventure_template_entry/{id}`
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
        getAdventureTemplateEntriesBulk: async (body: CrudcontrollersBulkFetchByIdsGetRequest, options: any = {}): Promise<RequestArgs> => {
            if (body === null || body === undefined) {
                throw new RequiredError('body','Required parameter body was null or undefined when calling getAdventureTemplateEntriesBulk.');
            }
            const localVarPath = `/adventure_template_entries/bulk`;
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
        getAdventureTemplateEntry: async (id: number, includes?: string, select?: string, options: any = {}): Promise<RequestArgs> => {
            if (id === null || id === undefined) {
                throw new RequiredError('id','Required parameter id was null or undefined when calling getAdventureTemplateEntry.');
            }
            const localVarPath = `/adventure_template_entry/{id}`
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
        listAdventureTemplateEntries: async (includes?: string, where?: string, whereOr?: string, groupBy?: string, limit?: string, page?: number, orderBy?: string, orderDirection?: string, select?: string, options: any = {}): Promise<RequestArgs> => {
            const localVarPath = `/adventure_template_entries`;
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
        updateAdventureTemplateEntry: async (id: number, adventureTemplateEntry: ModelsAdventureTemplateEntry, options: any = {}): Promise<RequestArgs> => {
            if (id === null || id === undefined) {
                throw new RequiredError('id','Required parameter id was null or undefined when calling updateAdventureTemplateEntry.');
            }
            if (adventureTemplateEntry === null || adventureTemplateEntry === undefined) {
                throw new RequiredError('adventureTemplateEntry','Required parameter adventureTemplateEntry was null or undefined when calling updateAdventureTemplateEntry.');
            }
            const localVarPath = `/adventure_template_entry/{id}`
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
            const nonString = typeof adventureTemplateEntry !== 'string';
            const needsSerialization = nonString && configuration && configuration.isJsonMime
                ? configuration.isJsonMime(localVarRequestOptions.headers['Content-Type'])
                : nonString;
            localVarRequestOptions.data =  needsSerialization
                ? JSON.stringify(adventureTemplateEntry !== undefined ? adventureTemplateEntry : {})
                : (adventureTemplateEntry || "");
            return {
                url: localVarUrlObj.pathname + localVarUrlObj.search + localVarUrlObj.hash,
                options: localVarRequestOptions,
            };
        },
    }
};
export const AdventureTemplateEntryApiFp = function(configuration?: Configuration) {
    return {
        async createAdventureTemplateEntry(adventureTemplateEntry: ModelsAdventureTemplateEntry, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsAdventureTemplateEntry>>> {
            const localVarAxiosArgs = await AdventureTemplateEntryApiAxiosParamCreator(configuration).createAdventureTemplateEntry(adventureTemplateEntry, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        async deleteAdventureTemplateEntry(id: number, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<string>> {
            const localVarAxiosArgs = await AdventureTemplateEntryApiAxiosParamCreator(configuration).deleteAdventureTemplateEntry(id, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        async getAdventureTemplateEntriesBulk(body: CrudcontrollersBulkFetchByIdsGetRequest, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsAdventureTemplateEntry>>> {
            const localVarAxiosArgs = await AdventureTemplateEntryApiAxiosParamCreator(configuration).getAdventureTemplateEntriesBulk(body, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        async getAdventureTemplateEntry(id: number, includes?: string, select?: string, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsAdventureTemplateEntry>>> {
            const localVarAxiosArgs = await AdventureTemplateEntryApiAxiosParamCreator(configuration).getAdventureTemplateEntry(id, includes, select, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        async listAdventureTemplateEntries(includes?: string, where?: string, whereOr?: string, groupBy?: string, limit?: string, page?: number, orderBy?: string, orderDirection?: string, select?: string, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsAdventureTemplateEntry>>> {
            const localVarAxiosArgs = await AdventureTemplateEntryApiAxiosParamCreator(configuration).listAdventureTemplateEntries(includes, where, whereOr, groupBy, limit, page, orderBy, orderDirection, select, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        async updateAdventureTemplateEntry(id: number, adventureTemplateEntry: ModelsAdventureTemplateEntry, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsAdventureTemplateEntry>>> {
            const localVarAxiosArgs = await AdventureTemplateEntryApiAxiosParamCreator(configuration).updateAdventureTemplateEntry(id, adventureTemplateEntry, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
    }
};
export const AdventureTemplateEntryApiFactory = function (configuration?: Configuration, basePath?: string, axios?: AxiosInstance) {
    return {
        createAdventureTemplateEntry(adventureTemplateEntry: ModelsAdventureTemplateEntry, options?: any): AxiosPromise<Array<ModelsAdventureTemplateEntry>> {
            return AdventureTemplateEntryApiFp(configuration).createAdventureTemplateEntry(adventureTemplateEntry, options).then((request) => request(axios, basePath));
        },
        deleteAdventureTemplateEntry(id: number, options?: any): AxiosPromise<string> {
            return AdventureTemplateEntryApiFp(configuration).deleteAdventureTemplateEntry(id, options).then((request) => request(axios, basePath));
        },
        getAdventureTemplateEntriesBulk(body: CrudcontrollersBulkFetchByIdsGetRequest, options?: any): AxiosPromise<Array<ModelsAdventureTemplateEntry>> {
            return AdventureTemplateEntryApiFp(configuration).getAdventureTemplateEntriesBulk(body, options).then((request) => request(axios, basePath));
        },
        getAdventureTemplateEntry(id: number, includes?: string, select?: string, options?: any): AxiosPromise<Array<ModelsAdventureTemplateEntry>> {
            return AdventureTemplateEntryApiFp(configuration).getAdventureTemplateEntry(id, includes, select, options).then((request) => request(axios, basePath));
        },
        listAdventureTemplateEntries(includes?: string, where?: string, whereOr?: string, groupBy?: string, limit?: string, page?: number, orderBy?: string, orderDirection?: string, select?: string, options?: any): AxiosPromise<Array<ModelsAdventureTemplateEntry>> {
            return AdventureTemplateEntryApiFp(configuration).listAdventureTemplateEntries(includes, where, whereOr, groupBy, limit, page, orderBy, orderDirection, select, options).then((request) => request(axios, basePath));
        },
        updateAdventureTemplateEntry(id: number, adventureTemplateEntry: ModelsAdventureTemplateEntry, options?: any): AxiosPromise<Array<ModelsAdventureTemplateEntry>> {
            return AdventureTemplateEntryApiFp(configuration).updateAdventureTemplateEntry(id, adventureTemplateEntry, options).then((request) => request(axios, basePath));
        },
    };
};
export interface AdventureTemplateEntryApiCreateAdventureTemplateEntryRequest {
    readonly adventureTemplateEntry: ModelsAdventureTemplateEntry
}
export interface AdventureTemplateEntryApiDeleteAdventureTemplateEntryRequest {
    readonly id: number
}
export interface AdventureTemplateEntryApiGetAdventureTemplateEntriesBulkRequest {
    readonly body: CrudcontrollersBulkFetchByIdsGetRequest
}
export interface AdventureTemplateEntryApiGetAdventureTemplateEntryRequest {
    readonly id: number
    readonly includes?: string
    readonly select?: string
}
export interface AdventureTemplateEntryApiListAdventureTemplateEntriesRequest {
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
export interface AdventureTemplateEntryApiUpdateAdventureTemplateEntryRequest {
    readonly id: number
    readonly adventureTemplateEntry: ModelsAdventureTemplateEntry
}
export class AdventureTemplateEntryApi extends BaseAPI {
    public createAdventureTemplateEntry(requestParameters: AdventureTemplateEntryApiCreateAdventureTemplateEntryRequest, options?: any) {
        return AdventureTemplateEntryApiFp(this.configuration).createAdventureTemplateEntry(requestParameters.adventureTemplateEntry, options).then((request) => request(this.axios, this.basePath));
    }
    public deleteAdventureTemplateEntry(requestParameters: AdventureTemplateEntryApiDeleteAdventureTemplateEntryRequest, options?: any) {
        return AdventureTemplateEntryApiFp(this.configuration).deleteAdventureTemplateEntry(requestParameters.id, options).then((request) => request(this.axios, this.basePath));
    }
    public getAdventureTemplateEntriesBulk(requestParameters: AdventureTemplateEntryApiGetAdventureTemplateEntriesBulkRequest, options?: any) {
        return AdventureTemplateEntryApiFp(this.configuration).getAdventureTemplateEntriesBulk(requestParameters.body, options).then((request) => request(this.axios, this.basePath));
    }
    public getAdventureTemplateEntry(requestParameters: AdventureTemplateEntryApiGetAdventureTemplateEntryRequest, options?: any) {
        return AdventureTemplateEntryApiFp(this.configuration).getAdventureTemplateEntry(requestParameters.id, requestParameters.includes, requestParameters.select, options).then((request) => request(this.axios, this.basePath));
    }
    public listAdventureTemplateEntries(requestParameters: AdventureTemplateEntryApiListAdventureTemplateEntriesRequest = {}, options?: any) {
        return AdventureTemplateEntryApiFp(this.configuration).listAdventureTemplateEntries(requestParameters.includes, requestParameters.where, requestParameters.whereOr, requestParameters.groupBy, requestParameters.limit, requestParameters.page, requestParameters.orderBy, requestParameters.orderDirection, requestParameters.select, options).then((request) => request(this.axios, this.basePath));
    }
    public updateAdventureTemplateEntry(requestParameters: AdventureTemplateEntryApiUpdateAdventureTemplateEntryRequest, options?: any) {
        return AdventureTemplateEntryApiFp(this.configuration).updateAdventureTemplateEntry(requestParameters.id, requestParameters.adventureTemplateEntry, options).then((request) => request(this.axios, this.basePath));
    }
}
