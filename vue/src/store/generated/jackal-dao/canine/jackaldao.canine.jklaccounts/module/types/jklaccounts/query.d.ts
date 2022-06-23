import { Reader, Writer } from "protobufjs/minimal";
import { Params } from "../jklaccounts/params";
import { Accounts } from "../jklaccounts/accounts";
import { PageRequest, PageResponse } from "../cosmos/base/query/v1beta1/pagination";
export declare const protobufPackage = "jackaldao.canine.jklaccounts";
/** QueryParamsRequest is request type for the Query/Params RPC method. */
export interface QueryParamsRequest {
}
/** QueryParamsResponse is response type for the Query/Params RPC method. */
export interface QueryParamsResponse {
    /** params holds all the parameters of this module. */
    params: Params | undefined;
}
export interface QueryGetAccountsRequest {
    address: string;
}
export interface QueryGetAccountsResponse {
    accounts: Accounts | undefined;
}
export interface QueryAllAccountsRequest {
    pagination: PageRequest | undefined;
}
export interface QueryAllAccountsResponse {
    accounts: Accounts[];
    pagination: PageResponse | undefined;
}
export declare const QueryParamsRequest: {
    encode(_: QueryParamsRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryParamsRequest;
    fromJSON(_: any): QueryParamsRequest;
    toJSON(_: QueryParamsRequest): unknown;
    fromPartial(_: DeepPartial<QueryParamsRequest>): QueryParamsRequest;
};
export declare const QueryParamsResponse: {
    encode(message: QueryParamsResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryParamsResponse;
    fromJSON(object: any): QueryParamsResponse;
    toJSON(message: QueryParamsResponse): unknown;
    fromPartial(object: DeepPartial<QueryParamsResponse>): QueryParamsResponse;
};
export declare const QueryGetAccountsRequest: {
    encode(message: QueryGetAccountsRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetAccountsRequest;
    fromJSON(object: any): QueryGetAccountsRequest;
    toJSON(message: QueryGetAccountsRequest): unknown;
    fromPartial(object: DeepPartial<QueryGetAccountsRequest>): QueryGetAccountsRequest;
};
export declare const QueryGetAccountsResponse: {
    encode(message: QueryGetAccountsResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetAccountsResponse;
    fromJSON(object: any): QueryGetAccountsResponse;
    toJSON(message: QueryGetAccountsResponse): unknown;
    fromPartial(object: DeepPartial<QueryGetAccountsResponse>): QueryGetAccountsResponse;
};
export declare const QueryAllAccountsRequest: {
    encode(message: QueryAllAccountsRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryAllAccountsRequest;
    fromJSON(object: any): QueryAllAccountsRequest;
    toJSON(message: QueryAllAccountsRequest): unknown;
    fromPartial(object: DeepPartial<QueryAllAccountsRequest>): QueryAllAccountsRequest;
};
export declare const QueryAllAccountsResponse: {
    encode(message: QueryAllAccountsResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryAllAccountsResponse;
    fromJSON(object: any): QueryAllAccountsResponse;
    toJSON(message: QueryAllAccountsResponse): unknown;
    fromPartial(object: DeepPartial<QueryAllAccountsResponse>): QueryAllAccountsResponse;
};
/** Query defines the gRPC querier service. */
export interface Query {
    /** Parameters queries the parameters of the module. */
    Params(request: QueryParamsRequest): Promise<QueryParamsResponse>;
    /** Queries a Accounts by index. */
    Accounts(request: QueryGetAccountsRequest): Promise<QueryGetAccountsResponse>;
    /** Queries a list of Accounts items. */
    AccountsAll(request: QueryAllAccountsRequest): Promise<QueryAllAccountsResponse>;
}
export declare class QueryClientImpl implements Query {
    private readonly rpc;
    constructor(rpc: Rpc);
    Params(request: QueryParamsRequest): Promise<QueryParamsResponse>;
    Accounts(request: QueryGetAccountsRequest): Promise<QueryGetAccountsResponse>;
    AccountsAll(request: QueryAllAccountsRequest): Promise<QueryAllAccountsResponse>;
}
interface Rpc {
    request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
