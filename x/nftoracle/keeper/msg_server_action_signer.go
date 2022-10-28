package keeper

import (
	"context"
	"encoding/base64"
	"strconv"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/thesixnetwork/sixnft/x/nftoracle/types"
	nftadmintypes "github.com/thesixnetwork/sixnft/x/nftadmin/types"
)

func (k msgServer) CreateActionSigner(goCtx context.Context, msg *types.MsgCreateActionSigner) (*types.MsgCreateActionSignerResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	binder, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	//check if signer is authorized
	isBinder := k.nftadminKeeper.HasPermission(ctx, nftadmintypes.KeyPermissionBinder, binder)
	if !isBinder {
		return nil, nftadmintypes.ErrUnauthorized
	}

	signerActionBz, err := base64.StdEncoding.DecodeString(msg.Base64EncodedSetSignerAction)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrInvalidBase64, msg.Base64EncodedSetSignerAction)
	}

	data := types.SetSignerSignature{}
	err = k.cdc.(*codec.ProtoCodec).UnmarshalJSON(signerActionBz, &data)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrParsingBase64, err.Error())
	}
	_signerParams, signer, err := k.ValidatesetSignernature(data)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrVerifyingSignature, err.Error())
	}

	// Check if the value already exists
	_, isFound := k.GetActionSigner(
		ctx,
		_signerParams.ActorAddress,
	)

	if isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "index already set")
	}

	createdAt := ctx.BlockTime()

	var actionSigner = types.ActionSigner{
		Creator:      msg.Creator,
		ActorAddress: _signerParams.ActorAddress,
		OwnerAddress: *signer,
		CreatedAt:    createdAt,
		ExpiredAt:    _signerParams.ExpiredAt,
	}

	k.SetActionSigner(
		ctx,
		actionSigner,
	)
	return &types.MsgCreateActionSignerResponse{}, nil
}

func (k msgServer) UpdateActionSigner(goCtx context.Context, msg *types.MsgUpdateActionSigner) (*types.MsgUpdateActionSignerResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	binder, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, err
	}

	//check if signer is authorized
	isBinder := k.nftadminKeeper.HasPermission(ctx, nftadmintypes.KeyPermissionBinder, binder)
	if !isBinder {
		return nil, nftadmintypes.ErrUnauthorized
	}

	signerActionBz, err := base64.StdEncoding.DecodeString(msg.Base64EncodedSetSignerAction)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrInvalidBase64, msg.Base64EncodedSetSignerAction)
	}

	data := types.SetSignerSignature{}
	err = k.cdc.(*codec.ProtoCodec).UnmarshalJSON(signerActionBz, &data)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrParsingSetSignerSignature, err.Error())
	}
	_signerParams, signer, err := k.ValidatesetSignernature(data)
	if err != nil {
		return nil, err
	}

	// Check if the value exists
	valFound, isFound := k.GetActionSigner(
		ctx,
		_signerParams.ActorAddress,
	)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	// Checks if the the msg creator is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	updatedAt := ctx.BlockTime()

	var actionSigner = types.ActionSigner{
		Creator:      msg.Creator,
		ActorAddress: _signerParams.ActorAddress,
		OwnerAddress: *signer,
		CreatedAt:    updatedAt,
		ExpiredAt:    _signerParams.ExpiredAt,
	}

	k.SetActionSigner(ctx, actionSigner)

	return &types.MsgUpdateActionSignerResponse{}, nil
}

func (k msgServer) DeleteActionSigner(goCtx context.Context, msg *types.MsgDeleteActionSigner) (*types.MsgDeleteActionSignerResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	binder, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, err
	}

	//check if signer is authorized
	isBinder := k.nftadminKeeper.HasPermission(ctx, nftadmintypes.KeyPermissionBinder, binder)
	if !isBinder {
		return nil, nftadmintypes.ErrUnauthorized
	}

	signerActionBz, err := base64.StdEncoding.DecodeString(msg.Base64EncodedSetSignerAction)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrInvalidBase64, msg.Base64EncodedSetSignerAction)
	}

	data := types.SetSignerSignature{}
	err = k.cdc.(*codec.ProtoCodec).UnmarshalJSON(signerActionBz, &data)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrParsingSetSignerSignature, err.Error())
	}
	_signerParams, _, err := k.ValidatesetSignernature(data)
	if err != nil {
		return nil, err
	}

	// Check if the value exists
	valFound, isFound := k.GetActionSigner(
		ctx,
		_signerParams.ActorAddress,
	)

	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	// Checks if the the msg creator is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveActionSigner(
		ctx,
		_signerParams.ActorAddress,
	)

	return &types.MsgDeleteActionSignerResponse{}, nil
}

func (k msgServer) ValidatesetSignernature(setSigner types.SetSignerSignature) (*types.SetSignerParams, *string, error) {

	sign_msg := "\x19Ethereum Signed Message:\n" + strconv.FormatInt(int64(len(setSigner.Message)), 10) + setSigner.Message

	data := []byte(sign_msg)
	hash := crypto.Keccak256Hash(data)
	var hash_bytes = hash.Bytes()

	setSignerParams := &types.SetSignerParams{}
	setSignerTypeBz, err := base64.StdEncoding.DecodeString(setSigner.Message)
	if err != nil {
		return nil, nil, err
	}
	err = k.cdc.(*codec.ProtoCodec).UnmarshalJSON(setSignerTypeBz, setSignerParams)
	if err != nil {
		return nil, nil, sdkerrors.Wrap(types.ErrParsingSetSignerSignature, err.Error())
	}

	//validate signature format
	decode_signature, err := hexutil.Decode(setSigner.Signature)
	if err != nil {
		// log.Fatalf("Failed to decode signature: %v", msg.Signature)
		return nil, nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid signature")
	}
	signature_with_revocery_id := decode_signature

	// get pulic key from signature
	sigPublicKey, err := crypto.Ecrecover(hash_bytes, decode_signature) //recover publickey from signature and hash
	if err != nil {
		return nil, nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid signature or message")
	}
	// get address from public key
	pubEDCA, err := crypto.UnmarshalPubkey(sigPublicKey)
	if err != nil {
		return nil, nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "faild to unmarshal public key")
	}
	eth_address_from_pubkey := crypto.PubkeyToAddress(*pubEDCA)

	signatureNoRecoverID := signature_with_revocery_id[:len(signature_with_revocery_id)-1] // remove recovery id
	if verified := crypto.VerifySignature(sigPublicKey, hash.Bytes(), signatureNoRecoverID); !verified {
		return nil, nil, sdkerrors.Wrap(types.ErrVerifyingSignature, "invalid signature")
	}
	signer := eth_address_from_pubkey.Hex()
	return setSignerParams, &signer, nil
}
