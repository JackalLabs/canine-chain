/* eslint-disable */
import { Reader, Writer } from "protobufjs/minimal";

export const protobufPackage = "jackaldao.canine.storage";

export interface MsgPostContract {
  creator: string;
  priceamt: string;
  pricedenom: string;
  merkle: string;
  signee: string;
  duration: string;
  filesize: string;
  fid: string;
}

export interface MsgPostContractResponse {}

export interface MsgCreateContracts {
  creator: string;
  cid: string;
  priceamt: string;
  pricedenom: string;
  chunks: string;
  merkle: string;
  signee: string;
  duration: string;
  filesize: string;
  fid: string;
}

export interface MsgCreateContractsResponse {}

export interface MsgUpdateContracts {
  creator: string;
  cid: string;
  chunks: string;
  merkle: string;
  signee: string;
  duration: string;
  filesize: string;
  fid: string;
}

export interface MsgUpdateContractsResponse {}

export interface MsgDeleteContracts {
  creator: string;
  cid: string;
}

export interface MsgDeleteContractsResponse {}

export interface MsgCreateProofs {
  creator: string;
  cid: string;
  item: string;
  hashes: string;
}

export interface MsgCreateProofsResponse {}

export interface MsgUpdateProofs {
  creator: string;
  cid: string;
  item: string;
  hashes: string;
}

export interface MsgUpdateProofsResponse {}

export interface MsgDeleteProofs {
  creator: string;
  cid: string;
}

export interface MsgDeleteProofsResponse {}

export interface MsgItem {
  creator: string;
  hashlist: string;
}

export interface MsgItemResponse {}

export interface MsgPostproof {
  creator: string;
  item: string;
  hashlist: string;
  cid: string;
}

export interface MsgPostproofResponse {
  merkle: string;
}

export interface MsgCreateActiveDeals {
  creator: string;
  cid: string;
  signee: string;
  miner: string;
  startblock: string;
  endblock: string;
  filesize: string;
  proofverified: string;
  proofsmissed: string;
  blocktoprove: string;
}

export interface MsgCreateActiveDealsResponse {}

export interface MsgUpdateActiveDeals {
  creator: string;
  cid: string;
  signee: string;
  miner: string;
  startblock: string;
  endblock: string;
  filesize: string;
  proofverified: string;
  proofsmissed: string;
  blocktoprove: string;
}

export interface MsgUpdateActiveDealsResponse {}

export interface MsgDeleteActiveDeals {
  creator: string;
  cid: string;
}

export interface MsgDeleteActiveDealsResponse {}

export interface MsgSignContract {
  creator: string;
  cid: string;
}

export interface MsgSignContractResponse {}

export interface MsgCreateMiners {
  creator: string;
  address: string;
  ip: string;
  totalspace: string;
}

export interface MsgCreateMinersResponse {}

export interface MsgUpdateMiners {
  creator: string;
  address: string;
  ip: string;
  totalspace: string;
}

export interface MsgUpdateMinersResponse {}

export interface MsgDeleteMiners {
  creator: string;
  address: string;
}

export interface MsgDeleteMinersResponse {}

export interface MsgSetMinerIp {
  creator: string;
  ip: string;
}

export interface MsgSetMinerIpResponse {}

export interface MsgSetMinerTotalspace {
  creator: string;
  space: string;
}

export interface MsgSetMinerTotalspaceResponse {}

export interface MsgInitMiner {
  creator: string;
  ip: string;
  totalspace: string;
}

export interface MsgInitMinerResponse {}

export interface MsgCancelContract {
  creator: string;
  cid: string;
}

export interface MsgCancelContractResponse {}

export interface MsgBuyStorage {
  creator: string;
  forAddress: string;
  duration: string;
  bytes: string;
  paymentDenom: string;
}

export interface MsgBuyStorageResponse {}

const baseMsgPostContract: object = {
  creator: "",
  priceamt: "",
  pricedenom: "",
  merkle: "",
  signee: "",
  duration: "",
  filesize: "",
  fid: "",
};

export const MsgPostContract = {
  encode(message: MsgPostContract, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.priceamt !== "") {
      writer.uint32(18).string(message.priceamt);
    }
    if (message.pricedenom !== "") {
      writer.uint32(26).string(message.pricedenom);
    }
    if (message.merkle !== "") {
      writer.uint32(34).string(message.merkle);
    }
    if (message.signee !== "") {
      writer.uint32(42).string(message.signee);
    }
    if (message.duration !== "") {
      writer.uint32(50).string(message.duration);
    }
    if (message.filesize !== "") {
      writer.uint32(58).string(message.filesize);
    }
    if (message.fid !== "") {
      writer.uint32(66).string(message.fid);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgPostContract {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgPostContract } as MsgPostContract;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.priceamt = reader.string();
          break;
        case 3:
          message.pricedenom = reader.string();
          break;
        case 4:
          message.merkle = reader.string();
          break;
        case 5:
          message.signee = reader.string();
          break;
        case 6:
          message.duration = reader.string();
          break;
        case 7:
          message.filesize = reader.string();
          break;
        case 8:
          message.fid = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgPostContract {
    const message = { ...baseMsgPostContract } as MsgPostContract;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
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
    return message;
  },

  toJSON(message: MsgPostContract): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.priceamt !== undefined && (obj.priceamt = message.priceamt);
    message.pricedenom !== undefined && (obj.pricedenom = message.pricedenom);
    message.merkle !== undefined && (obj.merkle = message.merkle);
    message.signee !== undefined && (obj.signee = message.signee);
    message.duration !== undefined && (obj.duration = message.duration);
    message.filesize !== undefined && (obj.filesize = message.filesize);
    message.fid !== undefined && (obj.fid = message.fid);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgPostContract>): MsgPostContract {
    const message = { ...baseMsgPostContract } as MsgPostContract;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
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
    return message;
  },
};

const baseMsgPostContractResponse: object = {};

export const MsgPostContractResponse = {
  encode(_: MsgPostContractResponse, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgPostContractResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgPostContractResponse,
    } as MsgPostContractResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): MsgPostContractResponse {
    const message = {
      ...baseMsgPostContractResponse,
    } as MsgPostContractResponse;
    return message;
  },

  toJSON(_: MsgPostContractResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgPostContractResponse>
  ): MsgPostContractResponse {
    const message = {
      ...baseMsgPostContractResponse,
    } as MsgPostContractResponse;
    return message;
  },
};

const baseMsgCreateContracts: object = {
  creator: "",
  cid: "",
  priceamt: "",
  pricedenom: "",
  chunks: "",
  merkle: "",
  signee: "",
  duration: "",
  filesize: "",
  fid: "",
};

