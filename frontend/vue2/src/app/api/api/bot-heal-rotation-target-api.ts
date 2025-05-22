import globalAxios, { AxiosPromise, AxiosInstance } from 'axios';
import { Configuration } from '../configuration';
// Some imports not used depending on template conditions
// @ts-ignore
import { BASE_PATH, COLLECTION_FORMATS, RequestArgs, BaseAPI, RequiredError } from '../base';
// @ts-ignore
import { CrudcontrollersBulkFetchByIdsGetRequest } from '../models';
// @ts-ignore
import { ModelsBotHealRotationTarget } from '../models';
export const BotHealRotationTargetApiAxiosParamCreator = function (configuration?: Configuration) {
    return {
        createBotHealRotationTarget: async (botHealRotationTarget: ModelsBotHealRotationTarget, options: any = {}): Promise<RequestArgs> => {
            if (botHealRotationTarget === null || botHealRotationTarget === undefined) {
                throw new RequiredError('botHealRotationTarget','Required parameter botHealRotationTarget was null or undefined when calling createBotHealRotationTarget.');
            }
            const localVarPath = `/bot_heal_rotation_target`;
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
            const nonString = typeof botHealRotationTarget !== 'string';
            const needsSerialization = nonString && configuration && configuration.isJsonMime
                ? configuration.isJsonMime(localVarRequestOptions.headers['Content-Type'])
                : nonString;
            localVarRequestOptions.data =  needsSerialization
                ? JSON.stringify(botHealRotationTarget !== undefined ? botHealRotationTarget : {})
                : (botHealRotationTarget || "");
            return {
                url: localVarUrlObj.pathname + localVarUrlObj.search + localVarUrlObj.hash,
                options: localVarRequestOptions,
            };
        },
        deleteBotHealRotationTarget: async (id: number, options: any = {}): Promise<RequestArgs> => {
            if (id === null || id === undefined) {
                throw new RequiredError('id','Required parameter id was null or undefined when calling deleteBotHealRotationTarget.');
            }
            const localVarPath = `/bot_heal_rotation_target/{id}`
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
        getBotHealRotationTarget: async (id: number, includes?: string, select?: string, options: any = {}): Promise<RequestArgs> => {
            if (id === null || id === undefined) {
                throw new RequiredError('id','Required parameter id was null or undefined when calling getBotHealRotationTarget.');
            }
            const localVarPath = `/bot_heal_rotation_target/{id}`
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
        getBotHealRotationTargetsBulk: async (body: CrudcontrollersBulkFetchByIdsGetRequest, options: any = {}): Promise<RequestArgs> => {
            if (body === null || body === undefined) {
                throw new RequiredError('body','Required parameter body was null or undefined when calling getBotHealRotationTargetsBulk.');
            }
            const localVarPath = `/bot_heal_rotation_targets/bulk`;
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
        getBotHealRotationTargetsCount: async (includes?: string, where?: string, whereOr?: string, groupBy?: string, limit?: string, page?: number, orderBy?: string, orderDirection?: string, select?: string, options: any = {}): Promise<RequestArgs> => {
            const localVarPath = `/bot_heal_rotation_targets/count`;
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
        listBotHealRotationTargets: async (includes?: string, where?: string, whereOr?: string, groupBy?: string, limit?: string, page?: number, orderBy?: string, orderDirection?: string, select?: string, options: any = {}): Promise<RequestArgs> => {
            const localVarPath = `/bot_heal_rotation_targets`;
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
        updateBotHealRotationTarget: async (id: number, botHealRotationTarget: ModelsBotHealRotationTarget, options: any = {}): Promise<RequestArgs> => {
            if (id === null || id === undefined) {
                throw new RequiredError('id','Required parameter id was null or undefined when calling updateBotHealRotationTarget.');
            }
            if (botHealRotationTarget === null || botHealRotationTarget === undefined) {
                throw new RequiredError('botHealRotationTarget','Required parameter botHealRotationTarget was null or undefined when calling updateBotHealRotationTarget.');
            }
            const localVarPath = `/bot_heal_rotation_target/{id}`
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
            const nonString = typeof botHealRotationTarget !== 'string';
            const needsSerialization = nonString && configuration && configuration.isJsonMime
                ? configuration.isJsonMime(localVarRequestOptions.headers['Content-Type'])
                : nonString;
            localVarRequestOptions.data =  needsSerialization
                ? JSON.stringify(botHealRotationTarget !== undefined ? botHealRotationTarget : {})
                : (botHealRotationTarget || "");
            return {
                url: localVarUrlObj.pathname + localVarUrlObj.search + localVarUrlObj.hash,
                options: localVarRequestOptions,
            };
        },
    }
};
export const BotHealRotationTargetApiFp = function(configuration?: Configuration) {
    return {
        async createBotHealRotationTarget(botHealRotationTarget: ModelsBotHealRotationTarget, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsBotHealRotationTarget>>> {
            const localVarAxiosArgs = await BotHealRotationTargetApiAxiosParamCreator(configuration).createBotHealRotationTarget(botHealRotationTarget, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        async deleteBotHealRotationTarget(id: number, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<string>> {
            const localVarAxiosArgs = await BotHealRotationTargetApiAxiosParamCreator(configuration).deleteBotHealRotationTarget(id, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        async getBotHealRotationTarget(id: number, includes?: string, select?: string, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsBotHealRotationTarget>>> {
            const localVarAxiosArgs = await BotHealRotationTargetApiAxiosParamCreator(configuration).getBotHealRotationTarget(id, includes, select, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        async getBotHealRotationTargetsBulk(body: CrudcontrollersBulkFetchByIdsGetRequest, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsBotHealRotationTarget>>> {
            const localVarAxiosArgs = await BotHealRotationTargetApiAxiosParamCreator(configuration).getBotHealRotationTargetsBulk(body, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        async getBotHealRotationTargetsCount(includes?: string, where?: string, whereOr?: string, groupBy?: string, limit?: string, page?: number, orderBy?: string, orderDirection?: string, select?: string, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsBotHealRotationTarget>>> {
            const localVarAxiosArgs = await BotHealRotationTargetApiAxiosParamCreator(configuration).getBotHealRotationTargetsCount(includes, where, whereOr, groupBy, limit, page, orderBy, orderDirection, select, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        async listBotHealRotationTargets(includes?: string, where?: string, whereOr?: string, groupBy?: string, limit?: string, page?: number, orderBy?: string, orderDirection?: string, select?: string, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsBotHealRotationTarget>>> {
            const localVarAxiosArgs = await BotHealRotationTargetApiAxiosParamCreator(configuration).listBotHealRotationTargets(includes, where, whereOr, groupBy, limit, page, orderBy, orderDirection, select, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        async updateBotHealRotationTarget(id: number, botHealRotationTarget: ModelsBotHealRotationTarget, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsBotHealRotationTarget>>> {
            const localVarAxiosArgs = await BotHealRotationTargetApiAxiosParamCreator(configuration).updateBotHealRotationTarget(id, botHealRotationTarget, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
    }
};
export const BotHealRotationTargetApiFactory = function (configuration?: Configuration, basePath?: string, axios?: AxiosInstance) {
    return {
        createBotHealRotationTarget(botHealRotationTarget: ModelsBotHealRotationTarget, options?: any): AxiosPromise<Array<ModelsBotHealRotationTarget>> {
            return BotHealRotationTargetApiFp(configuration).createBotHealRotationTarget(botHealRotationTarget, options).then((request) => request(axios, basePath));
        },
        deleteBotHealRotationTarget(id: number, options?: any): AxiosPromise<string> {
            return BotHealRotationTargetApiFp(configuration).deleteBotHealRotationTarget(id, options).then((request) => request(axios, basePath));
        },
        getBotHealRotationTarget(id: number, includes?: string, select?: string, options?: any): AxiosPromise<Array<ModelsBotHealRotationTarget>> {
            return BotHealRotationTargetApiFp(configuration).getBotHealRotationTarget(id, includes, select, options).then((request) => request(axios, basePath));
        },
        getBotHealRotationTargetsBulk(body: CrudcontrollersBulkFetchByIdsGetRequest, options?: any): AxiosPromise<Array<ModelsBotHealRotationTarget>> {
            return BotHealRotationTargetApiFp(configuration).getBotHealRotationTargetsBulk(body, options).then((request) => request(axios, basePath));
        },
        getBotHealRotationTargetsCount(includes?: string, where?: string, whereOr?: string, groupBy?: string, limit?: string, page?: number, orderBy?: string, orderDirection?: string, select?: string, options?: any): AxiosPromise<Array<ModelsBotHealRotationTarget>> {
            return BotHealRotationTargetApiFp(configuration).getBotHealRotationTargetsCount(includes, where, whereOr, groupBy, limit, page, orderBy, orderDirection, select, options).then((request) => request(axios, basePath));
        },
        listBotHealRotationTargets(includes?: string, where?: string, whereOr?: string, groupBy?: string, limit?: string, page?: number, orderBy?: string, orderDirection?: string, select?: string, options?: any): AxiosPromise<Array<ModelsBotHealRotationTarget>> {
            return BotHealRotationTargetApiFp(configuration).listBotHealRotationTargets(includes, where, whereOr, groupBy, limit, page, orderBy, orderDirection, select, options).then((request) => request(axios, basePath));
        },
        updateBotHealRotationTarget(id: number, botHealRotationTarget: ModelsBotHealRotationTarget, options?: any): AxiosPromise<Array<ModelsBotHealRotationTarget>> {
            return BotHealRotationTargetApiFp(configuration).updateBotHealRotationTarget(id, botHealRotationTarget, options).then((request) => request(axios, basePath));
        },
    };
};
export interface BotHealRotationTargetApiCreateBotHealRotationTargetRequest {
    readonly botHealRotationTarget: ModelsBotHealRotationTarget
}
export interface BotHealRotationTargetApiDeleteBotHealRotationTargetRequest {
    readonly id: number
}
export interface BotHealRotationTargetApiGetBotHealRotationTargetRequest {
    readonly id: number
    readonly includes?: string
    readonly select?: string
}
export interface BotHealRotationTargetApiGetBotHealRotationTargetsBulkRequest {
    readonly body: CrudcontrollersBulkFetchByIdsGetRequest
}
export interface BotHealRotationTargetApiGetBotHealRotationTargetsCountRequest {
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
export interface BotHealRotationTargetApiListBotHealRotationTargetsRequest {
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
export interface BotHealRotationTargetApiUpdateBotHealRotationTargetRequest {
    readonly id: number
    readonly botHealRotationTarget: ModelsBotHealRotationTarget
}
export class BotHealRotationTargetApi extends BaseAPI {
    public createBotHealRotationTarget(requestParameters: BotHealRotationTargetApiCreateBotHealRotationTargetRequest, options?: any) {
        return BotHealRotationTargetApiFp(this.configuration).createBotHealRotationTarget(requestParameters.botHealRotationTarget, options).then((request) => request(this.axios, this.basePath));
    }
    public deleteBotHealRotationTarget(requestParameters: BotHealRotationTargetApiDeleteBotHealRotationTargetRequest, options?: any) {
        return BotHealRotationTargetApiFp(this.configuration).deleteBotHealRotationTarget(requestParameters.id, options).then((request) => request(this.axios, this.basePath));
    }
    public getBotHealRotationTarget(requestParameters: BotHealRotationTargetApiGetBotHealRotationTargetRequest, options?: any) {
        return BotHealRotationTargetApiFp(this.configuration).getBotHealRotationTarget(requestParameters.id, requestParameters.includes, requestParameters.select, options).then((request) => request(this.axios, this.basePath));
    }
    public getBotHealRotationTargetsBulk(requestParameters: BotHealRotationTargetApiGetBotHealRotationTargetsBulkRequest, options?: any) {
        return BotHealRotationTargetApiFp(this.configuration).getBotHealRotationTargetsBulk(requestParameters.body, options).then((request) => request(this.axios, this.basePath));
    }
    public getBotHealRotationTargetsCount(requestParameters: BotHealRotationTargetApiGetBotHealRotationTargetsCountRequest = {}, options?: any) {
        return BotHealRotationTargetApiFp(this.configuration).getBotHealRotationTargetsCount(requestParameters.includes, requestParameters.where, requestParameters.whereOr, requestParameters.groupBy, requestParameters.limit, requestParameters.page, requestParameters.orderBy, requestParameters.orderDirection, requestParameters.select, options).then((request) => request(this.axios, this.basePath));
    }
    public listBotHealRotationTargets(requestParameters: BotHealRotationTargetApiListBotHealRotationTargetsRequest = {}, options?: any) {
        return BotHealRotationTargetApiFp(this.configuration).listBotHealRotationTargets(requestParameters.includes, requestParameters.where, requestParameters.whereOr, requestParameters.groupBy, requestParameters.limit, requestParameters.page, requestParameters.orderBy, requestParameters.orderDirection, requestParameters.select, options).then((request) => request(this.axios, this.basePath));
    }
    public updateBotHealRotationTarget(requestParameters: BotHealRotationTargetApiUpdateBotHealRotationTargetRequest, options?: any) {
        return BotHealRotationTargetApiFp(this.configuration).updateBotHealRotationTarget(requestParameters.id, requestParameters.botHealRotationTarget, options).then((request) => request(this.axios, this.basePath));
    }
}
