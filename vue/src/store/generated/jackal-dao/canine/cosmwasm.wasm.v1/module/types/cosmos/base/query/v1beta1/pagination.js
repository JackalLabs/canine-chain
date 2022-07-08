/* eslint-disable */
import * as Long from "long";
import { util, configure, Writer, Reader } from "protobufjs/minimal";
export const protobufPackage = "cosmos.base.query.v1beta1";
const basePageRequest = {
    offset: 0,
    limit: 0,
    count_total: false,
    reverse: false,
};
export const PageRequest = {
    encode(message, writer = Writer.create()) {
        if (message.key.length !== 0) {
            writer.uint32(10).bytes(message.key);
        }
        if (message.offset !== 0) {
            writer.uint32(16).uint64(message.offset);
        }
        if (message.limit !== 0) {
            writer.uint32(24).uint64(message.limit);
        }
        if (message.count_total === true) {
            writer.uint32(32).bool(message.count_total);
        }
        if (message.reverse === true) {
            writer.uint32(40).bool(message.reverse);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...basePageRequest };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.key = reader.bytes();
                    break;
                case 2:
                    message.offset = longToNumber(reader.uint64());
                    break;
                case 3:
                    message.limit = longToNumber(reader.uint64());
                    break;
                case 4:
                    message.count_total = reader.bool();
                    break;
                case 5:
                    message.reverse = reader.bool();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...basePageRequest };
        if (object.key !== undefined && object.key !== null) {
            message.key = bytesFromBase64(object.key);
        }
        if (object.offset !== undefined && object.offset !== null) {
            message.offset = Number(object.offset);
        }
        else {
            message.offset = 0;
        }
        if (object.limit !== undefined && object.limit !== null) {
            message.limit = Number(object.limit);
        }
        else {
            message.limit = 0;
        }
        if (object.count_total !== undefined && object.count_total !== null) {
            message.count_total = Boolean(object.count_total);
        }
        else {
            message.count_total = false;
        }
        if (object.reverse !== undefined && object.reverse !== null) {
            message.reverse = Boolean(object.reverse);
        }
        else {
            message.reverse = false;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.key !== undefined &&
            (obj.key = base64FromBytes(message.key !== undefined ? message.key : new Uint8Array()));
        message.offset !== undefined && (obj.offset = message.offset);
        message.limit !== undefined && (obj.limit = message.limit);
        message.count_total !== undefined &&
            (obj.count_total = message.count_total);
        message.reverse !== undefined && (obj.reverse = message.reverse);
        return obj;
    },
    fromPartial(object) {
        const message = { ...basePageRequest };
        if (object.key !== undefined && object.key !== null) {
            message.key = object.key;
        }
        else {
            message.key = new Uint8Array();
        }
        if (object.offset !== undefined && object.offset !== null) {
            message.offset = object.offset;
        }
        else {
            message.offset = 0;
        }
        if (object.limit !== undefined && object.limit !== null) {
            message.limit = object.limit;
        }
        else {
            message.limit = 0;
        }
        if (object.count_total !== undefined && object.count_total !== null) {
            message.count_total = object.count_total;
        }
        else {
            message.count_total = false;
        }
        if (object.reverse !== undefined && object.reverse !== null) {
            message.reverse = object.reverse;
        }
        else {
            message.reverse = false;
        }
        return message;
    },
};
const basePageResponse = { total: 0 };
export const PageResponse = {
    encode(message, writer = Writer.create()) {
        if (message.next_key.length !== 0) {
            writer.uint32(10).bytes(message.next_key);
        }
        if (message.total !== 0) {
            writer.uint32(16).uint64(message.total);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...basePageResponse };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.next_key = reader.bytes();
                    break;
                case 2:
                    message.total = longToNumber(reader.uint64());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...basePageResponse };
        if (object.next_key !== undefined && object.next_key !== null) {
            message.next_key = bytesFromBase64(object.next_key);
        }
        if (object.total !== undefined && object.total !== null) {
            message.total = Number(object.total);
        }
        else {
            message.total = 0;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.next_key !== undefined &&
            (obj.next_key = base64FromBytes(message.next_key !== undefined ? message.next_key : new Uint8Array()));
        message.total !== undefined && (obj.total = message.total);
        return obj;
    },
    fromPartial(object) {
        const message = { ...basePageResponse };
        if (object.next_key !== undefined && object.next_key !== null) {
            message.next_key = object.next_key;
        }
        else {
            message.next_key = new Uint8Array();
        }
        if (object.total !== undefined && object.total !== null) {
            message.total = object.total;
        }
        else {
            message.total = 0;
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
const atob = globalThis.atob ||
    ((b64) => globalThis.Buffer.from(b64, "base64").toString("binary"));
function bytesFromBase64(b64) {
    const bin = atob(b64);
    const arr = new Uint8Array(bin.length);
    for (let i = 0; i < bin.length; ++i) {
        arr[i] = bin.charCodeAt(i);
    }
    return arr;
}
const btoa = globalThis.btoa ||
    ((bin) => globalThis.Buffer.from(bin, "binary").toString("base64"));
function base64FromBytes(arr) {
    const bin = [];
    for (let i = 0; i < arr.byteLength; ++i) {
        bin.push(String.fromCharCode(arr[i]));
    }
    return btoa(bin.join(""));
}
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
