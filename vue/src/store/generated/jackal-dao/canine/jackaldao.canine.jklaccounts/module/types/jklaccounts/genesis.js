/* eslint-disable */
import { Params } from "../jklaccounts/params";
import { Accounts } from "../jklaccounts/accounts";
import { Writer, Reader } from "protobufjs/minimal";
export const protobufPackage = "jackaldao.canine.jklaccounts";
const baseGenesisState = {};
export const GenesisState = {
    encode(message, writer = Writer.create()) {
        if (message.params !== undefined) {
            Params.encode(message.params, writer.uint32(10).fork()).ldelim();
        }
        for (const v of message.accountsList) {
            Accounts.encode(v, writer.uint32(18).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseGenesisState };
        message.accountsList = [];
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.params = Params.decode(reader, reader.uint32());
                    break;
                case 2:
                    message.accountsList.push(Accounts.decode(reader, reader.uint32()));
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
        message.accountsList = [];
        if (object.params !== undefined && object.params !== null) {
            message.params = Params.fromJSON(object.params);
        }
        else {
            message.params = undefined;
        }
        if (object.accountsList !== undefined && object.accountsList !== null) {
            for (const e of object.accountsList) {
                message.accountsList.push(Accounts.fromJSON(e));
            }
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.params !== undefined &&
            (obj.params = message.params ? Params.toJSON(message.params) : undefined);
        if (message.accountsList) {
            obj.accountsList = message.accountsList.map((e) => e ? Accounts.toJSON(e) : undefined);
        }
        else {
            obj.accountsList = [];
        }
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseGenesisState };
        message.accountsList = [];
        if (object.params !== undefined && object.params !== null) {
            message.params = Params.fromPartial(object.params);
        }
        else {
            message.params = undefined;
        }
        if (object.accountsList !== undefined && object.accountsList !== null) {
            for (const e of object.accountsList) {
                message.accountsList.push(Accounts.fromPartial(e));
            }
        }
        return message;
    },
};
