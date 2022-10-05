/* eslint-disable */
import { Reader, util, configure, Writer } from "protobufjs/minimal";
import * as Long from "long";
import { DecCoin } from "../cosmos/base/v1beta1/coin";

export const protobufPackage = "jackaldao.canine.lp";

export interface MsgCreateLPool {
  creator: string;
  /**
   * Creator needs to deposit coins to create pool.
   * Input format:
   * "{amount0}{denomination},...,{amountN}{denominationN}"
   */
  coins: DecCoin[];
  /** AMM model id used to balance the pool. */
  amm_Id: number;
  /**
   * Swap fee charged to swap transaction (swap x swap_fee_multi).
   * Is converted to type sdk.Dec
   */
  swap_fee_multi: string;
  /**
   * Liquidity pool token (LPToken) lock duration in seconds.
   * Liquidity provider's LPToken is locked when they contribute to a pool.
   * Penalty is applied when LPToken is burned before the lock duration is over.
   */
  min_lock_duration: number;
  /**
   * Penalty multiplier applied to LPToken when provider wishes to burn LPToken
   * before lock duration is over (LPToken x penalty_multi) -> pool_tokens.
   * Is converted to type sdk.Dec
   */
  penalty_multi: string;
}

export interface MsgCreateLPoolResponse {
  /** Pool id */
  id: string;
}

export interface MsgDepositLPool {
  creator: string;
  poolName: string;
  /**
   * Input format:
   * "{amount0}{denomination},...,{amountN}{denominationN}"
   */
  coins: DecCoin[];
  /** The contributor can choose lock duration */
  lockDuration: number;
}

export interface MsgDepositLPoolResponse {
  /** Amount of shares given to pool contributor. */
  shares: number;
}

export interface MsgWithdrawLPool {
  creator: string;
  poolName: string;
  shares: number;
}

export interface MsgWithdrawLPoolResponse {}

export interface MsgSwap {
  creator: string;
  poolName: string;
  coin: DecCoin | undefined;
}

export interface MsgSwapResponse {}

const baseMsgCreateLPool: object = {
  creator: "",
  amm_Id: 0,
  swap_fee_multi: "",
  min_lock_duration: 0,
  penalty_multi: "",
};

