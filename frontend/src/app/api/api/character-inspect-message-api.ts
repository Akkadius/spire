import globalAxios, { AxiosPromise, AxiosInstance } from 'axios';
import { Configuration } from '../configuration';
// Some imports not used depending on template conditions
// @ts-ignore
import { BASE_PATH, COLLECTION_FORMATS, RequestArgs, BaseAPI, RequiredError } from '../base';
// @ts-ignore
import { CrudcontrollersBulkFetchByIdsGetRequest } from '../models';
// @ts-ignore
import { ModelsCharacterInspectMessage } from '../models';
export const CharacterInspectMessageApiAxiosParamCreator = function (configuration?: Configuration) {
    return {
        createCharacterInspectMessage: async (characterInspectMessage: ModelsCharacterInspectMessage, options: any = {}): Promise<RequestArgs> => {
            if (characterInspectMessage === null || characterInspectMessage === undefined) {
                throw new RequiredError('characterInspectMessage','Required parameter characterInspectMessage was null or undefined when calling createCharacterInspectMessage.');
            }
            const localVarPath = `/character_inspect_message`;
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
            const nonString = typeof characterInspectMessage !== 'string';
            const needsSerialization = nonString && configuration && configuration.isJsonMime
                ? configuration.isJsonMime(localVarRequestOptions.headers['Content-Type'])
                : nonString;
            localVarRequestOptions.data =  needsSerialization
                ? JSON.stringify(characterInspectMessage !== undefined ? characterInspectMessage : {})
                : (characterInspectMessage || "");
            return {
                url: localVarUrlObj.pathname + localVarUrlObj.search + localVarUrlObj.hash,
                options: localVarRequestOptions,
            };
        },
        deleteCharacterInspectMessage: async (id: number, options: any = {}): Promise<RequestArgs> => {
            if (id === null || id === undefined) {
                throw new RequiredError('id','Required parameter id was null or undefined when calling deleteCharacterInspectMessage.');
            }
            const localVarPath = `/character_inspect_message/{id}`
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
        getCharacterInspectMessage: async (id: number, includes?: string, select?: string, options: any = {}): Promise<RequestArgs> => {
            if (id === null || id === undefined) {
                throw new RequiredError('id','Required parameter id was null or undefined when calling getCharacterInspectMessage.');
            }
            const localVarPath = `/character_inspect_message/{id}`
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
        getCharacterInspectMessagesBulk: async (body: CrudcontrollersBulkFetchByIdsGetRequest, options: any = {}): Promise<RequestArgs> => {
            if (body === null || body === undefined) {
                throw new RequiredError('body','Required parameter body was null or undefined when calling getCharacterInspectMessagesBulk.');
            }
            const localVarPath = `/character_inspect_messages/bulk`;
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
        listCharacterInspectMessages: async (includes?: string, where?: string, whereOr?: string, groupBy?: string, limit?: string, page?: number, orderBy?: string, orderDirection?: string, select?: string, options: any = {}): Promise<RequestArgs> => {
            const localVarPath = `/character_inspect_messages`;
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
        updateCharacterInspectMessage: async (id: number, characterInspectMessage: ModelsCharacterInspectMessage, options: any = {}): Promise<RequestArgs> => {
            if (id === null || id === undefined) {
                throw new RequiredError('id','Required parameter id was null or undefined when calling updateCharacterInspectMessage.');
            }
            if (characterInspectMessage === null || characterInspectMessage === undefined) {
                throw new RequiredError('characterInspectMessage','Required parameter characterInspectMessage was null or undefined when calling updateCharacterInspectMessage.');
            }
            const localVarPath = `/character_inspect_message/{id}`
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
            const nonString = typeof characterInspectMessage !== 'string';
            const needsSerialization = nonString && configuration && configuration.isJsonMime
                ? configuration.isJsonMime(localVarRequestOptions.headers['Content-Type'])
                : nonString;
            localVarRequestOptions.data =  needsSerialization
                ? JSON.stringify(characterInspectMessage !== undefined ? characterInspectMessage : {})
                : (characterInspectMessage || "");
            return {
                url: localVarUrlObj.pathname + localVarUrlObj.search + localVarUrlObj.hash,
                options: localVarRequestOptions,
            };
        },
    }
};
export const CharacterInspectMessageApiFp = function(configuration?: Configuration) {
    return {
        async createCharacterInspectMessage(characterInspectMessage: ModelsCharacterInspectMessage, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsCharacterInspectMessage>>> {
            const localVarAxiosArgs = await CharacterInspectMessageApiAxiosParamCreator(configuration).createCharacterInspectMessage(characterInspectMessage, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        async deleteCharacterInspectMessage(id: number, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<string>> {
            const localVarAxiosArgs = await CharacterInspectMessageApiAxiosParamCreator(configuration).deleteCharacterInspectMessage(id, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        async getCharacterInspectMessage(id: number, includes?: string, select?: string, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsCharacterInspectMessage>>> {
            const localVarAxiosArgs = await CharacterInspectMessageApiAxiosParamCreator(configuration).getCharacterInspectMessage(id, includes, select, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        async getCharacterInspectMessagesBulk(body: CrudcontrollersBulkFetchByIdsGetRequest, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsCharacterInspectMessage>>> {
            const localVarAxiosArgs = await CharacterInspectMessageApiAxiosParamCreator(configuration).getCharacterInspectMessagesBulk(body, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        async listCharacterInspectMessages(includes?: string, where?: string, whereOr?: string, groupBy?: string, limit?: string, page?: number, orderBy?: string, orderDirection?: string, select?: string, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsCharacterInspectMessage>>> {
            const localVarAxiosArgs = await CharacterInspectMessageApiAxiosParamCreator(configuration).listCharacterInspectMessages(includes, where, whereOr, groupBy, limit, page, orderBy, orderDirection, select, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        async updateCharacterInspectMessage(id: number, characterInspectMessage: ModelsCharacterInspectMessage, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsCharacterInspectMessage>>> {
            const localVarAxiosArgs = await CharacterInspectMessageApiAxiosParamCreator(configuration).updateCharacterInspectMessage(id, characterInspectMessage, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
    }
};
export const CharacterInspectMessageApiFactory = function (configuration?: Configuration, basePath?: string, axios?: AxiosInstance) {
    return {
        createCharacterInspectMessage(characterInspectMessage: ModelsCharacterInspectMessage, options?: any): AxiosPromise<Array<ModelsCharacterInspectMessage>> {
            return CharacterInspectMessageApiFp(configuration).createCharacterInspectMessage(characterInspectMessage, options).then((request) => request(axios, basePath));
        },
        deleteCharacterInspectMessage(id: number, options?: any): AxiosPromise<string> {
            return CharacterInspectMessageApiFp(configuration).deleteCharacterInspectMessage(id, options).then((request) => request(axios, basePath));
        },
        getCharacterInspectMessage(id: number, includes?: string, select?: string, options?: any): AxiosPromise<Array<ModelsCharacterInspectMessage>> {
            return CharacterInspectMessageApiFp(configuration).getCharacterInspectMessage(id, includes, select, options).then((request) => request(axios, basePath));
        },
        getCharacterInspectMessagesBulk(body: CrudcontrollersBulkFetchByIdsGetRequest, options?: any): AxiosPromise<Array<ModelsCharacterInspectMessage>> {
            return CharacterInspectMessageApiFp(configuration).getCharacterInspectMessagesBulk(body, options).then((request) => request(axios, basePath));
        },
        listCharacterInspectMessages(includes?: string, where?: string, whereOr?: string, groupBy?: string, limit?: string, page?: number, orderBy?: string, orderDirection?: string, select?: string, options?: any): AxiosPromise<Array<ModelsCharacterInspectMessage>> {
            return CharacterInspectMessageApiFp(configuration).listCharacterInspectMessages(includes, where, whereOr, groupBy, limit, page, orderBy, orderDirection, select, options).then((request) => request(axios, basePath));
        },
        updateCharacterInspectMessage(id: number, characterInspectMessage: ModelsCharacterInspectMessage, options?: any): AxiosPromise<Array<ModelsCharacterInspectMessage>> {
            return CharacterInspectMessageApiFp(configuration).updateCharacterInspectMessage(id, characterInspectMessage, options).then((request) => request(axios, basePath));
        },
    };
};
export interface CharacterInspectMessageApiCreateCharacterInspectMessageRequest {
    readonly characterInspectMessage: ModelsCharacterInspectMessage
}
export interface CharacterInspectMessageApiDeleteCharacterInspectMessageRequest {
    readonly id: number
}
export interface CharacterInspectMessageApiGetCharacterInspectMessageRequest {
    readonly id: number
    readonly includes?: string
    readonly select?: string
}
export interface CharacterInspectMessageApiGetCharacterInspectMessagesBulkRequest {
    readonly body: CrudcontrollersBulkFetchByIdsGetRequest
}
export interface CharacterInspectMessageApiListCharacterInspectMessagesRequest {
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
export interface CharacterInspectMessageApiUpdateCharacterInspectMessageRequest {
    readonly id: number
    readonly characterInspectMessage: ModelsCharacterInspectMessage
}
export class CharacterInspectMessageApi extends BaseAPI {
    public createCharacterInspectMessage(requestParameters: CharacterInspectMessageApiCreateCharacterInspectMessageRequest, options?: any) {
        return CharacterInspectMessageApiFp(this.configuration).createCharacterInspectMessage(requestParameters.characterInspectMessage, options).then((request) => request(this.axios, this.basePath));
    }
    public deleteCharacterInspectMessage(requestParameters: CharacterInspectMessageApiDeleteCharacterInspectMessageRequest, options?: any) {
        return CharacterInspectMessageApiFp(this.configuration).deleteCharacterInspectMessage(requestParameters.id, options).then((request) => request(this.axios, this.basePath));
    }
    public getCharacterInspectMessage(requestParameters: CharacterInspectMessageApiGetCharacterInspectMessageRequest, options?: any) {
        return CharacterInspectMessageApiFp(this.configuration).getCharacterInspectMessage(requestParameters.id, requestParameters.includes, requestParameters.select, options).then((request) => request(this.axios, this.basePath));
    }
    public getCharacterInspectMessagesBulk(requestParameters: CharacterInspectMessageApiGetCharacterInspectMessagesBulkRequest, options?: any) {
        return CharacterInspectMessageApiFp(this.configuration).getCharacterInspectMessagesBulk(requestParameters.body, options).then((request) => request(this.axios, this.basePath));
    }
    public listCharacterInspectMessages(requestParameters: CharacterInspectMessageApiListCharacterInspectMessagesRequest = {}, options?: any) {
        return CharacterInspectMessageApiFp(this.configuration).listCharacterInspectMessages(requestParameters.includes, requestParameters.where, requestParameters.whereOr, requestParameters.groupBy, requestParameters.limit, requestParameters.page, requestParameters.orderBy, requestParameters.orderDirection, requestParameters.select, options).then((request) => request(this.axios, this.basePath));
    }
    public updateCharacterInspectMessage(requestParameters: CharacterInspectMessageApiUpdateCharacterInspectMessageRequest, options?: any) {
        return CharacterInspectMessageApiFp(this.configuration).updateCharacterInspectMessage(requestParameters.id, requestParameters.characterInspectMessage, options).then((request) => request(this.axios, this.basePath));
    }
}
