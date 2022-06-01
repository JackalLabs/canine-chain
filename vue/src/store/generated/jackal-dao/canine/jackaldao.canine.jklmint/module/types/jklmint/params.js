/* eslint-disable */
import { Writer, Reader } from "protobufjs/minimal";
export const protobufPackage = "jackaldao.canine.jklmint";
const baseParams = { mintDenom: "" };
export const Params = {
    encode(message, writer = Writer.create()) {
        if (message.mintDenom !== "") {
            writer.uint32(10).string(message.mintDenom);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseParams };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.mintDenom = reader.string();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseParams };
        if (object.mintDenom !== undefined && object.mintDenom !== null) {
            message.mintDenom = String(object.mintDenom);
        }
        else {
            message.mintDenom = "";
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.mintDenom !== undefined && (obj.mintDenom = message.mintDenom);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseParams };
        if (object.mintDenom !== undefined && object.mintDenom !== null) {
            message.mintDenom = object.mintDenom;
        }
        else {
            message.mintDenom = "";
        }
        return message;
    },
};
