/* eslint-disable */
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "jackaldao.canine.filetree";

export interface Files {
  address: string;
  contents: string;
  owner: string;
  viewingAccess: string;
  editAccess: string;
}

const baseFiles: object = {
  address: "",
  contents: "",
  owner: "",
  viewingAccess: "",
  editAccess: "",
};

export const Files = {
  encode(message: Files, writer: Writer = Writer.create()): Writer {
    if (message.address !== "") {
      writer.uint32(10).string(message.address);
    }
    if (message.contents !== "") {
      writer.uint32(18).string(message.contents);
    }
    if (message.owner !== "") {
      writer.uint32(26).string(message.owner);
    }
    if (message.viewingAccess !== "") {
      writer.uint32(34).string(message.viewingAccess);
    }
    if (message.editAccess !== "") {
      writer.uint32(42).string(message.editAccess);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): Files {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseFiles } as Files;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.address = reader.string();
          break;
        case 2:
          message.contents = reader.string();
          break;
        case 3:
          message.owner = reader.string();
          break;
        case 4:
          message.viewingAccess = reader.string();
          break;
        case 5:
          message.editAccess = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Files {
    const message = { ...baseFiles } as Files;
    if (object.address !== undefined && object.address !== null) {
      message.address = String(object.address);
    } else {
      message.address = "";
    }
    if (object.contents !== undefined && object.contents !== null) {
      message.contents = String(object.contents);
    } else {
      message.contents = "";
    }
    if (object.owner !== undefined && object.owner !== null) {
      message.owner = String(object.owner);
    } else {
      message.owner = "";
    }
    if (object.viewingAccess !== undefined && object.viewingAccess !== null) {
      message.viewingAccess = String(object.viewingAccess);
    } else {
      message.viewingAccess = "";
    }
    if (object.editAccess !== undefined && object.editAccess !== null) {
      message.editAccess = String(object.editAccess);
    } else {
      message.editAccess = "";
    }
    return message;
  },

  toJSON(message: Files): unknown {
    const obj: any = {};
    message.address !== undefined && (obj.address = message.address);
    message.contents !== undefined && (obj.contents = message.contents);
    message.owner !== undefined && (obj.owner = message.owner);
    message.viewingAccess !== undefined &&
      (obj.viewingAccess = message.viewingAccess);
    message.editAccess !== undefined && (obj.editAccess = message.editAccess);
    return obj;
  },

  fromPartial(object: DeepPartial<Files>): Files {
    const message = { ...baseFiles } as Files;
    if (object.address !== undefined && object.address !== null) {
      message.address = object.address;
    } else {
      message.address = "";
    }
    if (object.contents !== undefined && object.contents !== null) {
      message.contents = object.contents;
    } else {
      message.contents = "";
    }
    if (object.owner !== undefined && object.owner !== null) {
      message.owner = object.owner;
    } else {
      message.owner = "";
    }
    if (object.viewingAccess !== undefined && object.viewingAccess !== null) {
      message.viewingAccess = object.viewingAccess;
    } else {
      message.viewingAccess = "";
    }
    if (object.editAccess !== undefined && object.editAccess !== null) {
      message.editAccess = object.editAccess;
    } else {
      message.editAccess = "";
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
