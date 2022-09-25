import globalAxios, { AxiosPromise, AxiosInstance } from 'axios';
import { Configuration } from '../configuration';
// Some imports not used depending on template conditions
// @ts-ignore
import { BASE_PATH, COLLECTION_FORMATS, RequestArgs, BaseAPI, RequiredError } from '../base';
// @ts-ignore
import { CrudcontrollersBulkFetchByIdsGetRequest } from '../models';
// @ts-ignore
import { ModelsInventorySnapshot } from '../models';
export const InventorySnapshotApiAxiosParamCreator = function (configuration?: Configuration) {
    return {
        createInventorySnapshot: async (inventorySnapshot: ModelsInventorySnapshot, options: any = {}): Promise<RequestArgs> => {
            if (inventorySnapshot === null || inventorySnapshot === undefined) {
                throw new RequiredError('inventorySnapshot','Required parameter inventorySnapshot was null or undefined when calling createInventorySnapshot.');
            }
            const localVarPath = `/inventory_snapshot`;
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
            const nonString = typeof inventorySnapshot !== 'string';
            const needsSerialization = nonString && configuration && configuration.isJsonMime
                ? configuration.isJsonMime(localVarRequestOptions.headers['Content-Type'])
                : nonString;
            localVarRequestOptions.data =  needsSerialization
                ? JSON.stringify(inventorySnapshot !== undefined ? inventorySnapshot : {})
                : (inventorySnapshot || "");
            return {
                url: localVarUrlObj.pathname + localVarUrlObj.search + localVarUrlObj.hash,
                options: localVarRequestOptions,
            };
        },
        deleteInventorySnapshot: async (id: number, options: any = {}): Promise<RequestArgs> => {
            if (id === null || id === undefined) {
                throw new RequiredError('id','Required parameter id was null or undefined when calling deleteInventorySnapshot.');
            }
            const localVarPath = `/inventory_snapshot/{id}`
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
        getInventorySnapshot: async (id: number, includes?: string, select?: string, options: any = {}): Promise<RequestArgs> => {
            if (id === null || id === undefined) {
                throw new RequiredError('id','Required parameter id was null or undefined when calling getInventorySnapshot.');
            }
            const localVarPath = `/inventory_snapshot/{id}`
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
        getInventorySnapshotsBulk: async (body: CrudcontrollersBulkFetchByIdsGetRequest, options: any = {}): Promise<RequestArgs> => {
            if (body === null || body === undefined) {
                throw new RequiredError('body','Required parameter body was null or undefined when calling getInventorySnapshotsBulk.');
            }
            const localVarPath = `/inventory_snapshots/bulk`;
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
        listInventorySnapshots: async (includes?: string, where?: string, whereOr?: string, groupBy?: string, limit?: string, page?: number, orderBy?: string, orderDirection?: string, select?: string, options: any = {}): Promise<RequestArgs> => {
            const localVarPath = `/inventory_snapshots`;
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
        updateInventorySnapshot: async (id: number, inventorySnapshot: ModelsInventorySnapshot, options: any = {}): Promise<RequestArgs> => {
            if (id === null || id === undefined) {
                throw new RequiredError('id','Required parameter id was null or undefined when calling updateInventorySnapshot.');
            }
            if (inventorySnapshot === null || inventorySnapshot === undefined) {
                throw new RequiredError('inventorySnapshot','Required parameter inventorySnapshot was null or undefined when calling updateInventorySnapshot.');
            }
            const localVarPath = `/inventory_snapshot/{id}`
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
            const nonString = typeof inventorySnapshot !== 'string';
            const needsSerialization = nonString && configuration && configuration.isJsonMime
                ? configuration.isJsonMime(localVarRequestOptions.headers['Content-Type'])
                : nonString;
            localVarRequestOptions.data =  needsSerialization
                ? JSON.stringify(inventorySnapshot !== undefined ? inventorySnapshot : {})
                : (inventorySnapshot || "");
            return {
                url: localVarUrlObj.pathname + localVarUrlObj.search + localVarUrlObj.hash,
                options: localVarRequestOptions,
            };
        },
    }
};
export const InventorySnapshotApiFp = function(configuration?: Configuration) {
    return {
        async createInventorySnapshot(inventorySnapshot: ModelsInventorySnapshot, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsInventorySnapshot>>> {
            const localVarAxiosArgs = await InventorySnapshotApiAxiosParamCreator(configuration).createInventorySnapshot(inventorySnapshot, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        async deleteInventorySnapshot(id: number, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<string>> {
            const localVarAxiosArgs = await InventorySnapshotApiAxiosParamCreator(configuration).deleteInventorySnapshot(id, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        async getInventorySnapshot(id: number, includes?: string, select?: string, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsInventorySnapshot>>> {
            const localVarAxiosArgs = await InventorySnapshotApiAxiosParamCreator(configuration).getInventorySnapshot(id, includes, select, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        async getInventorySnapshotsBulk(body: CrudcontrollersBulkFetchByIdsGetRequest, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsInventorySnapshot>>> {
            const localVarAxiosArgs = await InventorySnapshotApiAxiosParamCreator(configuration).getInventorySnapshotsBulk(body, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        async listInventorySnapshots(includes?: string, where?: string, whereOr?: string, groupBy?: string, limit?: string, page?: number, orderBy?: string, orderDirection?: string, select?: string, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsInventorySnapshot>>> {
            const localVarAxiosArgs = await InventorySnapshotApiAxiosParamCreator(configuration).listInventorySnapshots(includes, where, whereOr, groupBy, limit, page, orderBy, orderDirection, select, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        async updateInventorySnapshot(id: number, inventorySnapshot: ModelsInventorySnapshot, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsInventorySnapshot>>> {
            const localVarAxiosArgs = await InventorySnapshotApiAxiosParamCreator(configuration).updateInventorySnapshot(id, inventorySnapshot, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
    }
};
export const InventorySnapshotApiFactory = function (configuration?: Configuration, basePath?: string, axios?: AxiosInstance) {
    return {
        createInventorySnapshot(inventorySnapshot: ModelsInventorySnapshot, options?: any): AxiosPromise<Array<ModelsInventorySnapshot>> {
            return InventorySnapshotApiFp(configuration).createInventorySnapshot(inventorySnapshot, options).then((request) => request(axios, basePath));
        },
        deleteInventorySnapshot(id: number, options?: any): AxiosPromise<string> {
            return InventorySnapshotApiFp(configuration).deleteInventorySnapshot(id, options).then((request) => request(axios, basePath));
        },
        getInventorySnapshot(id: number, includes?: string, select?: string, options?: any): AxiosPromise<Array<ModelsInventorySnapshot>> {
            return InventorySnapshotApiFp(configuration).getInventorySnapshot(id, includes, select, options).then((request) => request(axios, basePath));
        },
        getInventorySnapshotsBulk(body: CrudcontrollersBulkFetchByIdsGetRequest, options?: any): AxiosPromise<Array<ModelsInventorySnapshot>> {
            return InventorySnapshotApiFp(configuration).getInventorySnapshotsBulk(body, options).then((request) => request(axios, basePath));
        },
        listInventorySnapshots(includes?: string, where?: string, whereOr?: string, groupBy?: string, limit?: string, page?: number, orderBy?: string, orderDirection?: string, select?: string, options?: any): AxiosPromise<Array<ModelsInventorySnapshot>> {
            return InventorySnapshotApiFp(configuration).listInventorySnapshots(includes, where, whereOr, groupBy, limit, page, orderBy, orderDirection, select, options).then((request) => request(axios, basePath));
        },
        updateInventorySnapshot(id: number, inventorySnapshot: ModelsInventorySnapshot, options?: any): AxiosPromise<Array<ModelsInventorySnapshot>> {
            return InventorySnapshotApiFp(configuration).updateInventorySnapshot(id, inventorySnapshot, options).then((request) => request(axios, basePath));
        },
    };
};
export interface InventorySnapshotApiCreateInventorySnapshotRequest {
    readonly inventorySnapshot: ModelsInventorySnapshot
}
export interface InventorySnapshotApiDeleteInventorySnapshotRequest {
    readonly id: number
}
export interface InventorySnapshotApiGetInventorySnapshotRequest {
    readonly id: number
    readonly includes?: string
    readonly select?: string
}
export interface InventorySnapshotApiGetInventorySnapshotsBulkRequest {
    readonly body: CrudcontrollersBulkFetchByIdsGetRequest
}
export interface InventorySnapshotApiListInventorySnapshotsRequest {
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
export interface InventorySnapshotApiUpdateInventorySnapshotRequest {
    readonly id: number
    readonly inventorySnapshot: ModelsInventorySnapshot
}
export class InventorySnapshotApi extends BaseAPI {
    public createInventorySnapshot(requestParameters: InventorySnapshotApiCreateInventorySnapshotRequest, options?: any) {
        return InventorySnapshotApiFp(this.configuration).createInventorySnapshot(requestParameters.inventorySnapshot, options).then((request) => request(this.axios, this.basePath));
    }
    public deleteInventorySnapshot(requestParameters: InventorySnapshotApiDeleteInventorySnapshotRequest, options?: any) {
        return InventorySnapshotApiFp(this.configuration).deleteInventorySnapshot(requestParameters.id, options).then((request) => request(this.axios, this.basePath));
    }
    public getInventorySnapshot(requestParameters: InventorySnapshotApiGetInventorySnapshotRequest, options?: any) {
        return InventorySnapshotApiFp(this.configuration).getInventorySnapshot(requestParameters.id, requestParameters.includes, requestParameters.select, options).then((request) => request(this.axios, this.basePath));
    }
    public getInventorySnapshotsBulk(requestParameters: InventorySnapshotApiGetInventorySnapshotsBulkRequest, options?: any) {
        return InventorySnapshotApiFp(this.configuration).getInventorySnapshotsBulk(requestParameters.body, options).then((request) => request(this.axios, this.basePath));
    }
    public listInventorySnapshots(requestParameters: InventorySnapshotApiListInventorySnapshotsRequest = {}, options?: any) {
        return InventorySnapshotApiFp(this.configuration).listInventorySnapshots(requestParameters.includes, requestParameters.where, requestParameters.whereOr, requestParameters.groupBy, requestParameters.limit, requestParameters.page, requestParameters.orderBy, requestParameters.orderDirection, requestParameters.select, options).then((request) => request(this.axios, this.basePath));
    }
    public updateInventorySnapshot(requestParameters: InventorySnapshotApiUpdateInventorySnapshotRequest, options?: any) {
        return InventorySnapshotApiFp(this.configuration).updateInventorySnapshot(requestParameters.id, requestParameters.inventorySnapshot, options).then((request) => request(this.axios, this.basePath));
    }
}
