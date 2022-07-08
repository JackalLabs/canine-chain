/* eslint-disable */
import { Reader, util, configure, Writer } from "protobufjs/minimal";
import * as Long from "long";
import { ContractInfo, ContractCodeHistoryEntry, Model, AccessConfig, } from "../../../cosmwasm/wasm/v1/types";
import { PageRequest, PageResponse, } from "../../../cosmos/base/query/v1beta1/pagination";
export const protobufPackage = "cosmwasm.wasm.v1";
const baseQueryContractInfoRequest = { address: "" };
export const QueryContractInfoRequest = {
    encode(message, writer = Writer.create()) {
        if (message.address !== "") {
            writer.uint32(10).string(message.address);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = {
            ...baseQueryContractInfoRequest,
        };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.address = reader.string();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = {
            ...baseQueryContractInfoRequest,
        };
        if (object.address !== undefined && object.address !== null) {
            message.address = String(object.address);
        }
        else {
            message.address = "";
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.address !== undefined && (obj.address = message.address);
        return obj;
    },
    fromPartial(object) {
        const message = {
            ...baseQueryContractInfoRequest,
        };
        if (object.address !== undefined && object.address !== null) {
            message.address = object.address;
        }
        else {
            message.address = "";
        }
        return message;
    },
};
const baseQueryContractInfoResponse = { address: "" };
export const QueryContractInfoResponse = {
    encode(message, writer = Writer.create()) {
        if (message.address !== "") {
            writer.uint32(10).string(message.address);
        }
        if (message.contract_info !== undefined) {
            ContractInfo.encode(message.contract_info, writer.uint32(18).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = {
            ...baseQueryContractInfoResponse,
        };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.address = reader.string();
                    break;
                case 2:
                    message.contract_info = ContractInfo.decode(reader, reader.uint32());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = {
            ...baseQueryContractInfoResponse,
        };
        if (object.address !== undefined && object.address !== null) {
            message.address = String(object.address);
        }
        else {
            message.address = "";
        }
        if (object.contract_info !== undefined && object.contract_info !== null) {
            message.contract_info = ContractInfo.fromJSON(object.contract_info);
        }
        else {
            message.contract_info = undefined;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.address !== undefined && (obj.address = message.address);
        message.contract_info !== undefined &&
            (obj.contract_info = message.contract_info
                ? ContractInfo.toJSON(message.contract_info)
                : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = {
            ...baseQueryContractInfoResponse,
        };
        if (object.address !== undefined && object.address !== null) {
            message.address = object.address;
        }
        else {
            message.address = "";
        }
        if (object.contract_info !== undefined && object.contract_info !== null) {
            message.contract_info = ContractInfo.fromPartial(object.contract_info);
        }
        else {
            message.contract_info = undefined;
        }
        return message;
    },
};
const baseQueryContractHistoryRequest = { address: "" };
export const QueryContractHistoryRequest = {
    encode(message, writer = Writer.create()) {
        if (message.address !== "") {
            writer.uint32(10).string(message.address);
        }
        if (message.pagination !== undefined) {
            PageRequest.encode(message.pagination, writer.uint32(18).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = {
            ...baseQueryContractHistoryRequest,
        };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.address = reader.string();
                    break;
                case 2:
                    message.pagination = PageRequest.decode(reader, reader.uint32());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = {
            ...baseQueryContractHistoryRequest,
        };
        if (object.address !== undefined && object.address !== null) {
            message.address = String(object.address);
        }
        else {
            message.address = "";
        }
        if (object.pagination !== undefined && object.pagination !== null) {
            message.pagination = PageRequest.fromJSON(object.pagination);
        }
        else {
            message.pagination = undefined;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.address !== undefined && (obj.address = message.address);
        message.pagination !== undefined &&
            (obj.pagination = message.pagination
                ? PageRequest.toJSON(message.pagination)
                : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = {
            ...baseQueryContractHistoryRequest,
        };
        if (object.address !== undefined && object.address !== null) {
            message.address = object.address;
        }
        else {
            message.address = "";
        }
        if (object.pagination !== undefined && object.pagination !== null) {
            message.pagination = PageRequest.fromPartial(object.pagination);
        }
        else {
            message.pagination = undefined;
        }
        return message;
    },
};
const baseQueryContractHistoryResponse = {};
export const QueryContractHistoryResponse = {
    encode(message, writer = Writer.create()) {
        for (const v of message.entries) {
            ContractCodeHistoryEntry.encode(v, writer.uint32(10).fork()).ldelim();
        }
        if (message.pagination !== undefined) {
            PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = {
            ...baseQueryContractHistoryResponse,
        };
        message.entries = [];
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.entries.push(ContractCodeHistoryEntry.decode(reader, reader.uint32()));
                    break;
                case 2:
                    message.pagination = PageResponse.decode(reader, reader.uint32());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = {
            ...baseQueryContractHistoryResponse,
        };
        message.entries = [];
        if (object.entries !== undefined && object.entries !== null) {
            for (const e of object.entries) {
                message.entries.push(ContractCodeHistoryEntry.fromJSON(e));
            }
        }
        if (object.pagination !== undefined && object.pagination !== null) {
            message.pagination = PageResponse.fromJSON(object.pagination);
        }
        else {
            message.pagination = undefined;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        if (message.entries) {
            obj.entries = message.entries.map((e) => e ? ContractCodeHistoryEntry.toJSON(e) : undefined);
        }
        else {
            obj.entries = [];
        }
        message.pagination !== undefined &&
            (obj.pagination = message.pagination
                ? PageResponse.toJSON(message.pagination)
                : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = {
            ...baseQueryContractHistoryResponse,
        };
        message.entries = [];
        if (object.entries !== undefined && object.entries !== null) {
            for (const e of object.entries) {
                message.entries.push(ContractCodeHistoryEntry.fromPartial(e));
            }
        }
        if (object.pagination !== undefined && object.pagination !== null) {
            message.pagination = PageResponse.fromPartial(object.pagination);
        }
        else {
            message.pagination = undefined;
        }
        return message;
    },
};
const baseQueryContractsByCodeRequest = { code_id: 0 };
export const QueryContractsByCodeRequest = {
    encode(message, writer = Writer.create()) {
        if (message.code_id !== 0) {
            writer.uint32(8).uint64(message.code_id);
        }
        if (message.pagination !== undefined) {
            PageRequest.encode(message.pagination, writer.uint32(18).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = {
            ...baseQueryContractsByCodeRequest,
        };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.code_id = longToNumber(reader.uint64());
                    break;
                case 2:
                    message.pagination = PageRequest.decode(reader, reader.uint32());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = {
            ...baseQueryContractsByCodeRequest,
        };
        if (object.code_id !== undefined && object.code_id !== null) {
            message.code_id = Number(object.code_id);
        }
        else {
            message.code_id = 0;
        }
        if (object.pagination !== undefined && object.pagination !== null) {
            message.pagination = PageRequest.fromJSON(object.pagination);
        }
        else {
            message.pagination = undefined;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.code_id !== undefined && (obj.code_id = message.code_id);
        message.pagination !== undefined &&
            (obj.pagination = message.pagination
                ? PageRequest.toJSON(message.pagination)
                : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = {
            ...baseQueryContractsByCodeRequest,
        };
        if (object.code_id !== undefined && object.code_id !== null) {
            message.code_id = object.code_id;
        }
        else {
            message.code_id = 0;
        }
        if (object.pagination !== undefined && object.pagination !== null) {
            message.pagination = PageRequest.fromPartial(object.pagination);
        }
        else {
            message.pagination = undefined;
        }
        return message;
    },
};
const baseQueryContractsByCodeResponse = { contracts: "" };
export const QueryContractsByCodeResponse = {
    encode(message, writer = Writer.create()) {
        for (const v of message.contracts) {
            writer.uint32(10).string(v);
        }
        if (message.pagination !== undefined) {
            PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = {
            ...baseQueryContractsByCodeResponse,
        };
        message.contracts = [];
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.contracts.push(reader.string());
                    break;
                case 2:
                    message.pagination = PageResponse.decode(reader, reader.uint32());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = {
            ...baseQueryContractsByCodeResponse,
        };
        message.contracts = [];
        if (object.contracts !== undefined && object.contracts !== null) {
            for (const e of object.contracts) {
                message.contracts.push(String(e));
            }
        }
        if (object.pagination !== undefined && object.pagination !== null) {
            message.pagination = PageResponse.fromJSON(object.pagination);
        }
        else {
            message.pagination = undefined;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        if (message.contracts) {
            obj.contracts = message.contracts.map((e) => e);
        }
        else {
            obj.contracts = [];
        }
        message.pagination !== undefined &&
            (obj.pagination = message.pagination
                ? PageResponse.toJSON(message.pagination)
                : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = {
            ...baseQueryContractsByCodeResponse,
        };
        message.contracts = [];
        if (object.contracts !== undefined && object.contracts !== null) {
            for (const e of object.contracts) {
                message.contracts.push(e);
            }
        }
        if (object.pagination !== undefined && object.pagination !== null) {
            message.pagination = PageResponse.fromPartial(object.pagination);
        }
        else {
            message.pagination = undefined;
        }
        return message;
    },
};
const baseQueryAllContractStateRequest = { address: "" };
export const QueryAllContractStateRequest = {
    encode(message, writer = Writer.create()) {
        if (message.address !== "") {
            writer.uint32(10).string(message.address);
        }
        if (message.pagination !== undefined) {
            PageRequest.encode(message.pagination, writer.uint32(18).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = {
            ...baseQueryAllContractStateRequest,
        };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.address = reader.string();
                    break;
                case 2:
                    message.pagination = PageRequest.decode(reader, reader.uint32());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = {
            ...baseQueryAllContractStateRequest,
        };
        if (object.address !== undefined && object.address !== null) {
            message.address = String(object.address);
        }
        else {
            message.address = "";
        }
        if (object.pagination !== undefined && object.pagination !== null) {
            message.pagination = PageRequest.fromJSON(object.pagination);
        }
        else {
            message.pagination = undefined;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.address !== undefined && (obj.address = message.address);
        message.pagination !== undefined &&
            (obj.pagination = message.pagination
                ? PageRequest.toJSON(message.pagination)
                : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = {
            ...baseQueryAllContractStateRequest,
        };
        if (object.address !== undefined && object.address !== null) {
            message.address = object.address;
        }
        else {
            message.address = "";
        }
        if (object.pagination !== undefined && object.pagination !== null) {
            message.pagination = PageRequest.fromPartial(object.pagination);
        }
        else {
            message.pagination = undefined;
        }
        return message;
    },
};
const baseQueryAllContractStateResponse = {};
export const QueryAllContractStateResponse = {
    encode(message, writer = Writer.create()) {
        for (const v of message.models) {
            Model.encode(v, writer.uint32(10).fork()).ldelim();
        }
        if (message.pagination !== undefined) {
            PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = {
            ...baseQueryAllContractStateResponse,
        };
        message.models = [];
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.models.push(Model.decode(reader, reader.uint32()));
                    break;
                case 2:
                    message.pagination = PageResponse.decode(reader, reader.uint32());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = {
            ...baseQueryAllContractStateResponse,
        };
        message.models = [];
        if (object.models !== undefined && object.models !== null) {
            for (const e of object.models) {
                message.models.push(Model.fromJSON(e));
            }
        }
        if (object.pagination !== undefined && object.pagination !== null) {
            message.pagination = PageResponse.fromJSON(object.pagination);
        }
        else {
            message.pagination = undefined;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        if (message.models) {
            obj.models = message.models.map((e) => (e ? Model.toJSON(e) : undefined));
        }
        else {
            obj.models = [];
        }
        message.pagination !== undefined &&
            (obj.pagination = message.pagination
                ? PageResponse.toJSON(message.pagination)
                : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = {
            ...baseQueryAllContractStateResponse,
        };
        message.models = [];
        if (object.models !== undefined && object.models !== null) {
            for (const e of object.models) {
                message.models.push(Model.fromPartial(e));
            }
        }
        if (object.pagination !== undefined && object.pagination !== null) {
            message.pagination = PageResponse.fromPartial(object.pagination);
        }
        else {
            message.pagination = undefined;
        }
        return message;
    },
};
const baseQueryRawContractStateRequest = { address: "" };
export const QueryRawContractStateRequest = {
    encode(message, writer = Writer.create()) {
        if (message.address !== "") {
            writer.uint32(10).string(message.address);
        }
        if (message.query_data.length !== 0) {
            writer.uint32(18).bytes(message.query_data);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = {
            ...baseQueryRawContractStateRequest,
        };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.address = reader.string();
                    break;
                case 2:
                    message.query_data = reader.bytes();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = {
            ...baseQueryRawContractStateRequest,
        };
        if (object.address !== undefined && object.address !== null) {
            message.address = String(object.address);
        }
        else {
            message.address = "";
        }
        if (object.query_data !== undefined && object.query_data !== null) {
            message.query_data = bytesFromBase64(object.query_data);
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.address !== undefined && (obj.address = message.address);
        message.query_data !== undefined &&
            (obj.query_data = base64FromBytes(message.query_data !== undefined ? message.query_data : new Uint8Array()));
        return obj;
    },
    fromPartial(object) {
        const message = {
            ...baseQueryRawContractStateRequest,
        };
        if (object.address !== undefined && object.address !== null) {
            message.address = object.address;
        }
        else {
            message.address = "";
        }
        if (object.query_data !== undefined && object.query_data !== null) {
            message.query_data = object.query_data;
        }
        else {
            message.query_data = new Uint8Array();
        }
        return message;
    },
};
const baseQueryRawContractStateResponse = {};
export const QueryRawContractStateResponse = {
    encode(message, writer = Writer.create()) {
        if (message.data.length !== 0) {
            writer.uint32(10).bytes(message.data);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = {
            ...baseQueryRawContractStateResponse,
        };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.data = reader.bytes();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = {
            ...baseQueryRawContractStateResponse,
        };
        if (object.data !== undefined && object.data !== null) {
            message.data = bytesFromBase64(object.data);
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.data !== undefined &&
            (obj.data = base64FromBytes(message.data !== undefined ? message.data : new Uint8Array()));
        return obj;
    },
    fromPartial(object) {
        const message = {
            ...baseQueryRawContractStateResponse,
        };
        if (object.data !== undefined && object.data !== null) {
            message.data = object.data;
        }
        else {
            message.data = new Uint8Array();
        }
        return message;
    },
};
const baseQuerySmartContractStateRequest = { address: "" };
export const QuerySmartContractStateRequest = {
    encode(message, writer = Writer.create()) {
        if (message.address !== "") {
            writer.uint32(10).string(message.address);
        }
        if (message.query_data.length !== 0) {
            writer.uint32(18).bytes(message.query_data);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = {
            ...baseQuerySmartContractStateRequest,
        };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.address = reader.string();
                    break;
                case 2:
                    message.query_data = reader.bytes();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = {
            ...baseQuerySmartContractStateRequest,
        };
        if (object.address !== undefined && object.address !== null) {
            message.address = String(object.address);
        }
        else {
            message.address = "";
        }
        if (object.query_data !== undefined && object.query_data !== null) {
            message.query_data = bytesFromBase64(object.query_data);
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.address !== undefined && (obj.address = message.address);
        message.query_data !== undefined &&
            (obj.query_data = base64FromBytes(message.query_data !== undefined ? message.query_data : new Uint8Array()));
        return obj;
    },
    fromPartial(object) {
        const message = {
            ...baseQuerySmartContractStateRequest,
        };
        if (object.address !== undefined && object.address !== null) {
            message.address = object.address;
        }
        else {
            message.address = "";
        }
        if (object.query_data !== undefined && object.query_data !== null) {
            message.query_data = object.query_data;
        }
        else {
            message.query_data = new Uint8Array();
        }
        return message;
    },
};
const baseQuerySmartContractStateResponse = {};
export const QuerySmartContractStateResponse = {
    encode(message, writer = Writer.create()) {
        if (message.data.length !== 0) {
            writer.uint32(10).bytes(message.data);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = {
            ...baseQuerySmartContractStateResponse,
        };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.data = reader.bytes();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = {
            ...baseQuerySmartContractStateResponse,
        };
        if (object.data !== undefined && object.data !== null) {
            message.data = bytesFromBase64(object.data);
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.data !== undefined &&
            (obj.data = base64FromBytes(message.data !== undefined ? message.data : new Uint8Array()));
        return obj;
    },
    fromPartial(object) {
        const message = {
            ...baseQuerySmartContractStateResponse,
        };
        if (object.data !== undefined && object.data !== null) {
            message.data = object.data;
        }
        else {
            message.data = new Uint8Array();
        }
        return message;
    },
};
const baseQueryCodeRequest = { code_id: 0 };
export const QueryCodeRequest = {
    encode(message, writer = Writer.create()) {
        if (message.code_id !== 0) {
            writer.uint32(8).uint64(message.code_id);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryCodeRequest };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.code_id = longToNumber(reader.uint64());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseQueryCodeRequest };
        if (object.code_id !== undefined && object.code_id !== null) {
            message.code_id = Number(object.code_id);
        }
        else {
            message.code_id = 0;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.code_id !== undefined && (obj.code_id = message.code_id);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryCodeRequest };
        if (object.code_id !== undefined && object.code_id !== null) {
            message.code_id = object.code_id;
        }
        else {
            message.code_id = 0;
        }
        return message;
    },
};
const baseCodeInfoResponse = { code_id: 0, creator: "" };
export const CodeInfoResponse = {
    encode(message, writer = Writer.create()) {
        if (message.code_id !== 0) {
            writer.uint32(8).uint64(message.code_id);
        }
        if (message.creator !== "") {
            writer.uint32(18).string(message.creator);
        }
        if (message.data_hash.length !== 0) {
            writer.uint32(26).bytes(message.data_hash);
        }
        if (message.instantiate_permission !== undefined) {
            AccessConfig.encode(message.instantiate_permission, writer.uint32(50).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseCodeInfoResponse };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.code_id = longToNumber(reader.uint64());
                    break;
                case 2:
                    message.creator = reader.string();
                    break;
                case 3:
                    message.data_hash = reader.bytes();
                    break;
                case 6:
                    message.instantiate_permission = AccessConfig.decode(reader, reader.uint32());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseCodeInfoResponse };
        if (object.code_id !== undefined && object.code_id !== null) {
            message.code_id = Number(object.code_id);
        }
        else {
            message.code_id = 0;
        }
        if (object.creator !== undefined && object.creator !== null) {
            message.creator = String(object.creator);
        }
        else {
            message.creator = "";
        }
        if (object.data_hash !== undefined && object.data_hash !== null) {
            message.data_hash = bytesFromBase64(object.data_hash);
        }
        if (object.instantiate_permission !== undefined &&
            object.instantiate_permission !== null) {
            message.instantiate_permission = AccessConfig.fromJSON(object.instantiate_permission);
        }
        else {
            message.instantiate_permission = undefined;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.code_id !== undefined && (obj.code_id = message.code_id);
        message.creator !== undefined && (obj.creator = message.creator);
        message.data_hash !== undefined &&
            (obj.data_hash = base64FromBytes(message.data_hash !== undefined ? message.data_hash : new Uint8Array()));
        message.instantiate_permission !== undefined &&
            (obj.instantiate_permission = message.instantiate_permission
                ? AccessConfig.toJSON(message.instantiate_permission)
                : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseCodeInfoResponse };
        if (object.code_id !== undefined && object.code_id !== null) {
            message.code_id = object.code_id;
        }
        else {
            message.code_id = 0;
        }
        if (object.creator !== undefined && object.creator !== null) {
            message.creator = object.creator;
        }
        else {
            message.creator = "";
        }
        if (object.data_hash !== undefined && object.data_hash !== null) {
            message.data_hash = object.data_hash;
        }
        else {
            message.data_hash = new Uint8Array();
        }
        if (object.instantiate_permission !== undefined &&
            object.instantiate_permission !== null) {
            message.instantiate_permission = AccessConfig.fromPartial(object.instantiate_permission);
        }
        else {
            message.instantiate_permission = undefined;
        }
        return message;
    },
};
const baseQueryCodeResponse = {};
export const QueryCodeResponse = {
    encode(message, writer = Writer.create()) {
        if (message.code_info !== undefined) {
            CodeInfoResponse.encode(message.code_info, writer.uint32(10).fork()).ldelim();
        }
        if (message.data.length !== 0) {
            writer.uint32(18).bytes(message.data);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryCodeResponse };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.code_info = CodeInfoResponse.decode(reader, reader.uint32());
                    break;
                case 2:
                    message.data = reader.bytes();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseQueryCodeResponse };
        if (object.code_info !== undefined && object.code_info !== null) {
            message.code_info = CodeInfoResponse.fromJSON(object.code_info);
        }
        else {
            message.code_info = undefined;
        }
        if (object.data !== undefined && object.data !== null) {
            message.data = bytesFromBase64(object.data);
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.code_info !== undefined &&
            (obj.code_info = message.code_info
                ? CodeInfoResponse.toJSON(message.code_info)
                : undefined);
        message.data !== undefined &&
            (obj.data = base64FromBytes(message.data !== undefined ? message.data : new Uint8Array()));
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryCodeResponse };
        if (object.code_info !== undefined && object.code_info !== null) {
            message.code_info = CodeInfoResponse.fromPartial(object.code_info);
        }
        else {
            message.code_info = undefined;
        }
        if (object.data !== undefined && object.data !== null) {
            message.data = object.data;
        }
        else {
            message.data = new Uint8Array();
        }
        return message;
    },
};
const baseQueryCodesRequest = {};
export const QueryCodesRequest = {
    encode(message, writer = Writer.create()) {
        if (message.pagination !== undefined) {
            PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryCodesRequest };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.pagination = PageRequest.decode(reader, reader.uint32());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseQueryCodesRequest };
        if (object.pagination !== undefined && object.pagination !== null) {
            message.pagination = PageRequest.fromJSON(object.pagination);
        }
        else {
            message.pagination = undefined;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.pagination !== undefined &&
            (obj.pagination = message.pagination
                ? PageRequest.toJSON(message.pagination)
                : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryCodesRequest };
        if (object.pagination !== undefined && object.pagination !== null) {
            message.pagination = PageRequest.fromPartial(object.pagination);
        }
        else {
            message.pagination = undefined;
        }
        return message;
    },
};
const baseQueryCodesResponse = {};
export const QueryCodesResponse = {
    encode(message, writer = Writer.create()) {
        for (const v of message.code_infos) {
            CodeInfoResponse.encode(v, writer.uint32(10).fork()).ldelim();
        }
        if (message.pagination !== undefined) {
            PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryCodesResponse };
        message.code_infos = [];
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.code_infos.push(CodeInfoResponse.decode(reader, reader.uint32()));
                    break;
                case 2:
                    message.pagination = PageResponse.decode(reader, reader.uint32());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseQueryCodesResponse };
        message.code_infos = [];
        if (object.code_infos !== undefined && object.code_infos !== null) {
            for (const e of object.code_infos) {
                message.code_infos.push(CodeInfoResponse.fromJSON(e));
            }
        }
        if (object.pagination !== undefined && object.pagination !== null) {
            message.pagination = PageResponse.fromJSON(object.pagination);
        }
        else {
            message.pagination = undefined;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        if (message.code_infos) {
            obj.code_infos = message.code_infos.map((e) => e ? CodeInfoResponse.toJSON(e) : undefined);
        }
        else {
            obj.code_infos = [];
        }
        message.pagination !== undefined &&
            (obj.pagination = message.pagination
                ? PageResponse.toJSON(message.pagination)
                : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryCodesResponse };
        message.code_infos = [];
        if (object.code_infos !== undefined && object.code_infos !== null) {
            for (const e of object.code_infos) {
                message.code_infos.push(CodeInfoResponse.fromPartial(e));
            }
        }
        if (object.pagination !== undefined && object.pagination !== null) {
            message.pagination = PageResponse.fromPartial(object.pagination);
        }
        else {
            message.pagination = undefined;
        }
        return message;
    },
};
const baseQueryPinnedCodesRequest = {};
export const QueryPinnedCodesRequest = {
    encode(message, writer = Writer.create()) {
        if (message.pagination !== undefined) {
            PageRequest.encode(message.pagination, writer.uint32(18).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = {
            ...baseQueryPinnedCodesRequest,
        };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 2:
                    message.pagination = PageRequest.decode(reader, reader.uint32());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = {
            ...baseQueryPinnedCodesRequest,
        };
        if (object.pagination !== undefined && object.pagination !== null) {
            message.pagination = PageRequest.fromJSON(object.pagination);
        }
        else {
            message.pagination = undefined;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.pagination !== undefined &&
            (obj.pagination = message.pagination
                ? PageRequest.toJSON(message.pagination)
                : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = {
            ...baseQueryPinnedCodesRequest,
        };
        if (object.pagination !== undefined && object.pagination !== null) {
            message.pagination = PageRequest.fromPartial(object.pagination);
        }
        else {
            message.pagination = undefined;
        }
        return message;
    },
};
const baseQueryPinnedCodesResponse = { code_ids: 0 };
export const QueryPinnedCodesResponse = {
    encode(message, writer = Writer.create()) {
        writer.uint32(10).fork();
        for (const v of message.code_ids) {
            writer.uint64(v);
        }
        writer.ldelim();
        if (message.pagination !== undefined) {
            PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = {
            ...baseQueryPinnedCodesResponse,
        };
        message.code_ids = [];
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    if ((tag & 7) === 2) {
                        const end2 = reader.uint32() + reader.pos;
                        while (reader.pos < end2) {
                            message.code_ids.push(longToNumber(reader.uint64()));
                        }
                    }
                    else {
                        message.code_ids.push(longToNumber(reader.uint64()));
                    }
                    break;
                case 2:
                    message.pagination = PageResponse.decode(reader, reader.uint32());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = {
            ...baseQueryPinnedCodesResponse,
        };
        message.code_ids = [];
        if (object.code_ids !== undefined && object.code_ids !== null) {
            for (const e of object.code_ids) {
                message.code_ids.push(Number(e));
            }
        }
        if (object.pagination !== undefined && object.pagination !== null) {
            message.pagination = PageResponse.fromJSON(object.pagination);
        }
        else {
            message.pagination = undefined;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        if (message.code_ids) {
            obj.code_ids = message.code_ids.map((e) => e);
        }
        else {
            obj.code_ids = [];
        }
        message.pagination !== undefined &&
            (obj.pagination = message.pagination
                ? PageResponse.toJSON(message.pagination)
                : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = {
            ...baseQueryPinnedCodesResponse,
        };
        message.code_ids = [];
        if (object.code_ids !== undefined && object.code_ids !== null) {
            for (const e of object.code_ids) {
                message.code_ids.push(e);
            }
        }
        if (object.pagination !== undefined && object.pagination !== null) {
            message.pagination = PageResponse.fromPartial(object.pagination);
        }
        else {
            message.pagination = undefined;
        }
        return message;
    },
};
export class QueryClientImpl {
    constructor(rpc) {
        this.rpc = rpc;
    }
    ContractInfo(request) {
        const data = QueryContractInfoRequest.encode(request).finish();
        const promise = this.rpc.request("cosmwasm.wasm.v1.Query", "ContractInfo", data);
        return promise.then((data) => QueryContractInfoResponse.decode(new Reader(data)));
    }
    ContractHistory(request) {
        const data = QueryContractHistoryRequest.encode(request).finish();
        const promise = this.rpc.request("cosmwasm.wasm.v1.Query", "ContractHistory", data);
        return promise.then((data) => QueryContractHistoryResponse.decode(new Reader(data)));
    }
    ContractsByCode(request) {
        const data = QueryContractsByCodeRequest.encode(request).finish();
        const promise = this.rpc.request("cosmwasm.wasm.v1.Query", "ContractsByCode", data);
        return promise.then((data) => QueryContractsByCodeResponse.decode(new Reader(data)));
    }
    AllContractState(request) {
        const data = QueryAllContractStateRequest.encode(request).finish();
        const promise = this.rpc.request("cosmwasm.wasm.v1.Query", "AllContractState", data);
        return promise.then((data) => QueryAllContractStateResponse.decode(new Reader(data)));
    }
    RawContractState(request) {
        const data = QueryRawContractStateRequest.encode(request).finish();
        const promise = this.rpc.request("cosmwasm.wasm.v1.Query", "RawContractState", data);
        return promise.then((data) => QueryRawContractStateResponse.decode(new Reader(data)));
    }
    SmartContractState(request) {
        const data = QuerySmartContractStateRequest.encode(request).finish();
        const promise = this.rpc.request("cosmwasm.wasm.v1.Query", "SmartContractState", data);
        return promise.then((data) => QuerySmartContractStateResponse.decode(new Reader(data)));
    }
    Code(request) {
        const data = QueryCodeRequest.encode(request).finish();
        const promise = this.rpc.request("cosmwasm.wasm.v1.Query", "Code", data);
        return promise.then((data) => QueryCodeResponse.decode(new Reader(data)));
    }
    Codes(request) {
        const data = QueryCodesRequest.encode(request).finish();
        const promise = this.rpc.request("cosmwasm.wasm.v1.Query", "Codes", data);
        return promise.then((data) => QueryCodesResponse.decode(new Reader(data)));
    }
    PinnedCodes(request) {
        const data = QueryPinnedCodesRequest.encode(request).finish();
        const promise = this.rpc.request("cosmwasm.wasm.v1.Query", "PinnedCodes", data);
        return promise.then((data) => QueryPinnedCodesResponse.decode(new Reader(data)));
    }
}
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
