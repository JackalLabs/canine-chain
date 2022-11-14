/* eslint-disable */
/* tslint:disable */
/*
 * ---------------------------------------------------------------
 * ## THIS FILE WAS GENERATED VIA SWAGGER-TYPESCRIPT-API        ##
 * ##                                                           ##
 * ## AUTHOR: acacode                                           ##
 * ## SOURCE: https://github.com/acacode/swagger-typescript-api ##
 * ---------------------------------------------------------------
 */

export interface ProtobufAny {
  "@type"?: string;
}

export interface RpcStatus {
  /** @format int32 */
  code?: number;
  message?: string;
  details?: ProtobufAny[];
}

export interface StorageActiveDeals {
  cid?: string;
  signee?: string;
  miner?: string;
  startblock?: string;
  endblock?: string;
  filesize?: string;
  proofverified?: string;
  proofsmissed?: string;
  blocktoprove?: string;
  creator?: string;
  merkle?: string;
  fid?: string;
}

export interface StorageClientUsage {
  address?: string;
  usage?: string;
}

export interface StorageContracts {
  cid?: string;
  priceamt?: string;
  pricedenom?: string;
  merkle?: string;
  signee?: string;
  duration?: string;
  filesize?: string;
  fid?: string;
  creator?: string;
}

export interface StorageMiners {
  address?: string;
  ip?: string;
  totalspace?: string;
  burned_contracts?: string;
  creator?: string;
}

export type StorageMsgBuyStorageResponse = object;

export type StorageMsgCancelContractResponse = object;

export type StorageMsgInitMinerResponse = object;

export type StorageMsgPostContractResponse = object;

export interface StorageMsgPostproofResponse {
  merkle?: string;
}

export type StorageMsgSetMinerIpResponse = object;

export type StorageMsgSetMinerTotalspaceResponse = object;

export type StorageMsgSignContractResponse = object;

/**
 * Params defines the parameters for the module.
 */
export type StorageParams = object;

export interface StoragePayBlocks {
  blockid?: string;
  bytes?: string;
  blocktype?: string;
  blocknum?: string;
}

export interface StorageProofs {
  cid?: string;
  item?: string;
  hashes?: string;
  creator?: string;
}

export interface StorageQueryAllActiveDealsResponse {
  activeDeals?: StorageActiveDeals[];

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

export interface StorageQueryAllClientUsageResponse {
  clientUsage?: StorageClientUsage[];

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

export interface StorageQueryAllContractsResponse {
  contracts?: StorageContracts[];

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

export interface StorageQueryAllMinersResponse {
  miners?: StorageMiners[];

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

export interface StorageQueryAllPayBlocksResponse {
  payBlocks?: StoragePayBlocks[];

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

export interface StorageQueryAllProofsResponse {
  proofs?: StorageProofs[];

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

export interface StorageQueryAllStraysResponse {
  strays?: StorageStrays[];

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

export interface StorageQueryFindFileResponse {
  minerIps?: string;
}

export interface StorageQueryFreespaceResponse {
  space?: string;
}

export interface StorageQueryGetActiveDealsResponse {
  activeDeals?: StorageActiveDeals;
}

export interface StorageQueryGetClientFreeSpaceResponse {
  bytesfree?: string;
}

export interface StorageQueryGetClientUsageResponse {
  clientUsage?: StorageClientUsage;
}

export interface StorageQueryGetContractsResponse {
  contracts?: StorageContracts;
}

export interface StorageQueryGetMinersResponse {
  miners?: StorageMiners;
}

export interface StorageQueryGetPayBlocksResponse {
  payBlocks?: StoragePayBlocks;
}

export interface StorageQueryGetProofsResponse {
  proofs?: StorageProofs;
}

export interface StorageQueryGetStraysResponse {
  strays?: StorageStrays;
}

/**
 * QueryParamsResponse is response type for the Query/Params RPC method.
 */
export interface StorageQueryParamsResponse {
  /** params holds all the parameters of this module. */
  params?: StorageParams;
}

export interface StorageStrays {
  cid?: string;
  fid?: string;
  signee?: string;
  filesize?: string;
  merkle?: string;
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

export type QueryParamsType = Record<string | number, any>;
export type ResponseFormat = keyof Omit<Body, "body" | "bodyUsed">;

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

export type RequestParams = Omit<FullRequestParams, "body" | "method" | "query" | "path">;

export interface ApiConfig<SecurityDataType = unknown> {
  baseUrl?: string;
  baseApiParams?: Omit<RequestParams, "baseUrl" | "cancelToken" | "signal">;
  securityWorker?: (securityData: SecurityDataType) => RequestParams | void;
}

export interface HttpResponse<D extends unknown, E extends unknown = unknown> extends Response {
  data: D;
  error: E;
}

type CancelToken = Symbol | string | number;

export enum ContentType {
  Json = "application/json",
  FormData = "multipart/form-data",
  UrlEncoded = "application/x-www-form-urlencoded",
}

export class HttpClient<SecurityDataType = unknown> {
  public baseUrl: string = "";
  private securityData: SecurityDataType = null as any;
  private securityWorker: null | ApiConfig<SecurityDataType>["securityWorker"] = null;
  private abortControllers = new Map<CancelToken, AbortController>();

