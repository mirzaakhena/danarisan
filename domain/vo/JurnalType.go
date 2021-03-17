package vo

import (
	"strings"

	"github.com/mirzaakhena/danarisan/application/apperror"
)

type JurnalType string

const (
	TambahModalJurnalTypeEnum  JurnalType = "TAMBAHMODAL"
	SetorTagihanJurnalTypeEnum JurnalType = "SETORTAGIHAN"
	MenangUndianJurnalTypeEnum JurnalType = "MENANGUNDIAN"
	PenyesuaianJurnalTypeEnum  JurnalType = "PENYESUAIAN"
)

var enumJurnalType = map[JurnalType]JurnalTypeDetail{
	TambahModalJurnalTypeEnum:  {},
	SetorTagihanJurnalTypeEnum: {},
	MenangUndianJurnalTypeEnum: {},
	PenyesuaianJurnalTypeEnum:  {},
}

type JurnalTypeDetail struct { //
}

func NewJurnalType(name string) (JurnalType, error) {
	name = strings.ToUpper(name)

	if _, exist := enumJurnalType[JurnalType(name)]; !exist {
		return "", apperror.UnrecognizedEnum.Var(name, "JurnalType")
	}

	return JurnalType(name), nil
}

func (r JurnalType) GetDetail() JurnalTypeDetail {
	return enumJurnalType[r]
}

func (r JurnalType) PossibleValues() []JurnalType {
	res := []JurnalType{}
	for key, _ := range enumJurnalType {
		res = append(res, key)
	}
	return res
}
