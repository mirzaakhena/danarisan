package vo

import (
	"strings"

	"github.com/mirzaakhena/danarisan/application/apperror"
)

type UndanganState string

const (
	DitawarkanUndanganStateEnum UndanganState = "DITAWARKAN"
	TerimaUndanganStateEnum     UndanganState = "TERIMA"
	TolakUndanganStateEnum      UndanganState = "TOLAK"
)

var enumUndanganState = map[UndanganState]UndanganStateDetail{
	DitawarkanUndanganStateEnum: {},
	TerimaUndanganStateEnum:     {},
	TolakUndanganStateEnum:      {},
}

type UndanganStateDetail struct { //
}

func NewUndanganState(name string) (UndanganState, error) {
	name = strings.ToUpper(name)

	if _, exist := enumUndanganState[UndanganState(name)]; !exist {
		return "", apperror.UnrecognizedEnum.Var(name, "UndanganState")
	}

	return UndanganState(name), nil
}

func (r UndanganState) GetDetail() UndanganStateDetail {
	return enumUndanganState[r]
}

func (r UndanganState) PossibleValues() []UndanganState {
	res := []UndanganState{}
	for key, _ := range enumUndanganState {
		res = append(res, key)
	}
	return res
}
