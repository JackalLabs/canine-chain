import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgSwap } from "./types/lp/tx";
import { MsgWithdrawLPool } from "./types/lp/tx";
import { MsgCreateLPool } from "./types/lp/tx";
import { MsgDepositLPool } from "./types/lp/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/jackaldao.canine.lp.MsgSwap", MsgSwap],
    ["/jackaldao.canine.lp.MsgWithdrawLPool", MsgWithdrawLPool],
    ["/jackaldao.canine.lp.MsgCreateLPool", MsgCreateLPool],
    ["/jackaldao.canine.lp.MsgDepositLPool", MsgDepositLPool],
    
];

export { msgTypes }