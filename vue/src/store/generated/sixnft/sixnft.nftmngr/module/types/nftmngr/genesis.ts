/* eslint-disable */
import { Params } from "../nftmngr/params";
import { NFTSchema } from "../nftmngr/nft_schema";
import { NftData } from "../nftmngr/nft_data";
import { ActionByRefId } from "../nftmngr/action_by_ref_id";
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "sixnft.nftmngr";

/** GenesisState defines the nftmngr module's genesis state. */
export interface GenesisState {
  params: Params | undefined;
  nFTSchemaList: NFTSchema[];
  nftDataList: NftData[];
  /** this line is used by starport scaffolding # genesis/proto/state */
  actionByRefIdList: ActionByRefId[];
}

const baseGenesisState: object = {};

export const GenesisState = {
  encode(message: GenesisState, writer: Writer = Writer.create()): Writer {
    if (message.params !== undefined) {
      Params.encode(message.params, writer.uint32(10).fork()).ldelim();
    }
    for (const v of message.nFTSchemaList) {
      NFTSchema.encode(v!, writer.uint32(18).fork()).ldelim();
    }
    for (const v of message.nftDataList) {
      NftData.encode(v!, writer.uint32(26).fork()).ldelim();
    }
    for (const v of message.actionByRefIdList) {
      ActionByRefId.encode(v!, writer.uint32(34).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): GenesisState {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseGenesisState } as GenesisState;
    message.nFTSchemaList = [];
    message.nftDataList = [];
    message.actionByRefIdList = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.params = Params.decode(reader, reader.uint32());
          break;
        case 2:
          message.nFTSchemaList.push(NFTSchema.decode(reader, reader.uint32()));
          break;
        case 3:
          message.nftDataList.push(NftData.decode(reader, reader.uint32()));
          break;
        case 4:
          message.actionByRefIdList.push(
            ActionByRefId.decode(reader, reader.uint32())
          );
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): GenesisState {
    const message = { ...baseGenesisState } as GenesisState;
    message.nFTSchemaList = [];
    message.nftDataList = [];
    message.actionByRefIdList = [];
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromJSON(object.params);
    } else {
      message.params = undefined;
    }
    if (object.nFTSchemaList !== undefined && object.nFTSchemaList !== null) {
      for (const e of object.nFTSchemaList) {
        message.nFTSchemaList.push(NFTSchema.fromJSON(e));
      }
    }
    if (object.nftDataList !== undefined && object.nftDataList !== null) {
      for (const e of object.nftDataList) {
        message.nftDataList.push(NftData.fromJSON(e));
      }
    }
    if (
      object.actionByRefIdList !== undefined &&
      object.actionByRefIdList !== null
    ) {
      for (const e of object.actionByRefIdList) {
        message.actionByRefIdList.push(ActionByRefId.fromJSON(e));
      }
    }
    return message;
  },

  toJSON(message: GenesisState): unknown {
    const obj: any = {};
    message.params !== undefined &&
      (obj.params = message.params ? Params.toJSON(message.params) : undefined);
    if (message.nFTSchemaList) {
      obj.nFTSchemaList = message.nFTSchemaList.map((e) =>
        e ? NFTSchema.toJSON(e) : undefined
      );
    } else {
      obj.nFTSchemaList = [];
    }
    if (message.nftDataList) {
      obj.nftDataList = message.nftDataList.map((e) =>
        e ? NftData.toJSON(e) : undefined
      );
    } else {
      obj.nftDataList = [];
    }
    if (message.actionByRefIdList) {
      obj.actionByRefIdList = message.actionByRefIdList.map((e) =>
        e ? ActionByRefId.toJSON(e) : undefined
      );
    } else {
      obj.actionByRefIdList = [];
    }
    return obj;
  },

  fromPartial(object: DeepPartial<GenesisState>): GenesisState {
    const message = { ...baseGenesisState } as GenesisState;
    message.nFTSchemaList = [];
    message.nftDataList = [];
    message.actionByRefIdList = [];
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromPartial(object.params);
    } else {
      message.params = undefined;
    }
    if (object.nFTSchemaList !== undefined && object.nFTSchemaList !== null) {
      for (const e of object.nFTSchemaList) {
        message.nFTSchemaList.push(NFTSchema.fromPartial(e));
      }
    }
    if (object.nftDataList !== undefined && object.nftDataList !== null) {
      for (const e of object.nftDataList) {
        message.nftDataList.push(NftData.fromPartial(e));
      }
    }
    if (
      object.actionByRefIdList !== undefined &&
      object.actionByRefIdList !== null
    ) {
      for (const e of object.actionByRefIdList) {
        message.actionByRefIdList.push(ActionByRefId.fromPartial(e));
      }
    }
    return message;
  },
};

type Builtin = Date | Function | Uint8Array | string | number | undefined;
export type DeepPartial<T> = T extends Builtin
  ? T
  : T extends Array<infer U>
  ? Array<DeepPartial<U>>
  : T extends ReadonlyArray<infer U>
  ? ReadonlyArray<DeepPartial<U>>
  : T extends {}
  ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;
