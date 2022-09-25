import globalAxios, { AxiosPromise, AxiosInstance } from 'axios';
import { Configuration } from '../configuration';
// Some imports not used depending on template conditions
// @ts-ignore
import { BASE_PATH, COLLECTION_FORMATS, RequestArgs, BaseAPI, RequiredError } from '../base';
// @ts-ignore
import { CrudcontrollersBulkFetchByIdsGetRequest } from '../models';
// @ts-ignore
import { ModelsTradeskillRecipeEntry } from '../models';
export const TradeskillRecipeEntryApiAxiosParamCreator = function (configuration?: Configuration) {
    return {
        createTradeskillRecipeEntry: async (tradeskillRecipeEntry: ModelsTradeskillRecipeEntry, options: any = {}): Promise<RequestArgs> => {
            if (tradeskillRecipeEntry === null || tradeskillRecipeEntry === undefined) {
                throw new RequiredError('tradeskillRecipeEntry','Required parameter tradeskillRecipeEntry was null or undefined when calling createTradeskillRecipeEntry.');
            }
            const localVarPath = `/tradeskill_recipe_entry`;
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
            const nonString = typeof tradeskillRecipeEntry !== 'string';
            const needsSerialization = nonString && configuration && configuration.isJsonMime
                ? configuration.isJsonMime(localVarRequestOptions.headers['Content-Type'])
                : nonString;
            localVarRequestOptions.data =  needsSerialization
                ? JSON.stringify(tradeskillRecipeEntry !== undefined ? tradeskillRecipeEntry : {})
                : (tradeskillRecipeEntry || "");
            return {
                url: localVarUrlObj.pathname + localVarUrlObj.search + localVarUrlObj.hash,
                options: localVarRequestOptions,
            };
        },
        deleteTradeskillRecipeEntry: async (id: number, options: any = {}): Promise<RequestArgs> => {
            if (id === null || id === undefined) {
                throw new RequiredError('id','Required parameter id was null or undefined when calling deleteTradeskillRecipeEntry.');
            }
            const localVarPath = `/tradeskill_recipe_entry/{id}`
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
        getTradeskillRecipeEntriesBulk: async (body: CrudcontrollersBulkFetchByIdsGetRequest, options: any = {}): Promise<RequestArgs> => {
            if (body === null || body === undefined) {
                throw new RequiredError('body','Required parameter body was null or undefined when calling getTradeskillRecipeEntriesBulk.');
            }
            const localVarPath = `/tradeskill_recipe_entries/bulk`;
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
        getTradeskillRecipeEntry: async (id: number, includes?: string, select?: string, options: any = {}): Promise<RequestArgs> => {
            if (id === null || id === undefined) {
                throw new RequiredError('id','Required parameter id was null or undefined when calling getTradeskillRecipeEntry.');
            }
            const localVarPath = `/tradeskill_recipe_entry/{id}`
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
        listTradeskillRecipeEntries: async (includes?: string, where?: string, whereOr?: string, groupBy?: string, limit?: string, page?: number, orderBy?: string, orderDirection?: string, select?: string, options: any = {}): Promise<RequestArgs> => {
            const localVarPath = `/tradeskill_recipe_entries`;
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
        updateTradeskillRecipeEntry: async (id: number, tradeskillRecipeEntry: ModelsTradeskillRecipeEntry, options: any = {}): Promise<RequestArgs> => {
            if (id === null || id === undefined) {
                throw new RequiredError('id','Required parameter id was null or undefined when calling updateTradeskillRecipeEntry.');
            }
            if (tradeskillRecipeEntry === null || tradeskillRecipeEntry === undefined) {
                throw new RequiredError('tradeskillRecipeEntry','Required parameter tradeskillRecipeEntry was null or undefined when calling updateTradeskillRecipeEntry.');
            }
            const localVarPath = `/tradeskill_recipe_entry/{id}`
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
            const nonString = typeof tradeskillRecipeEntry !== 'string';
            const needsSerialization = nonString && configuration && configuration.isJsonMime
                ? configuration.isJsonMime(localVarRequestOptions.headers['Content-Type'])
                : nonString;
            localVarRequestOptions.data =  needsSerialization
                ? JSON.stringify(tradeskillRecipeEntry !== undefined ? tradeskillRecipeEntry : {})
                : (tradeskillRecipeEntry || "");
            return {
                url: localVarUrlObj.pathname + localVarUrlObj.search + localVarUrlObj.hash,
                options: localVarRequestOptions,
            };
        },
    }
};
export const TradeskillRecipeEntryApiFp = function(configuration?: Configuration) {
    return {
        async createTradeskillRecipeEntry(tradeskillRecipeEntry: ModelsTradeskillRecipeEntry, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsTradeskillRecipeEntry>>> {
            const localVarAxiosArgs = await TradeskillRecipeEntryApiAxiosParamCreator(configuration).createTradeskillRecipeEntry(tradeskillRecipeEntry, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        async deleteTradeskillRecipeEntry(id: number, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<string>> {
            const localVarAxiosArgs = await TradeskillRecipeEntryApiAxiosParamCreator(configuration).deleteTradeskillRecipeEntry(id, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        async getTradeskillRecipeEntriesBulk(body: CrudcontrollersBulkFetchByIdsGetRequest, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsTradeskillRecipeEntry>>> {
            const localVarAxiosArgs = await TradeskillRecipeEntryApiAxiosParamCreator(configuration).getTradeskillRecipeEntriesBulk(body, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        async getTradeskillRecipeEntry(id: number, includes?: string, select?: string, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsTradeskillRecipeEntry>>> {
            const localVarAxiosArgs = await TradeskillRecipeEntryApiAxiosParamCreator(configuration).getTradeskillRecipeEntry(id, includes, select, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        async listTradeskillRecipeEntries(includes?: string, where?: string, whereOr?: string, groupBy?: string, limit?: string, page?: number, orderBy?: string, orderDirection?: string, select?: string, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsTradeskillRecipeEntry>>> {
            const localVarAxiosArgs = await TradeskillRecipeEntryApiAxiosParamCreator(configuration).listTradeskillRecipeEntries(includes, where, whereOr, groupBy, limit, page, orderBy, orderDirection, select, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        async updateTradeskillRecipeEntry(id: number, tradeskillRecipeEntry: ModelsTradeskillRecipeEntry, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsTradeskillRecipeEntry>>> {
            const localVarAxiosArgs = await TradeskillRecipeEntryApiAxiosParamCreator(configuration).updateTradeskillRecipeEntry(id, tradeskillRecipeEntry, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
    }
};
export const TradeskillRecipeEntryApiFactory = function (configuration?: Configuration, basePath?: string, axios?: AxiosInstance) {
    return {
        createTradeskillRecipeEntry(tradeskillRecipeEntry: ModelsTradeskillRecipeEntry, options?: any): AxiosPromise<Array<ModelsTradeskillRecipeEntry>> {
            return TradeskillRecipeEntryApiFp(configuration).createTradeskillRecipeEntry(tradeskillRecipeEntry, options).then((request) => request(axios, basePath));
        },
        deleteTradeskillRecipeEntry(id: number, options?: any): AxiosPromise<string> {
            return TradeskillRecipeEntryApiFp(configuration).deleteTradeskillRecipeEntry(id, options).then((request) => request(axios, basePath));
        },
        getTradeskillRecipeEntriesBulk(body: CrudcontrollersBulkFetchByIdsGetRequest, options?: any): AxiosPromise<Array<ModelsTradeskillRecipeEntry>> {
            return TradeskillRecipeEntryApiFp(configuration).getTradeskillRecipeEntriesBulk(body, options).then((request) => request(axios, basePath));
        },
        getTradeskillRecipeEntry(id: number, includes?: string, select?: string, options?: any): AxiosPromise<Array<ModelsTradeskillRecipeEntry>> {
            return TradeskillRecipeEntryApiFp(configuration).getTradeskillRecipeEntry(id, includes, select, options).then((request) => request(axios, basePath));
        },
        listTradeskillRecipeEntries(includes?: string, where?: string, whereOr?: string, groupBy?: string, limit?: string, page?: number, orderBy?: string, orderDirection?: string, select?: string, options?: any): AxiosPromise<Array<ModelsTradeskillRecipeEntry>> {
            return TradeskillRecipeEntryApiFp(configuration).listTradeskillRecipeEntries(includes, where, whereOr, groupBy, limit, page, orderBy, orderDirection, select, options).then((request) => request(axios, basePath));
        },
        updateTradeskillRecipeEntry(id: number, tradeskillRecipeEntry: ModelsTradeskillRecipeEntry, options?: any): AxiosPromise<Array<ModelsTradeskillRecipeEntry>> {
            return TradeskillRecipeEntryApiFp(configuration).updateTradeskillRecipeEntry(id, tradeskillRecipeEntry, options).then((request) => request(axios, basePath));
        },
    };
};
export interface TradeskillRecipeEntryApiCreateTradeskillRecipeEntryRequest {
    readonly tradeskillRecipeEntry: ModelsTradeskillRecipeEntry
}
export interface TradeskillRecipeEntryApiDeleteTradeskillRecipeEntryRequest {
    readonly id: number
}
export interface TradeskillRecipeEntryApiGetTradeskillRecipeEntriesBulkRequest {
    readonly body: CrudcontrollersBulkFetchByIdsGetRequest
}
export interface TradeskillRecipeEntryApiGetTradeskillRecipeEntryRequest {
    readonly id: number
    readonly includes?: string
    readonly select?: string
}
export interface TradeskillRecipeEntryApiListTradeskillRecipeEntriesRequest {
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
export interface TradeskillRecipeEntryApiUpdateTradeskillRecipeEntryRequest {
    readonly id: number
    readonly tradeskillRecipeEntry: ModelsTradeskillRecipeEntry
}
export class TradeskillRecipeEntryApi extends BaseAPI {
    public createTradeskillRecipeEntry(requestParameters: TradeskillRecipeEntryApiCreateTradeskillRecipeEntryRequest, options?: any) {
        return TradeskillRecipeEntryApiFp(this.configuration).createTradeskillRecipeEntry(requestParameters.tradeskillRecipeEntry, options).then((request) => request(this.axios, this.basePath));
    }
    public deleteTradeskillRecipeEntry(requestParameters: TradeskillRecipeEntryApiDeleteTradeskillRecipeEntryRequest, options?: any) {
        return TradeskillRecipeEntryApiFp(this.configuration).deleteTradeskillRecipeEntry(requestParameters.id, options).then((request) => request(this.axios, this.basePath));
    }
    public getTradeskillRecipeEntriesBulk(requestParameters: TradeskillRecipeEntryApiGetTradeskillRecipeEntriesBulkRequest, options?: any) {
        return TradeskillRecipeEntryApiFp(this.configuration).getTradeskillRecipeEntriesBulk(requestParameters.body, options).then((request) => request(this.axios, this.basePath));
    }
    public getTradeskillRecipeEntry(requestParameters: TradeskillRecipeEntryApiGetTradeskillRecipeEntryRequest, options?: any) {
        return TradeskillRecipeEntryApiFp(this.configuration).getTradeskillRecipeEntry(requestParameters.id, requestParameters.includes, requestParameters.select, options).then((request) => request(this.axios, this.basePath));
    }
    public listTradeskillRecipeEntries(requestParameters: TradeskillRecipeEntryApiListTradeskillRecipeEntriesRequest = {}, options?: any) {
        return TradeskillRecipeEntryApiFp(this.configuration).listTradeskillRecipeEntries(requestParameters.includes, requestParameters.where, requestParameters.whereOr, requestParameters.groupBy, requestParameters.limit, requestParameters.page, requestParameters.orderBy, requestParameters.orderDirection, requestParameters.select, options).then((request) => request(this.axios, this.basePath));
    }
    public updateTradeskillRecipeEntry(requestParameters: TradeskillRecipeEntryApiUpdateTradeskillRecipeEntryRequest, options?: any) {
        return TradeskillRecipeEntryApiFp(this.configuration).updateTradeskillRecipeEntry(requestParameters.id, requestParameters.tradeskillRecipeEntry, options).then((request) => request(this.axios, this.basePath));
    }
}
