/* eslint-disable */
import * as Long from "long";
import { util, configure, Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "jackaldao.canine.rns";

export interface Names {
  name: string;
  expires: number;
  value: string;
  data: string;
  subdomains: Names[];
  tld: string;
  locked: number;
}

const baseNames: object = {
  name: "",
  expires: 0,
  value: "",
  data: "",
  tld: "",
  locked: 0,
};

export const Names = {
  encode(message: Names, writer: Writer = Writer.create()): Writer {
    if (message.name !== "") {
      writer.uint32(10).string(message.name);
    }
    if (message.expires !== 0) {
      writer.uint32(16).int64(message.expires);
    }
    if (message.value !== "") {
      writer.uint32(26).string(message.value);
    }
    if (message.data !== "") {
      writer.uint32(34).string(message.data);
    }
    for (const v of message.subdomains) {
      Names.encode(v!, writer.uint32(42).fork()).ldelim();
    }
    if (message.tld !== "") {
      writer.uint32(50).string(message.tld);
    }
    if (message.locked !== 0) {
      writer.uint32(56).int64(message.locked);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): Names {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseNames } as Names;
    message.subdomains = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.name = reader.string();
          break;
        case 2:
          message.expires = longToNumber(reader.int64() as Long);
          break;
        case 3:
          message.value = reader.string();
          break;
        case 4:
          message.data = reader.string();
          break;
        case 5:
          message.subdomains.push(Names.decode(reader, reader.uint32()));
          break;
        case 6:
          message.tld = reader.string();
          break;
        case 7:
          message.locked = longToNumber(reader.int64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Names {
    const message = { ...baseNames } as Names;
    message.subdomains = [];
    if (object.name !== undefined && object.name !== null) {
      message.name = String(object.name);
    } else {
      message.name = "";
    }
    if (object.expires !== undefined && object.expires !== null) {
      message.expires = Number(object.expires);
    } else {
      message.expires = 0;
    }
    if (object.value !== undefined && object.value !== null) {
      message.value = String(object.value);
    } else {
      message.value = "";
    }
    if (object.data !== undefined && object.data !== null) {
      message.data = String(object.data);
    } else {
      message.data = "";
    }
    if (object.subdomains !== undefined && object.subdomains !== null) {
      for (const e of object.subdomains) {
        message.subdomains.push(Names.fromJSON(e));
      }
    }
    if (object.tld !== undefined && object.tld !== null) {
      message.tld = String(object.tld);
    } else {
      message.tld = "";
    }
    if (object.locked !== undefined && object.locked !== null) {
      message.locked = Number(object.locked);
    } else {
      message.locked = 0;
    }
    return message;
  },

  toJSON(message: Names): unknown {
    const obj: any = {};
    message.name !== undefined && (obj.name = message.name);
    message.expires !== undefined && (obj.expires = message.expires);
    message.value !== undefined && (obj.value = message.value);
    message.data !== undefined && (obj.data = message.data);
    if (message.subdomains) {
      obj.subdomains = message.subdomains.map((e) =>
        e ? Names.toJSON(e) : undefined
      );
    } else {
      obj.subdomains = [];
    }
    message.tld !== undefined && (obj.tld = message.tld);
    message.locked !== undefined && (obj.locked = message.locked);
    return obj;
  },

  fromPartial(object: DeepPartial<Names>): Names {
    const message = { ...baseNames } as Names;
    message.subdomains = [];
    if (object.name !== undefined && object.name !== null) {
      message.name = object.name;
    } else {
      message.name = "";
    }
    if (object.expires !== undefined && object.expires !== null) {
      message.expires = object.expires;
    } else {
      message.expires = 0;
    }
    if (object.value !== undefined && object.value !== null) {
      message.value = object.value;
    } else {
      message.value = "";
    }
    if (object.data !== undefined && object.data !== null) {
      message.data = object.data;
    } else {
      message.data = "";
    }
    if (object.subdomains !== undefined && object.subdomains !== null) {
      for (const e of object.subdomains) {
        message.subdomains.push(Names.fromPartial(e));
      }
    }
    if (object.tld !== undefined && object.tld !== null) {
      message.tld = object.tld;
    } else {
      message.tld = "";
    }
    if (object.locked !== undefined && object.locked !== null) {
      message.locked = object.locked;
    } else {
      message.locked = 0;
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
