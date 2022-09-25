import globalAxios, { AxiosPromise, AxiosInstance } from 'axios';
import { Configuration } from '../configuration';
// Some imports not used depending on template conditions
// @ts-ignore
import { BASE_PATH, COLLECTION_FORMATS, RequestArgs, BaseAPI, RequiredError } from '../base';
// @ts-ignore
import { CrudcontrollersBulkFetchByIdsGetRequest } from '../models';
// @ts-ignore
import { ModelsMerchantlistTemp } from '../models';
export const MerchantlistTempApiAxiosParamCreator = function (configuration?: Configuration) {
    return {
        createMerchantlistTemp: async (merchantlistTemp: ModelsMerchantlistTemp, options: any = {}): Promise<RequestArgs> => {
            if (merchantlistTemp === null || merchantlistTemp === undefined) {
                throw new RequiredError('merchantlistTemp','Required parameter merchantlistTemp was null or undefined when calling createMerchantlistTemp.');
            }
            const localVarPath = `/merchantlist_temp`;
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
            const nonString = typeof merchantlistTemp !== 'string';
            const needsSerialization = nonString && configuration && configuration.isJsonMime
                ? configuration.isJsonMime(localVarRequestOptions.headers['Content-Type'])
                : nonString;
            localVarRequestOptions.data =  needsSerialization
                ? JSON.stringify(merchantlistTemp !== undefined ? merchantlistTemp : {})
                : (merchantlistTemp || "");
            return {
                url: localVarUrlObj.pathname + localVarUrlObj.search + localVarUrlObj.hash,
                options: localVarRequestOptions,
            };
        },
        deleteMerchantlistTemp: async (id: number, options: any = {}): Promise<RequestArgs> => {
            if (id === null || id === undefined) {
                throw new RequiredError('id','Required parameter id was null or undefined when calling deleteMerchantlistTemp.');
            }
            const localVarPath = `/merchantlist_temp/{id}`
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
        getMerchantlistTemp: async (id: number, includes?: string, select?: string, options: any = {}): Promise<RequestArgs> => {
            if (id === null || id === undefined) {
                throw new RequiredError('id','Required parameter id was null or undefined when calling getMerchantlistTemp.');
            }
            const localVarPath = `/merchantlist_temp/{id}`
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
        getMerchantlistTempsBulk: async (body: CrudcontrollersBulkFetchByIdsGetRequest, options: any = {}): Promise<RequestArgs> => {
            if (body === null || body === undefined) {
                throw new RequiredError('body','Required parameter body was null or undefined when calling getMerchantlistTempsBulk.');
            }
            const localVarPath = `/merchantlist_temps/bulk`;
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
        listMerchantlistTemps: async (includes?: string, where?: string, whereOr?: string, groupBy?: string, limit?: string, page?: number, orderBy?: string, orderDirection?: string, select?: string, options: any = {}): Promise<RequestArgs> => {
            const localVarPath = `/merchantlist_temps`;
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
        updateMerchantlistTemp: async (id: number, merchantlistTemp: ModelsMerchantlistTemp, options: any = {}): Promise<RequestArgs> => {
            if (id === null || id === undefined) {
                throw new RequiredError('id','Required parameter id was null or undefined when calling updateMerchantlistTemp.');
            }
            if (merchantlistTemp === null || merchantlistTemp === undefined) {
                throw new RequiredError('merchantlistTemp','Required parameter merchantlistTemp was null or undefined when calling updateMerchantlistTemp.');
            }
            const localVarPath = `/merchantlist_temp/{id}`
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
            const nonString = typeof merchantlistTemp !== 'string';
            const needsSerialization = nonString && configuration && configuration.isJsonMime
                ? configuration.isJsonMime(localVarRequestOptions.headers['Content-Type'])
                : nonString;
            localVarRequestOptions.data =  needsSerialization
                ? JSON.stringify(merchantlistTemp !== undefined ? merchantlistTemp : {})
                : (merchantlistTemp || "");
            return {
                url: localVarUrlObj.pathname + localVarUrlObj.search + localVarUrlObj.hash,
                options: localVarRequestOptions,
            };
        },
    }
};
export const MerchantlistTempApiFp = function(configuration?: Configuration) {
    return {
        async createMerchantlistTemp(merchantlistTemp: ModelsMerchantlistTemp, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsMerchantlistTemp>>> {
            const localVarAxiosArgs = await MerchantlistTempApiAxiosParamCreator(configuration).createMerchantlistTemp(merchantlistTemp, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        async deleteMerchantlistTemp(id: number, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<string>> {
            const localVarAxiosArgs = await MerchantlistTempApiAxiosParamCreator(configuration).deleteMerchantlistTemp(id, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        async getMerchantlistTemp(id: number, includes?: string, select?: string, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsMerchantlistTemp>>> {
            const localVarAxiosArgs = await MerchantlistTempApiAxiosParamCreator(configuration).getMerchantlistTemp(id, includes, select, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        async getMerchantlistTempsBulk(body: CrudcontrollersBulkFetchByIdsGetRequest, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsMerchantlistTemp>>> {
            const localVarAxiosArgs = await MerchantlistTempApiAxiosParamCreator(configuration).getMerchantlistTempsBulk(body, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        async listMerchantlistTemps(includes?: string, where?: string, whereOr?: string, groupBy?: string, limit?: string, page?: number, orderBy?: string, orderDirection?: string, select?: string, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsMerchantlistTemp>>> {
            const localVarAxiosArgs = await MerchantlistTempApiAxiosParamCreator(configuration).listMerchantlistTemps(includes, where, whereOr, groupBy, limit, page, orderBy, orderDirection, select, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        async updateMerchantlistTemp(id: number, merchantlistTemp: ModelsMerchantlistTemp, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsMerchantlistTemp>>> {
            const localVarAxiosArgs = await MerchantlistTempApiAxiosParamCreator(configuration).updateMerchantlistTemp(id, merchantlistTemp, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
    }
};
export const MerchantlistTempApiFactory = function (configuration?: Configuration, basePath?: string, axios?: AxiosInstance) {
    return {
        createMerchantlistTemp(merchantlistTemp: ModelsMerchantlistTemp, options?: any): AxiosPromise<Array<ModelsMerchantlistTemp>> {
            return MerchantlistTempApiFp(configuration).createMerchantlistTemp(merchantlistTemp, options).then((request) => request(axios, basePath));
        },
        deleteMerchantlistTemp(id: number, options?: any): AxiosPromise<string> {
            return MerchantlistTempApiFp(configuration).deleteMerchantlistTemp(id, options).then((request) => request(axios, basePath));
        },
        getMerchantlistTemp(id: number, includes?: string, select?: string, options?: any): AxiosPromise<Array<ModelsMerchantlistTemp>> {
            return MerchantlistTempApiFp(configuration).getMerchantlistTemp(id, includes, select, options).then((request) => request(axios, basePath));
        },
        getMerchantlistTempsBulk(body: CrudcontrollersBulkFetchByIdsGetRequest, options?: any): AxiosPromise<Array<ModelsMerchantlistTemp>> {
            return MerchantlistTempApiFp(configuration).getMerchantlistTempsBulk(body, options).then((request) => request(axios, basePath));
        },
        listMerchantlistTemps(includes?: string, where?: string, whereOr?: string, groupBy?: string, limit?: string, page?: number, orderBy?: string, orderDirection?: string, select?: string, options?: any): AxiosPromise<Array<ModelsMerchantlistTemp>> {
            return MerchantlistTempApiFp(configuration).listMerchantlistTemps(includes, where, whereOr, groupBy, limit, page, orderBy, orderDirection, select, options).then((request) => request(axios, basePath));
        },
        updateMerchantlistTemp(id: number, merchantlistTemp: ModelsMerchantlistTemp, options?: any): AxiosPromise<Array<ModelsMerchantlistTemp>> {
            return MerchantlistTempApiFp(configuration).updateMerchantlistTemp(id, merchantlistTemp, options).then((request) => request(axios, basePath));
        },
    };
};
export interface MerchantlistTempApiCreateMerchantlistTempRequest {
    readonly merchantlistTemp: ModelsMerchantlistTemp
}
export interface MerchantlistTempApiDeleteMerchantlistTempRequest {
    readonly id: number
}
export interface MerchantlistTempApiGetMerchantlistTempRequest {
    readonly id: number
    readonly includes?: string
    readonly select?: string
}
export interface MerchantlistTempApiGetMerchantlistTempsBulkRequest {
    readonly body: CrudcontrollersBulkFetchByIdsGetRequest
}
export interface MerchantlistTempApiListMerchantlistTempsRequest {
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
export interface MerchantlistTempApiUpdateMerchantlistTempRequest {
    readonly id: number
    readonly merchantlistTemp: ModelsMerchantlistTemp
}
export class MerchantlistTempApi extends BaseAPI {
    public createMerchantlistTemp(requestParameters: MerchantlistTempApiCreateMerchantlistTempRequest, options?: any) {
        return MerchantlistTempApiFp(this.configuration).createMerchantlistTemp(requestParameters.merchantlistTemp, options).then((request) => request(this.axios, this.basePath));
    }
    public deleteMerchantlistTemp(requestParameters: MerchantlistTempApiDeleteMerchantlistTempRequest, options?: any) {
        return MerchantlistTempApiFp(this.configuration).deleteMerchantlistTemp(requestParameters.id, options).then((request) => request(this.axios, this.basePath));
    }
    public getMerchantlistTemp(requestParameters: MerchantlistTempApiGetMerchantlistTempRequest, options?: any) {
        return MerchantlistTempApiFp(this.configuration).getMerchantlistTemp(requestParameters.id, requestParameters.includes, requestParameters.select, options).then((request) => request(this.axios, this.basePath));
    }
    public getMerchantlistTempsBulk(requestParameters: MerchantlistTempApiGetMerchantlistTempsBulkRequest, options?: any) {
        return MerchantlistTempApiFp(this.configuration).getMerchantlistTempsBulk(requestParameters.body, options).then((request) => request(this.axios, this.basePath));
    }
    public listMerchantlistTemps(requestParameters: MerchantlistTempApiListMerchantlistTempsRequest = {}, options?: any) {
        return MerchantlistTempApiFp(this.configuration).listMerchantlistTemps(requestParameters.includes, requestParameters.where, requestParameters.whereOr, requestParameters.groupBy, requestParameters.limit, requestParameters.page, requestParameters.orderBy, requestParameters.orderDirection, requestParameters.select, options).then((request) => request(this.axios, this.basePath));
    }
    public updateMerchantlistTemp(requestParameters: MerchantlistTempApiUpdateMerchantlistTempRequest, options?: any) {
        return MerchantlistTempApiFp(this.configuration).updateMerchantlistTemp(requestParameters.id, requestParameters.merchantlistTemp, options).then((request) => request(this.axios, this.basePath));
    }
}
