/* eslint-disable */
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "jackaldao.canine.storage";

export interface Miners {
  address: string;
  ip: string;
  totalspace: string;
  burned_contracts: string;
  creator: string;
}

const baseMiners: object = {
  address: "",
  ip: "",
  totalspace: "",
  burned_contracts: "",
  creator: "",
};

export const Miners = {
  encode(message: Miners, writer: Writer = Writer.create()): Writer {
    if (message.address !== "") {
      writer.uint32(10).string(message.address);
    }
    if (message.ip !== "") {
      writer.uint32(18).string(message.ip);
    }
    if (message.totalspace !== "") {
      writer.uint32(26).string(message.totalspace);
    }
    if (message.burned_contracts !== "") {
      writer.uint32(34).string(message.burned_contracts);
    }
    if (message.creator !== "") {
      writer.uint32(42).string(message.creator);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): Miners {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMiners } as Miners;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.address = reader.string();
          break;
        case 2:
          message.ip = reader.string();
          break;
        case 3:
          message.totalspace = reader.string();
          break;
        case 4:
          message.burned_contracts = reader.string();
          break;
        case 5:
          message.creator = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Miners {
    const message = { ...baseMiners } as Miners;
    if (object.address !== undefined && object.address !== null) {
      message.address = String(object.address);
    } else {
      message.address = "";
    }
    if (object.ip !== undefined && object.ip !== null) {
      message.ip = String(object.ip);
    } else {
      message.ip = "";
    }
    if (object.totalspace !== undefined && object.totalspace !== null) {
      message.totalspace = String(object.totalspace);
    } else {
      message.totalspace = "";
    }
    if (
      object.burned_contracts !== undefined &&
      object.burned_contracts !== null
    ) {
      message.burned_contracts = String(object.burned_contracts);
    } else {
      message.burned_contracts = "";
    }
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    return message;
  },

  toJSON(message: Miners): unknown {
    const obj: any = {};
    message.address !== undefined && (obj.address = message.address);
    message.ip !== undefined && (obj.ip = message.ip);
    message.totalspace !== undefined && (obj.totalspace = message.totalspace);
    message.burned_contracts !== undefined &&
      (obj.burned_contracts = message.burned_contracts);
    message.creator !== undefined && (obj.creator = message.creator);
    return obj;
  },

  fromPartial(object: DeepPartial<Miners>): Miners {
    const message = { ...baseMiners } as Miners;
    if (object.address !== undefined && object.address !== null) {
      message.address = object.address;
    } else {
      message.address = "";
    }
    if (object.ip !== undefined && object.ip !== null) {
      message.ip = object.ip;
    } else {
      message.ip = "";
    }
    if (object.totalspace !== undefined && object.totalspace !== null) {
      message.totalspace = object.totalspace;
    } else {
      message.totalspace = "";
    }
    if (
      object.burned_contracts !== undefined &&
      object.burned_contracts !== null
    ) {
      message.burned_contracts = object.burned_contracts;
    } else {
      message.burned_contracts = "";
    }
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
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
