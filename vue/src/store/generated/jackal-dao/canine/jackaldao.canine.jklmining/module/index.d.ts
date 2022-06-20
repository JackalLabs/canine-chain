import { StdFee } from "@cosmjs/launchpad";
import { Registry, OfflineSigner, EncodeObject } from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgCreateMiners } from "./types/jklmining/tx";
import { MsgDeleteMiners } from "./types/jklmining/tx";
import { MsgAllowSave } from "./types/jklmining/tx";
import { MsgCreateMinerClaims } from "./types/jklmining/tx";
import { MsgDeleteMinerClaims } from "./types/jklmining/tx";
import { MsgUpdateMiners } from "./types/jklmining/tx";
import { MsgClaimSave } from "./types/jklmining/tx";
import { MsgUpdateMinerClaims } from "./types/jklmining/tx";
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
    msgCreateMiners: (data: MsgCreateMiners) => EncodeObject;
    msgDeleteMiners: (data: MsgDeleteMiners) => EncodeObject;
    msgAllowSave: (data: MsgAllowSave) => EncodeObject;
    msgCreateMinerClaims: (data: MsgCreateMinerClaims) => EncodeObject;
    msgDeleteMinerClaims: (data: MsgDeleteMinerClaims) => EncodeObject;
    msgUpdateMiners: (data: MsgUpdateMiners) => EncodeObject;
    msgClaimSave: (data: MsgClaimSave) => EncodeObject;
    msgUpdateMinerClaims: (data: MsgUpdateMinerClaims) => EncodeObject;
}>;
interface QueryClientOptions {
    addr: string;
}
declare const queryClient: ({ addr: addr }?: QueryClientOptions) => Promise<Api<unknown>>;
export { txClient, queryClient, };
