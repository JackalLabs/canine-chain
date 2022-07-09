import { Writer, Reader } from "protobufjs/minimal";
export declare const protobufPackage = "jackaldao.canine.rns";
export interface Bids {
    index: string;
    name: string;
    bidder: string;
    price: string;
}
export declare const Bids: {
    encode(message: Bids, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): Bids;
    fromJSON(object: any): Bids;
    toJSON(message: Bids): unknown;
    fromPartial(object: DeepPartial<Bids>): Bids;
};
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
