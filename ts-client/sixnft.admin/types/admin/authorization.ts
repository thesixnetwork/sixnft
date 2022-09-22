/* eslint-disable */
import { Permissions } from "../admin/permissions";
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "sixnft.admin";

export interface Authorization {
  rootAdmin: string;
  permissions: Permissions | undefined;
}

const baseAuthorization: object = { rootAdmin: "" };

export const Authorization = {
  encode(message: Authorization, writer: Writer = Writer.create()): Writer {
    if (message.rootAdmin !== "") {
      writer.uint32(10).string(message.rootAdmin);
    }
    if (message.permissions !== undefined) {
      Permissions.encode(
        message.permissions,
        writer.uint32(18).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): Authorization {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseAuthorization } as Authorization;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.rootAdmin = reader.string();
          break;
        case 2:
          message.permissions = Permissions.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Authorization {
    const message = { ...baseAuthorization } as Authorization;
    if (object.rootAdmin !== undefined && object.rootAdmin !== null) {
      message.rootAdmin = String(object.rootAdmin);
    } else {
      message.rootAdmin = "";
    }
    if (object.permissions !== undefined && object.permissions !== null) {
      message.permissions = Permissions.fromJSON(object.permissions);
    } else {
      message.permissions = undefined;
    }
    return message;
  },

  toJSON(message: Authorization): unknown {
    const obj: any = {};
    message.rootAdmin !== undefined && (obj.rootAdmin = message.rootAdmin);
    message.permissions !== undefined &&
      (obj.permissions = message.permissions
        ? Permissions.toJSON(message.permissions)
        : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<Authorization>): Authorization {
    const message = { ...baseAuthorization } as Authorization;
    if (object.rootAdmin !== undefined && object.rootAdmin !== null) {
      message.rootAdmin = object.rootAdmin;
    } else {
      message.rootAdmin = "";
    }
    if (object.permissions !== undefined && object.permissions !== null) {
      message.permissions = Permissions.fromPartial(object.permissions);
    } else {
      message.permissions = undefined;
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
