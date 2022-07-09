import { StdFee } from "@cosmjs/launchpad";
import { Registry, OfflineSigner, EncodeObject } from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgAcceptBid } from "./types/rns/tx";
import { MsgBid } from "./types/rns/tx";
import { MsgTransfer } from "./types/rns/tx";
import { MsgList } from "./types/rns/tx";
import { MsgRegister } from "./types/rns/tx";
import { MsgBuy } from "./types/rns/tx";
import { MsgCancelBid } from "./types/rns/tx";
import { MsgDelist } from "./types/rns/tx";
export declare const MissingWalletError: Error;
export declare const registry: Registry;
interface TxClientOptions {
    addr: string;
}
interface SignAndBroadcastOptions {
    fee: StdFee;
    memo?: string;
}
declare const txClient: (wallet: OfflineSigner, { addr: addr }?: TxClientOptions) => Promise<{
    signAndBroadcast: (msgs: EncodeObject[], { fee, memo }?: SignAndBroadcastOptions) => any;
    msgAcceptBid: (data: MsgAcceptBid) => EncodeObject;
    msgBid: (data: MsgBid) => EncodeObject;
    msgTransfer: (data: MsgTransfer) => EncodeObject;
    msgList: (data: MsgList) => EncodeObject;
    msgRegister: (data: MsgRegister) => EncodeObject;
    msgBuy: (data: MsgBuy) => EncodeObject;
    msgCancelBid: (data: MsgCancelBid) => EncodeObject;
    msgDelist: (data: MsgDelist) => EncodeObject;
}>;
interface QueryClientOptions {
    addr: string;
}
declare const queryClient: ({ addr: addr }?: QueryClientOptions) => Promise<Api<unknown>>;
export { txClient, queryClient, };
