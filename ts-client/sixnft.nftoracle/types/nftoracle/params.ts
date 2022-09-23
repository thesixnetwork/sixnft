/* eslint-disable */
import { Duration } from "../google/protobuf/duration";
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "sixnft.nftoracle";

/** Params defines the parameters for the module. */
export interface Params {
  mintRequestActiveDuration: Duration | undefined;
  actionRequestActiveDuration: Duration | undefined;
}

const baseParams: object = {};

export const Params = {
  encode(message: Params, writer: Writer = Writer.create()): Writer {
    if (message.mintRequestActiveDuration !== undefined) {
      Duration.encode(
        message.mintRequestActiveDuration,
        writer.uint32(10).fork()
      ).ldelim();
    }
    if (message.actionRequestActiveDuration !== undefined) {
      Duration.encode(
        message.actionRequestActiveDuration,
        writer.uint32(18).fork()
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
          message.mintRequestActiveDuration = Duration.decode(
            reader,
            reader.uint32()
          );
          break;
        case 2:
          message.actionRequestActiveDuration = Duration.decode(
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
      object.mintRequestActiveDuration !== undefined &&
      object.mintRequestActiveDuration !== null
    ) {
      message.mintRequestActiveDuration = Duration.fromJSON(
        object.mintRequestActiveDuration
      );
    } else {
      message.mintRequestActiveDuration = undefined;
    }
    if (
      object.actionRequestActiveDuration !== undefined &&
      object.actionRequestActiveDuration !== null
    ) {
      message.actionRequestActiveDuration = Duration.fromJSON(
        object.actionRequestActiveDuration
      );
    } else {
      message.actionRequestActiveDuration = undefined;
    }
    return message;
  },

  toJSON(message: Params): unknown {
    const obj: any = {};
    message.mintRequestActiveDuration !== undefined &&
      (obj.mintRequestActiveDuration = message.mintRequestActiveDuration
        ? Duration.toJSON(message.mintRequestActiveDuration)
        : undefined);
    message.actionRequestActiveDuration !== undefined &&
      (obj.actionRequestActiveDuration = message.actionRequestActiveDuration
        ? Duration.toJSON(message.actionRequestActiveDuration)
        : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<Params>): Params {
    const message = { ...baseParams } as Params;
    if (
      object.mintRequestActiveDuration !== undefined &&
      object.mintRequestActiveDuration !== null
    ) {
      message.mintRequestActiveDuration = Duration.fromPartial(
        object.mintRequestActiveDuration
      );
    } else {
      message.mintRequestActiveDuration = undefined;
    }
    if (
      object.actionRequestActiveDuration !== undefined &&
      object.actionRequestActiveDuration !== null
    ) {
      message.actionRequestActiveDuration = Duration.fromPartial(
        object.actionRequestActiveDuration
      );
    } else {
      message.actionRequestActiveDuration = undefined;
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
