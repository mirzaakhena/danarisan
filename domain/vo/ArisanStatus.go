package vo

import (
	"strings"

	"github.com/mirzaakhena/danarisan/application/apperror"
)

type ArisanStatus string

const (
	TerimaPesertaArisanStatusEnum ArisanStatus = "TERIMAPESERTA"
	MulaiArisanStatusEnum         ArisanStatus = "MULAI"
	SelesaiArisanStatusEnum       ArisanStatus = "SELESAI"
)

var enumArisanStatus = map[ArisanStatus]ArisanStatusDetail{
	TerimaPesertaArisanStatusEnum: {},
	MulaiArisanStatusEnum:         {},
	SelesaiArisanStatusEnum:       {},
}

type ArisanStatusDetail struct { //
}

func NewArisanStatus(name string) (ArisanStatus, error) {
	name = strings.ToUpper(name)

	if _, exist := enumArisanStatus[ArisanStatus(name)]; !exist {
		return "", apperror.UnrecognizedEnum.Var(name, "ArisanStatus")
	}

	return ArisanStatus(name), nil
}

func (r ArisanStatus) GetDetail() ArisanStatusDetail {
	return enumArisanStatus[r]
}

func (r ArisanStatus) String() string {
	return string(r)
}

func (r ArisanStatus) PossibleValues() []ArisanStatus {
	res := []ArisanStatus{}
	for key, _ := range enumArisanStatus {
		res = append(res, key)
	}
	return res
}
