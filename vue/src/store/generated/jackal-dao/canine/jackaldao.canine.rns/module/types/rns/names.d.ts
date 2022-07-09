import { Writer, Reader } from "protobufjs/minimal";
export declare const protobufPackage = "jackaldao.canine.rns";
export interface Names {
    index: string;
    name: string;
    expires: string;
    value: string;
    data: string;
}
export declare const Names: {
    encode(message: Names, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): Names;
    fromJSON(object: any): Names;
    toJSON(message: Names): unknown;
    fromPartial(object: DeepPartial<Names>): Names;
};
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
