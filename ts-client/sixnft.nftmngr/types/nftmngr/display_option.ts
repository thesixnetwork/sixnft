/* eslint-disable */
import { OpenseaDisplayOption } from "../nftmngr/opensea_display_option";
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "sixnft.nftmngr";

export interface DisplayOption {
  boolTrueValue: string;
  boolFalseValue: string;
  opensea: OpenseaDisplayOption | undefined;
}

const baseDisplayOption: object = { boolTrueValue: "", boolFalseValue: "" };

export const DisplayOption = {
  encode(message: DisplayOption, writer: Writer = Writer.create()): Writer {
    if (message.boolTrueValue !== "") {
      writer.uint32(10).string(message.boolTrueValue);
    }
    if (message.boolFalseValue !== "") {
      writer.uint32(18).string(message.boolFalseValue);
    }
    if (message.opensea !== undefined) {
      OpenseaDisplayOption.encode(
        message.opensea,
        writer.uint32(26).fork()
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
          message.boolTrueValue = reader.string();
          break;
        case 2:
          message.boolFalseValue = reader.string();
          break;
        case 3:
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
    if (object.boolTrueValue !== undefined && object.boolTrueValue !== null) {
      message.boolTrueValue = String(object.boolTrueValue);
    } else {
      message.boolTrueValue = "";
    }
    if (object.boolFalseValue !== undefined && object.boolFalseValue !== null) {
      message.boolFalseValue = String(object.boolFalseValue);
    } else {
      message.boolFalseValue = "";
    }
    if (object.opensea !== undefined && object.opensea !== null) {
      message.opensea = OpenseaDisplayOption.fromJSON(object.opensea);
    } else {
      message.opensea = undefined;
    }
    return message;
  },

  toJSON(message: DisplayOption): unknown {
    const obj: any = {};
    message.boolTrueValue !== undefined &&
      (obj.boolTrueValue = message.boolTrueValue);
    message.boolFalseValue !== undefined &&
      (obj.boolFalseValue = message.boolFalseValue);
    message.opensea !== undefined &&
      (obj.opensea = message.opensea
        ? OpenseaDisplayOption.toJSON(message.opensea)
        : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<DisplayOption>): DisplayOption {
    const message = { ...baseDisplayOption } as DisplayOption;
    if (object.boolTrueValue !== undefined && object.boolTrueValue !== null) {
      message.boolTrueValue = object.boolTrueValue;
    } else {
      message.boolTrueValue = "";
    }
    if (object.boolFalseValue !== undefined && object.boolFalseValue !== null) {
      message.boolFalseValue = object.boolFalseValue;
    } else {
      message.boolFalseValue = "";
    }
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
