import globalAxios, { AxiosPromise, AxiosInstance } from 'axios';
import { Configuration } from '../configuration';
// Some imports not used depending on template conditions
// @ts-ignore
import { BASE_PATH, COLLECTION_FORMATS, RequestArgs, BaseAPI, RequiredError } from '../base';
// @ts-ignore
import { CrudcontrollersBulkFetchByIdsGetRequest } from '../models';
// @ts-ignore
import { ModelsCharacterStatsRecord } from '../models';
export const CharacterStatsRecordApiAxiosParamCreator = function (configuration?: Configuration) {
    return {
        createCharacterStatsRecord: async (characterStatsRecord: ModelsCharacterStatsRecord, options: any = {}): Promise<RequestArgs> => {
            if (characterStatsRecord === null || characterStatsRecord === undefined) {
                throw new RequiredError('characterStatsRecord','Required parameter characterStatsRecord was null or undefined when calling createCharacterStatsRecord.');
            }
            const localVarPath = `/character_stats_record`;
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
            const nonString = typeof characterStatsRecord !== 'string';
            const needsSerialization = nonString && configuration && configuration.isJsonMime
                ? configuration.isJsonMime(localVarRequestOptions.headers['Content-Type'])
                : nonString;
            localVarRequestOptions.data =  needsSerialization
                ? JSON.stringify(characterStatsRecord !== undefined ? characterStatsRecord : {})
                : (characterStatsRecord || "");
            return {
                url: localVarUrlObj.pathname + localVarUrlObj.search + localVarUrlObj.hash,
                options: localVarRequestOptions,
            };
        },
        deleteCharacterStatsRecord: async (id: number, options: any = {}): Promise<RequestArgs> => {
            if (id === null || id === undefined) {
                throw new RequiredError('id','Required parameter id was null or undefined when calling deleteCharacterStatsRecord.');
            }
            const localVarPath = `/character_stats_record/{id}`
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
        getCharacterStatsRecord: async (id: number, includes?: string, select?: string, options: any = {}): Promise<RequestArgs> => {
            if (id === null || id === undefined) {
                throw new RequiredError('id','Required parameter id was null or undefined when calling getCharacterStatsRecord.');
            }
            const localVarPath = `/character_stats_record/{id}`
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
        getCharacterStatsRecordsBulk: async (body: CrudcontrollersBulkFetchByIdsGetRequest, options: any = {}): Promise<RequestArgs> => {
            if (body === null || body === undefined) {
                throw new RequiredError('body','Required parameter body was null or undefined when calling getCharacterStatsRecordsBulk.');
            }
            const localVarPath = `/character_stats_records/bulk`;
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
        getCharacterStatsRecordsCount: async (includes?: string, where?: string, whereOr?: string, groupBy?: string, limit?: string, page?: number, orderBy?: string, orderDirection?: string, select?: string, options: any = {}): Promise<RequestArgs> => {
            const localVarPath = `/character_stats_records/count`;
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
        listCharacterStatsRecords: async (includes?: string, where?: string, whereOr?: string, groupBy?: string, limit?: string, page?: number, orderBy?: string, orderDirection?: string, select?: string, options: any = {}): Promise<RequestArgs> => {
            const localVarPath = `/character_stats_records`;
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
        updateCharacterStatsRecord: async (id: number, characterStatsRecord: ModelsCharacterStatsRecord, options: any = {}): Promise<RequestArgs> => {
            if (id === null || id === undefined) {
                throw new RequiredError('id','Required parameter id was null or undefined when calling updateCharacterStatsRecord.');
            }
            if (characterStatsRecord === null || characterStatsRecord === undefined) {
                throw new RequiredError('characterStatsRecord','Required parameter characterStatsRecord was null or undefined when calling updateCharacterStatsRecord.');
            }
            const localVarPath = `/character_stats_record/{id}`
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
            const nonString = typeof characterStatsRecord !== 'string';
            const needsSerialization = nonString && configuration && configuration.isJsonMime
                ? configuration.isJsonMime(localVarRequestOptions.headers['Content-Type'])
                : nonString;
            localVarRequestOptions.data =  needsSerialization
                ? JSON.stringify(characterStatsRecord !== undefined ? characterStatsRecord : {})
                : (characterStatsRecord || "");
            return {
                url: localVarUrlObj.pathname + localVarUrlObj.search + localVarUrlObj.hash,
                options: localVarRequestOptions,
            };
        },
    }
};
export const CharacterStatsRecordApiFp = function(configuration?: Configuration) {
    return {
        async createCharacterStatsRecord(characterStatsRecord: ModelsCharacterStatsRecord, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsCharacterStatsRecord>>> {
            const localVarAxiosArgs = await CharacterStatsRecordApiAxiosParamCreator(configuration).createCharacterStatsRecord(characterStatsRecord, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        async deleteCharacterStatsRecord(id: number, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<string>> {
            const localVarAxiosArgs = await CharacterStatsRecordApiAxiosParamCreator(configuration).deleteCharacterStatsRecord(id, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        async getCharacterStatsRecord(id: number, includes?: string, select?: string, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsCharacterStatsRecord>>> {
            const localVarAxiosArgs = await CharacterStatsRecordApiAxiosParamCreator(configuration).getCharacterStatsRecord(id, includes, select, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        async getCharacterStatsRecordsBulk(body: CrudcontrollersBulkFetchByIdsGetRequest, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsCharacterStatsRecord>>> {
            const localVarAxiosArgs = await CharacterStatsRecordApiAxiosParamCreator(configuration).getCharacterStatsRecordsBulk(body, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        async getCharacterStatsRecordsCount(includes?: string, where?: string, whereOr?: string, groupBy?: string, limit?: string, page?: number, orderBy?: string, orderDirection?: string, select?: string, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsCharacterStatsRecord>>> {
            const localVarAxiosArgs = await CharacterStatsRecordApiAxiosParamCreator(configuration).getCharacterStatsRecordsCount(includes, where, whereOr, groupBy, limit, page, orderBy, orderDirection, select, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        async listCharacterStatsRecords(includes?: string, where?: string, whereOr?: string, groupBy?: string, limit?: string, page?: number, orderBy?: string, orderDirection?: string, select?: string, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsCharacterStatsRecord>>> {
            const localVarAxiosArgs = await CharacterStatsRecordApiAxiosParamCreator(configuration).listCharacterStatsRecords(includes, where, whereOr, groupBy, limit, page, orderBy, orderDirection, select, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        async updateCharacterStatsRecord(id: number, characterStatsRecord: ModelsCharacterStatsRecord, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsCharacterStatsRecord>>> {
            const localVarAxiosArgs = await CharacterStatsRecordApiAxiosParamCreator(configuration).updateCharacterStatsRecord(id, characterStatsRecord, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
    }
};
export const CharacterStatsRecordApiFactory = function (configuration?: Configuration, basePath?: string, axios?: AxiosInstance) {
    return {
        createCharacterStatsRecord(characterStatsRecord: ModelsCharacterStatsRecord, options?: any): AxiosPromise<Array<ModelsCharacterStatsRecord>> {
            return CharacterStatsRecordApiFp(configuration).createCharacterStatsRecord(characterStatsRecord, options).then((request) => request(axios, basePath));
        },
        deleteCharacterStatsRecord(id: number, options?: any): AxiosPromise<string> {
            return CharacterStatsRecordApiFp(configuration).deleteCharacterStatsRecord(id, options).then((request) => request(axios, basePath));
        },
        getCharacterStatsRecord(id: number, includes?: string, select?: string, options?: any): AxiosPromise<Array<ModelsCharacterStatsRecord>> {
            return CharacterStatsRecordApiFp(configuration).getCharacterStatsRecord(id, includes, select, options).then((request) => request(axios, basePath));
        },
        getCharacterStatsRecordsBulk(body: CrudcontrollersBulkFetchByIdsGetRequest, options?: any): AxiosPromise<Array<ModelsCharacterStatsRecord>> {
            return CharacterStatsRecordApiFp(configuration).getCharacterStatsRecordsBulk(body, options).then((request) => request(axios, basePath));
        },
        getCharacterStatsRecordsCount(includes?: string, where?: string, whereOr?: string, groupBy?: string, limit?: string, page?: number, orderBy?: string, orderDirection?: string, select?: string, options?: any): AxiosPromise<Array<ModelsCharacterStatsRecord>> {
            return CharacterStatsRecordApiFp(configuration).getCharacterStatsRecordsCount(includes, where, whereOr, groupBy, limit, page, orderBy, orderDirection, select, options).then((request) => request(axios, basePath));
        },
        listCharacterStatsRecords(includes?: string, where?: string, whereOr?: string, groupBy?: string, limit?: string, page?: number, orderBy?: string, orderDirection?: string, select?: string, options?: any): AxiosPromise<Array<ModelsCharacterStatsRecord>> {
            return CharacterStatsRecordApiFp(configuration).listCharacterStatsRecords(includes, where, whereOr, groupBy, limit, page, orderBy, orderDirection, select, options).then((request) => request(axios, basePath));
        },
        updateCharacterStatsRecord(id: number, characterStatsRecord: ModelsCharacterStatsRecord, options?: any): AxiosPromise<Array<ModelsCharacterStatsRecord>> {
            return CharacterStatsRecordApiFp(configuration).updateCharacterStatsRecord(id, characterStatsRecord, options).then((request) => request(axios, basePath));
        },
    };
};
export interface CharacterStatsRecordApiCreateCharacterStatsRecordRequest {
    readonly characterStatsRecord: ModelsCharacterStatsRecord
}
export interface CharacterStatsRecordApiDeleteCharacterStatsRecordRequest {
    readonly id: number
}
export interface CharacterStatsRecordApiGetCharacterStatsRecordRequest {
    readonly id: number
    readonly includes?: string
    readonly select?: string
}
export interface CharacterStatsRecordApiGetCharacterStatsRecordsBulkRequest {
    readonly body: CrudcontrollersBulkFetchByIdsGetRequest
}
export interface CharacterStatsRecordApiGetCharacterStatsRecordsCountRequest {
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
export interface CharacterStatsRecordApiListCharacterStatsRecordsRequest {
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
export interface CharacterStatsRecordApiUpdateCharacterStatsRecordRequest {
    readonly id: number
    readonly characterStatsRecord: ModelsCharacterStatsRecord
}
export class CharacterStatsRecordApi extends BaseAPI {
    public createCharacterStatsRecord(requestParameters: CharacterStatsRecordApiCreateCharacterStatsRecordRequest, options?: any) {
        return CharacterStatsRecordApiFp(this.configuration).createCharacterStatsRecord(requestParameters.characterStatsRecord, options).then((request) => request(this.axios, this.basePath));
    }
    public deleteCharacterStatsRecord(requestParameters: CharacterStatsRecordApiDeleteCharacterStatsRecordRequest, options?: any) {
        return CharacterStatsRecordApiFp(this.configuration).deleteCharacterStatsRecord(requestParameters.id, options).then((request) => request(this.axios, this.basePath));
    }
    public getCharacterStatsRecord(requestParameters: CharacterStatsRecordApiGetCharacterStatsRecordRequest, options?: any) {
        return CharacterStatsRecordApiFp(this.configuration).getCharacterStatsRecord(requestParameters.id, requestParameters.includes, requestParameters.select, options).then((request) => request(this.axios, this.basePath));
    }
    public getCharacterStatsRecordsBulk(requestParameters: CharacterStatsRecordApiGetCharacterStatsRecordsBulkRequest, options?: any) {
        return CharacterStatsRecordApiFp(this.configuration).getCharacterStatsRecordsBulk(requestParameters.body, options).then((request) => request(this.axios, this.basePath));
    }
    public getCharacterStatsRecordsCount(requestParameters: CharacterStatsRecordApiGetCharacterStatsRecordsCountRequest = {}, options?: any) {
        return CharacterStatsRecordApiFp(this.configuration).getCharacterStatsRecordsCount(requestParameters.includes, requestParameters.where, requestParameters.whereOr, requestParameters.groupBy, requestParameters.limit, requestParameters.page, requestParameters.orderBy, requestParameters.orderDirection, requestParameters.select, options).then((request) => request(this.axios, this.basePath));
    }
    public listCharacterStatsRecords(requestParameters: CharacterStatsRecordApiListCharacterStatsRecordsRequest = {}, options?: any) {
        return CharacterStatsRecordApiFp(this.configuration).listCharacterStatsRecords(requestParameters.includes, requestParameters.where, requestParameters.whereOr, requestParameters.groupBy, requestParameters.limit, requestParameters.page, requestParameters.orderBy, requestParameters.orderDirection, requestParameters.select, options).then((request) => request(this.axios, this.basePath));
    }
    public updateCharacterStatsRecord(requestParameters: CharacterStatsRecordApiUpdateCharacterStatsRecordRequest, options?: any) {
        return CharacterStatsRecordApiFp(this.configuration).updateCharacterStatsRecord(requestParameters.id, requestParameters.characterStatsRecord, options).then((request) => request(this.axios, this.basePath));
    }
}
