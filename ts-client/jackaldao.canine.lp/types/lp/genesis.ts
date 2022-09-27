/* eslint-disable */
import { Params } from "../lp/params";
import { LPool } from "../lp/l_pool";
import { LProviderRecord } from "../lp/l_provider_record";
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "jackaldao.canine.lp";

/** GenesisState defines the lp module's genesis state. */
export interface GenesisState {
  params: Params | undefined;
  lPoolList: LPool[];
  /** this line is used by starport scaffolding # genesis/proto/state */
  lProviderRecordList: LProviderRecord[];
}

const baseGenesisState: object = {};

export const GenesisState = {
  encode(message: GenesisState, writer: Writer = Writer.create()): Writer {
    if (message.params !== undefined) {
      Params.encode(message.params, writer.uint32(10).fork()).ldelim();
    }
    for (const v of message.lPoolList) {
      LPool.encode(v!, writer.uint32(18).fork()).ldelim();
    }
    for (const v of message.lProviderRecordList) {
      LProviderRecord.encode(v!, writer.uint32(26).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): GenesisState {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseGenesisState } as GenesisState;
    message.lPoolList = [];
    message.lProviderRecordList = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.params = Params.decode(reader, reader.uint32());
          break;
        case 2:
          message.lPoolList.push(LPool.decode(reader, reader.uint32()));
          break;
        case 3:
          message.lProviderRecordList.push(
            LProviderRecord.decode(reader, reader.uint32())
          );
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
    message.lPoolList = [];
    message.lProviderRecordList = [];
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromJSON(object.params);
    } else {
      message.params = undefined;
    }
    if (object.lPoolList !== undefined && object.lPoolList !== null) {
      for (const e of object.lPoolList) {
        message.lPoolList.push(LPool.fromJSON(e));
      }
    }
    if (
      object.lProviderRecordList !== undefined &&
      object.lProviderRecordList !== null
    ) {
      for (const e of object.lProviderRecordList) {
        message.lProviderRecordList.push(LProviderRecord.fromJSON(e));
      }
    }
    return message;
  },

  toJSON(message: GenesisState): unknown {
    const obj: any = {};
    message.params !== undefined &&
      (obj.params = message.params ? Params.toJSON(message.params) : undefined);
    if (message.lPoolList) {
      obj.lPoolList = message.lPoolList.map((e) =>
        e ? LPool.toJSON(e) : undefined
      );
    } else {
      obj.lPoolList = [];
    }
    if (message.lProviderRecordList) {
      obj.lProviderRecordList = message.lProviderRecordList.map((e) =>
        e ? LProviderRecord.toJSON(e) : undefined
      );
    } else {
      obj.lProviderRecordList = [];
    }
    return obj;
  },

  fromPartial(object: DeepPartial<GenesisState>): GenesisState {
    const message = { ...baseGenesisState } as GenesisState;
    message.lPoolList = [];
    message.lProviderRecordList = [];
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromPartial(object.params);
    } else {
      message.params = undefined;
    }
    if (object.lPoolList !== undefined && object.lPoolList !== null) {
      for (const e of object.lPoolList) {
        message.lPoolList.push(LPool.fromPartial(e));
      }
    }
    if (
      object.lProviderRecordList !== undefined &&
      object.lProviderRecordList !== null
    ) {
      for (const e of object.lProviderRecordList) {
        message.lProviderRecordList.push(LProviderRecord.fromPartial(e));
      }
    }
    return message;
  },
};

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
