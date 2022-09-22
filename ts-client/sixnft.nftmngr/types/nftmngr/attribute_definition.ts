/* eslint-disable */
import * as Long from "long";
import { util, configure, Writer, Reader } from "protobufjs/minimal";
import {
  NumberAttributeValue,
  StringAttributeValue,
  BooleanAttributeValue,
  FloatAttributeValue,
} from "../nftmngr/nft_attribute_value";
import { DisplayOption } from "../nftmngr/display_option";

export const protobufPackage = "sixnft.nftmngr";

export interface DefaultMintValue {
  numberAttributeValue: NumberAttributeValue | undefined;
  stringAttributeValue: StringAttributeValue | undefined;
  booleanAttributeValue: BooleanAttributeValue | undefined;
  floatAttributeValue: FloatAttributeValue | undefined;
}

export interface AttributeDefinition {
  name: string;
  dataType: string;
  required: boolean;
  displayValueField: string;
  displayOption: DisplayOption | undefined;
  defaultMintValue: DefaultMintValue | undefined;
  hiddenToMarketplace: boolean;
  attributeId: number;
}

const baseDefaultMintValue: object = {};

export const DefaultMintValue = {
  encode(message: DefaultMintValue, writer: Writer = Writer.create()): Writer {
    if (message.numberAttributeValue !== undefined) {
      NumberAttributeValue.encode(
        message.numberAttributeValue,
        writer.uint32(10).fork()
      ).ldelim();
    }
    if (message.stringAttributeValue !== undefined) {
      StringAttributeValue.encode(
        message.stringAttributeValue,
        writer.uint32(18).fork()
      ).ldelim();
    }
    if (message.booleanAttributeValue !== undefined) {
      BooleanAttributeValue.encode(
        message.booleanAttributeValue,
        writer.uint32(26).fork()
      ).ldelim();
    }
    if (message.floatAttributeValue !== undefined) {
      FloatAttributeValue.encode(
        message.floatAttributeValue,
        writer.uint32(34).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): DefaultMintValue {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseDefaultMintValue } as DefaultMintValue;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.numberAttributeValue = NumberAttributeValue.decode(
            reader,
            reader.uint32()
          );
          break;
        case 2:
          message.stringAttributeValue = StringAttributeValue.decode(
            reader,
            reader.uint32()
          );
          break;
        case 3:
          message.booleanAttributeValue = BooleanAttributeValue.decode(
            reader,
            reader.uint32()
          );
          break;
        case 4:
          message.floatAttributeValue = FloatAttributeValue.decode(
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

  fromJSON(object: any): DefaultMintValue {
    const message = { ...baseDefaultMintValue } as DefaultMintValue;
    if (
      object.numberAttributeValue !== undefined &&
      object.numberAttributeValue !== null
    ) {
      message.numberAttributeValue = NumberAttributeValue.fromJSON(
        object.numberAttributeValue
      );
    } else {
      message.numberAttributeValue = undefined;
    }
    if (
      object.stringAttributeValue !== undefined &&
      object.stringAttributeValue !== null
    ) {
      message.stringAttributeValue = StringAttributeValue.fromJSON(
        object.stringAttributeValue
      );
    } else {
      message.stringAttributeValue = undefined;
    }
    if (
      object.booleanAttributeValue !== undefined &&
      object.booleanAttributeValue !== null
    ) {
      message.booleanAttributeValue = BooleanAttributeValue.fromJSON(
        object.booleanAttributeValue
      );
    } else {
      message.booleanAttributeValue = undefined;
    }
    if (
      object.floatAttributeValue !== undefined &&
      object.floatAttributeValue !== null
    ) {
      message.floatAttributeValue = FloatAttributeValue.fromJSON(
        object.floatAttributeValue
      );
    } else {
      message.floatAttributeValue = undefined;
    }
    return message;
  },

  toJSON(message: DefaultMintValue): unknown {
    const obj: any = {};
    message.numberAttributeValue !== undefined &&
      (obj.numberAttributeValue = message.numberAttributeValue
        ? NumberAttributeValue.toJSON(message.numberAttributeValue)
        : undefined);
    message.stringAttributeValue !== undefined &&
      (obj.stringAttributeValue = message.stringAttributeValue
        ? StringAttributeValue.toJSON(message.stringAttributeValue)
        : undefined);
    message.booleanAttributeValue !== undefined &&
      (obj.booleanAttributeValue = message.booleanAttributeValue
        ? BooleanAttributeValue.toJSON(message.booleanAttributeValue)
        : undefined);
    message.floatAttributeValue !== undefined &&
      (obj.floatAttributeValue = message.floatAttributeValue
        ? FloatAttributeValue.toJSON(message.floatAttributeValue)
        : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<DefaultMintValue>): DefaultMintValue {
    const message = { ...baseDefaultMintValue } as DefaultMintValue;
    if (
      object.numberAttributeValue !== undefined &&
      object.numberAttributeValue !== null
    ) {
      message.numberAttributeValue = NumberAttributeValue.fromPartial(
        object.numberAttributeValue
      );
    } else {
      message.numberAttributeValue = undefined;
    }
    if (
      object.stringAttributeValue !== undefined &&
      object.stringAttributeValue !== null
    ) {
      message.stringAttributeValue = StringAttributeValue.fromPartial(
        object.stringAttributeValue
      );
    } else {
      message.stringAttributeValue = undefined;
    }
    if (
      object.booleanAttributeValue !== undefined &&
      object.booleanAttributeValue !== null
    ) {
      message.booleanAttributeValue = BooleanAttributeValue.fromPartial(
        object.booleanAttributeValue
      );
    } else {
      message.booleanAttributeValue = undefined;
    }
    if (
      object.floatAttributeValue !== undefined &&
      object.floatAttributeValue !== null
    ) {
      message.floatAttributeValue = FloatAttributeValue.fromPartial(
        object.floatAttributeValue
      );
    } else {
      message.floatAttributeValue = undefined;
    }
    return message;
  },
};

const baseAttributeDefinition: object = {
  name: "",
  dataType: "",
  required: false,
  displayValueField: "",
  hiddenToMarketplace: false,
  attributeId: 0,
};

export const AttributeDefinition = {
  encode(
    message: AttributeDefinition,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.name !== "") {
      writer.uint32(10).string(message.name);
    }
    if (message.dataType !== "") {
      writer.uint32(18).string(message.dataType);
    }
    if (message.required === true) {
      writer.uint32(24).bool(message.required);
    }
    if (message.displayValueField !== "") {
      writer.uint32(34).string(message.displayValueField);
    }
    if (message.displayOption !== undefined) {
      DisplayOption.encode(
        message.displayOption,
        writer.uint32(42).fork()
      ).ldelim();
    }
    if (message.defaultMintValue !== undefined) {
      DefaultMintValue.encode(
        message.defaultMintValue,
        writer.uint32(50).fork()
      ).ldelim();
    }
    if (message.hiddenToMarketplace === true) {
      writer.uint32(56).bool(message.hiddenToMarketplace);
    }
    if (message.attributeId !== 0) {
      writer.uint32(64).uint64(message.attributeId);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): AttributeDefinition {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseAttributeDefinition } as AttributeDefinition;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.name = reader.string();
          break;
        case 2:
          message.dataType = reader.string();
          break;
        case 3:
          message.required = reader.bool();
          break;
        case 4:
          message.displayValueField = reader.string();
          break;
        case 5:
          message.displayOption = DisplayOption.decode(reader, reader.uint32());
          break;
        case 6:
          message.defaultMintValue = DefaultMintValue.decode(
            reader,
            reader.uint32()
          );
          break;
        case 7:
          message.hiddenToMarketplace = reader.bool();
          break;
        case 8:
          message.attributeId = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): AttributeDefinition {
    const message = { ...baseAttributeDefinition } as AttributeDefinition;
    if (object.name !== undefined && object.name !== null) {
      message.name = String(object.name);
    } else {
      message.name = "";
    }
    if (object.dataType !== undefined && object.dataType !== null) {
      message.dataType = String(object.dataType);
    } else {
      message.dataType = "";
    }
    if (object.required !== undefined && object.required !== null) {
      message.required = Boolean(object.required);
    } else {
      message.required = false;
    }
    if (
      object.displayValueField !== undefined &&
      object.displayValueField !== null
    ) {
      message.displayValueField = String(object.displayValueField);
    } else {
      message.displayValueField = "";
    }
    if (object.displayOption !== undefined && object.displayOption !== null) {
      message.displayOption = DisplayOption.fromJSON(object.displayOption);
    } else {
      message.displayOption = undefined;
    }
    if (
      object.defaultMintValue !== undefined &&
      object.defaultMintValue !== null
    ) {
      message.defaultMintValue = DefaultMintValue.fromJSON(
        object.defaultMintValue
      );
    } else {
      message.defaultMintValue = undefined;
    }
    if (
      object.hiddenToMarketplace !== undefined &&
      object.hiddenToMarketplace !== null
    ) {
      message.hiddenToMarketplace = Boolean(object.hiddenToMarketplace);
    } else {
      message.hiddenToMarketplace = false;
    }
    if (object.attributeId !== undefined && object.attributeId !== null) {
      message.attributeId = Number(object.attributeId);
    } else {
      message.attributeId = 0;
    }
    return message;
  },

  toJSON(message: AttributeDefinition): unknown {
    const obj: any = {};
    message.name !== undefined && (obj.name = message.name);
    message.dataType !== undefined && (obj.dataType = message.dataType);
    message.required !== undefined && (obj.required = message.required);
    message.displayValueField !== undefined &&
      (obj.displayValueField = message.displayValueField);
    message.displayOption !== undefined &&
      (obj.displayOption = message.displayOption
        ? DisplayOption.toJSON(message.displayOption)
        : undefined);
    message.defaultMintValue !== undefined &&
      (obj.defaultMintValue = message.defaultMintValue
        ? DefaultMintValue.toJSON(message.defaultMintValue)
        : undefined);
    message.hiddenToMarketplace !== undefined &&
      (obj.hiddenToMarketplace = message.hiddenToMarketplace);
    message.attributeId !== undefined &&
      (obj.attributeId = message.attributeId);
    return obj;
  },

  fromPartial(object: DeepPartial<AttributeDefinition>): AttributeDefinition {
    const message = { ...baseAttributeDefinition } as AttributeDefinition;
    if (object.name !== undefined && object.name !== null) {
      message.name = object.name;
    } else {
      message.name = "";
    }
    if (object.dataType !== undefined && object.dataType !== null) {
      message.dataType = object.dataType;
    } else {
      message.dataType = "";
    }
    if (object.required !== undefined && object.required !== null) {
      message.required = object.required;
    } else {
      message.required = false;
    }
    if (
      object.displayValueField !== undefined &&
      object.displayValueField !== null
    ) {
      message.displayValueField = object.displayValueField;
    } else {
      message.displayValueField = "";
    }
    if (object.displayOption !== undefined && object.displayOption !== null) {
      message.displayOption = DisplayOption.fromPartial(object.displayOption);
    } else {
      message.displayOption = undefined;
    }
    if (
      object.defaultMintValue !== undefined &&
      object.defaultMintValue !== null
    ) {
      message.defaultMintValue = DefaultMintValue.fromPartial(
        object.defaultMintValue
      );
    } else {
      message.defaultMintValue = undefined;
    }
    if (
      object.hiddenToMarketplace !== undefined &&
      object.hiddenToMarketplace !== null
    ) {
      message.hiddenToMarketplace = object.hiddenToMarketplace;
    } else {
      message.hiddenToMarketplace = false;
    }
    if (object.attributeId !== undefined && object.attributeId !== null) {
      message.attributeId = object.attributeId;
    } else {
      message.attributeId = 0;
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
