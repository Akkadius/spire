import globalAxios, { AxiosPromise, AxiosInstance } from 'axios';
import { Configuration } from '../configuration';
// Some imports not used depending on template conditions
// @ts-ignore
import { BASE_PATH, COLLECTION_FORMATS, RequestArgs, BaseAPI, RequiredError } from '../base';
// @ts-ignore
import { CrudcontrollersBulkFetchByIdsGetRequest } from '../models';
// @ts-ignore
import { ModelsInstanceListPlayer } from '../models';
export const InstanceListPlayerApiAxiosParamCreator = function (configuration?: Configuration) {
    return {
        createInstanceListPlayer: async (instanceListPlayer: ModelsInstanceListPlayer, options: any = {}): Promise<RequestArgs> => {
            if (instanceListPlayer === null || instanceListPlayer === undefined) {
                throw new RequiredError('instanceListPlayer','Required parameter instanceListPlayer was null or undefined when calling createInstanceListPlayer.');
            }
            const localVarPath = `/instance_list_player`;
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
            const nonString = typeof instanceListPlayer !== 'string';
            const needsSerialization = nonString && configuration && configuration.isJsonMime
                ? configuration.isJsonMime(localVarRequestOptions.headers['Content-Type'])
                : nonString;
            localVarRequestOptions.data =  needsSerialization
                ? JSON.stringify(instanceListPlayer !== undefined ? instanceListPlayer : {})
                : (instanceListPlayer || "");
            return {
                url: localVarUrlObj.pathname + localVarUrlObj.search + localVarUrlObj.hash,
                options: localVarRequestOptions,
            };
        },
        deleteInstanceListPlayer: async (id: number, options: any = {}): Promise<RequestArgs> => {
            if (id === null || id === undefined) {
                throw new RequiredError('id','Required parameter id was null or undefined when calling deleteInstanceListPlayer.');
            }
            const localVarPath = `/instance_list_player/{id}`
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
        getInstanceListPlayer: async (id: number, includes?: string, select?: string, options: any = {}): Promise<RequestArgs> => {
            if (id === null || id === undefined) {
                throw new RequiredError('id','Required parameter id was null or undefined when calling getInstanceListPlayer.');
            }
            const localVarPath = `/instance_list_player/{id}`
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
        getInstanceListPlayersBulk: async (body: CrudcontrollersBulkFetchByIdsGetRequest, options: any = {}): Promise<RequestArgs> => {
            if (body === null || body === undefined) {
                throw new RequiredError('body','Required parameter body was null or undefined when calling getInstanceListPlayersBulk.');
            }
            const localVarPath = `/instance_list_players/bulk`;
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
        listInstanceListPlayers: async (includes?: string, where?: string, whereOr?: string, groupBy?: string, limit?: string, page?: number, orderBy?: string, orderDirection?: string, select?: string, options: any = {}): Promise<RequestArgs> => {
            const localVarPath = `/instance_list_players`;
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
        updateInstanceListPlayer: async (id: number, instanceListPlayer: ModelsInstanceListPlayer, options: any = {}): Promise<RequestArgs> => {
            if (id === null || id === undefined) {
                throw new RequiredError('id','Required parameter id was null or undefined when calling updateInstanceListPlayer.');
            }
            if (instanceListPlayer === null || instanceListPlayer === undefined) {
                throw new RequiredError('instanceListPlayer','Required parameter instanceListPlayer was null or undefined when calling updateInstanceListPlayer.');
            }
            const localVarPath = `/instance_list_player/{id}`
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
            const nonString = typeof instanceListPlayer !== 'string';
            const needsSerialization = nonString && configuration && configuration.isJsonMime
                ? configuration.isJsonMime(localVarRequestOptions.headers['Content-Type'])
                : nonString;
            localVarRequestOptions.data =  needsSerialization
                ? JSON.stringify(instanceListPlayer !== undefined ? instanceListPlayer : {})
                : (instanceListPlayer || "");
            return {
                url: localVarUrlObj.pathname + localVarUrlObj.search + localVarUrlObj.hash,
                options: localVarRequestOptions,
            };
        },
    }
};
export const InstanceListPlayerApiFp = function(configuration?: Configuration) {
    return {
        async createInstanceListPlayer(instanceListPlayer: ModelsInstanceListPlayer, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsInstanceListPlayer>>> {
            const localVarAxiosArgs = await InstanceListPlayerApiAxiosParamCreator(configuration).createInstanceListPlayer(instanceListPlayer, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        async deleteInstanceListPlayer(id: number, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<string>> {
            const localVarAxiosArgs = await InstanceListPlayerApiAxiosParamCreator(configuration).deleteInstanceListPlayer(id, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        async getInstanceListPlayer(id: number, includes?: string, select?: string, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsInstanceListPlayer>>> {
            const localVarAxiosArgs = await InstanceListPlayerApiAxiosParamCreator(configuration).getInstanceListPlayer(id, includes, select, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        async getInstanceListPlayersBulk(body: CrudcontrollersBulkFetchByIdsGetRequest, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsInstanceListPlayer>>> {
            const localVarAxiosArgs = await InstanceListPlayerApiAxiosParamCreator(configuration).getInstanceListPlayersBulk(body, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        async listInstanceListPlayers(includes?: string, where?: string, whereOr?: string, groupBy?: string, limit?: string, page?: number, orderBy?: string, orderDirection?: string, select?: string, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsInstanceListPlayer>>> {
            const localVarAxiosArgs = await InstanceListPlayerApiAxiosParamCreator(configuration).listInstanceListPlayers(includes, where, whereOr, groupBy, limit, page, orderBy, orderDirection, select, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        async updateInstanceListPlayer(id: number, instanceListPlayer: ModelsInstanceListPlayer, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsInstanceListPlayer>>> {
            const localVarAxiosArgs = await InstanceListPlayerApiAxiosParamCreator(configuration).updateInstanceListPlayer(id, instanceListPlayer, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
    }
};
export const InstanceListPlayerApiFactory = function (configuration?: Configuration, basePath?: string, axios?: AxiosInstance) {
    return {
        createInstanceListPlayer(instanceListPlayer: ModelsInstanceListPlayer, options?: any): AxiosPromise<Array<ModelsInstanceListPlayer>> {
            return InstanceListPlayerApiFp(configuration).createInstanceListPlayer(instanceListPlayer, options).then((request) => request(axios, basePath));
        },
        deleteInstanceListPlayer(id: number, options?: any): AxiosPromise<string> {
            return InstanceListPlayerApiFp(configuration).deleteInstanceListPlayer(id, options).then((request) => request(axios, basePath));
        },
        getInstanceListPlayer(id: number, includes?: string, select?: string, options?: any): AxiosPromise<Array<ModelsInstanceListPlayer>> {
            return InstanceListPlayerApiFp(configuration).getInstanceListPlayer(id, includes, select, options).then((request) => request(axios, basePath));
        },
        getInstanceListPlayersBulk(body: CrudcontrollersBulkFetchByIdsGetRequest, options?: any): AxiosPromise<Array<ModelsInstanceListPlayer>> {
            return InstanceListPlayerApiFp(configuration).getInstanceListPlayersBulk(body, options).then((request) => request(axios, basePath));
        },
        listInstanceListPlayers(includes?: string, where?: string, whereOr?: string, groupBy?: string, limit?: string, page?: number, orderBy?: string, orderDirection?: string, select?: string, options?: any): AxiosPromise<Array<ModelsInstanceListPlayer>> {
            return InstanceListPlayerApiFp(configuration).listInstanceListPlayers(includes, where, whereOr, groupBy, limit, page, orderBy, orderDirection, select, options).then((request) => request(axios, basePath));
        },
        updateInstanceListPlayer(id: number, instanceListPlayer: ModelsInstanceListPlayer, options?: any): AxiosPromise<Array<ModelsInstanceListPlayer>> {
            return InstanceListPlayerApiFp(configuration).updateInstanceListPlayer(id, instanceListPlayer, options).then((request) => request(axios, basePath));
        },
    };
};
export interface InstanceListPlayerApiCreateInstanceListPlayerRequest {
    readonly instanceListPlayer: ModelsInstanceListPlayer
}
export interface InstanceListPlayerApiDeleteInstanceListPlayerRequest {
    readonly id: number
}
export interface InstanceListPlayerApiGetInstanceListPlayerRequest {
    readonly id: number
    readonly includes?: string
    readonly select?: string
}
export interface InstanceListPlayerApiGetInstanceListPlayersBulkRequest {
    readonly body: CrudcontrollersBulkFetchByIdsGetRequest
}
export interface InstanceListPlayerApiListInstanceListPlayersRequest {
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
export interface InstanceListPlayerApiUpdateInstanceListPlayerRequest {
    readonly id: number
    readonly instanceListPlayer: ModelsInstanceListPlayer
}
export class InstanceListPlayerApi extends BaseAPI {
    public createInstanceListPlayer(requestParameters: InstanceListPlayerApiCreateInstanceListPlayerRequest, options?: any) {
        return InstanceListPlayerApiFp(this.configuration).createInstanceListPlayer(requestParameters.instanceListPlayer, options).then((request) => request(this.axios, this.basePath));
    }
    public deleteInstanceListPlayer(requestParameters: InstanceListPlayerApiDeleteInstanceListPlayerRequest, options?: any) {
        return InstanceListPlayerApiFp(this.configuration).deleteInstanceListPlayer(requestParameters.id, options).then((request) => request(this.axios, this.basePath));
    }
    public getInstanceListPlayer(requestParameters: InstanceListPlayerApiGetInstanceListPlayerRequest, options?: any) {
        return InstanceListPlayerApiFp(this.configuration).getInstanceListPlayer(requestParameters.id, requestParameters.includes, requestParameters.select, options).then((request) => request(this.axios, this.basePath));
    }
    public getInstanceListPlayersBulk(requestParameters: InstanceListPlayerApiGetInstanceListPlayersBulkRequest, options?: any) {
        return InstanceListPlayerApiFp(this.configuration).getInstanceListPlayersBulk(requestParameters.body, options).then((request) => request(this.axios, this.basePath));
    }
    public listInstanceListPlayers(requestParameters: InstanceListPlayerApiListInstanceListPlayersRequest = {}, options?: any) {
        return InstanceListPlayerApiFp(this.configuration).listInstanceListPlayers(requestParameters.includes, requestParameters.where, requestParameters.whereOr, requestParameters.groupBy, requestParameters.limit, requestParameters.page, requestParameters.orderBy, requestParameters.orderDirection, requestParameters.select, options).then((request) => request(this.axios, this.basePath));
    }
    public updateInstanceListPlayer(requestParameters: InstanceListPlayerApiUpdateInstanceListPlayerRequest, options?: any) {
        return InstanceListPlayerApiFp(this.configuration).updateInstanceListPlayer(requestParameters.id, requestParameters.instanceListPlayer, options).then((request) => request(this.axios, this.basePath));
    }
}
