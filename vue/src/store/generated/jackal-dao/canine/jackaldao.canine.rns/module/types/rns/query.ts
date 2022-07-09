/* eslint-disable */
import { Reader, Writer } from "protobufjs/minimal";
import { Params } from "../rns/params";
import { Whois } from "../rns/whois";
import {
  PageRequest,
  PageResponse,
} from "../cosmos/base/query/v1beta1/pagination";
import { Names } from "../rns/names";
import { Bids } from "../rns/bids";
import { Forsale } from "../rns/forsale";

export const protobufPackage = "jackaldao.canine.rns";

/** QueryParamsRequest is request type for the Query/Params RPC method. */
export interface QueryParamsRequest {}

/** QueryParamsResponse is response type for the Query/Params RPC method. */
export interface QueryParamsResponse {
  /** params holds all the parameters of this module. */
  params: Params | undefined;
}

export interface QueryGetWhoisRequest {
  index: string;
}

export interface QueryGetWhoisResponse {
  whois: Whois | undefined;
}

export interface QueryAllWhoisRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllWhoisResponse {
  whois: Whois[];
  pagination: PageResponse | undefined;
}

export interface QueryGetNamesRequest {
  index: string;
}

export interface QueryGetNamesResponse {
  names: Names | undefined;
}

export interface QueryAllNamesRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllNamesResponse {
  names: Names[];
  pagination: PageResponse | undefined;
}

export interface QueryGetBidsRequest {
  index: string;
}

export interface QueryGetBidsResponse {
  bids: Bids | undefined;
}

export interface QueryAllBidsRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllBidsResponse {
  bids: Bids[];
  pagination: PageResponse | undefined;
}

export interface QueryGetForsaleRequest {
  name: string;
}

export interface QueryGetForsaleResponse {
  forsale: Forsale | undefined;
}

