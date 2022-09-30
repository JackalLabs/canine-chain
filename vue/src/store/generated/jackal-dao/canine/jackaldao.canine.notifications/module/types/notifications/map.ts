/* eslint-disable */
import * as Long from "long";
import { util, configure, Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "jackaldao.canine.notifications";

export interface Map {
  creator: string;
  notificationsList: Map_Notifications[];
  notificationsCount: number;
}

export interface Map_Notifications {
  id: number;
  actualmsg: string;
  sender: string;
}

const baseMap: object = { creator: "", notificationsCount: 0 };

export const Map = {
  encode(message: Map, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    for (const v of message.notificationsList) {
      Map_Notifications.encode(v!, writer.uint32(18).fork()).ldelim();
    }
    if (message.notificationsCount !== 0) {
      writer.uint32(24).uint64(message.notificationsCount);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): Map {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMap } as Map;
    message.notificationsList = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.notificationsList.push(
            Map_Notifications.decode(reader, reader.uint32())
          );
          break;
        case 3:
          message.notificationsCount = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Map {
    const message = { ...baseMap } as Map;
    message.notificationsList = [];
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (
      object.notificationsList !== undefined &&
      object.notificationsList !== null
    ) {
      for (const e of object.notificationsList) {
        message.notificationsList.push(Map_Notifications.fromJSON(e));
      }
    }
    if (
      object.notificationsCount !== undefined &&
      object.notificationsCount !== null
    ) {
      message.notificationsCount = Number(object.notificationsCount);
    } else {
      message.notificationsCount = 0;
    }
    return message;
  },

  toJSON(message: Map): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    if (message.notificationsList) {
      obj.notificationsList = message.notificationsList.map((e) =>
        e ? Map_Notifications.toJSON(e) : undefined
      );
    } else {
      obj.notificationsList = [];
    }
    message.notificationsCount !== undefined &&
      (obj.notificationsCount = message.notificationsCount);
    return obj;
  },

  fromPartial(object: DeepPartial<Map>): Map {
    const message = { ...baseMap } as Map;
    message.notificationsList = [];
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (
      object.notificationsList !== undefined &&
      object.notificationsList !== null
    ) {
      for (const e of object.notificationsList) {
        message.notificationsList.push(Map_Notifications.fromPartial(e));
      }
    }
    if (
      object.notificationsCount !== undefined &&
      object.notificationsCount !== null
    ) {
      message.notificationsCount = object.notificationsCount;
    } else {
      message.notificationsCount = 0;
    }
    return message;
  },
};

const baseMap_Notifications: object = { id: 0, actualmsg: "", sender: "" };

export const Map_Notifications = {
  encode(message: Map_Notifications, writer: Writer = Writer.create()): Writer {
    if (message.id !== 0) {
      writer.uint32(8).uint64(message.id);
    }
    if (message.actualmsg !== "") {
      writer.uint32(18).string(message.actualmsg);
    }
    if (message.sender !== "") {
      writer.uint32(26).string(message.sender);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): Map_Notifications {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMap_Notifications } as Map_Notifications;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = longToNumber(reader.uint64() as Long);
          break;
        case 2:
          message.actualmsg = reader.string();
          break;
        case 3:
          message.sender = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Map_Notifications {
    const message = { ...baseMap_Notifications } as Map_Notifications;
    if (object.id !== undefined && object.id !== null) {
      message.id = Number(object.id);
    } else {
      message.id = 0;
    }
    if (object.actualmsg !== undefined && object.actualmsg !== null) {
      message.actualmsg = String(object.actualmsg);
    } else {
      message.actualmsg = "";
    }
    if (object.sender !== undefined && object.sender !== null) {
      message.sender = String(object.sender);
    } else {
      message.sender = "";
    }
    return message;
  },

  toJSON(message: Map_Notifications): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    message.actualmsg !== undefined && (obj.actualmsg = message.actualmsg);
    message.sender !== undefined && (obj.sender = message.sender);
    return obj;
  },

  fromPartial(object: DeepPartial<Map_Notifications>): Map_Notifications {
    const message = { ...baseMap_Notifications } as Map_Notifications;
    if (object.id !== undefined && object.id !== null) {
      message.id = object.id;
    } else {
      message.id = 0;
    }
    if (object.actualmsg !== undefined && object.actualmsg !== null) {
      message.actualmsg = object.actualmsg;
    } else {
      message.actualmsg = "";
    }
    if (object.sender !== undefined && object.sender !== null) {
      message.sender = object.sender;
    } else {
      message.sender = "";
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
