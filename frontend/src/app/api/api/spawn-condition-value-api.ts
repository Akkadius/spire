import globalAxios, { AxiosPromise, AxiosInstance } from 'axios';
import { Configuration } from '../configuration';
// Some imports not used depending on template conditions
// @ts-ignore
import { BASE_PATH, COLLECTION_FORMATS, RequestArgs, BaseAPI, RequiredError } from '../base';
// @ts-ignore
import { CrudcontrollersBulkFetchByIdsGetRequest } from '../models';
// @ts-ignore
import { ModelsSpawnConditionValue } from '../models';
export const SpawnConditionValueApiAxiosParamCreator = function (configuration?: Configuration) {
    return {
        createSpawnConditionValue: async (spawnConditionValue: ModelsSpawnConditionValue, options: any = {}): Promise<RequestArgs> => {
            if (spawnConditionValue === null || spawnConditionValue === undefined) {
                throw new RequiredError('spawnConditionValue','Required parameter spawnConditionValue was null or undefined when calling createSpawnConditionValue.');
            }
            const localVarPath = `/spawn_condition_value`;
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
            const nonString = typeof spawnConditionValue !== 'string';
            const needsSerialization = nonString && configuration && configuration.isJsonMime
                ? configuration.isJsonMime(localVarRequestOptions.headers['Content-Type'])
                : nonString;
            localVarRequestOptions.data =  needsSerialization
                ? JSON.stringify(spawnConditionValue !== undefined ? spawnConditionValue : {})
                : (spawnConditionValue || "");
            return {
                url: localVarUrlObj.pathname + localVarUrlObj.search + localVarUrlObj.hash,
                options: localVarRequestOptions,
            };
        },
        deleteSpawnConditionValue: async (id: number, options: any = {}): Promise<RequestArgs> => {
            if (id === null || id === undefined) {
                throw new RequiredError('id','Required parameter id was null or undefined when calling deleteSpawnConditionValue.');
            }
            const localVarPath = `/spawn_condition_value/{id}`
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
        getSpawnConditionValue: async (id: number, includes?: string, select?: string, options: any = {}): Promise<RequestArgs> => {
            if (id === null || id === undefined) {
                throw new RequiredError('id','Required parameter id was null or undefined when calling getSpawnConditionValue.');
            }
            const localVarPath = `/spawn_condition_value/{id}`
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
        getSpawnConditionValuesBulk: async (body: CrudcontrollersBulkFetchByIdsGetRequest, options: any = {}): Promise<RequestArgs> => {
            if (body === null || body === undefined) {
                throw new RequiredError('body','Required parameter body was null or undefined when calling getSpawnConditionValuesBulk.');
            }
            const localVarPath = `/spawn_condition_values/bulk`;
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
        listSpawnConditionValues: async (includes?: string, where?: string, whereOr?: string, groupBy?: string, limit?: string, page?: number, orderBy?: string, orderDirection?: string, select?: string, options: any = {}): Promise<RequestArgs> => {
            const localVarPath = `/spawn_condition_values`;
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
        updateSpawnConditionValue: async (id: number, spawnConditionValue: ModelsSpawnConditionValue, options: any = {}): Promise<RequestArgs> => {
            if (id === null || id === undefined) {
                throw new RequiredError('id','Required parameter id was null or undefined when calling updateSpawnConditionValue.');
            }
            if (spawnConditionValue === null || spawnConditionValue === undefined) {
                throw new RequiredError('spawnConditionValue','Required parameter spawnConditionValue was null or undefined when calling updateSpawnConditionValue.');
            }
            const localVarPath = `/spawn_condition_value/{id}`
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
            const nonString = typeof spawnConditionValue !== 'string';
            const needsSerialization = nonString && configuration && configuration.isJsonMime
                ? configuration.isJsonMime(localVarRequestOptions.headers['Content-Type'])
                : nonString;
            localVarRequestOptions.data =  needsSerialization
                ? JSON.stringify(spawnConditionValue !== undefined ? spawnConditionValue : {})
                : (spawnConditionValue || "");
            return {
                url: localVarUrlObj.pathname + localVarUrlObj.search + localVarUrlObj.hash,
                options: localVarRequestOptions,
            };
        },
    }
};
export const SpawnConditionValueApiFp = function(configuration?: Configuration) {
    return {
        async createSpawnConditionValue(spawnConditionValue: ModelsSpawnConditionValue, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsSpawnConditionValue>>> {
            const localVarAxiosArgs = await SpawnConditionValueApiAxiosParamCreator(configuration).createSpawnConditionValue(spawnConditionValue, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        async deleteSpawnConditionValue(id: number, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<string>> {
            const localVarAxiosArgs = await SpawnConditionValueApiAxiosParamCreator(configuration).deleteSpawnConditionValue(id, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        async getSpawnConditionValue(id: number, includes?: string, select?: string, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsSpawnConditionValue>>> {
            const localVarAxiosArgs = await SpawnConditionValueApiAxiosParamCreator(configuration).getSpawnConditionValue(id, includes, select, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        async getSpawnConditionValuesBulk(body: CrudcontrollersBulkFetchByIdsGetRequest, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsSpawnConditionValue>>> {
            const localVarAxiosArgs = await SpawnConditionValueApiAxiosParamCreator(configuration).getSpawnConditionValuesBulk(body, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        async listSpawnConditionValues(includes?: string, where?: string, whereOr?: string, groupBy?: string, limit?: string, page?: number, orderBy?: string, orderDirection?: string, select?: string, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsSpawnConditionValue>>> {
            const localVarAxiosArgs = await SpawnConditionValueApiAxiosParamCreator(configuration).listSpawnConditionValues(includes, where, whereOr, groupBy, limit, page, orderBy, orderDirection, select, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        async updateSpawnConditionValue(id: number, spawnConditionValue: ModelsSpawnConditionValue, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsSpawnConditionValue>>> {
            const localVarAxiosArgs = await SpawnConditionValueApiAxiosParamCreator(configuration).updateSpawnConditionValue(id, spawnConditionValue, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
    }
};
export const SpawnConditionValueApiFactory = function (configuration?: Configuration, basePath?: string, axios?: AxiosInstance) {
    return {
        createSpawnConditionValue(spawnConditionValue: ModelsSpawnConditionValue, options?: any): AxiosPromise<Array<ModelsSpawnConditionValue>> {
            return SpawnConditionValueApiFp(configuration).createSpawnConditionValue(spawnConditionValue, options).then((request) => request(axios, basePath));
        },
        deleteSpawnConditionValue(id: number, options?: any): AxiosPromise<string> {
            return SpawnConditionValueApiFp(configuration).deleteSpawnConditionValue(id, options).then((request) => request(axios, basePath));
        },
        getSpawnConditionValue(id: number, includes?: string, select?: string, options?: any): AxiosPromise<Array<ModelsSpawnConditionValue>> {
            return SpawnConditionValueApiFp(configuration).getSpawnConditionValue(id, includes, select, options).then((request) => request(axios, basePath));
        },
        getSpawnConditionValuesBulk(body: CrudcontrollersBulkFetchByIdsGetRequest, options?: any): AxiosPromise<Array<ModelsSpawnConditionValue>> {
            return SpawnConditionValueApiFp(configuration).getSpawnConditionValuesBulk(body, options).then((request) => request(axios, basePath));
        },
        listSpawnConditionValues(includes?: string, where?: string, whereOr?: string, groupBy?: string, limit?: string, page?: number, orderBy?: string, orderDirection?: string, select?: string, options?: any): AxiosPromise<Array<ModelsSpawnConditionValue>> {
            return SpawnConditionValueApiFp(configuration).listSpawnConditionValues(includes, where, whereOr, groupBy, limit, page, orderBy, orderDirection, select, options).then((request) => request(axios, basePath));
        },
        updateSpawnConditionValue(id: number, spawnConditionValue: ModelsSpawnConditionValue, options?: any): AxiosPromise<Array<ModelsSpawnConditionValue>> {
            return SpawnConditionValueApiFp(configuration).updateSpawnConditionValue(id, spawnConditionValue, options).then((request) => request(axios, basePath));
        },
    };
};
export interface SpawnConditionValueApiCreateSpawnConditionValueRequest {
    readonly spawnConditionValue: ModelsSpawnConditionValue
}
export interface SpawnConditionValueApiDeleteSpawnConditionValueRequest {
    readonly id: number
}
export interface SpawnConditionValueApiGetSpawnConditionValueRequest {
    readonly id: number
    readonly includes?: string
    readonly select?: string
}
export interface SpawnConditionValueApiGetSpawnConditionValuesBulkRequest {
    readonly body: CrudcontrollersBulkFetchByIdsGetRequest
}
export interface SpawnConditionValueApiListSpawnConditionValuesRequest {
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
export interface SpawnConditionValueApiUpdateSpawnConditionValueRequest {
    readonly id: number
    readonly spawnConditionValue: ModelsSpawnConditionValue
}
export class SpawnConditionValueApi extends BaseAPI {
    public createSpawnConditionValue(requestParameters: SpawnConditionValueApiCreateSpawnConditionValueRequest, options?: any) {
        return SpawnConditionValueApiFp(this.configuration).createSpawnConditionValue(requestParameters.spawnConditionValue, options).then((request) => request(this.axios, this.basePath));
    }
    public deleteSpawnConditionValue(requestParameters: SpawnConditionValueApiDeleteSpawnConditionValueRequest, options?: any) {
        return SpawnConditionValueApiFp(this.configuration).deleteSpawnConditionValue(requestParameters.id, options).then((request) => request(this.axios, this.basePath));
    }
    public getSpawnConditionValue(requestParameters: SpawnConditionValueApiGetSpawnConditionValueRequest, options?: any) {
        return SpawnConditionValueApiFp(this.configuration).getSpawnConditionValue(requestParameters.id, requestParameters.includes, requestParameters.select, options).then((request) => request(this.axios, this.basePath));
    }
    public getSpawnConditionValuesBulk(requestParameters: SpawnConditionValueApiGetSpawnConditionValuesBulkRequest, options?: any) {
        return SpawnConditionValueApiFp(this.configuration).getSpawnConditionValuesBulk(requestParameters.body, options).then((request) => request(this.axios, this.basePath));
    }
    public listSpawnConditionValues(requestParameters: SpawnConditionValueApiListSpawnConditionValuesRequest = {}, options?: any) {
        return SpawnConditionValueApiFp(this.configuration).listSpawnConditionValues(requestParameters.includes, requestParameters.where, requestParameters.whereOr, requestParameters.groupBy, requestParameters.limit, requestParameters.page, requestParameters.orderBy, requestParameters.orderDirection, requestParameters.select, options).then((request) => request(this.axios, this.basePath));
    }
    public updateSpawnConditionValue(requestParameters: SpawnConditionValueApiUpdateSpawnConditionValueRequest, options?: any) {
        return SpawnConditionValueApiFp(this.configuration).updateSpawnConditionValue(requestParameters.id, requestParameters.spawnConditionValue, options).then((request) => request(this.axios, this.basePath));
    }
}
