import { Writer, Reader } from "protobufjs/minimal";
export declare const protobufPackage = "jackaldao.canine.jklmining";
export interface Miners {
    address: string;
    ip: string;
    creator: string;
}
export declare const Miners: {
    encode(message: Miners, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): Miners;
    fromJSON(object: any): Miners;
    toJSON(message: Miners): unknown;
    fromPartial(object: DeepPartial<Miners>): Miners;
};
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
