/* eslint-disable */
import { Reader, Writer } from "protobufjs/minimal";

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

/** Msg defines the Msg service. */
export interface Msg {
  CreateNFTSchema(
    request: MsgCreateNFTSchema
  ): Promise<MsgCreateNFTSchemaResponse>;
  /** this line is used by starport scaffolding # proto/tx/rpc */
  CreateMetadata(
    request: MsgCreateMetadata
  ): Promise<MsgCreateMetadataResponse>;
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
