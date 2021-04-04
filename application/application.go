package application

type RegistryContract interface {
	RegisterUsecase()
	RunApplication()
}

func Run(rv RegistryContract) {
	if rv != nil {
		rv.RegisterUsecase()
		rv.RunApplication()
	}
}

type RegistryContract2 interface {
	Register()
}

type Runner interface {
	Start()
}

func RunApp(r Runner) {
	r.Start()
}