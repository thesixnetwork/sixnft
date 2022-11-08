package keeper

import (
	"context"
	"encoding/base64"
	"strconv"
	"time"

	"github.com/thesixnetwork/sixnft/x/nftoracle/types"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
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

	actionParam, signer, err := k.ValidateActionSignature(data)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrVerifyingSignature, err.Error())
	}

	// Check if nft_schema_code exists
	_, found := k.nftmngrKeeper.GetNFTSchema(ctx, actionParam.NftSchemaCode)
	if !found {
		return nil, sdkerrors.Wrap(types.ErrNFTSchemaNotFound, actionParam.NftSchemaCode)
	}
	// Check if the token is already Actioned
	_, found = k.nftmngrKeeper.GetNftData(ctx, actionParam.NftSchemaCode, actionParam.TokenId)
	if !found {
		return nil, sdkerrors.Wrap(types.ErrMetadataNotExists, actionParam.NftSchemaCode)
	}

	// Check action with reference exists
	if actionParam.RefId != "" {

		_, found := k.nftmngrKeeper.GetActionByRefId(ctx, actionParam.RefId)
		if found {
			return nil, sdkerrors.Wrap(types.ErrRefIdAlreadyExists, actionParam.RefId)
		}
	}

	oracleConfig, found := k.GetOracleConfig(ctx)
	if !found {
		return nil, sdkerrors.Wrap(types.ErrOracleConfigNotFound, "")
	}

	// Verify msg.RequiredConfirmations is less than or equal to oracleConfig.MinimumConfirmation
	if int32(msg.RequiredConfirm) < oracleConfig.MinimumConfirmation {
		return nil, sdkerrors.Wrap(types.ErrRequiredConfirmTooLess, strconv.Itoa(int(oracleConfig.MinimumConfirmation)))
	}

	createdAt := ctx.BlockTime()
	endTime := createdAt.Add(k.ActionRequestActiveDuration(ctx))

	// validate time given format as RFC3339
	_, err = time.Parse(time.RFC3339, actionParam.ExpiredAt.UTC().Format(time.RFC3339))
	if err != nil || len(actionParam.ExpiredAt.String()) == 0 || actionParam.ExpiredAt.Before(time.Now().UTC()){
		actionParam.ExpiredAt = endTime
	}

	if len(actionParam.ExpiredAt.String()) == 0 || actionParam.ExpiredAt.Before(time.Now().UTC()){
		actionParam.ExpiredAt = endTime
	}

	id_ := k.Keeper.AppendActionRequest(ctx, types.ActionRequest{
		NftSchemaCode:   actionParam.NftSchemaCode,
		TokenId:         actionParam.TokenId,
		RequiredConfirm: msg.RequiredConfirm,
		Caller:          *signer,
		Action:          actionParam.Action,
		RefId:           actionParam.RefId,
		Status:          types.RequestStatus_PENDING,
		CurrentConfirm:  0,
		CreatedAt:       createdAt,
		ValidUntil:      actionParam.ExpiredAt,
		Confirmers:      make([]string, 0),
		DataHashes:      make([]*types.DataHash, 0),
	})

	k.Keeper.InsertActiveActionRequestQueue(ctx, id_, actionParam.ExpiredAt)

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeActionRequestCreated,
			sdk.NewAttribute(types.AttributeKeyActionRequestID, strconv.FormatUint(id_, 10)),
			sdk.NewAttribute(types.AttributeKeyNftSchemaCode, actionParam.NftSchemaCode),
			sdk.NewAttribute(types.AttributeKeyTokenID, actionParam.TokenId),
			sdk.NewAttribute(types.AttributeKeyRequiredConfirm, strconv.FormatUint(msg.RequiredConfirm, 10)),
		),
	})

	return &types.MsgCreateActionRequestResponse{
		Id: id_,
	}, nil
}

func (k msgServer) ValidateActionSignature(actionSig types.ActionSignature) (*types.ActionParam, *string, error) {

	sign_msg := "\x19Ethereum Signed Message:\n" + strconv.FormatInt(int64(len(actionSig.Message)), 10) + actionSig.Message

	data := []byte(sign_msg)
	hash := crypto.Keccak256Hash(data)
	var hash_bytes = hash.Bytes()

	actionParam := &types.ActionParam{}
	actionParamBz, err := base64.StdEncoding.DecodeString(actionSig.Message)
	if err != nil {
		return nil, nil, err
	}
	err = k.cdc.(*codec.ProtoCodec).UnmarshalJSON(actionParamBz, actionParam)
	if err != nil {
		return nil, nil, sdkerrors.Wrap(types.ErrParsingActionParam, err.Error())
	}

	//validate signature format
	decode_signature, err := hexutil.Decode(actionSig.Signature)
	if err != nil {
		// log.Fatalf("Failed to decode signature: %v", msg.Signature)
		return nil, nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid signature")
	}
	signature_with_revocery_id := decode_signature
	// remove last byte coz is is recovery id

	// get pulic key from signature
	sigPublicKey, err := crypto.Ecrecover(hash_bytes, decode_signature) //recover publickey from signature and hash
	if err != nil {
		return nil, nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid signature or message")
	}

	// fmt.Println("sigPublicKey:", hexutil.Encode(sigPublicKey))
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
	return actionParam, &signer, nil
}
