/* eslint-disable */
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "jackaldao.canine.dsig";

export interface Form {
  ffid: string;
  cid: string;
  fid: string;
  signees: string;
}

const baseForm: object = { ffid: "", cid: "", fid: "", signees: "" };

export const Form = {
  encode(message: Form, writer: Writer = Writer.create()): Writer {
    if (message.ffid !== "") {
      writer.uint32(10).string(message.ffid);
    }
    if (message.cid !== "") {
      writer.uint32(18).string(message.cid);
    }
    if (message.fid !== "") {
      writer.uint32(26).string(message.fid);
    }
    if (message.signees !== "") {
      writer.uint32(34).string(message.signees);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): Form {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseForm } as Form;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.ffid = reader.string();
          break;
        case 2:
          message.cid = reader.string();
          break;
        case 3:
          message.fid = reader.string();
          break;
        case 4:
          message.signees = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Form {
    const message = { ...baseForm } as Form;
    if (object.ffid !== undefined && object.ffid !== null) {
      message.ffid = String(object.ffid);
    } else {
      message.ffid = "";
    }
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
    if (object.signees !== undefined && object.signees !== null) {
      message.signees = String(object.signees);
    } else {
      message.signees = "";
    }
    return message;
  },

  toJSON(message: Form): unknown {
    const obj: any = {};
    message.ffid !== undefined && (obj.ffid = message.ffid);
    message.cid !== undefined && (obj.cid = message.cid);
    message.fid !== undefined && (obj.fid = message.fid);
    message.signees !== undefined && (obj.signees = message.signees);
    return obj;
  },

  fromPartial(object: DeepPartial<Form>): Form {
    const message = { ...baseForm } as Form;
    if (object.ffid !== undefined && object.ffid !== null) {
      message.ffid = object.ffid;
    } else {
      message.ffid = "";
    }
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
    if (object.signees !== undefined && object.signees !== null) {
      message.signees = object.signees;
    } else {
      message.signees = "";
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
