// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.

import { StdFee } from "@cosmjs/launchpad";
import { SigningStargateClient } from "@cosmjs/stargate";
import { Registry, OfflineSigner, EncodeObject, DirectSecp256k1HdWallet } from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgDeleteActiveDeals } from "./types/storage/tx";
import { MsgItem } from "./types/storage/tx";
import { MsgSetProviderTotalspace } from "./types/storage/tx";
import { MsgInitProvider } from "./types/storage/tx";
import { MsgBuyStorage } from "./types/storage/tx";
import { MsgDeleteProviders } from "./types/storage/tx";
import { MsgCreateProofs } from "./types/storage/tx";
import { MsgDeleteContracts } from "./types/storage/tx";
import { MsgSignContract } from "./types/storage/tx";
import { MsgSetProviderIp } from "./types/storage/tx";
import { MsgPostContract } from "./types/storage/tx";
import { MsgDeleteProofs } from "./types/storage/tx";
import { MsgCreateActiveDeals } from "./types/storage/tx";
import { MsgUpdateProviders } from "./types/storage/tx";
import { MsgCancelContract } from "./types/storage/tx";
import { MsgUpdateContracts } from "./types/storage/tx";
import { MsgCreateProviders } from "./types/storage/tx";
import { MsgCreateContracts } from "./types/storage/tx";
import { MsgPostproof } from "./types/storage/tx";
import { MsgUpdateProofs } from "./types/storage/tx";
import { MsgUpdateActiveDeals } from "./types/storage/tx";


