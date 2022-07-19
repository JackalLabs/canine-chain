/* eslint-disable */
import { Reader, Writer } from "protobufjs/minimal";
import { Params } from "../storage/params";
import { Contracts } from "../storage/contracts";
import {
  PageRequest,
  PageResponse,
} from "../cosmos/base/query/v1beta1/pagination";
import { Proofs } from "../storage/proofs";
import { ActiveDeals } from "../storage/active_deals";
import { Miners } from "../storage/miners";

export const protobufPackage = "jackaldao.canine.storage";

/** QueryParamsRequest is request type for the Query/Params RPC method. */
export interface QueryParamsRequest {}

/** QueryParamsResponse is response type for the Query/Params RPC method. */
export interface QueryParamsResponse {
  /** params holds all the parameters of this module. */
  params: Params | undefined;
}

export interface QueryGetContractsRequest {
  cid: string;
}

export interface QueryGetContractsResponse {
  contracts: Contracts | undefined;
}

export interface QueryAllContractsRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllContractsResponse {
  contracts: Contracts[];
  pagination: PageResponse | undefined;
}

export interface QueryGetProofsRequest {
  cid: string;
}

export interface QueryGetProofsResponse {
  proofs: Proofs | undefined;
}

export interface QueryAllProofsRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllProofsResponse {
  proofs: Proofs[];
  pagination: PageResponse | undefined;
}

export interface QueryGetActiveDealsRequest {
  cid: string;
}

export interface QueryGetActiveDealsResponse {
  activeDeals: ActiveDeals | undefined;
}

