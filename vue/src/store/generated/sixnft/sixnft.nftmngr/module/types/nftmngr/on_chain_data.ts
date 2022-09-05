/* eslint-disable */
import { AttributeDefinition } from "../nftmngr/attribute_definition";
import { Action } from "../nftmngr/action";
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "sixnft.nftmngr";

export interface OnChainData {
  nft_attributes: AttributeDefinition[];
  token_attributes: AttributeDefinition[];
  actions: Action[];
}

const baseOnChainData: object = {};

export const OnChainData = {
  encode(message: OnChainData, writer: Writer = Writer.create()): Writer {
    for (const v of message.nft_attributes) {
      AttributeDefinition.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    for (const v of message.token_attributes) {
      AttributeDefinition.encode(v!, writer.uint32(18).fork()).ldelim();
    }
    for (const v of message.actions) {
      Action.encode(v!, writer.uint32(26).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): OnChainData {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseOnChainData } as OnChainData;
    message.nft_attributes = [];
    message.token_attributes = [];
    message.actions = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.nft_attributes.push(
            AttributeDefinition.decode(reader, reader.uint32())
          );
          break;
        case 2:
          message.token_attributes.push(
            AttributeDefinition.decode(reader, reader.uint32())
          );
          break;
        case 3:
          message.actions.push(Action.decode(reader, reader.uint32()));
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
    message.nft_attributes = [];
    message.token_attributes = [];
    message.actions = [];
    if (object.nft_attributes !== undefined && object.nft_attributes !== null) {
      for (const e of object.nft_attributes) {
        message.nft_attributes.push(AttributeDefinition.fromJSON(e));
      }
    }
    if (
      object.token_attributes !== undefined &&
      object.token_attributes !== null
    ) {
      for (const e of object.token_attributes) {
        message.token_attributes.push(AttributeDefinition.fromJSON(e));
      }
    }
    if (object.actions !== undefined && object.actions !== null) {
      for (const e of object.actions) {
        message.actions.push(Action.fromJSON(e));
      }
    }
    return message;
  },

  toJSON(message: OnChainData): unknown {
    const obj: any = {};
    if (message.nft_attributes) {
      obj.nft_attributes = message.nft_attributes.map((e) =>
        e ? AttributeDefinition.toJSON(e) : undefined
      );
    } else {
      obj.nft_attributes = [];
    }
    if (message.token_attributes) {
      obj.token_attributes = message.token_attributes.map((e) =>
        e ? AttributeDefinition.toJSON(e) : undefined
      );
    } else {
      obj.token_attributes = [];
    }
    if (message.actions) {
      obj.actions = message.actions.map((e) =>
        e ? Action.toJSON(e) : undefined
      );
    } else {
      obj.actions = [];
    }
    return obj;
  },

  fromPartial(object: DeepPartial<OnChainData>): OnChainData {
    const message = { ...baseOnChainData } as OnChainData;
    message.nft_attributes = [];
    message.token_attributes = [];
    message.actions = [];
    if (object.nft_attributes !== undefined && object.nft_attributes !== null) {
      for (const e of object.nft_attributes) {
        message.nft_attributes.push(AttributeDefinition.fromPartial(e));
      }
    }
    if (
      object.token_attributes !== undefined &&
      object.token_attributes !== null
    ) {
      for (const e of object.token_attributes) {
        message.token_attributes.push(AttributeDefinition.fromPartial(e));
      }
    }
    if (object.actions !== undefined && object.actions !== null) {
      for (const e of object.actions) {
        message.actions.push(Action.fromPartial(e));
      }
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
