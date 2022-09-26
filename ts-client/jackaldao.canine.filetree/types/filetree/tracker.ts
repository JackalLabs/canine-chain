/* eslint-disable */
import * as Long from "long";
import { util, configure, Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "jackaldao.canine.filetree";

export interface Tracker {
  trackingNumber: number;
}

const baseTracker: object = { trackingNumber: 0 };

export const Tracker = {
  encode(message: Tracker, writer: Writer = Writer.create()): Writer {
    if (message.trackingNumber !== 0) {
      writer.uint32(8).uint64(message.trackingNumber);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): Tracker {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseTracker } as Tracker;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.trackingNumber = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Tracker {
    const message = { ...baseTracker } as Tracker;
    if (object.trackingNumber !== undefined && object.trackingNumber !== null) {
      message.trackingNumber = Number(object.trackingNumber);
    } else {
      message.trackingNumber = 0;
    }
    return message;
  },

  toJSON(message: Tracker): unknown {
    const obj: any = {};
    message.trackingNumber !== undefined &&
      (obj.trackingNumber = message.trackingNumber);
    return obj;
  },

  fromPartial(object: DeepPartial<Tracker>): Tracker {
    const message = { ...baseTracker } as Tracker;
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
