import { Reader, Writer } from "protobufjs/minimal";
import { Params } from "../rns/params";
import { Whois } from "../rns/whois";
import { PageRequest, PageResponse } from "../cosmos/base/query/v1beta1/pagination";
import { Names } from "../rns/names";
import { Bids } from "../rns/bids";
import { Forsale } from "../rns/forsale";
export declare const protobufPackage = "jackaldao.canine.rns";
/** QueryParamsRequest is request type for the Query/Params RPC method. */
export interface QueryParamsRequest {
}
/** QueryParamsResponse is response type for the Query/Params RPC method. */
export interface QueryParamsResponse {
    /** params holds all the parameters of this module. */
    params: Params | undefined;
}
export interface QueryGetWhoisRequest {
    index: string;
}
export interface QueryGetWhoisResponse {
    whois: Whois | undefined;
}
export interface QueryAllWhoisRequest {
    pagination: PageRequest | undefined;
}
export interface QueryAllWhoisResponse {
    whois: Whois[];
    pagination: PageResponse | undefined;
}
export interface QueryGetNamesRequest {
    index: string;
}
export interface QueryGetNamesResponse {
    names: Names | undefined;
}
export interface QueryAllNamesRequest {
    pagination: PageRequest | undefined;
}
export interface QueryAllNamesResponse {
    names: Names[];
    pagination: PageResponse | undefined;
}
export interface QueryGetBidsRequest {
    index: string;
}
export interface QueryGetBidsResponse {
    bids: Bids | undefined;
}
export interface QueryAllBidsRequest {
    pagination: PageRequest | undefined;
}
export interface QueryAllBidsResponse {
    bids: Bids[];
    pagination: PageResponse | undefined;
}
export interface QueryGetForsaleRequest {
    name: string;
}
export interface QueryGetForsaleResponse {
    forsale: Forsale | undefined;
}
export interface QueryAllForsaleRequest {
    pagination: PageRequest | undefined;
}
export interface QueryAllForsaleResponse {
    forsale: Forsale[];
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
export declare const QueryGetWhoisRequest: {
    encode(message: QueryGetWhoisRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetWhoisRequest;
    fromJSON(object: any): QueryGetWhoisRequest;
    toJSON(message: QueryGetWhoisRequest): unknown;
    fromPartial(object: DeepPartial<QueryGetWhoisRequest>): QueryGetWhoisRequest;
};
export declare const QueryGetWhoisResponse: {
    encode(message: QueryGetWhoisResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetWhoisResponse;
    fromJSON(object: any): QueryGetWhoisResponse;
    toJSON(message: QueryGetWhoisResponse): unknown;
    fromPartial(object: DeepPartial<QueryGetWhoisResponse>): QueryGetWhoisResponse;
};
export declare const QueryAllWhoisRequest: {
    encode(message: QueryAllWhoisRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryAllWhoisRequest;
    fromJSON(object: any): QueryAllWhoisRequest;
    toJSON(message: QueryAllWhoisRequest): unknown;
    fromPartial(object: DeepPartial<QueryAllWhoisRequest>): QueryAllWhoisRequest;
};
export declare const QueryAllWhoisResponse: {
    encode(message: QueryAllWhoisResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryAllWhoisResponse;
    fromJSON(object: any): QueryAllWhoisResponse;
    toJSON(message: QueryAllWhoisResponse): unknown;
    fromPartial(object: DeepPartial<QueryAllWhoisResponse>): QueryAllWhoisResponse;
};
export declare const QueryGetNamesRequest: {
    encode(message: QueryGetNamesRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetNamesRequest;
    fromJSON(object: any): QueryGetNamesRequest;
    toJSON(message: QueryGetNamesRequest): unknown;
    fromPartial(object: DeepPartial<QueryGetNamesRequest>): QueryGetNamesRequest;
};
export declare const QueryGetNamesResponse: {
    encode(message: QueryGetNamesResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetNamesResponse;
    fromJSON(object: any): QueryGetNamesResponse;
    toJSON(message: QueryGetNamesResponse): unknown;
    fromPartial(object: DeepPartial<QueryGetNamesResponse>): QueryGetNamesResponse;
};
export declare const QueryAllNamesRequest: {
    encode(message: QueryAllNamesRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryAllNamesRequest;
    fromJSON(object: any): QueryAllNamesRequest;
    toJSON(message: QueryAllNamesRequest): unknown;
    fromPartial(object: DeepPartial<QueryAllNamesRequest>): QueryAllNamesRequest;
};
export declare const QueryAllNamesResponse: {
    encode(message: QueryAllNamesResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryAllNamesResponse;
    fromJSON(object: any): QueryAllNamesResponse;
    toJSON(message: QueryAllNamesResponse): unknown;
    fromPartial(object: DeepPartial<QueryAllNamesResponse>): QueryAllNamesResponse;
};
export declare const QueryGetBidsRequest: {
    encode(message: QueryGetBidsRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetBidsRequest;
    fromJSON(object: any): QueryGetBidsRequest;
    toJSON(message: QueryGetBidsRequest): unknown;
    fromPartial(object: DeepPartial<QueryGetBidsRequest>): QueryGetBidsRequest;
};
export declare const QueryGetBidsResponse: {
    encode(message: QueryGetBidsResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetBidsResponse;
    fromJSON(object: any): QueryGetBidsResponse;
    toJSON(message: QueryGetBidsResponse): unknown;
    fromPartial(object: DeepPartial<QueryGetBidsResponse>): QueryGetBidsResponse;
};
export declare const QueryAllBidsRequest: {
    encode(message: QueryAllBidsRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryAllBidsRequest;
    fromJSON(object: any): QueryAllBidsRequest;
    toJSON(message: QueryAllBidsRequest): unknown;
    fromPartial(object: DeepPartial<QueryAllBidsRequest>): QueryAllBidsRequest;
};
export declare const QueryAllBidsResponse: {
    encode(message: QueryAllBidsResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryAllBidsResponse;
    fromJSON(object: any): QueryAllBidsResponse;
    toJSON(message: QueryAllBidsResponse): unknown;
    fromPartial(object: DeepPartial<QueryAllBidsResponse>): QueryAllBidsResponse;
};
export declare const QueryGetForsaleRequest: {
    encode(message: QueryGetForsaleRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetForsaleRequest;
    fromJSON(object: any): QueryGetForsaleRequest;
    toJSON(message: QueryGetForsaleRequest): unknown;
    fromPartial(object: DeepPartial<QueryGetForsaleRequest>): QueryGetForsaleRequest;
};
export declare const QueryGetForsaleResponse: {
    encode(message: QueryGetForsaleResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetForsaleResponse;
    fromJSON(object: any): QueryGetForsaleResponse;
    toJSON(message: QueryGetForsaleResponse): unknown;
    fromPartial(object: DeepPartial<QueryGetForsaleResponse>): QueryGetForsaleResponse;
};
export declare const QueryAllForsaleRequest: {
    encode(message: QueryAllForsaleRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryAllForsaleRequest;
    fromJSON(object: any): QueryAllForsaleRequest;
    toJSON(message: QueryAllForsaleRequest): unknown;
    fromPartial(object: DeepPartial<QueryAllForsaleRequest>): QueryAllForsaleRequest;
};
export declare const QueryAllForsaleResponse: {
    encode(message: QueryAllForsaleResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryAllForsaleResponse;
    fromJSON(object: any): QueryAllForsaleResponse;
    toJSON(message: QueryAllForsaleResponse): unknown;
    fromPartial(object: DeepPartial<QueryAllForsaleResponse>): QueryAllForsaleResponse;
};
/** Query defines the gRPC querier service. */
export interface Query {
    /** Parameters queries the parameters of the module. */
    Params(request: QueryParamsRequest): Promise<QueryParamsResponse>;
    /** Queries a Name by index. */
    Names(request: QueryGetNamesRequest): Promise<QueryGetNamesResponse>;
    /** Queries a list of Names. */
    NamesAll(request: QueryAllNamesRequest): Promise<QueryAllNamesResponse>;
    /** Queries a Bid by index. */
    Bids(request: QueryGetBidsRequest): Promise<QueryGetBidsResponse>;
    /** Queries a list of Bids. */
    BidsAll(request: QueryAllBidsRequest): Promise<QueryAllBidsResponse>;
    /** Queries a Listing by index. */
    Forsale(request: QueryGetForsaleRequest): Promise<QueryGetForsaleResponse>;
    /** Queries all Listings. */
    ForsaleAll(request: QueryAllForsaleRequest): Promise<QueryAllForsaleResponse>;
}
export declare class QueryClientImpl implements Query {
    private readonly rpc;
    constructor(rpc: Rpc);
    Params(request: QueryParamsRequest): Promise<QueryParamsResponse>;
    Names(request: QueryGetNamesRequest): Promise<QueryGetNamesResponse>;
    NamesAll(request: QueryAllNamesRequest): Promise<QueryAllNamesResponse>;
    Bids(request: QueryGetBidsRequest): Promise<QueryGetBidsResponse>;
    BidsAll(request: QueryAllBidsRequest): Promise<QueryAllBidsResponse>;
    Forsale(request: QueryGetForsaleRequest): Promise<QueryGetForsaleResponse>;
    ForsaleAll(request: QueryAllForsaleRequest): Promise<QueryAllForsaleResponse>;
}
interface Rpc {
    request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
