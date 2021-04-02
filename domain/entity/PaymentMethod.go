package entity

import (
	"strings"

	"github.com/mirzaakhena/danarisan/application/apperror"
)

type PaymentMethod string

const (
	LinkAjaPaymentMethodEnum PaymentMethod = "LINKAJA"
	DANAPaymentMethodEnum    PaymentMethod = "DANA"
	GopayPaymentMethodEnum   PaymentMethod = "GOPAY"
)

var enumPaymentMethod = map[PaymentMethod]PaymentMethodDetail{
	LinkAjaPaymentMethodEnum: {},
	DANAPaymentMethodEnum:    {},
	GopayPaymentMethodEnum:   {},
}

type PaymentMethodDetail struct { //
}

func NewPaymentMethod(name string) (PaymentMethod, error) {
	name = strings.ToUpper(name)

	if _, exist := enumPaymentMethod[PaymentMethod(name)]; !exist {
		return "", apperror.UnrecognizedEnum.Var(name, "PaymentMethod")
	}

	return PaymentMethod(name), nil
}

func (r PaymentMethod) GetDetail() PaymentMethodDetail {
	return enumPaymentMethod[r]
}

func (r PaymentMethod) PossibleValues() []PaymentMethod {
	res := []PaymentMethod{}
	for key, _ := range enumPaymentMethod {
		res = append(res, key)
	}
	return res
}
