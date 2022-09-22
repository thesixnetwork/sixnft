/* eslint-disable */
import * as Long from "long";
import { util, configure, Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "sixnft.nftmngr";

export interface OpenseaDisplayOption {
  displayType: string;
  traitType: string;
  maxValue: number;
}

const baseOpenseaDisplayOption: object = {
  displayType: "",
  traitType: "",
  maxValue: 0,
};

export const OpenseaDisplayOption = {
  encode(
    message: OpenseaDisplayOption,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.displayType !== "") {
      writer.uint32(10).string(message.displayType);
    }
    if (message.traitType !== "") {
      writer.uint32(18).string(message.traitType);
    }
    if (message.maxValue !== 0) {
      writer.uint32(24).uint64(message.maxValue);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): OpenseaDisplayOption {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseOpenseaDisplayOption } as OpenseaDisplayOption;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.displayType = reader.string();
          break;
        case 2:
          message.traitType = reader.string();
          break;
        case 3:
          message.maxValue = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): OpenseaDisplayOption {
    const message = { ...baseOpenseaDisplayOption } as OpenseaDisplayOption;
    if (object.displayType !== undefined && object.displayType !== null) {
      message.displayType = String(object.displayType);
    } else {
      message.displayType = "";
    }
    if (object.traitType !== undefined && object.traitType !== null) {
      message.traitType = String(object.traitType);
    } else {
      message.traitType = "";
    }
    if (object.maxValue !== undefined && object.maxValue !== null) {
      message.maxValue = Number(object.maxValue);
    } else {
      message.maxValue = 0;
    }
    return message;
  },

  toJSON(message: OpenseaDisplayOption): unknown {
    const obj: any = {};
    message.displayType !== undefined &&
      (obj.displayType = message.displayType);
    message.traitType !== undefined && (obj.traitType = message.traitType);
    message.maxValue !== undefined && (obj.maxValue = message.maxValue);
    return obj;
  },

  fromPartial(object: DeepPartial<OpenseaDisplayOption>): OpenseaDisplayOption {
    const message = { ...baseOpenseaDisplayOption } as OpenseaDisplayOption;
    if (object.displayType !== undefined && object.displayType !== null) {
      message.displayType = object.displayType;
    } else {
      message.displayType = "";
    }
    if (object.traitType !== undefined && object.traitType !== null) {
      message.traitType = object.traitType;
    } else {
      message.traitType = "";
    }
    if (object.maxValue !== undefined && object.maxValue !== null) {
      message.maxValue = object.maxValue;
    } else {
      message.maxValue = 0;
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
