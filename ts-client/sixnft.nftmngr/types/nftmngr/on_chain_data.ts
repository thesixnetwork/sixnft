/* eslint-disable */
import { AttributeDefinition } from "../nftmngr/attribute_definition";
import { Action } from "../nftmngr/action";
import { NftAttributeValue } from "../nftmngr/nft_attribute_value";
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "sixnft.nftmngr";

export interface OnChainData {
  revealRequired: boolean;
  revealSecret: Uint8Array;
  nftAttributes: AttributeDefinition[];
  tokenAttributes: AttributeDefinition[];
  actions: Action[];
  status: { [key: string]: boolean };
  onOffSwitch: { [key: string]: boolean };
  nftAttributesValue: NftAttributeValue[];
}

export interface OnChainData_StatusEntry {
  key: string;
  value: boolean;
}

export interface OnChainData_OnOffSwitchEntry {
  key: string;
  value: boolean;
}

const baseOnChainData: object = { revealRequired: false };

export const OnChainData = {
  encode(message: OnChainData, writer: Writer = Writer.create()): Writer {
    if (message.revealRequired === true) {
      writer.uint32(8).bool(message.revealRequired);
    }
    if (message.revealSecret.length !== 0) {
      writer.uint32(18).bytes(message.revealSecret);
    }
    for (const v of message.nftAttributes) {
      AttributeDefinition.encode(v!, writer.uint32(26).fork()).ldelim();
    }
    for (const v of message.tokenAttributes) {
      AttributeDefinition.encode(v!, writer.uint32(34).fork()).ldelim();
    }
    for (const v of message.actions) {
      Action.encode(v!, writer.uint32(42).fork()).ldelim();
    }
    Object.entries(message.status).forEach(([key, value]) => {
      OnChainData_StatusEntry.encode(
        { key: key as any, value },
        writer.uint32(50).fork()
      ).ldelim();
    });
    Object.entries(message.onOffSwitch).forEach(([key, value]) => {
      OnChainData_OnOffSwitchEntry.encode(
        { key: key as any, value },
        writer.uint32(58).fork()
      ).ldelim();
    });
    for (const v of message.nftAttributesValue) {
      NftAttributeValue.encode(v!, writer.uint32(66).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): OnChainData {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseOnChainData } as OnChainData;
    message.nftAttributes = [];
    message.tokenAttributes = [];
    message.actions = [];
    message.status = {};
    message.onOffSwitch = {};
    message.nftAttributesValue = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.revealRequired = reader.bool();
          break;
        case 2:
          message.revealSecret = reader.bytes();
          break;
        case 3:
          message.nftAttributes.push(
            AttributeDefinition.decode(reader, reader.uint32())
          );
          break;
        case 4:
          message.tokenAttributes.push(
            AttributeDefinition.decode(reader, reader.uint32())
          );
          break;
        case 5:
          message.actions.push(Action.decode(reader, reader.uint32()));
          break;
        case 6:
          const entry6 = OnChainData_StatusEntry.decode(
            reader,
            reader.uint32()
          );
          if (entry6.value !== undefined) {
            message.status[entry6.key] = entry6.value;
          }
          break;
        case 7:
          const entry7 = OnChainData_OnOffSwitchEntry.decode(
            reader,
            reader.uint32()
          );
          if (entry7.value !== undefined) {
            message.onOffSwitch[entry7.key] = entry7.value;
          }
          break;
        case 8:
          message.nftAttributesValue.push(
            NftAttributeValue.decode(reader, reader.uint32())
          );
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): OnChainData {
    const message = { ...baseOnChainData } as OnChainData;
    message.nftAttributes = [];
    message.tokenAttributes = [];
    message.actions = [];
    message.status = {};
    message.onOffSwitch = {};
    message.nftAttributesValue = [];
    if (object.revealRequired !== undefined && object.revealRequired !== null) {
      message.revealRequired = Boolean(object.revealRequired);
    } else {
      message.revealRequired = false;
    }
    if (object.revealSecret !== undefined && object.revealSecret !== null) {
      message.revealSecret = bytesFromBase64(object.revealSecret);
    }
    if (object.nftAttributes !== undefined && object.nftAttributes !== null) {
      for (const e of object.nftAttributes) {
        message.nftAttributes.push(AttributeDefinition.fromJSON(e));
      }
    }
    if (
      object.tokenAttributes !== undefined &&
      object.tokenAttributes !== null
    ) {
      for (const e of object.tokenAttributes) {
        message.tokenAttributes.push(AttributeDefinition.fromJSON(e));
      }
    }
    if (object.actions !== undefined && object.actions !== null) {
      for (const e of object.actions) {
        message.actions.push(Action.fromJSON(e));
      }
    }
    if (object.status !== undefined && object.status !== null) {
      Object.entries(object.status).forEach(([key, value]) => {
        message.status[key] = Boolean(value);
      });
    }
    if (object.onOffSwitch !== undefined && object.onOffSwitch !== null) {
      Object.entries(object.onOffSwitch).forEach(([key, value]) => {
        message.onOffSwitch[key] = Boolean(value);
      });
    }
    if (
      object.nftAttributesValue !== undefined &&
      object.nftAttributesValue !== null
    ) {
      for (const e of object.nftAttributesValue) {
        message.nftAttributesValue.push(NftAttributeValue.fromJSON(e));
      }
    }
    return message;
  },

  toJSON(message: OnChainData): unknown {
    const obj: any = {};
    message.revealRequired !== undefined &&
      (obj.revealRequired = message.revealRequired);
    message.revealSecret !== undefined &&
      (obj.revealSecret = base64FromBytes(
        message.revealSecret !== undefined
          ? message.revealSecret
          : new Uint8Array()
      ));
    if (message.nftAttributes) {
      obj.nftAttributes = message.nftAttributes.map((e) =>
        e ? AttributeDefinition.toJSON(e) : undefined
      );
    } else {
      obj.nftAttributes = [];
    }
    if (message.tokenAttributes) {
      obj.tokenAttributes = message.tokenAttributes.map((e) =>
        e ? AttributeDefinition.toJSON(e) : undefined
      );
    } else {
      obj.tokenAttributes = [];
    }
    if (message.actions) {
      obj.actions = message.actions.map((e) =>
        e ? Action.toJSON(e) : undefined
      );
    } else {
      obj.actions = [];
    }
    obj.status = {};
    if (message.status) {
      Object.entries(message.status).forEach(([k, v]) => {
        obj.status[k] = v;
      });
    }
    obj.onOffSwitch = {};
    if (message.onOffSwitch) {
      Object.entries(message.onOffSwitch).forEach(([k, v]) => {
        obj.onOffSwitch[k] = v;
      });
    }
    if (message.nftAttributesValue) {
      obj.nftAttributesValue = message.nftAttributesValue.map((e) =>
        e ? NftAttributeValue.toJSON(e) : undefined
      );
    } else {
      obj.nftAttributesValue = [];
    }
    return obj;
  },

  fromPartial(object: DeepPartial<OnChainData>): OnChainData {
    const message = { ...baseOnChainData } as OnChainData;
    message.nftAttributes = [];
    message.tokenAttributes = [];
    message.actions = [];
    message.status = {};
    message.onOffSwitch = {};
    message.nftAttributesValue = [];
    if (object.revealRequired !== undefined && object.revealRequired !== null) {
      message.revealRequired = object.revealRequired;
    } else {
      message.revealRequired = false;
    }
    if (object.revealSecret !== undefined && object.revealSecret !== null) {
      message.revealSecret = object.revealSecret;
    } else {
      message.revealSecret = new Uint8Array();
    }
    if (object.nftAttributes !== undefined && object.nftAttributes !== null) {
      for (const e of object.nftAttributes) {
        message.nftAttributes.push(AttributeDefinition.fromPartial(e));
      }
    }
    if (
      object.tokenAttributes !== undefined &&
      object.tokenAttributes !== null
    ) {
      for (const e of object.tokenAttributes) {
        message.tokenAttributes.push(AttributeDefinition.fromPartial(e));
      }
    }
    if (object.actions !== undefined && object.actions !== null) {
      for (const e of object.actions) {
        message.actions.push(Action.fromPartial(e));
      }
    }
    if (object.status !== undefined && object.status !== null) {
      Object.entries(object.status).forEach(([key, value]) => {
        if (value !== undefined) {
          message.status[key] = Boolean(value);
        }
      });
    }
    if (object.onOffSwitch !== undefined && object.onOffSwitch !== null) {
      Object.entries(object.onOffSwitch).forEach(([key, value]) => {
        if (value !== undefined) {
          message.onOffSwitch[key] = Boolean(value);
        }
      });
    }
    if (
      object.nftAttributesValue !== undefined &&
      object.nftAttributesValue !== null
    ) {
      for (const e of object.nftAttributesValue) {
        message.nftAttributesValue.push(NftAttributeValue.fromPartial(e));
      }
    }
    return message;
  },
};

const baseOnChainData_StatusEntry: object = { key: "", value: false };

export const OnChainData_StatusEntry = {
  encode(
    message: OnChainData_StatusEntry,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.key !== "") {
      writer.uint32(10).string(message.key);
    }
    if (message.value === true) {
      writer.uint32(16).bool(message.value);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): OnChainData_StatusEntry {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseOnChainData_StatusEntry,
    } as OnChainData_StatusEntry;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.key = reader.string();
          break;
        case 2:
          message.value = reader.bool();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): OnChainData_StatusEntry {
    const message = {
      ...baseOnChainData_StatusEntry,
    } as OnChainData_StatusEntry;
    if (object.key !== undefined && object.key !== null) {
      message.key = String(object.key);
    } else {
      message.key = "";
    }
    if (object.value !== undefined && object.value !== null) {
      message.value = Boolean(object.value);
    } else {
      message.value = false;
    }
    return message;
  },

  toJSON(message: OnChainData_StatusEntry): unknown {
    const obj: any = {};
    message.key !== undefined && (obj.key = message.key);
    message.value !== undefined && (obj.value = message.value);
    return obj;
  },

  fromPartial(
    object: DeepPartial<OnChainData_StatusEntry>
  ): OnChainData_StatusEntry {
    const message = {
      ...baseOnChainData_StatusEntry,
    } as OnChainData_StatusEntry;
    if (object.key !== undefined && object.key !== null) {
      message.key = object.key;
    } else {
      message.key = "";
    }
    if (object.value !== undefined && object.value !== null) {
      message.value = object.value;
    } else {
      message.value = false;
    }
    return message;
  },
};

