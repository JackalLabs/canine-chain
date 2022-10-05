/* eslint-disable */
import { Reader, Writer } from "protobufjs/minimal";

export const protobufPackage = "jackaldao.canine.dsig";

export interface MsgUploadfile {
  creator: string;
  fid: string;
}

export interface MsgUploadfileResponse {}

export interface MsgCreateform {
  creator: string;
  fid: string;
  signees: string;
}

export interface MsgCreateformResponse {
  ffid: string;
}

export interface MsgSignform {
  creator: string;
  ffid: string;
  vote: number;
}

export interface MsgSignformResponse {}

const baseMsgUploadfile: object = { creator: "", fid: "" };

export const MsgUploadfile = {
  encode(message: MsgUploadfile, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.fid !== "") {
      writer.uint32(18).string(message.fid);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgUploadfile {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgUploadfile } as MsgUploadfile;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.fid = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgUploadfile {
    const message = { ...baseMsgUploadfile } as MsgUploadfile;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.fid !== undefined && object.fid !== null) {
      message.fid = String(object.fid);
    } else {
      message.fid = "";
    }
    return message;
  },

  toJSON(message: MsgUploadfile): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.fid !== undefined && (obj.fid = message.fid);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgUploadfile>): MsgUploadfile {
    const message = { ...baseMsgUploadfile } as MsgUploadfile;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.fid !== undefined && object.fid !== null) {
      message.fid = object.fid;
    } else {
      message.fid = "";
    }
    return message;
  },
};

const baseMsgUploadfileResponse: object = {};

export const MsgUploadfileResponse = {
  encode(_: MsgUploadfileResponse, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgUploadfileResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgUploadfileResponse } as MsgUploadfileResponse;
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

  fromJSON(_: any): MsgUploadfileResponse {
    const message = { ...baseMsgUploadfileResponse } as MsgUploadfileResponse;
    return message;
  },

  toJSON(_: MsgUploadfileResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(_: DeepPartial<MsgUploadfileResponse>): MsgUploadfileResponse {
    const message = { ...baseMsgUploadfileResponse } as MsgUploadfileResponse;
    return message;
  },
};

const baseMsgCreateform: object = { creator: "", fid: "", signees: "" };

export const MsgCreateform = {
  encode(message: MsgCreateform, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.fid !== "") {
      writer.uint32(18).string(message.fid);
    }
    if (message.signees !== "") {
      writer.uint32(26).string(message.signees);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgCreateform {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgCreateform } as MsgCreateform;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.fid = reader.string();
          break;
        case 3:
          message.signees = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgCreateform {
    const message = { ...baseMsgCreateform } as MsgCreateform;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
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

  toJSON(message: MsgCreateform): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.fid !== undefined && (obj.fid = message.fid);
    message.signees !== undefined && (obj.signees = message.signees);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgCreateform>): MsgCreateform {
    const message = { ...baseMsgCreateform } as MsgCreateform;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
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

const baseMsgCreateformResponse: object = { ffid: "" };

export const MsgCreateformResponse = {
  encode(
    message: MsgCreateformResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.ffid !== "") {
      writer.uint32(10).string(message.ffid);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgCreateformResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgCreateformResponse } as MsgCreateformResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.ffid = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgCreateformResponse {
    const message = { ...baseMsgCreateformResponse } as MsgCreateformResponse;
    if (object.ffid !== undefined && object.ffid !== null) {
      message.ffid = String(object.ffid);
    } else {
      message.ffid = "";
    }
    return message;
  },

  toJSON(message: MsgCreateformResponse): unknown {
    const obj: any = {};
    message.ffid !== undefined && (obj.ffid = message.ffid);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgCreateformResponse>
  ): MsgCreateformResponse {
    const message = { ...baseMsgCreateformResponse } as MsgCreateformResponse;
    if (object.ffid !== undefined && object.ffid !== null) {
      message.ffid = object.ffid;
    } else {
      message.ffid = "";
    }
    return message;
  },
};

const baseMsgSignform: object = { creator: "", ffid: "", vote: 0 };

export const MsgSignform = {
  encode(message: MsgSignform, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.ffid !== "") {
      writer.uint32(18).string(message.ffid);
    }
    if (message.vote !== 0) {
      writer.uint32(24).int32(message.vote);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgSignform {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgSignform } as MsgSignform;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.ffid = reader.string();
          break;
        case 3:
          message.vote = reader.int32();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgSignform {
    const message = { ...baseMsgSignform } as MsgSignform;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.ffid !== undefined && object.ffid !== null) {
      message.ffid = String(object.ffid);
    } else {
      message.ffid = "";
    }
    if (object.vote !== undefined && object.vote !== null) {
      message.vote = Number(object.vote);
    } else {
      message.vote = 0;
    }
    return message;
  },

  toJSON(message: MsgSignform): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.ffid !== undefined && (obj.ffid = message.ffid);
    message.vote !== undefined && (obj.vote = message.vote);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgSignform>): MsgSignform {
    const message = { ...baseMsgSignform } as MsgSignform;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.ffid !== undefined && object.ffid !== null) {
      message.ffid = object.ffid;
    } else {
      message.ffid = "";
    }
    if (object.vote !== undefined && object.vote !== null) {
      message.vote = object.vote;
    } else {
      message.vote = 0;
    }
    return message;
  },
};

const baseMsgSignformResponse: object = {};

export const MsgSignformResponse = {
  encode(_: MsgSignformResponse, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgSignformResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgSignformResponse } as MsgSignformResponse;
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

  fromJSON(_: any): MsgSignformResponse {
    const message = { ...baseMsgSignformResponse } as MsgSignformResponse;
    return message;
  },

  toJSON(_: MsgSignformResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(_: DeepPartial<MsgSignformResponse>): MsgSignformResponse {
    const message = { ...baseMsgSignformResponse } as MsgSignformResponse;
    return message;
  },
};

/** Msg defines the Msg service. */
export interface Msg {
  Uploadfile(request: MsgUploadfile): Promise<MsgUploadfileResponse>;
  Createform(request: MsgCreateform): Promise<MsgCreateformResponse>;
  /** this line is used by starport scaffolding # proto/tx/rpc */
  Signform(request: MsgSignform): Promise<MsgSignformResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  Uploadfile(request: MsgUploadfile): Promise<MsgUploadfileResponse> {
    const data = MsgUploadfile.encode(request).finish();
    const promise = this.rpc.request(
      "jackaldao.canine.dsig.Msg",
      "Uploadfile",
      data
    );
    return promise.then((data) =>
      MsgUploadfileResponse.decode(new Reader(data))
    );
  }

  Createform(request: MsgCreateform): Promise<MsgCreateformResponse> {
    const data = MsgCreateform.encode(request).finish();
    const promise = this.rpc.request(
      "jackaldao.canine.dsig.Msg",
      "Createform",
      data
    );
    return promise.then((data) =>
      MsgCreateformResponse.decode(new Reader(data))
    );
  }

  Signform(request: MsgSignform): Promise<MsgSignformResponse> {
    const data = MsgSignform.encode(request).finish();
    const promise = this.rpc.request(
      "jackaldao.canine.dsig.Msg",
      "Signform",
      data
    );
    return promise.then((data) => MsgSignformResponse.decode(new Reader(data)));
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
