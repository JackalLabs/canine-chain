/* eslint-disable */
import * as Long from "long";
import { util, configure, Writer, Reader } from "protobufjs/minimal";
import { AccessConfig } from "../../../cosmwasm/wasm/v1/types";
import { Coin } from "../../../cosmos/base/v1beta1/coin";
export const protobufPackage = "cosmwasm.wasm.v1";
const baseStoreCodeProposal = {
    title: "",
    description: "",
    run_as: "",
};
export const StoreCodeProposal = {
    encode(message, writer = Writer.create()) {
        if (message.title !== "") {
            writer.uint32(10).string(message.title);
        }
        if (message.description !== "") {
            writer.uint32(18).string(message.description);
        }
        if (message.run_as !== "") {
            writer.uint32(26).string(message.run_as);
        }
        if (message.wasm_byte_code.length !== 0) {
            writer.uint32(34).bytes(message.wasm_byte_code);
        }
        if (message.instantiate_permission !== undefined) {
            AccessConfig.encode(message.instantiate_permission, writer.uint32(58).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseStoreCodeProposal };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.title = reader.string();
                    break;
                case 2:
                    message.description = reader.string();
                    break;
                case 3:
                    message.run_as = reader.string();
                    break;
                case 4:
                    message.wasm_byte_code = reader.bytes();
                    break;
                case 7:
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
        const message = { ...baseStoreCodeProposal };
        if (object.title !== undefined && object.title !== null) {
            message.title = String(object.title);
        }
        else {
            message.title = "";
        }
        if (object.description !== undefined && object.description !== null) {
            message.description = String(object.description);
        }
        else {
            message.description = "";
        }
        if (object.run_as !== undefined && object.run_as !== null) {
            message.run_as = String(object.run_as);
        }
        else {
            message.run_as = "";
        }
        if (object.wasm_byte_code !== undefined && object.wasm_byte_code !== null) {
            message.wasm_byte_code = bytesFromBase64(object.wasm_byte_code);
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
        message.title !== undefined && (obj.title = message.title);
        message.description !== undefined &&
            (obj.description = message.description);
        message.run_as !== undefined && (obj.run_as = message.run_as);
        message.wasm_byte_code !== undefined &&
            (obj.wasm_byte_code = base64FromBytes(message.wasm_byte_code !== undefined
                ? message.wasm_byte_code
                : new Uint8Array()));
        message.instantiate_permission !== undefined &&
            (obj.instantiate_permission = message.instantiate_permission
                ? AccessConfig.toJSON(message.instantiate_permission)
                : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseStoreCodeProposal };
        if (object.title !== undefined && object.title !== null) {
            message.title = object.title;
        }
        else {
            message.title = "";
        }
        if (object.description !== undefined && object.description !== null) {
            message.description = object.description;
        }
        else {
            message.description = "";
        }
        if (object.run_as !== undefined && object.run_as !== null) {
            message.run_as = object.run_as;
        }
        else {
            message.run_as = "";
        }
        if (object.wasm_byte_code !== undefined && object.wasm_byte_code !== null) {
            message.wasm_byte_code = object.wasm_byte_code;
        }
        else {
            message.wasm_byte_code = new Uint8Array();
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
const baseInstantiateContractProposal = {
    title: "",
    description: "",
    run_as: "",
    admin: "",
    code_id: 0,
    label: "",
};
export const InstantiateContractProposal = {
    encode(message, writer = Writer.create()) {
        if (message.title !== "") {
            writer.uint32(10).string(message.title);
        }
        if (message.description !== "") {
            writer.uint32(18).string(message.description);
        }
        if (message.run_as !== "") {
            writer.uint32(26).string(message.run_as);
        }
        if (message.admin !== "") {
            writer.uint32(34).string(message.admin);
        }
        if (message.code_id !== 0) {
            writer.uint32(40).uint64(message.code_id);
        }
        if (message.label !== "") {
            writer.uint32(50).string(message.label);
        }
        if (message.msg.length !== 0) {
            writer.uint32(58).bytes(message.msg);
        }
        for (const v of message.funds) {
            Coin.encode(v, writer.uint32(66).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = {
            ...baseInstantiateContractProposal,
        };
        message.funds = [];
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.title = reader.string();
                    break;
                case 2:
                    message.description = reader.string();
                    break;
                case 3:
                    message.run_as = reader.string();
                    break;
                case 4:
                    message.admin = reader.string();
                    break;
                case 5:
                    message.code_id = longToNumber(reader.uint64());
                    break;
                case 6:
                    message.label = reader.string();
                    break;
                case 7:
                    message.msg = reader.bytes();
                    break;
                case 8:
                    message.funds.push(Coin.decode(reader, reader.uint32()));
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
            ...baseInstantiateContractProposal,
        };
        message.funds = [];
        if (object.title !== undefined && object.title !== null) {
            message.title = String(object.title);
        }
        else {
            message.title = "";
        }
        if (object.description !== undefined && object.description !== null) {
            message.description = String(object.description);
        }
        else {
            message.description = "";
        }
        if (object.run_as !== undefined && object.run_as !== null) {
            message.run_as = String(object.run_as);
        }
        else {
            message.run_as = "";
        }
        if (object.admin !== undefined && object.admin !== null) {
            message.admin = String(object.admin);
        }
        else {
            message.admin = "";
        }
        if (object.code_id !== undefined && object.code_id !== null) {
            message.code_id = Number(object.code_id);
        }
        else {
            message.code_id = 0;
        }
        if (object.label !== undefined && object.label !== null) {
            message.label = String(object.label);
        }
        else {
            message.label = "";
        }
        if (object.msg !== undefined && object.msg !== null) {
            message.msg = bytesFromBase64(object.msg);
        }
        if (object.funds !== undefined && object.funds !== null) {
            for (const e of object.funds) {
                message.funds.push(Coin.fromJSON(e));
            }
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.title !== undefined && (obj.title = message.title);
        message.description !== undefined &&
            (obj.description = message.description);
        message.run_as !== undefined && (obj.run_as = message.run_as);
        message.admin !== undefined && (obj.admin = message.admin);
        message.code_id !== undefined && (obj.code_id = message.code_id);
        message.label !== undefined && (obj.label = message.label);
        message.msg !== undefined &&
            (obj.msg = base64FromBytes(message.msg !== undefined ? message.msg : new Uint8Array()));
        if (message.funds) {
            obj.funds = message.funds.map((e) => (e ? Coin.toJSON(e) : undefined));
        }
        else {
            obj.funds = [];
        }
        return obj;
    },
    fromPartial(object) {
        const message = {
            ...baseInstantiateContractProposal,
        };
        message.funds = [];
        if (object.title !== undefined && object.title !== null) {
            message.title = object.title;
        }
        else {
            message.title = "";
        }
        if (object.description !== undefined && object.description !== null) {
            message.description = object.description;
        }
        else {
            message.description = "";
        }
        if (object.run_as !== undefined && object.run_as !== null) {
            message.run_as = object.run_as;
        }
        else {
            message.run_as = "";
        }
        if (object.admin !== undefined && object.admin !== null) {
            message.admin = object.admin;
        }
        else {
            message.admin = "";
        }
        if (object.code_id !== undefined && object.code_id !== null) {
            message.code_id = object.code_id;
        }
        else {
            message.code_id = 0;
        }
        if (object.label !== undefined && object.label !== null) {
            message.label = object.label;
        }
        else {
            message.label = "";
        }
        if (object.msg !== undefined && object.msg !== null) {
            message.msg = object.msg;
        }
        else {
            message.msg = new Uint8Array();
        }
        if (object.funds !== undefined && object.funds !== null) {
            for (const e of object.funds) {
                message.funds.push(Coin.fromPartial(e));
            }
        }
        return message;
    },
};
const baseMigrateContractProposal = {
    title: "",
    description: "",
    contract: "",
    code_id: 0,
};
export const MigrateContractProposal = {
    encode(message, writer = Writer.create()) {
        if (message.title !== "") {
            writer.uint32(10).string(message.title);
        }
        if (message.description !== "") {
            writer.uint32(18).string(message.description);
        }
        if (message.contract !== "") {
            writer.uint32(34).string(message.contract);
        }
        if (message.code_id !== 0) {
            writer.uint32(40).uint64(message.code_id);
        }
        if (message.msg.length !== 0) {
            writer.uint32(50).bytes(message.msg);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = {
            ...baseMigrateContractProposal,
        };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.title = reader.string();
                    break;
                case 2:
                    message.description = reader.string();
                    break;
                case 4:
                    message.contract = reader.string();
                    break;
                case 5:
                    message.code_id = longToNumber(reader.uint64());
                    break;
                case 6:
                    message.msg = reader.bytes();
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
            ...baseMigrateContractProposal,
        };
        if (object.title !== undefined && object.title !== null) {
            message.title = String(object.title);
        }
        else {
            message.title = "";
        }
        if (object.description !== undefined && object.description !== null) {
            message.description = String(object.description);
        }
        else {
            message.description = "";
        }
        if (object.contract !== undefined && object.contract !== null) {
            message.contract = String(object.contract);
        }
        else {
            message.contract = "";
        }
        if (object.code_id !== undefined && object.code_id !== null) {
            message.code_id = Number(object.code_id);
        }
        else {
            message.code_id = 0;
        }
        if (object.msg !== undefined && object.msg !== null) {
            message.msg = bytesFromBase64(object.msg);
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.title !== undefined && (obj.title = message.title);
        message.description !== undefined &&
            (obj.description = message.description);
        message.contract !== undefined && (obj.contract = message.contract);
        message.code_id !== undefined && (obj.code_id = message.code_id);
        message.msg !== undefined &&
            (obj.msg = base64FromBytes(message.msg !== undefined ? message.msg : new Uint8Array()));
        return obj;
    },
    fromPartial(object) {
        const message = {
            ...baseMigrateContractProposal,
        };
        if (object.title !== undefined && object.title !== null) {
            message.title = object.title;
        }
        else {
            message.title = "";
        }
        if (object.description !== undefined && object.description !== null) {
            message.description = object.description;
        }
        else {
            message.description = "";
        }
        if (object.contract !== undefined && object.contract !== null) {
            message.contract = object.contract;
        }
        else {
            message.contract = "";
        }
        if (object.code_id !== undefined && object.code_id !== null) {
            message.code_id = object.code_id;
        }
        else {
            message.code_id = 0;
        }
        if (object.msg !== undefined && object.msg !== null) {
            message.msg = object.msg;
        }
        else {
            message.msg = new Uint8Array();
        }
        return message;
    },
};
const baseSudoContractProposal = {
    title: "",
    description: "",
    contract: "",
};
export const SudoContractProposal = {
    encode(message, writer = Writer.create()) {
        if (message.title !== "") {
            writer.uint32(10).string(message.title);
        }
        if (message.description !== "") {
            writer.uint32(18).string(message.description);
        }
        if (message.contract !== "") {
            writer.uint32(26).string(message.contract);
        }
        if (message.msg.length !== 0) {
            writer.uint32(34).bytes(message.msg);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseSudoContractProposal };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.title = reader.string();
                    break;
                case 2:
                    message.description = reader.string();
                    break;
                case 3:
                    message.contract = reader.string();
                    break;
                case 4:
                    message.msg = reader.bytes();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseSudoContractProposal };
        if (object.title !== undefined && object.title !== null) {
            message.title = String(object.title);
        }
        else {
            message.title = "";
        }
        if (object.description !== undefined && object.description !== null) {
            message.description = String(object.description);
        }
        else {
            message.description = "";
        }
        if (object.contract !== undefined && object.contract !== null) {
            message.contract = String(object.contract);
        }
        else {
            message.contract = "";
        }
        if (object.msg !== undefined && object.msg !== null) {
            message.msg = bytesFromBase64(object.msg);
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.title !== undefined && (obj.title = message.title);
        message.description !== undefined &&
            (obj.description = message.description);
        message.contract !== undefined && (obj.contract = message.contract);
        message.msg !== undefined &&
            (obj.msg = base64FromBytes(message.msg !== undefined ? message.msg : new Uint8Array()));
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseSudoContractProposal };
        if (object.title !== undefined && object.title !== null) {
            message.title = object.title;
        }
        else {
            message.title = "";
        }
        if (object.description !== undefined && object.description !== null) {
            message.description = object.description;
        }
        else {
            message.description = "";
        }
        if (object.contract !== undefined && object.contract !== null) {
            message.contract = object.contract;
        }
        else {
            message.contract = "";
        }
        if (object.msg !== undefined && object.msg !== null) {
            message.msg = object.msg;
        }
        else {
            message.msg = new Uint8Array();
        }
        return message;
    },
};
const baseExecuteContractProposal = {
    title: "",
    description: "",
    run_as: "",
    contract: "",
};
export const ExecuteContractProposal = {
    encode(message, writer = Writer.create()) {
        if (message.title !== "") {
            writer.uint32(10).string(message.title);
        }
        if (message.description !== "") {
            writer.uint32(18).string(message.description);
        }
        if (message.run_as !== "") {
            writer.uint32(26).string(message.run_as);
        }
        if (message.contract !== "") {
            writer.uint32(34).string(message.contract);
        }
        if (message.msg.length !== 0) {
            writer.uint32(42).bytes(message.msg);
        }
        for (const v of message.funds) {
            Coin.encode(v, writer.uint32(50).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = {
            ...baseExecuteContractProposal,
        };
        message.funds = [];
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.title = reader.string();
                    break;
                case 2:
                    message.description = reader.string();
                    break;
                case 3:
                    message.run_as = reader.string();
                    break;
                case 4:
                    message.contract = reader.string();
                    break;
                case 5:
                    message.msg = reader.bytes();
                    break;
                case 6:
                    message.funds.push(Coin.decode(reader, reader.uint32()));
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
            ...baseExecuteContractProposal,
        };
        message.funds = [];
        if (object.title !== undefined && object.title !== null) {
            message.title = String(object.title);
        }
        else {
            message.title = "";
        }
        if (object.description !== undefined && object.description !== null) {
            message.description = String(object.description);
        }
        else {
            message.description = "";
        }
        if (object.run_as !== undefined && object.run_as !== null) {
            message.run_as = String(object.run_as);
        }
        else {
            message.run_as = "";
        }
        if (object.contract !== undefined && object.contract !== null) {
            message.contract = String(object.contract);
        }
        else {
            message.contract = "";
        }
        if (object.msg !== undefined && object.msg !== null) {
            message.msg = bytesFromBase64(object.msg);
        }
        if (object.funds !== undefined && object.funds !== null) {
            for (const e of object.funds) {
                message.funds.push(Coin.fromJSON(e));
            }
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.title !== undefined && (obj.title = message.title);
        message.description !== undefined &&
            (obj.description = message.description);
        message.run_as !== undefined && (obj.run_as = message.run_as);
        message.contract !== undefined && (obj.contract = message.contract);
        message.msg !== undefined &&
            (obj.msg = base64FromBytes(message.msg !== undefined ? message.msg : new Uint8Array()));
        if (message.funds) {
            obj.funds = message.funds.map((e) => (e ? Coin.toJSON(e) : undefined));
        }
        else {
            obj.funds = [];
        }
        return obj;
    },
    fromPartial(object) {
        const message = {
            ...baseExecuteContractProposal,
        };
        message.funds = [];
        if (object.title !== undefined && object.title !== null) {
            message.title = object.title;
        }
        else {
            message.title = "";
        }
        if (object.description !== undefined && object.description !== null) {
            message.description = object.description;
        }
        else {
            message.description = "";
        }
        if (object.run_as !== undefined && object.run_as !== null) {
            message.run_as = object.run_as;
        }
        else {
            message.run_as = "";
        }
        if (object.contract !== undefined && object.contract !== null) {
            message.contract = object.contract;
        }
        else {
            message.contract = "";
        }
        if (object.msg !== undefined && object.msg !== null) {
            message.msg = object.msg;
        }
        else {
            message.msg = new Uint8Array();
        }
        if (object.funds !== undefined && object.funds !== null) {
            for (const e of object.funds) {
                message.funds.push(Coin.fromPartial(e));
            }
        }
        return message;
    },
};
const baseUpdateAdminProposal = {
    title: "",
    description: "",
    new_admin: "",
    contract: "",
};
export const UpdateAdminProposal = {
    encode(message, writer = Writer.create()) {
        if (message.title !== "") {
            writer.uint32(10).string(message.title);
        }
        if (message.description !== "") {
            writer.uint32(18).string(message.description);
        }
        if (message.new_admin !== "") {
            writer.uint32(26).string(message.new_admin);
        }
        if (message.contract !== "") {
            writer.uint32(34).string(message.contract);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseUpdateAdminProposal };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.title = reader.string();
                    break;
                case 2:
                    message.description = reader.string();
                    break;
                case 3:
                    message.new_admin = reader.string();
                    break;
                case 4:
                    message.contract = reader.string();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseUpdateAdminProposal };
        if (object.title !== undefined && object.title !== null) {
            message.title = String(object.title);
        }
        else {
            message.title = "";
        }
        if (object.description !== undefined && object.description !== null) {
            message.description = String(object.description);
        }
        else {
            message.description = "";
        }
        if (object.new_admin !== undefined && object.new_admin !== null) {
            message.new_admin = String(object.new_admin);
        }
        else {
            message.new_admin = "";
        }
        if (object.contract !== undefined && object.contract !== null) {
            message.contract = String(object.contract);
        }
        else {
            message.contract = "";
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.title !== undefined && (obj.title = message.title);
        message.description !== undefined &&
            (obj.description = message.description);
        message.new_admin !== undefined && (obj.new_admin = message.new_admin);
        message.contract !== undefined && (obj.contract = message.contract);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseUpdateAdminProposal };
        if (object.title !== undefined && object.title !== null) {
            message.title = object.title;
        }
        else {
            message.title = "";
        }
        if (object.description !== undefined && object.description !== null) {
            message.description = object.description;
        }
        else {
            message.description = "";
        }
        if (object.new_admin !== undefined && object.new_admin !== null) {
            message.new_admin = object.new_admin;
        }
        else {
            message.new_admin = "";
        }
        if (object.contract !== undefined && object.contract !== null) {
            message.contract = object.contract;
        }
        else {
            message.contract = "";
        }
        return message;
    },
};
const baseClearAdminProposal = {
    title: "",
    description: "",
    contract: "",
};
export const ClearAdminProposal = {
    encode(message, writer = Writer.create()) {
        if (message.title !== "") {
            writer.uint32(10).string(message.title);
        }
        if (message.description !== "") {
            writer.uint32(18).string(message.description);
        }
        if (message.contract !== "") {
            writer.uint32(26).string(message.contract);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseClearAdminProposal };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.title = reader.string();
                    break;
                case 2:
                    message.description = reader.string();
                    break;
                case 3:
                    message.contract = reader.string();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseClearAdminProposal };
        if (object.title !== undefined && object.title !== null) {
            message.title = String(object.title);
        }
        else {
            message.title = "";
        }
        if (object.description !== undefined && object.description !== null) {
            message.description = String(object.description);
        }
        else {
            message.description = "";
        }
        if (object.contract !== undefined && object.contract !== null) {
            message.contract = String(object.contract);
        }
        else {
            message.contract = "";
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.title !== undefined && (obj.title = message.title);
        message.description !== undefined &&
            (obj.description = message.description);
        message.contract !== undefined && (obj.contract = message.contract);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseClearAdminProposal };
        if (object.title !== undefined && object.title !== null) {
            message.title = object.title;
        }
        else {
            message.title = "";
        }
        if (object.description !== undefined && object.description !== null) {
            message.description = object.description;
        }
        else {
            message.description = "";
        }
        if (object.contract !== undefined && object.contract !== null) {
            message.contract = object.contract;
        }
        else {
            message.contract = "";
        }
        return message;
    },
};
const basePinCodesProposal = {
    title: "",
    description: "",
    code_ids: 0,
};
export const PinCodesProposal = {
    encode(message, writer = Writer.create()) {
        if (message.title !== "") {
            writer.uint32(10).string(message.title);
        }
        if (message.description !== "") {
            writer.uint32(18).string(message.description);
        }
        writer.uint32(26).fork();
        for (const v of message.code_ids) {
            writer.uint64(v);
        }
        writer.ldelim();
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...basePinCodesProposal };
        message.code_ids = [];
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.title = reader.string();
                    break;
                case 2:
                    message.description = reader.string();
                    break;
                case 3:
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
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...basePinCodesProposal };
        message.code_ids = [];
        if (object.title !== undefined && object.title !== null) {
            message.title = String(object.title);
        }
        else {
            message.title = "";
        }
        if (object.description !== undefined && object.description !== null) {
            message.description = String(object.description);
        }
        else {
            message.description = "";
        }
        if (object.code_ids !== undefined && object.code_ids !== null) {
            for (const e of object.code_ids) {
                message.code_ids.push(Number(e));
            }
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.title !== undefined && (obj.title = message.title);
        message.description !== undefined &&
            (obj.description = message.description);
        if (message.code_ids) {
            obj.code_ids = message.code_ids.map((e) => e);
        }
        else {
            obj.code_ids = [];
        }
        return obj;
    },
    fromPartial(object) {
        const message = { ...basePinCodesProposal };
        message.code_ids = [];
        if (object.title !== undefined && object.title !== null) {
            message.title = object.title;
        }
        else {
            message.title = "";
        }
        if (object.description !== undefined && object.description !== null) {
            message.description = object.description;
        }
        else {
            message.description = "";
        }
        if (object.code_ids !== undefined && object.code_ids !== null) {
            for (const e of object.code_ids) {
                message.code_ids.push(e);
            }
        }
        return message;
    },
};
const baseUnpinCodesProposal = {
    title: "",
    description: "",
    code_ids: 0,
};
export const UnpinCodesProposal = {
    encode(message, writer = Writer.create()) {
        if (message.title !== "") {
            writer.uint32(10).string(message.title);
        }
        if (message.description !== "") {
            writer.uint32(18).string(message.description);
        }
        writer.uint32(26).fork();
        for (const v of message.code_ids) {
            writer.uint64(v);
        }
        writer.ldelim();
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseUnpinCodesProposal };
        message.code_ids = [];
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.title = reader.string();
                    break;
                case 2:
                    message.description = reader.string();
                    break;
                case 3:
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
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseUnpinCodesProposal };
        message.code_ids = [];
        if (object.title !== undefined && object.title !== null) {
            message.title = String(object.title);
        }
        else {
            message.title = "";
        }
        if (object.description !== undefined && object.description !== null) {
            message.description = String(object.description);
        }
        else {
            message.description = "";
        }
        if (object.code_ids !== undefined && object.code_ids !== null) {
            for (const e of object.code_ids) {
                message.code_ids.push(Number(e));
            }
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.title !== undefined && (obj.title = message.title);
        message.description !== undefined &&
            (obj.description = message.description);
        if (message.code_ids) {
            obj.code_ids = message.code_ids.map((e) => e);
        }
        else {
            obj.code_ids = [];
        }
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseUnpinCodesProposal };
        message.code_ids = [];
        if (object.title !== undefined && object.title !== null) {
            message.title = object.title;
        }
        else {
            message.title = "";
        }
        if (object.description !== undefined && object.description !== null) {
            message.description = object.description;
        }
        else {
            message.description = "";
        }
        if (object.code_ids !== undefined && object.code_ids !== null) {
            for (const e of object.code_ids) {
                message.code_ids.push(e);
            }
        }
        return message;
    },
};
const baseAccessConfigUpdate = { code_id: 0 };
export const AccessConfigUpdate = {
    encode(message, writer = Writer.create()) {
        if (message.code_id !== 0) {
            writer.uint32(8).uint64(message.code_id);
        }
        if (message.instantiate_permission !== undefined) {
            AccessConfig.encode(message.instantiate_permission, writer.uint32(18).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseAccessConfigUpdate };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.code_id = longToNumber(reader.uint64());
                    break;
                case 2:
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
        const message = { ...baseAccessConfigUpdate };
        if (object.code_id !== undefined && object.code_id !== null) {
            message.code_id = Number(object.code_id);
        }
        else {
            message.code_id = 0;
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
        message.instantiate_permission !== undefined &&
            (obj.instantiate_permission = message.instantiate_permission
                ? AccessConfig.toJSON(message.instantiate_permission)
                : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseAccessConfigUpdate };
        if (object.code_id !== undefined && object.code_id !== null) {
            message.code_id = object.code_id;
        }
        else {
            message.code_id = 0;
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
const baseUpdateInstantiateConfigProposal = {
    title: "",
    description: "",
};
export const UpdateInstantiateConfigProposal = {
    encode(message, writer = Writer.create()) {
        if (message.title !== "") {
            writer.uint32(10).string(message.title);
        }
        if (message.description !== "") {
            writer.uint32(18).string(message.description);
        }
        for (const v of message.access_config_updates) {
            AccessConfigUpdate.encode(v, writer.uint32(26).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = {
            ...baseUpdateInstantiateConfigProposal,
        };
        message.access_config_updates = [];
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.title = reader.string();
                    break;
                case 2:
                    message.description = reader.string();
                    break;
                case 3:
                    message.access_config_updates.push(AccessConfigUpdate.decode(reader, reader.uint32()));
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
            ...baseUpdateInstantiateConfigProposal,
        };
        message.access_config_updates = [];
        if (object.title !== undefined && object.title !== null) {
            message.title = String(object.title);
        }
        else {
            message.title = "";
        }
        if (object.description !== undefined && object.description !== null) {
            message.description = String(object.description);
        }
        else {
            message.description = "";
        }
        if (object.access_config_updates !== undefined &&
            object.access_config_updates !== null) {
            for (const e of object.access_config_updates) {
                message.access_config_updates.push(AccessConfigUpdate.fromJSON(e));
            }
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.title !== undefined && (obj.title = message.title);
        message.description !== undefined &&
            (obj.description = message.description);
        if (message.access_config_updates) {
            obj.access_config_updates = message.access_config_updates.map((e) => e ? AccessConfigUpdate.toJSON(e) : undefined);
        }
        else {
            obj.access_config_updates = [];
        }
        return obj;
    },
    fromPartial(object) {
        const message = {
            ...baseUpdateInstantiateConfigProposal,
        };
        message.access_config_updates = [];
        if (object.title !== undefined && object.title !== null) {
            message.title = object.title;
        }
        else {
            message.title = "";
        }
        if (object.description !== undefined && object.description !== null) {
            message.description = object.description;
        }
        else {
            message.description = "";
        }
        if (object.access_config_updates !== undefined &&
            object.access_config_updates !== null) {
            for (const e of object.access_config_updates) {
                message.access_config_updates.push(AccessConfigUpdate.fromPartial(e));
            }
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
