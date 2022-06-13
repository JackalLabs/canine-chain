/* eslint-disable */
import { Reader, Writer } from "protobufjs/minimal";

export const protobufPackage = "jackaldao.canine.jklmining";

export interface MsgAllowSave {
  creator: string;
  passkey: string;
  size: string;
}

export interface MsgAllowSaveResponse {}

export interface MsgCreateSaveRequests {
  creator: string;
  index: string;
  size: string;
  approved: string;
}

export interface MsgCreateSaveRequestsResponse {}

export interface MsgUpdateSaveRequests {
  creator: string;
  index: string;
  size: string;
  approved: string;
}

export interface MsgUpdateSaveRequestsResponse {}

export interface MsgDeleteSaveRequests {
  creator: string;
  index: string;
}

export interface MsgDeleteSaveRequestsResponse {}

const baseMsgAllowSave: object = { creator: "", passkey: "", size: "" };

export const MsgAllowSave = {
  encode(message: MsgAllowSave, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.passkey !== "") {
      writer.uint32(18).string(message.passkey);
    }
    if (message.size !== "") {
      writer.uint32(26).string(message.size);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgAllowSave {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgAllowSave } as MsgAllowSave;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.passkey = reader.string();
          break;
        case 3:
          message.size = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgAllowSave {
    const message = { ...baseMsgAllowSave } as MsgAllowSave;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.passkey !== undefined && object.passkey !== null) {
      message.passkey = String(object.passkey);
    } else {
      message.passkey = "";
    }
    if (object.size !== undefined && object.size !== null) {
      message.size = String(object.size);
    } else {
      message.size = "";
    }
    return message;
  },

  toJSON(message: MsgAllowSave): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.passkey !== undefined && (obj.passkey = message.passkey);
    message.size !== undefined && (obj.size = message.size);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgAllowSave>): MsgAllowSave {
    const message = { ...baseMsgAllowSave } as MsgAllowSave;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.passkey !== undefined && object.passkey !== null) {
      message.passkey = object.passkey;
    } else {
      message.passkey = "";
    }
    if (object.size !== undefined && object.size !== null) {
      message.size = object.size;
    } else {
      message.size = "";
    }
    return message;
  },
};

const baseMsgAllowSaveResponse: object = {};

export const MsgAllowSaveResponse = {
  encode(_: MsgAllowSaveResponse, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgAllowSaveResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgAllowSaveResponse } as MsgAllowSaveResponse;
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

  fromJSON(_: any): MsgAllowSaveResponse {
    const message = { ...baseMsgAllowSaveResponse } as MsgAllowSaveResponse;
    return message;
  },

  toJSON(_: MsgAllowSaveResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(_: DeepPartial<MsgAllowSaveResponse>): MsgAllowSaveResponse {
    const message = { ...baseMsgAllowSaveResponse } as MsgAllowSaveResponse;
    return message;
  },
};

const baseMsgCreateSaveRequests: object = {
  creator: "",
  index: "",
  size: "",
  approved: "",
};

export const MsgCreateSaveRequests = {
  encode(
    message: MsgCreateSaveRequests,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.index !== "") {
      writer.uint32(18).string(message.index);
    }
    if (message.size !== "") {
      writer.uint32(26).string(message.size);
    }
    if (message.approved !== "") {
      writer.uint32(34).string(message.approved);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgCreateSaveRequests {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgCreateSaveRequests } as MsgCreateSaveRequests;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.index = reader.string();
          break;
        case 3:
          message.size = reader.string();
          break;
        case 4:
          message.approved = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgCreateSaveRequests {
    const message = { ...baseMsgCreateSaveRequests } as MsgCreateSaveRequests;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
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
    return message;
  },

  toJSON(message: MsgCreateSaveRequests): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.index !== undefined && (obj.index = message.index);
    message.size !== undefined && (obj.size = message.size);
    message.approved !== undefined && (obj.approved = message.approved);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgCreateSaveRequests>
  ): MsgCreateSaveRequests {
    const message = { ...baseMsgCreateSaveRequests } as MsgCreateSaveRequests;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
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
    return message;
  },
};

const baseMsgCreateSaveRequestsResponse: object = {};

export const MsgCreateSaveRequestsResponse = {
  encode(
    _: MsgCreateSaveRequestsResponse,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgCreateSaveRequestsResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgCreateSaveRequestsResponse,
    } as MsgCreateSaveRequestsResponse;
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

  fromJSON(_: any): MsgCreateSaveRequestsResponse {
    const message = {
      ...baseMsgCreateSaveRequestsResponse,
    } as MsgCreateSaveRequestsResponse;
    return message;
  },

  toJSON(_: MsgCreateSaveRequestsResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgCreateSaveRequestsResponse>
  ): MsgCreateSaveRequestsResponse {
    const message = {
      ...baseMsgCreateSaveRequestsResponse,
    } as MsgCreateSaveRequestsResponse;
    return message;
  },
};

const baseMsgUpdateSaveRequests: object = {
  creator: "",
  index: "",
  size: "",
  approved: "",
};

export const MsgUpdateSaveRequests = {
  encode(
    message: MsgUpdateSaveRequests,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.index !== "") {
      writer.uint32(18).string(message.index);
    }
    if (message.size !== "") {
      writer.uint32(26).string(message.size);
    }
    if (message.approved !== "") {
      writer.uint32(34).string(message.approved);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgUpdateSaveRequests {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgUpdateSaveRequests } as MsgUpdateSaveRequests;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.index = reader.string();
          break;
        case 3:
          message.size = reader.string();
          break;
        case 4:
          message.approved = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgUpdateSaveRequests {
    const message = { ...baseMsgUpdateSaveRequests } as MsgUpdateSaveRequests;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
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
    return message;
  },

  toJSON(message: MsgUpdateSaveRequests): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.index !== undefined && (obj.index = message.index);
    message.size !== undefined && (obj.size = message.size);
    message.approved !== undefined && (obj.approved = message.approved);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgUpdateSaveRequests>
  ): MsgUpdateSaveRequests {
    const message = { ...baseMsgUpdateSaveRequests } as MsgUpdateSaveRequests;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
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
    return message;
  },
};

const baseMsgUpdateSaveRequestsResponse: object = {};

export const MsgUpdateSaveRequestsResponse = {
  encode(
    _: MsgUpdateSaveRequestsResponse,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgUpdateSaveRequestsResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgUpdateSaveRequestsResponse,
    } as MsgUpdateSaveRequestsResponse;
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

  fromJSON(_: any): MsgUpdateSaveRequestsResponse {
    const message = {
      ...baseMsgUpdateSaveRequestsResponse,
    } as MsgUpdateSaveRequestsResponse;
    return message;
  },

  toJSON(_: MsgUpdateSaveRequestsResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgUpdateSaveRequestsResponse>
  ): MsgUpdateSaveRequestsResponse {
    const message = {
      ...baseMsgUpdateSaveRequestsResponse,
    } as MsgUpdateSaveRequestsResponse;
    return message;
  },
};

const baseMsgDeleteSaveRequests: object = { creator: "", index: "" };

export const MsgDeleteSaveRequests = {
  encode(
    message: MsgDeleteSaveRequests,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.index !== "") {
      writer.uint32(18).string(message.index);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgDeleteSaveRequests {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgDeleteSaveRequests } as MsgDeleteSaveRequests;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.index = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgDeleteSaveRequests {
    const message = { ...baseMsgDeleteSaveRequests } as MsgDeleteSaveRequests;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.index !== undefined && object.index !== null) {
      message.index = String(object.index);
    } else {
      message.index = "";
    }
    return message;
  },

  toJSON(message: MsgDeleteSaveRequests): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.index !== undefined && (obj.index = message.index);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgDeleteSaveRequests>
  ): MsgDeleteSaveRequests {
    const message = { ...baseMsgDeleteSaveRequests } as MsgDeleteSaveRequests;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.index !== undefined && object.index !== null) {
      message.index = object.index;
    } else {
      message.index = "";
    }
    return message;
  },
};

