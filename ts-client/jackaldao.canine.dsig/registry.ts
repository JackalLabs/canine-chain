import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgUploadfile } from "./types/dsig/tx";
import { MsgCreateform } from "./types/dsig/tx";
import { MsgSignform } from "./types/dsig/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/jackaldao.canine.dsig.MsgUploadfile", MsgUploadfile],
    ["/jackaldao.canine.dsig.MsgCreateform", MsgCreateform],
    ["/jackaldao.canine.dsig.MsgSignform", MsgSignform],
    
];

export { msgTypes }