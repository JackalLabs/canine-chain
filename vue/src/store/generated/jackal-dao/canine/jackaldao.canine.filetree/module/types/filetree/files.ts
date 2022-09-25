/* eslint-disable */
import * as Long from "long";
import { util, configure, Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "jackaldao.canine.filetree";

export interface Files {
  address: string;
  contents: string;
  owner: string;
  viewingAccess: string;
  editAccess: string;
  trackingNumber: number;
}

const baseFiles: object = {
  address: "",
  contents: "",
  owner: "",
  viewingAccess: "",
  editAccess: "",
  trackingNumber: 0,
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
    if (message.trackingNumber !== 0) {
      writer.uint32(48).uint64(message.trackingNumber);
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
        case 6:
          message.trackingNumber = longToNumber(reader.uint64() as Long);
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
    if (object.trackingNumber !== undefined && object.trackingNumber !== null) {
      message.trackingNumber = Number(object.trackingNumber);
    } else {
      message.trackingNumber = 0;
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
    message.trackingNumber !== undefined &&
      (obj.trackingNumber = message.trackingNumber);
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
    if (object.trackingNumber !== undefined && object.trackingNumber !== null) {
      message.trackingNumber = object.trackingNumber;
    } else {
      message.trackingNumber = 0;
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
