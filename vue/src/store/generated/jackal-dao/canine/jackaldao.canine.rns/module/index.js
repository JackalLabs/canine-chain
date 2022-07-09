// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.
import { SigningStargateClient } from "@cosmjs/stargate";
import { Registry } from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgAcceptBid } from "./types/rns/tx";
import { MsgBid } from "./types/rns/tx";
import { MsgTransfer } from "./types/rns/tx";
import { MsgList } from "./types/rns/tx";
import { MsgRegister } from "./types/rns/tx";
import { MsgBuy } from "./types/rns/tx";
import { MsgCancelBid } from "./types/rns/tx";
import { MsgDelist } from "./types/rns/tx";
const types = [
    ["/jackaldao.canine.rns.MsgAcceptBid", MsgAcceptBid],
    ["/jackaldao.canine.rns.MsgBid", MsgBid],
    ["/jackaldao.canine.rns.MsgTransfer", MsgTransfer],
    ["/jackaldao.canine.rns.MsgList", MsgList],
    ["/jackaldao.canine.rns.MsgRegister", MsgRegister],
    ["/jackaldao.canine.rns.MsgBuy", MsgBuy],
    ["/jackaldao.canine.rns.MsgCancelBid", MsgCancelBid],
    ["/jackaldao.canine.rns.MsgDelist", MsgDelist],
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
        msgAcceptBid: (data) => ({ typeUrl: "/jackaldao.canine.rns.MsgAcceptBid", value: MsgAcceptBid.fromPartial(data) }),
        msgBid: (data) => ({ typeUrl: "/jackaldao.canine.rns.MsgBid", value: MsgBid.fromPartial(data) }),
        msgTransfer: (data) => ({ typeUrl: "/jackaldao.canine.rns.MsgTransfer", value: MsgTransfer.fromPartial(data) }),
        msgList: (data) => ({ typeUrl: "/jackaldao.canine.rns.MsgList", value: MsgList.fromPartial(data) }),
        msgRegister: (data) => ({ typeUrl: "/jackaldao.canine.rns.MsgRegister", value: MsgRegister.fromPartial(data) }),
        msgBuy: (data) => ({ typeUrl: "/jackaldao.canine.rns.MsgBuy", value: MsgBuy.fromPartial(data) }),
        msgCancelBid: (data) => ({ typeUrl: "/jackaldao.canine.rns.MsgCancelBid", value: MsgCancelBid.fromPartial(data) }),
        msgDelist: (data) => ({ typeUrl: "/jackaldao.canine.rns.MsgDelist", value: MsgDelist.fromPartial(data) }),
    };
};
const queryClient = async ({ addr: addr } = { addr: "http://localhost:1317" }) => {
    return new Api({ baseUrl: addr });
};
export { txClient, queryClient, };
