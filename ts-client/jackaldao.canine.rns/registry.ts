import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgBuy } from "./types/rns/tx";
import { MsgAcceptBid } from "./types/rns/tx";
import { MsgRegister } from "./types/rns/tx";
import { MsgDelist } from "./types/rns/tx";
import { MsgInit } from "./types/rns/tx";
import { MsgTransfer } from "./types/rns/tx";
import { MsgAddRecord } from "./types/rns/tx";
import { MsgBid } from "./types/rns/tx";
import { MsgCancelBid } from "./types/rns/tx";
import { MsgList } from "./types/rns/tx";
import { MsgDelRecord } from "./types/rns/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/jackaldao.canine.rns.MsgBuy", MsgBuy],
    ["/jackaldao.canine.rns.MsgAcceptBid", MsgAcceptBid],
    ["/jackaldao.canine.rns.MsgRegister", MsgRegister],
    ["/jackaldao.canine.rns.MsgDelist", MsgDelist],
    ["/jackaldao.canine.rns.MsgInit", MsgInit],
    ["/jackaldao.canine.rns.MsgTransfer", MsgTransfer],
    ["/jackaldao.canine.rns.MsgAddRecord", MsgAddRecord],
    ["/jackaldao.canine.rns.MsgBid", MsgBid],
    ["/jackaldao.canine.rns.MsgCancelBid", MsgCancelBid],
    ["/jackaldao.canine.rns.MsgList", MsgList],
    ["/jackaldao.canine.rns.MsgDelRecord", MsgDelRecord],
    
];

export { msgTypes }