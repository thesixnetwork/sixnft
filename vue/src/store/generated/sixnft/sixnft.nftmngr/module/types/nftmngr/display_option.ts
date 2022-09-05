/* eslint-disable */
import { OpenseaDisplayOption } from "../nftmngr/opensea_display_option";
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "sixnft.nftmngr";

export interface DisplayOption {
  opensea: OpenseaDisplayOption | undefined;
}

const baseDisplayOption: object = {};

export const DisplayOption = {
  encode(message: DisplayOption, writer: Writer = Writer.create()): Writer {
    if (message.opensea !== undefined) {
      OpenseaDisplayOption.encode(
        message.opensea,
        writer.uint32(10).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): DisplayOption {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseDisplayOption } as DisplayOption;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.opensea = OpenseaDisplayOption.decode(
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

  fromJSON(object: any): DisplayOption {
    const message = { ...baseDisplayOption } as DisplayOption;
    if (object.opensea !== undefined && object.opensea !== null) {
      message.opensea = OpenseaDisplayOption.fromJSON(object.opensea);
    } else {
      message.opensea = undefined;
    }
    return message;
  },

  toJSON(message: DisplayOption): unknown {
    const obj: any = {};
    message.opensea !== undefined &&
      (obj.opensea = message.opensea
        ? OpenseaDisplayOption.toJSON(message.opensea)
        : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<DisplayOption>): DisplayOption {
    const message = { ...baseDisplayOption } as DisplayOption;
    if (object.opensea !== undefined && object.opensea !== null) {
      message.opensea = OpenseaDisplayOption.fromPartial(object.opensea);
    } else {
      message.opensea = undefined;
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