  private baseApiParams: RequestParams = {
    credentials: "same-origin",
    headers: {},
    redirect: "follow",
    referrerPolicy: "no-referrer",
  };

  constructor(apiConfig: ApiConfig<SecurityDataType> = {}) {
    Object.assign(this, apiConfig);
  }

  public setSecurityData = (data: SecurityDataType) => {
    this.securityData = data;
  };

  private addQueryParam(query: QueryParamsType, key: string) {
    const value = query[key];

    return (
      encodeURIComponent(key) +
      "=" +
      encodeURIComponent(Array.isArray(value) ? value.join(",") : typeof value === "number" ? value : `${value}`)
    );
  }

  protected toQueryString(rawQuery?: QueryParamsType): string {
    const query = rawQuery || {};
    const keys = Object.keys(query).filter((key) => "undefined" !== typeof query[key]);
    return keys
      .map((key) =>
        typeof query[key] === "object" && !Array.isArray(query[key])
          ? this.toQueryString(query[key] as QueryParamsType)
          : this.addQueryParam(query, key),
      )
      .join("&");
  }

  protected addQueryParams(rawQuery?: QueryParamsType): string {
    const queryString = this.toQueryString(rawQuery);
    return queryString ? `?${queryString}` : "";
  }

  private contentFormatters: Record<ContentType, (input: any) => any> = {
    [ContentType.Json]: (input: any) =>
      input !== null && (typeof input === "object" || typeof input === "string") ? JSON.stringify(input) : input,
    [ContentType.FormData]: (input: any) =>
      Object.keys(input || {}).reduce((data, key) => {
        data.append(key, input[key]);
        return data;
      }, new FormData()),
    [ContentType.UrlEncoded]: (input: any) => this.toQueryString(input),
  };

  private mergeRequestParams(params1: RequestParams, params2?: RequestParams): RequestParams {
    return {
      ...this.baseApiParams,
      ...params1,
      ...(params2 || {}),
      headers: {
        ...(this.baseApiParams.headers || {}),
        ...(params1.headers || {}),
        ...((params2 && params2.headers) || {}),
      },
    };
  }

  private createAbortSignal = (cancelToken: CancelToken): AbortSignal | undefined => {
    if (this.abortControllers.has(cancelToken)) {
      const abortController = this.abortControllers.get(cancelToken);
      if (abortController) {
        return abortController.signal;
      }
      return void 0;
    }

    const abortController = new AbortController();
    this.abortControllers.set(cancelToken, abortController);
    return abortController.signal;
  };

  public abortRequest = (cancelToken: CancelToken) => {
    const abortController = this.abortControllers.get(cancelToken);

    if (abortController) {
      abortController.abort();
      this.abortControllers.delete(cancelToken);
    }
  };