export const MsgCreateLPool = {
  encode(message: MsgCreateLPool, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    for (const v of message.coins) {
      DecCoin.encode(v!, writer.uint32(18).fork()).ldelim();
    }
    if (message.amm_Id !== 0) {
      writer.uint32(24).uint32(message.amm_Id);
    }
    if (message.swap_fee_multi !== "") {
      writer.uint32(34).string(message.swap_fee_multi);
    }
    if (message.min_lock_duration !== 0) {
      writer.uint32(40).int64(message.min_lock_duration);
    }
    if (message.penalty_multi !== "") {
      writer.uint32(50).string(message.penalty_multi);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgCreateLPool {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgCreateLPool } as MsgCreateLPool;
    message.coins = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.coins.push(DecCoin.decode(reader, reader.uint32()));
          break;
        case 3:
          message.amm_Id = reader.uint32();
          break;
        case 4:
          message.swap_fee_multi = reader.string();
          break;
        case 5:
          message.min_lock_duration = longToNumber(reader.int64() as Long);
          break;
        case 6:
          message.penalty_multi = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgCreateLPool {
    const message = { ...baseMsgCreateLPool } as MsgCreateLPool;
    message.coins = [];
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.coins !== undefined && object.coins !== null) {
      for (const e of object.coins) {
        message.coins.push(DecCoin.fromJSON(e));
      }
    }
    if (object.amm_Id !== undefined && object.amm_Id !== null) {
      message.amm_Id = Number(object.amm_Id);
    } else {
      message.amm_Id = 0;
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
    return message;
  },

  toJSON(message: MsgCreateLPool): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    if (message.coins) {
      obj.coins = message.coins.map((e) => (e ? DecCoin.toJSON(e) : undefined));
    } else {
      obj.coins = [];
    }
    message.amm_Id !== undefined && (obj.amm_Id = message.amm_Id);
    message.swap_fee_multi !== undefined &&
      (obj.swap_fee_multi = message.swap_fee_multi);
    message.min_lock_duration !== undefined &&
      (obj.min_lock_duration = message.min_lock_duration);
    message.penalty_multi !== undefined &&
      (obj.penalty_multi = message.penalty_multi);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgCreateLPool>): MsgCreateLPool {
    const message = { ...baseMsgCreateLPool } as MsgCreateLPool;
    message.coins = [];
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.coins !== undefined && object.coins !== null) {
      for (const e of object.coins) {
        message.coins.push(DecCoin.fromPartial(e));
      }
    }
    if (object.amm_Id !== undefined && object.amm_Id !== null) {
      message.amm_Id = object.amm_Id;
    } else {
      message.amm_Id = 0;
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
    return message;
  },
};

const baseMsgCreateLPoolResponse: object = { id: "" };

export const MsgCreateLPoolResponse = {
  encode(
    message: MsgCreateLPoolResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.id !== "") {
      writer.uint32(10).string(message.id);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgCreateLPoolResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgCreateLPoolResponse } as MsgCreateLPoolResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgCreateLPoolResponse {
    const message = { ...baseMsgCreateLPoolResponse } as MsgCreateLPoolResponse;
    if (object.id !== undefined && object.id !== null) {
      message.id = String(object.id);
    } else {
      message.id = "";
    }
    return message;
  },

  toJSON(message: MsgCreateLPoolResponse): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgCreateLPoolResponse>
  ): MsgCreateLPoolResponse {
    const message = { ...baseMsgCreateLPoolResponse } as MsgCreateLPoolResponse;
    if (object.id !== undefined && object.id !== null) {
      message.id = object.id;
    } else {
      message.id = "";
    }
    return message;
  },
};

const baseMsgDepositLPool: object = {
  creator: "",
  poolName: "",
  lockDuration: 0,
};

export const MsgDepositLPool = {
  encode(message: MsgDepositLPool, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.poolName !== "") {
      writer.uint32(18).string(message.poolName);
    }
    for (const v of message.coins) {
      DecCoin.encode(v!, writer.uint32(26).fork()).ldelim();
    }
    if (message.lockDuration !== 0) {
      writer.uint32(32).int64(message.lockDuration);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgDepositLPool {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgDepositLPool } as MsgDepositLPool;
    message.coins = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.poolName = reader.string();
          break;
        case 3:
          message.coins.push(DecCoin.decode(reader, reader.uint32()));
          break;
        case 4:
          message.lockDuration = longToNumber(reader.int64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgDepositLPool {
    const message = { ...baseMsgDepositLPool } as MsgDepositLPool;
    message.coins = [];
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.poolName !== undefined && object.poolName !== null) {
      message.poolName = String(object.poolName);
    } else {
      message.poolName = "";
    }
    if (object.coins !== undefined && object.coins !== null) {
      for (const e of object.coins) {
        message.coins.push(DecCoin.fromJSON(e));
      }
    }
    if (object.lockDuration !== undefined && object.lockDuration !== null) {
      message.lockDuration = Number(object.lockDuration);
    } else {
      message.lockDuration = 0;
    }
    return message;
  },

  toJSON(message: MsgDepositLPool): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.poolName !== undefined && (obj.poolName = message.poolName);
    if (message.coins) {
      obj.coins = message.coins.map((e) => (e ? DecCoin.toJSON(e) : undefined));
    } else {
      obj.coins = [];
    }
    message.lockDuration !== undefined &&
      (obj.lockDuration = message.lockDuration);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgDepositLPool>): MsgDepositLPool {
    const message = { ...baseMsgDepositLPool } as MsgDepositLPool;
    message.coins = [];
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.poolName !== undefined && object.poolName !== null) {
      message.poolName = object.poolName;
    } else {
      message.poolName = "";
    }
    if (object.coins !== undefined && object.coins !== null) {
      for (const e of object.coins) {
        message.coins.push(DecCoin.fromPartial(e));
      }
    }
    if (object.lockDuration !== undefined && object.lockDuration !== null) {
      message.lockDuration = object.lockDuration;
    } else {
      message.lockDuration = 0;
    }
    return message;
  },
};