export const MsgCreateContracts = {
  encode(
    message: MsgCreateContracts,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.cid !== "") {
      writer.uint32(18).string(message.cid);
    }
    if (message.priceamt !== "") {
      writer.uint32(26).string(message.priceamt);
    }
    if (message.pricedenom !== "") {
      writer.uint32(34).string(message.pricedenom);
    }
    if (message.chunks !== "") {
      writer.uint32(42).string(message.chunks);
    }
    if (message.merkle !== "") {
      writer.uint32(50).string(message.merkle);
    }
    if (message.signee !== "") {
      writer.uint32(58).string(message.signee);
    }
    if (message.duration !== "") {
      writer.uint32(66).string(message.duration);
    }
    if (message.filesize !== "") {
      writer.uint32(74).string(message.filesize);
    }
    if (message.fid !== "") {
      writer.uint32(82).string(message.fid);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgCreateContracts {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgCreateContracts } as MsgCreateContracts;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.cid = reader.string();
          break;
        case 3:
          message.priceamt = reader.string();
          break;
        case 4:
          message.pricedenom = reader.string();
          break;
        case 5:
          message.chunks = reader.string();
          break;
        case 6:
          message.merkle = reader.string();
          break;
        case 7:
          message.signee = reader.string();
          break;
        case 8:
          message.duration = reader.string();
          break;
        case 9:
          message.filesize = reader.string();
          break;
        case 10:
          message.fid = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgCreateContracts {
    const message = { ...baseMsgCreateContracts } as MsgCreateContracts;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
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
    if (object.chunks !== undefined && object.chunks !== null) {
      message.chunks = String(object.chunks);
    } else {
      message.chunks = "";
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
    return message;
  },

  toJSON(message: MsgCreateContracts): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.cid !== undefined && (obj.cid = message.cid);
    message.priceamt !== undefined && (obj.priceamt = message.priceamt);
    message.pricedenom !== undefined && (obj.pricedenom = message.pricedenom);
    message.chunks !== undefined && (obj.chunks = message.chunks);
    message.merkle !== undefined && (obj.merkle = message.merkle);
    message.signee !== undefined && (obj.signee = message.signee);
    message.duration !== undefined && (obj.duration = message.duration);
    message.filesize !== undefined && (obj.filesize = message.filesize);
    message.fid !== undefined && (obj.fid = message.fid);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgCreateContracts>): MsgCreateContracts {
    const message = { ...baseMsgCreateContracts } as MsgCreateContracts;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
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
    if (object.chunks !== undefined && object.chunks !== null) {
      message.chunks = object.chunks;
    } else {
      message.chunks = "";
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
    return message;
  },
};

const baseMsgCreateContractsResponse: object = {};

export const MsgCreateContractsResponse = {
  encode(
    _: MsgCreateContractsResponse,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgCreateContractsResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgCreateContractsResponse,
    } as MsgCreateContractsResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): MsgCreateContractsResponse {
    const message = {
      ...baseMsgCreateContractsResponse,
    } as MsgCreateContractsResponse;
    return message;
  },

  toJSON(_: MsgCreateContractsResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgCreateContractsResponse>
  ): MsgCreateContractsResponse {
    const message = {
      ...baseMsgCreateContractsResponse,
    } as MsgCreateContractsResponse;
    return message;
  },
};

const baseMsgUpdateContracts: object = {
  creator: "",
  cid: "",
  chunks: "",
  merkle: "",
  signee: "",
  duration: "",
  filesize: "",
  fid: "",
};

export const MsgUpdateContracts = {
  encode(
    message: MsgUpdateContracts,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.cid !== "") {
      writer.uint32(18).string(message.cid);
    }
    if (message.chunks !== "") {
      writer.uint32(42).string(message.chunks);
    }
    if (message.merkle !== "") {
      writer.uint32(50).string(message.merkle);
    }
    if (message.signee !== "") {
      writer.uint32(58).string(message.signee);
    }
    if (message.duration !== "") {
      writer.uint32(66).string(message.duration);
    }
    if (message.filesize !== "") {
      writer.uint32(34).string(message.filesize);
    }
    if (message.fid !== "") {
      writer.uint32(26).string(message.fid);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgUpdateContracts {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgUpdateContracts } as MsgUpdateContracts;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.cid = reader.string();
          break;
        case 5:
          message.chunks = reader.string();
          break;
        case 6:
          message.merkle = reader.string();
          break;
        case 7:
          message.signee = reader.string();
          break;
        case 8:
          message.duration = reader.string();
          break;
        case 4:
          message.filesize = reader.string();
          break;
        case 3:
          message.fid = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgUpdateContracts {
    const message = { ...baseMsgUpdateContracts } as MsgUpdateContracts;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.cid !== undefined && object.cid !== null) {
      message.cid = String(object.cid);
    } else {
      message.cid = "";
    }
    if (object.chunks !== undefined && object.chunks !== null) {
      message.chunks = String(object.chunks);
    } else {
      message.chunks = "";
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
    return message;
  },

  toJSON(message: MsgUpdateContracts): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.cid !== undefined && (obj.cid = message.cid);
    message.chunks !== undefined && (obj.chunks = message.chunks);
    message.merkle !== undefined && (obj.merkle = message.merkle);
    message.signee !== undefined && (obj.signee = message.signee);
    message.duration !== undefined && (obj.duration = message.duration);
    message.filesize !== undefined && (obj.filesize = message.filesize);
    message.fid !== undefined && (obj.fid = message.fid);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgUpdateContracts>): MsgUpdateContracts {
    const message = { ...baseMsgUpdateContracts } as MsgUpdateContracts;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.cid !== undefined && object.cid !== null) {
      message.cid = object.cid;
    } else {
      message.cid = "";
    }
    if (object.chunks !== undefined && object.chunks !== null) {
      message.chunks = object.chunks;
    } else {
      message.chunks = "";
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
    return message;
  },
};

const baseMsgUpdateContractsResponse: object = {};

export const MsgUpdateContractsResponse = {
  encode(
    _: MsgUpdateContractsResponse,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgUpdateContractsResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgUpdateContractsResponse,
    } as MsgUpdateContractsResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): MsgUpdateContractsResponse {
    const message = {
      ...baseMsgUpdateContractsResponse,
    } as MsgUpdateContractsResponse;
    return message;
  },

  toJSON(_: MsgUpdateContractsResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgUpdateContractsResponse>
  ): MsgUpdateContractsResponse {
    const message = {
      ...baseMsgUpdateContractsResponse,
    } as MsgUpdateContractsResponse;
    return message;
  },
};

const baseMsgDeleteContracts: object = { creator: "", cid: "" };

export const MsgDeleteContracts = {
  encode(
    message: MsgDeleteContracts,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.cid !== "") {
      writer.uint32(18).string(message.cid);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgDeleteContracts {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgDeleteContracts } as MsgDeleteContracts;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.cid = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgDeleteContracts {
    const message = { ...baseMsgDeleteContracts } as MsgDeleteContracts;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.cid !== undefined && object.cid !== null) {
      message.cid = String(object.cid);
    } else {
      message.cid = "";
    }
    return message;
  },

  toJSON(message: MsgDeleteContracts): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.cid !== undefined && (obj.cid = message.cid);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgDeleteContracts>): MsgDeleteContracts {
    const message = { ...baseMsgDeleteContracts } as MsgDeleteContracts;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.cid !== undefined && object.cid !== null) {
      message.cid = object.cid;
    } else {
      message.cid = "";
    }
    return message;
  },
};

