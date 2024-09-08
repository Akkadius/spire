import globalAxios, { AxiosPromise, AxiosInstance } from 'axios';
import { Configuration } from '../configuration';
// Some imports not used depending on template conditions
// @ts-ignore
import { BASE_PATH, COLLECTION_FORMATS, RequestArgs, BaseAPI, RequiredError } from '../base';
// @ts-ignore
import { CrudcontrollersBulkFetchByIdsGetRequest } from '../models';
// @ts-ignore
import { ModelsBotStartingItem } from '../models';
export const BotStartingItemApiAxiosParamCreator = function (configuration?: Configuration) {
    return {
        createBotStartingItem: async (botStartingItem: ModelsBotStartingItem, options: any = {}): Promise<RequestArgs> => {
            if (botStartingItem === null || botStartingItem === undefined) {
                throw new RequiredError('botStartingItem','Required parameter botStartingItem was null or undefined when calling createBotStartingItem.');
            }
            const localVarPath = `/bot_starting_item`;
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
            const nonString = typeof botStartingItem !== 'string';
            const needsSerialization = nonString && configuration && configuration.isJsonMime
                ? configuration.isJsonMime(localVarRequestOptions.headers['Content-Type'])
                : nonString;
            localVarRequestOptions.data =  needsSerialization
                ? JSON.stringify(botStartingItem !== undefined ? botStartingItem : {})
                : (botStartingItem || "");
            return {
                url: localVarUrlObj.pathname + localVarUrlObj.search + localVarUrlObj.hash,
                options: localVarRequestOptions,
            };
        },
        deleteBotStartingItem: async (id: number, options: any = {}): Promise<RequestArgs> => {
            if (id === null || id === undefined) {
                throw new RequiredError('id','Required parameter id was null or undefined when calling deleteBotStartingItem.');
            }
            const localVarPath = `/bot_starting_item/{id}`
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
        getBotStartingItem: async (id: number, includes?: string, select?: string, options: any = {}): Promise<RequestArgs> => {
            if (id === null || id === undefined) {
                throw new RequiredError('id','Required parameter id was null or undefined when calling getBotStartingItem.');
            }
            const localVarPath = `/bot_starting_item/{id}`
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
        getBotStartingItemsBulk: async (body: CrudcontrollersBulkFetchByIdsGetRequest, options: any = {}): Promise<RequestArgs> => {
            if (body === null || body === undefined) {
                throw new RequiredError('body','Required parameter body was null or undefined when calling getBotStartingItemsBulk.');
            }
            const localVarPath = `/bot_starting_items/bulk`;
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
        getBotStartingItemsCount: async (includes?: string, where?: string, whereOr?: string, groupBy?: string, limit?: string, page?: number, orderBy?: string, orderDirection?: string, select?: string, options: any = {}): Promise<RequestArgs> => {
            const localVarPath = `/bot_starting_items/count`;
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
        listBotStartingItems: async (includes?: string, where?: string, whereOr?: string, groupBy?: string, limit?: string, page?: number, orderBy?: string, orderDirection?: string, select?: string, options: any = {}): Promise<RequestArgs> => {
            const localVarPath = `/bot_starting_items`;
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
        updateBotStartingItem: async (id: number, botStartingItem: ModelsBotStartingItem, options: any = {}): Promise<RequestArgs> => {
            if (id === null || id === undefined) {
                throw new RequiredError('id','Required parameter id was null or undefined when calling updateBotStartingItem.');
            }
            if (botStartingItem === null || botStartingItem === undefined) {
                throw new RequiredError('botStartingItem','Required parameter botStartingItem was null or undefined when calling updateBotStartingItem.');
            }
            const localVarPath = `/bot_starting_item/{id}`
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
            const nonString = typeof botStartingItem !== 'string';
            const needsSerialization = nonString && configuration && configuration.isJsonMime
                ? configuration.isJsonMime(localVarRequestOptions.headers['Content-Type'])
                : nonString;
            localVarRequestOptions.data =  needsSerialization
                ? JSON.stringify(botStartingItem !== undefined ? botStartingItem : {})
                : (botStartingItem || "");
            return {
                url: localVarUrlObj.pathname + localVarUrlObj.search + localVarUrlObj.hash,
                options: localVarRequestOptions,
            };
        },
    }
};
export const BotStartingItemApiFp = function(configuration?: Configuration) {
    return {
        async createBotStartingItem(botStartingItem: ModelsBotStartingItem, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsBotStartingItem>>> {
            const localVarAxiosArgs = await BotStartingItemApiAxiosParamCreator(configuration).createBotStartingItem(botStartingItem, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        async deleteBotStartingItem(id: number, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<string>> {
            const localVarAxiosArgs = await BotStartingItemApiAxiosParamCreator(configuration).deleteBotStartingItem(id, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        async getBotStartingItem(id: number, includes?: string, select?: string, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsBotStartingItem>>> {
            const localVarAxiosArgs = await BotStartingItemApiAxiosParamCreator(configuration).getBotStartingItem(id, includes, select, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        async getBotStartingItemsBulk(body: CrudcontrollersBulkFetchByIdsGetRequest, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsBotStartingItem>>> {
            const localVarAxiosArgs = await BotStartingItemApiAxiosParamCreator(configuration).getBotStartingItemsBulk(body, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        async getBotStartingItemsCount(includes?: string, where?: string, whereOr?: string, groupBy?: string, limit?: string, page?: number, orderBy?: string, orderDirection?: string, select?: string, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsBotStartingItem>>> {
            const localVarAxiosArgs = await BotStartingItemApiAxiosParamCreator(configuration).getBotStartingItemsCount(includes, where, whereOr, groupBy, limit, page, orderBy, orderDirection, select, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        async listBotStartingItems(includes?: string, where?: string, whereOr?: string, groupBy?: string, limit?: string, page?: number, orderBy?: string, orderDirection?: string, select?: string, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsBotStartingItem>>> {
            const localVarAxiosArgs = await BotStartingItemApiAxiosParamCreator(configuration).listBotStartingItems(includes, where, whereOr, groupBy, limit, page, orderBy, orderDirection, select, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        async updateBotStartingItem(id: number, botStartingItem: ModelsBotStartingItem, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsBotStartingItem>>> {
            const localVarAxiosArgs = await BotStartingItemApiAxiosParamCreator(configuration).updateBotStartingItem(id, botStartingItem, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
    }
};
export const BotStartingItemApiFactory = function (configuration?: Configuration, basePath?: string, axios?: AxiosInstance) {
    return {
        createBotStartingItem(botStartingItem: ModelsBotStartingItem, options?: any): AxiosPromise<Array<ModelsBotStartingItem>> {
            return BotStartingItemApiFp(configuration).createBotStartingItem(botStartingItem, options).then((request) => request(axios, basePath));
        },
        deleteBotStartingItem(id: number, options?: any): AxiosPromise<string> {
            return BotStartingItemApiFp(configuration).deleteBotStartingItem(id, options).then((request) => request(axios, basePath));
        },
        getBotStartingItem(id: number, includes?: string, select?: string, options?: any): AxiosPromise<Array<ModelsBotStartingItem>> {
            return BotStartingItemApiFp(configuration).getBotStartingItem(id, includes, select, options).then((request) => request(axios, basePath));
        },
        getBotStartingItemsBulk(body: CrudcontrollersBulkFetchByIdsGetRequest, options?: any): AxiosPromise<Array<ModelsBotStartingItem>> {
            return BotStartingItemApiFp(configuration).getBotStartingItemsBulk(body, options).then((request) => request(axios, basePath));
        },
        getBotStartingItemsCount(includes?: string, where?: string, whereOr?: string, groupBy?: string, limit?: string, page?: number, orderBy?: string, orderDirection?: string, select?: string, options?: any): AxiosPromise<Array<ModelsBotStartingItem>> {
            return BotStartingItemApiFp(configuration).getBotStartingItemsCount(includes, where, whereOr, groupBy, limit, page, orderBy, orderDirection, select, options).then((request) => request(axios, basePath));
        },
        listBotStartingItems(includes?: string, where?: string, whereOr?: string, groupBy?: string, limit?: string, page?: number, orderBy?: string, orderDirection?: string, select?: string, options?: any): AxiosPromise<Array<ModelsBotStartingItem>> {
            return BotStartingItemApiFp(configuration).listBotStartingItems(includes, where, whereOr, groupBy, limit, page, orderBy, orderDirection, select, options).then((request) => request(axios, basePath));
        },
        updateBotStartingItem(id: number, botStartingItem: ModelsBotStartingItem, options?: any): AxiosPromise<Array<ModelsBotStartingItem>> {
            return BotStartingItemApiFp(configuration).updateBotStartingItem(id, botStartingItem, options).then((request) => request(axios, basePath));
        },
    };
};
export interface BotStartingItemApiCreateBotStartingItemRequest {
    readonly botStartingItem: ModelsBotStartingItem
}
export interface BotStartingItemApiDeleteBotStartingItemRequest {
    readonly id: number
}
export interface BotStartingItemApiGetBotStartingItemRequest {
    readonly id: number
    readonly includes?: string
    readonly select?: string
}
export interface BotStartingItemApiGetBotStartingItemsBulkRequest {
    readonly body: CrudcontrollersBulkFetchByIdsGetRequest
}
export interface BotStartingItemApiGetBotStartingItemsCountRequest {
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
export interface BotStartingItemApiListBotStartingItemsRequest {
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
export interface BotStartingItemApiUpdateBotStartingItemRequest {
    readonly id: number
    readonly botStartingItem: ModelsBotStartingItem
}
export class BotStartingItemApi extends BaseAPI {
    public createBotStartingItem(requestParameters: BotStartingItemApiCreateBotStartingItemRequest, options?: any) {
        return BotStartingItemApiFp(this.configuration).createBotStartingItem(requestParameters.botStartingItem, options).then((request) => request(this.axios, this.basePath));
    }
    public deleteBotStartingItem(requestParameters: BotStartingItemApiDeleteBotStartingItemRequest, options?: any) {
        return BotStartingItemApiFp(this.configuration).deleteBotStartingItem(requestParameters.id, options).then((request) => request(this.axios, this.basePath));
    }
    public getBotStartingItem(requestParameters: BotStartingItemApiGetBotStartingItemRequest, options?: any) {
        return BotStartingItemApiFp(this.configuration).getBotStartingItem(requestParameters.id, requestParameters.includes, requestParameters.select, options).then((request) => request(this.axios, this.basePath));
    }
    public getBotStartingItemsBulk(requestParameters: BotStartingItemApiGetBotStartingItemsBulkRequest, options?: any) {
        return BotStartingItemApiFp(this.configuration).getBotStartingItemsBulk(requestParameters.body, options).then((request) => request(this.axios, this.basePath));
    }
    public getBotStartingItemsCount(requestParameters: BotStartingItemApiGetBotStartingItemsCountRequest = {}, options?: any) {
        return BotStartingItemApiFp(this.configuration).getBotStartingItemsCount(requestParameters.includes, requestParameters.where, requestParameters.whereOr, requestParameters.groupBy, requestParameters.limit, requestParameters.page, requestParameters.orderBy, requestParameters.orderDirection, requestParameters.select, options).then((request) => request(this.axios, this.basePath));
    }
    public listBotStartingItems(requestParameters: BotStartingItemApiListBotStartingItemsRequest = {}, options?: any) {
        return BotStartingItemApiFp(this.configuration).listBotStartingItems(requestParameters.includes, requestParameters.where, requestParameters.whereOr, requestParameters.groupBy, requestParameters.limit, requestParameters.page, requestParameters.orderBy, requestParameters.orderDirection, requestParameters.select, options).then((request) => request(this.axios, this.basePath));
    }
    public updateBotStartingItem(requestParameters: BotStartingItemApiUpdateBotStartingItemRequest, options?: any) {
        return BotStartingItemApiFp(this.configuration).updateBotStartingItem(requestParameters.id, requestParameters.botStartingItem, options).then((request) => request(this.axios, this.basePath));
    }
}
