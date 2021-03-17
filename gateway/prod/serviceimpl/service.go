package serviceimpl

import (
	"github.com/google/uuid"
	"sync"
)

var onceServiceImpl sync.Once

var serviceImplObj ServiceImplementation

type ServiceImplementation struct {
}

func SingletonServiceImplementation() *ServiceImplementation {
	onceServiceImpl.Do(func() {
		serviceImplObj = ServiceImplementation{}
	})
	return &serviceImplObj
}

func (r *ServiceImplementation) Generate() string {
	return uuid.NewString()
}
