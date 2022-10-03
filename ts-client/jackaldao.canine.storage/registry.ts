import { GeneratedType } from "@cosmjs/proto-signing";
<<<<<<< HEAD
import { MsgSetMinerIp } from "./types/storage/tx";
import { MsgCreateActiveDeals } from "./types/storage/tx";
import { MsgBuyStorage } from "./types/storage/tx";
import { MsgCreateContracts } from "./types/storage/tx";
import { MsgCreateProofs } from "./types/storage/tx";
import { MsgSignContract } from "./types/storage/tx";
import { MsgSetMinerTotalspace } from "./types/storage/tx";
import { MsgUpdateContracts } from "./types/storage/tx";
import { MsgCreateMiners } from "./types/storage/tx";
import { MsgDeleteActiveDeals } from "./types/storage/tx";
import { MsgUpdateMiners } from "./types/storage/tx";
import { MsgItem } from "./types/storage/tx";
import { MsgInitMiner } from "./types/storage/tx";
import { MsgCancelContract } from "./types/storage/tx";
import { MsgPostContract } from "./types/storage/tx";
import { MsgDeleteContracts } from "./types/storage/tx";
import { MsgUpdateProofs } from "./types/storage/tx";
import { MsgDeleteMiners } from "./types/storage/tx";
import { MsgPostproof } from "./types/storage/tx";
import { MsgDeleteProofs } from "./types/storage/tx";
import { MsgUpdateActiveDeals } from "./types/storage/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/jackaldao.canine.storage.MsgSetMinerIp", MsgSetMinerIp],
    ["/jackaldao.canine.storage.MsgCreateActiveDeals", MsgCreateActiveDeals],
    ["/jackaldao.canine.storage.MsgBuyStorage", MsgBuyStorage],
    ["/jackaldao.canine.storage.MsgCreateContracts", MsgCreateContracts],
    ["/jackaldao.canine.storage.MsgCreateProofs", MsgCreateProofs],
    ["/jackaldao.canine.storage.MsgSignContract", MsgSignContract],
    ["/jackaldao.canine.storage.MsgSetMinerTotalspace", MsgSetMinerTotalspace],
    ["/jackaldao.canine.storage.MsgUpdateContracts", MsgUpdateContracts],
    ["/jackaldao.canine.storage.MsgCreateMiners", MsgCreateMiners],
    ["/jackaldao.canine.storage.MsgDeleteActiveDeals", MsgDeleteActiveDeals],
    ["/jackaldao.canine.storage.MsgUpdateMiners", MsgUpdateMiners],
    ["/jackaldao.canine.storage.MsgItem", MsgItem],
    ["/jackaldao.canine.storage.MsgInitMiner", MsgInitMiner],
    ["/jackaldao.canine.storage.MsgCancelContract", MsgCancelContract],
    ["/jackaldao.canine.storage.MsgPostContract", MsgPostContract],
    ["/jackaldao.canine.storage.MsgDeleteContracts", MsgDeleteContracts],
    ["/jackaldao.canine.storage.MsgUpdateProofs", MsgUpdateProofs],
    ["/jackaldao.canine.storage.MsgDeleteMiners", MsgDeleteMiners],
    ["/jackaldao.canine.storage.MsgPostproof", MsgPostproof],
    ["/jackaldao.canine.storage.MsgDeleteProofs", MsgDeleteProofs],
    ["/jackaldao.canine.storage.MsgUpdateActiveDeals", MsgUpdateActiveDeals],
=======
import { MsgDeleteMiners } from "./types/storage/tx";
import { MsgSignContract } from "./types/storage/tx";
import { MsgUpdateContracts } from "./types/storage/tx";
import { MsgDeleteProofs } from "./types/storage/tx";
import { MsgCreateActiveDeals } from "./types/storage/tx";
import { MsgSetMinerIp } from "./types/storage/tx";
import { MsgSetMinerTotalspace } from "./types/storage/tx";
import { MsgCreateMiners } from "./types/storage/tx";
import { MsgPostContract } from "./types/storage/tx";
import { MsgCreateProofs } from "./types/storage/tx";
import { MsgDeleteContracts } from "./types/storage/tx";
import { MsgBuyStorage } from "./types/storage/tx";
import { MsgDeleteActiveDeals } from "./types/storage/tx";
import { MsgUpdateMiners } from "./types/storage/tx";
import { MsgItem } from "./types/storage/tx";
import { MsgPostproof } from "./types/storage/tx";
import { MsgUpdateActiveDeals } from "./types/storage/tx";
import { MsgCancelContract } from "./types/storage/tx";
import { MsgUpdateProofs } from "./types/storage/tx";
import { MsgCreateContracts } from "./types/storage/tx";
import { MsgInitMiner } from "./types/storage/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/jackaldao.canine.storage.MsgDeleteMiners", MsgDeleteMiners],
    ["/jackaldao.canine.storage.MsgSignContract", MsgSignContract],
    ["/jackaldao.canine.storage.MsgUpdateContracts", MsgUpdateContracts],
    ["/jackaldao.canine.storage.MsgDeleteProofs", MsgDeleteProofs],
    ["/jackaldao.canine.storage.MsgCreateActiveDeals", MsgCreateActiveDeals],
    ["/jackaldao.canine.storage.MsgSetMinerIp", MsgSetMinerIp],
    ["/jackaldao.canine.storage.MsgSetMinerTotalspace", MsgSetMinerTotalspace],
    ["/jackaldao.canine.storage.MsgCreateMiners", MsgCreateMiners],
    ["/jackaldao.canine.storage.MsgPostContract", MsgPostContract],
    ["/jackaldao.canine.storage.MsgCreateProofs", MsgCreateProofs],
    ["/jackaldao.canine.storage.MsgDeleteContracts", MsgDeleteContracts],
    ["/jackaldao.canine.storage.MsgBuyStorage", MsgBuyStorage],
    ["/jackaldao.canine.storage.MsgDeleteActiveDeals", MsgDeleteActiveDeals],
    ["/jackaldao.canine.storage.MsgUpdateMiners", MsgUpdateMiners],
    ["/jackaldao.canine.storage.MsgItem", MsgItem],
    ["/jackaldao.canine.storage.MsgPostproof", MsgPostproof],
    ["/jackaldao.canine.storage.MsgUpdateActiveDeals", MsgUpdateActiveDeals],
    ["/jackaldao.canine.storage.MsgCancelContract", MsgCancelContract],
    ["/jackaldao.canine.storage.MsgUpdateProofs", MsgUpdateProofs],
    ["/jackaldao.canine.storage.MsgCreateContracts", MsgCreateContracts],
    ["/jackaldao.canine.storage.MsgInitMiner", MsgInitMiner],
>>>>>>> master
    
];

export { msgTypes }