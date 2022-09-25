import globalAxios, { AxiosPromise, AxiosInstance } from 'axios';
import { Configuration } from '../configuration';
// Some imports not used depending on template conditions
// @ts-ignore
import { BASE_PATH, COLLECTION_FORMATS, RequestArgs, BaseAPI, RequiredError } from '../base';
// @ts-ignore
import { CrudcontrollersBulkFetchByIdsGetRequest } from '../models';
// @ts-ignore
import { ModelsCharacterPotionbelt } from '../models';
export const CharacterPotionbeltApiAxiosParamCreator = function (configuration?: Configuration) {
    return {
        createCharacterPotionbelt: async (characterPotionbelt: ModelsCharacterPotionbelt, options: any = {}): Promise<RequestArgs> => {
            if (characterPotionbelt === null || characterPotionbelt === undefined) {
                throw new RequiredError('characterPotionbelt','Required parameter characterPotionbelt was null or undefined when calling createCharacterPotionbelt.');
            }
            const localVarPath = `/character_potionbelt`;
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
            const nonString = typeof characterPotionbelt !== 'string';
            const needsSerialization = nonString && configuration && configuration.isJsonMime
                ? configuration.isJsonMime(localVarRequestOptions.headers['Content-Type'])
                : nonString;
            localVarRequestOptions.data =  needsSerialization
                ? JSON.stringify(characterPotionbelt !== undefined ? characterPotionbelt : {})
                : (characterPotionbelt || "");
            return {
                url: localVarUrlObj.pathname + localVarUrlObj.search + localVarUrlObj.hash,
                options: localVarRequestOptions,
            };
        },
        deleteCharacterPotionbelt: async (id: number, options: any = {}): Promise<RequestArgs> => {
            if (id === null || id === undefined) {
                throw new RequiredError('id','Required parameter id was null or undefined when calling deleteCharacterPotionbelt.');
            }
            const localVarPath = `/character_potionbelt/{id}`
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
        getCharacterPotionbelt: async (id: number, includes?: string, select?: string, options: any = {}): Promise<RequestArgs> => {
            if (id === null || id === undefined) {
                throw new RequiredError('id','Required parameter id was null or undefined when calling getCharacterPotionbelt.');
            }
            const localVarPath = `/character_potionbelt/{id}`
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
        getCharacterPotionbeltsBulk: async (body: CrudcontrollersBulkFetchByIdsGetRequest, options: any = {}): Promise<RequestArgs> => {
            if (body === null || body === undefined) {
                throw new RequiredError('body','Required parameter body was null or undefined when calling getCharacterPotionbeltsBulk.');
            }
            const localVarPath = `/character_potionbelts/bulk`;
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
        listCharacterPotionbelts: async (includes?: string, where?: string, whereOr?: string, groupBy?: string, limit?: string, page?: number, orderBy?: string, orderDirection?: string, select?: string, options: any = {}): Promise<RequestArgs> => {
            const localVarPath = `/character_potionbelts`;
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
        updateCharacterPotionbelt: async (id: number, characterPotionbelt: ModelsCharacterPotionbelt, options: any = {}): Promise<RequestArgs> => {
            if (id === null || id === undefined) {
                throw new RequiredError('id','Required parameter id was null or undefined when calling updateCharacterPotionbelt.');
            }
            if (characterPotionbelt === null || characterPotionbelt === undefined) {
                throw new RequiredError('characterPotionbelt','Required parameter characterPotionbelt was null or undefined when calling updateCharacterPotionbelt.');
            }
            const localVarPath = `/character_potionbelt/{id}`
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
            const nonString = typeof characterPotionbelt !== 'string';
            const needsSerialization = nonString && configuration && configuration.isJsonMime
                ? configuration.isJsonMime(localVarRequestOptions.headers['Content-Type'])
                : nonString;
            localVarRequestOptions.data =  needsSerialization
                ? JSON.stringify(characterPotionbelt !== undefined ? characterPotionbelt : {})
                : (characterPotionbelt || "");
            return {
                url: localVarUrlObj.pathname + localVarUrlObj.search + localVarUrlObj.hash,
                options: localVarRequestOptions,
            };
        },
    }
};
export const CharacterPotionbeltApiFp = function(configuration?: Configuration) {
    return {
        async createCharacterPotionbelt(characterPotionbelt: ModelsCharacterPotionbelt, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsCharacterPotionbelt>>> {
            const localVarAxiosArgs = await CharacterPotionbeltApiAxiosParamCreator(configuration).createCharacterPotionbelt(characterPotionbelt, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        async deleteCharacterPotionbelt(id: number, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<string>> {
            const localVarAxiosArgs = await CharacterPotionbeltApiAxiosParamCreator(configuration).deleteCharacterPotionbelt(id, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        async getCharacterPotionbelt(id: number, includes?: string, select?: string, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsCharacterPotionbelt>>> {
            const localVarAxiosArgs = await CharacterPotionbeltApiAxiosParamCreator(configuration).getCharacterPotionbelt(id, includes, select, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        async getCharacterPotionbeltsBulk(body: CrudcontrollersBulkFetchByIdsGetRequest, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsCharacterPotionbelt>>> {
            const localVarAxiosArgs = await CharacterPotionbeltApiAxiosParamCreator(configuration).getCharacterPotionbeltsBulk(body, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        async listCharacterPotionbelts(includes?: string, where?: string, whereOr?: string, groupBy?: string, limit?: string, page?: number, orderBy?: string, orderDirection?: string, select?: string, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsCharacterPotionbelt>>> {
            const localVarAxiosArgs = await CharacterPotionbeltApiAxiosParamCreator(configuration).listCharacterPotionbelts(includes, where, whereOr, groupBy, limit, page, orderBy, orderDirection, select, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        async updateCharacterPotionbelt(id: number, characterPotionbelt: ModelsCharacterPotionbelt, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsCharacterPotionbelt>>> {
            const localVarAxiosArgs = await CharacterPotionbeltApiAxiosParamCreator(configuration).updateCharacterPotionbelt(id, characterPotionbelt, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
    }
};
export const CharacterPotionbeltApiFactory = function (configuration?: Configuration, basePath?: string, axios?: AxiosInstance) {
    return {
        createCharacterPotionbelt(characterPotionbelt: ModelsCharacterPotionbelt, options?: any): AxiosPromise<Array<ModelsCharacterPotionbelt>> {
            return CharacterPotionbeltApiFp(configuration).createCharacterPotionbelt(characterPotionbelt, options).then((request) => request(axios, basePath));
        },
        deleteCharacterPotionbelt(id: number, options?: any): AxiosPromise<string> {
            return CharacterPotionbeltApiFp(configuration).deleteCharacterPotionbelt(id, options).then((request) => request(axios, basePath));
        },
        getCharacterPotionbelt(id: number, includes?: string, select?: string, options?: any): AxiosPromise<Array<ModelsCharacterPotionbelt>> {
            return CharacterPotionbeltApiFp(configuration).getCharacterPotionbelt(id, includes, select, options).then((request) => request(axios, basePath));
        },
        getCharacterPotionbeltsBulk(body: CrudcontrollersBulkFetchByIdsGetRequest, options?: any): AxiosPromise<Array<ModelsCharacterPotionbelt>> {
            return CharacterPotionbeltApiFp(configuration).getCharacterPotionbeltsBulk(body, options).then((request) => request(axios, basePath));
        },
        listCharacterPotionbelts(includes?: string, where?: string, whereOr?: string, groupBy?: string, limit?: string, page?: number, orderBy?: string, orderDirection?: string, select?: string, options?: any): AxiosPromise<Array<ModelsCharacterPotionbelt>> {
            return CharacterPotionbeltApiFp(configuration).listCharacterPotionbelts(includes, where, whereOr, groupBy, limit, page, orderBy, orderDirection, select, options).then((request) => request(axios, basePath));
        },
        updateCharacterPotionbelt(id: number, characterPotionbelt: ModelsCharacterPotionbelt, options?: any): AxiosPromise<Array<ModelsCharacterPotionbelt>> {
            return CharacterPotionbeltApiFp(configuration).updateCharacterPotionbelt(id, characterPotionbelt, options).then((request) => request(axios, basePath));
        },
    };
};
export interface CharacterPotionbeltApiCreateCharacterPotionbeltRequest {
    readonly characterPotionbelt: ModelsCharacterPotionbelt
}
export interface CharacterPotionbeltApiDeleteCharacterPotionbeltRequest {
    readonly id: number
}
export interface CharacterPotionbeltApiGetCharacterPotionbeltRequest {
    readonly id: number
    readonly includes?: string
    readonly select?: string
}
export interface CharacterPotionbeltApiGetCharacterPotionbeltsBulkRequest {
    readonly body: CrudcontrollersBulkFetchByIdsGetRequest
}
export interface CharacterPotionbeltApiListCharacterPotionbeltsRequest {
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
export interface CharacterPotionbeltApiUpdateCharacterPotionbeltRequest {
    readonly id: number
    readonly characterPotionbelt: ModelsCharacterPotionbelt
}
export class CharacterPotionbeltApi extends BaseAPI {
    public createCharacterPotionbelt(requestParameters: CharacterPotionbeltApiCreateCharacterPotionbeltRequest, options?: any) {
        return CharacterPotionbeltApiFp(this.configuration).createCharacterPotionbelt(requestParameters.characterPotionbelt, options).then((request) => request(this.axios, this.basePath));
    }
    public deleteCharacterPotionbelt(requestParameters: CharacterPotionbeltApiDeleteCharacterPotionbeltRequest, options?: any) {
        return CharacterPotionbeltApiFp(this.configuration).deleteCharacterPotionbelt(requestParameters.id, options).then((request) => request(this.axios, this.basePath));
    }
    public getCharacterPotionbelt(requestParameters: CharacterPotionbeltApiGetCharacterPotionbeltRequest, options?: any) {
        return CharacterPotionbeltApiFp(this.configuration).getCharacterPotionbelt(requestParameters.id, requestParameters.includes, requestParameters.select, options).then((request) => request(this.axios, this.basePath));
    }
    public getCharacterPotionbeltsBulk(requestParameters: CharacterPotionbeltApiGetCharacterPotionbeltsBulkRequest, options?: any) {
        return CharacterPotionbeltApiFp(this.configuration).getCharacterPotionbeltsBulk(requestParameters.body, options).then((request) => request(this.axios, this.basePath));
    }
    public listCharacterPotionbelts(requestParameters: CharacterPotionbeltApiListCharacterPotionbeltsRequest = {}, options?: any) {
        return CharacterPotionbeltApiFp(this.configuration).listCharacterPotionbelts(requestParameters.includes, requestParameters.where, requestParameters.whereOr, requestParameters.groupBy, requestParameters.limit, requestParameters.page, requestParameters.orderBy, requestParameters.orderDirection, requestParameters.select, options).then((request) => request(this.axios, this.basePath));
    }
    public updateCharacterPotionbelt(requestParameters: CharacterPotionbeltApiUpdateCharacterPotionbeltRequest, options?: any) {
        return CharacterPotionbeltApiFp(this.configuration).updateCharacterPotionbelt(requestParameters.id, requestParameters.characterPotionbelt, options).then((request) => request(this.axios, this.basePath));
    }
}
