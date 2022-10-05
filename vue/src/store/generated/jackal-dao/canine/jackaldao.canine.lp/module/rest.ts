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

export interface LpLPool {
  index?: string;
  name?: string;
  coins?: V1Beta1Coin[];

  /** @format int64 */
  aMM_Id?: number;
  swap_fee_multi?: string;

  /** @format int64 */
  min_lock_duration?: string;
  penalty_multi?: string;
  lptoken_denom?: string;
  LPTokenBalance?: string;
}

/**
* LProviderRecord is a record of a liquidity provider depositing to a pool.
It is used to enforce withdraw panelty and calculate rewards collected. 
This record is created only once when provider contributes to a pool and
only updated after witdrawal or deposit.
It is deleted when the provider burns all of the liquidity pool token.
This is stored at KVStore with 
	{LProviderRecordKeyPrefix}{poolName}{provider} key.
*/
export interface LpLProviderRecord {
  /** Provider is the account address of the provider. */
  provider?: string;

  /** A pool that the provider contributed to. */
  poolName?: string;

  /**
   * Burning LP token is locked for certain duration the after provider
   * deposits to the pool. Unlock time is updated every succeeding deposits.
   * The provider can burn their LP token during lock time but has to take
   * certain amount of panelty.
   * Unlock time is blocktime + lockDuration at time of contribution.
   */
  unlockTime?: string;
  lockDuration?: string;
}

export interface LpMsgCreateLPoolResponse {
  id?: string;
}

export interface LpMsgDepositLPoolResponse {
  /**
   * Amount of shares given to pool contributor.
   * @format uint64
   */
  shares?: string;
}

export type LpMsgSwapResponse = object;

export type LpMsgWithdrawLPoolResponse = object;

/**
 * Params defines the parameters for the module.
 */
export interface LpParams {
  /** @format uint64 */
  MinInitPoolDeposit?: string;

  /** @format int64 */
  MaxPoolDenomCount?: number;

  /** @format int64 */
  LPTokenUnit?: number;
}

