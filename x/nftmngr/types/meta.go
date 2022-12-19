package types

import (
	"regexp"
	"strconv"
	"time"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

type MetadataChange struct {
	Key           string `json:"key,omitempty"`
	PreviousValue string `json:"previous_value,omitempty"`
	NewValue      string `json:"new_value,omitempty"`
}

type Metadata struct {
	nftData                *NftData
	schema                 *NFTSchema
	attributeOverring      AttributeOverriding
	ChangeList             []*MetadataChange `json:"change_list,omitempty"`
	MapAllKey              map[string]*MetadataAttribute
	OtherUpdatedTokenDatas map[string]*NftData
	NftDataFunction        func(tokenId string) (*NftData, error)
	GetBlockTimeFunction   func() time.Time
	GetBlockHeightFunction func() int64
}

type MetadataAttribute struct {
	AttributeValue *NftAttributeValue
	From           string
	Index          int
}

// var MapAllKey = map[string]*MetadataAttribute{}

func NewMetadata(schema *NFTSchema, tokenData *NftData, attributeOverring AttributeOverriding) *Metadata {
	meta := &Metadata{
		nftData:           tokenData,
		schema:            schema,
		attributeOverring: attributeOverring,
		ChangeList:        []*MetadataChange{},
	}

	meta.MapAllKey = map[string]*MetadataAttribute{}
	meta.OtherUpdatedTokenDatas = map[string]*NftData{}

	// Parse the metadata
	for i, attri := range schema.GetOnchainData().NftAttributesValue {
		meta.MapAllKey[attri.Name] = &MetadataAttribute{
			AttributeValue: attri,
			From:           "nft",
			Index:          i,
		}
	}
	for i, attri := range tokenData.OriginAttributes {
		meta.MapAllKey[attri.Name] = &MetadataAttribute{
			AttributeValue: attri,
			From:           "origin",
			Index:          i,
		}
	}
	for i, attri := range tokenData.OnchainAttributes {
		if _, ok := meta.MapAllKey[attri.Name]; ok {
			if attributeOverring == AttributeOverriding_CHAIN {
				meta.MapAllKey[attri.Name] = &MetadataAttribute{
					AttributeValue: attri,
					From:           "chain",
					Index:          i,
				}
			}
		} else {
			meta.MapAllKey[attri.Name] = &MetadataAttribute{
				AttributeValue: attri,
				From:           "chain",
				Index:          i,
			}
		}

	}
	return meta
}

func (m *Metadata) SetGetBlockHeightFunction(f func() int64) {
	m.GetBlockHeightFunction = f
}

func (m *Metadata) SetGetBlockTimeFunction(f func() time.Time) {
	m.GetBlockTimeFunction = f
}

func (m *Metadata) SetGetNFTFunction(f func(tokenId string) (*NftData, error)) {
	m.NftDataFunction = f
}

func (m *Metadata) CurrentTimestamp() int64 {
	return time.Now().Unix()
}

func (m *Metadata) GetBaseURI() string {
	return m.schema.OriginData.OriginBaseUri
}

func (m *Metadata) GetTokenURI() string {
	return m.nftData.TokenUri
}

func (m *Metadata) SetTokenURI(uri string) {
	m.ChangeList = append(m.ChangeList, &MetadataChange{
		Key:           "tokenURI",
		PreviousValue: m.nftData.TokenUri,
		NewValue:      uri,
	})
	m.nftData.TokenUri = uri
}

func (m *Metadata) GetImage() string {
	if m.nftData.OnchainImage != "" {
		return m.nftData.OnchainImage
	}
	return m.nftData.OriginImage
}

func (m *Metadata) SetImage(image string) {
	currentImage := m.nftData.OnchainImage
	if currentImage == "" {
		currentImage = m.nftData.OriginImage
	}
	m.ChangeList = append(m.ChangeList, &MetadataChange{
		Key:           "image",
		PreviousValue: currentImage,
		NewValue:      image,
	})
	m.nftData.OnchainImage = image
}

func (m *Metadata) GetNumber(key string) int64 {
	v, err := m.MustGetNumber(key)
	if err != nil {
		panic(err)
	}
	return v
}

func (m *Metadata) MustGetNumber(key string) (int64, error) {
	attri := m.MapAllKey[key]
	if attri == nil {
		return 0, sdkerrors.Wrap(ErrAttributeNotFoundForAction, key)
	}
	if _, ok := attri.AttributeValue.GetValue().(*NftAttributeValue_NumberAttributeValue); ok {
		// Number
		return int64(attri.AttributeValue.GetNumberAttributeValue().Value), nil
	}
	return 0, sdkerrors.Wrap(ErrAttributeTypeNotMatch, attri.AttributeValue.Name)
}

func (m *Metadata) SetNumber(key string, value int64) error {
	// m.mapNumber[key] = value
	attri := m.MapAllKey[key]
	if attri == nil {
		panic(sdkerrors.Wrap(ErrAttributeNotFoundForAction, key))
	}
	if _, ok := attri.AttributeValue.GetValue().(*NftAttributeValue_NumberAttributeValue); ok {
		// Number
		newAttributeValue := &NftAttributeValue{
			Name: attri.AttributeValue.Name,
			Value: &NftAttributeValue_NumberAttributeValue{
				NumberAttributeValue: &NumberAttributeValue{
					Value: uint64(value),
				},
			},
		}
		if attri.From == "chain" {
			m.ChangeList = append(m.ChangeList, &MetadataChange{
				Key:           key,
				PreviousValue: strconv.FormatUint(attri.AttributeValue.GetNumberAttributeValue().Value, 10),
				NewValue:      strconv.FormatUint(uint64(value), 10),
			})
			m.MapAllKey[key].AttributeValue = newAttributeValue
			m.nftData.OnchainAttributes[attri.Index] = newAttributeValue
		} else {
			return sdkerrors.Wrap(ErrAttributeOverriding, "can not override the origin attribute or nft attribute")
		}
	} else {
		return sdkerrors.Wrap(ErrAttributeTypeNotMatch, attri.AttributeValue.Name)
	}
	return nil
}
func (m *Metadata) GetString(key string) string {
	v, err := m.MustGetString(key)
	if err != nil {
		panic(err)
	}
	return v
}

func (m *Metadata) MustGetString(key string) (string, error) {
	attri := m.MapAllKey[key]
	if attri == nil {
		return "", sdkerrors.Wrap(ErrAttributeNotFoundForAction, key)
	}
	if _, ok := attri.AttributeValue.GetValue().(*NftAttributeValue_StringAttributeValue); ok {
		// Number
		return attri.AttributeValue.GetStringAttributeValue().Value, nil
	}
	return "", sdkerrors.Wrap(ErrAttributeTypeNotMatch, attri.AttributeValue.Name)
}

func (m *Metadata) SetString(key string, value string) error {
	// m.mapString[key] = value
	attri := m.MapAllKey[key]
	if attri == nil {
		panic(sdkerrors.Wrap(ErrAttributeNotFoundForAction, key))
	}
	if _, ok := attri.AttributeValue.GetValue().(*NftAttributeValue_StringAttributeValue); ok {
		// Number
		newAttributeValue := &NftAttributeValue{
			Name: attri.AttributeValue.Name,
			Value: &NftAttributeValue_StringAttributeValue{
				StringAttributeValue: &StringAttributeValue{
					Value: value,
				},
			},
		}
		if attri.From == "chain" {
			m.ChangeList = append(m.ChangeList, &MetadataChange{
				Key:           key,
				PreviousValue: attri.AttributeValue.GetStringAttributeValue().Value,
				NewValue:      value,
			})
			m.MapAllKey[key].AttributeValue = newAttributeValue
			m.nftData.OnchainAttributes[attri.Index] = newAttributeValue
		} else {
			return sdkerrors.Wrap(ErrAttributeOverriding, "can not override the origin attribute")
		}
	} else {
		return sdkerrors.Wrap(ErrAttributeTypeNotMatch, attri.AttributeValue.Name)
	}
	return nil
}

func (m *Metadata) GetFloat(key string) float64 {
	v, err := m.MustGetFloat(key)
	if err != nil {
		panic(err)
	}
	return v
}

func (m *Metadata) MustGetFloat(key string) (float64, error) {
	// return m.mapFloat[key]
	attri := m.MapAllKey[key]
	if attri == nil {
		return 0, sdkerrors.Wrap(ErrAttributeNotFoundForAction, key)
	}
	if _, ok := attri.AttributeValue.GetValue().(*NftAttributeValue_FloatAttributeValue); ok {
		// Number
		return attri.AttributeValue.GetFloatAttributeValue().Value, nil
	}
	return 0, sdkerrors.Wrap(ErrAttributeTypeNotMatch, attri.AttributeValue.Name)
}

func (m *Metadata) SetFloat(key string, value float64) error {
	// m.mapFloat[key] = value
	attri := m.MapAllKey[key]
	if attri == nil {
		panic(sdkerrors.Wrap(ErrAttributeNotFoundForAction, key))
	}
	if _, ok := attri.AttributeValue.GetValue().(*NftAttributeValue_FloatAttributeValue); ok {
		// Number
		newAttributeValue := &NftAttributeValue{
			Name: attri.AttributeValue.Name,
			Value: &NftAttributeValue_FloatAttributeValue{
				FloatAttributeValue: &FloatAttributeValue{
					Value: value,
				},
			},
		}
		if attri.From == "chain" {
			m.ChangeList = append(m.ChangeList, &MetadataChange{
				Key:           key,
				PreviousValue: strconv.FormatFloat(attri.AttributeValue.GetFloatAttributeValue().Value, 'f', -1, 64),
				NewValue:      strconv.FormatFloat(value, 'f', -1, 64),
			})
			m.MapAllKey[key].AttributeValue = newAttributeValue
			m.nftData.OnchainAttributes[attri.Index] = newAttributeValue
		} else {
			return sdkerrors.Wrap(ErrAttributeOverriding, "can not override the origin attribute")
		}
	} else {
		return sdkerrors.Wrap(ErrAttributeTypeNotMatch, attri.AttributeValue.Name)
	}
	return nil
}

func (m *Metadata) GetBoolean(key string) bool {
	v, err := m.MustGetBool(key)
	if err != nil {
		panic(err)
	}
	return v
}

func (m *Metadata) MustGetBool(key string) (bool, error) {
	// return m.mapBool[key]
	attri := m.MapAllKey[key]
	if attri == nil {
		return false, sdkerrors.Wrap(ErrAttributeNotFoundForAction, key)
	}
	if _, ok := attri.AttributeValue.GetValue().(*NftAttributeValue_BooleanAttributeValue); ok {
		// Number
		return attri.AttributeValue.GetBooleanAttributeValue().Value, nil
	}
	return false, sdkerrors.Wrap(ErrAttributeTypeNotMatch, attri.AttributeValue.Name)
}

func (m *Metadata) SetBoolean(key string, value bool) error {
	// m.mapBool[key] = value
	attri := m.MapAllKey[key]
	if attri == nil {
		panic(sdkerrors.Wrap(ErrAttributeNotFoundForAction, key))
	}
	if _, ok := attri.AttributeValue.GetValue().(*NftAttributeValue_BooleanAttributeValue); ok {
		// Number
		newAttributeValue := &NftAttributeValue{
			Name: attri.AttributeValue.Name,
			Value: &NftAttributeValue_BooleanAttributeValue{
				BooleanAttributeValue: &BooleanAttributeValue{
					Value: value,
				},
			},
		}
		if attri.From == "chain" {
			m.ChangeList = append(m.ChangeList, &MetadataChange{
				Key:           key,
				PreviousValue: strconv.FormatBool(attri.AttributeValue.GetBooleanAttributeValue().Value),
				NewValue:      strconv.FormatBool(value),
			})
			m.MapAllKey[key].AttributeValue = newAttributeValue
			m.nftData.OnchainAttributes[attri.Index] = newAttributeValue
		} else {
			return sdkerrors.Wrap(ErrAttributeOverriding, "can not override the origin attribute")
		}
	} else {
		return sdkerrors.Wrap(ErrAttributeTypeNotMatch, attri.AttributeValue.Name)
	}
	return nil
}

func (m *Metadata) SetDisplayArribute(key string, value string) error {
	bool_val, _ := strconv.ParseBool(value)
	attri := m.MapAllKey[key]
	if attri == nil {
		return sdkerrors.Wrap(ErrAttributeNotFoundForAction, key)
	}
	if _bool := attri.AttributeValue.GetHiddenToMarketplace() == bool_val; _bool {
		return nil
	}
	if attri.From == "chain" {
		newAttributeValue := &NftAttributeValue{
			Name:                attri.AttributeValue.Name,
			HiddenToMarketplace: bool_val,
		}

		m.ChangeList = append(m.ChangeList, &MetadataChange{
			Key:           key,
			PreviousValue: strconv.FormatBool(attri.AttributeValue.GetHiddenToMarketplace()),
			NewValue:      strconv.FormatBool(bool_val),
		})
		m.MapAllKey[key].AttributeValue = newAttributeValue
		m.nftData.OnchainAttributes[attri.Index] = newAttributeValue
	} else {
		return sdkerrors.Wrap(ErrAttributeOverriding, "can not override the origin attribute")
	}
	return nil
}

func (m *Metadata) ReplaceAllString(intput string, regexpStr string, replaceStr string) string {
	reg := regexp.MustCompile(regexpStr)
	return reg.ReplaceAllString(intput, replaceStr)
}

func (p *ActionParameter) MustGetNumber(key string) (uint64, error) {
	v, err := strconv.ParseUint(p.Value, 10, 64)
	if err != nil {
		return 0, sdkerrors.Wrap(ErrAttributeTypeNotMatch, key)
	}
	return v, nil
}

func (p *ActionParameter) GetNumber() uint64 {
	v, err := p.MustGetNumber(p.Name)
	if err != nil {
		panic(err)
	}
	return v
}

func (p *ActionParameter) GetString() string {
	return p.Value
}
