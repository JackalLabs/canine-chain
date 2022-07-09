export interface ProtobufAny {
    "@type"?: string;
}
export interface RnsBids {
    index?: string;
    name?: string;
    bidder?: string;
    price?: string;
}
export interface RnsForsale {
    name?: string;
    price?: string;
    owner?: string;
}
export declare type RnsMsgAcceptBidResponse = object;
export declare type RnsMsgBidResponse = object;
export declare type RnsMsgBuyResponse = object;
export declare type RnsMsgCancelBidResponse = object;
export declare type RnsMsgDelistResponse = object;
export declare type RnsMsgListResponse = object;
export declare type RnsMsgRegisterResponse = object;
export declare type RnsMsgTransferResponse = object;
export interface RnsNames {
    index?: string;
    name?: string;
    expires?: string;
    value?: string;
    data?: string;
}
/**
 * Params defines the parameters for the module.
 */
export declare type RnsParams = object;
export interface RnsQueryAllBidsResponse {
    bids?: RnsBids[];
    pagination?: V1Beta1PageResponse;
}
export interface RnsQueryAllForsaleResponse {
    forsale?: RnsForsale[];
    pagination?: V1Beta1PageResponse;
}
export interface RnsQueryAllNamesResponse {
    names?: RnsNames[];
    pagination?: V1Beta1PageResponse;
}
export interface RnsQueryGetBidsResponse {
    bids?: RnsBids;
}
export interface RnsQueryGetForsaleResponse {
    forsale?: RnsForsale;
}
export interface RnsQueryGetNamesResponse {
    names?: RnsNames;
}
/**
 * QueryParamsResponse is response type for the Query/Params RPC method.
 */
export interface RnsQueryParamsResponse {
    /** params holds all the parameters of this module. */
    params?: RnsParams;
}
export interface RpcStatus {
    /** @format int32 */
    code?: number;
    message?: string;
    details?: ProtobufAny[];
}
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
    reverse?: boolean;
}
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
 * @title rns/bids.proto
 * @version version not set
 */
export declare class Api<SecurityDataType extends unknown> extends HttpClient<SecurityDataType> {
    /**
     * No description
     *
     * @tags Query
     * @name QueryBidsAll
     * @summary Queries a list of Bids.
     * @request GET:/jackal-dao/canine/rnsbids
     */
    queryBidsAll: (query?: {
        "pagination.key"?: string;
        "pagination.offset"?: string;
        "pagination.limit"?: string;
        "pagination.count_total"?: boolean;
        "pagination.reverse"?: boolean;
    }, params?: RequestParams) => Promise<HttpResponse<RnsQueryAllBidsResponse, RpcStatus>>;
    /**
     * No description
     *
     * @tags Query
     * @name QueryBids
     * @summary Queries a Bid by index.
     * @request GET:/jackal-dao/canine/rnsbids/{index}
     */
    queryBids: (index: string, params?: RequestParams) => Promise<HttpResponse<RnsQueryGetBidsResponse, RpcStatus>>;
    /**
     * No description
     *
     * @tags Query
     * @name QueryForsaleAll
     * @summary Queries all Listings.
     * @request GET:/jackal-dao/canine/rnsforsale
     */
    queryForsaleAll: (query?: {
        "pagination.key"?: string;
        "pagination.offset"?: string;
        "pagination.limit"?: string;
        "pagination.count_total"?: boolean;
        "pagination.reverse"?: boolean;
    }, params?: RequestParams) => Promise<HttpResponse<RnsQueryAllForsaleResponse, RpcStatus>>;
    /**
     * No description
     *
     * @tags Query
     * @name QueryForsale
     * @summary Queries a Listing by index.
     * @request GET:/jackal-dao/canine/rnsforsale/{name}
     */
    queryForsale: (name: string, params?: RequestParams) => Promise<HttpResponse<RnsQueryGetForsaleResponse, RpcStatus>>;
    /**
     * No description
     *
     * @tags Query
     * @name QueryNamesAll
     * @summary Queries a list of Names.
     * @request GET:/jackal-dao/canine/rnsnames
     */
    queryNamesAll: (query?: {
        "pagination.key"?: string;
        "pagination.offset"?: string;
        "pagination.limit"?: string;
        "pagination.count_total"?: boolean;
        "pagination.reverse"?: boolean;
    }, params?: RequestParams) => Promise<HttpResponse<RnsQueryAllNamesResponse, RpcStatus>>;
    /**
     * No description
     *
     * @tags Query
     * @name QueryNames
     * @summary Queries a Name by index.
     * @request GET:/jackal-dao/canine/rnsnames/{index}
     */
    queryNames: (index: string, params?: RequestParams) => Promise<HttpResponse<RnsQueryGetNamesResponse, RpcStatus>>;
    /**
     * No description
     *
     * @tags Query
     * @name QueryParams
     * @summary Parameters queries the parameters of the module.
     * @request GET:/jackal-dao/canine/rnsparams
     */
    queryParams: (params?: RequestParams) => Promise<HttpResponse<RnsQueryParamsResponse, RpcStatus>>;
}
export {};