const baseMsgDeleteContractsResponse: object = {};

export const MsgDeleteContractsResponse = {
  encode(
    _: MsgDeleteContractsResponse,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgDeleteContractsResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgDeleteContractsResponse,
    } as MsgDeleteContractsResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): MsgDeleteContractsResponse {
    const message = {
      ...baseMsgDeleteContractsResponse,
    } as MsgDeleteContractsResponse;
    return message;
  },

  toJSON(_: MsgDeleteContractsResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgDeleteContractsResponse>
  ): MsgDeleteContractsResponse {
    const message = {
      ...baseMsgDeleteContractsResponse,
    } as MsgDeleteContractsResponse;
    return message;
  },
};

const baseMsgCreateProofs: object = {
  creator: "",
  cid: "",
  item: "",
  hashes: "",
};

export const MsgCreateProofs = {
  encode(message: MsgCreateProofs, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.cid !== "") {
      writer.uint32(18).string(message.cid);
    }
    if (message.item !== "") {
      writer.uint32(26).string(message.item);
    }
    if (message.hashes !== "") {
      writer.uint32(34).string(message.hashes);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgCreateProofs {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgCreateProofs } as MsgCreateProofs;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.cid = reader.string();
          break;
        case 3:
          message.item = reader.string();
          break;
        case 4:
          message.hashes = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgCreateProofs {
    const message = { ...baseMsgCreateProofs } as MsgCreateProofs;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
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
    return message;
  },

  toJSON(message: MsgCreateProofs): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.cid !== undefined && (obj.cid = message.cid);
    message.item !== undefined && (obj.item = message.item);
    message.hashes !== undefined && (obj.hashes = message.hashes);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgCreateProofs>): MsgCreateProofs {
    const message = { ...baseMsgCreateProofs } as MsgCreateProofs;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
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
    return message;
  },
};

const baseMsgCreateProofsResponse: object = {};

export const MsgCreateProofsResponse = {
  encode(_: MsgCreateProofsResponse, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgCreateProofsResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgCreateProofsResponse,
    } as MsgCreateProofsResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): MsgCreateProofsResponse {
    const message = {
      ...baseMsgCreateProofsResponse,
    } as MsgCreateProofsResponse;
    return message;
  },

  toJSON(_: MsgCreateProofsResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgCreateProofsResponse>
  ): MsgCreateProofsResponse {
    const message = {
      ...baseMsgCreateProofsResponse,
    } as MsgCreateProofsResponse;
    return message;
  },
};

const baseMsgUpdateProofs: object = {
  creator: "",
  cid: "",
  item: "",
  hashes: "",
};

export const MsgUpdateProofs = {
  encode(message: MsgUpdateProofs, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.cid !== "") {
      writer.uint32(18).string(message.cid);
    }
    if (message.item !== "") {
      writer.uint32(26).string(message.item);
    }
    if (message.hashes !== "") {
      writer.uint32(34).string(message.hashes);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgUpdateProofs {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgUpdateProofs } as MsgUpdateProofs;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.cid = reader.string();
          break;
        case 3:
          message.item = reader.string();
          break;
        case 4:
          message.hashes = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgUpdateProofs {
    const message = { ...baseMsgUpdateProofs } as MsgUpdateProofs;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
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
    return message;
  },

  toJSON(message: MsgUpdateProofs): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.cid !== undefined && (obj.cid = message.cid);
    message.item !== undefined && (obj.item = message.item);
    message.hashes !== undefined && (obj.hashes = message.hashes);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgUpdateProofs>): MsgUpdateProofs {
    const message = { ...baseMsgUpdateProofs } as MsgUpdateProofs;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
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
    return message;
  },
};

const baseMsgUpdateProofsResponse: object = {};

export const MsgUpdateProofsResponse = {
  encode(_: MsgUpdateProofsResponse, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgUpdateProofsResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgUpdateProofsResponse,
    } as MsgUpdateProofsResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): MsgUpdateProofsResponse {
    const message = {
      ...baseMsgUpdateProofsResponse,
    } as MsgUpdateProofsResponse;
    return message;
  },

  toJSON(_: MsgUpdateProofsResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgUpdateProofsResponse>
  ): MsgUpdateProofsResponse {
    const message = {
      ...baseMsgUpdateProofsResponse,
    } as MsgUpdateProofsResponse;
    return message;
  },
};

const baseMsgDeleteProofs: object = { creator: "", cid: "" };

export const MsgDeleteProofs = {
  encode(message: MsgDeleteProofs, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.cid !== "") {
      writer.uint32(18).string(message.cid);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgDeleteProofs {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgDeleteProofs } as MsgDeleteProofs;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.cid = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgDeleteProofs {
    const message = { ...baseMsgDeleteProofs } as MsgDeleteProofs;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.cid !== undefined && object.cid !== null) {
      message.cid = String(object.cid);
    } else {
      message.cid = "";
    }
    return message;
  },

  toJSON(message: MsgDeleteProofs): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.cid !== undefined && (obj.cid = message.cid);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgDeleteProofs>): MsgDeleteProofs {
    const message = { ...baseMsgDeleteProofs } as MsgDeleteProofs;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.cid !== undefined && object.cid !== null) {
      message.cid = object.cid;
    } else {
      message.cid = "";
    }
    return message;
  },
};

const baseMsgDeleteProofsResponse: object = {};

export const MsgDeleteProofsResponse = {
  encode(_: MsgDeleteProofsResponse, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgDeleteProofsResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgDeleteProofsResponse,
    } as MsgDeleteProofsResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): MsgDeleteProofsResponse {
    const message = {
      ...baseMsgDeleteProofsResponse,
    } as MsgDeleteProofsResponse;
    return message;
  },

  toJSON(_: MsgDeleteProofsResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgDeleteProofsResponse>
  ): MsgDeleteProofsResponse {
    const message = {
      ...baseMsgDeleteProofsResponse,
    } as MsgDeleteProofsResponse;
    return message;
  },
};

const baseMsgItem: object = { creator: "", hashlist: "" };

export const MsgItem = {
  encode(message: MsgItem, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.hashlist !== "") {
      writer.uint32(18).string(message.hashlist);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgItem {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgItem } as MsgItem;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.hashlist = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgItem {
    const message = { ...baseMsgItem } as MsgItem;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.hashlist !== undefined && object.hashlist !== null) {
      message.hashlist = String(object.hashlist);
    } else {
      message.hashlist = "";
    }
    return message;
  },

  toJSON(message: MsgItem): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.hashlist !== undefined && (obj.hashlist = message.hashlist);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgItem>): MsgItem {
    const message = { ...baseMsgItem } as MsgItem;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.hashlist !== undefined && object.hashlist !== null) {
      message.hashlist = object.hashlist;
    } else {
      message.hashlist = "";
    }
    return message;
  },
};

