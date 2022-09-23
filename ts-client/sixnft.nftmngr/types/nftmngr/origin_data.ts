/* eslint-disable */
import { AttributeDefinition } from "../nftmngr/attribute_definition";
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "sixnft.nftmngr";

export enum AttributeOverriding {
  ORIGIN = 0,
  CHAIN = 1,
  UNRECOGNIZED = -1,
}

export function attributeOverridingFromJSON(object: any): AttributeOverriding {
  switch (object) {
    case 0:
    case "ORIGIN":
      return AttributeOverriding.ORIGIN;
    case 1:
    case "CHAIN":
      return AttributeOverriding.CHAIN;
    case -1:
    case "UNRECOGNIZED":
    default:
      return AttributeOverriding.UNRECOGNIZED;
  }
}

export function attributeOverridingToJSON(object: AttributeOverriding): string {
  switch (object) {
    case AttributeOverriding.ORIGIN:
      return "ORIGIN";
    case AttributeOverriding.CHAIN:
      return "CHAIN";
    default:
      return "UNKNOWN";
  }
}

export enum URIRetrievalMethod {
  BASE = 0,
  TOKEN = 1,
  UNRECOGNIZED = -1,
}

export function uRIRetrievalMethodFromJSON(object: any): URIRetrievalMethod {
  switch (object) {
    case 0:
    case "BASE":
      return URIRetrievalMethod.BASE;
    case 1:
    case "TOKEN":
      return URIRetrievalMethod.TOKEN;
    case -1:
    case "UNRECOGNIZED":
    default:
      return URIRetrievalMethod.UNRECOGNIZED;
  }
}

export function uRIRetrievalMethodToJSON(object: URIRetrievalMethod): string {
  switch (object) {
    case URIRetrievalMethod.BASE:
      return "BASE";
    case URIRetrievalMethod.TOKEN:
      return "TOKEN";
    default:
      return "UNKNOWN";
  }
}

export interface OriginData {
  originChain: string;
  originContractAddress: string;
  originBaseUri: string;
  attributeOverriding: AttributeOverriding;
  metadataFormat: string;
  originAttributes: AttributeDefinition[];
  uriRetrievalMethod: URIRetrievalMethod;
}

const baseOriginData: object = {
  originChain: "",
  originContractAddress: "",
  originBaseUri: "",
  attributeOverriding: 0,
  metadataFormat: "",
  uriRetrievalMethod: 0,
};

