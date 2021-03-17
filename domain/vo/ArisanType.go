package vo

import (
	"strings"

	"github.com/mirzaakhena/danarisan/application/apperror"
)

type ArisanType string

const (
	SingleSlotArisanTypeEnum ArisanType = "SINGLESLOT"
	MultiSlotArisanTypeEnum  ArisanType = "MULTISLOT"
	GroupSlotArisanTypeEnum  ArisanType = "GROUPSLOT"
)

var enumArisanType = map[ArisanType]ArisanTypeDetail{
	SingleSlotArisanTypeEnum: {},
	MultiSlotArisanTypeEnum:  {},
	GroupSlotArisanTypeEnum:  {},
}

type ArisanTypeDetail struct { //
}

func NewArisanType(name string) (ArisanType, error) {
	name = strings.ToUpper(name)

	if _, exist := enumArisanType[ArisanType(name)]; !exist {
		return "", apperror.UnrecognizedEnum.Var(name, "ArisanType")
	}

	return ArisanType(name), nil
}

func (r ArisanType) GetDetail() ArisanTypeDetail {
	return enumArisanType[r]
}

func (r ArisanType) PossibleValues() []ArisanType {
	res := []ArisanType{}
	for key := range enumArisanType {
		res = append(res, key)
	}
	return res
}
