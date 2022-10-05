/* eslint-disable */
import { Reader, Writer } from "protobufjs/minimal";
import { Params } from "../dsig/params";
import { UserUploads } from "../dsig/user_uploads";
import {
  PageRequest,
  PageResponse,
} from "../cosmos/base/query/v1beta1/pagination";
import { Form } from "../dsig/form";

export const protobufPackage = "jackaldao.canine.dsig";

/** QueryParamsRequest is request type for the Query/Params RPC method. */
export interface QueryParamsRequest {}

/** QueryParamsResponse is response type for the Query/Params RPC method. */
export interface QueryParamsResponse {
  /** params holds all the parameters of this module. */
  params: Params | undefined;
}

export interface QueryGetUserUploadsRequest {
  fid: string;
}

export interface QueryGetUserUploadsResponse {
  userUploads: UserUploads | undefined;
}

export interface QueryAllUserUploadsRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllUserUploadsResponse {
  userUploads: UserUploads[];
  pagination: PageResponse | undefined;
}

export interface QueryGetFormRequest {
  ffid: string;
}

export interface QueryGetFormResponse {
  form: Form | undefined;
}

export interface QueryAllFormRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllFormResponse {
  form: Form[];
  pagination: PageResponse | undefined;
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

const baseQueryGetUserUploadsRequest: object = { fid: "" };

export const QueryGetUserUploadsRequest = {
  encode(
    message: QueryGetUserUploadsRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.fid !== "") {
      writer.uint32(10).string(message.fid);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetUserUploadsRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetUserUploadsRequest,
    } as QueryGetUserUploadsRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.fid = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetUserUploadsRequest {
    const message = {
      ...baseQueryGetUserUploadsRequest,
    } as QueryGetUserUploadsRequest;
    if (object.fid !== undefined && object.fid !== null) {
      message.fid = String(object.fid);
    } else {
      message.fid = "";
    }
    return message;
  },

  toJSON(message: QueryGetUserUploadsRequest): unknown {
    const obj: any = {};
    message.fid !== undefined && (obj.fid = message.fid);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetUserUploadsRequest>
  ): QueryGetUserUploadsRequest {
    const message = {
      ...baseQueryGetUserUploadsRequest,
    } as QueryGetUserUploadsRequest;
    if (object.fid !== undefined && object.fid !== null) {
      message.fid = object.fid;
    } else {
      message.fid = "";
    }
    return message;
  },
};

const baseQueryGetUserUploadsResponse: object = {};

export const QueryGetUserUploadsResponse = {
  encode(
    message: QueryGetUserUploadsResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.userUploads !== undefined) {
      UserUploads.encode(
        message.userUploads,
        writer.uint32(10).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetUserUploadsResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetUserUploadsResponse,
    } as QueryGetUserUploadsResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.userUploads = UserUploads.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetUserUploadsResponse {
    const message = {
      ...baseQueryGetUserUploadsResponse,
    } as QueryGetUserUploadsResponse;
    if (object.userUploads !== undefined && object.userUploads !== null) {
      message.userUploads = UserUploads.fromJSON(object.userUploads);
    } else {
      message.userUploads = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetUserUploadsResponse): unknown {
    const obj: any = {};
    message.userUploads !== undefined &&
      (obj.userUploads = message.userUploads
        ? UserUploads.toJSON(message.userUploads)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetUserUploadsResponse>
  ): QueryGetUserUploadsResponse {
    const message = {
      ...baseQueryGetUserUploadsResponse,
    } as QueryGetUserUploadsResponse;
    if (object.userUploads !== undefined && object.userUploads !== null) {
      message.userUploads = UserUploads.fromPartial(object.userUploads);
    } else {
      message.userUploads = undefined;
    }
    return message;
  },
};

const baseQueryAllUserUploadsRequest: object = {};

export const QueryAllUserUploadsRequest = {
  encode(
    message: QueryAllUserUploadsRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryAllUserUploadsRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllUserUploadsRequest,
    } as QueryAllUserUploadsRequest;
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

  fromJSON(object: any): QueryAllUserUploadsRequest {
    const message = {
      ...baseQueryAllUserUploadsRequest,
    } as QueryAllUserUploadsRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllUserUploadsRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageRequest.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllUserUploadsRequest>
  ): QueryAllUserUploadsRequest {
    const message = {
      ...baseQueryAllUserUploadsRequest,
    } as QueryAllUserUploadsRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryAllUserUploadsResponse: object = {};

export const QueryAllUserUploadsResponse = {
  encode(
    message: QueryAllUserUploadsResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.userUploads) {
      UserUploads.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(
        message.pagination,
        writer.uint32(18).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryAllUserUploadsResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllUserUploadsResponse,
    } as QueryAllUserUploadsResponse;
    message.userUploads = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.userUploads.push(UserUploads.decode(reader, reader.uint32()));
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

  fromJSON(object: any): QueryAllUserUploadsResponse {
    const message = {
      ...baseQueryAllUserUploadsResponse,
    } as QueryAllUserUploadsResponse;
    message.userUploads = [];
    if (object.userUploads !== undefined && object.userUploads !== null) {
      for (const e of object.userUploads) {
        message.userUploads.push(UserUploads.fromJSON(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllUserUploadsResponse): unknown {
    const obj: any = {};
    if (message.userUploads) {
      obj.userUploads = message.userUploads.map((e) =>
        e ? UserUploads.toJSON(e) : undefined
      );
    } else {
      obj.userUploads = [];
    }
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageResponse.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllUserUploadsResponse>
  ): QueryAllUserUploadsResponse {
    const message = {
      ...baseQueryAllUserUploadsResponse,
    } as QueryAllUserUploadsResponse;
    message.userUploads = [];
    if (object.userUploads !== undefined && object.userUploads !== null) {
      for (const e of object.userUploads) {
        message.userUploads.push(UserUploads.fromPartial(e));
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

const baseQueryGetFormRequest: object = { ffid: "" };

export const QueryGetFormRequest = {
  encode(
    message: QueryGetFormRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.ffid !== "") {
      writer.uint32(10).string(message.ffid);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetFormRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryGetFormRequest } as QueryGetFormRequest;
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

  fromJSON(object: any): QueryGetFormRequest {
    const message = { ...baseQueryGetFormRequest } as QueryGetFormRequest;
    if (object.ffid !== undefined && object.ffid !== null) {
      message.ffid = String(object.ffid);
    } else {
      message.ffid = "";
    }
    return message;
  },

  toJSON(message: QueryGetFormRequest): unknown {
    const obj: any = {};
    message.ffid !== undefined && (obj.ffid = message.ffid);
    return obj;
  },

  fromPartial(object: DeepPartial<QueryGetFormRequest>): QueryGetFormRequest {
    const message = { ...baseQueryGetFormRequest } as QueryGetFormRequest;
    if (object.ffid !== undefined && object.ffid !== null) {
      message.ffid = object.ffid;
    } else {
      message.ffid = "";
    }
    return message;
  },
};

const baseQueryGetFormResponse: object = {};

export const QueryGetFormResponse = {
  encode(
    message: QueryGetFormResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.form !== undefined) {
      Form.encode(message.form, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetFormResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryGetFormResponse } as QueryGetFormResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.form = Form.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetFormResponse {
    const message = { ...baseQueryGetFormResponse } as QueryGetFormResponse;
    if (object.form !== undefined && object.form !== null) {
      message.form = Form.fromJSON(object.form);
    } else {
      message.form = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetFormResponse): unknown {
    const obj: any = {};
    message.form !== undefined &&
      (obj.form = message.form ? Form.toJSON(message.form) : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<QueryGetFormResponse>): QueryGetFormResponse {
    const message = { ...baseQueryGetFormResponse } as QueryGetFormResponse;
    if (object.form !== undefined && object.form !== null) {
      message.form = Form.fromPartial(object.form);
    } else {
      message.form = undefined;
    }
    return message;
  },
};

const baseQueryAllFormRequest: object = {};

export const QueryAllFormRequest = {
  encode(
    message: QueryAllFormRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllFormRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryAllFormRequest } as QueryAllFormRequest;
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

  fromJSON(object: any): QueryAllFormRequest {
    const message = { ...baseQueryAllFormRequest } as QueryAllFormRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllFormRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageRequest.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<QueryAllFormRequest>): QueryAllFormRequest {
    const message = { ...baseQueryAllFormRequest } as QueryAllFormRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryAllFormResponse: object = {};

export const QueryAllFormResponse = {
  encode(
    message: QueryAllFormResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.form) {
      Form.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(
        message.pagination,
        writer.uint32(18).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllFormResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryAllFormResponse } as QueryAllFormResponse;
    message.form = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.form.push(Form.decode(reader, reader.uint32()));
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

  fromJSON(object: any): QueryAllFormResponse {
    const message = { ...baseQueryAllFormResponse } as QueryAllFormResponse;
    message.form = [];
    if (object.form !== undefined && object.form !== null) {
      for (const e of object.form) {
        message.form.push(Form.fromJSON(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllFormResponse): unknown {
    const obj: any = {};
    if (message.form) {
      obj.form = message.form.map((e) => (e ? Form.toJSON(e) : undefined));
    } else {
      obj.form = [];
    }
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageResponse.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<QueryAllFormResponse>): QueryAllFormResponse {
    const message = { ...baseQueryAllFormResponse } as QueryAllFormResponse;
    message.form = [];
    if (object.form !== undefined && object.form !== null) {
      for (const e of object.form) {
        message.form.push(Form.fromPartial(e));
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

/** Query defines the gRPC querier service. */
export interface Query {
  /** Parameters queries the parameters of the module. */
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse>;
  /** Queries a UserUploads by index. */
  UserUploads(
    request: QueryGetUserUploadsRequest
  ): Promise<QueryGetUserUploadsResponse>;
  /** Queries a list of UserUploads items. */
  UserUploadsAll(
    request: QueryAllUserUploadsRequest
  ): Promise<QueryAllUserUploadsResponse>;
  /** Queries a Form by index. */
  Form(request: QueryGetFormRequest): Promise<QueryGetFormResponse>;
  /** Queries a list of Form items. */
  FormAll(request: QueryAllFormRequest): Promise<QueryAllFormResponse>;
}

export class QueryClientImpl implements Query {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse> {
    const data = QueryParamsRequest.encode(request).finish();
    const promise = this.rpc.request(
      "jackaldao.canine.dsig.Query",
      "Params",
      data
    );
    return promise.then((data) => QueryParamsResponse.decode(new Reader(data)));
  }

  UserUploads(
    request: QueryGetUserUploadsRequest
  ): Promise<QueryGetUserUploadsResponse> {
    const data = QueryGetUserUploadsRequest.encode(request).finish();
    const promise = this.rpc.request(
      "jackaldao.canine.dsig.Query",
      "UserUploads",
      data
    );
    return promise.then((data) =>
      QueryGetUserUploadsResponse.decode(new Reader(data))
    );
  }

  UserUploadsAll(
    request: QueryAllUserUploadsRequest
  ): Promise<QueryAllUserUploadsResponse> {
    const data = QueryAllUserUploadsRequest.encode(request).finish();
    const promise = this.rpc.request(
      "jackaldao.canine.dsig.Query",
      "UserUploadsAll",
      data
    );
    return promise.then((data) =>
      QueryAllUserUploadsResponse.decode(new Reader(data))
    );
  }

  Form(request: QueryGetFormRequest): Promise<QueryGetFormResponse> {
    const data = QueryGetFormRequest.encode(request).finish();
    const promise = this.rpc.request(
      "jackaldao.canine.dsig.Query",
      "Form",
      data
    );
    return promise.then((data) =>
      QueryGetFormResponse.decode(new Reader(data))
    );
  }

  FormAll(request: QueryAllFormRequest): Promise<QueryAllFormResponse> {
    const data = QueryAllFormRequest.encode(request).finish();
    const promise = this.rpc.request(
      "jackaldao.canine.dsig.Query",
      "FormAll",
      data
    );
    return promise.then((data) =>
      QueryAllFormResponse.decode(new Reader(data))
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
