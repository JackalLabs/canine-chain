/* eslint-disable */
import { Reader, util, configure, Writer } from "protobufjs/minimal";
import * as Long from "long";
import { Params } from "../jklmining/params";
import { SaveRequests } from "../jklmining/save_requests";
import {
  PageRequest,
  PageResponse,
} from "../cosmos/base/query/v1beta1/pagination";
import { Miners } from "../jklmining/miners";
import { Mined } from "../jklmining/mined";

export const protobufPackage = "jackaldao.canine.jklmining";

/** QueryParamsRequest is request type for the Query/Params RPC method. */
export interface QueryParamsRequest {}

/** QueryParamsResponse is response type for the Query/Params RPC method. */
export interface QueryParamsResponse {
  /** params holds all the parameters of this module. */
  params: Params | undefined;
}

export interface QueryGetSaveRequestsRequest {
  index: string;
}

export interface QueryGetSaveRequestsResponse {
  saveRequests: SaveRequests | undefined;
}

export interface QueryAllSaveRequestsRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllSaveRequestsResponse {
  saveRequests: SaveRequests[];
  pagination: PageResponse | undefined;
}

export interface QueryGetMinersRequest {
  address: string;
}

export interface QueryGetMinersResponse {
  miners: Miners | undefined;
}

export interface QueryAllMinersRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllMinersResponse {
  miners: Miners[];
  pagination: PageResponse | undefined;
}

export interface QueryGetMinedRequest {
  id: number;
}

export interface QueryGetMinedResponse {
  Mined: Mined | undefined;
}

