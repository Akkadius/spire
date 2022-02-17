/* tslint:disable */
/* eslint-disable */
/**
 * Spire
 * Spire API documentation
 *
 * The version of the OpenAPI document: 3.0
 * Contact: akkadius1@gmail.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


import globalAxios, { AxiosPromise, AxiosInstance } from 'axios';
import { Configuration } from '../configuration';
// Some imports not used depending on template conditions
// @ts-ignore
import { BASE_PATH, COLLECTION_FORMATS, RequestArgs, BaseAPI, RequiredError } from '../base';
// @ts-ignore
import { CrudcontrollersBulkFetchByIdsGetRequest } from '../models';
// @ts-ignore
import { ModelsCharacterInspectMessage } from '../models';
/**
 * CharacterInspectMessageApi - axios parameter creator
 * @export
 */
export const CharacterInspectMessageApiAxiosParamCreator = function (configuration?: Configuration) {
    return {
        /**
         * 
         * @summary Creates CharacterInspectMessage
         * @param {ModelsCharacterInspectMessage} characterInspectMessage CharacterInspectMessage
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        createCharacterInspectMessage: async (characterInspectMessage: ModelsCharacterInspectMessage, options: any = {}): Promise<RequestArgs> => {
            // verify required parameter 'characterInspectMessage' is not null or undefined
            if (characterInspectMessage === null || characterInspectMessage === undefined) {
                throw new RequiredError('characterInspectMessage','Required parameter characterInspectMessage was null or undefined when calling createCharacterInspectMessage.');
            }
            const localVarPath = `/character_inspect_message`;
            // use dummy base URL string because the URL constructor only accepts absolute URLs.
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
        /**
         * 
         * @summary Deletes CharacterInspectMessage
         * @param {number} id id
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        deleteCharacterInspectMessage: async (id: number, options: any = {}): Promise<RequestArgs> => {
            // verify required parameter 'id' is not null or undefined
            if (id === null || id === undefined) {
                throw new RequiredError('id','Required parameter id was null or undefined when calling deleteCharacterInspectMessage.');
            }
            const localVarPath = `/character_inspect_message/{id}`
                .replace(`{${"id"}}`, encodeURIComponent(String(id)));
            // use dummy base URL string because the URL constructor only accepts absolute URLs.
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
        /**
         * 
         * @summary Gets CharacterInspectMessage
         * @param {number} id Id
         * @param {string} [includes] Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names 
         * @param {string} [select] Column names [.] separated to fetch specific fields in response
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        getCharacterInspectMessage: async (id: number, includes?: string, select?: string, options: any = {}): Promise<RequestArgs> => {
            // verify required parameter 'id' is not null or undefined
            if (id === null || id === undefined) {
                throw new RequiredError('id','Required parameter id was null or undefined when calling getCharacterInspectMessage.');
            }
            const localVarPath = `/character_inspect_message/{id}`
                .replace(`{${"id"}}`, encodeURIComponent(String(id)));
            // use dummy base URL string because the URL constructor only accepts absolute URLs.
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
        /**
         * 
         * @summary Gets CharacterInspectMessages in bulk
         * @param {CrudcontrollersBulkFetchByIdsGetRequest} body body
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        getCharacterInspectMessagesBulk: async (body: CrudcontrollersBulkFetchByIdsGetRequest, options: any = {}): Promise<RequestArgs> => {
            // verify required parameter 'body' is not null or undefined
            if (body === null || body === undefined) {
                throw new RequiredError('body','Required parameter body was null or undefined when calling getCharacterInspectMessagesBulk.');
            }
            const localVarPath = `/character_inspect_messages/bulk`;
            // use dummy base URL string because the URL constructor only accepts absolute URLs.
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
        /**
         * 
         * @summary Lists CharacterInspectMessages
         * @param {string} [includes] Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names 
         * @param {string} [where] Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2
         * @param {string} [whereOr] Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2
         * @param {string} [groupBy] Group by field. Multiple conditions [.] separated Example: field1.field2
         * @param {string} [limit] Rows to limit in response (Default: 10,000)
         * @param {string} [orderBy] Order by [field]
         * @param {string} [orderDirection] Order by field direction
         * @param {string} [select] Column names [.] separated to fetch specific fields in response
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        listCharacterInspectMessages: async (includes?: string, where?: string, whereOr?: string, groupBy?: string, limit?: string, orderBy?: string, orderDirection?: string, select?: string, options: any = {}): Promise<RequestArgs> => {
            const localVarPath = `/character_inspect_messages`;
            // use dummy base URL string because the URL constructor only accepts absolute URLs.
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
        /**
         * 
         * @summary Updates CharacterInspectMessage
         * @param {number} id Id
         * @param {ModelsCharacterInspectMessage} characterInspectMessage CharacterInspectMessage
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        updateCharacterInspectMessage: async (id: number, characterInspectMessage: ModelsCharacterInspectMessage, options: any = {}): Promise<RequestArgs> => {
            // verify required parameter 'id' is not null or undefined
            if (id === null || id === undefined) {
                throw new RequiredError('id','Required parameter id was null or undefined when calling updateCharacterInspectMessage.');
            }
            // verify required parameter 'characterInspectMessage' is not null or undefined
            if (characterInspectMessage === null || characterInspectMessage === undefined) {
                throw new RequiredError('characterInspectMessage','Required parameter characterInspectMessage was null or undefined when calling updateCharacterInspectMessage.');
            }
            const localVarPath = `/character_inspect_message/{id}`
                .replace(`{${"id"}}`, encodeURIComponent(String(id)));
            // use dummy base URL string because the URL constructor only accepts absolute URLs.
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

/**
 * CharacterInspectMessageApi - functional programming interface
 * @export
 */
export const CharacterInspectMessageApiFp = function(configuration?: Configuration) {
    return {
        /**
         * 
         * @summary Creates CharacterInspectMessage
         * @param {ModelsCharacterInspectMessage} characterInspectMessage CharacterInspectMessage
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async createCharacterInspectMessage(characterInspectMessage: ModelsCharacterInspectMessage, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsCharacterInspectMessage>>> {
            const localVarAxiosArgs = await CharacterInspectMessageApiAxiosParamCreator(configuration).createCharacterInspectMessage(characterInspectMessage, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        /**
         * 
         * @summary Deletes CharacterInspectMessage
         * @param {number} id id
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async deleteCharacterInspectMessage(id: number, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<string>> {
            const localVarAxiosArgs = await CharacterInspectMessageApiAxiosParamCreator(configuration).deleteCharacterInspectMessage(id, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        /**
         * 
         * @summary Gets CharacterInspectMessage
         * @param {number} id Id
         * @param {string} [includes] Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names 
         * @param {string} [select] Column names [.] separated to fetch specific fields in response
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async getCharacterInspectMessage(id: number, includes?: string, select?: string, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsCharacterInspectMessage>>> {
            const localVarAxiosArgs = await CharacterInspectMessageApiAxiosParamCreator(configuration).getCharacterInspectMessage(id, includes, select, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        /**
         * 
         * @summary Gets CharacterInspectMessages in bulk
         * @param {CrudcontrollersBulkFetchByIdsGetRequest} body body
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async getCharacterInspectMessagesBulk(body: CrudcontrollersBulkFetchByIdsGetRequest, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsCharacterInspectMessage>>> {
            const localVarAxiosArgs = await CharacterInspectMessageApiAxiosParamCreator(configuration).getCharacterInspectMessagesBulk(body, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        /**
         * 
         * @summary Lists CharacterInspectMessages
         * @param {string} [includes] Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names 
         * @param {string} [where] Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2
         * @param {string} [whereOr] Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2
         * @param {string} [groupBy] Group by field. Multiple conditions [.] separated Example: field1.field2
         * @param {string} [limit] Rows to limit in response (Default: 10,000)
         * @param {string} [orderBy] Order by [field]
         * @param {string} [orderDirection] Order by field direction
         * @param {string} [select] Column names [.] separated to fetch specific fields in response
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async listCharacterInspectMessages(includes?: string, where?: string, whereOr?: string, groupBy?: string, limit?: string, orderBy?: string, orderDirection?: string, select?: string, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsCharacterInspectMessage>>> {
            const localVarAxiosArgs = await CharacterInspectMessageApiAxiosParamCreator(configuration).listCharacterInspectMessages(includes, where, whereOr, groupBy, limit, orderBy, orderDirection, select, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        /**
         * 
         * @summary Updates CharacterInspectMessage
         * @param {number} id Id
         * @param {ModelsCharacterInspectMessage} characterInspectMessage CharacterInspectMessage
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async updateCharacterInspectMessage(id: number, characterInspectMessage: ModelsCharacterInspectMessage, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsCharacterInspectMessage>>> {
            const localVarAxiosArgs = await CharacterInspectMessageApiAxiosParamCreator(configuration).updateCharacterInspectMessage(id, characterInspectMessage, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
    }
};

/**
 * CharacterInspectMessageApi - factory interface
 * @export
 */
export const CharacterInspectMessageApiFactory = function (configuration?: Configuration, basePath?: string, axios?: AxiosInstance) {
    return {
        /**
         * 
         * @summary Creates CharacterInspectMessage
         * @param {ModelsCharacterInspectMessage} characterInspectMessage CharacterInspectMessage
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        createCharacterInspectMessage(characterInspectMessage: ModelsCharacterInspectMessage, options?: any): AxiosPromise<Array<ModelsCharacterInspectMessage>> {
            return CharacterInspectMessageApiFp(configuration).createCharacterInspectMessage(characterInspectMessage, options).then((request) => request(axios, basePath));
        },
        /**
         * 
         * @summary Deletes CharacterInspectMessage
         * @param {number} id id
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        deleteCharacterInspectMessage(id: number, options?: any): AxiosPromise<string> {
            return CharacterInspectMessageApiFp(configuration).deleteCharacterInspectMessage(id, options).then((request) => request(axios, basePath));
        },
        /**
         * 
         * @summary Gets CharacterInspectMessage
         * @param {number} id Id
         * @param {string} [includes] Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names 
         * @param {string} [select] Column names [.] separated to fetch specific fields in response
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        getCharacterInspectMessage(id: number, includes?: string, select?: string, options?: any): AxiosPromise<Array<ModelsCharacterInspectMessage>> {
            return CharacterInspectMessageApiFp(configuration).getCharacterInspectMessage(id, includes, select, options).then((request) => request(axios, basePath));
        },
        /**
         * 
         * @summary Gets CharacterInspectMessages in bulk
         * @param {CrudcontrollersBulkFetchByIdsGetRequest} body body
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        getCharacterInspectMessagesBulk(body: CrudcontrollersBulkFetchByIdsGetRequest, options?: any): AxiosPromise<Array<ModelsCharacterInspectMessage>> {
            return CharacterInspectMessageApiFp(configuration).getCharacterInspectMessagesBulk(body, options).then((request) => request(axios, basePath));
        },
        /**
         * 
         * @summary Lists CharacterInspectMessages
         * @param {string} [includes] Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names 
         * @param {string} [where] Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2
         * @param {string} [whereOr] Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2
         * @param {string} [groupBy] Group by field. Multiple conditions [.] separated Example: field1.field2
         * @param {string} [limit] Rows to limit in response (Default: 10,000)
         * @param {string} [orderBy] Order by [field]
         * @param {string} [orderDirection] Order by field direction
         * @param {string} [select] Column names [.] separated to fetch specific fields in response
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        listCharacterInspectMessages(includes?: string, where?: string, whereOr?: string, groupBy?: string, limit?: string, orderBy?: string, orderDirection?: string, select?: string, options?: any): AxiosPromise<Array<ModelsCharacterInspectMessage>> {
            return CharacterInspectMessageApiFp(configuration).listCharacterInspectMessages(includes, where, whereOr, groupBy, limit, orderBy, orderDirection, select, options).then((request) => request(axios, basePath));
        },
        /**
         * 
         * @summary Updates CharacterInspectMessage
         * @param {number} id Id
         * @param {ModelsCharacterInspectMessage} characterInspectMessage CharacterInspectMessage
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        updateCharacterInspectMessage(id: number, characterInspectMessage: ModelsCharacterInspectMessage, options?: any): AxiosPromise<Array<ModelsCharacterInspectMessage>> {
            return CharacterInspectMessageApiFp(configuration).updateCharacterInspectMessage(id, characterInspectMessage, options).then((request) => request(axios, basePath));
        },
    };
};

/**
 * Request parameters for createCharacterInspectMessage operation in CharacterInspectMessageApi.
 * @export
 * @interface CharacterInspectMessageApiCreateCharacterInspectMessageRequest
 */
export interface CharacterInspectMessageApiCreateCharacterInspectMessageRequest {
    /**
     * CharacterInspectMessage
     * @type {ModelsCharacterInspectMessage}
     * @memberof CharacterInspectMessageApiCreateCharacterInspectMessage
     */
    readonly characterInspectMessage: ModelsCharacterInspectMessage
}

/**
 * Request parameters for deleteCharacterInspectMessage operation in CharacterInspectMessageApi.
 * @export
 * @interface CharacterInspectMessageApiDeleteCharacterInspectMessageRequest
 */
export interface CharacterInspectMessageApiDeleteCharacterInspectMessageRequest {
    /**
     * id
     * @type {number}
     * @memberof CharacterInspectMessageApiDeleteCharacterInspectMessage
     */
    readonly id: number
}

/**
 * Request parameters for getCharacterInspectMessage operation in CharacterInspectMessageApi.
 * @export
 * @interface CharacterInspectMessageApiGetCharacterInspectMessageRequest
 */
export interface CharacterInspectMessageApiGetCharacterInspectMessageRequest {
    /**
     * Id
     * @type {number}
     * @memberof CharacterInspectMessageApiGetCharacterInspectMessage
     */
    readonly id: number

    /**
     * Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names 
     * @type {string}
     * @memberof CharacterInspectMessageApiGetCharacterInspectMessage
     */
    readonly includes?: string

    /**
     * Column names [.] separated to fetch specific fields in response
     * @type {string}
     * @memberof CharacterInspectMessageApiGetCharacterInspectMessage
     */
    readonly select?: string
}

/**
 * Request parameters for getCharacterInspectMessagesBulk operation in CharacterInspectMessageApi.
 * @export
 * @interface CharacterInspectMessageApiGetCharacterInspectMessagesBulkRequest
 */
export interface CharacterInspectMessageApiGetCharacterInspectMessagesBulkRequest {
    /**
     * body
     * @type {CrudcontrollersBulkFetchByIdsGetRequest}
     * @memberof CharacterInspectMessageApiGetCharacterInspectMessagesBulk
     */
    readonly body: CrudcontrollersBulkFetchByIdsGetRequest
}

/**
 * Request parameters for listCharacterInspectMessages operation in CharacterInspectMessageApi.
 * @export
 * @interface CharacterInspectMessageApiListCharacterInspectMessagesRequest
 */
export interface CharacterInspectMessageApiListCharacterInspectMessagesRequest {
    /**
     * Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names 
     * @type {string}
     * @memberof CharacterInspectMessageApiListCharacterInspectMessages
     */
    readonly includes?: string

    /**
     * Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2
     * @type {string}
     * @memberof CharacterInspectMessageApiListCharacterInspectMessages
     */
    readonly where?: string

    /**
     * Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2
     * @type {string}
     * @memberof CharacterInspectMessageApiListCharacterInspectMessages
     */
    readonly whereOr?: string

    /**
     * Group by field. Multiple conditions [.] separated Example: field1.field2
     * @type {string}
     * @memberof CharacterInspectMessageApiListCharacterInspectMessages
     */
    readonly groupBy?: string

    /**
     * Rows to limit in response (Default: 10,000)
     * @type {string}
     * @memberof CharacterInspectMessageApiListCharacterInspectMessages
     */
    readonly limit?: string

    /**
     * Order by [field]
     * @type {string}
     * @memberof CharacterInspectMessageApiListCharacterInspectMessages
     */
    readonly orderBy?: string

    /**
     * Order by field direction
     * @type {string}
     * @memberof CharacterInspectMessageApiListCharacterInspectMessages
     */
    readonly orderDirection?: string

    /**
     * Column names [.] separated to fetch specific fields in response
     * @type {string}
     * @memberof CharacterInspectMessageApiListCharacterInspectMessages
     */
    readonly select?: string
}

/**
 * Request parameters for updateCharacterInspectMessage operation in CharacterInspectMessageApi.
 * @export
 * @interface CharacterInspectMessageApiUpdateCharacterInspectMessageRequest
 */
export interface CharacterInspectMessageApiUpdateCharacterInspectMessageRequest {
    /**
     * Id
     * @type {number}
     * @memberof CharacterInspectMessageApiUpdateCharacterInspectMessage
     */
    readonly id: number

    /**
     * CharacterInspectMessage
     * @type {ModelsCharacterInspectMessage}
     * @memberof CharacterInspectMessageApiUpdateCharacterInspectMessage
     */
    readonly characterInspectMessage: ModelsCharacterInspectMessage
}

/**
 * CharacterInspectMessageApi - object-oriented interface
 * @export
 * @class CharacterInspectMessageApi
 * @extends {BaseAPI}
 */
export class CharacterInspectMessageApi extends BaseAPI {
    /**
     * 
     * @summary Creates CharacterInspectMessage
     * @param {CharacterInspectMessageApiCreateCharacterInspectMessageRequest} requestParameters Request parameters.
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof CharacterInspectMessageApi
     */
    public createCharacterInspectMessage(requestParameters: CharacterInspectMessageApiCreateCharacterInspectMessageRequest, options?: any) {
        return CharacterInspectMessageApiFp(this.configuration).createCharacterInspectMessage(requestParameters.characterInspectMessage, options).then((request) => request(this.axios, this.basePath));
    }

    /**
     * 
     * @summary Deletes CharacterInspectMessage
     * @param {CharacterInspectMessageApiDeleteCharacterInspectMessageRequest} requestParameters Request parameters.
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof CharacterInspectMessageApi
     */
    public deleteCharacterInspectMessage(requestParameters: CharacterInspectMessageApiDeleteCharacterInspectMessageRequest, options?: any) {
        return CharacterInspectMessageApiFp(this.configuration).deleteCharacterInspectMessage(requestParameters.id, options).then((request) => request(this.axios, this.basePath));
    }

    /**
     * 
     * @summary Gets CharacterInspectMessage
     * @param {CharacterInspectMessageApiGetCharacterInspectMessageRequest} requestParameters Request parameters.
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof CharacterInspectMessageApi
     */
    public getCharacterInspectMessage(requestParameters: CharacterInspectMessageApiGetCharacterInspectMessageRequest, options?: any) {
        return CharacterInspectMessageApiFp(this.configuration).getCharacterInspectMessage(requestParameters.id, requestParameters.includes, requestParameters.select, options).then((request) => request(this.axios, this.basePath));
    }

    /**
     * 
     * @summary Gets CharacterInspectMessages in bulk
     * @param {CharacterInspectMessageApiGetCharacterInspectMessagesBulkRequest} requestParameters Request parameters.
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof CharacterInspectMessageApi
     */
    public getCharacterInspectMessagesBulk(requestParameters: CharacterInspectMessageApiGetCharacterInspectMessagesBulkRequest, options?: any) {
        return CharacterInspectMessageApiFp(this.configuration).getCharacterInspectMessagesBulk(requestParameters.body, options).then((request) => request(this.axios, this.basePath));
    }

    /**
     * 
     * @summary Lists CharacterInspectMessages
     * @param {CharacterInspectMessageApiListCharacterInspectMessagesRequest} requestParameters Request parameters.
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof CharacterInspectMessageApi
     */
    public listCharacterInspectMessages(requestParameters: CharacterInspectMessageApiListCharacterInspectMessagesRequest = {}, options?: any) {
        return CharacterInspectMessageApiFp(this.configuration).listCharacterInspectMessages(requestParameters.includes, requestParameters.where, requestParameters.whereOr, requestParameters.groupBy, requestParameters.limit, requestParameters.orderBy, requestParameters.orderDirection, requestParameters.select, options).then((request) => request(this.axios, this.basePath));
    }

    /**
     * 
     * @summary Updates CharacterInspectMessage
     * @param {CharacterInspectMessageApiUpdateCharacterInspectMessageRequest} requestParameters Request parameters.
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof CharacterInspectMessageApi
     */
    public updateCharacterInspectMessage(requestParameters: CharacterInspectMessageApiUpdateCharacterInspectMessageRequest, options?: any) {
        return CharacterInspectMessageApiFp(this.configuration).updateCharacterInspectMessage(requestParameters.id, requestParameters.characterInspectMessage, options).then((request) => request(this.axios, this.basePath));
    }
}
