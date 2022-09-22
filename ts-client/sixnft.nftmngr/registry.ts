import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgAddTokenAttribute } from "./types/nftmngr/tx";
import { MsgPerformActionByAdmin } from "./types/nftmngr/tx";
import { MsgAddAttribute } from "./types/nftmngr/tx";
import { MsgSetNFTAttribute } from "./types/nftmngr/tx";
import { MsgCreateNFTSchema } from "./types/nftmngr/tx";
import { MsgAddAction } from "./types/nftmngr/tx";
import { MsgCreateMetadata } from "./types/nftmngr/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/sixnft.nftmngr.MsgAddTokenAttribute", MsgAddTokenAttribute],
    ["/sixnft.nftmngr.MsgPerformActionByAdmin", MsgPerformActionByAdmin],
    ["/sixnft.nftmngr.MsgAddAttribute", MsgAddAttribute],
    ["/sixnft.nftmngr.MsgSetNFTAttribute", MsgSetNFTAttribute],
    ["/sixnft.nftmngr.MsgCreateNFTSchema", MsgCreateNFTSchema],
    ["/sixnft.nftmngr.MsgAddAction", MsgAddAction],
    ["/sixnft.nftmngr.MsgCreateMetadata", MsgCreateMetadata],
    
];

export { msgTypes }