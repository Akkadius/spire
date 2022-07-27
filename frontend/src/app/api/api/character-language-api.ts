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
import { ModelsCharacterLanguage } from '../models';
/**
 * CharacterLanguageApi - axios parameter creator
 * @export
 */
export const CharacterLanguageApiAxiosParamCreator = function (configuration?: Configuration) {
    return {
        /**
         * 
         * @summary Creates CharacterLanguage
         * @param {ModelsCharacterLanguage} characterLanguage CharacterLanguage
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        createCharacterLanguage: async (characterLanguage: ModelsCharacterLanguage, options: any = {}): Promise<RequestArgs> => {
            // verify required parameter 'characterLanguage' is not null or undefined
            if (characterLanguage === null || characterLanguage === undefined) {
                throw new RequiredError('characterLanguage','Required parameter characterLanguage was null or undefined when calling createCharacterLanguage.');
            }
            const localVarPath = `/character_language`;
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
            const nonString = typeof characterLanguage !== 'string';
            const needsSerialization = nonString && configuration && configuration.isJsonMime
                ? configuration.isJsonMime(localVarRequestOptions.headers['Content-Type'])
                : nonString;
            localVarRequestOptions.data =  needsSerialization
                ? JSON.stringify(characterLanguage !== undefined ? characterLanguage : {})
                : (characterLanguage || "");

            return {
                url: localVarUrlObj.pathname + localVarUrlObj.search + localVarUrlObj.hash,
                options: localVarRequestOptions,
            };
        },
        /**
         * 
         * @summary Deletes CharacterLanguage
         * @param {number} id id
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        deleteCharacterLanguage: async (id: number, options: any = {}): Promise<RequestArgs> => {
            // verify required parameter 'id' is not null or undefined
            if (id === null || id === undefined) {
                throw new RequiredError('id','Required parameter id was null or undefined when calling deleteCharacterLanguage.');
            }
            const localVarPath = `/character_language/{id}`
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
         * @summary Gets CharacterLanguage
         * @param {number} id Id
         * @param {string} [includes] Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names 
         * @param {string} [select] Column names [.] separated to fetch specific fields in response
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        getCharacterLanguage: async (id: number, includes?: string, select?: string, options: any = {}): Promise<RequestArgs> => {
            // verify required parameter 'id' is not null or undefined
            if (id === null || id === undefined) {
                throw new RequiredError('id','Required parameter id was null or undefined when calling getCharacterLanguage.');
            }
            const localVarPath = `/character_language/{id}`
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
         * @summary Gets CharacterLanguages in bulk
         * @param {CrudcontrollersBulkFetchByIdsGetRequest} body body
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        getCharacterLanguagesBulk: async (body: CrudcontrollersBulkFetchByIdsGetRequest, options: any = {}): Promise<RequestArgs> => {
            // verify required parameter 'body' is not null or undefined
            if (body === null || body === undefined) {
                throw new RequiredError('body','Required parameter body was null or undefined when calling getCharacterLanguagesBulk.');
            }
            const localVarPath = `/character_languages/bulk`;
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
         * @summary Lists CharacterLanguages
         * @param {string} [includes] Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names 
         * @param {string} [where] Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2
         * @param {string} [whereOr] Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2
         * @param {string} [groupBy] Group by field. Multiple conditions [.] separated Example: field1.field2
         * @param {string} [limit] Rows to limit in response (Default: 10,000)
         * @param {number} [page] Pagination page
         * @param {string} [orderBy] Order by [field]
         * @param {string} [orderDirection] Order by field direction
         * @param {string} [select] Column names [.] separated to fetch specific fields in response
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        listCharacterLanguages: async (includes?: string, where?: string, whereOr?: string, groupBy?: string, limit?: string, page?: number, orderBy?: string, orderDirection?: string, select?: string, options: any = {}): Promise<RequestArgs> => {
            const localVarPath = `/character_languages`;
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
        /**
         * 
         * @summary Updates CharacterLanguage
         * @param {number} id Id
         * @param {ModelsCharacterLanguage} characterLanguage CharacterLanguage
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        updateCharacterLanguage: async (id: number, characterLanguage: ModelsCharacterLanguage, options: any = {}): Promise<RequestArgs> => {
            // verify required parameter 'id' is not null or undefined
            if (id === null || id === undefined) {
                throw new RequiredError('id','Required parameter id was null or undefined when calling updateCharacterLanguage.');
            }
            // verify required parameter 'characterLanguage' is not null or undefined
            if (characterLanguage === null || characterLanguage === undefined) {
                throw new RequiredError('characterLanguage','Required parameter characterLanguage was null or undefined when calling updateCharacterLanguage.');
            }
            const localVarPath = `/character_language/{id}`
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
            const nonString = typeof characterLanguage !== 'string';
            const needsSerialization = nonString && configuration && configuration.isJsonMime
                ? configuration.isJsonMime(localVarRequestOptions.headers['Content-Type'])
                : nonString;
            localVarRequestOptions.data =  needsSerialization
                ? JSON.stringify(characterLanguage !== undefined ? characterLanguage : {})
                : (characterLanguage || "");

            return {
                url: localVarUrlObj.pathname + localVarUrlObj.search + localVarUrlObj.hash,
                options: localVarRequestOptions,
            };
        },
    }
};

/**
 * CharacterLanguageApi - functional programming interface
 * @export
 */
export const CharacterLanguageApiFp = function(configuration?: Configuration) {
    return {
        /**
         * 
         * @summary Creates CharacterLanguage
         * @param {ModelsCharacterLanguage} characterLanguage CharacterLanguage
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async createCharacterLanguage(characterLanguage: ModelsCharacterLanguage, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsCharacterLanguage>>> {
            const localVarAxiosArgs = await CharacterLanguageApiAxiosParamCreator(configuration).createCharacterLanguage(characterLanguage, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        /**
         * 
         * @summary Deletes CharacterLanguage
         * @param {number} id id
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async deleteCharacterLanguage(id: number, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<string>> {
            const localVarAxiosArgs = await CharacterLanguageApiAxiosParamCreator(configuration).deleteCharacterLanguage(id, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        /**
         * 
         * @summary Gets CharacterLanguage
         * @param {number} id Id
         * @param {string} [includes] Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names 
         * @param {string} [select] Column names [.] separated to fetch specific fields in response
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async getCharacterLanguage(id: number, includes?: string, select?: string, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsCharacterLanguage>>> {
            const localVarAxiosArgs = await CharacterLanguageApiAxiosParamCreator(configuration).getCharacterLanguage(id, includes, select, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        /**
         * 
         * @summary Gets CharacterLanguages in bulk
         * @param {CrudcontrollersBulkFetchByIdsGetRequest} body body
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async getCharacterLanguagesBulk(body: CrudcontrollersBulkFetchByIdsGetRequest, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsCharacterLanguage>>> {
            const localVarAxiosArgs = await CharacterLanguageApiAxiosParamCreator(configuration).getCharacterLanguagesBulk(body, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        /**
         * 
         * @summary Lists CharacterLanguages
         * @param {string} [includes] Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names 
         * @param {string} [where] Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2
         * @param {string} [whereOr] Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2
         * @param {string} [groupBy] Group by field. Multiple conditions [.] separated Example: field1.field2
         * @param {string} [limit] Rows to limit in response (Default: 10,000)
         * @param {number} [page] Pagination page
         * @param {string} [orderBy] Order by [field]
         * @param {string} [orderDirection] Order by field direction
         * @param {string} [select] Column names [.] separated to fetch specific fields in response
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async listCharacterLanguages(includes?: string, where?: string, whereOr?: string, groupBy?: string, limit?: string, page?: number, orderBy?: string, orderDirection?: string, select?: string, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsCharacterLanguage>>> {
            const localVarAxiosArgs = await CharacterLanguageApiAxiosParamCreator(configuration).listCharacterLanguages(includes, where, whereOr, groupBy, limit, page, orderBy, orderDirection, select, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        /**
         * 
         * @summary Updates CharacterLanguage
         * @param {number} id Id
         * @param {ModelsCharacterLanguage} characterLanguage CharacterLanguage
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async updateCharacterLanguage(id: number, characterLanguage: ModelsCharacterLanguage, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsCharacterLanguage>>> {
            const localVarAxiosArgs = await CharacterLanguageApiAxiosParamCreator(configuration).updateCharacterLanguage(id, characterLanguage, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
    }
};

/**
 * CharacterLanguageApi - factory interface
 * @export
 */
export const CharacterLanguageApiFactory = function (configuration?: Configuration, basePath?: string, axios?: AxiosInstance) {
    return {
        /**
         * 
         * @summary Creates CharacterLanguage
         * @param {ModelsCharacterLanguage} characterLanguage CharacterLanguage
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        createCharacterLanguage(characterLanguage: ModelsCharacterLanguage, options?: any): AxiosPromise<Array<ModelsCharacterLanguage>> {
            return CharacterLanguageApiFp(configuration).createCharacterLanguage(characterLanguage, options).then((request) => request(axios, basePath));
        },
        /**
         * 
         * @summary Deletes CharacterLanguage
         * @param {number} id id
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        deleteCharacterLanguage(id: number, options?: any): AxiosPromise<string> {
            return CharacterLanguageApiFp(configuration).deleteCharacterLanguage(id, options).then((request) => request(axios, basePath));
        },
        /**
         * 
         * @summary Gets CharacterLanguage
         * @param {number} id Id
         * @param {string} [includes] Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names 
         * @param {string} [select] Column names [.] separated to fetch specific fields in response
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        getCharacterLanguage(id: number, includes?: string, select?: string, options?: any): AxiosPromise<Array<ModelsCharacterLanguage>> {
            return CharacterLanguageApiFp(configuration).getCharacterLanguage(id, includes, select, options).then((request) => request(axios, basePath));
        },
        /**
         * 
         * @summary Gets CharacterLanguages in bulk
         * @param {CrudcontrollersBulkFetchByIdsGetRequest} body body
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        getCharacterLanguagesBulk(body: CrudcontrollersBulkFetchByIdsGetRequest, options?: any): AxiosPromise<Array<ModelsCharacterLanguage>> {
            return CharacterLanguageApiFp(configuration).getCharacterLanguagesBulk(body, options).then((request) => request(axios, basePath));
        },
        /**
         * 
         * @summary Lists CharacterLanguages
         * @param {string} [includes] Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names 
         * @param {string} [where] Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2
         * @param {string} [whereOr] Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2
         * @param {string} [groupBy] Group by field. Multiple conditions [.] separated Example: field1.field2
         * @param {string} [limit] Rows to limit in response (Default: 10,000)
         * @param {number} [page] Pagination page
         * @param {string} [orderBy] Order by [field]
         * @param {string} [orderDirection] Order by field direction
         * @param {string} [select] Column names [.] separated to fetch specific fields in response
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        listCharacterLanguages(includes?: string, where?: string, whereOr?: string, groupBy?: string, limit?: string, page?: number, orderBy?: string, orderDirection?: string, select?: string, options?: any): AxiosPromise<Array<ModelsCharacterLanguage>> {
            return CharacterLanguageApiFp(configuration).listCharacterLanguages(includes, where, whereOr, groupBy, limit, page, orderBy, orderDirection, select, options).then((request) => request(axios, basePath));
        },
        /**
         * 
         * @summary Updates CharacterLanguage
         * @param {number} id Id
         * @param {ModelsCharacterLanguage} characterLanguage CharacterLanguage
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        updateCharacterLanguage(id: number, characterLanguage: ModelsCharacterLanguage, options?: any): AxiosPromise<Array<ModelsCharacterLanguage>> {
            return CharacterLanguageApiFp(configuration).updateCharacterLanguage(id, characterLanguage, options).then((request) => request(axios, basePath));
        },
    };
};

/**
 * Request parameters for createCharacterLanguage operation in CharacterLanguageApi.
 * @export
 * @interface CharacterLanguageApiCreateCharacterLanguageRequest
 */
export interface CharacterLanguageApiCreateCharacterLanguageRequest {
    /**
     * CharacterLanguage
     * @type {ModelsCharacterLanguage}
     * @memberof CharacterLanguageApiCreateCharacterLanguage
     */
    readonly characterLanguage: ModelsCharacterLanguage
}

/**
 * Request parameters for deleteCharacterLanguage operation in CharacterLanguageApi.
 * @export
 * @interface CharacterLanguageApiDeleteCharacterLanguageRequest
 */
export interface CharacterLanguageApiDeleteCharacterLanguageRequest {
    /**
     * id
     * @type {number}
     * @memberof CharacterLanguageApiDeleteCharacterLanguage
     */
    readonly id: number
}

/**
 * Request parameters for getCharacterLanguage operation in CharacterLanguageApi.
 * @export
 * @interface CharacterLanguageApiGetCharacterLanguageRequest
 */
export interface CharacterLanguageApiGetCharacterLanguageRequest {
    /**
     * Id
     * @type {number}
     * @memberof CharacterLanguageApiGetCharacterLanguage
     */
    readonly id: number

    /**
     * Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names 
     * @type {string}
     * @memberof CharacterLanguageApiGetCharacterLanguage
     */
    readonly includes?: string

    /**
     * Column names [.] separated to fetch specific fields in response
     * @type {string}
     * @memberof CharacterLanguageApiGetCharacterLanguage
     */
    readonly select?: string
}

/**
 * Request parameters for getCharacterLanguagesBulk operation in CharacterLanguageApi.
 * @export
 * @interface CharacterLanguageApiGetCharacterLanguagesBulkRequest
 */
export interface CharacterLanguageApiGetCharacterLanguagesBulkRequest {
    /**
     * body
     * @type {CrudcontrollersBulkFetchByIdsGetRequest}
     * @memberof CharacterLanguageApiGetCharacterLanguagesBulk
     */
    readonly body: CrudcontrollersBulkFetchByIdsGetRequest
}

/**
 * Request parameters for listCharacterLanguages operation in CharacterLanguageApi.
 * @export
 * @interface CharacterLanguageApiListCharacterLanguagesRequest
 */
export interface CharacterLanguageApiListCharacterLanguagesRequest {
    /**
     * Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names 
     * @type {string}
     * @memberof CharacterLanguageApiListCharacterLanguages
     */
    readonly includes?: string

    /**
     * Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2
     * @type {string}
     * @memberof CharacterLanguageApiListCharacterLanguages
     */
    readonly where?: string

    /**
     * Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2
     * @type {string}
     * @memberof CharacterLanguageApiListCharacterLanguages
     */
    readonly whereOr?: string

    /**
     * Group by field. Multiple conditions [.] separated Example: field1.field2
     * @type {string}
     * @memberof CharacterLanguageApiListCharacterLanguages
     */
    readonly groupBy?: string

    /**
     * Rows to limit in response (Default: 10,000)
     * @type {string}
     * @memberof CharacterLanguageApiListCharacterLanguages
     */
    readonly limit?: string

    /**
     * Pagination page
     * @type {number}
     * @memberof CharacterLanguageApiListCharacterLanguages
     */
    readonly page?: number

    /**
     * Order by [field]
     * @type {string}
     * @memberof CharacterLanguageApiListCharacterLanguages
     */
    readonly orderBy?: string

    /**
     * Order by field direction
     * @type {string}
     * @memberof CharacterLanguageApiListCharacterLanguages
     */
    readonly orderDirection?: string

    /**
     * Column names [.] separated to fetch specific fields in response
     * @type {string}
     * @memberof CharacterLanguageApiListCharacterLanguages
     */
    readonly select?: string
}

/**
 * Request parameters for updateCharacterLanguage operation in CharacterLanguageApi.
 * @export
 * @interface CharacterLanguageApiUpdateCharacterLanguageRequest
 */
export interface CharacterLanguageApiUpdateCharacterLanguageRequest {
    /**
     * Id
     * @type {number}
     * @memberof CharacterLanguageApiUpdateCharacterLanguage
     */
    readonly id: number

    /**
     * CharacterLanguage
     * @type {ModelsCharacterLanguage}
     * @memberof CharacterLanguageApiUpdateCharacterLanguage
     */
    readonly characterLanguage: ModelsCharacterLanguage
}

/**
 * CharacterLanguageApi - object-oriented interface
 * @export
 * @class CharacterLanguageApi
 * @extends {BaseAPI}
 */
export class CharacterLanguageApi extends BaseAPI {
    /**
     * 
     * @summary Creates CharacterLanguage
     * @param {CharacterLanguageApiCreateCharacterLanguageRequest} requestParameters Request parameters.
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof CharacterLanguageApi
     */
    public createCharacterLanguage(requestParameters: CharacterLanguageApiCreateCharacterLanguageRequest, options?: any) {
        return CharacterLanguageApiFp(this.configuration).createCharacterLanguage(requestParameters.characterLanguage, options).then((request) => request(this.axios, this.basePath));
    }

    /**
     * 
     * @summary Deletes CharacterLanguage
     * @param {CharacterLanguageApiDeleteCharacterLanguageRequest} requestParameters Request parameters.
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof CharacterLanguageApi
     */
    public deleteCharacterLanguage(requestParameters: CharacterLanguageApiDeleteCharacterLanguageRequest, options?: any) {
        return CharacterLanguageApiFp(this.configuration).deleteCharacterLanguage(requestParameters.id, options).then((request) => request(this.axios, this.basePath));
    }

    /**
     * 
     * @summary Gets CharacterLanguage
     * @param {CharacterLanguageApiGetCharacterLanguageRequest} requestParameters Request parameters.
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof CharacterLanguageApi
     */
    public getCharacterLanguage(requestParameters: CharacterLanguageApiGetCharacterLanguageRequest, options?: any) {
        return CharacterLanguageApiFp(this.configuration).getCharacterLanguage(requestParameters.id, requestParameters.includes, requestParameters.select, options).then((request) => request(this.axios, this.basePath));
    }

    /**
     * 
     * @summary Gets CharacterLanguages in bulk
     * @param {CharacterLanguageApiGetCharacterLanguagesBulkRequest} requestParameters Request parameters.
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof CharacterLanguageApi
     */
    public getCharacterLanguagesBulk(requestParameters: CharacterLanguageApiGetCharacterLanguagesBulkRequest, options?: any) {
        return CharacterLanguageApiFp(this.configuration).getCharacterLanguagesBulk(requestParameters.body, options).then((request) => request(this.axios, this.basePath));
    }

    /**
     * 
     * @summary Lists CharacterLanguages
     * @param {CharacterLanguageApiListCharacterLanguagesRequest} requestParameters Request parameters.
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof CharacterLanguageApi
     */
    public listCharacterLanguages(requestParameters: CharacterLanguageApiListCharacterLanguagesRequest = {}, options?: any) {
        return CharacterLanguageApiFp(this.configuration).listCharacterLanguages(requestParameters.includes, requestParameters.where, requestParameters.whereOr, requestParameters.groupBy, requestParameters.limit, requestParameters.page, requestParameters.orderBy, requestParameters.orderDirection, requestParameters.select, options).then((request) => request(this.axios, this.basePath));
    }

    /**
     * 
     * @summary Updates CharacterLanguage
     * @param {CharacterLanguageApiUpdateCharacterLanguageRequest} requestParameters Request parameters.
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof CharacterLanguageApi
     */
    public updateCharacterLanguage(requestParameters: CharacterLanguageApiUpdateCharacterLanguageRequest, options?: any) {
        return CharacterLanguageApiFp(this.configuration).updateCharacterLanguage(requestParameters.id, requestParameters.characterLanguage, options).then((request) => request(this.axios, this.basePath));
    }
}