const baseMsgItemResponse: object = {};

export const MsgItemResponse = {
  encode(_: MsgItemResponse, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgItemResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgItemResponse } as MsgItemResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): MsgItemResponse {
    const message = { ...baseMsgItemResponse } as MsgItemResponse;
    return message;
  },

  toJSON(_: MsgItemResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(_: DeepPartial<MsgItemResponse>): MsgItemResponse {
    const message = { ...baseMsgItemResponse } as MsgItemResponse;
    return message;
  },
};

const baseMsgPostproof: object = {
  creator: "",
  item: "",
  hashlist: "",
  cid: "",
};

export const MsgPostproof = {
  encode(message: MsgPostproof, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.item !== "") {
      writer.uint32(18).string(message.item);
    }
    if (message.hashlist !== "") {
      writer.uint32(26).string(message.hashlist);
    }
    if (message.cid !== "") {
      writer.uint32(34).string(message.cid);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgPostproof {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgPostproof } as MsgPostproof;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.item = reader.string();
          break;
        case 3:
          message.hashlist = reader.string();
          break;
        case 4:
          message.cid = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgPostproof {
    const message = { ...baseMsgPostproof } as MsgPostproof;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.item !== undefined && object.item !== null) {
      message.item = String(object.item);
    } else {
      message.item = "";
    }
    if (object.hashlist !== undefined && object.hashlist !== null) {
      message.hashlist = String(object.hashlist);
    } else {
      message.hashlist = "";
    }
    if (object.cid !== undefined && object.cid !== null) {
      message.cid = String(object.cid);
    } else {
      message.cid = "";
    }
    return message;
  },

  toJSON(message: MsgPostproof): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.item !== undefined && (obj.item = message.item);
    message.hashlist !== undefined && (obj.hashlist = message.hashlist);
    message.cid !== undefined && (obj.cid = message.cid);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgPostproof>): MsgPostproof {
    const message = { ...baseMsgPostproof } as MsgPostproof;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.item !== undefined && object.item !== null) {
      message.item = object.item;
    } else {
      message.item = "";
    }
    if (object.hashlist !== undefined && object.hashlist !== null) {
      message.hashlist = object.hashlist;
    } else {
      message.hashlist = "";
    }
    if (object.cid !== undefined && object.cid !== null) {
      message.cid = object.cid;
    } else {
      message.cid = "";
    }
    return message;
  },
};

const baseMsgPostproofResponse: object = { merkle: "" };

export const MsgPostproofResponse = {
  encode(
    message: MsgPostproofResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.merkle !== "") {
      writer.uint32(10).string(message.merkle);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgPostproofResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgPostproofResponse } as MsgPostproofResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.merkle = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgPostproofResponse {
    const message = { ...baseMsgPostproofResponse } as MsgPostproofResponse;
    if (object.merkle !== undefined && object.merkle !== null) {
      message.merkle = String(object.merkle);
    } else {
      message.merkle = "";
    }
    return message;
  },

  toJSON(message: MsgPostproofResponse): unknown {
    const obj: any = {};
    message.merkle !== undefined && (obj.merkle = message.merkle);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgPostproofResponse>): MsgPostproofResponse {
    const message = { ...baseMsgPostproofResponse } as MsgPostproofResponse;
    if (object.merkle !== undefined && object.merkle !== null) {
      message.merkle = object.merkle;
    } else {
      message.merkle = "";
    }
    return message;
  },
};

const baseMsgCreateActiveDeals: object = {
  creator: "",
  cid: "",
  signee: "",
  miner: "",
  startblock: "",
  endblock: "",
  filesize: "",
  proofverified: "",
  proofsmissed: "",
  blocktoprove: "",
};

export const MsgCreateActiveDeals = {
  encode(
    message: MsgCreateActiveDeals,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.cid !== "") {
      writer.uint32(18).string(message.cid);
    }
    if (message.signee !== "") {
      writer.uint32(26).string(message.signee);
    }
    if (message.miner !== "") {
      writer.uint32(34).string(message.miner);
    }
    if (message.startblock !== "") {
      writer.uint32(42).string(message.startblock);
    }
    if (message.endblock !== "") {
      writer.uint32(50).string(message.endblock);
    }
    if (message.filesize !== "") {
      writer.uint32(58).string(message.filesize);
    }
    if (message.proofverified !== "") {
      writer.uint32(66).string(message.proofverified);
    }
    if (message.proofsmissed !== "") {
      writer.uint32(74).string(message.proofsmissed);
    }
    if (message.blocktoprove !== "") {
      writer.uint32(82).string(message.blocktoprove);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgCreateActiveDeals {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgCreateActiveDeals } as MsgCreateActiveDeals;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.cid = reader.string();
          break;
        case 3:
          message.signee = reader.string();
          break;
        case 4:
          message.miner = reader.string();
          break;
        case 5:
          message.startblock = reader.string();
          break;
        case 6:
          message.endblock = reader.string();
          break;
        case 7:
          message.filesize = reader.string();
          break;
        case 8:
          message.proofverified = reader.string();
          break;
        case 9:
          message.proofsmissed = reader.string();
          break;
        case 10:
          message.blocktoprove = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgCreateActiveDeals {
    const message = { ...baseMsgCreateActiveDeals } as MsgCreateActiveDeals;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
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
    return message;
  },

  toJSON(message: MsgCreateActiveDeals): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
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
    return obj;
  },

  fromPartial(object: DeepPartial<MsgCreateActiveDeals>): MsgCreateActiveDeals {
    const message = { ...baseMsgCreateActiveDeals } as MsgCreateActiveDeals;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
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
    return message;
  },
};

const baseMsgCreateActiveDealsResponse: object = {};

export const MsgCreateActiveDealsResponse = {
  encode(
    _: MsgCreateActiveDealsResponse,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgCreateActiveDealsResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgCreateActiveDealsResponse,
    } as MsgCreateActiveDealsResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): MsgCreateActiveDealsResponse {
    const message = {
      ...baseMsgCreateActiveDealsResponse,
    } as MsgCreateActiveDealsResponse;
    return message;
  },

  toJSON(_: MsgCreateActiveDealsResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgCreateActiveDealsResponse>
  ): MsgCreateActiveDealsResponse {
    const message = {
      ...baseMsgCreateActiveDealsResponse,
    } as MsgCreateActiveDealsResponse;
    return message;
  },
};