export interface QueryAllForsaleRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllForsaleResponse {
  forsale: Forsale[];
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

const baseQueryGetWhoisRequest: object = { index: "" };

export const QueryGetWhoisRequest = {
  encode(
    message: QueryGetWhoisRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.index !== "") {
      writer.uint32(10).string(message.index);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetWhoisRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryGetWhoisRequest } as QueryGetWhoisRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.index = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetWhoisRequest {
    const message = { ...baseQueryGetWhoisRequest } as QueryGetWhoisRequest;
    if (object.index !== undefined && object.index !== null) {
      message.index = String(object.index);
    } else {
      message.index = "";
    }
    return message;
  },

  toJSON(message: QueryGetWhoisRequest): unknown {
    const obj: any = {};
    message.index !== undefined && (obj.index = message.index);
    return obj;
  },

  fromPartial(object: DeepPartial<QueryGetWhoisRequest>): QueryGetWhoisRequest {
    const message = { ...baseQueryGetWhoisRequest } as QueryGetWhoisRequest;
    if (object.index !== undefined && object.index !== null) {
      message.index = object.index;
    } else {
      message.index = "";
    }
    return message;
  },
};

const baseQueryGetWhoisResponse: object = {};

export const QueryGetWhoisResponse = {
  encode(
    message: QueryGetWhoisResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.whois !== undefined) {
      Whois.encode(message.whois, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetWhoisResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryGetWhoisResponse } as QueryGetWhoisResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.whois = Whois.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetWhoisResponse {
    const message = { ...baseQueryGetWhoisResponse } as QueryGetWhoisResponse;
    if (object.whois !== undefined && object.whois !== null) {
      message.whois = Whois.fromJSON(object.whois);
    } else {
      message.whois = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetWhoisResponse): unknown {
    const obj: any = {};
    message.whois !== undefined &&
      (obj.whois = message.whois ? Whois.toJSON(message.whois) : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetWhoisResponse>
  ): QueryGetWhoisResponse {
    const message = { ...baseQueryGetWhoisResponse } as QueryGetWhoisResponse;
    if (object.whois !== undefined && object.whois !== null) {
      message.whois = Whois.fromPartial(object.whois);
    } else {
      message.whois = undefined;
    }
    return message;
  },
};

const baseQueryAllWhoisRequest: object = {};

export const QueryAllWhoisRequest = {
  encode(
    message: QueryAllWhoisRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllWhoisRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryAllWhoisRequest } as QueryAllWhoisRequest;
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

  fromJSON(object: any): QueryAllWhoisRequest {
    const message = { ...baseQueryAllWhoisRequest } as QueryAllWhoisRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllWhoisRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageRequest.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<QueryAllWhoisRequest>): QueryAllWhoisRequest {
    const message = { ...baseQueryAllWhoisRequest } as QueryAllWhoisRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryAllWhoisResponse: object = {};

export const QueryAllWhoisResponse = {
  encode(
    message: QueryAllWhoisResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.whois) {
      Whois.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(
        message.pagination,
        writer.uint32(18).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllWhoisResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryAllWhoisResponse } as QueryAllWhoisResponse;
    message.whois = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.whois.push(Whois.decode(reader, reader.uint32()));
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

  fromJSON(object: any): QueryAllWhoisResponse {
    const message = { ...baseQueryAllWhoisResponse } as QueryAllWhoisResponse;
    message.whois = [];
    if (object.whois !== undefined && object.whois !== null) {
      for (const e of object.whois) {
        message.whois.push(Whois.fromJSON(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllWhoisResponse): unknown {
    const obj: any = {};
    if (message.whois) {
      obj.whois = message.whois.map((e) => (e ? Whois.toJSON(e) : undefined));
    } else {
      obj.whois = [];
    }
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageResponse.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllWhoisResponse>
  ): QueryAllWhoisResponse {
    const message = { ...baseQueryAllWhoisResponse } as QueryAllWhoisResponse;
    message.whois = [];
    if (object.whois !== undefined && object.whois !== null) {
      for (const e of object.whois) {
        message.whois.push(Whois.fromPartial(e));
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

const baseQueryGetNamesRequest: object = { index: "" };

export const QueryGetNamesRequest = {
  encode(
    message: QueryGetNamesRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.index !== "") {
      writer.uint32(10).string(message.index);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetNamesRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryGetNamesRequest } as QueryGetNamesRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.index = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetNamesRequest {
    const message = { ...baseQueryGetNamesRequest } as QueryGetNamesRequest;
    if (object.index !== undefined && object.index !== null) {
      message.index = String(object.index);
    } else {
      message.index = "";
    }
    return message;
  },

  toJSON(message: QueryGetNamesRequest): unknown {
    const obj: any = {};
    message.index !== undefined && (obj.index = message.index);
    return obj;
  },

  fromPartial(object: DeepPartial<QueryGetNamesRequest>): QueryGetNamesRequest {
    const message = { ...baseQueryGetNamesRequest } as QueryGetNamesRequest;
    if (object.index !== undefined && object.index !== null) {
      message.index = object.index;
    } else {
      message.index = "";
    }
    return message;
  },
};

const baseQueryGetNamesResponse: object = {};

export const QueryGetNamesResponse = {
  encode(
    message: QueryGetNamesResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.names !== undefined) {
      Names.encode(message.names, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetNamesResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryGetNamesResponse } as QueryGetNamesResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.names = Names.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetNamesResponse {
    const message = { ...baseQueryGetNamesResponse } as QueryGetNamesResponse;
    if (object.names !== undefined && object.names !== null) {
      message.names = Names.fromJSON(object.names);
    } else {
      message.names = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetNamesResponse): unknown {
    const obj: any = {};
    message.names !== undefined &&
      (obj.names = message.names ? Names.toJSON(message.names) : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetNamesResponse>
  ): QueryGetNamesResponse {
    const message = { ...baseQueryGetNamesResponse } as QueryGetNamesResponse;
    if (object.names !== undefined && object.names !== null) {
      message.names = Names.fromPartial(object.names);
    } else {
      message.names = undefined;
    }
    return message;
  },
};

const baseQueryAllNamesRequest: object = {};

export const QueryAllNamesRequest = {
  encode(
    message: QueryAllNamesRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllNamesRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryAllNamesRequest } as QueryAllNamesRequest;
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

  fromJSON(object: any): QueryAllNamesRequest {
    const message = { ...baseQueryAllNamesRequest } as QueryAllNamesRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllNamesRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageRequest.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<QueryAllNamesRequest>): QueryAllNamesRequest {
    const message = { ...baseQueryAllNamesRequest } as QueryAllNamesRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryAllNamesResponse: object = {};

export const QueryAllNamesResponse = {
  encode(
    message: QueryAllNamesResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.names) {
      Names.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(
        message.pagination,
        writer.uint32(18).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllNamesResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryAllNamesResponse } as QueryAllNamesResponse;
    message.names = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.names.push(Names.decode(reader, reader.uint32()));
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

  fromJSON(object: any): QueryAllNamesResponse {
    const message = { ...baseQueryAllNamesResponse } as QueryAllNamesResponse;
    message.names = [];
    if (object.names !== undefined && object.names !== null) {
      for (const e of object.names) {
        message.names.push(Names.fromJSON(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllNamesResponse): unknown {
    const obj: any = {};
    if (message.names) {
      obj.names = message.names.map((e) => (e ? Names.toJSON(e) : undefined));
    } else {
      obj.names = [];
    }
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageResponse.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllNamesResponse>
  ): QueryAllNamesResponse {
    const message = { ...baseQueryAllNamesResponse } as QueryAllNamesResponse;
    message.names = [];
    if (object.names !== undefined && object.names !== null) {
      for (const e of object.names) {
        message.names.push(Names.fromPartial(e));
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

const baseQueryGetBidsRequest: object = { index: "" };

export const QueryGetBidsRequest = {
  encode(
    message: QueryGetBidsRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.index !== "") {
      writer.uint32(10).string(message.index);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetBidsRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryGetBidsRequest } as QueryGetBidsRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.index = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetBidsRequest {
    const message = { ...baseQueryGetBidsRequest } as QueryGetBidsRequest;
    if (object.index !== undefined && object.index !== null) {
      message.index = String(object.index);
    } else {
      message.index = "";
    }
    return message;
  },

  toJSON(message: QueryGetBidsRequest): unknown {
    const obj: any = {};
    message.index !== undefined && (obj.index = message.index);
    return obj;
  },

  fromPartial(object: DeepPartial<QueryGetBidsRequest>): QueryGetBidsRequest {
    const message = { ...baseQueryGetBidsRequest } as QueryGetBidsRequest;
    if (object.index !== undefined && object.index !== null) {
      message.index = object.index;
    } else {
      message.index = "";
    }
    return message;
  },
};

const baseQueryGetBidsResponse: object = {};

export const QueryGetBidsResponse = {
  encode(
    message: QueryGetBidsResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.bids !== undefined) {
      Bids.encode(message.bids, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetBidsResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryGetBidsResponse } as QueryGetBidsResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.bids = Bids.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetBidsResponse {
    const message = { ...baseQueryGetBidsResponse } as QueryGetBidsResponse;
    if (object.bids !== undefined && object.bids !== null) {
      message.bids = Bids.fromJSON(object.bids);
    } else {
      message.bids = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetBidsResponse): unknown {
    const obj: any = {};
    message.bids !== undefined &&
      (obj.bids = message.bids ? Bids.toJSON(message.bids) : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<QueryGetBidsResponse>): QueryGetBidsResponse {
    const message = { ...baseQueryGetBidsResponse } as QueryGetBidsResponse;
    if (object.bids !== undefined && object.bids !== null) {
      message.bids = Bids.fromPartial(object.bids);
    } else {
      message.bids = undefined;
    }
    return message;
  },
};

const baseQueryAllBidsRequest: object = {};

export const QueryAllBidsRequest = {
  encode(
    message: QueryAllBidsRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllBidsRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryAllBidsRequest } as QueryAllBidsRequest;
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

  fromJSON(object: any): QueryAllBidsRequest {
    const message = { ...baseQueryAllBidsRequest } as QueryAllBidsRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllBidsRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageRequest.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<QueryAllBidsRequest>): QueryAllBidsRequest {
    const message = { ...baseQueryAllBidsRequest } as QueryAllBidsRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryAllBidsResponse: object = {};

export const QueryAllBidsResponse = {
  encode(
    message: QueryAllBidsResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.bids) {
      Bids.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(
        message.pagination,
        writer.uint32(18).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllBidsResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryAllBidsResponse } as QueryAllBidsResponse;
    message.bids = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.bids.push(Bids.decode(reader, reader.uint32()));
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

  fromJSON(object: any): QueryAllBidsResponse {
    const message = { ...baseQueryAllBidsResponse } as QueryAllBidsResponse;
    message.bids = [];
    if (object.bids !== undefined && object.bids !== null) {
      for (const e of object.bids) {
        message.bids.push(Bids.fromJSON(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllBidsResponse): unknown {
    const obj: any = {};
    if (message.bids) {
      obj.bids = message.bids.map((e) => (e ? Bids.toJSON(e) : undefined));
    } else {
      obj.bids = [];
    }
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageResponse.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<QueryAllBidsResponse>): QueryAllBidsResponse {
    const message = { ...baseQueryAllBidsResponse } as QueryAllBidsResponse;
    message.bids = [];
    if (object.bids !== undefined && object.bids !== null) {
      for (const e of object.bids) {
        message.bids.push(Bids.fromPartial(e));
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

const baseQueryGetForsaleRequest: object = { name: "" };

export const QueryGetForsaleRequest = {
  encode(
    message: QueryGetForsaleRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.name !== "") {
      writer.uint32(10).string(message.name);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetForsaleRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryGetForsaleRequest } as QueryGetForsaleRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.name = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetForsaleRequest {
    const message = { ...baseQueryGetForsaleRequest } as QueryGetForsaleRequest;
    if (object.name !== undefined && object.name !== null) {
      message.name = String(object.name);
    } else {
      message.name = "";
    }
    return message;
  },

  toJSON(message: QueryGetForsaleRequest): unknown {
    const obj: any = {};
    message.name !== undefined && (obj.name = message.name);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetForsaleRequest>
  ): QueryGetForsaleRequest {
    const message = { ...baseQueryGetForsaleRequest } as QueryGetForsaleRequest;
    if (object.name !== undefined && object.name !== null) {
      message.name = object.name;
    } else {
      message.name = "";
    }
    return message;
  },
};

const baseQueryGetForsaleResponse: object = {};

export const QueryGetForsaleResponse = {
  encode(
    message: QueryGetForsaleResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.forsale !== undefined) {
      Forsale.encode(message.forsale, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetForsaleResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetForsaleResponse,
    } as QueryGetForsaleResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.forsale = Forsale.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetForsaleResponse {
    const message = {
      ...baseQueryGetForsaleResponse,
    } as QueryGetForsaleResponse;
    if (object.forsale !== undefined && object.forsale !== null) {
      message.forsale = Forsale.fromJSON(object.forsale);
    } else {
      message.forsale = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetForsaleResponse): unknown {
    const obj: any = {};
    message.forsale !== undefined &&
      (obj.forsale = message.forsale
        ? Forsale.toJSON(message.forsale)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetForsaleResponse>
  ): QueryGetForsaleResponse {
    const message = {
      ...baseQueryGetForsaleResponse,
    } as QueryGetForsaleResponse;
    if (object.forsale !== undefined && object.forsale !== null) {
      message.forsale = Forsale.fromPartial(object.forsale);
    } else {
      message.forsale = undefined;
    }
    return message;
  },
};

const baseQueryAllForsaleRequest: object = {};

export const QueryAllForsaleRequest = {
  encode(
    message: QueryAllForsaleRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllForsaleRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryAllForsaleRequest } as QueryAllForsaleRequest;
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

  fromJSON(object: any): QueryAllForsaleRequest {
    const message = { ...baseQueryAllForsaleRequest } as QueryAllForsaleRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllForsaleRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageRequest.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllForsaleRequest>
  ): QueryAllForsaleRequest {
    const message = { ...baseQueryAllForsaleRequest } as QueryAllForsaleRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryAllForsaleResponse: object = {};

export const QueryAllForsaleResponse = {
  encode(
    message: QueryAllForsaleResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.forsale) {
      Forsale.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(
        message.pagination,
        writer.uint32(18).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllForsaleResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllForsaleResponse,
    } as QueryAllForsaleResponse;
    message.forsale = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.forsale.push(Forsale.decode(reader, reader.uint32()));
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

  fromJSON(object: any): QueryAllForsaleResponse {
    const message = {
      ...baseQueryAllForsaleResponse,
    } as QueryAllForsaleResponse;
    message.forsale = [];
    if (object.forsale !== undefined && object.forsale !== null) {
      for (const e of object.forsale) {
        message.forsale.push(Forsale.fromJSON(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllForsaleResponse): unknown {
    const obj: any = {};
    if (message.forsale) {
      obj.forsale = message.forsale.map((e) =>
        e ? Forsale.toJSON(e) : undefined
      );
    } else {
      obj.forsale = [];
    }
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageResponse.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllForsaleResponse>
  ): QueryAllForsaleResponse {
    const message = {
      ...baseQueryAllForsaleResponse,
    } as QueryAllForsaleResponse;
    message.forsale = [];
    if (object.forsale !== undefined && object.forsale !== null) {
      for (const e of object.forsale) {
        message.forsale.push(Forsale.fromPartial(e));
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
  /** Queries a Name by index. */
  Names(request: QueryGetNamesRequest): Promise<QueryGetNamesResponse>;
  /** Queries a list of Names. */
  NamesAll(request: QueryAllNamesRequest): Promise<QueryAllNamesResponse>;
  /** Queries a Bid by index. */
  Bids(request: QueryGetBidsRequest): Promise<QueryGetBidsResponse>;
  /** Queries a list of Bids. */
  BidsAll(request: QueryAllBidsRequest): Promise<QueryAllBidsResponse>;
  /** Queries a Listing by index. */
  Forsale(request: QueryGetForsaleRequest): Promise<QueryGetForsaleResponse>;
  /** Queries all Listings. */
  ForsaleAll(request: QueryAllForsaleRequest): Promise<QueryAllForsaleResponse>;
}

export class QueryClientImpl implements Query {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse> {
    const data = QueryParamsRequest.encode(request).finish();
    const promise = this.rpc.request(
      "jackaldao.canine.rns.Query",
      "Params",
      data
    );
    return promise.then((data) => QueryParamsResponse.decode(new Reader(data)));
  }

  Names(request: QueryGetNamesRequest): Promise<QueryGetNamesResponse> {
    const data = QueryGetNamesRequest.encode(request).finish();
    const promise = this.rpc.request(
      "jackaldao.canine.rns.Query",
      "Names",
      data
    );
    return promise.then((data) =>
      QueryGetNamesResponse.decode(new Reader(data))
    );
  }

  NamesAll(request: QueryAllNamesRequest): Promise<QueryAllNamesResponse> {
    const data = QueryAllNamesRequest.encode(request).finish();
    const promise = this.rpc.request(
      "jackaldao.canine.rns.Query",
      "NamesAll",
      data
    );
    return promise.then((data) =>
      QueryAllNamesResponse.decode(new Reader(data))
    );
  }

  Bids(request: QueryGetBidsRequest): Promise<QueryGetBidsResponse> {
    const data = QueryGetBidsRequest.encode(request).finish();
    const promise = this.rpc.request(
      "jackaldao.canine.rns.Query",
      "Bids",
      data
    );
    return promise.then((data) =>
      QueryGetBidsResponse.decode(new Reader(data))
    );
  }

  BidsAll(request: QueryAllBidsRequest): Promise<QueryAllBidsResponse> {
    const data = QueryAllBidsRequest.encode(request).finish();
    const promise = this.rpc.request(
      "jackaldao.canine.rns.Query",
      "BidsAll",
      data
    );
    return promise.then((data) =>
      QueryAllBidsResponse.decode(new Reader(data))
    );
  }

  Forsale(request: QueryGetForsaleRequest): Promise<QueryGetForsaleResponse> {
    const data = QueryGetForsaleRequest.encode(request).finish();
    const promise = this.rpc.request(
      "jackaldao.canine.rns.Query",
      "Forsale",
      data
    );
    return promise.then((data) =>
      QueryGetForsaleResponse.decode(new Reader(data))
    );
  }

  ForsaleAll(
    request: QueryAllForsaleRequest
  ): Promise<QueryAllForsaleResponse> {
    const data = QueryAllForsaleRequest.encode(request).finish();
    const promise = this.rpc.request(
      "jackaldao.canine.rns.Query",
      "ForsaleAll",
      data
    );
    return promise.then((data) =>
      QueryAllForsaleResponse.decode(new Reader(data))
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