export const OriginData = {
  encode(message: OriginData, writer: Writer = Writer.create()): Writer {
    if (message.originChain !== "") {
      writer.uint32(10).string(message.originChain);
    }
    if (message.originContractAddress !== "") {
      writer.uint32(18).string(message.originContractAddress);
    }
    if (message.originBaseUri !== "") {
      writer.uint32(26).string(message.originBaseUri);
    }
    if (message.attributeOverriding !== 0) {
      writer.uint32(32).int32(message.attributeOverriding);
    }
    if (message.metadataFormat !== "") {
      writer.uint32(42).string(message.metadataFormat);
    }
    for (const v of message.originAttributes) {
      AttributeDefinition.encode(v!, writer.uint32(50).fork()).ldelim();
    }
    if (message.uriRetrievalMethod !== 0) {
      writer.uint32(56).int32(message.uriRetrievalMethod);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): OriginData {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseOriginData } as OriginData;
    message.originAttributes = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.originChain = reader.string();
          break;
        case 2:
          message.originContractAddress = reader.string();
          break;
        case 3:
          message.originBaseUri = reader.string();
          break;
        case 4:
          message.attributeOverriding = reader.int32() as any;
          break;
        case 5:
          message.metadataFormat = reader.string();
          break;
        case 6:
          message.originAttributes.push(
            AttributeDefinition.decode(reader, reader.uint32())
          );
          break;
        case 7:
          message.uriRetrievalMethod = reader.int32() as any;
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): OriginData {
    const message = { ...baseOriginData } as OriginData;
    message.originAttributes = [];
    if (object.originChain !== undefined && object.originChain !== null) {
      message.originChain = String(object.originChain);
    } else {
      message.originChain = "";
    }
    if (
      object.originContractAddress !== undefined &&
      object.originContractAddress !== null
    ) {
      message.originContractAddress = String(object.originContractAddress);
    } else {
      message.originContractAddress = "";
    }
    if (object.originBaseUri !== undefined && object.originBaseUri !== null) {
      message.originBaseUri = String(object.originBaseUri);
    } else {
      message.originBaseUri = "";
    }
    if (
      object.attributeOverriding !== undefined &&
      object.attributeOverriding !== null
    ) {
      message.attributeOverriding = attributeOverridingFromJSON(
        object.attributeOverriding
      );
    } else {
      message.attributeOverriding = 0;
    }
    if (object.metadataFormat !== undefined && object.metadataFormat !== null) {
      message.metadataFormat = String(object.metadataFormat);
    } else {
      message.metadataFormat = "";
    }
    if (
      object.originAttributes !== undefined &&
      object.originAttributes !== null
    ) {
      for (const e of object.originAttributes) {
        message.originAttributes.push(AttributeDefinition.fromJSON(e));
      }
    }
    if (
      object.uriRetrievalMethod !== undefined &&
      object.uriRetrievalMethod !== null
    ) {
      message.uriRetrievalMethod = uRIRetrievalMethodFromJSON(
        object.uriRetrievalMethod
      );
    } else {
      message.uriRetrievalMethod = 0;
    }
    return message;
  },

  toJSON(message: OriginData): unknown {
    const obj: any = {};
    message.originChain !== undefined &&
      (obj.originChain = message.originChain);
    message.originContractAddress !== undefined &&
      (obj.originContractAddress = message.originContractAddress);
    message.originBaseUri !== undefined &&
      (obj.originBaseUri = message.originBaseUri);
    message.attributeOverriding !== undefined &&
      (obj.attributeOverriding = attributeOverridingToJSON(
        message.attributeOverriding
      ));
    message.metadataFormat !== undefined &&
      (obj.metadataFormat = message.metadataFormat);
    if (message.originAttributes) {
      obj.originAttributes = message.originAttributes.map((e) =>
        e ? AttributeDefinition.toJSON(e) : undefined
      );
    } else {
      obj.originAttributes = [];
    }
    message.uriRetrievalMethod !== undefined &&
      (obj.uriRetrievalMethod = uRIRetrievalMethodToJSON(
        message.uriRetrievalMethod
      ));
    return obj;
  },

  fromPartial(object: DeepPartial<OriginData>): OriginData {
    const message = { ...baseOriginData } as OriginData;
    message.originAttributes = [];
    if (object.originChain !== undefined && object.originChain !== null) {
      message.originChain = object.originChain;
    } else {
      message.originChain = "";
    }
    if (
      object.originContractAddress !== undefined &&
      object.originContractAddress !== null
    ) {
      message.originContractAddress = object.originContractAddress;
    } else {
      message.originContractAddress = "";
    }
    if (object.originBaseUri !== undefined && object.originBaseUri !== null) {
      message.originBaseUri = object.originBaseUri;
    } else {
      message.originBaseUri = "";
    }
    if (
      object.attributeOverriding !== undefined &&
      object.attributeOverriding !== null
    ) {
      message.attributeOverriding = object.attributeOverriding;
    } else {
      message.attributeOverriding = 0;
    }
    if (object.metadataFormat !== undefined && object.metadataFormat !== null) {
      message.metadataFormat = object.metadataFormat;
    } else {
      message.metadataFormat = "";
    }
    if (
      object.originAttributes !== undefined &&
      object.originAttributes !== null
    ) {
      for (const e of object.originAttributes) {
        message.originAttributes.push(AttributeDefinition.fromPartial(e));
      }
    }
    if (
      object.uriRetrievalMethod !== undefined &&
      object.uriRetrievalMethod !== null
    ) {
      message.uriRetrievalMethod = object.uriRetrievalMethod;
    } else {
      message.uriRetrievalMethod = 0;
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
