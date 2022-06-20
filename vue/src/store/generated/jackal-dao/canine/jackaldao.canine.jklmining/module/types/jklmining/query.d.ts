import { Reader, Writer } from "protobufjs/minimal";
import { Params } from "../jklmining/params";
import { SaveRequests } from "../jklmining/save_requests";
import { PageRequest, PageResponse } from "../cosmos/base/query/v1beta1/pagination";
import { Miners } from "../jklmining/miners";
import { Mined } from "../jklmining/mined";
import { MinerClaims } from "../jklmining/miner_claims";
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
export interface QueryCheckMinerIndexRequest {
}
export interface QueryCheckMinerIndexResponse {
}
export interface QueryGetMinerIndexRequest {
    index: string;
}
export interface QueryGetMinerIndexResponse {
}
export interface QueryGetMinerStartRequest {
}
export interface QueryGetMinerStartResponse {
    index: string;
}
export interface QueryGetMinerClaimsRequest {
    hash: string;
}
export interface QueryGetMinerClaimsResponse {
    minerClaims: MinerClaims | undefined;
}
export interface QueryAllMinerClaimsRequest {
    pagination: PageRequest | undefined;
}
export interface QueryAllMinerClaimsResponse {
    minerClaims: MinerClaims[];
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
export declare const QueryCheckMinerIndexRequest: {
    encode(_: QueryCheckMinerIndexRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryCheckMinerIndexRequest;
    fromJSON(_: any): QueryCheckMinerIndexRequest;
    toJSON(_: QueryCheckMinerIndexRequest): unknown;
    fromPartial(_: DeepPartial<QueryCheckMinerIndexRequest>): QueryCheckMinerIndexRequest;
};
export declare const QueryCheckMinerIndexResponse: {
    encode(_: QueryCheckMinerIndexResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryCheckMinerIndexResponse;
    fromJSON(_: any): QueryCheckMinerIndexResponse;
    toJSON(_: QueryCheckMinerIndexResponse): unknown;
    fromPartial(_: DeepPartial<QueryCheckMinerIndexResponse>): QueryCheckMinerIndexResponse;
};
export declare const QueryGetMinerIndexRequest: {
    encode(message: QueryGetMinerIndexRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetMinerIndexRequest;
    fromJSON(object: any): QueryGetMinerIndexRequest;
    toJSON(message: QueryGetMinerIndexRequest): unknown;
    fromPartial(object: DeepPartial<QueryGetMinerIndexRequest>): QueryGetMinerIndexRequest;
};
export declare const QueryGetMinerIndexResponse: {
    encode(_: QueryGetMinerIndexResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetMinerIndexResponse;
    fromJSON(_: any): QueryGetMinerIndexResponse;
    toJSON(_: QueryGetMinerIndexResponse): unknown;
    fromPartial(_: DeepPartial<QueryGetMinerIndexResponse>): QueryGetMinerIndexResponse;
};
export declare const QueryGetMinerStartRequest: {
    encode(_: QueryGetMinerStartRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetMinerStartRequest;
    fromJSON(_: any): QueryGetMinerStartRequest;
    toJSON(_: QueryGetMinerStartRequest): unknown;
    fromPartial(_: DeepPartial<QueryGetMinerStartRequest>): QueryGetMinerStartRequest;
};
export declare const QueryGetMinerStartResponse: {
    encode(message: QueryGetMinerStartResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetMinerStartResponse;
    fromJSON(object: any): QueryGetMinerStartResponse;
    toJSON(message: QueryGetMinerStartResponse): unknown;
    fromPartial(object: DeepPartial<QueryGetMinerStartResponse>): QueryGetMinerStartResponse;
};
export declare const QueryGetMinerClaimsRequest: {
    encode(message: QueryGetMinerClaimsRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetMinerClaimsRequest;
    fromJSON(object: any): QueryGetMinerClaimsRequest;
    toJSON(message: QueryGetMinerClaimsRequest): unknown;
    fromPartial(object: DeepPartial<QueryGetMinerClaimsRequest>): QueryGetMinerClaimsRequest;
};
export declare const QueryGetMinerClaimsResponse: {
    encode(message: QueryGetMinerClaimsResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetMinerClaimsResponse;
    fromJSON(object: any): QueryGetMinerClaimsResponse;
    toJSON(message: QueryGetMinerClaimsResponse): unknown;
    fromPartial(object: DeepPartial<QueryGetMinerClaimsResponse>): QueryGetMinerClaimsResponse;
};
export declare const QueryAllMinerClaimsRequest: {
    encode(message: QueryAllMinerClaimsRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryAllMinerClaimsRequest;
    fromJSON(object: any): QueryAllMinerClaimsRequest;
    toJSON(message: QueryAllMinerClaimsRequest): unknown;
    fromPartial(object: DeepPartial<QueryAllMinerClaimsRequest>): QueryAllMinerClaimsRequest;
};
export declare const QueryAllMinerClaimsResponse: {
    encode(message: QueryAllMinerClaimsResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryAllMinerClaimsResponse;
    fromJSON(object: any): QueryAllMinerClaimsResponse;
    toJSON(message: QueryAllMinerClaimsResponse): unknown;
    fromPartial(object: DeepPartial<QueryAllMinerClaimsResponse>): QueryAllMinerClaimsResponse;
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
    /** Queries a list of CheckMinerIndex items. */
    CheckMinerIndex(request: QueryCheckMinerIndexRequest): Promise<QueryCheckMinerIndexResponse>;
    /** Queries a list of GetMinerIndex items. */
    GetMinerIndex(request: QueryGetMinerIndexRequest): Promise<QueryGetMinerIndexResponse>;
    /** Queries a list of GetMinerStart items. */
    GetMinerStart(request: QueryGetMinerStartRequest): Promise<QueryGetMinerStartResponse>;
    /** Queries a MinerClaims by index. */
    MinerClaims(request: QueryGetMinerClaimsRequest): Promise<QueryGetMinerClaimsResponse>;
    /** Queries a list of MinerClaims items. */
    MinerClaimsAll(request: QueryAllMinerClaimsRequest): Promise<QueryAllMinerClaimsResponse>;
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
    CheckMinerIndex(request: QueryCheckMinerIndexRequest): Promise<QueryCheckMinerIndexResponse>;
    GetMinerIndex(request: QueryGetMinerIndexRequest): Promise<QueryGetMinerIndexResponse>;
    GetMinerStart(request: QueryGetMinerStartRequest): Promise<QueryGetMinerStartResponse>;
    MinerClaims(request: QueryGetMinerClaimsRequest): Promise<QueryGetMinerClaimsResponse>;
    MinerClaimsAll(request: QueryAllMinerClaimsRequest): Promise<QueryAllMinerClaimsResponse>;
}
interface Rpc {
    request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
