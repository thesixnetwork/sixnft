/* eslint-disable */
import { Timestamp } from "../google/protobuf/timestamp";
import * as Long from "long";
import { util, configure, Writer, Reader } from "protobufjs/minimal";
import { Trait } from "../nftoracle/opensea";

export const protobufPackage = "sixnft.nftoracle";

export enum RequestStatus {
  PENDING = 0,
  SUCCESS_WITH_CONSENSUS = 1,
  FAILED_WITHOUT_CONCENSUS = 2,
  UNRECOGNIZED = -1,
}

export function requestStatusFromJSON(object: any): RequestStatus {
  switch (object) {
    case 0:
    case "PENDING":
      return RequestStatus.PENDING;
    case 1:
    case "SUCCESS_WITH_CONSENSUS":
      return RequestStatus.SUCCESS_WITH_CONSENSUS;
    case 2:
    case "FAILED_WITHOUT_CONCENSUS":
      return RequestStatus.FAILED_WITHOUT_CONCENSUS;
    case -1:
    case "UNRECOGNIZED":
    default:
      return RequestStatus.UNRECOGNIZED;
  }
}

export function requestStatusToJSON(object: RequestStatus): string {
  switch (object) {
    case RequestStatus.PENDING:
      return "PENDING";
    case RequestStatus.SUCCESS_WITH_CONSENSUS:
      return "SUCCESS_WITH_CONSENSUS";
    case RequestStatus.FAILED_WITHOUT_CONCENSUS:
      return "FAILED_WITHOUT_CONCENSUS";
    default:
      return "UNKNOWN";
  }
}

export interface NftOriginData {
  image: string;
  holder_address: string;
  traits: Trait[];
}

export interface MintRequest {
  id: number;
  nft_schema_code: string;
  token_id: string;
  required_confirm: number;
  status: RequestStatus;
  current_confirm: number;
  nft_origin_data: NftOriginData | undefined;
  created_at: Date | undefined;
  valid_until: Date | undefined;
  data_hash: Uint8Array;
}

const baseNftOriginData: object = { image: "", holder_address: "" };

