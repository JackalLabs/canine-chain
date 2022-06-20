import { Writer, Reader } from "protobufjs/minimal";
export declare const protobufPackage = "jackaldao.canine.jklmining";
export interface Mined {
    id: number;
    datasize: string;
    hash: string;
    pcount: string;
}
export declare const Mined: {
    encode(message: Mined, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): Mined;
    fromJSON(object: any): Mined;
    toJSON(message: Mined): unknown;
    fromPartial(object: DeepPartial<Mined>): Mined;
};
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
