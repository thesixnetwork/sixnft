/* eslint-disable */
import { NftAttributeValue } from "../nftmngr/nft_attribute_value";
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "sixnft.nftmngr";

export interface NftData {
  nft_schema_code: string;
  token_id: string;
  token_owner: string;
  origin_image: string;
  onchain_image: string;
  origin_attributes: NftAttributeValue[];
  onchain_attributes: NftAttributeValue[];
}

const baseNftData: object = {
  nft_schema_code: "",
  token_id: "",
  token_owner: "",
  origin_image: "",
  onchain_image: "",
};

export const NftData = {
  encode(message: NftData, writer: Writer = Writer.create()): Writer {
    if (message.nft_schema_code !== "") {
      writer.uint32(10).string(message.nft_schema_code);
    }
    if (message.token_id !== "") {
      writer.uint32(18).string(message.token_id);
    }
    if (message.token_owner !== "") {
      writer.uint32(26).string(message.token_owner);
    }
    if (message.origin_image !== "") {
      writer.uint32(34).string(message.origin_image);
    }
    if (message.onchain_image !== "") {
      writer.uint32(42).string(message.onchain_image);
    }
    for (const v of message.origin_attributes) {
      NftAttributeValue.encode(v!, writer.uint32(50).fork()).ldelim();
    }
    for (const v of message.onchain_attributes) {
      NftAttributeValue.encode(v!, writer.uint32(58).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): NftData {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseNftData } as NftData;
    message.origin_attributes = [];
    message.onchain_attributes = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.nft_schema_code = reader.string();
          break;
        case 2:
          message.token_id = reader.string();
          break;
        case 3:
          message.token_owner = reader.string();
          break;
        case 4:
          message.origin_image = reader.string();
          break;
        case 5:
          message.onchain_image = reader.string();
          break;
        case 6:
          message.origin_attributes.push(
            NftAttributeValue.decode(reader, reader.uint32())
          );
          break;
        case 7:
          message.onchain_attributes.push(
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
    message.origin_attributes = [];
    message.onchain_attributes = [];
    if (
      object.nft_schema_code !== undefined &&
      object.nft_schema_code !== null
    ) {
      message.nft_schema_code = String(object.nft_schema_code);
    } else {
      message.nft_schema_code = "";
    }
    if (object.token_id !== undefined && object.token_id !== null) {
      message.token_id = String(object.token_id);
    } else {
      message.token_id = "";
    }
    if (object.token_owner !== undefined && object.token_owner !== null) {
      message.token_owner = String(object.token_owner);
    } else {
      message.token_owner = "";
    }
    if (object.origin_image !== undefined && object.origin_image !== null) {
      message.origin_image = String(object.origin_image);
    } else {
      message.origin_image = "";
    }
    if (object.onchain_image !== undefined && object.onchain_image !== null) {
      message.onchain_image = String(object.onchain_image);
    } else {
      message.onchain_image = "";
    }
    if (
      object.origin_attributes !== undefined &&
      object.origin_attributes !== null
    ) {
      for (const e of object.origin_attributes) {
        message.origin_attributes.push(NftAttributeValue.fromJSON(e));
      }
    }
    if (
      object.onchain_attributes !== undefined &&
      object.onchain_attributes !== null
    ) {
      for (const e of object.onchain_attributes) {
        message.onchain_attributes.push(NftAttributeValue.fromJSON(e));
      }
    }
    return message;
  },

  toJSON(message: NftData): unknown {
    const obj: any = {};
    message.nft_schema_code !== undefined &&
      (obj.nft_schema_code = message.nft_schema_code);
    message.token_id !== undefined && (obj.token_id = message.token_id);
    message.token_owner !== undefined &&
      (obj.token_owner = message.token_owner);
    message.origin_image !== undefined &&
      (obj.origin_image = message.origin_image);
    message.onchain_image !== undefined &&
      (obj.onchain_image = message.onchain_image);
    if (message.origin_attributes) {
      obj.origin_attributes = message.origin_attributes.map((e) =>
        e ? NftAttributeValue.toJSON(e) : undefined
      );
    } else {
      obj.origin_attributes = [];
    }
    if (message.onchain_attributes) {
      obj.onchain_attributes = message.onchain_attributes.map((e) =>
        e ? NftAttributeValue.toJSON(e) : undefined
      );
    } else {
      obj.onchain_attributes = [];
    }
    return obj;
  },

  fromPartial(object: DeepPartial<NftData>): NftData {
    const message = { ...baseNftData } as NftData;
    message.origin_attributes = [];
    message.onchain_attributes = [];
    if (
      object.nft_schema_code !== undefined &&
      object.nft_schema_code !== null
    ) {
      message.nft_schema_code = object.nft_schema_code;
    } else {
      message.nft_schema_code = "";
    }
    if (object.token_id !== undefined && object.token_id !== null) {
      message.token_id = object.token_id;
    } else {
      message.token_id = "";
    }
    if (object.token_owner !== undefined && object.token_owner !== null) {
      message.token_owner = object.token_owner;
    } else {
      message.token_owner = "";
    }
    if (object.origin_image !== undefined && object.origin_image !== null) {
      message.origin_image = object.origin_image;
    } else {
      message.origin_image = "";
    }
    if (object.onchain_image !== undefined && object.onchain_image !== null) {
      message.onchain_image = object.onchain_image;
    } else {
      message.onchain_image = "";
    }
    if (
      object.origin_attributes !== undefined &&
      object.origin_attributes !== null
    ) {
      for (const e of object.origin_attributes) {
        message.origin_attributes.push(NftAttributeValue.fromPartial(e));
      }
    }
    if (
      object.onchain_attributes !== undefined &&
      object.onchain_attributes !== null
    ) {
      for (const e of object.onchain_attributes) {
        message.onchain_attributes.push(NftAttributeValue.fromPartial(e));
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
