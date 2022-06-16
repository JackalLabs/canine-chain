import { StdFee } from "@cosmjs/launchpad";
import { Registry, OfflineSigner, EncodeObject } from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgClaimSave } from "./types/jklmining/tx";
import { MsgCreateMiners } from "./types/jklmining/tx";
import { MsgUpdateMiners } from "./types/jklmining/tx";
import { MsgAllowSave } from "./types/jklmining/tx";
import { MsgDeleteMiners } from "./types/jklmining/tx";
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
    msgClaimSave: (data: MsgClaimSave) => EncodeObject;
    msgCreateMiners: (data: MsgCreateMiners) => EncodeObject;
    msgUpdateMiners: (data: MsgUpdateMiners) => EncodeObject;
    msgAllowSave: (data: MsgAllowSave) => EncodeObject;
    msgDeleteMiners: (data: MsgDeleteMiners) => EncodeObject;
}>;
interface QueryClientOptions {
    addr: string;
}
declare const queryClient: ({ addr: addr }?: QueryClientOptions) => Promise<Api<unknown>>;
export { txClient, queryClient, };
