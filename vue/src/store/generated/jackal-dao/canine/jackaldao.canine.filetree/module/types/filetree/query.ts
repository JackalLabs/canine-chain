/* eslint-disable */
import { Reader, Writer } from "protobufjs/minimal";
import { Params } from "../filetree/params";
import { Files } from "../filetree/files";
import {
  PageRequest,
  PageResponse,
} from "../cosmos/base/query/v1beta1/pagination";

export const protobufPackage = "jackaldao.canine.filetree";

/** QueryParamsRequest is request type for the Query/Params RPC method. */
export interface QueryParamsRequest {}

/** QueryParamsResponse is response type for the Query/Params RPC method. */
export interface QueryParamsResponse {
  /** params holds all the parameters of this module. */
  params: Params | undefined;
}

export interface QueryEncryptRequest {
  address: string;
  message: string;
}

export interface QueryEncryptResponse {
  encryptionData: string;
}

export interface QueryDecryptRequest {
  message: string;
}

export interface QueryDecryptResponse {
  data: string;
}

export interface QueryGetFilesRequest {
  address: string;
}

export interface QueryGetFilesResponse {
  files: Files | undefined;
}

export interface QueryAllFilesRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllFilesResponse {
  files: Files[];
  pagination: PageResponse | undefined;
}

export interface QueryGetKeysRequest {
  hashpath: string;
}

export interface QueryGetKeysResponse {
  keys: string;
}

export interface QueryGetKeyRequest {
  filepath: string;
  owner: string;
}

const baseQueryParamsRequest: object = {};

export const QueryParamsRequest = {
  encode(_: QueryParamsRequest, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryParamsRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryParamsRequest } as QueryParamsRequest;
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

  fromJSON(_: any): QueryParamsRequest {
    const message = { ...baseQueryParamsRequest } as QueryParamsRequest;
    return message;
  },

  toJSON(_: QueryParamsRequest): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(_: DeepPartial<QueryParamsRequest>): QueryParamsRequest {
    const message = { ...baseQueryParamsRequest } as QueryParamsRequest;
    return message;
  },
};

const baseQueryParamsResponse: object = {};

export const QueryParamsResponse = {
  encode(
    message: QueryParamsResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.params !== undefined) {
      Params.encode(message.params, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryParamsResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryParamsResponse } as QueryParamsResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.params = Params.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryParamsResponse {
    const message = { ...baseQueryParamsResponse } as QueryParamsResponse;
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromJSON(object.params);
    } else {
      message.params = undefined;
    }
    return message;
  },

  toJSON(message: QueryParamsResponse): unknown {
    const obj: any = {};
    message.params !== undefined &&
      (obj.params = message.params ? Params.toJSON(message.params) : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<QueryParamsResponse>): QueryParamsResponse {
    const message = { ...baseQueryParamsResponse } as QueryParamsResponse;
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromPartial(object.params);
    } else {
      message.params = undefined;
    }
    return message;
  },
};

const baseQueryEncryptRequest: object = { address: "", message: "" };

export const QueryEncryptRequest = {
  encode(
    message: QueryEncryptRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.address !== "") {
      writer.uint32(10).string(message.address);
    }
    if (message.message !== "") {
      writer.uint32(18).string(message.message);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryEncryptRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryEncryptRequest } as QueryEncryptRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.address = reader.string();
          break;
        case 2:
          message.message = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryEncryptRequest {
    const message = { ...baseQueryEncryptRequest } as QueryEncryptRequest;
    if (object.address !== undefined && object.address !== null) {
      message.address = String(object.address);
    } else {
      message.address = "";
    }
    if (object.message !== undefined && object.message !== null) {
      message.message = String(object.message);
    } else {
      message.message = "";
    }
    return message;
  },

  toJSON(message: QueryEncryptRequest): unknown {
    const obj: any = {};
    message.address !== undefined && (obj.address = message.address);
    message.message !== undefined && (obj.message = message.message);
    return obj;
  },

  fromPartial(object: DeepPartial<QueryEncryptRequest>): QueryEncryptRequest {
    const message = { ...baseQueryEncryptRequest } as QueryEncryptRequest;
    if (object.address !== undefined && object.address !== null) {
      message.address = object.address;
    } else {
      message.address = "";
    }
    if (object.message !== undefined && object.message !== null) {
      message.message = object.message;
    } else {
      message.message = "";
    }
    return message;
  },
};

const baseQueryEncryptResponse: object = { encryptionData: "" };

