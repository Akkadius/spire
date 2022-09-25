import globalAxios, { AxiosPromise, AxiosInstance } from 'axios';
import { Configuration } from '../configuration';
// Some imports not used depending on template conditions
// @ts-ignore
import { BASE_PATH, COLLECTION_FORMATS, RequestArgs, BaseAPI, RequiredError } from '../base';
// @ts-ignore
import { CrudcontrollersBulkFetchByIdsGetRequest } from '../models';
// @ts-ignore
import { ModelsLoginServerAdmin } from '../models';
export const LoginServerAdminApiAxiosParamCreator = function (configuration?: Configuration) {
    return {
        createLoginServerAdmin: async (loginServerAdmin: ModelsLoginServerAdmin, options: any = {}): Promise<RequestArgs> => {
            if (loginServerAdmin === null || loginServerAdmin === undefined) {
                throw new RequiredError('loginServerAdmin','Required parameter loginServerAdmin was null or undefined when calling createLoginServerAdmin.');
            }
            const localVarPath = `/login_server_admin`;
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
            const nonString = typeof loginServerAdmin !== 'string';
            const needsSerialization = nonString && configuration && configuration.isJsonMime
                ? configuration.isJsonMime(localVarRequestOptions.headers['Content-Type'])
                : nonString;
            localVarRequestOptions.data =  needsSerialization
                ? JSON.stringify(loginServerAdmin !== undefined ? loginServerAdmin : {})
                : (loginServerAdmin || "");
            return {
                url: localVarUrlObj.pathname + localVarUrlObj.search + localVarUrlObj.hash,
                options: localVarRequestOptions,
            };
        },
        deleteLoginServerAdmin: async (id: number, options: any = {}): Promise<RequestArgs> => {
            if (id === null || id === undefined) {
                throw new RequiredError('id','Required parameter id was null or undefined when calling deleteLoginServerAdmin.');
            }
            const localVarPath = `/login_server_admin/{id}`
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
        getLoginServerAdmin: async (id: number, includes?: string, select?: string, options: any = {}): Promise<RequestArgs> => {
            if (id === null || id === undefined) {
                throw new RequiredError('id','Required parameter id was null or undefined when calling getLoginServerAdmin.');
            }
            const localVarPath = `/login_server_admin/{id}`
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
        getLoginServerAdminsBulk: async (body: CrudcontrollersBulkFetchByIdsGetRequest, options: any = {}): Promise<RequestArgs> => {
            if (body === null || body === undefined) {
                throw new RequiredError('body','Required parameter body was null or undefined when calling getLoginServerAdminsBulk.');
            }
            const localVarPath = `/login_server_admins/bulk`;
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
        listLoginServerAdmins: async (includes?: string, where?: string, whereOr?: string, groupBy?: string, limit?: string, page?: number, orderBy?: string, orderDirection?: string, select?: string, options: any = {}): Promise<RequestArgs> => {
            const localVarPath = `/login_server_admins`;
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
        updateLoginServerAdmin: async (id: number, loginServerAdmin: ModelsLoginServerAdmin, options: any = {}): Promise<RequestArgs> => {
            if (id === null || id === undefined) {
                throw new RequiredError('id','Required parameter id was null or undefined when calling updateLoginServerAdmin.');
            }
            if (loginServerAdmin === null || loginServerAdmin === undefined) {
                throw new RequiredError('loginServerAdmin','Required parameter loginServerAdmin was null or undefined when calling updateLoginServerAdmin.');
            }
            const localVarPath = `/login_server_admin/{id}`
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
            const nonString = typeof loginServerAdmin !== 'string';
            const needsSerialization = nonString && configuration && configuration.isJsonMime
                ? configuration.isJsonMime(localVarRequestOptions.headers['Content-Type'])
                : nonString;
            localVarRequestOptions.data =  needsSerialization
                ? JSON.stringify(loginServerAdmin !== undefined ? loginServerAdmin : {})
                : (loginServerAdmin || "");
            return {
                url: localVarUrlObj.pathname + localVarUrlObj.search + localVarUrlObj.hash,
                options: localVarRequestOptions,
            };
        },
    }
};
export const LoginServerAdminApiFp = function(configuration?: Configuration) {
    return {
        async createLoginServerAdmin(loginServerAdmin: ModelsLoginServerAdmin, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsLoginServerAdmin>>> {
            const localVarAxiosArgs = await LoginServerAdminApiAxiosParamCreator(configuration).createLoginServerAdmin(loginServerAdmin, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        async deleteLoginServerAdmin(id: number, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<string>> {
            const localVarAxiosArgs = await LoginServerAdminApiAxiosParamCreator(configuration).deleteLoginServerAdmin(id, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        async getLoginServerAdmin(id: number, includes?: string, select?: string, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsLoginServerAdmin>>> {
            const localVarAxiosArgs = await LoginServerAdminApiAxiosParamCreator(configuration).getLoginServerAdmin(id, includes, select, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        async getLoginServerAdminsBulk(body: CrudcontrollersBulkFetchByIdsGetRequest, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsLoginServerAdmin>>> {
            const localVarAxiosArgs = await LoginServerAdminApiAxiosParamCreator(configuration).getLoginServerAdminsBulk(body, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        async listLoginServerAdmins(includes?: string, where?: string, whereOr?: string, groupBy?: string, limit?: string, page?: number, orderBy?: string, orderDirection?: string, select?: string, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsLoginServerAdmin>>> {
            const localVarAxiosArgs = await LoginServerAdminApiAxiosParamCreator(configuration).listLoginServerAdmins(includes, where, whereOr, groupBy, limit, page, orderBy, orderDirection, select, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        async updateLoginServerAdmin(id: number, loginServerAdmin: ModelsLoginServerAdmin, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsLoginServerAdmin>>> {
            const localVarAxiosArgs = await LoginServerAdminApiAxiosParamCreator(configuration).updateLoginServerAdmin(id, loginServerAdmin, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
    }
};
export const LoginServerAdminApiFactory = function (configuration?: Configuration, basePath?: string, axios?: AxiosInstance) {
    return {
        createLoginServerAdmin(loginServerAdmin: ModelsLoginServerAdmin, options?: any): AxiosPromise<Array<ModelsLoginServerAdmin>> {
            return LoginServerAdminApiFp(configuration).createLoginServerAdmin(loginServerAdmin, options).then((request) => request(axios, basePath));
        },
        deleteLoginServerAdmin(id: number, options?: any): AxiosPromise<string> {
            return LoginServerAdminApiFp(configuration).deleteLoginServerAdmin(id, options).then((request) => request(axios, basePath));
        },
        getLoginServerAdmin(id: number, includes?: string, select?: string, options?: any): AxiosPromise<Array<ModelsLoginServerAdmin>> {
            return LoginServerAdminApiFp(configuration).getLoginServerAdmin(id, includes, select, options).then((request) => request(axios, basePath));
        },
        getLoginServerAdminsBulk(body: CrudcontrollersBulkFetchByIdsGetRequest, options?: any): AxiosPromise<Array<ModelsLoginServerAdmin>> {
            return LoginServerAdminApiFp(configuration).getLoginServerAdminsBulk(body, options).then((request) => request(axios, basePath));
        },
        listLoginServerAdmins(includes?: string, where?: string, whereOr?: string, groupBy?: string, limit?: string, page?: number, orderBy?: string, orderDirection?: string, select?: string, options?: any): AxiosPromise<Array<ModelsLoginServerAdmin>> {
            return LoginServerAdminApiFp(configuration).listLoginServerAdmins(includes, where, whereOr, groupBy, limit, page, orderBy, orderDirection, select, options).then((request) => request(axios, basePath));
        },
        updateLoginServerAdmin(id: number, loginServerAdmin: ModelsLoginServerAdmin, options?: any): AxiosPromise<Array<ModelsLoginServerAdmin>> {
            return LoginServerAdminApiFp(configuration).updateLoginServerAdmin(id, loginServerAdmin, options).then((request) => request(axios, basePath));
        },
    };
};
export interface LoginServerAdminApiCreateLoginServerAdminRequest {
    readonly loginServerAdmin: ModelsLoginServerAdmin
}
export interface LoginServerAdminApiDeleteLoginServerAdminRequest {
    readonly id: number
}
export interface LoginServerAdminApiGetLoginServerAdminRequest {
    readonly id: number
    readonly includes?: string
    readonly select?: string
}
export interface LoginServerAdminApiGetLoginServerAdminsBulkRequest {
    readonly body: CrudcontrollersBulkFetchByIdsGetRequest
}
export interface LoginServerAdminApiListLoginServerAdminsRequest {
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
export interface LoginServerAdminApiUpdateLoginServerAdminRequest {
    readonly id: number
    readonly loginServerAdmin: ModelsLoginServerAdmin
}
export class LoginServerAdminApi extends BaseAPI {
    public createLoginServerAdmin(requestParameters: LoginServerAdminApiCreateLoginServerAdminRequest, options?: any) {
        return LoginServerAdminApiFp(this.configuration).createLoginServerAdmin(requestParameters.loginServerAdmin, options).then((request) => request(this.axios, this.basePath));
    }
    public deleteLoginServerAdmin(requestParameters: LoginServerAdminApiDeleteLoginServerAdminRequest, options?: any) {
        return LoginServerAdminApiFp(this.configuration).deleteLoginServerAdmin(requestParameters.id, options).then((request) => request(this.axios, this.basePath));
    }
    public getLoginServerAdmin(requestParameters: LoginServerAdminApiGetLoginServerAdminRequest, options?: any) {
        return LoginServerAdminApiFp(this.configuration).getLoginServerAdmin(requestParameters.id, requestParameters.includes, requestParameters.select, options).then((request) => request(this.axios, this.basePath));
    }
    public getLoginServerAdminsBulk(requestParameters: LoginServerAdminApiGetLoginServerAdminsBulkRequest, options?: any) {
        return LoginServerAdminApiFp(this.configuration).getLoginServerAdminsBulk(requestParameters.body, options).then((request) => request(this.axios, this.basePath));
    }
    public listLoginServerAdmins(requestParameters: LoginServerAdminApiListLoginServerAdminsRequest = {}, options?: any) {
        return LoginServerAdminApiFp(this.configuration).listLoginServerAdmins(requestParameters.includes, requestParameters.where, requestParameters.whereOr, requestParameters.groupBy, requestParameters.limit, requestParameters.page, requestParameters.orderBy, requestParameters.orderDirection, requestParameters.select, options).then((request) => request(this.axios, this.basePath));
    }
    public updateLoginServerAdmin(requestParameters: LoginServerAdminApiUpdateLoginServerAdminRequest, options?: any) {
        return LoginServerAdminApiFp(this.configuration).updateLoginServerAdmin(requestParameters.id, requestParameters.loginServerAdmin, options).then((request) => request(this.axios, this.basePath));
    }
}
