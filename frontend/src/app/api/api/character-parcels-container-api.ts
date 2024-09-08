import globalAxios, { AxiosPromise, AxiosInstance } from 'axios';
import { Configuration } from '../configuration';
// Some imports not used depending on template conditions
// @ts-ignore
import { BASE_PATH, COLLECTION_FORMATS, RequestArgs, BaseAPI, RequiredError } from '../base';
// @ts-ignore
import { CrudcontrollersBulkFetchByIdsGetRequest } from '../models';
// @ts-ignore
import { ModelsCharacterParcelsContainer } from '../models';
export const CharacterParcelsContainerApiAxiosParamCreator = function (configuration?: Configuration) {
    return {
        createCharacterParcelsContainer: async (characterParcelsContainer: ModelsCharacterParcelsContainer, options: any = {}): Promise<RequestArgs> => {
            if (characterParcelsContainer === null || characterParcelsContainer === undefined) {
                throw new RequiredError('characterParcelsContainer','Required parameter characterParcelsContainer was null or undefined when calling createCharacterParcelsContainer.');
            }
            const localVarPath = `/character_parcels_container`;
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
            const nonString = typeof characterParcelsContainer !== 'string';
            const needsSerialization = nonString && configuration && configuration.isJsonMime
                ? configuration.isJsonMime(localVarRequestOptions.headers['Content-Type'])
                : nonString;
            localVarRequestOptions.data =  needsSerialization
                ? JSON.stringify(characterParcelsContainer !== undefined ? characterParcelsContainer : {})
                : (characterParcelsContainer || "");
            return {
                url: localVarUrlObj.pathname + localVarUrlObj.search + localVarUrlObj.hash,
                options: localVarRequestOptions,
            };
        },
        deleteCharacterParcelsContainer: async (id: number, options: any = {}): Promise<RequestArgs> => {
            if (id === null || id === undefined) {
                throw new RequiredError('id','Required parameter id was null or undefined when calling deleteCharacterParcelsContainer.');
            }
            const localVarPath = `/character_parcels_container/{id}`
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
        getCharacterParcelsContainer: async (id: number, includes?: string, select?: string, options: any = {}): Promise<RequestArgs> => {
            if (id === null || id === undefined) {
                throw new RequiredError('id','Required parameter id was null or undefined when calling getCharacterParcelsContainer.');
            }
            const localVarPath = `/character_parcels_container/{id}`
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
        getCharacterParcelsContainersBulk: async (body: CrudcontrollersBulkFetchByIdsGetRequest, options: any = {}): Promise<RequestArgs> => {
            if (body === null || body === undefined) {
                throw new RequiredError('body','Required parameter body was null or undefined when calling getCharacterParcelsContainersBulk.');
            }
            const localVarPath = `/character_parcels_containers/bulk`;
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
        getCharacterParcelsContainersCount: async (includes?: string, where?: string, whereOr?: string, groupBy?: string, limit?: string, page?: number, orderBy?: string, orderDirection?: string, select?: string, options: any = {}): Promise<RequestArgs> => {
            const localVarPath = `/character_parcels_containers/count`;
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
        listCharacterParcelsContainers: async (includes?: string, where?: string, whereOr?: string, groupBy?: string, limit?: string, page?: number, orderBy?: string, orderDirection?: string, select?: string, options: any = {}): Promise<RequestArgs> => {
            const localVarPath = `/character_parcels_containers`;
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
        updateCharacterParcelsContainer: async (id: number, characterParcelsContainer: ModelsCharacterParcelsContainer, options: any = {}): Promise<RequestArgs> => {
            if (id === null || id === undefined) {
                throw new RequiredError('id','Required parameter id was null or undefined when calling updateCharacterParcelsContainer.');
            }
            if (characterParcelsContainer === null || characterParcelsContainer === undefined) {
                throw new RequiredError('characterParcelsContainer','Required parameter characterParcelsContainer was null or undefined when calling updateCharacterParcelsContainer.');
            }
            const localVarPath = `/character_parcels_container/{id}`
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
            const nonString = typeof characterParcelsContainer !== 'string';
            const needsSerialization = nonString && configuration && configuration.isJsonMime
                ? configuration.isJsonMime(localVarRequestOptions.headers['Content-Type'])
                : nonString;
            localVarRequestOptions.data =  needsSerialization
                ? JSON.stringify(characterParcelsContainer !== undefined ? characterParcelsContainer : {})
                : (characterParcelsContainer || "");
            return {
                url: localVarUrlObj.pathname + localVarUrlObj.search + localVarUrlObj.hash,
                options: localVarRequestOptions,
            };
        },
    }
};
export const CharacterParcelsContainerApiFp = function(configuration?: Configuration) {
    return {
        async createCharacterParcelsContainer(characterParcelsContainer: ModelsCharacterParcelsContainer, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsCharacterParcelsContainer>>> {
            const localVarAxiosArgs = await CharacterParcelsContainerApiAxiosParamCreator(configuration).createCharacterParcelsContainer(characterParcelsContainer, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        async deleteCharacterParcelsContainer(id: number, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<string>> {
            const localVarAxiosArgs = await CharacterParcelsContainerApiAxiosParamCreator(configuration).deleteCharacterParcelsContainer(id, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        async getCharacterParcelsContainer(id: number, includes?: string, select?: string, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsCharacterParcelsContainer>>> {
            const localVarAxiosArgs = await CharacterParcelsContainerApiAxiosParamCreator(configuration).getCharacterParcelsContainer(id, includes, select, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        async getCharacterParcelsContainersBulk(body: CrudcontrollersBulkFetchByIdsGetRequest, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsCharacterParcelsContainer>>> {
            const localVarAxiosArgs = await CharacterParcelsContainerApiAxiosParamCreator(configuration).getCharacterParcelsContainersBulk(body, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        async getCharacterParcelsContainersCount(includes?: string, where?: string, whereOr?: string, groupBy?: string, limit?: string, page?: number, orderBy?: string, orderDirection?: string, select?: string, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsCharacterParcelsContainer>>> {
            const localVarAxiosArgs = await CharacterParcelsContainerApiAxiosParamCreator(configuration).getCharacterParcelsContainersCount(includes, where, whereOr, groupBy, limit, page, orderBy, orderDirection, select, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        async listCharacterParcelsContainers(includes?: string, where?: string, whereOr?: string, groupBy?: string, limit?: string, page?: number, orderBy?: string, orderDirection?: string, select?: string, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsCharacterParcelsContainer>>> {
            const localVarAxiosArgs = await CharacterParcelsContainerApiAxiosParamCreator(configuration).listCharacterParcelsContainers(includes, where, whereOr, groupBy, limit, page, orderBy, orderDirection, select, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        async updateCharacterParcelsContainer(id: number, characterParcelsContainer: ModelsCharacterParcelsContainer, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsCharacterParcelsContainer>>> {
            const localVarAxiosArgs = await CharacterParcelsContainerApiAxiosParamCreator(configuration).updateCharacterParcelsContainer(id, characterParcelsContainer, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
    }
};
export const CharacterParcelsContainerApiFactory = function (configuration?: Configuration, basePath?: string, axios?: AxiosInstance) {
    return {
        createCharacterParcelsContainer(characterParcelsContainer: ModelsCharacterParcelsContainer, options?: any): AxiosPromise<Array<ModelsCharacterParcelsContainer>> {
            return CharacterParcelsContainerApiFp(configuration).createCharacterParcelsContainer(characterParcelsContainer, options).then((request) => request(axios, basePath));
        },
        deleteCharacterParcelsContainer(id: number, options?: any): AxiosPromise<string> {
            return CharacterParcelsContainerApiFp(configuration).deleteCharacterParcelsContainer(id, options).then((request) => request(axios, basePath));
        },
        getCharacterParcelsContainer(id: number, includes?: string, select?: string, options?: any): AxiosPromise<Array<ModelsCharacterParcelsContainer>> {
            return CharacterParcelsContainerApiFp(configuration).getCharacterParcelsContainer(id, includes, select, options).then((request) => request(axios, basePath));
        },
        getCharacterParcelsContainersBulk(body: CrudcontrollersBulkFetchByIdsGetRequest, options?: any): AxiosPromise<Array<ModelsCharacterParcelsContainer>> {
            return CharacterParcelsContainerApiFp(configuration).getCharacterParcelsContainersBulk(body, options).then((request) => request(axios, basePath));
        },
        getCharacterParcelsContainersCount(includes?: string, where?: string, whereOr?: string, groupBy?: string, limit?: string, page?: number, orderBy?: string, orderDirection?: string, select?: string, options?: any): AxiosPromise<Array<ModelsCharacterParcelsContainer>> {
            return CharacterParcelsContainerApiFp(configuration).getCharacterParcelsContainersCount(includes, where, whereOr, groupBy, limit, page, orderBy, orderDirection, select, options).then((request) => request(axios, basePath));
        },
        listCharacterParcelsContainers(includes?: string, where?: string, whereOr?: string, groupBy?: string, limit?: string, page?: number, orderBy?: string, orderDirection?: string, select?: string, options?: any): AxiosPromise<Array<ModelsCharacterParcelsContainer>> {
            return CharacterParcelsContainerApiFp(configuration).listCharacterParcelsContainers(includes, where, whereOr, groupBy, limit, page, orderBy, orderDirection, select, options).then((request) => request(axios, basePath));
        },
        updateCharacterParcelsContainer(id: number, characterParcelsContainer: ModelsCharacterParcelsContainer, options?: any): AxiosPromise<Array<ModelsCharacterParcelsContainer>> {
            return CharacterParcelsContainerApiFp(configuration).updateCharacterParcelsContainer(id, characterParcelsContainer, options).then((request) => request(axios, basePath));
        },
    };
};
export interface CharacterParcelsContainerApiCreateCharacterParcelsContainerRequest {
    readonly characterParcelsContainer: ModelsCharacterParcelsContainer
}
export interface CharacterParcelsContainerApiDeleteCharacterParcelsContainerRequest {
    readonly id: number
}
export interface CharacterParcelsContainerApiGetCharacterParcelsContainerRequest {
    readonly id: number
    readonly includes?: string
    readonly select?: string
}
export interface CharacterParcelsContainerApiGetCharacterParcelsContainersBulkRequest {
    readonly body: CrudcontrollersBulkFetchByIdsGetRequest
}
export interface CharacterParcelsContainerApiGetCharacterParcelsContainersCountRequest {
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
export interface CharacterParcelsContainerApiListCharacterParcelsContainersRequest {
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
export interface CharacterParcelsContainerApiUpdateCharacterParcelsContainerRequest {
    readonly id: number
    readonly characterParcelsContainer: ModelsCharacterParcelsContainer
}
export class CharacterParcelsContainerApi extends BaseAPI {
    public createCharacterParcelsContainer(requestParameters: CharacterParcelsContainerApiCreateCharacterParcelsContainerRequest, options?: any) {
        return CharacterParcelsContainerApiFp(this.configuration).createCharacterParcelsContainer(requestParameters.characterParcelsContainer, options).then((request) => request(this.axios, this.basePath));
    }
    public deleteCharacterParcelsContainer(requestParameters: CharacterParcelsContainerApiDeleteCharacterParcelsContainerRequest, options?: any) {
        return CharacterParcelsContainerApiFp(this.configuration).deleteCharacterParcelsContainer(requestParameters.id, options).then((request) => request(this.axios, this.basePath));
    }
    public getCharacterParcelsContainer(requestParameters: CharacterParcelsContainerApiGetCharacterParcelsContainerRequest, options?: any) {
        return CharacterParcelsContainerApiFp(this.configuration).getCharacterParcelsContainer(requestParameters.id, requestParameters.includes, requestParameters.select, options).then((request) => request(this.axios, this.basePath));
    }
    public getCharacterParcelsContainersBulk(requestParameters: CharacterParcelsContainerApiGetCharacterParcelsContainersBulkRequest, options?: any) {
        return CharacterParcelsContainerApiFp(this.configuration).getCharacterParcelsContainersBulk(requestParameters.body, options).then((request) => request(this.axios, this.basePath));
    }
    public getCharacterParcelsContainersCount(requestParameters: CharacterParcelsContainerApiGetCharacterParcelsContainersCountRequest = {}, options?: any) {
        return CharacterParcelsContainerApiFp(this.configuration).getCharacterParcelsContainersCount(requestParameters.includes, requestParameters.where, requestParameters.whereOr, requestParameters.groupBy, requestParameters.limit, requestParameters.page, requestParameters.orderBy, requestParameters.orderDirection, requestParameters.select, options).then((request) => request(this.axios, this.basePath));
    }
    public listCharacterParcelsContainers(requestParameters: CharacterParcelsContainerApiListCharacterParcelsContainersRequest = {}, options?: any) {
        return CharacterParcelsContainerApiFp(this.configuration).listCharacterParcelsContainers(requestParameters.includes, requestParameters.where, requestParameters.whereOr, requestParameters.groupBy, requestParameters.limit, requestParameters.page, requestParameters.orderBy, requestParameters.orderDirection, requestParameters.select, options).then((request) => request(this.axios, this.basePath));
    }
    public updateCharacterParcelsContainer(requestParameters: CharacterParcelsContainerApiUpdateCharacterParcelsContainerRequest, options?: any) {
        return CharacterParcelsContainerApiFp(this.configuration).updateCharacterParcelsContainer(requestParameters.id, requestParameters.characterParcelsContainer, options).then((request) => request(this.axios, this.basePath));
    }
}
