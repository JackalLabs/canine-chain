export interface JklminingMined {
    /** @format uint64 */
    id?: string;
    datasize?: string;
    hash?: string;
    pcount?: string;
}
export interface JklminingMinerClaims {
    hash?: string;
    creator?: string;
}
export interface JklminingMiners {
    address?: string;
    ip?: string;
    creator?: string;
}
export declare type JklminingMsgAllowSaveResponse = object;
export declare type JklminingMsgClaimSaveResponse = object;
export declare type JklminingMsgCreateMinerClaimsResponse = object;
export declare type JklminingMsgCreateMinersResponse = object;
export declare type JklminingMsgCreateSaveRequestsResponse = object;
export declare type JklminingMsgDeleteMinerClaimsResponse = object;
export declare type JklminingMsgDeleteMinersResponse = object;
export declare type JklminingMsgDeleteSaveRequestsResponse = object;
export declare type JklminingMsgUpdateMinerClaimsResponse = object;
export declare type JklminingMsgUpdateMinersResponse = object;
export declare type JklminingMsgUpdateSaveRequestsResponse = object;
/**
 * Params defines the parameters for the module.
 */
export declare type JklminingParams = object;
export interface JklminingQueryAllMinedResponse {
    Mined?: JklminingMined[];
    /**
     * PageResponse is to be embedded in gRPC response messages where the
     * corresponding request message has used PageRequest.
     *
     *  message SomeResponse {
     *          repeated Bar results = 1;
     *          PageResponse page = 2;
     *  }
     */
    pagination?: V1Beta1PageResponse;
}
export interface JklminingQueryAllMinerClaimsResponse {
    minerClaims?: JklminingMinerClaims[];
    /**
     * PageResponse is to be embedded in gRPC response messages where the
     * corresponding request message has used PageRequest.
     *
     *  message SomeResponse {
     *          repeated Bar results = 1;
     *          PageResponse page = 2;
     *  }
     */
    pagination?: V1Beta1PageResponse;
}
export interface JklminingQueryAllMinersResponse {
    miners?: JklminingMiners[];
    /**
     * PageResponse is to be embedded in gRPC response messages where the
     * corresponding request message has used PageRequest.
     *
     *  message SomeResponse {
     *          repeated Bar results = 1;
     *          PageResponse page = 2;
     *  }
     */
    pagination?: V1Beta1PageResponse;
}
export interface JklminingQueryAllSaveRequestsResponse {
    saveRequests?: JklminingSaveRequests[];
    /**
     * PageResponse is to be embedded in gRPC response messages where the
     * corresponding request message has used PageRequest.
     *
     *  message SomeResponse {
     *          repeated Bar results = 1;
     *          PageResponse page = 2;
     *  }
     */
    pagination?: V1Beta1PageResponse;
}
export declare type JklminingQueryCheckMinerIndexResponse = object;
export interface JklminingQueryGetMinedResponse {
    Mined?: JklminingMined;
}
export interface JklminingQueryGetMinerClaimsResponse {
    minerClaims?: JklminingMinerClaims;
}
export declare type JklminingQueryGetMinerIndexResponse = object;
export interface JklminingQueryGetMinerStartResponse {
    index?: string;
}
export interface JklminingQueryGetMinersResponse {
    miners?: JklminingMiners;
}
export interface JklminingQueryGetSaveRequestsResponse {
    saveRequests?: JklminingSaveRequests;
}
/**
 * QueryParamsResponse is response type for the Query/Params RPC method.
 */
