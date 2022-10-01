import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgCreateNotifications } from "./types/notifications/tx";
import { MsgUpdateNotifications } from "./types/notifications/tx";
import { MsgDeleteNotifications } from "./types/notifications/tx";
import { MsgSetCounter } from "./types/notifications/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/jackaldao.canine.notifications.MsgCreateNotifications", MsgCreateNotifications],
    ["/jackaldao.canine.notifications.MsgUpdateNotifications", MsgUpdateNotifications],
    ["/jackaldao.canine.notifications.MsgDeleteNotifications", MsgDeleteNotifications],
    ["/jackaldao.canine.notifications.MsgSetCounter", MsgSetCounter],
    
];

export { msgTypes }