const baseOnChainData_OnOffSwitchEntry: object = { key: "", value: false };

export const OnChainData_OnOffSwitchEntry = {
  encode(
    message: OnChainData_OnOffSwitchEntry,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.key !== "") {
      writer.uint32(10).string(message.key);
    }
    if (message.value === true) {
      writer.uint32(16).bool(message.value);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): OnChainData_OnOffSwitchEntry {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseOnChainData_OnOffSwitchEntry,
    } as OnChainData_OnOffSwitchEntry;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.key = reader.string();
          break;
        case 2:
          message.value = reader.bool();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): OnChainData_OnOffSwitchEntry {
    const message = {
      ...baseOnChainData_OnOffSwitchEntry,
    } as OnChainData_OnOffSwitchEntry;
    if (object.key !== undefined && object.key !== null) {
      message.key = String(object.key);
    } else {
      message.key = "";
    }
    if (object.value !== undefined && object.value !== null) {
      message.value = Boolean(object.value);
    } else {
      message.value = false;
    }
    return message;
  },

  toJSON(message: OnChainData_OnOffSwitchEntry): unknown {
    const obj: any = {};
    message.key !== undefined && (obj.key = message.key);
    message.value !== undefined && (obj.value = message.value);
    return obj;
  },

  fromPartial(
    object: DeepPartial<OnChainData_OnOffSwitchEntry>
  ): OnChainData_OnOffSwitchEntry {
    const message = {
      ...baseOnChainData_OnOffSwitchEntry,
    } as OnChainData_OnOffSwitchEntry;
    if (object.key !== undefined && object.key !== null) {
      message.key = object.key;
    } else {
      message.key = "";
    }
    if (object.value !== undefined && object.value !== null) {
      message.value = object.value;
    } else {
      message.value = false;
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

const atob: (b64: string) => string =
  globalThis.atob ||
  ((b64) => globalThis.Buffer.from(b64, "base64").toString("binary"));
function bytesFromBase64(b64: string): Uint8Array {
  const bin = atob(b64);
  const arr = new Uint8Array(bin.length);
  for (let i = 0; i < bin.length; ++i) {
    arr[i] = bin.charCodeAt(i);
  }
  return arr;
}

const btoa: (bin: string) => string =
  globalThis.btoa ||
  ((bin) => globalThis.Buffer.from(bin, "binary").toString("base64"));
function base64FromBytes(arr: Uint8Array): string {
  const bin: string[] = [];
  for (let i = 0; i < arr.byteLength; ++i) {
    bin.push(String.fromCharCode(arr[i]));
  }
  return btoa(bin.join(""));
}

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
