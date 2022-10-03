import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgSignform } from "./types/dsig/tx";
import { MsgUploadfile } from "./types/dsig/tx";
import { MsgCreateform } from "./types/dsig/tx";
import { MsgUploadfile } from "./types/dsig/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/jackaldao.canine.dsig.MsgSignform", MsgSignform],
    ["/jackaldao.canine.dsig.MsgUploadfile", MsgUploadfile],
    ["/jackaldao.canine.dsig.MsgCreateform", MsgCreateform],
    ["/jackaldao.canine.dsig.MsgUploadfile", MsgUploadfile],
    
];

export { msgTypes }