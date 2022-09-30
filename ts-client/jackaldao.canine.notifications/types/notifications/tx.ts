/* eslint-disable */
import { Reader, util, configure, Writer } from "protobufjs/minimal";
import * as Long from "long";

export const protobufPackage = "jackaldao.canine.notifications";

export interface MsgCreateNotifications {
  creator: string;
  count: number;
  notification: string;
  address: string;
}

export interface MsgCreateNotificationsResponse {}

export interface MsgUpdateNotifications {
  creator: string;
  count: number;
  notification: string;
  address: string;
}

export interface MsgUpdateNotificationsResponse {}

export interface MsgDeleteNotifications {
  creator: string;
  count: number;
}

export interface MsgDeleteNotificationsResponse {}

export interface MsgSetCounter {
  creator: string;
}

export interface MsgSetCounterResponse {}

const baseMsgCreateNotifications: object = {
  creator: "",
  count: 0,
  notification: "",
  address: "",
};

export const MsgCreateNotifications = {
  encode(
    message: MsgCreateNotifications,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.count !== 0) {
      writer.uint32(16).uint64(message.count);
    }
    if (message.notification !== "") {
      writer.uint32(26).string(message.notification);
    }
    if (message.address !== "") {
      writer.uint32(34).string(message.address);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgCreateNotifications {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgCreateNotifications } as MsgCreateNotifications;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.count = longToNumber(reader.uint64() as Long);
          break;
        case 3:
          message.notification = reader.string();
          break;
        case 4:
          message.address = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgCreateNotifications {
    const message = { ...baseMsgCreateNotifications } as MsgCreateNotifications;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
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
    return message;
  },

  toJSON(message: MsgCreateNotifications): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.count !== undefined && (obj.count = message.count);
    message.notification !== undefined &&
      (obj.notification = message.notification);
    message.address !== undefined && (obj.address = message.address);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgCreateNotifications>
  ): MsgCreateNotifications {
    const message = { ...baseMsgCreateNotifications } as MsgCreateNotifications;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
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
    return message;
  },
};

const baseMsgCreateNotificationsResponse: object = {};

export const MsgCreateNotificationsResponse = {
  encode(
    _: MsgCreateNotificationsResponse,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgCreateNotificationsResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgCreateNotificationsResponse,
    } as MsgCreateNotificationsResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): MsgCreateNotificationsResponse {
    const message = {
      ...baseMsgCreateNotificationsResponse,
    } as MsgCreateNotificationsResponse;
    return message;
  },

  toJSON(_: MsgCreateNotificationsResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgCreateNotificationsResponse>
  ): MsgCreateNotificationsResponse {
    const message = {
      ...baseMsgCreateNotificationsResponse,
    } as MsgCreateNotificationsResponse;
    return message;
  },
};

const baseMsgUpdateNotifications: object = {
  creator: "",
  count: 0,
  notification: "",
  address: "",
};

export const MsgUpdateNotifications = {
  encode(
    message: MsgUpdateNotifications,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.count !== 0) {
      writer.uint32(16).uint64(message.count);
    }
    if (message.notification !== "") {
      writer.uint32(26).string(message.notification);
    }
    if (message.address !== "") {
      writer.uint32(34).string(message.address);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgUpdateNotifications {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgUpdateNotifications } as MsgUpdateNotifications;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.count = longToNumber(reader.uint64() as Long);
          break;
        case 3:
          message.notification = reader.string();
          break;
        case 4:
          message.address = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgUpdateNotifications {
    const message = { ...baseMsgUpdateNotifications } as MsgUpdateNotifications;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
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
    return message;
  },

  toJSON(message: MsgUpdateNotifications): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.count !== undefined && (obj.count = message.count);
    message.notification !== undefined &&
      (obj.notification = message.notification);
    message.address !== undefined && (obj.address = message.address);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgUpdateNotifications>
  ): MsgUpdateNotifications {
    const message = { ...baseMsgUpdateNotifications } as MsgUpdateNotifications;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
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
    return message;
  },
};

const baseMsgUpdateNotificationsResponse: object = {};

export const MsgUpdateNotificationsResponse = {
  encode(
    _: MsgUpdateNotificationsResponse,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgUpdateNotificationsResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgUpdateNotificationsResponse,
    } as MsgUpdateNotificationsResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): MsgUpdateNotificationsResponse {
    const message = {
      ...baseMsgUpdateNotificationsResponse,
    } as MsgUpdateNotificationsResponse;
    return message;
  },

  toJSON(_: MsgUpdateNotificationsResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgUpdateNotificationsResponse>
  ): MsgUpdateNotificationsResponse {
    const message = {
      ...baseMsgUpdateNotificationsResponse,
    } as MsgUpdateNotificationsResponse;
    return message;
  },
};

const baseMsgDeleteNotifications: object = { creator: "", count: 0 };

