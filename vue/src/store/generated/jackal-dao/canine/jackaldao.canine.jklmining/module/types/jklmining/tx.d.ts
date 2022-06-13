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
/** Msg defines the Msg service. */
export interface Msg {
    AllowSave(request: MsgAllowSave): Promise<MsgAllowSaveResponse>;
    CreateSaveRequests(request: MsgCreateSaveRequests): Promise<MsgCreateSaveRequestsResponse>;
    UpdateSaveRequests(request: MsgUpdateSaveRequests): Promise<MsgUpdateSaveRequestsResponse>;
    /** this line is used by starport scaffolding # proto/tx/rpc */
    DeleteSaveRequests(request: MsgDeleteSaveRequests): Promise<MsgDeleteSaveRequestsResponse>;
}
export declare class MsgClientImpl implements Msg {
    private readonly rpc;
    constructor(rpc: Rpc);
    AllowSave(request: MsgAllowSave): Promise<MsgAllowSaveResponse>;
    CreateSaveRequests(request: MsgCreateSaveRequests): Promise<MsgCreateSaveRequestsResponse>;
    UpdateSaveRequests(request: MsgUpdateSaveRequests): Promise<MsgUpdateSaveRequestsResponse>;
    DeleteSaveRequests(request: MsgDeleteSaveRequests): Promise<MsgDeleteSaveRequestsResponse>;
}
interface Rpc {
    request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
