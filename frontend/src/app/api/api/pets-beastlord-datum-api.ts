import globalAxios, { AxiosPromise, AxiosInstance } from 'axios';
import { Configuration } from '../configuration';
// Some imports not used depending on template conditions
// @ts-ignore
import { BASE_PATH, COLLECTION_FORMATS, RequestArgs, BaseAPI, RequiredError } from '../base';
// @ts-ignore
import { CrudcontrollersBulkFetchByIdsGetRequest } from '../models';
// @ts-ignore
import { ModelsPetsBeastlordDatum } from '../models';
export const PetsBeastlordDatumApiAxiosParamCreator = function (configuration?: Configuration) {
    return {
        createPetsBeastlordDatum: async (petsBeastlordDatum: ModelsPetsBeastlordDatum, options: any = {}): Promise<RequestArgs> => {
            if (petsBeastlordDatum === null || petsBeastlordDatum === undefined) {
                throw new RequiredError('petsBeastlordDatum','Required parameter petsBeastlordDatum was null or undefined when calling createPetsBeastlordDatum.');
            }
            const localVarPath = `/pets_beastlord_datum`;
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
            const nonString = typeof petsBeastlordDatum !== 'string';
            const needsSerialization = nonString && configuration && configuration.isJsonMime
                ? configuration.isJsonMime(localVarRequestOptions.headers['Content-Type'])
                : nonString;
            localVarRequestOptions.data =  needsSerialization
                ? JSON.stringify(petsBeastlordDatum !== undefined ? petsBeastlordDatum : {})
                : (petsBeastlordDatum || "");
            return {
                url: localVarUrlObj.pathname + localVarUrlObj.search + localVarUrlObj.hash,
                options: localVarRequestOptions,
            };
        },
        deletePetsBeastlordDatum: async (id: number, options: any = {}): Promise<RequestArgs> => {
            if (id === null || id === undefined) {
                throw new RequiredError('id','Required parameter id was null or undefined when calling deletePetsBeastlordDatum.');
            }
            const localVarPath = `/pets_beastlord_datum/{id}`
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
        getPetsBeastlordDataBulk: async (body: CrudcontrollersBulkFetchByIdsGetRequest, options: any = {}): Promise<RequestArgs> => {
            if (body === null || body === undefined) {
                throw new RequiredError('body','Required parameter body was null or undefined when calling getPetsBeastlordDataBulk.');
            }
            const localVarPath = `/pets_beastlord_data/bulk`;
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
        getPetsBeastlordDatum: async (id: number, includes?: string, select?: string, options: any = {}): Promise<RequestArgs> => {
            if (id === null || id === undefined) {
                throw new RequiredError('id','Required parameter id was null or undefined when calling getPetsBeastlordDatum.');
            }
            const localVarPath = `/pets_beastlord_datum/{id}`
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
        listPetsBeastlordData: async (includes?: string, where?: string, whereOr?: string, groupBy?: string, limit?: string, page?: number, orderBy?: string, orderDirection?: string, select?: string, options: any = {}): Promise<RequestArgs> => {
            const localVarPath = `/pets_beastlord_data`;
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
        updatePetsBeastlordDatum: async (id: number, petsBeastlordDatum: ModelsPetsBeastlordDatum, options: any = {}): Promise<RequestArgs> => {
            if (id === null || id === undefined) {
                throw new RequiredError('id','Required parameter id was null or undefined when calling updatePetsBeastlordDatum.');
            }
            if (petsBeastlordDatum === null || petsBeastlordDatum === undefined) {
                throw new RequiredError('petsBeastlordDatum','Required parameter petsBeastlordDatum was null or undefined when calling updatePetsBeastlordDatum.');
            }
            const localVarPath = `/pets_beastlord_datum/{id}`
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
            const nonString = typeof petsBeastlordDatum !== 'string';
            const needsSerialization = nonString && configuration && configuration.isJsonMime
                ? configuration.isJsonMime(localVarRequestOptions.headers['Content-Type'])
                : nonString;
            localVarRequestOptions.data =  needsSerialization
                ? JSON.stringify(petsBeastlordDatum !== undefined ? petsBeastlordDatum : {})
                : (petsBeastlordDatum || "");
            return {
                url: localVarUrlObj.pathname + localVarUrlObj.search + localVarUrlObj.hash,
                options: localVarRequestOptions,
            };
        },
    }
};
export const PetsBeastlordDatumApiFp = function(configuration?: Configuration) {
    return {
        async createPetsBeastlordDatum(petsBeastlordDatum: ModelsPetsBeastlordDatum, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsPetsBeastlordDatum>>> {
            const localVarAxiosArgs = await PetsBeastlordDatumApiAxiosParamCreator(configuration).createPetsBeastlordDatum(petsBeastlordDatum, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        async deletePetsBeastlordDatum(id: number, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<string>> {
            const localVarAxiosArgs = await PetsBeastlordDatumApiAxiosParamCreator(configuration).deletePetsBeastlordDatum(id, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        async getPetsBeastlordDataBulk(body: CrudcontrollersBulkFetchByIdsGetRequest, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsPetsBeastlordDatum>>> {
            const localVarAxiosArgs = await PetsBeastlordDatumApiAxiosParamCreator(configuration).getPetsBeastlordDataBulk(body, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        async getPetsBeastlordDatum(id: number, includes?: string, select?: string, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsPetsBeastlordDatum>>> {
            const localVarAxiosArgs = await PetsBeastlordDatumApiAxiosParamCreator(configuration).getPetsBeastlordDatum(id, includes, select, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        async listPetsBeastlordData(includes?: string, where?: string, whereOr?: string, groupBy?: string, limit?: string, page?: number, orderBy?: string, orderDirection?: string, select?: string, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsPetsBeastlordDatum>>> {
            const localVarAxiosArgs = await PetsBeastlordDatumApiAxiosParamCreator(configuration).listPetsBeastlordData(includes, where, whereOr, groupBy, limit, page, orderBy, orderDirection, select, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        async updatePetsBeastlordDatum(id: number, petsBeastlordDatum: ModelsPetsBeastlordDatum, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsPetsBeastlordDatum>>> {
            const localVarAxiosArgs = await PetsBeastlordDatumApiAxiosParamCreator(configuration).updatePetsBeastlordDatum(id, petsBeastlordDatum, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
    }
};
export const PetsBeastlordDatumApiFactory = function (configuration?: Configuration, basePath?: string, axios?: AxiosInstance) {
    return {
        createPetsBeastlordDatum(petsBeastlordDatum: ModelsPetsBeastlordDatum, options?: any): AxiosPromise<Array<ModelsPetsBeastlordDatum>> {
            return PetsBeastlordDatumApiFp(configuration).createPetsBeastlordDatum(petsBeastlordDatum, options).then((request) => request(axios, basePath));
        },
        deletePetsBeastlordDatum(id: number, options?: any): AxiosPromise<string> {
            return PetsBeastlordDatumApiFp(configuration).deletePetsBeastlordDatum(id, options).then((request) => request(axios, basePath));
        },
        getPetsBeastlordDataBulk(body: CrudcontrollersBulkFetchByIdsGetRequest, options?: any): AxiosPromise<Array<ModelsPetsBeastlordDatum>> {
            return PetsBeastlordDatumApiFp(configuration).getPetsBeastlordDataBulk(body, options).then((request) => request(axios, basePath));
        },
        getPetsBeastlordDatum(id: number, includes?: string, select?: string, options?: any): AxiosPromise<Array<ModelsPetsBeastlordDatum>> {
            return PetsBeastlordDatumApiFp(configuration).getPetsBeastlordDatum(id, includes, select, options).then((request) => request(axios, basePath));
        },
        listPetsBeastlordData(includes?: string, where?: string, whereOr?: string, groupBy?: string, limit?: string, page?: number, orderBy?: string, orderDirection?: string, select?: string, options?: any): AxiosPromise<Array<ModelsPetsBeastlordDatum>> {
            return PetsBeastlordDatumApiFp(configuration).listPetsBeastlordData(includes, where, whereOr, groupBy, limit, page, orderBy, orderDirection, select, options).then((request) => request(axios, basePath));
        },
        updatePetsBeastlordDatum(id: number, petsBeastlordDatum: ModelsPetsBeastlordDatum, options?: any): AxiosPromise<Array<ModelsPetsBeastlordDatum>> {
            return PetsBeastlordDatumApiFp(configuration).updatePetsBeastlordDatum(id, petsBeastlordDatum, options).then((request) => request(axios, basePath));
        },
    };
};
export interface PetsBeastlordDatumApiCreatePetsBeastlordDatumRequest {
    readonly petsBeastlordDatum: ModelsPetsBeastlordDatum
}
export interface PetsBeastlordDatumApiDeletePetsBeastlordDatumRequest {
    readonly id: number
}
export interface PetsBeastlordDatumApiGetPetsBeastlordDataBulkRequest {
    readonly body: CrudcontrollersBulkFetchByIdsGetRequest
}
export interface PetsBeastlordDatumApiGetPetsBeastlordDatumRequest {
    readonly id: number
    readonly includes?: string
    readonly select?: string
}
export interface PetsBeastlordDatumApiListPetsBeastlordDataRequest {
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
export interface PetsBeastlordDatumApiUpdatePetsBeastlordDatumRequest {
    readonly id: number
    readonly petsBeastlordDatum: ModelsPetsBeastlordDatum
}
export class PetsBeastlordDatumApi extends BaseAPI {
    public createPetsBeastlordDatum(requestParameters: PetsBeastlordDatumApiCreatePetsBeastlordDatumRequest, options?: any) {
        return PetsBeastlordDatumApiFp(this.configuration).createPetsBeastlordDatum(requestParameters.petsBeastlordDatum, options).then((request) => request(this.axios, this.basePath));
    }
    public deletePetsBeastlordDatum(requestParameters: PetsBeastlordDatumApiDeletePetsBeastlordDatumRequest, options?: any) {
        return PetsBeastlordDatumApiFp(this.configuration).deletePetsBeastlordDatum(requestParameters.id, options).then((request) => request(this.axios, this.basePath));
    }
    public getPetsBeastlordDataBulk(requestParameters: PetsBeastlordDatumApiGetPetsBeastlordDataBulkRequest, options?: any) {
        return PetsBeastlordDatumApiFp(this.configuration).getPetsBeastlordDataBulk(requestParameters.body, options).then((request) => request(this.axios, this.basePath));
    }
    public getPetsBeastlordDatum(requestParameters: PetsBeastlordDatumApiGetPetsBeastlordDatumRequest, options?: any) {
        return PetsBeastlordDatumApiFp(this.configuration).getPetsBeastlordDatum(requestParameters.id, requestParameters.includes, requestParameters.select, options).then((request) => request(this.axios, this.basePath));
    }
    public listPetsBeastlordData(requestParameters: PetsBeastlordDatumApiListPetsBeastlordDataRequest = {}, options?: any) {
        return PetsBeastlordDatumApiFp(this.configuration).listPetsBeastlordData(requestParameters.includes, requestParameters.where, requestParameters.whereOr, requestParameters.groupBy, requestParameters.limit, requestParameters.page, requestParameters.orderBy, requestParameters.orderDirection, requestParameters.select, options).then((request) => request(this.axios, this.basePath));
    }
    public updatePetsBeastlordDatum(requestParameters: PetsBeastlordDatumApiUpdatePetsBeastlordDatumRequest, options?: any) {
        return PetsBeastlordDatumApiFp(this.configuration).updatePetsBeastlordDatum(requestParameters.id, requestParameters.petsBeastlordDatum, options).then((request) => request(this.axios, this.basePath));
    }
}