const baseMsgDepositLPoolResponse: object = { shares: 0 };

export const MsgDepositLPoolResponse = {
  encode(
    message: MsgDepositLPoolResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.shares !== 0) {
      writer.uint32(8).uint64(message.shares);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgDepositLPoolResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgDepositLPoolResponse,
    } as MsgDepositLPoolResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.shares = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgDepositLPoolResponse {
    const message = {
      ...baseMsgDepositLPoolResponse,
    } as MsgDepositLPoolResponse;
    if (object.shares !== undefined && object.shares !== null) {
      message.shares = Number(object.shares);
    } else {
      message.shares = 0;
    }
    return message;
  },

  toJSON(message: MsgDepositLPoolResponse): unknown {
    const obj: any = {};
    message.shares !== undefined && (obj.shares = message.shares);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgDepositLPoolResponse>
  ): MsgDepositLPoolResponse {
    const message = {
      ...baseMsgDepositLPoolResponse,
    } as MsgDepositLPoolResponse;
    if (object.shares !== undefined && object.shares !== null) {
      message.shares = object.shares;
    } else {
      message.shares = 0;
    }
    return message;
  },
};

const baseMsgWithdrawLPool: object = { creator: "", poolName: "", shares: 0 };

export const MsgWithdrawLPool = {
  encode(message: MsgWithdrawLPool, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.poolName !== "") {
      writer.uint32(18).string(message.poolName);
    }
    if (message.shares !== 0) {
      writer.uint32(24).int64(message.shares);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgWithdrawLPool {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgWithdrawLPool } as MsgWithdrawLPool;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.poolName = reader.string();
          break;
        case 3:
          message.shares = longToNumber(reader.int64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgWithdrawLPool {
    const message = { ...baseMsgWithdrawLPool } as MsgWithdrawLPool;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.poolName !== undefined && object.poolName !== null) {
      message.poolName = String(object.poolName);
    } else {
      message.poolName = "";
    }
    if (object.shares !== undefined && object.shares !== null) {
      message.shares = Number(object.shares);
    } else {
      message.shares = 0;
    }
    return message;
  },

  toJSON(message: MsgWithdrawLPool): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.poolName !== undefined && (obj.poolName = message.poolName);
    message.shares !== undefined && (obj.shares = message.shares);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgWithdrawLPool>): MsgWithdrawLPool {
    const message = { ...baseMsgWithdrawLPool } as MsgWithdrawLPool;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.poolName !== undefined && object.poolName !== null) {
      message.poolName = object.poolName;
    } else {
      message.poolName = "";
    }
    if (object.shares !== undefined && object.shares !== null) {
      message.shares = object.shares;
    } else {
      message.shares = 0;
    }
    return message;
  },
};

const baseMsgWithdrawLPoolResponse: object = {};

export const MsgWithdrawLPoolResponse = {
  encode(
    _: MsgWithdrawLPoolResponse,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgWithdrawLPoolResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgWithdrawLPoolResponse,
    } as MsgWithdrawLPoolResponse;
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

  fromJSON(_: any): MsgWithdrawLPoolResponse {
    const message = {
      ...baseMsgWithdrawLPoolResponse,
    } as MsgWithdrawLPoolResponse;
    return message;
  },

  toJSON(_: MsgWithdrawLPoolResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgWithdrawLPoolResponse>
  ): MsgWithdrawLPoolResponse {
    const message = {
      ...baseMsgWithdrawLPoolResponse,
    } as MsgWithdrawLPoolResponse;
    return message;
  },
};

const baseMsgSwap: object = { creator: "", poolName: "" };

