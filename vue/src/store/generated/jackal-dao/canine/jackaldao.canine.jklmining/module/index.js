// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.
import { SigningStargateClient } from "@cosmjs/stargate";
import { Registry } from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgAllowSave } from "./types/jklmining/tx";
import { MsgUpdateMinerClaims } from "./types/jklmining/tx";
import { MsgCreateMiners } from "./types/jklmining/tx";
import { MsgClaimSave } from "./types/jklmining/tx";
import { MsgDeleteMiners } from "./types/jklmining/tx";
import { MsgCreateMinerClaims } from "./types/jklmining/tx";
import { MsgUpdateMiners } from "./types/jklmining/tx";
import { MsgDeleteMinerClaims } from "./types/jklmining/tx";
const types = [
    ["/jackaldao.canine.jklmining.MsgAllowSave", MsgAllowSave],
    ["/jackaldao.canine.jklmining.MsgUpdateMinerClaims", MsgUpdateMinerClaims],
    ["/jackaldao.canine.jklmining.MsgCreateMiners", MsgCreateMiners],
    ["/jackaldao.canine.jklmining.MsgClaimSave", MsgClaimSave],
    ["/jackaldao.canine.jklmining.MsgDeleteMiners", MsgDeleteMiners],
    ["/jackaldao.canine.jklmining.MsgCreateMinerClaims", MsgCreateMinerClaims],
    ["/jackaldao.canine.jklmining.MsgUpdateMiners", MsgUpdateMiners],
    ["/jackaldao.canine.jklmining.MsgDeleteMinerClaims", MsgDeleteMinerClaims],
];
export const MissingWalletError = new Error("wallet is required");
export const registry = new Registry(types);
const defaultFee = {
    amount: [],
    gas: "200000",
};
const txClient = async (wallet, { addr: addr } = { addr: "http://localhost:26657" }) => {
    if (!wallet)
        throw MissingWalletError;
    let client;
    if (addr) {
        client = await SigningStargateClient.connectWithSigner(addr, wallet, { registry });
    }
    else {
        client = await SigningStargateClient.offline(wallet, { registry });
    }
    const { address } = (await wallet.getAccounts())[0];
    return {
        signAndBroadcast: (msgs, { fee, memo } = { fee: defaultFee, memo: "" }) => client.signAndBroadcast(address, msgs, fee, memo),
        msgAllowSave: (data) => ({ typeUrl: "/jackaldao.canine.jklmining.MsgAllowSave", value: MsgAllowSave.fromPartial(data) }),
        msgUpdateMinerClaims: (data) => ({ typeUrl: "/jackaldao.canine.jklmining.MsgUpdateMinerClaims", value: MsgUpdateMinerClaims.fromPartial(data) }),
        msgCreateMiners: (data) => ({ typeUrl: "/jackaldao.canine.jklmining.MsgCreateMiners", value: MsgCreateMiners.fromPartial(data) }),
        msgClaimSave: (data) => ({ typeUrl: "/jackaldao.canine.jklmining.MsgClaimSave", value: MsgClaimSave.fromPartial(data) }),
        msgDeleteMiners: (data) => ({ typeUrl: "/jackaldao.canine.jklmining.MsgDeleteMiners", value: MsgDeleteMiners.fromPartial(data) }),
        msgCreateMinerClaims: (data) => ({ typeUrl: "/jackaldao.canine.jklmining.MsgCreateMinerClaims", value: MsgCreateMinerClaims.fromPartial(data) }),
        msgUpdateMiners: (data) => ({ typeUrl: "/jackaldao.canine.jklmining.MsgUpdateMiners", value: MsgUpdateMiners.fromPartial(data) }),
        msgDeleteMinerClaims: (data) => ({ typeUrl: "/jackaldao.canine.jklmining.MsgDeleteMinerClaims", value: MsgDeleteMinerClaims.fromPartial(data) }),
    };
};
const queryClient = async ({ addr: addr } = { addr: "http://localhost:1317" }) => {
    return new Api({ baseUrl: addr });
};
export { txClient, queryClient, };
