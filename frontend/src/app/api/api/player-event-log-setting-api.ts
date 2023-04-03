import globalAxios, { AxiosPromise, AxiosInstance } from 'axios';
import { Configuration } from '../configuration';
// Some imports not used depending on template conditions
// @ts-ignore
import { BASE_PATH, COLLECTION_FORMATS, RequestArgs, BaseAPI, RequiredError } from '../base';
// @ts-ignore
import { CrudcontrollersBulkFetchByIdsGetRequest } from '../models';
// @ts-ignore
import { ModelsPlayerEventLogSetting } from '../models';
export const PlayerEventLogSettingApiAxiosParamCreator = function (configuration?: Configuration) {
    return {
        createPlayerEventLogSetting: async (playerEventLogSetting: ModelsPlayerEventLogSetting, options: any = {}): Promise<RequestArgs> => {
            if (playerEventLogSetting === null || playerEventLogSetting === undefined) {
                throw new RequiredError('playerEventLogSetting','Required parameter playerEventLogSetting was null or undefined when calling createPlayerEventLogSetting.');
            }
            const localVarPath = `/player_event_log_setting`;
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
            const nonString = typeof playerEventLogSetting !== 'string';
            const needsSerialization = nonString && configuration && configuration.isJsonMime
                ? configuration.isJsonMime(localVarRequestOptions.headers['Content-Type'])
                : nonString;
            localVarRequestOptions.data =  needsSerialization
                ? JSON.stringify(playerEventLogSetting !== undefined ? playerEventLogSetting : {})
                : (playerEventLogSetting || "");
            return {
                url: localVarUrlObj.pathname + localVarUrlObj.search + localVarUrlObj.hash,
                options: localVarRequestOptions,
            };
        },
        deletePlayerEventLogSetting: async (id: number, options: any = {}): Promise<RequestArgs> => {
            if (id === null || id === undefined) {
                throw new RequiredError('id','Required parameter id was null or undefined when calling deletePlayerEventLogSetting.');
            }
            const localVarPath = `/player_event_log_setting/{id}`
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
        getPlayerEventLogSetting: async (id: number, includes?: string, select?: string, options: any = {}): Promise<RequestArgs> => {
            if (id === null || id === undefined) {
                throw new RequiredError('id','Required parameter id was null or undefined when calling getPlayerEventLogSetting.');
            }
            const localVarPath = `/player_event_log_setting/{id}`
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
        getPlayerEventLogSettingsBulk: async (body: CrudcontrollersBulkFetchByIdsGetRequest, options: any = {}): Promise<RequestArgs> => {
            if (body === null || body === undefined) {
                throw new RequiredError('body','Required parameter body was null or undefined when calling getPlayerEventLogSettingsBulk.');
            }
            const localVarPath = `/player_event_log_settings/bulk`;
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
        getPlayerEventLogSettingsCount: async (includes?: string, where?: string, whereOr?: string, groupBy?: string, limit?: string, page?: number, orderBy?: string, orderDirection?: string, select?: string, options: any = {}): Promise<RequestArgs> => {
            const localVarPath = `/player_event_log_settings/count`;
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
        listPlayerEventLogSettings: async (includes?: string, where?: string, whereOr?: string, groupBy?: string, limit?: string, page?: number, orderBy?: string, orderDirection?: string, select?: string, options: any = {}): Promise<RequestArgs> => {
            const localVarPath = `/player_event_log_settings`;
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
        updatePlayerEventLogSetting: async (id: number, playerEventLogSetting: ModelsPlayerEventLogSetting, options: any = {}): Promise<RequestArgs> => {
            if (id === null || id === undefined) {
                throw new RequiredError('id','Required parameter id was null or undefined when calling updatePlayerEventLogSetting.');
            }
            if (playerEventLogSetting === null || playerEventLogSetting === undefined) {
                throw new RequiredError('playerEventLogSetting','Required parameter playerEventLogSetting was null or undefined when calling updatePlayerEventLogSetting.');
            }
            const localVarPath = `/player_event_log_setting/{id}`
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
            const nonString = typeof playerEventLogSetting !== 'string';
            const needsSerialization = nonString && configuration && configuration.isJsonMime
                ? configuration.isJsonMime(localVarRequestOptions.headers['Content-Type'])
                : nonString;
            localVarRequestOptions.data =  needsSerialization
                ? JSON.stringify(playerEventLogSetting !== undefined ? playerEventLogSetting : {})
                : (playerEventLogSetting || "");
            return {
                url: localVarUrlObj.pathname + localVarUrlObj.search + localVarUrlObj.hash,
                options: localVarRequestOptions,
            };
        },
    }
};
export const PlayerEventLogSettingApiFp = function(configuration?: Configuration) {
    return {
        async createPlayerEventLogSetting(playerEventLogSetting: ModelsPlayerEventLogSetting, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsPlayerEventLogSetting>>> {
            const localVarAxiosArgs = await PlayerEventLogSettingApiAxiosParamCreator(configuration).createPlayerEventLogSetting(playerEventLogSetting, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        async deletePlayerEventLogSetting(id: number, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<string>> {
            const localVarAxiosArgs = await PlayerEventLogSettingApiAxiosParamCreator(configuration).deletePlayerEventLogSetting(id, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        async getPlayerEventLogSetting(id: number, includes?: string, select?: string, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsPlayerEventLogSetting>>> {
            const localVarAxiosArgs = await PlayerEventLogSettingApiAxiosParamCreator(configuration).getPlayerEventLogSetting(id, includes, select, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        async getPlayerEventLogSettingsBulk(body: CrudcontrollersBulkFetchByIdsGetRequest, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsPlayerEventLogSetting>>> {
            const localVarAxiosArgs = await PlayerEventLogSettingApiAxiosParamCreator(configuration).getPlayerEventLogSettingsBulk(body, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        async getPlayerEventLogSettingsCount(includes?: string, where?: string, whereOr?: string, groupBy?: string, limit?: string, page?: number, orderBy?: string, orderDirection?: string, select?: string, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsPlayerEventLogSetting>>> {
            const localVarAxiosArgs = await PlayerEventLogSettingApiAxiosParamCreator(configuration).getPlayerEventLogSettingsCount(includes, where, whereOr, groupBy, limit, page, orderBy, orderDirection, select, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        async listPlayerEventLogSettings(includes?: string, where?: string, whereOr?: string, groupBy?: string, limit?: string, page?: number, orderBy?: string, orderDirection?: string, select?: string, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsPlayerEventLogSetting>>> {
            const localVarAxiosArgs = await PlayerEventLogSettingApiAxiosParamCreator(configuration).listPlayerEventLogSettings(includes, where, whereOr, groupBy, limit, page, orderBy, orderDirection, select, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        async updatePlayerEventLogSetting(id: number, playerEventLogSetting: ModelsPlayerEventLogSetting, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsPlayerEventLogSetting>>> {
            const localVarAxiosArgs = await PlayerEventLogSettingApiAxiosParamCreator(configuration).updatePlayerEventLogSetting(id, playerEventLogSetting, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
    }
};
export const PlayerEventLogSettingApiFactory = function (configuration?: Configuration, basePath?: string, axios?: AxiosInstance) {
    return {
        createPlayerEventLogSetting(playerEventLogSetting: ModelsPlayerEventLogSetting, options?: any): AxiosPromise<Array<ModelsPlayerEventLogSetting>> {
            return PlayerEventLogSettingApiFp(configuration).createPlayerEventLogSetting(playerEventLogSetting, options).then((request) => request(axios, basePath));
        },
        deletePlayerEventLogSetting(id: number, options?: any): AxiosPromise<string> {
            return PlayerEventLogSettingApiFp(configuration).deletePlayerEventLogSetting(id, options).then((request) => request(axios, basePath));
        },
        getPlayerEventLogSetting(id: number, includes?: string, select?: string, options?: any): AxiosPromise<Array<ModelsPlayerEventLogSetting>> {
            return PlayerEventLogSettingApiFp(configuration).getPlayerEventLogSetting(id, includes, select, options).then((request) => request(axios, basePath));
        },
        getPlayerEventLogSettingsBulk(body: CrudcontrollersBulkFetchByIdsGetRequest, options?: any): AxiosPromise<Array<ModelsPlayerEventLogSetting>> {
            return PlayerEventLogSettingApiFp(configuration).getPlayerEventLogSettingsBulk(body, options).then((request) => request(axios, basePath));
        },
        getPlayerEventLogSettingsCount(includes?: string, where?: string, whereOr?: string, groupBy?: string, limit?: string, page?: number, orderBy?: string, orderDirection?: string, select?: string, options?: any): AxiosPromise<Array<ModelsPlayerEventLogSetting>> {
            return PlayerEventLogSettingApiFp(configuration).getPlayerEventLogSettingsCount(includes, where, whereOr, groupBy, limit, page, orderBy, orderDirection, select, options).then((request) => request(axios, basePath));
        },
        listPlayerEventLogSettings(includes?: string, where?: string, whereOr?: string, groupBy?: string, limit?: string, page?: number, orderBy?: string, orderDirection?: string, select?: string, options?: any): AxiosPromise<Array<ModelsPlayerEventLogSetting>> {
            return PlayerEventLogSettingApiFp(configuration).listPlayerEventLogSettings(includes, where, whereOr, groupBy, limit, page, orderBy, orderDirection, select, options).then((request) => request(axios, basePath));
        },
        updatePlayerEventLogSetting(id: number, playerEventLogSetting: ModelsPlayerEventLogSetting, options?: any): AxiosPromise<Array<ModelsPlayerEventLogSetting>> {
            return PlayerEventLogSettingApiFp(configuration).updatePlayerEventLogSetting(id, playerEventLogSetting, options).then((request) => request(axios, basePath));
        },
    };
};
export interface PlayerEventLogSettingApiCreatePlayerEventLogSettingRequest {
    readonly playerEventLogSetting: ModelsPlayerEventLogSetting
}
export interface PlayerEventLogSettingApiDeletePlayerEventLogSettingRequest {
    readonly id: number
}
export interface PlayerEventLogSettingApiGetPlayerEventLogSettingRequest {
    readonly id: number
    readonly includes?: string
    readonly select?: string
}
export interface PlayerEventLogSettingApiGetPlayerEventLogSettingsBulkRequest {
    readonly body: CrudcontrollersBulkFetchByIdsGetRequest
}
export interface PlayerEventLogSettingApiGetPlayerEventLogSettingsCountRequest {
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
export interface PlayerEventLogSettingApiListPlayerEventLogSettingsRequest {
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
export interface PlayerEventLogSettingApiUpdatePlayerEventLogSettingRequest {
    readonly id: number
    readonly playerEventLogSetting: ModelsPlayerEventLogSetting
}
export class PlayerEventLogSettingApi extends BaseAPI {
    public createPlayerEventLogSetting(requestParameters: PlayerEventLogSettingApiCreatePlayerEventLogSettingRequest, options?: any) {
        return PlayerEventLogSettingApiFp(this.configuration).createPlayerEventLogSetting(requestParameters.playerEventLogSetting, options).then((request) => request(this.axios, this.basePath));
    }
    public deletePlayerEventLogSetting(requestParameters: PlayerEventLogSettingApiDeletePlayerEventLogSettingRequest, options?: any) {
        return PlayerEventLogSettingApiFp(this.configuration).deletePlayerEventLogSetting(requestParameters.id, options).then((request) => request(this.axios, this.basePath));
    }
    public getPlayerEventLogSetting(requestParameters: PlayerEventLogSettingApiGetPlayerEventLogSettingRequest, options?: any) {
        return PlayerEventLogSettingApiFp(this.configuration).getPlayerEventLogSetting(requestParameters.id, requestParameters.includes, requestParameters.select, options).then((request) => request(this.axios, this.basePath));
    }
    public getPlayerEventLogSettingsBulk(requestParameters: PlayerEventLogSettingApiGetPlayerEventLogSettingsBulkRequest, options?: any) {
        return PlayerEventLogSettingApiFp(this.configuration).getPlayerEventLogSettingsBulk(requestParameters.body, options).then((request) => request(this.axios, this.basePath));
    }
    public getPlayerEventLogSettingsCount(requestParameters: PlayerEventLogSettingApiGetPlayerEventLogSettingsCountRequest = {}, options?: any) {
        return PlayerEventLogSettingApiFp(this.configuration).getPlayerEventLogSettingsCount(requestParameters.includes, requestParameters.where, requestParameters.whereOr, requestParameters.groupBy, requestParameters.limit, requestParameters.page, requestParameters.orderBy, requestParameters.orderDirection, requestParameters.select, options).then((request) => request(this.axios, this.basePath));
    }
    public listPlayerEventLogSettings(requestParameters: PlayerEventLogSettingApiListPlayerEventLogSettingsRequest = {}, options?: any) {
        return PlayerEventLogSettingApiFp(this.configuration).listPlayerEventLogSettings(requestParameters.includes, requestParameters.where, requestParameters.whereOr, requestParameters.groupBy, requestParameters.limit, requestParameters.page, requestParameters.orderBy, requestParameters.orderDirection, requestParameters.select, options).then((request) => request(this.axios, this.basePath));
    }
    public updatePlayerEventLogSetting(requestParameters: PlayerEventLogSettingApiUpdatePlayerEventLogSettingRequest, options?: any) {
        return PlayerEventLogSettingApiFp(this.configuration).updatePlayerEventLogSetting(requestParameters.id, requestParameters.playerEventLogSetting, options).then((request) => request(this.axios, this.basePath));
    }
}