const baseMsgUpdateActiveDeals: object = {
  creator: "",
  cid: "",
  signee: "",
  miner: "",
  startblock: "",
  endblock: "",
  filesize: "",
  proofverified: "",
  proofsmissed: "",
  blocktoprove: "",
};

export const MsgUpdateActiveDeals = {
  encode(
    message: MsgUpdateActiveDeals,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.cid !== "") {
      writer.uint32(18).string(message.cid);
    }
    if (message.signee !== "") {
      writer.uint32(26).string(message.signee);
    }
    if (message.miner !== "") {
      writer.uint32(34).string(message.miner);
    }
    if (message.startblock !== "") {
      writer.uint32(42).string(message.startblock);
    }
    if (message.endblock !== "") {
      writer.uint32(50).string(message.endblock);
    }
    if (message.filesize !== "") {
      writer.uint32(58).string(message.filesize);
    }
    if (message.proofverified !== "") {
      writer.uint32(66).string(message.proofverified);
    }
    if (message.proofsmissed !== "") {
      writer.uint32(74).string(message.proofsmissed);
    }
    if (message.blocktoprove !== "") {
      writer.uint32(82).string(message.blocktoprove);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgUpdateActiveDeals {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgUpdateActiveDeals } as MsgUpdateActiveDeals;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.cid = reader.string();
          break;
        case 3:
          message.signee = reader.string();
          break;
        case 4:
          message.miner = reader.string();
          break;
        case 5:
          message.startblock = reader.string();
          break;
        case 6:
          message.endblock = reader.string();
          break;
        case 7:
          message.filesize = reader.string();
          break;
        case 8:
          message.proofverified = reader.string();
          break;
        case 9:
          message.proofsmissed = reader.string();
          break;
        case 10:
          message.blocktoprove = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgUpdateActiveDeals {
    const message = { ...baseMsgUpdateActiveDeals } as MsgUpdateActiveDeals;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
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
    return message;
  },

  toJSON(message: MsgUpdateActiveDeals): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
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
    return obj;
  },

  fromPartial(object: DeepPartial<MsgUpdateActiveDeals>): MsgUpdateActiveDeals {
    const message = { ...baseMsgUpdateActiveDeals } as MsgUpdateActiveDeals;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
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
    return message;
  },
};

const baseMsgUpdateActiveDealsResponse: object = {};

export const MsgUpdateActiveDealsResponse = {
  encode(
    _: MsgUpdateActiveDealsResponse,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgUpdateActiveDealsResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgUpdateActiveDealsResponse,
    } as MsgUpdateActiveDealsResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): MsgUpdateActiveDealsResponse {
    const message = {
      ...baseMsgUpdateActiveDealsResponse,
    } as MsgUpdateActiveDealsResponse;
    return message;
  },

  toJSON(_: MsgUpdateActiveDealsResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgUpdateActiveDealsResponse>
  ): MsgUpdateActiveDealsResponse {
    const message = {
      ...baseMsgUpdateActiveDealsResponse,
    } as MsgUpdateActiveDealsResponse;
    return message;
  },
};

const baseMsgDeleteActiveDeals: object = { creator: "", cid: "" };

export const MsgDeleteActiveDeals = {
  encode(
    message: MsgDeleteActiveDeals,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.cid !== "") {
      writer.uint32(18).string(message.cid);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgDeleteActiveDeals {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgDeleteActiveDeals } as MsgDeleteActiveDeals;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.cid = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgDeleteActiveDeals {
    const message = { ...baseMsgDeleteActiveDeals } as MsgDeleteActiveDeals;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.cid !== undefined && object.cid !== null) {
      message.cid = String(object.cid);
    } else {
      message.cid = "";
    }
    return message;
  },

  toJSON(message: MsgDeleteActiveDeals): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.cid !== undefined && (obj.cid = message.cid);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgDeleteActiveDeals>): MsgDeleteActiveDeals {
    const message = { ...baseMsgDeleteActiveDeals } as MsgDeleteActiveDeals;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.cid !== undefined && object.cid !== null) {
      message.cid = object.cid;
    } else {
      message.cid = "";
    }
    return message;
  },
};

const baseMsgDeleteActiveDealsResponse: object = {};

export const MsgDeleteActiveDealsResponse = {
  encode(
    _: MsgDeleteActiveDealsResponse,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgDeleteActiveDealsResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgDeleteActiveDealsResponse,
    } as MsgDeleteActiveDealsResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): MsgDeleteActiveDealsResponse {
    const message = {
      ...baseMsgDeleteActiveDealsResponse,
    } as MsgDeleteActiveDealsResponse;
    return message;
  },

  toJSON(_: MsgDeleteActiveDealsResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgDeleteActiveDealsResponse>
  ): MsgDeleteActiveDealsResponse {
    const message = {
      ...baseMsgDeleteActiveDealsResponse,
    } as MsgDeleteActiveDealsResponse;
    return message;
  },
};

const baseMsgSignContract: object = { creator: "", cid: "" };

export const MsgSignContract = {
  encode(message: MsgSignContract, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.cid !== "") {
      writer.uint32(18).string(message.cid);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgSignContract {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgSignContract } as MsgSignContract;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.cid = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgSignContract {
    const message = { ...baseMsgSignContract } as MsgSignContract;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.cid !== undefined && object.cid !== null) {
      message.cid = String(object.cid);
    } else {
      message.cid = "";
    }
    return message;
  },

  toJSON(message: MsgSignContract): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.cid !== undefined && (obj.cid = message.cid);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgSignContract>): MsgSignContract {
    const message = { ...baseMsgSignContract } as MsgSignContract;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.cid !== undefined && object.cid !== null) {
      message.cid = object.cid;
    } else {
      message.cid = "";
    }
    return message;
  },
};

const baseMsgSignContractResponse: object = {};

export const MsgSignContractResponse = {
  encode(_: MsgSignContractResponse, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgSignContractResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgSignContractResponse,
    } as MsgSignContractResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): MsgSignContractResponse {
    const message = {
      ...baseMsgSignContractResponse,
    } as MsgSignContractResponse;
    return message;
  },

  toJSON(_: MsgSignContractResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgSignContractResponse>
  ): MsgSignContractResponse {
    const message = {
      ...baseMsgSignContractResponse,
    } as MsgSignContractResponse;
    return message;
  },
};

const baseMsgCreateMiners: object = {
  creator: "",
  address: "",
  ip: "",
  totalspace: "",
};

