import { Writer, Reader } from "protobufjs/minimal";
export declare const protobufPackage = "jackaldao.canine.jklmining";
export interface SaveRequests {
    index: string;
    size: string;
    approved: string;
    creator: string;
}
export declare const SaveRequests: {
    encode(message: SaveRequests, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): SaveRequests;
    fromJSON(object: any): SaveRequests;
    toJSON(message: SaveRequests): unknown;
    fromPartial(object: DeepPartial<SaveRequests>): SaveRequests;
};
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
