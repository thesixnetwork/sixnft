/* eslint-disable */
import {
  RequestStatus,
  DataHash,
  requestStatusFromJSON,
  requestStatusToJSON,
} from "../nftoracle/request";
import { Timestamp } from "../google/protobuf/timestamp";
import * as Long from "long";
import { util, configure, Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "sixnft.nftoracle";

export interface ActionParam {
  nftSchemaCode: string;
  tokenId: string;
  action: string;
  refId: string;
  expiredAt: Date | undefined;
}

export interface ActionRequest {
  id: number;
  nftSchemaCode: string;
  tokenId: string;
  action: string;
  caller: string;
  refId: string;
  requiredConfirm: number;
  status: RequestStatus;
  currentConfirm: number;
  confirmers: { [key: string]: boolean };
  createdAt: Date | undefined;
  validUntil: Date | undefined;
  dataHashes: DataHash[];
  expiredHeight: number;
  executionErrorMessage: string;
}

export interface ActionRequest_ConfirmersEntry {
  key: string;
  value: boolean;
}

const baseActionParam: object = {
  nftSchemaCode: "",
  tokenId: "",
  action: "",
  refId: "",
};

export const ActionParam = {
  encode(message: ActionParam, writer: Writer = Writer.create()): Writer {
    if (message.nftSchemaCode !== "") {
      writer.uint32(10).string(message.nftSchemaCode);
    }
    if (message.tokenId !== "") {
      writer.uint32(18).string(message.tokenId);
    }
    if (message.action !== "") {
      writer.uint32(26).string(message.action);
    }
    if (message.refId !== "") {
      writer.uint32(34).string(message.refId);
    }
    if (message.expiredAt !== undefined) {
      Timestamp.encode(
        toTimestamp(message.expiredAt),
        writer.uint32(42).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): ActionParam {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseActionParam } as ActionParam;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.nftSchemaCode = reader.string();
          break;
        case 2:
          message.tokenId = reader.string();
          break;
        case 3:
          message.action = reader.string();
          break;
        case 4:
          message.refId = reader.string();
          break;
        case 5:
          message.expiredAt = fromTimestamp(
            Timestamp.decode(reader, reader.uint32())
          );
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): ActionParam {
    const message = { ...baseActionParam } as ActionParam;
    if (object.nftSchemaCode !== undefined && object.nftSchemaCode !== null) {
      message.nftSchemaCode = String(object.nftSchemaCode);
    } else {
      message.nftSchemaCode = "";
    }
    if (object.tokenId !== undefined && object.tokenId !== null) {
      message.tokenId = String(object.tokenId);
    } else {
      message.tokenId = "";
    }
    if (object.action !== undefined && object.action !== null) {
      message.action = String(object.action);
    } else {
      message.action = "";
    }
    if (object.refId !== undefined && object.refId !== null) {
      message.refId = String(object.refId);
    } else {
      message.refId = "";
    }
    if (object.expiredAt !== undefined && object.expiredAt !== null) {
      message.expiredAt = fromJsonTimestamp(object.expiredAt);
    } else {
      message.expiredAt = undefined;
    }
    return message;
  },

  toJSON(message: ActionParam): unknown {
    const obj: any = {};
    message.nftSchemaCode !== undefined &&
      (obj.nftSchemaCode = message.nftSchemaCode);
    message.tokenId !== undefined && (obj.tokenId = message.tokenId);
    message.action !== undefined && (obj.action = message.action);
    message.refId !== undefined && (obj.refId = message.refId);
    message.expiredAt !== undefined &&
      (obj.expiredAt =
        message.expiredAt !== undefined
          ? message.expiredAt.toISOString()
          : null);
    return obj;
  },

  fromPartial(object: DeepPartial<ActionParam>): ActionParam {
    const message = { ...baseActionParam } as ActionParam;
    if (object.nftSchemaCode !== undefined && object.nftSchemaCode !== null) {
      message.nftSchemaCode = object.nftSchemaCode;
    } else {
      message.nftSchemaCode = "";
    }
    if (object.tokenId !== undefined && object.tokenId !== null) {
      message.tokenId = object.tokenId;
    } else {
      message.tokenId = "";
    }
    if (object.action !== undefined && object.action !== null) {
      message.action = object.action;
    } else {
      message.action = "";
    }
    if (object.refId !== undefined && object.refId !== null) {
      message.refId = object.refId;
    } else {
      message.refId = "";
    }
    if (object.expiredAt !== undefined && object.expiredAt !== null) {
      message.expiredAt = object.expiredAt;
    } else {
      message.expiredAt = undefined;
    }
    return message;
  },
};

const baseActionRequest: object = {
  id: 0,
  nftSchemaCode: "",
  tokenId: "",
  action: "",
  caller: "",
  refId: "",
  requiredConfirm: 0,
  status: 0,
  currentConfirm: 0,
  expiredHeight: 0,
  executionErrorMessage: "",
};

export const ActionRequest = {
  encode(message: ActionRequest, writer: Writer = Writer.create()): Writer {
    if (message.id !== 0) {
      writer.uint32(8).uint64(message.id);
    }
    if (message.nftSchemaCode !== "") {
      writer.uint32(18).string(message.nftSchemaCode);
    }
    if (message.tokenId !== "") {
      writer.uint32(26).string(message.tokenId);
    }
    if (message.action !== "") {
      writer.uint32(34).string(message.action);
    }
    if (message.caller !== "") {
      writer.uint32(42).string(message.caller);
    }
    if (message.refId !== "") {
      writer.uint32(50).string(message.refId);
    }
    if (message.requiredConfirm !== 0) {
      writer.uint32(56).uint64(message.requiredConfirm);
    }
    if (message.status !== 0) {
      writer.uint32(64).int32(message.status);
    }
    if (message.currentConfirm !== 0) {
      writer.uint32(72).uint64(message.currentConfirm);
    }
    Object.entries(message.confirmers).forEach(([key, value]) => {
      ActionRequest_ConfirmersEntry.encode(
        { key: key as any, value },
        writer.uint32(82).fork()
      ).ldelim();
    });
    if (message.createdAt !== undefined) {
      Timestamp.encode(
        toTimestamp(message.createdAt),
        writer.uint32(90).fork()
      ).ldelim();
    }
    if (message.validUntil !== undefined) {
      Timestamp.encode(
        toTimestamp(message.validUntil),
        writer.uint32(98).fork()
      ).ldelim();
    }
    for (const v of message.dataHashes) {
      DataHash.encode(v!, writer.uint32(106).fork()).ldelim();
    }
    if (message.expiredHeight !== 0) {
      writer.uint32(112).int64(message.expiredHeight);
    }
    if (message.executionErrorMessage !== "") {
      writer.uint32(122).string(message.executionErrorMessage);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): ActionRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseActionRequest } as ActionRequest;
    message.confirmers = {};
    message.dataHashes = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = longToNumber(reader.uint64() as Long);
          break;
        case 2:
          message.nftSchemaCode = reader.string();
          break;
        case 3:
          message.tokenId = reader.string();
          break;
        case 4:
          message.action = reader.string();
          break;
        case 5:
          message.caller = reader.string();
          break;
        case 6:
          message.refId = reader.string();
          break;
        case 7:
          message.requiredConfirm = longToNumber(reader.uint64() as Long);
          break;
        case 8:
          message.status = reader.int32() as any;
          break;
        case 9:
          message.currentConfirm = longToNumber(reader.uint64() as Long);
          break;
        case 10:
          const entry10 = ActionRequest_ConfirmersEntry.decode(
            reader,
            reader.uint32()
          );
          if (entry10.value !== undefined) {
            message.confirmers[entry10.key] = entry10.value;
          }
          break;
        case 11:
          message.createdAt = fromTimestamp(
            Timestamp.decode(reader, reader.uint32())
          );
          break;
        case 12:
          message.validUntil = fromTimestamp(
            Timestamp.decode(reader, reader.uint32())
          );
          break;
        case 13:
          message.dataHashes.push(DataHash.decode(reader, reader.uint32()));
          break;
        case 14:
          message.expiredHeight = longToNumber(reader.int64() as Long);
          break;
        case 15:
          message.executionErrorMessage = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): ActionRequest {
    const message = { ...baseActionRequest } as ActionRequest;
    message.confirmers = {};
    message.dataHashes = [];
    if (object.id !== undefined && object.id !== null) {
      message.id = Number(object.id);
    } else {
      message.id = 0;
    }
    if (object.nftSchemaCode !== undefined && object.nftSchemaCode !== null) {
      message.nftSchemaCode = String(object.nftSchemaCode);
    } else {
      message.nftSchemaCode = "";
    }
    if (object.tokenId !== undefined && object.tokenId !== null) {
      message.tokenId = String(object.tokenId);
    } else {
      message.tokenId = "";
    }
    if (object.action !== undefined && object.action !== null) {
      message.action = String(object.action);
    } else {
      message.action = "";
    }
    if (object.caller !== undefined && object.caller !== null) {
      message.caller = String(object.caller);
    } else {
      message.caller = "";
    }
    if (object.refId !== undefined && object.refId !== null) {
      message.refId = String(object.refId);
    } else {
      message.refId = "";
    }
    if (
      object.requiredConfirm !== undefined &&
      object.requiredConfirm !== null
    ) {
      message.requiredConfirm = Number(object.requiredConfirm);
    } else {
      message.requiredConfirm = 0;
    }
    if (object.status !== undefined && object.status !== null) {
      message.status = requestStatusFromJSON(object.status);
    } else {
      message.status = 0;
    }
    if (object.currentConfirm !== undefined && object.currentConfirm !== null) {
      message.currentConfirm = Number(object.currentConfirm);
    } else {
      message.currentConfirm = 0;
    }
    if (object.confirmers !== undefined && object.confirmers !== null) {
      Object.entries(object.confirmers).forEach(([key, value]) => {
        message.confirmers[key] = Boolean(value);
      });
    }
    if (object.createdAt !== undefined && object.createdAt !== null) {
      message.createdAt = fromJsonTimestamp(object.createdAt);
    } else {
      message.createdAt = undefined;
    }
    if (object.validUntil !== undefined && object.validUntil !== null) {
      message.validUntil = fromJsonTimestamp(object.validUntil);
    } else {
      message.validUntil = undefined;
    }
    if (object.dataHashes !== undefined && object.dataHashes !== null) {
      for (const e of object.dataHashes) {
        message.dataHashes.push(DataHash.fromJSON(e));
      }
    }
    if (object.expiredHeight !== undefined && object.expiredHeight !== null) {
      message.expiredHeight = Number(object.expiredHeight);
    } else {
      message.expiredHeight = 0;
    }
    if (
      object.executionErrorMessage !== undefined &&
      object.executionErrorMessage !== null
    ) {
      message.executionErrorMessage = String(object.executionErrorMessage);
    } else {
      message.executionErrorMessage = "";
    }
    return message;
  },

  toJSON(message: ActionRequest): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    message.nftSchemaCode !== undefined &&
      (obj.nftSchemaCode = message.nftSchemaCode);
    message.tokenId !== undefined && (obj.tokenId = message.tokenId);
    message.action !== undefined && (obj.action = message.action);
    message.caller !== undefined && (obj.caller = message.caller);
    message.refId !== undefined && (obj.refId = message.refId);
    message.requiredConfirm !== undefined &&
      (obj.requiredConfirm = message.requiredConfirm);
    message.status !== undefined &&
      (obj.status = requestStatusToJSON(message.status));
    message.currentConfirm !== undefined &&
      (obj.currentConfirm = message.currentConfirm);
    obj.confirmers = {};
    if (message.confirmers) {
      Object.entries(message.confirmers).forEach(([k, v]) => {
        obj.confirmers[k] = v;
      });
    }
    message.createdAt !== undefined &&
      (obj.createdAt =
        message.createdAt !== undefined
          ? message.createdAt.toISOString()
          : null);
    message.validUntil !== undefined &&
      (obj.validUntil =
        message.validUntil !== undefined
          ? message.validUntil.toISOString()
          : null);
    if (message.dataHashes) {
      obj.dataHashes = message.dataHashes.map((e) =>
        e ? DataHash.toJSON(e) : undefined
      );
    } else {
      obj.dataHashes = [];
    }
    message.expiredHeight !== undefined &&
      (obj.expiredHeight = message.expiredHeight);
    message.executionErrorMessage !== undefined &&
      (obj.executionErrorMessage = message.executionErrorMessage);
    return obj;
  },

  fromPartial(object: DeepPartial<ActionRequest>): ActionRequest {
    const message = { ...baseActionRequest } as ActionRequest;
    message.confirmers = {};
    message.dataHashes = [];
    if (object.id !== undefined && object.id !== null) {
      message.id = object.id;
    } else {
      message.id = 0;
    }
    if (object.nftSchemaCode !== undefined && object.nftSchemaCode !== null) {
      message.nftSchemaCode = object.nftSchemaCode;
    } else {
      message.nftSchemaCode = "";
    }
    if (object.tokenId !== undefined && object.tokenId !== null) {
      message.tokenId = object.tokenId;
    } else {
      message.tokenId = "";
    }
    if (object.action !== undefined && object.action !== null) {
      message.action = object.action;
    } else {
      message.action = "";
    }
    if (object.caller !== undefined && object.caller !== null) {
      message.caller = object.caller;
    } else {
      message.caller = "";
    }
    if (object.refId !== undefined && object.refId !== null) {
      message.refId = object.refId;
    } else {
      message.refId = "";
    }
    if (
      object.requiredConfirm !== undefined &&
      object.requiredConfirm !== null
    ) {
      message.requiredConfirm = object.requiredConfirm;
    } else {
      message.requiredConfirm = 0;
    }
    if (object.status !== undefined && object.status !== null) {
      message.status = object.status;
    } else {
      message.status = 0;
    }
    if (object.currentConfirm !== undefined && object.currentConfirm !== null) {
      message.currentConfirm = object.currentConfirm;
    } else {
      message.currentConfirm = 0;
    }
    if (object.confirmers !== undefined && object.confirmers !== null) {
      Object.entries(object.confirmers).forEach(([key, value]) => {
        if (value !== undefined) {
          message.confirmers[key] = Boolean(value);
        }
      });
    }
    if (object.createdAt !== undefined && object.createdAt !== null) {
      message.createdAt = object.createdAt;
    } else {
      message.createdAt = undefined;
    }
    if (object.validUntil !== undefined && object.validUntil !== null) {
      message.validUntil = object.validUntil;
    } else {
      message.validUntil = undefined;
    }
    if (object.dataHashes !== undefined && object.dataHashes !== null) {
      for (const e of object.dataHashes) {
        message.dataHashes.push(DataHash.fromPartial(e));
      }
    }
    if (object.expiredHeight !== undefined && object.expiredHeight !== null) {
      message.expiredHeight = object.expiredHeight;
    } else {
      message.expiredHeight = 0;
    }
    if (
      object.executionErrorMessage !== undefined &&
      object.executionErrorMessage !== null
    ) {
      message.executionErrorMessage = object.executionErrorMessage;
    } else {
      message.executionErrorMessage = "";
    }
    return message;
  },
};

