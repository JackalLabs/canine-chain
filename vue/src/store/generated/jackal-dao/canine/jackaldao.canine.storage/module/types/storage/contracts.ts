/* eslint-disable */
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "jackaldao.canine.storage";

export interface Contracts {
  cid: string;
  priceamt: string;
  pricedenom: string;
  merkle: string;
  signee: string;
  duration: string;
  filesize: string;
  fid: string;
  creator: string;
}

const baseContracts: object = {
  cid: "",
  priceamt: "",
  pricedenom: "",
  merkle: "",
  signee: "",
  duration: "",
  filesize: "",
  fid: "",
  creator: "",
};

export const Contracts = {
  encode(message: Contracts, writer: Writer = Writer.create()): Writer {
    if (message.cid !== "") {
      writer.uint32(10).string(message.cid);
    }
    if (message.priceamt !== "") {
      writer.uint32(18).string(message.priceamt);
    }
    if (message.pricedenom !== "") {
      writer.uint32(26).string(message.pricedenom);
    }
    if (message.merkle !== "") {
      writer.uint32(42).string(message.merkle);
    }
    if (message.signee !== "") {
      writer.uint32(50).string(message.signee);
    }
    if (message.duration !== "") {
      writer.uint32(58).string(message.duration);
    }
    if (message.filesize !== "") {
      writer.uint32(66).string(message.filesize);
    }
    if (message.fid !== "") {
      writer.uint32(74).string(message.fid);
    }
    if (message.creator !== "") {
      writer.uint32(34).string(message.creator);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): Contracts {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseContracts } as Contracts;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.cid = reader.string();
          break;
        case 2:
          message.priceamt = reader.string();
          break;
        case 3:
          message.pricedenom = reader.string();
          break;
        case 5:
          message.merkle = reader.string();
          break;
        case 6:
          message.signee = reader.string();
          break;
        case 7:
          message.duration = reader.string();
          break;
        case 8:
          message.filesize = reader.string();
          break;
        case 9:
          message.fid = reader.string();
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

  fromJSON(object: any): Contracts {
    const message = { ...baseContracts } as Contracts;
    if (object.cid !== undefined && object.cid !== null) {
      message.cid = String(object.cid);
    } else {
      message.cid = "";
    }
    if (object.priceamt !== undefined && object.priceamt !== null) {
      message.priceamt = String(object.priceamt);
    } else {
      message.priceamt = "";
    }
    if (object.pricedenom !== undefined && object.pricedenom !== null) {
      message.pricedenom = String(object.pricedenom);
    } else {
      message.pricedenom = "";
    }
    if (object.merkle !== undefined && object.merkle !== null) {
      message.merkle = String(object.merkle);
    } else {
      message.merkle = "";
    }
    if (object.signee !== undefined && object.signee !== null) {
      message.signee = String(object.signee);
    } else {
      message.signee = "";
    }
    if (object.duration !== undefined && object.duration !== null) {
      message.duration = String(object.duration);
    } else {
      message.duration = "";
    }
    if (object.filesize !== undefined && object.filesize !== null) {
      message.filesize = String(object.filesize);
    } else {
      message.filesize = "";
    }
    if (object.fid !== undefined && object.fid !== null) {
      message.fid = String(object.fid);
    } else {
      message.fid = "";
    }
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    return message;
  },

  toJSON(message: Contracts): unknown {
    const obj: any = {};
    message.cid !== undefined && (obj.cid = message.cid);
    message.priceamt !== undefined && (obj.priceamt = message.priceamt);
    message.pricedenom !== undefined && (obj.pricedenom = message.pricedenom);
    message.merkle !== undefined && (obj.merkle = message.merkle);
    message.signee !== undefined && (obj.signee = message.signee);
    message.duration !== undefined && (obj.duration = message.duration);
    message.filesize !== undefined && (obj.filesize = message.filesize);
    message.fid !== undefined && (obj.fid = message.fid);
    message.creator !== undefined && (obj.creator = message.creator);
    return obj;
  },

  fromPartial(object: DeepPartial<Contracts>): Contracts {
    const message = { ...baseContracts } as Contracts;
    if (object.cid !== undefined && object.cid !== null) {
      message.cid = object.cid;
    } else {
      message.cid = "";
    }
    if (object.priceamt !== undefined && object.priceamt !== null) {
      message.priceamt = object.priceamt;
    } else {
      message.priceamt = "";
    }
    if (object.pricedenom !== undefined && object.pricedenom !== null) {
      message.pricedenom = object.pricedenom;
    } else {
      message.pricedenom = "";
    }
    if (object.merkle !== undefined && object.merkle !== null) {
      message.merkle = object.merkle;
    } else {
      message.merkle = "";
    }
    if (object.signee !== undefined && object.signee !== null) {
      message.signee = object.signee;
    } else {
      message.signee = "";
    }
    if (object.duration !== undefined && object.duration !== null) {
      message.duration = object.duration;
    } else {
      message.duration = "";
    }
    if (object.filesize !== undefined && object.filesize !== null) {
      message.filesize = object.filesize;
    } else {
      message.filesize = "";
    }
    if (object.fid !== undefined && object.fid !== null) {
      message.fid = object.fid;
    } else {
      message.fid = "";
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
