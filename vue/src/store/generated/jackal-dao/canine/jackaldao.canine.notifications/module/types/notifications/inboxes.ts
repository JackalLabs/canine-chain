/* eslint-disable */
import * as Long from "long";
import { util, configure, Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "jackaldao.canine.notifications";

export interface Inboxes {
  index: string;
  creator: string;
  notifications: { [key: number]: Inboxes_Notifications };
  notificationsCount: number;
}

export interface Inboxes_NotificationsEntry {
  key: number;
  value: Inboxes_Notifications | undefined;
}

export interface Inboxes_Notifications {
  id: number;
  notification: string;
  sender: string;
}

const baseInboxes: object = { index: "", creator: "", notificationsCount: 0 };

export const Inboxes = {
  encode(message: Inboxes, writer: Writer = Writer.create()): Writer {
    if (message.index !== "") {
      writer.uint32(10).string(message.index);
    }
    if (message.creator !== "") {
      writer.uint32(18).string(message.creator);
    }
    Object.entries(message.notifications).forEach(([key, value]) => {
      Inboxes_NotificationsEntry.encode(
        { key: key as any, value },
        writer.uint32(26).fork()
      ).ldelim();
    });
    if (message.notificationsCount !== 0) {
      writer.uint32(32).uint64(message.notificationsCount);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): Inboxes {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseInboxes } as Inboxes;
    message.notifications = {};
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.index = reader.string();
          break;
        case 2:
          message.creator = reader.string();
          break;
        case 3:
          const entry3 = Inboxes_NotificationsEntry.decode(
            reader,
            reader.uint32()
          );
          if (entry3.value !== undefined) {
            message.notifications[entry3.key] = entry3.value;
          }
          break;
        case 4:
          message.notificationsCount = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Inboxes {
    const message = { ...baseInboxes } as Inboxes;
    message.notifications = {};
    if (object.index !== undefined && object.index !== null) {
      message.index = String(object.index);
    } else {
      message.index = "";
    }
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.notifications !== undefined && object.notifications !== null) {
      Object.entries(object.notifications).forEach(([key, value]) => {
        message.notifications[Number(key)] = Inboxes_Notifications.fromJSON(
          value
        );
      });
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

  toJSON(message: Inboxes): unknown {
    const obj: any = {};
    message.index !== undefined && (obj.index = message.index);
    message.creator !== undefined && (obj.creator = message.creator);
    obj.notifications = {};
    if (message.notifications) {
      Object.entries(message.notifications).forEach(([k, v]) => {
        obj.notifications[k] = Inboxes_Notifications.toJSON(v);
      });
    }
    message.notificationsCount !== undefined &&
      (obj.notificationsCount = message.notificationsCount);
    return obj;
  },

  fromPartial(object: DeepPartial<Inboxes>): Inboxes {
    const message = { ...baseInboxes } as Inboxes;
    message.notifications = {};
    if (object.index !== undefined && object.index !== null) {
      message.index = object.index;
    } else {
      message.index = "";
    }
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.notifications !== undefined && object.notifications !== null) {
      Object.entries(object.notifications).forEach(([key, value]) => {
        if (value !== undefined) {
          message.notifications[
            Number(key)
          ] = Inboxes_Notifications.fromPartial(value);
        }
      });
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

const baseInboxes_NotificationsEntry: object = { key: 0 };

export const Inboxes_NotificationsEntry = {
  encode(
    message: Inboxes_NotificationsEntry,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.key !== 0) {
      writer.uint32(8).uint64(message.key);
    }
    if (message.value !== undefined) {
      Inboxes_Notifications.encode(
        message.value,
        writer.uint32(18).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): Inboxes_NotificationsEntry {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseInboxes_NotificationsEntry,
    } as Inboxes_NotificationsEntry;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.key = longToNumber(reader.uint64() as Long);
          break;
        case 2:
          message.value = Inboxes_Notifications.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Inboxes_NotificationsEntry {
    const message = {
      ...baseInboxes_NotificationsEntry,
    } as Inboxes_NotificationsEntry;
    if (object.key !== undefined && object.key !== null) {
      message.key = Number(object.key);
    } else {
      message.key = 0;
    }
    if (object.value !== undefined && object.value !== null) {
      message.value = Inboxes_Notifications.fromJSON(object.value);
    } else {
      message.value = undefined;
    }
    return message;
  },

  toJSON(message: Inboxes_NotificationsEntry): unknown {
    const obj: any = {};
    message.key !== undefined && (obj.key = message.key);
    message.value !== undefined &&
      (obj.value = message.value
        ? Inboxes_Notifications.toJSON(message.value)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<Inboxes_NotificationsEntry>
  ): Inboxes_NotificationsEntry {
    const message = {
      ...baseInboxes_NotificationsEntry,
    } as Inboxes_NotificationsEntry;
    if (object.key !== undefined && object.key !== null) {
      message.key = object.key;
    } else {
      message.key = 0;
    }
    if (object.value !== undefined && object.value !== null) {
      message.value = Inboxes_Notifications.fromPartial(object.value);
    } else {
      message.value = undefined;
    }
    return message;
  },
};

const baseInboxes_Notifications: object = {
  id: 0,
  notification: "",
  sender: "",
};

export const Inboxes_Notifications = {
  encode(
    message: Inboxes_Notifications,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.id !== 0) {
      writer.uint32(8).uint64(message.id);
    }
    if (message.notification !== "") {
      writer.uint32(18).string(message.notification);
    }
    if (message.sender !== "") {
      writer.uint32(26).string(message.sender);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): Inboxes_Notifications {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseInboxes_Notifications } as Inboxes_Notifications;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = longToNumber(reader.uint64() as Long);
          break;
        case 2:
          message.notification = reader.string();
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

  fromJSON(object: any): Inboxes_Notifications {
    const message = { ...baseInboxes_Notifications } as Inboxes_Notifications;
    if (object.id !== undefined && object.id !== null) {
      message.id = Number(object.id);
    } else {
      message.id = 0;
    }
    if (object.notification !== undefined && object.notification !== null) {
      message.notification = String(object.notification);
    } else {
      message.notification = "";
    }
    if (object.sender !== undefined && object.sender !== null) {
      message.sender = String(object.sender);
    } else {
      message.sender = "";
    }
    return message;
  },

  toJSON(message: Inboxes_Notifications): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    message.notification !== undefined &&
      (obj.notification = message.notification);
    message.sender !== undefined && (obj.sender = message.sender);
    return obj;
  },

  fromPartial(
    object: DeepPartial<Inboxes_Notifications>
  ): Inboxes_Notifications {
    const message = { ...baseInboxes_Notifications } as Inboxes_Notifications;
    if (object.id !== undefined && object.id !== null) {
      message.id = object.id;
    } else {
      message.id = 0;
    }
    if (object.notification !== undefined && object.notification !== null) {
      message.notification = object.notification;
    } else {
      message.notification = "";
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
