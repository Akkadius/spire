import globalAxios, { AxiosPromise, AxiosInstance } from 'axios';
import { Configuration } from '../configuration';
// Some imports not used depending on template conditions
// @ts-ignore
import { BASE_PATH, COLLECTION_FORMATS, RequestArgs, BaseAPI, RequiredError } from '../base';
// @ts-ignore
import { CrudcontrollersBulkFetchByIdsGetRequest } from '../models';
// @ts-ignore
import { ModelsCharacterCorpseItem } from '../models';
export const CharacterCorpseItemApiAxiosParamCreator = function (configuration?: Configuration) {
    return {
        createCharacterCorpseItem: async (characterCorpseItem: ModelsCharacterCorpseItem, options: any = {}): Promise<RequestArgs> => {
            if (characterCorpseItem === null || characterCorpseItem === undefined) {
                throw new RequiredError('characterCorpseItem','Required parameter characterCorpseItem was null or undefined when calling createCharacterCorpseItem.');
            }
            const localVarPath = `/character_corpse_item`;
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
            const nonString = typeof characterCorpseItem !== 'string';
            const needsSerialization = nonString && configuration && configuration.isJsonMime
                ? configuration.isJsonMime(localVarRequestOptions.headers['Content-Type'])
                : nonString;
            localVarRequestOptions.data =  needsSerialization
                ? JSON.stringify(characterCorpseItem !== undefined ? characterCorpseItem : {})
                : (characterCorpseItem || "");
            return {
                url: localVarUrlObj.pathname + localVarUrlObj.search + localVarUrlObj.hash,
                options: localVarRequestOptions,
            };
        },
        deleteCharacterCorpseItem: async (id: number, options: any = {}): Promise<RequestArgs> => {
            if (id === null || id === undefined) {
                throw new RequiredError('id','Required parameter id was null or undefined when calling deleteCharacterCorpseItem.');
            }
            const localVarPath = `/character_corpse_item/{id}`
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
        getCharacterCorpseItem: async (id: number, includes?: string, select?: string, options: any = {}): Promise<RequestArgs> => {
            if (id === null || id === undefined) {
                throw new RequiredError('id','Required parameter id was null or undefined when calling getCharacterCorpseItem.');
            }
            const localVarPath = `/character_corpse_item/{id}`
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
        getCharacterCorpseItemsBulk: async (body: CrudcontrollersBulkFetchByIdsGetRequest, options: any = {}): Promise<RequestArgs> => {
            if (body === null || body === undefined) {
                throw new RequiredError('body','Required parameter body was null or undefined when calling getCharacterCorpseItemsBulk.');
            }
            const localVarPath = `/character_corpse_items/bulk`;
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
        getCharacterCorpseItemsCount: async (includes?: string, where?: string, whereOr?: string, groupBy?: string, limit?: string, page?: number, orderBy?: string, orderDirection?: string, select?: string, options: any = {}): Promise<RequestArgs> => {
            const localVarPath = `/character_corpse_items/count`;
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
        listCharacterCorpseItems: async (includes?: string, where?: string, whereOr?: string, groupBy?: string, limit?: string, page?: number, orderBy?: string, orderDirection?: string, select?: string, options: any = {}): Promise<RequestArgs> => {
            const localVarPath = `/character_corpse_items`;
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
        updateCharacterCorpseItem: async (id: number, characterCorpseItem: ModelsCharacterCorpseItem, options: any = {}): Promise<RequestArgs> => {
            if (id === null || id === undefined) {
                throw new RequiredError('id','Required parameter id was null or undefined when calling updateCharacterCorpseItem.');
            }
            if (characterCorpseItem === null || characterCorpseItem === undefined) {
                throw new RequiredError('characterCorpseItem','Required parameter characterCorpseItem was null or undefined when calling updateCharacterCorpseItem.');
            }
            const localVarPath = `/character_corpse_item/{id}`
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
            const nonString = typeof characterCorpseItem !== 'string';
            const needsSerialization = nonString && configuration && configuration.isJsonMime
                ? configuration.isJsonMime(localVarRequestOptions.headers['Content-Type'])
                : nonString;
            localVarRequestOptions.data =  needsSerialization
                ? JSON.stringify(characterCorpseItem !== undefined ? characterCorpseItem : {})
                : (characterCorpseItem || "");
            return {
                url: localVarUrlObj.pathname + localVarUrlObj.search + localVarUrlObj.hash,
                options: localVarRequestOptions,
            };
        },
    }
};
export const CharacterCorpseItemApiFp = function(configuration?: Configuration) {
    return {
        async createCharacterCorpseItem(characterCorpseItem: ModelsCharacterCorpseItem, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsCharacterCorpseItem>>> {
            const localVarAxiosArgs = await CharacterCorpseItemApiAxiosParamCreator(configuration).createCharacterCorpseItem(characterCorpseItem, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        async deleteCharacterCorpseItem(id: number, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<string>> {
            const localVarAxiosArgs = await CharacterCorpseItemApiAxiosParamCreator(configuration).deleteCharacterCorpseItem(id, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        async getCharacterCorpseItem(id: number, includes?: string, select?: string, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsCharacterCorpseItem>>> {
            const localVarAxiosArgs = await CharacterCorpseItemApiAxiosParamCreator(configuration).getCharacterCorpseItem(id, includes, select, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        async getCharacterCorpseItemsBulk(body: CrudcontrollersBulkFetchByIdsGetRequest, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsCharacterCorpseItem>>> {
            const localVarAxiosArgs = await CharacterCorpseItemApiAxiosParamCreator(configuration).getCharacterCorpseItemsBulk(body, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        async getCharacterCorpseItemsCount(includes?: string, where?: string, whereOr?: string, groupBy?: string, limit?: string, page?: number, orderBy?: string, orderDirection?: string, select?: string, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsCharacterCorpseItem>>> {
            const localVarAxiosArgs = await CharacterCorpseItemApiAxiosParamCreator(configuration).getCharacterCorpseItemsCount(includes, where, whereOr, groupBy, limit, page, orderBy, orderDirection, select, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        async listCharacterCorpseItems(includes?: string, where?: string, whereOr?: string, groupBy?: string, limit?: string, page?: number, orderBy?: string, orderDirection?: string, select?: string, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsCharacterCorpseItem>>> {
            const localVarAxiosArgs = await CharacterCorpseItemApiAxiosParamCreator(configuration).listCharacterCorpseItems(includes, where, whereOr, groupBy, limit, page, orderBy, orderDirection, select, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        async updateCharacterCorpseItem(id: number, characterCorpseItem: ModelsCharacterCorpseItem, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsCharacterCorpseItem>>> {
            const localVarAxiosArgs = await CharacterCorpseItemApiAxiosParamCreator(configuration).updateCharacterCorpseItem(id, characterCorpseItem, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
    }
};
export const CharacterCorpseItemApiFactory = function (configuration?: Configuration, basePath?: string, axios?: AxiosInstance) {
    return {
        createCharacterCorpseItem(characterCorpseItem: ModelsCharacterCorpseItem, options?: any): AxiosPromise<Array<ModelsCharacterCorpseItem>> {
            return CharacterCorpseItemApiFp(configuration).createCharacterCorpseItem(characterCorpseItem, options).then((request) => request(axios, basePath));
        },
        deleteCharacterCorpseItem(id: number, options?: any): AxiosPromise<string> {
            return CharacterCorpseItemApiFp(configuration).deleteCharacterCorpseItem(id, options).then((request) => request(axios, basePath));
        },
        getCharacterCorpseItem(id: number, includes?: string, select?: string, options?: any): AxiosPromise<Array<ModelsCharacterCorpseItem>> {
            return CharacterCorpseItemApiFp(configuration).getCharacterCorpseItem(id, includes, select, options).then((request) => request(axios, basePath));
        },
        getCharacterCorpseItemsBulk(body: CrudcontrollersBulkFetchByIdsGetRequest, options?: any): AxiosPromise<Array<ModelsCharacterCorpseItem>> {
            return CharacterCorpseItemApiFp(configuration).getCharacterCorpseItemsBulk(body, options).then((request) => request(axios, basePath));
        },
        getCharacterCorpseItemsCount(includes?: string, where?: string, whereOr?: string, groupBy?: string, limit?: string, page?: number, orderBy?: string, orderDirection?: string, select?: string, options?: any): AxiosPromise<Array<ModelsCharacterCorpseItem>> {
            return CharacterCorpseItemApiFp(configuration).getCharacterCorpseItemsCount(includes, where, whereOr, groupBy, limit, page, orderBy, orderDirection, select, options).then((request) => request(axios, basePath));
        },
        listCharacterCorpseItems(includes?: string, where?: string, whereOr?: string, groupBy?: string, limit?: string, page?: number, orderBy?: string, orderDirection?: string, select?: string, options?: any): AxiosPromise<Array<ModelsCharacterCorpseItem>> {
            return CharacterCorpseItemApiFp(configuration).listCharacterCorpseItems(includes, where, whereOr, groupBy, limit, page, orderBy, orderDirection, select, options).then((request) => request(axios, basePath));
        },
        updateCharacterCorpseItem(id: number, characterCorpseItem: ModelsCharacterCorpseItem, options?: any): AxiosPromise<Array<ModelsCharacterCorpseItem>> {
            return CharacterCorpseItemApiFp(configuration).updateCharacterCorpseItem(id, characterCorpseItem, options).then((request) => request(axios, basePath));
        },
    };
};
export interface CharacterCorpseItemApiCreateCharacterCorpseItemRequest {
    readonly characterCorpseItem: ModelsCharacterCorpseItem
}
export interface CharacterCorpseItemApiDeleteCharacterCorpseItemRequest {
    readonly id: number
}
export interface CharacterCorpseItemApiGetCharacterCorpseItemRequest {
    readonly id: number
    readonly includes?: string
    readonly select?: string
}
export interface CharacterCorpseItemApiGetCharacterCorpseItemsBulkRequest {
    readonly body: CrudcontrollersBulkFetchByIdsGetRequest
}
export interface CharacterCorpseItemApiGetCharacterCorpseItemsCountRequest {
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
export interface CharacterCorpseItemApiListCharacterCorpseItemsRequest {
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
export interface CharacterCorpseItemApiUpdateCharacterCorpseItemRequest {
    readonly id: number
    readonly characterCorpseItem: ModelsCharacterCorpseItem
}
export class CharacterCorpseItemApi extends BaseAPI {
    public createCharacterCorpseItem(requestParameters: CharacterCorpseItemApiCreateCharacterCorpseItemRequest, options?: any) {
        return CharacterCorpseItemApiFp(this.configuration).createCharacterCorpseItem(requestParameters.characterCorpseItem, options).then((request) => request(this.axios, this.basePath));
    }
    public deleteCharacterCorpseItem(requestParameters: CharacterCorpseItemApiDeleteCharacterCorpseItemRequest, options?: any) {
        return CharacterCorpseItemApiFp(this.configuration).deleteCharacterCorpseItem(requestParameters.id, options).then((request) => request(this.axios, this.basePath));
    }
    public getCharacterCorpseItem(requestParameters: CharacterCorpseItemApiGetCharacterCorpseItemRequest, options?: any) {
        return CharacterCorpseItemApiFp(this.configuration).getCharacterCorpseItem(requestParameters.id, requestParameters.includes, requestParameters.select, options).then((request) => request(this.axios, this.basePath));
    }
    public getCharacterCorpseItemsBulk(requestParameters: CharacterCorpseItemApiGetCharacterCorpseItemsBulkRequest, options?: any) {
        return CharacterCorpseItemApiFp(this.configuration).getCharacterCorpseItemsBulk(requestParameters.body, options).then((request) => request(this.axios, this.basePath));
    }
    public getCharacterCorpseItemsCount(requestParameters: CharacterCorpseItemApiGetCharacterCorpseItemsCountRequest = {}, options?: any) {
        return CharacterCorpseItemApiFp(this.configuration).getCharacterCorpseItemsCount(requestParameters.includes, requestParameters.where, requestParameters.whereOr, requestParameters.groupBy, requestParameters.limit, requestParameters.page, requestParameters.orderBy, requestParameters.orderDirection, requestParameters.select, options).then((request) => request(this.axios, this.basePath));
    }
    public listCharacterCorpseItems(requestParameters: CharacterCorpseItemApiListCharacterCorpseItemsRequest = {}, options?: any) {
        return CharacterCorpseItemApiFp(this.configuration).listCharacterCorpseItems(requestParameters.includes, requestParameters.where, requestParameters.whereOr, requestParameters.groupBy, requestParameters.limit, requestParameters.page, requestParameters.orderBy, requestParameters.orderDirection, requestParameters.select, options).then((request) => request(this.axios, this.basePath));
    }
    public updateCharacterCorpseItem(requestParameters: CharacterCorpseItemApiUpdateCharacterCorpseItemRequest, options?: any) {
        return CharacterCorpseItemApiFp(this.configuration).updateCharacterCorpseItem(requestParameters.id, requestParameters.characterCorpseItem, options).then((request) => request(this.axios, this.basePath));
    }
}