  public request = <T = any, E = any>({
    body,
    secure,
    path,
    type,
    query,
    format = "json",
    baseUrl,
    cancelToken,
    ...params
  }: FullRequestParams): Promise<HttpResponse<T, E>> => {
    const secureParams = (secure && this.securityWorker && this.securityWorker(this.securityData)) || {};
    const requestParams = this.mergeRequestParams(params, secureParams);
    const queryString = query && this.toQueryString(query);
    const payloadFormatter = this.contentFormatters[type || ContentType.Json];

    return fetch(`${baseUrl || this.baseUrl || ""}${path}${queryString ? `?${queryString}` : ""}`, {
      ...requestParams,
      headers: {
        ...(type && type !== ContentType.FormData ? { "Content-Type": type } : {}),
        ...(requestParams.headers || {}),
      },
      signal: cancelToken ? this.createAbortSignal(cancelToken) : void 0,
      body: typeof body === "undefined" || body === null ? null : payloadFormatter(body),
    }).then(async (response) => {
      const r = response as HttpResponse<T, E>;
      r.data = (null as unknown) as T;
      r.error = (null as unknown) as E;

      const data = await response[format]()
        .then((data) => {
          if (r.ok) {
            r.data = data;
          } else {
            r.error = data;
          }
          return r;
        })
        .catch((e) => {
          r.error = e;
          return r;
        });

      if (cancelToken) {
        this.abortControllers.delete(cancelToken);
      }

      if (!response.ok) throw data;
      return data;
    });
  };
}

/**
 * @title storage/active_deals.proto
 * @version version not set
 */
export class Api<SecurityDataType extends unknown> extends HttpClient<SecurityDataType> {
  /**
   * No description
   *
   * @tags Query
   * @name QueryActiveDealsAll
   * @summary Queries a list of ActiveDeals items.
   * @request GET:/jackal-dao/canine/storage/active_deals
   */
  queryActiveDealsAll = (
    query?: {
      "pagination.key"?: string;
      "pagination.offset"?: string;
      "pagination.limit"?: string;
      "pagination.count_total"?: boolean;
      "pagination.reverse"?: boolean;
    },
    params: RequestParams = {},
  ) =>
    this.request<StorageQueryAllActiveDealsResponse, RpcStatus>({
      path: `/jackal-dao/canine/storage/active_deals`,
      method: "GET",
      query: query,
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryActiveDeals
   * @summary Queries a ActiveDeals by index.
   * @request GET:/jackal-dao/canine/storage/active_deals/{cid}
   */
  queryActiveDeals = (cid: string, params: RequestParams = {}) =>
    this.request<StorageQueryGetActiveDealsResponse, RpcStatus>({
      path: `/jackal-dao/canine/storage/active_deals/${cid}`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryClientUsageAll
   * @summary Queries a list of ClientUsage items.
   * @request GET:/jackal-dao/canine/storage/client_usage
   */
  queryClientUsageAll = (
    query?: {
      "pagination.key"?: string;
      "pagination.offset"?: string;
      "pagination.limit"?: string;
      "pagination.count_total"?: boolean;
      "pagination.reverse"?: boolean;
    },
    params: RequestParams = {},
  ) =>
    this.request<StorageQueryAllClientUsageResponse, RpcStatus>({
      path: `/jackal-dao/canine/storage/client_usage`,
      method: "GET",
      query: query,
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryClientUsage
   * @summary Queries a ClientUsage by index.
   * @request GET:/jackal-dao/canine/storage/client_usage/{address}
   */
  queryClientUsage = (address: string, params: RequestParams = {}) =>
    this.request<StorageQueryGetClientUsageResponse, RpcStatus>({
      path: `/jackal-dao/canine/storage/client_usage/${address}`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryContractsAll
   * @summary Queries a list of Contracts items.
   * @request GET:/jackal-dao/canine/storage/contracts
   */
  queryContractsAll = (
    query?: {
      "pagination.key"?: string;
      "pagination.offset"?: string;
      "pagination.limit"?: string;
      "pagination.count_total"?: boolean;
      "pagination.reverse"?: boolean;
    },
    params: RequestParams = {},
  ) =>
    this.request<StorageQueryAllContractsResponse, RpcStatus>({
      path: `/jackal-dao/canine/storage/contracts`,
      method: "GET",
      query: query,
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryContracts
   * @summary Queries a Contracts by index.
   * @request GET:/jackal-dao/canine/storage/contracts/{cid}
   */
  queryContracts = (cid: string, params: RequestParams = {}) =>
    this.request<StorageQueryGetContractsResponse, RpcStatus>({
      path: `/jackal-dao/canine/storage/contracts/${cid}`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryFindFile
   * @summary Queries a list of FindFile items.
   * @request GET:/jackal-dao/canine/storage/find_file/{fid}
   */
  queryFindFile = (fid: string, params: RequestParams = {}) =>
    this.request<StorageQueryFindFileResponse, RpcStatus>({
      path: `/jackal-dao/canine/storage/find_file/${fid}`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryFreespace
   * @summary Queries a list of Freespace items.
   * @request GET:/jackal-dao/canine/storage/freespace/{address}
   */
  queryFreespace = (address: string, params: RequestParams = {}) =>
    this.request<StorageQueryFreespaceResponse, RpcStatus>({
      path: `/jackal-dao/canine/storage/freespace/${address}`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryGetClientFreeSpace
   * @summary Queries a list of GetClientFreeSpace items.
   * @request GET:/jackal-dao/canine/storage/get_client_free_space/{address}
   */
  queryGetClientFreeSpace = (address: string, params: RequestParams = {}) =>
    this.request<StorageQueryGetClientFreeSpaceResponse, RpcStatus>({
      path: `/jackal-dao/canine/storage/get_client_free_space/${address}`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryMinersAll
   * @summary Queries a list of Miners items.
   * @request GET:/jackal-dao/canine/storage/miners
   */
  queryMinersAll = (
    query?: {
      "pagination.key"?: string;
      "pagination.offset"?: string;
      "pagination.limit"?: string;
      "pagination.count_total"?: boolean;
      "pagination.reverse"?: boolean;
    },
    params: RequestParams = {},
  ) =>
    this.request<StorageQueryAllMinersResponse, RpcStatus>({
      path: `/jackal-dao/canine/storage/miners`,
      method: "GET",
      query: query,
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryMiners
   * @summary Queries a Miners by index.
   * @request GET:/jackal-dao/canine/storage/miners/{address}
   */
  queryMiners = (address: string, params: RequestParams = {}) =>
    this.request<StorageQueryGetMinersResponse, RpcStatus>({
      path: `/jackal-dao/canine/storage/miners/${address}`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryParams
   * @summary Parameters queries the parameters of the module.
   * @request GET:/jackal-dao/canine/storage/params
   */
  queryParams = (params: RequestParams = {}) =>
    this.request<StorageQueryParamsResponse, RpcStatus>({
      path: `/jackal-dao/canine/storage/params`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryPayBlocksAll
   * @summary Queries a list of PayBlocks items.
   * @request GET:/jackal-dao/canine/storage/pay_blocks
   */
  queryPayBlocksAll = (
    query?: {
      "pagination.key"?: string;
      "pagination.offset"?: string;
      "pagination.limit"?: string;
      "pagination.count_total"?: boolean;
      "pagination.reverse"?: boolean;
    },
    params: RequestParams = {},
  ) =>
    this.request<StorageQueryAllPayBlocksResponse, RpcStatus>({
      path: `/jackal-dao/canine/storage/pay_blocks`,
      method: "GET",
      query: query,
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryPayBlocks
   * @summary Queries a PayBlocks by index.
   * @request GET:/jackal-dao/canine/storage/pay_blocks/{blockid}
   */
  queryPayBlocks = (blockid: string, params: RequestParams = {}) =>
    this.request<StorageQueryGetPayBlocksResponse, RpcStatus>({
      path: `/jackal-dao/canine/storage/pay_blocks/${blockid}`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryProofsAll
   * @summary Queries a list of Proofs items.
   * @request GET:/jackal-dao/canine/storage/proofs
   */
  queryProofsAll = (
    query?: {
      "pagination.key"?: string;
      "pagination.offset"?: string;
      "pagination.limit"?: string;
      "pagination.count_total"?: boolean;
      "pagination.reverse"?: boolean;
    },
    params: RequestParams = {},
  ) =>
    this.request<StorageQueryAllProofsResponse, RpcStatus>({
      path: `/jackal-dao/canine/storage/proofs`,
      method: "GET",
      query: query,
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryProofs
   * @summary Queries a Proofs by index.
   * @request GET:/jackal-dao/canine/storage/proofs/{cid}
   */
  queryProofs = (cid: string, params: RequestParams = {}) =>
    this.request<StorageQueryGetProofsResponse, RpcStatus>({
      path: `/jackal-dao/canine/storage/proofs/${cid}`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryStraysAll
   * @summary Queries a list of Strays items.
   * @request GET:/jackal-dao/canine/storage/strays
   */
  queryStraysAll = (
    query?: {
      "pagination.key"?: string;
      "pagination.offset"?: string;
      "pagination.limit"?: string;
      "pagination.count_total"?: boolean;
      "pagination.reverse"?: boolean;
    },
    params: RequestParams = {},
  ) =>
    this.request<StorageQueryAllStraysResponse, RpcStatus>({
      path: `/jackal-dao/canine/storage/strays`,
      method: "GET",
      query: query,
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryStrays
   * @summary Queries a Strays by index.
   * @request GET:/jackal-dao/canine/storage/strays/{cid}
   */
  queryStrays = (cid: string, params: RequestParams = {}) =>
    this.request<StorageQueryGetStraysResponse, RpcStatus>({
      path: `/jackal-dao/canine/storage/strays/${cid}`,
      method: "GET",
      format: "json",
      ...params,
    });
}