export interface JklminingQueryParamsResponse {
    /** params holds all the parameters of this module. */
    params?: JklminingParams;
}
export interface JklminingSaveRequests {
    index?: string;
    size?: string;
    approved?: string;
    creator?: string;
}
export interface ProtobufAny {
    "@type"?: string;
}
export interface RpcStatus {
    /** @format int32 */
    code?: number;
    message?: string;
    details?: ProtobufAny[];
}
/**
* message SomeRequest {
         Foo some_parameter = 1;
         PageRequest pagination = 2;
 }
*/
export interface V1Beta1PageRequest {
    /**
     * key is a value returned in PageResponse.next_key to begin
     * querying the next page most efficiently. Only one of offset or key
     * should be set.
     * @format byte
     */
    key?: string;
    /**
     * offset is a numeric offset that can be used when key is unavailable.
     * It is less efficient than using key. Only one of offset or key should
     * be set.
     * @format uint64
     */
    offset?: string;
    /**
     * limit is the total number of results to be returned in the result page.
     * If left empty it will default to a value to be set by each app.
     * @format uint64
     */
    limit?: string;
    /**
     * count_total is set to true  to indicate that the result set should include
     * a count of the total number of items available for pagination in UIs.
     * count_total is only respected when offset is used. It is ignored when key
     * is set.
     */
    count_total?: boolean;
    /**
     * reverse is set to true if results are to be returned in the descending order.
     *
     * Since: cosmos-sdk 0.43
     */
    reverse?: boolean;
}
/**
* PageResponse is to be embedded in gRPC response messages where the
corresponding request message has used PageRequest.

 message SomeResponse {
         repeated Bar results = 1;
         PageResponse page = 2;
 }
*/
export interface V1Beta1PageResponse {
    /** @format byte */
    next_key?: string;
    /** @format uint64 */
    total?: string;
}
export declare type QueryParamsType = Record<string | number, any>;
export declare type ResponseFormat = keyof Omit<Body, "body" | "bodyUsed">;
export interface FullRequestParams extends Omit<RequestInit, "body"> {
    /** set parameter to `true` for call `securityWorker` for this request */
    secure?: boolean;
    /** request path */
    path: string;
    /** content type of request body */
    type?: ContentType;
    /** query params */
    query?: QueryParamsType;
    /** format of response (i.e. response.json() -> format: "json") */
    format?: keyof Omit<Body, "body" | "bodyUsed">;
    /** request body */
    body?: unknown;
    /** base url */
    baseUrl?: string;
    /** request cancellation token */
    cancelToken?: CancelToken;
}
export declare type RequestParams = Omit<FullRequestParams, "body" | "method" | "query" | "path">;
export interface ApiConfig<SecurityDataType = unknown> {
    baseUrl?: string;
    baseApiParams?: Omit<RequestParams, "baseUrl" | "cancelToken" | "signal">;
    securityWorker?: (securityData: SecurityDataType) => RequestParams | void;
}
export interface HttpResponse<D extends unknown, E extends unknown = unknown> extends Response {
    data: D;
    error: E;
}
declare type CancelToken = Symbol | string | number;
export declare enum ContentType {
    Json = "application/json",
    FormData = "multipart/form-data",
    UrlEncoded = "application/x-www-form-urlencoded"
}
export declare class HttpClient<SecurityDataType = unknown> {
    baseUrl: string;
    private securityData;
    private securityWorker;
    private abortControllers;
    private baseApiParams;
    constructor(apiConfig?: ApiConfig<SecurityDataType>);
    setSecurityData: (data: SecurityDataType) => void;
    private addQueryParam;
    protected toQueryString(rawQuery?: QueryParamsType): string;
    protected addQueryParams(rawQuery?: QueryParamsType): string;
    private contentFormatters;
    private mergeRequestParams;
    private createAbortSignal;
    abortRequest: (cancelToken: CancelToken) => void;
    request: <T = any, E = any>({ body, secure, path, type, query, format, baseUrl, cancelToken, ...params }: FullRequestParams) => Promise<HttpResponse<T, E>>;
}
/**
 * @title jklmining/genesis.proto
 * @version version not set
 */
