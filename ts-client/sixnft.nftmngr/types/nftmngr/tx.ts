/* eslint-disable */
import { Reader, Writer } from "protobufjs/minimal";
import { Any } from "../google/protobuf/any";
import { OnChainData } from "../nftmngr/on_chain_data";

export const protobufPackage = "sixnft.nftmngr";

export interface MsgCreateNFTSchema {
  creator: string;
  nftSchemaBase64: string;
}

export interface MsgCreateNFTSchemaResponse {
  code: string;
}

export interface MsgCreateMetadata {
  creator: string;
  nftSchemaCode: string;
  tokenId: string;
  base64NFTData: string;
}

export interface MsgCreateMetadataResponse {
  nftSchemaCode: string;
  tokenId: string;
}

export interface OpenseaAttribute {
  traitType: string;
  value: Any | undefined;
}

export interface UpdatedOpenseaAttributes {
  attributes: OpenseaAttribute[];
}

export interface UpdatedOriginData {
  opensea: UpdatedOpenseaAttributes | undefined;
}

export interface MsgPerformActionByAdmin {
  creator: string;
  nftSchemaCode: string;
  tokenId: string;
  action: string;
  refId: string;
}

export interface MsgPerformActionByAdminResponse {
  nftSchemaCode: string;
  tokenId: string;
}

export interface MsgAddAttribute {
  creator: string;
  code: string;
  base64NewAttriuteDefenition: string;
}

export interface MsgAddAttributeResponse {
  code: string;
  name: string;
  onchainData: OnChainData | undefined;
}

export interface MsgAddTokenAttribute {
  creator: string;
  code: string;
  base64NewTokenAttriute: string;
}

export interface MsgAddTokenAttributeResponse {
  code: string;
  name: string;
  onchainData: OnChainData | undefined;
}

export interface MsgAddAction {
  creator: string;
  code: string;
  base64NewAction: string;
}

export interface MsgAddActionResponse {
  code: string;
  name: string;
  onchainData: OnChainData | undefined;
}

export interface MsgSetNFTAttribute {
  creator: string;
  nftSchemaCode: string;
  base64NftAttributeValue: string;
}

export interface MsgSetNFTAttributeResponse {
  nftSchemaCode: string;
  attributeName: string;
  nftAttributeValue: string;
}

const baseMsgCreateNFTSchema: object = { creator: "", nftSchemaBase64: "" };

export const MsgCreateNFTSchema = {
  encode(
    message: MsgCreateNFTSchema,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.nftSchemaBase64 !== "") {
      writer.uint32(18).string(message.nftSchemaBase64);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgCreateNFTSchema {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgCreateNFTSchema } as MsgCreateNFTSchema;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.nftSchemaBase64 = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgCreateNFTSchema {
    const message = { ...baseMsgCreateNFTSchema } as MsgCreateNFTSchema;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (
      object.nftSchemaBase64 !== undefined &&
      object.nftSchemaBase64 !== null
    ) {
      message.nftSchemaBase64 = String(object.nftSchemaBase64);
    } else {
      message.nftSchemaBase64 = "";
    }
    return message;
  },

  toJSON(message: MsgCreateNFTSchema): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.nftSchemaBase64 !== undefined &&
      (obj.nftSchemaBase64 = message.nftSchemaBase64);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgCreateNFTSchema>): MsgCreateNFTSchema {
    const message = { ...baseMsgCreateNFTSchema } as MsgCreateNFTSchema;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (
      object.nftSchemaBase64 !== undefined &&
      object.nftSchemaBase64 !== null
    ) {
      message.nftSchemaBase64 = object.nftSchemaBase64;
    } else {
      message.nftSchemaBase64 = "";
    }
    return message;
  },
};

const baseMsgCreateNFTSchemaResponse: object = { code: "" };

export const MsgCreateNFTSchemaResponse = {
  encode(
    message: MsgCreateNFTSchemaResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.code !== "") {
      writer.uint32(10).string(message.code);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgCreateNFTSchemaResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgCreateNFTSchemaResponse,
    } as MsgCreateNFTSchemaResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.code = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgCreateNFTSchemaResponse {
    const message = {
      ...baseMsgCreateNFTSchemaResponse,
    } as MsgCreateNFTSchemaResponse;
    if (object.code !== undefined && object.code !== null) {
      message.code = String(object.code);
    } else {
      message.code = "";
    }
    return message;
  },

  toJSON(message: MsgCreateNFTSchemaResponse): unknown {
    const obj: any = {};
    message.code !== undefined && (obj.code = message.code);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgCreateNFTSchemaResponse>
  ): MsgCreateNFTSchemaResponse {
    const message = {
      ...baseMsgCreateNFTSchemaResponse,
    } as MsgCreateNFTSchemaResponse;
    if (object.code !== undefined && object.code !== null) {
      message.code = object.code;
    } else {
      message.code = "";
    }
    return message;
  },
};

