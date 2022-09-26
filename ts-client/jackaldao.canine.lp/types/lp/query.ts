/* eslint-disable */
import { Reader, Writer } from "protobufjs/minimal";
import { Params } from "../lp/params";
import { LPool } from "../lp/l_pool";
import {
  PageRequest,
  PageResponse,
} from "../cosmos/base/query/v1beta1/pagination";
import { LProviderRecord } from "../lp/l_provider_record";
import { Coin } from "../cosmos/base/v1beta1/coin";

export const protobufPackage = "jackaldao.canine.lp";

/** QueryParamsRequest is request type for the Query/Params RPC method. */
export interface QueryParamsRequest {}

/** QueryParamsResponse is response type for the Query/Params RPC method. */
export interface QueryParamsResponse {
  /** params holds all the parameters of this module. */
  params: Params | undefined;
}

export interface QueryGetLPoolRequest {
  index: string;
}

export interface QueryGetLPoolResponse {
  lPool: LPool | undefined;
}

export interface QueryAllLPoolRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllLPoolResponse {
  lPool: LPool[];
  pagination: PageResponse | undefined;
}

export interface QueryGetLProviderRecordRequest {
  provider: string;
  poolName: string;
}

export interface QueryGetLProviderRecordResponse {
  lProviderRecord: LProviderRecord | undefined;
}

export interface QueryEstimateSwapOutRequest {
  poolName: string;
  inputCoin: string;
}

export interface QueryEstimateSwapOutResponse {
  outputCoin: Coin | undefined;
}

export interface QueryEstimateSwapInRequest {
  poolName: string;
  outputCoins: string;
}

export interface QueryEstimateSwapInResponse {
  inputCoins: Coin | undefined;
}

/** Estimate amount of coins to deposit to get desired amount of LPToken */
export interface QueryEstimateContributionRequest {
  poolName: string;
  desiredAmount: string;
}

export interface QueryEstimateContributionResponse {
  coins: Coin[];
}

/** Query amount of coins to deposit to make a valid liquidity pair */
export interface QueryMakeValidPairRequest {
  poolName: string;
  coin: string;
}

export interface QueryMakeValidPairResponse {
  coin: Coin | undefined;
}

/** Estimate pool coins returned by burning LPToken */
export interface QueryEstimatePoolRemoveRequest {
  poolName: string;
  amount: string;
}

export interface QueryEstimatePoolRemoveResponse {
  coins: Coin[];
}

export interface QueryListRecordsFromPoolRequest {
  poolName: string;
}

