import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgSetCounter } from "./types/notifications/tx";
import { MsgDeleteNotifications } from "./types/notifications/tx";
import { MsgCreateNotifications } from "./types/notifications/tx";
import { MsgUpdateNotifications } from "./types/notifications/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/jackaldao.canine.notifications.MsgSetCounter", MsgSetCounter],
    ["/jackaldao.canine.notifications.MsgDeleteNotifications", MsgDeleteNotifications],
    ["/jackaldao.canine.notifications.MsgCreateNotifications", MsgCreateNotifications],
    ["/jackaldao.canine.notifications.MsgUpdateNotifications", MsgUpdateNotifications],
    
];

export { msgTypes }