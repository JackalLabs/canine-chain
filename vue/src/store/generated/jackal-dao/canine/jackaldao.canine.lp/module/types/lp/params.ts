/* eslint-disable */
import * as Long from "long";
import { util, configure, Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "jackaldao.canine.lp";

/** Params defines the parameters for the module. */
export interface Params {
  MinInitPoolDeposit: number;
  MaxPoolDenomCount: number;
  LPTokenUnit: number;
}

const baseParams: object = {
  MinInitPoolDeposit: 0,
  MaxPoolDenomCount: 0,
  LPTokenUnit: 0,
};

export const Params = {
  encode(message: Params, writer: Writer = Writer.create()): Writer {
    if (message.MinInitPoolDeposit !== 0) {
      writer.uint32(8).uint64(message.MinInitPoolDeposit);
    }
    if (message.MaxPoolDenomCount !== 0) {
      writer.uint32(16).uint32(message.MaxPoolDenomCount);
    }
    if (message.LPTokenUnit !== 0) {
      writer.uint32(24).uint32(message.LPTokenUnit);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): Params {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseParams } as Params;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.MinInitPoolDeposit = longToNumber(reader.uint64() as Long);
          break;
        case 2:
          message.MaxPoolDenomCount = reader.uint32();
          break;
        case 3:
          message.LPTokenUnit = reader.uint32();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Params {
    const message = { ...baseParams } as Params;
    if (
      object.MinInitPoolDeposit !== undefined &&
      object.MinInitPoolDeposit !== null
    ) {
      message.MinInitPoolDeposit = Number(object.MinInitPoolDeposit);
    } else {
      message.MinInitPoolDeposit = 0;
    }
    if (
      object.MaxPoolDenomCount !== undefined &&
      object.MaxPoolDenomCount !== null
    ) {
      message.MaxPoolDenomCount = Number(object.MaxPoolDenomCount);
    } else {
      message.MaxPoolDenomCount = 0;
    }
    if (object.LPTokenUnit !== undefined && object.LPTokenUnit !== null) {
      message.LPTokenUnit = Number(object.LPTokenUnit);
    } else {
      message.LPTokenUnit = 0;
    }
    return message;
  },

  toJSON(message: Params): unknown {
    const obj: any = {};
    message.MinInitPoolDeposit !== undefined &&
      (obj.MinInitPoolDeposit = message.MinInitPoolDeposit);
    message.MaxPoolDenomCount !== undefined &&
      (obj.MaxPoolDenomCount = message.MaxPoolDenomCount);
    message.LPTokenUnit !== undefined &&
      (obj.LPTokenUnit = message.LPTokenUnit);
    return obj;
  },

  fromPartial(object: DeepPartial<Params>): Params {
    const message = { ...baseParams } as Params;
    if (
      object.MinInitPoolDeposit !== undefined &&
      object.MinInitPoolDeposit !== null
    ) {
      message.MinInitPoolDeposit = object.MinInitPoolDeposit;
    } else {
      message.MinInitPoolDeposit = 0;
    }
    if (
      object.MaxPoolDenomCount !== undefined &&
      object.MaxPoolDenomCount !== null
    ) {
      message.MaxPoolDenomCount = object.MaxPoolDenomCount;
    } else {
      message.MaxPoolDenomCount = 0;
    }
    if (object.LPTokenUnit !== undefined && object.LPTokenUnit !== null) {
      message.LPTokenUnit = object.LPTokenUnit;
    } else {
      message.LPTokenUnit = 0;
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