const baseActionRequest_ConfirmersEntry: object = { key: "", value: false };

export const ActionRequest_ConfirmersEntry = {
  encode(
    message: ActionRequest_ConfirmersEntry,
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
  ): ActionRequest_ConfirmersEntry {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseActionRequest_ConfirmersEntry,
    } as ActionRequest_ConfirmersEntry;
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

  fromJSON(object: any): ActionRequest_ConfirmersEntry {
    const message = {
      ...baseActionRequest_ConfirmersEntry,
    } as ActionRequest_ConfirmersEntry;
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

  toJSON(message: ActionRequest_ConfirmersEntry): unknown {
    const obj: any = {};
    message.key !== undefined && (obj.key = message.key);
    message.value !== undefined && (obj.value = message.value);
    return obj;
  },

  fromPartial(
    object: DeepPartial<ActionRequest_ConfirmersEntry>
  ): ActionRequest_ConfirmersEntry {
    const message = {
      ...baseActionRequest_ConfirmersEntry,
    } as ActionRequest_ConfirmersEntry;
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

function toTimestamp(date: Date): Timestamp {
  const seconds = date.getTime() / 1_000;
  const nanos = (date.getTime() % 1_000) * 1_000_000;
  return { seconds, nanos };
}

function fromTimestamp(t: Timestamp): Date {
  let millis = t.seconds * 1_000;
  millis += t.nanos / 1_000_000;
  return new Date(millis);
}

function fromJsonTimestamp(o: any): Date {
  if (o instanceof Date) {
    return o;
  } else if (typeof o === "string") {
    return new Date(o);
  } else {
    return fromTimestamp(Timestamp.fromJSON(o));
  }
}

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
