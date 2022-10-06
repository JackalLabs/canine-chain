/* eslint-disable */
import { Reader, util, configure, Writer } from "protobufjs/minimal";
import * as Long from "long";
import { Params } from "../notifications/params";
import { Notifications } from "../notifications/notifications";
import {
  PageRequest,
  PageResponse,
} from "../cosmos/base/query/v1beta1/pagination";
import { NotiCounter } from "../notifications/noti_counter";

export const protobufPackage = "jackaldao.canine.notifications";

/** QueryParamsRequest is request type for the Query/Params RPC method. */
export interface QueryParamsRequest {}

/** QueryParamsResponse is response type for the Query/Params RPC method. */
export interface QueryParamsResponse {
  /** params holds all the parameters of this module. */
  params: Params | undefined;
}

export interface QueryGetNotificationsRequest {
  count: number;
  address: string;
}

export interface QueryGetNotificationsResponse {
  notifications: Notifications | undefined;
}

export interface QueryAllNotificationsRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllNotificationsResponse {
  notifications: Notifications[];
  pagination: PageResponse | undefined;
}

export interface QueryFilteredNotificationsRequest {
  address: string;
}

export interface QueryFilteredNotificationsResponse {
  /** could turn it back to 'repeated Notifications notifications' */
  notifications: string[];
}

export interface QueryGetNotiCounterRequest {
  address: string;
}

export interface QueryGetNotiCounterResponse {
  notiCounter: NotiCounter | undefined;
}

