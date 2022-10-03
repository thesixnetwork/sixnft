package keeper

import (
	"context"
	"encoding/base64"
	"strconv"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/thesixnetwork/sixnft/x/nftoracle/types"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

func (k msgServer) CreateVerifyCollectionOwnerRequest(goCtx context.Context, msg *types.MsgCreateVerifyCollectionOwnerRequest) (*types.MsgCreateVerifyCollectionOwnerRequestResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	collectionOwnerBz, err := base64.StdEncoding.DecodeString(msg.Base64VerifierSignature)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrInvalidBase64, msg.Base64VerifierSignature)
	}
	data := types.CollectionOwnerSignature{}
	err = k.cdc.(*codec.ProtoCodec).UnmarshalJSON(collectionOwnerBz, &data)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrParsingCollectionOwnerSignature, err.Error())
	}
	verifyParam, signer, err := k.ValidateCollectionOwnerSignature(data)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrVerifyingSignature, err.Error())
	}

	// Check if nft_schema_code exists
	_schema, found := k.nftmngrKeeper.GetNFTSchema(ctx, msg.NftSchemaCode)
	if !found {
		return nil, sdkerrors.Wrap(types.ErrNFTSchemaNotFound, msg.NftSchemaCode)
	}
	
	// Check if nft_schema_code already verified
	if _schema.IsVerified {
		return nil, sdkerrors.Wrap(types.ErrNFTSchemaAlreadyVerified, msg.NftSchemaCode)
	}

	// check if signer is collection owner
	if _schema.Owner != *signer {
		return nil, sdkerrors.Wrap(types.ErrNotCollectionOwner, _schema.CollectionOwner.String())
	}

	

	// TODO: Handling the message
	_ = ctx

	return &types.MsgCreateVerifyCollectionOwnerRequestResponse{}, nil
}

func (k msgServer) ValidateCollectionOwnerSignature(collectionOwnerSig types.CollectionOwnerSignature) (*types.ActionParam, *string, error) {

	sign_msg := "\x19Ethereum Signed Message:\n" + strconv.FormatInt(int64(len(collectionOwnerSig.Message)), 10) + collectionOwnerSig.Message
	// fmt.Println("sign_msg: ", sign_msg)

	data := []byte(sign_msg)
	hash := crypto.Keccak256Hash(data)
	var hash_bytes = hash.Bytes()

	// fmt.Println("data: ", hexutil.Encode(data))

	actionParam := &types.ActionParam{}
	actionParamBz, err := base64.StdEncoding.DecodeString(collectionOwnerSig.Message)
	if err != nil {
		return nil, nil, err
	}
	err = k.cdc.(*codec.ProtoCodec).UnmarshalJSON(actionParamBz, actionParam)
	if err != nil {
		return nil, nil, sdkerrors.Wrap(types.ErrParsingActionParam, err.Error())
	}

	//validate signature format
	decode_signature, err := hexutil.Decode(collectionOwnerSig.Signature)
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
	// fmt.Println("######################### signer", signer)
	return actionParam, &signer, nil
}