export const QueryEncryptResponse = {
  encode(
    message: QueryEncryptResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.encryptionData !== "") {
      writer.uint32(10).string(message.encryptionData);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryEncryptResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryEncryptResponse } as QueryEncryptResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.encryptionData = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryEncryptResponse {
    const message = { ...baseQueryEncryptResponse } as QueryEncryptResponse;
    if (object.encryptionData !== undefined && object.encryptionData !== null) {
      message.encryptionData = String(object.encryptionData);
    } else {
      message.encryptionData = "";
    }
    return message;
  },

  toJSON(message: QueryEncryptResponse): unknown {
    const obj: any = {};
    message.encryptionData !== undefined &&
      (obj.encryptionData = message.encryptionData);
    return obj;
  },

  fromPartial(object: DeepPartial<QueryEncryptResponse>): QueryEncryptResponse {
    const message = { ...baseQueryEncryptResponse } as QueryEncryptResponse;
    if (object.encryptionData !== undefined && object.encryptionData !== null) {
      message.encryptionData = object.encryptionData;
    } else {
      message.encryptionData = "";
    }
    return message;
  },
};

const baseQueryDecryptRequest: object = { message: "" };

export const QueryDecryptRequest = {
  encode(
    message: QueryDecryptRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.message !== "") {
      writer.uint32(10).string(message.message);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryDecryptRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryDecryptRequest } as QueryDecryptRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.message = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryDecryptRequest {
    const message = { ...baseQueryDecryptRequest } as QueryDecryptRequest;
    if (object.message !== undefined && object.message !== null) {
      message.message = String(object.message);
    } else {
      message.message = "";
    }
    return message;
  },

  toJSON(message: QueryDecryptRequest): unknown {
    const obj: any = {};
    message.message !== undefined && (obj.message = message.message);
    return obj;
  },

  fromPartial(object: DeepPartial<QueryDecryptRequest>): QueryDecryptRequest {
    const message = { ...baseQueryDecryptRequest } as QueryDecryptRequest;
    if (object.message !== undefined && object.message !== null) {
      message.message = object.message;
    } else {
      message.message = "";
    }
    return message;
  },
};

const baseQueryDecryptResponse: object = { data: "" };

export const QueryDecryptResponse = {
  encode(
    message: QueryDecryptResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.data !== "") {
      writer.uint32(10).string(message.data);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryDecryptResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryDecryptResponse } as QueryDecryptResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.data = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryDecryptResponse {
    const message = { ...baseQueryDecryptResponse } as QueryDecryptResponse;
    if (object.data !== undefined && object.data !== null) {
      message.data = String(object.data);
    } else {
      message.data = "";
    }
    return message;
  },

  toJSON(message: QueryDecryptResponse): unknown {
    const obj: any = {};
    message.data !== undefined && (obj.data = message.data);
    return obj;
  },

  fromPartial(object: DeepPartial<QueryDecryptResponse>): QueryDecryptResponse {
    const message = { ...baseQueryDecryptResponse } as QueryDecryptResponse;
    if (object.data !== undefined && object.data !== null) {
      message.data = object.data;
    } else {
      message.data = "";
    }
    return message;
  },
};

const baseQueryGetFilesRequest: object = { address: "" };

export const QueryGetFilesRequest = {
  encode(
    message: QueryGetFilesRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.address !== "") {
      writer.uint32(10).string(message.address);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetFilesRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryGetFilesRequest } as QueryGetFilesRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.address = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetFilesRequest {
    const message = { ...baseQueryGetFilesRequest } as QueryGetFilesRequest;
    if (object.address !== undefined && object.address !== null) {
      message.address = String(object.address);
    } else {
      message.address = "";
    }
    return message;
  },

  toJSON(message: QueryGetFilesRequest): unknown {
    const obj: any = {};
    message.address !== undefined && (obj.address = message.address);
    return obj;
  },

  fromPartial(object: DeepPartial<QueryGetFilesRequest>): QueryGetFilesRequest {
    const message = { ...baseQueryGetFilesRequest } as QueryGetFilesRequest;
    if (object.address !== undefined && object.address !== null) {
      message.address = object.address;
    } else {
      message.address = "";
    }
    return message;
  },
};

const baseQueryGetFilesResponse: object = {};

