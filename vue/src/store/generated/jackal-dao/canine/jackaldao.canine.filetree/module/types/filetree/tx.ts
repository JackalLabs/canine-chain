/* eslint-disable */
import { Reader, util, configure, Writer } from "protobufjs/minimal";
import * as Long from "long";

export const protobufPackage = "jackaldao.canine.filetree";

export interface MsgPostFile {
  creator: string;
  account: string;
  hashParent: string;
  hashChild: string;
  contents: string;
  viewers: string;
  editors: string;
  trackingNumber: number;
}

export interface MsgPostFileResponse {
  path: string;
}

export interface MsgAddViewers {
  creator: string;
  viewerIds: string;
  viewerKeys: string;
  address: string;
  fileowner: string;
}

export interface MsgAddViewersResponse {}

export interface MsgPostkey {
  creator: string;
  key: string;
}

export interface MsgPostkeyResponse {}

export interface MsgInitAccount {
  creator: string;
  account: string;
  rootHashpath: string;
  editors: string;
  key: string;
  trackingNumber: number;
}

export interface MsgInitAccountResponse {
  trackingNumber: number;
}

export interface MsgDeleteFile {
  creator: string;
  hashPath: string;
  account: string;
}

export interface MsgDeleteFileResponse {}

export interface MsgInitAll {
  creator: string;
  name: string;
  pubkey: string;
}

export interface MsgInitAllResponse {}

const baseMsgPostFile: object = {
  creator: "",
  account: "",
  hashParent: "",
  hashChild: "",
  contents: "",
  viewers: "",
  editors: "",
  trackingNumber: 0,
};