export const MsgDeleteNotifications = {
  encode(
    message: MsgDeleteNotifications,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.count !== 0) {
      writer.uint32(16).uint64(message.count);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgDeleteNotifications {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgDeleteNotifications } as MsgDeleteNotifications;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.count = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgDeleteNotifications {
    const message = { ...baseMsgDeleteNotifications } as MsgDeleteNotifications;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.count !== undefined && object.count !== null) {
      message.count = Number(object.count);
    } else {
      message.count = 0;
    }
    return message;
  },

  toJSON(message: MsgDeleteNotifications): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.count !== undefined && (obj.count = message.count);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgDeleteNotifications>
  ): MsgDeleteNotifications {
    const message = { ...baseMsgDeleteNotifications } as MsgDeleteNotifications;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.count !== undefined && object.count !== null) {
      message.count = object.count;
    } else {
      message.count = 0;
    }
    return message;
  },
};

const baseMsgDeleteNotificationsResponse: object = {};

export const MsgDeleteNotificationsResponse = {
  encode(
    _: MsgDeleteNotificationsResponse,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgDeleteNotificationsResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgDeleteNotificationsResponse,
    } as MsgDeleteNotificationsResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): MsgDeleteNotificationsResponse {
    const message = {
      ...baseMsgDeleteNotificationsResponse,
    } as MsgDeleteNotificationsResponse;
    return message;
  },

  toJSON(_: MsgDeleteNotificationsResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgDeleteNotificationsResponse>
  ): MsgDeleteNotificationsResponse {
    const message = {
      ...baseMsgDeleteNotificationsResponse,
    } as MsgDeleteNotificationsResponse;
    return message;
  },
};

const baseMsgSetCounter: object = { creator: "" };

export const MsgSetCounter = {
  encode(message: MsgSetCounter, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgSetCounter {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgSetCounter } as MsgSetCounter;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgSetCounter {
    const message = { ...baseMsgSetCounter } as MsgSetCounter;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    return message;
  },

  toJSON(message: MsgSetCounter): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgSetCounter>): MsgSetCounter {
    const message = { ...baseMsgSetCounter } as MsgSetCounter;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    return message;
  },
};

const baseMsgSetCounterResponse: object = {};

export const MsgSetCounterResponse = {
  encode(_: MsgSetCounterResponse, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgSetCounterResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgSetCounterResponse } as MsgSetCounterResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): MsgSetCounterResponse {
    const message = { ...baseMsgSetCounterResponse } as MsgSetCounterResponse;
    return message;
  },

  toJSON(_: MsgSetCounterResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(_: DeepPartial<MsgSetCounterResponse>): MsgSetCounterResponse {
    const message = { ...baseMsgSetCounterResponse } as MsgSetCounterResponse;
    return message;
  },
};

/** Msg defines the Msg service. */
export interface Msg {
  CreateNotifications(
    request: MsgCreateNotifications
  ): Promise<MsgCreateNotificationsResponse>;
  UpdateNotifications(
    request: MsgUpdateNotifications
  ): Promise<MsgUpdateNotificationsResponse>;
  DeleteNotifications(
    request: MsgDeleteNotifications
  ): Promise<MsgDeleteNotificationsResponse>;
  /** this line is used by starport scaffolding # proto/tx/rpc */
  SetCounter(request: MsgSetCounter): Promise<MsgSetCounterResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  CreateNotifications(
    request: MsgCreateNotifications
  ): Promise<MsgCreateNotificationsResponse> {
    const data = MsgCreateNotifications.encode(request).finish();
    const promise = this.rpc.request(
      "jackaldao.canine.notifications.Msg",
      "CreateNotifications",
      data
    );
    return promise.then((data) =>
      MsgCreateNotificationsResponse.decode(new Reader(data))
    );
  }

  UpdateNotifications(
    request: MsgUpdateNotifications
  ): Promise<MsgUpdateNotificationsResponse> {
    const data = MsgUpdateNotifications.encode(request).finish();
    const promise = this.rpc.request(
      "jackaldao.canine.notifications.Msg",
      "UpdateNotifications",
      data
    );
    return promise.then((data) =>
      MsgUpdateNotificationsResponse.decode(new Reader(data))
    );
  }

  DeleteNotifications(
    request: MsgDeleteNotifications
  ): Promise<MsgDeleteNotificationsResponse> {
    const data = MsgDeleteNotifications.encode(request).finish();
    const promise = this.rpc.request(
      "jackaldao.canine.notifications.Msg",
      "DeleteNotifications",
      data
    );
    return promise.then((data) =>
      MsgDeleteNotificationsResponse.decode(new Reader(data))
    );
  }

  SetCounter(request: MsgSetCounter): Promise<MsgSetCounterResponse> {
    const data = MsgSetCounter.encode(request).finish();
    const promise = this.rpc.request(
      "jackaldao.canine.notifications.Msg",
      "SetCounter",
      data
    );
    return promise.then((data) =>
      MsgSetCounterResponse.decode(new Reader(data))
    );
  }
}

interface Rpc {
  request(
    service: string,
    method: string,
    data: Uint8Array
  ): Promise<Uint8Array>;
}

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
