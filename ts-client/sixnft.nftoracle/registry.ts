import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgCreateActionRequest } from "./types/nftoracle/tx";
import { MsgSubmitMintResponse } from "./types/nftoracle/tx";
import { MsgSubmitActionResponse } from "./types/nftoracle/tx";
import { MsgCreateMintRequest } from "./types/nftoracle/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/sixnft.nftoracle.MsgCreateActionRequest", MsgCreateActionRequest],
    ["/sixnft.nftoracle.MsgSubmitMintResponse", MsgSubmitMintResponse],
    ["/sixnft.nftoracle.MsgSubmitActionResponse", MsgSubmitActionResponse],
    ["/sixnft.nftoracle.MsgCreateMintRequest", MsgCreateMintRequest],
    
];

export { msgTypes }