/* eslint-disable */
import { Reader, Writer } from "protobufjs/minimal";
import { Params } from "../nftmngr/params";
import { NFTSchema } from "../nftmngr/nft_schema";
import {
  PageRequest,
  PageResponse,
} from "../cosmos/base/query/v1beta1/pagination";

export const protobufPackage = "sixnft.nftmngr";

/** QueryParamsRequest is request type for the Query/Params RPC method. */
export interface QueryParamsRequest {}

/** QueryParamsResponse is response type for the Query/Params RPC method. */
export interface QueryParamsResponse {
  /** params holds all the parameters of this module. */
  params: Params | undefined;
}

export interface QueryGetNFTSchemaRequest {
  code: string;
}

export interface QueryGetNFTSchemaResponse {
  nFTSchema: NFTSchema | undefined;
}

export interface QueryAllNFTSchemaRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllNFTSchemaResponse {
  nFTSchema: NFTSchema[];
  pagination: PageResponse | undefined;
}

const baseQueryParamsRequest: object = {};

export const QueryParamsRequest = {
  encode(_: QueryParamsRequest, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryParamsRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryParamsRequest } as QueryParamsRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): QueryParamsRequest {
    const message = { ...baseQueryParamsRequest } as QueryParamsRequest;
    return message;
  },

  toJSON(_: QueryParamsRequest): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(_: DeepPartial<QueryParamsRequest>): QueryParamsRequest {
    const message = { ...baseQueryParamsRequest } as QueryParamsRequest;
    return message;
  },
};

const baseQueryParamsResponse: object = {};

export const QueryParamsResponse = {
  encode(
    message: QueryParamsResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.params !== undefined) {
      Params.encode(message.params, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryParamsResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryParamsResponse } as QueryParamsResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.params = Params.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryParamsResponse {
    const message = { ...baseQueryParamsResponse } as QueryParamsResponse;
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromJSON(object.params);
    } else {
      message.params = undefined;
    }
    return message;
  },

  toJSON(message: QueryParamsResponse): unknown {
    const obj: any = {};
    message.params !== undefined &&
      (obj.params = message.params ? Params.toJSON(message.params) : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<QueryParamsResponse>): QueryParamsResponse {
    const message = { ...baseQueryParamsResponse } as QueryParamsResponse;
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromPartial(object.params);
    } else {
      message.params = undefined;
    }
    return message;
  },
};

const baseQueryGetNFTSchemaRequest: object = { code: "" };

export const QueryGetNFTSchemaRequest = {
  encode(
    message: QueryGetNFTSchemaRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.code !== "") {
      writer.uint32(10).string(message.code);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetNFTSchemaRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetNFTSchemaRequest,
    } as QueryGetNFTSchemaRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.code = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetNFTSchemaRequest {
    const message = {
      ...baseQueryGetNFTSchemaRequest,
    } as QueryGetNFTSchemaRequest;
    if (object.code !== undefined && object.code !== null) {
      message.code = String(object.code);
    } else {
      message.code = "";
    }
    return message;
  },

  toJSON(message: QueryGetNFTSchemaRequest): unknown {
    const obj: any = {};
    message.code !== undefined && (obj.code = message.code);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetNFTSchemaRequest>
  ): QueryGetNFTSchemaRequest {
    const message = {
      ...baseQueryGetNFTSchemaRequest,
    } as QueryGetNFTSchemaRequest;
    if (object.code !== undefined && object.code !== null) {
      message.code = object.code;
    } else {
      message.code = "";
    }
    return message;
  },
};

const baseQueryGetNFTSchemaResponse: object = {};

