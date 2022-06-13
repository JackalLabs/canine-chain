/* eslint-disable */
import { Reader, Writer } from "protobufjs/minimal";
export const protobufPackage = "jackaldao.canine.jklmining";
const baseMsgAllowSave = { creator: "", passkey: "", size: "" };
export const MsgAllowSave = {
    encode(message, writer = Writer.create()) {
        if (message.creator !== "") {
            writer.uint32(10).string(message.creator);
        }
        if (message.passkey !== "") {
            writer.uint32(18).string(message.passkey);
        }
        if (message.size !== "") {
            writer.uint32(26).string(message.size);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseMsgAllowSave };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.creator = reader.string();
                    break;
                case 2:
                    message.passkey = reader.string();
                    break;
                case 3:
                    message.size = reader.string();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseMsgAllowSave };
        if (object.creator !== undefined && object.creator !== null) {
            message.creator = String(object.creator);
        }
        else {
            message.creator = "";
        }
        if (object.passkey !== undefined && object.passkey !== null) {
            message.passkey = String(object.passkey);
        }
        else {
            message.passkey = "";
        }
        if (object.size !== undefined && object.size !== null) {
            message.size = String(object.size);
        }
        else {
            message.size = "";
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.creator !== undefined && (obj.creator = message.creator);
        message.passkey !== undefined && (obj.passkey = message.passkey);
        message.size !== undefined && (obj.size = message.size);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseMsgAllowSave };
        if (object.creator !== undefined && object.creator !== null) {
            message.creator = object.creator;
        }
        else {
            message.creator = "";
        }
        if (object.passkey !== undefined && object.passkey !== null) {
            message.passkey = object.passkey;
        }
        else {
            message.passkey = "";
        }
        if (object.size !== undefined && object.size !== null) {
            message.size = object.size;
        }
        else {
            message.size = "";
        }
        return message;
    },
};
const baseMsgAllowSaveResponse = {};
export const MsgAllowSaveResponse = {
    encode(_, writer = Writer.create()) {
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseMsgAllowSaveResponse };
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
        const message = { ...baseMsgAllowSaveResponse };
        return message;
    },
    toJSON(_) {
        const obj = {};
        return obj;
    },
    fromPartial(_) {
        const message = { ...baseMsgAllowSaveResponse };
        return message;
    },
};
const baseMsgCreateSaveRequests = {
    creator: "",
    index: "",
    size: "",
    approved: "",
};
export const MsgCreateSaveRequests = {
    encode(message, writer = Writer.create()) {
        if (message.creator !== "") {
            writer.uint32(10).string(message.creator);
        }
        if (message.index !== "") {
            writer.uint32(18).string(message.index);
        }
        if (message.size !== "") {
            writer.uint32(26).string(message.size);
        }
        if (message.approved !== "") {
            writer.uint32(34).string(message.approved);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseMsgCreateSaveRequests };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.creator = reader.string();
                    break;
                case 2:
                    message.index = reader.string();
                    break;
                case 3:
                    message.size = reader.string();
                    break;
                case 4:
                    message.approved = reader.string();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseMsgCreateSaveRequests };
        if (object.creator !== undefined && object.creator !== null) {
            message.creator = String(object.creator);
        }
        else {
            message.creator = "";
        }
        if (object.index !== undefined && object.index !== null) {
            message.index = String(object.index);
        }
        else {
            message.index = "";
        }
        if (object.size !== undefined && object.size !== null) {
            message.size = String(object.size);
        }
        else {
            message.size = "";
        }
        if (object.approved !== undefined && object.approved !== null) {
            message.approved = String(object.approved);
        }
        else {
            message.approved = "";
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.creator !== undefined && (obj.creator = message.creator);
        message.index !== undefined && (obj.index = message.index);
        message.size !== undefined && (obj.size = message.size);
        message.approved !== undefined && (obj.approved = message.approved);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseMsgCreateSaveRequests };
        if (object.creator !== undefined && object.creator !== null) {
            message.creator = object.creator;
        }
        else {
            message.creator = "";
        }
        if (object.index !== undefined && object.index !== null) {
            message.index = object.index;
        }
        else {
            message.index = "";
        }
        if (object.size !== undefined && object.size !== null) {
            message.size = object.size;
        }
        else {
            message.size = "";
        }
        if (object.approved !== undefined && object.approved !== null) {
            message.approved = object.approved;
        }
        else {
            message.approved = "";
        }
        return message;
    },
};
const baseMsgCreateSaveRequestsResponse = {};
export const MsgCreateSaveRequestsResponse = {
    encode(_, writer = Writer.create()) {
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = {
            ...baseMsgCreateSaveRequestsResponse,
        };
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
        const message = {
            ...baseMsgCreateSaveRequestsResponse,
        };
        return message;
    },
    toJSON(_) {
        const obj = {};
        return obj;
    },
    fromPartial(_) {
        const message = {
            ...baseMsgCreateSaveRequestsResponse,
        };
        return message;
    },
};
const baseMsgUpdateSaveRequests = {
    creator: "",
    index: "",
    size: "",
    approved: "",
};
export const MsgUpdateSaveRequests = {
    encode(message, writer = Writer.create()) {
        if (message.creator !== "") {
            writer.uint32(10).string(message.creator);
        }
        if (message.index !== "") {
            writer.uint32(18).string(message.index);
        }
        if (message.size !== "") {
            writer.uint32(26).string(message.size);
        }
        if (message.approved !== "") {
            writer.uint32(34).string(message.approved);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseMsgUpdateSaveRequests };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.creator = reader.string();
                    break;
                case 2:
                    message.index = reader.string();
                    break;
                case 3:
                    message.size = reader.string();
                    break;
                case 4:
                    message.approved = reader.string();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseMsgUpdateSaveRequests };
        if (object.creator !== undefined && object.creator !== null) {
            message.creator = String(object.creator);
        }
        else {
            message.creator = "";
        }
        if (object.index !== undefined && object.index !== null) {
            message.index = String(object.index);
        }
        else {
            message.index = "";
        }
        if (object.size !== undefined && object.size !== null) {
            message.size = String(object.size);
        }
        else {
            message.size = "";
        }
        if (object.approved !== undefined && object.approved !== null) {
            message.approved = String(object.approved);
        }
        else {
            message.approved = "";
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.creator !== undefined && (obj.creator = message.creator);
        message.index !== undefined && (obj.index = message.index);
        message.size !== undefined && (obj.size = message.size);
        message.approved !== undefined && (obj.approved = message.approved);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseMsgUpdateSaveRequests };
        if (object.creator !== undefined && object.creator !== null) {
            message.creator = object.creator;
        }
        else {
            message.creator = "";
        }
        if (object.index !== undefined && object.index !== null) {
            message.index = object.index;
        }
        else {
            message.index = "";
        }
        if (object.size !== undefined && object.size !== null) {
            message.size = object.size;
        }
        else {
            message.size = "";
        }
        if (object.approved !== undefined && object.approved !== null) {
            message.approved = object.approved;
        }
        else {
            message.approved = "";
        }
        return message;
    },
};
const baseMsgUpdateSaveRequestsResponse = {};
export const MsgUpdateSaveRequestsResponse = {
    encode(_, writer = Writer.create()) {
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = {
            ...baseMsgUpdateSaveRequestsResponse,
        };
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
        const message = {
            ...baseMsgUpdateSaveRequestsResponse,
        };
        return message;
    },
    toJSON(_) {
        const obj = {};
        return obj;
    },
    fromPartial(_) {
        const message = {
            ...baseMsgUpdateSaveRequestsResponse,
        };
        return message;
    },
};
const baseMsgDeleteSaveRequests = { creator: "", index: "" };
export const MsgDeleteSaveRequests = {
    encode(message, writer = Writer.create()) {
        if (message.creator !== "") {
            writer.uint32(10).string(message.creator);
        }
        if (message.index !== "") {
            writer.uint32(18).string(message.index);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseMsgDeleteSaveRequests };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.creator = reader.string();
                    break;
                case 2:
                    message.index = reader.string();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseMsgDeleteSaveRequests };
        if (object.creator !== undefined && object.creator !== null) {
            message.creator = String(object.creator);
        }
        else {
            message.creator = "";
        }
        if (object.index !== undefined && object.index !== null) {
            message.index = String(object.index);
        }
        else {
            message.index = "";
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.creator !== undefined && (obj.creator = message.creator);
        message.index !== undefined && (obj.index = message.index);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseMsgDeleteSaveRequests };
        if (object.creator !== undefined && object.creator !== null) {
            message.creator = object.creator;
        }
        else {
            message.creator = "";
        }
        if (object.index !== undefined && object.index !== null) {
            message.index = object.index;
        }
        else {
            message.index = "";
        }
        return message;
    },
};
const baseMsgDeleteSaveRequestsResponse = {};
export const MsgDeleteSaveRequestsResponse = {
    encode(_, writer = Writer.create()) {
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = {
            ...baseMsgDeleteSaveRequestsResponse,
        };
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
        const message = {
            ...baseMsgDeleteSaveRequestsResponse,
        };
        return message;
    },
    toJSON(_) {
        const obj = {};
        return obj;
    },
    fromPartial(_) {
        const message = {
            ...baseMsgDeleteSaveRequestsResponse,
        };
        return message;
    },
};
export class MsgClientImpl {
    constructor(rpc) {
        this.rpc = rpc;
    }
    AllowSave(request) {
        const data = MsgAllowSave.encode(request).finish();
        const promise = this.rpc.request("jackaldao.canine.jklmining.Msg", "AllowSave", data);
        return promise.then((data) => MsgAllowSaveResponse.decode(new Reader(data)));
    }
    CreateSaveRequests(request) {
        const data = MsgCreateSaveRequests.encode(request).finish();
        const promise = this.rpc.request("jackaldao.canine.jklmining.Msg", "CreateSaveRequests", data);
        return promise.then((data) => MsgCreateSaveRequestsResponse.decode(new Reader(data)));
    }
    UpdateSaveRequests(request) {
        const data = MsgUpdateSaveRequests.encode(request).finish();
        const promise = this.rpc.request("jackaldao.canine.jklmining.Msg", "UpdateSaveRequests", data);
        return promise.then((data) => MsgUpdateSaveRequestsResponse.decode(new Reader(data)));
    }
    DeleteSaveRequests(request) {
        const data = MsgDeleteSaveRequests.encode(request).finish();
        const promise = this.rpc.request("jackaldao.canine.jklmining.Msg", "DeleteSaveRequests", data);
        return promise.then((data) => MsgDeleteSaveRequestsResponse.decode(new Reader(data)));
    }
}
