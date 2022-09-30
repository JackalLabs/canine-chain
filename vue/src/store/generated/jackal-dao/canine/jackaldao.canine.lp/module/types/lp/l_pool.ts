/* eslint-disable */
import * as Long from "long";
import { util, configure, Writer, Reader } from "protobufjs/minimal";
import { Coin } from "../cosmos/base/v1beta1/coin";

export const protobufPackage = "jackaldao.canine.lp";

export interface LPool {
  index: string;
  name: string;
  coins: Coin[];
  aMM_Id: number;
  /** sdk.Dec in string representation */
  swap_fee_multi: string;
  min_lock_duration: number;
  /** sdk.Dec in string representation */
  penalty_multi: string;
  lptoken_denom: string;
  LPTokenBalance: string;
}

const baseLPool: object = {
  index: "",
  name: "",
  aMM_Id: 0,
  swap_fee_multi: "",
  min_lock_duration: 0,
  penalty_multi: "",
  lptoken_denom: "",
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
    if (message.aMM_Id !== 0) {
      writer.uint32(32).uint32(message.aMM_Id);
    }
    if (message.swap_fee_multi !== "") {
      writer.uint32(42).string(message.swap_fee_multi);
    }
    if (message.min_lock_duration !== 0) {
      writer.uint32(48).int64(message.min_lock_duration);
    }
    if (message.penalty_multi !== "") {
      writer.uint32(58).string(message.penalty_multi);
    }
    if (message.lptoken_denom !== "") {
      writer.uint32(66).string(message.lptoken_denom);
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
          message.aMM_Id = reader.uint32();
          break;
        case 5:
          message.swap_fee_multi = reader.string();
          break;
        case 6:
          message.min_lock_duration = longToNumber(reader.int64() as Long);
          break;
        case 7:
          message.penalty_multi = reader.string();
          break;
        case 8:
          message.lptoken_denom = reader.string();
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
    if (object.aMM_Id !== undefined && object.aMM_Id !== null) {
      message.aMM_Id = Number(object.aMM_Id);
    } else {
      message.aMM_Id = 0;
    }
    if (object.swap_fee_multi !== undefined && object.swap_fee_multi !== null) {
      message.swap_fee_multi = String(object.swap_fee_multi);
    } else {
      message.swap_fee_multi = "";
    }
    if (
      object.min_lock_duration !== undefined &&
      object.min_lock_duration !== null
    ) {
      message.min_lock_duration = Number(object.min_lock_duration);
    } else {
      message.min_lock_duration = 0;
    }
    if (object.penalty_multi !== undefined && object.penalty_multi !== null) {
      message.penalty_multi = String(object.penalty_multi);
    } else {
      message.penalty_multi = "";
    }
    if (object.lptoken_denom !== undefined && object.lptoken_denom !== null) {
      message.lptoken_denom = String(object.lptoken_denom);
    } else {
      message.lptoken_denom = "";
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
    message.aMM_Id !== undefined && (obj.aMM_Id = message.aMM_Id);
    message.swap_fee_multi !== undefined &&
      (obj.swap_fee_multi = message.swap_fee_multi);
    message.min_lock_duration !== undefined &&
      (obj.min_lock_duration = message.min_lock_duration);
    message.penalty_multi !== undefined &&
      (obj.penalty_multi = message.penalty_multi);
    message.lptoken_denom !== undefined &&
      (obj.lptoken_denom = message.lptoken_denom);
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
    if (object.aMM_Id !== undefined && object.aMM_Id !== null) {
      message.aMM_Id = object.aMM_Id;
    } else {
      message.aMM_Id = 0;
    }
    if (object.swap_fee_multi !== undefined && object.swap_fee_multi !== null) {
      message.swap_fee_multi = object.swap_fee_multi;
    } else {
      message.swap_fee_multi = "";
    }
    if (
      object.min_lock_duration !== undefined &&
      object.min_lock_duration !== null
    ) {
      message.min_lock_duration = object.min_lock_duration;
    } else {
      message.min_lock_duration = 0;
    }
    if (object.penalty_multi !== undefined && object.penalty_multi !== null) {
      message.penalty_multi = object.penalty_multi;
    } else {
      message.penalty_multi = "";
    }
    if (object.lptoken_denom !== undefined && object.lptoken_denom !== null) {
      message.lptoken_denom = object.lptoken_denom;
    } else {
      message.lptoken_denom = "";
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