export interface QueryAllMinedRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllMinedResponse {
  Mined: Mined[];
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

const baseQueryGetSaveRequestsRequest: object = { index: "" };

export const QueryGetSaveRequestsRequest = {
  encode(
    message: QueryGetSaveRequestsRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.index !== "") {
      writer.uint32(10).string(message.index);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetSaveRequestsRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetSaveRequestsRequest,
    } as QueryGetSaveRequestsRequest;
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

  fromJSON(object: any): QueryGetSaveRequestsRequest {
    const message = {
      ...baseQueryGetSaveRequestsRequest,
    } as QueryGetSaveRequestsRequest;
    if (object.index !== undefined && object.index !== null) {
      message.index = String(object.index);
    } else {
      message.index = "";
    }
    return message;
  },

  toJSON(message: QueryGetSaveRequestsRequest): unknown {
    const obj: any = {};
    message.index !== undefined && (obj.index = message.index);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetSaveRequestsRequest>
  ): QueryGetSaveRequestsRequest {
    const message = {
      ...baseQueryGetSaveRequestsRequest,
    } as QueryGetSaveRequestsRequest;
    if (object.index !== undefined && object.index !== null) {
      message.index = object.index;
    } else {
      message.index = "";
    }
    return message;
  },
};

const baseQueryGetSaveRequestsResponse: object = {};

export const QueryGetSaveRequestsResponse = {
  encode(
    message: QueryGetSaveRequestsResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.saveRequests !== undefined) {
      SaveRequests.encode(
        message.saveRequests,
        writer.uint32(10).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetSaveRequestsResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetSaveRequestsResponse,
    } as QueryGetSaveRequestsResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.saveRequests = SaveRequests.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetSaveRequestsResponse {
    const message = {
      ...baseQueryGetSaveRequestsResponse,
    } as QueryGetSaveRequestsResponse;
    if (object.saveRequests !== undefined && object.saveRequests !== null) {
      message.saveRequests = SaveRequests.fromJSON(object.saveRequests);
    } else {
      message.saveRequests = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetSaveRequestsResponse): unknown {
    const obj: any = {};
    message.saveRequests !== undefined &&
      (obj.saveRequests = message.saveRequests
        ? SaveRequests.toJSON(message.saveRequests)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetSaveRequestsResponse>
  ): QueryGetSaveRequestsResponse {
    const message = {
      ...baseQueryGetSaveRequestsResponse,
    } as QueryGetSaveRequestsResponse;
    if (object.saveRequests !== undefined && object.saveRequests !== null) {
      message.saveRequests = SaveRequests.fromPartial(object.saveRequests);
    } else {
      message.saveRequests = undefined;
    }
    return message;
  },
};

const baseQueryAllSaveRequestsRequest: object = {};

export const QueryAllSaveRequestsRequest = {
  encode(
    message: QueryAllSaveRequestsRequest,
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
  ): QueryAllSaveRequestsRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllSaveRequestsRequest,
    } as QueryAllSaveRequestsRequest;
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

  fromJSON(object: any): QueryAllSaveRequestsRequest {
    const message = {
      ...baseQueryAllSaveRequestsRequest,
    } as QueryAllSaveRequestsRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllSaveRequestsRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageRequest.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllSaveRequestsRequest>
  ): QueryAllSaveRequestsRequest {
    const message = {
      ...baseQueryAllSaveRequestsRequest,
    } as QueryAllSaveRequestsRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryAllSaveRequestsResponse: object = {};

export const QueryAllSaveRequestsResponse = {
  encode(
    message: QueryAllSaveRequestsResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.saveRequests) {
      SaveRequests.encode(v!, writer.uint32(10).fork()).ldelim();
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
  ): QueryAllSaveRequestsResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllSaveRequestsResponse,
    } as QueryAllSaveRequestsResponse;
    message.saveRequests = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.saveRequests.push(
            SaveRequests.decode(reader, reader.uint32())
          );
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

  fromJSON(object: any): QueryAllSaveRequestsResponse {
    const message = {
      ...baseQueryAllSaveRequestsResponse,
    } as QueryAllSaveRequestsResponse;
    message.saveRequests = [];
    if (object.saveRequests !== undefined && object.saveRequests !== null) {
      for (const e of object.saveRequests) {
        message.saveRequests.push(SaveRequests.fromJSON(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllSaveRequestsResponse): unknown {
    const obj: any = {};
    if (message.saveRequests) {
      obj.saveRequests = message.saveRequests.map((e) =>
        e ? SaveRequests.toJSON(e) : undefined
      );
    } else {
      obj.saveRequests = [];
    }
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageResponse.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllSaveRequestsResponse>
  ): QueryAllSaveRequestsResponse {
    const message = {
      ...baseQueryAllSaveRequestsResponse,
    } as QueryAllSaveRequestsResponse;
    message.saveRequests = [];
    if (object.saveRequests !== undefined && object.saveRequests !== null) {
      for (const e of object.saveRequests) {
        message.saveRequests.push(SaveRequests.fromPartial(e));
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

const baseQueryGetMinersRequest: object = { address: "" };

export const QueryGetMinersRequest = {
  encode(
    message: QueryGetMinersRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.address !== "") {
      writer.uint32(10).string(message.address);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetMinersRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryGetMinersRequest } as QueryGetMinersRequest;
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

  fromJSON(object: any): QueryGetMinersRequest {
    const message = { ...baseQueryGetMinersRequest } as QueryGetMinersRequest;
    if (object.address !== undefined && object.address !== null) {
      message.address = String(object.address);
    } else {
      message.address = "";
    }
    return message;
  },

  toJSON(message: QueryGetMinersRequest): unknown {
    const obj: any = {};
    message.address !== undefined && (obj.address = message.address);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetMinersRequest>
  ): QueryGetMinersRequest {
    const message = { ...baseQueryGetMinersRequest } as QueryGetMinersRequest;
    if (object.address !== undefined && object.address !== null) {
      message.address = object.address;
    } else {
      message.address = "";
    }
    return message;
  },
};

const baseQueryGetMinersResponse: object = {};

export const QueryGetMinersResponse = {
  encode(
    message: QueryGetMinersResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.miners !== undefined) {
      Miners.encode(message.miners, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetMinersResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryGetMinersResponse } as QueryGetMinersResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.miners = Miners.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetMinersResponse {
    const message = { ...baseQueryGetMinersResponse } as QueryGetMinersResponse;
    if (object.miners !== undefined && object.miners !== null) {
      message.miners = Miners.fromJSON(object.miners);
    } else {
      message.miners = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetMinersResponse): unknown {
    const obj: any = {};
    message.miners !== undefined &&
      (obj.miners = message.miners ? Miners.toJSON(message.miners) : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetMinersResponse>
  ): QueryGetMinersResponse {
    const message = { ...baseQueryGetMinersResponse } as QueryGetMinersResponse;
    if (object.miners !== undefined && object.miners !== null) {
      message.miners = Miners.fromPartial(object.miners);
    } else {
      message.miners = undefined;
    }
    return message;
  },
};

const baseQueryAllMinersRequest: object = {};

export const QueryAllMinersRequest = {
  encode(
    message: QueryAllMinersRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllMinersRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryAllMinersRequest } as QueryAllMinersRequest;
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

  fromJSON(object: any): QueryAllMinersRequest {
    const message = { ...baseQueryAllMinersRequest } as QueryAllMinersRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllMinersRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageRequest.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllMinersRequest>
  ): QueryAllMinersRequest {
    const message = { ...baseQueryAllMinersRequest } as QueryAllMinersRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryAllMinersResponse: object = {};

export const QueryAllMinersResponse = {
  encode(
    message: QueryAllMinersResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.miners) {
      Miners.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(
        message.pagination,
        writer.uint32(18).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllMinersResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryAllMinersResponse } as QueryAllMinersResponse;
    message.miners = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.miners.push(Miners.decode(reader, reader.uint32()));
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

  fromJSON(object: any): QueryAllMinersResponse {
    const message = { ...baseQueryAllMinersResponse } as QueryAllMinersResponse;
    message.miners = [];
    if (object.miners !== undefined && object.miners !== null) {
      for (const e of object.miners) {
        message.miners.push(Miners.fromJSON(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllMinersResponse): unknown {
    const obj: any = {};
    if (message.miners) {
      obj.miners = message.miners.map((e) =>
        e ? Miners.toJSON(e) : undefined
      );
    } else {
      obj.miners = [];
    }
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageResponse.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllMinersResponse>
  ): QueryAllMinersResponse {
    const message = { ...baseQueryAllMinersResponse } as QueryAllMinersResponse;
    message.miners = [];
    if (object.miners !== undefined && object.miners !== null) {
      for (const e of object.miners) {
        message.miners.push(Miners.fromPartial(e));
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

const baseQueryGetMinedRequest: object = { id: 0 };

export const QueryGetMinedRequest = {
  encode(
    message: QueryGetMinedRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.id !== 0) {
      writer.uint32(8).uint64(message.id);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetMinedRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryGetMinedRequest } as QueryGetMinedRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetMinedRequest {
    const message = { ...baseQueryGetMinedRequest } as QueryGetMinedRequest;
    if (object.id !== undefined && object.id !== null) {
      message.id = Number(object.id);
    } else {
      message.id = 0;
    }
    return message;
  },

  toJSON(message: QueryGetMinedRequest): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    return obj;
  },

  fromPartial(object: DeepPartial<QueryGetMinedRequest>): QueryGetMinedRequest {
    const message = { ...baseQueryGetMinedRequest } as QueryGetMinedRequest;
    if (object.id !== undefined && object.id !== null) {
      message.id = object.id;
    } else {
      message.id = 0;
    }
    return message;
  },
};

const baseQueryGetMinedResponse: object = {};

export const QueryGetMinedResponse = {
  encode(
    message: QueryGetMinedResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.Mined !== undefined) {
      Mined.encode(message.Mined, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetMinedResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryGetMinedResponse } as QueryGetMinedResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.Mined = Mined.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetMinedResponse {
    const message = { ...baseQueryGetMinedResponse } as QueryGetMinedResponse;
    if (object.Mined !== undefined && object.Mined !== null) {
      message.Mined = Mined.fromJSON(object.Mined);
    } else {
      message.Mined = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetMinedResponse): unknown {
    const obj: any = {};
    message.Mined !== undefined &&
      (obj.Mined = message.Mined ? Mined.toJSON(message.Mined) : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetMinedResponse>
  ): QueryGetMinedResponse {
    const message = { ...baseQueryGetMinedResponse } as QueryGetMinedResponse;
    if (object.Mined !== undefined && object.Mined !== null) {
      message.Mined = Mined.fromPartial(object.Mined);
    } else {
      message.Mined = undefined;
    }
    return message;
  },
};

const baseQueryAllMinedRequest: object = {};

export const QueryAllMinedRequest = {
  encode(
    message: QueryAllMinedRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllMinedRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryAllMinedRequest } as QueryAllMinedRequest;
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

  fromJSON(object: any): QueryAllMinedRequest {
    const message = { ...baseQueryAllMinedRequest } as QueryAllMinedRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllMinedRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageRequest.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<QueryAllMinedRequest>): QueryAllMinedRequest {
    const message = { ...baseQueryAllMinedRequest } as QueryAllMinedRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryAllMinedResponse: object = {};

export const QueryAllMinedResponse = {
  encode(
    message: QueryAllMinedResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.Mined) {
      Mined.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(
        message.pagination,
        writer.uint32(18).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllMinedResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryAllMinedResponse } as QueryAllMinedResponse;
    message.Mined = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.Mined.push(Mined.decode(reader, reader.uint32()));
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

  fromJSON(object: any): QueryAllMinedResponse {
    const message = { ...baseQueryAllMinedResponse } as QueryAllMinedResponse;
    message.Mined = [];
    if (object.Mined !== undefined && object.Mined !== null) {
      for (const e of object.Mined) {
        message.Mined.push(Mined.fromJSON(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllMinedResponse): unknown {
    const obj: any = {};
    if (message.Mined) {
      obj.Mined = message.Mined.map((e) => (e ? Mined.toJSON(e) : undefined));
    } else {
      obj.Mined = [];
    }
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageResponse.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllMinedResponse>
  ): QueryAllMinedResponse {
    const message = { ...baseQueryAllMinedResponse } as QueryAllMinedResponse;
    message.Mined = [];
    if (object.Mined !== undefined && object.Mined !== null) {
      for (const e of object.Mined) {
        message.Mined.push(Mined.fromPartial(e));
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
  /** Queries a SaveRequests by index. */
  SaveRequests(
    request: QueryGetSaveRequestsRequest
  ): Promise<QueryGetSaveRequestsResponse>;
  /** Queries a list of SaveRequests items. */
  SaveRequestsAll(
    request: QueryAllSaveRequestsRequest
  ): Promise<QueryAllSaveRequestsResponse>;
  /** Queries a Miners by index. */
  Miners(request: QueryGetMinersRequest): Promise<QueryGetMinersResponse>;
  /** Queries a list of Miners items. */
  MinersAll(request: QueryAllMinersRequest): Promise<QueryAllMinersResponse>;
  /** Queries a Mined by id. */
  Mined(request: QueryGetMinedRequest): Promise<QueryGetMinedResponse>;
  /** Queries a list of Mined items. */
  MinedAll(request: QueryAllMinedRequest): Promise<QueryAllMinedResponse>;
}

export class QueryClientImpl implements Query {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse> {
    const data = QueryParamsRequest.encode(request).finish();
    const promise = this.rpc.request(
      "jackaldao.canine.jklmining.Query",
      "Params",
      data
    );
    return promise.then((data) => QueryParamsResponse.decode(new Reader(data)));
  }

  SaveRequests(
    request: QueryGetSaveRequestsRequest
  ): Promise<QueryGetSaveRequestsResponse> {
    const data = QueryGetSaveRequestsRequest.encode(request).finish();
    const promise = this.rpc.request(
      "jackaldao.canine.jklmining.Query",
      "SaveRequests",
      data
    );
    return promise.then((data) =>
      QueryGetSaveRequestsResponse.decode(new Reader(data))
    );
  }

  SaveRequestsAll(
    request: QueryAllSaveRequestsRequest
  ): Promise<QueryAllSaveRequestsResponse> {
    const data = QueryAllSaveRequestsRequest.encode(request).finish();
    const promise = this.rpc.request(
      "jackaldao.canine.jklmining.Query",
      "SaveRequestsAll",
      data
    );
    return promise.then((data) =>
      QueryAllSaveRequestsResponse.decode(new Reader(data))
    );
  }

  Miners(request: QueryGetMinersRequest): Promise<QueryGetMinersResponse> {
    const data = QueryGetMinersRequest.encode(request).finish();
    const promise = this.rpc.request(
      "jackaldao.canine.jklmining.Query",
      "Miners",
      data
    );
    return promise.then((data) =>
      QueryGetMinersResponse.decode(new Reader(data))
    );
  }

  MinersAll(request: QueryAllMinersRequest): Promise<QueryAllMinersResponse> {
    const data = QueryAllMinersRequest.encode(request).finish();
    const promise = this.rpc.request(
      "jackaldao.canine.jklmining.Query",
      "MinersAll",
      data
    );
    return promise.then((data) =>
      QueryAllMinersResponse.decode(new Reader(data))
    );
  }

  Mined(request: QueryGetMinedRequest): Promise<QueryGetMinedResponse> {
    const data = QueryGetMinedRequest.encode(request).finish();
    const promise = this.rpc.request(
      "jackaldao.canine.jklmining.Query",
      "Mined",
      data
    );
    return promise.then((data) =>
      QueryGetMinedResponse.decode(new Reader(data))
    );
  }

  MinedAll(request: QueryAllMinedRequest): Promise<QueryAllMinedResponse> {
    const data = QueryAllMinedRequest.encode(request).finish();
    const promise = this.rpc.request(
      "jackaldao.canine.jklmining.Query",
      "MinedAll",
      data
    );
    return promise.then((data) =>
      QueryAllMinedResponse.decode(new Reader(data))
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
