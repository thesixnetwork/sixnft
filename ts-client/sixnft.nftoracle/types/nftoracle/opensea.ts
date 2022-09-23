/* eslint-disable */
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "sixnft.nftoracle";

export interface Trait {
  traitType: string;
  value: string;
  displayType: string;
  maxValue: string;
}

const baseTrait: object = {
  traitType: "",
  value: "",
  displayType: "",
  maxValue: "",
};

export const Trait = {
  encode(message: Trait, writer: Writer = Writer.create()): Writer {
    if (message.traitType !== "") {
      writer.uint32(10).string(message.traitType);
    }
    if (message.value !== "") {
      writer.uint32(18).string(message.value);
    }
    if (message.displayType !== "") {
      writer.uint32(26).string(message.displayType);
    }
    if (message.maxValue !== "") {
      writer.uint32(34).string(message.maxValue);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): Trait {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseTrait } as Trait;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.traitType = reader.string();
          break;
        case 2:
          message.value = reader.string();
          break;
        case 3:
          message.displayType = reader.string();
          break;
        case 4:
          message.maxValue = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Trait {
    const message = { ...baseTrait } as Trait;
    if (object.traitType !== undefined && object.traitType !== null) {
      message.traitType = String(object.traitType);
    } else {
      message.traitType = "";
    }
    if (object.value !== undefined && object.value !== null) {
      message.value = String(object.value);
    } else {
      message.value = "";
    }
    if (object.displayType !== undefined && object.displayType !== null) {
      message.displayType = String(object.displayType);
    } else {
      message.displayType = "";
    }
    if (object.maxValue !== undefined && object.maxValue !== null) {
      message.maxValue = String(object.maxValue);
    } else {
      message.maxValue = "";
    }
    return message;
  },

  toJSON(message: Trait): unknown {
    const obj: any = {};
    message.traitType !== undefined && (obj.traitType = message.traitType);
    message.value !== undefined && (obj.value = message.value);
    message.displayType !== undefined &&
      (obj.displayType = message.displayType);
    message.maxValue !== undefined && (obj.maxValue = message.maxValue);
    return obj;
  },

  fromPartial(object: DeepPartial<Trait>): Trait {
    const message = { ...baseTrait } as Trait;
    if (object.traitType !== undefined && object.traitType !== null) {
      message.traitType = object.traitType;
    } else {
      message.traitType = "";
    }
    if (object.value !== undefined && object.value !== null) {
      message.value = object.value;
    } else {
      message.value = "";
    }
    if (object.displayType !== undefined && object.displayType !== null) {
      message.displayType = object.displayType;
    } else {
      message.displayType = "";
    }
    if (object.maxValue !== undefined && object.maxValue !== null) {
      message.maxValue = object.maxValue;
    } else {
      message.maxValue = "";
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