const baseMsgDeleteSaveRequestsResponse: object = {};

export const MsgDeleteSaveRequestsResponse = {
  encode(
    _: MsgDeleteSaveRequestsResponse,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgDeleteSaveRequestsResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgDeleteSaveRequestsResponse,
    } as MsgDeleteSaveRequestsResponse;
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

  fromJSON(_: any): MsgDeleteSaveRequestsResponse {
    const message = {
      ...baseMsgDeleteSaveRequestsResponse,
    } as MsgDeleteSaveRequestsResponse;
    return message;
  },

  toJSON(_: MsgDeleteSaveRequestsResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgDeleteSaveRequestsResponse>
  ): MsgDeleteSaveRequestsResponse {
    const message = {
      ...baseMsgDeleteSaveRequestsResponse,
    } as MsgDeleteSaveRequestsResponse;
    return message;
  },
};

/** Msg defines the Msg service. */
export interface Msg {
  AllowSave(request: MsgAllowSave): Promise<MsgAllowSaveResponse>;
  CreateSaveRequests(
    request: MsgCreateSaveRequests
  ): Promise<MsgCreateSaveRequestsResponse>;
  UpdateSaveRequests(
    request: MsgUpdateSaveRequests
  ): Promise<MsgUpdateSaveRequestsResponse>;
  /** this line is used by starport scaffolding # proto/tx/rpc */
  DeleteSaveRequests(
    request: MsgDeleteSaveRequests
  ): Promise<MsgDeleteSaveRequestsResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  AllowSave(request: MsgAllowSave): Promise<MsgAllowSaveResponse> {
    const data = MsgAllowSave.encode(request).finish();
    const promise = this.rpc.request(
      "jackaldao.canine.jklmining.Msg",
      "AllowSave",
      data
    );
    return promise.then((data) =>
      MsgAllowSaveResponse.decode(new Reader(data))
    );
  }

  CreateSaveRequests(
    request: MsgCreateSaveRequests
  ): Promise<MsgCreateSaveRequestsResponse> {
    const data = MsgCreateSaveRequests.encode(request).finish();
    const promise = this.rpc.request(
      "jackaldao.canine.jklmining.Msg",
      "CreateSaveRequests",
      data
    );
    return promise.then((data) =>
      MsgCreateSaveRequestsResponse.decode(new Reader(data))
    );
  }

  UpdateSaveRequests(
    request: MsgUpdateSaveRequests
  ): Promise<MsgUpdateSaveRequestsResponse> {
    const data = MsgUpdateSaveRequests.encode(request).finish();
    const promise = this.rpc.request(
      "jackaldao.canine.jklmining.Msg",
      "UpdateSaveRequests",
      data
    );
    return promise.then((data) =>
      MsgUpdateSaveRequestsResponse.decode(new Reader(data))
    );
  }

  DeleteSaveRequests(
    request: MsgDeleteSaveRequests
  ): Promise<MsgDeleteSaveRequestsResponse> {
    const data = MsgDeleteSaveRequests.encode(request).finish();
    const promise = this.rpc.request(
      "jackaldao.canine.jklmining.Msg",
      "DeleteSaveRequests",
      data
    );
    return promise.then((data) =>
      MsgDeleteSaveRequestsResponse.decode(new Reader(data))
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
