/* eslint-disable */
import * as Long from "long";
import { util, configure, Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "jackaldao.canine.notifications";

export interface Notifications {
  count: number;
  notification: string;
  address: string;
  sender: string;
  hashPath: string;
  hashPathOwner: string;
}

const baseNotifications: object = {
  count: 0,
  notification: "",
  address: "",
  sender: "",
  hashPath: "",
  hashPathOwner: "",
};

export const Notifications = {
  encode(message: Notifications, writer: Writer = Writer.create()): Writer {
    if (message.count !== 0) {
      writer.uint32(8).uint64(message.count);
    }
    if (message.notification !== "") {
      writer.uint32(18).string(message.notification);
    }
    if (message.address !== "") {
      writer.uint32(26).string(message.address);
    }
    if (message.sender !== "") {
      writer.uint32(34).string(message.sender);
    }
    if (message.hashPath !== "") {
      writer.uint32(42).string(message.hashPath);
    }
    if (message.hashPathOwner !== "") {
      writer.uint32(50).string(message.hashPathOwner);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): Notifications {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseNotifications } as Notifications;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.count = longToNumber(reader.uint64() as Long);
          break;
        case 2:
          message.notification = reader.string();
          break;
        case 3:
          message.address = reader.string();
          break;
        case 4:
          message.sender = reader.string();
          break;
        case 5:
          message.hashPath = reader.string();
          break;
        case 6:
          message.hashPathOwner = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Notifications {
    const message = { ...baseNotifications } as Notifications;
    if (object.count !== undefined && object.count !== null) {
      message.count = Number(object.count);
    } else {
      message.count = 0;
    }
    if (object.notification !== undefined && object.notification !== null) {
      message.notification = String(object.notification);
    } else {
      message.notification = "";
    }
    if (object.address !== undefined && object.address !== null) {
      message.address = String(object.address);
    } else {
      message.address = "";
    }
    if (object.sender !== undefined && object.sender !== null) {
      message.sender = String(object.sender);
    } else {
      message.sender = "";
    }
    if (object.hashPath !== undefined && object.hashPath !== null) {
      message.hashPath = String(object.hashPath);
    } else {
      message.hashPath = "";
    }
    if (object.hashPathOwner !== undefined && object.hashPathOwner !== null) {
      message.hashPathOwner = String(object.hashPathOwner);
    } else {
      message.hashPathOwner = "";
    }
    return message;
  },

  toJSON(message: Notifications): unknown {
    const obj: any = {};
    message.count !== undefined && (obj.count = message.count);
    message.notification !== undefined &&
      (obj.notification = message.notification);
    message.address !== undefined && (obj.address = message.address);
    message.sender !== undefined && (obj.sender = message.sender);
    message.hashPath !== undefined && (obj.hashPath = message.hashPath);
    message.hashPathOwner !== undefined &&
      (obj.hashPathOwner = message.hashPathOwner);
    return obj;
  },

  fromPartial(object: DeepPartial<Notifications>): Notifications {
    const message = { ...baseNotifications } as Notifications;
    if (object.count !== undefined && object.count !== null) {
      message.count = object.count;
    } else {
      message.count = 0;
    }
    if (object.notification !== undefined && object.notification !== null) {
      message.notification = object.notification;
    } else {
      message.notification = "";
    }
    if (object.address !== undefined && object.address !== null) {
      message.address = object.address;
    } else {
      message.address = "";
    }
    if (object.sender !== undefined && object.sender !== null) {
      message.sender = object.sender;
    } else {
      message.sender = "";
    }
    if (object.hashPath !== undefined && object.hashPath !== null) {
      message.hashPath = object.hashPath;
    } else {
      message.hashPath = "";
    }
    if (object.hashPathOwner !== undefined && object.hashPathOwner !== null) {
      message.hashPathOwner = object.hashPathOwner;
    } else {
      message.hashPathOwner = "";
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