export const MsgPostFile = {
  encode(message: MsgPostFile, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.account !== "") {
      writer.uint32(18).string(message.account);
    }
    if (message.hashParent !== "") {
      writer.uint32(26).string(message.hashParent);
    }
    if (message.hashChild !== "") {
      writer.uint32(34).string(message.hashChild);
    }
    if (message.contents !== "") {
      writer.uint32(42).string(message.contents);
    }
    if (message.viewers !== "") {
      writer.uint32(50).string(message.viewers);
    }
    if (message.editors !== "") {
      writer.uint32(58).string(message.editors);
    }
    if (message.trackingNumber !== 0) {
      writer.uint32(64).uint64(message.trackingNumber);
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
          message.account = reader.string();
          break;
        case 3:
          message.hashParent = reader.string();
          break;
        case 4:
          message.hashChild = reader.string();
          break;
        case 5:
          message.contents = reader.string();
          break;
        case 6:
          message.viewers = reader.string();
          break;
        case 7:
          message.editors = reader.string();
          break;
        case 8:
          message.trackingNumber = longToNumber(reader.uint64() as Long);
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
    if (object.account !== undefined && object.account !== null) {
      message.account = String(object.account);
    } else {
      message.account = "";
    }
    if (object.hashParent !== undefined && object.hashParent !== null) {
      message.hashParent = String(object.hashParent);
    } else {
      message.hashParent = "";
    }
    if (object.hashChild !== undefined && object.hashChild !== null) {
      message.hashChild = String(object.hashChild);
    } else {
      message.hashChild = "";
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
    if (object.trackingNumber !== undefined && object.trackingNumber !== null) {
      message.trackingNumber = Number(object.trackingNumber);
    } else {
      message.trackingNumber = 0;
    }
    return message;
  },

  toJSON(message: MsgPostFile): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.account !== undefined && (obj.account = message.account);
    message.hashParent !== undefined && (obj.hashParent = message.hashParent);
    message.hashChild !== undefined && (obj.hashChild = message.hashChild);
    message.contents !== undefined && (obj.contents = message.contents);
    message.viewers !== undefined && (obj.viewers = message.viewers);
    message.editors !== undefined && (obj.editors = message.editors);
    message.trackingNumber !== undefined &&
      (obj.trackingNumber = message.trackingNumber);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgPostFile>): MsgPostFile {
    const message = { ...baseMsgPostFile } as MsgPostFile;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.account !== undefined && object.account !== null) {
      message.account = object.account;
    } else {
      message.account = "";
    }
    if (object.hashParent !== undefined && object.hashParent !== null) {
      message.hashParent = object.hashParent;
    } else {
      message.hashParent = "";
    }
    if (object.hashChild !== undefined && object.hashChild !== null) {
      message.hashChild = object.hashChild;
    } else {
      message.hashChild = "";
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
    if (object.trackingNumber !== undefined && object.trackingNumber !== null) {
      message.trackingNumber = object.trackingNumber;
    } else {
      message.trackingNumber = 0;
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

const baseMsgAddViewers: object = {
  creator: "",
  viewerIds: "",
  viewerKeys: "",
  address: "",
  fileowner: "",
};

export const MsgAddViewers = {
  encode(message: MsgAddViewers, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.viewerIds !== "") {
      writer.uint32(18).string(message.viewerIds);
    }
    if (message.viewerKeys !== "") {
      writer.uint32(26).string(message.viewerKeys);
    }
    if (message.address !== "") {
      writer.uint32(34).string(message.address);
    }
    if (message.fileowner !== "") {
      writer.uint32(42).string(message.fileowner);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgAddViewers {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgAddViewers } as MsgAddViewers;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.viewerIds = reader.string();
          break;
        case 3:
          message.viewerKeys = reader.string();
          break;
        case 4:
          message.address = reader.string();
          break;
        case 5:
          message.fileowner = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgAddViewers {
    const message = { ...baseMsgAddViewers } as MsgAddViewers;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.viewerIds !== undefined && object.viewerIds !== null) {
      message.viewerIds = String(object.viewerIds);
    } else {
      message.viewerIds = "";
    }
    if (object.viewerKeys !== undefined && object.viewerKeys !== null) {
      message.viewerKeys = String(object.viewerKeys);
    } else {
      message.viewerKeys = "";
    }
    if (object.address !== undefined && object.address !== null) {
      message.address = String(object.address);
    } else {
      message.address = "";
    }
    if (object.fileowner !== undefined && object.fileowner !== null) {
      message.fileowner = String(object.fileowner);
    } else {
      message.fileowner = "";
    }
    return message;
  },

  toJSON(message: MsgAddViewers): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.viewerIds !== undefined && (obj.viewerIds = message.viewerIds);
    message.viewerKeys !== undefined && (obj.viewerKeys = message.viewerKeys);
    message.address !== undefined && (obj.address = message.address);
    message.fileowner !== undefined && (obj.fileowner = message.fileowner);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgAddViewers>): MsgAddViewers {
    const message = { ...baseMsgAddViewers } as MsgAddViewers;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.viewerIds !== undefined && object.viewerIds !== null) {
      message.viewerIds = object.viewerIds;
    } else {
      message.viewerIds = "";
    }
    if (object.viewerKeys !== undefined && object.viewerKeys !== null) {
      message.viewerKeys = object.viewerKeys;
    } else {
      message.viewerKeys = "";
    }
    if (object.address !== undefined && object.address !== null) {
      message.address = object.address;
    } else {
      message.address = "";
    }
    if (object.fileowner !== undefined && object.fileowner !== null) {
      message.fileowner = object.fileowner;
    } else {
      message.fileowner = "";
    }
    return message;
  },
};

const baseMsgAddViewersResponse: object = {};

export const MsgAddViewersResponse = {
  encode(_: MsgAddViewersResponse, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgAddViewersResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgAddViewersResponse } as MsgAddViewersResponse;
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

  fromJSON(_: any): MsgAddViewersResponse {
    const message = { ...baseMsgAddViewersResponse } as MsgAddViewersResponse;
    return message;
  },

  toJSON(_: MsgAddViewersResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(_: DeepPartial<MsgAddViewersResponse>): MsgAddViewersResponse {
    const message = { ...baseMsgAddViewersResponse } as MsgAddViewersResponse;
    return message;
  },
};

const baseMsgPostkey: object = { creator: "", key: "" };

export const MsgPostkey = {
  encode(message: MsgPostkey, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.key !== "") {
      writer.uint32(18).string(message.key);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgPostkey {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgPostkey } as MsgPostkey;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.key = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgPostkey {
    const message = { ...baseMsgPostkey } as MsgPostkey;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.key !== undefined && object.key !== null) {
      message.key = String(object.key);
    } else {
      message.key = "";
    }
    return message;
  },

  toJSON(message: MsgPostkey): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.key !== undefined && (obj.key = message.key);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgPostkey>): MsgPostkey {
    const message = { ...baseMsgPostkey } as MsgPostkey;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.key !== undefined && object.key !== null) {
      message.key = object.key;
    } else {
      message.key = "";
    }
    return message;
  },
};

