import globalAxios, { AxiosPromise, AxiosInstance } from 'axios';
import { Configuration } from '../configuration';
// Some imports not used depending on template conditions
// @ts-ignore
import { BASE_PATH, COLLECTION_FORMATS, RequestArgs, BaseAPI, RequiredError } from '../base';
// @ts-ignore
import { CrudcontrollersBulkFetchByIdsGetRequest } from '../models';
// @ts-ignore
import { ModelsPetsEquipmentsetEntry } from '../models';
export const PetsEquipmentsetEntryApiAxiosParamCreator = function (configuration?: Configuration) {
    return {
        createPetsEquipmentsetEntry: async (petsEquipmentsetEntry: ModelsPetsEquipmentsetEntry, options: any = {}): Promise<RequestArgs> => {
            if (petsEquipmentsetEntry === null || petsEquipmentsetEntry === undefined) {
                throw new RequiredError('petsEquipmentsetEntry','Required parameter petsEquipmentsetEntry was null or undefined when calling createPetsEquipmentsetEntry.');
            }
            const localVarPath = `/pets_equipmentset_entry`;
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
            const nonString = typeof petsEquipmentsetEntry !== 'string';
            const needsSerialization = nonString && configuration && configuration.isJsonMime
                ? configuration.isJsonMime(localVarRequestOptions.headers['Content-Type'])
                : nonString;
            localVarRequestOptions.data =  needsSerialization
                ? JSON.stringify(petsEquipmentsetEntry !== undefined ? petsEquipmentsetEntry : {})
                : (petsEquipmentsetEntry || "");
            return {
                url: localVarUrlObj.pathname + localVarUrlObj.search + localVarUrlObj.hash,
                options: localVarRequestOptions,
            };
        },
        deletePetsEquipmentsetEntry: async (id: number, options: any = {}): Promise<RequestArgs> => {
            if (id === null || id === undefined) {
                throw new RequiredError('id','Required parameter id was null or undefined when calling deletePetsEquipmentsetEntry.');
            }
            const localVarPath = `/pets_equipmentset_entry/{id}`
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
        getPetsEquipmentsetEntriesBulk: async (body: CrudcontrollersBulkFetchByIdsGetRequest, options: any = {}): Promise<RequestArgs> => {
            if (body === null || body === undefined) {
                throw new RequiredError('body','Required parameter body was null or undefined when calling getPetsEquipmentsetEntriesBulk.');
            }
            const localVarPath = `/pets_equipmentset_entries/bulk`;
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
        getPetsEquipmentsetEntriesCount: async (includes?: string, where?: string, whereOr?: string, groupBy?: string, limit?: string, page?: number, orderBy?: string, orderDirection?: string, select?: string, options: any = {}): Promise<RequestArgs> => {
            const localVarPath = `/pets_equipmentset_entries/count`;
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
        getPetsEquipmentsetEntry: async (id: number, includes?: string, select?: string, options: any = {}): Promise<RequestArgs> => {
            if (id === null || id === undefined) {
                throw new RequiredError('id','Required parameter id was null or undefined when calling getPetsEquipmentsetEntry.');
            }
            const localVarPath = `/pets_equipmentset_entry/{id}`
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
        listPetsEquipmentsetEntries: async (includes?: string, where?: string, whereOr?: string, groupBy?: string, limit?: string, page?: number, orderBy?: string, orderDirection?: string, select?: string, options: any = {}): Promise<RequestArgs> => {
            const localVarPath = `/pets_equipmentset_entries`;
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
        updatePetsEquipmentsetEntry: async (id: number, petsEquipmentsetEntry: ModelsPetsEquipmentsetEntry, options: any = {}): Promise<RequestArgs> => {
            if (id === null || id === undefined) {
                throw new RequiredError('id','Required parameter id was null or undefined when calling updatePetsEquipmentsetEntry.');
            }
            if (petsEquipmentsetEntry === null || petsEquipmentsetEntry === undefined) {
                throw new RequiredError('petsEquipmentsetEntry','Required parameter petsEquipmentsetEntry was null or undefined when calling updatePetsEquipmentsetEntry.');
            }
            const localVarPath = `/pets_equipmentset_entry/{id}`
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
            const nonString = typeof petsEquipmentsetEntry !== 'string';
            const needsSerialization = nonString && configuration && configuration.isJsonMime
                ? configuration.isJsonMime(localVarRequestOptions.headers['Content-Type'])
                : nonString;
            localVarRequestOptions.data =  needsSerialization
                ? JSON.stringify(petsEquipmentsetEntry !== undefined ? petsEquipmentsetEntry : {})
                : (petsEquipmentsetEntry || "");
            return {
                url: localVarUrlObj.pathname + localVarUrlObj.search + localVarUrlObj.hash,
                options: localVarRequestOptions,
            };
        },
    }
};
export const PetsEquipmentsetEntryApiFp = function(configuration?: Configuration) {
    return {
        async createPetsEquipmentsetEntry(petsEquipmentsetEntry: ModelsPetsEquipmentsetEntry, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsPetsEquipmentsetEntry>>> {
            const localVarAxiosArgs = await PetsEquipmentsetEntryApiAxiosParamCreator(configuration).createPetsEquipmentsetEntry(petsEquipmentsetEntry, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        async deletePetsEquipmentsetEntry(id: number, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<string>> {
            const localVarAxiosArgs = await PetsEquipmentsetEntryApiAxiosParamCreator(configuration).deletePetsEquipmentsetEntry(id, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        async getPetsEquipmentsetEntriesBulk(body: CrudcontrollersBulkFetchByIdsGetRequest, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsPetsEquipmentsetEntry>>> {
            const localVarAxiosArgs = await PetsEquipmentsetEntryApiAxiosParamCreator(configuration).getPetsEquipmentsetEntriesBulk(body, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        async getPetsEquipmentsetEntriesCount(includes?: string, where?: string, whereOr?: string, groupBy?: string, limit?: string, page?: number, orderBy?: string, orderDirection?: string, select?: string, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsPetsEquipmentsetEntry>>> {
            const localVarAxiosArgs = await PetsEquipmentsetEntryApiAxiosParamCreator(configuration).getPetsEquipmentsetEntriesCount(includes, where, whereOr, groupBy, limit, page, orderBy, orderDirection, select, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        async getPetsEquipmentsetEntry(id: number, includes?: string, select?: string, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsPetsEquipmentsetEntry>>> {
            const localVarAxiosArgs = await PetsEquipmentsetEntryApiAxiosParamCreator(configuration).getPetsEquipmentsetEntry(id, includes, select, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        async listPetsEquipmentsetEntries(includes?: string, where?: string, whereOr?: string, groupBy?: string, limit?: string, page?: number, orderBy?: string, orderDirection?: string, select?: string, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsPetsEquipmentsetEntry>>> {
            const localVarAxiosArgs = await PetsEquipmentsetEntryApiAxiosParamCreator(configuration).listPetsEquipmentsetEntries(includes, where, whereOr, groupBy, limit, page, orderBy, orderDirection, select, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        async updatePetsEquipmentsetEntry(id: number, petsEquipmentsetEntry: ModelsPetsEquipmentsetEntry, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsPetsEquipmentsetEntry>>> {
            const localVarAxiosArgs = await PetsEquipmentsetEntryApiAxiosParamCreator(configuration).updatePetsEquipmentsetEntry(id, petsEquipmentsetEntry, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
    }
};
export const PetsEquipmentsetEntryApiFactory = function (configuration?: Configuration, basePath?: string, axios?: AxiosInstance) {
    return {
        createPetsEquipmentsetEntry(petsEquipmentsetEntry: ModelsPetsEquipmentsetEntry, options?: any): AxiosPromise<Array<ModelsPetsEquipmentsetEntry>> {
            return PetsEquipmentsetEntryApiFp(configuration).createPetsEquipmentsetEntry(petsEquipmentsetEntry, options).then((request) => request(axios, basePath));
        },
        deletePetsEquipmentsetEntry(id: number, options?: any): AxiosPromise<string> {
            return PetsEquipmentsetEntryApiFp(configuration).deletePetsEquipmentsetEntry(id, options).then((request) => request(axios, basePath));
        },
        getPetsEquipmentsetEntriesBulk(body: CrudcontrollersBulkFetchByIdsGetRequest, options?: any): AxiosPromise<Array<ModelsPetsEquipmentsetEntry>> {
            return PetsEquipmentsetEntryApiFp(configuration).getPetsEquipmentsetEntriesBulk(body, options).then((request) => request(axios, basePath));
        },
        getPetsEquipmentsetEntriesCount(includes?: string, where?: string, whereOr?: string, groupBy?: string, limit?: string, page?: number, orderBy?: string, orderDirection?: string, select?: string, options?: any): AxiosPromise<Array<ModelsPetsEquipmentsetEntry>> {
            return PetsEquipmentsetEntryApiFp(configuration).getPetsEquipmentsetEntriesCount(includes, where, whereOr, groupBy, limit, page, orderBy, orderDirection, select, options).then((request) => request(axios, basePath));
        },
        getPetsEquipmentsetEntry(id: number, includes?: string, select?: string, options?: any): AxiosPromise<Array<ModelsPetsEquipmentsetEntry>> {
            return PetsEquipmentsetEntryApiFp(configuration).getPetsEquipmentsetEntry(id, includes, select, options).then((request) => request(axios, basePath));
        },
        listPetsEquipmentsetEntries(includes?: string, where?: string, whereOr?: string, groupBy?: string, limit?: string, page?: number, orderBy?: string, orderDirection?: string, select?: string, options?: any): AxiosPromise<Array<ModelsPetsEquipmentsetEntry>> {
            return PetsEquipmentsetEntryApiFp(configuration).listPetsEquipmentsetEntries(includes, where, whereOr, groupBy, limit, page, orderBy, orderDirection, select, options).then((request) => request(axios, basePath));
        },
        updatePetsEquipmentsetEntry(id: number, petsEquipmentsetEntry: ModelsPetsEquipmentsetEntry, options?: any): AxiosPromise<Array<ModelsPetsEquipmentsetEntry>> {
            return PetsEquipmentsetEntryApiFp(configuration).updatePetsEquipmentsetEntry(id, petsEquipmentsetEntry, options).then((request) => request(axios, basePath));
        },
    };
};
export interface PetsEquipmentsetEntryApiCreatePetsEquipmentsetEntryRequest {
    readonly petsEquipmentsetEntry: ModelsPetsEquipmentsetEntry
}
export interface PetsEquipmentsetEntryApiDeletePetsEquipmentsetEntryRequest {
    readonly id: number
}
export interface PetsEquipmentsetEntryApiGetPetsEquipmentsetEntriesBulkRequest {
    readonly body: CrudcontrollersBulkFetchByIdsGetRequest
}
export interface PetsEquipmentsetEntryApiGetPetsEquipmentsetEntriesCountRequest {
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
export interface PetsEquipmentsetEntryApiGetPetsEquipmentsetEntryRequest {
    readonly id: number
    readonly includes?: string
    readonly select?: string
}
export interface PetsEquipmentsetEntryApiListPetsEquipmentsetEntriesRequest {
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
export interface PetsEquipmentsetEntryApiUpdatePetsEquipmentsetEntryRequest {
    readonly id: number
    readonly petsEquipmentsetEntry: ModelsPetsEquipmentsetEntry
}
export class PetsEquipmentsetEntryApi extends BaseAPI {
    public createPetsEquipmentsetEntry(requestParameters: PetsEquipmentsetEntryApiCreatePetsEquipmentsetEntryRequest, options?: any) {
        return PetsEquipmentsetEntryApiFp(this.configuration).createPetsEquipmentsetEntry(requestParameters.petsEquipmentsetEntry, options).then((request) => request(this.axios, this.basePath));
    }
    public deletePetsEquipmentsetEntry(requestParameters: PetsEquipmentsetEntryApiDeletePetsEquipmentsetEntryRequest, options?: any) {
        return PetsEquipmentsetEntryApiFp(this.configuration).deletePetsEquipmentsetEntry(requestParameters.id, options).then((request) => request(this.axios, this.basePath));
    }
    public getPetsEquipmentsetEntriesBulk(requestParameters: PetsEquipmentsetEntryApiGetPetsEquipmentsetEntriesBulkRequest, options?: any) {
        return PetsEquipmentsetEntryApiFp(this.configuration).getPetsEquipmentsetEntriesBulk(requestParameters.body, options).then((request) => request(this.axios, this.basePath));
    }
    public getPetsEquipmentsetEntriesCount(requestParameters: PetsEquipmentsetEntryApiGetPetsEquipmentsetEntriesCountRequest = {}, options?: any) {
        return PetsEquipmentsetEntryApiFp(this.configuration).getPetsEquipmentsetEntriesCount(requestParameters.includes, requestParameters.where, requestParameters.whereOr, requestParameters.groupBy, requestParameters.limit, requestParameters.page, requestParameters.orderBy, requestParameters.orderDirection, requestParameters.select, options).then((request) => request(this.axios, this.basePath));
    }
    public getPetsEquipmentsetEntry(requestParameters: PetsEquipmentsetEntryApiGetPetsEquipmentsetEntryRequest, options?: any) {
        return PetsEquipmentsetEntryApiFp(this.configuration).getPetsEquipmentsetEntry(requestParameters.id, requestParameters.includes, requestParameters.select, options).then((request) => request(this.axios, this.basePath));
    }
    public listPetsEquipmentsetEntries(requestParameters: PetsEquipmentsetEntryApiListPetsEquipmentsetEntriesRequest = {}, options?: any) {
        return PetsEquipmentsetEntryApiFp(this.configuration).listPetsEquipmentsetEntries(requestParameters.includes, requestParameters.where, requestParameters.whereOr, requestParameters.groupBy, requestParameters.limit, requestParameters.page, requestParameters.orderBy, requestParameters.orderDirection, requestParameters.select, options).then((request) => request(this.axios, this.basePath));
    }
    public updatePetsEquipmentsetEntry(requestParameters: PetsEquipmentsetEntryApiUpdatePetsEquipmentsetEntryRequest, options?: any) {
        return PetsEquipmentsetEntryApiFp(this.configuration).updatePetsEquipmentsetEntry(requestParameters.id, requestParameters.petsEquipmentsetEntry, options).then((request) => request(this.axios, this.basePath));
    }
}