const types = [
  ["/jackaldao.canine.storage.MsgDeleteActiveDeals", MsgDeleteActiveDeals],
  ["/jackaldao.canine.storage.MsgItem", MsgItem],
  ["/jackaldao.canine.storage.MsgSetProviderTotalspace", MsgSetProviderTotalspace],
  ["/jackaldao.canine.storage.MsgInitProvider", MsgInitProvider],
  ["/jackaldao.canine.storage.MsgBuyStorage", MsgBuyStorage],
  ["/jackaldao.canine.storage.MsgDeleteProviders", MsgDeleteProviders],
  ["/jackaldao.canine.storage.MsgCreateProofs", MsgCreateProofs],
  ["/jackaldao.canine.storage.MsgDeleteContracts", MsgDeleteContracts],
  ["/jackaldao.canine.storage.MsgSignContract", MsgSignContract],
  ["/jackaldao.canine.storage.MsgSetProviderIp", MsgSetProviderIp],
  ["/jackaldao.canine.storage.MsgPostContract", MsgPostContract],
  ["/jackaldao.canine.storage.MsgDeleteProofs", MsgDeleteProofs],
  ["/jackaldao.canine.storage.MsgCreateActiveDeals", MsgCreateActiveDeals],
  ["/jackaldao.canine.storage.MsgUpdateProviders", MsgUpdateProviders],
  ["/jackaldao.canine.storage.MsgCancelContract", MsgCancelContract],
  ["/jackaldao.canine.storage.MsgUpdateContracts", MsgUpdateContracts],
  ["/jackaldao.canine.storage.MsgCreateProviders", MsgCreateProviders],
  ["/jackaldao.canine.storage.MsgCreateContracts", MsgCreateContracts],
  ["/jackaldao.canine.storage.MsgPostproof", MsgPostproof],
  ["/jackaldao.canine.storage.MsgUpdateProofs", MsgUpdateProofs],
  ["/jackaldao.canine.storage.MsgUpdateActiveDeals", MsgUpdateActiveDeals],
  
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
    msgDeleteActiveDeals: (data: MsgDeleteActiveDeals): EncodeObject => ({ typeUrl: "/jackaldao.canine.storage.MsgDeleteActiveDeals", value: MsgDeleteActiveDeals.fromPartial( data ) }),
    msgItem: (data: MsgItem): EncodeObject => ({ typeUrl: "/jackaldao.canine.storage.MsgItem", value: MsgItem.fromPartial( data ) }),
    msgSetProviderTotalspace: (data: MsgSetProviderTotalspace): EncodeObject => ({ typeUrl: "/jackaldao.canine.storage.MsgSetProviderTotalspace", value: MsgSetProviderTotalspace.fromPartial( data ) }),
    msgInitProvider: (data: MsgInitProvider): EncodeObject => ({ typeUrl: "/jackaldao.canine.storage.MsgInitProvider", value: MsgInitProvider.fromPartial( data ) }),
    msgBuyStorage: (data: MsgBuyStorage): EncodeObject => ({ typeUrl: "/jackaldao.canine.storage.MsgBuyStorage", value: MsgBuyStorage.fromPartial( data ) }),
    msgDeleteProviders: (data: MsgDeleteProviders): EncodeObject => ({ typeUrl: "/jackaldao.canine.storage.MsgDeleteProviders", value: MsgDeleteProviders.fromPartial( data ) }),
    msgCreateProofs: (data: MsgCreateProofs): EncodeObject => ({ typeUrl: "/jackaldao.canine.storage.MsgCreateProofs", value: MsgCreateProofs.fromPartial( data ) }),
    msgDeleteContracts: (data: MsgDeleteContracts): EncodeObject => ({ typeUrl: "/jackaldao.canine.storage.MsgDeleteContracts", value: MsgDeleteContracts.fromPartial( data ) }),
    msgSignContract: (data: MsgSignContract): EncodeObject => ({ typeUrl: "/jackaldao.canine.storage.MsgSignContract", value: MsgSignContract.fromPartial( data ) }),
    msgSetProviderIp: (data: MsgSetProviderIp): EncodeObject => ({ typeUrl: "/jackaldao.canine.storage.MsgSetProviderIp", value: MsgSetProviderIp.fromPartial( data ) }),
    msgPostContract: (data: MsgPostContract): EncodeObject => ({ typeUrl: "/jackaldao.canine.storage.MsgPostContract", value: MsgPostContract.fromPartial( data ) }),
    msgDeleteProofs: (data: MsgDeleteProofs): EncodeObject => ({ typeUrl: "/jackaldao.canine.storage.MsgDeleteProofs", value: MsgDeleteProofs.fromPartial( data ) }),
    msgCreateActiveDeals: (data: MsgCreateActiveDeals): EncodeObject => ({ typeUrl: "/jackaldao.canine.storage.MsgCreateActiveDeals", value: MsgCreateActiveDeals.fromPartial( data ) }),
    msgUpdateProviders: (data: MsgUpdateProviders): EncodeObject => ({ typeUrl: "/jackaldao.canine.storage.MsgUpdateProviders", value: MsgUpdateProviders.fromPartial( data ) }),
    msgCancelContract: (data: MsgCancelContract): EncodeObject => ({ typeUrl: "/jackaldao.canine.storage.MsgCancelContract", value: MsgCancelContract.fromPartial( data ) }),
    msgUpdateContracts: (data: MsgUpdateContracts): EncodeObject => ({ typeUrl: "/jackaldao.canine.storage.MsgUpdateContracts", value: MsgUpdateContracts.fromPartial( data ) }),
    msgCreateProviders: (data: MsgCreateProviders): EncodeObject => ({ typeUrl: "/jackaldao.canine.storage.MsgCreateProviders", value: MsgCreateProviders.fromPartial( data ) }),
    msgCreateContracts: (data: MsgCreateContracts): EncodeObject => ({ typeUrl: "/jackaldao.canine.storage.MsgCreateContracts", value: MsgCreateContracts.fromPartial( data ) }),
    msgPostproof: (data: MsgPostproof): EncodeObject => ({ typeUrl: "/jackaldao.canine.storage.MsgPostproof", value: MsgPostproof.fromPartial( data ) }),
    msgUpdateProofs: (data: MsgUpdateProofs): EncodeObject => ({ typeUrl: "/jackaldao.canine.storage.MsgUpdateProofs", value: MsgUpdateProofs.fromPartial( data ) }),
    msgUpdateActiveDeals: (data: MsgUpdateActiveDeals): EncodeObject => ({ typeUrl: "/jackaldao.canine.storage.MsgUpdateActiveDeals", value: MsgUpdateActiveDeals.fromPartial( data ) }),
    
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