const baseMsgPostkeyResponse: object = {};

export const MsgPostkeyResponse = {
  encode(_: MsgPostkeyResponse, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgPostkeyResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgPostkeyResponse } as MsgPostkeyResponse;
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

  fromJSON(_: any): MsgPostkeyResponse {
    const message = { ...baseMsgPostkeyResponse } as MsgPostkeyResponse;
    return message;
  },

  toJSON(_: MsgPostkeyResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(_: DeepPartial<MsgPostkeyResponse>): MsgPostkeyResponse {
    const message = { ...baseMsgPostkeyResponse } as MsgPostkeyResponse;
    return message;
  },
};

const baseMsgInitAccount: object = {
  creator: "",
  account: "",
  rootHashpath: "",
  editors: "",
  key: "",
  trackingNumber: 0,
};

export const MsgInitAccount = {
  encode(message: MsgInitAccount, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.account !== "") {
      writer.uint32(18).string(message.account);
    }
    if (message.rootHashpath !== "") {
      writer.uint32(26).string(message.rootHashpath);
    }
    if (message.editors !== "") {
      writer.uint32(34).string(message.editors);
    }
    if (message.key !== "") {
      writer.uint32(42).string(message.key);
    }
    if (message.trackingNumber !== 0) {
      writer.uint32(48).uint64(message.trackingNumber);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgInitAccount {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgInitAccount } as MsgInitAccount;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.account = reader.string();
          break;
        case 3:
          message.rootHashpath = reader.string();
          break;
        case 4:
          message.editors = reader.string();
          break;
        case 5:
          message.key = reader.string();
          break;
        case 6:
          message.trackingNumber = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgInitAccount {
    const message = { ...baseMsgInitAccount } as MsgInitAccount;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.account !== undefined && object.account !== null) {
      message.account = String(object.account);
    } else {
      message.account = "";
    }
    if (object.rootHashpath !== undefined && object.rootHashpath !== null) {
      message.rootHashpath = String(object.rootHashpath);
    } else {
      message.rootHashpath = "";
    }
    if (object.editors !== undefined && object.editors !== null) {
      message.editors = String(object.editors);
    } else {
      message.editors = "";
    }
    if (object.key !== undefined && object.key !== null) {
      message.key = String(object.key);
    } else {
      message.key = "";
    }
    if (object.trackingNumber !== undefined && object.trackingNumber !== null) {
      message.trackingNumber = Number(object.trackingNumber);
    } else {
      message.trackingNumber = 0;
    }
    return message;
  },

  toJSON(message: MsgInitAccount): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.account !== undefined && (obj.account = message.account);
    message.rootHashpath !== undefined &&
      (obj.rootHashpath = message.rootHashpath);
    message.editors !== undefined && (obj.editors = message.editors);
    message.key !== undefined && (obj.key = message.key);
    message.trackingNumber !== undefined &&
      (obj.trackingNumber = message.trackingNumber);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgInitAccount>): MsgInitAccount {
    const message = { ...baseMsgInitAccount } as MsgInitAccount;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.account !== undefined && object.account !== null) {
      message.account = object.account;
    } else {
      message.account = "";
    }
    if (object.rootHashpath !== undefined && object.rootHashpath !== null) {
      message.rootHashpath = object.rootHashpath;
    } else {
      message.rootHashpath = "";
    }
    if (object.editors !== undefined && object.editors !== null) {
      message.editors = object.editors;
    } else {
      message.editors = "";
    }
    if (object.key !== undefined && object.key !== null) {
      message.key = object.key;
    } else {
      message.key = "";
    }
    if (object.trackingNumber !== undefined && object.trackingNumber !== null) {
      message.trackingNumber = object.trackingNumber;
    } else {
      message.trackingNumber = 0;
    }
    return message;
  },
};

