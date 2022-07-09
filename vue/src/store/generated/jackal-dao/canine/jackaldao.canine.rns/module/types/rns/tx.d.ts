import { Reader, Writer } from "protobufjs/minimal";
export declare const protobufPackage = "jackaldao.canine.rns";
export interface MsgRegister {
    creator: string;
    name: string;
    years: string;
    data: string;
}
export interface MsgRegisterResponse {
}
export interface MsgBid {
    creator: string;
    name: string;
    bid: string;
}
export interface MsgBidResponse {
}
export interface MsgAcceptBid {
    creator: string;
    name: string;
    from: string;
}
export interface MsgAcceptBidResponse {
}
export interface MsgCancelBid {
    creator: string;
    name: string;
}
export interface MsgCancelBidResponse {
}
export interface MsgList {
    creator: string;
    name: string;
    price: string;
}
export interface MsgListResponse {
}
export interface MsgBuy {
    creator: string;
    name: string;
}
export interface MsgBuyResponse {
}
export interface MsgDelist {
    creator: string;
    name: string;
}
export interface MsgDelistResponse {
}
export interface MsgTransfer {
    creator: string;
    name: string;
    reciever: string;
}
export interface MsgTransferResponse {
}
export declare const MsgRegister: {
    encode(message: MsgRegister, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgRegister;
    fromJSON(object: any): MsgRegister;
    toJSON(message: MsgRegister): unknown;
    fromPartial(object: DeepPartial<MsgRegister>): MsgRegister;
};
export declare const MsgRegisterResponse: {
    encode(_: MsgRegisterResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgRegisterResponse;
    fromJSON(_: any): MsgRegisterResponse;
    toJSON(_: MsgRegisterResponse): unknown;
    fromPartial(_: DeepPartial<MsgRegisterResponse>): MsgRegisterResponse;
};
export declare const MsgBid: {
    encode(message: MsgBid, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgBid;
    fromJSON(object: any): MsgBid;
    toJSON(message: MsgBid): unknown;
    fromPartial(object: DeepPartial<MsgBid>): MsgBid;
};
export declare const MsgBidResponse: {
    encode(_: MsgBidResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgBidResponse;
    fromJSON(_: any): MsgBidResponse;
    toJSON(_: MsgBidResponse): unknown;
    fromPartial(_: DeepPartial<MsgBidResponse>): MsgBidResponse;
};
export declare const MsgAcceptBid: {
    encode(message: MsgAcceptBid, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgAcceptBid;
    fromJSON(object: any): MsgAcceptBid;
    toJSON(message: MsgAcceptBid): unknown;
    fromPartial(object: DeepPartial<MsgAcceptBid>): MsgAcceptBid;
};
export declare const MsgAcceptBidResponse: {
    encode(_: MsgAcceptBidResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgAcceptBidResponse;
    fromJSON(_: any): MsgAcceptBidResponse;
    toJSON(_: MsgAcceptBidResponse): unknown;
    fromPartial(_: DeepPartial<MsgAcceptBidResponse>): MsgAcceptBidResponse;
};
export declare const MsgCancelBid: {
    encode(message: MsgCancelBid, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgCancelBid;
    fromJSON(object: any): MsgCancelBid;
    toJSON(message: MsgCancelBid): unknown;
    fromPartial(object: DeepPartial<MsgCancelBid>): MsgCancelBid;
};
export declare const MsgCancelBidResponse: {
    encode(_: MsgCancelBidResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgCancelBidResponse;
    fromJSON(_: any): MsgCancelBidResponse;
    toJSON(_: MsgCancelBidResponse): unknown;
    fromPartial(_: DeepPartial<MsgCancelBidResponse>): MsgCancelBidResponse;
};
export declare const MsgList: {
    encode(message: MsgList, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgList;
    fromJSON(object: any): MsgList;
    toJSON(message: MsgList): unknown;
    fromPartial(object: DeepPartial<MsgList>): MsgList;
};
export declare const MsgListResponse: {
    encode(_: MsgListResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgListResponse;
    fromJSON(_: any): MsgListResponse;
    toJSON(_: MsgListResponse): unknown;
    fromPartial(_: DeepPartial<MsgListResponse>): MsgListResponse;
};
export declare const MsgBuy: {
    encode(message: MsgBuy, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgBuy;
    fromJSON(object: any): MsgBuy;
    toJSON(message: MsgBuy): unknown;
    fromPartial(object: DeepPartial<MsgBuy>): MsgBuy;
};
export declare const MsgBuyResponse: {
    encode(_: MsgBuyResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgBuyResponse;
    fromJSON(_: any): MsgBuyResponse;
    toJSON(_: MsgBuyResponse): unknown;
    fromPartial(_: DeepPartial<MsgBuyResponse>): MsgBuyResponse;
};
export declare const MsgDelist: {
    encode(message: MsgDelist, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgDelist;
    fromJSON(object: any): MsgDelist;
    toJSON(message: MsgDelist): unknown;
    fromPartial(object: DeepPartial<MsgDelist>): MsgDelist;
};
export declare const MsgDelistResponse: {
    encode(_: MsgDelistResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgDelistResponse;
    fromJSON(_: any): MsgDelistResponse;
    toJSON(_: MsgDelistResponse): unknown;
    fromPartial(_: DeepPartial<MsgDelistResponse>): MsgDelistResponse;
};
export declare const MsgTransfer: {
    encode(message: MsgTransfer, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgTransfer;
    fromJSON(object: any): MsgTransfer;
    toJSON(message: MsgTransfer): unknown;
    fromPartial(object: DeepPartial<MsgTransfer>): MsgTransfer;
};
export declare const MsgTransferResponse: {
    encode(_: MsgTransferResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgTransferResponse;
    fromJSON(_: any): MsgTransferResponse;
    toJSON(_: MsgTransferResponse): unknown;
    fromPartial(_: DeepPartial<MsgTransferResponse>): MsgTransferResponse;
};
/** Msg defines the Msg service. */
export interface Msg {
    Register(request: MsgRegister): Promise<MsgRegisterResponse>;
    Bid(request: MsgBid): Promise<MsgBidResponse>;
    AcceptBid(request: MsgAcceptBid): Promise<MsgAcceptBidResponse>;
    CancelBid(request: MsgCancelBid): Promise<MsgCancelBidResponse>;
    List(request: MsgList): Promise<MsgListResponse>;
    Buy(request: MsgBuy): Promise<MsgBuyResponse>;
    Delist(request: MsgDelist): Promise<MsgDelistResponse>;
    /** this line is used by starport scaffolding # proto/tx/rpc */
    Transfer(request: MsgTransfer): Promise<MsgTransferResponse>;
}
export declare class MsgClientImpl implements Msg {
    private readonly rpc;
    constructor(rpc: Rpc);
    Register(request: MsgRegister): Promise<MsgRegisterResponse>;
    Bid(request: MsgBid): Promise<MsgBidResponse>;
    AcceptBid(request: MsgAcceptBid): Promise<MsgAcceptBidResponse>;
    CancelBid(request: MsgCancelBid): Promise<MsgCancelBidResponse>;
    List(request: MsgList): Promise<MsgListResponse>;
    Buy(request: MsgBuy): Promise<MsgBuyResponse>;
    Delist(request: MsgDelist): Promise<MsgDelistResponse>;
    Transfer(request: MsgTransfer): Promise<MsgTransferResponse>;
}
interface Rpc {
    request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
