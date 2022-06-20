import { Reader, Writer } from "protobufjs/minimal";
export declare const protobufPackage = "jackaldao.canine.jklmining";
export interface MsgAllowSave {
    creator: string;
    passkey: string;
    size: string;
}
export interface MsgAllowSaveResponse {
}
export interface MsgCreateSaveRequests {
    creator: string;
    index: string;
    size: string;
    approved: string;
}
export interface MsgCreateSaveRequestsResponse {
}
export interface MsgUpdateSaveRequests {
    creator: string;
    index: string;
    size: string;
    approved: string;
}
export interface MsgUpdateSaveRequestsResponse {
}
export interface MsgDeleteSaveRequests {
    creator: string;
    index: string;
}
export interface MsgDeleteSaveRequestsResponse {
}
export interface MsgCreateMiners {
    creator: string;
    address: string;
    ip: string;
}
export interface MsgCreateMinersResponse {
}
export interface MsgUpdateMiners {
    creator: string;
    address: string;
    ip: string;
}
export interface MsgUpdateMinersResponse {
}
export interface MsgDeleteMiners {
    creator: string;
    address: string;
}
export interface MsgDeleteMinersResponse {
}
export interface MsgClaimSave {
    creator: string;
    saveindex: string;
    key: string;
}
export interface MsgClaimSaveResponse {
}
export interface MsgCreateMinerClaims {
    creator: string;
    hash: string;
}
export interface MsgCreateMinerClaimsResponse {
}
export interface MsgUpdateMinerClaims {
    creator: string;
    hash: string;
}
export interface MsgUpdateMinerClaimsResponse {
}
export interface MsgDeleteMinerClaims {
    creator: string;
    hash: string;
}
export interface MsgDeleteMinerClaimsResponse {
}
export declare const MsgAllowSave: {
    encode(message: MsgAllowSave, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgAllowSave;
    fromJSON(object: any): MsgAllowSave;
    toJSON(message: MsgAllowSave): unknown;
    fromPartial(object: DeepPartial<MsgAllowSave>): MsgAllowSave;
};
export declare const MsgAllowSaveResponse: {
    encode(_: MsgAllowSaveResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgAllowSaveResponse;
    fromJSON(_: any): MsgAllowSaveResponse;
    toJSON(_: MsgAllowSaveResponse): unknown;
    fromPartial(_: DeepPartial<MsgAllowSaveResponse>): MsgAllowSaveResponse;
};
export declare const MsgCreateSaveRequests: {
    encode(message: MsgCreateSaveRequests, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgCreateSaveRequests;
    fromJSON(object: any): MsgCreateSaveRequests;
    toJSON(message: MsgCreateSaveRequests): unknown;
    fromPartial(object: DeepPartial<MsgCreateSaveRequests>): MsgCreateSaveRequests;
};
export declare const MsgCreateSaveRequestsResponse: {
    encode(_: MsgCreateSaveRequestsResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgCreateSaveRequestsResponse;
    fromJSON(_: any): MsgCreateSaveRequestsResponse;
    toJSON(_: MsgCreateSaveRequestsResponse): unknown;
    fromPartial(_: DeepPartial<MsgCreateSaveRequestsResponse>): MsgCreateSaveRequestsResponse;
};
export declare const MsgUpdateSaveRequests: {
    encode(message: MsgUpdateSaveRequests, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgUpdateSaveRequests;
    fromJSON(object: any): MsgUpdateSaveRequests;
    toJSON(message: MsgUpdateSaveRequests): unknown;
    fromPartial(object: DeepPartial<MsgUpdateSaveRequests>): MsgUpdateSaveRequests;
};
export declare const MsgUpdateSaveRequestsResponse: {
    encode(_: MsgUpdateSaveRequestsResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgUpdateSaveRequestsResponse;
    fromJSON(_: any): MsgUpdateSaveRequestsResponse;
    toJSON(_: MsgUpdateSaveRequestsResponse): unknown;
    fromPartial(_: DeepPartial<MsgUpdateSaveRequestsResponse>): MsgUpdateSaveRequestsResponse;
};
export declare const MsgDeleteSaveRequests: {
    encode(message: MsgDeleteSaveRequests, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgDeleteSaveRequests;
    fromJSON(object: any): MsgDeleteSaveRequests;
    toJSON(message: MsgDeleteSaveRequests): unknown;
    fromPartial(object: DeepPartial<MsgDeleteSaveRequests>): MsgDeleteSaveRequests;
};
export declare const MsgDeleteSaveRequestsResponse: {
    encode(_: MsgDeleteSaveRequestsResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgDeleteSaveRequestsResponse;
    fromJSON(_: any): MsgDeleteSaveRequestsResponse;
    toJSON(_: MsgDeleteSaveRequestsResponse): unknown;
    fromPartial(_: DeepPartial<MsgDeleteSaveRequestsResponse>): MsgDeleteSaveRequestsResponse;
};
export declare const MsgCreateMiners: {
    encode(message: MsgCreateMiners, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgCreateMiners;
    fromJSON(object: any): MsgCreateMiners;
    toJSON(message: MsgCreateMiners): unknown;
    fromPartial(object: DeepPartial<MsgCreateMiners>): MsgCreateMiners;
};
export declare const MsgCreateMinersResponse: {
    encode(_: MsgCreateMinersResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgCreateMinersResponse;
    fromJSON(_: any): MsgCreateMinersResponse;
    toJSON(_: MsgCreateMinersResponse): unknown;
    fromPartial(_: DeepPartial<MsgCreateMinersResponse>): MsgCreateMinersResponse;
};
export declare const MsgUpdateMiners: {
    encode(message: MsgUpdateMiners, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgUpdateMiners;
    fromJSON(object: any): MsgUpdateMiners;
    toJSON(message: MsgUpdateMiners): unknown;
    fromPartial(object: DeepPartial<MsgUpdateMiners>): MsgUpdateMiners;
};
export declare const MsgUpdateMinersResponse: {
    encode(_: MsgUpdateMinersResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgUpdateMinersResponse;
    fromJSON(_: any): MsgUpdateMinersResponse;
    toJSON(_: MsgUpdateMinersResponse): unknown;
    fromPartial(_: DeepPartial<MsgUpdateMinersResponse>): MsgUpdateMinersResponse;
};
export declare const MsgDeleteMiners: {
    encode(message: MsgDeleteMiners, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgDeleteMiners;
    fromJSON(object: any): MsgDeleteMiners;
    toJSON(message: MsgDeleteMiners): unknown;
    fromPartial(object: DeepPartial<MsgDeleteMiners>): MsgDeleteMiners;
};
export declare const MsgDeleteMinersResponse: {
    encode(_: MsgDeleteMinersResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgDeleteMinersResponse;
    fromJSON(_: any): MsgDeleteMinersResponse;
    toJSON(_: MsgDeleteMinersResponse): unknown;
    fromPartial(_: DeepPartial<MsgDeleteMinersResponse>): MsgDeleteMinersResponse;
};
export declare const MsgClaimSave: {
    encode(message: MsgClaimSave, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgClaimSave;
    fromJSON(object: any): MsgClaimSave;
    toJSON(message: MsgClaimSave): unknown;
    fromPartial(object: DeepPartial<MsgClaimSave>): MsgClaimSave;
};
export declare const MsgClaimSaveResponse: {
    encode(_: MsgClaimSaveResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgClaimSaveResponse;
    fromJSON(_: any): MsgClaimSaveResponse;
    toJSON(_: MsgClaimSaveResponse): unknown;
    fromPartial(_: DeepPartial<MsgClaimSaveResponse>): MsgClaimSaveResponse;
};
export declare const MsgCreateMinerClaims: {
    encode(message: MsgCreateMinerClaims, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgCreateMinerClaims;
    fromJSON(object: any): MsgCreateMinerClaims;
    toJSON(message: MsgCreateMinerClaims): unknown;
    fromPartial(object: DeepPartial<MsgCreateMinerClaims>): MsgCreateMinerClaims;
};
export declare const MsgCreateMinerClaimsResponse: {
    encode(_: MsgCreateMinerClaimsResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgCreateMinerClaimsResponse;
    fromJSON(_: any): MsgCreateMinerClaimsResponse;
    toJSON(_: MsgCreateMinerClaimsResponse): unknown;
    fromPartial(_: DeepPartial<MsgCreateMinerClaimsResponse>): MsgCreateMinerClaimsResponse;
};
export declare const MsgUpdateMinerClaims: {
    encode(message: MsgUpdateMinerClaims, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgUpdateMinerClaims;
    fromJSON(object: any): MsgUpdateMinerClaims;
    toJSON(message: MsgUpdateMinerClaims): unknown;
    fromPartial(object: DeepPartial<MsgUpdateMinerClaims>): MsgUpdateMinerClaims;
};
export declare const MsgUpdateMinerClaimsResponse: {
    encode(_: MsgUpdateMinerClaimsResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgUpdateMinerClaimsResponse;
    fromJSON(_: any): MsgUpdateMinerClaimsResponse;
    toJSON(_: MsgUpdateMinerClaimsResponse): unknown;
    fromPartial(_: DeepPartial<MsgUpdateMinerClaimsResponse>): MsgUpdateMinerClaimsResponse;
};
export declare const MsgDeleteMinerClaims: {
    encode(message: MsgDeleteMinerClaims, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgDeleteMinerClaims;
    fromJSON(object: any): MsgDeleteMinerClaims;
    toJSON(message: MsgDeleteMinerClaims): unknown;
    fromPartial(object: DeepPartial<MsgDeleteMinerClaims>): MsgDeleteMinerClaims;
};
export declare const MsgDeleteMinerClaimsResponse: {
    encode(_: MsgDeleteMinerClaimsResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgDeleteMinerClaimsResponse;
    fromJSON(_: any): MsgDeleteMinerClaimsResponse;
    toJSON(_: MsgDeleteMinerClaimsResponse): unknown;
    fromPartial(_: DeepPartial<MsgDeleteMinerClaimsResponse>): MsgDeleteMinerClaimsResponse;
};
/** Msg defines the Msg service. */
export interface Msg {
    AllowSave(request: MsgAllowSave): Promise<MsgAllowSaveResponse>;
    CreateSaveRequests(request: MsgCreateSaveRequests): Promise<MsgCreateSaveRequestsResponse>;
    UpdateSaveRequests(request: MsgUpdateSaveRequests): Promise<MsgUpdateSaveRequestsResponse>;
    DeleteSaveRequests(request: MsgDeleteSaveRequests): Promise<MsgDeleteSaveRequestsResponse>;
    CreateMiners(request: MsgCreateMiners): Promise<MsgCreateMinersResponse>;
    UpdateMiners(request: MsgUpdateMiners): Promise<MsgUpdateMinersResponse>;
    DeleteMiners(request: MsgDeleteMiners): Promise<MsgDeleteMinersResponse>;
    ClaimSave(request: MsgClaimSave): Promise<MsgClaimSaveResponse>;
    CreateMinerClaims(request: MsgCreateMinerClaims): Promise<MsgCreateMinerClaimsResponse>;
    UpdateMinerClaims(request: MsgUpdateMinerClaims): Promise<MsgUpdateMinerClaimsResponse>;
    /** this line is used by starport scaffolding # proto/tx/rpc */
    DeleteMinerClaims(request: MsgDeleteMinerClaims): Promise<MsgDeleteMinerClaimsResponse>;
}
export declare class MsgClientImpl implements Msg {
    private readonly rpc;
    constructor(rpc: Rpc);
    AllowSave(request: MsgAllowSave): Promise<MsgAllowSaveResponse>;
    CreateSaveRequests(request: MsgCreateSaveRequests): Promise<MsgCreateSaveRequestsResponse>;
    UpdateSaveRequests(request: MsgUpdateSaveRequests): Promise<MsgUpdateSaveRequestsResponse>;
    DeleteSaveRequests(request: MsgDeleteSaveRequests): Promise<MsgDeleteSaveRequestsResponse>;
    CreateMiners(request: MsgCreateMiners): Promise<MsgCreateMinersResponse>;
    UpdateMiners(request: MsgUpdateMiners): Promise<MsgUpdateMinersResponse>;
    DeleteMiners(request: MsgDeleteMiners): Promise<MsgDeleteMinersResponse>;
    ClaimSave(request: MsgClaimSave): Promise<MsgClaimSaveResponse>;
    CreateMinerClaims(request: MsgCreateMinerClaims): Promise<MsgCreateMinerClaimsResponse>;
    UpdateMinerClaims(request: MsgUpdateMinerClaims): Promise<MsgUpdateMinerClaimsResponse>;
    DeleteMinerClaims(request: MsgDeleteMinerClaims): Promise<MsgDeleteMinerClaimsResponse>;
}
interface Rpc {
    request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
