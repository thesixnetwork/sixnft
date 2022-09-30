/* eslint-disable */
import { NftData } from "../nftmngr/nft_data";
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "thesixnetwork.sixnft.nftmngr";

export interface NftCollection {
  nftSchemaCode: string;
  nftData: NftData | undefined;
}

const baseNftCollection: object = { nftSchemaCode: "" };

export const NftCollection = {
  encode(message: NftCollection, writer: Writer = Writer.create()): Writer {
    if (message.nftSchemaCode !== "") {
      writer.uint32(10).string(message.nftSchemaCode);
    }
    if (message.nftData !== undefined) {
      NftData.encode(message.nftData, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): NftCollection {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseNftCollection } as NftCollection;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.nftSchemaCode = reader.string();
          break;
        case 2:
          message.nftData = NftData.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): NftCollection {
    const message = { ...baseNftCollection } as NftCollection;
    if (object.nftSchemaCode !== undefined && object.nftSchemaCode !== null) {
      message.nftSchemaCode = String(object.nftSchemaCode);
    } else {
      message.nftSchemaCode = "";
    }
    if (object.nftData !== undefined && object.nftData !== null) {
      message.nftData = NftData.fromJSON(object.nftData);
    } else {
      message.nftData = undefined;
    }
    return message;
  },

  toJSON(message: NftCollection): unknown {
    const obj: any = {};
    message.nftSchemaCode !== undefined &&
      (obj.nftSchemaCode = message.nftSchemaCode);
    message.nftData !== undefined &&
      (obj.nftData = message.nftData
        ? NftData.toJSON(message.nftData)
        : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<NftCollection>): NftCollection {
    const message = { ...baseNftCollection } as NftCollection;
    if (object.nftSchemaCode !== undefined && object.nftSchemaCode !== null) {
      message.nftSchemaCode = object.nftSchemaCode;
    } else {
      message.nftSchemaCode = "";
    }
    if (object.nftData !== undefined && object.nftData !== null) {
      message.nftData = NftData.fromPartial(object.nftData);
    } else {
      message.nftData = undefined;
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
