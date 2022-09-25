import globalAxios, { AxiosPromise, AxiosInstance } from 'axios';
import { Configuration } from '../configuration';
// Some imports not used depending on template conditions
// @ts-ignore
import { BASE_PATH, COLLECTION_FORMATS, RequestArgs, BaseAPI, RequiredError } from '../base';
// @ts-ignore
import { CrudcontrollersBulkFetchByIdsGetRequest } from '../models';
// @ts-ignore
import { ModelsCharacterLeadershipAbility } from '../models';
export const CharacterLeadershipAbilityApiAxiosParamCreator = function (configuration?: Configuration) {
    return {
        createCharacterLeadershipAbility: async (characterLeadershipAbility: ModelsCharacterLeadershipAbility, options: any = {}): Promise<RequestArgs> => {
            if (characterLeadershipAbility === null || characterLeadershipAbility === undefined) {
                throw new RequiredError('characterLeadershipAbility','Required parameter characterLeadershipAbility was null or undefined when calling createCharacterLeadershipAbility.');
            }
            const localVarPath = `/character_leadership_ability`;
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
            const nonString = typeof characterLeadershipAbility !== 'string';
            const needsSerialization = nonString && configuration && configuration.isJsonMime
                ? configuration.isJsonMime(localVarRequestOptions.headers['Content-Type'])
                : nonString;
            localVarRequestOptions.data =  needsSerialization
                ? JSON.stringify(characterLeadershipAbility !== undefined ? characterLeadershipAbility : {})
                : (characterLeadershipAbility || "");
            return {
                url: localVarUrlObj.pathname + localVarUrlObj.search + localVarUrlObj.hash,
                options: localVarRequestOptions,
            };
        },
        deleteCharacterLeadershipAbility: async (id: number, options: any = {}): Promise<RequestArgs> => {
            if (id === null || id === undefined) {
                throw new RequiredError('id','Required parameter id was null or undefined when calling deleteCharacterLeadershipAbility.');
            }
            const localVarPath = `/character_leadership_ability/{id}`
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
        getCharacterLeadershipAbilitiesBulk: async (body: CrudcontrollersBulkFetchByIdsGetRequest, options: any = {}): Promise<RequestArgs> => {
            if (body === null || body === undefined) {
                throw new RequiredError('body','Required parameter body was null or undefined when calling getCharacterLeadershipAbilitiesBulk.');
            }
            const localVarPath = `/character_leadership_abilities/bulk`;
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
        getCharacterLeadershipAbility: async (id: number, includes?: string, select?: string, options: any = {}): Promise<RequestArgs> => {
            if (id === null || id === undefined) {
                throw new RequiredError('id','Required parameter id was null or undefined when calling getCharacterLeadershipAbility.');
            }
            const localVarPath = `/character_leadership_ability/{id}`
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
        listCharacterLeadershipAbilities: async (includes?: string, where?: string, whereOr?: string, groupBy?: string, limit?: string, page?: number, orderBy?: string, orderDirection?: string, select?: string, options: any = {}): Promise<RequestArgs> => {
            const localVarPath = `/character_leadership_abilities`;
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
        updateCharacterLeadershipAbility: async (id: number, characterLeadershipAbility: ModelsCharacterLeadershipAbility, options: any = {}): Promise<RequestArgs> => {
            if (id === null || id === undefined) {
                throw new RequiredError('id','Required parameter id was null or undefined when calling updateCharacterLeadershipAbility.');
            }
            if (characterLeadershipAbility === null || characterLeadershipAbility === undefined) {
                throw new RequiredError('characterLeadershipAbility','Required parameter characterLeadershipAbility was null or undefined when calling updateCharacterLeadershipAbility.');
            }
            const localVarPath = `/character_leadership_ability/{id}`
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
            const nonString = typeof characterLeadershipAbility !== 'string';
            const needsSerialization = nonString && configuration && configuration.isJsonMime
                ? configuration.isJsonMime(localVarRequestOptions.headers['Content-Type'])
                : nonString;
            localVarRequestOptions.data =  needsSerialization
                ? JSON.stringify(characterLeadershipAbility !== undefined ? characterLeadershipAbility : {})
                : (characterLeadershipAbility || "");
            return {
                url: localVarUrlObj.pathname + localVarUrlObj.search + localVarUrlObj.hash,
                options: localVarRequestOptions,
            };
        },
    }
};
export const CharacterLeadershipAbilityApiFp = function(configuration?: Configuration) {
    return {
        async createCharacterLeadershipAbility(characterLeadershipAbility: ModelsCharacterLeadershipAbility, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsCharacterLeadershipAbility>>> {
            const localVarAxiosArgs = await CharacterLeadershipAbilityApiAxiosParamCreator(configuration).createCharacterLeadershipAbility(characterLeadershipAbility, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        async deleteCharacterLeadershipAbility(id: number, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<string>> {
            const localVarAxiosArgs = await CharacterLeadershipAbilityApiAxiosParamCreator(configuration).deleteCharacterLeadershipAbility(id, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        async getCharacterLeadershipAbilitiesBulk(body: CrudcontrollersBulkFetchByIdsGetRequest, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsCharacterLeadershipAbility>>> {
            const localVarAxiosArgs = await CharacterLeadershipAbilityApiAxiosParamCreator(configuration).getCharacterLeadershipAbilitiesBulk(body, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        async getCharacterLeadershipAbility(id: number, includes?: string, select?: string, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsCharacterLeadershipAbility>>> {
            const localVarAxiosArgs = await CharacterLeadershipAbilityApiAxiosParamCreator(configuration).getCharacterLeadershipAbility(id, includes, select, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        async listCharacterLeadershipAbilities(includes?: string, where?: string, whereOr?: string, groupBy?: string, limit?: string, page?: number, orderBy?: string, orderDirection?: string, select?: string, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsCharacterLeadershipAbility>>> {
            const localVarAxiosArgs = await CharacterLeadershipAbilityApiAxiosParamCreator(configuration).listCharacterLeadershipAbilities(includes, where, whereOr, groupBy, limit, page, orderBy, orderDirection, select, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        async updateCharacterLeadershipAbility(id: number, characterLeadershipAbility: ModelsCharacterLeadershipAbility, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsCharacterLeadershipAbility>>> {
            const localVarAxiosArgs = await CharacterLeadershipAbilityApiAxiosParamCreator(configuration).updateCharacterLeadershipAbility(id, characterLeadershipAbility, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
    }
};
export const CharacterLeadershipAbilityApiFactory = function (configuration?: Configuration, basePath?: string, axios?: AxiosInstance) {
    return {
        createCharacterLeadershipAbility(characterLeadershipAbility: ModelsCharacterLeadershipAbility, options?: any): AxiosPromise<Array<ModelsCharacterLeadershipAbility>> {
            return CharacterLeadershipAbilityApiFp(configuration).createCharacterLeadershipAbility(characterLeadershipAbility, options).then((request) => request(axios, basePath));
        },
        deleteCharacterLeadershipAbility(id: number, options?: any): AxiosPromise<string> {
            return CharacterLeadershipAbilityApiFp(configuration).deleteCharacterLeadershipAbility(id, options).then((request) => request(axios, basePath));
        },
        getCharacterLeadershipAbilitiesBulk(body: CrudcontrollersBulkFetchByIdsGetRequest, options?: any): AxiosPromise<Array<ModelsCharacterLeadershipAbility>> {
            return CharacterLeadershipAbilityApiFp(configuration).getCharacterLeadershipAbilitiesBulk(body, options).then((request) => request(axios, basePath));
        },
        getCharacterLeadershipAbility(id: number, includes?: string, select?: string, options?: any): AxiosPromise<Array<ModelsCharacterLeadershipAbility>> {
            return CharacterLeadershipAbilityApiFp(configuration).getCharacterLeadershipAbility(id, includes, select, options).then((request) => request(axios, basePath));
        },
        listCharacterLeadershipAbilities(includes?: string, where?: string, whereOr?: string, groupBy?: string, limit?: string, page?: number, orderBy?: string, orderDirection?: string, select?: string, options?: any): AxiosPromise<Array<ModelsCharacterLeadershipAbility>> {
            return CharacterLeadershipAbilityApiFp(configuration).listCharacterLeadershipAbilities(includes, where, whereOr, groupBy, limit, page, orderBy, orderDirection, select, options).then((request) => request(axios, basePath));
        },
        updateCharacterLeadershipAbility(id: number, characterLeadershipAbility: ModelsCharacterLeadershipAbility, options?: any): AxiosPromise<Array<ModelsCharacterLeadershipAbility>> {
            return CharacterLeadershipAbilityApiFp(configuration).updateCharacterLeadershipAbility(id, characterLeadershipAbility, options).then((request) => request(axios, basePath));
        },
    };
};
export interface CharacterLeadershipAbilityApiCreateCharacterLeadershipAbilityRequest {
    readonly characterLeadershipAbility: ModelsCharacterLeadershipAbility
}
export interface CharacterLeadershipAbilityApiDeleteCharacterLeadershipAbilityRequest {
    readonly id: number
}
export interface CharacterLeadershipAbilityApiGetCharacterLeadershipAbilitiesBulkRequest {
    readonly body: CrudcontrollersBulkFetchByIdsGetRequest
}
export interface CharacterLeadershipAbilityApiGetCharacterLeadershipAbilityRequest {
    readonly id: number
    readonly includes?: string
    readonly select?: string
}
export interface CharacterLeadershipAbilityApiListCharacterLeadershipAbilitiesRequest {
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
export interface CharacterLeadershipAbilityApiUpdateCharacterLeadershipAbilityRequest {
    readonly id: number
    readonly characterLeadershipAbility: ModelsCharacterLeadershipAbility
}
export class CharacterLeadershipAbilityApi extends BaseAPI {
    public createCharacterLeadershipAbility(requestParameters: CharacterLeadershipAbilityApiCreateCharacterLeadershipAbilityRequest, options?: any) {
        return CharacterLeadershipAbilityApiFp(this.configuration).createCharacterLeadershipAbility(requestParameters.characterLeadershipAbility, options).then((request) => request(this.axios, this.basePath));
    }
    public deleteCharacterLeadershipAbility(requestParameters: CharacterLeadershipAbilityApiDeleteCharacterLeadershipAbilityRequest, options?: any) {
        return CharacterLeadershipAbilityApiFp(this.configuration).deleteCharacterLeadershipAbility(requestParameters.id, options).then((request) => request(this.axios, this.basePath));
    }
    public getCharacterLeadershipAbilitiesBulk(requestParameters: CharacterLeadershipAbilityApiGetCharacterLeadershipAbilitiesBulkRequest, options?: any) {
        return CharacterLeadershipAbilityApiFp(this.configuration).getCharacterLeadershipAbilitiesBulk(requestParameters.body, options).then((request) => request(this.axios, this.basePath));
    }
    public getCharacterLeadershipAbility(requestParameters: CharacterLeadershipAbilityApiGetCharacterLeadershipAbilityRequest, options?: any) {
        return CharacterLeadershipAbilityApiFp(this.configuration).getCharacterLeadershipAbility(requestParameters.id, requestParameters.includes, requestParameters.select, options).then((request) => request(this.axios, this.basePath));
    }
    public listCharacterLeadershipAbilities(requestParameters: CharacterLeadershipAbilityApiListCharacterLeadershipAbilitiesRequest = {}, options?: any) {
        return CharacterLeadershipAbilityApiFp(this.configuration).listCharacterLeadershipAbilities(requestParameters.includes, requestParameters.where, requestParameters.whereOr, requestParameters.groupBy, requestParameters.limit, requestParameters.page, requestParameters.orderBy, requestParameters.orderDirection, requestParameters.select, options).then((request) => request(this.axios, this.basePath));
    }
    public updateCharacterLeadershipAbility(requestParameters: CharacterLeadershipAbilityApiUpdateCharacterLeadershipAbilityRequest, options?: any) {
        return CharacterLeadershipAbilityApiFp(this.configuration).updateCharacterLeadershipAbility(requestParameters.id, requestParameters.characterLeadershipAbility, options).then((request) => request(this.axios, this.basePath));
    }
}