export interface LpQueryAllLPoolResponse {
  lPool?: LpLPool[];

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

export interface LpQueryEstimateContributionResponse {
  coins?: V1Beta1Coin[];
}

export interface LpQueryEstimatePoolRemoveResponse {
  coins?: V1Beta1Coin[];
}

export interface LpQueryEstimateSwapInResponse {
  /**
   * Coin defines a token with a denomination and an amount.
   *
   * NOTE: The amount field is an Int which implements the custom method
   * signatures required by gogoproto.
   */
  inputCoins?: V1Beta1Coin;
}

export interface LpQueryEstimateSwapOutResponse {
  /**
   * Coin defines a token with a denomination and an amount.
   *
   * NOTE: The amount field is an Int which implements the custom method
   * signatures required by gogoproto.
   */
  outputCoin?: V1Beta1Coin;
}

export interface LpQueryGetLPoolResponse {
  lPool?: LpLPool;
}

export interface LpQueryGetLProviderRecordResponse {
  /**
   * LProviderRecord is a record of a liquidity provider depositing to a pool.
   * It is used to enforce withdraw panelty and calculate rewards collected.
   * This record is created only once when provider contributes to a pool and
   * only updated after witdrawal or deposit.
   * It is deleted when the provider burns all of the liquidity pool token.
   * This is stored at KVStore with
   * 	{LProviderRecordKeyPrefix}{poolName}{provider} key.
   */
  lProviderRecord?: LpLProviderRecord;
}

export interface LpQueryListRecordsFromPoolResponse {
  records?: LpLProviderRecord[];
}

export interface LpQueryMakeValidPairResponse {
  /**
   * Coin defines a token with a denomination and an amount.
   *
   * NOTE: The amount field is an Int which implements the custom method
   * signatures required by gogoproto.
   */
  coin?: V1Beta1Coin;
}

/**
 * QueryParamsResponse is response type for the Query/Params RPC method.
 */
export interface LpQueryParamsResponse {
  /** params holds all the parameters of this module. */
  params?: LpParams;
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
* Coin defines a token with a denomination and an amount.

NOTE: The amount field is an Int which implements the custom method
signatures required by gogoproto.
*/
export interface V1Beta1Coin {
  denom?: string;
  amount?: string;
}

/**
* DecCoin defines a token with a denomination and a decimal amount.

NOTE: The amount field is an Dec which implements the custom method
signatures required by gogoproto.
*/
export interface V1Beta1DecCoin {
  denom?: string;
  amount?: string;
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
 * @title lp/genesis.proto
 * @version version not set
 */
export class Api<SecurityDataType extends unknown> extends HttpClient<SecurityDataType> {
  /**
   * No description
   *
   * @tags Query
   * @name QueryEstimateContribution
   * @summary Estimate coin inputs to get desired amount of LPToken.
   * @request GET:/jackal-dao/canine/lp/estimate_contribution/{poolName}/{desiredAmount}
   */
  queryEstimateContribution = (poolName: string, desiredAmount: string, params: RequestParams = {}) =>
    this.request<LpQueryEstimateContributionResponse, RpcStatus>({
      path: `/jackal-dao/canine/lp/estimate_contribution/${poolName}/${desiredAmount}`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryEstimateSwapIn
   * @summary Estimate coin input to get desired coin output from a swap.
   * @request GET:/jackal-dao/canine/lp/estimate_in/{poolName}/{outputCoins}
   */
  queryEstimateSwapIn = (poolName: string, outputCoins: string, params: RequestParams = {}) =>
    this.request<LpQueryEstimateSwapInResponse, RpcStatus>({
      path: `/jackal-dao/canine/lp/estimate_in/${poolName}/${outputCoins}`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryEstimateSwapOut
   * @summary Estimate coin output from a swap.
   * @request GET:/jackal-dao/canine/lp/estimate_out/{poolName}/{inputCoin}
   */
  queryEstimateSwapOut = (poolName: string, inputCoin: string, params: RequestParams = {}) =>
    this.request<LpQueryEstimateSwapOutResponse, RpcStatus>({
      path: `/jackal-dao/canine/lp/estimate_out/${poolName}/${inputCoin}`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryEstimatePoolRemove
   * @summary Estimate amoutn of coins returned by burning a LPToken.
   * @request GET:/jackal-dao/canine/lp/estimate_pool_remove/{amount}
   */
  queryEstimatePoolRemove = (amount: string, query?: { poolName?: string }, params: RequestParams = {}) =>
    this.request<LpQueryEstimatePoolRemoveResponse, RpcStatus>({
      path: `/jackal-dao/canine/lp/estimate_pool_remove/${amount}`,
      method: "GET",
      query: query,
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryLPoolAll
   * @summary Queries a list of LPool items.
   * @request GET:/jackal-dao/canine/lp/l_pool
   */
  queryLPoolAll = (
    query?: {
      "pagination.key"?: string;
      "pagination.offset"?: string;
      "pagination.limit"?: string;
      "pagination.count_total"?: boolean;
      "pagination.reverse"?: boolean;
    },
    params: RequestParams = {},
  ) =>
    this.request<LpQueryAllLPoolResponse, RpcStatus>({
      path: `/jackal-dao/canine/lp/l_pool`,
      method: "GET",
      query: query,
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryLPool
   * @summary Queries a LPool by index.
   * @request GET:/jackal-dao/canine/lp/l_pool/{index}
   */
  queryLPool = (index: string, params: RequestParams = {}) =>
    this.request<LpQueryGetLPoolResponse, RpcStatus>({
      path: `/jackal-dao/canine/lp/l_pool/${index}`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryLProviderRecord
   * @summary Queries a LProviderRecord by pool name and provider address.
   * @request GET:/jackal-dao/canine/lp/l_provider_record/{poolName}/{provider}
   */
  queryLProviderRecord = (poolName: string, provider: string, params: RequestParams = {}) =>
    this.request<LpQueryGetLProviderRecordResponse, RpcStatus>({
      path: `/jackal-dao/canine/lp/l_provider_record/${poolName}/${provider}`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryListRecordsFromPool
   * @summary Queries a list of ListRecordsFromPool items.
   * @request GET:/jackal-dao/canine/lp/list_records_from_pool/{poolName}
   */
  queryListRecordsFromPool = (poolName: string, params: RequestParams = {}) =>
    this.request<LpQueryListRecordsFromPoolResponse, RpcStatus>({
      path: `/jackal-dao/canine/lp/list_records_from_pool/${poolName}`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
 * No description
 * 
 * @tags Query
 * @name QueryMakeValidPair
 * @summary Query coins to deposit to make valid liquidity pair.
Input one coin and get other coins to deposit to make liquidity pair.
 * @request GET:/jackal-dao/canine/lp/make_pair/{poolName}/{coin}
 */
  queryMakeValidPair = (poolName: string, coin: string, params: RequestParams = {}) =>
    this.request<LpQueryMakeValidPairResponse, RpcStatus>({
      path: `/jackal-dao/canine/lp/make_pair/${poolName}/${coin}`,
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
   * @request GET:/jackal-dao/canine/lp/params
   */
  queryParams = (params: RequestParams = {}) =>
    this.request<LpQueryParamsResponse, RpcStatus>({
      path: `/jackal-dao/canine/lp/params`,
      method: "GET",
      format: "json",
      ...params,
    });
}
