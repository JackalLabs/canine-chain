/* eslint-disable */
import { Reader, Writer } from "protobufjs/minimal";
export const protobufPackage = "jackaldao.canine.rns";
const baseMsgTransfer = { creator: "", name: "", reciever: "" };
export const MsgTransfer = {
    encode(message, writer = Writer.create()) {
        if (message.creator !== "") {
            writer.uint32(10).string(message.creator);
        }
        if (message.name !== "") {
            writer.uint32(18).string(message.name);
        }
        if (message.reciever !== "") {
            writer.uint32(26).string(message.reciever);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseMsgTransfer };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.creator = reader.string();
                    break;
                case 2:
                    message.name = reader.string();
                    break;
                case 3:
                    message.reciever = reader.string();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseMsgTransfer };
        if (object.creator !== undefined && object.creator !== null) {
            message.creator = String(object.creator);
        }
        else {
            message.creator = "";
        }
        if (object.name !== undefined && object.name !== null) {
            message.name = String(object.name);
        }
        else {
            message.name = "";
        }
        if (object.reciever !== undefined && object.reciever !== null) {
            message.reciever = String(object.reciever);
        }
        else {
            message.reciever = "";
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.creator !== undefined && (obj.creator = message.creator);
        message.name !== undefined && (obj.name = message.name);
        message.reciever !== undefined && (obj.reciever = message.reciever);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseMsgTransfer };
        if (object.creator !== undefined && object.creator !== null) {
            message.creator = object.creator;
        }
        else {
            message.creator = "";
        }
        if (object.name !== undefined && object.name !== null) {
            message.name = object.name;
        }
        else {
            message.name = "";
        }
        if (object.reciever !== undefined && object.reciever !== null) {
            message.reciever = object.reciever;
        }
        else {
            message.reciever = "";
        }
        return message;
    },
};
const baseMsgTransferResponse = {};
export const MsgTransferResponse = {
    encode(_, writer = Writer.create()) {
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseMsgTransferResponse };
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
    fromJSON(_) {
        const message = { ...baseMsgTransferResponse };
        return message;
    },
    toJSON(_) {
        const obj = {};
        return obj;
    },
    fromPartial(_) {
        const message = { ...baseMsgTransferResponse };
        return message;
    },
};
export class MsgClientImpl {
    constructor(rpc) {
        this.rpc = rpc;
    }
    Transfer(request) {
        const data = MsgTransfer.encode(request).finish();
        const promise = this.rpc.request("jackaldao.canine.rns.Msg", "Transfer", data);
        return promise.then((data) => MsgTransferResponse.decode(new Reader(data)));
    }
}
