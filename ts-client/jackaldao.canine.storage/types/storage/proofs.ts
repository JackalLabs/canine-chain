/* eslint-disable */
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "jackaldao.canine.storage";

export interface Proofs {
  cid: string;
  item: string;
  hashes: string;
  creator: string;
}

const baseProofs: object = { cid: "", item: "", hashes: "", creator: "" };

export const Proofs = {
  encode(message: Proofs, writer: Writer = Writer.create()): Writer {
    if (message.cid !== "") {
      writer.uint32(10).string(message.cid);
    }
    if (message.item !== "") {
      writer.uint32(18).string(message.item);
    }
    if (message.hashes !== "") {
      writer.uint32(26).string(message.hashes);
    }
    if (message.creator !== "") {
      writer.uint32(34).string(message.creator);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): Proofs {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseProofs } as Proofs;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.cid = reader.string();
          break;
        case 2:
          message.item = reader.string();
          break;
        case 3:
          message.hashes = reader.string();
          break;
        case 4:
          message.creator = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Proofs {
    const message = { ...baseProofs } as Proofs;
    if (object.cid !== undefined && object.cid !== null) {
      message.cid = String(object.cid);
    } else {
      message.cid = "";
    }
    if (object.item !== undefined && object.item !== null) {
      message.item = String(object.item);
    } else {
      message.item = "";
    }
    if (object.hashes !== undefined && object.hashes !== null) {
      message.hashes = String(object.hashes);
    } else {
      message.hashes = "";
    }
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    return message;
  },

  toJSON(message: Proofs): unknown {
    const obj: any = {};
    message.cid !== undefined && (obj.cid = message.cid);
    message.item !== undefined && (obj.item = message.item);
    message.hashes !== undefined && (obj.hashes = message.hashes);
    message.creator !== undefined && (obj.creator = message.creator);
    return obj;
  },

  fromPartial(object: DeepPartial<Proofs>): Proofs {
    const message = { ...baseProofs } as Proofs;
    if (object.cid !== undefined && object.cid !== null) {
      message.cid = object.cid;
    } else {
      message.cid = "";
    }
    if (object.item !== undefined && object.item !== null) {
      message.item = object.item;
    } else {
      message.item = "";
    }
    if (object.hashes !== undefined && object.hashes !== null) {
      message.hashes = object.hashes;
    } else {
      message.hashes = "";
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
