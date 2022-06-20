/* eslint-disable */
import * as Long from "long";
import { util, configure, Writer, Reader } from "protobufjs/minimal";
import { Params } from "../jklmining/params";
import { SaveRequests } from "../jklmining/save_requests";
import { Miners } from "../jklmining/miners";
import { Mined } from "../jklmining/mined";
import { MinerClaims } from "../jklmining/miner_claims";
export const protobufPackage = "jackaldao.canine.jklmining";
const baseGenesisState = { minedCount: 0 };
export const GenesisState = {
    encode(message, writer = Writer.create()) {
        if (message.params !== undefined) {
            Params.encode(message.params, writer.uint32(10).fork()).ldelim();
        }
        for (const v of message.saveRequestsList) {
            SaveRequests.encode(v, writer.uint32(18).fork()).ldelim();
        }
        for (const v of message.minersList) {
            Miners.encode(v, writer.uint32(26).fork()).ldelim();
        }
        for (const v of message.minedList) {
            Mined.encode(v, writer.uint32(34).fork()).ldelim();
        }
        if (message.minedCount !== 0) {
            writer.uint32(40).uint64(message.minedCount);
        }
        for (const v of message.minerClaimsList) {
            MinerClaims.encode(v, writer.uint32(50).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseGenesisState };
        message.saveRequestsList = [];
        message.minersList = [];
        message.minedList = [];
        message.minerClaimsList = [];
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.params = Params.decode(reader, reader.uint32());
                    break;
                case 2:
                    message.saveRequestsList.push(SaveRequests.decode(reader, reader.uint32()));
                    break;
                case 3:
                    message.minersList.push(Miners.decode(reader, reader.uint32()));
                    break;
                case 4:
                    message.minedList.push(Mined.decode(reader, reader.uint32()));
                    break;
                case 5:
                    message.minedCount = longToNumber(reader.uint64());
                    break;
                case 6:
                    message.minerClaimsList.push(MinerClaims.decode(reader, reader.uint32()));
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseGenesisState };
        message.saveRequestsList = [];
        message.minersList = [];
        message.minedList = [];
        message.minerClaimsList = [];
        if (object.params !== undefined && object.params !== null) {
            message.params = Params.fromJSON(object.params);
        }
        else {
            message.params = undefined;
        }
        if (object.saveRequestsList !== undefined &&
            object.saveRequestsList !== null) {
            for (const e of object.saveRequestsList) {
                message.saveRequestsList.push(SaveRequests.fromJSON(e));
            }
        }
        if (object.minersList !== undefined && object.minersList !== null) {
            for (const e of object.minersList) {
                message.minersList.push(Miners.fromJSON(e));
            }
        }
        if (object.minedList !== undefined && object.minedList !== null) {
            for (const e of object.minedList) {
                message.minedList.push(Mined.fromJSON(e));
            }
        }
        if (object.minedCount !== undefined && object.minedCount !== null) {
            message.minedCount = Number(object.minedCount);
        }
        else {
            message.minedCount = 0;
        }
        if (object.minerClaimsList !== undefined &&
            object.minerClaimsList !== null) {
            for (const e of object.minerClaimsList) {
                message.minerClaimsList.push(MinerClaims.fromJSON(e));
            }
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.params !== undefined &&
            (obj.params = message.params ? Params.toJSON(message.params) : undefined);
        if (message.saveRequestsList) {
            obj.saveRequestsList = message.saveRequestsList.map((e) => e ? SaveRequests.toJSON(e) : undefined);
        }
        else {
            obj.saveRequestsList = [];
        }
        if (message.minersList) {
            obj.minersList = message.minersList.map((e) => e ? Miners.toJSON(e) : undefined);
        }
        else {
            obj.minersList = [];
        }
        if (message.minedList) {
            obj.minedList = message.minedList.map((e) => e ? Mined.toJSON(e) : undefined);
        }
        else {
            obj.minedList = [];
        }
        message.minedCount !== undefined && (obj.minedCount = message.minedCount);
        if (message.minerClaimsList) {
            obj.minerClaimsList = message.minerClaimsList.map((e) => e ? MinerClaims.toJSON(e) : undefined);
        }
        else {
            obj.minerClaimsList = [];
        }
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseGenesisState };
        message.saveRequestsList = [];
        message.minersList = [];
        message.minedList = [];
        message.minerClaimsList = [];
        if (object.params !== undefined && object.params !== null) {
            message.params = Params.fromPartial(object.params);
        }
        else {
            message.params = undefined;
        }
        if (object.saveRequestsList !== undefined &&
            object.saveRequestsList !== null) {
            for (const e of object.saveRequestsList) {
                message.saveRequestsList.push(SaveRequests.fromPartial(e));
            }
        }
        if (object.minersList !== undefined && object.minersList !== null) {
            for (const e of object.minersList) {
                message.minersList.push(Miners.fromPartial(e));
            }
        }
        if (object.minedList !== undefined && object.minedList !== null) {
            for (const e of object.minedList) {
                message.minedList.push(Mined.fromPartial(e));
            }
        }
        if (object.minedCount !== undefined && object.minedCount !== null) {
            message.minedCount = object.minedCount;
        }
        else {
            message.minedCount = 0;
        }
        if (object.minerClaimsList !== undefined &&
            object.minerClaimsList !== null) {
            for (const e of object.minerClaimsList) {
                message.minerClaimsList.push(MinerClaims.fromPartial(e));
            }
        }
        return message;
    },
};
var globalThis = (() => {
    if (typeof globalThis !== "undefined")
        return globalThis;
    if (typeof self !== "undefined")
        return self;
    if (typeof window !== "undefined")
        return window;
    if (typeof global !== "undefined")
        return global;
    throw "Unable to locate global object";
})();
function longToNumber(long) {
    if (long.gt(Number.MAX_SAFE_INTEGER)) {
        throw new globalThis.Error("Value is larger than Number.MAX_SAFE_INTEGER");
    }
    return long.toNumber();
}
if (util.Long !== Long) {
    util.Long = Long;
    configure();
}