export const MsgCreateMiners = {
  encode(message: MsgCreateMiners, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.address !== "") {
      writer.uint32(18).string(message.address);
    }
    if (message.ip !== "") {
      writer.uint32(26).string(message.ip);
    }
    if (message.totalspace !== "") {
      writer.uint32(34).string(message.totalspace);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgCreateMiners {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgCreateMiners } as MsgCreateMiners;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.address = reader.string();
          break;
        case 3:
          message.ip = reader.string();
          break;
        case 4:
          message.totalspace = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgCreateMiners {
    const message = { ...baseMsgCreateMiners } as MsgCreateMiners;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.address !== undefined && object.address !== null) {
      message.address = String(object.address);
    } else {
      message.address = "";
    }
    if (object.ip !== undefined && object.ip !== null) {
      message.ip = String(object.ip);
    } else {
      message.ip = "";
    }
    if (object.totalspace !== undefined && object.totalspace !== null) {
      message.totalspace = String(object.totalspace);
    } else {
      message.totalspace = "";
    }
    return message;
  },

  toJSON(message: MsgCreateMiners): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.address !== undefined && (obj.address = message.address);
    message.ip !== undefined && (obj.ip = message.ip);
    message.totalspace !== undefined && (obj.totalspace = message.totalspace);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgCreateMiners>): MsgCreateMiners {
    const message = { ...baseMsgCreateMiners } as MsgCreateMiners;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.address !== undefined && object.address !== null) {
      message.address = object.address;
    } else {
      message.address = "";
    }
    if (object.ip !== undefined && object.ip !== null) {
      message.ip = object.ip;
    } else {
      message.ip = "";
    }
    if (object.totalspace !== undefined && object.totalspace !== null) {
      message.totalspace = object.totalspace;
    } else {
      message.totalspace = "";
    }
    return message;
  },
};

const baseMsgCreateMinersResponse: object = {};

export const MsgCreateMinersResponse = {
  encode(_: MsgCreateMinersResponse, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgCreateMinersResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgCreateMinersResponse,
    } as MsgCreateMinersResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): MsgCreateMinersResponse {
    const message = {
      ...baseMsgCreateMinersResponse,
    } as MsgCreateMinersResponse;
    return message;
  },

  toJSON(_: MsgCreateMinersResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgCreateMinersResponse>
  ): MsgCreateMinersResponse {
    const message = {
      ...baseMsgCreateMinersResponse,
    } as MsgCreateMinersResponse;
    return message;
  },
};

const baseMsgUpdateMiners: object = {
  creator: "",
  address: "",
  ip: "",
  totalspace: "",
};

export const MsgUpdateMiners = {
  encode(message: MsgUpdateMiners, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.address !== "") {
      writer.uint32(18).string(message.address);
    }
    if (message.ip !== "") {
      writer.uint32(26).string(message.ip);
    }
    if (message.totalspace !== "") {
      writer.uint32(34).string(message.totalspace);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgUpdateMiners {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgUpdateMiners } as MsgUpdateMiners;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.address = reader.string();
          break;
        case 3:
          message.ip = reader.string();
          break;
        case 4:
          message.totalspace = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgUpdateMiners {
    const message = { ...baseMsgUpdateMiners } as MsgUpdateMiners;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.address !== undefined && object.address !== null) {
      message.address = String(object.address);
    } else {
      message.address = "";
    }
    if (object.ip !== undefined && object.ip !== null) {
      message.ip = String(object.ip);
    } else {
      message.ip = "";
    }
    if (object.totalspace !== undefined && object.totalspace !== null) {
      message.totalspace = String(object.totalspace);
    } else {
      message.totalspace = "";
    }
    return message;
  },

  toJSON(message: MsgUpdateMiners): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.address !== undefined && (obj.address = message.address);
    message.ip !== undefined && (obj.ip = message.ip);
    message.totalspace !== undefined && (obj.totalspace = message.totalspace);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgUpdateMiners>): MsgUpdateMiners {
    const message = { ...baseMsgUpdateMiners } as MsgUpdateMiners;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.address !== undefined && object.address !== null) {
      message.address = object.address;
    } else {
      message.address = "";
    }
    if (object.ip !== undefined && object.ip !== null) {
      message.ip = object.ip;
    } else {
      message.ip = "";
    }
    if (object.totalspace !== undefined && object.totalspace !== null) {
      message.totalspace = object.totalspace;
    } else {
      message.totalspace = "";
    }
    return message;
  },
};

const baseMsgUpdateMinersResponse: object = {};

export const MsgUpdateMinersResponse = {
  encode(_: MsgUpdateMinersResponse, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgUpdateMinersResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgUpdateMinersResponse,
    } as MsgUpdateMinersResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): MsgUpdateMinersResponse {
    const message = {
      ...baseMsgUpdateMinersResponse,
    } as MsgUpdateMinersResponse;
    return message;
  },

  toJSON(_: MsgUpdateMinersResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgUpdateMinersResponse>
  ): MsgUpdateMinersResponse {
    const message = {
      ...baseMsgUpdateMinersResponse,
    } as MsgUpdateMinersResponse;
    return message;
  },
};

const baseMsgDeleteMiners: object = { creator: "", address: "" };

export const MsgDeleteMiners = {
  encode(message: MsgDeleteMiners, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.address !== "") {
      writer.uint32(18).string(message.address);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgDeleteMiners {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgDeleteMiners } as MsgDeleteMiners;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.address = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgDeleteMiners {
    const message = { ...baseMsgDeleteMiners } as MsgDeleteMiners;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.address !== undefined && object.address !== null) {
      message.address = String(object.address);
    } else {
      message.address = "";
    }
    return message;
  },

  toJSON(message: MsgDeleteMiners): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.address !== undefined && (obj.address = message.address);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgDeleteMiners>): MsgDeleteMiners {
    const message = { ...baseMsgDeleteMiners } as MsgDeleteMiners;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.address !== undefined && object.address !== null) {
      message.address = object.address;
    } else {
      message.address = "";
    }
    return message;
  },
};

const baseMsgDeleteMinersResponse: object = {};

export const MsgDeleteMinersResponse = {
  encode(_: MsgDeleteMinersResponse, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgDeleteMinersResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgDeleteMinersResponse,
    } as MsgDeleteMinersResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): MsgDeleteMinersResponse {
    const message = {
      ...baseMsgDeleteMinersResponse,
    } as MsgDeleteMinersResponse;
    return message;
  },

  toJSON(_: MsgDeleteMinersResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgDeleteMinersResponse>
  ): MsgDeleteMinersResponse {
    const message = {
      ...baseMsgDeleteMinersResponse,
    } as MsgDeleteMinersResponse;
    return message;
  },
};

const baseMsgSetMinerIp: object = { creator: "", ip: "" };

export const MsgSetMinerIp = {
  encode(message: MsgSetMinerIp, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.ip !== "") {
      writer.uint32(18).string(message.ip);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgSetMinerIp {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgSetMinerIp } as MsgSetMinerIp;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.ip = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgSetMinerIp {
    const message = { ...baseMsgSetMinerIp } as MsgSetMinerIp;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.ip !== undefined && object.ip !== null) {
      message.ip = String(object.ip);
    } else {
      message.ip = "";
    }
    return message;
  },

  toJSON(message: MsgSetMinerIp): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.ip !== undefined && (obj.ip = message.ip);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgSetMinerIp>): MsgSetMinerIp {
    const message = { ...baseMsgSetMinerIp } as MsgSetMinerIp;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.ip !== undefined && object.ip !== null) {
      message.ip = object.ip;
    } else {
      message.ip = "";
    }
    return message;
  },
};

