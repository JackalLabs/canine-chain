/* eslint-disable */
import { Writer, Reader } from "protobufjs/minimal";
export const protobufPackage = "jackaldao.canine.jklaccounts";
const baseAccounts = {
    address: "",
    available: "",
    used: "",
    expireBlock: "",
};
export const Accounts = {
    encode(message, writer = Writer.create()) {
        if (message.address !== "") {
            writer.uint32(10).string(message.address);
        }
        if (message.available !== "") {
            writer.uint32(18).string(message.available);
        }
        if (message.used !== "") {
            writer.uint32(26).string(message.used);
        }
        if (message.expireBlock !== "") {
            writer.uint32(34).string(message.expireBlock);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseAccounts };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.address = reader.string();
                    break;
                case 2:
                    message.available = reader.string();
                    break;
                case 3:
                    message.used = reader.string();
                    break;
                case 4:
                    message.expireBlock = reader.string();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseAccounts };
        if (object.address !== undefined && object.address !== null) {
            message.address = String(object.address);
        }
        else {
            message.address = "";
        }
        if (object.available !== undefined && object.available !== null) {
            message.available = String(object.available);
        }
        else {
            message.available = "";
        }
        if (object.used !== undefined && object.used !== null) {
            message.used = String(object.used);
        }
        else {
            message.used = "";
        }
        if (object.expireBlock !== undefined && object.expireBlock !== null) {
            message.expireBlock = String(object.expireBlock);
        }
        else {
            message.expireBlock = "";
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.address !== undefined && (obj.address = message.address);
        message.available !== undefined && (obj.available = message.available);
        message.used !== undefined && (obj.used = message.used);
        message.expireBlock !== undefined &&
            (obj.expireBlock = message.expireBlock);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseAccounts };
        if (object.address !== undefined && object.address !== null) {
            message.address = object.address;
        }
        else {
            message.address = "";
        }
        if (object.available !== undefined && object.available !== null) {
            message.available = object.available;
        }
        else {
            message.available = "";
        }
        if (object.used !== undefined && object.used !== null) {
            message.used = object.used;
        }
        else {
            message.used = "";
        }
        if (object.expireBlock !== undefined && object.expireBlock !== null) {
            message.expireBlock = object.expireBlock;
        }
        else {
            message.expireBlock = "";
        }
        return message;
    },
};
