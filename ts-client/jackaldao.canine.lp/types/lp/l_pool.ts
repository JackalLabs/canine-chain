/* eslint-disable */
import * as Long from "long";
import { util, configure, Writer, Reader } from "protobufjs/minimal";
import { Coin } from "../cosmos/base/v1beta1/coin";

export const protobufPackage = "jackaldao.canine.lp";

export interface LPool {
  index: string;
  name: string;
  coins: Coin[];
  aMMId: number;
  /** sdk.Dec in string representation */
  swapFeeMulti: string;
  minLockDuration: number;
  /** sdk.Dec in string representation */
  penaltyMulti: string;
  lptokenDenom: string;
  LPTokenBalance: string;
}

const baseLPool: object = {
  index: "",
  name: "",
  aMMId: 0,
  swapFeeMulti: "",
  minLockDuration: 0,
  penaltyMulti: "",
  lptokenDenom: "",
  LPTokenBalance: "",
};

export const LPool = {
  encode(message: LPool, writer: Writer = Writer.create()): Writer {
    if (message.index !== "") {
      writer.uint32(10).string(message.index);
    }
    if (message.name !== "") {
      writer.uint32(18).string(message.name);
    }
    for (const v of message.coins) {
      Coin.encode(v!, writer.uint32(26).fork()).ldelim();
    }
    if (message.aMMId !== 0) {
      writer.uint32(32).uint32(message.aMMId);
    }
    if (message.swapFeeMulti !== "") {
      writer.uint32(42).string(message.swapFeeMulti);
    }
    if (message.minLockDuration !== 0) {
      writer.uint32(48).int64(message.minLockDuration);
    }
    if (message.penaltyMulti !== "") {
      writer.uint32(58).string(message.penaltyMulti);
    }
    if (message.lptokenDenom !== "") {
      writer.uint32(66).string(message.lptokenDenom);
    }
    if (message.LPTokenBalance !== "") {
      writer.uint32(74).string(message.LPTokenBalance);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): LPool {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseLPool } as LPool;
    message.coins = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.index = reader.string();
          break;
        case 2:
          message.name = reader.string();
          break;
        case 3:
          message.coins.push(Coin.decode(reader, reader.uint32()));
          break;
        case 4:
          message.aMMId = reader.uint32();
          break;
        case 5:
          message.swapFeeMulti = reader.string();
          break;
        case 6:
          message.minLockDuration = longToNumber(reader.int64() as Long);
          break;
        case 7:
          message.penaltyMulti = reader.string();
          break;
        case 8:
          message.lptokenDenom = reader.string();
          break;
        case 9:
          message.LPTokenBalance = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): LPool {
    const message = { ...baseLPool } as LPool;
    message.coins = [];
    if (object.index !== undefined && object.index !== null) {
      message.index = String(object.index);
    } else {
      message.index = "";
    }
    if (object.name !== undefined && object.name !== null) {
      message.name = String(object.name);
    } else {
      message.name = "";
    }
    if (object.coins !== undefined && object.coins !== null) {
      for (const e of object.coins) {
        message.coins.push(Coin.fromJSON(e));
      }
    }
    if (object.aMMId !== undefined && object.aMMId !== null) {
      message.aMMId = Number(object.aMMId);
    } else {
      message.aMMId = 0;
    }
    if (object.swapFeeMulti !== undefined && object.swapFeeMulti !== null) {
      message.swapFeeMulti = String(object.swapFeeMulti);
    } else {
      message.swapFeeMulti = "";
    }
    if (
      object.minLockDuration !== undefined &&
      object.minLockDuration !== null
    ) {
      message.minLockDuration = Number(object.minLockDuration);
    } else {
      message.minLockDuration = 0;
    }
    if (object.penaltyMulti !== undefined && object.penaltyMulti !== null) {
      message.penaltyMulti = String(object.penaltyMulti);
    } else {
      message.penaltyMulti = "";
    }
    if (object.lptokenDenom !== undefined && object.lptokenDenom !== null) {
      message.lptokenDenom = String(object.lptokenDenom);
    } else {
      message.lptokenDenom = "";
    }
    if (object.LPTokenBalance !== undefined && object.LPTokenBalance !== null) {
      message.LPTokenBalance = String(object.LPTokenBalance);
    } else {
      message.LPTokenBalance = "";
    }
    return message;
  },

  toJSON(message: LPool): unknown {
    const obj: any = {};
    message.index !== undefined && (obj.index = message.index);
    message.name !== undefined && (obj.name = message.name);
    if (message.coins) {
      obj.coins = message.coins.map((e) => (e ? Coin.toJSON(e) : undefined));
    } else {
      obj.coins = [];
    }
    message.aMMId !== undefined && (obj.aMMId = message.aMMId);
    message.swapFeeMulti !== undefined &&
      (obj.swapFeeMulti = message.swapFeeMulti);
    message.minLockDuration !== undefined &&
      (obj.minLockDuration = message.minLockDuration);
    message.penaltyMulti !== undefined &&
      (obj.penaltyMulti = message.penaltyMulti);
    message.lptokenDenom !== undefined &&
      (obj.lptokenDenom = message.lptokenDenom);
    message.LPTokenBalance !== undefined &&
      (obj.LPTokenBalance = message.LPTokenBalance);
    return obj;
  },

  fromPartial(object: DeepPartial<LPool>): LPool {
    const message = { ...baseLPool } as LPool;
    message.coins = [];
    if (object.index !== undefined && object.index !== null) {
      message.index = object.index;
    } else {
      message.index = "";
    }
    if (object.name !== undefined && object.name !== null) {
      message.name = object.name;
    } else {
      message.name = "";
    }
    if (object.coins !== undefined && object.coins !== null) {
      for (const e of object.coins) {
        message.coins.push(Coin.fromPartial(e));
      }
    }
    if (object.aMMId !== undefined && object.aMMId !== null) {
      message.aMMId = object.aMMId;
    } else {
      message.aMMId = 0;
    }
    if (object.swapFeeMulti !== undefined && object.swapFeeMulti !== null) {
      message.swapFeeMulti = object.swapFeeMulti;
    } else {
      message.swapFeeMulti = "";
    }
    if (
      object.minLockDuration !== undefined &&
      object.minLockDuration !== null
    ) {
      message.minLockDuration = object.minLockDuration;
    } else {
      message.minLockDuration = 0;
    }
    if (object.penaltyMulti !== undefined && object.penaltyMulti !== null) {
      message.penaltyMulti = object.penaltyMulti;
    } else {
      message.penaltyMulti = "";
    }
    if (object.lptokenDenom !== undefined && object.lptokenDenom !== null) {
      message.lptokenDenom = object.lptokenDenom;
    } else {
      message.lptokenDenom = "";
    }
    if (object.LPTokenBalance !== undefined && object.LPTokenBalance !== null) {
      message.LPTokenBalance = object.LPTokenBalance;
    } else {
      message.LPTokenBalance = "";
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
