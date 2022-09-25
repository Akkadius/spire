import globalAxios, { AxiosPromise, AxiosInstance } from 'axios';
import { Configuration } from '../configuration';
// Some imports not used depending on template conditions
// @ts-ignore
import { BASE_PATH, COLLECTION_FORMATS, RequestArgs, BaseAPI, RequiredError } from '../base';
// @ts-ignore
import { CrudcontrollersBulkFetchByIdsGetRequest } from '../models';
// @ts-ignore
import { ModelsCompletedSharedTask } from '../models';
export const CompletedSharedTaskApiAxiosParamCreator = function (configuration?: Configuration) {
    return {
        createCompletedSharedTask: async (completedSharedTask: ModelsCompletedSharedTask, options: any = {}): Promise<RequestArgs> => {
            if (completedSharedTask === null || completedSharedTask === undefined) {
                throw new RequiredError('completedSharedTask','Required parameter completedSharedTask was null or undefined when calling createCompletedSharedTask.');
            }
            const localVarPath = `/completed_shared_task`;
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
            const nonString = typeof completedSharedTask !== 'string';
            const needsSerialization = nonString && configuration && configuration.isJsonMime
                ? configuration.isJsonMime(localVarRequestOptions.headers['Content-Type'])
                : nonString;
            localVarRequestOptions.data =  needsSerialization
                ? JSON.stringify(completedSharedTask !== undefined ? completedSharedTask : {})
                : (completedSharedTask || "");
            return {
                url: localVarUrlObj.pathname + localVarUrlObj.search + localVarUrlObj.hash,
                options: localVarRequestOptions,
            };
        },
        deleteCompletedSharedTask: async (id: number, options: any = {}): Promise<RequestArgs> => {
            if (id === null || id === undefined) {
                throw new RequiredError('id','Required parameter id was null or undefined when calling deleteCompletedSharedTask.');
            }
            const localVarPath = `/completed_shared_task/{id}`
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
        getCompletedSharedTask: async (id: number, includes?: string, select?: string, options: any = {}): Promise<RequestArgs> => {
            if (id === null || id === undefined) {
                throw new RequiredError('id','Required parameter id was null or undefined when calling getCompletedSharedTask.');
            }
            const localVarPath = `/completed_shared_task/{id}`
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
        getCompletedSharedTasksBulk: async (body: CrudcontrollersBulkFetchByIdsGetRequest, options: any = {}): Promise<RequestArgs> => {
            if (body === null || body === undefined) {
                throw new RequiredError('body','Required parameter body was null or undefined when calling getCompletedSharedTasksBulk.');
            }
            const localVarPath = `/completed_shared_tasks/bulk`;
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
        listCompletedSharedTasks: async (includes?: string, where?: string, whereOr?: string, groupBy?: string, limit?: string, page?: number, orderBy?: string, orderDirection?: string, select?: string, options: any = {}): Promise<RequestArgs> => {
            const localVarPath = `/completed_shared_tasks`;
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
        updateCompletedSharedTask: async (id: number, completedSharedTask: ModelsCompletedSharedTask, options: any = {}): Promise<RequestArgs> => {
            if (id === null || id === undefined) {
                throw new RequiredError('id','Required parameter id was null or undefined when calling updateCompletedSharedTask.');
            }
            if (completedSharedTask === null || completedSharedTask === undefined) {
                throw new RequiredError('completedSharedTask','Required parameter completedSharedTask was null or undefined when calling updateCompletedSharedTask.');
            }
            const localVarPath = `/completed_shared_task/{id}`
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
            const nonString = typeof completedSharedTask !== 'string';
            const needsSerialization = nonString && configuration && configuration.isJsonMime
                ? configuration.isJsonMime(localVarRequestOptions.headers['Content-Type'])
                : nonString;
            localVarRequestOptions.data =  needsSerialization
                ? JSON.stringify(completedSharedTask !== undefined ? completedSharedTask : {})
                : (completedSharedTask || "");
            return {
                url: localVarUrlObj.pathname + localVarUrlObj.search + localVarUrlObj.hash,
                options: localVarRequestOptions,
            };
        },
    }
};
export const CompletedSharedTaskApiFp = function(configuration?: Configuration) {
    return {
        async createCompletedSharedTask(completedSharedTask: ModelsCompletedSharedTask, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsCompletedSharedTask>>> {
            const localVarAxiosArgs = await CompletedSharedTaskApiAxiosParamCreator(configuration).createCompletedSharedTask(completedSharedTask, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        async deleteCompletedSharedTask(id: number, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<string>> {
            const localVarAxiosArgs = await CompletedSharedTaskApiAxiosParamCreator(configuration).deleteCompletedSharedTask(id, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        async getCompletedSharedTask(id: number, includes?: string, select?: string, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsCompletedSharedTask>>> {
            const localVarAxiosArgs = await CompletedSharedTaskApiAxiosParamCreator(configuration).getCompletedSharedTask(id, includes, select, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        async getCompletedSharedTasksBulk(body: CrudcontrollersBulkFetchByIdsGetRequest, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsCompletedSharedTask>>> {
            const localVarAxiosArgs = await CompletedSharedTaskApiAxiosParamCreator(configuration).getCompletedSharedTasksBulk(body, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        async listCompletedSharedTasks(includes?: string, where?: string, whereOr?: string, groupBy?: string, limit?: string, page?: number, orderBy?: string, orderDirection?: string, select?: string, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsCompletedSharedTask>>> {
            const localVarAxiosArgs = await CompletedSharedTaskApiAxiosParamCreator(configuration).listCompletedSharedTasks(includes, where, whereOr, groupBy, limit, page, orderBy, orderDirection, select, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        async updateCompletedSharedTask(id: number, completedSharedTask: ModelsCompletedSharedTask, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsCompletedSharedTask>>> {
            const localVarAxiosArgs = await CompletedSharedTaskApiAxiosParamCreator(configuration).updateCompletedSharedTask(id, completedSharedTask, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
    }
};
export const CompletedSharedTaskApiFactory = function (configuration?: Configuration, basePath?: string, axios?: AxiosInstance) {
    return {
        createCompletedSharedTask(completedSharedTask: ModelsCompletedSharedTask, options?: any): AxiosPromise<Array<ModelsCompletedSharedTask>> {
            return CompletedSharedTaskApiFp(configuration).createCompletedSharedTask(completedSharedTask, options).then((request) => request(axios, basePath));
        },
        deleteCompletedSharedTask(id: number, options?: any): AxiosPromise<string> {
            return CompletedSharedTaskApiFp(configuration).deleteCompletedSharedTask(id, options).then((request) => request(axios, basePath));
        },
        getCompletedSharedTask(id: number, includes?: string, select?: string, options?: any): AxiosPromise<Array<ModelsCompletedSharedTask>> {
            return CompletedSharedTaskApiFp(configuration).getCompletedSharedTask(id, includes, select, options).then((request) => request(axios, basePath));
        },
        getCompletedSharedTasksBulk(body: CrudcontrollersBulkFetchByIdsGetRequest, options?: any): AxiosPromise<Array<ModelsCompletedSharedTask>> {
            return CompletedSharedTaskApiFp(configuration).getCompletedSharedTasksBulk(body, options).then((request) => request(axios, basePath));
        },
        listCompletedSharedTasks(includes?: string, where?: string, whereOr?: string, groupBy?: string, limit?: string, page?: number, orderBy?: string, orderDirection?: string, select?: string, options?: any): AxiosPromise<Array<ModelsCompletedSharedTask>> {
            return CompletedSharedTaskApiFp(configuration).listCompletedSharedTasks(includes, where, whereOr, groupBy, limit, page, orderBy, orderDirection, select, options).then((request) => request(axios, basePath));
        },
        updateCompletedSharedTask(id: number, completedSharedTask: ModelsCompletedSharedTask, options?: any): AxiosPromise<Array<ModelsCompletedSharedTask>> {
            return CompletedSharedTaskApiFp(configuration).updateCompletedSharedTask(id, completedSharedTask, options).then((request) => request(axios, basePath));
        },
    };
};
export interface CompletedSharedTaskApiCreateCompletedSharedTaskRequest {
    readonly completedSharedTask: ModelsCompletedSharedTask
}
export interface CompletedSharedTaskApiDeleteCompletedSharedTaskRequest {
    readonly id: number
}
export interface CompletedSharedTaskApiGetCompletedSharedTaskRequest {
    readonly id: number
    readonly includes?: string
    readonly select?: string
}
export interface CompletedSharedTaskApiGetCompletedSharedTasksBulkRequest {
    readonly body: CrudcontrollersBulkFetchByIdsGetRequest
}
export interface CompletedSharedTaskApiListCompletedSharedTasksRequest {
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
export interface CompletedSharedTaskApiUpdateCompletedSharedTaskRequest {
    readonly id: number
    readonly completedSharedTask: ModelsCompletedSharedTask
}
export class CompletedSharedTaskApi extends BaseAPI {
    public createCompletedSharedTask(requestParameters: CompletedSharedTaskApiCreateCompletedSharedTaskRequest, options?: any) {
        return CompletedSharedTaskApiFp(this.configuration).createCompletedSharedTask(requestParameters.completedSharedTask, options).then((request) => request(this.axios, this.basePath));
    }
    public deleteCompletedSharedTask(requestParameters: CompletedSharedTaskApiDeleteCompletedSharedTaskRequest, options?: any) {
        return CompletedSharedTaskApiFp(this.configuration).deleteCompletedSharedTask(requestParameters.id, options).then((request) => request(this.axios, this.basePath));
    }
    public getCompletedSharedTask(requestParameters: CompletedSharedTaskApiGetCompletedSharedTaskRequest, options?: any) {
        return CompletedSharedTaskApiFp(this.configuration).getCompletedSharedTask(requestParameters.id, requestParameters.includes, requestParameters.select, options).then((request) => request(this.axios, this.basePath));
    }
    public getCompletedSharedTasksBulk(requestParameters: CompletedSharedTaskApiGetCompletedSharedTasksBulkRequest, options?: any) {
        return CompletedSharedTaskApiFp(this.configuration).getCompletedSharedTasksBulk(requestParameters.body, options).then((request) => request(this.axios, this.basePath));
    }
    public listCompletedSharedTasks(requestParameters: CompletedSharedTaskApiListCompletedSharedTasksRequest = {}, options?: any) {
        return CompletedSharedTaskApiFp(this.configuration).listCompletedSharedTasks(requestParameters.includes, requestParameters.where, requestParameters.whereOr, requestParameters.groupBy, requestParameters.limit, requestParameters.page, requestParameters.orderBy, requestParameters.orderDirection, requestParameters.select, options).then((request) => request(this.axios, this.basePath));
    }
    public updateCompletedSharedTask(requestParameters: CompletedSharedTaskApiUpdateCompletedSharedTaskRequest, options?: any) {
        return CompletedSharedTaskApiFp(this.configuration).updateCompletedSharedTask(requestParameters.id, requestParameters.completedSharedTask, options).then((request) => request(this.axios, this.basePath));
    }
}
