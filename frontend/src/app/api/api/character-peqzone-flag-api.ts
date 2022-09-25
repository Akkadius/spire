import globalAxios, { AxiosPromise, AxiosInstance } from 'axios';
import { Configuration } from '../configuration';
// Some imports not used depending on template conditions
// @ts-ignore
import { BASE_PATH, COLLECTION_FORMATS, RequestArgs, BaseAPI, RequiredError } from '../base';
// @ts-ignore
import { CrudcontrollersBulkFetchByIdsGetRequest } from '../models';
// @ts-ignore
import { ModelsCharacterPeqzoneFlag } from '../models';
export const CharacterPeqzoneFlagApiAxiosParamCreator = function (configuration?: Configuration) {
    return {
        createCharacterPeqzoneFlag: async (characterPeqzoneFlag: ModelsCharacterPeqzoneFlag, options: any = {}): Promise<RequestArgs> => {
            if (characterPeqzoneFlag === null || characterPeqzoneFlag === undefined) {
                throw new RequiredError('characterPeqzoneFlag','Required parameter characterPeqzoneFlag was null or undefined when calling createCharacterPeqzoneFlag.');
            }
            const localVarPath = `/character_peqzone_flag`;
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
            const nonString = typeof characterPeqzoneFlag !== 'string';
            const needsSerialization = nonString && configuration && configuration.isJsonMime
                ? configuration.isJsonMime(localVarRequestOptions.headers['Content-Type'])
                : nonString;
            localVarRequestOptions.data =  needsSerialization
                ? JSON.stringify(characterPeqzoneFlag !== undefined ? characterPeqzoneFlag : {})
                : (characterPeqzoneFlag || "");
            return {
                url: localVarUrlObj.pathname + localVarUrlObj.search + localVarUrlObj.hash,
                options: localVarRequestOptions,
            };
        },
        deleteCharacterPeqzoneFlag: async (id: number, options: any = {}): Promise<RequestArgs> => {
            if (id === null || id === undefined) {
                throw new RequiredError('id','Required parameter id was null or undefined when calling deleteCharacterPeqzoneFlag.');
            }
            const localVarPath = `/character_peqzone_flag/{id}`
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
        getCharacterPeqzoneFlag: async (id: number, includes?: string, select?: string, options: any = {}): Promise<RequestArgs> => {
            if (id === null || id === undefined) {
                throw new RequiredError('id','Required parameter id was null or undefined when calling getCharacterPeqzoneFlag.');
            }
            const localVarPath = `/character_peqzone_flag/{id}`
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
        getCharacterPeqzoneFlagsBulk: async (body: CrudcontrollersBulkFetchByIdsGetRequest, options: any = {}): Promise<RequestArgs> => {
            if (body === null || body === undefined) {
                throw new RequiredError('body','Required parameter body was null or undefined when calling getCharacterPeqzoneFlagsBulk.');
            }
            const localVarPath = `/character_peqzone_flags/bulk`;
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
        listCharacterPeqzoneFlags: async (includes?: string, where?: string, whereOr?: string, groupBy?: string, limit?: string, page?: number, orderBy?: string, orderDirection?: string, select?: string, options: any = {}): Promise<RequestArgs> => {
            const localVarPath = `/character_peqzone_flags`;
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
        updateCharacterPeqzoneFlag: async (id: number, characterPeqzoneFlag: ModelsCharacterPeqzoneFlag, options: any = {}): Promise<RequestArgs> => {
            if (id === null || id === undefined) {
                throw new RequiredError('id','Required parameter id was null or undefined when calling updateCharacterPeqzoneFlag.');
            }
            if (characterPeqzoneFlag === null || characterPeqzoneFlag === undefined) {
                throw new RequiredError('characterPeqzoneFlag','Required parameter characterPeqzoneFlag was null or undefined when calling updateCharacterPeqzoneFlag.');
            }
            const localVarPath = `/character_peqzone_flag/{id}`
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
            const nonString = typeof characterPeqzoneFlag !== 'string';
            const needsSerialization = nonString && configuration && configuration.isJsonMime
                ? configuration.isJsonMime(localVarRequestOptions.headers['Content-Type'])
                : nonString;
            localVarRequestOptions.data =  needsSerialization
                ? JSON.stringify(characterPeqzoneFlag !== undefined ? characterPeqzoneFlag : {})
                : (characterPeqzoneFlag || "");
            return {
                url: localVarUrlObj.pathname + localVarUrlObj.search + localVarUrlObj.hash,
                options: localVarRequestOptions,
            };
        },
    }
};
export const CharacterPeqzoneFlagApiFp = function(configuration?: Configuration) {
    return {
        async createCharacterPeqzoneFlag(characterPeqzoneFlag: ModelsCharacterPeqzoneFlag, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsCharacterPeqzoneFlag>>> {
            const localVarAxiosArgs = await CharacterPeqzoneFlagApiAxiosParamCreator(configuration).createCharacterPeqzoneFlag(characterPeqzoneFlag, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        async deleteCharacterPeqzoneFlag(id: number, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<string>> {
            const localVarAxiosArgs = await CharacterPeqzoneFlagApiAxiosParamCreator(configuration).deleteCharacterPeqzoneFlag(id, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        async getCharacterPeqzoneFlag(id: number, includes?: string, select?: string, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsCharacterPeqzoneFlag>>> {
            const localVarAxiosArgs = await CharacterPeqzoneFlagApiAxiosParamCreator(configuration).getCharacterPeqzoneFlag(id, includes, select, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        async getCharacterPeqzoneFlagsBulk(body: CrudcontrollersBulkFetchByIdsGetRequest, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsCharacterPeqzoneFlag>>> {
            const localVarAxiosArgs = await CharacterPeqzoneFlagApiAxiosParamCreator(configuration).getCharacterPeqzoneFlagsBulk(body, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        async listCharacterPeqzoneFlags(includes?: string, where?: string, whereOr?: string, groupBy?: string, limit?: string, page?: number, orderBy?: string, orderDirection?: string, select?: string, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsCharacterPeqzoneFlag>>> {
            const localVarAxiosArgs = await CharacterPeqzoneFlagApiAxiosParamCreator(configuration).listCharacterPeqzoneFlags(includes, where, whereOr, groupBy, limit, page, orderBy, orderDirection, select, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        async updateCharacterPeqzoneFlag(id: number, characterPeqzoneFlag: ModelsCharacterPeqzoneFlag, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsCharacterPeqzoneFlag>>> {
            const localVarAxiosArgs = await CharacterPeqzoneFlagApiAxiosParamCreator(configuration).updateCharacterPeqzoneFlag(id, characterPeqzoneFlag, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
    }
};
export const CharacterPeqzoneFlagApiFactory = function (configuration?: Configuration, basePath?: string, axios?: AxiosInstance) {
    return {
        createCharacterPeqzoneFlag(characterPeqzoneFlag: ModelsCharacterPeqzoneFlag, options?: any): AxiosPromise<Array<ModelsCharacterPeqzoneFlag>> {
            return CharacterPeqzoneFlagApiFp(configuration).createCharacterPeqzoneFlag(characterPeqzoneFlag, options).then((request) => request(axios, basePath));
        },
        deleteCharacterPeqzoneFlag(id: number, options?: any): AxiosPromise<string> {
            return CharacterPeqzoneFlagApiFp(configuration).deleteCharacterPeqzoneFlag(id, options).then((request) => request(axios, basePath));
        },
        getCharacterPeqzoneFlag(id: number, includes?: string, select?: string, options?: any): AxiosPromise<Array<ModelsCharacterPeqzoneFlag>> {
            return CharacterPeqzoneFlagApiFp(configuration).getCharacterPeqzoneFlag(id, includes, select, options).then((request) => request(axios, basePath));
        },
        getCharacterPeqzoneFlagsBulk(body: CrudcontrollersBulkFetchByIdsGetRequest, options?: any): AxiosPromise<Array<ModelsCharacterPeqzoneFlag>> {
            return CharacterPeqzoneFlagApiFp(configuration).getCharacterPeqzoneFlagsBulk(body, options).then((request) => request(axios, basePath));
        },
        listCharacterPeqzoneFlags(includes?: string, where?: string, whereOr?: string, groupBy?: string, limit?: string, page?: number, orderBy?: string, orderDirection?: string, select?: string, options?: any): AxiosPromise<Array<ModelsCharacterPeqzoneFlag>> {
            return CharacterPeqzoneFlagApiFp(configuration).listCharacterPeqzoneFlags(includes, where, whereOr, groupBy, limit, page, orderBy, orderDirection, select, options).then((request) => request(axios, basePath));
        },
        updateCharacterPeqzoneFlag(id: number, characterPeqzoneFlag: ModelsCharacterPeqzoneFlag, options?: any): AxiosPromise<Array<ModelsCharacterPeqzoneFlag>> {
            return CharacterPeqzoneFlagApiFp(configuration).updateCharacterPeqzoneFlag(id, characterPeqzoneFlag, options).then((request) => request(axios, basePath));
        },
    };
};
export interface CharacterPeqzoneFlagApiCreateCharacterPeqzoneFlagRequest {
    readonly characterPeqzoneFlag: ModelsCharacterPeqzoneFlag
}
export interface CharacterPeqzoneFlagApiDeleteCharacterPeqzoneFlagRequest {
    readonly id: number
}
export interface CharacterPeqzoneFlagApiGetCharacterPeqzoneFlagRequest {
    readonly id: number
    readonly includes?: string
    readonly select?: string
}
export interface CharacterPeqzoneFlagApiGetCharacterPeqzoneFlagsBulkRequest {
    readonly body: CrudcontrollersBulkFetchByIdsGetRequest
}
export interface CharacterPeqzoneFlagApiListCharacterPeqzoneFlagsRequest {
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
export interface CharacterPeqzoneFlagApiUpdateCharacterPeqzoneFlagRequest {
    readonly id: number
    readonly characterPeqzoneFlag: ModelsCharacterPeqzoneFlag
}
export class CharacterPeqzoneFlagApi extends BaseAPI {
    public createCharacterPeqzoneFlag(requestParameters: CharacterPeqzoneFlagApiCreateCharacterPeqzoneFlagRequest, options?: any) {
        return CharacterPeqzoneFlagApiFp(this.configuration).createCharacterPeqzoneFlag(requestParameters.characterPeqzoneFlag, options).then((request) => request(this.axios, this.basePath));
    }
    public deleteCharacterPeqzoneFlag(requestParameters: CharacterPeqzoneFlagApiDeleteCharacterPeqzoneFlagRequest, options?: any) {
        return CharacterPeqzoneFlagApiFp(this.configuration).deleteCharacterPeqzoneFlag(requestParameters.id, options).then((request) => request(this.axios, this.basePath));
    }
    public getCharacterPeqzoneFlag(requestParameters: CharacterPeqzoneFlagApiGetCharacterPeqzoneFlagRequest, options?: any) {
        return CharacterPeqzoneFlagApiFp(this.configuration).getCharacterPeqzoneFlag(requestParameters.id, requestParameters.includes, requestParameters.select, options).then((request) => request(this.axios, this.basePath));
    }
    public getCharacterPeqzoneFlagsBulk(requestParameters: CharacterPeqzoneFlagApiGetCharacterPeqzoneFlagsBulkRequest, options?: any) {
        return CharacterPeqzoneFlagApiFp(this.configuration).getCharacterPeqzoneFlagsBulk(requestParameters.body, options).then((request) => request(this.axios, this.basePath));
    }
    public listCharacterPeqzoneFlags(requestParameters: CharacterPeqzoneFlagApiListCharacterPeqzoneFlagsRequest = {}, options?: any) {
        return CharacterPeqzoneFlagApiFp(this.configuration).listCharacterPeqzoneFlags(requestParameters.includes, requestParameters.where, requestParameters.whereOr, requestParameters.groupBy, requestParameters.limit, requestParameters.page, requestParameters.orderBy, requestParameters.orderDirection, requestParameters.select, options).then((request) => request(this.axios, this.basePath));
    }
    public updateCharacterPeqzoneFlag(requestParameters: CharacterPeqzoneFlagApiUpdateCharacterPeqzoneFlagRequest, options?: any) {
        return CharacterPeqzoneFlagApiFp(this.configuration).updateCharacterPeqzoneFlag(requestParameters.id, requestParameters.characterPeqzoneFlag, options).then((request) => request(this.axios, this.basePath));
    }
}
