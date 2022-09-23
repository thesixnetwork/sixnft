import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgBindAddress } from "./types/evmsupport/tx";
import { MsgRemoveBinding } from "./types/evmsupport/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/sixnft.evmsupport.MsgBindAddress", MsgBindAddress],
    ["/sixnft.evmsupport.MsgRemoveBinding", MsgRemoveBinding],
    
];

export { msgTypes }