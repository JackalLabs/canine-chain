import { Writer, Reader } from "protobufjs/minimal";
export declare const protobufPackage = "jackaldao.canine.jklaccounts";
export interface Accounts {
    address: string;
    available: string;
    used: string;
    expireBlock: string;
}
export declare const Accounts: {
    encode(message: Accounts, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): Accounts;
    fromJSON(object: any): Accounts;
    toJSON(message: Accounts): unknown;
    fromPartial(object: DeepPartial<Accounts>): Accounts;
};
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