export const QueryGetFilesResponse = {
  encode(
    message: QueryGetFilesResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.files !== undefined) {
      Files.encode(message.files, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetFilesResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryGetFilesResponse } as QueryGetFilesResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.files = Files.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetFilesResponse {
    const message = { ...baseQueryGetFilesResponse } as QueryGetFilesResponse;
    if (object.files !== undefined && object.files !== null) {
      message.files = Files.fromJSON(object.files);
    } else {
      message.files = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetFilesResponse): unknown {
    const obj: any = {};
    message.files !== undefined &&
      (obj.files = message.files ? Files.toJSON(message.files) : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetFilesResponse>
  ): QueryGetFilesResponse {
    const message = { ...baseQueryGetFilesResponse } as QueryGetFilesResponse;
    if (object.files !== undefined && object.files !== null) {
      message.files = Files.fromPartial(object.files);
    } else {
      message.files = undefined;
    }
    return message;
  },
};

const baseQueryAllFilesRequest: object = {};

export const QueryAllFilesRequest = {
  encode(
    message: QueryAllFilesRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllFilesRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryAllFilesRequest } as QueryAllFilesRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.pagination = PageRequest.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllFilesRequest {
    const message = { ...baseQueryAllFilesRequest } as QueryAllFilesRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllFilesRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageRequest.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<QueryAllFilesRequest>): QueryAllFilesRequest {
    const message = { ...baseQueryAllFilesRequest } as QueryAllFilesRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryAllFilesResponse: object = {};

export const QueryAllFilesResponse = {
  encode(
    message: QueryAllFilesResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.files) {
      Files.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(
        message.pagination,
        writer.uint32(18).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllFilesResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryAllFilesResponse } as QueryAllFilesResponse;
    message.files = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.files.push(Files.decode(reader, reader.uint32()));
          break;
        case 2:
          message.pagination = PageResponse.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllFilesResponse {
    const message = { ...baseQueryAllFilesResponse } as QueryAllFilesResponse;
    message.files = [];
    if (object.files !== undefined && object.files !== null) {
      for (const e of object.files) {
        message.files.push(Files.fromJSON(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllFilesResponse): unknown {
    const obj: any = {};
    if (message.files) {
      obj.files = message.files.map((e) => (e ? Files.toJSON(e) : undefined));
    } else {
      obj.files = [];
    }
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageResponse.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllFilesResponse>
  ): QueryAllFilesResponse {
    const message = { ...baseQueryAllFilesResponse } as QueryAllFilesResponse;
    message.files = [];
    if (object.files !== undefined && object.files !== null) {
      for (const e of object.files) {
        message.files.push(Files.fromPartial(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryGetKeysRequest: object = { hashpath: "" };

export const QueryGetKeysRequest = {
  encode(
    message: QueryGetKeysRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.hashpath !== "") {
      writer.uint32(10).string(message.hashpath);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetKeysRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryGetKeysRequest } as QueryGetKeysRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.hashpath = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetKeysRequest {
    const message = { ...baseQueryGetKeysRequest } as QueryGetKeysRequest;
    if (object.hashpath !== undefined && object.hashpath !== null) {
      message.hashpath = String(object.hashpath);
    } else {
      message.hashpath = "";
    }
    return message;
  },

  toJSON(message: QueryGetKeysRequest): unknown {
    const obj: any = {};
    message.hashpath !== undefined && (obj.hashpath = message.hashpath);
    return obj;
  },

  fromPartial(object: DeepPartial<QueryGetKeysRequest>): QueryGetKeysRequest {
    const message = { ...baseQueryGetKeysRequest } as QueryGetKeysRequest;
    if (object.hashpath !== undefined && object.hashpath !== null) {
      message.hashpath = object.hashpath;
    } else {
      message.hashpath = "";
    }
    return message;
  },
};

const baseQueryGetKeysResponse: object = { keys: "" };

export const QueryGetKeysResponse = {
  encode(
    message: QueryGetKeysResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.keys !== "") {
      writer.uint32(10).string(message.keys);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetKeysResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryGetKeysResponse } as QueryGetKeysResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.keys = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetKeysResponse {
    const message = { ...baseQueryGetKeysResponse } as QueryGetKeysResponse;
    if (object.keys !== undefined && object.keys !== null) {
      message.keys = String(object.keys);
    } else {
      message.keys = "";
    }
    return message;
  },

  toJSON(message: QueryGetKeysResponse): unknown {
    const obj: any = {};
    message.keys !== undefined && (obj.keys = message.keys);
    return obj;
  },

  fromPartial(object: DeepPartial<QueryGetKeysResponse>): QueryGetKeysResponse {
    const message = { ...baseQueryGetKeysResponse } as QueryGetKeysResponse;
    if (object.keys !== undefined && object.keys !== null) {
      message.keys = object.keys;
    } else {
      message.keys = "";
    }
    return message;
  },
};

const baseQueryGetKeyRequest: object = { filepath: "", owner: "" };

export const QueryGetKeyRequest = {
  encode(
    message: QueryGetKeyRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.filepath !== "") {
      writer.uint32(10).string(message.filepath);
    }
    if (message.owner !== "") {
      writer.uint32(18).string(message.owner);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetKeyRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryGetKeyRequest } as QueryGetKeyRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.filepath = reader.string();
          break;
        case 2:
          message.owner = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetKeyRequest {
    const message = { ...baseQueryGetKeyRequest } as QueryGetKeyRequest;
    if (object.filepath !== undefined && object.filepath !== null) {
      message.filepath = String(object.filepath);
    } else {
      message.filepath = "";
    }
    if (object.owner !== undefined && object.owner !== null) {
      message.owner = String(object.owner);
    } else {
      message.owner = "";
    }
    return message;
  },

  toJSON(message: QueryGetKeyRequest): unknown {
    const obj: any = {};
    message.filepath !== undefined && (obj.filepath = message.filepath);
    message.owner !== undefined && (obj.owner = message.owner);
    return obj;
  },

  fromPartial(object: DeepPartial<QueryGetKeyRequest>): QueryGetKeyRequest {
    const message = { ...baseQueryGetKeyRequest } as QueryGetKeyRequest;
    if (object.filepath !== undefined && object.filepath !== null) {
      message.filepath = object.filepath;
    } else {
      message.filepath = "";
    }
    if (object.owner !== undefined && object.owner !== null) {
      message.owner = object.owner;
    } else {
      message.owner = "";
    }
    return message;
  },
};

/** Query defines the gRPC querier service. */
export interface Query {
  /** Parameters queries the parameters of the module. */
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse>;
  /** Queries a list of Encrypt items. */
  Encrypt(request: QueryEncryptRequest): Promise<QueryEncryptResponse>;
  /** Queries a list of Decrypt items. */
  Decrypt(request: QueryDecryptRequest): Promise<QueryDecryptResponse>;
  /** Queries a Files by index. */
  Files(request: QueryGetFilesRequest): Promise<QueryGetFilesResponse>;
  /** Queries a list of Files items. */
  FilesAll(request: QueryAllFilesRequest): Promise<QueryAllFilesResponse>;
  /** Queries a list of GetKeys items. */
  GetKeys(request: QueryGetKeysRequest): Promise<QueryGetKeysResponse>;
}

export class QueryClientImpl implements Query {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse> {
    const data = QueryParamsRequest.encode(request).finish();
    const promise = this.rpc.request(
      "jackaldao.canine.filetree.Query",
      "Params",
      data
    );
    return promise.then((data) => QueryParamsResponse.decode(new Reader(data)));
  }

  Encrypt(request: QueryEncryptRequest): Promise<QueryEncryptResponse> {
    const data = QueryEncryptRequest.encode(request).finish();
    const promise = this.rpc.request(
      "jackaldao.canine.filetree.Query",
      "Encrypt",
      data
    );
    return promise.then((data) =>
      QueryEncryptResponse.decode(new Reader(data))
    );
  }

  Decrypt(request: QueryDecryptRequest): Promise<QueryDecryptResponse> {
    const data = QueryDecryptRequest.encode(request).finish();
    const promise = this.rpc.request(
      "jackaldao.canine.filetree.Query",
      "Decrypt",
      data
    );
    return promise.then((data) =>
      QueryDecryptResponse.decode(new Reader(data))
    );
  }

  Files(request: QueryGetFilesRequest): Promise<QueryGetFilesResponse> {
    const data = QueryGetFilesRequest.encode(request).finish();
    const promise = this.rpc.request(
      "jackaldao.canine.filetree.Query",
      "Files",
      data
    );
    return promise.then((data) =>
      QueryGetFilesResponse.decode(new Reader(data))
    );
  }

  FilesAll(request: QueryAllFilesRequest): Promise<QueryAllFilesResponse> {
    const data = QueryAllFilesRequest.encode(request).finish();
    const promise = this.rpc.request(
      "jackaldao.canine.filetree.Query",
      "FilesAll",
      data
    );
    return promise.then((data) =>
      QueryAllFilesResponse.decode(new Reader(data))
    );
  }

  GetKeys(request: QueryGetKeysRequest): Promise<QueryGetKeysResponse> {
    const data = QueryGetKeysRequest.encode(request).finish();
    const promise = this.rpc.request(
      "jackaldao.canine.filetree.Query",
      "GetKeys",
      data
    );
    return promise.then((data) =>
      QueryGetKeysResponse.decode(new Reader(data))
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
