/* eslint-disable */
import { Params } from "../jklmining/params";
import { SaveRequests } from "../jklmining/save_requests";
import { Miners } from "../jklmining/miners";
import { Writer, Reader } from "protobufjs/minimal";
export const protobufPackage = "jackaldao.canine.jklmining";
const baseGenesisState = {};
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
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseGenesisState };
        message.saveRequestsList = [];
        message.minersList = [];
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
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseGenesisState };
        message.saveRequestsList = [];
        message.minersList = [];
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
        return message;
    },
};