export interface QueryAllActiveDealsRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllActiveDealsResponse {
  activeDeals: ActiveDeals[];
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

export interface QueryFreespaceRequest {
  address: string;
}

export interface QueryFreespaceResponse {
  space: string;
}

export interface QueryFindFileRequest {
  fid: string;
}

export interface QueryFindFileResponse {
  minerIps: string;
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

const baseQueryGetContractsRequest: object = { cid: "" };

export const QueryGetContractsRequest = {
  encode(
    message: QueryGetContractsRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.cid !== "") {
      writer.uint32(10).string(message.cid);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetContractsRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetContractsRequest,
    } as QueryGetContractsRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.cid = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetContractsRequest {
    const message = {
      ...baseQueryGetContractsRequest,
    } as QueryGetContractsRequest;
    if (object.cid !== undefined && object.cid !== null) {
      message.cid = String(object.cid);
    } else {
      message.cid = "";
    }
    return message;
  },

  toJSON(message: QueryGetContractsRequest): unknown {
    const obj: any = {};
    message.cid !== undefined && (obj.cid = message.cid);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetContractsRequest>
  ): QueryGetContractsRequest {
    const message = {
      ...baseQueryGetContractsRequest,
    } as QueryGetContractsRequest;
    if (object.cid !== undefined && object.cid !== null) {
      message.cid = object.cid;
    } else {
      message.cid = "";
    }
    return message;
  },
};

const baseQueryGetContractsResponse: object = {};

export const QueryGetContractsResponse = {
  encode(
    message: QueryGetContractsResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.contracts !== undefined) {
      Contracts.encode(message.contracts, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetContractsResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetContractsResponse,
    } as QueryGetContractsResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.contracts = Contracts.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetContractsResponse {
    const message = {
      ...baseQueryGetContractsResponse,
    } as QueryGetContractsResponse;
    if (object.contracts !== undefined && object.contracts !== null) {
      message.contracts = Contracts.fromJSON(object.contracts);
    } else {
      message.contracts = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetContractsResponse): unknown {
    const obj: any = {};
    message.contracts !== undefined &&
      (obj.contracts = message.contracts
        ? Contracts.toJSON(message.contracts)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetContractsResponse>
  ): QueryGetContractsResponse {
    const message = {
      ...baseQueryGetContractsResponse,
    } as QueryGetContractsResponse;
    if (object.contracts !== undefined && object.contracts !== null) {
      message.contracts = Contracts.fromPartial(object.contracts);
    } else {
      message.contracts = undefined;
    }
    return message;
  },
};

const baseQueryAllContractsRequest: object = {};

export const QueryAllContractsRequest = {
  encode(
    message: QueryAllContractsRequest,
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
  ): QueryAllContractsRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllContractsRequest,
    } as QueryAllContractsRequest;
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

  fromJSON(object: any): QueryAllContractsRequest {
    const message = {
      ...baseQueryAllContractsRequest,
    } as QueryAllContractsRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllContractsRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageRequest.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllContractsRequest>
  ): QueryAllContractsRequest {
    const message = {
      ...baseQueryAllContractsRequest,
    } as QueryAllContractsRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryAllContractsResponse: object = {};

export const QueryAllContractsResponse = {
  encode(
    message: QueryAllContractsResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.contracts) {
      Contracts.encode(v!, writer.uint32(10).fork()).ldelim();
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
  ): QueryAllContractsResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllContractsResponse,
    } as QueryAllContractsResponse;
    message.contracts = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.contracts.push(Contracts.decode(reader, reader.uint32()));
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

  fromJSON(object: any): QueryAllContractsResponse {
    const message = {
      ...baseQueryAllContractsResponse,
    } as QueryAllContractsResponse;
    message.contracts = [];
    if (object.contracts !== undefined && object.contracts !== null) {
      for (const e of object.contracts) {
        message.contracts.push(Contracts.fromJSON(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllContractsResponse): unknown {
    const obj: any = {};
    if (message.contracts) {
      obj.contracts = message.contracts.map((e) =>
        e ? Contracts.toJSON(e) : undefined
      );
    } else {
      obj.contracts = [];
    }
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageResponse.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllContractsResponse>
  ): QueryAllContractsResponse {
    const message = {
      ...baseQueryAllContractsResponse,
    } as QueryAllContractsResponse;
    message.contracts = [];
    if (object.contracts !== undefined && object.contracts !== null) {
      for (const e of object.contracts) {
        message.contracts.push(Contracts.fromPartial(e));
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

const baseQueryGetProofsRequest: object = { cid: "" };

export const QueryGetProofsRequest = {
  encode(
    message: QueryGetProofsRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.cid !== "") {
      writer.uint32(10).string(message.cid);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetProofsRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryGetProofsRequest } as QueryGetProofsRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.cid = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetProofsRequest {
    const message = { ...baseQueryGetProofsRequest } as QueryGetProofsRequest;
    if (object.cid !== undefined && object.cid !== null) {
      message.cid = String(object.cid);
    } else {
      message.cid = "";
    }
    return message;
  },

  toJSON(message: QueryGetProofsRequest): unknown {
    const obj: any = {};
    message.cid !== undefined && (obj.cid = message.cid);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetProofsRequest>
  ): QueryGetProofsRequest {
    const message = { ...baseQueryGetProofsRequest } as QueryGetProofsRequest;
    if (object.cid !== undefined && object.cid !== null) {
      message.cid = object.cid;
    } else {
      message.cid = "";
    }
    return message;
  },
};

const baseQueryGetProofsResponse: object = {};

export const QueryGetProofsResponse = {
  encode(
    message: QueryGetProofsResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.proofs !== undefined) {
      Proofs.encode(message.proofs, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetProofsResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryGetProofsResponse } as QueryGetProofsResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.proofs = Proofs.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetProofsResponse {
    const message = { ...baseQueryGetProofsResponse } as QueryGetProofsResponse;
    if (object.proofs !== undefined && object.proofs !== null) {
      message.proofs = Proofs.fromJSON(object.proofs);
    } else {
      message.proofs = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetProofsResponse): unknown {
    const obj: any = {};
    message.proofs !== undefined &&
      (obj.proofs = message.proofs ? Proofs.toJSON(message.proofs) : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetProofsResponse>
  ): QueryGetProofsResponse {
    const message = { ...baseQueryGetProofsResponse } as QueryGetProofsResponse;
    if (object.proofs !== undefined && object.proofs !== null) {
      message.proofs = Proofs.fromPartial(object.proofs);
    } else {
      message.proofs = undefined;
    }
    return message;
  },
};

const baseQueryAllProofsRequest: object = {};

export const QueryAllProofsRequest = {
  encode(
    message: QueryAllProofsRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllProofsRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryAllProofsRequest } as QueryAllProofsRequest;
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

  fromJSON(object: any): QueryAllProofsRequest {
    const message = { ...baseQueryAllProofsRequest } as QueryAllProofsRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllProofsRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageRequest.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllProofsRequest>
  ): QueryAllProofsRequest {
    const message = { ...baseQueryAllProofsRequest } as QueryAllProofsRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryAllProofsResponse: object = {};

export const QueryAllProofsResponse = {
  encode(
    message: QueryAllProofsResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.proofs) {
      Proofs.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(
        message.pagination,
        writer.uint32(18).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllProofsResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryAllProofsResponse } as QueryAllProofsResponse;
    message.proofs = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.proofs.push(Proofs.decode(reader, reader.uint32()));
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

  fromJSON(object: any): QueryAllProofsResponse {
    const message = { ...baseQueryAllProofsResponse } as QueryAllProofsResponse;
    message.proofs = [];
    if (object.proofs !== undefined && object.proofs !== null) {
      for (const e of object.proofs) {
        message.proofs.push(Proofs.fromJSON(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllProofsResponse): unknown {
    const obj: any = {};
    if (message.proofs) {
      obj.proofs = message.proofs.map((e) =>
        e ? Proofs.toJSON(e) : undefined
      );
    } else {
      obj.proofs = [];
    }
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageResponse.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllProofsResponse>
  ): QueryAllProofsResponse {
    const message = { ...baseQueryAllProofsResponse } as QueryAllProofsResponse;
    message.proofs = [];
    if (object.proofs !== undefined && object.proofs !== null) {
      for (const e of object.proofs) {
        message.proofs.push(Proofs.fromPartial(e));
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

const baseQueryGetActiveDealsRequest: object = { cid: "" };

export const QueryGetActiveDealsRequest = {
  encode(
    message: QueryGetActiveDealsRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.cid !== "") {
      writer.uint32(10).string(message.cid);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetActiveDealsRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetActiveDealsRequest,
    } as QueryGetActiveDealsRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.cid = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetActiveDealsRequest {
    const message = {
      ...baseQueryGetActiveDealsRequest,
    } as QueryGetActiveDealsRequest;
    if (object.cid !== undefined && object.cid !== null) {
      message.cid = String(object.cid);
    } else {
      message.cid = "";
    }
    return message;
  },

  toJSON(message: QueryGetActiveDealsRequest): unknown {
    const obj: any = {};
    message.cid !== undefined && (obj.cid = message.cid);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetActiveDealsRequest>
  ): QueryGetActiveDealsRequest {
    const message = {
      ...baseQueryGetActiveDealsRequest,
    } as QueryGetActiveDealsRequest;
    if (object.cid !== undefined && object.cid !== null) {
      message.cid = object.cid;
    } else {
      message.cid = "";
    }
    return message;
  },
};

const baseQueryGetActiveDealsResponse: object = {};

export const QueryGetActiveDealsResponse = {
  encode(
    message: QueryGetActiveDealsResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.activeDeals !== undefined) {
      ActiveDeals.encode(
        message.activeDeals,
        writer.uint32(10).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetActiveDealsResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetActiveDealsResponse,
    } as QueryGetActiveDealsResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.activeDeals = ActiveDeals.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetActiveDealsResponse {
    const message = {
      ...baseQueryGetActiveDealsResponse,
    } as QueryGetActiveDealsResponse;
    if (object.activeDeals !== undefined && object.activeDeals !== null) {
      message.activeDeals = ActiveDeals.fromJSON(object.activeDeals);
    } else {
      message.activeDeals = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetActiveDealsResponse): unknown {
    const obj: any = {};
    message.activeDeals !== undefined &&
      (obj.activeDeals = message.activeDeals
        ? ActiveDeals.toJSON(message.activeDeals)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetActiveDealsResponse>
  ): QueryGetActiveDealsResponse {
    const message = {
      ...baseQueryGetActiveDealsResponse,
    } as QueryGetActiveDealsResponse;
    if (object.activeDeals !== undefined && object.activeDeals !== null) {
      message.activeDeals = ActiveDeals.fromPartial(object.activeDeals);
    } else {
      message.activeDeals = undefined;
    }
    return message;
  },
};

const baseQueryAllActiveDealsRequest: object = {};

export const QueryAllActiveDealsRequest = {
  encode(
    message: QueryAllActiveDealsRequest,
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
  ): QueryAllActiveDealsRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllActiveDealsRequest,
    } as QueryAllActiveDealsRequest;
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

  fromJSON(object: any): QueryAllActiveDealsRequest {
    const message = {
      ...baseQueryAllActiveDealsRequest,
    } as QueryAllActiveDealsRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllActiveDealsRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageRequest.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllActiveDealsRequest>
  ): QueryAllActiveDealsRequest {
    const message = {
      ...baseQueryAllActiveDealsRequest,
    } as QueryAllActiveDealsRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryAllActiveDealsResponse: object = {};

export const QueryAllActiveDealsResponse = {
  encode(
    message: QueryAllActiveDealsResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.activeDeals) {
      ActiveDeals.encode(v!, writer.uint32(10).fork()).ldelim();
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
  ): QueryAllActiveDealsResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllActiveDealsResponse,
    } as QueryAllActiveDealsResponse;
    message.activeDeals = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.activeDeals.push(ActiveDeals.decode(reader, reader.uint32()));
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

  fromJSON(object: any): QueryAllActiveDealsResponse {
    const message = {
      ...baseQueryAllActiveDealsResponse,
    } as QueryAllActiveDealsResponse;
    message.activeDeals = [];
    if (object.activeDeals !== undefined && object.activeDeals !== null) {
      for (const e of object.activeDeals) {
        message.activeDeals.push(ActiveDeals.fromJSON(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllActiveDealsResponse): unknown {
    const obj: any = {};
    if (message.activeDeals) {
      obj.activeDeals = message.activeDeals.map((e) =>
        e ? ActiveDeals.toJSON(e) : undefined
      );
    } else {
      obj.activeDeals = [];
    }
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageResponse.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllActiveDealsResponse>
  ): QueryAllActiveDealsResponse {
    const message = {
      ...baseQueryAllActiveDealsResponse,
    } as QueryAllActiveDealsResponse;
    message.activeDeals = [];
    if (object.activeDeals !== undefined && object.activeDeals !== null) {
      for (const e of object.activeDeals) {
        message.activeDeals.push(ActiveDeals.fromPartial(e));
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

const baseQueryFreespaceRequest: object = { address: "" };

export const QueryFreespaceRequest = {
  encode(
    message: QueryFreespaceRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.address !== "") {
      writer.uint32(10).string(message.address);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryFreespaceRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryFreespaceRequest } as QueryFreespaceRequest;
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

  fromJSON(object: any): QueryFreespaceRequest {
    const message = { ...baseQueryFreespaceRequest } as QueryFreespaceRequest;
    if (object.address !== undefined && object.address !== null) {
      message.address = String(object.address);
    } else {
      message.address = "";
    }
    return message;
  },

  toJSON(message: QueryFreespaceRequest): unknown {
    const obj: any = {};
    message.address !== undefined && (obj.address = message.address);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryFreespaceRequest>
  ): QueryFreespaceRequest {
    const message = { ...baseQueryFreespaceRequest } as QueryFreespaceRequest;
    if (object.address !== undefined && object.address !== null) {
      message.address = object.address;
    } else {
      message.address = "";
    }
    return message;
  },
};

const baseQueryFreespaceResponse: object = { space: "" };

export const QueryFreespaceResponse = {
  encode(
    message: QueryFreespaceResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.space !== "") {
      writer.uint32(10).string(message.space);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryFreespaceResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryFreespaceResponse } as QueryFreespaceResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.space = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryFreespaceResponse {
    const message = { ...baseQueryFreespaceResponse } as QueryFreespaceResponse;
    if (object.space !== undefined && object.space !== null) {
      message.space = String(object.space);
    } else {
      message.space = "";
    }
    return message;
  },

  toJSON(message: QueryFreespaceResponse): unknown {
    const obj: any = {};
    message.space !== undefined && (obj.space = message.space);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryFreespaceResponse>
  ): QueryFreespaceResponse {
    const message = { ...baseQueryFreespaceResponse } as QueryFreespaceResponse;
    if (object.space !== undefined && object.space !== null) {
      message.space = object.space;
    } else {
      message.space = "";
    }
    return message;
  },
};

const baseQueryFindFileRequest: object = { fid: "" };

export const QueryFindFileRequest = {
  encode(
    message: QueryFindFileRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.fid !== "") {
      writer.uint32(10).string(message.fid);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryFindFileRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryFindFileRequest } as QueryFindFileRequest;
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

  fromJSON(object: any): QueryFindFileRequest {
    const message = { ...baseQueryFindFileRequest } as QueryFindFileRequest;
    if (object.fid !== undefined && object.fid !== null) {
      message.fid = String(object.fid);
    } else {
      message.fid = "";
    }
    return message;
  },

  toJSON(message: QueryFindFileRequest): unknown {
    const obj: any = {};
    message.fid !== undefined && (obj.fid = message.fid);
    return obj;
  },

  fromPartial(object: DeepPartial<QueryFindFileRequest>): QueryFindFileRequest {
    const message = { ...baseQueryFindFileRequest } as QueryFindFileRequest;
    if (object.fid !== undefined && object.fid !== null) {
      message.fid = object.fid;
    } else {
      message.fid = "";
    }
    return message;
  },
};

const baseQueryFindFileResponse: object = { minerIps: "" };

export const QueryFindFileResponse = {
  encode(
    message: QueryFindFileResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.minerIps !== "") {
      writer.uint32(10).string(message.minerIps);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryFindFileResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryFindFileResponse } as QueryFindFileResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.minerIps = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryFindFileResponse {
    const message = { ...baseQueryFindFileResponse } as QueryFindFileResponse;
    if (object.minerIps !== undefined && object.minerIps !== null) {
      message.minerIps = String(object.minerIps);
    } else {
      message.minerIps = "";
    }
    return message;
  },

  toJSON(message: QueryFindFileResponse): unknown {
    const obj: any = {};
    message.minerIps !== undefined && (obj.minerIps = message.minerIps);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryFindFileResponse>
  ): QueryFindFileResponse {
    const message = { ...baseQueryFindFileResponse } as QueryFindFileResponse;
    if (object.minerIps !== undefined && object.minerIps !== null) {
      message.minerIps = object.minerIps;
    } else {
      message.minerIps = "";
    }
    return message;
  },
};

/** Query defines the gRPC querier service. */
export interface Query {
  /** Parameters queries the parameters of the module. */
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse>;
  /** Queries a Contracts by index. */
  Contracts(
    request: QueryGetContractsRequest
  ): Promise<QueryGetContractsResponse>;
  /** Queries a list of Contracts items. */
  ContractsAll(
    request: QueryAllContractsRequest
  ): Promise<QueryAllContractsResponse>;
  /** Queries a Proofs by index. */
  Proofs(request: QueryGetProofsRequest): Promise<QueryGetProofsResponse>;
  /** Queries a list of Proofs items. */
  ProofsAll(request: QueryAllProofsRequest): Promise<QueryAllProofsResponse>;
  /** Queries a ActiveDeals by index. */
  ActiveDeals(
    request: QueryGetActiveDealsRequest
  ): Promise<QueryGetActiveDealsResponse>;
  /** Queries a list of ActiveDeals items. */
  ActiveDealsAll(
    request: QueryAllActiveDealsRequest
  ): Promise<QueryAllActiveDealsResponse>;
  /** Queries a Miners by index. */
  Miners(request: QueryGetMinersRequest): Promise<QueryGetMinersResponse>;
  /** Queries a list of Miners items. */
  MinersAll(request: QueryAllMinersRequest): Promise<QueryAllMinersResponse>;
  /** Queries a list of Freespace items. */
  Freespace(request: QueryFreespaceRequest): Promise<QueryFreespaceResponse>;
  /** Queries a list of FindFile items. */
  FindFile(request: QueryFindFileRequest): Promise<QueryFindFileResponse>;
}

export class QueryClientImpl implements Query {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse> {
    const data = QueryParamsRequest.encode(request).finish();
    const promise = this.rpc.request(
      "jackaldao.canine.storage.Query",
      "Params",
      data
    );
    return promise.then((data) => QueryParamsResponse.decode(new Reader(data)));
  }

  Contracts(
    request: QueryGetContractsRequest
  ): Promise<QueryGetContractsResponse> {
    const data = QueryGetContractsRequest.encode(request).finish();
    const promise = this.rpc.request(
      "jackaldao.canine.storage.Query",
      "Contracts",
      data
    );
    return promise.then((data) =>
      QueryGetContractsResponse.decode(new Reader(data))
    );
  }

  ContractsAll(
    request: QueryAllContractsRequest
  ): Promise<QueryAllContractsResponse> {
    const data = QueryAllContractsRequest.encode(request).finish();
    const promise = this.rpc.request(
      "jackaldao.canine.storage.Query",
      "ContractsAll",
      data
    );
    return promise.then((data) =>
      QueryAllContractsResponse.decode(new Reader(data))
    );
  }

  Proofs(request: QueryGetProofsRequest): Promise<QueryGetProofsResponse> {
    const data = QueryGetProofsRequest.encode(request).finish();
    const promise = this.rpc.request(
      "jackaldao.canine.storage.Query",
      "Proofs",
      data
    );
    return promise.then((data) =>
      QueryGetProofsResponse.decode(new Reader(data))
    );
  }

  ProofsAll(request: QueryAllProofsRequest): Promise<QueryAllProofsResponse> {
    const data = QueryAllProofsRequest.encode(request).finish();
    const promise = this.rpc.request(
      "jackaldao.canine.storage.Query",
      "ProofsAll",
      data
    );
    return promise.then((data) =>
      QueryAllProofsResponse.decode(new Reader(data))
    );
  }

  ActiveDeals(
    request: QueryGetActiveDealsRequest
  ): Promise<QueryGetActiveDealsResponse> {
    const data = QueryGetActiveDealsRequest.encode(request).finish();
    const promise = this.rpc.request(
      "jackaldao.canine.storage.Query",
      "ActiveDeals",
      data
    );
    return promise.then((data) =>
      QueryGetActiveDealsResponse.decode(new Reader(data))
    );
  }

  ActiveDealsAll(
    request: QueryAllActiveDealsRequest
  ): Promise<QueryAllActiveDealsResponse> {
    const data = QueryAllActiveDealsRequest.encode(request).finish();
    const promise = this.rpc.request(
      "jackaldao.canine.storage.Query",
      "ActiveDealsAll",
      data
    );
    return promise.then((data) =>
      QueryAllActiveDealsResponse.decode(new Reader(data))
    );
  }

  Miners(request: QueryGetMinersRequest): Promise<QueryGetMinersResponse> {
    const data = QueryGetMinersRequest.encode(request).finish();
    const promise = this.rpc.request(
      "jackaldao.canine.storage.Query",
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
      "jackaldao.canine.storage.Query",
      "MinersAll",
      data
    );
    return promise.then((data) =>
      QueryAllMinersResponse.decode(new Reader(data))
    );
  }

  Freespace(request: QueryFreespaceRequest): Promise<QueryFreespaceResponse> {
    const data = QueryFreespaceRequest.encode(request).finish();
    const promise = this.rpc.request(
      "jackaldao.canine.storage.Query",
      "Freespace",
      data
    );
    return promise.then((data) =>
      QueryFreespaceResponse.decode(new Reader(data))
    );
  }

  FindFile(request: QueryFindFileRequest): Promise<QueryFindFileResponse> {
    const data = QueryFindFileRequest.encode(request).finish();
    const promise = this.rpc.request(
      "jackaldao.canine.storage.Query",
      "FindFile",
      data
    );
    return promise.then((data) =>
      QueryFindFileResponse.decode(new Reader(data))
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
