package vo

import (
	"strings"

	"github.com/mirzaakhena/danarisan/application/apperror"
)

type SideType string

const (
	Activa  = "ACTIVA"
	Passiva = "PASSIVA"
)

type AkunType string

const (
	HartaAkunTypeEnum       AkunType = "HARTA"
	PiutangAkunTypeEnum     AkunType = "PIUTANG"
	UtangAkunTypeEnum       AkunType = "UTANG"
	ModalAkunTypeEnum       AkunType = "MODAL"
	BiayaAdminAkunTypeEnum  AkunType = "BIAYAADMIN"
	BiayaArisanAkunTypeEnum AkunType = "BIAYAARISAN"
)

var enumAkunType = map[AkunType]AkunTypeDetail{
	HartaAkunTypeEnum:       {Side: Activa},
	PiutangAkunTypeEnum:     {Side: Activa},
	UtangAkunTypeEnum:       {Side: Passiva},
	ModalAkunTypeEnum:       {Side: Passiva},
	BiayaAdminAkunTypeEnum:  {Side: Activa},
	BiayaArisanAkunTypeEnum: {Side: Activa},
}

type AkunTypeDetail struct { //
	Side SideType
}

func NewAkunType(name string) (AkunType, error) {
	name = strings.ToUpper(name)

	if _, exist := enumAkunType[AkunType(name)]; !exist {
		return "", apperror.UnrecognizedEnum.Var(name, "AkunType")
	}

	return AkunType(name), nil
}

func (r AkunType) GetDetail() AkunTypeDetail {
	return enumAkunType[r]
}

func (r AkunType) PossibleValues() []AkunType {
	res := []AkunType{}
	for key, _ := range enumAkunType {
		res = append(res, key)
	}
	return res
}
