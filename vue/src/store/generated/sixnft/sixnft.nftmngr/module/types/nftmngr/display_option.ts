/* eslint-disable */
import { OpenseaDisplayOption } from "../nftmngr/opensea_display_option";
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "sixnft.nftmngr";

export enum BoolDisplayType {
  PRESENT_ABSENT = 0,
  YES_NO = 1,
  TRUE_FALSE = 2,
  UNRECOGNIZED = -1,
}

export function boolDisplayTypeFromJSON(object: any): BoolDisplayType {
  switch (object) {
    case 0:
    case "PRESENT_ABSENT":
      return BoolDisplayType.PRESENT_ABSENT;
    case 1:
    case "YES_NO":
      return BoolDisplayType.YES_NO;
    case 2:
    case "TRUE_FALSE":
      return BoolDisplayType.TRUE_FALSE;
    case -1:
    case "UNRECOGNIZED":
    default:
      return BoolDisplayType.UNRECOGNIZED;
  }
}

export function boolDisplayTypeToJSON(object: BoolDisplayType): string {
  switch (object) {
    case BoolDisplayType.PRESENT_ABSENT:
      return "PRESENT_ABSENT";
    case BoolDisplayType.YES_NO:
      return "YES_NO";
    case BoolDisplayType.TRUE_FALSE:
      return "TRUE_FALSE";
    default:
      return "UNKNOWN";
  }
}

export interface DisplayOption {
  bool_display_type: BoolDisplayType;
  opensea: OpenseaDisplayOption | undefined;
}

const baseDisplayOption: object = { bool_display_type: 0 };

export const DisplayOption = {
  encode(message: DisplayOption, writer: Writer = Writer.create()): Writer {
    if (message.bool_display_type !== 0) {
      writer.uint32(8).int32(message.bool_display_type);
    }
    if (message.opensea !== undefined) {
      OpenseaDisplayOption.encode(
        message.opensea,
        writer.uint32(18).fork()
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
          message.bool_display_type = reader.int32() as any;
          break;
        case 2:
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
    if (
      object.bool_display_type !== undefined &&
      object.bool_display_type !== null
    ) {
      message.bool_display_type = boolDisplayTypeFromJSON(
        object.bool_display_type
      );
    } else {
      message.bool_display_type = 0;
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
    message.bool_display_type !== undefined &&
      (obj.bool_display_type = boolDisplayTypeToJSON(
        message.bool_display_type
      ));
    message.opensea !== undefined &&
      (obj.opensea = message.opensea
        ? OpenseaDisplayOption.toJSON(message.opensea)
        : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<DisplayOption>): DisplayOption {
    const message = { ...baseDisplayOption } as DisplayOption;
    if (
      object.bool_display_type !== undefined &&
      object.bool_display_type !== null
    ) {
      message.bool_display_type = object.bool_display_type;
    } else {
      message.bool_display_type = 0;
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