const baseMsgSetMinerIpResponse: object = {};

export const MsgSetMinerIpResponse = {
  encode(_: MsgSetMinerIpResponse, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgSetMinerIpResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgSetMinerIpResponse } as MsgSetMinerIpResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): MsgSetMinerIpResponse {
    const message = { ...baseMsgSetMinerIpResponse } as MsgSetMinerIpResponse;
    return message;
  },

  toJSON(_: MsgSetMinerIpResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(_: DeepPartial<MsgSetMinerIpResponse>): MsgSetMinerIpResponse {
    const message = { ...baseMsgSetMinerIpResponse } as MsgSetMinerIpResponse;
    return message;
  },
};

const baseMsgSetMinerTotalspace: object = { creator: "", space: "" };

export const MsgSetMinerTotalspace = {
  encode(
    message: MsgSetMinerTotalspace,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.space !== "") {
      writer.uint32(18).string(message.space);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgSetMinerTotalspace {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgSetMinerTotalspace } as MsgSetMinerTotalspace;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.space = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgSetMinerTotalspace {
    const message = { ...baseMsgSetMinerTotalspace } as MsgSetMinerTotalspace;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.space !== undefined && object.space !== null) {
      message.space = String(object.space);
    } else {
      message.space = "";
    }
    return message;
  },

  toJSON(message: MsgSetMinerTotalspace): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.space !== undefined && (obj.space = message.space);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgSetMinerTotalspace>
  ): MsgSetMinerTotalspace {
    const message = { ...baseMsgSetMinerTotalspace } as MsgSetMinerTotalspace;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.space !== undefined && object.space !== null) {
      message.space = object.space;
    } else {
      message.space = "";
    }
    return message;
  },
};

const baseMsgSetMinerTotalspaceResponse: object = {};

export const MsgSetMinerTotalspaceResponse = {
  encode(
    _: MsgSetMinerTotalspaceResponse,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgSetMinerTotalspaceResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgSetMinerTotalspaceResponse,
    } as MsgSetMinerTotalspaceResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): MsgSetMinerTotalspaceResponse {
    const message = {
      ...baseMsgSetMinerTotalspaceResponse,
    } as MsgSetMinerTotalspaceResponse;
    return message;
  },

  toJSON(_: MsgSetMinerTotalspaceResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgSetMinerTotalspaceResponse>
  ): MsgSetMinerTotalspaceResponse {
    const message = {
      ...baseMsgSetMinerTotalspaceResponse,
    } as MsgSetMinerTotalspaceResponse;
    return message;
  },
};

const baseMsgInitMiner: object = { creator: "", ip: "", totalspace: "" };

export const MsgInitMiner = {
  encode(message: MsgInitMiner, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.ip !== "") {
      writer.uint32(18).string(message.ip);
    }
    if (message.totalspace !== "") {
      writer.uint32(26).string(message.totalspace);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgInitMiner {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgInitMiner } as MsgInitMiner;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.ip = reader.string();
          break;
        case 3:
          message.totalspace = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgInitMiner {
    const message = { ...baseMsgInitMiner } as MsgInitMiner;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.ip !== undefined && object.ip !== null) {
      message.ip = String(object.ip);
    } else {
      message.ip = "";
    }
    if (object.totalspace !== undefined && object.totalspace !== null) {
      message.totalspace = String(object.totalspace);
    } else {
      message.totalspace = "";
    }
    return message;
  },

  toJSON(message: MsgInitMiner): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.ip !== undefined && (obj.ip = message.ip);
    message.totalspace !== undefined && (obj.totalspace = message.totalspace);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgInitMiner>): MsgInitMiner {
    const message = { ...baseMsgInitMiner } as MsgInitMiner;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.ip !== undefined && object.ip !== null) {
      message.ip = object.ip;
    } else {
      message.ip = "";
    }
    if (object.totalspace !== undefined && object.totalspace !== null) {
      message.totalspace = object.totalspace;
    } else {
      message.totalspace = "";
    }
    return message;
  },
};

const baseMsgInitMinerResponse: object = {};

export const MsgInitMinerResponse = {
  encode(_: MsgInitMinerResponse, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgInitMinerResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgInitMinerResponse } as MsgInitMinerResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): MsgInitMinerResponse {
    const message = { ...baseMsgInitMinerResponse } as MsgInitMinerResponse;
    return message;
  },

  toJSON(_: MsgInitMinerResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(_: DeepPartial<MsgInitMinerResponse>): MsgInitMinerResponse {
    const message = { ...baseMsgInitMinerResponse } as MsgInitMinerResponse;
    return message;
  },
};

const baseMsgCancelContract: object = { creator: "", cid: "" };

export const MsgCancelContract = {
  encode(message: MsgCancelContract, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.cid !== "") {
      writer.uint32(18).string(message.cid);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgCancelContract {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgCancelContract } as MsgCancelContract;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.cid = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgCancelContract {
    const message = { ...baseMsgCancelContract } as MsgCancelContract;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.cid !== undefined && object.cid !== null) {
      message.cid = String(object.cid);
    } else {
      message.cid = "";
    }
    return message;
  },

  toJSON(message: MsgCancelContract): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.cid !== undefined && (obj.cid = message.cid);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgCancelContract>): MsgCancelContract {
    const message = { ...baseMsgCancelContract } as MsgCancelContract;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.cid !== undefined && object.cid !== null) {
      message.cid = object.cid;
    } else {
      message.cid = "";
    }
    return message;
  },
};

const baseMsgCancelContractResponse: object = {};

export const MsgCancelContractResponse = {
  encode(
    _: MsgCancelContractResponse,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgCancelContractResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgCancelContractResponse,
    } as MsgCancelContractResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): MsgCancelContractResponse {
    const message = {
      ...baseMsgCancelContractResponse,
    } as MsgCancelContractResponse;
    return message;
  },

  toJSON(_: MsgCancelContractResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgCancelContractResponse>
  ): MsgCancelContractResponse {
    const message = {
      ...baseMsgCancelContractResponse,
    } as MsgCancelContractResponse;
    return message;
  },
};

const baseMsgBuyStorage: object = {
  creator: "",
  forAddress: "",
  duration: "",
  bytes: "",
  paymentDenom: "",
};

