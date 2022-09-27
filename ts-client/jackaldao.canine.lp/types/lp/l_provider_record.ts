/* eslint-disable */
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "jackaldao.canine.lp";

/**
 * LProviderRecord is a record of a liquidity provider depositing to a pool.
 * It is used to enforce withdraw panelty and calculate rewards collected.
 * This record is created only once when provider contributes to a pool and
 * only updated after witdrawal or deposit.
 * It is deleted when the provider burns all of the liquidity pool token.
 * This is stored at KVStore with
 * 	{LProviderRecordKeyPrefix}{poolName}{provider} key.
 */
export interface LProviderRecord {
  /** Provider is the account address of the provider. */
  provider: string;
  /** A pool that the provider contributed to. */
  poolName: string;
  /**
   * Burning LP token is locked for certain duration the after provider
   * deposits to the pool. Unlock time is updated every succeeding deposits.
   * The provider can burn their LP token during lock time but has to take
   * certain amount of panelty.
   * Unlock time is blocktime + lockDuration at time of contribution.
   */
  unlockTime: string;
  lockDuration: string;
}

const baseLProviderRecord: object = {
  provider: "",
  poolName: "",
  unlockTime: "",
  lockDuration: "",
};

export const LProviderRecord = {
  encode(message: LProviderRecord, writer: Writer = Writer.create()): Writer {
    if (message.provider !== "") {
      writer.uint32(10).string(message.provider);
    }
    if (message.poolName !== "") {
      writer.uint32(18).string(message.poolName);
    }
    if (message.unlockTime !== "") {
      writer.uint32(26).string(message.unlockTime);
    }
    if (message.lockDuration !== "") {
      writer.uint32(34).string(message.lockDuration);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): LProviderRecord {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseLProviderRecord } as LProviderRecord;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.provider = reader.string();
          break;
        case 2:
          message.poolName = reader.string();
          break;
        case 3:
          message.unlockTime = reader.string();
          break;
        case 4:
          message.lockDuration = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): LProviderRecord {
    const message = { ...baseLProviderRecord } as LProviderRecord;
    if (object.provider !== undefined && object.provider !== null) {
      message.provider = String(object.provider);
    } else {
      message.provider = "";
    }
    if (object.poolName !== undefined && object.poolName !== null) {
      message.poolName = String(object.poolName);
    } else {
      message.poolName = "";
    }
    if (object.unlockTime !== undefined && object.unlockTime !== null) {
      message.unlockTime = String(object.unlockTime);
    } else {
      message.unlockTime = "";
    }
    if (object.lockDuration !== undefined && object.lockDuration !== null) {
      message.lockDuration = String(object.lockDuration);
    } else {
      message.lockDuration = "";
    }
    return message;
  },

  toJSON(message: LProviderRecord): unknown {
    const obj: any = {};
    message.provider !== undefined && (obj.provider = message.provider);
    message.poolName !== undefined && (obj.poolName = message.poolName);
    message.unlockTime !== undefined && (obj.unlockTime = message.unlockTime);
    message.lockDuration !== undefined &&
      (obj.lockDuration = message.lockDuration);
    return obj;
  },

  fromPartial(object: DeepPartial<LProviderRecord>): LProviderRecord {
    const message = { ...baseLProviderRecord } as LProviderRecord;
    if (object.provider !== undefined && object.provider !== null) {
      message.provider = object.provider;
    } else {
      message.provider = "";
    }
    if (object.poolName !== undefined && object.poolName !== null) {
      message.poolName = object.poolName;
    } else {
      message.poolName = "";
    }
    if (object.unlockTime !== undefined && object.unlockTime !== null) {
      message.unlockTime = object.unlockTime;
    } else {
      message.unlockTime = "";
    }
    if (object.lockDuration !== undefined && object.lockDuration !== null) {
      message.lockDuration = object.lockDuration;
    } else {
      message.lockDuration = "";
    }
    return message;
  },
};

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
