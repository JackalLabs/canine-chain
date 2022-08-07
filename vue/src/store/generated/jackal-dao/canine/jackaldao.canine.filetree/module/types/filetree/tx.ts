/* eslint-disable */
import { Reader, Writer } from "protobufjs/minimal";

export const protobufPackage = "jackaldao.canine.filetree";

export interface MsgPostFile {
  creator: string;
  hashpath: string;
  contents: string;
  viewers: string;
  editors: string;
}

export interface MsgPostFileResponse {
  path: string;
}

const baseMsgPostFile: object = {
  creator: "",
  hashpath: "",
  contents: "",
  viewers: "",
  editors: "",
};

export const MsgPostFile = {
  encode(message: MsgPostFile, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.hashpath !== "") {
      writer.uint32(18).string(message.hashpath);
    }
    if (message.contents !== "") {
      writer.uint32(26).string(message.contents);
    }
    if (message.viewers !== "") {
      writer.uint32(34).string(message.viewers);
    }
    if (message.editors !== "") {
      writer.uint32(42).string(message.editors);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgPostFile {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgPostFile } as MsgPostFile;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.hashpath = reader.string();
          break;
        case 3:
          message.contents = reader.string();
          break;
        case 4:
          message.viewers = reader.string();
          break;
        case 5:
          message.editors = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgPostFile {
    const message = { ...baseMsgPostFile } as MsgPostFile;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.hashpath !== undefined && object.hashpath !== null) {
      message.hashpath = String(object.hashpath);
    } else {
      message.hashpath = "";
    }
    if (object.contents !== undefined && object.contents !== null) {
      message.contents = String(object.contents);
    } else {
      message.contents = "";
    }
    if (object.viewers !== undefined && object.viewers !== null) {
      message.viewers = String(object.viewers);
    } else {
      message.viewers = "";
    }
    if (object.editors !== undefined && object.editors !== null) {
      message.editors = String(object.editors);
    } else {
      message.editors = "";
    }
    return message;
  },

  toJSON(message: MsgPostFile): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.hashpath !== undefined && (obj.hashpath = message.hashpath);
    message.contents !== undefined && (obj.contents = message.contents);
    message.viewers !== undefined && (obj.viewers = message.viewers);
    message.editors !== undefined && (obj.editors = message.editors);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgPostFile>): MsgPostFile {
    const message = { ...baseMsgPostFile } as MsgPostFile;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.hashpath !== undefined && object.hashpath !== null) {
      message.hashpath = object.hashpath;
    } else {
      message.hashpath = "";
    }
    if (object.contents !== undefined && object.contents !== null) {
      message.contents = object.contents;
    } else {
      message.contents = "";
    }
    if (object.viewers !== undefined && object.viewers !== null) {
      message.viewers = object.viewers;
    } else {
      message.viewers = "";
    }
    if (object.editors !== undefined && object.editors !== null) {
      message.editors = object.editors;
    } else {
      message.editors = "";
    }
    return message;
  },
};

const baseMsgPostFileResponse: object = { path: "" };

export const MsgPostFileResponse = {
  encode(
    message: MsgPostFileResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.path !== "") {
      writer.uint32(10).string(message.path);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgPostFileResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgPostFileResponse } as MsgPostFileResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.path = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgPostFileResponse {
    const message = { ...baseMsgPostFileResponse } as MsgPostFileResponse;
    if (object.path !== undefined && object.path !== null) {
      message.path = String(object.path);
    } else {
      message.path = "";
    }
    return message;
  },

  toJSON(message: MsgPostFileResponse): unknown {
    const obj: any = {};
    message.path !== undefined && (obj.path = message.path);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgPostFileResponse>): MsgPostFileResponse {
    const message = { ...baseMsgPostFileResponse } as MsgPostFileResponse;
    if (object.path !== undefined && object.path !== null) {
      message.path = object.path;
    } else {
      message.path = "";
    }
    return message;
  },
};

/** Msg defines the Msg service. */
export interface Msg {
  /** this line is used by starport scaffolding # proto/tx/rpc */
  PostFile(request: MsgPostFile): Promise<MsgPostFileResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  PostFile(request: MsgPostFile): Promise<MsgPostFileResponse> {
    const data = MsgPostFile.encode(request).finish();
    const promise = this.rpc.request(
      "jackaldao.canine.filetree.Msg",
      "PostFile",
      data
    );
    return promise.then((data) => MsgPostFileResponse.decode(new Reader(data)));
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
