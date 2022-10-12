/* eslint-disable */
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "thesixnetwork.sixnft.nftmngr";

export interface NFTFeeConfig {}

const baseNFTFeeConfig: object = {};

export const NFTFeeConfig = {
  encode(_: NFTFeeConfig, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): NFTFeeConfig {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseNFTFeeConfig } as NFTFeeConfig;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): NFTFeeConfig {
    const message = { ...baseNFTFeeConfig } as NFTFeeConfig;
    return message;
  },

  toJSON(_: NFTFeeConfig): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(_: DeepPartial<NFTFeeConfig>): NFTFeeConfig {
    const message = { ...baseNFTFeeConfig } as NFTFeeConfig;
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