export const MsgBuyStorage = {
  encode(message: MsgBuyStorage, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.forAddress !== "") {
      writer.uint32(18).string(message.forAddress);
    }
    if (message.duration !== "") {
      writer.uint32(26).string(message.duration);
    }
    if (message.bytes !== "") {
      writer.uint32(34).string(message.bytes);
    }
    if (message.paymentDenom !== "") {
      writer.uint32(42).string(message.paymentDenom);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgBuyStorage {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgBuyStorage } as MsgBuyStorage;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.forAddress = reader.string();
          break;
        case 3:
          message.duration = reader.string();
          break;
        case 4:
          message.bytes = reader.string();
          break;
        case 5:
          message.paymentDenom = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgBuyStorage {
    const message = { ...baseMsgBuyStorage } as MsgBuyStorage;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.forAddress !== undefined && object.forAddress !== null) {
      message.forAddress = String(object.forAddress);
    } else {
      message.forAddress = "";
    }
    if (object.duration !== undefined && object.duration !== null) {
      message.duration = String(object.duration);
    } else {
      message.duration = "";
    }
    if (object.bytes !== undefined && object.bytes !== null) {
      message.bytes = String(object.bytes);
    } else {
      message.bytes = "";
    }
    if (object.paymentDenom !== undefined && object.paymentDenom !== null) {
      message.paymentDenom = String(object.paymentDenom);
    } else {
      message.paymentDenom = "";
    }
    return message;
  },

  toJSON(message: MsgBuyStorage): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.forAddress !== undefined && (obj.forAddress = message.forAddress);
    message.duration !== undefined && (obj.duration = message.duration);
    message.bytes !== undefined && (obj.bytes = message.bytes);
    message.paymentDenom !== undefined &&
      (obj.paymentDenom = message.paymentDenom);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgBuyStorage>): MsgBuyStorage {
    const message = { ...baseMsgBuyStorage } as MsgBuyStorage;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.forAddress !== undefined && object.forAddress !== null) {
      message.forAddress = object.forAddress;
    } else {
      message.forAddress = "";
    }
    if (object.duration !== undefined && object.duration !== null) {
      message.duration = object.duration;
    } else {
      message.duration = "";
    }
    if (object.bytes !== undefined && object.bytes !== null) {
      message.bytes = object.bytes;
    } else {
      message.bytes = "";
    }
    if (object.paymentDenom !== undefined && object.paymentDenom !== null) {
      message.paymentDenom = object.paymentDenom;
    } else {
      message.paymentDenom = "";
    }
    return message;
  },
};

const baseMsgBuyStorageResponse: object = {};

export const MsgBuyStorageResponse = {
  encode(_: MsgBuyStorageResponse, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgBuyStorageResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgBuyStorageResponse } as MsgBuyStorageResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): MsgBuyStorageResponse {
    const message = { ...baseMsgBuyStorageResponse } as MsgBuyStorageResponse;
    return message;
  },

  toJSON(_: MsgBuyStorageResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(_: DeepPartial<MsgBuyStorageResponse>): MsgBuyStorageResponse {
    const message = { ...baseMsgBuyStorageResponse } as MsgBuyStorageResponse;
    return message;
  },
};

/** Msg defines the Msg service. */
export interface Msg {
  PostContract(request: MsgPostContract): Promise<MsgPostContractResponse>;
  Postproof(request: MsgPostproof): Promise<MsgPostproofResponse>;
  SignContract(request: MsgSignContract): Promise<MsgSignContractResponse>;
  SetMinerIp(request: MsgSetMinerIp): Promise<MsgSetMinerIpResponse>;
  SetMinerTotalspace(
    request: MsgSetMinerTotalspace
  ): Promise<MsgSetMinerTotalspaceResponse>;
  InitMiner(request: MsgInitMiner): Promise<MsgInitMinerResponse>;
  CancelContract(
    request: MsgCancelContract
  ): Promise<MsgCancelContractResponse>;
  /** this line is used by starport scaffolding # proto/tx/rpc */
  BuyStorage(request: MsgBuyStorage): Promise<MsgBuyStorageResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  PostContract(request: MsgPostContract): Promise<MsgPostContractResponse> {
    const data = MsgPostContract.encode(request).finish();
    const promise = this.rpc.request(
      "jackaldao.canine.storage.Msg",
      "PostContract",
      data
    );
    return promise.then((data) =>
      MsgPostContractResponse.decode(new Reader(data))
    );
  }

  Postproof(request: MsgPostproof): Promise<MsgPostproofResponse> {
    const data = MsgPostproof.encode(request).finish();
    const promise = this.rpc.request(
      "jackaldao.canine.storage.Msg",
      "Postproof",
      data
    );
    return promise.then((data) =>
      MsgPostproofResponse.decode(new Reader(data))
    );
  }

  SignContract(request: MsgSignContract): Promise<MsgSignContractResponse> {
    const data = MsgSignContract.encode(request).finish();
    const promise = this.rpc.request(
      "jackaldao.canine.storage.Msg",
      "SignContract",
      data
    );
    return promise.then((data) =>
      MsgSignContractResponse.decode(new Reader(data))
    );
  }

  SetMinerIp(request: MsgSetMinerIp): Promise<MsgSetMinerIpResponse> {
    const data = MsgSetMinerIp.encode(request).finish();
    const promise = this.rpc.request(
      "jackaldao.canine.storage.Msg",
      "SetMinerIp",
      data
    );
    return promise.then((data) =>
      MsgSetMinerIpResponse.decode(new Reader(data))
    );
  }

  SetMinerTotalspace(
    request: MsgSetMinerTotalspace
  ): Promise<MsgSetMinerTotalspaceResponse> {
    const data = MsgSetMinerTotalspace.encode(request).finish();
    const promise = this.rpc.request(
      "jackaldao.canine.storage.Msg",
      "SetMinerTotalspace",
      data
    );
    return promise.then((data) =>
      MsgSetMinerTotalspaceResponse.decode(new Reader(data))
    );
  }

  InitMiner(request: MsgInitMiner): Promise<MsgInitMinerResponse> {
    const data = MsgInitMiner.encode(request).finish();
    const promise = this.rpc.request(
      "jackaldao.canine.storage.Msg",
      "InitMiner",
      data
    );
    return promise.then((data) =>
      MsgInitMinerResponse.decode(new Reader(data))
    );
  }

  CancelContract(
    request: MsgCancelContract
  ): Promise<MsgCancelContractResponse> {
    const data = MsgCancelContract.encode(request).finish();
    const promise = this.rpc.request(
      "jackaldao.canine.storage.Msg",
      "CancelContract",
      data
    );
    return promise.then((data) =>
      MsgCancelContractResponse.decode(new Reader(data))
    );
  }

  BuyStorage(request: MsgBuyStorage): Promise<MsgBuyStorageResponse> {
    const data = MsgBuyStorage.encode(request).finish();
    const promise = this.rpc.request(
      "jackaldao.canine.storage.Msg",
      "BuyStorage",
      data
    );
    return promise.then((data) =>
      MsgBuyStorageResponse.decode(new Reader(data))
    );
  }
}

interface Rpc {
  request(
    service: string,
    method: string,
    data: Uint8Array
  ): Promise<Uint8Array>;
}

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
