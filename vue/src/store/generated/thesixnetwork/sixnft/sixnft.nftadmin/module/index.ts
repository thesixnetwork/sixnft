// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.

import { StdFee } from "@cosmjs/launchpad";
import { SigningStargateClient } from "@cosmjs/stargate";
import { Registry, OfflineSigner, EncodeObject, DirectSecp256k1HdWallet } from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgBurn } from "./types/nftadmin/tx";
import { MsgGrantPermission } from "./types/nftadmin/tx";
import { MsgMint } from "./types/nftadmin/tx";
import { MsgRevokePermission } from "./types/nftadmin/tx";


const types = [
  ["/sixnft.nftadmin.MsgBurn", MsgBurn],
  ["/sixnft.nftadmin.MsgGrantPermission", MsgGrantPermission],
  ["/sixnft.nftadmin.MsgMint", MsgMint],
  ["/sixnft.nftadmin.MsgRevokePermission", MsgRevokePermission],
  
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
    msgBurn: (data: MsgBurn): EncodeObject => ({ typeUrl: "/sixnft.nftadmin.MsgBurn", value: MsgBurn.fromPartial( data ) }),
    msgGrantPermission: (data: MsgGrantPermission): EncodeObject => ({ typeUrl: "/sixnft.nftadmin.MsgGrantPermission", value: MsgGrantPermission.fromPartial( data ) }),
    msgMint: (data: MsgMint): EncodeObject => ({ typeUrl: "/sixnft.nftadmin.MsgMint", value: MsgMint.fromPartial( data ) }),
    msgRevokePermission: (data: MsgRevokePermission): EncodeObject => ({ typeUrl: "/sixnft.nftadmin.MsgRevokePermission", value: MsgRevokePermission.fromPartial( data ) }),
    
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