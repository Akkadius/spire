import globalAxios, { AxiosPromise, AxiosInstance } from 'axios';
import { Configuration } from '../configuration';
// Some imports not used depending on template conditions
// @ts-ignore
import { BASE_PATH, COLLECTION_FORMATS, RequestArgs, BaseAPI, RequiredError } from '../base';
// @ts-ignore
import { CrudcontrollersBulkFetchByIdsGetRequest } from '../models';
// @ts-ignore
import { ModelsAdventureTemplateEntryFlavor } from '../models';
export const AdventureTemplateEntryFlavorApiAxiosParamCreator = function (configuration?: Configuration) {
    return {
        createAdventureTemplateEntryFlavor: async (adventureTemplateEntryFlavor: ModelsAdventureTemplateEntryFlavor, options: any = {}): Promise<RequestArgs> => {
            if (adventureTemplateEntryFlavor === null || adventureTemplateEntryFlavor === undefined) {
                throw new RequiredError('adventureTemplateEntryFlavor','Required parameter adventureTemplateEntryFlavor was null or undefined when calling createAdventureTemplateEntryFlavor.');
            }
            const localVarPath = `/adventure_template_entry_flavor`;
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
            const nonString = typeof adventureTemplateEntryFlavor !== 'string';
            const needsSerialization = nonString && configuration && configuration.isJsonMime
                ? configuration.isJsonMime(localVarRequestOptions.headers['Content-Type'])
                : nonString;
            localVarRequestOptions.data =  needsSerialization
                ? JSON.stringify(adventureTemplateEntryFlavor !== undefined ? adventureTemplateEntryFlavor : {})
                : (adventureTemplateEntryFlavor || "");
            return {
                url: localVarUrlObj.pathname + localVarUrlObj.search + localVarUrlObj.hash,
                options: localVarRequestOptions,
            };
        },
        deleteAdventureTemplateEntryFlavor: async (id: number, options: any = {}): Promise<RequestArgs> => {
            if (id === null || id === undefined) {
                throw new RequiredError('id','Required parameter id was null or undefined when calling deleteAdventureTemplateEntryFlavor.');
            }
            const localVarPath = `/adventure_template_entry_flavor/{id}`
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
        getAdventureTemplateEntryFlavor: async (id: number, includes?: string, select?: string, options: any = {}): Promise<RequestArgs> => {
            if (id === null || id === undefined) {
                throw new RequiredError('id','Required parameter id was null or undefined when calling getAdventureTemplateEntryFlavor.');
            }
            const localVarPath = `/adventure_template_entry_flavor/{id}`
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
        getAdventureTemplateEntryFlavorsBulk: async (body: CrudcontrollersBulkFetchByIdsGetRequest, options: any = {}): Promise<RequestArgs> => {
            if (body === null || body === undefined) {
                throw new RequiredError('body','Required parameter body was null or undefined when calling getAdventureTemplateEntryFlavorsBulk.');
            }
            const localVarPath = `/adventure_template_entry_flavors/bulk`;
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
        getAdventureTemplateEntryFlavorsCount: async (includes?: string, where?: string, whereOr?: string, groupBy?: string, limit?: string, page?: number, orderBy?: string, orderDirection?: string, select?: string, options: any = {}): Promise<RequestArgs> => {
            const localVarPath = `/adventure_template_entry_flavors/count`;
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
        listAdventureTemplateEntryFlavors: async (includes?: string, where?: string, whereOr?: string, groupBy?: string, limit?: string, page?: number, orderBy?: string, orderDirection?: string, select?: string, options: any = {}): Promise<RequestArgs> => {
            const localVarPath = `/adventure_template_entry_flavors`;
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
        updateAdventureTemplateEntryFlavor: async (id: number, adventureTemplateEntryFlavor: ModelsAdventureTemplateEntryFlavor, options: any = {}): Promise<RequestArgs> => {
            if (id === null || id === undefined) {
                throw new RequiredError('id','Required parameter id was null or undefined when calling updateAdventureTemplateEntryFlavor.');
            }
            if (adventureTemplateEntryFlavor === null || adventureTemplateEntryFlavor === undefined) {
                throw new RequiredError('adventureTemplateEntryFlavor','Required parameter adventureTemplateEntryFlavor was null or undefined when calling updateAdventureTemplateEntryFlavor.');
            }
            const localVarPath = `/adventure_template_entry_flavor/{id}`
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
            const nonString = typeof adventureTemplateEntryFlavor !== 'string';
            const needsSerialization = nonString && configuration && configuration.isJsonMime
                ? configuration.isJsonMime(localVarRequestOptions.headers['Content-Type'])
                : nonString;
            localVarRequestOptions.data =  needsSerialization
                ? JSON.stringify(adventureTemplateEntryFlavor !== undefined ? adventureTemplateEntryFlavor : {})
                : (adventureTemplateEntryFlavor || "");
            return {
                url: localVarUrlObj.pathname + localVarUrlObj.search + localVarUrlObj.hash,
                options: localVarRequestOptions,
            };
        },
    }
};
export const AdventureTemplateEntryFlavorApiFp = function(configuration?: Configuration) {
    return {
        async createAdventureTemplateEntryFlavor(adventureTemplateEntryFlavor: ModelsAdventureTemplateEntryFlavor, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsAdventureTemplateEntryFlavor>>> {
            const localVarAxiosArgs = await AdventureTemplateEntryFlavorApiAxiosParamCreator(configuration).createAdventureTemplateEntryFlavor(adventureTemplateEntryFlavor, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        async deleteAdventureTemplateEntryFlavor(id: number, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<string>> {
            const localVarAxiosArgs = await AdventureTemplateEntryFlavorApiAxiosParamCreator(configuration).deleteAdventureTemplateEntryFlavor(id, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        async getAdventureTemplateEntryFlavor(id: number, includes?: string, select?: string, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsAdventureTemplateEntryFlavor>>> {
            const localVarAxiosArgs = await AdventureTemplateEntryFlavorApiAxiosParamCreator(configuration).getAdventureTemplateEntryFlavor(id, includes, select, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        async getAdventureTemplateEntryFlavorsBulk(body: CrudcontrollersBulkFetchByIdsGetRequest, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsAdventureTemplateEntryFlavor>>> {
            const localVarAxiosArgs = await AdventureTemplateEntryFlavorApiAxiosParamCreator(configuration).getAdventureTemplateEntryFlavorsBulk(body, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        async getAdventureTemplateEntryFlavorsCount(includes?: string, where?: string, whereOr?: string, groupBy?: string, limit?: string, page?: number, orderBy?: string, orderDirection?: string, select?: string, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsAdventureTemplateEntryFlavor>>> {
            const localVarAxiosArgs = await AdventureTemplateEntryFlavorApiAxiosParamCreator(configuration).getAdventureTemplateEntryFlavorsCount(includes, where, whereOr, groupBy, limit, page, orderBy, orderDirection, select, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        async listAdventureTemplateEntryFlavors(includes?: string, where?: string, whereOr?: string, groupBy?: string, limit?: string, page?: number, orderBy?: string, orderDirection?: string, select?: string, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsAdventureTemplateEntryFlavor>>> {
            const localVarAxiosArgs = await AdventureTemplateEntryFlavorApiAxiosParamCreator(configuration).listAdventureTemplateEntryFlavors(includes, where, whereOr, groupBy, limit, page, orderBy, orderDirection, select, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        async updateAdventureTemplateEntryFlavor(id: number, adventureTemplateEntryFlavor: ModelsAdventureTemplateEntryFlavor, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsAdventureTemplateEntryFlavor>>> {
            const localVarAxiosArgs = await AdventureTemplateEntryFlavorApiAxiosParamCreator(configuration).updateAdventureTemplateEntryFlavor(id, adventureTemplateEntryFlavor, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
    }
};
export const AdventureTemplateEntryFlavorApiFactory = function (configuration?: Configuration, basePath?: string, axios?: AxiosInstance) {
    return {
        createAdventureTemplateEntryFlavor(adventureTemplateEntryFlavor: ModelsAdventureTemplateEntryFlavor, options?: any): AxiosPromise<Array<ModelsAdventureTemplateEntryFlavor>> {
            return AdventureTemplateEntryFlavorApiFp(configuration).createAdventureTemplateEntryFlavor(adventureTemplateEntryFlavor, options).then((request) => request(axios, basePath));
        },
        deleteAdventureTemplateEntryFlavor(id: number, options?: any): AxiosPromise<string> {
            return AdventureTemplateEntryFlavorApiFp(configuration).deleteAdventureTemplateEntryFlavor(id, options).then((request) => request(axios, basePath));
        },
        getAdventureTemplateEntryFlavor(id: number, includes?: string, select?: string, options?: any): AxiosPromise<Array<ModelsAdventureTemplateEntryFlavor>> {
            return AdventureTemplateEntryFlavorApiFp(configuration).getAdventureTemplateEntryFlavor(id, includes, select, options).then((request) => request(axios, basePath));
        },
        getAdventureTemplateEntryFlavorsBulk(body: CrudcontrollersBulkFetchByIdsGetRequest, options?: any): AxiosPromise<Array<ModelsAdventureTemplateEntryFlavor>> {
            return AdventureTemplateEntryFlavorApiFp(configuration).getAdventureTemplateEntryFlavorsBulk(body, options).then((request) => request(axios, basePath));
        },
        getAdventureTemplateEntryFlavorsCount(includes?: string, where?: string, whereOr?: string, groupBy?: string, limit?: string, page?: number, orderBy?: string, orderDirection?: string, select?: string, options?: any): AxiosPromise<Array<ModelsAdventureTemplateEntryFlavor>> {
            return AdventureTemplateEntryFlavorApiFp(configuration).getAdventureTemplateEntryFlavorsCount(includes, where, whereOr, groupBy, limit, page, orderBy, orderDirection, select, options).then((request) => request(axios, basePath));
        },
        listAdventureTemplateEntryFlavors(includes?: string, where?: string, whereOr?: string, groupBy?: string, limit?: string, page?: number, orderBy?: string, orderDirection?: string, select?: string, options?: any): AxiosPromise<Array<ModelsAdventureTemplateEntryFlavor>> {
            return AdventureTemplateEntryFlavorApiFp(configuration).listAdventureTemplateEntryFlavors(includes, where, whereOr, groupBy, limit, page, orderBy, orderDirection, select, options).then((request) => request(axios, basePath));
        },
        updateAdventureTemplateEntryFlavor(id: number, adventureTemplateEntryFlavor: ModelsAdventureTemplateEntryFlavor, options?: any): AxiosPromise<Array<ModelsAdventureTemplateEntryFlavor>> {
            return AdventureTemplateEntryFlavorApiFp(configuration).updateAdventureTemplateEntryFlavor(id, adventureTemplateEntryFlavor, options).then((request) => request(axios, basePath));
        },
    };
};
export interface AdventureTemplateEntryFlavorApiCreateAdventureTemplateEntryFlavorRequest {
    readonly adventureTemplateEntryFlavor: ModelsAdventureTemplateEntryFlavor
}
export interface AdventureTemplateEntryFlavorApiDeleteAdventureTemplateEntryFlavorRequest {
    readonly id: number
}
export interface AdventureTemplateEntryFlavorApiGetAdventureTemplateEntryFlavorRequest {
    readonly id: number
    readonly includes?: string
    readonly select?: string
}
export interface AdventureTemplateEntryFlavorApiGetAdventureTemplateEntryFlavorsBulkRequest {
    readonly body: CrudcontrollersBulkFetchByIdsGetRequest
}
export interface AdventureTemplateEntryFlavorApiGetAdventureTemplateEntryFlavorsCountRequest {
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
export interface AdventureTemplateEntryFlavorApiListAdventureTemplateEntryFlavorsRequest {
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
export interface AdventureTemplateEntryFlavorApiUpdateAdventureTemplateEntryFlavorRequest {
    readonly id: number
    readonly adventureTemplateEntryFlavor: ModelsAdventureTemplateEntryFlavor
}
export class AdventureTemplateEntryFlavorApi extends BaseAPI {
    public createAdventureTemplateEntryFlavor(requestParameters: AdventureTemplateEntryFlavorApiCreateAdventureTemplateEntryFlavorRequest, options?: any) {
        return AdventureTemplateEntryFlavorApiFp(this.configuration).createAdventureTemplateEntryFlavor(requestParameters.adventureTemplateEntryFlavor, options).then((request) => request(this.axios, this.basePath));
    }
    public deleteAdventureTemplateEntryFlavor(requestParameters: AdventureTemplateEntryFlavorApiDeleteAdventureTemplateEntryFlavorRequest, options?: any) {
        return AdventureTemplateEntryFlavorApiFp(this.configuration).deleteAdventureTemplateEntryFlavor(requestParameters.id, options).then((request) => request(this.axios, this.basePath));
    }
    public getAdventureTemplateEntryFlavor(requestParameters: AdventureTemplateEntryFlavorApiGetAdventureTemplateEntryFlavorRequest, options?: any) {
        return AdventureTemplateEntryFlavorApiFp(this.configuration).getAdventureTemplateEntryFlavor(requestParameters.id, requestParameters.includes, requestParameters.select, options).then((request) => request(this.axios, this.basePath));
    }
    public getAdventureTemplateEntryFlavorsBulk(requestParameters: AdventureTemplateEntryFlavorApiGetAdventureTemplateEntryFlavorsBulkRequest, options?: any) {
        return AdventureTemplateEntryFlavorApiFp(this.configuration).getAdventureTemplateEntryFlavorsBulk(requestParameters.body, options).then((request) => request(this.axios, this.basePath));
    }
    public getAdventureTemplateEntryFlavorsCount(requestParameters: AdventureTemplateEntryFlavorApiGetAdventureTemplateEntryFlavorsCountRequest = {}, options?: any) {
        return AdventureTemplateEntryFlavorApiFp(this.configuration).getAdventureTemplateEntryFlavorsCount(requestParameters.includes, requestParameters.where, requestParameters.whereOr, requestParameters.groupBy, requestParameters.limit, requestParameters.page, requestParameters.orderBy, requestParameters.orderDirection, requestParameters.select, options).then((request) => request(this.axios, this.basePath));
    }
    public listAdventureTemplateEntryFlavors(requestParameters: AdventureTemplateEntryFlavorApiListAdventureTemplateEntryFlavorsRequest = {}, options?: any) {
        return AdventureTemplateEntryFlavorApiFp(this.configuration).listAdventureTemplateEntryFlavors(requestParameters.includes, requestParameters.where, requestParameters.whereOr, requestParameters.groupBy, requestParameters.limit, requestParameters.page, requestParameters.orderBy, requestParameters.orderDirection, requestParameters.select, options).then((request) => request(this.axios, this.basePath));
    }
    public updateAdventureTemplateEntryFlavor(requestParameters: AdventureTemplateEntryFlavorApiUpdateAdventureTemplateEntryFlavorRequest, options?: any) {
        return AdventureTemplateEntryFlavorApiFp(this.configuration).updateAdventureTemplateEntryFlavor(requestParameters.id, requestParameters.adventureTemplateEntryFlavor, options).then((request) => request(this.axios, this.basePath));
    }
}
