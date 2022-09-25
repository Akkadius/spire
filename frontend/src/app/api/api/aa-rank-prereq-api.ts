import globalAxios, { AxiosPromise, AxiosInstance } from 'axios';
import { Configuration } from '../configuration';
// Some imports not used depending on template conditions
// @ts-ignore
import { BASE_PATH, COLLECTION_FORMATS, RequestArgs, BaseAPI, RequiredError } from '../base';
// @ts-ignore
import { CrudcontrollersBulkFetchByIdsGetRequest } from '../models';
// @ts-ignore
import { ModelsAaRankPrereq } from '../models';
export const AaRankPrereqApiAxiosParamCreator = function (configuration?: Configuration) {
    return {
        createAaRankPrereq: async (aaRankPrereq: ModelsAaRankPrereq, options: any = {}): Promise<RequestArgs> => {
            if (aaRankPrereq === null || aaRankPrereq === undefined) {
                throw new RequiredError('aaRankPrereq','Required parameter aaRankPrereq was null or undefined when calling createAaRankPrereq.');
            }
            const localVarPath = `/aa_rank_prereq`;
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
            const nonString = typeof aaRankPrereq !== 'string';
            const needsSerialization = nonString && configuration && configuration.isJsonMime
                ? configuration.isJsonMime(localVarRequestOptions.headers['Content-Type'])
                : nonString;
            localVarRequestOptions.data =  needsSerialization
                ? JSON.stringify(aaRankPrereq !== undefined ? aaRankPrereq : {})
                : (aaRankPrereq || "");
            return {
                url: localVarUrlObj.pathname + localVarUrlObj.search + localVarUrlObj.hash,
                options: localVarRequestOptions,
            };
        },
        deleteAaRankPrereq: async (id: number, options: any = {}): Promise<RequestArgs> => {
            if (id === null || id === undefined) {
                throw new RequiredError('id','Required parameter id was null or undefined when calling deleteAaRankPrereq.');
            }
            const localVarPath = `/aa_rank_prereq/{id}`
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
        getAaRankPrereq: async (id: number, includes?: string, select?: string, options: any = {}): Promise<RequestArgs> => {
            if (id === null || id === undefined) {
                throw new RequiredError('id','Required parameter id was null or undefined when calling getAaRankPrereq.');
            }
            const localVarPath = `/aa_rank_prereq/{id}`
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
        getAaRankPrereqsBulk: async (body: CrudcontrollersBulkFetchByIdsGetRequest, options: any = {}): Promise<RequestArgs> => {
            if (body === null || body === undefined) {
                throw new RequiredError('body','Required parameter body was null or undefined when calling getAaRankPrereqsBulk.');
            }
            const localVarPath = `/aa_rank_prereqs/bulk`;
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
        listAaRankPrereqs: async (includes?: string, where?: string, whereOr?: string, groupBy?: string, limit?: string, page?: number, orderBy?: string, orderDirection?: string, select?: string, options: any = {}): Promise<RequestArgs> => {
            const localVarPath = `/aa_rank_prereqs`;
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
        updateAaRankPrereq: async (id: number, aaRankPrereq: ModelsAaRankPrereq, options: any = {}): Promise<RequestArgs> => {
            if (id === null || id === undefined) {
                throw new RequiredError('id','Required parameter id was null or undefined when calling updateAaRankPrereq.');
            }
            if (aaRankPrereq === null || aaRankPrereq === undefined) {
                throw new RequiredError('aaRankPrereq','Required parameter aaRankPrereq was null or undefined when calling updateAaRankPrereq.');
            }
            const localVarPath = `/aa_rank_prereq/{id}`
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
            const nonString = typeof aaRankPrereq !== 'string';
            const needsSerialization = nonString && configuration && configuration.isJsonMime
                ? configuration.isJsonMime(localVarRequestOptions.headers['Content-Type'])
                : nonString;
            localVarRequestOptions.data =  needsSerialization
                ? JSON.stringify(aaRankPrereq !== undefined ? aaRankPrereq : {})
                : (aaRankPrereq || "");
            return {
                url: localVarUrlObj.pathname + localVarUrlObj.search + localVarUrlObj.hash,
                options: localVarRequestOptions,
            };
        },
    }
};
export const AaRankPrereqApiFp = function(configuration?: Configuration) {
    return {
        async createAaRankPrereq(aaRankPrereq: ModelsAaRankPrereq, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsAaRankPrereq>>> {
            const localVarAxiosArgs = await AaRankPrereqApiAxiosParamCreator(configuration).createAaRankPrereq(aaRankPrereq, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        async deleteAaRankPrereq(id: number, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<string>> {
            const localVarAxiosArgs = await AaRankPrereqApiAxiosParamCreator(configuration).deleteAaRankPrereq(id, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        async getAaRankPrereq(id: number, includes?: string, select?: string, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsAaRankPrereq>>> {
            const localVarAxiosArgs = await AaRankPrereqApiAxiosParamCreator(configuration).getAaRankPrereq(id, includes, select, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        async getAaRankPrereqsBulk(body: CrudcontrollersBulkFetchByIdsGetRequest, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsAaRankPrereq>>> {
            const localVarAxiosArgs = await AaRankPrereqApiAxiosParamCreator(configuration).getAaRankPrereqsBulk(body, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        async listAaRankPrereqs(includes?: string, where?: string, whereOr?: string, groupBy?: string, limit?: string, page?: number, orderBy?: string, orderDirection?: string, select?: string, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsAaRankPrereq>>> {
            const localVarAxiosArgs = await AaRankPrereqApiAxiosParamCreator(configuration).listAaRankPrereqs(includes, where, whereOr, groupBy, limit, page, orderBy, orderDirection, select, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        async updateAaRankPrereq(id: number, aaRankPrereq: ModelsAaRankPrereq, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsAaRankPrereq>>> {
            const localVarAxiosArgs = await AaRankPrereqApiAxiosParamCreator(configuration).updateAaRankPrereq(id, aaRankPrereq, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
    }
};
export const AaRankPrereqApiFactory = function (configuration?: Configuration, basePath?: string, axios?: AxiosInstance) {
    return {
        createAaRankPrereq(aaRankPrereq: ModelsAaRankPrereq, options?: any): AxiosPromise<Array<ModelsAaRankPrereq>> {
            return AaRankPrereqApiFp(configuration).createAaRankPrereq(aaRankPrereq, options).then((request) => request(axios, basePath));
        },
        deleteAaRankPrereq(id: number, options?: any): AxiosPromise<string> {
            return AaRankPrereqApiFp(configuration).deleteAaRankPrereq(id, options).then((request) => request(axios, basePath));
        },
        getAaRankPrereq(id: number, includes?: string, select?: string, options?: any): AxiosPromise<Array<ModelsAaRankPrereq>> {
            return AaRankPrereqApiFp(configuration).getAaRankPrereq(id, includes, select, options).then((request) => request(axios, basePath));
        },
        getAaRankPrereqsBulk(body: CrudcontrollersBulkFetchByIdsGetRequest, options?: any): AxiosPromise<Array<ModelsAaRankPrereq>> {
            return AaRankPrereqApiFp(configuration).getAaRankPrereqsBulk(body, options).then((request) => request(axios, basePath));
        },
        listAaRankPrereqs(includes?: string, where?: string, whereOr?: string, groupBy?: string, limit?: string, page?: number, orderBy?: string, orderDirection?: string, select?: string, options?: any): AxiosPromise<Array<ModelsAaRankPrereq>> {
            return AaRankPrereqApiFp(configuration).listAaRankPrereqs(includes, where, whereOr, groupBy, limit, page, orderBy, orderDirection, select, options).then((request) => request(axios, basePath));
        },
        updateAaRankPrereq(id: number, aaRankPrereq: ModelsAaRankPrereq, options?: any): AxiosPromise<Array<ModelsAaRankPrereq>> {
            return AaRankPrereqApiFp(configuration).updateAaRankPrereq(id, aaRankPrereq, options).then((request) => request(axios, basePath));
        },
    };
};
export interface AaRankPrereqApiCreateAaRankPrereqRequest {
    readonly aaRankPrereq: ModelsAaRankPrereq
}
export interface AaRankPrereqApiDeleteAaRankPrereqRequest {
    readonly id: number
}
export interface AaRankPrereqApiGetAaRankPrereqRequest {
    readonly id: number
    readonly includes?: string
    readonly select?: string
}
export interface AaRankPrereqApiGetAaRankPrereqsBulkRequest {
    readonly body: CrudcontrollersBulkFetchByIdsGetRequest
}
export interface AaRankPrereqApiListAaRankPrereqsRequest {
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
export interface AaRankPrereqApiUpdateAaRankPrereqRequest {
    readonly id: number
    readonly aaRankPrereq: ModelsAaRankPrereq
}
export class AaRankPrereqApi extends BaseAPI {
    public createAaRankPrereq(requestParameters: AaRankPrereqApiCreateAaRankPrereqRequest, options?: any) {
        return AaRankPrereqApiFp(this.configuration).createAaRankPrereq(requestParameters.aaRankPrereq, options).then((request) => request(this.axios, this.basePath));
    }
    public deleteAaRankPrereq(requestParameters: AaRankPrereqApiDeleteAaRankPrereqRequest, options?: any) {
        return AaRankPrereqApiFp(this.configuration).deleteAaRankPrereq(requestParameters.id, options).then((request) => request(this.axios, this.basePath));
    }
    public getAaRankPrereq(requestParameters: AaRankPrereqApiGetAaRankPrereqRequest, options?: any) {
        return AaRankPrereqApiFp(this.configuration).getAaRankPrereq(requestParameters.id, requestParameters.includes, requestParameters.select, options).then((request) => request(this.axios, this.basePath));
    }
    public getAaRankPrereqsBulk(requestParameters: AaRankPrereqApiGetAaRankPrereqsBulkRequest, options?: any) {
        return AaRankPrereqApiFp(this.configuration).getAaRankPrereqsBulk(requestParameters.body, options).then((request) => request(this.axios, this.basePath));
    }
    public listAaRankPrereqs(requestParameters: AaRankPrereqApiListAaRankPrereqsRequest = {}, options?: any) {
        return AaRankPrereqApiFp(this.configuration).listAaRankPrereqs(requestParameters.includes, requestParameters.where, requestParameters.whereOr, requestParameters.groupBy, requestParameters.limit, requestParameters.page, requestParameters.orderBy, requestParameters.orderDirection, requestParameters.select, options).then((request) => request(this.axios, this.basePath));
    }
    public updateAaRankPrereq(requestParameters: AaRankPrereqApiUpdateAaRankPrereqRequest, options?: any) {
        return AaRankPrereqApiFp(this.configuration).updateAaRankPrereq(requestParameters.id, requestParameters.aaRankPrereq, options).then((request) => request(this.axios, this.basePath));
    }
}