const baseMsgCreateMetadata: object = {
  creator: "",
  nftSchemaCode: "",
  tokenId: "",
  base64NFTData: "",
};

export const MsgCreateMetadata = {
  encode(message: MsgCreateMetadata, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.nftSchemaCode !== "") {
      writer.uint32(18).string(message.nftSchemaCode);
    }
    if (message.tokenId !== "") {
      writer.uint32(26).string(message.tokenId);
    }
    if (message.base64NFTData !== "") {
      writer.uint32(34).string(message.base64NFTData);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgCreateMetadata {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgCreateMetadata } as MsgCreateMetadata;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.nftSchemaCode = reader.string();
          break;
        case 3:
          message.tokenId = reader.string();
          break;
        case 4:
          message.base64NFTData = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgCreateMetadata {
    const message = { ...baseMsgCreateMetadata } as MsgCreateMetadata;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
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
    if (object.base64NFTData !== undefined && object.base64NFTData !== null) {
      message.base64NFTData = String(object.base64NFTData);
    } else {
      message.base64NFTData = "";
    }
    return message;
  },

  toJSON(message: MsgCreateMetadata): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.nftSchemaCode !== undefined &&
      (obj.nftSchemaCode = message.nftSchemaCode);
    message.tokenId !== undefined && (obj.tokenId = message.tokenId);
    message.base64NFTData !== undefined &&
      (obj.base64NFTData = message.base64NFTData);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgCreateMetadata>): MsgCreateMetadata {
    const message = { ...baseMsgCreateMetadata } as MsgCreateMetadata;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
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
    if (object.base64NFTData !== undefined && object.base64NFTData !== null) {
      message.base64NFTData = object.base64NFTData;
    } else {
      message.base64NFTData = "";
    }
    return message;
  },
};

const baseMsgCreateMetadataResponse: object = {
  nftSchemaCode: "",
  tokenId: "",
};

export const MsgCreateMetadataResponse = {
  encode(
    message: MsgCreateMetadataResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.nftSchemaCode !== "") {
      writer.uint32(10).string(message.nftSchemaCode);
    }
    if (message.tokenId !== "") {
      writer.uint32(18).string(message.tokenId);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgCreateMetadataResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgCreateMetadataResponse,
    } as MsgCreateMetadataResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.nftSchemaCode = reader.string();
          break;
        case 2:
          message.tokenId = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgCreateMetadataResponse {
    const message = {
      ...baseMsgCreateMetadataResponse,
    } as MsgCreateMetadataResponse;
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
    return message;
  },

  toJSON(message: MsgCreateMetadataResponse): unknown {
    const obj: any = {};
    message.nftSchemaCode !== undefined &&
      (obj.nftSchemaCode = message.nftSchemaCode);
    message.tokenId !== undefined && (obj.tokenId = message.tokenId);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgCreateMetadataResponse>
  ): MsgCreateMetadataResponse {
    const message = {
      ...baseMsgCreateMetadataResponse,
    } as MsgCreateMetadataResponse;
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
    return message;
  },
};

const baseOpenseaAttribute: object = { traitType: "" };

