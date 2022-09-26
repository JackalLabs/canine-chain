/* eslint-disable */
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "jackaldao.canine.storage";

export interface PayBlocks {
  blockid: string;
  bytes: string;
  blocktype: string;
  blocknum: string;
}

const basePayBlocks: object = {
  blockid: "",
  bytes: "",
  blocktype: "",
  blocknum: "",
};

export const PayBlocks = {
  encode(message: PayBlocks, writer: Writer = Writer.create()): Writer {
    if (message.blockid !== "") {
      writer.uint32(10).string(message.blockid);
    }
    if (message.bytes !== "") {
      writer.uint32(18).string(message.bytes);
    }
    if (message.blocktype !== "") {
      writer.uint32(26).string(message.blocktype);
    }
    if (message.blocknum !== "") {
      writer.uint32(34).string(message.blocknum);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): PayBlocks {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...basePayBlocks } as PayBlocks;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.blockid = reader.string();
          break;
        case 2:
          message.bytes = reader.string();
          break;
        case 3:
          message.blocktype = reader.string();
          break;
        case 4:
          message.blocknum = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): PayBlocks {
    const message = { ...basePayBlocks } as PayBlocks;
    if (object.blockid !== undefined && object.blockid !== null) {
      message.blockid = String(object.blockid);
    } else {
      message.blockid = "";
    }
    if (object.bytes !== undefined && object.bytes !== null) {
      message.bytes = String(object.bytes);
    } else {
      message.bytes = "";
    }
    if (object.blocktype !== undefined && object.blocktype !== null) {
      message.blocktype = String(object.blocktype);
    } else {
      message.blocktype = "";
    }
    if (object.blocknum !== undefined && object.blocknum !== null) {
      message.blocknum = String(object.blocknum);
    } else {
      message.blocknum = "";
    }
    return message;
  },

  toJSON(message: PayBlocks): unknown {
    const obj: any = {};
    message.blockid !== undefined && (obj.blockid = message.blockid);
    message.bytes !== undefined && (obj.bytes = message.bytes);
    message.blocktype !== undefined && (obj.blocktype = message.blocktype);
    message.blocknum !== undefined && (obj.blocknum = message.blocknum);
    return obj;
  },

  fromPartial(object: DeepPartial<PayBlocks>): PayBlocks {
    const message = { ...basePayBlocks } as PayBlocks;
    if (object.blockid !== undefined && object.blockid !== null) {
      message.blockid = object.blockid;
    } else {
      message.blockid = "";
    }
    if (object.bytes !== undefined && object.bytes !== null) {
      message.bytes = object.bytes;
    } else {
      message.bytes = "";
    }
    if (object.blocktype !== undefined && object.blocktype !== null) {
      message.blocktype = object.blocktype;
    } else {
      message.blocktype = "";
    }
    if (object.blocknum !== undefined && object.blocknum !== null) {
      message.blocknum = object.blocknum;
    } else {
      message.blocknum = "";
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
