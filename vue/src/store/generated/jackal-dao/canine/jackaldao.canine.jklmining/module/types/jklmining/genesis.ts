/* eslint-disable */
import * as Long from "long";
import { util, configure, Writer, Reader } from "protobufjs/minimal";
import { Params } from "../jklmining/params";
import { SaveRequests } from "../jklmining/save_requests";
import { Miners } from "../jklmining/miners";
import { Mined } from "../jklmining/mined";

export const protobufPackage = "jackaldao.canine.jklmining";

/** GenesisState defines the jklmining module's genesis state. */
export interface GenesisState {
  params: Params | undefined;
  saveRequestsList: SaveRequests[];
  minersList: Miners[];
  minedList: Mined[];
  /** this line is used by starport scaffolding # genesis/proto/state */
  minedCount: number;
}

const baseGenesisState: object = { minedCount: 0 };

export const GenesisState = {
  encode(message: GenesisState, writer: Writer = Writer.create()): Writer {
    if (message.params !== undefined) {
      Params.encode(message.params, writer.uint32(10).fork()).ldelim();
    }
    for (const v of message.saveRequestsList) {
      SaveRequests.encode(v!, writer.uint32(18).fork()).ldelim();
    }
    for (const v of message.minersList) {
      Miners.encode(v!, writer.uint32(26).fork()).ldelim();
    }
    for (const v of message.minedList) {
      Mined.encode(v!, writer.uint32(34).fork()).ldelim();
    }
    if (message.minedCount !== 0) {
      writer.uint32(40).uint64(message.minedCount);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): GenesisState {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseGenesisState } as GenesisState;
    message.saveRequestsList = [];
    message.minersList = [];
    message.minedList = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.params = Params.decode(reader, reader.uint32());
          break;
        case 2:
          message.saveRequestsList.push(
            SaveRequests.decode(reader, reader.uint32())
          );
          break;
        case 3:
          message.minersList.push(Miners.decode(reader, reader.uint32()));
          break;
        case 4:
          message.minedList.push(Mined.decode(reader, reader.uint32()));
          break;
        case 5:
          message.minedCount = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): GenesisState {
    const message = { ...baseGenesisState } as GenesisState;
    message.saveRequestsList = [];
    message.minersList = [];
    message.minedList = [];
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromJSON(object.params);
    } else {
      message.params = undefined;
    }
    if (
      object.saveRequestsList !== undefined &&
      object.saveRequestsList !== null
    ) {
      for (const e of object.saveRequestsList) {
        message.saveRequestsList.push(SaveRequests.fromJSON(e));
      }
    }
    if (object.minersList !== undefined && object.minersList !== null) {
      for (const e of object.minersList) {
        message.minersList.push(Miners.fromJSON(e));
      }
    }
    if (object.minedList !== undefined && object.minedList !== null) {
      for (const e of object.minedList) {
        message.minedList.push(Mined.fromJSON(e));
      }
    }
    if (object.minedCount !== undefined && object.minedCount !== null) {
      message.minedCount = Number(object.minedCount);
    } else {
      message.minedCount = 0;
    }
    return message;
  },

  toJSON(message: GenesisState): unknown {
    const obj: any = {};
    message.params !== undefined &&
      (obj.params = message.params ? Params.toJSON(message.params) : undefined);
    if (message.saveRequestsList) {
      obj.saveRequestsList = message.saveRequestsList.map((e) =>
        e ? SaveRequests.toJSON(e) : undefined
      );
    } else {
      obj.saveRequestsList = [];
    }
    if (message.minersList) {
      obj.minersList = message.minersList.map((e) =>
        e ? Miners.toJSON(e) : undefined
      );
    } else {
      obj.minersList = [];
    }
    if (message.minedList) {
      obj.minedList = message.minedList.map((e) =>
        e ? Mined.toJSON(e) : undefined
      );
    } else {
      obj.minedList = [];
    }
    message.minedCount !== undefined && (obj.minedCount = message.minedCount);
    return obj;
  },

  fromPartial(object: DeepPartial<GenesisState>): GenesisState {
    const message = { ...baseGenesisState } as GenesisState;
    message.saveRequestsList = [];
    message.minersList = [];
    message.minedList = [];
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromPartial(object.params);
    } else {
      message.params = undefined;
    }
    if (
      object.saveRequestsList !== undefined &&
      object.saveRequestsList !== null
    ) {
      for (const e of object.saveRequestsList) {
        message.saveRequestsList.push(SaveRequests.fromPartial(e));
      }
    }
    if (object.minersList !== undefined && object.minersList !== null) {
      for (const e of object.minersList) {
        message.minersList.push(Miners.fromPartial(e));
      }
    }
    if (object.minedList !== undefined && object.minedList !== null) {
      for (const e of object.minedList) {
        message.minedList.push(Mined.fromPartial(e));
      }
    }
    if (object.minedCount !== undefined && object.minedCount !== null) {
      message.minedCount = object.minedCount;
    } else {
      message.minedCount = 0;
    }
    return message;
  },
};

declare var self: any | undefined;
declare var window: any | undefined;
var globalThis: any = (() => {
  if (typeof globalThis !== "undefined") return globalThis;
  if (typeof self !== "undefined") return self;
  if (typeof window !== "undefined") return window;
  if (typeof global !== "undefined") return global;
  throw "Unable to locate global object";
})();

type Builtin = Date | Function | Uint8Array | string | number | undefined;
export type DeepPartial<T> = T extends Builtin
  ? T
  : T extends Array<infer U>
  ? Array<DeepPartial<U>>
  : T extends ReadonlyArray<infer U>
  ? ReadonlyArray<DeepPartial<U>>
  : T extends {}
  ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

function longToNumber(long: Long): number {
  if (long.gt(Number.MAX_SAFE_INTEGER)) {
    throw new globalThis.Error("Value is larger than Number.MAX_SAFE_INTEGER");
  }
  return long.toNumber();
}

if (util.Long !== Long) {
  util.Long = Long as any;
  configure();
}