const baseMsgInitAccountResponse: object = { trackingNumber: 0 };

export const MsgInitAccountResponse = {
  encode(
    message: MsgInitAccountResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.trackingNumber !== 0) {
      writer.uint32(8).uint64(message.trackingNumber);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgInitAccountResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgInitAccountResponse } as MsgInitAccountResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.trackingNumber = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgInitAccountResponse {
    const message = { ...baseMsgInitAccountResponse } as MsgInitAccountResponse;
    if (object.trackingNumber !== undefined && object.trackingNumber !== null) {
      message.trackingNumber = Number(object.trackingNumber);
    } else {
      message.trackingNumber = 0;
    }
    return message;
  },

  toJSON(message: MsgInitAccountResponse): unknown {
    const obj: any = {};
    message.trackingNumber !== undefined &&
      (obj.trackingNumber = message.trackingNumber);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgInitAccountResponse>
  ): MsgInitAccountResponse {
    const message = { ...baseMsgInitAccountResponse } as MsgInitAccountResponse;
    if (object.trackingNumber !== undefined && object.trackingNumber !== null) {
      message.trackingNumber = object.trackingNumber;
    } else {
      message.trackingNumber = 0;
    }
    return message;
  },
};

const baseMsgDeleteFile: object = { creator: "", hashPath: "", account: "" };

export const MsgDeleteFile = {
  encode(message: MsgDeleteFile, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.hashPath !== "") {
      writer.uint32(18).string(message.hashPath);
    }
    if (message.account !== "") {
      writer.uint32(26).string(message.account);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgDeleteFile {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgDeleteFile } as MsgDeleteFile;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.hashPath = reader.string();
          break;
        case 3:
          message.account = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgDeleteFile {
    const message = { ...baseMsgDeleteFile } as MsgDeleteFile;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.hashPath !== undefined && object.hashPath !== null) {
      message.hashPath = String(object.hashPath);
    } else {
      message.hashPath = "";
    }
    if (object.account !== undefined && object.account !== null) {
      message.account = String(object.account);
    } else {
      message.account = "";
    }
    return message;
  },

  toJSON(message: MsgDeleteFile): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.hashPath !== undefined && (obj.hashPath = message.hashPath);
    message.account !== undefined && (obj.account = message.account);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgDeleteFile>): MsgDeleteFile {
    const message = { ...baseMsgDeleteFile } as MsgDeleteFile;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.hashPath !== undefined && object.hashPath !== null) {
      message.hashPath = object.hashPath;
    } else {
      message.hashPath = "";
    }
    if (object.account !== undefined && object.account !== null) {
      message.account = object.account;
    } else {
      message.account = "";
    }
    return message;
  },
};

const baseMsgDeleteFileResponse: object = {};

export const MsgDeleteFileResponse = {
  encode(_: MsgDeleteFileResponse, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgDeleteFileResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgDeleteFileResponse } as MsgDeleteFileResponse;
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

  fromJSON(_: any): MsgDeleteFileResponse {
    const message = { ...baseMsgDeleteFileResponse } as MsgDeleteFileResponse;
    return message;
  },

  toJSON(_: MsgDeleteFileResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(_: DeepPartial<MsgDeleteFileResponse>): MsgDeleteFileResponse {
    const message = { ...baseMsgDeleteFileResponse } as MsgDeleteFileResponse;
    return message;
  },
};

const baseMsgInitAll: object = { creator: "", name: "", pubkey: "" };

export const MsgInitAll = {
  encode(message: MsgInitAll, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.name !== "") {
      writer.uint32(18).string(message.name);
    }
    if (message.pubkey !== "") {
      writer.uint32(26).string(message.pubkey);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgInitAll {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgInitAll } as MsgInitAll;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.name = reader.string();
          break;
        case 3:
          message.pubkey = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgInitAll {
    const message = { ...baseMsgInitAll } as MsgInitAll;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.name !== undefined && object.name !== null) {
      message.name = String(object.name);
    } else {
      message.name = "";
    }
    if (object.pubkey !== undefined && object.pubkey !== null) {
      message.pubkey = String(object.pubkey);
    } else {
      message.pubkey = "";
    }
    return message;
  },

  toJSON(message: MsgInitAll): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.name !== undefined && (obj.name = message.name);
    message.pubkey !== undefined && (obj.pubkey = message.pubkey);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgInitAll>): MsgInitAll {
    const message = { ...baseMsgInitAll } as MsgInitAll;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.name !== undefined && object.name !== null) {
      message.name = object.name;
    } else {
      message.name = "";
    }
    if (object.pubkey !== undefined && object.pubkey !== null) {
      message.pubkey = object.pubkey;
    } else {
      message.pubkey = "";
    }
    return message;
  },
};

