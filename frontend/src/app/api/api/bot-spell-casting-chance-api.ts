import globalAxios, { AxiosPromise, AxiosInstance } from 'axios';
import { Configuration } from '../configuration';
// Some imports not used depending on template conditions
// @ts-ignore
import { BASE_PATH, COLLECTION_FORMATS, RequestArgs, BaseAPI, RequiredError } from '../base';
// @ts-ignore
import { CrudcontrollersBulkFetchByIdsGetRequest } from '../models';
// @ts-ignore
import { ModelsBotSpellCastingChance } from '../models';
export const BotSpellCastingChanceApiAxiosParamCreator = function (configuration?: Configuration) {
    return {
        createBotSpellCastingChance: async (botSpellCastingChance: ModelsBotSpellCastingChance, options: any = {}): Promise<RequestArgs> => {
            if (botSpellCastingChance === null || botSpellCastingChance === undefined) {
                throw new RequiredError('botSpellCastingChance','Required parameter botSpellCastingChance was null or undefined when calling createBotSpellCastingChance.');
            }
            const localVarPath = `/bot_spell_casting_chance`;
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
            const nonString = typeof botSpellCastingChance !== 'string';
            const needsSerialization = nonString && configuration && configuration.isJsonMime
                ? configuration.isJsonMime(localVarRequestOptions.headers['Content-Type'])
                : nonString;
            localVarRequestOptions.data =  needsSerialization
                ? JSON.stringify(botSpellCastingChance !== undefined ? botSpellCastingChance : {})
                : (botSpellCastingChance || "");
            return {
                url: localVarUrlObj.pathname + localVarUrlObj.search + localVarUrlObj.hash,
                options: localVarRequestOptions,
            };
        },
        deleteBotSpellCastingChance: async (id: number, options: any = {}): Promise<RequestArgs> => {
            if (id === null || id === undefined) {
                throw new RequiredError('id','Required parameter id was null or undefined when calling deleteBotSpellCastingChance.');
            }
            const localVarPath = `/bot_spell_casting_chance/{id}`
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
        getBotSpellCastingChance: async (id: number, includes?: string, select?: string, options: any = {}): Promise<RequestArgs> => {
            if (id === null || id === undefined) {
                throw new RequiredError('id','Required parameter id was null or undefined when calling getBotSpellCastingChance.');
            }
            const localVarPath = `/bot_spell_casting_chance/{id}`
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
        getBotSpellCastingChancesBulk: async (body: CrudcontrollersBulkFetchByIdsGetRequest, options: any = {}): Promise<RequestArgs> => {
            if (body === null || body === undefined) {
                throw new RequiredError('body','Required parameter body was null or undefined when calling getBotSpellCastingChancesBulk.');
            }
            const localVarPath = `/bot_spell_casting_chances/bulk`;
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
        getBotSpellCastingChancesCount: async (includes?: string, where?: string, whereOr?: string, groupBy?: string, limit?: string, page?: number, orderBy?: string, orderDirection?: string, select?: string, options: any = {}): Promise<RequestArgs> => {
            const localVarPath = `/bot_spell_casting_chances/count`;
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
        listBotSpellCastingChances: async (includes?: string, where?: string, whereOr?: string, groupBy?: string, limit?: string, page?: number, orderBy?: string, orderDirection?: string, select?: string, options: any = {}): Promise<RequestArgs> => {
            const localVarPath = `/bot_spell_casting_chances`;
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
        updateBotSpellCastingChance: async (id: number, botSpellCastingChance: ModelsBotSpellCastingChance, options: any = {}): Promise<RequestArgs> => {
            if (id === null || id === undefined) {
                throw new RequiredError('id','Required parameter id was null or undefined when calling updateBotSpellCastingChance.');
            }
            if (botSpellCastingChance === null || botSpellCastingChance === undefined) {
                throw new RequiredError('botSpellCastingChance','Required parameter botSpellCastingChance was null or undefined when calling updateBotSpellCastingChance.');
            }
            const localVarPath = `/bot_spell_casting_chance/{id}`
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
            const nonString = typeof botSpellCastingChance !== 'string';
            const needsSerialization = nonString && configuration && configuration.isJsonMime
                ? configuration.isJsonMime(localVarRequestOptions.headers['Content-Type'])
                : nonString;
            localVarRequestOptions.data =  needsSerialization
                ? JSON.stringify(botSpellCastingChance !== undefined ? botSpellCastingChance : {})
                : (botSpellCastingChance || "");
            return {
                url: localVarUrlObj.pathname + localVarUrlObj.search + localVarUrlObj.hash,
                options: localVarRequestOptions,
            };
        },
    }
};
export const BotSpellCastingChanceApiFp = function(configuration?: Configuration) {
    return {
        async createBotSpellCastingChance(botSpellCastingChance: ModelsBotSpellCastingChance, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsBotSpellCastingChance>>> {
            const localVarAxiosArgs = await BotSpellCastingChanceApiAxiosParamCreator(configuration).createBotSpellCastingChance(botSpellCastingChance, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        async deleteBotSpellCastingChance(id: number, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<string>> {
            const localVarAxiosArgs = await BotSpellCastingChanceApiAxiosParamCreator(configuration).deleteBotSpellCastingChance(id, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        async getBotSpellCastingChance(id: number, includes?: string, select?: string, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsBotSpellCastingChance>>> {
            const localVarAxiosArgs = await BotSpellCastingChanceApiAxiosParamCreator(configuration).getBotSpellCastingChance(id, includes, select, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        async getBotSpellCastingChancesBulk(body: CrudcontrollersBulkFetchByIdsGetRequest, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsBotSpellCastingChance>>> {
            const localVarAxiosArgs = await BotSpellCastingChanceApiAxiosParamCreator(configuration).getBotSpellCastingChancesBulk(body, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        async getBotSpellCastingChancesCount(includes?: string, where?: string, whereOr?: string, groupBy?: string, limit?: string, page?: number, orderBy?: string, orderDirection?: string, select?: string, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsBotSpellCastingChance>>> {
            const localVarAxiosArgs = await BotSpellCastingChanceApiAxiosParamCreator(configuration).getBotSpellCastingChancesCount(includes, where, whereOr, groupBy, limit, page, orderBy, orderDirection, select, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        async listBotSpellCastingChances(includes?: string, where?: string, whereOr?: string, groupBy?: string, limit?: string, page?: number, orderBy?: string, orderDirection?: string, select?: string, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsBotSpellCastingChance>>> {
            const localVarAxiosArgs = await BotSpellCastingChanceApiAxiosParamCreator(configuration).listBotSpellCastingChances(includes, where, whereOr, groupBy, limit, page, orderBy, orderDirection, select, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        async updateBotSpellCastingChance(id: number, botSpellCastingChance: ModelsBotSpellCastingChance, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsBotSpellCastingChance>>> {
            const localVarAxiosArgs = await BotSpellCastingChanceApiAxiosParamCreator(configuration).updateBotSpellCastingChance(id, botSpellCastingChance, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
    }
};
export const BotSpellCastingChanceApiFactory = function (configuration?: Configuration, basePath?: string, axios?: AxiosInstance) {
    return {
        createBotSpellCastingChance(botSpellCastingChance: ModelsBotSpellCastingChance, options?: any): AxiosPromise<Array<ModelsBotSpellCastingChance>> {
            return BotSpellCastingChanceApiFp(configuration).createBotSpellCastingChance(botSpellCastingChance, options).then((request) => request(axios, basePath));
        },
        deleteBotSpellCastingChance(id: number, options?: any): AxiosPromise<string> {
            return BotSpellCastingChanceApiFp(configuration).deleteBotSpellCastingChance(id, options).then((request) => request(axios, basePath));
        },
        getBotSpellCastingChance(id: number, includes?: string, select?: string, options?: any): AxiosPromise<Array<ModelsBotSpellCastingChance>> {
            return BotSpellCastingChanceApiFp(configuration).getBotSpellCastingChance(id, includes, select, options).then((request) => request(axios, basePath));
        },
        getBotSpellCastingChancesBulk(body: CrudcontrollersBulkFetchByIdsGetRequest, options?: any): AxiosPromise<Array<ModelsBotSpellCastingChance>> {
            return BotSpellCastingChanceApiFp(configuration).getBotSpellCastingChancesBulk(body, options).then((request) => request(axios, basePath));
        },
        getBotSpellCastingChancesCount(includes?: string, where?: string, whereOr?: string, groupBy?: string, limit?: string, page?: number, orderBy?: string, orderDirection?: string, select?: string, options?: any): AxiosPromise<Array<ModelsBotSpellCastingChance>> {
            return BotSpellCastingChanceApiFp(configuration).getBotSpellCastingChancesCount(includes, where, whereOr, groupBy, limit, page, orderBy, orderDirection, select, options).then((request) => request(axios, basePath));
        },
        listBotSpellCastingChances(includes?: string, where?: string, whereOr?: string, groupBy?: string, limit?: string, page?: number, orderBy?: string, orderDirection?: string, select?: string, options?: any): AxiosPromise<Array<ModelsBotSpellCastingChance>> {
            return BotSpellCastingChanceApiFp(configuration).listBotSpellCastingChances(includes, where, whereOr, groupBy, limit, page, orderBy, orderDirection, select, options).then((request) => request(axios, basePath));
        },
        updateBotSpellCastingChance(id: number, botSpellCastingChance: ModelsBotSpellCastingChance, options?: any): AxiosPromise<Array<ModelsBotSpellCastingChance>> {
            return BotSpellCastingChanceApiFp(configuration).updateBotSpellCastingChance(id, botSpellCastingChance, options).then((request) => request(axios, basePath));
        },
    };
};
export interface BotSpellCastingChanceApiCreateBotSpellCastingChanceRequest {
    readonly botSpellCastingChance: ModelsBotSpellCastingChance
}
export interface BotSpellCastingChanceApiDeleteBotSpellCastingChanceRequest {
    readonly id: number
}
export interface BotSpellCastingChanceApiGetBotSpellCastingChanceRequest {
    readonly id: number
    readonly includes?: string
    readonly select?: string
}
export interface BotSpellCastingChanceApiGetBotSpellCastingChancesBulkRequest {
    readonly body: CrudcontrollersBulkFetchByIdsGetRequest
}
export interface BotSpellCastingChanceApiGetBotSpellCastingChancesCountRequest {
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
export interface BotSpellCastingChanceApiListBotSpellCastingChancesRequest {
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
export interface BotSpellCastingChanceApiUpdateBotSpellCastingChanceRequest {
    readonly id: number
    readonly botSpellCastingChance: ModelsBotSpellCastingChance
}
export class BotSpellCastingChanceApi extends BaseAPI {
    public createBotSpellCastingChance(requestParameters: BotSpellCastingChanceApiCreateBotSpellCastingChanceRequest, options?: any) {
        return BotSpellCastingChanceApiFp(this.configuration).createBotSpellCastingChance(requestParameters.botSpellCastingChance, options).then((request) => request(this.axios, this.basePath));
    }
    public deleteBotSpellCastingChance(requestParameters: BotSpellCastingChanceApiDeleteBotSpellCastingChanceRequest, options?: any) {
        return BotSpellCastingChanceApiFp(this.configuration).deleteBotSpellCastingChance(requestParameters.id, options).then((request) => request(this.axios, this.basePath));
    }
    public getBotSpellCastingChance(requestParameters: BotSpellCastingChanceApiGetBotSpellCastingChanceRequest, options?: any) {
        return BotSpellCastingChanceApiFp(this.configuration).getBotSpellCastingChance(requestParameters.id, requestParameters.includes, requestParameters.select, options).then((request) => request(this.axios, this.basePath));
    }
    public getBotSpellCastingChancesBulk(requestParameters: BotSpellCastingChanceApiGetBotSpellCastingChancesBulkRequest, options?: any) {
        return BotSpellCastingChanceApiFp(this.configuration).getBotSpellCastingChancesBulk(requestParameters.body, options).then((request) => request(this.axios, this.basePath));
    }
    public getBotSpellCastingChancesCount(requestParameters: BotSpellCastingChanceApiGetBotSpellCastingChancesCountRequest = {}, options?: any) {
        return BotSpellCastingChanceApiFp(this.configuration).getBotSpellCastingChancesCount(requestParameters.includes, requestParameters.where, requestParameters.whereOr, requestParameters.groupBy, requestParameters.limit, requestParameters.page, requestParameters.orderBy, requestParameters.orderDirection, requestParameters.select, options).then((request) => request(this.axios, this.basePath));
    }
    public listBotSpellCastingChances(requestParameters: BotSpellCastingChanceApiListBotSpellCastingChancesRequest = {}, options?: any) {
        return BotSpellCastingChanceApiFp(this.configuration).listBotSpellCastingChances(requestParameters.includes, requestParameters.where, requestParameters.whereOr, requestParameters.groupBy, requestParameters.limit, requestParameters.page, requestParameters.orderBy, requestParameters.orderDirection, requestParameters.select, options).then((request) => request(this.axios, this.basePath));
    }
    public updateBotSpellCastingChance(requestParameters: BotSpellCastingChanceApiUpdateBotSpellCastingChanceRequest, options?: any) {
        return BotSpellCastingChanceApiFp(this.configuration).updateBotSpellCastingChance(requestParameters.id, requestParameters.botSpellCastingChance, options).then((request) => request(this.axios, this.basePath));
    }
}
