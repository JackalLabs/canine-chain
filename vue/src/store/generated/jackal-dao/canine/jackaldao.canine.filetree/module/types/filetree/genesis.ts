/* eslint-disable */
import { Params } from "../filetree/params";
import { Files } from "../filetree/files";
import { Pubkey } from "../filetree/pubkey";
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "jackaldao.canine.filetree";

/** GenesisState defines the filetree module's genesis state. */
export interface GenesisState {
  params: Params | undefined;
  filesList: Files[];
  /** this line is used by starport scaffolding # genesis/proto/state */
  pubkeyList: Pubkey[];
}

const baseGenesisState: object = {};

export const GenesisState = {
  encode(message: GenesisState, writer: Writer = Writer.create()): Writer {
    if (message.params !== undefined) {
      Params.encode(message.params, writer.uint32(10).fork()).ldelim();
    }
    for (const v of message.filesList) {
      Files.encode(v!, writer.uint32(18).fork()).ldelim();
    }
    for (const v of message.pubkeyList) {
      Pubkey.encode(v!, writer.uint32(26).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): GenesisState {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseGenesisState } as GenesisState;
    message.filesList = [];
    message.pubkeyList = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.params = Params.decode(reader, reader.uint32());
          break;
        case 2:
          message.filesList.push(Files.decode(reader, reader.uint32()));
          break;
        case 3:
          message.pubkeyList.push(Pubkey.decode(reader, reader.uint32()));
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
    message.filesList = [];
    message.pubkeyList = [];
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromJSON(object.params);
    } else {
      message.params = undefined;
    }
    if (object.filesList !== undefined && object.filesList !== null) {
      for (const e of object.filesList) {
        message.filesList.push(Files.fromJSON(e));
      }
    }
    if (object.pubkeyList !== undefined && object.pubkeyList !== null) {
      for (const e of object.pubkeyList) {
        message.pubkeyList.push(Pubkey.fromJSON(e));
      }
    }
    return message;
  },

  toJSON(message: GenesisState): unknown {
    const obj: any = {};
    message.params !== undefined &&
      (obj.params = message.params ? Params.toJSON(message.params) : undefined);
    if (message.filesList) {
      obj.filesList = message.filesList.map((e) =>
        e ? Files.toJSON(e) : undefined
      );
    } else {
      obj.filesList = [];
    }
    if (message.pubkeyList) {
      obj.pubkeyList = message.pubkeyList.map((e) =>
        e ? Pubkey.toJSON(e) : undefined
      );
    } else {
      obj.pubkeyList = [];
    }
    return obj;
  },

  fromPartial(object: DeepPartial<GenesisState>): GenesisState {
    const message = { ...baseGenesisState } as GenesisState;
    message.filesList = [];
    message.pubkeyList = [];
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromPartial(object.params);
    } else {
      message.params = undefined;
    }
    if (object.filesList !== undefined && object.filesList !== null) {
      for (const e of object.filesList) {
        message.filesList.push(Files.fromPartial(e));
      }
    }
    if (object.pubkeyList !== undefined && object.pubkeyList !== null) {
      for (const e of object.pubkeyList) {
        message.pubkeyList.push(Pubkey.fromPartial(e));
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
