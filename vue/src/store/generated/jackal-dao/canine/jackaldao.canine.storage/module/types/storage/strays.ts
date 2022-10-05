/* eslint-disable */
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "jackaldao.canine.storage";

export interface Strays {
  cid: string;
  fid: string;
  signee: string;
  filesize: string;
  merkle: string;
}

const baseStrays: object = {
  cid: "",
  fid: "",
  signee: "",
  filesize: "",
  merkle: "",
};

export const Strays = {
  encode(message: Strays, writer: Writer = Writer.create()): Writer {
    if (message.cid !== "") {
      writer.uint32(10).string(message.cid);
    }
    if (message.fid !== "") {
      writer.uint32(18).string(message.fid);
    }
    if (message.signee !== "") {
      writer.uint32(26).string(message.signee);
    }
    if (message.filesize !== "") {
      writer.uint32(34).string(message.filesize);
    }
    if (message.merkle !== "") {
      writer.uint32(42).string(message.merkle);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): Strays {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseStrays } as Strays;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.cid = reader.string();
          break;
        case 2:
          message.fid = reader.string();
          break;
        case 3:
          message.signee = reader.string();
          break;
        case 4:
          message.filesize = reader.string();
          break;
        case 5:
          message.merkle = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Strays {
    const message = { ...baseStrays } as Strays;
    if (object.cid !== undefined && object.cid !== null) {
      message.cid = String(object.cid);
    } else {
      message.cid = "";
    }
    if (object.fid !== undefined && object.fid !== null) {
      message.fid = String(object.fid);
    } else {
      message.fid = "";
    }
    if (object.signee !== undefined && object.signee !== null) {
      message.signee = String(object.signee);
    } else {
      message.signee = "";
    }
    if (object.filesize !== undefined && object.filesize !== null) {
      message.filesize = String(object.filesize);
    } else {
      message.filesize = "";
    }
    if (object.merkle !== undefined && object.merkle !== null) {
      message.merkle = String(object.merkle);
    } else {
      message.merkle = "";
    }
    return message;
  },

  toJSON(message: Strays): unknown {
    const obj: any = {};
    message.cid !== undefined && (obj.cid = message.cid);
    message.fid !== undefined && (obj.fid = message.fid);
    message.signee !== undefined && (obj.signee = message.signee);
    message.filesize !== undefined && (obj.filesize = message.filesize);
    message.merkle !== undefined && (obj.merkle = message.merkle);
    return obj;
  },

  fromPartial(object: DeepPartial<Strays>): Strays {
    const message = { ...baseStrays } as Strays;
    if (object.cid !== undefined && object.cid !== null) {
      message.cid = object.cid;
    } else {
      message.cid = "";
    }
    if (object.fid !== undefined && object.fid !== null) {
      message.fid = object.fid;
    } else {
      message.fid = "";
    }
    if (object.signee !== undefined && object.signee !== null) {
      message.signee = object.signee;
    } else {
      message.signee = "";
    }
    if (object.filesize !== undefined && object.filesize !== null) {
      message.filesize = object.filesize;
    } else {
      message.filesize = "";
    }
    if (object.merkle !== undefined && object.merkle !== null) {
      message.merkle = object.merkle;
    } else {
      message.merkle = "";
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
