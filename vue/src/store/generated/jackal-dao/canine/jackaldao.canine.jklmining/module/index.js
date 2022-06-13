// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.
import { SigningStargateClient } from "@cosmjs/stargate";
import { Registry } from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgAllowSave } from "./types/jklmining/tx";
import { MsgDeleteMiners } from "./types/jklmining/tx";
import { MsgCreateMiners } from "./types/jklmining/tx";
import { MsgUpdateMiners } from "./types/jklmining/tx";
const types = [
    ["/jackaldao.canine.jklmining.MsgAllowSave", MsgAllowSave],
    ["/jackaldao.canine.jklmining.MsgDeleteMiners", MsgDeleteMiners],
    ["/jackaldao.canine.jklmining.MsgCreateMiners", MsgCreateMiners],
    ["/jackaldao.canine.jklmining.MsgUpdateMiners", MsgUpdateMiners],
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
        msgDeleteMiners: (data) => ({ typeUrl: "/jackaldao.canine.jklmining.MsgDeleteMiners", value: MsgDeleteMiners.fromPartial(data) }),
        msgCreateMiners: (data) => ({ typeUrl: "/jackaldao.canine.jklmining.MsgCreateMiners", value: MsgCreateMiners.fromPartial(data) }),
        msgUpdateMiners: (data) => ({ typeUrl: "/jackaldao.canine.jklmining.MsgUpdateMiners", value: MsgUpdateMiners.fromPartial(data) }),
    };
};
const queryClient = async ({ addr: addr } = { addr: "http://localhost:1317" }) => {
    return new Api({ baseUrl: addr });
};
export { txClient, queryClient, };
