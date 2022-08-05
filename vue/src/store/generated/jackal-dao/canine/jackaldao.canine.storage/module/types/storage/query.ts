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
import { PayBlocks } from "../storage/pay_blocks";
import { ClientUsage } from "../storage/client_usage";
import { Strays } from "../storage/strays";

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

export interface QueryGetPayBlocksRequest {
  blockid: string;
}

export interface QueryGetPayBlocksResponse {
  payBlocks: PayBlocks | undefined;
}

export interface QueryAllPayBlocksRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllPayBlocksResponse {
  payBlocks: PayBlocks[];
  pagination: PageResponse | undefined;
}

export interface QueryGetClientUsageRequest {
  address: string;
}

export interface QueryGetClientUsageResponse {
  clientUsage: ClientUsage | undefined;
}

export interface QueryAllClientUsageRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllClientUsageResponse {
  clientUsage: ClientUsage[];
  pagination: PageResponse | undefined;
}

export interface QueryGetStraysRequest {
  cid: string;
}

export interface QueryGetStraysResponse {
  strays: Strays | undefined;
}

export interface QueryAllStraysRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllStraysResponse {
  strays: Strays[];
  pagination: PageResponse | undefined;
}

export interface QueryGetClientFreeSpaceRequest {
  address: string;
}

