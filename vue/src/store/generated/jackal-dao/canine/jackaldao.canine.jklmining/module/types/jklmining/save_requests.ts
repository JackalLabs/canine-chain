/* eslint-disable */
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "jackaldao.canine.jklmining";

export interface SaveRequests {
  index: string;
  size: string;
  approved: string;
  creator: string;
}

const baseSaveRequests: object = {
  index: "",
  size: "",
  approved: "",
  creator: "",
};

export const SaveRequests = {
  encode(message: SaveRequests, writer: Writer = Writer.create()): Writer {
    if (message.index !== "") {
      writer.uint32(10).string(message.index);
    }
    if (message.size !== "") {
      writer.uint32(18).string(message.size);
    }
    if (message.approved !== "") {
      writer.uint32(26).string(message.approved);
    }
    if (message.creator !== "") {
      writer.uint32(34).string(message.creator);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): SaveRequests {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseSaveRequests } as SaveRequests;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.index = reader.string();
          break;
        case 2:
          message.size = reader.string();
          break;
        case 3:
          message.approved = reader.string();
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

  fromJSON(object: any): SaveRequests {
    const message = { ...baseSaveRequests } as SaveRequests;
    if (object.index !== undefined && object.index !== null) {
      message.index = String(object.index);
    } else {
      message.index = "";
    }
    if (object.size !== undefined && object.size !== null) {
      message.size = String(object.size);
    } else {
      message.size = "";
    }
    if (object.approved !== undefined && object.approved !== null) {
      message.approved = String(object.approved);
    } else {
      message.approved = "";
    }
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    return message;
  },

  toJSON(message: SaveRequests): unknown {
    const obj: any = {};
    message.index !== undefined && (obj.index = message.index);
    message.size !== undefined && (obj.size = message.size);
    message.approved !== undefined && (obj.approved = message.approved);
    message.creator !== undefined && (obj.creator = message.creator);
    return obj;
  },

  fromPartial(object: DeepPartial<SaveRequests>): SaveRequests {
    const message = { ...baseSaveRequests } as SaveRequests;
    if (object.index !== undefined && object.index !== null) {
      message.index = object.index;
    } else {
      message.index = "";
    }
    if (object.size !== undefined && object.size !== null) {
      message.size = object.size;
    } else {
      message.size = "";
    }
    if (object.approved !== undefined && object.approved !== null) {
      message.approved = object.approved;
    } else {
      message.approved = "";
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
