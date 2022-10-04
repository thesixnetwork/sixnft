// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.

import { StdFee } from "@cosmjs/launchpad";
import { SigningStargateClient } from "@cosmjs/stargate";
import { Registry, OfflineSigner, EncodeObject, DirectSecp256k1HdWallet } from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgAddTokenAttribute } from "./types/nftmngr/tx";
import { MsgSetNFTAttribute } from "./types/nftmngr/tx";
import { MsgChangeSchemaOwner } from "./types/nftmngr/tx";
import { MsgAddAction } from "./types/nftmngr/tx";
import { MsgPerformActionByAdmin } from "./types/nftmngr/tx";
import { MsgCreateMetadata } from "./types/nftmngr/tx";
import { MsgAddSystemActioner } from "./types/nftmngr/tx";
import { MsgRemoveSystemActioner } from "./types/nftmngr/tx";
import { MsgSetBaseUri } from "./types/nftmngr/tx";
import { MsgToggleAction } from "./types/nftmngr/tx";
import { MsgAddAttribute } from "./types/nftmngr/tx";
import { MsgCreateNFTSchema } from "./types/nftmngr/tx";


const types = [
  ["/sixnft.nftmngr.MsgAddTokenAttribute", MsgAddTokenAttribute],
  ["/sixnft.nftmngr.MsgSetNFTAttribute", MsgSetNFTAttribute],
  ["/sixnft.nftmngr.MsgChangeSchemaOwner", MsgChangeSchemaOwner],
  ["/sixnft.nftmngr.MsgAddAction", MsgAddAction],
  ["/sixnft.nftmngr.MsgPerformActionByAdmin", MsgPerformActionByAdmin],
  ["/sixnft.nftmngr.MsgCreateMetadata", MsgCreateMetadata],
  ["/sixnft.nftmngr.MsgAddSystemActioner", MsgAddSystemActioner],
  ["/sixnft.nftmngr.MsgRemoveSystemActioner", MsgRemoveSystemActioner],
  ["/sixnft.nftmngr.MsgSetBaseUri", MsgSetBaseUri],
  ["/sixnft.nftmngr.MsgToggleAction", MsgToggleAction],
  ["/sixnft.nftmngr.MsgAddAttribute", MsgAddAttribute],
  ["/sixnft.nftmngr.MsgCreateNFTSchema", MsgCreateNFTSchema],
  
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
    msgAddTokenAttribute: (data: MsgAddTokenAttribute): EncodeObject => ({ typeUrl: "/sixnft.nftmngr.MsgAddTokenAttribute", value: MsgAddTokenAttribute.fromPartial( data ) }),
    msgSetNFTAttribute: (data: MsgSetNFTAttribute): EncodeObject => ({ typeUrl: "/sixnft.nftmngr.MsgSetNFTAttribute", value: MsgSetNFTAttribute.fromPartial( data ) }),
    msgChangeSchemaOwner: (data: MsgChangeSchemaOwner): EncodeObject => ({ typeUrl: "/sixnft.nftmngr.MsgChangeSchemaOwner", value: MsgChangeSchemaOwner.fromPartial( data ) }),
    msgAddAction: (data: MsgAddAction): EncodeObject => ({ typeUrl: "/sixnft.nftmngr.MsgAddAction", value: MsgAddAction.fromPartial( data ) }),
    msgPerformActionByAdmin: (data: MsgPerformActionByAdmin): EncodeObject => ({ typeUrl: "/sixnft.nftmngr.MsgPerformActionByAdmin", value: MsgPerformActionByAdmin.fromPartial( data ) }),
    msgCreateMetadata: (data: MsgCreateMetadata): EncodeObject => ({ typeUrl: "/sixnft.nftmngr.MsgCreateMetadata", value: MsgCreateMetadata.fromPartial( data ) }),
    msgAddSystemActioner: (data: MsgAddSystemActioner): EncodeObject => ({ typeUrl: "/sixnft.nftmngr.MsgAddSystemActioner", value: MsgAddSystemActioner.fromPartial( data ) }),
    msgRemoveSystemActioner: (data: MsgRemoveSystemActioner): EncodeObject => ({ typeUrl: "/sixnft.nftmngr.MsgRemoveSystemActioner", value: MsgRemoveSystemActioner.fromPartial( data ) }),
    msgSetBaseUri: (data: MsgSetBaseUri): EncodeObject => ({ typeUrl: "/sixnft.nftmngr.MsgSetBaseUri", value: MsgSetBaseUri.fromPartial( data ) }),
    msgToggleAction: (data: MsgToggleAction): EncodeObject => ({ typeUrl: "/sixnft.nftmngr.MsgToggleAction", value: MsgToggleAction.fromPartial( data ) }),
    msgAddAttribute: (data: MsgAddAttribute): EncodeObject => ({ typeUrl: "/sixnft.nftmngr.MsgAddAttribute", value: MsgAddAttribute.fromPartial( data ) }),
    msgCreateNFTSchema: (data: MsgCreateNFTSchema): EncodeObject => ({ typeUrl: "/sixnft.nftmngr.MsgCreateNFTSchema", value: MsgCreateNFTSchema.fromPartial( data ) }),
    
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
