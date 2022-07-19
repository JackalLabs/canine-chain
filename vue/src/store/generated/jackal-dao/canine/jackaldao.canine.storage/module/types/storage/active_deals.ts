/* eslint-disable */
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "jackaldao.canine.storage";

export interface ActiveDeals {
  cid: string;
  signee: string;
  miner: string;
  startblock: string;
  endblock: string;
  filesize: string;
  proofverified: string;
  proofsmissed: string;
  blocktoprove: string;
  creator: string;
  merkle: string;
  fid: string;
}

const baseActiveDeals: object = {
  cid: "",
  signee: "",
  miner: "",
  startblock: "",
  endblock: "",
  filesize: "",
  proofverified: "",
  proofsmissed: "",
  blocktoprove: "",
  creator: "",
  merkle: "",
  fid: "",
};

export const ActiveDeals = {
  encode(message: ActiveDeals, writer: Writer = Writer.create()): Writer {
    if (message.cid !== "") {
      writer.uint32(10).string(message.cid);
    }
    if (message.signee !== "") {
      writer.uint32(18).string(message.signee);
    }
    if (message.miner !== "") {
      writer.uint32(26).string(message.miner);
    }
    if (message.startblock !== "") {
      writer.uint32(34).string(message.startblock);
    }
    if (message.endblock !== "") {
      writer.uint32(42).string(message.endblock);
    }
    if (message.filesize !== "") {
      writer.uint32(50).string(message.filesize);
    }
    if (message.proofverified !== "") {
      writer.uint32(58).string(message.proofverified);
    }
    if (message.proofsmissed !== "") {
      writer.uint32(66).string(message.proofsmissed);
    }
    if (message.blocktoprove !== "") {
      writer.uint32(74).string(message.blocktoprove);
    }
    if (message.creator !== "") {
      writer.uint32(82).string(message.creator);
    }
    if (message.merkle !== "") {
      writer.uint32(90).string(message.merkle);
    }
    if (message.fid !== "") {
      writer.uint32(98).string(message.fid);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): ActiveDeals {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseActiveDeals } as ActiveDeals;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.cid = reader.string();
          break;
        case 2:
          message.signee = reader.string();
          break;
        case 3:
          message.miner = reader.string();
          break;
        case 4:
          message.startblock = reader.string();
          break;
        case 5:
          message.endblock = reader.string();
          break;
        case 6:
          message.filesize = reader.string();
          break;
        case 7:
          message.proofverified = reader.string();
          break;
        case 8:
          message.proofsmissed = reader.string();
          break;
        case 9:
          message.blocktoprove = reader.string();
          break;
        case 10:
          message.creator = reader.string();
          break;
        case 11:
          message.merkle = reader.string();
          break;
        case 12:
          message.fid = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): ActiveDeals {
    const message = { ...baseActiveDeals } as ActiveDeals;
    if (object.cid !== undefined && object.cid !== null) {
      message.cid = String(object.cid);
    } else {
      message.cid = "";
    }
    if (object.signee !== undefined && object.signee !== null) {
      message.signee = String(object.signee);
    } else {
      message.signee = "";
    }
    if (object.miner !== undefined && object.miner !== null) {
      message.miner = String(object.miner);
    } else {
      message.miner = "";
    }
    if (object.startblock !== undefined && object.startblock !== null) {
      message.startblock = String(object.startblock);
    } else {
      message.startblock = "";
    }
    if (object.endblock !== undefined && object.endblock !== null) {
      message.endblock = String(object.endblock);
    } else {
      message.endblock = "";
    }
    if (object.filesize !== undefined && object.filesize !== null) {
      message.filesize = String(object.filesize);
    } else {
      message.filesize = "";
    }
    if (object.proofverified !== undefined && object.proofverified !== null) {
      message.proofverified = String(object.proofverified);
    } else {
      message.proofverified = "";
    }
    if (object.proofsmissed !== undefined && object.proofsmissed !== null) {
      message.proofsmissed = String(object.proofsmissed);
    } else {
      message.proofsmissed = "";
    }
    if (object.blocktoprove !== undefined && object.blocktoprove !== null) {
      message.blocktoprove = String(object.blocktoprove);
    } else {
      message.blocktoprove = "";
    }
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.merkle !== undefined && object.merkle !== null) {
      message.merkle = String(object.merkle);
    } else {
      message.merkle = "";
    }
    if (object.fid !== undefined && object.fid !== null) {
      message.fid = String(object.fid);
    } else {
      message.fid = "";
    }
    return message;
  },

  toJSON(message: ActiveDeals): unknown {
    const obj: any = {};
    message.cid !== undefined && (obj.cid = message.cid);
    message.signee !== undefined && (obj.signee = message.signee);
    message.miner !== undefined && (obj.miner = message.miner);
    message.startblock !== undefined && (obj.startblock = message.startblock);
    message.endblock !== undefined && (obj.endblock = message.endblock);
    message.filesize !== undefined && (obj.filesize = message.filesize);
    message.proofverified !== undefined &&
      (obj.proofverified = message.proofverified);
    message.proofsmissed !== undefined &&
      (obj.proofsmissed = message.proofsmissed);
    message.blocktoprove !== undefined &&
      (obj.blocktoprove = message.blocktoprove);
    message.creator !== undefined && (obj.creator = message.creator);
    message.merkle !== undefined && (obj.merkle = message.merkle);
    message.fid !== undefined && (obj.fid = message.fid);
    return obj;
  },

  fromPartial(object: DeepPartial<ActiveDeals>): ActiveDeals {
    const message = { ...baseActiveDeals } as ActiveDeals;
    if (object.cid !== undefined && object.cid !== null) {
      message.cid = object.cid;
    } else {
      message.cid = "";
    }
    if (object.signee !== undefined && object.signee !== null) {
      message.signee = object.signee;
    } else {
      message.signee = "";
    }
    if (object.miner !== undefined && object.miner !== null) {
      message.miner = object.miner;
    } else {
      message.miner = "";
    }
    if (object.startblock !== undefined && object.startblock !== null) {
      message.startblock = object.startblock;
    } else {
      message.startblock = "";
    }
    if (object.endblock !== undefined && object.endblock !== null) {
      message.endblock = object.endblock;
    } else {
      message.endblock = "";
    }
    if (object.filesize !== undefined && object.filesize !== null) {
      message.filesize = object.filesize;
    } else {
      message.filesize = "";
    }
    if (object.proofverified !== undefined && object.proofverified !== null) {
      message.proofverified = object.proofverified;
    } else {
      message.proofverified = "";
    }
    if (object.proofsmissed !== undefined && object.proofsmissed !== null) {
      message.proofsmissed = object.proofsmissed;
    } else {
      message.proofsmissed = "";
    }
    if (object.blocktoprove !== undefined && object.blocktoprove !== null) {
      message.blocktoprove = object.blocktoprove;
    } else {
      message.blocktoprove = "";
    }
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.merkle !== undefined && object.merkle !== null) {
      message.merkle = object.merkle;
    } else {
      message.merkle = "";
    }
    if (object.fid !== undefined && object.fid !== null) {
      message.fid = object.fid;
    } else {
      message.fid = "";
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
