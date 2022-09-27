/* eslint-disable */
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "jackaldao.canine.rns";

export interface Init {
  address: string;
  complete: boolean;
}

const baseInit: object = { address: "", complete: false };

export const Init = {
  encode(message: Init, writer: Writer = Writer.create()): Writer {
    if (message.address !== "") {
      writer.uint32(10).string(message.address);
    }
    if (message.complete === true) {
      writer.uint32(16).bool(message.complete);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): Init {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseInit } as Init;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.address = reader.string();
          break;
        case 2:
          message.complete = reader.bool();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Init {
    const message = { ...baseInit } as Init;
    if (object.address !== undefined && object.address !== null) {
      message.address = String(object.address);
    } else {
      message.address = "";
    }
    if (object.complete !== undefined && object.complete !== null) {
      message.complete = Boolean(object.complete);
    } else {
      message.complete = false;
    }
    return message;
  },

  toJSON(message: Init): unknown {
    const obj: any = {};
    message.address !== undefined && (obj.address = message.address);
    message.complete !== undefined && (obj.complete = message.complete);
    return obj;
  },

  fromPartial(object: DeepPartial<Init>): Init {
    const message = { ...baseInit } as Init;
    if (object.address !== undefined && object.address !== null) {
      message.address = object.address;
    } else {
      message.address = "";
    }
    if (object.complete !== undefined && object.complete !== null) {
      message.complete = object.complete;
    } else {
      message.complete = false;
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