export const OpenseaAttribute = {
  encode(message: OpenseaAttribute, writer: Writer = Writer.create()): Writer {
    if (message.traitType !== "") {
      writer.uint32(10).string(message.traitType);
    }
    if (message.value !== undefined) {
      Any.encode(message.value, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): OpenseaAttribute {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseOpenseaAttribute } as OpenseaAttribute;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.traitType = reader.string();
          break;
        case 2:
          message.value = Any.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): OpenseaAttribute {
    const message = { ...baseOpenseaAttribute } as OpenseaAttribute;
    if (object.traitType !== undefined && object.traitType !== null) {
      message.traitType = String(object.traitType);
    } else {
      message.traitType = "";
    }
    if (object.value !== undefined && object.value !== null) {
      message.value = Any.fromJSON(object.value);
    } else {
      message.value = undefined;
    }
    return message;
  },

  toJSON(message: OpenseaAttribute): unknown {
    const obj: any = {};
    message.traitType !== undefined && (obj.traitType = message.traitType);
    message.value !== undefined &&
      (obj.value = message.value ? Any.toJSON(message.value) : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<OpenseaAttribute>): OpenseaAttribute {
    const message = { ...baseOpenseaAttribute } as OpenseaAttribute;
    if (object.traitType !== undefined && object.traitType !== null) {
      message.traitType = object.traitType;
    } else {
      message.traitType = "";
    }
    if (object.value !== undefined && object.value !== null) {
      message.value = Any.fromPartial(object.value);
    } else {
      message.value = undefined;
    }
    return message;
  },
};

const baseUpdatedOpenseaAttributes: object = {};

export const UpdatedOpenseaAttributes = {
  encode(
    message: UpdatedOpenseaAttributes,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.attributes) {
      OpenseaAttribute.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): UpdatedOpenseaAttributes {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseUpdatedOpenseaAttributes,
    } as UpdatedOpenseaAttributes;
    message.attributes = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.attributes.push(
            OpenseaAttribute.decode(reader, reader.uint32())
          );
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): UpdatedOpenseaAttributes {
    const message = {
      ...baseUpdatedOpenseaAttributes,
    } as UpdatedOpenseaAttributes;
    message.attributes = [];
    if (object.attributes !== undefined && object.attributes !== null) {
      for (const e of object.attributes) {
        message.attributes.push(OpenseaAttribute.fromJSON(e));
      }
    }
    return message;
  },

  toJSON(message: UpdatedOpenseaAttributes): unknown {
    const obj: any = {};
    if (message.attributes) {
      obj.attributes = message.attributes.map((e) =>
        e ? OpenseaAttribute.toJSON(e) : undefined
      );
    } else {
      obj.attributes = [];
    }
    return obj;
  },

  fromPartial(
    object: DeepPartial<UpdatedOpenseaAttributes>
  ): UpdatedOpenseaAttributes {
    const message = {
      ...baseUpdatedOpenseaAttributes,
    } as UpdatedOpenseaAttributes;
    message.attributes = [];
    if (object.attributes !== undefined && object.attributes !== null) {
      for (const e of object.attributes) {
        message.attributes.push(OpenseaAttribute.fromPartial(e));
      }
    }
    return message;
  },
};

const baseUpdatedOriginData: object = {};

export const UpdatedOriginData = {
  encode(message: UpdatedOriginData, writer: Writer = Writer.create()): Writer {
    if (message.opensea !== undefined) {
      UpdatedOpenseaAttributes.encode(
        message.opensea,
        writer.uint32(10).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): UpdatedOriginData {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseUpdatedOriginData } as UpdatedOriginData;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.opensea = UpdatedOpenseaAttributes.decode(
            reader,
            reader.uint32()
          );
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): UpdatedOriginData {
    const message = { ...baseUpdatedOriginData } as UpdatedOriginData;
    if (object.opensea !== undefined && object.opensea !== null) {
      message.opensea = UpdatedOpenseaAttributes.fromJSON(object.opensea);
    } else {
      message.opensea = undefined;
    }
    return message;
  },

  toJSON(message: UpdatedOriginData): unknown {
    const obj: any = {};
    message.opensea !== undefined &&
      (obj.opensea = message.opensea
        ? UpdatedOpenseaAttributes.toJSON(message.opensea)
        : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<UpdatedOriginData>): UpdatedOriginData {
    const message = { ...baseUpdatedOriginData } as UpdatedOriginData;
    if (object.opensea !== undefined && object.opensea !== null) {
      message.opensea = UpdatedOpenseaAttributes.fromPartial(object.opensea);
    } else {
      message.opensea = undefined;
    }
    return message;
  },
};

const baseMsgPerformActionByAdmin: object = {
  creator: "",
  nftSchemaCode: "",
  tokenId: "",
  action: "",
  refId: "",
};

export const MsgPerformActionByAdmin = {
  encode(
    message: MsgPerformActionByAdmin,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.nftSchemaCode !== "") {
      writer.uint32(18).string(message.nftSchemaCode);
    }
    if (message.tokenId !== "") {
      writer.uint32(26).string(message.tokenId);
    }
    if (message.action !== "") {
      writer.uint32(34).string(message.action);
    }
    if (message.refId !== "") {
      writer.uint32(42).string(message.refId);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgPerformActionByAdmin {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgPerformActionByAdmin,
    } as MsgPerformActionByAdmin;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.nftSchemaCode = reader.string();
          break;
        case 3:
          message.tokenId = reader.string();
          break;
        case 4:
          message.action = reader.string();
          break;
        case 5:
          message.refId = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgPerformActionByAdmin {
    const message = {
      ...baseMsgPerformActionByAdmin,
    } as MsgPerformActionByAdmin;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
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
    if (object.action !== undefined && object.action !== null) {
      message.action = String(object.action);
    } else {
      message.action = "";
    }
    if (object.refId !== undefined && object.refId !== null) {
      message.refId = String(object.refId);
    } else {
      message.refId = "";
    }
    return message;
  },

  toJSON(message: MsgPerformActionByAdmin): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.nftSchemaCode !== undefined &&
      (obj.nftSchemaCode = message.nftSchemaCode);
    message.tokenId !== undefined && (obj.tokenId = message.tokenId);
    message.action !== undefined && (obj.action = message.action);
    message.refId !== undefined && (obj.refId = message.refId);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgPerformActionByAdmin>
  ): MsgPerformActionByAdmin {
    const message = {
      ...baseMsgPerformActionByAdmin,
    } as MsgPerformActionByAdmin;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
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
    if (object.action !== undefined && object.action !== null) {
      message.action = object.action;
    } else {
      message.action = "";
    }
    if (object.refId !== undefined && object.refId !== null) {
      message.refId = object.refId;
    } else {
      message.refId = "";
    }
    return message;
  },
};

const baseMsgPerformActionByAdminResponse: object = {
  nftSchemaCode: "",
  tokenId: "",
};

export const MsgPerformActionByAdminResponse = {
  encode(
    message: MsgPerformActionByAdminResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.nftSchemaCode !== "") {
      writer.uint32(10).string(message.nftSchemaCode);
    }
    if (message.tokenId !== "") {
      writer.uint32(18).string(message.tokenId);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgPerformActionByAdminResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgPerformActionByAdminResponse,
    } as MsgPerformActionByAdminResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.nftSchemaCode = reader.string();
          break;
        case 2:
          message.tokenId = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgPerformActionByAdminResponse {
    const message = {
      ...baseMsgPerformActionByAdminResponse,
    } as MsgPerformActionByAdminResponse;
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
    return message;
  },

  toJSON(message: MsgPerformActionByAdminResponse): unknown {
    const obj: any = {};
    message.nftSchemaCode !== undefined &&
      (obj.nftSchemaCode = message.nftSchemaCode);
    message.tokenId !== undefined && (obj.tokenId = message.tokenId);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgPerformActionByAdminResponse>
  ): MsgPerformActionByAdminResponse {
    const message = {
      ...baseMsgPerformActionByAdminResponse,
    } as MsgPerformActionByAdminResponse;
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
    return message;
  },
};

const baseMsgAddAttribute: object = {
  creator: "",
  code: "",
  base64NewAttriuteDefenition: "",
};

export const MsgAddAttribute = {
  encode(message: MsgAddAttribute, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.code !== "") {
      writer.uint32(18).string(message.code);
    }
    if (message.base64NewAttriuteDefenition !== "") {
      writer.uint32(26).string(message.base64NewAttriuteDefenition);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgAddAttribute {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgAddAttribute } as MsgAddAttribute;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.code = reader.string();
          break;
        case 3:
          message.base64NewAttriuteDefenition = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgAddAttribute {
    const message = { ...baseMsgAddAttribute } as MsgAddAttribute;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.code !== undefined && object.code !== null) {
      message.code = String(object.code);
    } else {
      message.code = "";
    }
    if (
      object.base64NewAttriuteDefenition !== undefined &&
      object.base64NewAttriuteDefenition !== null
    ) {
      message.base64NewAttriuteDefenition = String(
        object.base64NewAttriuteDefenition
      );
    } else {
      message.base64NewAttriuteDefenition = "";
    }
    return message;
  },

  toJSON(message: MsgAddAttribute): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.code !== undefined && (obj.code = message.code);
    message.base64NewAttriuteDefenition !== undefined &&
      (obj.base64NewAttriuteDefenition = message.base64NewAttriuteDefenition);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgAddAttribute>): MsgAddAttribute {
    const message = { ...baseMsgAddAttribute } as MsgAddAttribute;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.code !== undefined && object.code !== null) {
      message.code = object.code;
    } else {
      message.code = "";
    }
    if (
      object.base64NewAttriuteDefenition !== undefined &&
      object.base64NewAttriuteDefenition !== null
    ) {
      message.base64NewAttriuteDefenition = object.base64NewAttriuteDefenition;
    } else {
      message.base64NewAttriuteDefenition = "";
    }
    return message;
  },
};

const baseMsgAddAttributeResponse: object = { code: "", name: "" };

export const MsgAddAttributeResponse = {
  encode(
    message: MsgAddAttributeResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.code !== "") {
      writer.uint32(10).string(message.code);
    }
    if (message.name !== "") {
      writer.uint32(18).string(message.name);
    }
    if (message.onchainData !== undefined) {
      OnChainData.encode(
        message.onchainData,
        writer.uint32(26).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgAddAttributeResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgAddAttributeResponse,
    } as MsgAddAttributeResponse;
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
          message.onchainData = OnChainData.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgAddAttributeResponse {
    const message = {
      ...baseMsgAddAttributeResponse,
    } as MsgAddAttributeResponse;
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
    if (object.onchainData !== undefined && object.onchainData !== null) {
      message.onchainData = OnChainData.fromJSON(object.onchainData);
    } else {
      message.onchainData = undefined;
    }
    return message;
  },

  toJSON(message: MsgAddAttributeResponse): unknown {
    const obj: any = {};
    message.code !== undefined && (obj.code = message.code);
    message.name !== undefined && (obj.name = message.name);
    message.onchainData !== undefined &&
      (obj.onchainData = message.onchainData
        ? OnChainData.toJSON(message.onchainData)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgAddAttributeResponse>
  ): MsgAddAttributeResponse {
    const message = {
      ...baseMsgAddAttributeResponse,
    } as MsgAddAttributeResponse;
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
    if (object.onchainData !== undefined && object.onchainData !== null) {
      message.onchainData = OnChainData.fromPartial(object.onchainData);
    } else {
      message.onchainData = undefined;
    }
    return message;
  },
};

const baseMsgAddTokenAttribute: object = {
  creator: "",
  code: "",
  base64NewTokenAttriute: "",
};

export const MsgAddTokenAttribute = {
  encode(
    message: MsgAddTokenAttribute,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.code !== "") {
      writer.uint32(18).string(message.code);
    }
    if (message.base64NewTokenAttriute !== "") {
      writer.uint32(26).string(message.base64NewTokenAttriute);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgAddTokenAttribute {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgAddTokenAttribute } as MsgAddTokenAttribute;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.code = reader.string();
          break;
        case 3:
          message.base64NewTokenAttriute = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgAddTokenAttribute {
    const message = { ...baseMsgAddTokenAttribute } as MsgAddTokenAttribute;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.code !== undefined && object.code !== null) {
      message.code = String(object.code);
    } else {
      message.code = "";
    }
    if (
      object.base64NewTokenAttriute !== undefined &&
      object.base64NewTokenAttriute !== null
    ) {
      message.base64NewTokenAttriute = String(object.base64NewTokenAttriute);
    } else {
      message.base64NewTokenAttriute = "";
    }
    return message;
  },

  toJSON(message: MsgAddTokenAttribute): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.code !== undefined && (obj.code = message.code);
    message.base64NewTokenAttriute !== undefined &&
      (obj.base64NewTokenAttriute = message.base64NewTokenAttriute);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgAddTokenAttribute>): MsgAddTokenAttribute {
    const message = { ...baseMsgAddTokenAttribute } as MsgAddTokenAttribute;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.code !== undefined && object.code !== null) {
      message.code = object.code;
    } else {
      message.code = "";
    }
    if (
      object.base64NewTokenAttriute !== undefined &&
      object.base64NewTokenAttriute !== null
    ) {
      message.base64NewTokenAttriute = object.base64NewTokenAttriute;
    } else {
      message.base64NewTokenAttriute = "";
    }
    return message;
  },
};

const baseMsgAddTokenAttributeResponse: object = { code: "", name: "" };

export const MsgAddTokenAttributeResponse = {
  encode(
    message: MsgAddTokenAttributeResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.code !== "") {
      writer.uint32(10).string(message.code);
    }
    if (message.name !== "") {
      writer.uint32(18).string(message.name);
    }
    if (message.onchainData !== undefined) {
      OnChainData.encode(
        message.onchainData,
        writer.uint32(26).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgAddTokenAttributeResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgAddTokenAttributeResponse,
    } as MsgAddTokenAttributeResponse;
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
          message.onchainData = OnChainData.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgAddTokenAttributeResponse {
    const message = {
      ...baseMsgAddTokenAttributeResponse,
    } as MsgAddTokenAttributeResponse;
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
    if (object.onchainData !== undefined && object.onchainData !== null) {
      message.onchainData = OnChainData.fromJSON(object.onchainData);
    } else {
      message.onchainData = undefined;
    }
    return message;
  },

  toJSON(message: MsgAddTokenAttributeResponse): unknown {
    const obj: any = {};
    message.code !== undefined && (obj.code = message.code);
    message.name !== undefined && (obj.name = message.name);
    message.onchainData !== undefined &&
      (obj.onchainData = message.onchainData
        ? OnChainData.toJSON(message.onchainData)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgAddTokenAttributeResponse>
  ): MsgAddTokenAttributeResponse {
    const message = {
      ...baseMsgAddTokenAttributeResponse,
    } as MsgAddTokenAttributeResponse;
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
    if (object.onchainData !== undefined && object.onchainData !== null) {
      message.onchainData = OnChainData.fromPartial(object.onchainData);
    } else {
      message.onchainData = undefined;
    }
    return message;
  },
};

const baseMsgAddAction: object = { creator: "", code: "", base64NewAction: "" };

export const MsgAddAction = {
  encode(message: MsgAddAction, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.code !== "") {
      writer.uint32(18).string(message.code);
    }
    if (message.base64NewAction !== "") {
      writer.uint32(26).string(message.base64NewAction);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgAddAction {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgAddAction } as MsgAddAction;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.code = reader.string();
          break;
        case 3:
          message.base64NewAction = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgAddAction {
    const message = { ...baseMsgAddAction } as MsgAddAction;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.code !== undefined && object.code !== null) {
      message.code = String(object.code);
    } else {
      message.code = "";
    }
    if (
      object.base64NewAction !== undefined &&
      object.base64NewAction !== null
    ) {
      message.base64NewAction = String(object.base64NewAction);
    } else {
      message.base64NewAction = "";
    }
    return message;
  },

  toJSON(message: MsgAddAction): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.code !== undefined && (obj.code = message.code);
    message.base64NewAction !== undefined &&
      (obj.base64NewAction = message.base64NewAction);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgAddAction>): MsgAddAction {
    const message = { ...baseMsgAddAction } as MsgAddAction;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.code !== undefined && object.code !== null) {
      message.code = object.code;
    } else {
      message.code = "";
    }
    if (
      object.base64NewAction !== undefined &&
      object.base64NewAction !== null
    ) {
      message.base64NewAction = object.base64NewAction;
    } else {
      message.base64NewAction = "";
    }
    return message;
  },
};

const baseMsgAddActionResponse: object = { code: "", name: "" };

export const MsgAddActionResponse = {
  encode(
    message: MsgAddActionResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.code !== "") {
      writer.uint32(10).string(message.code);
    }
    if (message.name !== "") {
      writer.uint32(18).string(message.name);
    }
    if (message.onchainData !== undefined) {
      OnChainData.encode(
        message.onchainData,
        writer.uint32(26).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgAddActionResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgAddActionResponse } as MsgAddActionResponse;
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
          message.onchainData = OnChainData.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgAddActionResponse {
    const message = { ...baseMsgAddActionResponse } as MsgAddActionResponse;
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
    if (object.onchainData !== undefined && object.onchainData !== null) {
      message.onchainData = OnChainData.fromJSON(object.onchainData);
    } else {
      message.onchainData = undefined;
    }
    return message;
  },

  toJSON(message: MsgAddActionResponse): unknown {
    const obj: any = {};
    message.code !== undefined && (obj.code = message.code);
    message.name !== undefined && (obj.name = message.name);
    message.onchainData !== undefined &&
      (obj.onchainData = message.onchainData
        ? OnChainData.toJSON(message.onchainData)
        : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgAddActionResponse>): MsgAddActionResponse {
    const message = { ...baseMsgAddActionResponse } as MsgAddActionResponse;
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
    if (object.onchainData !== undefined && object.onchainData !== null) {
      message.onchainData = OnChainData.fromPartial(object.onchainData);
    } else {
      message.onchainData = undefined;
    }
    return message;
  },
};

const baseMsgSetNFTAttribute: object = {
  creator: "",
  nftSchemaCode: "",
  base64NftAttributeValue: "",
};

export const MsgSetNFTAttribute = {
  encode(
    message: MsgSetNFTAttribute,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.nftSchemaCode !== "") {
      writer.uint32(18).string(message.nftSchemaCode);
    }
    if (message.base64NftAttributeValue !== "") {
      writer.uint32(26).string(message.base64NftAttributeValue);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgSetNFTAttribute {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgSetNFTAttribute } as MsgSetNFTAttribute;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.nftSchemaCode = reader.string();
          break;
        case 3:
          message.base64NftAttributeValue = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgSetNFTAttribute {
    const message = { ...baseMsgSetNFTAttribute } as MsgSetNFTAttribute;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.nftSchemaCode !== undefined && object.nftSchemaCode !== null) {
      message.nftSchemaCode = String(object.nftSchemaCode);
    } else {
      message.nftSchemaCode = "";
    }
    if (
      object.base64NftAttributeValue !== undefined &&
      object.base64NftAttributeValue !== null
    ) {
      message.base64NftAttributeValue = String(object.base64NftAttributeValue);
    } else {
      message.base64NftAttributeValue = "";
    }
    return message;
  },

  toJSON(message: MsgSetNFTAttribute): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.nftSchemaCode !== undefined &&
      (obj.nftSchemaCode = message.nftSchemaCode);
    message.base64NftAttributeValue !== undefined &&
      (obj.base64NftAttributeValue = message.base64NftAttributeValue);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgSetNFTAttribute>): MsgSetNFTAttribute {
    const message = { ...baseMsgSetNFTAttribute } as MsgSetNFTAttribute;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.nftSchemaCode !== undefined && object.nftSchemaCode !== null) {
      message.nftSchemaCode = object.nftSchemaCode;
    } else {
      message.nftSchemaCode = "";
    }
    if (
      object.base64NftAttributeValue !== undefined &&
      object.base64NftAttributeValue !== null
    ) {
      message.base64NftAttributeValue = object.base64NftAttributeValue;
    } else {
      message.base64NftAttributeValue = "";
    }
    return message;
  },
};

const baseMsgSetNFTAttributeResponse: object = {
  nftSchemaCode: "",
  attributeName: "",
  nftAttributeValue: "",
};

export const MsgSetNFTAttributeResponse = {
  encode(
    message: MsgSetNFTAttributeResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.nftSchemaCode !== "") {
      writer.uint32(10).string(message.nftSchemaCode);
    }
    if (message.attributeName !== "") {
      writer.uint32(18).string(message.attributeName);
    }
    if (message.nftAttributeValue !== "") {
      writer.uint32(26).string(message.nftAttributeValue);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgSetNFTAttributeResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgSetNFTAttributeResponse,
    } as MsgSetNFTAttributeResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.nftSchemaCode = reader.string();
          break;
        case 2:
          message.attributeName = reader.string();
          break;
        case 3:
          message.nftAttributeValue = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgSetNFTAttributeResponse {
    const message = {
      ...baseMsgSetNFTAttributeResponse,
    } as MsgSetNFTAttributeResponse;
    if (object.nftSchemaCode !== undefined && object.nftSchemaCode !== null) {
      message.nftSchemaCode = String(object.nftSchemaCode);
    } else {
      message.nftSchemaCode = "";
    }
    if (object.attributeName !== undefined && object.attributeName !== null) {
      message.attributeName = String(object.attributeName);
    } else {
      message.attributeName = "";
    }
    if (
      object.nftAttributeValue !== undefined &&
      object.nftAttributeValue !== null
    ) {
      message.nftAttributeValue = String(object.nftAttributeValue);
    } else {
      message.nftAttributeValue = "";
    }
    return message;
  },

  toJSON(message: MsgSetNFTAttributeResponse): unknown {
    const obj: any = {};
    message.nftSchemaCode !== undefined &&
      (obj.nftSchemaCode = message.nftSchemaCode);
    message.attributeName !== undefined &&
      (obj.attributeName = message.attributeName);
    message.nftAttributeValue !== undefined &&
      (obj.nftAttributeValue = message.nftAttributeValue);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgSetNFTAttributeResponse>
  ): MsgSetNFTAttributeResponse {
    const message = {
      ...baseMsgSetNFTAttributeResponse,
    } as MsgSetNFTAttributeResponse;
    if (object.nftSchemaCode !== undefined && object.nftSchemaCode !== null) {
      message.nftSchemaCode = object.nftSchemaCode;
    } else {
      message.nftSchemaCode = "";
    }
    if (object.attributeName !== undefined && object.attributeName !== null) {
      message.attributeName = object.attributeName;
    } else {
      message.attributeName = "";
    }
    if (
      object.nftAttributeValue !== undefined &&
      object.nftAttributeValue !== null
    ) {
      message.nftAttributeValue = object.nftAttributeValue;
    } else {
      message.nftAttributeValue = "";
    }
    return message;
  },
};

/** Msg defines the Msg service. */
export interface Msg {
  CreateNFTSchema(
    request: MsgCreateNFTSchema
  ): Promise<MsgCreateNFTSchemaResponse>;
  CreateMetadata(
    request: MsgCreateMetadata
  ): Promise<MsgCreateMetadataResponse>;
  PerformActionByAdmin(
    request: MsgPerformActionByAdmin
  ): Promise<MsgPerformActionByAdminResponse>;
  AddAttribute(request: MsgAddAttribute): Promise<MsgAddAttributeResponse>;
  AddTokenAttribute(
    request: MsgAddTokenAttribute
  ): Promise<MsgAddTokenAttributeResponse>;
  AddAction(request: MsgAddAction): Promise<MsgAddActionResponse>;
  /** this line is used by starport scaffolding # proto/tx/rpc */
  SetNFTAttribute(
    request: MsgSetNFTAttribute
  ): Promise<MsgSetNFTAttributeResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  CreateNFTSchema(
    request: MsgCreateNFTSchema
  ): Promise<MsgCreateNFTSchemaResponse> {
    const data = MsgCreateNFTSchema.encode(request).finish();
    const promise = this.rpc.request(
      "sixnft.nftmngr.Msg",
      "CreateNFTSchema",
      data
    );
    return promise.then((data) =>
      MsgCreateNFTSchemaResponse.decode(new Reader(data))
    );
  }

  CreateMetadata(
    request: MsgCreateMetadata
  ): Promise<MsgCreateMetadataResponse> {
    const data = MsgCreateMetadata.encode(request).finish();
    const promise = this.rpc.request(
      "sixnft.nftmngr.Msg",
      "CreateMetadata",
      data
    );
    return promise.then((data) =>
      MsgCreateMetadataResponse.decode(new Reader(data))
    );
  }

  PerformActionByAdmin(
    request: MsgPerformActionByAdmin
  ): Promise<MsgPerformActionByAdminResponse> {
    const data = MsgPerformActionByAdmin.encode(request).finish();
    const promise = this.rpc.request(
      "sixnft.nftmngr.Msg",
      "PerformActionByAdmin",
      data
    );
    return promise.then((data) =>
      MsgPerformActionByAdminResponse.decode(new Reader(data))
    );
  }

  AddAttribute(request: MsgAddAttribute): Promise<MsgAddAttributeResponse> {
    const data = MsgAddAttribute.encode(request).finish();
    const promise = this.rpc.request(
      "sixnft.nftmngr.Msg",
      "AddAttribute",
      data
    );
    return promise.then((data) =>
      MsgAddAttributeResponse.decode(new Reader(data))
    );
  }

  AddTokenAttribute(
    request: MsgAddTokenAttribute
  ): Promise<MsgAddTokenAttributeResponse> {
    const data = MsgAddTokenAttribute.encode(request).finish();
    const promise = this.rpc.request(
      "sixnft.nftmngr.Msg",
      "AddTokenAttribute",
      data
    );
    return promise.then((data) =>
      MsgAddTokenAttributeResponse.decode(new Reader(data))
    );
  }

  AddAction(request: MsgAddAction): Promise<MsgAddActionResponse> {
    const data = MsgAddAction.encode(request).finish();
    const promise = this.rpc.request("sixnft.nftmngr.Msg", "AddAction", data);
    return promise.then((data) =>
      MsgAddActionResponse.decode(new Reader(data))
    );
  }

  SetNFTAttribute(
    request: MsgSetNFTAttribute
  ): Promise<MsgSetNFTAttributeResponse> {
    const data = MsgSetNFTAttribute.encode(request).finish();
    const promise = this.rpc.request(
      "sixnft.nftmngr.Msg",
      "SetNFTAttribute",
      data
    );
    return promise.then((data) =>
      MsgSetNFTAttributeResponse.decode(new Reader(data))
    );
  }
}

interface Rpc {
  request(
    service: string,
    method: string,
    data: Uint8Array
  ): Promise<Uint8Array>;
}

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
