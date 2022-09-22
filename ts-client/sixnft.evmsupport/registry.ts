import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgRemoveBinding } from "./types/evmsupport/tx";
import { MsgBindAddress } from "./types/evmsupport/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/sixnft.evmsupport.MsgRemoveBinding", MsgRemoveBinding],
    ["/sixnft.evmsupport.MsgBindAddress", MsgBindAddress],
    
];

export { msgTypes }