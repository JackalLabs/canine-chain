/* eslint-disable */
import { Reader, Writer } from "protobufjs/minimal";

export const protobufPackage = "jackaldao.canine.jklaccounts";

export interface MsgChoosePlan {
  creator: string;
  tbCount: string;
  paymentDenom: string;
}

export interface MsgChoosePlanResponse {}

export interface MsgPayMonths {
  creator: string;
  address: string;
  months: string;
  paymentDenom: string;
}

export interface MsgPayMonthsResponse {}

const baseMsgChoosePlan: object = {
  creator: "",
  tbCount: "",
  paymentDenom: "",
};

export const MsgChoosePlan = {
  encode(message: MsgChoosePlan, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.tbCount !== "") {
      writer.uint32(18).string(message.tbCount);
    }
    if (message.paymentDenom !== "") {
      writer.uint32(26).string(message.paymentDenom);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgChoosePlan {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgChoosePlan } as MsgChoosePlan;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.tbCount = reader.string();
          break;
        case 3:
          message.paymentDenom = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgChoosePlan {
    const message = { ...baseMsgChoosePlan } as MsgChoosePlan;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.tbCount !== undefined && object.tbCount !== null) {
      message.tbCount = String(object.tbCount);
    } else {
      message.tbCount = "";
    }
    if (object.paymentDenom !== undefined && object.paymentDenom !== null) {
      message.paymentDenom = String(object.paymentDenom);
    } else {
      message.paymentDenom = "";
    }
    return message;
  },

  toJSON(message: MsgChoosePlan): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.tbCount !== undefined && (obj.tbCount = message.tbCount);
    message.paymentDenom !== undefined &&
      (obj.paymentDenom = message.paymentDenom);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgChoosePlan>): MsgChoosePlan {
    const message = { ...baseMsgChoosePlan } as MsgChoosePlan;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.tbCount !== undefined && object.tbCount !== null) {
      message.tbCount = object.tbCount;
    } else {
      message.tbCount = "";
    }
    if (object.paymentDenom !== undefined && object.paymentDenom !== null) {
      message.paymentDenom = object.paymentDenom;
    } else {
      message.paymentDenom = "";
    }
    return message;
  },
};

const baseMsgChoosePlanResponse: object = {};

export const MsgChoosePlanResponse = {
  encode(_: MsgChoosePlanResponse, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgChoosePlanResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgChoosePlanResponse } as MsgChoosePlanResponse;
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

  fromJSON(_: any): MsgChoosePlanResponse {
    const message = { ...baseMsgChoosePlanResponse } as MsgChoosePlanResponse;
    return message;
  },

  toJSON(_: MsgChoosePlanResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(_: DeepPartial<MsgChoosePlanResponse>): MsgChoosePlanResponse {
    const message = { ...baseMsgChoosePlanResponse } as MsgChoosePlanResponse;
    return message;
  },
};

const baseMsgPayMonths: object = {
  creator: "",
  address: "",
  months: "",
  paymentDenom: "",
};

export const MsgPayMonths = {
  encode(message: MsgPayMonths, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.address !== "") {
      writer.uint32(18).string(message.address);
    }
    if (message.months !== "") {
      writer.uint32(26).string(message.months);
    }
    if (message.paymentDenom !== "") {
      writer.uint32(34).string(message.paymentDenom);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgPayMonths {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgPayMonths } as MsgPayMonths;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.address = reader.string();
          break;
        case 3:
          message.months = reader.string();
          break;
        case 4:
          message.paymentDenom = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgPayMonths {
    const message = { ...baseMsgPayMonths } as MsgPayMonths;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.address !== undefined && object.address !== null) {
      message.address = String(object.address);
    } else {
      message.address = "";
    }
    if (object.months !== undefined && object.months !== null) {
      message.months = String(object.months);
    } else {
      message.months = "";
    }
    if (object.paymentDenom !== undefined && object.paymentDenom !== null) {
      message.paymentDenom = String(object.paymentDenom);
    } else {
      message.paymentDenom = "";
    }
    return message;
  },

  toJSON(message: MsgPayMonths): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.address !== undefined && (obj.address = message.address);
    message.months !== undefined && (obj.months = message.months);
    message.paymentDenom !== undefined &&
      (obj.paymentDenom = message.paymentDenom);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgPayMonths>): MsgPayMonths {
    const message = { ...baseMsgPayMonths } as MsgPayMonths;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.address !== undefined && object.address !== null) {
      message.address = object.address;
    } else {
      message.address = "";
    }
    if (object.months !== undefined && object.months !== null) {
      message.months = object.months;
    } else {
      message.months = "";
    }
    if (object.paymentDenom !== undefined && object.paymentDenom !== null) {
      message.paymentDenom = object.paymentDenom;
    } else {
      message.paymentDenom = "";
    }
    return message;
  },
};

const baseMsgPayMonthsResponse: object = {};

export const MsgPayMonthsResponse = {
  encode(_: MsgPayMonthsResponse, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgPayMonthsResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgPayMonthsResponse } as MsgPayMonthsResponse;
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

  fromJSON(_: any): MsgPayMonthsResponse {
    const message = { ...baseMsgPayMonthsResponse } as MsgPayMonthsResponse;
    return message;
  },

  toJSON(_: MsgPayMonthsResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(_: DeepPartial<MsgPayMonthsResponse>): MsgPayMonthsResponse {
    const message = { ...baseMsgPayMonthsResponse } as MsgPayMonthsResponse;
    return message;
  },
};

/** Msg defines the Msg service. */
export interface Msg {
  ChoosePlan(request: MsgChoosePlan): Promise<MsgChoosePlanResponse>;
  /** this line is used by starport scaffolding # proto/tx/rpc */
  PayMonths(request: MsgPayMonths): Promise<MsgPayMonthsResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  ChoosePlan(request: MsgChoosePlan): Promise<MsgChoosePlanResponse> {
    const data = MsgChoosePlan.encode(request).finish();
    const promise = this.rpc.request(
      "jackaldao.canine.jklaccounts.Msg",
      "ChoosePlan",
      data
    );
    return promise.then((data) =>
      MsgChoosePlanResponse.decode(new Reader(data))
    );
  }

  PayMonths(request: MsgPayMonths): Promise<MsgPayMonthsResponse> {
    const data = MsgPayMonths.encode(request).finish();
    const promise = this.rpc.request(
      "jackaldao.canine.jklaccounts.Msg",
      "PayMonths",
      data
    );
    return promise.then((data) =>
      MsgPayMonthsResponse.decode(new Reader(data))
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
