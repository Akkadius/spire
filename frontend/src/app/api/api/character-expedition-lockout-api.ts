import globalAxios, { AxiosPromise, AxiosInstance } from 'axios';
import { Configuration } from '../configuration';
// Some imports not used depending on template conditions
// @ts-ignore
import { BASE_PATH, COLLECTION_FORMATS, RequestArgs, BaseAPI, RequiredError } from '../base';
// @ts-ignore
import { CrudcontrollersBulkFetchByIdsGetRequest } from '../models';
// @ts-ignore
import { ModelsCharacterExpeditionLockout } from '../models';
export const CharacterExpeditionLockoutApiAxiosParamCreator = function (configuration?: Configuration) {
    return {
        createCharacterExpeditionLockout: async (characterExpeditionLockout: ModelsCharacterExpeditionLockout, options: any = {}): Promise<RequestArgs> => {
            if (characterExpeditionLockout === null || characterExpeditionLockout === undefined) {
                throw new RequiredError('characterExpeditionLockout','Required parameter characterExpeditionLockout was null or undefined when calling createCharacterExpeditionLockout.');
            }
            const localVarPath = `/character_expedition_lockout`;
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
            const nonString = typeof characterExpeditionLockout !== 'string';
            const needsSerialization = nonString && configuration && configuration.isJsonMime
                ? configuration.isJsonMime(localVarRequestOptions.headers['Content-Type'])
                : nonString;
            localVarRequestOptions.data =  needsSerialization
                ? JSON.stringify(characterExpeditionLockout !== undefined ? characterExpeditionLockout : {})
                : (characterExpeditionLockout || "");
            return {
                url: localVarUrlObj.pathname + localVarUrlObj.search + localVarUrlObj.hash,
                options: localVarRequestOptions,
            };
        },
        deleteCharacterExpeditionLockout: async (id: number, options: any = {}): Promise<RequestArgs> => {
            if (id === null || id === undefined) {
                throw new RequiredError('id','Required parameter id was null or undefined when calling deleteCharacterExpeditionLockout.');
            }
            const localVarPath = `/character_expedition_lockout/{id}`
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
        getCharacterExpeditionLockout: async (id: number, includes?: string, select?: string, options: any = {}): Promise<RequestArgs> => {
            if (id === null || id === undefined) {
                throw new RequiredError('id','Required parameter id was null or undefined when calling getCharacterExpeditionLockout.');
            }
            const localVarPath = `/character_expedition_lockout/{id}`
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
        getCharacterExpeditionLockoutsBulk: async (body: CrudcontrollersBulkFetchByIdsGetRequest, options: any = {}): Promise<RequestArgs> => {
            if (body === null || body === undefined) {
                throw new RequiredError('body','Required parameter body was null or undefined when calling getCharacterExpeditionLockoutsBulk.');
            }
            const localVarPath = `/character_expedition_lockouts/bulk`;
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
        listCharacterExpeditionLockouts: async (includes?: string, where?: string, whereOr?: string, groupBy?: string, limit?: string, page?: number, orderBy?: string, orderDirection?: string, select?: string, options: any = {}): Promise<RequestArgs> => {
            const localVarPath = `/character_expedition_lockouts`;
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
        updateCharacterExpeditionLockout: async (id: number, characterExpeditionLockout: ModelsCharacterExpeditionLockout, options: any = {}): Promise<RequestArgs> => {
            if (id === null || id === undefined) {
                throw new RequiredError('id','Required parameter id was null or undefined when calling updateCharacterExpeditionLockout.');
            }
            if (characterExpeditionLockout === null || characterExpeditionLockout === undefined) {
                throw new RequiredError('characterExpeditionLockout','Required parameter characterExpeditionLockout was null or undefined when calling updateCharacterExpeditionLockout.');
            }
            const localVarPath = `/character_expedition_lockout/{id}`
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
            const nonString = typeof characterExpeditionLockout !== 'string';
            const needsSerialization = nonString && configuration && configuration.isJsonMime
                ? configuration.isJsonMime(localVarRequestOptions.headers['Content-Type'])
                : nonString;
            localVarRequestOptions.data =  needsSerialization
                ? JSON.stringify(characterExpeditionLockout !== undefined ? characterExpeditionLockout : {})
                : (characterExpeditionLockout || "");
            return {
                url: localVarUrlObj.pathname + localVarUrlObj.search + localVarUrlObj.hash,
                options: localVarRequestOptions,
            };
        },
    }
};
export const CharacterExpeditionLockoutApiFp = function(configuration?: Configuration) {
    return {
        async createCharacterExpeditionLockout(characterExpeditionLockout: ModelsCharacterExpeditionLockout, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsCharacterExpeditionLockout>>> {
            const localVarAxiosArgs = await CharacterExpeditionLockoutApiAxiosParamCreator(configuration).createCharacterExpeditionLockout(characterExpeditionLockout, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        async deleteCharacterExpeditionLockout(id: number, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<string>> {
            const localVarAxiosArgs = await CharacterExpeditionLockoutApiAxiosParamCreator(configuration).deleteCharacterExpeditionLockout(id, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        async getCharacterExpeditionLockout(id: number, includes?: string, select?: string, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsCharacterExpeditionLockout>>> {
            const localVarAxiosArgs = await CharacterExpeditionLockoutApiAxiosParamCreator(configuration).getCharacterExpeditionLockout(id, includes, select, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        async getCharacterExpeditionLockoutsBulk(body: CrudcontrollersBulkFetchByIdsGetRequest, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsCharacterExpeditionLockout>>> {
            const localVarAxiosArgs = await CharacterExpeditionLockoutApiAxiosParamCreator(configuration).getCharacterExpeditionLockoutsBulk(body, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        async listCharacterExpeditionLockouts(includes?: string, where?: string, whereOr?: string, groupBy?: string, limit?: string, page?: number, orderBy?: string, orderDirection?: string, select?: string, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsCharacterExpeditionLockout>>> {
            const localVarAxiosArgs = await CharacterExpeditionLockoutApiAxiosParamCreator(configuration).listCharacterExpeditionLockouts(includes, where, whereOr, groupBy, limit, page, orderBy, orderDirection, select, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        async updateCharacterExpeditionLockout(id: number, characterExpeditionLockout: ModelsCharacterExpeditionLockout, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsCharacterExpeditionLockout>>> {
            const localVarAxiosArgs = await CharacterExpeditionLockoutApiAxiosParamCreator(configuration).updateCharacterExpeditionLockout(id, characterExpeditionLockout, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
    }
};
export const CharacterExpeditionLockoutApiFactory = function (configuration?: Configuration, basePath?: string, axios?: AxiosInstance) {
    return {
        createCharacterExpeditionLockout(characterExpeditionLockout: ModelsCharacterExpeditionLockout, options?: any): AxiosPromise<Array<ModelsCharacterExpeditionLockout>> {
            return CharacterExpeditionLockoutApiFp(configuration).createCharacterExpeditionLockout(characterExpeditionLockout, options).then((request) => request(axios, basePath));
        },
        deleteCharacterExpeditionLockout(id: number, options?: any): AxiosPromise<string> {
            return CharacterExpeditionLockoutApiFp(configuration).deleteCharacterExpeditionLockout(id, options).then((request) => request(axios, basePath));
        },
        getCharacterExpeditionLockout(id: number, includes?: string, select?: string, options?: any): AxiosPromise<Array<ModelsCharacterExpeditionLockout>> {
            return CharacterExpeditionLockoutApiFp(configuration).getCharacterExpeditionLockout(id, includes, select, options).then((request) => request(axios, basePath));
        },
        getCharacterExpeditionLockoutsBulk(body: CrudcontrollersBulkFetchByIdsGetRequest, options?: any): AxiosPromise<Array<ModelsCharacterExpeditionLockout>> {
            return CharacterExpeditionLockoutApiFp(configuration).getCharacterExpeditionLockoutsBulk(body, options).then((request) => request(axios, basePath));
        },
        listCharacterExpeditionLockouts(includes?: string, where?: string, whereOr?: string, groupBy?: string, limit?: string, page?: number, orderBy?: string, orderDirection?: string, select?: string, options?: any): AxiosPromise<Array<ModelsCharacterExpeditionLockout>> {
            return CharacterExpeditionLockoutApiFp(configuration).listCharacterExpeditionLockouts(includes, where, whereOr, groupBy, limit, page, orderBy, orderDirection, select, options).then((request) => request(axios, basePath));
        },
        updateCharacterExpeditionLockout(id: number, characterExpeditionLockout: ModelsCharacterExpeditionLockout, options?: any): AxiosPromise<Array<ModelsCharacterExpeditionLockout>> {
            return CharacterExpeditionLockoutApiFp(configuration).updateCharacterExpeditionLockout(id, characterExpeditionLockout, options).then((request) => request(axios, basePath));
        },
    };
};
export interface CharacterExpeditionLockoutApiCreateCharacterExpeditionLockoutRequest {
    readonly characterExpeditionLockout: ModelsCharacterExpeditionLockout
}
export interface CharacterExpeditionLockoutApiDeleteCharacterExpeditionLockoutRequest {
    readonly id: number
}
export interface CharacterExpeditionLockoutApiGetCharacterExpeditionLockoutRequest {
    readonly id: number
    readonly includes?: string
    readonly select?: string
}
export interface CharacterExpeditionLockoutApiGetCharacterExpeditionLockoutsBulkRequest {
    readonly body: CrudcontrollersBulkFetchByIdsGetRequest
}
export interface CharacterExpeditionLockoutApiListCharacterExpeditionLockoutsRequest {
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
export interface CharacterExpeditionLockoutApiUpdateCharacterExpeditionLockoutRequest {
    readonly id: number
    readonly characterExpeditionLockout: ModelsCharacterExpeditionLockout
}
export class CharacterExpeditionLockoutApi extends BaseAPI {
    public createCharacterExpeditionLockout(requestParameters: CharacterExpeditionLockoutApiCreateCharacterExpeditionLockoutRequest, options?: any) {
        return CharacterExpeditionLockoutApiFp(this.configuration).createCharacterExpeditionLockout(requestParameters.characterExpeditionLockout, options).then((request) => request(this.axios, this.basePath));
    }
    public deleteCharacterExpeditionLockout(requestParameters: CharacterExpeditionLockoutApiDeleteCharacterExpeditionLockoutRequest, options?: any) {
        return CharacterExpeditionLockoutApiFp(this.configuration).deleteCharacterExpeditionLockout(requestParameters.id, options).then((request) => request(this.axios, this.basePath));
    }
    public getCharacterExpeditionLockout(requestParameters: CharacterExpeditionLockoutApiGetCharacterExpeditionLockoutRequest, options?: any) {
        return CharacterExpeditionLockoutApiFp(this.configuration).getCharacterExpeditionLockout(requestParameters.id, requestParameters.includes, requestParameters.select, options).then((request) => request(this.axios, this.basePath));
    }
    public getCharacterExpeditionLockoutsBulk(requestParameters: CharacterExpeditionLockoutApiGetCharacterExpeditionLockoutsBulkRequest, options?: any) {
        return CharacterExpeditionLockoutApiFp(this.configuration).getCharacterExpeditionLockoutsBulk(requestParameters.body, options).then((request) => request(this.axios, this.basePath));
    }
    public listCharacterExpeditionLockouts(requestParameters: CharacterExpeditionLockoutApiListCharacterExpeditionLockoutsRequest = {}, options?: any) {
        return CharacterExpeditionLockoutApiFp(this.configuration).listCharacterExpeditionLockouts(requestParameters.includes, requestParameters.where, requestParameters.whereOr, requestParameters.groupBy, requestParameters.limit, requestParameters.page, requestParameters.orderBy, requestParameters.orderDirection, requestParameters.select, options).then((request) => request(this.axios, this.basePath));
    }
    public updateCharacterExpeditionLockout(requestParameters: CharacterExpeditionLockoutApiUpdateCharacterExpeditionLockoutRequest, options?: any) {
        return CharacterExpeditionLockoutApiFp(this.configuration).updateCharacterExpeditionLockout(requestParameters.id, requestParameters.characterExpeditionLockout, options).then((request) => request(this.axios, this.basePath));
    }
}
