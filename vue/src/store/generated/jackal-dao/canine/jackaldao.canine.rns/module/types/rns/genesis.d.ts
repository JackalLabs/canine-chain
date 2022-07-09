import { Params } from "../rns/params";
import { Whois } from "../rns/whois";
import { Names } from "../rns/names";
import { Bids } from "../rns/bids";
import { Forsale } from "../rns/forsale";
import { Writer, Reader } from "protobufjs/minimal";
export declare const protobufPackage = "jackaldao.canine.rns";
/** GenesisState defines the rns module's genesis state. */
export interface GenesisState {
    params: Params | undefined;
    whoisList: Whois[];
    namesList: Names[];
    bidsList: Bids[];
    /** this line is used by starport scaffolding # genesis/proto/state */
    forsaleList: Forsale[];
}
export declare const GenesisState: {
    encode(message: GenesisState, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): GenesisState;
    fromJSON(object: any): GenesisState;
    toJSON(message: GenesisState): unknown;
    fromPartial(object: DeepPartial<GenesisState>): GenesisState;
};
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