export declare class Api<SecurityDataType extends unknown> extends HttpClient<SecurityDataType> {
    /**
     * No description
     *
     * @tags Query
     * @name QueryCheckMinerIndex
     * @summary Queries a list of CheckMinerIndex items.
     * @request GET:/jackal-dao/canine/jklmining/check_miner_index
     */
    queryCheckMinerIndex: (params?: RequestParams) => Promise<HttpResponse<object, RpcStatus>>;
    /**
     * No description
     *
     * @tags Query
     * @name QueryGetMinerIndex
     * @summary Queries a list of GetMinerIndex items.
     * @request GET:/jackal-dao/canine/jklmining/get_miner_index/{index}
     */
    queryGetMinerIndex: (index: string, params?: RequestParams) => Promise<HttpResponse<object, RpcStatus>>;
    /**
     * No description
     *
     * @tags Query
     * @name QueryGetMinerStart
     * @summary Queries a list of GetMinerStart items.
     * @request GET:/jackal-dao/canine/jklmining/get_miner_start
     */
    queryGetMinerStart: (params?: RequestParams) => Promise<HttpResponse<JklminingQueryGetMinerStartResponse, RpcStatus>>;
    /**
     * No description
     *
     * @tags Query
     * @name QueryMinedAll
     * @summary Queries a list of Mined items.
     * @request GET:/jackal-dao/canine/jklmining/mined
     */
    queryMinedAll: (query?: {
        "pagination.key"?: string;
        "pagination.offset"?: string;
        "pagination.limit"?: string;
        "pagination.count_total"?: boolean;
        "pagination.reverse"?: boolean;
    }, params?: RequestParams) => Promise<HttpResponse<JklminingQueryAllMinedResponse, RpcStatus>>;
    /**
     * No description
     *
     * @tags Query
     * @name QueryMined
     * @summary Queries a Mined by id.
     * @request GET:/jackal-dao/canine/jklmining/mined/{id}
     */
    queryMined: (id: string, params?: RequestParams) => Promise<HttpResponse<JklminingQueryGetMinedResponse, RpcStatus>>;
    /**
     * No description
     *
     * @tags Query
     * @name QueryMinerClaimsAll
     * @summary Queries a list of MinerClaims items.
     * @request GET:/jackal-dao/canine/jklmining/miner_claims
     */
    queryMinerClaimsAll: (query?: {
        "pagination.key"?: string;
        "pagination.offset"?: string;
        "pagination.limit"?: string;
        "pagination.count_total"?: boolean;
        "pagination.reverse"?: boolean;
    }, params?: RequestParams) => Promise<HttpResponse<JklminingQueryAllMinerClaimsResponse, RpcStatus>>;
    /**
     * No description
     *
     * @tags Query
     * @name QueryMinerClaims
     * @summary Queries a MinerClaims by index.
     * @request GET:/jackal-dao/canine/jklmining/miner_claims/{hash}
     */
    queryMinerClaims: (hash: string, params?: RequestParams) => Promise<HttpResponse<JklminingQueryGetMinerClaimsResponse, RpcStatus>>;
    /**
     * No description
     *
     * @tags Query
     * @name QueryMinersAll
     * @summary Queries a list of Miners items.
     * @request GET:/jackal-dao/canine/jklmining/miners
     */
    queryMinersAll: (query?: {
        "pagination.key"?: string;
        "pagination.offset"?: string;
        "pagination.limit"?: string;
        "pagination.count_total"?: boolean;
        "pagination.reverse"?: boolean;
    }, params?: RequestParams) => Promise<HttpResponse<JklminingQueryAllMinersResponse, RpcStatus>>;
    /**
     * No description
     *
     * @tags Query
     * @name QueryMiners
     * @summary Queries a Miners by index.
     * @request GET:/jackal-dao/canine/jklmining/miners/{address}
     */
    queryMiners: (address: string, params?: RequestParams) => Promise<HttpResponse<JklminingQueryGetMinersResponse, RpcStatus>>;
    /**
     * No description
     *
     * @tags Query
     * @name QuerySaveRequestsAll
     * @summary Queries a list of SaveRequests items.
     * @request GET:/jackal-dao/canine/jklmining/save_requests
     */
    querySaveRequestsAll: (query?: {
        "pagination.key"?: string;
        "pagination.offset"?: string;
        "pagination.limit"?: string;
        "pagination.count_total"?: boolean;
        "pagination.reverse"?: boolean;
    }, params?: RequestParams) => Promise<HttpResponse<JklminingQueryAllSaveRequestsResponse, RpcStatus>>;
    /**
     * No description
     *
     * @tags Query
     * @name QuerySaveRequests
     * @summary Queries a SaveRequests by index.
     * @request GET:/jackal-dao/canine/jklmining/save_requests/{index}
     */
    querySaveRequests: (index: string, params?: RequestParams) => Promise<HttpResponse<JklminingQueryGetSaveRequestsResponse, RpcStatus>>;
    /**
     * No description
     *
     * @tags Query
     * @name QueryParams
     * @summary Parameters queries the parameters of the module.
     * @request GET:/jackaldao/canine/jklmining/params
     */
    queryParams: (params?: RequestParams) => Promise<HttpResponse<JklminingQueryParamsResponse, RpcStatus>>;
}
export {};
