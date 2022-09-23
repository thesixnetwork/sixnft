import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgGrantPermission } from "./types/admin/tx";
import { MsgBurn } from "./types/admin/tx";
import { MsgMint } from "./types/admin/tx";
import { MsgRevokePermission } from "./types/admin/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/sixnft.admin.MsgGrantPermission", MsgGrantPermission],
    ["/sixnft.admin.MsgBurn", MsgBurn],
    ["/sixnft.admin.MsgMint", MsgMint],
    ["/sixnft.admin.MsgRevokePermission", MsgRevokePermission],
    
];

export { msgTypes }