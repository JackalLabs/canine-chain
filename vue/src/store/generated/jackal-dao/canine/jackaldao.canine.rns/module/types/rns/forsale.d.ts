import { Writer, Reader } from "protobufjs/minimal";
export declare const protobufPackage = "jackaldao.canine.rns";
export interface Forsale {
    name: string;
    price: string;
    owner: string;
}
export declare const Forsale: {
    encode(message: Forsale, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): Forsale;
    fromJSON(object: any): Forsale;
    toJSON(message: Forsale): unknown;
    fromPartial(object: DeepPartial<Forsale>): Forsale;
};
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
