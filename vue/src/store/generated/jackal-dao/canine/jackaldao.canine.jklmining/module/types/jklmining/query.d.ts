import { Reader, Writer } from "protobufjs/minimal";
import { Params } from "../jklmining/params";
import { SaveRequests } from "../jklmining/save_requests";
import { PageRequest, PageResponse } from "../cosmos/base/query/v1beta1/pagination";
import { Miners } from "../jklmining/miners";
import { Mined } from "../jklmining/mined";
export declare const protobufPackage = "jackaldao.canine.jklmining";
/** QueryParamsRequest is request type for the Query/Params RPC method. */
export interface QueryParamsRequest {
}
/** QueryParamsResponse is response type for the Query/Params RPC method. */
export interface QueryParamsResponse {
    /** params holds all the parameters of this module. */
    params: Params | undefined;
}
export interface QueryGetSaveRequestsRequest {
    index: string;
}
export interface QueryGetSaveRequestsResponse {
    saveRequests: SaveRequests | undefined;
}
export interface QueryAllSaveRequestsRequest {
    pagination: PageRequest | undefined;
}
export interface QueryAllSaveRequestsResponse {
    saveRequests: SaveRequests[];
    pagination: PageResponse | undefined;
}
export interface QueryGetMinersRequest {
    address: string;
}
export interface QueryGetMinersResponse {
    miners: Miners | undefined;
}
export interface QueryAllMinersRequest {
    pagination: PageRequest | undefined;
}
export interface QueryAllMinersResponse {
    miners: Miners[];
    pagination: PageResponse | undefined;
}
export interface QueryGetMinedRequest {
    id: number;
}
export interface QueryGetMinedResponse {
    Mined: Mined | undefined;
}
export interface QueryAllMinedRequest {
    pagination: PageRequest | undefined;
}
export interface QueryAllMinedResponse {
    Mined: Mined[];
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
export declare const QueryGetSaveRequestsRequest: {
    encode(message: QueryGetSaveRequestsRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetSaveRequestsRequest;
    fromJSON(object: any): QueryGetSaveRequestsRequest;
    toJSON(message: QueryGetSaveRequestsRequest): unknown;
    fromPartial(object: DeepPartial<QueryGetSaveRequestsRequest>): QueryGetSaveRequestsRequest;
};
export declare const QueryGetSaveRequestsResponse: {
    encode(message: QueryGetSaveRequestsResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetSaveRequestsResponse;
    fromJSON(object: any): QueryGetSaveRequestsResponse;
    toJSON(message: QueryGetSaveRequestsResponse): unknown;
    fromPartial(object: DeepPartial<QueryGetSaveRequestsResponse>): QueryGetSaveRequestsResponse;
};
export declare const QueryAllSaveRequestsRequest: {
    encode(message: QueryAllSaveRequestsRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryAllSaveRequestsRequest;
    fromJSON(object: any): QueryAllSaveRequestsRequest;
    toJSON(message: QueryAllSaveRequestsRequest): unknown;
    fromPartial(object: DeepPartial<QueryAllSaveRequestsRequest>): QueryAllSaveRequestsRequest;
};
export declare const QueryAllSaveRequestsResponse: {
    encode(message: QueryAllSaveRequestsResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryAllSaveRequestsResponse;
    fromJSON(object: any): QueryAllSaveRequestsResponse;
    toJSON(message: QueryAllSaveRequestsResponse): unknown;
    fromPartial(object: DeepPartial<QueryAllSaveRequestsResponse>): QueryAllSaveRequestsResponse;
};
export declare const QueryGetMinersRequest: {
    encode(message: QueryGetMinersRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetMinersRequest;
    fromJSON(object: any): QueryGetMinersRequest;
    toJSON(message: QueryGetMinersRequest): unknown;
    fromPartial(object: DeepPartial<QueryGetMinersRequest>): QueryGetMinersRequest;
};
export declare const QueryGetMinersResponse: {
    encode(message: QueryGetMinersResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetMinersResponse;
    fromJSON(object: any): QueryGetMinersResponse;
    toJSON(message: QueryGetMinersResponse): unknown;
    fromPartial(object: DeepPartial<QueryGetMinersResponse>): QueryGetMinersResponse;
};
export declare const QueryAllMinersRequest: {
    encode(message: QueryAllMinersRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryAllMinersRequest;
    fromJSON(object: any): QueryAllMinersRequest;
    toJSON(message: QueryAllMinersRequest): unknown;
    fromPartial(object: DeepPartial<QueryAllMinersRequest>): QueryAllMinersRequest;
};
export declare const QueryAllMinersResponse: {
    encode(message: QueryAllMinersResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryAllMinersResponse;
    fromJSON(object: any): QueryAllMinersResponse;
    toJSON(message: QueryAllMinersResponse): unknown;
    fromPartial(object: DeepPartial<QueryAllMinersResponse>): QueryAllMinersResponse;
};
export declare const QueryGetMinedRequest: {
    encode(message: QueryGetMinedRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetMinedRequest;
    fromJSON(object: any): QueryGetMinedRequest;
    toJSON(message: QueryGetMinedRequest): unknown;
    fromPartial(object: DeepPartial<QueryGetMinedRequest>): QueryGetMinedRequest;
};
export declare const QueryGetMinedResponse: {
    encode(message: QueryGetMinedResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetMinedResponse;
    fromJSON(object: any): QueryGetMinedResponse;
    toJSON(message: QueryGetMinedResponse): unknown;
    fromPartial(object: DeepPartial<QueryGetMinedResponse>): QueryGetMinedResponse;
};
export declare const QueryAllMinedRequest: {
    encode(message: QueryAllMinedRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryAllMinedRequest;
    fromJSON(object: any): QueryAllMinedRequest;
    toJSON(message: QueryAllMinedRequest): unknown;
    fromPartial(object: DeepPartial<QueryAllMinedRequest>): QueryAllMinedRequest;
};
export declare const QueryAllMinedResponse: {
    encode(message: QueryAllMinedResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryAllMinedResponse;
    fromJSON(object: any): QueryAllMinedResponse;
    toJSON(message: QueryAllMinedResponse): unknown;
    fromPartial(object: DeepPartial<QueryAllMinedResponse>): QueryAllMinedResponse;
};
/** Query defines the gRPC querier service. */
export interface Query {
    /** Parameters queries the parameters of the module. */
    Params(request: QueryParamsRequest): Promise<QueryParamsResponse>;
    /** Queries a SaveRequests by index. */
    SaveRequests(request: QueryGetSaveRequestsRequest): Promise<QueryGetSaveRequestsResponse>;
    /** Queries a list of SaveRequests items. */
    SaveRequestsAll(request: QueryAllSaveRequestsRequest): Promise<QueryAllSaveRequestsResponse>;
    /** Queries a Miners by index. */
    Miners(request: QueryGetMinersRequest): Promise<QueryGetMinersResponse>;
    /** Queries a list of Miners items. */
    MinersAll(request: QueryAllMinersRequest): Promise<QueryAllMinersResponse>;
    /** Queries a Mined by id. */
    Mined(request: QueryGetMinedRequest): Promise<QueryGetMinedResponse>;
    /** Queries a list of Mined items. */
    MinedAll(request: QueryAllMinedRequest): Promise<QueryAllMinedResponse>;
}
export declare class QueryClientImpl implements Query {
    private readonly rpc;
    constructor(rpc: Rpc);
    Params(request: QueryParamsRequest): Promise<QueryParamsResponse>;
    SaveRequests(request: QueryGetSaveRequestsRequest): Promise<QueryGetSaveRequestsResponse>;
    SaveRequestsAll(request: QueryAllSaveRequestsRequest): Promise<QueryAllSaveRequestsResponse>;
    Miners(request: QueryGetMinersRequest): Promise<QueryGetMinersResponse>;
    MinersAll(request: QueryAllMinersRequest): Promise<QueryAllMinersResponse>;
    Mined(request: QueryGetMinedRequest): Promise<QueryGetMinedResponse>;
    MinedAll(request: QueryAllMinedRequest): Promise<QueryAllMinedResponse>;
}
interface Rpc {
    request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
