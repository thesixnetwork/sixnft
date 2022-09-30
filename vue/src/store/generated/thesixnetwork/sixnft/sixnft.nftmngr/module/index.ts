// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.

import { StdFee } from "@cosmjs/launchpad";
import { SigningStargateClient } from "@cosmjs/stargate";
import { Registry, OfflineSigner, EncodeObject, DirectSecp256k1HdWallet } from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgPerformActionByAdmin } from "./types/nftmngr/tx";
import { MsgChangeSchemaOwner } from "./types/nftmngr/tx";
import { MsgAddAction } from "./types/nftmngr/tx";
import { MsgSetBaseUri } from "./types/nftmngr/tx";
import { MsgCreateMetadata } from "./types/nftmngr/tx";
import { MsgCreateNFTSchema } from "./types/nftmngr/tx";
import { MsgAddTokenAttribute } from "./types/nftmngr/tx";
import { MsgAddAttribute } from "./types/nftmngr/tx";
import { MsgSetNFTAttribute } from "./types/nftmngr/tx";
import { MsgToggleAction } from "./types/nftmngr/tx";


const types = [
  ["/sixnft.nftmngr.MsgPerformActionByAdmin", MsgPerformActionByAdmin],
  ["/sixnft.nftmngr.MsgChangeSchemaOwner", MsgChangeSchemaOwner],
  ["/sixnft.nftmngr.MsgAddAction", MsgAddAction],
  ["/sixnft.nftmngr.MsgSetBaseUri", MsgSetBaseUri],
  ["/sixnft.nftmngr.MsgCreateMetadata", MsgCreateMetadata],
  ["/sixnft.nftmngr.MsgCreateNFTSchema", MsgCreateNFTSchema],
  ["/sixnft.nftmngr.MsgAddTokenAttribute", MsgAddTokenAttribute],
  ["/sixnft.nftmngr.MsgAddAttribute", MsgAddAttribute],
  ["/sixnft.nftmngr.MsgSetNFTAttribute", MsgSetNFTAttribute],
  ["/sixnft.nftmngr.MsgToggleAction", MsgToggleAction],
  
];
export const MissingWalletError = new Error("wallet is required");

export const registry = new Registry(<any>types);

const defaultFee = {
  amount: [],
  gas: "200000",
};

interface TxClientOptions {
  addr: string
}

interface SignAndBroadcastOptions {
  fee: StdFee,
  memo?: string
}

const txClient = async (wallet: OfflineSigner, { addr: addr }: TxClientOptions = { addr: "http://localhost:26657" }) => {
  if (!wallet) throw MissingWalletError;
  let client;
  if (addr) {
    client = await SigningStargateClient.connectWithSigner(addr, wallet, { registry });
  }else{
    client = await SigningStargateClient.offline( wallet, { registry });
  }
  const { address } = (await wallet.getAccounts())[0];

  return {
    signAndBroadcast: (msgs: EncodeObject[], { fee, memo }: SignAndBroadcastOptions = {fee: defaultFee, memo: ""}) => client.signAndBroadcast(address, msgs, fee,memo),
    msgPerformActionByAdmin: (data: MsgPerformActionByAdmin): EncodeObject => ({ typeUrl: "/sixnft.nftmngr.MsgPerformActionByAdmin", value: MsgPerformActionByAdmin.fromPartial( data ) }),
    msgChangeSchemaOwner: (data: MsgChangeSchemaOwner): EncodeObject => ({ typeUrl: "/sixnft.nftmngr.MsgChangeSchemaOwner", value: MsgChangeSchemaOwner.fromPartial( data ) }),
    msgAddAction: (data: MsgAddAction): EncodeObject => ({ typeUrl: "/sixnft.nftmngr.MsgAddAction", value: MsgAddAction.fromPartial( data ) }),
    msgSetBaseUri: (data: MsgSetBaseUri): EncodeObject => ({ typeUrl: "/sixnft.nftmngr.MsgSetBaseUri", value: MsgSetBaseUri.fromPartial( data ) }),
    msgCreateMetadata: (data: MsgCreateMetadata): EncodeObject => ({ typeUrl: "/sixnft.nftmngr.MsgCreateMetadata", value: MsgCreateMetadata.fromPartial( data ) }),
    msgCreateNFTSchema: (data: MsgCreateNFTSchema): EncodeObject => ({ typeUrl: "/sixnft.nftmngr.MsgCreateNFTSchema", value: MsgCreateNFTSchema.fromPartial( data ) }),
    msgAddTokenAttribute: (data: MsgAddTokenAttribute): EncodeObject => ({ typeUrl: "/sixnft.nftmngr.MsgAddTokenAttribute", value: MsgAddTokenAttribute.fromPartial( data ) }),
    msgAddAttribute: (data: MsgAddAttribute): EncodeObject => ({ typeUrl: "/sixnft.nftmngr.MsgAddAttribute", value: MsgAddAttribute.fromPartial( data ) }),
    msgSetNFTAttribute: (data: MsgSetNFTAttribute): EncodeObject => ({ typeUrl: "/sixnft.nftmngr.MsgSetNFTAttribute", value: MsgSetNFTAttribute.fromPartial( data ) }),
    msgToggleAction: (data: MsgToggleAction): EncodeObject => ({ typeUrl: "/sixnft.nftmngr.MsgToggleAction", value: MsgToggleAction.fromPartial( data ) }),
    
  };
};

interface QueryClientOptions {
  addr: string
}

const queryClient = async ({ addr: addr }: QueryClientOptions = { addr: "http://localhost:1317" }) => {
  return new Api({ baseUrl: addr });
};

export {
  txClient,
  queryClient,
};
