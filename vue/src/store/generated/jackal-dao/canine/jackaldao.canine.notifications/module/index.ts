// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.

import { StdFee } from "@cosmjs/launchpad";
import { SigningStargateClient } from "@cosmjs/stargate";
import { Registry, OfflineSigner, EncodeObject, DirectSecp256k1HdWallet } from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgUpdateNotifications } from "./types/notifications/tx";
import { MsgSetCounter } from "./types/notifications/tx";
import { MsgCreateNotifications } from "./types/notifications/tx";
import { MsgAddSenders } from "./types/notifications/tx";
import { MsgDeleteNotifications } from "./types/notifications/tx";


const types = [
  ["/jackaldao.canine.notifications.MsgUpdateNotifications", MsgUpdateNotifications],
  ["/jackaldao.canine.notifications.MsgSetCounter", MsgSetCounter],
  ["/jackaldao.canine.notifications.MsgCreateNotifications", MsgCreateNotifications],
  ["/jackaldao.canine.notifications.MsgAddSenders", MsgAddSenders],
  ["/jackaldao.canine.notifications.MsgDeleteNotifications", MsgDeleteNotifications],
  
];
export const MissingWalletError = new Error("wallet is required");

export const registry = new Registry(<any>types);

const defaultFee = {
  amount: [],
  gas: "200000",
};

interface TxClientOptions {
  addr: string
}

interface SignAndBroadcastOptions {
  fee: StdFee,
  memo?: string
}

const txClient = async (wallet: OfflineSigner, { addr: addr }: TxClientOptions = { addr: "http://localhost:26657" }) => {
  if (!wallet) throw MissingWalletError;
  let client;
  if (addr) {
    client = await SigningStargateClient.connectWithSigner(addr, wallet, { registry });
  }else{
    client = await SigningStargateClient.offline( wallet, { registry });
  }
  const { address } = (await wallet.getAccounts())[0];

  return {
    signAndBroadcast: (msgs: EncodeObject[], { fee, memo }: SignAndBroadcastOptions = {fee: defaultFee, memo: ""}) => client.signAndBroadcast(address, msgs, fee,memo),
    msgUpdateNotifications: (data: MsgUpdateNotifications): EncodeObject => ({ typeUrl: "/jackaldao.canine.notifications.MsgUpdateNotifications", value: MsgUpdateNotifications.fromPartial( data ) }),
    msgSetCounter: (data: MsgSetCounter): EncodeObject => ({ typeUrl: "/jackaldao.canine.notifications.MsgSetCounter", value: MsgSetCounter.fromPartial( data ) }),
    msgCreateNotifications: (data: MsgCreateNotifications): EncodeObject => ({ typeUrl: "/jackaldao.canine.notifications.MsgCreateNotifications", value: MsgCreateNotifications.fromPartial( data ) }),
    msgAddSenders: (data: MsgAddSenders): EncodeObject => ({ typeUrl: "/jackaldao.canine.notifications.MsgAddSenders", value: MsgAddSenders.fromPartial( data ) }),
    msgDeleteNotifications: (data: MsgDeleteNotifications): EncodeObject => ({ typeUrl: "/jackaldao.canine.notifications.MsgDeleteNotifications", value: MsgDeleteNotifications.fromPartial( data ) }),
    
  };
};

interface QueryClientOptions {
  addr: string
}

const queryClient = async ({ addr: addr }: QueryClientOptions = { addr: "http://localhost:1317" }) => {
  return new Api({ baseUrl: addr });
};

export {
  txClient,
  queryClient,
};
