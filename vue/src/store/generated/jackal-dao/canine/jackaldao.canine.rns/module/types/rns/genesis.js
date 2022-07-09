/* eslint-disable */
import { Params } from "../rns/params";
import { Whois } from "../rns/whois";
import { Names } from "../rns/names";
import { Bids } from "../rns/bids";
import { Forsale } from "../rns/forsale";
import { Writer, Reader } from "protobufjs/minimal";
export const protobufPackage = "jackaldao.canine.rns";
const baseGenesisState = {};
export const GenesisState = {
    encode(message, writer = Writer.create()) {
        if (message.params !== undefined) {
            Params.encode(message.params, writer.uint32(10).fork()).ldelim();
        }
        for (const v of message.whoisList) {
            Whois.encode(v, writer.uint32(18).fork()).ldelim();
        }
        for (const v of message.namesList) {
            Names.encode(v, writer.uint32(26).fork()).ldelim();
        }
        for (const v of message.bidsList) {
            Bids.encode(v, writer.uint32(34).fork()).ldelim();
        }
        for (const v of message.forsaleList) {
            Forsale.encode(v, writer.uint32(42).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseGenesisState };
        message.whoisList = [];
        message.namesList = [];
        message.bidsList = [];
        message.forsaleList = [];
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.params = Params.decode(reader, reader.uint32());
                    break;
                case 2:
                    message.whoisList.push(Whois.decode(reader, reader.uint32()));
                    break;
                case 3:
                    message.namesList.push(Names.decode(reader, reader.uint32()));
                    break;
                case 4:
                    message.bidsList.push(Bids.decode(reader, reader.uint32()));
                    break;
                case 5:
                    message.forsaleList.push(Forsale.decode(reader, reader.uint32()));
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
        message.whoisList = [];
        message.namesList = [];
        message.bidsList = [];
        message.forsaleList = [];
        if (object.params !== undefined && object.params !== null) {
            message.params = Params.fromJSON(object.params);
        }
        else {
            message.params = undefined;
        }
        if (object.whoisList !== undefined && object.whoisList !== null) {
            for (const e of object.whoisList) {
                message.whoisList.push(Whois.fromJSON(e));
            }
        }
        if (object.namesList !== undefined && object.namesList !== null) {
            for (const e of object.namesList) {
                message.namesList.push(Names.fromJSON(e));
            }
        }
        if (object.bidsList !== undefined && object.bidsList !== null) {
            for (const e of object.bidsList) {
                message.bidsList.push(Bids.fromJSON(e));
            }
        }
        if (object.forsaleList !== undefined && object.forsaleList !== null) {
            for (const e of object.forsaleList) {
                message.forsaleList.push(Forsale.fromJSON(e));
            }
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.params !== undefined &&
            (obj.params = message.params ? Params.toJSON(message.params) : undefined);
        if (message.whoisList) {
            obj.whoisList = message.whoisList.map((e) => e ? Whois.toJSON(e) : undefined);
        }
        else {
            obj.whoisList = [];
        }
        if (message.namesList) {
            obj.namesList = message.namesList.map((e) => e ? Names.toJSON(e) : undefined);
        }
        else {
            obj.namesList = [];
        }
        if (message.bidsList) {
            obj.bidsList = message.bidsList.map((e) => e ? Bids.toJSON(e) : undefined);
        }
        else {
            obj.bidsList = [];
        }
        if (message.forsaleList) {
            obj.forsaleList = message.forsaleList.map((e) => e ? Forsale.toJSON(e) : undefined);
        }
        else {
            obj.forsaleList = [];
        }
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseGenesisState };
        message.whoisList = [];
        message.namesList = [];
        message.bidsList = [];
        message.forsaleList = [];
        if (object.params !== undefined && object.params !== null) {
            message.params = Params.fromPartial(object.params);
        }
        else {
            message.params = undefined;
        }
        if (object.whoisList !== undefined && object.whoisList !== null) {
            for (const e of object.whoisList) {
                message.whoisList.push(Whois.fromPartial(e));
            }
        }
        if (object.namesList !== undefined && object.namesList !== null) {
            for (const e of object.namesList) {
                message.namesList.push(Names.fromPartial(e));
            }
        }
        if (object.bidsList !== undefined && object.bidsList !== null) {
            for (const e of object.bidsList) {
                message.bidsList.push(Bids.fromPartial(e));
            }
        }
        if (object.forsaleList !== undefined && object.forsaleList !== null) {
            for (const e of object.forsaleList) {
                message.forsaleList.push(Forsale.fromPartial(e));
            }
        }
        return message;
    },
};
