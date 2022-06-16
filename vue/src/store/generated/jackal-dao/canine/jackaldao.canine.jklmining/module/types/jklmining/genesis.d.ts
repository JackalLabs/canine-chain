import { Writer, Reader } from "protobufjs/minimal";
import { Params } from "../jklmining/params";
import { SaveRequests } from "../jklmining/save_requests";
import { Miners } from "../jklmining/miners";
import { Mined } from "../jklmining/mined";
export declare const protobufPackage = "jackaldao.canine.jklmining";
/** GenesisState defines the jklmining module's genesis state. */
export interface GenesisState {
    params: Params | undefined;
    saveRequestsList: SaveRequests[];
    minersList: Miners[];
    minedList: Mined[];
    /** this line is used by starport scaffolding # genesis/proto/state */
    minedCount: number;
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
