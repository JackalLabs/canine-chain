import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgPostFile } from "./types/filetree/tx";
import { MsgPostkey } from "./types/filetree/tx";
import { MsgAddViewers } from "./types/filetree/tx";
import { MsgInitAccount } from "./types/filetree/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/jackaldao.canine.filetree.MsgPostFile", MsgPostFile],
    ["/jackaldao.canine.filetree.MsgPostkey", MsgPostkey],
    ["/jackaldao.canine.filetree.MsgAddViewers", MsgAddViewers],
    ["/jackaldao.canine.filetree.MsgInitAccount", MsgInitAccount],
    
];

export { msgTypes }