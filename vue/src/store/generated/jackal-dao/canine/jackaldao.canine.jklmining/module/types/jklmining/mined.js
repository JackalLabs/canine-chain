/* eslint-disable */
import * as Long from "long";
import { util, configure, Writer, Reader } from "protobufjs/minimal";
export const protobufPackage = "jackaldao.canine.jklmining";
const baseMined = { id: 0, datasize: "", hash: "", pcount: "" };
export const Mined = {
    encode(message, writer = Writer.create()) {
        if (message.id !== 0) {
            writer.uint32(8).uint64(message.id);
        }
        if (message.datasize !== "") {
            writer.uint32(18).string(message.datasize);
        }
        if (message.hash !== "") {
            writer.uint32(26).string(message.hash);
        }
        if (message.pcount !== "") {
            writer.uint32(34).string(message.pcount);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseMined };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.id = longToNumber(reader.uint64());
                    break;
                case 2:
                    message.datasize = reader.string();
                    break;
                case 3:
                    message.hash = reader.string();
                    break;
                case 4:
                    message.pcount = reader.string();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseMined };
        if (object.id !== undefined && object.id !== null) {
            message.id = Number(object.id);
        }
        else {
            message.id = 0;
        }
        if (object.datasize !== undefined && object.datasize !== null) {
            message.datasize = String(object.datasize);
        }
        else {
            message.datasize = "";
        }
        if (object.hash !== undefined && object.hash !== null) {
            message.hash = String(object.hash);
        }
        else {
            message.hash = "";
        }
        if (object.pcount !== undefined && object.pcount !== null) {
            message.pcount = String(object.pcount);
        }
        else {
            message.pcount = "";
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.id !== undefined && (obj.id = message.id);
        message.datasize !== undefined && (obj.datasize = message.datasize);
        message.hash !== undefined && (obj.hash = message.hash);
        message.pcount !== undefined && (obj.pcount = message.pcount);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseMined };
        if (object.id !== undefined && object.id !== null) {
            message.id = object.id;
        }
        else {
            message.id = 0;
        }
        if (object.datasize !== undefined && object.datasize !== null) {
            message.datasize = object.datasize;
        }
        else {
            message.datasize = "";
        }
        if (object.hash !== undefined && object.hash !== null) {
            message.hash = object.hash;
        }
        else {
            message.hash = "";
        }
        if (object.pcount !== undefined && object.pcount !== null) {
            message.pcount = object.pcount;
        }
        else {
            message.pcount = "";
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
