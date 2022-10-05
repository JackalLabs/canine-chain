// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.

import { StdFee } from "@cosmjs/launchpad";
import { SigningStargateClient } from "@cosmjs/stargate";
import { Registry, OfflineSigner, EncodeObject, DirectSecp256k1HdWallet } from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgDepositLPool } from "./types/lp/tx";
import { MsgWithdrawLPool } from "./types/lp/tx";
import { MsgCreateLPool } from "./types/lp/tx";
import { MsgSwap } from "./types/lp/tx";


const types = [
  ["/jackaldao.canine.lp.MsgDepositLPool", MsgDepositLPool],
  ["/jackaldao.canine.lp.MsgWithdrawLPool", MsgWithdrawLPool],
  ["/jackaldao.canine.lp.MsgCreateLPool", MsgCreateLPool],
  ["/jackaldao.canine.lp.MsgSwap", MsgSwap],
  
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
    msgDepositLPool: (data: MsgDepositLPool): EncodeObject => ({ typeUrl: "/jackaldao.canine.lp.MsgDepositLPool", value: MsgDepositLPool.fromPartial( data ) }),
    msgWithdrawLPool: (data: MsgWithdrawLPool): EncodeObject => ({ typeUrl: "/jackaldao.canine.lp.MsgWithdrawLPool", value: MsgWithdrawLPool.fromPartial( data ) }),
    msgCreateLPool: (data: MsgCreateLPool): EncodeObject => ({ typeUrl: "/jackaldao.canine.lp.MsgCreateLPool", value: MsgCreateLPool.fromPartial( data ) }),
    msgSwap: (data: MsgSwap): EncodeObject => ({ typeUrl: "/jackaldao.canine.lp.MsgSwap", value: MsgSwap.fromPartial( data ) }),
    
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