const baseMsgInitAllResponse: object = {};

export const MsgInitAllResponse = {
  encode(_: MsgInitAllResponse, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgInitAllResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgInitAllResponse } as MsgInitAllResponse;
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

  fromJSON(_: any): MsgInitAllResponse {
    const message = { ...baseMsgInitAllResponse } as MsgInitAllResponse;
    return message;
  },

  toJSON(_: MsgInitAllResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(_: DeepPartial<MsgInitAllResponse>): MsgInitAllResponse {
    const message = { ...baseMsgInitAllResponse } as MsgInitAllResponse;
    return message;
  },
};

/** Msg defines the Msg service. */
export interface Msg {
  PostFile(request: MsgPostFile): Promise<MsgPostFileResponse>;
  AddViewers(request: MsgAddViewers): Promise<MsgAddViewersResponse>;
  Postkey(request: MsgPostkey): Promise<MsgPostkeyResponse>;
  InitAccount(request: MsgInitAccount): Promise<MsgInitAccountResponse>;
  DeleteFile(request: MsgDeleteFile): Promise<MsgDeleteFileResponse>;
  /** this line is used by starport scaffolding # proto/tx/rpc */
  InitAll(request: MsgInitAll): Promise<MsgInitAllResponse>;
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

  AddViewers(request: MsgAddViewers): Promise<MsgAddViewersResponse> {
    const data = MsgAddViewers.encode(request).finish();
    const promise = this.rpc.request(
      "jackaldao.canine.filetree.Msg",
      "AddViewers",
      data
    );
    return promise.then((data) =>
      MsgAddViewersResponse.decode(new Reader(data))
    );
  }

  Postkey(request: MsgPostkey): Promise<MsgPostkeyResponse> {
    const data = MsgPostkey.encode(request).finish();
    const promise = this.rpc.request(
      "jackaldao.canine.filetree.Msg",
      "Postkey",
      data
    );
    return promise.then((data) => MsgPostkeyResponse.decode(new Reader(data)));
  }

  InitAccount(request: MsgInitAccount): Promise<MsgInitAccountResponse> {
    const data = MsgInitAccount.encode(request).finish();
    const promise = this.rpc.request(
      "jackaldao.canine.filetree.Msg",
      "InitAccount",
      data
    );
    return promise.then((data) =>
      MsgInitAccountResponse.decode(new Reader(data))
    );
  }

  DeleteFile(request: MsgDeleteFile): Promise<MsgDeleteFileResponse> {
    const data = MsgDeleteFile.encode(request).finish();
    const promise = this.rpc.request(
      "jackaldao.canine.filetree.Msg",
      "DeleteFile",
      data
    );
    return promise.then((data) =>
      MsgDeleteFileResponse.decode(new Reader(data))
    );
  }

  InitAll(request: MsgInitAll): Promise<MsgInitAllResponse> {
    const data = MsgInitAll.encode(request).finish();
    const promise = this.rpc.request(
      "jackaldao.canine.filetree.Msg",
      "InitAll",
      data
    );
    return promise.then((data) => MsgInitAllResponse.decode(new Reader(data)));
  }
}

interface Rpc {
  request(
    service: string,
    method: string,
    data: Uint8Array
  ): Promise<Uint8Array>;
}

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