export const NftOriginData = {
  encode(message: NftOriginData, writer: Writer = Writer.create()): Writer {
    if (message.image !== "") {
      writer.uint32(10).string(message.image);
    }
    if (message.holder_address !== "") {
      writer.uint32(18).string(message.holder_address);
    }
    for (const v of message.traits) {
      Trait.encode(v!, writer.uint32(26).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): NftOriginData {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseNftOriginData } as NftOriginData;
    message.traits = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.image = reader.string();
          break;
        case 2:
          message.holder_address = reader.string();
          break;
        case 3:
          message.traits.push(Trait.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): NftOriginData {
    const message = { ...baseNftOriginData } as NftOriginData;
    message.traits = [];
    if (object.image !== undefined && object.image !== null) {
      message.image = String(object.image);
    } else {
      message.image = "";
    }
    if (object.holder_address !== undefined && object.holder_address !== null) {
      message.holder_address = String(object.holder_address);
    } else {
      message.holder_address = "";
    }
    if (object.traits !== undefined && object.traits !== null) {
      for (const e of object.traits) {
        message.traits.push(Trait.fromJSON(e));
      }
    }
    return message;
  },

  toJSON(message: NftOriginData): unknown {
    const obj: any = {};
    message.image !== undefined && (obj.image = message.image);
    message.holder_address !== undefined &&
      (obj.holder_address = message.holder_address);
    if (message.traits) {
      obj.traits = message.traits.map((e) => (e ? Trait.toJSON(e) : undefined));
    } else {
      obj.traits = [];
    }
    return obj;
  },

  fromPartial(object: DeepPartial<NftOriginData>): NftOriginData {
    const message = { ...baseNftOriginData } as NftOriginData;
    message.traits = [];
    if (object.image !== undefined && object.image !== null) {
      message.image = object.image;
    } else {
      message.image = "";
    }
    if (object.holder_address !== undefined && object.holder_address !== null) {
      message.holder_address = object.holder_address;
    } else {
      message.holder_address = "";
    }
    if (object.traits !== undefined && object.traits !== null) {
      for (const e of object.traits) {
        message.traits.push(Trait.fromPartial(e));
      }
    }
    return message;
  },
};

const baseMintRequest: object = {
  id: 0,
  nft_schema_code: "",
  token_id: "",
  required_confirm: 0,
  status: 0,
  current_confirm: 0,
};

export const MintRequest = {
  encode(message: MintRequest, writer: Writer = Writer.create()): Writer {
    if (message.id !== 0) {
      writer.uint32(8).uint64(message.id);
    }
    if (message.nft_schema_code !== "") {
      writer.uint32(18).string(message.nft_schema_code);
    }
    if (message.token_id !== "") {
      writer.uint32(26).string(message.token_id);
    }
    if (message.required_confirm !== 0) {
      writer.uint32(32).uint64(message.required_confirm);
    }
    if (message.status !== 0) {
      writer.uint32(40).int32(message.status);
    }
    if (message.current_confirm !== 0) {
      writer.uint32(48).uint64(message.current_confirm);
    }
    if (message.nft_origin_data !== undefined) {
      NftOriginData.encode(
        message.nft_origin_data,
        writer.uint32(58).fork()
      ).ldelim();
    }
    if (message.created_at !== undefined) {
      Timestamp.encode(
        toTimestamp(message.created_at),
        writer.uint32(66).fork()
      ).ldelim();
    }
    if (message.valid_until !== undefined) {
      Timestamp.encode(
        toTimestamp(message.valid_until),
        writer.uint32(74).fork()
      ).ldelim();
    }
    if (message.data_hash.length !== 0) {
      writer.uint32(82).bytes(message.data_hash);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MintRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMintRequest } as MintRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = longToNumber(reader.uint64() as Long);
          break;
        case 2:
          message.nft_schema_code = reader.string();
          break;
        case 3:
          message.token_id = reader.string();
          break;
        case 4:
          message.required_confirm = longToNumber(reader.uint64() as Long);
          break;
        case 5:
          message.status = reader.int32() as any;
          break;
        case 6:
          message.current_confirm = longToNumber(reader.uint64() as Long);
          break;
        case 7:
          message.nft_origin_data = NftOriginData.decode(
            reader,
            reader.uint32()
          );
          break;
        case 8:
          message.created_at = fromTimestamp(
            Timestamp.decode(reader, reader.uint32())
          );
          break;
        case 9:
          message.valid_until = fromTimestamp(
            Timestamp.decode(reader, reader.uint32())
          );
          break;
        case 10:
          message.data_hash = reader.bytes();
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
    if (object.id !== undefined && object.id !== null) {
      message.id = Number(object.id);
    } else {
      message.id = 0;
    }
    if (
      object.nft_schema_code !== undefined &&
      object.nft_schema_code !== null
    ) {
      message.nft_schema_code = String(object.nft_schema_code);
    } else {
      message.nft_schema_code = "";
    }
    if (object.token_id !== undefined && object.token_id !== null) {
      message.token_id = String(object.token_id);
    } else {
      message.token_id = "";
    }
    if (
      object.required_confirm !== undefined &&
      object.required_confirm !== null
    ) {
      message.required_confirm = Number(object.required_confirm);
    } else {
      message.required_confirm = 0;
    }
    if (object.status !== undefined && object.status !== null) {
      message.status = requestStatusFromJSON(object.status);
    } else {
      message.status = 0;
    }
    if (
      object.current_confirm !== undefined &&
      object.current_confirm !== null
    ) {
      message.current_confirm = Number(object.current_confirm);
    } else {
      message.current_confirm = 0;
    }
    if (
      object.nft_origin_data !== undefined &&
      object.nft_origin_data !== null
    ) {
      message.nft_origin_data = NftOriginData.fromJSON(object.nft_origin_data);
    } else {
      message.nft_origin_data = undefined;
    }
    if (object.created_at !== undefined && object.created_at !== null) {
      message.created_at = fromJsonTimestamp(object.created_at);
    } else {
      message.created_at = undefined;
    }
    if (object.valid_until !== undefined && object.valid_until !== null) {
      message.valid_until = fromJsonTimestamp(object.valid_until);
    } else {
      message.valid_until = undefined;
    }
    if (object.data_hash !== undefined && object.data_hash !== null) {
      message.data_hash = bytesFromBase64(object.data_hash);
    }
    return message;
  },

  toJSON(message: MintRequest): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    message.nft_schema_code !== undefined &&
      (obj.nft_schema_code = message.nft_schema_code);
    message.token_id !== undefined && (obj.token_id = message.token_id);
    message.required_confirm !== undefined &&
      (obj.required_confirm = message.required_confirm);
    message.status !== undefined &&
      (obj.status = requestStatusToJSON(message.status));
    message.current_confirm !== undefined &&
      (obj.current_confirm = message.current_confirm);
    message.nft_origin_data !== undefined &&
      (obj.nft_origin_data = message.nft_origin_data
        ? NftOriginData.toJSON(message.nft_origin_data)
        : undefined);
    message.created_at !== undefined &&
      (obj.created_at =
        message.created_at !== undefined
          ? message.created_at.toISOString()
          : null);
    message.valid_until !== undefined &&
      (obj.valid_until =
        message.valid_until !== undefined
          ? message.valid_until.toISOString()
          : null);
    message.data_hash !== undefined &&
      (obj.data_hash = base64FromBytes(
        message.data_hash !== undefined ? message.data_hash : new Uint8Array()
      ));
    return obj;
  },

  fromPartial(object: DeepPartial<MintRequest>): MintRequest {
    const message = { ...baseMintRequest } as MintRequest;
    if (object.id !== undefined && object.id !== null) {
      message.id = object.id;
    } else {
      message.id = 0;
    }
    if (
      object.nft_schema_code !== undefined &&
      object.nft_schema_code !== null
    ) {
      message.nft_schema_code = object.nft_schema_code;
    } else {
      message.nft_schema_code = "";
    }
    if (object.token_id !== undefined && object.token_id !== null) {
      message.token_id = object.token_id;
    } else {
      message.token_id = "";
    }
    if (
      object.required_confirm !== undefined &&
      object.required_confirm !== null
    ) {
      message.required_confirm = object.required_confirm;
    } else {
      message.required_confirm = 0;
    }
    if (object.status !== undefined && object.status !== null) {
      message.status = object.status;
    } else {
      message.status = 0;
    }
    if (
      object.current_confirm !== undefined &&
      object.current_confirm !== null
    ) {
      message.current_confirm = object.current_confirm;
    } else {
      message.current_confirm = 0;
    }
    if (
      object.nft_origin_data !== undefined &&
      object.nft_origin_data !== null
    ) {
      message.nft_origin_data = NftOriginData.fromPartial(
        object.nft_origin_data
      );
    } else {
      message.nft_origin_data = undefined;
    }
    if (object.created_at !== undefined && object.created_at !== null) {
      message.created_at = object.created_at;
    } else {
      message.created_at = undefined;
    }
    if (object.valid_until !== undefined && object.valid_until !== null) {
      message.valid_until = object.valid_until;
    } else {
      message.valid_until = undefined;
    }
    if (object.data_hash !== undefined && object.data_hash !== null) {
      message.data_hash = object.data_hash;
    } else {
      message.data_hash = new Uint8Array();
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
