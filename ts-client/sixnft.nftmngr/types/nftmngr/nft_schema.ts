/* eslint-disable */
import { OriginData } from "../nftmngr/origin_data";
import { OnChainData } from "../nftmngr/on_chain_data";
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "sixnft.nftmngr";

export interface NFTSchema {
  code: string;
  name: string;
  owner: string;
  originData: OriginData | undefined;
  onchainData: OnChainData | undefined;
}

const baseNFTSchema: object = { code: "", name: "", owner: "" };

export const NFTSchema = {
  encode(message: NFTSchema, writer: Writer = Writer.create()): Writer {
    if (message.code !== "") {
      writer.uint32(10).string(message.code);
    }
    if (message.name !== "") {
      writer.uint32(18).string(message.name);
    }
    if (message.owner !== "") {
      writer.uint32(26).string(message.owner);
    }
    if (message.originData !== undefined) {
      OriginData.encode(message.originData, writer.uint32(34).fork()).ldelim();
    }
    if (message.onchainData !== undefined) {
      OnChainData.encode(
        message.onchainData,
        writer.uint32(42).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): NFTSchema {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseNFTSchema } as NFTSchema;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.code = reader.string();
          break;
        case 2:
          message.name = reader.string();
          break;
        case 3:
          message.owner = reader.string();
          break;
        case 4:
          message.originData = OriginData.decode(reader, reader.uint32());
          break;
        case 5:
          message.onchainData = OnChainData.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): NFTSchema {
    const message = { ...baseNFTSchema } as NFTSchema;
    if (object.code !== undefined && object.code !== null) {
      message.code = String(object.code);
    } else {
      message.code = "";
    }
    if (object.name !== undefined && object.name !== null) {
      message.name = String(object.name);
    } else {
      message.name = "";
    }
    if (object.owner !== undefined && object.owner !== null) {
      message.owner = String(object.owner);
    } else {
      message.owner = "";
    }
    if (object.originData !== undefined && object.originData !== null) {
      message.originData = OriginData.fromJSON(object.originData);
    } else {
      message.originData = undefined;
    }
    if (object.onchainData !== undefined && object.onchainData !== null) {
      message.onchainData = OnChainData.fromJSON(object.onchainData);
    } else {
      message.onchainData = undefined;
    }
    return message;
  },

  toJSON(message: NFTSchema): unknown {
    const obj: any = {};
    message.code !== undefined && (obj.code = message.code);
    message.name !== undefined && (obj.name = message.name);
    message.owner !== undefined && (obj.owner = message.owner);
    message.originData !== undefined &&
      (obj.originData = message.originData
        ? OriginData.toJSON(message.originData)
        : undefined);
    message.onchainData !== undefined &&
      (obj.onchainData = message.onchainData
        ? OnChainData.toJSON(message.onchainData)
        : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<NFTSchema>): NFTSchema {
    const message = { ...baseNFTSchema } as NFTSchema;
    if (object.code !== undefined && object.code !== null) {
      message.code = object.code;
    } else {
      message.code = "";
    }
    if (object.name !== undefined && object.name !== null) {
      message.name = object.name;
    } else {
      message.name = "";
    }
    if (object.owner !== undefined && object.owner !== null) {
      message.owner = object.owner;
    } else {
      message.owner = "";
    }
    if (object.originData !== undefined && object.originData !== null) {
      message.originData = OriginData.fromPartial(object.originData);
    } else {
      message.originData = undefined;
    }
    if (object.onchainData !== undefined && object.onchainData !== null) {
      message.onchainData = OnChainData.fromPartial(object.onchainData);
    } else {
      message.onchainData = undefined;
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