export interface QueryGetClientFreeSpaceResponse {
  bytesfree: string;
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

const baseQueryGetPayBlocksRequest: object = { blockid: "" };

export const QueryGetPayBlocksRequest = {
  encode(
    message: QueryGetPayBlocksRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.blockid !== "") {
      writer.uint32(10).string(message.blockid);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetPayBlocksRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetPayBlocksRequest,
    } as QueryGetPayBlocksRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.blockid = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetPayBlocksRequest {
    const message = {
      ...baseQueryGetPayBlocksRequest,
    } as QueryGetPayBlocksRequest;
    if (object.blockid !== undefined && object.blockid !== null) {
      message.blockid = String(object.blockid);
    } else {
      message.blockid = "";
    }
    return message;
  },

  toJSON(message: QueryGetPayBlocksRequest): unknown {
    const obj: any = {};
    message.blockid !== undefined && (obj.blockid = message.blockid);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetPayBlocksRequest>
  ): QueryGetPayBlocksRequest {
    const message = {
      ...baseQueryGetPayBlocksRequest,
    } as QueryGetPayBlocksRequest;
    if (object.blockid !== undefined && object.blockid !== null) {
      message.blockid = object.blockid;
    } else {
      message.blockid = "";
    }
    return message;
  },
};

const baseQueryGetPayBlocksResponse: object = {};

export const QueryGetPayBlocksResponse = {
  encode(
    message: QueryGetPayBlocksResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.payBlocks !== undefined) {
      PayBlocks.encode(message.payBlocks, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetPayBlocksResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetPayBlocksResponse,
    } as QueryGetPayBlocksResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.payBlocks = PayBlocks.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetPayBlocksResponse {
    const message = {
      ...baseQueryGetPayBlocksResponse,
    } as QueryGetPayBlocksResponse;
    if (object.payBlocks !== undefined && object.payBlocks !== null) {
      message.payBlocks = PayBlocks.fromJSON(object.payBlocks);
    } else {
      message.payBlocks = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetPayBlocksResponse): unknown {
    const obj: any = {};
    message.payBlocks !== undefined &&
      (obj.payBlocks = message.payBlocks
        ? PayBlocks.toJSON(message.payBlocks)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetPayBlocksResponse>
  ): QueryGetPayBlocksResponse {
    const message = {
      ...baseQueryGetPayBlocksResponse,
    } as QueryGetPayBlocksResponse;
    if (object.payBlocks !== undefined && object.payBlocks !== null) {
      message.payBlocks = PayBlocks.fromPartial(object.payBlocks);
    } else {
      message.payBlocks = undefined;
    }
    return message;
  },
};

const baseQueryAllPayBlocksRequest: object = {};

export const QueryAllPayBlocksRequest = {
  encode(
    message: QueryAllPayBlocksRequest,
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
  ): QueryAllPayBlocksRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllPayBlocksRequest,
    } as QueryAllPayBlocksRequest;
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

  fromJSON(object: any): QueryAllPayBlocksRequest {
    const message = {
      ...baseQueryAllPayBlocksRequest,
    } as QueryAllPayBlocksRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllPayBlocksRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageRequest.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllPayBlocksRequest>
  ): QueryAllPayBlocksRequest {
    const message = {
      ...baseQueryAllPayBlocksRequest,
    } as QueryAllPayBlocksRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryAllPayBlocksResponse: object = {};

export const QueryAllPayBlocksResponse = {
  encode(
    message: QueryAllPayBlocksResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.payBlocks) {
      PayBlocks.encode(v!, writer.uint32(10).fork()).ldelim();
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
  ): QueryAllPayBlocksResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllPayBlocksResponse,
    } as QueryAllPayBlocksResponse;
    message.payBlocks = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.payBlocks.push(PayBlocks.decode(reader, reader.uint32()));
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

  fromJSON(object: any): QueryAllPayBlocksResponse {
    const message = {
      ...baseQueryAllPayBlocksResponse,
    } as QueryAllPayBlocksResponse;
    message.payBlocks = [];
    if (object.payBlocks !== undefined && object.payBlocks !== null) {
      for (const e of object.payBlocks) {
        message.payBlocks.push(PayBlocks.fromJSON(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllPayBlocksResponse): unknown {
    const obj: any = {};
    if (message.payBlocks) {
      obj.payBlocks = message.payBlocks.map((e) =>
        e ? PayBlocks.toJSON(e) : undefined
      );
    } else {
      obj.payBlocks = [];
    }
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageResponse.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllPayBlocksResponse>
  ): QueryAllPayBlocksResponse {
    const message = {
      ...baseQueryAllPayBlocksResponse,
    } as QueryAllPayBlocksResponse;
    message.payBlocks = [];
    if (object.payBlocks !== undefined && object.payBlocks !== null) {
      for (const e of object.payBlocks) {
        message.payBlocks.push(PayBlocks.fromPartial(e));
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

const baseQueryGetClientUsageRequest: object = { address: "" };

export const QueryGetClientUsageRequest = {
  encode(
    message: QueryGetClientUsageRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.address !== "") {
      writer.uint32(10).string(message.address);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetClientUsageRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetClientUsageRequest,
    } as QueryGetClientUsageRequest;
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

  fromJSON(object: any): QueryGetClientUsageRequest {
    const message = {
      ...baseQueryGetClientUsageRequest,
    } as QueryGetClientUsageRequest;
    if (object.address !== undefined && object.address !== null) {
      message.address = String(object.address);
    } else {
      message.address = "";
    }
    return message;
  },

  toJSON(message: QueryGetClientUsageRequest): unknown {
    const obj: any = {};
    message.address !== undefined && (obj.address = message.address);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetClientUsageRequest>
  ): QueryGetClientUsageRequest {
    const message = {
      ...baseQueryGetClientUsageRequest,
    } as QueryGetClientUsageRequest;
    if (object.address !== undefined && object.address !== null) {
      message.address = object.address;
    } else {
      message.address = "";
    }
    return message;
  },
};

const baseQueryGetClientUsageResponse: object = {};

export const QueryGetClientUsageResponse = {
  encode(
    message: QueryGetClientUsageResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.clientUsage !== undefined) {
      ClientUsage.encode(
        message.clientUsage,
        writer.uint32(10).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetClientUsageResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetClientUsageResponse,
    } as QueryGetClientUsageResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.clientUsage = ClientUsage.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetClientUsageResponse {
    const message = {
      ...baseQueryGetClientUsageResponse,
    } as QueryGetClientUsageResponse;
    if (object.clientUsage !== undefined && object.clientUsage !== null) {
      message.clientUsage = ClientUsage.fromJSON(object.clientUsage);
    } else {
      message.clientUsage = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetClientUsageResponse): unknown {
    const obj: any = {};
    message.clientUsage !== undefined &&
      (obj.clientUsage = message.clientUsage
        ? ClientUsage.toJSON(message.clientUsage)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetClientUsageResponse>
  ): QueryGetClientUsageResponse {
    const message = {
      ...baseQueryGetClientUsageResponse,
    } as QueryGetClientUsageResponse;
    if (object.clientUsage !== undefined && object.clientUsage !== null) {
      message.clientUsage = ClientUsage.fromPartial(object.clientUsage);
    } else {
      message.clientUsage = undefined;
    }
    return message;
  },
};

const baseQueryAllClientUsageRequest: object = {};

export const QueryAllClientUsageRequest = {
  encode(
    message: QueryAllClientUsageRequest,
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
  ): QueryAllClientUsageRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllClientUsageRequest,
    } as QueryAllClientUsageRequest;
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

  fromJSON(object: any): QueryAllClientUsageRequest {
    const message = {
      ...baseQueryAllClientUsageRequest,
    } as QueryAllClientUsageRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllClientUsageRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageRequest.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllClientUsageRequest>
  ): QueryAllClientUsageRequest {
    const message = {
      ...baseQueryAllClientUsageRequest,
    } as QueryAllClientUsageRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryAllClientUsageResponse: object = {};

export const QueryAllClientUsageResponse = {
  encode(
    message: QueryAllClientUsageResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.clientUsage) {
      ClientUsage.encode(v!, writer.uint32(10).fork()).ldelim();
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
  ): QueryAllClientUsageResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllClientUsageResponse,
    } as QueryAllClientUsageResponse;
    message.clientUsage = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.clientUsage.push(ClientUsage.decode(reader, reader.uint32()));
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

  fromJSON(object: any): QueryAllClientUsageResponse {
    const message = {
      ...baseQueryAllClientUsageResponse,
    } as QueryAllClientUsageResponse;
    message.clientUsage = [];
    if (object.clientUsage !== undefined && object.clientUsage !== null) {
      for (const e of object.clientUsage) {
        message.clientUsage.push(ClientUsage.fromJSON(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllClientUsageResponse): unknown {
    const obj: any = {};
    if (message.clientUsage) {
      obj.clientUsage = message.clientUsage.map((e) =>
        e ? ClientUsage.toJSON(e) : undefined
      );
    } else {
      obj.clientUsage = [];
    }
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageResponse.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllClientUsageResponse>
  ): QueryAllClientUsageResponse {
    const message = {
      ...baseQueryAllClientUsageResponse,
    } as QueryAllClientUsageResponse;
    message.clientUsage = [];
    if (object.clientUsage !== undefined && object.clientUsage !== null) {
      for (const e of object.clientUsage) {
        message.clientUsage.push(ClientUsage.fromPartial(e));
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

const baseQueryGetStraysRequest: object = { cid: "" };

export const QueryGetStraysRequest = {
  encode(
    message: QueryGetStraysRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.cid !== "") {
      writer.uint32(10).string(message.cid);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetStraysRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryGetStraysRequest } as QueryGetStraysRequest;
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

  fromJSON(object: any): QueryGetStraysRequest {
    const message = { ...baseQueryGetStraysRequest } as QueryGetStraysRequest;
    if (object.cid !== undefined && object.cid !== null) {
      message.cid = String(object.cid);
    } else {
      message.cid = "";
    }
    return message;
  },

  toJSON(message: QueryGetStraysRequest): unknown {
    const obj: any = {};
    message.cid !== undefined && (obj.cid = message.cid);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetStraysRequest>
  ): QueryGetStraysRequest {
    const message = { ...baseQueryGetStraysRequest } as QueryGetStraysRequest;
    if (object.cid !== undefined && object.cid !== null) {
      message.cid = object.cid;
    } else {
      message.cid = "";
    }
    return message;
  },
};

const baseQueryGetStraysResponse: object = {};

export const QueryGetStraysResponse = {
  encode(
    message: QueryGetStraysResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.strays !== undefined) {
      Strays.encode(message.strays, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetStraysResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryGetStraysResponse } as QueryGetStraysResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.strays = Strays.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetStraysResponse {
    const message = { ...baseQueryGetStraysResponse } as QueryGetStraysResponse;
    if (object.strays !== undefined && object.strays !== null) {
      message.strays = Strays.fromJSON(object.strays);
    } else {
      message.strays = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetStraysResponse): unknown {
    const obj: any = {};
    message.strays !== undefined &&
      (obj.strays = message.strays ? Strays.toJSON(message.strays) : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetStraysResponse>
  ): QueryGetStraysResponse {
    const message = { ...baseQueryGetStraysResponse } as QueryGetStraysResponse;
    if (object.strays !== undefined && object.strays !== null) {
      message.strays = Strays.fromPartial(object.strays);
    } else {
      message.strays = undefined;
    }
    return message;
  },
};

const baseQueryAllStraysRequest: object = {};

export const QueryAllStraysRequest = {
  encode(
    message: QueryAllStraysRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllStraysRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryAllStraysRequest } as QueryAllStraysRequest;
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

  fromJSON(object: any): QueryAllStraysRequest {
    const message = { ...baseQueryAllStraysRequest } as QueryAllStraysRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllStraysRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageRequest.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllStraysRequest>
  ): QueryAllStraysRequest {
    const message = { ...baseQueryAllStraysRequest } as QueryAllStraysRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryAllStraysResponse: object = {};

export const QueryAllStraysResponse = {
  encode(
    message: QueryAllStraysResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.strays) {
      Strays.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(
        message.pagination,
        writer.uint32(18).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllStraysResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryAllStraysResponse } as QueryAllStraysResponse;
    message.strays = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.strays.push(Strays.decode(reader, reader.uint32()));
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

  fromJSON(object: any): QueryAllStraysResponse {
    const message = { ...baseQueryAllStraysResponse } as QueryAllStraysResponse;
    message.strays = [];
    if (object.strays !== undefined && object.strays !== null) {
      for (const e of object.strays) {
        message.strays.push(Strays.fromJSON(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllStraysResponse): unknown {
    const obj: any = {};
    if (message.strays) {
      obj.strays = message.strays.map((e) =>
        e ? Strays.toJSON(e) : undefined
      );
    } else {
      obj.strays = [];
    }
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageResponse.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllStraysResponse>
  ): QueryAllStraysResponse {
    const message = { ...baseQueryAllStraysResponse } as QueryAllStraysResponse;
    message.strays = [];
    if (object.strays !== undefined && object.strays !== null) {
      for (const e of object.strays) {
        message.strays.push(Strays.fromPartial(e));
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

const baseQueryGetClientFreeSpaceRequest: object = { address: "" };

export const QueryGetClientFreeSpaceRequest = {
  encode(
    message: QueryGetClientFreeSpaceRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.address !== "") {
      writer.uint32(10).string(message.address);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetClientFreeSpaceRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetClientFreeSpaceRequest,
    } as QueryGetClientFreeSpaceRequest;
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

  fromJSON(object: any): QueryGetClientFreeSpaceRequest {
    const message = {
      ...baseQueryGetClientFreeSpaceRequest,
    } as QueryGetClientFreeSpaceRequest;
    if (object.address !== undefined && object.address !== null) {
      message.address = String(object.address);
    } else {
      message.address = "";
    }
    return message;
  },

  toJSON(message: QueryGetClientFreeSpaceRequest): unknown {
    const obj: any = {};
    message.address !== undefined && (obj.address = message.address);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetClientFreeSpaceRequest>
  ): QueryGetClientFreeSpaceRequest {
    const message = {
      ...baseQueryGetClientFreeSpaceRequest,
    } as QueryGetClientFreeSpaceRequest;
    if (object.address !== undefined && object.address !== null) {
      message.address = object.address;
    } else {
      message.address = "";
    }
    return message;
  },
};

const baseQueryGetClientFreeSpaceResponse: object = { bytesfree: "" };

export const QueryGetClientFreeSpaceResponse = {
  encode(
    message: QueryGetClientFreeSpaceResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.bytesfree !== "") {
      writer.uint32(10).string(message.bytesfree);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetClientFreeSpaceResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetClientFreeSpaceResponse,
    } as QueryGetClientFreeSpaceResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.bytesfree = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetClientFreeSpaceResponse {
    const message = {
      ...baseQueryGetClientFreeSpaceResponse,
    } as QueryGetClientFreeSpaceResponse;
    if (object.bytesfree !== undefined && object.bytesfree !== null) {
      message.bytesfree = String(object.bytesfree);
    } else {
      message.bytesfree = "";
    }
    return message;
  },

  toJSON(message: QueryGetClientFreeSpaceResponse): unknown {
    const obj: any = {};
    message.bytesfree !== undefined && (obj.bytesfree = message.bytesfree);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetClientFreeSpaceResponse>
  ): QueryGetClientFreeSpaceResponse {
    const message = {
      ...baseQueryGetClientFreeSpaceResponse,
    } as QueryGetClientFreeSpaceResponse;
    if (object.bytesfree !== undefined && object.bytesfree !== null) {
      message.bytesfree = object.bytesfree;
    } else {
      message.bytesfree = "";
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
  /** Queries a PayBlocks by index. */
  PayBlocks(
    request: QueryGetPayBlocksRequest
  ): Promise<QueryGetPayBlocksResponse>;
  /** Queries a list of PayBlocks items. */
  PayBlocksAll(
    request: QueryAllPayBlocksRequest
  ): Promise<QueryAllPayBlocksResponse>;
  /** Queries a ClientUsage by index. */
  ClientUsage(
    request: QueryGetClientUsageRequest
  ): Promise<QueryGetClientUsageResponse>;
  /** Queries a list of ClientUsage items. */
  ClientUsageAll(
    request: QueryAllClientUsageRequest
  ): Promise<QueryAllClientUsageResponse>;
  /** Queries a Strays by index. */
  Strays(request: QueryGetStraysRequest): Promise<QueryGetStraysResponse>;
  /** Queries a list of Strays items. */
  StraysAll(request: QueryAllStraysRequest): Promise<QueryAllStraysResponse>;
  /** Queries a list of GetClientFreeSpace items. */
  GetClientFreeSpace(
    request: QueryGetClientFreeSpaceRequest
  ): Promise<QueryGetClientFreeSpaceResponse>;
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

  PayBlocks(
    request: QueryGetPayBlocksRequest
  ): Promise<QueryGetPayBlocksResponse> {
    const data = QueryGetPayBlocksRequest.encode(request).finish();
    const promise = this.rpc.request(
      "jackaldao.canine.storage.Query",
      "PayBlocks",
      data
    );
    return promise.then((data) =>
      QueryGetPayBlocksResponse.decode(new Reader(data))
    );
  }

  PayBlocksAll(
    request: QueryAllPayBlocksRequest
  ): Promise<QueryAllPayBlocksResponse> {
    const data = QueryAllPayBlocksRequest.encode(request).finish();
    const promise = this.rpc.request(
      "jackaldao.canine.storage.Query",
      "PayBlocksAll",
      data
    );
    return promise.then((data) =>
      QueryAllPayBlocksResponse.decode(new Reader(data))
    );
  }

  ClientUsage(
    request: QueryGetClientUsageRequest
  ): Promise<QueryGetClientUsageResponse> {
    const data = QueryGetClientUsageRequest.encode(request).finish();
    const promise = this.rpc.request(
      "jackaldao.canine.storage.Query",
      "ClientUsage",
      data
    );
    return promise.then((data) =>
      QueryGetClientUsageResponse.decode(new Reader(data))
    );
  }

  ClientUsageAll(
    request: QueryAllClientUsageRequest
  ): Promise<QueryAllClientUsageResponse> {
    const data = QueryAllClientUsageRequest.encode(request).finish();
    const promise = this.rpc.request(
      "jackaldao.canine.storage.Query",
      "ClientUsageAll",
      data
    );
    return promise.then((data) =>
      QueryAllClientUsageResponse.decode(new Reader(data))
    );
  }

  Strays(request: QueryGetStraysRequest): Promise<QueryGetStraysResponse> {
    const data = QueryGetStraysRequest.encode(request).finish();
    const promise = this.rpc.request(
      "jackaldao.canine.storage.Query",
      "Strays",
      data
    );
    return promise.then((data) =>
      QueryGetStraysResponse.decode(new Reader(data))
    );
  }

  StraysAll(request: QueryAllStraysRequest): Promise<QueryAllStraysResponse> {
    const data = QueryAllStraysRequest.encode(request).finish();
    const promise = this.rpc.request(
      "jackaldao.canine.storage.Query",
      "StraysAll",
      data
    );
    return promise.then((data) =>
      QueryAllStraysResponse.decode(new Reader(data))
    );
  }

  GetClientFreeSpace(
    request: QueryGetClientFreeSpaceRequest
  ): Promise<QueryGetClientFreeSpaceResponse> {
    const data = QueryGetClientFreeSpaceRequest.encode(request).finish();
    const promise = this.rpc.request(
      "jackaldao.canine.storage.Query",
      "GetClientFreeSpace",
      data
    );
    return promise.then((data) =>
      QueryGetClientFreeSpaceResponse.decode(new Reader(data))
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
