/* eslint-disable */
import { Params } from "../storage/params";
import { Contracts } from "../storage/contracts";
import { Proofs } from "../storage/proofs";
import { ActiveDeals } from "../storage/active_deals";
import { Miners } from "../storage/miners";
import { PayBlocks } from "../storage/pay_blocks";
import { ClientUsage } from "../storage/client_usage";
import { Strays } from "../storage/strays";
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "jackaldao.canine.storage";

/** GenesisState defines the storage module's genesis state. */
export interface GenesisState {
  params: Params | undefined;
  contractsList: Contracts[];
  proofsList: Proofs[];
  activeDealsList: ActiveDeals[];
  minersList: Miners[];
  payBlocksList: PayBlocks[];
  clientUsageList: ClientUsage[];
  /** this line is used by starport scaffolding # genesis/proto/state */
  straysList: Strays[];
}

const baseGenesisState: object = {};

export const GenesisState = {
  encode(message: GenesisState, writer: Writer = Writer.create()): Writer {
    if (message.params !== undefined) {
      Params.encode(message.params, writer.uint32(10).fork()).ldelim();
    }
    for (const v of message.contractsList) {
      Contracts.encode(v!, writer.uint32(18).fork()).ldelim();
    }
    for (const v of message.proofsList) {
      Proofs.encode(v!, writer.uint32(26).fork()).ldelim();
    }
    for (const v of message.activeDealsList) {
      ActiveDeals.encode(v!, writer.uint32(34).fork()).ldelim();
    }
    for (const v of message.minersList) {
      Miners.encode(v!, writer.uint32(42).fork()).ldelim();
    }
    for (const v of message.payBlocksList) {
      PayBlocks.encode(v!, writer.uint32(50).fork()).ldelim();
    }
    for (const v of message.clientUsageList) {
      ClientUsage.encode(v!, writer.uint32(58).fork()).ldelim();
    }
    for (const v of message.straysList) {
      Strays.encode(v!, writer.uint32(66).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): GenesisState {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseGenesisState } as GenesisState;
    message.contractsList = [];
    message.proofsList = [];
    message.activeDealsList = [];
    message.minersList = [];
    message.payBlocksList = [];
    message.clientUsageList = [];
    message.straysList = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.params = Params.decode(reader, reader.uint32());
          break;
        case 2:
          message.contractsList.push(Contracts.decode(reader, reader.uint32()));
          break;
        case 3:
          message.proofsList.push(Proofs.decode(reader, reader.uint32()));
          break;
        case 4:
          message.activeDealsList.push(
            ActiveDeals.decode(reader, reader.uint32())
          );
          break;
        case 5:
          message.minersList.push(Miners.decode(reader, reader.uint32()));
          break;
        case 6:
          message.payBlocksList.push(PayBlocks.decode(reader, reader.uint32()));
          break;
        case 7:
          message.clientUsageList.push(
            ClientUsage.decode(reader, reader.uint32())
          );
          break;
        case 8:
          message.straysList.push(Strays.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): GenesisState {
    const message = { ...baseGenesisState } as GenesisState;
    message.contractsList = [];
    message.proofsList = [];
    message.activeDealsList = [];
    message.minersList = [];
    message.payBlocksList = [];
    message.clientUsageList = [];
    message.straysList = [];
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromJSON(object.params);
    } else {
      message.params = undefined;
    }
    if (object.contractsList !== undefined && object.contractsList !== null) {
      for (const e of object.contractsList) {
        message.contractsList.push(Contracts.fromJSON(e));
      }
    }
    if (object.proofsList !== undefined && object.proofsList !== null) {
      for (const e of object.proofsList) {
        message.proofsList.push(Proofs.fromJSON(e));
      }
    }
    if (
      object.activeDealsList !== undefined &&
      object.activeDealsList !== null
    ) {
      for (const e of object.activeDealsList) {
        message.activeDealsList.push(ActiveDeals.fromJSON(e));
      }
    }
    if (object.minersList !== undefined && object.minersList !== null) {
      for (const e of object.minersList) {
        message.minersList.push(Miners.fromJSON(e));
      }
    }
    if (object.payBlocksList !== undefined && object.payBlocksList !== null) {
      for (const e of object.payBlocksList) {
        message.payBlocksList.push(PayBlocks.fromJSON(e));
      }
    }
    if (
      object.clientUsageList !== undefined &&
      object.clientUsageList !== null
    ) {
      for (const e of object.clientUsageList) {
        message.clientUsageList.push(ClientUsage.fromJSON(e));
      }
    }
    if (object.straysList !== undefined && object.straysList !== null) {
      for (const e of object.straysList) {
        message.straysList.push(Strays.fromJSON(e));
      }
    }
    return message;
  },

  toJSON(message: GenesisState): unknown {
    const obj: any = {};
    message.params !== undefined &&
      (obj.params = message.params ? Params.toJSON(message.params) : undefined);
    if (message.contractsList) {
      obj.contractsList = message.contractsList.map((e) =>
        e ? Contracts.toJSON(e) : undefined
      );
    } else {
      obj.contractsList = [];
    }
    if (message.proofsList) {
      obj.proofsList = message.proofsList.map((e) =>
        e ? Proofs.toJSON(e) : undefined
      );
    } else {
      obj.proofsList = [];
    }
    if (message.activeDealsList) {
      obj.activeDealsList = message.activeDealsList.map((e) =>
        e ? ActiveDeals.toJSON(e) : undefined
      );
    } else {
      obj.activeDealsList = [];
    }
    if (message.minersList) {
      obj.minersList = message.minersList.map((e) =>
        e ? Miners.toJSON(e) : undefined
      );
    } else {
      obj.minersList = [];
    }
    if (message.payBlocksList) {
      obj.payBlocksList = message.payBlocksList.map((e) =>
        e ? PayBlocks.toJSON(e) : undefined
      );
    } else {
      obj.payBlocksList = [];
    }
    if (message.clientUsageList) {
      obj.clientUsageList = message.clientUsageList.map((e) =>
        e ? ClientUsage.toJSON(e) : undefined
      );
    } else {
      obj.clientUsageList = [];
    }
    if (message.straysList) {
      obj.straysList = message.straysList.map((e) =>
        e ? Strays.toJSON(e) : undefined
      );
    } else {
      obj.straysList = [];
    }
    return obj;
  },

  fromPartial(object: DeepPartial<GenesisState>): GenesisState {
    const message = { ...baseGenesisState } as GenesisState;
    message.contractsList = [];
    message.proofsList = [];
    message.activeDealsList = [];
    message.minersList = [];
    message.payBlocksList = [];
    message.clientUsageList = [];
    message.straysList = [];
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromPartial(object.params);
    } else {
      message.params = undefined;
    }
    if (object.contractsList !== undefined && object.contractsList !== null) {
      for (const e of object.contractsList) {
        message.contractsList.push(Contracts.fromPartial(e));
      }
    }
    if (object.proofsList !== undefined && object.proofsList !== null) {
      for (const e of object.proofsList) {
        message.proofsList.push(Proofs.fromPartial(e));
      }
    }
    if (
      object.activeDealsList !== undefined &&
      object.activeDealsList !== null
    ) {
      for (const e of object.activeDealsList) {
        message.activeDealsList.push(ActiveDeals.fromPartial(e));
      }
    }
    if (object.minersList !== undefined && object.minersList !== null) {
      for (const e of object.minersList) {
        message.minersList.push(Miners.fromPartial(e));
      }
    }
    if (object.payBlocksList !== undefined && object.payBlocksList !== null) {
      for (const e of object.payBlocksList) {
        message.payBlocksList.push(PayBlocks.fromPartial(e));
      }
    }
    if (
      object.clientUsageList !== undefined &&
      object.clientUsageList !== null
    ) {
      for (const e of object.clientUsageList) {
        message.clientUsageList.push(ClientUsage.fromPartial(e));
      }
    }
    if (object.straysList !== undefined && object.straysList !== null) {
      for (const e of object.straysList) {
        message.straysList.push(Strays.fromPartial(e));
      }
    }
    return message;
  },
};

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
