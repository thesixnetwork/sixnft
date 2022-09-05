/* eslint-disable */
import * as Long from "long";
import { util, configure, Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "sixnft.nftmngr";

export interface NftAttributeValue {
  name: string;
  number_attribute_value: number | undefined;
  string_attribute_value: string | undefined;
  boolean_attribute_value: boolean | undefined;
  float_attribute_value: number | undefined;
}

const baseNftAttributeValue: object = { name: "" };

export const NftAttributeValue = {
  encode(message: NftAttributeValue, writer: Writer = Writer.create()): Writer {
    if (message.name !== "") {
      writer.uint32(10).string(message.name);
    }
    if (message.number_attribute_value !== undefined) {
      writer.uint32(16).uint64(message.number_attribute_value);
    }
    if (message.string_attribute_value !== undefined) {
      writer.uint32(26).string(message.string_attribute_value);
    }
    if (message.boolean_attribute_value !== undefined) {
      writer.uint32(32).bool(message.boolean_attribute_value);
    }
    if (message.float_attribute_value !== undefined) {
      writer.uint32(45).float(message.float_attribute_value);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): NftAttributeValue {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseNftAttributeValue } as NftAttributeValue;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.name = reader.string();
          break;
        case 2:
          message.number_attribute_value = longToNumber(
            reader.uint64() as Long
          );
          break;
        case 3:
          message.string_attribute_value = reader.string();
          break;
        case 4:
          message.boolean_attribute_value = reader.bool();
          break;
        case 5:
          message.float_attribute_value = reader.float();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): NftAttributeValue {
    const message = { ...baseNftAttributeValue } as NftAttributeValue;
    if (object.name !== undefined && object.name !== null) {
      message.name = String(object.name);
    } else {
      message.name = "";
    }
    if (
      object.number_attribute_value !== undefined &&
      object.number_attribute_value !== null
    ) {
      message.number_attribute_value = Number(object.number_attribute_value);
    } else {
      message.number_attribute_value = undefined;
    }
    if (
      object.string_attribute_value !== undefined &&
      object.string_attribute_value !== null
    ) {
      message.string_attribute_value = String(object.string_attribute_value);
    } else {
      message.string_attribute_value = undefined;
    }
    if (
      object.boolean_attribute_value !== undefined &&
      object.boolean_attribute_value !== null
    ) {
      message.boolean_attribute_value = Boolean(object.boolean_attribute_value);
    } else {
      message.boolean_attribute_value = undefined;
    }
    if (
      object.float_attribute_value !== undefined &&
      object.float_attribute_value !== null
    ) {
      message.float_attribute_value = Number(object.float_attribute_value);
    } else {
      message.float_attribute_value = undefined;
    }
    return message;
  },

  toJSON(message: NftAttributeValue): unknown {
    const obj: any = {};
    message.name !== undefined && (obj.name = message.name);
    message.number_attribute_value !== undefined &&
      (obj.number_attribute_value = message.number_attribute_value);
    message.string_attribute_value !== undefined &&
      (obj.string_attribute_value = message.string_attribute_value);
    message.boolean_attribute_value !== undefined &&
      (obj.boolean_attribute_value = message.boolean_attribute_value);
    message.float_attribute_value !== undefined &&
      (obj.float_attribute_value = message.float_attribute_value);
    return obj;
  },

  fromPartial(object: DeepPartial<NftAttributeValue>): NftAttributeValue {
    const message = { ...baseNftAttributeValue } as NftAttributeValue;
    if (object.name !== undefined && object.name !== null) {
      message.name = object.name;
    } else {
      message.name = "";
    }
    if (
      object.number_attribute_value !== undefined &&
      object.number_attribute_value !== null
    ) {
      message.number_attribute_value = object.number_attribute_value;
    } else {
      message.number_attribute_value = undefined;
    }
    if (
      object.string_attribute_value !== undefined &&
      object.string_attribute_value !== null
    ) {
      message.string_attribute_value = object.string_attribute_value;
    } else {
      message.string_attribute_value = undefined;
    }
    if (
      object.boolean_attribute_value !== undefined &&
      object.boolean_attribute_value !== null
    ) {
      message.boolean_attribute_value = object.boolean_attribute_value;
    } else {
      message.boolean_attribute_value = undefined;
    }
    if (
      object.float_attribute_value !== undefined &&
      object.float_attribute_value !== null
    ) {
      message.float_attribute_value = object.float_attribute_value;
    } else {
      message.float_attribute_value = undefined;
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
