import { Writer, Reader } from "protobufjs/minimal";
export declare const protobufPackage = "jackaldao.canine.jklmining";
export interface MinerClaims {
    hash: string;
    creator: string;
}
export declare const MinerClaims: {
    encode(message: MinerClaims, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MinerClaims;
    fromJSON(object: any): MinerClaims;
    toJSON(message: MinerClaims): unknown;
    fromPartial(object: DeepPartial<MinerClaims>): MinerClaims;
};
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