export const QueryGetNFTSchemaResponse = {
  encode(
    message: QueryGetNFTSchemaResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.nFTSchema !== undefined) {
      NFTSchema.encode(message.nFTSchema, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetNFTSchemaResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetNFTSchemaResponse,
    } as QueryGetNFTSchemaResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.nFTSchema = NFTSchema.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetNFTSchemaResponse {
    const message = {
      ...baseQueryGetNFTSchemaResponse,
    } as QueryGetNFTSchemaResponse;
    if (object.nFTSchema !== undefined && object.nFTSchema !== null) {
      message.nFTSchema = NFTSchema.fromJSON(object.nFTSchema);
    } else {
      message.nFTSchema = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetNFTSchemaResponse): unknown {
    const obj: any = {};
    message.nFTSchema !== undefined &&
      (obj.nFTSchema = message.nFTSchema
        ? NFTSchema.toJSON(message.nFTSchema)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetNFTSchemaResponse>
  ): QueryGetNFTSchemaResponse {
    const message = {
      ...baseQueryGetNFTSchemaResponse,
    } as QueryGetNFTSchemaResponse;
    if (object.nFTSchema !== undefined && object.nFTSchema !== null) {
      message.nFTSchema = NFTSchema.fromPartial(object.nFTSchema);
    } else {
      message.nFTSchema = undefined;
    }
    return message;
  },
};

const baseQueryAllNFTSchemaRequest: object = {};

export const QueryAllNFTSchemaRequest = {
  encode(
    message: QueryAllNFTSchemaRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryAllNFTSchemaRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllNFTSchemaRequest,
    } as QueryAllNFTSchemaRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.pagination = PageRequest.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllNFTSchemaRequest {
    const message = {
      ...baseQueryAllNFTSchemaRequest,
    } as QueryAllNFTSchemaRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllNFTSchemaRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageRequest.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllNFTSchemaRequest>
  ): QueryAllNFTSchemaRequest {
    const message = {
      ...baseQueryAllNFTSchemaRequest,
    } as QueryAllNFTSchemaRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryAllNFTSchemaResponse: object = {};

export const QueryAllNFTSchemaResponse = {
  encode(
    message: QueryAllNFTSchemaResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.nFTSchema) {
      NFTSchema.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(
        message.pagination,
        writer.uint32(18).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryAllNFTSchemaResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllNFTSchemaResponse,
    } as QueryAllNFTSchemaResponse;
    message.nFTSchema = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.nFTSchema.push(NFTSchema.decode(reader, reader.uint32()));
          break;
        case 2:
          message.pagination = PageResponse.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllNFTSchemaResponse {
    const message = {
      ...baseQueryAllNFTSchemaResponse,
    } as QueryAllNFTSchemaResponse;
    message.nFTSchema = [];
    if (object.nFTSchema !== undefined && object.nFTSchema !== null) {
      for (const e of object.nFTSchema) {
        message.nFTSchema.push(NFTSchema.fromJSON(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllNFTSchemaResponse): unknown {
    const obj: any = {};
    if (message.nFTSchema) {
      obj.nFTSchema = message.nFTSchema.map((e) =>
        e ? NFTSchema.toJSON(e) : undefined
      );
    } else {
      obj.nFTSchema = [];
    }
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageResponse.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllNFTSchemaResponse>
  ): QueryAllNFTSchemaResponse {
    const message = {
      ...baseQueryAllNFTSchemaResponse,
    } as QueryAllNFTSchemaResponse;
    message.nFTSchema = [];
    if (object.nFTSchema !== undefined && object.nFTSchema !== null) {
      for (const e of object.nFTSchema) {
        message.nFTSchema.push(NFTSchema.fromPartial(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

/** Query defines the gRPC querier service. */
export interface Query {
  /** Parameters queries the parameters of the module. */
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse>;
  /** Queries a NFTSchema by index. */
  NFTSchema(
    request: QueryGetNFTSchemaRequest
  ): Promise<QueryGetNFTSchemaResponse>;
  /** Queries a list of NFTSchema items. */
  NFTSchemaAll(
    request: QueryAllNFTSchemaRequest
  ): Promise<QueryAllNFTSchemaResponse>;
}

export class QueryClientImpl implements Query {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse> {
    const data = QueryParamsRequest.encode(request).finish();
    const promise = this.rpc.request("sixnft.nftmngr.Query", "Params", data);
    return promise.then((data) => QueryParamsResponse.decode(new Reader(data)));
  }

  NFTSchema(
    request: QueryGetNFTSchemaRequest
  ): Promise<QueryGetNFTSchemaResponse> {
    const data = QueryGetNFTSchemaRequest.encode(request).finish();
    const promise = this.rpc.request("sixnft.nftmngr.Query", "NFTSchema", data);
    return promise.then((data) =>
      QueryGetNFTSchemaResponse.decode(new Reader(data))
    );
  }

  NFTSchemaAll(
    request: QueryAllNFTSchemaRequest
  ): Promise<QueryAllNFTSchemaResponse> {
    const data = QueryAllNFTSchemaRequest.encode(request).finish();
    const promise = this.rpc.request(
      "sixnft.nftmngr.Query",
      "NFTSchemaAll",
      data
    );
    return promise.then((data) =>
      QueryAllNFTSchemaResponse.decode(new Reader(data))
    );
  }
}

interface Rpc {
  request(
    service: string,
    method: string,
    data: Uint8Array
  ): Promise<Uint8Array>;
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