export interface QueryAllNotiCounterRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllNotiCounterResponse {
  notiCounter: NotiCounter[];
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

const baseQueryGetNotificationsRequest: object = { count: 0, address: "" };

export const QueryGetNotificationsRequest = {
  encode(
    message: QueryGetNotificationsRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.count !== 0) {
      writer.uint32(8).uint64(message.count);
    }
    if (message.address !== "") {
      writer.uint32(18).string(message.address);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetNotificationsRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetNotificationsRequest,
    } as QueryGetNotificationsRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.count = longToNumber(reader.uint64() as Long);
          break;
        case 2:
          message.address = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetNotificationsRequest {
    const message = {
      ...baseQueryGetNotificationsRequest,
    } as QueryGetNotificationsRequest;
    if (object.count !== undefined && object.count !== null) {
      message.count = Number(object.count);
    } else {
      message.count = 0;
    }
    if (object.address !== undefined && object.address !== null) {
      message.address = String(object.address);
    } else {
      message.address = "";
    }
    return message;
  },

  toJSON(message: QueryGetNotificationsRequest): unknown {
    const obj: any = {};
    message.count !== undefined && (obj.count = message.count);
    message.address !== undefined && (obj.address = message.address);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetNotificationsRequest>
  ): QueryGetNotificationsRequest {
    const message = {
      ...baseQueryGetNotificationsRequest,
    } as QueryGetNotificationsRequest;
    if (object.count !== undefined && object.count !== null) {
      message.count = object.count;
    } else {
      message.count = 0;
    }
    if (object.address !== undefined && object.address !== null) {
      message.address = object.address;
    } else {
      message.address = "";
    }
    return message;
  },
};

const baseQueryGetNotificationsResponse: object = {};

export const QueryGetNotificationsResponse = {
  encode(
    message: QueryGetNotificationsResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.notifications !== undefined) {
      Notifications.encode(
        message.notifications,
        writer.uint32(10).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetNotificationsResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetNotificationsResponse,
    } as QueryGetNotificationsResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.notifications = Notifications.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetNotificationsResponse {
    const message = {
      ...baseQueryGetNotificationsResponse,
    } as QueryGetNotificationsResponse;
    if (object.notifications !== undefined && object.notifications !== null) {
      message.notifications = Notifications.fromJSON(object.notifications);
    } else {
      message.notifications = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetNotificationsResponse): unknown {
    const obj: any = {};
    message.notifications !== undefined &&
      (obj.notifications = message.notifications
        ? Notifications.toJSON(message.notifications)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetNotificationsResponse>
  ): QueryGetNotificationsResponse {
    const message = {
      ...baseQueryGetNotificationsResponse,
    } as QueryGetNotificationsResponse;
    if (object.notifications !== undefined && object.notifications !== null) {
      message.notifications = Notifications.fromPartial(object.notifications);
    } else {
      message.notifications = undefined;
    }
    return message;
  },
};

const baseQueryAllNotificationsRequest: object = {};

export const QueryAllNotificationsRequest = {
  encode(
    message: QueryAllNotificationsRequest,
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
  ): QueryAllNotificationsRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllNotificationsRequest,
    } as QueryAllNotificationsRequest;
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

  fromJSON(object: any): QueryAllNotificationsRequest {
    const message = {
      ...baseQueryAllNotificationsRequest,
    } as QueryAllNotificationsRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllNotificationsRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageRequest.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllNotificationsRequest>
  ): QueryAllNotificationsRequest {
    const message = {
      ...baseQueryAllNotificationsRequest,
    } as QueryAllNotificationsRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryAllNotificationsResponse: object = {};

export const QueryAllNotificationsResponse = {
  encode(
    message: QueryAllNotificationsResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.notifications) {
      Notifications.encode(v!, writer.uint32(10).fork()).ldelim();
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
  ): QueryAllNotificationsResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllNotificationsResponse,
    } as QueryAllNotificationsResponse;
    message.notifications = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.notifications.push(
            Notifications.decode(reader, reader.uint32())
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

  fromJSON(object: any): QueryAllNotificationsResponse {
    const message = {
      ...baseQueryAllNotificationsResponse,
    } as QueryAllNotificationsResponse;
    message.notifications = [];
    if (object.notifications !== undefined && object.notifications !== null) {
      for (const e of object.notifications) {
        message.notifications.push(Notifications.fromJSON(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllNotificationsResponse): unknown {
    const obj: any = {};
    if (message.notifications) {
      obj.notifications = message.notifications.map((e) =>
        e ? Notifications.toJSON(e) : undefined
      );
    } else {
      obj.notifications = [];
    }
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageResponse.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllNotificationsResponse>
  ): QueryAllNotificationsResponse {
    const message = {
      ...baseQueryAllNotificationsResponse,
    } as QueryAllNotificationsResponse;
    message.notifications = [];
    if (object.notifications !== undefined && object.notifications !== null) {
      for (const e of object.notifications) {
        message.notifications.push(Notifications.fromPartial(e));
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

const baseQueryFilteredNotificationsRequest: object = { address: "" };

export const QueryFilteredNotificationsRequest = {
  encode(
    message: QueryFilteredNotificationsRequest,
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
  ): QueryFilteredNotificationsRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryFilteredNotificationsRequest,
    } as QueryFilteredNotificationsRequest;
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

  fromJSON(object: any): QueryFilteredNotificationsRequest {
    const message = {
      ...baseQueryFilteredNotificationsRequest,
    } as QueryFilteredNotificationsRequest;
    if (object.address !== undefined && object.address !== null) {
      message.address = String(object.address);
    } else {
      message.address = "";
    }
    return message;
  },

  toJSON(message: QueryFilteredNotificationsRequest): unknown {
    const obj: any = {};
    message.address !== undefined && (obj.address = message.address);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryFilteredNotificationsRequest>
  ): QueryFilteredNotificationsRequest {
    const message = {
      ...baseQueryFilteredNotificationsRequest,
    } as QueryFilteredNotificationsRequest;
    if (object.address !== undefined && object.address !== null) {
      message.address = object.address;
    } else {
      message.address = "";
    }
    return message;
  },
};

const baseQueryFilteredNotificationsResponse: object = { notifications: "" };

export const QueryFilteredNotificationsResponse = {
  encode(
    message: QueryFilteredNotificationsResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.notifications) {
      writer.uint32(10).string(v!);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryFilteredNotificationsResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryFilteredNotificationsResponse,
    } as QueryFilteredNotificationsResponse;
    message.notifications = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.notifications.push(reader.string());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryFilteredNotificationsResponse {
    const message = {
      ...baseQueryFilteredNotificationsResponse,
    } as QueryFilteredNotificationsResponse;
    message.notifications = [];
    if (object.notifications !== undefined && object.notifications !== null) {
      for (const e of object.notifications) {
        message.notifications.push(String(e));
      }
    }
    return message;
  },

  toJSON(message: QueryFilteredNotificationsResponse): unknown {
    const obj: any = {};
    if (message.notifications) {
      obj.notifications = message.notifications.map((e) => e);
    } else {
      obj.notifications = [];
    }
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryFilteredNotificationsResponse>
  ): QueryFilteredNotificationsResponse {
    const message = {
      ...baseQueryFilteredNotificationsResponse,
    } as QueryFilteredNotificationsResponse;
    message.notifications = [];
    if (object.notifications !== undefined && object.notifications !== null) {
      for (const e of object.notifications) {
        message.notifications.push(e);
      }
    }
    return message;
  },
};

const baseQueryGetNotiCounterRequest: object = { address: "" };

export const QueryGetNotiCounterRequest = {
  encode(
    message: QueryGetNotiCounterRequest,
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
  ): QueryGetNotiCounterRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetNotiCounterRequest,
    } as QueryGetNotiCounterRequest;
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

  fromJSON(object: any): QueryGetNotiCounterRequest {
    const message = {
      ...baseQueryGetNotiCounterRequest,
    } as QueryGetNotiCounterRequest;
    if (object.address !== undefined && object.address !== null) {
      message.address = String(object.address);
    } else {
      message.address = "";
    }
    return message;
  },

  toJSON(message: QueryGetNotiCounterRequest): unknown {
    const obj: any = {};
    message.address !== undefined && (obj.address = message.address);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetNotiCounterRequest>
  ): QueryGetNotiCounterRequest {
    const message = {
      ...baseQueryGetNotiCounterRequest,
    } as QueryGetNotiCounterRequest;
    if (object.address !== undefined && object.address !== null) {
      message.address = object.address;
    } else {
      message.address = "";
    }
    return message;
  },
};

const baseQueryGetNotiCounterResponse: object = {};

export const QueryGetNotiCounterResponse = {
  encode(
    message: QueryGetNotiCounterResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.notiCounter !== undefined) {
      NotiCounter.encode(
        message.notiCounter,
        writer.uint32(10).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetNotiCounterResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetNotiCounterResponse,
    } as QueryGetNotiCounterResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.notiCounter = NotiCounter.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetNotiCounterResponse {
    const message = {
      ...baseQueryGetNotiCounterResponse,
    } as QueryGetNotiCounterResponse;
    if (object.notiCounter !== undefined && object.notiCounter !== null) {
      message.notiCounter = NotiCounter.fromJSON(object.notiCounter);
    } else {
      message.notiCounter = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetNotiCounterResponse): unknown {
    const obj: any = {};
    message.notiCounter !== undefined &&
      (obj.notiCounter = message.notiCounter
        ? NotiCounter.toJSON(message.notiCounter)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetNotiCounterResponse>
  ): QueryGetNotiCounterResponse {
    const message = {
      ...baseQueryGetNotiCounterResponse,
    } as QueryGetNotiCounterResponse;
    if (object.notiCounter !== undefined && object.notiCounter !== null) {
      message.notiCounter = NotiCounter.fromPartial(object.notiCounter);
    } else {
      message.notiCounter = undefined;
    }
    return message;
  },
};

const baseQueryAllNotiCounterRequest: object = {};

export const QueryAllNotiCounterRequest = {
  encode(
    message: QueryAllNotiCounterRequest,
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
  ): QueryAllNotiCounterRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllNotiCounterRequest,
    } as QueryAllNotiCounterRequest;
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

  fromJSON(object: any): QueryAllNotiCounterRequest {
    const message = {
      ...baseQueryAllNotiCounterRequest,
    } as QueryAllNotiCounterRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllNotiCounterRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageRequest.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllNotiCounterRequest>
  ): QueryAllNotiCounterRequest {
    const message = {
      ...baseQueryAllNotiCounterRequest,
    } as QueryAllNotiCounterRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryAllNotiCounterResponse: object = {};

export const QueryAllNotiCounterResponse = {
  encode(
    message: QueryAllNotiCounterResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.notiCounter) {
      NotiCounter.encode(v!, writer.uint32(10).fork()).ldelim();
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
  ): QueryAllNotiCounterResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllNotiCounterResponse,
    } as QueryAllNotiCounterResponse;
    message.notiCounter = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.notiCounter.push(NotiCounter.decode(reader, reader.uint32()));
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

  fromJSON(object: any): QueryAllNotiCounterResponse {
    const message = {
      ...baseQueryAllNotiCounterResponse,
    } as QueryAllNotiCounterResponse;
    message.notiCounter = [];
    if (object.notiCounter !== undefined && object.notiCounter !== null) {
      for (const e of object.notiCounter) {
        message.notiCounter.push(NotiCounter.fromJSON(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllNotiCounterResponse): unknown {
    const obj: any = {};
    if (message.notiCounter) {
      obj.notiCounter = message.notiCounter.map((e) =>
        e ? NotiCounter.toJSON(e) : undefined
      );
    } else {
      obj.notiCounter = [];
    }
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageResponse.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllNotiCounterResponse>
  ): QueryAllNotiCounterResponse {
    const message = {
      ...baseQueryAllNotiCounterResponse,
    } as QueryAllNotiCounterResponse;
    message.notiCounter = [];
    if (object.notiCounter !== undefined && object.notiCounter !== null) {
      for (const e of object.notiCounter) {
        message.notiCounter.push(NotiCounter.fromPartial(e));
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
  /** Queries a Notifications by index. */
  Notifications(
    request: QueryGetNotificationsRequest
  ): Promise<QueryGetNotificationsResponse>;
  /** Queries a list of Notifications items. */
  NotificationsAll(
    request: QueryAllNotificationsRequest
  ): Promise<QueryAllNotificationsResponse>;
  /** Queries a list of FilteredNotifications items. */
  FilteredNotifications(
    request: QueryFilteredNotificationsRequest
  ): Promise<QueryFilteredNotificationsResponse>;
  /** Queries a NotiCounter by index. */
  NotiCounter(
    request: QueryGetNotiCounterRequest
  ): Promise<QueryGetNotiCounterResponse>;
  /** Queries a list of NotiCounter items. */
  NotiCounterAll(
    request: QueryAllNotiCounterRequest
  ): Promise<QueryAllNotiCounterResponse>;
}

export class QueryClientImpl implements Query {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse> {
    const data = QueryParamsRequest.encode(request).finish();
    const promise = this.rpc.request(
      "jackaldao.canine.notifications.Query",
      "Params",
      data
    );
    return promise.then((data) => QueryParamsResponse.decode(new Reader(data)));
  }

  Notifications(
    request: QueryGetNotificationsRequest
  ): Promise<QueryGetNotificationsResponse> {
    const data = QueryGetNotificationsRequest.encode(request).finish();
    const promise = this.rpc.request(
      "jackaldao.canine.notifications.Query",
      "Notifications",
      data
    );
    return promise.then((data) =>
      QueryGetNotificationsResponse.decode(new Reader(data))
    );
  }

  NotificationsAll(
    request: QueryAllNotificationsRequest
  ): Promise<QueryAllNotificationsResponse> {
    const data = QueryAllNotificationsRequest.encode(request).finish();
    const promise = this.rpc.request(
      "jackaldao.canine.notifications.Query",
      "NotificationsAll",
      data
    );
    return promise.then((data) =>
      QueryAllNotificationsResponse.decode(new Reader(data))
    );
  }

  FilteredNotifications(
    request: QueryFilteredNotificationsRequest
  ): Promise<QueryFilteredNotificationsResponse> {
    const data = QueryFilteredNotificationsRequest.encode(request).finish();
    const promise = this.rpc.request(
      "jackaldao.canine.notifications.Query",
      "FilteredNotifications",
      data
    );
    return promise.then((data) =>
      QueryFilteredNotificationsResponse.decode(new Reader(data))
    );
  }

  NotiCounter(
    request: QueryGetNotiCounterRequest
  ): Promise<QueryGetNotiCounterResponse> {
    const data = QueryGetNotiCounterRequest.encode(request).finish();
    const promise = this.rpc.request(
      "jackaldao.canine.notifications.Query",
      "NotiCounter",
      data
    );
    return promise.then((data) =>
      QueryGetNotiCounterResponse.decode(new Reader(data))
    );
  }

  NotiCounterAll(
    request: QueryAllNotiCounterRequest
  ): Promise<QueryAllNotiCounterResponse> {
    const data = QueryAllNotiCounterRequest.encode(request).finish();
    const promise = this.rpc.request(
      "jackaldao.canine.notifications.Query",
      "NotiCounterAll",
      data
    );
    return promise.then((data) =>
      QueryAllNotiCounterResponse.decode(new Reader(data))
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
