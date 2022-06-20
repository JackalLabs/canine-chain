// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.

import { StdFee } from "@cosmjs/launchpad";
import { SigningStargateClient } from "@cosmjs/stargate";
import { Registry, OfflineSigner, EncodeObject, DirectSecp256k1HdWallet } from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgAllowSave } from "./types/jklmining/tx";
import { MsgUpdateMinerClaims } from "./types/jklmining/tx";
import { MsgClaimSave } from "./types/jklmining/tx";
import { MsgUpdateMiners } from "./types/jklmining/tx";
import { MsgCreateMinerClaims } from "./types/jklmining/tx";
import { MsgCreateMiners } from "./types/jklmining/tx";
import { MsgDeleteMinerClaims } from "./types/jklmining/tx";
import { MsgDeleteMiners } from "./types/jklmining/tx";


const types = [
  ["/jackaldao.canine.jklmining.MsgAllowSave", MsgAllowSave],
  ["/jackaldao.canine.jklmining.MsgUpdateMinerClaims", MsgUpdateMinerClaims],
  ["/jackaldao.canine.jklmining.MsgClaimSave", MsgClaimSave],
  ["/jackaldao.canine.jklmining.MsgUpdateMiners", MsgUpdateMiners],
  ["/jackaldao.canine.jklmining.MsgCreateMinerClaims", MsgCreateMinerClaims],
  ["/jackaldao.canine.jklmining.MsgCreateMiners", MsgCreateMiners],
  ["/jackaldao.canine.jklmining.MsgDeleteMinerClaims", MsgDeleteMinerClaims],
  ["/jackaldao.canine.jklmining.MsgDeleteMiners", MsgDeleteMiners],
  
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
    msgAllowSave: (data: MsgAllowSave): EncodeObject => ({ typeUrl: "/jackaldao.canine.jklmining.MsgAllowSave", value: MsgAllowSave.fromPartial( data ) }),
    msgUpdateMinerClaims: (data: MsgUpdateMinerClaims): EncodeObject => ({ typeUrl: "/jackaldao.canine.jklmining.MsgUpdateMinerClaims", value: MsgUpdateMinerClaims.fromPartial( data ) }),
    msgClaimSave: (data: MsgClaimSave): EncodeObject => ({ typeUrl: "/jackaldao.canine.jklmining.MsgClaimSave", value: MsgClaimSave.fromPartial( data ) }),
    msgUpdateMiners: (data: MsgUpdateMiners): EncodeObject => ({ typeUrl: "/jackaldao.canine.jklmining.MsgUpdateMiners", value: MsgUpdateMiners.fromPartial( data ) }),
    msgCreateMinerClaims: (data: MsgCreateMinerClaims): EncodeObject => ({ typeUrl: "/jackaldao.canine.jklmining.MsgCreateMinerClaims", value: MsgCreateMinerClaims.fromPartial( data ) }),
    msgCreateMiners: (data: MsgCreateMiners): EncodeObject => ({ typeUrl: "/jackaldao.canine.jklmining.MsgCreateMiners", value: MsgCreateMiners.fromPartial( data ) }),
    msgDeleteMinerClaims: (data: MsgDeleteMinerClaims): EncodeObject => ({ typeUrl: "/jackaldao.canine.jklmining.MsgDeleteMinerClaims", value: MsgDeleteMinerClaims.fromPartial( data ) }),
    msgDeleteMiners: (data: MsgDeleteMiners): EncodeObject => ({ typeUrl: "/jackaldao.canine.jklmining.MsgDeleteMiners", value: MsgDeleteMiners.fromPartial( data ) }),
    
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