export const MsgSwap = {
  encode(message: MsgSwap, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.poolName !== "") {
      writer.uint32(18).string(message.poolName);
    }
    if (message.coin !== undefined) {
      DecCoin.encode(message.coin, writer.uint32(26).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgSwap {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgSwap } as MsgSwap;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.poolName = reader.string();
          break;
        case 3:
          message.coin = DecCoin.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgSwap {
    const message = { ...baseMsgSwap } as MsgSwap;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.poolName !== undefined && object.poolName !== null) {
      message.poolName = String(object.poolName);
    } else {
      message.poolName = "";
    }
    if (object.coin !== undefined && object.coin !== null) {
      message.coin = DecCoin.fromJSON(object.coin);
    } else {
      message.coin = undefined;
    }
    return message;
  },

  toJSON(message: MsgSwap): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.poolName !== undefined && (obj.poolName = message.poolName);
    message.coin !== undefined &&
      (obj.coin = message.coin ? DecCoin.toJSON(message.coin) : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgSwap>): MsgSwap {
    const message = { ...baseMsgSwap } as MsgSwap;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.poolName !== undefined && object.poolName !== null) {
      message.poolName = object.poolName;
    } else {
      message.poolName = "";
    }
    if (object.coin !== undefined && object.coin !== null) {
      message.coin = DecCoin.fromPartial(object.coin);
    } else {
      message.coin = undefined;
    }
    return message;
  },
};

const baseMsgSwapResponse: object = {};

export const MsgSwapResponse = {
  encode(_: MsgSwapResponse, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgSwapResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgSwapResponse } as MsgSwapResponse;
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

  fromJSON(_: any): MsgSwapResponse {
    const message = { ...baseMsgSwapResponse } as MsgSwapResponse;
    return message;
  },

  toJSON(_: MsgSwapResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(_: DeepPartial<MsgSwapResponse>): MsgSwapResponse {
    const message = { ...baseMsgSwapResponse } as MsgSwapResponse;
    return message;
  },
};

/** Msg defines the Msg service. */
export interface Msg {
  CreateLPool(request: MsgCreateLPool): Promise<MsgCreateLPoolResponse>;
  DepositLPool(request: MsgDepositLPool): Promise<MsgDepositLPoolResponse>;
  WithdrawLPool(request: MsgWithdrawLPool): Promise<MsgWithdrawLPoolResponse>;
  /** this line is used by starport scaffolding # proto/tx/rpc */
  Swap(request: MsgSwap): Promise<MsgSwapResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  CreateLPool(request: MsgCreateLPool): Promise<MsgCreateLPoolResponse> {
    const data = MsgCreateLPool.encode(request).finish();
    const promise = this.rpc.request(
      "jackaldao.canine.lp.Msg",
      "CreateLPool",
      data
    );
    return promise.then((data) =>
      MsgCreateLPoolResponse.decode(new Reader(data))
    );
  }

  DepositLPool(request: MsgDepositLPool): Promise<MsgDepositLPoolResponse> {
    const data = MsgDepositLPool.encode(request).finish();
    const promise = this.rpc.request(
      "jackaldao.canine.lp.Msg",
      "DepositLPool",
      data
    );
    return promise.then((data) =>
      MsgDepositLPoolResponse.decode(new Reader(data))
    );
  }

  WithdrawLPool(request: MsgWithdrawLPool): Promise<MsgWithdrawLPoolResponse> {
    const data = MsgWithdrawLPool.encode(request).finish();
    const promise = this.rpc.request(
      "jackaldao.canine.lp.Msg",
      "WithdrawLPool",
      data
    );
    return promise.then((data) =>
      MsgWithdrawLPoolResponse.decode(new Reader(data))
    );
  }

  Swap(request: MsgSwap): Promise<MsgSwapResponse> {
    const data = MsgSwap.encode(request).finish();
    const promise = this.rpc.request("jackaldao.canine.lp.Msg", "Swap", data);
    return promise.then((data) => MsgSwapResponse.decode(new Reader(data)));
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
