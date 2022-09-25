import globalAxios, { AxiosPromise, AxiosInstance } from 'axios';
import { Configuration } from '../configuration';
// Some imports not used depending on template conditions
// @ts-ignore
import { BASE_PATH, COLLECTION_FORMATS, RequestArgs, BaseAPI, RequiredError } from '../base';
// @ts-ignore
import { CrudcontrollersBulkFetchByIdsGetRequest } from '../models';
// @ts-ignore
import { ModelsCharacterInstanceSafereturn } from '../models';
export const CharacterInstanceSafereturnApiAxiosParamCreator = function (configuration?: Configuration) {
    return {
        createCharacterInstanceSafereturn: async (characterInstanceSafereturn: ModelsCharacterInstanceSafereturn, options: any = {}): Promise<RequestArgs> => {
            if (characterInstanceSafereturn === null || characterInstanceSafereturn === undefined) {
                throw new RequiredError('characterInstanceSafereturn','Required parameter characterInstanceSafereturn was null or undefined when calling createCharacterInstanceSafereturn.');
            }
            const localVarPath = `/character_instance_safereturn`;
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
            const nonString = typeof characterInstanceSafereturn !== 'string';
            const needsSerialization = nonString && configuration && configuration.isJsonMime
                ? configuration.isJsonMime(localVarRequestOptions.headers['Content-Type'])
                : nonString;
            localVarRequestOptions.data =  needsSerialization
                ? JSON.stringify(characterInstanceSafereturn !== undefined ? characterInstanceSafereturn : {})
                : (characterInstanceSafereturn || "");
            return {
                url: localVarUrlObj.pathname + localVarUrlObj.search + localVarUrlObj.hash,
                options: localVarRequestOptions,
            };
        },
        deleteCharacterInstanceSafereturn: async (id: number, options: any = {}): Promise<RequestArgs> => {
            if (id === null || id === undefined) {
                throw new RequiredError('id','Required parameter id was null or undefined when calling deleteCharacterInstanceSafereturn.');
            }
            const localVarPath = `/character_instance_safereturn/{id}`
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
        getCharacterInstanceSafereturn: async (id: number, includes?: string, select?: string, options: any = {}): Promise<RequestArgs> => {
            if (id === null || id === undefined) {
                throw new RequiredError('id','Required parameter id was null or undefined when calling getCharacterInstanceSafereturn.');
            }
            const localVarPath = `/character_instance_safereturn/{id}`
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
        getCharacterInstanceSafereturnsBulk: async (body: CrudcontrollersBulkFetchByIdsGetRequest, options: any = {}): Promise<RequestArgs> => {
            if (body === null || body === undefined) {
                throw new RequiredError('body','Required parameter body was null or undefined when calling getCharacterInstanceSafereturnsBulk.');
            }
            const localVarPath = `/character_instance_safereturns/bulk`;
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
        listCharacterInstanceSafereturns: async (includes?: string, where?: string, whereOr?: string, groupBy?: string, limit?: string, page?: number, orderBy?: string, orderDirection?: string, select?: string, options: any = {}): Promise<RequestArgs> => {
            const localVarPath = `/character_instance_safereturns`;
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
        updateCharacterInstanceSafereturn: async (id: number, characterInstanceSafereturn: ModelsCharacterInstanceSafereturn, options: any = {}): Promise<RequestArgs> => {
            if (id === null || id === undefined) {
                throw new RequiredError('id','Required parameter id was null or undefined when calling updateCharacterInstanceSafereturn.');
            }
            if (characterInstanceSafereturn === null || characterInstanceSafereturn === undefined) {
                throw new RequiredError('characterInstanceSafereturn','Required parameter characterInstanceSafereturn was null or undefined when calling updateCharacterInstanceSafereturn.');
            }
            const localVarPath = `/character_instance_safereturn/{id}`
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
            const nonString = typeof characterInstanceSafereturn !== 'string';
            const needsSerialization = nonString && configuration && configuration.isJsonMime
                ? configuration.isJsonMime(localVarRequestOptions.headers['Content-Type'])
                : nonString;
            localVarRequestOptions.data =  needsSerialization
                ? JSON.stringify(characterInstanceSafereturn !== undefined ? characterInstanceSafereturn : {})
                : (characterInstanceSafereturn || "");
            return {
                url: localVarUrlObj.pathname + localVarUrlObj.search + localVarUrlObj.hash,
                options: localVarRequestOptions,
            };
        },
    }
};
export const CharacterInstanceSafereturnApiFp = function(configuration?: Configuration) {
    return {
        async createCharacterInstanceSafereturn(characterInstanceSafereturn: ModelsCharacterInstanceSafereturn, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsCharacterInstanceSafereturn>>> {
            const localVarAxiosArgs = await CharacterInstanceSafereturnApiAxiosParamCreator(configuration).createCharacterInstanceSafereturn(characterInstanceSafereturn, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        async deleteCharacterInstanceSafereturn(id: number, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<string>> {
            const localVarAxiosArgs = await CharacterInstanceSafereturnApiAxiosParamCreator(configuration).deleteCharacterInstanceSafereturn(id, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        async getCharacterInstanceSafereturn(id: number, includes?: string, select?: string, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsCharacterInstanceSafereturn>>> {
            const localVarAxiosArgs = await CharacterInstanceSafereturnApiAxiosParamCreator(configuration).getCharacterInstanceSafereturn(id, includes, select, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        async getCharacterInstanceSafereturnsBulk(body: CrudcontrollersBulkFetchByIdsGetRequest, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsCharacterInstanceSafereturn>>> {
            const localVarAxiosArgs = await CharacterInstanceSafereturnApiAxiosParamCreator(configuration).getCharacterInstanceSafereturnsBulk(body, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        async listCharacterInstanceSafereturns(includes?: string, where?: string, whereOr?: string, groupBy?: string, limit?: string, page?: number, orderBy?: string, orderDirection?: string, select?: string, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsCharacterInstanceSafereturn>>> {
            const localVarAxiosArgs = await CharacterInstanceSafereturnApiAxiosParamCreator(configuration).listCharacterInstanceSafereturns(includes, where, whereOr, groupBy, limit, page, orderBy, orderDirection, select, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        async updateCharacterInstanceSafereturn(id: number, characterInstanceSafereturn: ModelsCharacterInstanceSafereturn, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsCharacterInstanceSafereturn>>> {
            const localVarAxiosArgs = await CharacterInstanceSafereturnApiAxiosParamCreator(configuration).updateCharacterInstanceSafereturn(id, characterInstanceSafereturn, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
    }
};
export const CharacterInstanceSafereturnApiFactory = function (configuration?: Configuration, basePath?: string, axios?: AxiosInstance) {
    return {
        createCharacterInstanceSafereturn(characterInstanceSafereturn: ModelsCharacterInstanceSafereturn, options?: any): AxiosPromise<Array<ModelsCharacterInstanceSafereturn>> {
            return CharacterInstanceSafereturnApiFp(configuration).createCharacterInstanceSafereturn(characterInstanceSafereturn, options).then((request) => request(axios, basePath));
        },
        deleteCharacterInstanceSafereturn(id: number, options?: any): AxiosPromise<string> {
            return CharacterInstanceSafereturnApiFp(configuration).deleteCharacterInstanceSafereturn(id, options).then((request) => request(axios, basePath));
        },
        getCharacterInstanceSafereturn(id: number, includes?: string, select?: string, options?: any): AxiosPromise<Array<ModelsCharacterInstanceSafereturn>> {
            return CharacterInstanceSafereturnApiFp(configuration).getCharacterInstanceSafereturn(id, includes, select, options).then((request) => request(axios, basePath));
        },
        getCharacterInstanceSafereturnsBulk(body: CrudcontrollersBulkFetchByIdsGetRequest, options?: any): AxiosPromise<Array<ModelsCharacterInstanceSafereturn>> {
            return CharacterInstanceSafereturnApiFp(configuration).getCharacterInstanceSafereturnsBulk(body, options).then((request) => request(axios, basePath));
        },
        listCharacterInstanceSafereturns(includes?: string, where?: string, whereOr?: string, groupBy?: string, limit?: string, page?: number, orderBy?: string, orderDirection?: string, select?: string, options?: any): AxiosPromise<Array<ModelsCharacterInstanceSafereturn>> {
            return CharacterInstanceSafereturnApiFp(configuration).listCharacterInstanceSafereturns(includes, where, whereOr, groupBy, limit, page, orderBy, orderDirection, select, options).then((request) => request(axios, basePath));
        },
        updateCharacterInstanceSafereturn(id: number, characterInstanceSafereturn: ModelsCharacterInstanceSafereturn, options?: any): AxiosPromise<Array<ModelsCharacterInstanceSafereturn>> {
            return CharacterInstanceSafereturnApiFp(configuration).updateCharacterInstanceSafereturn(id, characterInstanceSafereturn, options).then((request) => request(axios, basePath));
        },
    };
};
export interface CharacterInstanceSafereturnApiCreateCharacterInstanceSafereturnRequest {
    readonly characterInstanceSafereturn: ModelsCharacterInstanceSafereturn
}
export interface CharacterInstanceSafereturnApiDeleteCharacterInstanceSafereturnRequest {
    readonly id: number
}
export interface CharacterInstanceSafereturnApiGetCharacterInstanceSafereturnRequest {
    readonly id: number
    readonly includes?: string
    readonly select?: string
}
export interface CharacterInstanceSafereturnApiGetCharacterInstanceSafereturnsBulkRequest {
    readonly body: CrudcontrollersBulkFetchByIdsGetRequest
}
export interface CharacterInstanceSafereturnApiListCharacterInstanceSafereturnsRequest {
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
export interface CharacterInstanceSafereturnApiUpdateCharacterInstanceSafereturnRequest {
    readonly id: number
    readonly characterInstanceSafereturn: ModelsCharacterInstanceSafereturn
}
export class CharacterInstanceSafereturnApi extends BaseAPI {
    public createCharacterInstanceSafereturn(requestParameters: CharacterInstanceSafereturnApiCreateCharacterInstanceSafereturnRequest, options?: any) {
        return CharacterInstanceSafereturnApiFp(this.configuration).createCharacterInstanceSafereturn(requestParameters.characterInstanceSafereturn, options).then((request) => request(this.axios, this.basePath));
    }
    public deleteCharacterInstanceSafereturn(requestParameters: CharacterInstanceSafereturnApiDeleteCharacterInstanceSafereturnRequest, options?: any) {
        return CharacterInstanceSafereturnApiFp(this.configuration).deleteCharacterInstanceSafereturn(requestParameters.id, options).then((request) => request(this.axios, this.basePath));
    }
    public getCharacterInstanceSafereturn(requestParameters: CharacterInstanceSafereturnApiGetCharacterInstanceSafereturnRequest, options?: any) {
        return CharacterInstanceSafereturnApiFp(this.configuration).getCharacterInstanceSafereturn(requestParameters.id, requestParameters.includes, requestParameters.select, options).then((request) => request(this.axios, this.basePath));
    }
    public getCharacterInstanceSafereturnsBulk(requestParameters: CharacterInstanceSafereturnApiGetCharacterInstanceSafereturnsBulkRequest, options?: any) {
        return CharacterInstanceSafereturnApiFp(this.configuration).getCharacterInstanceSafereturnsBulk(requestParameters.body, options).then((request) => request(this.axios, this.basePath));
    }
    public listCharacterInstanceSafereturns(requestParameters: CharacterInstanceSafereturnApiListCharacterInstanceSafereturnsRequest = {}, options?: any) {
        return CharacterInstanceSafereturnApiFp(this.configuration).listCharacterInstanceSafereturns(requestParameters.includes, requestParameters.where, requestParameters.whereOr, requestParameters.groupBy, requestParameters.limit, requestParameters.page, requestParameters.orderBy, requestParameters.orderDirection, requestParameters.select, options).then((request) => request(this.axios, this.basePath));
    }
    public updateCharacterInstanceSafereturn(requestParameters: CharacterInstanceSafereturnApiUpdateCharacterInstanceSafereturnRequest, options?: any) {
        return CharacterInstanceSafereturnApiFp(this.configuration).updateCharacterInstanceSafereturn(requestParameters.id, requestParameters.characterInstanceSafereturn, options).then((request) => request(this.axios, this.basePath));
    }
}
