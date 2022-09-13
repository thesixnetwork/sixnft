package types

import (
	"regexp"
	"strconv"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

type MetadataChange struct {
	Key           string `json:"key,omitempty"`
	PreviousValue string `json:"previous_value,omitempty"`
	NewValue      string `json:"new_value,omitempty"`
}

type Metadata struct {
	nftData           *NftData
	attributeOverring AttributeOverriding
	ChangeList        []*MetadataChange `json:"change_list,omitempty"`
}

type MetadataAttribute struct {
	attributeValue *NftAttributeValue
	from           string
	index          int
}

var mapAllKey = map[string]*MetadataAttribute{}

func NewMetadata(tokenData *NftData, attributeOverring AttributeOverriding) *Metadata {
	meta := &Metadata{
		nftData:           tokenData,
		attributeOverring: attributeOverring,
		ChangeList:        []*MetadataChange{},
	}

	mapAllKey = map[string]*MetadataAttribute{}

	// Parse the metadata
	for i, attri := range tokenData.OriginAttributes {
		mapAllKey[attri.Name] = &MetadataAttribute{
			attributeValue: attri,
			from:           "origin",
			index:          i,
		}
	}
	for i, attri := range tokenData.OnchainAttributes {
		if _, ok := mapAllKey[attri.Name]; ok {
			if attributeOverring == AttributeOverriding_CHAIN {
				mapAllKey[attri.Name] = &MetadataAttribute{
					attributeValue: attri,
					from:           "chain",
					index:          i,
				}
			}
		} else {
			mapAllKey[attri.Name] = &MetadataAttribute{
				attributeValue: attri,
				from:           "chain",
				index:          i,
			}
		}

	}
	return meta
}

func (m *Metadata) GetImage() string {
	if m.nftData.OnchainImage != "" {
		return m.nftData.OnchainImage
	}
	return m.nftData.OriginImage
}

func (m *Metadata) SetImage(image string) {
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
	attri := mapAllKey[key]
	if attri == nil {
		return 0, sdkerrors.Wrap(ErrAttributeNotFoundForAction, key)
	}
	if _, ok := attri.attributeValue.GetValue().(*NftAttributeValue_NumberAttributeValue_); ok {
		// Number
		return int64(attri.attributeValue.GetNumberAttributeValue().Value), nil
	}
	return 0, sdkerrors.Wrap(ErrAttributeTypeNotMatch, attri.attributeValue.Name)
}

func (m *Metadata) SetNumber(key string, value int64) error {
	// m.mapNumber[key] = value
	attri := mapAllKey[key]
	if attri == nil {
		panic(sdkerrors.Wrap(ErrAttributeNotFoundForAction, key))
	}
	if _, ok := attri.attributeValue.GetValue().(*NftAttributeValue_NumberAttributeValue_); ok {
		// Number
		newAttributeValue := &NftAttributeValue{
			Name: attri.attributeValue.Name,
			Value: &NftAttributeValue_NumberAttributeValue_{
				NumberAttributeValue: &NftAttributeValue_NumberAttributeValue{
					Value: uint64(value),
				},
			},
		}
		if attri.from == "chain" {
			m.ChangeList = append(m.ChangeList, &MetadataChange{
				Key:           key,
				PreviousValue: strconv.FormatUint(attri.attributeValue.GetNumberAttributeValue().Value, 10),
				NewValue:      strconv.FormatUint(uint64(value), 10),
			})
			mapAllKey[key].attributeValue = newAttributeValue
			m.nftData.OnchainAttributes[attri.index] = newAttributeValue
		} else {
			return sdkerrors.Wrap(ErrAttributeOverriding, "can not override the origin attribute")
		}
	} else {
		return sdkerrors.Wrap(ErrAttributeTypeNotMatch, attri.attributeValue.Name)
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
	attri := mapAllKey[key]
	if attri == nil {
		return "", sdkerrors.Wrap(ErrAttributeNotFoundForAction, key)
	}
	if _, ok := attri.attributeValue.GetValue().(*NftAttributeValue_StringAttributeValue_); ok {
		// Number
		return attri.attributeValue.GetStringAttributeValue().Value, nil
	}
	return "", sdkerrors.Wrap(ErrAttributeTypeNotMatch, attri.attributeValue.Name)
}

func (m *Metadata) SetString(key string, value string) error {
	// m.mapString[key] = value
	attri := mapAllKey[key]
	if attri == nil {
		panic(sdkerrors.Wrap(ErrAttributeNotFoundForAction, key))
	}
	if _, ok := attri.attributeValue.GetValue().(*NftAttributeValue_StringAttributeValue_); ok {
		// Number
		newAttributeValue := &NftAttributeValue{
			Name: attri.attributeValue.Name,
			Value: &NftAttributeValue_StringAttributeValue_{
				StringAttributeValue: &NftAttributeValue_StringAttributeValue{
					Value: value,
				},
			},
		}
		if attri.from == "chain" {
			m.ChangeList = append(m.ChangeList, &MetadataChange{
				Key:           key,
				PreviousValue: attri.attributeValue.GetStringAttributeValue().Value,
				NewValue:      value,
			})
			mapAllKey[key].attributeValue = newAttributeValue
			m.nftData.OnchainAttributes[attri.index] = newAttributeValue
		} else {
			return sdkerrors.Wrap(ErrAttributeOverriding, "can not override the origin attribute")
		}
	} else {
		return sdkerrors.Wrap(ErrAttributeTypeNotMatch, attri.attributeValue.Name)
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
	attri := mapAllKey[key]
	if attri == nil {
		return 0, sdkerrors.Wrap(ErrAttributeNotFoundForAction, key)
	}
	if _, ok := attri.attributeValue.GetValue().(*NftAttributeValue_FloatAttributeValue_); ok {
		// Number
		return attri.attributeValue.GetFloatAttributeValue().Value, nil
	}
	return 0, sdkerrors.Wrap(ErrAttributeTypeNotMatch, attri.attributeValue.Name)
}

func (m *Metadata) SetFloat(key string, value float64) error {
	// m.mapFloat[key] = value
	attri := mapAllKey[key]
	if attri == nil {
		panic(sdkerrors.Wrap(ErrAttributeNotFoundForAction, key))
	}
	if _, ok := attri.attributeValue.GetValue().(*NftAttributeValue_FloatAttributeValue_); ok {
		// Number
		newAttributeValue := &NftAttributeValue{
			Name: attri.attributeValue.Name,
			Value: &NftAttributeValue_FloatAttributeValue_{
				FloatAttributeValue: &NftAttributeValue_FloatAttributeValue{
					Value: value,
				},
			},
		}
		if attri.from == "chain" {
			m.ChangeList = append(m.ChangeList, &MetadataChange{
				Key:           key,
				PreviousValue: strconv.FormatFloat(attri.attributeValue.GetFloatAttributeValue().Value, 'f', -1, 64),
				NewValue:      strconv.FormatFloat(value, 'f', -1, 64),
			})
			mapAllKey[key].attributeValue = newAttributeValue
			m.nftData.OnchainAttributes[attri.index] = newAttributeValue
		} else {
			return sdkerrors.Wrap(ErrAttributeOverriding, "can not override the origin attribute")
		}
	} else {
		return sdkerrors.Wrap(ErrAttributeTypeNotMatch, attri.attributeValue.Name)
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
	attri := mapAllKey[key]
	if attri == nil {
		return false, sdkerrors.Wrap(ErrAttributeNotFoundForAction, key)
	}
	if _, ok := attri.attributeValue.GetValue().(*NftAttributeValue_BooleanAttributeValue_); ok {
		// Number
		return attri.attributeValue.GetBooleanAttributeValue().Value, nil
	}
	return false, sdkerrors.Wrap(ErrAttributeTypeNotMatch, attri.attributeValue.Name)
}

func (m *Metadata) SetBoolean(key string, value bool) error {
	// m.mapBool[key] = value
	attri := mapAllKey[key]
	if attri == nil {
		panic(sdkerrors.Wrap(ErrAttributeNotFoundForAction, key))
	}
	if _, ok := attri.attributeValue.GetValue().(*NftAttributeValue_BooleanAttributeValue_); ok {
		// Number
		newAttributeValue := &NftAttributeValue{
			Name: attri.attributeValue.Name,
			Value: &NftAttributeValue_BooleanAttributeValue_{
				BooleanAttributeValue: &NftAttributeValue_BooleanAttributeValue{
					Value: value,
				},
			},
		}
		if attri.from == "chain" {
			m.ChangeList = append(m.ChangeList, &MetadataChange{
				Key:           key,
				PreviousValue: strconv.FormatBool(attri.attributeValue.GetBooleanAttributeValue().Value),
				NewValue:      strconv.FormatBool(value),
			})
			mapAllKey[key].attributeValue = newAttributeValue
			m.nftData.OnchainAttributes[attri.index] = newAttributeValue
		} else {
			return sdkerrors.Wrap(ErrAttributeOverriding, "can not override the origin attribute")
		}
	} else {
		return sdkerrors.Wrap(ErrAttributeTypeNotMatch, attri.attributeValue.Name)
	}
	return nil
}

func (m *Metadata) ReplaceAllString(intput string, regexpStr string, replaceStr string) string {
	reg := regexp.MustCompile(regexpStr)
	return reg.ReplaceAllString(intput, replaceStr)
}
