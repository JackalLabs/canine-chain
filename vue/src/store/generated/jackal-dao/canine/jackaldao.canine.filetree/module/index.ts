// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.

import { StdFee } from "@cosmjs/launchpad";
import { SigningStargateClient } from "@cosmjs/stargate";
import { Registry, OfflineSigner, EncodeObject, DirectSecp256k1HdWallet } from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgPostFile } from "./types/filetree/tx";
import { MsgAddViewers } from "./types/filetree/tx";
<<<<<<< HEAD
import { MsgPostkey } from "./types/filetree/tx";
=======
>>>>>>> d3584237d1295a036e3a1be837b66d909f3d7823


const types = [
  ["/jackaldao.canine.filetree.MsgPostFile", MsgPostFile],
  ["/jackaldao.canine.filetree.MsgAddViewers", MsgAddViewers],
<<<<<<< HEAD
  ["/jackaldao.canine.filetree.MsgPostkey", MsgPostkey],
=======
>>>>>>> d3584237d1295a036e3a1be837b66d909f3d7823
  
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
    msgPostFile: (data: MsgPostFile): EncodeObject => ({ typeUrl: "/jackaldao.canine.filetree.MsgPostFile", value: MsgPostFile.fromPartial( data ) }),
    msgAddViewers: (data: MsgAddViewers): EncodeObject => ({ typeUrl: "/jackaldao.canine.filetree.MsgAddViewers", value: MsgAddViewers.fromPartial( data ) }),
<<<<<<< HEAD
    msgPostkey: (data: MsgPostkey): EncodeObject => ({ typeUrl: "/jackaldao.canine.filetree.MsgPostkey", value: MsgPostkey.fromPartial( data ) }),
=======
>>>>>>> d3584237d1295a036e3a1be837b66d909f3d7823
    
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
