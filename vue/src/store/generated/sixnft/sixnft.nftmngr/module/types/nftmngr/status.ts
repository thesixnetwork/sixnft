/* eslint-disable */
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "sixnft.nftmngr";

export interface Status {
  firstMintComplete: string;
}

const baseStatus: object = { firstMintComplete: "" };

export const Status = {
  encode(message: Status, writer: Writer = Writer.create()): Writer {
    if (message.firstMintComplete !== "") {
      writer.uint32(10).string(message.firstMintComplete);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): Status {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseStatus } as Status;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.firstMintComplete = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Status {
    const message = { ...baseStatus } as Status;
    if (
      object.firstMintComplete !== undefined &&
      object.firstMintComplete !== null
    ) {
      message.firstMintComplete = String(object.firstMintComplete);
    } else {
      message.firstMintComplete = "";
    }
    return message;
  },

  toJSON(message: Status): unknown {
    const obj: any = {};
    message.firstMintComplete !== undefined &&
      (obj.firstMintComplete = message.firstMintComplete);
    return obj;
  },

  fromPartial(object: DeepPartial<Status>): Status {
    const message = { ...baseStatus } as Status;
    if (
      object.firstMintComplete !== undefined &&
      object.firstMintComplete !== null
    ) {
      message.firstMintComplete = object.firstMintComplete;
    } else {
      message.firstMintComplete = "";
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
