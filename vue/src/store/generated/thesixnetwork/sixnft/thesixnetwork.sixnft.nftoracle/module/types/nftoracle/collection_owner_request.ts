/* eslint-disable */
import * as Long from "long";
import { util, configure, Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "thesixnetwork.sixnft.nftoracle";

export interface CollectionOwnerRequest {
  id: number;
  nftSchemaCode: string;
  base64OwnerSignature: string;
}

const baseCollectionOwnerRequest: object = {
  id: 0,
  nftSchemaCode: "",
  base64OwnerSignature: "",
};

export const CollectionOwnerRequest = {
  encode(
    message: CollectionOwnerRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.id !== 0) {
      writer.uint32(8).uint64(message.id);
    }
    if (message.nftSchemaCode !== "") {
      writer.uint32(18).string(message.nftSchemaCode);
    }
    if (message.base64OwnerSignature !== "") {
      writer.uint32(26).string(message.base64OwnerSignature);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): CollectionOwnerRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseCollectionOwnerRequest } as CollectionOwnerRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = longToNumber(reader.uint64() as Long);
          break;
        case 2:
          message.nftSchemaCode = reader.string();
          break;
        case 3:
          message.base64OwnerSignature = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): CollectionOwnerRequest {
    const message = { ...baseCollectionOwnerRequest } as CollectionOwnerRequest;
    if (object.id !== undefined && object.id !== null) {
      message.id = Number(object.id);
    } else {
      message.id = 0;
    }
    if (object.nftSchemaCode !== undefined && object.nftSchemaCode !== null) {
      message.nftSchemaCode = String(object.nftSchemaCode);
    } else {
      message.nftSchemaCode = "";
    }
    if (
      object.base64OwnerSignature !== undefined &&
      object.base64OwnerSignature !== null
    ) {
      message.base64OwnerSignature = String(object.base64OwnerSignature);
    } else {
      message.base64OwnerSignature = "";
    }
    return message;
  },

  toJSON(message: CollectionOwnerRequest): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    message.nftSchemaCode !== undefined &&
      (obj.nftSchemaCode = message.nftSchemaCode);
    message.base64OwnerSignature !== undefined &&
      (obj.base64OwnerSignature = message.base64OwnerSignature);
    return obj;
  },

  fromPartial(
    object: DeepPartial<CollectionOwnerRequest>
  ): CollectionOwnerRequest {
    const message = { ...baseCollectionOwnerRequest } as CollectionOwnerRequest;
    if (object.id !== undefined && object.id !== null) {
      message.id = object.id;
    } else {
      message.id = 0;
    }
    if (object.nftSchemaCode !== undefined && object.nftSchemaCode !== null) {
      message.nftSchemaCode = object.nftSchemaCode;
    } else {
      message.nftSchemaCode = "";
    }
    if (
      object.base64OwnerSignature !== undefined &&
      object.base64OwnerSignature !== null
    ) {
      message.base64OwnerSignature = object.base64OwnerSignature;
    } else {
      message.base64OwnerSignature = "";
    }
    return message;
  },
};

declare var self: any | undefined;
declare var window: any | undefined;
var globalThis: any = (() => {
  if (typeof globalThis !== "undefined") return globalThis;
  if (typeof self !== "undefined") return self;
  if (typeof window !== "undefined") return window;
  if (typeof global !== "undefined") return global;
  throw "Unable to locate global object";
})();

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

function longToNumber(long: Long): number {
  if (long.gt(Number.MAX_SAFE_INTEGER)) {
    throw new globalThis.Error("Value is larger than Number.MAX_SAFE_INTEGER");
  }
  return long.toNumber();
}

if (util.Long !== Long) {
  util.Long = Long as any;
  configure();
}
