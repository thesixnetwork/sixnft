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

export interface MintRequest {
  id: number;
  nftSchemaCode: string;
  tokenId: string;
  requiredConfirm: number;
  status: RequestStatus;
  currentConfirm: number;
  confirmers: { [key: string]: boolean };
  /** NftOriginData nft_origin_data = 8; */
  createdAt: Date | undefined;
  validUntil: Date | undefined;
  dataHashes: DataHash[];
  expiredHeight: number;
}

export interface MintRequest_ConfirmersEntry {
  key: string;
  value: boolean;
}

const baseMintRequest: object = {
  id: 0,
  nftSchemaCode: "",
  tokenId: "",
  requiredConfirm: 0,
  status: 0,
  currentConfirm: 0,
  expiredHeight: 0,
};

export const MintRequest = {
  encode(message: MintRequest, writer: Writer = Writer.create()): Writer {
    if (message.id !== 0) {
      writer.uint32(8).uint64(message.id);
    }
    if (message.nftSchemaCode !== "") {
      writer.uint32(18).string(message.nftSchemaCode);
    }
    if (message.tokenId !== "") {
      writer.uint32(26).string(message.tokenId);
    }
    if (message.requiredConfirm !== 0) {
      writer.uint32(32).uint64(message.requiredConfirm);
    }
    if (message.status !== 0) {
      writer.uint32(40).int32(message.status);
    }
    if (message.currentConfirm !== 0) {
      writer.uint32(48).uint64(message.currentConfirm);
    }
    Object.entries(message.confirmers).forEach(([key, value]) => {
      MintRequest_ConfirmersEntry.encode(
        { key: key as any, value },
        writer.uint32(58).fork()
      ).ldelim();
    });
    if (message.createdAt !== undefined) {
      Timestamp.encode(
        toTimestamp(message.createdAt),
        writer.uint32(66).fork()
      ).ldelim();
    }
    if (message.validUntil !== undefined) {
      Timestamp.encode(
        toTimestamp(message.validUntil),
        writer.uint32(74).fork()
      ).ldelim();
    }
    for (const v of message.dataHashes) {
      DataHash.encode(v!, writer.uint32(82).fork()).ldelim();
    }
    if (message.expiredHeight !== 0) {
      writer.uint32(88).int64(message.expiredHeight);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MintRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMintRequest } as MintRequest;
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
          message.requiredConfirm = longToNumber(reader.uint64() as Long);
          break;
        case 5:
          message.status = reader.int32() as any;
          break;
        case 6:
          message.currentConfirm = longToNumber(reader.uint64() as Long);
          break;
        case 7:
          const entry7 = MintRequest_ConfirmersEntry.decode(
            reader,
            reader.uint32()
          );
          if (entry7.value !== undefined) {
            message.confirmers[entry7.key] = entry7.value;
          }
          break;
        case 8:
          message.createdAt = fromTimestamp(
            Timestamp.decode(reader, reader.uint32())
          );
          break;
        case 9:
          message.validUntil = fromTimestamp(
            Timestamp.decode(reader, reader.uint32())
          );
          break;
        case 10:
          message.dataHashes.push(DataHash.decode(reader, reader.uint32()));
          break;
        case 11:
          message.expiredHeight = longToNumber(reader.int64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MintRequest {
    const message = { ...baseMintRequest } as MintRequest;
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
    return message;
  },

  toJSON(message: MintRequest): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    message.nftSchemaCode !== undefined &&
      (obj.nftSchemaCode = message.nftSchemaCode);
    message.tokenId !== undefined && (obj.tokenId = message.tokenId);
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
    return obj;
  },

  fromPartial(object: DeepPartial<MintRequest>): MintRequest {
    const message = { ...baseMintRequest } as MintRequest;
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
    return message;
  },
};

const baseMintRequest_ConfirmersEntry: object = { key: "", value: false };

export const MintRequest_ConfirmersEntry = {
  encode(
    message: MintRequest_ConfirmersEntry,
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
  ): MintRequest_ConfirmersEntry {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMintRequest_ConfirmersEntry,
    } as MintRequest_ConfirmersEntry;
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

  fromJSON(object: any): MintRequest_ConfirmersEntry {
    const message = {
      ...baseMintRequest_ConfirmersEntry,
    } as MintRequest_ConfirmersEntry;
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

  toJSON(message: MintRequest_ConfirmersEntry): unknown {
    const obj: any = {};
    message.key !== undefined && (obj.key = message.key);
    message.value !== undefined && (obj.value = message.value);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MintRequest_ConfirmersEntry>
  ): MintRequest_ConfirmersEntry {
    const message = {
      ...baseMintRequest_ConfirmersEntry,
    } as MintRequest_ConfirmersEntry;
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
