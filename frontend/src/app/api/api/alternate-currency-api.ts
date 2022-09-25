import globalAxios, { AxiosPromise, AxiosInstance } from 'axios';
import { Configuration } from '../configuration';
// Some imports not used depending on template conditions
// @ts-ignore
import { BASE_PATH, COLLECTION_FORMATS, RequestArgs, BaseAPI, RequiredError } from '../base';
// @ts-ignore
import { CrudcontrollersBulkFetchByIdsGetRequest } from '../models';
// @ts-ignore
import { ModelsAlternateCurrency } from '../models';
export const AlternateCurrencyApiAxiosParamCreator = function (configuration?: Configuration) {
    return {
        createAlternateCurrency: async (alternateCurrency: ModelsAlternateCurrency, options: any = {}): Promise<RequestArgs> => {
            if (alternateCurrency === null || alternateCurrency === undefined) {
                throw new RequiredError('alternateCurrency','Required parameter alternateCurrency was null or undefined when calling createAlternateCurrency.');
            }
            const localVarPath = `/alternate_currency`;
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
            const nonString = typeof alternateCurrency !== 'string';
            const needsSerialization = nonString && configuration && configuration.isJsonMime
                ? configuration.isJsonMime(localVarRequestOptions.headers['Content-Type'])
                : nonString;
            localVarRequestOptions.data =  needsSerialization
                ? JSON.stringify(alternateCurrency !== undefined ? alternateCurrency : {})
                : (alternateCurrency || "");
            return {
                url: localVarUrlObj.pathname + localVarUrlObj.search + localVarUrlObj.hash,
                options: localVarRequestOptions,
            };
        },
        deleteAlternateCurrency: async (id: number, options: any = {}): Promise<RequestArgs> => {
            if (id === null || id === undefined) {
                throw new RequiredError('id','Required parameter id was null or undefined when calling deleteAlternateCurrency.');
            }
            const localVarPath = `/alternate_currency/{id}`
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
        getAlternateCurrenciesBulk: async (body: CrudcontrollersBulkFetchByIdsGetRequest, options: any = {}): Promise<RequestArgs> => {
            if (body === null || body === undefined) {
                throw new RequiredError('body','Required parameter body was null or undefined when calling getAlternateCurrenciesBulk.');
            }
            const localVarPath = `/alternate_currencies/bulk`;
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
        getAlternateCurrency: async (id: number, includes?: string, select?: string, options: any = {}): Promise<RequestArgs> => {
            if (id === null || id === undefined) {
                throw new RequiredError('id','Required parameter id was null or undefined when calling getAlternateCurrency.');
            }
            const localVarPath = `/alternate_currency/{id}`
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
        listAlternateCurrencies: async (includes?: string, where?: string, whereOr?: string, groupBy?: string, limit?: string, page?: number, orderBy?: string, orderDirection?: string, select?: string, options: any = {}): Promise<RequestArgs> => {
            const localVarPath = `/alternate_currencies`;
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
        updateAlternateCurrency: async (id: number, alternateCurrency: ModelsAlternateCurrency, options: any = {}): Promise<RequestArgs> => {
            if (id === null || id === undefined) {
                throw new RequiredError('id','Required parameter id was null or undefined when calling updateAlternateCurrency.');
            }
            if (alternateCurrency === null || alternateCurrency === undefined) {
                throw new RequiredError('alternateCurrency','Required parameter alternateCurrency was null or undefined when calling updateAlternateCurrency.');
            }
            const localVarPath = `/alternate_currency/{id}`
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
            const nonString = typeof alternateCurrency !== 'string';
            const needsSerialization = nonString && configuration && configuration.isJsonMime
                ? configuration.isJsonMime(localVarRequestOptions.headers['Content-Type'])
                : nonString;
            localVarRequestOptions.data =  needsSerialization
                ? JSON.stringify(alternateCurrency !== undefined ? alternateCurrency : {})
                : (alternateCurrency || "");
            return {
                url: localVarUrlObj.pathname + localVarUrlObj.search + localVarUrlObj.hash,
                options: localVarRequestOptions,
            };
        },
    }
};
export const AlternateCurrencyApiFp = function(configuration?: Configuration) {
    return {
        async createAlternateCurrency(alternateCurrency: ModelsAlternateCurrency, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsAlternateCurrency>>> {
            const localVarAxiosArgs = await AlternateCurrencyApiAxiosParamCreator(configuration).createAlternateCurrency(alternateCurrency, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        async deleteAlternateCurrency(id: number, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<string>> {
            const localVarAxiosArgs = await AlternateCurrencyApiAxiosParamCreator(configuration).deleteAlternateCurrency(id, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        async getAlternateCurrenciesBulk(body: CrudcontrollersBulkFetchByIdsGetRequest, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsAlternateCurrency>>> {
            const localVarAxiosArgs = await AlternateCurrencyApiAxiosParamCreator(configuration).getAlternateCurrenciesBulk(body, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        async getAlternateCurrency(id: number, includes?: string, select?: string, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsAlternateCurrency>>> {
            const localVarAxiosArgs = await AlternateCurrencyApiAxiosParamCreator(configuration).getAlternateCurrency(id, includes, select, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        async listAlternateCurrencies(includes?: string, where?: string, whereOr?: string, groupBy?: string, limit?: string, page?: number, orderBy?: string, orderDirection?: string, select?: string, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsAlternateCurrency>>> {
            const localVarAxiosArgs = await AlternateCurrencyApiAxiosParamCreator(configuration).listAlternateCurrencies(includes, where, whereOr, groupBy, limit, page, orderBy, orderDirection, select, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        async updateAlternateCurrency(id: number, alternateCurrency: ModelsAlternateCurrency, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsAlternateCurrency>>> {
            const localVarAxiosArgs = await AlternateCurrencyApiAxiosParamCreator(configuration).updateAlternateCurrency(id, alternateCurrency, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
    }
};
export const AlternateCurrencyApiFactory = function (configuration?: Configuration, basePath?: string, axios?: AxiosInstance) {
    return {
        createAlternateCurrency(alternateCurrency: ModelsAlternateCurrency, options?: any): AxiosPromise<Array<ModelsAlternateCurrency>> {
            return AlternateCurrencyApiFp(configuration).createAlternateCurrency(alternateCurrency, options).then((request) => request(axios, basePath));
        },
        deleteAlternateCurrency(id: number, options?: any): AxiosPromise<string> {
            return AlternateCurrencyApiFp(configuration).deleteAlternateCurrency(id, options).then((request) => request(axios, basePath));
        },
        getAlternateCurrenciesBulk(body: CrudcontrollersBulkFetchByIdsGetRequest, options?: any): AxiosPromise<Array<ModelsAlternateCurrency>> {
            return AlternateCurrencyApiFp(configuration).getAlternateCurrenciesBulk(body, options).then((request) => request(axios, basePath));
        },
        getAlternateCurrency(id: number, includes?: string, select?: string, options?: any): AxiosPromise<Array<ModelsAlternateCurrency>> {
            return AlternateCurrencyApiFp(configuration).getAlternateCurrency(id, includes, select, options).then((request) => request(axios, basePath));
        },
        listAlternateCurrencies(includes?: string, where?: string, whereOr?: string, groupBy?: string, limit?: string, page?: number, orderBy?: string, orderDirection?: string, select?: string, options?: any): AxiosPromise<Array<ModelsAlternateCurrency>> {
            return AlternateCurrencyApiFp(configuration).listAlternateCurrencies(includes, where, whereOr, groupBy, limit, page, orderBy, orderDirection, select, options).then((request) => request(axios, basePath));
        },
        updateAlternateCurrency(id: number, alternateCurrency: ModelsAlternateCurrency, options?: any): AxiosPromise<Array<ModelsAlternateCurrency>> {
            return AlternateCurrencyApiFp(configuration).updateAlternateCurrency(id, alternateCurrency, options).then((request) => request(axios, basePath));
        },
    };
};
export interface AlternateCurrencyApiCreateAlternateCurrencyRequest {
    readonly alternateCurrency: ModelsAlternateCurrency
}
export interface AlternateCurrencyApiDeleteAlternateCurrencyRequest {
    readonly id: number
}
export interface AlternateCurrencyApiGetAlternateCurrenciesBulkRequest {
    readonly body: CrudcontrollersBulkFetchByIdsGetRequest
}
export interface AlternateCurrencyApiGetAlternateCurrencyRequest {
    readonly id: number
    readonly includes?: string
    readonly select?: string
}
export interface AlternateCurrencyApiListAlternateCurrenciesRequest {
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
export interface AlternateCurrencyApiUpdateAlternateCurrencyRequest {
    readonly id: number
    readonly alternateCurrency: ModelsAlternateCurrency
}
export class AlternateCurrencyApi extends BaseAPI {
    public createAlternateCurrency(requestParameters: AlternateCurrencyApiCreateAlternateCurrencyRequest, options?: any) {
        return AlternateCurrencyApiFp(this.configuration).createAlternateCurrency(requestParameters.alternateCurrency, options).then((request) => request(this.axios, this.basePath));
    }
    public deleteAlternateCurrency(requestParameters: AlternateCurrencyApiDeleteAlternateCurrencyRequest, options?: any) {
        return AlternateCurrencyApiFp(this.configuration).deleteAlternateCurrency(requestParameters.id, options).then((request) => request(this.axios, this.basePath));
    }
    public getAlternateCurrenciesBulk(requestParameters: AlternateCurrencyApiGetAlternateCurrenciesBulkRequest, options?: any) {
        return AlternateCurrencyApiFp(this.configuration).getAlternateCurrenciesBulk(requestParameters.body, options).then((request) => request(this.axios, this.basePath));
    }
    public getAlternateCurrency(requestParameters: AlternateCurrencyApiGetAlternateCurrencyRequest, options?: any) {
        return AlternateCurrencyApiFp(this.configuration).getAlternateCurrency(requestParameters.id, requestParameters.includes, requestParameters.select, options).then((request) => request(this.axios, this.basePath));
    }
    public listAlternateCurrencies(requestParameters: AlternateCurrencyApiListAlternateCurrenciesRequest = {}, options?: any) {
        return AlternateCurrencyApiFp(this.configuration).listAlternateCurrencies(requestParameters.includes, requestParameters.where, requestParameters.whereOr, requestParameters.groupBy, requestParameters.limit, requestParameters.page, requestParameters.orderBy, requestParameters.orderDirection, requestParameters.select, options).then((request) => request(this.axios, this.basePath));
    }
    public updateAlternateCurrency(requestParameters: AlternateCurrencyApiUpdateAlternateCurrencyRequest, options?: any) {
        return AlternateCurrencyApiFp(this.configuration).updateAlternateCurrency(requestParameters.id, requestParameters.alternateCurrency, options).then((request) => request(this.axios, this.basePath));
    }
}
