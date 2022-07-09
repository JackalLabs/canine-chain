/* eslint-disable */
import * as Long from "long";
import { util, configure, Writer, Reader } from "protobufjs/minimal";
import { Params, CodeInfo, ContractInfo, Model, } from "../../../cosmwasm/wasm/v1/types";
import { MsgStoreCode, MsgInstantiateContract, MsgExecuteContract, } from "../../../cosmwasm/wasm/v1/tx";
export const protobufPackage = "cosmwasm.wasm.v1";
const baseGenesisState = {};
export const GenesisState = {
    encode(message, writer = Writer.create()) {
        if (message.params !== undefined) {
            Params.encode(message.params, writer.uint32(10).fork()).ldelim();
        }
        for (const v of message.codes) {
            Code.encode(v, writer.uint32(18).fork()).ldelim();
        }
        for (const v of message.contracts) {
            Contract.encode(v, writer.uint32(26).fork()).ldelim();
        }
        for (const v of message.sequences) {
            Sequence.encode(v, writer.uint32(34).fork()).ldelim();
        }
        for (const v of message.gen_msgs) {
            GenesisState_GenMsgs.encode(v, writer.uint32(42).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseGenesisState };
        message.codes = [];
        message.contracts = [];
        message.sequences = [];
        message.gen_msgs = [];
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.params = Params.decode(reader, reader.uint32());
                    break;
                case 2:
                    message.codes.push(Code.decode(reader, reader.uint32()));
                    break;
                case 3:
                    message.contracts.push(Contract.decode(reader, reader.uint32()));
                    break;
                case 4:
                    message.sequences.push(Sequence.decode(reader, reader.uint32()));
                    break;
                case 5:
                    message.gen_msgs.push(GenesisState_GenMsgs.decode(reader, reader.uint32()));
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
        message.codes = [];
        message.contracts = [];
        message.sequences = [];
        message.gen_msgs = [];
        if (object.params !== undefined && object.params !== null) {
            message.params = Params.fromJSON(object.params);
        }
        else {
            message.params = undefined;
        }
        if (object.codes !== undefined && object.codes !== null) {
            for (const e of object.codes) {
                message.codes.push(Code.fromJSON(e));
            }
        }
        if (object.contracts !== undefined && object.contracts !== null) {
            for (const e of object.contracts) {
                message.contracts.push(Contract.fromJSON(e));
            }
        }
        if (object.sequences !== undefined && object.sequences !== null) {
            for (const e of object.sequences) {
                message.sequences.push(Sequence.fromJSON(e));
            }
        }
        if (object.gen_msgs !== undefined && object.gen_msgs !== null) {
            for (const e of object.gen_msgs) {
                message.gen_msgs.push(GenesisState_GenMsgs.fromJSON(e));
            }
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.params !== undefined &&
            (obj.params = message.params ? Params.toJSON(message.params) : undefined);
        if (message.codes) {
            obj.codes = message.codes.map((e) => (e ? Code.toJSON(e) : undefined));
        }
        else {
            obj.codes = [];
        }
        if (message.contracts) {
            obj.contracts = message.contracts.map((e) => e ? Contract.toJSON(e) : undefined);
        }
        else {
            obj.contracts = [];
        }
        if (message.sequences) {
            obj.sequences = message.sequences.map((e) => e ? Sequence.toJSON(e) : undefined);
        }
        else {
            obj.sequences = [];
        }
        if (message.gen_msgs) {
            obj.gen_msgs = message.gen_msgs.map((e) => e ? GenesisState_GenMsgs.toJSON(e) : undefined);
        }
        else {
            obj.gen_msgs = [];
        }
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseGenesisState };
        message.codes = [];
        message.contracts = [];
        message.sequences = [];
        message.gen_msgs = [];
        if (object.params !== undefined && object.params !== null) {
            message.params = Params.fromPartial(object.params);
        }
        else {
            message.params = undefined;
        }
        if (object.codes !== undefined && object.codes !== null) {
            for (const e of object.codes) {
                message.codes.push(Code.fromPartial(e));
            }
        }
        if (object.contracts !== undefined && object.contracts !== null) {
            for (const e of object.contracts) {
                message.contracts.push(Contract.fromPartial(e));
            }
        }
        if (object.sequences !== undefined && object.sequences !== null) {
            for (const e of object.sequences) {
                message.sequences.push(Sequence.fromPartial(e));
            }
        }
        if (object.gen_msgs !== undefined && object.gen_msgs !== null) {
            for (const e of object.gen_msgs) {
                message.gen_msgs.push(GenesisState_GenMsgs.fromPartial(e));
            }
        }
        return message;
    },
};
const baseGenesisState_GenMsgs = {};
export const GenesisState_GenMsgs = {
    encode(message, writer = Writer.create()) {
        if (message.store_code !== undefined) {
            MsgStoreCode.encode(message.store_code, writer.uint32(10).fork()).ldelim();
        }
        if (message.instantiate_contract !== undefined) {
            MsgInstantiateContract.encode(message.instantiate_contract, writer.uint32(18).fork()).ldelim();
        }
        if (message.execute_contract !== undefined) {
            MsgExecuteContract.encode(message.execute_contract, writer.uint32(26).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseGenesisState_GenMsgs };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.store_code = MsgStoreCode.decode(reader, reader.uint32());
                    break;
                case 2:
                    message.instantiate_contract = MsgInstantiateContract.decode(reader, reader.uint32());
                    break;
                case 3:
                    message.execute_contract = MsgExecuteContract.decode(reader, reader.uint32());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseGenesisState_GenMsgs };
        if (object.store_code !== undefined && object.store_code !== null) {
            message.store_code = MsgStoreCode.fromJSON(object.store_code);
        }
        else {
            message.store_code = undefined;
        }
        if (object.instantiate_contract !== undefined &&
            object.instantiate_contract !== null) {
            message.instantiate_contract = MsgInstantiateContract.fromJSON(object.instantiate_contract);
        }
        else {
            message.instantiate_contract = undefined;
        }
        if (object.execute_contract !== undefined &&
            object.execute_contract !== null) {
            message.execute_contract = MsgExecuteContract.fromJSON(object.execute_contract);
        }
        else {
            message.execute_contract = undefined;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.store_code !== undefined &&
            (obj.store_code = message.store_code
                ? MsgStoreCode.toJSON(message.store_code)
                : undefined);
        message.instantiate_contract !== undefined &&
            (obj.instantiate_contract = message.instantiate_contract
                ? MsgInstantiateContract.toJSON(message.instantiate_contract)
                : undefined);
        message.execute_contract !== undefined &&
            (obj.execute_contract = message.execute_contract
                ? MsgExecuteContract.toJSON(message.execute_contract)
                : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseGenesisState_GenMsgs };
        if (object.store_code !== undefined && object.store_code !== null) {
            message.store_code = MsgStoreCode.fromPartial(object.store_code);
        }
        else {
            message.store_code = undefined;
        }
        if (object.instantiate_contract !== undefined &&
            object.instantiate_contract !== null) {
            message.instantiate_contract = MsgInstantiateContract.fromPartial(object.instantiate_contract);
        }
        else {
            message.instantiate_contract = undefined;
        }
        if (object.execute_contract !== undefined &&
            object.execute_contract !== null) {
            message.execute_contract = MsgExecuteContract.fromPartial(object.execute_contract);
        }
        else {
            message.execute_contract = undefined;
        }
        return message;
    },
};
const baseCode = { code_id: 0, pinned: false };
export const Code = {
    encode(message, writer = Writer.create()) {
        if (message.code_id !== 0) {
            writer.uint32(8).uint64(message.code_id);
        }
        if (message.code_info !== undefined) {
            CodeInfo.encode(message.code_info, writer.uint32(18).fork()).ldelim();
        }
        if (message.code_bytes.length !== 0) {
            writer.uint32(26).bytes(message.code_bytes);
        }
        if (message.pinned === true) {
            writer.uint32(32).bool(message.pinned);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseCode };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.code_id = longToNumber(reader.uint64());
                    break;
                case 2:
                    message.code_info = CodeInfo.decode(reader, reader.uint32());
                    break;
                case 3:
                    message.code_bytes = reader.bytes();
                    break;
                case 4:
                    message.pinned = reader.bool();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseCode };
        if (object.code_id !== undefined && object.code_id !== null) {
            message.code_id = Number(object.code_id);
        }
        else {
            message.code_id = 0;
        }
        if (object.code_info !== undefined && object.code_info !== null) {
            message.code_info = CodeInfo.fromJSON(object.code_info);
        }
        else {
            message.code_info = undefined;
        }
        if (object.code_bytes !== undefined && object.code_bytes !== null) {
            message.code_bytes = bytesFromBase64(object.code_bytes);
        }
        if (object.pinned !== undefined && object.pinned !== null) {
            message.pinned = Boolean(object.pinned);
        }
        else {
            message.pinned = false;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.code_id !== undefined && (obj.code_id = message.code_id);
        message.code_info !== undefined &&
            (obj.code_info = message.code_info
                ? CodeInfo.toJSON(message.code_info)
                : undefined);
        message.code_bytes !== undefined &&
            (obj.code_bytes = base64FromBytes(message.code_bytes !== undefined ? message.code_bytes : new Uint8Array()));
        message.pinned !== undefined && (obj.pinned = message.pinned);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseCode };
        if (object.code_id !== undefined && object.code_id !== null) {
            message.code_id = object.code_id;
        }
        else {
            message.code_id = 0;
        }
        if (object.code_info !== undefined && object.code_info !== null) {
            message.code_info = CodeInfo.fromPartial(object.code_info);
        }
        else {
            message.code_info = undefined;
        }
        if (object.code_bytes !== undefined && object.code_bytes !== null) {
            message.code_bytes = object.code_bytes;
        }
        else {
            message.code_bytes = new Uint8Array();
        }
        if (object.pinned !== undefined && object.pinned !== null) {
            message.pinned = object.pinned;
        }
        else {
            message.pinned = false;
        }
        return message;
    },
};
const baseContract = { contract_address: "" };
export const Contract = {
    encode(message, writer = Writer.create()) {
        if (message.contract_address !== "") {
            writer.uint32(10).string(message.contract_address);
        }
        if (message.contract_info !== undefined) {
            ContractInfo.encode(message.contract_info, writer.uint32(18).fork()).ldelim();
        }
        for (const v of message.contract_state) {
            Model.encode(v, writer.uint32(26).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseContract };
        message.contract_state = [];
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.contract_address = reader.string();
                    break;
                case 2:
                    message.contract_info = ContractInfo.decode(reader, reader.uint32());
                    break;
                case 3:
                    message.contract_state.push(Model.decode(reader, reader.uint32()));
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseContract };
        message.contract_state = [];
        if (object.contract_address !== undefined &&
            object.contract_address !== null) {
            message.contract_address = String(object.contract_address);
        }
        else {
            message.contract_address = "";
        }
        if (object.contract_info !== undefined && object.contract_info !== null) {
            message.contract_info = ContractInfo.fromJSON(object.contract_info);
        }
        else {
            message.contract_info = undefined;
        }
        if (object.contract_state !== undefined && object.contract_state !== null) {
            for (const e of object.contract_state) {
                message.contract_state.push(Model.fromJSON(e));
            }
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.contract_address !== undefined &&
            (obj.contract_address = message.contract_address);
        message.contract_info !== undefined &&
            (obj.contract_info = message.contract_info
                ? ContractInfo.toJSON(message.contract_info)
                : undefined);
        if (message.contract_state) {
            obj.contract_state = message.contract_state.map((e) => e ? Model.toJSON(e) : undefined);
        }
        else {
            obj.contract_state = [];
        }
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseContract };
        message.contract_state = [];
        if (object.contract_address !== undefined &&
            object.contract_address !== null) {
            message.contract_address = object.contract_address;
        }
        else {
            message.contract_address = "";
        }
        if (object.contract_info !== undefined && object.contract_info !== null) {
            message.contract_info = ContractInfo.fromPartial(object.contract_info);
        }
        else {
            message.contract_info = undefined;
        }
        if (object.contract_state !== undefined && object.contract_state !== null) {
            for (const e of object.contract_state) {
                message.contract_state.push(Model.fromPartial(e));
            }
        }
        return message;
    },
};
const baseSequence = { value: 0 };
export const Sequence = {
    encode(message, writer = Writer.create()) {
        if (message.id_key.length !== 0) {
            writer.uint32(10).bytes(message.id_key);
        }
        if (message.value !== 0) {
            writer.uint32(16).uint64(message.value);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseSequence };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.id_key = reader.bytes();
                    break;
                case 2:
                    message.value = longToNumber(reader.uint64());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseSequence };
        if (object.id_key !== undefined && object.id_key !== null) {
            message.id_key = bytesFromBase64(object.id_key);
        }
        if (object.value !== undefined && object.value !== null) {
            message.value = Number(object.value);
        }
        else {
            message.value = 0;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.id_key !== undefined &&
            (obj.id_key = base64FromBytes(message.id_key !== undefined ? message.id_key : new Uint8Array()));
        message.value !== undefined && (obj.value = message.value);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseSequence };
        if (object.id_key !== undefined && object.id_key !== null) {
            message.id_key = object.id_key;
        }
        else {
            message.id_key = new Uint8Array();
        }
        if (object.value !== undefined && object.value !== null) {
            message.value = object.value;
        }
        else {
            message.value = 0;
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
