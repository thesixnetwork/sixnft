package keeper

import (
	"bytes"
	"context"
	"encoding/base64"
	"fmt"
	"strconv"

	"sixnft/x/nftoracle/types"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

func (k msgServer) CreateActionRequest(goCtx context.Context, msg *types.MsgCreateActionRequest) (*types.MsgCreateActionRequestResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	actionSigBz, err := base64.StdEncoding.DecodeString(msg.Base64ActionSignature)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrInvalidBase64, msg.Base64ActionSignature)
	}
	data := types.ActionSignature{}
	err = k.cdc.(*codec.ProtoCodec).UnmarshalJSON(actionSigBz, &data)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrParsingActionSignature, err.Error())
	}

	actionParam, err := k.ValidateActionSignature(data)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrVerifyingSignature, err.Error())
	}

	// Check if nft_schema_code exists
	_, found := k.nftmngrKeeper.GetNFTSchema(ctx, actionParam.NftSchemaCode)
	if !found {
		return nil, sdkerrors.Wrap(types.ErrNFTSchemaNotFound, actionParam.NftSchemaCode)
	}
	// Check if the token is already minted
	_, found = k.nftmngrKeeper.GetNftData(ctx, actionParam.NftSchemaCode, actionParam.TokenId)
	if found {
		return nil, sdkerrors.Wrap(types.ErrMetadataAlreadyExists, actionParam.NftSchemaCode)
	}

	createdAt := ctx.BlockTime()
	endTime := createdAt.Add(k.MintRequestActiveDuration(ctx))

	id_ := k.Keeper.AppendActionRequest(ctx, types.ActionRequest{
		NftSchemaCode:   actionParam.NftSchemaCode,
		TokenId:         actionParam.TokenId,
		RequiredConfirm: msg.RequiredConfirm,
		Status:          types.RequestStatus_PENDING,
		CurrentConfirm:  0,
		CreatedAt:       createdAt,
		ValidUntil:      endTime,
		Confirmers:      make(map[string]bool),
		DataHashes:      make([]*types.DataHash, 0),
	})

	k.Keeper.InsertActiveMintRequestQueue(ctx, id_, endTime)

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeMintRequestCreated,
			sdk.NewAttribute(types.AttributeKeyMintRequestID, strconv.FormatUint(id_, 10)),
			sdk.NewAttribute(types.AttributeKeyNftSchemaCode, actionParam.NftSchemaCode),
			sdk.NewAttribute(types.AttributeKeyTokenID, actionParam.TokenId),
			sdk.NewAttribute(types.AttributeKeyRequiredConfirm, strconv.FormatUint(msg.RequiredConfirm, 10)),
		),
	})

	return &types.MsgCreateActionRequestResponse{
		Id: id_,
	}, nil
}

func (k msgServer) ValidateActionSignature(actionSig types.ActionSignature) (*types.ActionParam, error) {
	sign_msg := "\x19Ethereum Signed Message:\n" + strconv.FormatInt(int64(len(actionSig.Message)), 10) + actionSig.Message

	data := []byte(sign_msg)
	hash := crypto.Keccak256Hash(data)
	var hash_bytes = hash.Bytes()

	actionParam := &types.ActionParam{}
	actionParamBz, err := base64.StdEncoding.DecodeString(actionSig.Message)
	if err != nil {
		return nil, err
	}
	err = k.cdc.(*codec.ProtoCodec).UnmarshalJSON(actionParamBz, actionParam)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrParsingActionParam, err.Error())
	}

	//validate signature format
	decode_signature, err := hexutil.Decode(actionSig.Signature)
	if err != nil {
		// log.Fatalf("Failed to decode signature: %v", msg.Signature)
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid signature")
	}
	signature_with_revocery_id := decode_signature
	// remove last byte coz is is recovery id
	decode_signature[64] -= 27 // this on should be checked whether it can be a weak point later // remove recovery id

	// get pulic key from signature
	sigPublicKey, err := crypto.Ecrecover(hash_bytes, decode_signature) //recover publickey from signature and hash
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid signature or message")
	}

	// get address from public key
	pubEDCA, err := crypto.UnmarshalPubkey(sigPublicKey)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "faild to unmarshal public key")
	}
	eth_address_from_pubkey := crypto.PubkeyToAddress(*pubEDCA)

	common_eth_address := common.HexToAddress(actionSig.Signer)
	if matches := bytes.Equal([]byte(eth_address_from_pubkey.Hex()), []byte(common_eth_address.Hex())); !matches {
		var ret = fmt.Sprintf("eth_address_from_pubkey: %s ,eth_address %s", eth_address_from_pubkey.Hex(), common_eth_address.Hex())
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, ret)
	}

	signatureNoRecoverID := signature_with_revocery_id[:len(signature_with_revocery_id)-1] // remove recovery id
	if verified := crypto.VerifySignature(sigPublicKey, hash.Bytes(), signatureNoRecoverID); !verified {
		return nil, sdkerrors.Wrap(types.ErrVerifyingSignature, "invalid signature")
	}
	return actionParam, nil
}
