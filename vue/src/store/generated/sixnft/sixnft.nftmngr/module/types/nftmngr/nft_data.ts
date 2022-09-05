/* eslint-disable */
import { NftAttributeValue } from "../nftmngr/nft_attribute_value";
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "sixnft.nftmngr";

export interface NftData {
  nftSchemaCode: string;
  tokenId: string;
  tokenOwner: string;
  originImage: string;
  originAttributes: NftAttributeValue[];
  onchainAttributes: NftAttributeValue[];
}

const baseNftData: object = {
  nftSchemaCode: "",
  tokenId: "",
  tokenOwner: "",
  originImage: "",
};

export const NftData = {
  encode(message: NftData, writer: Writer = Writer.create()): Writer {
    if (message.nftSchemaCode !== "") {
      writer.uint32(10).string(message.nftSchemaCode);
    }
    if (message.tokenId !== "") {
      writer.uint32(18).string(message.tokenId);
    }
    if (message.tokenOwner !== "") {
      writer.uint32(26).string(message.tokenOwner);
    }
    if (message.originImage !== "") {
      writer.uint32(34).string(message.originImage);
    }
    for (const v of message.originAttributes) {
      NftAttributeValue.encode(v!, writer.uint32(42).fork()).ldelim();
    }
    for (const v of message.onchainAttributes) {
      NftAttributeValue.encode(v!, writer.uint32(50).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): NftData {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseNftData } as NftData;
    message.originAttributes = [];
    message.onchainAttributes = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.nftSchemaCode = reader.string();
          break;
        case 2:
          message.tokenId = reader.string();
          break;
        case 3:
          message.tokenOwner = reader.string();
          break;
        case 4:
          message.originImage = reader.string();
          break;
        case 5:
          message.originAttributes.push(
            NftAttributeValue.decode(reader, reader.uint32())
          );
          break;
        case 6:
          message.onchainAttributes.push(
            NftAttributeValue.decode(reader, reader.uint32())
          );
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): NftData {
    const message = { ...baseNftData } as NftData;
    message.originAttributes = [];
    message.onchainAttributes = [];
    if (object.nftSchemaCode !== undefined && object.nftSchemaCode !== null) {
      message.nftSchemaCode = String(object.nftSchemaCode);
    } else {
      message.nftSchemaCode = "";
    }
    if (object.tokenId !== undefined && object.tokenId !== null) {
      message.tokenId = String(object.tokenId);
    } else {
      message.tokenId = "";
    }
    if (object.tokenOwner !== undefined && object.tokenOwner !== null) {
      message.tokenOwner = String(object.tokenOwner);
    } else {
      message.tokenOwner = "";
    }
    if (object.originImage !== undefined && object.originImage !== null) {
      message.originImage = String(object.originImage);
    } else {
      message.originImage = "";
    }
    if (
      object.originAttributes !== undefined &&
      object.originAttributes !== null
    ) {
      for (const e of object.originAttributes) {
        message.originAttributes.push(NftAttributeValue.fromJSON(e));
      }
    }
    if (
      object.onchainAttributes !== undefined &&
      object.onchainAttributes !== null
    ) {
      for (const e of object.onchainAttributes) {
        message.onchainAttributes.push(NftAttributeValue.fromJSON(e));
      }
    }
    return message;
  },

  toJSON(message: NftData): unknown {
    const obj: any = {};
    message.nftSchemaCode !== undefined &&
      (obj.nftSchemaCode = message.nftSchemaCode);
    message.tokenId !== undefined && (obj.tokenId = message.tokenId);
    message.tokenOwner !== undefined && (obj.tokenOwner = message.tokenOwner);
    message.originImage !== undefined &&
      (obj.originImage = message.originImage);
    if (message.originAttributes) {
      obj.originAttributes = message.originAttributes.map((e) =>
        e ? NftAttributeValue.toJSON(e) : undefined
      );
    } else {
      obj.originAttributes = [];
    }
    if (message.onchainAttributes) {
      obj.onchainAttributes = message.onchainAttributes.map((e) =>
        e ? NftAttributeValue.toJSON(e) : undefined
      );
    } else {
      obj.onchainAttributes = [];
    }
    return obj;
  },

  fromPartial(object: DeepPartial<NftData>): NftData {
    const message = { ...baseNftData } as NftData;
    message.originAttributes = [];
    message.onchainAttributes = [];
    if (object.nftSchemaCode !== undefined && object.nftSchemaCode !== null) {
      message.nftSchemaCode = object.nftSchemaCode;
    } else {
      message.nftSchemaCode = "";
    }
    if (object.tokenId !== undefined && object.tokenId !== null) {
      message.tokenId = object.tokenId;
    } else {
      message.tokenId = "";
    }
    if (object.tokenOwner !== undefined && object.tokenOwner !== null) {
      message.tokenOwner = object.tokenOwner;
    } else {
      message.tokenOwner = "";
    }
    if (object.originImage !== undefined && object.originImage !== null) {
      message.originImage = object.originImage;
    } else {
      message.originImage = "";
    }
    if (
      object.originAttributes !== undefined &&
      object.originAttributes !== null
    ) {
      for (const e of object.originAttributes) {
        message.originAttributes.push(NftAttributeValue.fromPartial(e));
      }
    }
    if (
      object.onchainAttributes !== undefined &&
      object.onchainAttributes !== null
    ) {
      for (const e of object.onchainAttributes) {
        message.onchainAttributes.push(NftAttributeValue.fromPartial(e));
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
