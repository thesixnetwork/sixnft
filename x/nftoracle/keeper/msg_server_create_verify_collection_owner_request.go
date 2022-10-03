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

	// Check if nft_schema_code exists
	_schema, found := k.nftmngrKeeper.GetNFTSchema(ctx, msg.NftSchemaCode)
	if !found {
		return nil, sdkerrors.Wrap(types.ErrNFTSchemaNotFound, msg.NftSchemaCode)
	}

	// check if creator is owner of the collection
	if _schema.Owner != msg.Creator {
		return nil, sdkerrors.Wrap(types.ErrNotCollectionOwner, msg.Creator)
	}

	// Check if nft_schema_code already verified
	if _schema.IsVerified {
		return nil, sdkerrors.Wrap(types.ErrNFTSchemaAlreadyVerified, msg.NftSchemaCode)
	}

	collectionOwnerBz, err := base64.StdEncoding.DecodeString(msg.Base64VerifierSignature)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrInvalidBase64, msg.Base64VerifierSignature)
	}
	data := types.CollectionOwnerSignature{}
	err = k.cdc.(*codec.ProtoCodec).UnmarshalJSON(collectionOwnerBz, &data)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrParsingCollectionOwnerSignature, err.Error())
	}
	signer, err := k.ValidateCollectionOwnerSignature(data)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrVerifyingSignature, err.Error())
	}

	createdAt := ctx.BlockTime()
	endTime := createdAt.Add(k.MintRequestActiveDuration(ctx))
	
	id_ := k.Keeper.AppendCollectionOwnerRequest(ctx, types.CollectionOwnerRequest{
		NftSchemaCode:   msg.NftSchemaCode,
		Signer: 		*signer,
		RequiredConfirm: msg.RequiredConfirm,
		Status:          types.RequestStatus_PENDING,
		CurrentConfirm:  0,
		CreatedAt:       createdAt,
		ValidUntil:      endTime,
		Confirmers:      make(map[string]bool),
		DataHashes:      make([]*types.DataHash, 0),
	})


	return &types.MsgCreateVerifyCollectionOwnerRequestResponse{
		Id: id_,
		NftSchemaCode: _schema.Code,
		OwnerAddress: *signer,
	}, nil
}

func (k msgServer) ValidateCollectionOwnerSignature(collectionOwnerSig types.CollectionOwnerSignature) (*string, error) {

	sign_msg := "\x19Ethereum Signed Message:\n" + strconv.FormatInt(int64(len(collectionOwnerSig.Message)), 10) + collectionOwnerSig.Message

	data := []byte(sign_msg)
	hash := crypto.Keccak256Hash(data)
	var hash_bytes = hash.Bytes()

	collectionOwnerType := &types.CollectionOwnerSignature{}
	collectionOwnerTypeBz, err := base64.StdEncoding.DecodeString(collectionOwnerSig.Message)
	if err != nil {
		return nil,  err
	}
	err = k.cdc.(*codec.ProtoCodec).UnmarshalJSON(collectionOwnerTypeBz, collectionOwnerType)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrParsingActionParam, err.Error())
	}

	//validate signature format
	decode_signature, err := hexutil.Decode(collectionOwnerSig.Signature)
	if err != nil {
		// log.Fatalf("Failed to decode signature: %v", msg.Signature)
		return nil,  sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid signature")
	}
	signature_with_revocery_id := decode_signature

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

	signatureNoRecoverID := signature_with_revocery_id[:len(signature_with_revocery_id)-1] // remove recovery id
	if verified := crypto.VerifySignature(sigPublicKey, hash.Bytes(), signatureNoRecoverID); !verified {
		return  nil, sdkerrors.Wrap(types.ErrVerifyingSignature, "invalid signature")
	}
	signer := eth_address_from_pubkey.Hex()
	// fmt.Println("######################### signer", signer)
	return  &signer, nil
}
