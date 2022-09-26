/* eslint-disable */
import { Params } from "../dsig/params";
import { UserUploads } from "../dsig/user_uploads";
import { Form } from "../dsig/form";
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "jackaldao.canine.dsig";

/** GenesisState defines the dsig module's genesis state. */
export interface GenesisState {
  params: Params | undefined;
  userUploadsList: UserUploads[];
  /** this line is used by starport scaffolding # genesis/proto/state */
  formList: Form[];
}

const baseGenesisState: object = {};

export const GenesisState = {
  encode(message: GenesisState, writer: Writer = Writer.create()): Writer {
    if (message.params !== undefined) {
      Params.encode(message.params, writer.uint32(10).fork()).ldelim();
    }
    for (const v of message.userUploadsList) {
      UserUploads.encode(v!, writer.uint32(18).fork()).ldelim();
    }
    for (const v of message.formList) {
      Form.encode(v!, writer.uint32(26).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): GenesisState {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseGenesisState } as GenesisState;
    message.userUploadsList = [];
    message.formList = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.params = Params.decode(reader, reader.uint32());
          break;
        case 2:
          message.userUploadsList.push(
            UserUploads.decode(reader, reader.uint32())
          );
          break;
        case 3:
          message.formList.push(Form.decode(reader, reader.uint32()));
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
    message.userUploadsList = [];
    message.formList = [];
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromJSON(object.params);
    } else {
      message.params = undefined;
    }
    if (
      object.userUploadsList !== undefined &&
      object.userUploadsList !== null
    ) {
      for (const e of object.userUploadsList) {
        message.userUploadsList.push(UserUploads.fromJSON(e));
      }
    }
    if (object.formList !== undefined && object.formList !== null) {
      for (const e of object.formList) {
        message.formList.push(Form.fromJSON(e));
      }
    }
    return message;
  },

  toJSON(message: GenesisState): unknown {
    const obj: any = {};
    message.params !== undefined &&
      (obj.params = message.params ? Params.toJSON(message.params) : undefined);
    if (message.userUploadsList) {
      obj.userUploadsList = message.userUploadsList.map((e) =>
        e ? UserUploads.toJSON(e) : undefined
      );
    } else {
      obj.userUploadsList = [];
    }
    if (message.formList) {
      obj.formList = message.formList.map((e) =>
        e ? Form.toJSON(e) : undefined
      );
    } else {
      obj.formList = [];
    }
    return obj;
  },

  fromPartial(object: DeepPartial<GenesisState>): GenesisState {
    const message = { ...baseGenesisState } as GenesisState;
    message.userUploadsList = [];
    message.formList = [];
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromPartial(object.params);
    } else {
      message.params = undefined;
    }
    if (
      object.userUploadsList !== undefined &&
      object.userUploadsList !== null
    ) {
      for (const e of object.userUploadsList) {
        message.userUploadsList.push(UserUploads.fromPartial(e));
      }
    }
    if (object.formList !== undefined && object.formList !== null) {
      for (const e of object.formList) {
        message.formList.push(Form.fromPartial(e));
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