export interface QueryListRecordsFromPoolResponse {
  records: LProviderRecord[];
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

const baseQueryGetLPoolRequest: object = { index: "" };

export const QueryGetLPoolRequest = {
  encode(
    message: QueryGetLPoolRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.index !== "") {
      writer.uint32(10).string(message.index);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetLPoolRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryGetLPoolRequest } as QueryGetLPoolRequest;
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

  fromJSON(object: any): QueryGetLPoolRequest {
    const message = { ...baseQueryGetLPoolRequest } as QueryGetLPoolRequest;
    if (object.index !== undefined && object.index !== null) {
      message.index = String(object.index);
    } else {
      message.index = "";
    }
    return message;
  },

  toJSON(message: QueryGetLPoolRequest): unknown {
    const obj: any = {};
    message.index !== undefined && (obj.index = message.index);
    return obj;
  },

  fromPartial(object: DeepPartial<QueryGetLPoolRequest>): QueryGetLPoolRequest {
    const message = { ...baseQueryGetLPoolRequest } as QueryGetLPoolRequest;
    if (object.index !== undefined && object.index !== null) {
      message.index = object.index;
    } else {
      message.index = "";
    }
    return message;
  },
};

const baseQueryGetLPoolResponse: object = {};

export const QueryGetLPoolResponse = {
  encode(
    message: QueryGetLPoolResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.lPool !== undefined) {
      LPool.encode(message.lPool, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetLPoolResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryGetLPoolResponse } as QueryGetLPoolResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.lPool = LPool.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetLPoolResponse {
    const message = { ...baseQueryGetLPoolResponse } as QueryGetLPoolResponse;
    if (object.lPool !== undefined && object.lPool !== null) {
      message.lPool = LPool.fromJSON(object.lPool);
    } else {
      message.lPool = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetLPoolResponse): unknown {
    const obj: any = {};
    message.lPool !== undefined &&
      (obj.lPool = message.lPool ? LPool.toJSON(message.lPool) : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetLPoolResponse>
  ): QueryGetLPoolResponse {
    const message = { ...baseQueryGetLPoolResponse } as QueryGetLPoolResponse;
    if (object.lPool !== undefined && object.lPool !== null) {
      message.lPool = LPool.fromPartial(object.lPool);
    } else {
      message.lPool = undefined;
    }
    return message;
  },
};

const baseQueryAllLPoolRequest: object = {};

export const QueryAllLPoolRequest = {
  encode(
    message: QueryAllLPoolRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllLPoolRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryAllLPoolRequest } as QueryAllLPoolRequest;
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

  fromJSON(object: any): QueryAllLPoolRequest {
    const message = { ...baseQueryAllLPoolRequest } as QueryAllLPoolRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllLPoolRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageRequest.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<QueryAllLPoolRequest>): QueryAllLPoolRequest {
    const message = { ...baseQueryAllLPoolRequest } as QueryAllLPoolRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryAllLPoolResponse: object = {};

export const QueryAllLPoolResponse = {
  encode(
    message: QueryAllLPoolResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.lPool) {
      LPool.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(
        message.pagination,
        writer.uint32(18).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllLPoolResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryAllLPoolResponse } as QueryAllLPoolResponse;
    message.lPool = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.lPool.push(LPool.decode(reader, reader.uint32()));
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

  fromJSON(object: any): QueryAllLPoolResponse {
    const message = { ...baseQueryAllLPoolResponse } as QueryAllLPoolResponse;
    message.lPool = [];
    if (object.lPool !== undefined && object.lPool !== null) {
      for (const e of object.lPool) {
        message.lPool.push(LPool.fromJSON(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllLPoolResponse): unknown {
    const obj: any = {};
    if (message.lPool) {
      obj.lPool = message.lPool.map((e) => (e ? LPool.toJSON(e) : undefined));
    } else {
      obj.lPool = [];
    }
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageResponse.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllLPoolResponse>
  ): QueryAllLPoolResponse {
    const message = { ...baseQueryAllLPoolResponse } as QueryAllLPoolResponse;
    message.lPool = [];
    if (object.lPool !== undefined && object.lPool !== null) {
      for (const e of object.lPool) {
        message.lPool.push(LPool.fromPartial(e));
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

const baseQueryGetLProviderRecordRequest: object = {
  provider: "",
  poolName: "",
};

export const QueryGetLProviderRecordRequest = {
  encode(
    message: QueryGetLProviderRecordRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.provider !== "") {
      writer.uint32(10).string(message.provider);
    }
    if (message.poolName !== "") {
      writer.uint32(18).string(message.poolName);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetLProviderRecordRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetLProviderRecordRequest,
    } as QueryGetLProviderRecordRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.provider = reader.string();
          break;
        case 2:
          message.poolName = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetLProviderRecordRequest {
    const message = {
      ...baseQueryGetLProviderRecordRequest,
    } as QueryGetLProviderRecordRequest;
    if (object.provider !== undefined && object.provider !== null) {
      message.provider = String(object.provider);
    } else {
      message.provider = "";
    }
    if (object.poolName !== undefined && object.poolName !== null) {
      message.poolName = String(object.poolName);
    } else {
      message.poolName = "";
    }
    return message;
  },

  toJSON(message: QueryGetLProviderRecordRequest): unknown {
    const obj: any = {};
    message.provider !== undefined && (obj.provider = message.provider);
    message.poolName !== undefined && (obj.poolName = message.poolName);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetLProviderRecordRequest>
  ): QueryGetLProviderRecordRequest {
    const message = {
      ...baseQueryGetLProviderRecordRequest,
    } as QueryGetLProviderRecordRequest;
    if (object.provider !== undefined && object.provider !== null) {
      message.provider = object.provider;
    } else {
      message.provider = "";
    }
    if (object.poolName !== undefined && object.poolName !== null) {
      message.poolName = object.poolName;
    } else {
      message.poolName = "";
    }
    return message;
  },
};

const baseQueryGetLProviderRecordResponse: object = {};

export const QueryGetLProviderRecordResponse = {
  encode(
    message: QueryGetLProviderRecordResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.lProviderRecord !== undefined) {
      LProviderRecord.encode(
        message.lProviderRecord,
        writer.uint32(10).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetLProviderRecordResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetLProviderRecordResponse,
    } as QueryGetLProviderRecordResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.lProviderRecord = LProviderRecord.decode(
            reader,
            reader.uint32()
          );
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetLProviderRecordResponse {
    const message = {
      ...baseQueryGetLProviderRecordResponse,
    } as QueryGetLProviderRecordResponse;
    if (
      object.lProviderRecord !== undefined &&
      object.lProviderRecord !== null
    ) {
      message.lProviderRecord = LProviderRecord.fromJSON(
        object.lProviderRecord
      );
    } else {
      message.lProviderRecord = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetLProviderRecordResponse): unknown {
    const obj: any = {};
    message.lProviderRecord !== undefined &&
      (obj.lProviderRecord = message.lProviderRecord
        ? LProviderRecord.toJSON(message.lProviderRecord)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetLProviderRecordResponse>
  ): QueryGetLProviderRecordResponse {
    const message = {
      ...baseQueryGetLProviderRecordResponse,
    } as QueryGetLProviderRecordResponse;
    if (
      object.lProviderRecord !== undefined &&
      object.lProviderRecord !== null
    ) {
      message.lProviderRecord = LProviderRecord.fromPartial(
        object.lProviderRecord
      );
    } else {
      message.lProviderRecord = undefined;
    }
    return message;
  },
};

const baseQueryEstimateSwapOutRequest: object = { poolName: "", inputCoin: "" };

export const QueryEstimateSwapOutRequest = {
  encode(
    message: QueryEstimateSwapOutRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.poolName !== "") {
      writer.uint32(10).string(message.poolName);
    }
    if (message.inputCoin !== "") {
      writer.uint32(18).string(message.inputCoin);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryEstimateSwapOutRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryEstimateSwapOutRequest,
    } as QueryEstimateSwapOutRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.poolName = reader.string();
          break;
        case 2:
          message.inputCoin = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryEstimateSwapOutRequest {
    const message = {
      ...baseQueryEstimateSwapOutRequest,
    } as QueryEstimateSwapOutRequest;
    if (object.poolName !== undefined && object.poolName !== null) {
      message.poolName = String(object.poolName);
    } else {
      message.poolName = "";
    }
    if (object.inputCoin !== undefined && object.inputCoin !== null) {
      message.inputCoin = String(object.inputCoin);
    } else {
      message.inputCoin = "";
    }
    return message;
  },

  toJSON(message: QueryEstimateSwapOutRequest): unknown {
    const obj: any = {};
    message.poolName !== undefined && (obj.poolName = message.poolName);
    message.inputCoin !== undefined && (obj.inputCoin = message.inputCoin);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryEstimateSwapOutRequest>
  ): QueryEstimateSwapOutRequest {
    const message = {
      ...baseQueryEstimateSwapOutRequest,
    } as QueryEstimateSwapOutRequest;
    if (object.poolName !== undefined && object.poolName !== null) {
      message.poolName = object.poolName;
    } else {
      message.poolName = "";
    }
    if (object.inputCoin !== undefined && object.inputCoin !== null) {
      message.inputCoin = object.inputCoin;
    } else {
      message.inputCoin = "";
    }
    return message;
  },
};

const baseQueryEstimateSwapOutResponse: object = {};

export const QueryEstimateSwapOutResponse = {
  encode(
    message: QueryEstimateSwapOutResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.outputCoin !== undefined) {
      Coin.encode(message.outputCoin, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryEstimateSwapOutResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryEstimateSwapOutResponse,
    } as QueryEstimateSwapOutResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.outputCoin = Coin.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryEstimateSwapOutResponse {
    const message = {
      ...baseQueryEstimateSwapOutResponse,
    } as QueryEstimateSwapOutResponse;
    if (object.outputCoin !== undefined && object.outputCoin !== null) {
      message.outputCoin = Coin.fromJSON(object.outputCoin);
    } else {
      message.outputCoin = undefined;
    }
    return message;
  },

  toJSON(message: QueryEstimateSwapOutResponse): unknown {
    const obj: any = {};
    message.outputCoin !== undefined &&
      (obj.outputCoin = message.outputCoin
        ? Coin.toJSON(message.outputCoin)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryEstimateSwapOutResponse>
  ): QueryEstimateSwapOutResponse {
    const message = {
      ...baseQueryEstimateSwapOutResponse,
    } as QueryEstimateSwapOutResponse;
    if (object.outputCoin !== undefined && object.outputCoin !== null) {
      message.outputCoin = Coin.fromPartial(object.outputCoin);
    } else {
      message.outputCoin = undefined;
    }
    return message;
  },
};

const baseQueryEstimateSwapInRequest: object = {
  poolName: "",
  outputCoins: "",
};

export const QueryEstimateSwapInRequest = {
  encode(
    message: QueryEstimateSwapInRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.poolName !== "") {
      writer.uint32(10).string(message.poolName);
    }
    if (message.outputCoins !== "") {
      writer.uint32(18).string(message.outputCoins);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryEstimateSwapInRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryEstimateSwapInRequest,
    } as QueryEstimateSwapInRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.poolName = reader.string();
          break;
        case 2:
          message.outputCoins = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryEstimateSwapInRequest {
    const message = {
      ...baseQueryEstimateSwapInRequest,
    } as QueryEstimateSwapInRequest;
    if (object.poolName !== undefined && object.poolName !== null) {
      message.poolName = String(object.poolName);
    } else {
      message.poolName = "";
    }
    if (object.outputCoins !== undefined && object.outputCoins !== null) {
      message.outputCoins = String(object.outputCoins);
    } else {
      message.outputCoins = "";
    }
    return message;
  },

  toJSON(message: QueryEstimateSwapInRequest): unknown {
    const obj: any = {};
    message.poolName !== undefined && (obj.poolName = message.poolName);
    message.outputCoins !== undefined &&
      (obj.outputCoins = message.outputCoins);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryEstimateSwapInRequest>
  ): QueryEstimateSwapInRequest {
    const message = {
      ...baseQueryEstimateSwapInRequest,
    } as QueryEstimateSwapInRequest;
    if (object.poolName !== undefined && object.poolName !== null) {
      message.poolName = object.poolName;
    } else {
      message.poolName = "";
    }
    if (object.outputCoins !== undefined && object.outputCoins !== null) {
      message.outputCoins = object.outputCoins;
    } else {
      message.outputCoins = "";
    }
    return message;
  },
};

const baseQueryEstimateSwapInResponse: object = {};

export const QueryEstimateSwapInResponse = {
  encode(
    message: QueryEstimateSwapInResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.inputCoins !== undefined) {
      Coin.encode(message.inputCoins, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryEstimateSwapInResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryEstimateSwapInResponse,
    } as QueryEstimateSwapInResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.inputCoins = Coin.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryEstimateSwapInResponse {
    const message = {
      ...baseQueryEstimateSwapInResponse,
    } as QueryEstimateSwapInResponse;
    if (object.inputCoins !== undefined && object.inputCoins !== null) {
      message.inputCoins = Coin.fromJSON(object.inputCoins);
    } else {
      message.inputCoins = undefined;
    }
    return message;
  },

  toJSON(message: QueryEstimateSwapInResponse): unknown {
    const obj: any = {};
    message.inputCoins !== undefined &&
      (obj.inputCoins = message.inputCoins
        ? Coin.toJSON(message.inputCoins)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryEstimateSwapInResponse>
  ): QueryEstimateSwapInResponse {
    const message = {
      ...baseQueryEstimateSwapInResponse,
    } as QueryEstimateSwapInResponse;
    if (object.inputCoins !== undefined && object.inputCoins !== null) {
      message.inputCoins = Coin.fromPartial(object.inputCoins);
    } else {
      message.inputCoins = undefined;
    }
    return message;
  },
};

const baseQueryEstimateContributionRequest: object = {
  poolName: "",
  desiredAmount: "",
};

export const QueryEstimateContributionRequest = {
  encode(
    message: QueryEstimateContributionRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.poolName !== "") {
      writer.uint32(10).string(message.poolName);
    }
    if (message.desiredAmount !== "") {
      writer.uint32(18).string(message.desiredAmount);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryEstimateContributionRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryEstimateContributionRequest,
    } as QueryEstimateContributionRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.poolName = reader.string();
          break;
        case 2:
          message.desiredAmount = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryEstimateContributionRequest {
    const message = {
      ...baseQueryEstimateContributionRequest,
    } as QueryEstimateContributionRequest;
    if (object.poolName !== undefined && object.poolName !== null) {
      message.poolName = String(object.poolName);
    } else {
      message.poolName = "";
    }
    if (object.desiredAmount !== undefined && object.desiredAmount !== null) {
      message.desiredAmount = String(object.desiredAmount);
    } else {
      message.desiredAmount = "";
    }
    return message;
  },

  toJSON(message: QueryEstimateContributionRequest): unknown {
    const obj: any = {};
    message.poolName !== undefined && (obj.poolName = message.poolName);
    message.desiredAmount !== undefined &&
      (obj.desiredAmount = message.desiredAmount);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryEstimateContributionRequest>
  ): QueryEstimateContributionRequest {
    const message = {
      ...baseQueryEstimateContributionRequest,
    } as QueryEstimateContributionRequest;
    if (object.poolName !== undefined && object.poolName !== null) {
      message.poolName = object.poolName;
    } else {
      message.poolName = "";
    }
    if (object.desiredAmount !== undefined && object.desiredAmount !== null) {
      message.desiredAmount = object.desiredAmount;
    } else {
      message.desiredAmount = "";
    }
    return message;
  },
};

const baseQueryEstimateContributionResponse: object = {};

export const QueryEstimateContributionResponse = {
  encode(
    message: QueryEstimateContributionResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.coins) {
      Coin.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryEstimateContributionResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryEstimateContributionResponse,
    } as QueryEstimateContributionResponse;
    message.coins = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.coins.push(Coin.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryEstimateContributionResponse {
    const message = {
      ...baseQueryEstimateContributionResponse,
    } as QueryEstimateContributionResponse;
    message.coins = [];
    if (object.coins !== undefined && object.coins !== null) {
      for (const e of object.coins) {
        message.coins.push(Coin.fromJSON(e));
      }
    }
    return message;
  },

  toJSON(message: QueryEstimateContributionResponse): unknown {
    const obj: any = {};
    if (message.coins) {
      obj.coins = message.coins.map((e) => (e ? Coin.toJSON(e) : undefined));
    } else {
      obj.coins = [];
    }
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryEstimateContributionResponse>
  ): QueryEstimateContributionResponse {
    const message = {
      ...baseQueryEstimateContributionResponse,
    } as QueryEstimateContributionResponse;
    message.coins = [];
    if (object.coins !== undefined && object.coins !== null) {
      for (const e of object.coins) {
        message.coins.push(Coin.fromPartial(e));
      }
    }
    return message;
  },
};

const baseQueryMakeValidPairRequest: object = { poolName: "", coin: "" };

export const QueryMakeValidPairRequest = {
  encode(
    message: QueryMakeValidPairRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.poolName !== "") {
      writer.uint32(10).string(message.poolName);
    }
    if (message.coin !== "") {
      writer.uint32(18).string(message.coin);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryMakeValidPairRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryMakeValidPairRequest,
    } as QueryMakeValidPairRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.poolName = reader.string();
          break;
        case 2:
          message.coin = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryMakeValidPairRequest {
    const message = {
      ...baseQueryMakeValidPairRequest,
    } as QueryMakeValidPairRequest;
    if (object.poolName !== undefined && object.poolName !== null) {
      message.poolName = String(object.poolName);
    } else {
      message.poolName = "";
    }
    if (object.coin !== undefined && object.coin !== null) {
      message.coin = String(object.coin);
    } else {
      message.coin = "";
    }
    return message;
  },

  toJSON(message: QueryMakeValidPairRequest): unknown {
    const obj: any = {};
    message.poolName !== undefined && (obj.poolName = message.poolName);
    message.coin !== undefined && (obj.coin = message.coin);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryMakeValidPairRequest>
  ): QueryMakeValidPairRequest {
    const message = {
      ...baseQueryMakeValidPairRequest,
    } as QueryMakeValidPairRequest;
    if (object.poolName !== undefined && object.poolName !== null) {
      message.poolName = object.poolName;
    } else {
      message.poolName = "";
    }
    if (object.coin !== undefined && object.coin !== null) {
      message.coin = object.coin;
    } else {
      message.coin = "";
    }
    return message;
  },
};

const baseQueryMakeValidPairResponse: object = {};

export const QueryMakeValidPairResponse = {
  encode(
    message: QueryMakeValidPairResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.coin !== undefined) {
      Coin.encode(message.coin, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryMakeValidPairResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryMakeValidPairResponse,
    } as QueryMakeValidPairResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.coin = Coin.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryMakeValidPairResponse {
    const message = {
      ...baseQueryMakeValidPairResponse,
    } as QueryMakeValidPairResponse;
    if (object.coin !== undefined && object.coin !== null) {
      message.coin = Coin.fromJSON(object.coin);
    } else {
      message.coin = undefined;
    }
    return message;
  },

  toJSON(message: QueryMakeValidPairResponse): unknown {
    const obj: any = {};
    message.coin !== undefined &&
      (obj.coin = message.coin ? Coin.toJSON(message.coin) : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryMakeValidPairResponse>
  ): QueryMakeValidPairResponse {
    const message = {
      ...baseQueryMakeValidPairResponse,
    } as QueryMakeValidPairResponse;
    if (object.coin !== undefined && object.coin !== null) {
      message.coin = Coin.fromPartial(object.coin);
    } else {
      message.coin = undefined;
    }
    return message;
  },
};

const baseQueryEstimatePoolRemoveRequest: object = { poolName: "", amount: "" };

export const QueryEstimatePoolRemoveRequest = {
  encode(
    message: QueryEstimatePoolRemoveRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.poolName !== "") {
      writer.uint32(10).string(message.poolName);
    }
    if (message.amount !== "") {
      writer.uint32(18).string(message.amount);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryEstimatePoolRemoveRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryEstimatePoolRemoveRequest,
    } as QueryEstimatePoolRemoveRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.poolName = reader.string();
          break;
        case 2:
          message.amount = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryEstimatePoolRemoveRequest {
    const message = {
      ...baseQueryEstimatePoolRemoveRequest,
    } as QueryEstimatePoolRemoveRequest;
    if (object.poolName !== undefined && object.poolName !== null) {
      message.poolName = String(object.poolName);
    } else {
      message.poolName = "";
    }
    if (object.amount !== undefined && object.amount !== null) {
      message.amount = String(object.amount);
    } else {
      message.amount = "";
    }
    return message;
  },

  toJSON(message: QueryEstimatePoolRemoveRequest): unknown {
    const obj: any = {};
    message.poolName !== undefined && (obj.poolName = message.poolName);
    message.amount !== undefined && (obj.amount = message.amount);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryEstimatePoolRemoveRequest>
  ): QueryEstimatePoolRemoveRequest {
    const message = {
      ...baseQueryEstimatePoolRemoveRequest,
    } as QueryEstimatePoolRemoveRequest;
    if (object.poolName !== undefined && object.poolName !== null) {
      message.poolName = object.poolName;
    } else {
      message.poolName = "";
    }
    if (object.amount !== undefined && object.amount !== null) {
      message.amount = object.amount;
    } else {
      message.amount = "";
    }
    return message;
  },
};

const baseQueryEstimatePoolRemoveResponse: object = {};

export const QueryEstimatePoolRemoveResponse = {
  encode(
    message: QueryEstimatePoolRemoveResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.coins) {
      Coin.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryEstimatePoolRemoveResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryEstimatePoolRemoveResponse,
    } as QueryEstimatePoolRemoveResponse;
    message.coins = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.coins.push(Coin.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryEstimatePoolRemoveResponse {
    const message = {
      ...baseQueryEstimatePoolRemoveResponse,
    } as QueryEstimatePoolRemoveResponse;
    message.coins = [];
    if (object.coins !== undefined && object.coins !== null) {
      for (const e of object.coins) {
        message.coins.push(Coin.fromJSON(e));
      }
    }
    return message;
  },

  toJSON(message: QueryEstimatePoolRemoveResponse): unknown {
    const obj: any = {};
    if (message.coins) {
      obj.coins = message.coins.map((e) => (e ? Coin.toJSON(e) : undefined));
    } else {
      obj.coins = [];
    }
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryEstimatePoolRemoveResponse>
  ): QueryEstimatePoolRemoveResponse {
    const message = {
      ...baseQueryEstimatePoolRemoveResponse,
    } as QueryEstimatePoolRemoveResponse;
    message.coins = [];
    if (object.coins !== undefined && object.coins !== null) {
      for (const e of object.coins) {
        message.coins.push(Coin.fromPartial(e));
      }
    }
    return message;
  },
};

const baseQueryListRecordsFromPoolRequest: object = { poolName: "" };

export const QueryListRecordsFromPoolRequest = {
  encode(
    message: QueryListRecordsFromPoolRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.poolName !== "") {
      writer.uint32(10).string(message.poolName);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryListRecordsFromPoolRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryListRecordsFromPoolRequest,
    } as QueryListRecordsFromPoolRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.poolName = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryListRecordsFromPoolRequest {
    const message = {
      ...baseQueryListRecordsFromPoolRequest,
    } as QueryListRecordsFromPoolRequest;
    if (object.poolName !== undefined && object.poolName !== null) {
      message.poolName = String(object.poolName);
    } else {
      message.poolName = "";
    }
    return message;
  },

  toJSON(message: QueryListRecordsFromPoolRequest): unknown {
    const obj: any = {};
    message.poolName !== undefined && (obj.poolName = message.poolName);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryListRecordsFromPoolRequest>
  ): QueryListRecordsFromPoolRequest {
    const message = {
      ...baseQueryListRecordsFromPoolRequest,
    } as QueryListRecordsFromPoolRequest;
    if (object.poolName !== undefined && object.poolName !== null) {
      message.poolName = object.poolName;
    } else {
      message.poolName = "";
    }
    return message;
  },
};

const baseQueryListRecordsFromPoolResponse: object = {};

export const QueryListRecordsFromPoolResponse = {
  encode(
    message: QueryListRecordsFromPoolResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.records) {
      LProviderRecord.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryListRecordsFromPoolResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryListRecordsFromPoolResponse,
    } as QueryListRecordsFromPoolResponse;
    message.records = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.records.push(LProviderRecord.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryListRecordsFromPoolResponse {
    const message = {
      ...baseQueryListRecordsFromPoolResponse,
    } as QueryListRecordsFromPoolResponse;
    message.records = [];
    if (object.records !== undefined && object.records !== null) {
      for (const e of object.records) {
        message.records.push(LProviderRecord.fromJSON(e));
      }
    }
    return message;
  },

  toJSON(message: QueryListRecordsFromPoolResponse): unknown {
    const obj: any = {};
    if (message.records) {
      obj.records = message.records.map((e) =>
        e ? LProviderRecord.toJSON(e) : undefined
      );
    } else {
      obj.records = [];
    }
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryListRecordsFromPoolResponse>
  ): QueryListRecordsFromPoolResponse {
    const message = {
      ...baseQueryListRecordsFromPoolResponse,
    } as QueryListRecordsFromPoolResponse;
    message.records = [];
    if (object.records !== undefined && object.records !== null) {
      for (const e of object.records) {
        message.records.push(LProviderRecord.fromPartial(e));
      }
    }
    return message;
  },
};

/** Query defines the gRPC querier service. */
export interface Query {
  /** Parameters queries the parameters of the module. */
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse>;
  /** Queries a LPool by index. */
  LPool(request: QueryGetLPoolRequest): Promise<QueryGetLPoolResponse>;
  /** Queries a list of LPool items. */
  LPoolAll(request: QueryAllLPoolRequest): Promise<QueryAllLPoolResponse>;
  /** Queries a LProviderRecord by pool name and provider address. */
  LProviderRecord(
    request: QueryGetLProviderRecordRequest
  ): Promise<QueryGetLProviderRecordResponse>;
  /** Estimate coin output from a swap. */
  EstimateSwapOut(
    request: QueryEstimateSwapOutRequest
  ): Promise<QueryEstimateSwapOutResponse>;
  /** Estimate coin input to get desired coin output from a swap. */
  EstimateSwapIn(
    request: QueryEstimateSwapInRequest
  ): Promise<QueryEstimateSwapInResponse>;
  /** Estimate coin inputs to get desired amount of LPToken. */
  EstimateContribution(
    request: QueryEstimateContributionRequest
  ): Promise<QueryEstimateContributionResponse>;
  /**
   * Query coins to deposit to make valid liquidity pair.
   * Input one coin and get other coins to deposit to make liquidity pair.
   */
  MakeValidPair(
    request: QueryMakeValidPairRequest
  ): Promise<QueryMakeValidPairResponse>;
  /** Estimate amoutn of coins returned by burning a LPToken. */
  EstimatePoolRemove(
    request: QueryEstimatePoolRemoveRequest
  ): Promise<QueryEstimatePoolRemoveResponse>;
  /** Queries a list of ListRecordsFromPool items. */
  ListRecordsFromPool(
    request: QueryListRecordsFromPoolRequest
  ): Promise<QueryListRecordsFromPoolResponse>;
}

export class QueryClientImpl implements Query {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse> {
    const data = QueryParamsRequest.encode(request).finish();
    const promise = this.rpc.request(
      "jackaldao.canine.lp.Query",
      "Params",
      data
    );
    return promise.then((data) => QueryParamsResponse.decode(new Reader(data)));
  }

  LPool(request: QueryGetLPoolRequest): Promise<QueryGetLPoolResponse> {
    const data = QueryGetLPoolRequest.encode(request).finish();
    const promise = this.rpc.request(
      "jackaldao.canine.lp.Query",
      "LPool",
      data
    );
    return promise.then((data) =>
      QueryGetLPoolResponse.decode(new Reader(data))
    );
  }

  LPoolAll(request: QueryAllLPoolRequest): Promise<QueryAllLPoolResponse> {
    const data = QueryAllLPoolRequest.encode(request).finish();
    const promise = this.rpc.request(
      "jackaldao.canine.lp.Query",
      "LPoolAll",
      data
    );
    return promise.then((data) =>
      QueryAllLPoolResponse.decode(new Reader(data))
    );
  }

  LProviderRecord(
    request: QueryGetLProviderRecordRequest
  ): Promise<QueryGetLProviderRecordResponse> {
    const data = QueryGetLProviderRecordRequest.encode(request).finish();
    const promise = this.rpc.request(
      "jackaldao.canine.lp.Query",
      "LProviderRecord",
      data
    );
    return promise.then((data) =>
      QueryGetLProviderRecordResponse.decode(new Reader(data))
    );
  }

  EstimateSwapOut(
    request: QueryEstimateSwapOutRequest
  ): Promise<QueryEstimateSwapOutResponse> {
    const data = QueryEstimateSwapOutRequest.encode(request).finish();
    const promise = this.rpc.request(
      "jackaldao.canine.lp.Query",
      "EstimateSwapOut",
      data
    );
    return promise.then((data) =>
      QueryEstimateSwapOutResponse.decode(new Reader(data))
    );
  }

  EstimateSwapIn(
    request: QueryEstimateSwapInRequest
  ): Promise<QueryEstimateSwapInResponse> {
    const data = QueryEstimateSwapInRequest.encode(request).finish();
    const promise = this.rpc.request(
      "jackaldao.canine.lp.Query",
      "EstimateSwapIn",
      data
    );
    return promise.then((data) =>
      QueryEstimateSwapInResponse.decode(new Reader(data))
    );
  }

  EstimateContribution(
    request: QueryEstimateContributionRequest
  ): Promise<QueryEstimateContributionResponse> {
    const data = QueryEstimateContributionRequest.encode(request).finish();
    const promise = this.rpc.request(
      "jackaldao.canine.lp.Query",
      "EstimateContribution",
      data
    );
    return promise.then((data) =>
      QueryEstimateContributionResponse.decode(new Reader(data))
    );
  }

  MakeValidPair(
    request: QueryMakeValidPairRequest
  ): Promise<QueryMakeValidPairResponse> {
    const data = QueryMakeValidPairRequest.encode(request).finish();
    const promise = this.rpc.request(
      "jackaldao.canine.lp.Query",
      "MakeValidPair",
      data
    );
    return promise.then((data) =>
      QueryMakeValidPairResponse.decode(new Reader(data))
    );
  }

  EstimatePoolRemove(
    request: QueryEstimatePoolRemoveRequest
  ): Promise<QueryEstimatePoolRemoveResponse> {
    const data = QueryEstimatePoolRemoveRequest.encode(request).finish();
    const promise = this.rpc.request(
      "jackaldao.canine.lp.Query",
      "EstimatePoolRemove",
      data
    );
    return promise.then((data) =>
      QueryEstimatePoolRemoveResponse.decode(new Reader(data))
    );
  }

  ListRecordsFromPool(
    request: QueryListRecordsFromPoolRequest
  ): Promise<QueryListRecordsFromPoolResponse> {
    const data = QueryListRecordsFromPoolRequest.encode(request).finish();
    const promise = this.rpc.request(
      "jackaldao.canine.lp.Query",
      "ListRecordsFromPool",
      data
    );
    return promise.then((data) =>
      QueryListRecordsFromPoolResponse.decode(new Reader(data))
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
