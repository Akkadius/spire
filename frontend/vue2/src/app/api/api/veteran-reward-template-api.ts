import globalAxios, { AxiosPromise, AxiosInstance } from 'axios';
import { Configuration } from '../configuration';
// Some imports not used depending on template conditions
// @ts-ignore
import { BASE_PATH, COLLECTION_FORMATS, RequestArgs, BaseAPI, RequiredError } from '../base';
// @ts-ignore
import { CrudcontrollersBulkFetchByIdsGetRequest } from '../models';
// @ts-ignore
import { ModelsVeteranRewardTemplate } from '../models';
export const VeteranRewardTemplateApiAxiosParamCreator = function (configuration?: Configuration) {
    return {
        createVeteranRewardTemplate: async (veteranRewardTemplate: ModelsVeteranRewardTemplate, options: any = {}): Promise<RequestArgs> => {
            if (veteranRewardTemplate === null || veteranRewardTemplate === undefined) {
                throw new RequiredError('veteranRewardTemplate','Required parameter veteranRewardTemplate was null or undefined when calling createVeteranRewardTemplate.');
            }
            const localVarPath = `/veteran_reward_template`;
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
            const nonString = typeof veteranRewardTemplate !== 'string';
            const needsSerialization = nonString && configuration && configuration.isJsonMime
                ? configuration.isJsonMime(localVarRequestOptions.headers['Content-Type'])
                : nonString;
            localVarRequestOptions.data =  needsSerialization
                ? JSON.stringify(veteranRewardTemplate !== undefined ? veteranRewardTemplate : {})
                : (veteranRewardTemplate || "");
            return {
                url: localVarUrlObj.pathname + localVarUrlObj.search + localVarUrlObj.hash,
                options: localVarRequestOptions,
            };
        },
        deleteVeteranRewardTemplate: async (id: number, options: any = {}): Promise<RequestArgs> => {
            if (id === null || id === undefined) {
                throw new RequiredError('id','Required parameter id was null or undefined when calling deleteVeteranRewardTemplate.');
            }
            const localVarPath = `/veteran_reward_template/{id}`
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
        getVeteranRewardTemplate: async (id: number, includes?: string, select?: string, options: any = {}): Promise<RequestArgs> => {
            if (id === null || id === undefined) {
                throw new RequiredError('id','Required parameter id was null or undefined when calling getVeteranRewardTemplate.');
            }
            const localVarPath = `/veteran_reward_template/{id}`
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
        getVeteranRewardTemplatesBulk: async (body: CrudcontrollersBulkFetchByIdsGetRequest, options: any = {}): Promise<RequestArgs> => {
            if (body === null || body === undefined) {
                throw new RequiredError('body','Required parameter body was null or undefined when calling getVeteranRewardTemplatesBulk.');
            }
            const localVarPath = `/veteran_reward_templates/bulk`;
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
        getVeteranRewardTemplatesCount: async (includes?: string, where?: string, whereOr?: string, groupBy?: string, limit?: string, page?: number, orderBy?: string, orderDirection?: string, select?: string, options: any = {}): Promise<RequestArgs> => {
            const localVarPath = `/veteran_reward_templates/count`;
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
        listVeteranRewardTemplates: async (includes?: string, where?: string, whereOr?: string, groupBy?: string, limit?: string, page?: number, orderBy?: string, orderDirection?: string, select?: string, options: any = {}): Promise<RequestArgs> => {
            const localVarPath = `/veteran_reward_templates`;
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
        updateVeteranRewardTemplate: async (id: number, veteranRewardTemplate: ModelsVeteranRewardTemplate, options: any = {}): Promise<RequestArgs> => {
            if (id === null || id === undefined) {
                throw new RequiredError('id','Required parameter id was null or undefined when calling updateVeteranRewardTemplate.');
            }
            if (veteranRewardTemplate === null || veteranRewardTemplate === undefined) {
                throw new RequiredError('veteranRewardTemplate','Required parameter veteranRewardTemplate was null or undefined when calling updateVeteranRewardTemplate.');
            }
            const localVarPath = `/veteran_reward_template/{id}`
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
            const nonString = typeof veteranRewardTemplate !== 'string';
            const needsSerialization = nonString && configuration && configuration.isJsonMime
                ? configuration.isJsonMime(localVarRequestOptions.headers['Content-Type'])
                : nonString;
            localVarRequestOptions.data =  needsSerialization
                ? JSON.stringify(veteranRewardTemplate !== undefined ? veteranRewardTemplate : {})
                : (veteranRewardTemplate || "");
            return {
                url: localVarUrlObj.pathname + localVarUrlObj.search + localVarUrlObj.hash,
                options: localVarRequestOptions,
            };
        },
    }
};
export const VeteranRewardTemplateApiFp = function(configuration?: Configuration) {
    return {
        async createVeteranRewardTemplate(veteranRewardTemplate: ModelsVeteranRewardTemplate, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsVeteranRewardTemplate>>> {
            const localVarAxiosArgs = await VeteranRewardTemplateApiAxiosParamCreator(configuration).createVeteranRewardTemplate(veteranRewardTemplate, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        async deleteVeteranRewardTemplate(id: number, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<string>> {
            const localVarAxiosArgs = await VeteranRewardTemplateApiAxiosParamCreator(configuration).deleteVeteranRewardTemplate(id, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        async getVeteranRewardTemplate(id: number, includes?: string, select?: string, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsVeteranRewardTemplate>>> {
            const localVarAxiosArgs = await VeteranRewardTemplateApiAxiosParamCreator(configuration).getVeteranRewardTemplate(id, includes, select, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        async getVeteranRewardTemplatesBulk(body: CrudcontrollersBulkFetchByIdsGetRequest, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsVeteranRewardTemplate>>> {
            const localVarAxiosArgs = await VeteranRewardTemplateApiAxiosParamCreator(configuration).getVeteranRewardTemplatesBulk(body, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        async getVeteranRewardTemplatesCount(includes?: string, where?: string, whereOr?: string, groupBy?: string, limit?: string, page?: number, orderBy?: string, orderDirection?: string, select?: string, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsVeteranRewardTemplate>>> {
            const localVarAxiosArgs = await VeteranRewardTemplateApiAxiosParamCreator(configuration).getVeteranRewardTemplatesCount(includes, where, whereOr, groupBy, limit, page, orderBy, orderDirection, select, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        async listVeteranRewardTemplates(includes?: string, where?: string, whereOr?: string, groupBy?: string, limit?: string, page?: number, orderBy?: string, orderDirection?: string, select?: string, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsVeteranRewardTemplate>>> {
            const localVarAxiosArgs = await VeteranRewardTemplateApiAxiosParamCreator(configuration).listVeteranRewardTemplates(includes, where, whereOr, groupBy, limit, page, orderBy, orderDirection, select, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        async updateVeteranRewardTemplate(id: number, veteranRewardTemplate: ModelsVeteranRewardTemplate, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsVeteranRewardTemplate>>> {
            const localVarAxiosArgs = await VeteranRewardTemplateApiAxiosParamCreator(configuration).updateVeteranRewardTemplate(id, veteranRewardTemplate, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
    }
};
export const VeteranRewardTemplateApiFactory = function (configuration?: Configuration, basePath?: string, axios?: AxiosInstance) {
    return {
        createVeteranRewardTemplate(veteranRewardTemplate: ModelsVeteranRewardTemplate, options?: any): AxiosPromise<Array<ModelsVeteranRewardTemplate>> {
            return VeteranRewardTemplateApiFp(configuration).createVeteranRewardTemplate(veteranRewardTemplate, options).then((request) => request(axios, basePath));
        },
        deleteVeteranRewardTemplate(id: number, options?: any): AxiosPromise<string> {
            return VeteranRewardTemplateApiFp(configuration).deleteVeteranRewardTemplate(id, options).then((request) => request(axios, basePath));
        },
        getVeteranRewardTemplate(id: number, includes?: string, select?: string, options?: any): AxiosPromise<Array<ModelsVeteranRewardTemplate>> {
            return VeteranRewardTemplateApiFp(configuration).getVeteranRewardTemplate(id, includes, select, options).then((request) => request(axios, basePath));
        },
        getVeteranRewardTemplatesBulk(body: CrudcontrollersBulkFetchByIdsGetRequest, options?: any): AxiosPromise<Array<ModelsVeteranRewardTemplate>> {
            return VeteranRewardTemplateApiFp(configuration).getVeteranRewardTemplatesBulk(body, options).then((request) => request(axios, basePath));
        },
        getVeteranRewardTemplatesCount(includes?: string, where?: string, whereOr?: string, groupBy?: string, limit?: string, page?: number, orderBy?: string, orderDirection?: string, select?: string, options?: any): AxiosPromise<Array<ModelsVeteranRewardTemplate>> {
            return VeteranRewardTemplateApiFp(configuration).getVeteranRewardTemplatesCount(includes, where, whereOr, groupBy, limit, page, orderBy, orderDirection, select, options).then((request) => request(axios, basePath));
        },
        listVeteranRewardTemplates(includes?: string, where?: string, whereOr?: string, groupBy?: string, limit?: string, page?: number, orderBy?: string, orderDirection?: string, select?: string, options?: any): AxiosPromise<Array<ModelsVeteranRewardTemplate>> {
            return VeteranRewardTemplateApiFp(configuration).listVeteranRewardTemplates(includes, where, whereOr, groupBy, limit, page, orderBy, orderDirection, select, options).then((request) => request(axios, basePath));
        },
        updateVeteranRewardTemplate(id: number, veteranRewardTemplate: ModelsVeteranRewardTemplate, options?: any): AxiosPromise<Array<ModelsVeteranRewardTemplate>> {
            return VeteranRewardTemplateApiFp(configuration).updateVeteranRewardTemplate(id, veteranRewardTemplate, options).then((request) => request(axios, basePath));
        },
    };
};
export interface VeteranRewardTemplateApiCreateVeteranRewardTemplateRequest {
    readonly veteranRewardTemplate: ModelsVeteranRewardTemplate
}
export interface VeteranRewardTemplateApiDeleteVeteranRewardTemplateRequest {
    readonly id: number
}
export interface VeteranRewardTemplateApiGetVeteranRewardTemplateRequest {
    readonly id: number
    readonly includes?: string
    readonly select?: string
}
export interface VeteranRewardTemplateApiGetVeteranRewardTemplatesBulkRequest {
    readonly body: CrudcontrollersBulkFetchByIdsGetRequest
}
export interface VeteranRewardTemplateApiGetVeteranRewardTemplatesCountRequest {
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
export interface VeteranRewardTemplateApiListVeteranRewardTemplatesRequest {
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
export interface VeteranRewardTemplateApiUpdateVeteranRewardTemplateRequest {
    readonly id: number
    readonly veteranRewardTemplate: ModelsVeteranRewardTemplate
}
export class VeteranRewardTemplateApi extends BaseAPI {
    public createVeteranRewardTemplate(requestParameters: VeteranRewardTemplateApiCreateVeteranRewardTemplateRequest, options?: any) {
        return VeteranRewardTemplateApiFp(this.configuration).createVeteranRewardTemplate(requestParameters.veteranRewardTemplate, options).then((request) => request(this.axios, this.basePath));
    }
    public deleteVeteranRewardTemplate(requestParameters: VeteranRewardTemplateApiDeleteVeteranRewardTemplateRequest, options?: any) {
        return VeteranRewardTemplateApiFp(this.configuration).deleteVeteranRewardTemplate(requestParameters.id, options).then((request) => request(this.axios, this.basePath));
    }
    public getVeteranRewardTemplate(requestParameters: VeteranRewardTemplateApiGetVeteranRewardTemplateRequest, options?: any) {
        return VeteranRewardTemplateApiFp(this.configuration).getVeteranRewardTemplate(requestParameters.id, requestParameters.includes, requestParameters.select, options).then((request) => request(this.axios, this.basePath));
    }
    public getVeteranRewardTemplatesBulk(requestParameters: VeteranRewardTemplateApiGetVeteranRewardTemplatesBulkRequest, options?: any) {
        return VeteranRewardTemplateApiFp(this.configuration).getVeteranRewardTemplatesBulk(requestParameters.body, options).then((request) => request(this.axios, this.basePath));
    }
    public getVeteranRewardTemplatesCount(requestParameters: VeteranRewardTemplateApiGetVeteranRewardTemplatesCountRequest = {}, options?: any) {
        return VeteranRewardTemplateApiFp(this.configuration).getVeteranRewardTemplatesCount(requestParameters.includes, requestParameters.where, requestParameters.whereOr, requestParameters.groupBy, requestParameters.limit, requestParameters.page, requestParameters.orderBy, requestParameters.orderDirection, requestParameters.select, options).then((request) => request(this.axios, this.basePath));
    }
    public listVeteranRewardTemplates(requestParameters: VeteranRewardTemplateApiListVeteranRewardTemplatesRequest = {}, options?: any) {
        return VeteranRewardTemplateApiFp(this.configuration).listVeteranRewardTemplates(requestParameters.includes, requestParameters.where, requestParameters.whereOr, requestParameters.groupBy, requestParameters.limit, requestParameters.page, requestParameters.orderBy, requestParameters.orderDirection, requestParameters.select, options).then((request) => request(this.axios, this.basePath));
    }
    public updateVeteranRewardTemplate(requestParameters: VeteranRewardTemplateApiUpdateVeteranRewardTemplateRequest, options?: any) {
        return VeteranRewardTemplateApiFp(this.configuration).updateVeteranRewardTemplate(requestParameters.id, requestParameters.veteranRewardTemplate, options).then((request) => request(this.axios, this.basePath));
    }
}
