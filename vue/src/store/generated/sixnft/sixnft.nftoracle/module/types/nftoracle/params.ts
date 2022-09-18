/* eslint-disable */
import { Duration } from "../google/protobuf/duration";
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "sixnft.nftoracle";

/** Params defines the parameters for the module. */
export interface Params {
  mint_request_active_duration: Duration | undefined;
}

const baseParams: object = {};

export const Params = {
  encode(message: Params, writer: Writer = Writer.create()): Writer {
    if (message.mint_request_active_duration !== undefined) {
      Duration.encode(
        message.mint_request_active_duration,
        writer.uint32(10).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): Params {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseParams } as Params;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.mint_request_active_duration = Duration.decode(
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

  fromJSON(object: any): Params {
    const message = { ...baseParams } as Params;
    if (
      object.mint_request_active_duration !== undefined &&
      object.mint_request_active_duration !== null
    ) {
      message.mint_request_active_duration = Duration.fromJSON(
        object.mint_request_active_duration
      );
    } else {
      message.mint_request_active_duration = undefined;
    }
    return message;
  },

  toJSON(message: Params): unknown {
    const obj: any = {};
    message.mint_request_active_duration !== undefined &&
      (obj.mint_request_active_duration = message.mint_request_active_duration
        ? Duration.toJSON(message.mint_request_active_duration)
        : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<Params>): Params {
    const message = { ...baseParams } as Params;
    if (
      object.mint_request_active_duration !== undefined &&
      object.mint_request_active_duration !== null
    ) {
      message.mint_request_active_duration = Duration.fromPartial(
        object.mint_request_active_duration
      );
    } else {
      message.mint_request_active_duration = undefined;
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
