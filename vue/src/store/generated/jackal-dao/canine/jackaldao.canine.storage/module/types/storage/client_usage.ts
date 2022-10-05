/* eslint-disable */
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "jackaldao.canine.storage";

export interface ClientUsage {
  address: string;
  usage: string;
}

const baseClientUsage: object = { address: "", usage: "" };

export const ClientUsage = {
  encode(message: ClientUsage, writer: Writer = Writer.create()): Writer {
    if (message.address !== "") {
      writer.uint32(10).string(message.address);
    }
    if (message.usage !== "") {
      writer.uint32(18).string(message.usage);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): ClientUsage {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseClientUsage } as ClientUsage;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.address = reader.string();
          break;
        case 2:
          message.usage = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): ClientUsage {
    const message = { ...baseClientUsage } as ClientUsage;
    if (object.address !== undefined && object.address !== null) {
      message.address = String(object.address);
    } else {
      message.address = "";
    }
    if (object.usage !== undefined && object.usage !== null) {
      message.usage = String(object.usage);
    } else {
      message.usage = "";
    }
    return message;
  },

  toJSON(message: ClientUsage): unknown {
    const obj: any = {};
    message.address !== undefined && (obj.address = message.address);
    message.usage !== undefined && (obj.usage = message.usage);
    return obj;
  },

  fromPartial(object: DeepPartial<ClientUsage>): ClientUsage {
    const message = { ...baseClientUsage } as ClientUsage;
    if (object.address !== undefined && object.address !== null) {
      message.address = object.address;
    } else {
      message.address = "";
    }
    if (object.usage !== undefined && object.usage !== null) {
      message.usage = object.usage;
    } else {
      message.usage = "";
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
