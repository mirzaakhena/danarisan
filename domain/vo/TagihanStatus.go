package vo

import (
	"strings"

	"github.com/mirzaakhena/danarisan/application/apperror"
)

type TagihanStatus string

const (
	BelumDitagihTagihanStatusEnum       TagihanStatus = "BELUMDITAGIH"
	MenungguPembayaranTagihanStatusEnum TagihanStatus = "MENUNGGUPEMBAYARAN"
	LunasTagihanStatusEnum              TagihanStatus = "LUNAS"
	KadaluwarsaTagihanStatusEnum        TagihanStatus = "KADALUWARSA"
)

var enumTagihanStatus = map[TagihanStatus]TagihanStatusDetail{
	BelumDitagihTagihanStatusEnum:       {},
	MenungguPembayaranTagihanStatusEnum: {},
	LunasTagihanStatusEnum:              {},
	KadaluwarsaTagihanStatusEnum:        {},
}

type TagihanStatusDetail struct { //
}

func NewTagihanStatus(name string) (TagihanStatus, error) {
	name = strings.ToUpper(name)

	if _, exist := enumTagihanStatus[TagihanStatus(name)]; !exist {
		return "", apperror.UnrecognizedEnum.Var(name, "TagihanStatus")
	}

	return TagihanStatus(name), nil
}

func (r TagihanStatus) GetDetail() TagihanStatusDetail {
	return enumTagihanStatus[r]
}

func (r TagihanStatus) String() string {
	return string(r)
}

func (r TagihanStatus) PossibleValues() []TagihanStatus {
	res := []TagihanStatus{}
	for key, _ := range enumTagihanStatus {
		res = append(res, key)
	}
	return res
}
