import { GeneratedType } from "@cosmjs/proto-signing";
<<<<<<< HEAD
import { MsgInitAccount } from "./types/filetree/tx";
import { MsgDeleteFile } from "./types/filetree/tx";
import { MsgPostkey } from "./types/filetree/tx";
=======
import { MsgInitAll } from "./types/filetree/tx";
import { MsgPostkey } from "./types/filetree/tx";
import { MsgDeleteFile } from "./types/filetree/tx";
import { MsgInitAccount } from "./types/filetree/tx";
>>>>>>> master
import { MsgPostFile } from "./types/filetree/tx";
import { MsgAddViewers } from "./types/filetree/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
<<<<<<< HEAD
    ["/jackaldao.canine.filetree.MsgInitAccount", MsgInitAccount],
    ["/jackaldao.canine.filetree.MsgDeleteFile", MsgDeleteFile],
    ["/jackaldao.canine.filetree.MsgPostkey", MsgPostkey],
=======
    ["/jackaldao.canine.filetree.MsgInitAll", MsgInitAll],
    ["/jackaldao.canine.filetree.MsgPostkey", MsgPostkey],
    ["/jackaldao.canine.filetree.MsgDeleteFile", MsgDeleteFile],
    ["/jackaldao.canine.filetree.MsgInitAccount", MsgInitAccount],
>>>>>>> master
    ["/jackaldao.canine.filetree.MsgPostFile", MsgPostFile],
    ["/jackaldao.canine.filetree.MsgAddViewers", MsgAddViewers],
    
];

export { msgTypes }